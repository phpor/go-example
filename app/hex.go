package main

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
)

func main() {
	xxx()
}

func yyyyy() {
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

func xxx() {
	base64Encode := base64.StdEncoding
	v1, _ := hex.DecodeString("81418f163d3e3deffcb95b4277096126")
	fmt.Printf("v1: %s\n", base64Encode.EncodeToString(v1))
	v2, _ := hex.DecodeString("d0267736ae5258365f4ed04876bc9ff8")
	fmt.Printf("v2: %s\n", base64Encode.EncodeToString(v2))

	iv1, _ := hex.DecodeString("daa9eb481587fe83254da2b1e3f2708b")
	fmt.Printf("iv1: %s\n", base64Encode.EncodeToString(iv1))
	iv2, _ := hex.DecodeString("3810e99176d82c3745a3fe223fd8fa7b")
	fmt.Printf("iv2: %s\n", base64Encode.EncodeToString(iv2))
}
