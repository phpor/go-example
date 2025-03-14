package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	testRetry()
	//
	//res, err := http.Get("https://wetest.qq.com/app/testlab/cgi/v1/model/list/filter")
	//if err != nil {
	//	panic(err)
	//}
	//body, err := io.ReadAll(res.Body)
	//if err != nil {
	//	panic(err)
	//}
	//println(string(body))
}
func testHttp22() {
	hc := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			ForceAttemptHTTP2: true,
		},
	}
	client := resty.NewWithClient(hc)
	client.EnableTrace()
	client.Debug = true
	res, err := client.R().Get("http://localhost:8080/") //.SetResult(&result) 只有 JSON 或 XML 会解析到这个Result里面
	if err != nil {
		panic(err)
	}
	fmt.Println(res.String())
}

func testHttp2(url string) {
	h := getHttpClient()
	run := func() {
		req, _ := http.NewRequest("GET", url, nil)
		resp, err := h.Do(req)
		if err != nil {
			println(err)
			return
		}
		defer func() {
			_ = resp.Body.Close()
		}()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			println(err)
			return
		}
		fmt.Println(string(b))
		return

	}
	count, _ := strconv.ParseInt(os.Getenv("Count"), 10, 64)
	if count == 0 {
		count = 3
	}
	sleep, _ := strconv.ParseInt(os.Getenv("Sleep"), 10, 64)
	for i := 0; i < int(count); i++ {
		time.Sleep(time.Duration(sleep) * time.Second)
		run()
	}
}

func testHttp3() {
	count, _ := strconv.ParseInt(os.Getenv("Count"), 10, 64)
	if count == 0 {
		count = 3
	}
	sleep, _ := strconv.ParseInt(os.Getenv("Sleep"), 10, 64)
	for i := 0; i < int(count); i++ {
		time.Sleep(time.Duration(sleep) * time.Second)
		testHttp(os.Args[1])
	}
}
func testHttp(url string) {
	resp, err := http.Get(url)
	if err != nil {
		println(err)
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println(err)
		return
	}
	fmt.Println(string(b))
	return
}

func getHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, err error) {
				conn, err = net.DialTimeout(network, addr, 2*time.Second)
				if err != nil {
					return nil, err
				}
				return conn, err
			},
		},
	}
}

func testRetry() {
	// 从代码上来看 client.execute 来看，只有连接失败、写失败、读流失败，才会重试； 如果解析内容不符合预期不会自动重试
	// 对于不想重试的错误，都通过 wrapNoRetryErr() 来进行包装
	client := resty.New().SetTimeout(time.Second)
	client.SetRetryCount(3)
	result := &struct{}{}
	resp, err := client.R().ForceContentType("application/json").SetResult(result).Get("https://baidu.com/")
	if err != nil {
		println(err.Error())
	}
	if resp != nil {
		//fmt.Println(resp.String())
	}
}

func testResty() {
	client := resty.New().SetTimeout(time.Second)
	client.EnableTrace()
	//http.Debug = true
	try := func() {
		resp, err := client.R().Get("https://baidu.com/")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%#v\n\n", resp.Request.TraceInfo())

	}
	try()
	try()

}
