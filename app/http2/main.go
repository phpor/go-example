package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, HTTP/2!")
	})

	// 配置服务器以支持HTTP/2的明文通信
	server := &http.Server{
		Addr: ":8080",
		//Handler: h2c.NewHandler(http.DefaultServeMux, &http2.Server{}),
	}

	fmt.Println("Starting HTTP/2 server on :8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}

//
//func http2Server() {
//	var srv http.Server
//	//http2.VerboseLogs = true
//	srv.Addr = ":8080"
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("hello http2"))
//	})
//	err := http2.ConfigureServer(&srv, &http2.Server{})
//	if err != nil {
//		println(err.Error())
//	}
//	go func() {
//		time.Sleep(3 * time.Second)
//		srv.Shutdown(context.TODO())
//	}()
//	println(srv.ListenAndServe().Error())
//
//}
//
//func http2Server2() {
//	var srv http2.Server
//	//http2.VerboseLogs = true
//	srv.Addr = ":8080"
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("hello http2"))
//	})
//	err := http2.ConfigureServer(&srv, &http2.Server{})
//	if err != nil {
//		println(err.Error())
//	}
//	println(srv.ListenAndServe())
//
//}
