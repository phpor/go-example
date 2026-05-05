/**
 * 写一个程序，有两个功能：
 * 1. 加密：读取命令行参数提供的csv文件，一次可以接受多个文件，每个文件只有一列， 对每个文件的第一列进行加密，加密后保存到另一个文件，加密后的文件有8个字节的MAGIC，其余都是16个字节的MD5字符串的反转。
 * 2. 解密：读取命令行参数提供的加密文件，一次可以接受多个文件， 对每个文件进行解密，解密逻辑为： 1、校验8字节的magic是否正确 2、 循环读取16字节，并进行反转，然后做hex encode， 处理成32字节的MD5。
 * 加密规则：对每一个输入做md5加密，不做hex encode，然后做一个字符串反转操作。
 *
 */

package main

import (
	"bytes"
	"crypto/md5"
	"encoding/csv"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var magic = []byte{1, 0, 1, 0, 1, 0, 1, 0}

func main() {
	var err error
	for i := 1; i < len(os.Args); i++ {
		inputFile := os.Args[i]
		fmt.Printf("proccessing %s ", inputFile)
		if os.Getenv("DECRYPT_MODE") == "ON" && strings.HasSuffix(inputFile, ".encrypt") {
			err = decryptFile(inputFile)
		} else {
			err = encryptFile(inputFile)
		}
		if err != nil {
			fmt.Printf("fail: %s\n", err.Error())
		} else {
			fmt.Printf("success\n")
		}

	}
}

func encryptFile(inputFile string) error {
	in, err := os.OpenFile(inputFile, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	out, err := os.OpenFile(inputFile+".encrypt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	_, err = out.Write(magic)
	if err != nil {
		return err
	}
	r := csv.NewReader(in)
	for {
		ss, err := r.Read()
		if err == io.EOF {
			break
		}
		sb := md5.Sum([]byte(ss[0]))
		_, err = out.Write(reverseHex(sb[:]))
		if err != nil {
			return err
		}
	}
	_ = out.Close()
	_ = in.Close()

	return nil
}

func decryptFile(inputFile string) error {
	in, err := os.OpenFile(inputFile, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	out, err := os.OpenFile(strings.TrimSuffix(inputFile, ".encrypt")+".decrypt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	m := make([]byte, 8)
	_, err = in.Read(m[:])
	if err != nil {
		return err
	}
	if !bytes.Equal(magic, m) {
		return errors.New("not encrypt file")
	}
	record := make([]byte, 16)
	hexCode := make([]byte, 33)
	hexCode[32] = '\n'
	for {
		n, err := in.Read(record)
		if err == io.EOF {
			break
		}
		if n != 16 {
			return errors.New("file is broken")
		}
		record = reverseHex(record[:])
		hex.Encode(hexCode, record)
		_, err = out.Write(hexCode)
		if err != nil {
			return err
		}
	}
	_ = out.Close()
	_ = in.Close()

	return nil
}

func reverseHex(data []byte) []byte {
	reversed := make([]byte, len(data))
	for i, b := range data {
		reversed[len(data)-1-i] = b
	}
	return reversed
}
