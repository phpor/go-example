package main

import (
	"flag"
	"github.com/elazarl/goproxy"
	"github.com/elazarl/goproxy/ext/auth"
	"log"
	"net/http"
	"os"
)

func main() {
	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	addr := flag.String("addr", ":8686", "proxy listen address")
	username := flag.String("username", "", "proxy auth username")
	password := flag.String("password", "", "proxy auth password")
	flag.Parse()

	proxy := goproxy.NewProxyHttpServer()
	proxy.KeepHeader = true
	//proxy.OnRequest().HandleConnectFunc(func(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
	//	c := goproxy.OkConnect
	//	c.TLSConfig = func(host string, ctx *goproxy.ProxyCtx) (*tls.Config, error) {
	//
	//	}
	//	return c, ""
	//})
	if username != nil && *username != "" && password != nil && *password != "" {
		proxy.OnRequest().HandleConnect(auth.BasicConnect("my_realm", func(user, passwd string) bool {
			return user == *username && passwd == *password
		}))
		proxy.OnRequest().Do(auth.Basic("my_realm", func(user, passwd string) bool {
			return user == *username && passwd == *password
		}))
	}
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

	proxy.Verbose = *verbose

	log.Fatal(http.ListenAndServe(*addr, proxy))
}
