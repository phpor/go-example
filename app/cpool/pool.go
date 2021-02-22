package cpool

type pool struct {
	workerNum int
	ch        chan interface{}
}

func New(workerNum int) *pool {
	return &pool{
		workerNum: workerNum,
	}
}

func (p *pool) AddTask(task func(chan<- interface{})) (<-chan interface{}, error) {

	return nil, nil
}

func (p *pool) Start() {
	for i := 0; i < p.workerNum; i++ {
		go func() {
			for {
				func() {
					defer func() {

					}()
				}()
			}
		}()
	}
}
