package main

import (
	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "Hello, \x90\xA2\x8A\x45" // CP932 encoded version of "Hello, 世界" ， 这里的 s 是string类型，说明string没有字符集的概念

	r, _ := charset.NewReader("CP932", strings.NewReader(s)) // convert from CP932 to UTF-8
	s2_, _ := ioutil.ReadAll(r)
	s2 := string(s2_)
	fmt.Println(s2)                         // => Hello, 世界
	fmt.Println(len(s2))                    // => 13
	fmt.Println(utf8.RuneCountInString(s2)) // => 9
	fmt.Println(utf8.ValidString(s2))       // => true
	fmt.Println(utf8.ValidString(s))        // => false
	fmt.Printf("%T|%#v\n", s, s)            // 注意 %v 与 %#v 的区别

	ss := "This is not utf-8 string \xa1"
	fmt.Println(utf8.ValidString(ss)) // => false

	pice := []int32{20, 30, 40, 90}
	sss := string(pice)                                             // string 似乎执行了内存拷贝，但是不会涉及到字符集的处理（转换或校验）
	fmt.Printf("%T:%p %T:%p:%d\n", pice, pice, sss, &sss, len(sss)) // 为什么打印字符串变量的地址还需要取地址符

	tr, err := charset.TranslatorTo("windows-1252") //需要检查字符集列表
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, gbk, err2 := tr.Translate([]byte("utf-8汉字"), true)
	if err2 != nil {
		fmt.Println(err2)
		os.Exit(1)
	}
	fmt.Println(gbk)
}
