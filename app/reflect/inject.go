package main

import (
	"fmt"
	"reflect"
)

type Ctx interface {
	Printf(format string, a ...interface{})
}

type CliCtx struct {
}

func (cc *CliCtx) Printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

type impl struct {
}

func (impl *impl) ActionA(ctx Ctx) error {
	ctx.Printf("%s", "haha")
	return nil
}

func testInject() {
	o := &impl{}
	methodName := "ActionA"
	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)
	_, ok := t.MethodByName(methodName)
	if !ok {
		panic("method not exists")
	}
	m := v.MethodByName(methodName)
	i := m.Type().NumIn()
	if i == 0 {
		println("only one arg")
		return
	}
	aType := m.Type().In(0)
	// 该类型是否实现了某个接口
	if ok := aType.Implements(reflect.TypeOf((*Ctx)(nil)).Elem()); ok {
		println("可以用 reflect.Implements 的方式来检测")
	}

	// 该类型能否转换为另一种类型
	if ok := aType.ConvertibleTo(reflect.TypeOf((*Ctx)(nil))); ok {
		println("可以用 reflect.ConvertibleTo 的方式来检测")
	}

	// 该类型能否赋值给另一种类型
	if ok := aType.AssignableTo(reflect.TypeOf((*Ctx)(nil))); ok {
		println("可以用reflect.AssignableTo 的方式来检测")
	}

	m.Call([]reflect.Value{reflect.ValueOf(&CliCtx{})})

}
