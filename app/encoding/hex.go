package main

import (
	"encoding/hex"
	"fmt"
)

func main() {

	b, _ := hex.DecodeString("ac94e2b8f9294bb911c3e424efc593e4")
	fmt.Println(len(b))
}


