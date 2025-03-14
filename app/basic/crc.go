package main

import "hash/crc32"

func main() {
	println(crc32.ChecksumIEEE([]byte("xxxxxxx111")) % 100)
}
