package main

import (
	"flag"
	"net/http"
)

func main() {
	cfgFile := flag.String("c", "", "config file")
	addr := flag.String("addr", "8081", "addr")
	flag.Parse()

	if *cfgFile != "" {
		InitServiceByFile(*cfgFile)
	}

	http.HandleFunc("/services/", func(writer http.ResponseWriter, request *http.Request) {

	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		panic(err)
	}

}
