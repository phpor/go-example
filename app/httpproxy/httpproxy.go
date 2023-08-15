package main

import (
	"flag"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
	"os"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.KeepHeader = true
	proxy.NonproxyHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		req.URL.Scheme = "http"
		req.Host = req.Header.Get("Host")
		if req.Host == "" {
			req.Host = os.Getenv("DEFAULT_HOST")
			req.URL.Host = req.Host
			//req.URL.Path = "http://api.weibo.cn" + req.URL.Path
		}
		// 默认http协议
		proxy.ServeHTTP(w, req)
		//http.Error(w, "This is a proxy server. Does not respond to non-proxy requests.", 500)
	})
	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	addr := flag.String("addr", ":8686", "proxy listen address")
	flag.Parse()

	proxy.Verbose = *verbose
	log.Fatal(http.ListenAndServe(*addr, proxy))
}
