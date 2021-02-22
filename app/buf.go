package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

func main() {
	bufioReaderTest()
}

func bufTest1() {
	b := make([]byte, 18)
	b[0] = 'a'
	buf := bytes.NewBuffer(b)
	buf.Write([]byte{'b'}) // 这里的b没有挨着a的，但是NewBuffer 的时候不会copy一份内存，但是，会随着Write的执行而产生一个新的slice，不管如何bufer和slice不要混着用
	fmt.Printf("%v", buf)
}

func bufioReaderTest() {
	reader := bufio.NewReader(bytes.NewReader([]byte("1\n2\n")))
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF { // 一般来讲，都会是 \n 结尾的，所以，先判断err通常都不会有问题
			break
		}
		println(line)
	}
}
