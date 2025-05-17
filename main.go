package main

import (
	. "exp-raylib/shared"
	. "github.com/gen2brain/raylib-go/raylib"
	. "github.com/or-n/util-go"
)

func main() {
	MessageRegister()
	InitAudioDevice()
	MusicInit()
	InitWindow(1920, 1080, "Hello")
	defer func() {
		PlayerSave(PlayerFile, &MainPlayer)
		InputSave(InputFile, &Input)
		CloseWindow()
	}()
	WindowSize = MonitorSize()
	ToggleFullscreen()
	SetTargetFPS(600)
	InputLoad(InputFile, &Input)
	PlayerInit()
	NoiseInit()
	MapInit()
	CameraInit()
	CursorInit()
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
			ShowCursor()
			MenuDraw()
		case StateGame:
			if Joined {
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
			} else {
				ConnJoin()
				ShowCursor()
				DrawText("Joining", 400, 400, 20, White)
			}
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
