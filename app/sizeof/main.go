package main

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

type Man struct {
	Name string
	Desc string
	Age  int
}

// https://stackoverflow.com/questions/2113751/sizeof-struct-in-go
/*
Knowing this rule and remembering that:

bool, int8/uint8 take 1 byte
int16, uint16 - 2 bytes
int32, uint32, float32 - 4 bytes
int64, uint64, float64, pointer - 8 bytes
string - 16 bytes (2 alignments of 8 bytes)
any slice takes 24 bytes (3 alignments of 8 bytes). So []bool, [][][]string are the same (do not forget to reread the citation I added in the beginning)
array of length n takes n * type it takes of bytes.
*/
// 这个结构体的大小要对齐到8字节
// a + b 对齐8字节， s 自己16字节
type good struct {
	a bool
	b bool
	s string
}

// a ~ e 对齐到8字节 ， s字节16字节， 共24字节
type good2 struct {
	a bool
	b bool
	c bool
	d bool
	e bool
	s string
}

// s 使得a和b不能节省空间了
type good3 struct {
	a bool
	s string
	b bool
}

func main() {

	m := Man{Name: "Phpor", Desc: "111111111111111111111111111111", Age: 20}

	fmt.Println("man size:", unsafe.Sizeof(m))
	fmt.Println("name size:", unsafe.Sizeof(m.Name))
	fmt.Println("desc size:", unsafe.Sizeof(m.Desc))
	fmt.Println("age size:", unsafe.Sizeof(m.Age))

	fmt.Println("good size: ", unsafe.Sizeof(good{}))
	fmt.Println("good2 size: ", unsafe.Sizeof(good2{}))
	fmt.Println("good3 size: ", unsafe.Sizeof(good3{}))
	fmt.Println("good2 binary.Size: ", binary.Size(good2{}))
	fmt.Println("good2 binary.Size: ", binary.Size(good2{s: "string"}))
}
