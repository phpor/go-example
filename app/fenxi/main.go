package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 使用说明
// 1. 读取一个文本文件，文本文件内容为逗号分隔的两列，第一列是UID，第二列是时间戳，该文件是先按照UID排序，再按照时间戳排序
// 2. 统计连续的UID，如果两个UID的时间戳差值大于等于60秒，则认为是两个不同的记录，统计记录的数量； 换言之，相同的UID，间隔小于60s则视为同一个记录

func main() {
	var (
		prevUID       string
		prevTimestamp int64
		totalCount    int
	)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			fmt.Fprintf(os.Stderr, "Invalid line: %s\n", line)
			continue
		}

		uid := parts[0]
		timestampStr := parts[1]

		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid timestamp: %s\n", timestampStr)
			continue
		}

		if uid != prevUID || prevTimestamp == 0 {
			// New UID or first record
			totalCount++
		} else {
			// Same UID, check time difference
			timeDiff := timestamp - prevTimestamp
			if timeDiff >= 60 {
				totalCount++
			}
		}

		prevUID = uid
		prevTimestamp = timestamp
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Reading input: %s\n", err)
	}

	fmt.Println(totalCount)
}
