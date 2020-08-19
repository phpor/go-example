package worker

import (
	"testing"
	"time"
)

func TestNewWorker(t *testing.T) {
	w := NewWorker(NewSimpleProducer(func(ch chan interface{}, isShouldStop func() bool) {
		i := 0
		for {
			if isShouldStop() {
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
	go func() {
		time.Sleep(5 * time.Second)
		w.Stop()
	}()
	w.Start()
}
