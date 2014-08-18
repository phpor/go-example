package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	var counter = struct { // 这是一个匿名的结构体，匿名的结构体只能局部使用，没法（或者说比较麻烦）通过参数传递哦~
		sync.RWMutex                // 这个东西似乎可以自己初始化
			m            map[string]int // 但是map不能自己初始化
		}{m: make(map[string]int)}

	counter.m["some_key"] = 0

	go func() {
		fmt.Println("start read goroutine...")
		for {
			counter.RLock()
			n := counter.m["some_key"]
			fmt.Println("some_key:", n)
			time.Sleep(time.Second * 1)
			counter.RUnlock()
		}
	}()
	go func() {
		fmt.Println("start write goroutine...")
		for {
			counter.Lock()
			counter.m["some_key"]++
			counter.Unlock()
		}
	}()
	// 从输出结果来看，两个 goroutine 是交替执行的，调度显得很公平
	<-ch    //这里如果不等一等的上，上面的goroutine就没机会执行
	// 让进程在这里等待的合适的方法是什么？
}
