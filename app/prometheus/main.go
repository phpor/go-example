package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	var params struct {
		endpoint string
		username string
		password string
		query    string
	}
	flag.StringVar(&params.endpoint, "endpoint", os.Getenv("PROM_ENDPOINT"), "endpoint")
	flag.StringVar(&params.username, "username", os.Getenv("PROM_USERNAME"), "username")
	flag.StringVar(&params.password, "password", os.Getenv("PROM_PASSWORD"), "password")
	flag.StringVar(&params.query, "query", os.Getenv("PROM_QUERY"), "query")
	flag.Parse()
	c := NewClient(params.endpoint)
	if params.username != "" {
		c.SetBasicAuth(params.username, params.password)
	}
	info, err := c.Query(params.query)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := json.Marshal(info)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", string(b))
}
