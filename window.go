package main

import . "github.com/gen2brain/raylib-go/raylib"

var WindowSize Vector2

func WindowInit() {
    WindowSize = NewVector2(float32(GetMonitorWidth(0)), float32(GetMonitorHeight(0)))
}
