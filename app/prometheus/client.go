package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/ffmt.v1"
	"net/http"
	"net/url"
	"runtime/debug"
	"time"
)

type basicAuth struct {
	uername  string
	password string
}
type Client struct {
	endpoint string
	auth     *basicAuth
}

func NewClient(endpoint string) *Client {
	return &Client{endpoint: endpoint}
}

func (c *Client) SetBasicAuth(username, password string) *Client {
	c.auth = &basicAuth{
		uername:  username,
		password: password,
	}
	return c
}

type MetricType struct {
	NodeName string `json:"node_name"`
}

type ResultType struct {
	Metric MetricType    `json:"metric"`
	Value  []interface{} `json:"value"`
}

type QueryData struct {
	ResultType string       `json:"resultType"`
	Result     []ResultType `json:"result"`
}

type QueryInfo struct {
	Status string    `json:"status"`
	Data   QueryData `json:"data"`
}

func (c *Client) Query(query string) (*QueryInfo, error) {
	return c.queryMetric(query)
}

func (c *Client) query(url string, result interface{}) error {
	httpClient := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("Get", url, nil)
	if err != nil {
		return err
	}
	if c.auth != nil {
		req.SetBasicAuth(c.auth.uername, c.auth.password)
	}
	r, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(result)
	if err != nil {
		fmt.Printf("%s", debug.Stack())
		debug.PrintStack()
		return err
	}
	return nil
}

//queryMetric query metric by prom api
func (c *Client) queryMetric(query string) (*QueryInfo, error) {
	info := &QueryInfo{}
	ustr := c.endpoint + "/api/v1/query?query=" + query
	u, err := url.Parse(ustr)
	if err != nil {
		return info, err
	}
	u.RawQuery = u.Query().Encode()

	err = c.query(u.String(), &info)
	if err != nil {
		ffmt.Puts(info)
		return info, err
	}
	return info, nil
}
