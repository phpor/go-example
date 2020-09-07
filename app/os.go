package main

import (
	"fmt"
	"os"
)

func main() {

	info, err := os.Stat("/tmp/aaaaa")
	if err != nil {
		panic(err)
	}
	info.Mode().IsDir()

	os.Setenv("PWD", ".")
	pwd, _ := os.Getwd()
	fmt.Printf("%v\n", pwd)
	os.Chdir("/Users")
	pwd, _ = os.Getwd()
	fmt.Printf("%v", pwd)
}
