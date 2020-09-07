package main

import (
	"bufio"
	"flag"
	"github.com/phpor/go-example/app/worker"
	"io/ioutil"
	"os"
	"sort"
)

// 多核并发的外排序

func main() {
	filename := flag.String("file", "", "filename to sort")
	workerNum := flag.Int("worker-num", 1, "worker num")
	flag.Parse()
	if err := mySort(*filename, *workerNum); err != nil {
		panic(err)
	}

}

func mySort(filename string, workerNum int) error {
	fp, err := os.Open(filename)
	if err != nil {
		return err
	}

	var tmpFiles []*os.File
	getTmpFile := func() (*os.File, error) {
		f, err := ioutil.TempFile(os.TempDir(), "sort.*")
		if err != nil {
			return nil, err
		}
		tmpFiles = append(tmpFiles, f)
		return f, nil
	}
	w := worker.NewWorker(func(ch chan interface{}, isShouldStop func() bool) {
		var line []byte
		var err error
		r := bufio.NewReader(fp)

		for {
			line, err = r.ReadBytes('\n')
			if len(line) > 0 {
				ch <- line
			}
			if err != nil {
				break
			}
		}
	}, func() (worker.ConsumerFunc, error) {
		tmpFile, err := getTmpFile()
		if err != nil {
			return nil, err
		}
		bufLen := 1024
		sortBuf := make([]string, bufLen)
		i := 0
		return func(line interface{}) {
			sortBuf[i] = string(line.([]byte))
			if i == bufLen { // 这里无法知道处理到最后一个了
				sort.Strings(sortBuf)
			}
		}, nil
	})
	w.ConsumerNum = workerNum
	w.Start()

}
