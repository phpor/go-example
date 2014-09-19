package main

import (
	"fmt"
	"time"
)

func main() {
	// 格式化
	fmt.Println("Now: ", time.Now())
	fmt.Println("Now as Unix: ", time.Now().Unix())
	fmt.Println("Now as my format: ", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("One hour ago: ", time.Unix(time.Now().Unix()-3600, 0))
	fmt.Println("One hour ago: ", time.Unix(time.Now().Unix()-3600, 0).Format(time.RFC3339))
	fmt.Printf("%v\n", time.Now())
	fmt.Println(time.Now().Before(time.Now().Add(time.Second)))
}

