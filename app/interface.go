package main

import "fmt"

type mystdout struct{}

func (s *mystdout) Write(p []byte) (n int, err error) {
	return fmt.Print(string(p))
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

func main() {
	fmt.Fprint(&mystdout{}, "hello word")
}
