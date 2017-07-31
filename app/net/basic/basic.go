package basic

import (
	"net"
	"time"
)

func main()  {

}

type TimeoutConn struct {
	net.Conn
	timeout time.Duration
}

func NewTimeoutConn(conn net.Conn, timeout time.Duration) net.Conn{
	return &TimeoutConn{
		conn,
		timeout,
	}
}

func (c *TimeoutConn) Read(b []byte) (n int, err error) {
	c.SetReadDeadline(time.Now().Add(c.timeout))
	return c.Conn.Read(b)
}

func (c *TimeoutConn) Write(b []byte) (n int, err error) {
	c.SetWriteDeadline(time.Now().Add(c.timeout))
	return c.Conn.Write(b)
}

func timeout() {

}
