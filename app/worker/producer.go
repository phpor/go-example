package worker

import "sync"

type SimpleProducer struct {
	closeCh  chan struct{}
	f        ProducerFunc
	stopOnce sync.Once
}
type ProducerFunc func(ch chan interface{}, isShouldStop func() bool) // 比定义一个Interface要简洁一些

func NewSimpleProducer(f ProducerFunc) *SimpleProducer {
	sp := &SimpleProducer{
		closeCh: make(chan struct{}),
		f:       f,
	}
	return sp
}

func (sp *SimpleProducer) Init() error {
	return nil
}

func (sp *SimpleProducer) Do(ch chan interface{}) {
	sp.f(ch, func() bool { return sp.ShouldStop() })
}

func (sp *SimpleProducer) Close() {
	sp.stopOnce.Do(func() {
		close(sp.closeCh)
	})
}

func (sp *SimpleProducer) ShouldStop() bool {
	select {
	case <-sp.closeCh:
		return true
	default:

	}
	return false
}
