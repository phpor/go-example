package main

import (
	"runtime/debug"
	"time"
	"runtime"
	"fmt"
)

func main() {
	debug.SetGCPercent(-1) // disable gc
	go func() {
		eatmem := func() []byte {
			s := make([]byte, 1024)
			s[1023] = 0x80
			//			println(s)
			return s
		}
		for {
			eatmem()
			time.Sleep(1 * time.Millisecond)
		}
	}()
	memStats := &runtime.MemStats{}
	for {
		runtime.ReadMemStats(memStats)
		fmt.Println(memStats)
		time.Sleep(1 * time.Second)
	}
}

