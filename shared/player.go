package shared

import (
	"fmt"
	. "github.com/gen2brain/raylib-go/raylib"
	. "github.com/or-n/util-go"
)

var (
	PlayerSize = NewVector2(8, 12)
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
	player.Position = NewVector2(0, f32(50*TextureY))
	player.Grounded = false
	player.JumpTo = nil
	player.Inventory = 0
}

func PlayerGetRect(position Vector2) Rectangle {
	rec := Rectangle{}
	rec.X = position.X
	rec.Y = position.Y
	rec.Width = PlayerSize.X
	rec.Height = PlayerSize.Y
	return rec
}
