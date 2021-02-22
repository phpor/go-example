package main

import (
	"encoding/json"
	"fmt"
)

type person struct { // 表面上json包没法new一个person值，但是，反射可以new这个结构，这里的私有结构体的私有只是提示给编译器看的
	Name string
	Age  int
}

type Team struct {
	Leader   *person // 指针也能被填充
	Follower person
}

func main() {
	str := `{"leader":{"Name":"wang", "age":22},"Follower":{"Name":"sun", "age":20}}`
	t := &Team{}
	if err := json.Unmarshal([]byte(str), t); err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("leader: %#v, follower: %#v", t.Leader, t.Follower)
}
