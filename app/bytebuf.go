package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// 不能按照如下方式给其他地方的 struct 添加方法，错误提示为： cannot define new methods on non-local type bytes.Buffer
//func (this * bytes.Buffer) Seek(i int) {
//	fmt.Print("haha" + strconv.Itoa(i))
//}
func main() {
	var buf bytes.Buffer
	for i := 0; i < 10; i++ { //这样实现字符串的拼接是个比较高效的做法
		buf.WriteString("line" + strconv.Itoa(i) + "\n")
	}
	fmt.Println(buf.String())

	// bytes.Buffer 实现了Write 接口，这个在编码解码时非常方便,参考json.go gob.go
}
