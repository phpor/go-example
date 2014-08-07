package main

import (
	"fmt"
	"crypto/rc4"
	"os"
)

func main() {
	//rc4, err := rc4.NewCipher()
	//strings.
	key := []byte("12345")
	src := []byte("wo shi ming wen")
	encrypt := make([]byte, len(src))

	c, err := rc4.NewCipher(key)
	if (err != nil) {
		os.Exit(1)
	}

	c.XORKeyStream(encrypt, src)
	fmt.Println("key:", key)
	fmt.Println("src:", src)
	fmt.Println("encrypt:", encrypt)

	decrypt := make([]byte, len(encrypt))

	c, err = rc4.NewCipher(key)	// 原来的那个cipher是不能接着使用的，即使Reset也不行，就是说，cipher只能用一次

	c.XORKeyStream(decrypt, encrypt)
	fmt.Println("decrypt:", decrypt)
	fmt.Println("decrypt:", string(decrypt))


}
