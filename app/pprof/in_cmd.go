package main

import (
	"runtime/pprof"

	"net/http"
	"crypto/des"
	"os"
	"flag"
	"time"
)

func main() {
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
		println("use cpuprofile: " + *cpuprofile)
	}
	use_net()
}

func use_net() {
	for i := 0; i < 10; i++ {
		http.Get("http://www.sina.com.cn/")
		time.Sleep(1 * time.Second)
		//use_cpu()
	}
}

func use_cpu() {
	for i := 0; i < 10; i++ {
		c, err := des.NewCipher([]byte("12345678"))
		if err != nil {
			panic(err)
		}

		dst := make([]byte, des.BlockSize)
		c.Encrypt(dst, []byte("12345678"))
	}
}

