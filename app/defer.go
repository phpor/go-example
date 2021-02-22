package main

import (
	"fmt"
	"time"
)

func main() {
	defer TimeMeasure()()
	time.Sleep(time.Second)
}

func TimeMeasure() func() {
	now := time.Now()
	return func() {
		fmt.Printf("%s\n", time.Now().Sub(now))
	}
}
func test11() int {
	i := 1
	defer func() {
		i = 2 // 这个影响不到返回值，因为返回值已经被copy出去了
	}()
	return i
}
func deferTest11() int {
	i := 1
	defer func() int {
		i = 2    // 这个影响不到返回值，因为返回值已经被copy出去了
		return i // 这里的return和父函数中的return已经没有任何关系了
	}()
	return i
}
func deferTest12() *testDefer {
	s := &testDefer{i: 1}
	//如果在return之前中断，则返回值为nil
	return s
}
func deferTest13() (s *testDefer) {
	defer func() {
		if err := recover(); err != nil {
			//s.i = 3 // 这里可以给 s 赋予一个正常的值，
			fmt.Printf("recovered: %v\n", err)
		}
	}()
	s = &testDefer{i: 1}
	//如果在return之前中断，则返回值也至少不会是nil
	panic("111")
	s.i = 2
	return s
}
func test14() {
	i := test11()
	fmt.Printf("%d\n", i)
	s := test12()
	fmt.Printf("%d\n", s.i)
	fmt.Printf("%d\n", deferTest13().i)
}

type testDefer struct {
	i int
}

func test12() *testDefer {
	s := &testDefer{i: 1}

	defer func() {
		s.i = 2 // 这里会影响到返回值，因为返回的是s指针，这里修改的是指针指向的内容
	}()
	return s
}
