// 参考资料： http://www.tuicool.com/articles/rYVf2e	（写的比较到位）
package main

import (
	"fmt"
	"sort"
)

func slice_is_readonly(ss []string) {
	fmt.Printf("ss in func: %p\n", ss) // slice 通过参数传递时是传地址的
}
func main() {
	s := "hello_world"
	fmt.Println(s[:5]) //字符串默认可以当slice使用
	fmt.Println(s[5:6])
	fmt.Println(s[6:])

	ss := []string{"hello", "word"}
	fmt.Printf("ss in main: %p\n", ss)
	fmt.Printf("ss[1:] in main: %p\n", ss[1:])
	slice_is_readonly(ss[1:])

	sss := []string{"你好", "世界"}
	fmt.Println(sss[1:])
	fmt.Println(sss[:1])
	fmt.Println("golang 默认没有字符集转换类库，需要第三方类库")

	// 排序的示例
	ssss := []string{"c2", "c1", "b1", "b2", "a1", "a2"}
	fmt.Printf("before sort: %p\n", ssss)
	sort.Strings(ssss)
	fmt.Printf("after sort: %p\n", ssss)
	fmt.Println(ssss)
	fmt.Println(ssss[1])
	ssss[1] = "A2"
	fmt.Printf("%T :%p :%s\n", ssss, ssss, ssss) //slice 是可以被修改的

	println("slice or array:")
	arr := [2][2]int{{11, 12}, {21, 22}} // 课本上说，这个是数组，不是切片，但是很多时候用起来和切片一样一样的
	arr2 := [...]int{1, 2, 3}            // 这是个数组
	s6 := []int{1, 2, 3}                 // 这是个切片
	//var s7 []int = [...]int{1,2,3}		  // 这个赋值是错误的，类型不同，生命的是切片，赋值的是数组
	fmt.Printf("%T : %v :%v\n", arr, arr, arr[1])
	fmt.Printf("%T : %v :%v\n", arr2, arr2, arr2[1])
	fmt.Printf("%T : %v :%v\n", s6, s6, s6[1])

	println("append to slice:")
	//slice 的追加
	s1 := []string{"a", "b"}
	s2 := []string{"c", "d"}
	fmt.Printf("&s1: %p\n", s1)

	// 一般会这么写：
	s1 = append(s1, s2...)
	fmt.Printf("%p: %v \tcap(s1):%d\n", s1, s1, cap(s1)) // s1的地址可发生变化，也可能不变化，主要取决于s1的剩余容量够不够放下s2

	s3 := append(s1, s2...) // 注意：追加的是slice中的元素，而不是slice，所以这里用 ... 语法，将slice结构成多个元素参数
	fmt.Printf("%p: %v\n%p: %v\n", s1, s1, s3, s3)
}
