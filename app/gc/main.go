package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

func main() {
	test2()
}

func test1() {
	debug.SetGCPercent(-1) // disable gc
	go func() {
		eatmem := func() []byte {
			s := make([]byte, 1024)
			s[1023] = 0x80
			//			println(s)
			return s
		}
		for {
			eatmem()
			time.Sleep(1 * time.Millisecond)
		}
	}()
	for {
		printMemStats(time.Now().String())
		time.Sleep(1 * time.Second)
	}
}

func printMemStats(label string) {
	memStats := &runtime.MemStats{}
	runtime.ReadMemStats(memStats)
	fmt.Println(label, memStats.Alloc) // 这里的Alloc就是heapAlloc，指的是虚拟内存，而不是物理内存（rss）
}

// 由于 <= 10MB 的对象会直接放在栈上，这里方便测试，使用大于10MB的对象
const size = 1024*1024*10 + 1

type bigStruct struct {
	A struct {
		Name [size]byte
	}
}

func test2() {
	a := bigStruct{}
	printMemStats("before fill array")
	for i := 0; i < size; i++ {
		a.A.Name[i] = 'a'
	}
	printMemStats("after fill array") // 这里统计的是虚拟内存，而不是rss，所以，填充前后都一样
	func() {
		printMemStats("before assign")
		// 下面两种赋值方式对内存的消耗是不同的

		// 这种方式需要堆上申请内存，使用完再gc
		//newArr := [size]byte{}
		//a.A.Name = newArr

		// 这种方式并没有引入新的变量，也没有占用更多的内存空间，可能和原来的值能直接盖掉有关系
		a.A.Name = [size]byte{}

		printMemStats("after assign")
	}()

	printMemStats("after call func")
	runtime.GC()
	printMemStats("after gc")
	a.A.Name[1] = 'a'
}
