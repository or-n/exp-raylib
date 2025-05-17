package main

import (
	. "exp-raylib/shared"
	. "github.com/or-n/util-go"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

var (
	MapFile       = "data/map.gob"
	ServerMap     [MaxY][MaxX]Block
	ActivePlayers = make(map[net.Conn]*Player)
	mu            sync.Mutex
)

func handleConn(conn net.Conn) {
	defer func() {
		mu.Lock()
		delete(ActivePlayers, conn)
		mu.Unlock()
		conn.Close()
	}()
	player := &Player{}
	mu.Lock()
	ActivePlayers[conn] = player
	mu.Unlock()
	for {
		var msg Message
		if err := FromSeq(conn, &msg); err != nil {
			if err == io.EOF || strings.Contains(err.Error(), "connection reset by peer") {
				log.Println("Client disconnected")
				return
			}
			log.Println("Error decoding message:", err)
			return
		}
		log.Printf("Received message: %+v\n", msg)
		switch msg.Type {
		case ClientGreet:
			response := Message{
				Type: ServerGreet,
				Data: MapData{
					Map: ServerMap,
				},
			}
			err := ToSeq(conn, response)
			if err != nil {
				log.Println("Error sending ServerGreet:", err)
			}
			log.Println("Sent ServerGreet")
		case ClientChangeBlock:
			if data, ok := msg.Data.(ChangeBlockData); ok {
				ServerMap[data.Y][data.X] = data.Block
				response := Message{
					Type: ServerChangeBlock,
					Data: msg.Data,
				}
				Broadcast(response)
				log.Println("Sent ServerChangeBlock")
			}
		}
	}
}

func Broadcast(msg Message) {
	mu.Lock()
	defer mu.Unlock()
	for conn := range ActivePlayers {
		err := ToSeq(conn, msg)
		if err != nil {
			log.Println("Broadcast error:", err)
		}
	}
}

func main() {
	MessageRegister()
	MapLoad(MapFile, &ServerMap)
	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			MapSave(MapFile, &ServerMap)
			log.Println("Map backup saved")
		}
	}()
	ln, err := net.Listen("tcp", ServerPort)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	log.Println("Server is listening on port 1234")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}
		go handleConn(conn)
	}
}
