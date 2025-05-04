package main

import . "github.com/gen2brain/raylib-go/raylib"

var WindowSize Vector2

func WindowInit() {
	WindowSize = NewVector2(f32(GetMonitorWidth(0)), f32(GetMonitorHeight(0)))
}
