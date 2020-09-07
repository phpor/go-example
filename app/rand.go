package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//b := make([]byte, 20)
	//binary.LittleEndian.PutUint64(b,uint64(time.Now().Add(80*365*24*time.Hour).Unix()))
	//fmt.Printf("%v\n", b)
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//fmt.Printf("%08x\n",rand.Int31n(32767) << 16 | 15)

	fmt.Printf("%v\n", uuid())
	fmt.Printf("%v\n", uuid())
	fmt.Printf("%v\n", uuid())

	//fmt.Printf("%0x", uint64(rand.Int31n(32767) << 16) | uint64(time.Now().Add(80*365*24*time.Hour).Unix())&0x0000ffff)
}

func uuid() uint32 {
	r := rand.Int31n(32767)
	t := time.Now().Unix()
	c := []byte{byte(r >> 8), byte(r), byte(t >> 8), byte(t)}
	return binary.LittleEndian.Uint32(c)
}
