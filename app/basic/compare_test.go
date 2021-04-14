package main

import "testing"

func TestCompareArr(t *testing.T) {
	a := [1]byte{'a'}
	b := [1]byte{'a'}
	println(a == b) // 数组是可以直接比较的
}

// 基本类型如： 整型、bool、string、等都是可以直接比较的

// slice ,map, chan不能直接比较
