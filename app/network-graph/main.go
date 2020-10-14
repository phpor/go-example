package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:8181", "listen addr")
	root := flag.String("root", "./", "file root dir")
	flag.Parse()

	data := getData()
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

func getData() []byte {
	data := []byte("dinetwork { node [shape=box];")
	reader := bufio.NewReader(os.Stdin)

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
