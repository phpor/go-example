package main

import (
	redis "github.com/phpor/Go-Redis"
	"fmt"
)

func main() {
	async()
}
func async() {
	spec := redis.DefaultSpec()
	r, err := redis.NewAsynchClientWithSpec(spec.Host("localhost").Port(6379))
	if err != nil {
		println(err)
		return
	}
	fbAllKeys, e := r.Get("a")
	fb2AllKeys, e := r.Get("a")
	b0, e, timedout := fbAllKeys.TryGet(0)
	if timedout { // 注意： 连接spec中的 reqChanCap 如果为0 的话，异步将不会出现，这里将不会good
		println("good")
		b, _ := fbAllKeys.Get()
		fmt.Println(b)
	} else {
		println("aaaaaaa")
		fmt.Println(b0)
	}
	b2, e := fb2AllKeys.Get()
	if e != nil {
		println(e)
		return
	}
	println("ok")
	fmt.Println(b2)

}
