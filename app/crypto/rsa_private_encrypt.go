package main

// 源码参考：
// 1. http://play.golang.org/p/dGTl9siO8E
// 2. http://play.golang.org/p/jrqN2KnUEM
// 3. 上面链接来自该讨论： https://groups.google.com/forum/#!topic/Golang-Nuts/Vocj33WNhJQ

// 关于rsa算法的rfc： http://tools.ietf.org/html/rfc2313

// 公钥解密如何实现？

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"os/exec"
)

var (
	ErrInputSize  = errors.New("input size too large")
	ErrEncryption = errors.New("encryption error")
)

var key = map[string]string{
	"pub": `-----BEGIN PUBLIC KEY-----
MDMwDQYJKoZIhvcNAQEBBQADIgAwHwIYLMlPJojwz6CG59dhqBThmSvzDQSqCTER
AgMBAAE=
-----END PUBLIC KEY-----`,
	"pri": `-----BEGIN RSA PRIVATE KEY-----
MIGCAgEAAhgsyU8miPDPoIbn12GoFOGZK/MNBKoJMRECAwEAAQIYEZ21RmEC54gq
yDKNYMSf3sfLnpNKtHApAgxsDM9OncVcR4KpFJMCDGoccYPcO1TysFNuSwIMSB8w
zTgQga0V8QhTAgxCcq1jNXayK4fftyECDB1jgcnOU0DVWmJfFg==
-----END RSA PRIVATE KEY-----`,
}

func PrivateEncrypt(priv *rsa.PrivateKey, data []byte) (enc []byte, err error) {
	//	fmt.Println(priv)
	fmt.Println(data)
	k := (priv.N.BitLen() + 7) / 8
	tLen := len(data)
	// rfc2313, section 8:
	// The length of the data D shall not be more than k-11 octets
	// 参考 rsa.EncryptPKCS1v15(..) 了解如何使用随机数的
	if tLen > k-11 {
		err = ErrInputSize
		return
	}
	em := make([]byte, k)
	em[1] = 1
	for i := 2; i < k-tLen-1; i++ {
		em[i] = 0xff
	}
	fmt.Println(em)
	copy(em[k-tLen:k], data)
	fmt.Println(em)
	c := new(big.Int).SetBytes(em)
	if c.Cmp(priv.N) > 0 {
		err = ErrEncryption
		return
	}
	fmt.Println(c, priv.D, priv.N)
	var m *big.Int
	var ir *big.Int
	if priv.Precomputed.Dp == nil {
		m = new(big.Int).Exp(c, priv.D, priv.N)
	} else {
		// We have the precalculated values needed for the CRT.
		m = new(big.Int).Exp(c, priv.Precomputed.Dp, priv.Primes[0])
		m2 := new(big.Int).Exp(c, priv.Precomputed.Dq, priv.Primes[1])
		m.Sub(m, m2)
		if m.Sign() < 0 {
			m.Add(m, priv.Primes[0])
		}
		m.Mul(m, priv.Precomputed.Qinv)
		m.Mod(m, priv.Primes[0])
		m.Mul(m, priv.Primes[1])
		m.Add(m, m2)

		for i, values := range priv.Precomputed.CRTValues {
			prime := priv.Primes[2 + i]
			m2.Exp(c, values.Exp, prime)
			m2.Sub(m2, m)
			m2.Mul(m2, values.Coeff)
			m2.Mod(m2, prime)
			if m2.Sign() < 0 {
				m2.Add(m2, prime)
			}
			m2.Mul(m2, values.R)
			m.Add(m, m2)
		}
	}

	if ir != nil {
		// Unblind.
		m.Mul(m, ir)
		m.Mod(m, priv.N)
	}
	enc = m.Bytes()
	return
}

func verify_private_encrypt() {
	key_file := "d:\\temp\\rsa.key"
	data_file := "d:\\temp\\tmp.txt"

	// o is output from openssl
	o, _ := exec.Command("openssl", "rsautl", "-sign", "-inkey", key_file, "-in", data_file).Output()

	// t.key is private keyfile
	// in.txt is what to encode
	kt, _ := ioutil.ReadFile(key_file)
	e, _ := ioutil.ReadFile(data_file)
	block, _ := pem.Decode(kt)
	privkey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	encData, _ := PrivateEncrypt(privkey, e)
	fmt.Println(encData)

	fmt.Println(o)
	fmt.Println(string(o) == string(encData))
}
func enc(text string) {
	fmt.Println(text)
	data := []byte(text)
	block, _ := pem.Decode([]byte(key["pri"]))
	privkey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	encData, _ := PrivateEncrypt(privkey, data)
	fmt.Println(encData)

	block2, _ := pem.Decode([]byte(key["pub"]))
	pubinterface, err := x509.ParsePKIXPublicKey(block2.Bytes)
	if err != nil {
		println(err)
		return
	}
	plaindata, err := PublicDecrypt(pubinterface.(*rsa.PublicKey), encData)
	fmt.Println(string(plaindata))

}
func main() {

	enc("abcd")
	enc("abcd")
	//	enc("abcdef")
	//	enc("12345")

}
func verify_public_decrypt() {
	//	pubInterface,_ := x509.ParsePKIXPublicKey(block.Bytes)
	//	plain,_ := PublicDecrypt(pubInterface.(*rsa.PublicKey), encData)
	//	println(string(plain))
}
func PublicDecrypt(pubkey *rsa.PublicKey, enc []byte) ([]byte, error) {
	k := (pubkey.N.BitLen() + 7) / 8
	if k != len(enc) {
		return nil, errors.New("enc data length error")
	}
	m := new(big.Int).SetBytes(enc)

	if m.Cmp(pubkey.N) > 0 {
		return nil, errors.New("enc data too long")
	}
	m.Exp(m, big.NewInt(int64(pubkey.E)), pubkey.N)

	d := leftPad(m.Bytes(), k)

	if d[0] != 0 {
		return nil, errors.New("data broken, first byte is not zero")
	}

	if d[1] != 0 && d[1] != 1 {
		return nil, errors.New("data is not encrypt by private key")
	}

	fmt.Println(d)
	fmt.Println(len(d))

	var i = 2
	for ; i < len(d); i++ {
		if d[i] == 0 {
			break
		}
	}
	i++
	if i == len(d) {
		return nil, nil
	}

	fmt.Println(d[i:])
	return d[i:], nil
}

// copy from crypto/rsa/rsa.go
// leftPad returns a new slice of length size. The contents of input are right
// aligned in the new slice.
func leftPad(input []byte, size int) (out []byte) {
	n := len(input)
	if n > size {
		n = size
	}
	out = make([]byte, size)
	copy(out[len(out)-n:], input)
	return
}
