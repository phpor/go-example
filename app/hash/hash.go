package main

import (
	"time"
)

func main() {
	pool := map[string]int{
		"ip1:port": 2,
		"ip2:port": 2,
		"ip3:port": 2,
		"ip4:port": 2,
	}
	h := NewHashRing(1024)
	h.AddNodes(pool)
	start := time.Now()
	s := ""
	for i := 0; i < 100000; i++ {
		s = h.GetNode("key1")
	}
	println(time.Now().Sub(start).String())
	println(pool[s])
}
