// 参考： http://blog.golang.org/go-maps-in-action
// 1.
package main

import "fmt"

func main() {
	ref()
}
func ref() {
	m1 := map[string]map[string]int{"a1":{"b1":1}, "a2":{"b2":1}}
	fmt.Printf("%p\n", m1)
	m2 := m1["a1"]
	fmt.Printf("%p\n", m1["a1"]["b1"])    // 整型值没有地址
	fmt.Printf("%p\n", m1["a1"])    // 整型值没有地址
	fmt.Printf("%p\n", m2)
	m2["b1"] = 3
	fmt.Printf("%p\n", m1["a1"]["b1"])    // 整型值没有地址
}

func basic() {
	m1 := map[string]int{}  // 等价于
	m2 := make(map[string]int)
	var m3  map[string]int     // 这是一个未初始化的map，不能直接用哦

	m1["k"] = 1
	m2["k"] = 2
	//m3["k"] = 3	// 这里虽没有语法错误，但是，会有运行时错误的
	m3 = m1
	fmt.Println(m3)
	i2 := 2
	i1 := m1["k"]
	//	i2,_ = m2["j"]        // key 不存在是， i2被初始化为默认值 0
	if v, exists := m2["j"]; exists {
		i2 = v
	}

	fmt.Println("i1,i2: ", i1, i2)

	if i3, exists := m1["k"]; exists {    // i3 , exists 的作用于仅仅是该if语句
		fmt.Println("m1[k]: ", i3)
	}

	if i3, exists := m2["j"]; exists {
		fmt.Println("m2[j]: ", i3)
	}

	// 可以有更复杂的map，如： map[string]map[string]int
}
