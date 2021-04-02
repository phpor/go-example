package main

func main() {
	T(newFunc(func() {
		println("A")
	}, func() {
		println("B")
	}))
}

func T(ifunc Ifunc) {
	ifunc.A()
	ifunc.B()
}

type Ifunc interface {
	A()
	B()
}

type Impl struct {
	A1 func()
	B1 func()
}

func (impl *Impl) A() {
	impl.A1()
}

func (impl *Impl) B() {
	impl.B1()
}

func newFunc(A func(), B func()) Ifunc {
	return &Impl{
		A1: A,
		B1: B,
	}
}

// 下面的逻辑说明一个问题，就是，虽然接口签名都一样，但是接口的类型不一样就不能视为相同的接口
func F1(i I1) {

}

func F2(i I2) {

}

func F3(f func(I2)) {
	f(nil)
}

func F4() {
	F3(F2)
}

type I1 interface {
	Print()
}
type I2 interface {
	Print()
}
