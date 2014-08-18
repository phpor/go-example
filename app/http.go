package main

import (
	"fmt"
	"net/http"
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
	http.ListenAndServe(":"+port, nil)
}
