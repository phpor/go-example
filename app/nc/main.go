package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
)

func main() {
	host := flag.String("host", "", "host")
	port := flag.String("port", "", "port")

	flag.Parse()
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", *host, *port))
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
