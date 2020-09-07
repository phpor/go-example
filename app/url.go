package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlTest3()
}

func urlTest1() {
	str := "a=b&a=b2&c=d&e=f"
	v, err := url.ParseQuery(str)
	if err != nil {
		println(err)
	}
	fmt.Println("%v", v)
}
func urlTest2() {
	str := "%7b"
	v, err := url.QueryUnescape(str)
	if err != nil {
		println(err)
	}
	println(v)
}
func urlTest3() {
	str := "sql://username:password@protocol(address)/dbname?param=value"
	u, err := url.Parse(str)

	if err != nil {
		println(err)
	}
	q := u.Query()
	println(q)
}
