package main

import (
	_ "net/http/pprof"
	"net/http"
	"io"
	"crypto/des"
	"runtime"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func UseCpu(w http.ResponseWriter, req *http.Request) {
	use_cpu()
}
func UseNet(w http.ResponseWriter, req *http.Request) {
	use_net()
}
func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/usecpu", UseCpu)
	http.HandleFunc("/usenet", UseNet)
	http.ListenAndServe(":12345", nil)
}
func use_net() {
	for i := 0; i < 10; i++ {
		http.Get("http://www.sina.com.cn/")
	}
}
func use_cpu() {
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
