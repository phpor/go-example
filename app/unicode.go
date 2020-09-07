package main

import (
	"fmt"
	"regexp"
)

func main() {
	regx := regexp.MustCompile("^[1-9]+[0-9]*$")
	fmt.Printf("%t", regx.MatchString("123"))
}
