package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

var (
	WindowSize Vector2
	WindowBg   = NewColor(127, 31, 255, 255)
)

func MonitorSize() Vector2 {
	return NewVector2(f32(GetMonitorWidth(0)), f32(GetMonitorHeight(0)))
}
