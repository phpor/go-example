// 参考：
// http://xiezhenye.com/2011/11/go-%E8%AF%AD%E8%A8%80%E5%B9%B6%E5%8F%91%E6%9C%BA%E5%88%B6-goroutine-%E5%88%9D%E6%8E%A2.html
// http://hit9.org/post/2013-11-17-14-07.html
// http://concur.rspace.googlecode.com/hg/talk/concur.html	并行 vs 并发
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	test2()
}

func test2() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		tmp := make(map[string]int, 0)
		tmp["key"] = i
		wg.Add(1)
		go func() {
			fmt.Printf("%v\n", tmp)
			wg.Done()
		}()
	}
	wg.Wait()
}

func test1() {
	ch := make(chan int)
	task("A", ch)
	task("B", ch)
	fmt.Printf("begin\n")
	<-ch
	<-ch
}
func task(name string, ch chan int) {
	go func() {
		i := 1
		for {
			fmt.Printf("%s %d\n", name, i)
			i++
			time.Sleep(time.Second * 1)
		}
		ch <- 1
	}()
}
