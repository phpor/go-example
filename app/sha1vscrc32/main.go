package main

import (
	"crypto/sha1"
	"hash/crc32"
	"time"
)

func main() {
	s := time.Now()
	b := []byte("phpor")
	var c [20]byte
	for i := 0; i < 100000; i++ {
		c = sha1.Sum(b) // 这个要比 New => Write => Sum 要快50%
		//hash := sha1.New()
		//hash.Write(b)
		//_ = hash.Sum(nil)
	}
	e := time.Now()
	println(len(c))
	println(e.Sub(s).String())

	var r uint32
	s = time.Now()
	for i := 0; i < 100000; i++ {
		r = crc32.ChecksumIEEE(b) // crc32 要比sha1快不止10倍
	}
	e = time.Now()
	println(r)
	println(e.Sub(s).String())
}
