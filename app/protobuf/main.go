package main

/**
 * 1.  安装 protoc-gen-go: go install github.com/golang/protobuf/protoc-gen-go
 * 2.  生成pb文件： protoc --go_out=. test.proto
 */

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"proto-test/proto/test"
)

func main() {
	p := &test.Person{}
	p.Hobby = &test.Hobby{ // 这个指针虽然没有实质性内容，依然会在marshal中体现出来的
		//Name: "PHPor",
	}
	v, err := proto.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", p)
	p2 := &test.Person{}
	if err = proto.Unmarshal(v, p2); err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", p2)
}
