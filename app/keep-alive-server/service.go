package main

import (
	"errors"
	"fmt"
	"github.com/pelletier/go-toml"
)

type Service interface {
	Start() error
	Stop() error
	IsRunning() bool
	Name() string
}

type service struct {
	running bool
	name    string
}

func (s *service) Name() string {
	return s.name
}

func (s *service) Start() error {
	s.running = true
	return nil
}

func (s *service) Stop() error {
	s.running = false
	return nil
}

func (s *service) IsRunning() bool {
	return s.running
}

var services map[string]Service

func InitService(t *toml.Tree) error {
	v := t.Get("service")
	if v == nil {
		return nil
	}
	ss, ok := v.([]*toml.Tree)
	if !ok {
		return errors.New("config file not contains service slice")
	}

	for _, service := range ss {
		kind, ok := service.Get("kind").(string)
		if !ok {
			return fmt.Errorf("service kind must be string")
		}
		if kind == "http" {
			s, err := NewHttpServiceFromToml(t)
			if err != nil {
				return err
			}
			services[s.Name()] = s
		}
	}
	return nil
}

func InitServiceByFile(filename string) error {
	tree, err := toml.LoadFile(filename)
	if err != nil {
		return err
	}
	return InitService(tree)
}
