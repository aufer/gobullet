package main

import (
	"github.com/aufer/gobullet/model"
)

type Hub struct {
	bindings map[string](chan model.Action)
	logging  bool
}

func NewHub() Hub {
	return Hub{}
}

func (h *Hub) Run(config model.SockClientConfig) {
	socket := model.SockServer{
		Config: config,
	}

	socket.Listen()
}
