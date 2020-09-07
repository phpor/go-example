package simple_worker

import "sync"

type ProducerFunc func(ch chan interface{}, isShouldStop func() bool) // 比定义一个Interface要简洁一些
type ConsumerFunc func(interface{})

type Worker struct {
	Producer        ProducerFunc // todo: 这个应该可以被通知结束
	ConsumerCreator func() (ConsumerFunc, error)
	ProducerNum     int
	ConsumerNum     int
	stopOnce        sync.Once
	closeCh         chan struct{}
}

func NewWorker(producer ProducerFunc, consumer func() (ConsumerFunc, error)) *Worker {
	return &Worker{Producer: producer, ConsumerCreator: consumer, ProducerNum: 1, ConsumerNum: 1}
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
		c, err := w.ConsumerCreator()
		if err != nil {
			return err
		}
		go func(c ConsumerFunc) {
			for msg := range ch {
				c(msg)
			}
			wgConsumer.Done()
		}(c)
		i--
	}
	i = w.ProducerNum
	for i > 0 {
		wgProducer.Add(1)
		go func() {
			w.Producer(ch, func() bool {
				select {
				case <-closeCh:
					return true
				default:

				}
				return false
			})
			wgProducer.Done()
		}()
		i--
	}
	wgProducer.Wait()
	close(ch)
	wgConsumer.Wait()
	return nil
}

func (w *Worker) Stop() {
	w.stopOnce.Do(func() {
		close(w.closeCh)
	})
}
