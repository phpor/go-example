// 这里有太多好玩的东西
// 参考资料： http://hit9.org/post/2013-11-17-19-20.html
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ch := make(chan int)
	go loop(ch)
	go loop(ch)
	runtime.Gosched()	// 如果这里没有这个，将不会有任何输出

	//<-ch
	//<-ch
}

func loop(ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		//	runtime.Gosched()
	}
		runtime.Gosched()

	ch <- 0
}
