// 参考资料： http://www.alaiblog.com/golang/step-by-step-learning-golang-go-language-basics.html

package main

import "fmt"

func main() {
	swap()
}

// 交换两个变量的值，就是这么的简单
func swap() {
	x, y := 1, 2
	x, y = y, x
	fmt.Println(x, y)
}

