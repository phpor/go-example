package main
import (
	"net/http"
	"flag"
)

func main() {
	addr := flag.String("addr", ":8181", "listen addr")
	root := flag.String("root", "./", "file root dir")
	flag.Parse()
	http.Handle("/", http.FileServer(http.Dir(*root)))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		panic(err)
	}
}