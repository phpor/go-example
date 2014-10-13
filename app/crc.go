package main

import (
	"hash/crc32"
	"strconv"
)

func main() {
	testcrc32()
	println(strconv.Itoa(1234))
}
func testcrc32() {
	println(crc32.ChecksumIEEE([]byte("1234")))
}

