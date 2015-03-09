package main

import (
	redis "github.com/phpor/Go-Redis"
	"github.com/phpor/goutils/pools"
	"strconv"
	"strings"
	"time"
)

func main() {
	ch := make(chan int, 50)
	num := 10
	for i := 0; i < num; i++ {
		go work(chan <- int(ch))
	}
	for i := num; i > 0; i -- {
		<-ch
	}
}
func work(ch chan <- int) {
	p := getPool("127.0.0.1:6379")
	r, err := p.Get()
	if err != nil {
		panic(err)
	}
	rr := r.(redisRes).resource
	ss, err := rr.AllKeys()
	if err != nil {
		println(err)
	} else {
		for _, k := range ss {
			println(k)
		}
	}
	time.Sleep(1 * time.Second)
	p.Put(r)
	ch <- 1
}

type redisRes struct {
	resource redis.Client
}

func (res redisRes) Close() {
	res.resource.Quit()
}

var pool_map = map[string]*pools.ResourcePool{}

func getPool(addr string) *pools.ResourcePool {

	pool, exists := pool_map[addr]    //todo: 操作这个资源的时候需要加锁,否则，容易出现问题； 最好的办法是：初始化时把pool都先创建好
	if exists {
		return pool
	}
	println("create pool")
	pool = pools.NewResourcePool((pools.Factory)(func() (pools.Resource, error) {
			return createResource(addr)
		}), 3, 3, 10*time.Second)
	pool_map[addr] = pool
	return pool
}
func createResource(addr string) (pools.Resource, error) {
	s := strings.Split(addr, ":")
	port, _ := strconv.Atoi(s[1])
	spec := &redis.ConnectionSpec{}
	c, e := redis.NewSynchClientWithSpec(spec.Host(s[0]).Port(port))
	if e != nil {
		return nil, e
	}
	rr := redisRes{resource: c}
	return rr, nil
}
