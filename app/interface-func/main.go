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
