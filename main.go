package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

var (
	WindowSize Vector2
	update     = func() {
		BeginDrawing()
		ClearBackground(Black)
		EndDrawing()
	}
)

// func main() {
// 	InitWindow(1920, 999, "path")
// 	defer CloseWindow()
// 	ToggleFullscreen()
// 	w, h := GetMonitorWidth(0), GetMonitorHeight(0)
// 	WindowSize.X = float32(w)
// 	WindowSize.Y = float32(h)
// 	SetTargetFPS(100)
// 	SetExitKey(0)
// 	SetMainLoop(update)
// 	for !WindowShouldClose() {
// 		update()
// 	}
// }
