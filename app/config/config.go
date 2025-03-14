package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"time"
)

type ModuleConfig interface {
	Keys() []string
	Has(string) bool
	Get(string) ModuleConfig // 返回值：ModuleConfig
	Value() interface{}      // 返回值：ModuleConfig、[]ModuleConfig、内部数据类型
	UnMarshal(v interface{}) error
}

var (
	tomlConfigNil = &TomlConfig{nil}
)

// TomlConfig ...
type TomlConfig struct {
	value interface{}
}

// NewTomlConfig ...
func NewTomlConfig() *TomlConfig {
	return &TomlConfig{}
}

// Get ...
func (cfg *TomlConfig) Get(module string) ModuleConfig {
	if cfg == tomlConfigNil {
		return tomlConfigNil
	}
	if tree, ok := cfg.value.(*toml.Tree); ok {
		if v := tree.Get(module); v != nil {
			return &TomlConfig{v}
		}
	}
	return tomlConfigNil
}

// Value ...
func (cfg *TomlConfig) Value() interface{} {
	var v interface{}
	switch node := cfg.value.(type) {
	case *toml.Tree:
		v = &TomlConfig{node}
	case []*toml.Tree:
		var nodes []ModuleConfig
		for _, n := range node {
			nodes = append(nodes, &TomlConfig{n})
		}
		v = nodes
	default:
		return cfg.value
	}
	return v
}

// Has ...
func (cfg *TomlConfig) Has(key string) bool {
	if tree, ok := cfg.value.(*toml.Tree); ok {
		return tree.Has(key)
	}
	return false
}

// Keys ...
func (cfg *TomlConfig) Keys() []string {
	if tree, ok := cfg.value.(*toml.Tree); ok {
		return tree.Keys()
	}
	return nil
}

func (cfg *TomlConfig) GetDuration(key string) (time.Duration, bool, error) {
	v := cfg.Get(key).Value()
	if v == nil {
		return 0, false, nil
	}
	if s, ok := v.(string); !ok {
		return 0, true, fmt.Errorf("%s is not string", key)
	} else {
		d, err := time.ParseDuration(s)
		if err != nil {
			return 0, true, fmt.Errorf("parse %s as time.Duration fail: %s", key, err.Error())
		}
		return d, true, nil
	}
}

func (cfg *TomlConfig) ParseDuration(key string, d *time.Duration) (bool, error) {
	duration, exists, err := cfg.GetDuration(key)
	if !exists || err != nil {
		return exists, err
	}
	if d != nil {
		*d = duration
	}
	return true, nil
}

func (cfg *TomlConfig) UnMarshal(v interface{}) error {
	if cfg.value == nil {
		return nil
	}
	if t, ok := cfg.value.(*toml.Tree); ok {
		return t.Unmarshal(v)
	}
	return nil
}

// Test ...
func (cfg *TomlConfig) Test(filename string) error {
	return cfg.LoadFile(filename)
}

// LoadFile ...
func (cfg *TomlConfig) LoadFile(filename string) error {
	tree, err := toml.LoadFile(filename)
	if err != nil {
		return err
	}
	cfg.value = tree
	return nil
}

func (cfg *TomlConfig) LoadString(content string) error {
	tree, err := toml.Load(content)
	if err != nil {
		return err
	}
	cfg.value = tree
	return nil
}

func (cfg *TomlConfig) LoadMap(content map[string]interface{}) error {
	tree, err := toml.TreeFromMap(content)
	if err != nil {
		return err
	}
	cfg.value = tree
	return nil
}
