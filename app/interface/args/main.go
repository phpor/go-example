package main

import "fmt"

func echo(i interface{}) {
	fmt.Println("in echo", i)
	println(i)
	i = "ssss" // i 是一个口袋，进来的时候 i.data 是个字符串指针（其实就是一个地址值而已），现在i.data 是一个字符串结构体，和调用者没有任何关系
}

func main() {
	s := "abcd"
	fmt.Println("in main", &s)
	echo(&s)
	fmt.Println(s)
}
