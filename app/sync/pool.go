package main

import (
	"fmt"
	"runtime"
	"sync"
)

func pool() error {
	obj := 0
	p := sync.Pool{New: func() interface{} {
		obj++
		fmt.Printf("new %d\n", obj)
		return obj
	}}
	i := 0
	for i < 1000 {
		p.Put(i)
		i++
	}
	runtime.GC() // go1.11一次gc就可以把pool中的对象清空
	runtime.GC() // go1.14可以通过两次gc把pool中的对象清空
	w := func() {
		i := p.Get().(int)
		println(i)
		//p.Put(i)
	}
	w()
	w()
	w()

	return nil
}

func init() {
	cmds["pool"] = pool
}
