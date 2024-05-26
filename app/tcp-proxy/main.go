package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var port string
var replaces = replaceMap{}
var ch chan error

func main() {
	tcpAddr := ""
	httpAddr := ""
	flag.StringVar(&tcpAddr, "tcp_addr", ":9393", "tcp proxy listen addr")
	flag.StringVar(&httpAddr, "http_addr", ":9494", "http server listen addr")
	flag.Var(&replaces, "replace", "split by -> ; eg: aaa->bbb means replace aaa to bbb")
	flag.Parse()
	log.Printf("replaces: %s", replaces)

	ch = make(chan error, 1)
	go server(tcpAddr)
	if err := <-ch; err != nil {
		return
	}
	http.HandleFunc("/update-port", func(writer http.ResponseWriter, request *http.Request) {
		p := request.FormValue("port")
		if n, err := strconv.Atoi(p); err != nil || n <= 0 || n >= 65535 {
			log.Printf("update port to %s fail", p)
			_, _ = writer.Write([]byte(fmt.Sprintf("port %s is invalid", p)))
		} else {
			log.Println("update port:", n)
			port = p
			_, _ = writer.Write([]byte("ok"))
		}
		return
	})
	http.HandleFunc("/replace", func(writer http.ResponseWriter, request *http.Request) {
		from := request.FormValue("from")
		to := request.FormValue("to")
		if from == "" {
			_, _ = writer.Write([]byte("from can not empty"))
			return
		}
		replaces[from] = to
		_, _ = writer.Write([]byte(fmt.Sprintf("success: replace %s to %s", from, to)))
		return
	})
	http.HandleFunc("/replace/show", func(writer http.ResponseWriter, _ *http.Request) {
		_, _ = writer.Write([]byte(fmt.Sprintf("%s", replaces.String())))
		return
	})
	ln, err := net.Listen("tcp", httpAddr)
	if err != nil {
		log.Println(fmt.Sprintf("http server start fail: %s", err.Error()))
		return
	}
	log.Println(fmt.Sprintf("http server start success: %s", httpAddr))
	_ = http.Serve(ln, nil)
}

// server tcp listen 端口9393， 有连接进来时，
func server(tcpAddr string) {

	ln, err := net.Listen("tcp", tcpAddr)
	if err != nil {
		log.Println(fmt.Sprintf("tcp proxy server start fail: %s", err.Error()))
		ch <- err
		return
	}
	close(ch)
	log.Println("tcp poxy server started: " + tcpAddr)

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
	log.Printf("received connection %s -> %s\n", conn.RemoteAddr(), conn.LocalAddr())

	// 创建到上游的连接
	remote, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(fmt.Sprintf("connect %s fail: %s", addr, err.Error()))
		return
	}
	log.Printf("connect %s success", addr)

	wg := sync.WaitGroup{}
	wg.Add(2) // 任何一边断开就可以退出了
	// 启动两个 goroutine 来复制数据
	go func() {
		defer wg.Done()
		if _, err := io.Copy(conn, replaceFilter(remote)); err != nil {
			log.Println(err)
		}
		_ = remote.Close()
		log.Printf("closed connection %s -> %s\n", remote.LocalAddr(), remote.RemoteAddr())

	}()

	go func() {
		defer wg.Done()
		if _, err := io.Copy(remote, replaceFilter(conn)); err != nil {
			log.Println(err)
		}
		_ = conn.Close()
		log.Printf("closed connection %s -> %s\n", conn.RemoteAddr(), conn.LocalAddr())

	}()
	wg.Wait()
}

// replaceFilter 创建一个流过滤器，可以自定义替换流中的字符串。
func replaceFilter(r io.Reader) io.Reader {
	return &replaceReader{r: r, buf: make([]byte, 32*1024)}
}

type replaceReader struct {
	r       io.Reader
	buf     []byte
	replace map[string]string
}

func (r *replaceReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if err != nil {
		return n, err
	}
	log.Printf("before replace: %s\n", string(p[:n]))
	newData := replaceBuf(p[:n]) //todo: 如果replace之后变大就比较麻烦了
	log.Printf("after replace: %s\n\n", string(newData))

	// 将处理后的内容拷贝到输出缓冲区
	copied := copy(p, newData)
	return copied, nil
}

func replaceBuf(buf []byte) []byte {
	for k, v := range replaces {
		buf = bytes.ReplaceAll(buf, []byte(k), []byte(v))
	}
	return buf
}

type replaceMap map[string]string

func (r *replaceMap) String() string {
	return fmt.Sprintf("%v", *r)
}

func (r *replaceMap) Set(s string) error {
	ss := strings.Split(s, "->")
	if len(ss) != 2 {
		return errors.New("value must be like aaa->bbb")
	}
	if *r == nil {
		*r = map[string]string{}
	}
	(*r)[ss[0]] = ss[1]
	return nil
}
