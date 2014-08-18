package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	var buf bytes.Buffer
	for i := 0; i < 10; i++ { //这样实现字符串的拼接是个比较高效的做法
		buf.WriteString("line" + strconv.Itoa(i) + "\n")
	}
	fmt.Println(buf.String())

	// bytes.Buffer 实现了Write 接口，这个在编码解码时非常方便,参考json.go gob.go
}
