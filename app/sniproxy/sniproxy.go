package main

import (
	"flag"
	"net"
	"fmt"
	"log"
	"io"
	"github.com/phpor/goexample/app/tls"
)

func main() {
	addr := flag.String("addr", ":443", "host:port")
	flag.Parse()

	ln, err := net.Listen("tcp4", *addr)

	if err != nil {
		fmt.Println("error listening on tcp port ", *addr)
		fmt.Println(err)
		return
	}

	defer ln.Close()
	for {
		c, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting new connection - %v", err)
			continue
		}
		go func(c net.Conn) {
			clientHelloMsg, err := tls.ReadClientHello(c)
			if err != nil {
				log.Printf("Error accepting new connection - %v", err)
			}
			if clientHelloMsg.ServerName == "" {
				log.Printf("Cannot support non-SNI enabled clients")
				return
			}
			dst := clientHelloMsg.ServerName + ":443"
			conn, err := net.Dial("tcp", dst)
			if err != nil {
				log.Printf("connect %s fail", dst)
				return
			}
			defer conn.Close()
			io.WriteString(conn, string(clientHelloMsg.Raw))
			ch := make(chan int, 2)
			go func() {
				io.Copy(c, conn)
				ch <- 1
			}()
			go func() {
				io.Copy(conn, c)
				ch <- 1
			}()
			<-ch
			<-ch

		}(c)
	}
}
