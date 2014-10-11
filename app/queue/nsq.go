package main

//http://1.guotie.sinaapp.com/?p=533

import (
	"fmt"
	nsq "github.com/bitly/go-nsq"
)

var (
	wr *nsq.Writer
)

func openNsq(nsqaddr string) {
	if nsqaddr == "" {
		nsqaddr = "127.0.0.1:4150"
	}

}

func clicks(uid, nid int64) {
	s := fmt.Sprintf(`{"uid":%d, "nid":%d}`, uid, nid)
	wr.Publish("clicks", []byte(s))
}
