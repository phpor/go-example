package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	var addr string
	flag.StringVar(&addr, "addr","", "eg: tcp://127.0.0.1:80")

	flag.Parse()

	network := "tcp"
	arr := strings.Split(addr, "://")
	if len(arr) > 1 {
		network = arr[0]
		addr = arr[1]
	}
	conn, err := net.Dial(network, addr)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for {
			_, err := io.Copy(os.Stdout, conn)
			if err != nil {
				break
			}
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for {
			_, err := io.Copy(conn, os.Stdin)
			if err != nil {
				break
			}
		}
		wg.Done()
	}()
	wg.Wait()
}
