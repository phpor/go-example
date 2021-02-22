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
	cmdline := flag.String("cmd", "", "command")
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
		cmd.Stdin = c
		go func() {
			if err := cmd.Run(); err != nil {
				println(err.Error())
			}
			_ = c.Close()
		}()
	}
}
