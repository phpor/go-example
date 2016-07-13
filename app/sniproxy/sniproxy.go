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
		fmt.Println("error listening on tcp port \n", *addr)
		fmt.Println(err)
		return
	}

	defer ln.Close()
	for {
		c, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting new connection - %v\n", err)
			continue
		}
		go func(c net.Conn) {
			clientHelloMsg, err := tls.ReadClientHello(c)
			if err != nil {
				log.Printf("Error accepting new connection - %v\n", err)
			}
			if clientHelloMsg.ServerName == "" {
				log.Printf("Cannot support non-SNI enabled clients\n")
				return
			}
			dst := clientHelloMsg.ServerName + ":443"
			//dst := "220.181.112.244:443"
			conn, err := net.Dial("tcp", dst)
			if err != nil {
				log.Printf("connect %s fail\n", dst)
				return
			}
			log.Printf("connected to %s\n", dst)
			defer conn.Close()

			log.Printf("clientHellowMsg len: %d\n", len(clientHelloMsg.RawData))
			len, err := io.WriteString(conn, string(clientHelloMsg.RawData))
			log.Printf("1: write len: %d\n", len)
			if err != nil {
				log.Printf("1: error: %s\n", err.Error())
			}

			ch := make(chan int, 2)
			go func() {
				len ,err := io.Copy(c, conn)
				log.Printf("2: read len: %d\n", len)
				if err != nil {
					log.Printf("2: error: %s\n", err.Error())
				}
				ch <- 1
			}()
			go func() {
				len, err := io.Copy(conn, c)
				log.Printf("3: write len: %d\n", len)
				if err != nil {
					log.Printf("3: error: %s\n", err.Error())
				}
				ch <- 1
			}()
			<-ch
			<-ch

		}(c)
	}
}
