package main

func TestCopyStack() {
	var x [10]int
	//var x1 = x
	//fmt.Printf("&x1: %p\n", &x1)
	//fmt.Printf("&x: %p\n", &x)
	f1(x)
	println(&x)
	a(x)
	//fmt.Printf("&x1: %p\n", &x1)
	//fmt.Printf("&x: %p\n", &x)
	println(&x)
}

func f1(x interface{}) {
	println("func f1", x)
}

//go:noinline
func a(x [10]int) {
	println(`func a`)
	var y [100]int
	b(y)
}

//go:noinline
func b(x [100]int) {
	println(`func b`)
	var y [1000]int
	c(y)
}

//go:noinline
func c(x [1000]int) {
	println(`func c`)
}
