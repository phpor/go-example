// 通过openssl生成rsa密钥： openssl genrsa -3 31 >rsa.key
// 查看密钥信息：  openssl rsa -in rsa.key -text
/*
Private-Key: (31 bit)
modulus: 2055400273 (0x7a82eb51)
publicExponent: 3 (0x3)
privateExponent: 1370204739 (0x51abaa43)
prime1: 57287 (0xdfc7)
prime2: 35879 (0x8c27)
exponent1: 38191 (0x952f)
exponent2: 23919 (0x5d6f)
coefficient: 25237 (0x6295)
writing RSA key
-----BEGIN RSA PRIVATE KEY-----
MCkCAQACBHqC61ECAQMCBFGrqkMCAwDfxwIDAIwnAgMAlS8CAl1vAgJilQ==
-----END RSA PRIVATE KEY-----
 */
package main

import (
	"math/big"
	"fmt"
)

func main() {

	plaintext := big.NewInt(3)
	n := big.NewInt(2055400273)
	d := big.NewInt(1370204739)
	e := big.NewInt(3)
	encryptdata := new(big.Int).Exp(plaintext, d, n)

	decryptdata := new(big.Int).Exp(encryptdata, e, n)
	fmt.Println(*plaintext, *decryptdata)
}

