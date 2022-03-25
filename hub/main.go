package main

import (
	"github.com/aufer/gobullet/model"
)

var config = model.SockClientConfig{
	Prot: "tcp",
	Host: "0.0.0.0",
	Port: "8282",
}

func main() {
	hub := NewHub()

	hub.Run(config)
}
