package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

// 不能按照如下方式给其他地方的 struct 添加方法，错误提示为： cannot define new methods on non-local type bytes.Buffer
//func (this * bytes.Buffer) Seek(i int) {
//	fmt.Print("haha" + strconv.Itoa(i))
//}
func main() {
	bytesMd5()
	//trunkBytes()
}

func bytesEqual() {
	b1 := []byte("phpor")
	b2 := []byte("phpor")
	var o1 interface{}
	var o2 interface{}
	o1 = b1
	o2 = b2
	//println(b1 == b2)   字节数组 == 比较是编译不过去的，语法错误
	println(o1 == o2) // 虽然interface可以比较，但是这个更加不能用，会出现运行时异常；
	// 所以，永远不要直接比较interface，interface只能用于接口，想执行任何具体的操作，都需要转成具体的类型再操作
	// 如果interface包装的具体对象是可以比较的，则该interface比较也不会panic，但是也不要指望比较的解决符合直觉
}

func bytesMd5() {
	buf := md5.Sum(nil)
	fmt.Printf("%s", hex.EncodeToString(buf[:]))
}

func bytesTest1() {
	var buf bytes.Buffer
	for i := 0; i < 10; i++ { //这样实现字符串的拼接是个比较高效的做法
		buf.WriteString("line" + strconv.Itoa(i) + "\n")
	}
	fmt.Println(buf.String())

	// bytes.Buffer 实现了Write 接口，这个在编码解码时非常方便,参考json.go gob.go
}
func trunkBytes() {
	b := []byte{}
	c := b[1:]
	fmt.Printf("%s", c)
}

func byteInt() {
	a := 0x1011
	b := byte(a)
	println(b)
}
