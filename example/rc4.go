package main

import (
	"crypto/rc4"
	"fmt"
	"os"
)

func main() {
	key := []byte("this is key")
	src := []byte("this is plain text")
	en, err := encrypt(key, src)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("key:", key)
	fmt.Println("src:", src)
	fmt.Println("encrypt:", en)

	decrypt, err := encrypt(key, en)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("decrypt:", decrypt)
	fmt.Println("decrypt:", string(decrypt))

}

func encrypt(key, plain []byte) ([]byte, error) {
	c, err := rc4.NewCipher(key) // cipher不能重复使用的，即使Reset也不行，就是说，cipher只能用一次
	if err != nil {
		return nil, err
	}
	en := make([]byte, len(plain))
	c.XORKeyStream(en, plain)
	return en, nil
}
