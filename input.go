package main

import (
	"fmt"
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
	InputFile = "asset/input.gob"
	Input     map[Action]i32
	keys      = []i32{
		KeyW,
		KeyA,
		KeyS,
		KeyD,
		KeyLeftShift,
		KeyLeftControl,
		KeySpace,
	}
)

func InputInit() {
	if err := Load(InputFile, &Input); err != nil {
		fmt.Println("Error loading input:", err)
		Input = map[Action]i32{
			ActionNegY:   KeyW,
			ActionNegX:   KeyA,
			ActionY:      KeyS,
			ActionX:      KeyD,
			ActionJump:   KeyW,
			ActionSneak:  KeyLeftShift,
			ActionSprint: KeyLeftControl,
		}
		if err := Save(InputFile, Input); err != nil {
			fmt.Println("Failed to save input:", err)
		}
	}
}

func InputAxisX() i32 {
	return BoolI32(IsKeyDown(Input[ActionX])) - BoolI32(IsKeyDown(Input[ActionNegX]))
}
