package main

import (
	"fmt"
)

func main() {
	c := &CC{}
	c.f = func(ib IA) {
		ib.Say()
	}
}

type IA interface {
	Say()
}

type IB interface {
	Say()
}

type CC struct {
	f func(IA)
}

func interfaceEqual() {
	var o1 interface{}
	var o2 interface{}
	o1 = 1
	o2 = 1
	fmt.Printf("o1==o2: %t\n", o1 == o2)

	type c struct {
		name string
	}
	o1 = c{"phpor"}
	o2 = c{"phpor"}
	fmt.Printf("o1==o2: %t\n", o1 == o2)

	o1 = struct {
		name string
	}{"phpor"}

	o2 = struct {
		name string
	}{"phpor"}
	fmt.Printf("o1==o2: %t\n", o1 == o2)

	a := "phpor"
	b := "phpor"
	o1 = &a
	o2 = &b
	fmt.Printf("o1==o2: %t\n", o1 == o2) // false

	_ = func() { // 这种方式可以注释一段代码
		o1 = struct {
			name []byte
		}{[]byte("phpor")}

		o2 = struct {
			name []byte
		}{[]byte("phpor")}
		fmt.Printf("o1==o2: %t\n", o1 == o2) // 这个比较没有语法错误，但是暗藏panic，因为[]byte 不能直接应用于 ==
	}
}

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
		fmt.Printf("%p\n", &m) // 如果传入的是对象
	} else {
		fmt.Printf("%p\n", m) // 如果传入的是指针
	}
	fmt.Println(m.String())
}

func interfaceArg() {

	fmt.Println("传递对象时，进行了值的copy，可以看到地址不同")
	obj := SimpleMessage{}
	fmt.Printf("%p\n", &obj)
	printMessager(obj, true) // 传对象

	fmt.Println("传递指针时，没有做值的copy，可以看到地址相同")
	msg := &SimpleMessage{}
	fmt.Printf("%p\n", msg)
	printMessager(msg, false) // 传指针

}

// 不能用属性函数当做方法
//func interfaceFunc() {
//	w := struct {
//		Write func(p []byte) (n int, err error)
//	}{
//		Write: func(p []byte) (n int, err error) {
//			println(string(p))
//			return len(p), nil
//		},
//	}
//	fmt.Fprintf(w, "%s", "hello")
//}
