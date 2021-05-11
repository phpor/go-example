package main

import (
	"fmt"
	"unsafe"
)

type MyStr struct {
	str [1]uint8
	len int
}

func main() {
	substr()
}

func a() {
	m := &MyStr{}
	a := "111"
	b := a
	a = "222"
	ch := make(chan int, 1)
	println(ch)
	println(a)
	println(b)
	print(m)
}

type String struct {
	str    *byte
	length int
}

func (s *String) String() string {
	return fmt.Sprintf("str: %p\nlen: %d", s.str, s.length)
}

func inspectString(s *string) {
	a := (*String)(unsafe.Pointer(s))
	fmt.Printf("%s\n", a.String())
}

// 变量a的地址比变量b的地址大16字节，一方面体现了栈是向下走的，先分配a，在分配b
// 另一方面，也体现了字符串变量占用的内存大小是16字节，一个指针和一个int
func substr() {
	a := "12345"
	b := a[2:4]
	println("a:", &a)
	inspectString(&a)
	println("b:", &b)
	inspectString(&b)
}
