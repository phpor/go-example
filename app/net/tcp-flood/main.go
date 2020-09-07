package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	addr := flag.String("addr", "localhost:80", "host:port")
	concurrent := flag.Int("concurrent", 500, "concurrent")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Author: phpor <junjie.li@beebank.com>\nVersion: 0.1.0\nUsage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	tcpAddress, err := net.ResolveTCPAddr("tcp4", *addr)
	if err != nil {
		fmt.Println("addr error")
		return
	}
	i := 0
	var timeUse int
	arrConn := map[int]*net.TCPConn{}
	for {
		if conn, ok := arrConn[i]; ok {
			_ = conn.Close()
		}
		arrConn[i], err = net.DialTCP("tcp", nil, tcpAddress)

		if err != nil {
			fmt.Printf("   fail:%s%12d %12dms :%s\n", time.Now().Format(time.UnixDate), i, timeUse, err.Error())
			continue
		}
		i++
		if i >= *concurrent {
			i = 0
		}
	}
}
