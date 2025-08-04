package main

import (
	"context"
	"github.com/coder/websocket"
	. "github.com/or-n/util-go"
	"io"
	"log"
	"net"
	"net/http"
	. "shared"
	"strings"
	"sync"
	"time"
)

var (
	MapFile       = "data/map.gob"
	ServerMap     [MaxY][MaxX]Block
	ActivePlayers = make(map[net.Conn]string)
	Players       = make(map[string]Player)
	mu            sync.Mutex
)

func Send(conn net.Conn, msg Message) {
	err := ToSeq(conn, msg)
	if err != nil {
		log.Println("Error sending msg:", err)
	}
	log.Println("Sent msg")
}

func handleConn(conn net.Conn) {
	log.Println("handling conn")
	defer func() {
		mu.Lock()
		name := ActivePlayers[conn]
		player := Players[name]
		PlayerSave("data/"+name+".gob", &player)
		log.Println(name + " player backup saved")
		delete(ActivePlayers, conn)
		mu.Unlock()
		log.Println("closing conn")
		conn.Close()
	}()
	for {
		var msg Message
		if err := FromSeq(conn, &msg); err != nil {
			reset := strings.Contains(err.Error(), "connection reset by peer")
			if err == io.EOF || reset {
				log.Println("Client disconnected")
				return
			}
			log.Println("Error decoding message:", err)
			return
		}
		log.Printf("Received message: %+v\n", msg)
		if msg.Type == ClientGreet {
			if data, ok := msg.Data.(JoinData); ok {
				mu.Lock()
				ActivePlayers[conn] = data.Name
				mu.Unlock()
			}
		}
		responses, broadcast, err := Respond(msg, &ServerMap, Players)
		if err != nil {
			continue
		}
		if broadcast {
			for _, response := range responses {
				Broadcast(response)
			}
		} else {
			for _, response := range responses {
				Send(conn, response)
			}
		}
	}
}

func handleWS(w http.ResponseWriter, r *http.Request) {
	log.Println("Ws connection try")
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Println("WebSocket accept error:", err)
		return
	}
	log.Println("WebSocket client connected")
	go func() {
		conn := websocket.NetConn(context.Background(), c, websocket.MessageBinary)
		handleConn(conn)
	}()
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
		ticker := time.NewTicker(time.Second * 15)
		defer ticker.Stop()
		for range ticker.C {
			MapSave(MapFile, &ServerMap)
			log.Println("Map backup saved")
			for name, player := range Players {
				PlayerSave("data/"+name+".gob", &player)
				log.Println(name + " player backup saved")
			}
		}
	}()
	go func() {
		ln, err := net.Listen("tcp", PortTCP)
		if err != nil {
			log.Fatal(err)
		}
		defer ln.Close()
		log.Println("TCP server listening on", PortTCP)
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Println("Accept error:", err)
				continue
			}
			go handleConn(conn)
		}
	}()
	log.Println("Ws server is listening on", PortWs)
	http.HandleFunc("/ws", handleWS)
	err := http.ListenAndServe(PortWs, nil)
	if err != nil {
		log.Fatal("HTTP server error:", err)
	}
}
