package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	a := int64(0)

	wg := sync.WaitGroup{}
	wg.Add(2)
	add := func() {
		i := 0
		for {
			i++
			//a++
			atomic.AddInt64(&a, 1) // 这样是正确的，如果直接a++,结果就不正确，会丢掉很多很多
			if i >= 50000 {
				break
			}
		}
		wg.Done()
	}
	go add()
	go add()
	wg.Wait()
	println(a)
}
