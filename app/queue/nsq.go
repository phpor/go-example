package main

//http://1.guotie.sinaapp.com/?p=533

import (
	"fmt"
	nsq "github.com/bitly/go-nsq"

	"time"
	"sync"
	"strconv"
)

var host = "127.0.0.1"

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		productor()
		wg.Done()
	}()
	go func() {
		consummer()
		wg.Done()
	}()
	wg.Wait()
}
func productor() {
	cfg := nsq.NewConfig()
	q,_ := nsq.NewProducer(host + ":5150", cfg)
	for i := 0; i < 100; i++ {
		q.Publish("test", []byte("msg" + strconv.Itoa(i)))
		time.Sleep(1 * time.Second)
	}
	q.Stop()
}
func consummer() {
	cfg := nsq.NewConfig()
	c, _ := nsq.NewConsumer("test", "c1", cfg)
	c.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error{ // 这个比定义一个实现handler interface的struct要快捷的多
		fmt.Println(msg.Body)
		return nil
	}))
	c.ConnectToNSQLookupd(host + ":4161")
	<-c.StopChan

}
