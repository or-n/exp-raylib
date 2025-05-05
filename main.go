package main

import (
	"fmt"
	. "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	InitAudioDevice()
	MusicInit()
	InitWindow(1920, 1080, "Hello")
	defer func() {
		if err := Save(MapFile, Map); err != nil {
			fmt.Println("Failed to save map:", err)
		}
		if err := Save(PlayerFile, MainPlayer); err != nil {
			fmt.Println("Failed to save player:", err)
		}
		CloseWindow()
	}()
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
		switch SimulationState {
		case StateMenu:
			ShowCursor()
			MenuDraw()
		case StateGame:
			HideCursor()
			CameraUpdate()
			PlayerUpdate(&MainPlayer)
			MainCamera.Target = PlayerCenter(&MainPlayer)
			BeginMode2D(MainCamera)
			MapDraw()
			PlayerDraw(&MainPlayer)
			CursorDraw()
			EndMode2D()
			PlayerOverlayDraw(&MainPlayer)
		case StateOptions:
			ShowCursor()
			OptionsUpdate()
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
