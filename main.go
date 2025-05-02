package main

import . "github.com/gen2brain/raylib-go/raylib"
import (
	"math/rand"
	"time"
)

var WindowBg = NewColor(127, 31, 255, 255)

func main() {
    rand.Seed(time.Now().UnixNano())
    InitWindow(1920, 1920, "Hello")
    defer CloseWindow()
    ToggleFullscreen()
    // size := NewVector2(float32(GetMonitorWidth(0)), float32(GetMonitorHeight(0)))
    SetTargetFPS(600)
    PlayerInit()
    InputInit()
    MapInit()
    for !WindowShouldClose() {
        PlayerUpdate()
        BeginDrawing()
        ClearBackground(WindowBg)
        DrawFPS(30, 30)
        MapDraw()
        PlayerDraw()
        EndDrawing()
    }
}
