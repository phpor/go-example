package main

import (
	"fmt"
	"os"
)

var cmds = map[string]func() error{}

var defaultCmd = "pool"

func main() {
	cmd := defaultCmd
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}
	if f, ok := cmds[cmd]; ok {
		if err := f(); err != nil {
			fmt.Print(err.Error())
		}
	} else {
		fmt.Printf("%s not exists", cmd)
	}
	fmt.Println()
}
