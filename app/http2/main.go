package main

import (
	"golang.org/x/net/http2"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http2Server()

}

func httpClient() {
	url := "https://baidu.com/"
	if len(os.Args) >= 2 {
		url = os.Args[1]
	}
	resp, err := http.Get(url)
	if err != nil {
		println(err.Error())
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	println(string(b))
}

func http2Server() {
	var srv http.Server
	//http2.VerboseLogs = true
	srv.Addr = ":8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello http2"))
	})
	http2.ConfigureServer(&srv, &http2.Server{})
	println(srv.ListenAndServe())

}
