package main

import (
	"fmt"
	. "github.com/gen2brain/raylib-go/raylib"
	. "github.com/or-n/util-go"
)

func main() {
	InitAudioDevice()
	Music := LoadMusicStream("asset/TetrisThemeDubstep.ogg")
	PlayMusicStream(Music)
	InitWindow(1920, 1080, "Staker")
	defer func() {
		if err := Save(AccountFile, &MainAccount); err != nil {
			fmt.Println("Failed to save account:", err)
		}
		CloseWindow()
	}()
	WindowSize = MonitorSize()
	ToggleFullscreen()
	SetTargetFPS(100)
	FontInit()
	MenuInit()
	EventNew()
	AccountInit()
	GifInit()
	SetExitKey(0)
	for !WindowShouldClose() && SimulationState != StateExit {
		SetMusicVolume(Music, MusicVolume)
		UpdateMusicStream(Music)
		if IsKeyDown(KeyEscape) {
			SimulationState = StateMenu
		}
		BeginDrawing()
		ClearBackground(WindowBg)
		switch SimulationState {
		case StateMenu:
			MenuDraw()
		case StateGame:
			GifUpdate()
			AccountUpdate(&MainAccount)
			GifDraw()
			EventDraw()
			AccountDraw(&MainAccount)
		case StateOptions:
			OptionsDraw()
		}
		EndDrawing()
	}
}
