package a

import (
	"fmt"
	"github.com/phpor/go-example/app/package/b"
)

func A() {
	fmt.Println("I am in a")
}
func A1() {
	b.B()
}
