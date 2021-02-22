package main

import (
	"flag"
	"fmt"
)

func main() {
	s := flag.String("name", "", "")
	err := flag.CommandLine.Parse([]string{"aaa", "-name", "bbb"})
	if err != nil {
		println(err.Error())
		return
	}
	//flag.Parse()
	fmt.Printf("%v\n", flag.Args())
	println(*s)
}
