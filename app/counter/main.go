package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	d := flag.Int64("i", 1, "interval")
	flag.Parse()
	interval := time.Second * time.Duration(*d)
	ticker := time.NewTicker(interval)
	i := 0
	j := 0
	go func() {
		for {
			j = i
			<-ticker.C
			fmt.Printf("\rall: %15d dela: %12d  speed: %12d ", i, i-j, (i-j)/int(*d))
		}
	}()
	r := bufio.NewReader(os.Stdin)
	for {
		_, err := r.ReadBytes('\n')
		i++
		if err == io.EOF {
			break
		}
	}
}
