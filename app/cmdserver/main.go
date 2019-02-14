package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	addr := flag.String("addr", "192.168.56.1:8181", "listen addr")
	root := flag.String("root", "./", "file root dir")
	confFile := flag.String("confFile", "./cmd-config", "cmd config file")
	config := newConfig(*confFile)
	confToken := config.getToken()
	flag.Parse()
	http.Handle("/static/", http.FileServer(http.Dir(*root)))
	http.HandleFunc("/cmd/", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		token := request.Form.Get("token")
		if token != confToken {
			fmt.Println(token, ":", confToken)
			writer.Write([]byte("403"))
			return
		}
		alias := ""
		arr := strings.Split(request.RequestURI, "/")
		if len(arr) > 2 {
			alias = arr[2]
		}
		if alias == "" {
			writer.Write([]byte("400"))
			return
		}
		cmdline := config.getCmd(alias)
		if cmdline == "" {
			writer.Write([]byte("400"))
			return
		}
		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			writer.Header().Set("Content-type", "text/html; charset=gbk")
			cmd = exec.Command("powershell.exe", "-Command", cmdline)
		} else {
			cmd = exec.Command("/bin/bash", "-c", cmdline)
		}

		output, err := cmd.Output()
		if err != nil {
			writer.Write([]byte(err.Error()))
		} else {
			writer.Write(output)
		}
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		panic(err)
	}
}

type Config struct {
	file  string
	cmds  map[string]string
	token string
}

func newConfig(confFile string) *Config {
	c := &Config{
		file: confFile,
		cmds: map[string]string{},
	}
	b, err := ioutil.ReadFile(c.file)
	if err != nil {
		panic("Error: " + err.Error())
	}
	lines := strings.Split(fmt.Sprintf("%s", b), "\n")
	for _, line := range lines {
		arr := strings.SplitN(line, "=", 2)
		if len(arr) > 1 {
			key := strings.TrimSpace(arr[0])
			value := strings.TrimSpace(arr[1])

			if key == "token" {
				c.token = value
			} else {
				c.cmds[key] = value
			}
		}
	}
	if c.token == "" {
		panic("Error: must config token")
	}
	return c
}
func (c *Config) getToken() string {
	return c.token
}
func (c *Config) getCmd(alias string) (cmd string) {
	if v, exits := c.cmds[alias]; exits {
		return v
	}
	return ""
}
