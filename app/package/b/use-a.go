package b

import (
	"fmt"
	"github.com/phpor/go-example/app/package/a"
)

func B() {
	fmt.Println("I am in b")
}

func B1() {
	a.A()
}
