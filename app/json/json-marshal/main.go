package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Int struct {
	i           int
	isUndefined bool
}

func (i *Int) MarshalJSON() ([]byte, error) {
	if i.isUndefined {
		return []byte{}, nil
	}
	return []byte(strconv.Itoa(i.i)), nil
}

func main() {
	i := &Int{1, true}
	a := map[string]*Int{"age": i}
	b, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}
	fmt.Printf("%s", string(b))
}
