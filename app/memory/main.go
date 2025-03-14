package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/google/gops/agent"
	"io"
	"log"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

var units = []string{" bytes", "KB", "MB", "GB", "TB", "PB"}

// 测试内存模型
// go1.19.4上没有这个环境变量也能释放内存给系统，ubuntu上测试的
// export GODEBUG=madvdontneed=1

func main() {
	c()
}

func c() {
	fmt.Printf("GODEBUG=%s\n", os.Getenv("GODEBUG"))

	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}

	// 申请很多内存
	ss := make([]byte, 1024*1024*50)
	for i := 0; i < len(ss); i++ {
		ss[i] = 'a'
	}

	var memBytes []byte

	usage := func() {
		fmt.Print(`
q: exit
e: show env
h: show help
s: stat
m: malloc
r: release memory
g: gc
c: debug.FreeOsMemory
`)
	}
	usage()
	r := bufio.NewReader(os.Stdin)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		cmd := string(bytes.TrimSpace(line))
		if cmd == "q" {
			break
		}

		switch cmd {
		case "g":
			runtime.GC()
		case "c":
			debug.FreeOSMemory()
		case "s":
			printStats()
		case "e":
			fmt.Printf("%s\n", os.Getenv("GODEBUG"))
		case "m":
			memBytes = append(memBytes, ss...)
		case "r":
			memBytes = nil
		case "h":
			usage()

		}
	}

}

func printStats() {
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	conn := os.Stdout
	fmt.Fprintf(conn, "alloc: %v\n", formatBytes(s.Alloc))
	fmt.Fprintf(conn, "total-alloc: %v\n", formatBytes(s.TotalAlloc))
	fmt.Fprintf(conn, "sys: %v\n", formatBytes(s.Sys))
	fmt.Fprintf(conn, "lookups: %v\n", s.Lookups)
	fmt.Fprintf(conn, "mallocs: %v\n", s.Mallocs)
	fmt.Fprintf(conn, "frees: %v\n", s.Frees)
	fmt.Fprintf(conn, "heap-alloc: %v\n", formatBytes(s.HeapAlloc))
	fmt.Fprintf(conn, "heap-sys: %v\n", formatBytes(s.HeapSys))
	fmt.Fprintf(conn, "heap-idle: %v\n", formatBytes(s.HeapIdle))
	fmt.Fprintf(conn, "heap-in-use: %v\n", formatBytes(s.HeapInuse))
	fmt.Fprintf(conn, "heap-released: %v\n", formatBytes(s.HeapReleased))
	fmt.Fprintf(conn, "heap-objects: %v\n", s.HeapObjects)
	fmt.Fprintf(conn, "stack-in-use: %v\n", formatBytes(s.StackInuse))
	fmt.Fprintf(conn, "stack-sys: %v\n", formatBytes(s.StackSys))
	fmt.Fprintf(conn, "stack-mspan-inuse: %v\n", formatBytes(s.MSpanInuse))
	fmt.Fprintf(conn, "stack-mspan-sys: %v\n", formatBytes(s.MSpanSys))
	fmt.Fprintf(conn, "stack-mcache-inuse: %v\n", formatBytes(s.MCacheInuse))
	fmt.Fprintf(conn, "stack-mcache-sys: %v\n", formatBytes(s.MCacheSys))
	fmt.Fprintf(conn, "other-sys: %v\n", formatBytes(s.OtherSys))
	fmt.Fprintf(conn, "gc-sys: %v\n", formatBytes(s.GCSys))
	fmt.Fprintf(conn, "next-gc: when heap-alloc >= %v\n", formatBytes(s.NextGC))
	lastGC := "-"
	if s.LastGC != 0 {
		lastGC = fmt.Sprint(time.Unix(0, int64(s.LastGC)))
	}
	fmt.Fprintf(conn, "last-gc: %v\n", lastGC)
	fmt.Fprintf(conn, "gc-pause-total: %v\n", time.Duration(s.PauseTotalNs))
	fmt.Fprintf(conn, "gc-pause: %v\n", s.PauseNs[(s.NumGC+255)%256])
	fmt.Fprintf(conn, "gc-pause-end: %v\n", s.PauseEnd[(s.NumGC+255)%256])
	fmt.Fprintf(conn, "num-gc: %v\n", s.NumGC)
	fmt.Fprintf(conn, "num-forced-gc: %v\n", s.NumForcedGC)
	fmt.Fprintf(conn, "gc-cpu-fraction: %v\n", s.GCCPUFraction)
	fmt.Fprintf(conn, "enable-gc: %v\n", s.EnableGC)
	fmt.Fprintf(conn, "debug-gc: %v\n", s.DebugGC)

}

func formatBytes(val uint64) string {
	var i int
	var target uint64
	for i = range units {
		target = 1 << uint(10*(i+1))
		if val < target {
			break
		}
	}
	if i > 0 {
		return fmt.Sprintf("%0.2f%s (%d bytes)", float64(val)/(float64(target)/1024), units[i], val)
	}
	return fmt.Sprintf("%d bytes", val)
}
