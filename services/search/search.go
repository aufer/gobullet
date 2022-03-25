package main

import (
	"fmt"

	"github.com/aufer/gobullet/model"
)

var SearchHandler = model.CreateHandler(
	"SearchHandler",
	func(action model.Action) bool {
		return action.IsType("search")
	},
	func(action model.Action) (string, bool) {
		return fmt.Sprintf("http://google.com/search?q=%v", action.Payload), true
	},
	func(res string) []model.Action {
		return []model.Action{
			model.CreateActionWithSource("ui:search", res, "", "SearchHandler"),
		}
	},
	"",
)
