package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

func main() {
	addr := flag.String("addr", "unixgram:/opt/passport/sys-logservice/run/log.sock", "unix socket")
	file := flag.String("file", "/var/log/message", "file")

	flag.Parse()
	fp, err := os.OpenFile(*file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	arr := strings.SplitN(*addr, ":", 2)
	pro := "unix"
	path := *addr
	if len(arr) > 1 {
		pro = arr[0]
		path = arr[1]
	}
	defer func() {
		_ = os.Remove(path)
	}()
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-ch
		_ = os.Remove(path)
		os.Exit(0)
	}()

	_ = os.Remove(path)
	wg := sync.WaitGroup{}
	if pro == "unix" {
		wg.Add(1)
		go func() {
			if err := unixStream(path, fp); err != nil {
				panic(err)
			}
			wg.Done()
		}()
	}
	if pro == "unixgram" {
		wg.Add(1)
		go func() {
			if err := unixGram(path, fp); err != nil {
				panic(err)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func unixStream(addr string, fp io.Writer) error {

	l, err := net.Listen("unix", addr)
	if err != nil {
		return err
	}
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go func() {
			r := bufio.NewReader(c)
			for {
				c.SetReadDeadline(time.Now().Add(10 * time.Second))
				line, err := r.ReadBytes('\n')
				if len(line) > 0 {
					fp.Write(line)
					if line[len(line)-1] != '\n' {
						fp.Write([]byte{'\n'})
					}
				}
				if err == io.EOF {
					break
				}
			}
			c.Close()
		}()
	}
}

func unixGram(addr string, fp io.Writer) error {
	l, err := net.ListenPacket("unixgram", addr)
	if err != nil {
		return err
	}
	buf := make([]byte, 65535)
	for {
		l.SetReadDeadline(time.Now().Add(10 * time.Second))
		n, addr, _ := l.ReadFrom(buf)
		if n > 0 {
			from := "<nil>"
			if addr != nil {
				from = addr.String()
			}
			if buf[n-1] != '\n' {
				buf[n] = '\n'
				n++
			}
			fmt.Fprintf(fp, "%s %s", from, buf[:n])
		}
	}
}
