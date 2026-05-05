// this is a very simple fastcgi web server
package main

import (
	"flag"
	"github.com/yookoala/gofast"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {

	var addr, address, root, prefix string

	flag.StringVar(&addr, "addr", "127.0.0.1:8080", "listen address")
	flag.StringVar(&address, "fastcgi-addr", "127.0.0.1:9000", "fastcgi address")
	flag.StringVar(&root, "root", "/Users/junjie2/data1/git.weibo.com/sso-intra-management/", "www root")
	flag.StringVar(&prefix, "prefix", "/manager/", "url prefix")

	flag.Parse()
	connFactory := gofast.SimpleConnFactory("tcp", address)
	println(root)
	if !strings.HasSuffix(prefix, "/") {
		prefix = prefix + "/"
	}
	if !strings.HasPrefix(prefix, "/") {
		prefix = "/" + prefix
	}
	if !strings.HasPrefix(root, "/") {
		root = root + "/"
	}
	// 把 / 302 到 /manager/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, prefix, http.StatusFound)
	})
	for _, static := range []string{"js", "css", "img", "flash"} {
		path := prefix + static + "/"
		http.Handle(path, http.StripPrefix(path, http.FileServer(http.Dir(root+static+"/"))))
	}
	pool := gofast.NewClientPool(
		gofast.SimpleClientFactory(connFactory),
		1,               // buffer size for pre-created client-connection
		300*time.Second, // life span of a client before expire
	)
	// route all requests to a single php file
	http.Handle(prefix, gofast.NewHandler(
		gofast.NewFileEndpoint(root+"index.php")(gofast.BasicSession),
		pool.CreateClient,
	))

	log.Printf("go-fast-cgi ( http://%s%s ) started, \n", addr, prefix)

	// serve at 8080 port
	log.Fatal(http.ListenAndServe(addr, nil))
}
