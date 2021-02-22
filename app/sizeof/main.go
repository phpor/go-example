package main

import (
	"fmt"
	"unsafe"
)

type Man struct {
	Name string
	Desc string
	Age  int
}

func main() {

	m := Man{Name: "Phpor", Desc: "111111111111111111111111111111", Age: 20}

	fmt.Println("man size:", unsafe.Sizeof(m))
	fmt.Println("name size:", unsafe.Sizeof(m.Name))
	fmt.Println("desc size:", unsafe.Sizeof(m.Desc))
	fmt.Println("age size:", unsafe.Sizeof(m.Age))
}
