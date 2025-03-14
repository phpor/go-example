package main

import (
	"flag"
	"github.com/google/gops/agent"
	"log"
	"time"
)

func main() {

	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}
	cnt := 0

	flag.IntVar(&cnt, "count", 10, "go routine count")
	flag.Parse()

	for i := 0; i < cnt; i++ {
		go func() {
			time.Sleep(10 * time.Second)
		}()
	}
	time.Sleep(time.Hour)
}
