package main

import "sync"

func main() {
	concurrentRun(func() {
		println(11)
		println(12)
	}, func() {
		println(21)
		println(22)
	}, func() {
		println(31)
		println(32)
	})
}

func concurrentRun(funcs ...func()) {
	wg := sync.WaitGroup{}
	for _, f := range funcs {
		wg.Add(1)
		go func(f func()) {
			f()
			wg.Done()
		}(f)
	}
	wg.Wait()
}
