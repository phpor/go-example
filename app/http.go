package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"os"
)

func main() {
	get()
}

func get() {
	res, err := http.Get("http://www.baidu.com/")
	check_fail(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	check_fail(err)
	fmt.Println(string(body))
}

func http_with_timeout() {

	req, err := http.NewRequest("GET", "http://baidu.com/", nil)
	check_fail(err)
	resp, err := http.DefaultClient.Do(req)
	check_fail(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check_fail(err)
	println(body)


}
func check_fail(err error) {
	println(err)
	os.Exit(1)
}

