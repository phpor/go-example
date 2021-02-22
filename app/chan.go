// 参考资料：
// 1. http://blog.sina.com.cn/s/blog_630c58cb01016j1u.html
// 2. http://www.kankanews.com/ICkengine/archives/99037.shtml
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	chan4()

}

func chan4() {
	c := make(chan int, 1)
	c <- 1
	<-c
	go func() { // 如果没有这个协程，就得报deadlock而终止程序
		i := 0
		for {
			i++
		}
	}()
	c <- 1
	c <- 1
}

func chan3() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	msg := make(chan int, 10)
	go func() {
		time.Sleep(1 * time.Second)
		for c := range msg {
			fmt.Printf("%d\n", c)
		}
		wg.Done()
	}()
	i := 0
	for {
		if i > 5 {
			break
		}
		msg <- i
		i++
	}
	// chan 必须 close，否则，chan中会读出空值，而且，空值不会触发for循环终止，或许会检测到死锁而退出
	// chan 被 close 时，不会影响其中未被读取的消息
	close(msg)

	wg.Wait()
}

type chanInterface interface {
}
type chanStruct struct {
	A string
}

func chan2() {
	ch := make(chan chanInterface, 2)
	ch <- chanStruct{A: "aaa"}
	c := <-ch
	println(c)
	close(ch)
	timer := time.NewTimer(10 * time.Second)
	c = <-ch
	ch = nil
	ch <- 1
	select {
	case ch <- 1:
	case <-timer.C:
		break
	}

	fmt.Printf("%v", c == nil)
}

func syntaxChan() {

}
func onewayChan2() {
	ch := make(chan int, 10)
	ch2 := (<-chan int)(ch) // ch2 只读 // 注意: (<-chan int)(ch)  不等于 <- chan int (ch)
	ch3 := (chan<- int)(ch) // ch3 只写
	ch3 <- 1
	// ch2 <- 2  // 只读的chan不能写
	fmt.Print(<-ch2)
}
func onewayChan() {
	ch1 := make(chan int, 10)
	ch2 := make(<-chan int, 10) // 直接 make一个只读或只写的chan是没有意义的
	ch3 := make(chan<- int, 10)
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
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Millisecond * 10)
		}
		//close(ch) // 这里关闭的只是 写 操作（似乎只能关闭写，也没有关闭读的必要），如果不关闭，最终的结果是 死锁
		// 问题： 这里是单个goroutine在写，如果多个goroutine在写一个chan，则谁来关闭？当然可以弄一个计数器，当然还要有锁，麻烦
	}()
	for {
		j, ok := <-ch // 注意： 不是  j <- ch ; 该语法用于判断chan是否已关闭
		if ok {
			fmt.Println(j)
		} else {
			break
		}
	}
}
