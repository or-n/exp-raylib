package shared

import (
	"encoding/gob"
	"fmt"
	. "github.com/gen2brain/raylib-go/raylib"
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
	ClientChangePlayer
	ServerChangePlayer
)

type Message struct {
	Type MessageType
	Data any
}

type InitData struct {
	Map    [MaxY][MaxX]Block
	Player Player
}

type JoinData struct {
	Name string
}

type PlayerChangeBlock struct {
	Name   string
	Change ChangeBlockData
}

type ChangeBlockData struct {
	X, Y  int
	Block Block
}

type PlayerData struct {
	Name   string
	Player Player
}

func MessageRegister() {
	gob.Register(JoinData{})
	gob.Register(InitData{})
	gob.Register(ChangeBlockData{})
	gob.Register(PlayerData{})
	gob.Register(PlayerChangeBlock{})
}

func Respond(msg Message, Map *[MaxY][MaxX]Block, players map[string]Player) ([]Message, bool, error) {
	switch msg.Type {
	case ClientGreet:
		if data, ok := msg.Data.(JoinData); ok {
			var player Player
			PlayerLoad("data/"+data.Name+".gob", &player)
			players[data.Name] = player
			response := Message{
				Type: ServerGreet,
				Data: InitData{
					Map:    *Map,
					Player: player,
				},
			}
			return []Message{response}, false, nil
		}
		return []Message{}, false, fmt.Errorf("ClientGreet data")
	case ClientChangeBlock:
		if data, ok := msg.Data.(PlayerChangeBlock); ok {
			change := data.Change
			x, y := change.X, change.Y
			player := players[data.Name]
			ok := false
			if change.Block == Dirt {
				if Map[y][x] == Empty && player.Inventory > 0 {
					r := MapRect(x, y)
					p := PlayerGetRect(player.Position)
					if CheckCollisionRecs(p, r) {
						return []Message{}, false, nil
					}
					player.Inventory -= 1
					ok = true
				}
			} else {
				if Map[y][x] != Empty {
					player.Inventory += 1
					ok = true
				}
			}
			if ok {
				Map[y][x] = change.Block
				players[data.Name] = player
				response1 := Message{
					Type: ServerChangeBlock,
					Data: ChangeBlockData{X: x, Y: y, Block: change.Block},
				}
				response2 := Message{
					Type: ServerChangePlayer,
					Data: PlayerData{Name: data.Name, Player: player},
				}
				return []Message{response1, response2}, true, nil
			} else {
				return []Message{}, false, nil
			}
		}
		return []Message{}, false, fmt.Errorf("ClientChangeBlock data")
	case ClientChangePlayer:
		if data, ok := msg.Data.(PlayerData); ok {
			players[data.Name] = data.Player
			return []Message{}, false, nil
		}
	}
	return []Message{}, false, fmt.Errorf("no match")
}
