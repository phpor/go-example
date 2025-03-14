package main

import (
	"fmt"
	"github.com/olivere/elastic"
)

func main() {
	// 读标准输入，写入elasticsearch
	_, err := elastic.NewClient()
	if err != nil {
		println(err.Error())
	}

	m := map[string]int{}
	m["aa"]++
	fmt.Println(m)

}
