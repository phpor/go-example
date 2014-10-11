package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

func main() {
	addr := "127.0.0.1:6379"
	go func() {
		for {
			if p, exists := pool[addr]; exists {
				println(p.ActiveCount())

			}
			time.Sleep(100 * time.Microsecond)
		}
	}()
	for i := 0; i < 10; i++ {
		go func() {

		}()
	}
}

var pool = map[string]*redis.Pool{}

func getConn(addr string) (redis.Conn, error) {
	pool, exists := pool[addr]
	if !exists {
		pool = redis.NewPool(func() (redis.Conn, error) {
				return redis.Dial("tcp", addr)
			}, 10)
	}

	if c := pool.Get(); c.Err() == nil {
		return c, nil
	} else {
		return c, c.Err()
	}
}


