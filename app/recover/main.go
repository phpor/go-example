package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func main() {

	if err := worker(); err != nil {
		fmt.Printf("main: %s", err.Error())
		return
	}
	fmt.Printf("main: %s", "OK")

}

func worker() (err error) {
	defer println("hahah")
	defer catch(&err)

	panic("aaa")
	arr := []string{"aa"}
	println(arr[1])
	return

}

func catch(err *error) {
	debug.Stack()
	if e := recover(); e != nil {
		pos := getCallerFunction(0)
		if e1, ok := e.(error); ok {
			*err = fmt.Errorf("panic:%s:%s", pos, e1.Error())
		} else {
			*err = fmt.Errorf("panic:%s:%s", pos, e)
		}
	}
	if *err != nil {
		println((*err).Error())
	}
	return
}

func getCallerFunction(skip int) string {
	var pc [15]uintptr
	n := runtime.Callers(skip+0, pc[:])
	if n == 0 {
		return ""
	}
	frames := runtime.CallersFrames(pc[:n])
	for {
		println(11)
		f, more := frames.Next()
		println(f.File, f.Line, f.Func.Name())
		if !more {
			break
		}
		//if f.Function == "runtime.gopanic" {
		//	frames.Next()
		//	f, _ := frames.Next()
		//	println(f.File, f.Line, f.Func.Name())
		//
		//}
	}
	return ""
}
