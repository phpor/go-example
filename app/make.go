// 需要开辟内存的变量都需要make（或者直接初始化），如： channel、slice ; 数组不需要make
package main

import (
	"fmt"
)

func main() {
	var b []byte = make([]byte, 4) // 如果不make就没有被初始化，没有被初始化就不能使用，虽没有语法错误，但会有运行时错误
	b[1] = 1
	fmt.Println(b)

	var a [4]byte // 数组不需要make就可以直接用哦~，只是这里的4不可能是变量的哦~
	a[1] = 1
	fmt.Println(a)

	//	i := 4
	//	var c2 [i]byte	//声数组是不允许使用变量的哦~
	//	c2[1] = 1
	//	fmt.Println(c2)
}
