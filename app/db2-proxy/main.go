package main

import (
	"flag"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	addr := flag.String("addr", ":33333", "listen addr")
	flag.Parse()
	http.HandleFunc("/cmd/db2", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		//command := r.PostFormValue("db2")
		command := "db2"
		option := r.PostFormValue("option")
		args := r.PostFormValue("args")
		if args == "quit" {
			go func() {
				time.Sleep(1 * time.Second)
				os.Exit(0)
			}()
		}
		cmd := exec.Command(command, option, args)
		buf, err := cmd.Output()
		retcode := cmd.ProcessState.ExitCode()
		w.Header().Add("x-cmd-retcode", strconv.Itoa(retcode))
		if err != nil {
			w.WriteHeader(500)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		_, _ = w.Write(buf)
		return
	})
	var err error
	if (*addr)[0] == '/' {
		server := http.Server{}
		unixListener, err := net.Listen("unix", *addr)
		if err != nil {
			panic(err)
		}
		err = server.Serve(unixListener)

	} else {
		err = http.ListenAndServe(*addr, nil)
	}
	if err != nil {
		panic(err)
	}
}
