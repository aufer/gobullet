package model

import (
	"encoding/json"
	"log"
	"net/http"
	"unsafe"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WebSocketServer struct {
	socketClient  SockClient[Action]
	wsConnections map[string]websocket.Conn
}

func NewWebSocketServer(config SockClientConfig) WebSocketServer {
	return WebSocketServer{
		wsConnections: map[string]websocket.Conn{},
		socketClient: SockClient[Action]{
			Config:  config,
			Handler: wsHandler,
		},
	}
}

type IdExchange struct {
	ConnectionId string `json:"connectionId"`
}

var wsHandler = CreateHandler(
	"UiConnector",
	func(action Action) bool {
		return action.IsGroup("ui")
	},
	func(action Action) (Action, bool) {
		return action, true
	},
	func(res Action) []Action {
		f := (*Action)(unsafe.Pointer(&res))
		return []Action{*f}
	},
	"",
)

func wsConnectionHandler(wss *WebSocketServer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var upgrader = websocket.Upgrader{}
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		connection, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}

		defer connection.Close()

		connId := uuid.New().String()
		connection.WriteJSON(Action{
			Type: "UI:WELCOME",
			Payload: IdExchange{
				ConnectionId: connId,
			}},
		)
		wss.wsConnections[connId] = *connection

		// SockClient writes message to the web socket
		wss.socketClient.SendFn = func(action Action) {
			connection.WriteJSON(action)
		}

		for {
			_, message, err := connection.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			var uiAction Action
			json.Unmarshal(message, &uiAction)
			log.Printf("recv: %s\n", message)
			wss.socketClient.Send(uiAction)
		}
	}
}

func (wss *WebSocketServer) Run() {
	go wss.socketClient.Connect()

	http.HandleFunc("/ws", wsConnectionHandler(wss))
	log.Println("Starting WebSocket server on 0.0.0.0:8484")
	log.Fatal(http.ListenAndServe("0.0.0.0:8484", nil))
}
