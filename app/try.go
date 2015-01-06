package main

import (
	"fmt"
)

func main() {
	tryCatch(func() {
			panic("exception")
		}, func(e interface{}) {
			fmt.Println(e)
		})
	err := try(func() {
		panic("exception")
	})
	if err != nil { // use try for return quickly
		fmt.Println(err)
		return
	}
}

func tryCatch(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}
func try(fun func()) (err interface{}) {
	defer func() {
		err = recover()    //注意，recover返回的是任意值，不是error，实现了空的interface，不一定能实现error这个interface
	}()
	fun()
	return
}

