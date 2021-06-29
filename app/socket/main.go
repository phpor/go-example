package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

func main() {
	addr := "127.0.0.1:6161"
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
		return
	}
	conn1, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	n, err := fmt.Fprintf(conn1, "client hello\n")
	if err != nil {
		panic(err)
	}
	if n == 0 {
		panic("应该能发送成功的")
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(conn net.Conn) {
		defer wg.Done()
		r := bufio.NewReader(conn)
		l, _, _ := r.ReadLine()

		println("s:", string(l))
		conn.Close()
	}(conn1)

	//ln.Close() // 如果这里 Close的话，下面的Accept就会失败，尽管上面有一个已经创建成功的连接
	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}
	clientHello, _, _ := bufio.NewReader(conn).ReadLine()
	fmt.Println("c:", string(clientHello))
	fmt.Fprintf(conn, "server hello\n")
	conn.Close()
	wg.Wait()

}
