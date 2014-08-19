package main

import "fmt"

type myslice struct {
	s []string
}

func (s *myslice) each(f func(string)) {
	for _, str := range s.s {
		f(str)
	}
}

func main() {
	var s = &myslice{[]string{"hello", "world"}}
	s.each(func(s string) {
		fmt.Println(s)
	})

	// 匿名函数 1
	f := func(i, j int) (int) { // f 为函数地址
		return i + j
	}

	fmt.Printf("f = %v  f(1,3) = %v\n", f, f(1, 3))

	// 匿名函数 2
	x, y := func(i, j int) (m, n int) { // x y 为函数返回值
		return j, i
	}(1, 9) // 直接创建匿名函数并执行

	fmt.Printf("x = %d   y = %d\n", x, y)
}
