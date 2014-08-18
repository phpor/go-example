package main

import "fmt"

func returnvalue() (x, y, z string) { // 函数的返回值可以全部命名，也可以全部不命名，但是不能只是部分命名
	return "x", "y", "z"
}
func main() {
	x, y, z := returnvalue()
	fmt.Println(x, y, z) //这里输出的字符串之间自动添加了空格
}
