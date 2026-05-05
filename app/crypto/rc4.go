package main

import (
	"bytes"
	"crypto/rc4"
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	i := 10
	for i > 0 {
		plain := strings.Repeat("a", rand.Intn(1000))
		en := DIDEncrypt(plain)
		println("encrypt:\n", en)
		p := DIDDecrypt(en)
		println("info:\n", plain, "\n", p)

		if plain != p {
			println("fail:\n", plain, "\n", p)
		}
		i--
	}

}

func testEncrypt() {
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

var didSalt = []byte("dbe39d9d56f98d5f9cb22df15b285791")

func DIDEncrypt(did string) string {
	if did == "" {
		return ""
	}
	did = did + did[0:1] // 把第一位放到最后一位，作为校验位
	c, err := rc4.NewCipher(didSalt)
	if err != nil {
		return ""
	}
	dst := make([]byte, len(did))
	c.XORKeyStream(dst, []byte(did))
	return base64.RawURLEncoding.EncodeToString(dst)
}

func DIDDecrypt(did string) string {
	if did == "" {
		return ""
	}
	b, err := base64.RawURLEncoding.DecodeString(did)
	if err != nil {
		return ""
	}
	c, err := rc4.NewCipher(didSalt)
	if err != nil {
		return ""
	}
	dst := make([]byte, len(b))
	c.XORKeyStream(dst, b)
	if !bytes.HasSuffix(dst, dst[0:0]) {
		return ""
	}
	dst = dst[0 : len(dst)-1]
	return string(dst) // 去掉最后一个校验位
}
