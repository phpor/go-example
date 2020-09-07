package main

import (
	"github.com/phpor/go-example/app/package/a"
	"github.com/phpor/go-example/app/package/b"
)

func main() { // import cycle not allowed
	a.A1()
	b.B1()
}
