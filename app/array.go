package main

func main() {
	a := [3]int{}
	b := a
	c := a[:] // 数组可以几乎零成本的转成切片（不触发copy）

	b[0] = 1 // 这个不影响a，说明，数组的赋值是copy的
	println(a[0])

	c[0] = 1 // 这个的赋值影响a，说明，数组转换成切片本身并不产生值的copy
	println(a[0])

}
