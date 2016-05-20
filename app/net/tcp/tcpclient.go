package main
import (
	"flag"
	"net"
	"fmt"
	"time"
	"os"
)

func main() {
	addr := flag.String("addr", "localhost:80", "host:port")
	timeout := flag.Int("timeout", 500, "timeout in ms")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Author: phpor <junjie.li@beebank.com>\nUsage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	tcpAddress, err := net.ResolveTCPAddr("tcp4", *addr)
	if err != nil {
		fmt.Println("addr error")
		return
	}
	i := 0
	var timeStart,timeEnd time.Time
	var timeUse int
	var conn *net.TCPConn
	for {
		i++
		timeStart = time.Now()
		conn, err = net.DialTCP("tcp", nil, tcpAddress)
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
		if i % 1000 == 0 {
			fmt.Printf("%s%12d\n", time.Now().Format(time.UnixDate), i)
		}
	}
}
