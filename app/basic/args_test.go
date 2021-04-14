package main

import (
	"reflect"
	"testing"
)

// []string to []interface{}
// 居然没有一个便捷的方式来转换
// https://stackoverflow.com/questions/12990338/cannot-convert-string-to-interface
// https://phpor.net/blog/post/17302
func TestArgs(t *testing.T) {
	ss := []string{"a", "b", "c"}
	func(s string, ss ...interface{}) {
		println(reflect.TypeOf(ss).String())
		println(ss...)
	}("这样是可以的", ss[0], ss[1], ss[2])

	//fmt.Println([]string{"a","b","c"}...)
}
