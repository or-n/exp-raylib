package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

type Action int

const (
	ActionNegY Action = iota
	ActionNegX
	ActionY
	ActionX
	ActionJump
	ActionSneak
	ActionSprint
)

func actionString(a Action) string {
	switch a {
	case ActionNegY:
		return "neg_y"
	case ActionNegX:
		return "neg_x"
	case ActionY:
		return "y"
	case ActionX:
		return "x"
	case ActionJump:
		return "jump"
	case ActionSneak:
		return "sneak"
	case ActionSprint:
		return "sprint"
	}
	return ""
}

func KeyString(x i32) string {
	switch x {
	case KeyW:
		return "W"
	case KeyA:
		return "A"
	case KeyS:
		return "S"
	case KeyD:
		return "D"
	case KeyLeftShift:
		return "Left Shift"
	case KeyLeftControl:
		return "Left Control"
	case KeySpace:
		return "Space"
	}
	return ""
}

var (
	Input = map[Action]i32{
		ActionNegY:   KeyW,
		ActionNegX:   KeyA,
		ActionY:      KeyS,
		ActionX:      KeyD,
		ActionJump:   KeyW,
		ActionSneak:  KeyLeftShift,
		ActionSprint: KeyLeftControl,
	}
	keys = []i32{
		KeyW,
		KeyA,
		KeyS,
		KeyD,
		KeyLeftShift,
		KeyLeftControl,
		KeySpace,
	}
)

func InputAxisX() i32 {
	return BoolI32(IsKeyDown(Input[ActionX])) - BoolI32(IsKeyDown(Input[ActionNegX]))
}
