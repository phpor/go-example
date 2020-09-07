package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	scan()
}
func readString() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		l, err := stdin.ReadString('\n')
		if l != "" { // 这里总是应该先判断是否有数据，而不是先判断err，否则，最后一个非回车结尾的行会被丢掉
			fmt.Printf("%s", l)
		}
		if err == io.EOF {
			break
		}
	}
}

func scan() {
	//r := bytes.NewReader(bytes.Repeat([]byte("xxxxx yyyyyyyyyy\n"), 10000))
	r := os.Stdin
	scan := bufio.NewScanner(r) // 默认每次按行切割，每行最大64k，返回的内容不包含换行符，如果是\r\n 结尾的话，也会吧 \r 去掉
	i := 0
	for scan.Scan() { // 这个比ReadString要好用一些，不必关心错误
		_ = scan.Bytes() // 如果word不超过64k，就不是在没有换行的地方截断的，通常我们知道每行都不长，也不用去判断长度
		i++
		//fmt.Printf("%s\n", string(word))
	}
	fmt.Printf("i: %d", i)
	if err := scan.Err(); err != nil && err != io.EOF { //
		fmt.Printf("%s\n", err.Error())
	}
}
