package main

import (
	"fmt"
	"sort"
)

func slice_is_readonly(ss []string) {
	fmt.Printf("ss in func: %p\n", ss) // slice 通过参数传递时是传地址的
}
func main(){
	s := "hello_world"
	fmt.Println(s[:5])	//字符串默认可以当slice使用
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
	fmt.Printf("%T :%p :%s\n", ssss, ssss, ssss)    //slice 是可以被修改的


	arr := [2][2]int{{11, 12}, {21, 22}}        // 课本上说，这个是数组，不是切片，但是很多时候用起来和切片一样一样的
	fmt.Printf("%T : %v :%v\n", arr, arr, arr[1])
}
