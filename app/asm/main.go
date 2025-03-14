package main

import "fmt"

type Person struct {
	Base BaseInfo
}

type BaseInfo struct {
	name string
	age int
}
var p = Person{}
func main() {
	f1()
	f2()
	f3()
}
func f1() {
	p.Base = BaseInfo{}
	fmt.Printf("%#v", p)
}
func f2() {
	b := BaseInfo{}
	p.Base = b
	fmt.Printf("%#v", p)
}

func f3() {
	b := new(BaseInfo)
	p.Base = *b
	fmt.Printf("%#v", p)
}