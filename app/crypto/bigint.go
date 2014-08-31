package main

import (
	"math/big"
	"fmt"
)

func main() {
	i := big.NewInt(2)
	j := big.NewInt(2)
	k := big.NewInt(8)
	m := i.Exp(i, j, k)
	fmt.Printf("%+v,%+v,%+v,%+v\n", i, j, k, m)
	//	println(i, j, k, m)
	fmt.Printf("%v\n", i.Add(i, j))
}
