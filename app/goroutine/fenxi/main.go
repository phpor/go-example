package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
)

type item struct {
	name  string
	count int
}

type runnerState struct {
	name        string
	totalWorker int
	jobRunning  int
	workerState *workerState
}

type workerState struct {
	chanSend  int // chan send 状态，正在努力把job往chan里面发送
	waitGroup int // 执行中，等待依赖的job完成
	waitJob   int // 在select状态，等待job
	runnable  int // 可运行,有job 可以运行，但是还没有执行
	unknown   int // 未知状态
}

func (workerState *workerState) String() string {
	s := fmt.Sprintf("chan send: %d\nwait group: %d\nwait job:%d\nrunnable: %d\nunknown: %d", workerState.chanSend, workerState.waitGroup, workerState.waitJob, workerState.runnable, workerState.unknown)
	return s
}

func (rs *runnerState) String() string {
	s := fmt.Sprintf("name: %s\ntotal worker: %d\njob running: %d\n%s\n", rs.name, rs.totalWorker, rs.jobRunning, rs.workerState.String())
	return s
}

func (workerState *workerState) update(l0 []byte) {
	if bytes.Contains(l0, []byte("chan send")) {
		workerState.chanSend++
	} else if bytes.Contains(l0, []byte("semacquire")) {
		workerState.waitGroup++
	} else if bytes.Contains(l0, []byte("select")) {
		workerState.waitJob++
		//println(string(gr), "\n")
	} else if bytes.Contains(l0, []byte("runnable")) {
		workerState.runnable++
	} else {
		workerState.unknown++
	}
}

func getWorkerID(s []byte) string {
	ss := bytes.Split(s, []byte("job.(*Runner).work("))
	if len(ss) != 2 {
		return ""
	}
	ss = bytes.Split(ss[1], []byte(")"))
	return string(ss[0])
}

func main() {

	f, err := os.Open("/Users/junjie2/data1/github.com/phpor/go-example/app/goroutine/fenxi/stack.txt")
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	bs, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	rs := map[string]*runnerState{}
	gs := map[string]int{}
	grs := bytes.Split(bs, []byte("\n\n"))
	for _, gr := range grs {
		ll := bytes.Split(gr, []byte("\n"))
		l0 := ll[0]
		l1 := ll[1]
		if bytes.Contains(gr, []byte("job.(*Runner).work(")) {
			rid := getWorkerID(gr)
			_, ok := rs[rid]
			if !ok {
				rs[rid] = &runnerState{name: rid, workerState: &workerState{}}
			}
			rs[rid].totalWorker++
			rs[rid].workerState.update(l0)
			if bytes.Contains(gr, []byte("job/runner.go:244")) {
				rs[rid].jobRunning++
			}
		}
		i := bytes.LastIndexByte(l1, '(')
		if i == -1 {
			println(string(l1))
			continue
		}
		gsName := string(l1[:i])
		gs[gsName]++
	}
	var ss []item
	for k, v := range gs {
		//println(k, v)
		ss = append(ss, item{name: k, count: v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].count > ss[j].count
	})
	statistics := os.Getenv("result")
	if statistics == "true" {
		for _, s := range ss {
			println(s.name, s.count)
		}

	}
	for _, v := range rs {
		println(v.String())
		println()
	}
}
