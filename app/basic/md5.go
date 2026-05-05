package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
	s := "abcd"
	m := md5.Sum([]byte(s))
	b1 := hex.EncodeToString(m[:])
	b2 := sha256.Sum256([]byte(b1[:]))
	fmt.Println(hex.EncodeToString(b2[:]))
	t := time.Time{}
	fmt.Printf("%t\t%t\t%d\n", true, false, t.Unix())
}
