package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

func main() {
	// 定义一个字符串
	message := "Hello, from Go to C!\n"

	// 将 Go 字符串转换为 C 字符串
	// 为了避免golang gc后C依然在使用，所以go字符串转成C字符串会发生内存copy，所以，用完后需要使用C的方式来free
	cMessage := C.CString(message)
	defer C.free(unsafe.Pointer(cMessage)) // 释放 C 字符串

	// 调用 C 中的 printf 函数
	C.puts(cMessage) // 因为C里面的printf是宏定义，所以不能直接使用
}
