package main

import (
	"errors"
	"fmt"
	"github.com/pelletier/go-toml"
	"net/http"
)

type HttpService struct {
	name string
	addr string
	on   bool
}

func (h *HttpService) Name() string {
	return h.name
}

func (h *HttpService) Start() error {
	h.on = true
	return nil
}

func (h *HttpService) Stop() error {
	h.on = false
	return nil
}

func (h *HttpService) IsRunning() bool {
	if !h.on {
		return false
	}
	_, err := http.Get(fmt.Sprintf("http://%s/", h.addr))
	if err != nil {
		return false
	}
	return true
}

func NewHttpServiceFromToml(t *toml.Tree) (*HttpService, error) {
	hs := &HttpService{}
	var ok bool
	if hs.name, ok = t.Get("name").(string); !ok {
		return nil, errors.New("http service name must be string")
	}
	if hs.addr, ok = t.Get("addr").(string); !ok {
		return nil, errors.New("http service addr must be string")
	}
	return hs, nil
}
