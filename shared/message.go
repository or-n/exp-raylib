package shared

import (
	"encoding/gob"
	"fmt"
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

func Respond(msg Message, Map *[MaxY][MaxX]Block) (Message, bool, error) {
	switch msg.Type {
	case ClientGreet:
		response := Message{
			Type: ServerGreet,
			Data: MapData{
				Map: *Map,
			},
		}
		return response, false, nil
	case ClientChangeBlock:
		if data, ok := msg.Data.(ChangeBlockData); ok {
			Map[data.Y][data.X] = data.Block
			response := Message{
				Type: ServerChangeBlock,
				Data: msg.Data,
			}
			return response, true, nil
		}
		return Message{}, false, fmt.Errorf("ClientChangeBlock data")
	}
	return Message{}, false, fmt.Errorf("no match")
}
