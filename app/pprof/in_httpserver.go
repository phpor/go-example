package main

import (
	"crypto/des"
	"io"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func UseCpu(w http.ResponseWriter, req *http.Request) {
	useCpu()
}
func UseNet(w http.ResponseWriter, req *http.Request) {
	useNet()
}
func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/usecpu", UseCpu)
	http.HandleFunc("/usenet", UseNet)
	http.ListenAndServe(":12345", nil)
}
func useNet() {
	for i := 0; i < 10; i++ {
		http.Get("http://www.sina.com.cn/")
	}
}
func useCpu() {
	for i := 0; i < 10; i++ {
		c, err := des.NewCipher([]byte("12345678"))
		if err != nil {
			panic(err)
		}

		dst := make([]byte, des.BlockSize)
		c.Encrypt(dst, []byte("12345678"))
		runtime.Gosched()
	}
}
