package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
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

	defer func() {
		_ = ln.Close()
	}()
	for {
		c, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting new connection - %v\n", err)
			continue
		}
		go func(c net.Conn) {
			clientHelloMsg, err := ReadClientHello(c)
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
			defer func() {
				_ = conn.Close()
			}()

			log.Printf("clientHellowMsg len: %d\n", len(clientHelloMsg.RawData))
			length, err := io.WriteString(conn, string(clientHelloMsg.RawData))
			log.Printf("1: write len: %d\n", length)
			if err != nil {
				log.Printf("1: error: %s\n", err.Error())
			}

			ch := make(chan int, 2)
			go func() {
				length, err := io.Copy(c, conn)
				log.Printf("2: read len: %d\n", length)
				if err != nil {
					log.Printf("2: error: %s\n", err.Error())
				}
				ch <- 1
			}()
			go func() {
				length, err := io.Copy(conn, c)
				log.Printf("3: write len: %d\n", length)
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
