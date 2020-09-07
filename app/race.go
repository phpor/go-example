package main

import "sync"

// bool 值的并发读写是不需要加锁的； 整型的赋值也是不需要加锁的； 但是  a++ 之类的操作是有并发问题的，但是atomic的原子性处理要比加锁的效率高很多
func main() {
	b := true
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for {
			if b {
				continue
			}
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for {
			if b {
				continue
			}
		}
		wg.Done()
	}()
	for {
		b = !b
	}

}
