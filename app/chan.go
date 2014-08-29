// 参考资料： http://blog.sina.com.cn/s/blog_630c58cb01016j1u.html
package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 10)
	ch2 := (<-chan int)(ch) // ch2 只读 // 注意: (<-chan int)(ch)  不等于 <- chan int (ch)
	ch3 := (chan <- int)(ch) // ch3 只写
	ch3 <- 1
	// ch2 <- 2  // 只读的chan不能写
	fmt.Print(<-ch2)

}

func onewayChan() {
	ch1 := make(chan int, 10)
	ch2 := make(<-chan int, 10)    // 直接 make一个只读或只写的chan是没有意义的
	ch3 := make(chan <- int, 10)
	ch1 <- 1
	//	ch2 <- 2
	//	ch3 <- 3
	//	fmt.Println(<-ch1, <-ch2, <-ch3)

	fmt.Printf("%T\n%T\n%T\n", ch1, ch2, ch3)

}
func chanchan(ch chan (chan int)) int {
	ch2 := make(chan int, 10)
	ch <- ch2
	return <-ch2
}
func first() {
	ch := make(chan int, 1)

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
			time.Sleep(time.Millisecond * 10)
		}
		close(ch) // 这里关闭的只是 写 操作（似乎只能关闭写，也没有关闭读的必要），如果不关闭，最终的结果是 死锁
		// 问题： 这里是单个goroutine在写，如果多个goroutine在写一个chan，则谁来关闭？当然可以弄一个计数器，当然还要有锁，麻烦
	}()
	for {
		j, ok := <-ch // 注意： 不是  j <- ch
		if ok {
			fmt.Println(j)
		} else {
			break
		}
	}
}
