package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	t.In(time.FixedZone("china", 7))
	println(t.Format("2006-01-02 15:04:05 -0700"))
}

func zeroTime() {
	t1 := time.Time{}
	t2 := time.Unix(0, 0)
	fmt.Printf("%#v, %#v, seconds: %d\n", t1, t1.IsZero(), t1.Unix())
	fmt.Printf("%#v, %#v, seconds: %d\n", t2, t2.IsZero(), t2.Unix())
	fmt.Printf("%d, %s\n", 62135596800, time.Duration(62135596800))
}

func timeLocation() {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-01-05 13:13:13", time.Local)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", t)
}
