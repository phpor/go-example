package main

import "testing"

func TestGoto(t *testing.T) {
	for {
		for {
			println("使用goto跳出循环也可以，定义的标签的位置就和break很不一样了")
			goto l2
		}
	}
l2:
	println("done")
}
