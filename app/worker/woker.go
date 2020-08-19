package worker

import (
	"sync"
)

type Producer interface {
	Init() error
	Do(chan interface{})
	Close()
}

type Consumer interface {
	Init() error
	Do(interface{}) error
	Close()
}

type Worker struct {
	Producer    Producer
	Consumer    Consumer
	ProducerNum int
	ConsumerNum int
	stopOnce    sync.Once
	closeCh     chan struct{}
}

func NewWorker(producer Producer, consumer Consumer) *Worker {
	return &Worker{Producer: producer, Consumer: consumer, ProducerNum: 1, ConsumerNum: 1}
}

func (w *Worker) Start() error {
	ch := make(chan interface{}, w.ProducerNum+w.ConsumerNum)
	closeCh := make(chan struct{})
	w.closeCh = closeCh

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
				_ = w.Consumer.Do(msg)
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
		<-closeCh
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
