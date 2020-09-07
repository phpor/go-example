package main

import (
	"fmt"
)

type parent struct {
	// 这个函数属性叫函数，而绑在该struct上的方法叫做method，是非常有区别的，该函数可以重新实现
	// 类似于静态方法，只与参数相关，和本struct实例无关（所以可以重新实现）
	// 该函数也不可能返回调用该函数的实例（进而实现链式调用）
	// 这个似乎做钩子比较好使
	// 这俨然是一个抽象类
	abstractHello func(string) // 函数默认初始化成nil，而不是一个空函数，且这里也无法实现该函数

	// 即使可以通过添加一个参数的方式来创造一个访问本实例的机会，也依然只能访问本实例的Public的属性和方法，能力有限
	// 而且没有该参数依然可以实现
	abstractWithMe func(*parent, string)
}

func NewParent() *parent {
	return &parent{ //可以通过实例化方法来初始化函数属性为空函数
		abstractHello:  func(string) {},
		abstractWithMe: func(*parent, string) {},
	}
}

func (p *parent) hello() {
	fmt.Println("parent hello")
}
func (p *parent) say() {
	p.hello()
}

type c2 parent

// 如此定义，c可以访问parent结构中定义的属性（包括函数属性），但不能访问到绑定到parent上的方法；换言之，继承了属性，没继承方法
func (c c2) say() {
	println("hello in c2")
}

type child struct {
	parent
}

func (c *child) hello() {
	c.parent.say()
}

func use_abstract(a *parent, name string) {
	//	if a.abstractHello != nil {
	a.abstractHello(name)
	//	}
}
func main() {
	//	c2 := &child{}
	//	c2.say()
	//	use_abstract(&parent{abstractHello:func(Name string) {
	//		println("hello:", Name)
	//	}}, "phpor")

	use_abstract(NewParent(), "phpor") //使用空函数
	a := NewParent()
	a.abstractHello = func(name string) { //自定义方法
		println(name)
	}
	use_abstract(a, "phpor2")
	//	c2 := c2{}
	//	c2.say()

}
