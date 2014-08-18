package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp4", "www.baidu.com:80")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	reader := bufio.NewReader(conn)
	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err != nil || n == 0 {
			break
		}
		fmt.Print(string(buf))
	}
}
