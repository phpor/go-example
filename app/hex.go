package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
)

func main() {
	a := math.Pow(2, 17) + 65535
	fmt.Printf("%d\n", uint16(a+2))
	buf := littleEndian(0x0000)
	i := 7
	for i > 0 { // 就算全是0字节，也保留一个
		if buf[i] != byte(0x0) {
			break
		}
		i--
	}
	buf = buf[:i+1]
	b := [16]byte{}
	hex.Encode(b[:], []byte(buf))
	fmt.Printf("%s\n", string(b[:]))
}

func littleEndian(uid uint64) string {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uid)
	return string(buf)
}
