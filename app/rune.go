package main

import (
	"fmt"
	"strconv"
)

func main() {
	n, _ := strconv.ParseInt("65f6", 16, 64)
	fmt.Printf("%s", rune(n))
}
