package main

import (
	"flag"
	"fmt"
	"reflect"
)

func main() {
	s := ""
	age := 0
	flag.StringVar(&s,"name", "", "")
	flag.IntVar(&age,"age", 0, "")
	err := flag.CommandLine.Parse([]string{"-name", "bbb"})
	if err != nil {
		println(err.Error())
		return
	}
	//flag.Parse()
	fmt.Printf("%v\n", flag.Args())
	flag.VisitAll(func(f *flag.Flag) { // 遍历所有定义了的选项； Visit() 只遍历命令行明确提供的选项
		t := reflect.TypeOf(f.Value)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		fmt.Printf("f.Value.name : %s\n", t.Kind())
	})
	println(s)
}
