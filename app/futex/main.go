package main

import (
	"time"
)

func main() {
	//m := &sync.Map{}
	i := 0
	go func() {
		for {
			i++
			//m.Store("a", i)
		}
	}()
	go func() {
		for {
			i--
			//v, _ := m.Load("a")
			//m.Store("a", v)

		}
	}()
	time.Sleep(time.Hour)
}
