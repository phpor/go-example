package worker

import (
	"fmt"
	"io"
	"sync"
	"sync/atomic"
	"time"
)

type Producer interface {
	Init() error
	Do(chan<- interface{})
	Close()
}

type Consumer interface {
	Init() error
	Do(interface{}) error
	Close()
}

type Stats struct {
	CntConsumed        uint64
	CntConsumedSuccess uint64
	CntConsumedFail    uint64
	TimeStart          time.Time
	TimeEnd            time.Time
}

func (s *Stats) String() string {
	end := ""
	if !s.TimeEnd.IsZero() {
		end = fmt.Sprintf("\nElapsed: %v", s.TimeEnd.Sub(s.TimeStart).String())
	}
	return fmt.Sprintf("%s All:%d Success:%d Fail:%d%s", time.Now().Format("2006-01-02 15:04:05"), s.CntConsumed, s.CntConsumedSuccess, s.CntConsumedFail, end)
}

type Worker struct {
	Producer    Producer
	Consumer    Consumer
	ProducerNum int
	ConsumerNum int
	stopOnce    sync.Once
	closeCh     chan struct{}
	closedCh    chan struct{}
	stats       Stats
}

func NewWorker(producer Producer, consumer Consumer) *Worker {
	w := &Worker{Producer: producer, Consumer: consumer, ProducerNum: 1, ConsumerNum: 1}
	w.closeCh = make(chan struct{})
	w.closedCh = make(chan struct{})
	return w
}

func (w *Worker) Start() error {
	w.stats.TimeStart = time.Now()
	defer func() {
		w.stats.TimeEnd = time.Now()
		close(w.closedCh)
	}()
	ch := make(chan interface{}, w.ProducerNum+w.ConsumerNum)

	wgConsumer := sync.WaitGroup{}
	wgProducer := sync.WaitGroup{}
	i := w.ConsumerNum
	for i > 0 {
		wgConsumer.Add(1)
		err := w.Consumer.Init()
		if err != nil {
			return err
		}
		go func() {
			defer wgConsumer.Done()

			for msg := range ch {
				if err = w.Consumer.Do(msg); err == nil {
					atomic.AddUint64(&w.stats.CntConsumedSuccess, 1)
				} else {
					atomic.AddUint64(&w.stats.CntConsumedFail, 1)
				}
				atomic.AddUint64(&w.stats.CntConsumed, 1)
			}
			w.Consumer.Close()
		}()
		i--
	}
	i = w.ProducerNum
	for i > 0 {
		wgProducer.Add(1)
		err := w.Producer.Init()
		if err != nil {
			return err
		}
		go func() {
			defer wgProducer.Done()
			w.Producer.Do(ch)
		}()
		i--
	}
	go func() {
		<-w.closeCh
		w.Producer.Close()
	}()
	wgProducer.Wait()
	w.Producer.Close()

	close(ch)
	wgConsumer.Wait()
	return nil
}

func (w *Worker) Stop() {
	w.stopOnce.Do(func() {
		close(w.closeCh)
	})
}

func (w *Worker) Stats() *Stats {
	return &w.stats
}

func (w *Worker) LogStats(writer io.Writer, duration time.Duration) {
	ticker := time.NewTicker(duration)
	defer ticker.Stop()
	isStopped := false
	for {
		_, _ = fmt.Fprintf(writer, "%s\n", w.stats.String())
		if isStopped {
			break
		}
		select {
		case <-ticker.C:
		case <-w.closedCh:
			isStopped = true
		}
	}
}
