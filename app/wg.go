package main

import (
	"sync"
	"time"
)

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time.Sleep(2 * time.Second)
		wg.Done()
	}()
	//wg.Done()
	wg.Wait()

}
