package main

import "runtime"

func main() {
	f1()
}

func f1() {
	f2(0)
}

func f2(skip int) {
	var pc [1]uintptr
	n := runtime.Callers(skip+2, pc[:])
	if n == 0 {
		println("nil")
	}
	frames := runtime.CallersFrames(pc[:n])
	f, more := frames.Next()
	println(f.Function)
	println(more)
}
