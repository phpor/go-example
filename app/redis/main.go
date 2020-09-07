package main

import (
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net"
	"time"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:6379", "")
	key := flag.String("key", "a", "")
	len := flag.Int("len", 10, "value length")
	flag.Parse()
	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	c := redis.NewConn(conn, 2*time.Second, 2*time.Second)
	data := make([]byte, *len)
	str, err := redis.String(c.Do("Set", *key, data))
	if err != nil {
		fmt.Printf("%v", err)
		return

	}
	fmt.Printf("%s\n", str)
}
