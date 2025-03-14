package main

import (
	"flag"
	"fmt"
	"net"
	"sort"
	"time"
)

// Constants for the number of connections and timeout duration
const (
	timeout = 5 * time.Second
	numBins = 10    // Number of bins for distribution
	binSize = 0.025 // Size of each bin in seconds (e.g., 25ms)
)

func main1() {
	var (
		connTimes []float64
		addr      string
		count     int
	)
	flag.StringVar(&addr, "addr", "", "addr")
	flag.IntVar(&count, "count", 10, "count")
	flag.Parse()

	for i := 0; i < count; i++ {
		connTime := measureConnectionTime1(addr) // Replace with your IP and port
		connTimes = append(connTimes, connTime)
	}

	if len(connTimes) == 0 {
		fmt.Println("没有成功建立连接")
		return
	}

	sort.Float64s(connTimes)

	maxValue := connTimes[len(connTimes)-1]
	minValue := connTimes[0]

	var sum float64
	for _, t := range connTimes {
		sum += t
	}
	mean := sum / float64(len(connTimes))

	var varianceSum float64
	for _, t := range connTimes {
		diff := t - mean
		varianceSum += diff * diff
	}
	variance := varianceSum / float64(len(connTimes))

	fmt.Printf("最大值: %.3f 秒\n", maxValue)
	fmt.Printf("最小值: %.3f 秒\n", minValue)
	fmt.Printf("平均值: %.3f 秒\n", mean)
	fmt.Printf("方差: %.3f 秒\n", variance)

	// Calculate distribution in 10 bins
	bins := make([]int, numBins)
	for _, t := range connTimes {
		binIndex := int(t / binSize)
		if binIndex >= numBins {
			binIndex = numBins - 1
		}
		bins[binIndex]++
	}

	fmt.Println("连接时间分布（每个区段的次数）:")
	for i, count := range bins {
		lowerBound := float64(i) * binSize
		upperBound := lowerBound + binSize
		if upperBound > maxValue {
			upperBound = maxValue
		}
		fmt.Printf("%.3f - %.3f 秒: %d\n", lowerBound, upperBound, count)
	}
}

func measureConnectionTime1(addr string) float64 {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return 0
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error closing connection:", err)
		}
	}(conn)

	duration := time.Since(start).Seconds()
	return duration
}
