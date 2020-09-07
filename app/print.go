package main

import (
	"encoding/json"
	"fmt"
)

type errorObject struct {
	str string
}
type printStruct struct {
	Msg []errorObject
}

func main() {
	msg := &printStruct{Msg: []errorObject{{str: "hhhh"}}}
	s, _ := json.Marshal(msg)
	fmt.Printf("%s", string(s)) //	{"Msg":[{}]}   没有显示详细的错误对象，因为str是私有的

}
