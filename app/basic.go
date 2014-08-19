package main

import "fmt"

func main() {
	swap()
}

// 交换两个变量的值，就是这么的简单
func swap() {
	x, y := 1, 2
	x, y = y, x
	fmt.Println(x, y)
}

