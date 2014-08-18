package main

import "fmt"

/** 笔记
1. 如下可变参数函数，虽然x、y类型相同，也不能省略掉 x 参数的类型，即，不能写成： func variableargs(x, y ...string) {
*/
func variableargs(x string, y ...string) { //可变参数实质上是个slice，只是语法上以可变参数的方式表达了
	fmt.Printf("slice in sub func: %p\n", y)
	fmt.Println(x)
	for k, v := range y {
		fmt.Printf("%d => %s\n", k, v)
	}
}
func main() {
	variableargs("first:", "a", "b", "c")
	variableargs("second:", "x", "y")
	s := []string{"hello", "world"}
	fmt.Printf("slice in main: %p\n", s)
	variableargs("slice 解构成多个参数已然传址：", s...)
}
