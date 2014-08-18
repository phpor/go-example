package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

// 参考资料： http://blog.golang.org/gobs-of-data
// gob 似乎是一个Go语言的序列化（反序列化）工具；gob是面向go的，是个非其它语言友好的编码格式，类似于PHP的serialize
func main() {
	data := map[int16]string{1: "first", 2: "second"}
	en := gob.NewEncoder(os.Stdout)
	en.Encode(data)

	fmt.Println()
	var buf bytes.Buffer
	en2 := gob.NewEncoder(&buf)
	en2.Encode(data)
	fmt.Print(buf)

}
