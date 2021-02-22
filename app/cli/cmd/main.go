package main

import (
	"flag"
	"github.com/phpor/go-example/app/cli"
	"strings"
)

type cliTest struct {
	cli.Cli
}

func (ct *cliTest) ActionTestA() error {
	uid := ct.GetParam("uid").(string)
	println(uid)
	return nil
}

func main() {
	c := &cliTest{}

	uid := flag.String("uid", "103630", "uid")
	action := flag.String("action", "test-a", "[ "+strings.Join(cli.ActionList(c), " ")+" ]")
	flag.Parse()
	flag.VisitAll(func(f *flag.Flag) { // 这里并不方便把name、value的值统一写进去
		c.AddParam(f.Name, f.Value)
	})
	c.AddParam("uid", *uid)
	err := cli.Dispatch(c, *action)
	if err == cli.ErrActionNotFound {
		flag.Usage()
		return
	}
}
