package main

import (
	"flag"
	"net"
	"os"
	"os/exec"
	"strings"
)

func main() {
	addr := flag.String("addr", ":1234", "[host]:port")
	cmdline := flag.String("cmd", "/bin/ls", "command")
	flag.Parse()
	command := strings.SplitN(*cmdline, ",,", -1)
	cmdName := command[0]
	args := command[1:]

	l, err := net.Listen("tcp", *addr)
	if err != nil {
		println(err.Error())
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			break
		}
		cmd := exec.Command(cmdName, args...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = c
		cmd.Stdin = c // 这里如果给c的话，cmd.Run会等到client端把c关掉才会返回，如果给nil就不必client主动关掉了
		go func() {
			if err := cmd.Run(); err != nil {
				println(err.Error())
			}
			_ = c.Close()
		}()
	}
}
