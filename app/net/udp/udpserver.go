package main

import (
	"flag"
	"net"
	"fmt"
	"os"
	"time"
	_ "net/http/pprof"
	"net/http"
)

var (
	cntIn int
	cntOut int
	cntInErr int
	cntOutErr int
)
func main() {
	addr := flag.String("addr", ":12345", "host:port")
	flag.Parse()
	udpAddress, err := net.ResolveUDPAddr("udp4", *addr)

	if err != nil {
		fmt.Println("error resolving UDP address on ", *addr)
		fmt.Println(err)
		return
	}

	conn ,err := net.ListenUDP("udp", udpAddress)

	if err != nil {
		fmt.Println("error listening on UDP port ", *addr)
		fmt.Println(err)
		return
	}

	defer conn.Close()

	go run(conn)
	go func() {
		http.ListenAndServe(":7777", nil)
	}()
	showStatus()
}

func showStatus() {
	for{
		fmt.Printf("cntIn: %d  cntOut: %d  cntInerr: %d  cntOutErr: %d cntBlock: %d\n", cntIn, cntOut, cntInErr, cntOutErr, cntIn - cntOut)
		time.Sleep(2 * time.Second)
	}

}
func run(conn *net.UDPConn) {
	for {
		buf := make([]byte, 256)	//这里要给一个合适的大小，不必太大，对于异常数据可以扔掉
		n, address, err := conn.ReadFromUDP(buf)

		if err != nil {
			fmt.Fprintf(os.Stderr, "read fail %s \n" ,err)
			cntInErr++
			continue
		}
		if address == nil {
			cntInErr++
			fmt.Fprintf(os.Stderr, "address fail %v \n" ,address)
			continue
		}
		if n == 0 {
			fmt.Fprintf(os.Stderr, "data empty \n")
			cntInErr++
			continue
		}
		buf = buf[:n]
		go func(conn *net.UDPConn, address *net.UDPAddr, buf []byte) {
			doJob(conn, address, buf)
		}(conn, address, buf)
		cntIn++
	}

}
func doJob(conn *net.UDPConn, address *net.UDPAddr, buf []byte) {
	_, err = conn.WriteToUDP(buf, address)
	if err != nil {
		cntOutErr++
		fmt.Fprintf(os.Stderr, "write fail %s \n" ,err)
	}
	cntOut++
}
