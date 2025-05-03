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
    InputInit()
    MapInit()
    CameraInit()
    CursorInit()
    for !WindowShouldClose() {
    	MusicUpdate()
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
    }
}
