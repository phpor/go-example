package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	http_with_timeout()
}

func get() {
	res, err := http.Get("http://www.baidu.com/")
	check_fail(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	check_fail(err)
	fmt.Println(string(body))
}

func http_with_custom_header() {

	req, err := http.NewRequest("GET", "http://www.baidu.com/", nil)
	check_fail(err)
	req.Header.Set("x-debug", "on")
	resp, err := http.DefaultClient.Do(req)
	check_fail(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check_fail(err)
	println(string(body))

}

// 参考： http://1234n.com/?post/mwsw2r
func http_with_timeout() {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*1)
				if err != nil {
					return nil, err
				}
				return NewTimeoutConn(conn, time.Second*1), nil
			},
			ResponseHeaderTimeout: time.Second * 1,
		},
	}
	req, err := http.NewRequest("GET", "http://www.google.com/", nil)
	check_fail(err)
	resp, err := client.Do(req)
	check_fail(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check_fail(err)
	println(string(body))
}

type TimeoutConn struct {
	conn    net.Conn
	timeout time.Duration
}

func NewTimeoutConn(conn net.Conn, timeout time.Duration) *TimeoutConn {
	return &TimeoutConn{
		conn:    conn,
		timeout: timeout,
	}
}

func (c *TimeoutConn) Read(b []byte) (n int, err error) {
	c.SetReadDeadline(time.Now().Add(c.timeout))
	return c.conn.Read(b)
}

func (c *TimeoutConn) Write(b []byte) (n int, err error) {
	c.SetWriteDeadline(time.Now().Add(c.timeout))
	return c.conn.Write(b)
}

func (c *TimeoutConn) Close() error {
	return c.conn.Close()
}

func (c *TimeoutConn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *TimeoutConn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *TimeoutConn) SetDeadline(t time.Time) error {
	return c.conn.SetDeadline(t)
}

func (c *TimeoutConn) SetReadDeadline(t time.Time) error {
	return c.conn.SetReadDeadline(t)
}

func (c *TimeoutConn) SetWriteDeadline(t time.Time) error {
	return c.conn.SetWriteDeadline(t)
}

func check_fail(err error) {
	if err != nil {
		println(err)
		os.Exit(1)
	}
}
