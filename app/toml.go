package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
)

func main() {
	arr()
}

func arr() {
	// 作为一种先进的配置语言，toml是允许数组最后一项跟一个分隔符的
	t, err := toml.Load("[section]\narr = [ 1, 2, ]")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	arrInterface := t.Get("section")
	arr := arrInterface.(*toml.Tree).Get("arr").([]interface{})
	println(len(arr))
}
