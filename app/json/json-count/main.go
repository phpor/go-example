package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	envJqKeys := strings.TrimSpace(os.Getenv("JQ_KEYS"))
	if len(envJqKeys) == 0 {
		println("please set env JQ_KEYS")
		return
	}
	ss := strings.Split(envJqKeys, ",")
	ssCnt := map[string]int{}
	for _, v := range ss {
		ssCnt[v] = 0
	}
	r := bufio.NewReader(os.Stdin)
	for {
		m := map[string]string{}
		line, _, err := r.ReadLine()
		if len(line) == 0 {
			if err == io.EOF {
				break
			}
			continue
		}
		_ = json.Unmarshal(bytes.TrimSpace(line), &m)
		if len(m) == 0 {
			continue
		}
		for k := range ssCnt {
			if m[k] != "" {
				ssCnt[k]++
			}
		}
	}
	for k := range ssCnt {
		fmt.Printf("%s\t%d\n", k, ssCnt[k])
	}
}
