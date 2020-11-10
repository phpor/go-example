package main

import (
	"container/ring"
	"fmt"
)

// Ring 的正确使用，使用ring.New(100) 创建一个ring，然后，通过r.Next()的方法取下一个节点，并且赋值
// 查看的时候，可以通过Do的方法查看所有节点的Value
func main() {
	Usage()

}

func test1() {
	r := ring.New(5)
	printr(r)
	i := 0
	for {
		r.Link(&ring.Ring{Value: i})
		i++
		printr(r)
		if i > 12 {
			break
		}
	}
}

type MyRing struct { // 这就是一个ring buffer了
	r    *ring.Ring
	head *ring.Ring
	tail *ring.Ring
}

func New(n int) *MyRing {
	r := ring.New(n)
	return &MyRing{r: r, head: r, tail: r}
}

func (mr *MyRing) Push(v int) {
	mr.tail.Value = v
	mr.tail = mr.tail.Next()
}
func (mr *MyRing) Dump() {
	printr(mr.head)
}

func Usage() {
	r := New(5)
	r.Push(1)
	r.Push(2)
	r.Push(3)
	r.Dump()
	r.Push(4)
	r.Push(5)
	r.Push(6)
	r.Push(7)
	r.Dump()
}

func printr(r *ring.Ring) {
	fmt.Printf("r = [ ")
	r.Do(func(v interface{}) {
		fmt.Printf("%v ", v)
	})
	fmt.Printf("]\n")
}
