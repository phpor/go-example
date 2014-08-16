package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {

	ss := []string{}
	for i := 0; i < 10; i++ {
		ss = append(ss, "line", strconv.Itoa(i), "\n")    //这个要比 += 操作的效率高
	}
	fmt.Println(strings.Join(ss, ""))
}
