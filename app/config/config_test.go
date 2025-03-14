package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"testing"
	"time"
)

var content = `
daemon = true
pid_file = "run/device-server.pid"
start_timeout = "5s"

[[registry]]
type = "consul"
name = "consul-default"
`

var mapContent = map[string]interface{}{}

type config struct {
	Daemon       bool                `toml:"daemon"`
	PidFile      string              `toml:"pid_file"`
	LogFile      string              `toml:"log_file"`
	StartTimeout time.Duration       `toml:"start_timeout"`
	Registry     []map[string]string `toml:"registry"`
}

var expectConfig = &config{
	Daemon:       true,
	PidFile:      "run/device-server.pid",
	LogFile:      "",
	StartTimeout: 5 * time.Second,
	Registry: []map[string]string{
		{
			"type": "consul",
			"name": "consul-default",
		},
	},
}

func TestMain(m *testing.M) {
	if err := setUp(); err != nil {
		panic(err)
	}
	code := m.Run()
	os.Exit(code)
}

func TestTomlConfig_LoadMap(t *testing.T) {
	cfg := NewTomlConfig()
	err := cfg.LoadMap(mapContent)
	assert.Equal(t, nil, err, "LoadMap should be success")
	d, exists, err := cfg.GetDuration("start_timeout")
	assert.Equal(t, true, exists, "start_timeout expect exists")
	assert.Equal(t, nil, err, "start_timeout should be parse success")
	assert.Equal(t, 5*time.Second, d, "start_timeout should be 5s")
}

func TestTomlConfig_UnMarshal(t *testing.T) {
	c := &config{
		LogFile: "logs/device-server.log", // this is default value
	}
	expect := *expectConfig
	reflect.Copy(reflect.ValueOf(expect.Registry), reflect.ValueOf(expectConfig.Registry))

	expect.LogFile = c.LogFile

	cfg := NewTomlConfig()
	err := cfg.LoadString(content)
	assert.Equal(t, nil, err, "LoadString should be success")
	err = cfg.UnMarshal(c)
	assert.Equal(t, nil, err, "UnMarshal should be success")
	assert.EqualValues(t, &expect, c, "UnMarshal unexpected")

}

func setUp() error {
	err := toml.Unmarshal([]byte(content), &mapContent)
	if err != nil {
		return fmt.Errorf("init fail: %s", err.Error())
	}
	return nil
}
