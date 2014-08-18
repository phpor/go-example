package main

import "fmt"
import (
	"encoding/json"
	"github.com/phpor/goexample/sso"
)

func main() {
	content, err := sso.Login("ssologin4@sina.com", "ssologin5")
	if err != nil {
		fmt.Println(err)
		return
	}
	var f interface{}
	err = json.Unmarshal([]byte(content), &f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(f)
	m := f.(map[string]interface{})
	for k, v := range m {
		fmt.Println(k, v)
	}
}
