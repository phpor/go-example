package main

// 参考 net/http/doc.go
import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	// Usage: ./test_groupcache port
	//port := "" + os.Args[1]
	port := "8082"
	fmt.Println("start server at " + port + "\n")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			//		w.Write([]byte(r.Method))
			fmt.Fprintln(w, r)
			fmt.Fprintln(w, r.Method)
			fmt.Fprintln(w, r.Host)
			fmt.Fprintln(w, r.RequestURI)
			fmt.Fprintln(w, r.PostForm)
			fmt.Fprintln(w, r.Cookies())
			fmt.Fprintln(w, r.URL.RawQuery)

		})
	go func() {
		eatmem := func() []byte {
			s := make([]byte, 1024*1024)
			s[1024] = 0x80
			println(s)
			return s
		}
		for {
			eatmem()
			time.Sleep(1 * time.Millisecond)
		}
	}()
	http.ListenAndServe(":"+port, nil)

	//http.ListenAndServeTLS() 可以提供https服务
}
