package model

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"time"
)

var MAX_CONNECTION_RETRIES = 5
var RETRY_TIMEOUT = 5 * time.Second
var MSG_DELIM = "\n"

type SockClient[T any] struct {
	Config            SockClientConfig `validate:"required"`
	Handler           Handler[T]
	connectionRetries int
	comms             *bufio.ReadWriter
	SendFn            func(action Action)
}

func (sc *SockClient[T]) Connect() {
	log.Printf("Connecting with config %v\n", sc.Config)

	connection, err := net.Dial(sc.Config.Prot, sc.Config.ConnSrv())
	if err != nil {
		log.Printf("Error connecting to server with %v\n", err)
		sc.connectionRetries++

		if sc.connectionRetries < MAX_CONNECTION_RETRIES {
			log.Printf("Connection to the server lost. Retry %v/%v in %v\n", sc.connectionRetries, MAX_CONNECTION_RETRIES, RETRY_TIMEOUT)
			time.Sleep(RETRY_TIMEOUT)
			sc.Connect()
			return
		} else {
			log.Fatalln("Max retry count reached. Server not available. Exit -1")
		}
	}

	sc.comms = bufio.NewReadWriter(bufio.NewReader(connection), bufio.NewWriter(connection))

	sc.connectionRetries = 0

	sc.handleConnection(connection)
}

func (sc *SockClient[T]) Send(action Action) {
	msg, err := json.Marshal(action)
	if err != nil {
		log.Fatalf("Error marshalling action %v with error %v\n", action, err)
		return
	}

	bytesSent, sendError := sc.comms.WriteString(string(msg) + MSG_DELIM)
	if err != nil {
		log.Fatalf("Error sending action %v to server with error %v\n", action, sendError)
	}
	log.Printf("Sent %v bytes with message: %v\n", bytesSent, string(msg))
	sc.comms.Flush()
}

func (sc *SockClient[T]) ReadAction() Action {
	input, err := sc.comms.ReadString(byte('\n'))
	if err != nil {
		log.Fatalf("Error reading from server with %v\n", err)
	}

	var action Action
	parseErr := json.Unmarshal([]byte(input), &action)
	if parseErr != nil {
		log.Fatalf("Error parsing message from server with %v\n", err)
	}

	return action
}

func (sc *SockClient[T]) handleConnection(connection net.Conn) {
	defer connection.Close()

	for {
		receivedAction := sc.ReadAction()

		if !sc.Handler.Trigger(receivedAction) {
			continue
		}

		res, shouldProcess := sc.Handler.Request(receivedAction)
		if &res == nil || !shouldProcess {
			continue
		}

		actions := sc.Handler.Response(res)

		for _, action := range actions {
			action.SessionId = receivedAction.SessionId

			if sc.SendFn != nil {
				sc.SendFn(action)
			} else {
				sc.Send(action)
			}
		}
	}
}
