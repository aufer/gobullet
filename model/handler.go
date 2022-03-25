package model

import (
	"fmt"
	"log"
	"strings"
)

type Handler[T any] struct {
	Id       string
	Trigger  func(action Action) bool
	Request  func(action Action) (T, bool)
	Response func(res T) []Action
	Produces string
}

func sysActionHandler[T any](action Action, id string) (T, bool) {
	log.Printf("[%v:REQUEST][SYS]: %v\n", id, action)
	return *new(T), false
}

func CreateHandler[T any](
	id string,
	trigger func(action Action) bool,
	request func(action Action) (T, bool),
	response func(res T) []Action,
	produces string,
) Handler[T] {
	return Handler[T]{
		Id: id,
		Trigger: func(action Action) bool {
			log.Printf("[%v:TRIGGER] %v\n", id, action)
			return action.Source != id && isSystemAction(action) || trigger(action)
		},
		Request: func(action Action) (T, bool) {
			log.Printf("[%v:REQUEST] %v\n", id, action)
			if isSystemAction(action) {
				return sysActionHandler[T](action, id)
			}

			return request(action)
		},
		Response: func(res T) []Action {
			log.Printf("[%v:RESPONSE] %v\n", id, res)
			actions := response(res)
			for idx := range actions {
				actions[idx].Source = id
			}

			fmt.Printf("%v returns actions %v", id, actions)
			return actions
		},
		Produces: produces,
	}
}

func isSystemAction(action Action) bool {
	return strings.HasPrefix(action.Type, "system:")
}
