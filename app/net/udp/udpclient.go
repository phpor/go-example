package main

import (
	"fmt"
	"time"
	"net"
	"flag"
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
	data := "data"
	len := len(data)
	buf := make([]byte, 256)
	cnt := 0
	for {
		time.Sleep(time.Millisecond)
		n, err := conn.Write([]byte(data))
		if err != nil {
			fmt.Println("error writing data to server")
			fmt.Println(err)
			break
		}

		if n != len {
			println("send data fail")
			break
		}
		n, _, err = conn.ReadFromUDP(buf)	//能阻塞read吗？
		if err != nil {
			fmt.Println(err)
			break
		}
		if n != len {
			fmt.Printf("fail => received data len: %d\n", n)
			break
		}
		cnt += 1
		if cnt%10000 == 0 {
			fmt.Println(time.Now(), cnt)

		}
	}

}
