package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

type strLen int

func (s *strLen) UnmarshalJSON(bytes []byte) error {
	*s = strLen(len(bytes))
	return nil
}

type metrics struct {
	count int
	max   int
	min   int
	avg   int
}

func (m *metrics) String() string {
	return fmt.Sprintf("count\t%d\tmax\t%d\tmin\t%d\tavg\t%d", m.count, m.max, m.min, m.avg)
}
func main() {
	var action string
	var header bool
	flag.StringVar(&action, "action", "default", "default | statistics")
	flag.BoolVar(&header, "header", true, "out put header on statistics")
	flag.Parse()
	r := bufio.NewReaderSize(os.Stdin, 64*1024)
	var err error
	var line []byte
	mStatistics := map[string]*metrics{}
	var f func(m map[string]strLen)
	switch action {
	case "statistics":
		f = func(m map[string]strLen) {
			var s *metrics
			var length int
			for k, v := range m {
				length = int(v)
				if _, ok := mStatistics[k]; !ok {
					mStatistics[k] = &metrics{}
				}
				s = mStatistics[k]
				if length > s.max {
					s.max = length
				}
				if length < s.min {
					s.min = length
				}
				s.avg = (s.avg*s.count + length) / (s.count + 1)
				s.count++
			}
		}
	case "default":
		f = func(m map[string]strLen) {
			line, _ := json.Marshal(m)
			fmt.Println(string(line))
		}
	default:
		println("action only can be [ default | statistics ]")
		return
	}

	for {

		line, _, err = r.ReadLine()
		if len(line) == 0 && err == io.EOF {
			break
		}
		if len(line) == 0 {
			continue
		}
		m := map[string]strLen{}
		err = json.Unmarshal(line, &m)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: %s\t%s\n", err.Error(), string(line))
			continue
		}
		f(m)
	}
	if action == "statistics" {
		if header {
			fmt.Printf("%-32s%-8s%-6s%-6s%-6s\n", "name", "count", "avg", "max", "min")
		}
		for k, v := range mStatistics {
			fmt.Printf("%-32s%-8d%-6d%-6d%-6d\n", k, v.count, v.avg, v.max, v.min)
		}
	}
}
