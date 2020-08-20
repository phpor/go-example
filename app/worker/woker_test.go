package worker

import (
	"fmt"
	"testing"
	"time"
)

func TestNewWorker(t *testing.T) {
	var w *Worker
	w = NewWorker(NewSimpleProducer(func(ch chan<- interface{}, isShouldStop func() bool) {
		i := 0
		for {
			if isShouldStop() {
				break
			}
			if i > 5 { // 生产者有能力随时优雅退出，生产者能知道自己生产了多少，也能通过w.Stats() 来知道消费了多少
				break
			}
			i++
			ch <- i
			time.Sleep(time.Second)
		}
	}), NewSimpleConsumer(func(i interface{}) error {
		println(i.(int))
		// 如果这里想退出，则可以直接调用w.Stop()
		w.Stop()
		return nil
	}))
	w.ConsumerNum = 3
	go w.LogStats(time.Second, func() {
		fmt.Printf("%s\n", w.stats.String())
	})
	_ = w.Start()
}
