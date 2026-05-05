package main

import (
	"fmt"
	"time"
)

var DefaultTimeFormat = "2006-01-02 15:04:05"

func main() {
	now := time.Now()
	now2 := now.In(time.FixedZone("UTC", 0))
	fmt.Printf("%d   %d\n", now.Unix(), now.Hour())
	fmt.Printf("%d   %d\n", now2.Unix(), now2.Hour())
}

func timeXXX() {
	fmt.Printf("%#v\n", time.Now())
	LOCAL, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("%v\n", time.Local)

	fmt.Printf("old: %d\n", time.Now().Unix())
	secondsEastOfUTC := int((7 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	time.Local = beijing

	fmt.Printf("%v\n", LOCAL)
	fmt.Printf("aa: %s\n", time.Now().Format(DefaultTimeFormat))
	fmt.Printf("%v\n", time.Now().Location())
	fmt.Printf("%d\n", time.Now().In(beijing).Unix())
	fmt.Printf("new: %d\n", time.Now().Unix())
}
