package shared

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Inventory int
	Position  Vector2
	Grounded  bool
	JumpTo    *float32
}
