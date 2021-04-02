package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

// golang 的协程的stack 不再系统分配的栈空间中

func main() {
	A()
	B()
	C()
	D()
	E()
	F()
	addr()
}

func A() {
	a := [1024]byte{'a'} // 栈
	fmt.Printf("%s: %p\n", FuncName(), &a)
	return
}

func B() [1024]byte {
	a := [1024]byte{'a'}                   // 栈，看似逃逸，实际在返回时复制了一份出去的
	fmt.Printf("%s: %p\n", FuncName(), &a) // 取地址取的是变量存储的内容的地址，所以 &a 和 &(a[0])是一样的
	fmt.Printf("%s: %p\n", FuncName(), &(a[0]))
	return a
}

func C() [1024 * 1024 * 20]byte {
	a := [1024 * 1024 * 20]byte{'a'} // 堆，因为太大
	fmt.Printf("%s: %p\n", FuncName(), &a)
	return a
}

func D() []byte {
	a := []byte{'a'}
	fmt.Printf("%s: %p\n", FuncName(), &a)
	return a
}

func E() {
	a := make([]byte, 1024)
	fmt.Printf("%s: %p\n", FuncName(), a)
	return
}

func F() {
	a := make([]byte, 1024*1024*100)
	fmt.Printf("%s: %p\n", FuncName(), a)
	return

}

func addr() {
	debug.PrintStack()
	time.Sleep(100 * time.Second)
}
func FuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
