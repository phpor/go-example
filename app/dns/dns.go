package main

import (
	"net"
	"fmt"
	"github.com/phpor/godns"
)

func main() {
	addr, err := net.LookupHost("www.baidu.com")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%v", addr)
	option := &godns.LookupOptions{
		DNSServers: []string{"172.16.10.4"},
	}
	addr, err = godns.LookupHost("www.baidu.com", option)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%v", addr)
}
