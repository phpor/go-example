// 参考资料： http://www.alaiblog.com/golang/step-by-step-learning-golang-go-language-basics.html

package main

import "fmt"

func main() {
	typeAssert()
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
// 注意： 可以将一个interface类型的变量x来断言是否某种非interface类型的值，不能反过来检查非interface类型的值是否实现了某个接口
// 参考资料： http://blog.csdn.net/kjfcpua/article/details/18667255
func typeAssert() {
	var str interface{} = Str("abc")
	s, _ := str.(Str)
	fmt.Printf("%T, %+v", s, s)
}

// 这里有一个返回值的隐式类型转换
func yinshiTypeConvert() <-chan int {
	ch := make(chan int, 5)
	ch <- 1
	return ch   // 等同于 return <-chan int(ch)
}

