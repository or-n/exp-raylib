package main

import . "github.com/gen2brain/raylib-go/raylib"
import "fmt"

var WindowBg = NewColor(127, 31, 255, 255)

func main() {
    InitWindow(1920, 1920, "Hello")
    defer CloseWindow()
    ToggleFullscreen()
    size := NewVector2(float32(GetMonitorWidth(0)), float32(GetMonitorHeight(0)))
    fmt.Println(size)
    SetTargetFPS(600)
    PlayerInit()
    InputInit()
    for !WindowShouldClose() {
        PlayerUpdate()
        BeginDrawing()
        ClearBackground(WindowBg)
        DrawFPS(30, 30)
        PlayerDraw()
        EndDrawing()
    }
}
