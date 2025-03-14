package main

import "fmt"

func main() {
	scan()
}

func scan() {
	var name string
	var is bool
	var length int
	n, err := fmt.Sscanf("a ", "a %s %t %d", &name, &is, &length)
	if err != nil {
		//panic(err)
	}
	println("scan length: ", nhtt)
	fmt.Println(name, is, length)
}


