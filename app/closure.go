package main

import "fmt"

type myslice struct {
	s []string
}

func (s *myslice) each(f func(string)) {
	for _, str := range s.s {
		f(str)
	}
}

func main() {
	var s = &myslice{[]string{"hello", "world"}}
	s.each(func(s string) {
		fmt.Println(s)
	})
}
