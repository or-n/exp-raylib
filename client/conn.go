package main

import (
	"context"
	"encoding/gob"
	"github.com/coder/websocket"
	// . "github.com/or-n/util-go"
	"log"
	"net"
	. "shared"
	"syscall/js"
	"time"
)

var (
	MainConn net.Conn
	Incoming = make(chan Message, 32)
	Outgoing = make(chan Message, 32)
	Ip       string
)

const (
	local = "localhost"
)

func Remote() string {
	// 	os.Getenv("SERVER_IP")
	ip := js.Global().Get("SERVER_IP")
	if !ip.Truthy() {
		return local
	}
	return ip.String()
}

func ConnJoin() {
	SimulationState = StateJoining
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	url := "wss://" + Ip + "/ws"
	log.Println("Dialing WebSocket:", url)
	conn, _, err := websocket.Dial(ctx, url, nil)
	// 	conn, err := net.Dial("tcp", Ip+PortTCP)
	if err != nil {
		log.Println("WebSocket join error:", err)
		SimulationState = StateJoinError
		return
	}
	log.Printf("connected")
	MainConn = websocket.NetConn(context.Background(), conn, websocket.MessageBinary)
	SimulationState = StateGame
	Outgoing <- Message{Type: ClientGreet, Data: nil}
	go ConnReceive()
	go ConnSend()
}

func ConnReceive() {
	for {
		var msg Message
		Decoder := gob.NewDecoder(MainConn)
		if err := Decoder.Decode(&msg); err != nil {
			log.Println("Receiver error:", err)
			return
		}
		log.Printf("Received message: %+v\n", msg)
		switch msg.Type {
		case ServerGreet:
			data, ok := msg.Data.(MapData)
			if ok {
				Map = data.Map
				MapLoaded = true
			}
		case ServerChangeBlock:
			data, ok := msg.Data.(ChangeBlockData)
			if ok {
				Map[data.Y][data.X] = data.Block
			}
		}
	}
}

func ConnSend() {
	for msg := range Outgoing {
		Encoder := gob.NewEncoder(MainConn)
		if err := Encoder.Encode(msg); err != nil {
			log.Println("Sender error:", err)
			return
		}
	}
}

func ConnSendSingleplayer() {
	for msg := range Outgoing {
		response, _, err := Respond(msg, &Map)
		if err != nil {
			Incoming <- response
		}
	}
}
