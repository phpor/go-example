package main

import (
	"flag"
	"net"
	"fmt"
	"os"
	"time"
)


type job struct {
	addr *net.UDPAddr
	data []byte
}

var (
	cntIn int = 0
	cntOut int = 0
	cntInErr int = 0
	cntOutErr int = 0
)
func main() {
	addr := flag.String("addr", "localhost:12345", "host:port")
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

	var jobChan = make(chan *job, 50 * 10000)
	createProductor(conn, jobChan, 1)
	createWorker(conn, jobChan, 8)
	showStatus()
}

func showStatus() {
	for{
		fmt.Printf("cntIn: %d  cntOut: %d  cntInerr: %d  cntOutErr: %d cntBlock: %d\n", cntIn, cntOut, cntInErr, cntOutErr, cntIn - cntOut)
		time.Sleep(2 * time.Second)
	}
	
}
func createProductor(conn *net.UDPConn, jobChan chan *job, num int) {
	for num > 0 {
		go func(conn *net.UDPConn, jobChan chan *job) {
			for {
				buf := make([]byte, 256)	//这里要给一个合适的大小，不必太大，对于异常数据可以扔掉
				n, address, err := conn.ReadFromUDP(buf)

				if err != nil {
					fmt.Fprintf(os.Stderr, "read fail %s \n" ,err)
					cntInErr += 1
					continue
				}
				if address == nil {
					cntInErr += 1
					fmt.Fprintf(os.Stderr, "address fail %v \n" ,address)
					continue
				}
				if n == 0 {
					fmt.Fprintf(os.Stderr, "data empty \n" ,err)
					cntInErr += 1
					continue
				}
				buf = buf[:n]
				jobChan <- &job{address, buf}
				cntIn += 1
			}
		}(conn, jobChan)
		num -= 1
	}
	
}
func createWorker(conn *net.UDPConn, jobChan chan *job, num int) {
	for num > 0 {
		go func(conn *net.UDPConn, jobChan chan *job) {
			for {
				job := <-jobChan
				_, err := conn.WriteToUDP(job.data, job.addr)
				if err != nil {
					cntOutErr += 1
					fmt.Fprintf(os.Stderr, "write fail %s \n" ,err)
					continue
				}
				cntOut += 1
			}
		}(conn, jobChan)
		num -= 1
	}
}
