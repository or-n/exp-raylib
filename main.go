package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	InitAudioDevice()
	MusicInit()
	InitWindow(1920, 1080, "Hello")
	defer CloseWindow()
	WindowSize = MonitorSize()
	ToggleFullscreen()
	SetTargetFPS(600)
	FontInit()
	MenuInit()
	SetExitKey(0)
	for !WindowShouldClose() && SimulationState != StateExit {
		if IsKeyDown(KeyEscape) {
			SimulationState = StateMenu
		}
		BeginDrawing()
		ClearBackground(WindowBg)
		switch SimulationState {
		case StateMenu:
			MenuDraw()
		case StateGame:
			MapUpdate()
			MapDraw()
		case StateOptions:
			OptionsDraw()
		}
		position := NewVector2(20, 25)
		size := NewVector2(100, 30)
		color := NewColor(0, 0, 0, 127)
		DrawRectangleV(position, size, color)
		DrawFPS(30, 30)
		EndDrawing()
		MusicUpdate()
	}
}
