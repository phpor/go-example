package main

func haha() int16 {
	println("haha")
	return 1
}

func init() {
	println("init")
}

var a1 = haha()

// 注意：这里是可以执行的，而且先于init执行，而且和与init的书写的先后顺序无关

func main() {
	print(a1)
}
