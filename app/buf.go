package main

import (
	"bytes"
	"fmt"
)

func main() {
	bufTest1()
}

func bufTest1() {
	b := make([]byte, 18)
	b[0] = 'a'
	buf := bytes.NewBuffer(b)
	buf.Write([]byte{'b'}) // 这里的b没有挨着a的，但是NewBuffer 的时候不会copy一份内存，但是，会随着Write的执行而产生一个新的slice，不管如何bufer和slice不要混着用
	fmt.Printf("%v", buf)
}
