package worker

import (
	"os"
	"testing"
	"time"
)

func TestNewWorker(t *testing.T) {
	w := NewWorker(NewSimpleProducer(func(ch chan<- interface{}, isShouldStop func() bool) {
		i := 0
		for {
			if isShouldStop() {
				break
			}
			if i > 5 {
				break
			}
			i++
			ch <- i
			time.Sleep(time.Second)
		}
	}), NewSimpleConsumer(func(i interface{}) error {
		println(i.(int))
		return nil
	}))
	w.ConsumerNum = 3
	go w.LogStats(os.Stdout, time.Second)
	_ = w.Start()
}
