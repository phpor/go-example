// 参考资料：
// 1. http://www.alaiblog.com/golang/step-by-step-learning-golang-go-language-basics.html
// 2. https://github.com/astaxie/build-web-application-with-golang/blob/master/ebook/02.6.md

package main

import (
	"errors"
	"fmt"
	"net"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const (
	a = iota + 1
	b
	c
	d
)

func main() {
	p := PointerTest{}
	p.say() // 可以用值直接call指针的方法

	//PointerTest{}.say()   // 这个和上面的写法差不多，但是，这样写不行

	// 不能用指针直接call值的方法
	(&p).sayByValue() // 为什么这个也能调用？？？

}

type PointerTest struct{}

func (p *PointerTest) say() {
	println("haha")
}

func (p PointerTest) sayByValue() {
	println("haha")
}

func testVersion() {
	var osVersionExp = regexp.MustCompile("^[0-9][0-9.]+$")
	if osVersionExp.MatchString("10.1") {
		println("succ")
	}
}

func testString() {
	println(string([]byte{64}))
	b := []byte{1, 2, 0, '@', 5, 0, 7}
	arr := strings.Split(string(b), "@")
	println(len(arr[1]))
	c := arr[1]
	decodeC([]byte(c))
}
func decodeC(c []byte) {
	println(len(c))
}

func glob() {
	files, err := filepath.Glob("~/mfp.db")
	if err != nil {
		println(err.Error())
		return
	}
	if len(files) <= 0 {
		println("no such file")
		return
	}
	println(files[0])
	return
}
func stringConv() {
	for _, v := range "我" {
		println(v)
	}
	println(string(25105))
}

func basic101() {
	y0, m0, d0 := time.Unix(0, 0).Date()
	fmt.Printf("%v, %v, %v\n", y0, m0, d0)
	y, m, d1 := time.Now().Date()
	fmt.Printf("%d, %d, %d\n", y, m, d1)
	updateTime := 20200709
	y2 := updateTime / 10000
	m2 := updateTime / 100 % 100
	d2 := updateTime % 100
	fmt.Printf("%d, %d, %d\n", y2, m2, d2)

	d := time.Date(y2, time.Month(m2), d2, 0, 0, 0, 0, time.FixedZone("CST", int((8*time.Hour).Seconds())))

	c := 20200904 / 10000
	fmt.Printf("%d, %d\n", c, d.Unix())

	fmt.Printf("%s\n", time.Now().Unix())
	s := make([]string, 3)

	fmt.Printf("%s", s)
	//fmt.Printf("%v",returnErr())
	//fmt.Printf("%v\n", 11)
	//useIp()
	//printMany()
}

func basicChar() {
	//fmt.Printf("%s %d\n", []byte{0xff}, len([]byte{0xff}))
	//fmt.Printf("%c", 0xff)
	a := 0xff
	b := string(a)
	print(b)
	//if string(uint8(255)) == string(0xff) {
	//	println("OK")
	//}
	//if string(uint8(255)) == string(-2) {
	//	println("OK")
	//}
}

func returnErr() (err error) {
	a := 1

	if _, err := 2, errors.New("bad"); err != nil {
		println(err)
	}
	println(a)
	return
}

func useIp() {
	fmt.Println([]byte(net.ParseIP("1.2.3.4")))
}
func useError() {
	a := errors.New("")
	fmt.Print(a)
}
func assert() {
	var i interface{} = []string{"a", "b", "c2"}
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
	s, _ := str.(Str) // 注意： 圆括号中是一个“类型”，而不是一个变量
	fmt.Printf("%T, %+v", s, s)
}

// 关于类型（转换）、接口参数的用法
func callfunc() {
	var i int64 = 12345
	needInt8(int8(i)) // 这里需要类型转换，而且，可能丢失信息
	needInterface(i)  // 这里不需要类型转换，而且，不会丢失信息
}
func needInt8(i int8) {
	fmt.Println(i)
}
func needInterface(i interface{}) {
	// 这里想知道i是什么，需要用到类型断言（或反射）
	fmt.Printf("%d\n", i.(int64)+1) // 这里需要显示地把 i 转换成原本的类型
}

// 这里有一个返回值的隐式类型转换
func yinshiTypeConvert() <-chan int {
	ch := make(chan int, 5)
	ch <- 1
	return ch // 等同于 return <-chan int(ch)
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
