package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"reflect"
)

type Age int8

type P struct {
	Name   string
	Age    Age
	Custom interface{}
}

func main() {
	str := "Name: phpor\nage: 36\ncustom:\n - key: 1\n   val: 2"
	p := &P{}
	k := reflect.TypeOf(p.Age).Kind()
	fmt.Printf("%s\n", k.String())
	err := yaml.Unmarshal([]byte(str), p)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("P: %v\n", p)
	s, err := yaml.Marshal(p.Custom)
	b := []map[string]int{}
	yaml.Unmarshal(s, &b)
	fmt.Printf("%v", b)
}
