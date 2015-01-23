package main

import (
	"os"
	"os/signal"
	"fmt"
	"sync"
	"syscall"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	c := make(chan os.Signal, 10)
	signal.Notify(c, syscall.SIGINT, syscall.SIGHUP)
	go func(c <-chan os.Signal) { // 信号处理
		for {
			s := <-c
			fmt.Println(s)
			if s.String() == "interrupt" {


			}
			//		wg.Done()
		}
	}(c)
	fmt.Printf("pid %d\n", os.Getpid())
	wg.Wait()
}

