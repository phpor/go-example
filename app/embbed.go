package main

type  p struct {
	name string  
}

func (*p) say()  {
	println("I am p")
}

func (*p) run()  {
	println("p is running")
}

type s struct {
	p
}


func (*s) say()  {
	println("I am s")
}

func main()  {
	(&s{}).say()
	(&s{}).run()
}