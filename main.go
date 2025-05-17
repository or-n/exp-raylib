package main

import (
	"fmt"
	. "github.com/gen2brain/raylib-go/raylib"
	. "github.com/or-n/util-go"
)

func main() {
	InitAudioDevice()
	MusicInit()
	InitWindow(1920, 1080, "Hello")
	defer func() {
		if err := Save(AccountFile, &MainAccount); err != nil {
			fmt.Println("Failed to save account:", err)
		}
		CloseWindow()
	}()
	WindowSize = MonitorSize()
	ToggleFullscreen()
	SetTargetFPS(600)
	FontInit()
	MenuInit()
	EventNew()
	AccountInit()
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
			AccountUpdate(&MainAccount)
			EventDraw()
			AccountDraw(&MainAccount)
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
