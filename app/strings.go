package main

import (
	"fmt"
	"strings"
)

func main() {

	byte2string()
}

func bytesAndArr() {
	slice := []byte{0x31}
	//arr := [...]byte{0x31}
	println(string(slice))
	//println(string(arr))  // array 是不能向slice那样直接强制类型转换为string的
}

func byte2string() {
	b := []byte{0, 0, 0, 1, 0} // []byte 和 string 是可以自由转换的，但是每次转换都产生一次copy，
	// 重要的是，slice到slice的赋值以及传参都是传地址的，字符串是传值的；所以，尽量使用 []byte
	println(len(b))
	println(len(string(b)))

	str := "hello 中国"
	for i, ch := range str { // 这样不是按照字节遍历，而是按照unicode字符遍历，输出的是unicode码，而不是字符本身
		println(i, ch)
	}
	s := "123"
	b = []byte(s)
	b[0] = 'a' // 这个不会影响到s，看来这次转换会发生一次copy
	println(s, &s, &b)

	c := []byte{0x31, 0x32}
	ss := string(c)
	println(ss)
	c[1] = 'a' // 这个也不会影响ss，看来这次转换会发生一次copy
	println(ss)

}

func split() {
	arr := strings.SplitN("ab c d e", "c", 2)
	fmt.Printf("%v\n", arr)
}
