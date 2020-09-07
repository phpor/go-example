package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	//r := bytes.NewReader([]byte("1\tx\n1\ty\n2\ty"))
	r := os.Stdin
	c := NewClassicer2(r, func(i []byte) []byte {
		if len(i) == 0 {
			return nil
		}
		return bytes.Split(i, []byte("\t"))[0]
	})
	for c.Scan() {
		for l, v := range c.Lines() {
			fmt.Printf("%d: %s\n", l, string(v))
		}
		//fmt.Printf("%s\n", "---------")
	}
}

// 这个程序处理的文件的格式
// 12345\txxxxx
// 12345\tyyyyy
// 22222\txxxxx
// 33333\tuuuuu
// 如果第一列连续重复N次，则将这样的行分类到 class.N 的文件中
func classic() {
	stdin := bufio.NewReader(os.Stdin)
	var currentUid []byte
	var arrBytes [][]byte
	prefix := "classic."
	if len(os.Args) > 1 {
		prefix = os.Args[1]
	}

	arrFp := map[int]*os.File{}

	write := func() {
		if _, ok := arrFp[len(arrBytes)]; !ok {
			arrFp[len(arrBytes)], _ = os.Create(fmt.Sprintf("%s%d", prefix, len(arrBytes)))
		}
		for l := range arrBytes {
			arrFp[len(arrBytes)].Write(arrBytes[l])
		}
	}
	for {
		line, err := stdin.ReadBytes('\n')
		if len(line) > 0 {
			arr := bytes.Split(line, []byte{'\t'})
			if len(currentUid) == 0 {
				currentUid = arr[0]
			}
			if !bytes.Equal(arr[0], currentUid) {
				write()
				arrBytes = nil
				currentUid = nil
			}
			arrBytes = append(arrBytes, line)
		}
		if err != nil {
			break
		}
	}
	write()
}
