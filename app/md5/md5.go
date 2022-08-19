package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	for {
		line, _, err := in.ReadLine()
		if len(line) != 0 {
			arr := strings.Split(string(line), "\t")
			b := md5.Sum([]byte(arr[0]))
			r := append([]string{hex.EncodeToString(b[:])}, arr[1:]...)
			fmt.Printf("%s\n", strings.Join(r, "\t"))
		}
		if err == io.EOF {
			break
		}
	}
}
