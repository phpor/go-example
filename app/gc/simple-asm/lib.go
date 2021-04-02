package main

func main() {
	b := BigStruct{
		name: "phpor",
		addr: "",
		age:  18,
	}
	b.reset()
	b.reset2()
}

type BigStruct struct {
	name string
	addr string
	age  int
}

func (b *BigStruct) reset() {
	b.name = ""
	b.addr = ""
	b.age = 0
}

func (b *BigStruct) reset2() {
	*b = BigStruct{}
}
