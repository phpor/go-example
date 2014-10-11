package main

import (
	log "github.com/cihub/seelog"
	"time"
	"os"
	"fmt"
)

func main() {
	logger, err := log.LoggerFromConfigAsString("<seelog type=\"asynctimer\" asyncinterval=\"500000000\"/>")
	checkFail(err)
	log.ReplaceLogger(logger)
	defer log.Flush()
	println("start")

	log.Info("Hello from Seelog!")

	time.Sleep(time.Second * 10)
}

func checkFail(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

/**
1. seelog 支持同步和异步日志方式
2. seelog 的异步是通过list而不是chan来实现的
3. seelog 的异步有两种方式：
	3.1 定期flush（AsyncTimerLogger）： 不太实时
	3.2 通知flush（asyncLoopLogger）： 相对来说，更加实时
	3.3 不管何种方式，当消息达到一个值（10000）时，都会立即flush的
 */
