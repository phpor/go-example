package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"time"
)

type Bag interface {
	Name() string
}
type SmallBag struct {
	NameA string `toml:"name"`
}

func (b *SmallBag) UnmarshalJSON(bytes []byte) error {
	b.NameA = fmt.Sprintf("%s", string(bytes))
	return nil
}

//func (b *SmallBag) UnmarshalTOML(i interface{}) error {
//	b.NameA = fmt.Sprintf("%v", i)
//	return nil
//}

func (b *SmallBag) Name() string {
	return b.NameA
}

type ClientConfig struct {
	Name    string
	Addr    string
	Timeout time.Duration
	Bag     Bag
}

func main() {

	c := &ClientConfig{
		Timeout: 3 * time.Second,
	}
	content := `
daemon = false
pid_file = "run/device-server.pid"
log_file = "logs/device-server.log"
go_max_procs = 0
start_timeout = "5s"
start_timeout2 = "10s"

[[registry]]
type = "consul"
name = "consul-default"
`

	conf := NewTomlConfig()
	if err := conf.LoadString(content); err != nil {
		panic(err)
	}
	if timeout, exists, err := conf.GetDuration("start_timeout"); err != nil {
		panic(err)
	} else if exists {
		c.Timeout = timeout
	}
	println(c.Timeout.String())

	if exists, err := conf.ParseDuration("start_timeout2", &c.Timeout); err != nil {
		panic(err)
	} else if !exists {
		println("start_timeout2 not exists; you can ignore this case if you careless")
	}
	println(c.Timeout.String())

	c1 := &ClientConfig{Bag: &SmallBag{}}
	err := toml.Unmarshal([]byte("bag = 2"), c1)
	if err != nil {
		println(err)
	}
}
