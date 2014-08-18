// 参考： http://xiezhenye.com/2011/11/go-%E8%AF%AD%E8%A8%80%E5%B9%B6%E5%8F%91%E6%9C%BA%E5%88%B6-goroutine-%E5%88%9D%E6%8E%A2.html
package main

import (
	"fmt"
	"time"
)

func main() {
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
