package main

import (
	"fmt"

	"github.com/aufer/gobullet/model"
)

var LogHandler = model.CreateHandler(
	"LogHandler",
	func(action model.Action) bool {
		return true
	},
	func(action model.Action) (interface{}, bool) {
		fmt.Printf("LogHandler: [%v] with payload {%v} from %v\n", action.Type, action.Payload, action.Source)
		return nil, false
	},
	func(res interface{}) []model.Action {
		return []model.Action{}
	},
	"",
)
