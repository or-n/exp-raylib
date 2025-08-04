package shared

import (
	"fmt"
	. "github.com/gen2brain/raylib-go/raylib"
	. "github.com/or-n/util-go"
)

type Player struct {
	Inventory int
	Position  Vector2
	Grounded  bool
	JumpTo    *float32
}

func PlayerLoad(filename string, player *Player) {
	if err := Load(filename, player); err != nil {
		fmt.Println("Error loading player:", err)
		PlayerGen(player)
		PlayerSave(filename, player)
	}
}

func PlayerSave(filename string, player *Player) {
	if err := Save(filename, player); err != nil {
		fmt.Println("Failed to save player:", err)
	}
}

func PlayerGen(player *Player) {
	player.Position = NewVector2(0, f32(100*TextureY))
	player.Grounded = false
	player.JumpTo = nil
	player.Inventory = 0
}
