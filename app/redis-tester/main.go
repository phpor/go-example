package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net"
	"time"
)

func main() {
	addr := flag.String("addr", "redis:6379", "")
	key := flag.String("key", "a", "")
	action := flag.String("action", "set-get", "")
	count := flag.Int("count", 10, "repeat times")
	timeout := flag.Int("timeout", 1000, "读写超时时间 单位: ms")
	length := flag.Int("length", 10, "value length")
	flag.Parse()
	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	timeout1 := time.Millisecond * time.Duration(*timeout)
	c := redis.NewConn(conn, timeout1, timeout1)
	data := bytes.Repeat([]byte("a"), *length)
	i := *count
	var start time.Time
	end := func(str string, err error) {
		fmt.Printf("%s %d us\n", str, int64(time.Now().Sub(start))/1e3)
		if err != nil { // 出错重连
			c = redis.NewConn(conn, timeout1, timeout1)
		}
	}
	for {
		if *count > 0 {
			i--
			if i <= 0 {
				break
			}
		}
		start = time.Now()
		fmt.Printf("%s ", start.Format("2006-01-02 15:04:05"))
		if *action == "set-get" {
			str, err := redis.String(c.Do("Set", *key, data))
			if err != nil {
				end(fmt.Sprintf("%v", err), err)
				continue
			}
			end(fmt.Sprintf("set %s\t", str), nil)
		}
		res, err := redis.String(c.Do("Get", *key))
		if err != nil {
			end(fmt.Sprintf("get %v", err), err)
			continue
		}
		end(fmt.Sprintf("get %s", res), nil)
	}
}
