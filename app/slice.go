package main

import (
	"fmt"
)

func slice_is_readonly(ss []string) {
	fmt.Printf("ss in func: %p\n", ss)
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

}
