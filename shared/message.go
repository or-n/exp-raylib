package shared

import (
	"encoding/gob"
)

type MessageType int

const (
	PortTCP = ":1234"
	PortWs  = ":1235"
)

const (
	ClientGreet MessageType = iota
	ServerGreet
	ClientChangeBlock
	ServerChangeBlock
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
