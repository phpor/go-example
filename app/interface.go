package main

import (
	"fmt"
)

func main() {
	interfaceArg()
}

func interfaceArg2() {
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
}

type SimpleMessage struct {
	msg string
}

// 可以看出，无论对象还是对象指针都实现了String方法
// "对于接口参数在运行时，传递的是 值 还是 地址，取决于运行时传递的是 值 还是 地址，定义的时候不要写成接口的指针"
func (sm *SimpleMessage) String() string {
	return sm.msg
}

// 可以看出，无论对象还是对象指针都实现了String方法
func (sm *SimpleMessage) SetMsg(msg string)  {
	sm.msg = msg
}


func printMessager(m Messager, obj bool) {
	if obj {
		s := m.(SimpleMessage)
		fmt.Printf("%p\n", &s) // 如果传入的是对象
	} else {
		// 如果传入的是指针, 那么，函数外面打印指针是一个"指针值"，这里的%p 也仅仅是把 m 用指针的格式打印而已
		// 而且，这个 m 看起来是接口本身，但是，外面看到的是被fmt.Printf从接口的data中解析到的数据的真实值
		fmt.Printf("%p\n", m)
	}

}

func setMsg(m Messager, msg string) {
	if o, ok := m.(SimpleMessage); ok {
		(&o).SetMsg(msg)
	}
	if o, ok := m.(*SimpleMessage); ok {
		o.SetMsg(msg)
	}
}

func interfaceArg3() {
	obj := SimpleMessage{}
	setMsg(obj, "值传递") // 值传递是一个副本
	fmt.Println("值传递给接口参数时: ", obj.String())
	setMsg(&obj, "指针传递")
	fmt.Println("指针传递给接口参数时: ", obj.String())
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
