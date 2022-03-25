package main

import "github.com/aufer/gobullet/model"

var config = model.SockClientConfig{
	Prot: "tcp",
	Host: "127.0.0.1",
	Port: "8282",
}

func main() {
	var client = model.SockClient[string]{
		Config:  config,
		Handler: SearchHandler,
	}

	client.Connect()
}
