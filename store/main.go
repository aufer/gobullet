package main

import "github.com/aufer/gobullet/model"

var config = model.SockClientConfig{
	Prot: "tcp",
	Host: "127.0.0.1",
	Port: "8282",
}

func buildStoreHandler(store Store) model.Handler[interface{}] {
	return model.CreateHandler[interface{}](
		"StoreHandler",
		func(action model.Action) bool {
			return action.IsGroup("store")
		},
		func(action model.Action) (interface{}, bool) {

		},
		func(res interface{}) []model.Action {
			return []model.Action{}
		},
		"",
	)
}

func main() {
	store := Store{}

	var client = model.SockClient[interface{}]{
		Config:  config,
		Handler: buildStoreHandler(store),
	}

	client.Connect()
}
