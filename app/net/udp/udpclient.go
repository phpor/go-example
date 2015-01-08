package main

import (
	"fmt"
	"time"
	"net"
	"flag"
	"bytes"
)

func main() {
	addr := flag.String("addr", "localhost:12345", "host:port")
	flag.Parse()
	server := *addr
	udpAddress, err := net.ResolveUDPAddr("udp4", server)
	conn, err := net.DialUDP("udp",nil, udpAddress)

	if err != nil {
		fmt.Println("Could not resolve udp address or connect to it  on " , server)
		fmt.Println(err)
		return
	}

	defer conn.Close()

	buf := make([]byte, 256)
	cntAll := 0
	cntOk := 0

	for {
		mydata, _ := time.Now().GobEncode() // 用序号也行
		length := len(mydata)
		conn.SetWriteDeadline(time.Now().Add(100 * time.Millisecond)) // 设置 超时为 100ms
		n, err := conn.Write(mydata)
		if err != nil {
			fmt.Println("error writing data to server")
			fmt.Println(err)
			break
		}

		if n != length {
			fmt.Printf("send data fail %d != %d", n, length)
			break
		}

		cntAll += 1
		ok := false
		for {
			conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond)) // 设置 读超时为 100ms
			n, _, err = conn.ReadFromUDP(buf)    //能阻塞read吗？
			if err != nil {
				fmt.Println(err)
				break
			}
			if n != length {
				fmt.Printf("fail => received data len: %d\n", n)
				break
			}
			if cntAll%10000 == 0 {
				fmt.Println(time.Now(), cntAll)
			}

			if !bytes.Equal(buf[:n], mydata) {
				continue
			} else {
				ok = true
				break
			}
		}
		if ok {
			cntOk += 1
		}
		if cntAll%10000 == 0 {
			fmt.Printf("cntAll: %d   cntOk: %d\n", cntAll, cntOk)

		}
	}

}
