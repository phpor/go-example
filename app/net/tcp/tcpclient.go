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
	timeout := flag.Int("timeout", 500, "timeout in ms")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Author: phpor <junjie.li@beebank.com>\nVersion: 0.1.0\nUsage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	i := 0
	var timeStart, timeEnd time.Time
	var timeUse int
	dial := net.Dialer{Timeout: time.Millisecond * time.Duration(*timeout)}
	for {
		i++
		timeStart = time.Now()
		conn, err := dial.Dial("tcp", *addr)
		println(err.Error())
		timeEnd = time.Now()
		timeUse = int(timeEnd.Sub(timeStart).Nanoseconds() / 1e3)
		if timeUse > *timeout {
			fmt.Printf("timeout: %s%12d %12dms\n", time.Now().Format(time.UnixDate), i, timeUse)
			break
		}
		if err != nil {
			fmt.Printf("   fail:%s%12d %12dms :%s\n", time.Now().Format(time.UnixDate), i, timeUse, err.Error())
			break
		}
		conn.Close()
		if i%1000 == 0 {
			fmt.Printf("%s%12d\n", time.Now().Format(time.UnixDate), i)
		}
	}
}
