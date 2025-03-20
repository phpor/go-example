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
		forkAndExec(fds)
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
	wg.Wait()
}

func forkAndExec(fds [2]int) {
	cmd := exec.Command("/bin/cat")
	cmd.Stdin = os.NewFile(uintptr(fds[0]), "stdin")
	cmd.Stdout = os.Stdout
	// cmd.Stdout = os.NewFile(uintptr(fds[1]), "stdout")
	// 让 cmd 不继承文件描述符

	// cmd.ExtraFiles = []*os.File{os.NewFile(uintptr(fds[1]), "extra")}
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting child process:", err)
		os.Exit(1)
	}
	cmd.Wait()
}
