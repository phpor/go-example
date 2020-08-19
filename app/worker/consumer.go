package worker

type SimpleConsumer struct {
	f ConsumerFunc
}

type ConsumerFunc func(interface{}) error

func NewSimpleConsumer(f ConsumerFunc) *SimpleConsumer {
	sc := &SimpleConsumer{
		f: f,
	}
	return sc
}

func (sc *SimpleConsumer) Init() error {
	return nil
}

func (sc *SimpleConsumer) Do(msg interface{}) error {
	return sc.f(msg)
}

func (sc *SimpleConsumer) Close() {}
