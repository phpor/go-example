package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:8181", "listen addr")
	root := flag.String("root", "./", "file root dir")
	removeLeaf := flag.Bool("remove-leaf", false, "file root dir")
	flag.Parse()

	data := getData(*removeLeaf)
	http.Handle("/", http.FileServer(http.Dir(*root)))
	http.HandleFunc("/data", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(data)
	})
	fmt.Printf("http://%s/network.html\n", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		panic(err)
	}
}

func getData(removeLeaf bool) []byte {
	var reader *bufio.Reader
	if !removeLeaf {
		reader = bufio.NewReader(os.Stdin)
	} else {
		buf, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		arr := bytes.Split(buf, []byte("\n"))
		KeyMap := map[string]struct{}{}
		vMap := map[string]struct{}{}
		for _, l := range arr {
			arr1 := bytes.Split(l, []byte(" "))
			if len(arr1) < 2 {
				continue
			}
			KeyMap[string(arr1[0])] = struct{}{}
			vMap[string(arr1[1])] = struct{}{}
		}
		buf = nil
		for _, l := range arr {
			arr1 := bytes.Split(l, []byte(" "))
			if len(arr1) < 2 {
				continue
			}
			if _, ok := KeyMap[string(arr1[1])]; !ok {
				continue
			}
			buf = append(buf, '\n')
			buf = append(buf, l...)
		}
		reader = bufio.NewReader(bytes.NewReader(buf))
	}
	data := []byte("dinetwork { node [shape=box];")

	for {
		b, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		b = bytes.TrimSpace(b)
		if len(b) == 0 {
			continue
		}
		arr := bytes.Split(b, []byte(" "))
		if len(arr) < 2 {
			continue
		}
		data = append(data, '"')
		data = append(data, bytes.Join(arr, []byte("\"->\""))...)
		data = append(data, "\";"...)
	}
	data = append(data, '}')
	return data
}
