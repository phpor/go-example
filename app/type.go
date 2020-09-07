package main

import "fmt"

type flag uint8

const (
	flag1 flag = 0
	flag2 flag = 1
)

func main() {
	alias()
}
func testType1() {
	println(flag1 == flag2) // false
	println(flag(0) == flag(0))
	fmt.Printf("%v\n", flag2)
}
func (f *flag) String() string {
	switch *f {
	case flag1:
		return "[0]"
	case flag2:
		return "[1]"
	}
	return "unknown"
}

func alias() {
	type a struct {
		name string
	}
	type b a
	a1 := a{name: "a"}
	var b1 b
	b1 = b(a1)
	fmt.Printf("%v\n", b1)
}
