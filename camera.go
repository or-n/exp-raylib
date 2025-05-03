package main

import . "github.com/gen2brain/raylib-go/raylib"
import "math"

var (
	MainCamera Camera2D
)

func CameraInit() {
	MainCamera.Zoom = 2
}

func CameraRestart() {
	MainCamera.Zoom = 2
}

func CameraUpdate() {
	wheel := GetMouseWheelMove()
	MainCamera.Offset = Vector2Scale(WindowSize, 0.5)
	scale := 0.2 * wheel
	zoom := float32(math.Exp(math.Log(float64(MainCamera.Zoom)) + float64(scale)))
	if zoom < 0.125 {
		MainCamera.Zoom = 0.125
	} else {
		MainCamera.Zoom = min(zoom, 64)
	}
}

func CameraRect(offset float32) Rectangle {
	r := Rectangle{}
	r.Width = WindowSize.X / MainCamera.Zoom
	r.Height = WindowSize.Y / MainCamera.Zoom
	r.X = MainCamera.Target.X - r.Width*0.5 + offset
	r.Y = MainCamera.Target.Y - r.Height*0.5 + offset
	r.Width -= 2 * offset
	r.Height -= 2 * offset
	return r
}
