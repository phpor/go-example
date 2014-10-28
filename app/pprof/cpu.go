package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

func main() {
	goroutineNum := flag.Int("goroutinenum", 1, "goroutine num")
	cpuNum := flag.Int("cpunum", 1, "cpu num to use")
	flag.Parse()

	fmt.Printf("use cpu num %d\n", *cpuNum)
	fmt.Printf("goroutine num %d\n", *goroutineNum)

	runtime.GOMAXPROCS(*cpuNum)
	i := *goroutineNum
	for ; i > 0; i-- {
		go func() {
			for {
				runtime.Gosched()
			}
		}()
	}
	fmt.Printf("goroutine num: %d\n", runtime.NumGoroutine())
	for {
		time.Sleep(100 * time.Second)
	}
}

