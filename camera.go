package main

import . "github.com/gen2brain/raylib-go/raylib"
import "math"

var (
	MainCamera Camera2D
)

func CameraInit() {
	MainCamera.Zoom = 2.
}

func CameraUpdate() {
	wheel := GetMouseWheelMove()
    MainCamera.Offset = Vector2Scale(WindowSize, 0.5)
    scale := 0.2 * wheel
    zoom := float32(math.Exp(math.Log(float64(MainCamera.Zoom)) + float64(scale)))
    if zoom < 0.125 {
    	MainCamera.Zoom = 0.125
    } else if zoom > 64 {
    	MainCamera.Zoom = 64
    } else {
    	MainCamera.Zoom = zoom
    }
}
