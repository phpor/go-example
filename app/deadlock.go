// 参考资料： http://hit9.org/post/2013-11-17-14-07.html
// 知识点：
// 1. 向无数据写入的空信道读取数据会引起死锁 （Go知道有没有人往信道写入数据）
// 2. 向无读取的无buffer（或buffer写满）的信道写数据会导致死锁
package main

import "time"

// 该main函数会导致死锁,如果有一个能干活的协程存在，哪怕该协程在sleep，也不会deadlock
func main() {
	ch := make(chan int)
	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()
	<-ch
}
