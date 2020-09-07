package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime/debug"
)

func main() {
	test333()
}

func test332() {

	f, err := os.OpenFile("/tmp/appc.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("aaaa")
		panic(err)
	}
	//n, err := io.WriteString(f, "abcd")
	n, err := f.WriteString("abcd")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", n)
}

func test333() {
	ioutil.WriteFile("/tmp/debug.log", debug.Stack(), 0666)
}
