package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

var (
	WindowSize Vector2
	update     = func() {
		GameUpdate()
		BeginDrawing()
		ClearBackground(Black)
		GameDraw()
		EndDrawing()
	}
)

func main() {
	InitWindow(1920, 1080, "gen game")
	defer func() {
		CloseWindow()
	}()
	ToggleFullscreen()
	w, h := GetMonitorWidth(0), GetMonitorHeight(0)
	WindowSize.X = float32(w)
	WindowSize.Y = float32(h)
	GameInit()
	SetTargetFPS(100)
	// SetExitKey(0)
	SetMainLoop(update)
	for !WindowShouldClose() {
		update()
	}
}
