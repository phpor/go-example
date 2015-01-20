package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"sort"
	"sync"
	"time"
)

var (
	verbos    bool
	timeStart time.Time
)

func main() {
	addr := flag.String("addr", "localhost:12345", "host:port")
	count := flag.Int("count", 10000, "count to request")
	goroutineNum := flag.Int("grn", 10, "num of goroutine")
	data := flag.String("data", "", "data to send")
	ptrVerbos := flag.Bool("v", false, "print debug info")

	flag.Parse()
	verbos = *ptrVerbos
	server := *addr
	udpAddress, err := net.ResolveUDPAddr("udp4", server)
	if err != nil {
		fmt.Println("addr error")
		return
	}

	timeSlice := make([]int, *count)

	buf := make([]byte, 256)
	cntAll := 0
	cntOk := 0
	cntInvalid := 0

	expectResult := []byte(*data)
	flag := byte(1)
	num := *goroutineNum

	timeStart = time.Now()
	var wg sync.WaitGroup
	wg.Add(num)
	j := *count
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			conn, err := net.DialUDP("udp", nil, udpAddress)
			if err != nil {
				fmt.Printf("Could not resolve udp address or connect to it  on %s :", server)
				fmt.Println(err)
				return
			}
			defer conn.Close()

			for {
				j--
				if j < 0 {
					break
				}
				header := make([]byte, 8)
				binary.BigEndian.PutUint64(header, uint64(cntAll))
				mydata := append(header, flag)
				mydata = append(mydata, []byte(*data)...)
				length := len(mydata)
				conn.SetWriteDeadline(time.Now().Add(100 * time.Millisecond)) // 设置 超时为 100ms
				timeStart := time.Now()
				n, err := conn.Write(mydata)
				if err != nil {
					debug("error writing data to server")
					debug(err.Error())
					break
				}

				if n != length {
					debug(fmt.Sprintf("send data fail %d != %d", n, length))
					break
				}

				ok := false
				for {
					conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond)) // 设置 读超时为 100ms
					n, _, err = conn.ReadFromUDP(buf)                            //能阻塞read吗？
					if err != nil {
						debug(err.Error())
						break
					}
					if !bytes.Equal(buf[:8], header) {
						continue
					}
					// buf[8] 为1字节的flag
					if bytes.Equal(buf[9:n], expectResult) {
						ok = true
					} else {
						debug(fmt.Sprintf("fail: %s", string(buf[8:n])))
						cntInvalid++
					}
					break
				}
				if ok {
					cntOk++
				}
				timeEnd := time.Now()
				if cntAll >= len(timeSlice) {
					break
				}
				//              fmt.Println(timeEnd.Sub(timeStart).Nanoseconds())
				timeSlice[cntAll] = int(timeEnd.Sub(timeStart).Nanoseconds() / 1e3)

				cntAll++
				if cntAll%10000 == 0 {
					fmt.Printf("cntAll: %d   cntOk: %d cntInvalid: %d\n", cntAll, cntOk, cntInvalid)
				}

			}
		}()

	}
	wg.Wait()
	fmt.Printf("cntAll: %d   cntOk: %d cntInvalid: %d\n", cntAll, cntOk, cntInvalid)
	report(timeSlice)

}

// report show the benchmark result
func report(timeSlice []int) {
	timeEnd := time.Now()
	timeUse := int(timeEnd.Sub(timeStart).Nanoseconds() / 1e6) //ms
	num := len(timeSlice)
	fmt.Printf("request total: %d timeuse: %d ms  r/s: %d   us/r: %d\n", len(timeSlice), timeUse, num/timeUse*1000, timeUse*1000/num)
	sort.Ints(timeSlice)
	//      fmt.Println(len(timeSlice), timeSlice)
	l := len(timeSlice)
	for _, i := range []int{5, 10, 20, 50, 80, 90, 99, 100} {
		index := l * i / 100
		if index >= l {
			index = l - 1
		}
		fmt.Printf("%d%%\t%dus\n", i, timeSlice[index])
	}
}

func debug(str string) {
	if verbos {
		fmt.Println(str)
	}

}
