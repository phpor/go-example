package main

type M struct {
}

func (m *M) say1() {
	println("say1")
}

func (m M) say2() {
	println("say2")
}

func main() {
	var m *M
	m.say1()
	m.say2()
}
