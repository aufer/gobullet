package model

import (
	"fmt"
	"strings"
	"time"
)

const typeDelim = ":"

type Action struct {
	Type      string      `json:"type"`
	Payload   interface{} `json:"payload"`
	Timestamp int64       `json:"timestamp"`
	Source    string      `json:"source"`
	SessionId string      `json:"sessionId"`
}

func (a *Action) IsType(typ string) bool {
	return IsActionType(a, typ)
}

func (a *Action) IsGroup(group string) bool {
	u := strings.ToUpper
	return strings.HasPrefix(u(a.Type)+":", u(group))
}

func (a *Action) GetType() string {
	return strings.ToUpper(a.Type)
}

func (a *Action) HasPayload() bool {
	return a.Payload != nil
}

func (a *Action) ReqType() string {
	return strings.ToUpper(fmt.Sprintf("%s%s%s", strings.ToLower(a.Type), typeDelim, "request"))
}

func (a *Action) ResType() string {
	return strings.ToUpper(fmt.Sprintf("%s%s%s", strings.ToLower(a.Type), typeDelim, "response"))
}

func CreateAction(typ string, payload interface{}, sessionId string) Action {
	return Action{
		Type:      strings.ToUpper(typ),
		Payload:   payload,
		Timestamp: time.Now().UnixMilli(),
		SessionId: sessionId,
	}
}

func CreateActionWithSource(typ string, payload interface{}, sessionId string, source string) Action {
	return Action{
		Type:      strings.ToUpper(typ),
		Payload:   payload,
		Timestamp: time.Now().UnixMilli(),
		SessionId: sessionId,
		Source:    source,
	}
}

func IsActionType(action *Action, typ string) bool {
	return strings.ToUpper(action.Type) == strings.ToUpper(typ)
}
