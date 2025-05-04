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
	PlayerInit()
	MapInit()
	CameraInit()
	CursorInit()
	MenuInit()
	SetExitKey(0)
	for !WindowShouldClose() && SimulationState != StateExit {
		if IsKeyDown(KeyEscape) {
			SimulationState = StateMenu
		}
		BeginDrawing()
		ClearBackground(WindowBg)
		DrawFPS(30, 30)
		switch SimulationState {
		case StateMenu:
			ShowCursor()
			MenuDraw()
		case StateGame:
			HideCursor()
			CameraUpdate()
			PlayerUpdate()
			MainCamera.Target = PlayerCenter()
			BeginMode2D(MainCamera)
			MapDraw()
			PlayerDraw()
			CursorDraw()
			EndMode2D()
		case StateOptions:
			ShowCursor()
			OptionsUpdate()
			OptionsDraw()
		}
		EndDrawing()
		MusicUpdate()
	}
}
