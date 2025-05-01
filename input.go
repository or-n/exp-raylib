package main

import . "github.com/gen2brain/raylib-go/raylib"

var (
	InputNegY int32
	InputNegX int32
	InputY int32
	InputX int32
	InputJump int32
	InputSprint int32
)

func InputInit() {
	InputNegY = KeyW
	InputNegX = KeyA
	InputY = KeyS
	InputX = KeyD
	InputJump = KeyW
	InputSprint = KeyLeftShift
}

func boolI32(x bool) int32 {
	var result int32
	if x {
		result = 1
	}
	return result
}

func InputAxisX() int32 {
	return boolI32(IsKeyDown(InputX)) - boolI32(IsKeyDown(InputNegX))
}
