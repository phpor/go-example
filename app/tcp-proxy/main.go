package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
)

var port string
var replaceString = map[string]string{}
var ch chan error

func main() {
	addr := ""
	replaceStr := ""
	flag.StringVar(&addr, "addr", ":9393", "listen addr")
	flag.StringVar(&replaceStr, "replace", "aaa->bbb", "split by ->")
	flag.Parse()
	ss := strings.Split(replaceStr, "->")
	if len(ss) == 2 {
		replaceString[ss[0]] = ss[1]
	}

	ch = make(chan error, 1)
	go server(addr)
	if err := <-ch; err != nil {
		return
	}
	http.HandleFunc("/update-port", func(writer http.ResponseWriter, request *http.Request) {
		port = request.FormValue("port")
		log.Println("update port:", port)
		_, _ = writer.Write([]byte("ok"))
		return
	})
	_ = http.ListenAndServe(":9494", nil)
}

// server tcp listen 端口9393， 有连接进来时，
func server(addr string) {

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println(fmt.Sprintf("server start fail: %s", err.Error()))
		ch <- err
		return
	}
	close(ch)
	log.Println("server start")

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			handle(conn)
		}()
	}
	return
}

// handle 转发数据到 127.0.0.1 的 1234 端口，并替换转发数据中的 aaa 为 bbb
func handle(conn net.Conn) {
	addr := "127.0.0.1:" + port
	defer func() {
		_ = conn.Close()
		log.Printf("finished connection %s\n", addr)
	}()
	// 创建一个连接到 127.0.0.1 的 1234 端口的连接
	remote, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(fmt.Sprintf("connect %s fail: %s", addr, err.Error()))
		return
	}
	defer func() {
		_ = remote.Close()
	}()

	wg := sync.WaitGroup{}
	wg.Add(2)
	// 启动两个 goroutine 来复制数据
	go func() {
		defer wg.Done()
		if _, err := io.Copy(conn, replaceFilter(remote)); err != nil {
			log.Println(err)
		}

	}()

	go func() {
		defer wg.Done()
		if _, err := io.Copy(remote, replaceFilter(conn)); err != nil {
			log.Println(err)
		}
	}()
	wg.Wait()
}

// replaceFilter 创建一个流过滤器，将输入的io.Reader中的"aa"替换为"bbbb"。
func replaceFilter(r io.Reader) io.Reader {
	return &replaceReader{r: r, buf: make([]byte, 32*1024)}
}

type replaceReader struct {
	r   io.Reader
	buf []byte
}

func (r *replaceReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(r.buf)
	if err != nil {
		return n, err
	}
	replaced := r.buf[:n]
	for k, v := range replaceString {
		replaced = bytes.ReplaceAll(replaced, []byte(k), []byte(v))
	}
	newData := replaced

	// 将处理后的内容拷贝到输出缓冲区
	copied := copy(p, newData)
	return copied, nil
}
