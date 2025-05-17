package main

import (
	. "exp-raylib/shared"
	. "github.com/or-n/util-go"
	"log"
	"net"
)

var (
	Joined   bool
	MainConn net.Conn
	Incoming = make(chan Message, 32)
	Outgoing = make(chan Message, 32)
)

func ConnJoin() {
	conn, err := net.Dial("tcp", ServerAddress())
	MainConn = conn
	if err != nil {
		SimulationState = StateMenu
	}
	Outgoing <- Message{Type: ClientGreet, Data: nil}
	Joined = true
	go ConnReceive()
	go ConnSend()
}

func ConnReceive() {
	for {
		var msg Message
		if err := FromSeq(MainConn, &msg); err != nil {
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
		if err := ToSeq(MainConn, msg); err != nil {
			log.Println("Sender error:", err)
			return
		}
	}
}
