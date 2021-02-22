package main

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) == 0 {
		fmt.Printf("Usage: %s str\n")
		return
	}
	str := os.Args[1]
	b, err := hexDecode(str)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", b)
	fmt.Printf("%d\n", binary.BigEndian.Uint64(b))
	return
}

func hexDecode(str string) ([]byte, error) {
	strBytes := []byte(str)
	buf := make([]byte, len(str))
	i := 0
	j := 0
	length := len(strBytes)
	for i < length {
		if strBytes[i] == '\\' && strBytes[i+1] == 'x' {
			i += 2
			if i+2 > length {
				return nil, errors.New("invalid string")
			}
			_, err := hex.Decode(buf[j:], strBytes[i:i+2])
			if err != nil {
				return nil, err
			}
			i += 2
		} else {
			buf[j] = strBytes[i]
			i++
		}
		j++
	}
	return buf, nil
}
