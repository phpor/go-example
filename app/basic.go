// 参考资料：
// 1. http://www.alaiblog.com/golang/step-by-step-learning-golang-go-language-basics.html
// 2. https://github.com/astaxie/build-web-application-with-golang/blob/master/ebook/02.6.md

package main

import (
	"fmt"
	"time"
	"strings"
	"net"
)

const (
	a = iota + 1
	b
	c
	d
)

func main() {
	assert()
	//printMany()
}

func assert() {
	var i interface{} = []string{"a", "b", "c"}
	fmt.Printf("%T %v\n", i, i)
	s := i.([]string)
	fmt.Println(s)
}
func returnValOrRef() {

	fmt.Println("=========== string ==========")
	s := returnVal()
	fmt.Println(&s)

	ss := returnRef()
	fmt.Println(ss)

	fmt.Println("=========== struct ==========")
	idc := returnIdcVal()
	fmt.Printf("%p\n", &idc)

	idc2 := returnIdcRef()
	fmt.Printf("%p\n", idc2)
}

type idc struct {
	name string
	code uint8
}

func returnIdcVal() idc {
	idc := idc{
		name: "yf",
		code: 1,
	}
	fmt.Printf("%p\n", &idc)
	return idc
}
func returnIdcRef() *idc {
	idc := &idc{
		name: "yf",
		code: 1,
	}
	fmt.Printf("%p\n", idc)
	return idc
}
func returnVal() string {
	s := "string"
	fmt.Println(&s)
	return s
}

func returnRef() *string {
	s := "string"
	fmt.Println(&s)
	return &s
}
func parseIp(ip string) {
	if ipv4 := net.ParseIP(ip).To4(); ipv4 != nil {
		fmt.Println([]byte(ipv4))
	}
}

// 交换两个变量的值，就是这么的简单
func swap() {
	x, y := 1, 2
	x, y = y, x
	fmt.Println(x, y)
}

// 类型转换
func typeConvert() {
	str := "abc"
	//fmt.Printf("%T, %+v", str.([]byte),  str.([]byte)) // 这样是不对的
	fmt.Printf("%T, %+v", []byte(str), []byte(str))
}

type Str string

// 类型断言
// 注意： 可以将一个interface类型的变量x来断言是否某种非interface类型(或interface类型)的值，不能反过来检查非interface类型的值是否实现了某个接口
// 参考资料： http://blog.csdn.net/kjfcpua/article/details/18667255
func typeAssert() {
	var str interface{} = Str("abc")
	s, _ := str.(Str)        // 注意： 圆括号中是一个“类型”，而不是一个变量
	fmt.Printf("%T, %+v", s, s)
}

// 关于类型（转换）、接口参数的用法
func callfunc() {
	var i int64 = 12345
	needInt8(int8(i))    // 这里需要类型转换，而且，可能丢失信息
	needInterface(i)    // 这里不需要类型转换，而且，不会丢失信息
}
func needInt8(i int8) {
	fmt.Println(i)
}
func needInterface(i interface{}) {
	// 这里想知道i是什么，需要用到类型断言（或反射）
	fmt.Printf("%d\n", i.(int64)+1)    // 这里需要显示地把 i 转换成原本的类型
}

// 这里有一个返回值的隐式类型转换
func yinshiTypeConvert() <-chan int {
	ch := make(chan int, 5)
	ch <- 1
	return ch   // 等同于 return <-chan int(ch)
}

func printMany() {
	a := 1
	b := "abc"
	c := time.Now()
	fmt.Println(a, b, c)
	s := []interface{}{a, b, c}
	str := make([]string, len(s))
	for k, v := range s {
		str[k] = fmt.Sprint(v)
	}
	fmt.Println(strings.Join(str[:], "\t"))
}

