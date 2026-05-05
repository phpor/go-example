package main

import (
	"bytes"
	"flag"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := flag.String("addr", ":8181", "listen addr")
	root := flag.String("root", "./", "file root dir")
	geoFile := flag.String("file", "/tmp/geo.json", "file to safe geo")
	flag.Parse()
	f, err := os.OpenFile(*geoFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		println(err.Error())
		return
	}
	http.Handle("/", http.FileServer(http.Dir(*root)))
	http.HandleFunc("/protoc", ProtocHandlerFunc)
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(time.Millisecond * 20)
		_, _ = writer.Write([]byte("success"))
	})
	http.HandleFunc("/geo", func(writer http.ResponseWriter, request *http.Request) {
		if request.Body == nil {
			_, _ = writer.Write([]byte("empty"))
			return
		}
		defer func() {
			request.Body.Close()
		}()
		b, err := io.ReadAll(request.Body)
		if err != nil {
			_, _ = writer.Write([]byte("error"))
			return
		}
		_, err = f.Write(bytes.Trim(b, "\n"))

		if err != nil {
			_, _ = writer.Write([]byte("error"))
			return
		}
		_, _ = writer.Write([]byte("success"))
	})

	http.HandleFunc("/test2/*", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(request.RequestURI))
	})

	err = http.ListenAndServe(*addr, nil)
	if err != nil {
		panic(err)
	}
}
