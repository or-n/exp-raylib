package main

import . "github.com/gen2brain/raylib-go/raylib"
import (
	"math/rand"
	"time"
)

var WindowBg = NewColor(127, 31, 255, 255)

func main() {
	rand.Seed(time.Now().UnixNano())
	InitAudioDevice()
	MusicInit()
	InitWindow(1920, 1920, "Hello")
	defer CloseWindow()
	ToggleFullscreen()
	WindowInit()
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
		switch SimulationState {
		case StateMenu:
			ShowCursor()
			BeginDrawing()
			ClearBackground(WindowBg)
			DrawFPS(30, 30)
			MenuDraw()
			EndDrawing()
		case StateGame:
			HideCursor()
			CameraUpdate()
			PlayerUpdate()
			MainCamera.Target = PlayerCenter()
			BeginDrawing()
			ClearBackground(WindowBg)
			DrawFPS(30, 30)
			BeginMode2D(MainCamera)
			MapDraw()
			PlayerDraw()
			CursorDraw()
			EndMode2D()
			EndDrawing()
		case StateOptions:
			ShowCursor()
			BeginDrawing()
			ClearBackground(WindowBg)
			DrawFPS(30, 30)
			InputDraw()
			EndDrawing()
		}
		InputUpdate()
		MusicUpdate()
	}
}
