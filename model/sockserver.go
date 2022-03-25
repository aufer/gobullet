package model

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type SockServer struct {
	Config      SockClientConfig
	Connections []net.Conn
}

func (sc *SockServer) Listen() {
	fmt.Printf("Creating server with config %v\n", sc.Config)

	listener, err := net.Listen(sc.Config.Prot, sc.Config.ConnSrv())
	if err != nil {
		log.Fatalf("Error connection to server with %v\n", err)

	}

	defer listener.Close()

	for {
		con, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error on client connection %v\n", err)
		}

		sc.Connections = append(sc.Connections, con)

		handshake(con)
		go sc.handleConnection(&con)
	}
}

func (sc *SockServer) handleConnection(conn *net.Conn) {
	log.Printf("Waiting for message from %v\n", (*conn).RemoteAddr().String())

	buffer, err := bufio.NewReader(*conn).ReadString('\n')

	if err != nil {
		fmt.Println("Client disconnected.")
		(*conn).Close()
		return
	}

	log.Println("Client message:", string(buffer[:len(buffer)-1]))

	for _, connection := range sc.Connections {
		connection.Write([]byte(buffer))
	}

	sc.handleConnection(conn)
}

var handshakeMessage = Action{
	Type: "system:handshake",
}

func handshake(conn net.Conn) {
	log.Println("Say hello to ", conn)
	b, _ := json.Marshal(handshakeMessage)
	conn.Write(append(b, '\n'))
}
