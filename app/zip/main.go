package main

import (
	"compress/flate"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"os"
	"strings"
)

func main() {
	//xorDecodeTest()
	fmt.Println(parseThrow(os.Stdin))
}

func xorDecodeTest() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(xorEncode(string(b), "password"))
}
func decodeThrowTest() {
	r, err := os.Open("/Users/junjie2/data1/github.com/phpor/go-example/app/zip/throw.log")
	if err != nil {
		panic(err)
	}
	throw := parseThrow(r)
	validator := "4BUrOzJ94qMruc5t1gts5iDuBKKPZOzbYT+F816dWOo="

	//b, err := base64.StdEncoding.DecodeString(validator)
	//validator = string(b)

	fmt.Println(decodeThrow(throw, validator, "1"))
}

func parseThrow(r io.Reader) string {
	boundary := os.Getenv("boundary")
	mr := multipart.NewReader(r, boundary)
	f, err := mr.ReadForm(1024*1024*10)
	if err != nil && err != io.EOF {
		panic(err)
	}
	return url.Values(f.Value).Get("throw")
}

func decodeThrow(throw, validator, isZip string) string {
	if throw == "" {
		return ""
	}
	if isZip != "1" {
		return throw
	}

	if validator != "" {
		throw = xorEncode(throw, validator)
	}
	throw = throw[10:len(throw)-8]
	b, err := ioutil.ReadAll(flate.NewReader(strings.NewReader(throw)))
	if err != nil {
		return ""
	}
	throw = string(b)
	return throw
}

/**
/*

function xor_encode($data, $key)
{
    $dataLen = strlen($data);
    $keyLen = strlen($key);

    for ($text = '', $i = 0; $i < $dataLen; $i++) {
        $text .= $data{$i} ^ $key{$i % $keyLen};
    }

    return $text;
}

 */
func xorEncode(data , key string) string {
	res := []byte(data[:])
	if len(key) == 0 {
		return data
	}
	for i := range res {
		res[i] = res[i] ^ key[i%len(key)]
	}
	return string(res)
}

func substr(str string) string{
	return str[10:len(str)-8]
}

func deflate(str string) string {
	r := flate.NewReader(strings.NewReader(str))

	b,err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return string(b)
}
