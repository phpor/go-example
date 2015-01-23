package main

import "fmt"

type mystdout struct{}

func (s *mystdout) Write(p []byte) (n int, err error) {
	return fmt.Print(string(p))
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// ====================================================
type Messager interface {
	String() string
}
type SimpleMessage struct {
}

// 可以看出，无论对象还是对象指针都实现了String方法
func (this SimpleMessage) String() string {
	return "对于接口参数在运行时，传递的是 值 还是 地址，取决于运行时传递的是 值 还是 地址，定义的时候不要写成接口的指针"
}

func printMessager(m Messager, obj bool) {
	if obj {
		fmt.Printf("%p\n", &m)    // 如果传入的是对象
	} else {
		fmt.Printf("%p\n", m)    // 如果传入的是指针
	}
	fmt.Println(m.String())
}
func main() {

	fmt.Println("传递对象时，进行了值的copy，可以看到地址不同")
	obj := SimpleMessage{}
	fmt.Printf("%p\n", &obj)
	printMessager(obj, true)  // 传对象

	fmt.Println("传递指针时，没有做值的copy，可以看到地址相同")
	msg := &SimpleMessage{}
	fmt.Printf("%p\n", msg)
	printMessager(msg, false) // 传指针

}
