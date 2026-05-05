package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"sync"

	"golang.org/x/sys/unix"
)

func main() {
	fds, err := unix.Socketpair(unix.AF_UNIX, unix.SOCK_STREAM, 0)
	if err != nil {
		fmt.Println("Error creating socket pair:", err)
		os.Exit(1)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		forkAndExec(fds[1])
		wg.Done()
	}()
	in := bufio.NewReader(os.Stdin)
	for {
		// 读取标准输入os.Stdin , 如果不是quit就写入 fds[0]，否则就break
		if s, err := in.ReadString('\n'); err != nil {
			fmt.Println("Error reading from stdin:", err)
			os.Exit(1)
		} else if s == "quit\n" {
			break
		} else {
			_, err := unix.Write(fds[0], []byte(s))
			if err != nil {
				fmt.Println("Error writing to socket:", err)
				os.Exit(1)
			}
		}
	}
	_ = unix.Close(fds[0]) // 父进程关闭写入端
	wg.Wait()
	_ = unix.Close(fds[1]) // 子进程关闭读取端
}

func forkAndExec(fd int) {
	cmd := exec.Command("/bin/cat")
	cmd.Stdin = os.NewFile(uintptr(fd), "stdin") // 子进程从fds[1]读取
	cmd.Stdout = os.Stdout

	// 设置子进程只继承需要的文件描述符
	cmd.ExtraFiles = []*os.File{os.NewFile(uintptr(fd), "socket")}

	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting child process:", err)
		os.Exit(1)
	}
	cmd.Wait()
}
