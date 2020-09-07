package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []int{2, 1, 4, 3}
	b := a // 这是a和b的指向是一样的，a、b两个变量是copy on write的关系； 如果没有append，则 b 的排序影响a，因为排序修改的是变量指向的内容
	fmt.Printf("%v\n%v\n\n", a, b)
	//b = append(b, 5)     // 如果这样append，则b的内容和a不再有关系，b的排序也不再会影响a
	_ = append(b, 5) // 如果这样append，则b和a的指向依然是同一个；b的排序会影响a； 但是，这样的不安全的，因为append后可能会因为需要扩容而地址发生变化
	sort.Ints(b)
	fmt.Printf("%v\n%v\n\n", a, b)

}
