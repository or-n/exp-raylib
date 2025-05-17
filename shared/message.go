package shared

import (
	"encoding/gob"
	"fmt"
	"os"
)

type MessageType int

const (
	ClientGreet MessageType = iota
	ServerGreet
	ClientChangeBlock
	ServerChangeBlock
)

const (
	ServerPort = ":1234"
)

type Message struct {
	Type MessageType
	Data any
}

type MapData struct {
	Map [MaxY][MaxX]Block
}

type ChangeBlockData struct {
	X, Y  int
	Block Block
}

func MessageRegister() {
	gob.Register(MapData{})
	gob.Register(ChangeBlockData{})
}

func ServerAddress() string {
	ip := os.Getenv("SERVER_IP")
	if ip == "" {
		ip = "localhost"
	}
	return fmt.Sprintf("%s%s", ip, ServerPort)
}
