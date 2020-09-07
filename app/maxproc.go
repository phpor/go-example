package main

import (
	"os"
	"runtime"
	"strconv"
	"sync"
)

func main() {
	numCpu := int64(1)
	var err error
	if len(os.Args) > 1 {
		numCpu, err = strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			println(err)
			return
		}

	}
	println(runtime.NumCPU())
	wg := sync.WaitGroup{}
	for numCpu > 0 {
		numCpu--
		wg.Add(1)
		go func() {
			i := 1
			for {
				i++
			}
		}()
	}
	wg.Wait()
}
