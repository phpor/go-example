package main

import "fmt"

type object struct {
	name string
}

func (o *object) Name() string {
	return o.name
}

func getObject() *object {
	if true {
		return nil
	}
	return &object{}
}

func main() {
	o := getObject()
	fmt.Println(o.Name())
}
