package main

import (
	"flag"
	"fmt"
	"math"
	"net"
	"os"
	"sort"
	"time"
)

func main() {
	var (
		connTimes []float64

		addr  string
		count int
	)
	flag.StringVar(&addr, "addr", "", "addr")
	flag.IntVar(&count, "count", 10, "count")
	flag.Parse()
	for i := 0; i < count; i++ {
		connTime := measureConnectionTime(addr) // Replace with your IP and port
		connTimes = append(connTimes, connTime)
	}

	if len(connTimes) == 0 {
		fmt.Println("没有成功建立连接")
		return
	}

	sort.Float64s(connTimes)

	percentiles := []float64{10, 20, 30, 40, 50, 60, 70, 80, 90, 99, 99.9, 100}

	fmt.Println("连接时间分布:")
	for _, p := range percentiles {
		fmt.Printf("P%.1f: %.0f ms\n", p, calculatePercentile(connTimes, p))
	}
}

func measureConnectionTime(addr string) float64 {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "连接失败: %s", err)
		return 0
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("关闭连接失败:", err)
		}
	}(conn)

	duration := time.Since(start).Milliseconds()
	return float64(duration)
}

func calculatePercentile(data []float64, percentile float64) float64 {
	index := (percentile / 100.0) * float64(len(data)-1)
	if index < 0 {
		return data[0]
	}
	if index >= float64(len(data)) {
		return data[len(data)-1]
	}

	lowerIndex := int(math.Floor(index))
	upperIndex := lowerIndex + 1
	weight := index - float64(lowerIndex)

	if upperIndex >= len(data) {
		return data[lowerIndex]
	}

	return data[lowerIndex]*(1-weight) + data[upperIndex]*weight
}

func countValuesBelow(data []float64, value float64) int {
	count := 0
	for _, v := range data {
		if v <= value {
			count++
		}
	}
	return count
}
