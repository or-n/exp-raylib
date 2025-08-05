package main

import (
	"fmt"
	. "github.com/gen2brain/raylib-go/raylib"
	. "github.com/or-n/util-go"
	. "shared"
)

var (
	Multi  = true
	update = func() {
		// if IsKeyDown(KeyEscape) {
		// 	SimulationState = StateMenu
		// }
		for range 10 {
			handleIncomingOnce()
		}
		BeginDrawing()
		ClearBackground(WindowBg)
		switch SimulationState {
		case StateMenu:
			ShowCursor()
			MenuDraw()
		case StateJoining:
			text := fmt.Sprintf("Joining %s", Ip)
			DrawText(text, 400, 400, 20, White)
		case StateJoinError:
			text := fmt.Sprintf("Couldn't join %s", Ip)
			DrawText(text, 400, 400, 20, White)
		case StateGame:
			CameraUpdate()
			player := Players[PlayerName]
			PlayerUpdate(&player)
			MainCamera.Target = PlayerCenter(&player)
			BeginMode2D(MainCamera)
			MapDraw()
			PlayerDraw(&player)
			EndMode2D()
			PlayerOverlayDraw(&player)
			Players[PlayerName] = player
		case StateOptions:
			ShowCursor()
			OptionsUpdate()
			OptionsDraw()
		}
		DrawFPSTopLeft()
		EndDrawing()
		if SimulationState == StateExit {
			CloseWindow()
		}
	}
)

func DrawFPSTopLeft() {
	position := NewVector2(20, 25)
	size := NewVector2(100, 30)
	color := NewColor(0, 0, 0, 127)
	DrawRectangleV(position, size, color)
	DrawFPS(30, 30)
}

func main() {
	MessageRegister()
	InitAudioDevice()
	WindowSize = NewVector2(1920, 999)
	InitWindow(1920, 999, "multiplayer")
	// InitWindow(1920, 1080, "multiplayer")
	defer func() {
		// PlayerSave(PlayerFile, &MainPlayer)
		InputSave(InputFile, &Input)
		CloseWindow()
	}()
	// WindowSize = MonitorSize()
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
	if Multi {
		// Ip = Remote()
		Ip = local
		ConnJoin()
	} else {
		MapGen(&Map)
		MapLoaded = true
		SimulationState = StateGame
		go ConnSendSingleplayer()
	}
	SetExitKey(0)
	SetMainLoop(update)
	for !WindowShouldClose() {
		update()
	}
}
