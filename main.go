package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
    rl.InitWindow(800, 600, "Hello Raylib-Go")
    defer rl.CloseWindow()

    for !rl.WindowShouldClose() {
        rl.BeginDrawing()
        rl.ClearBackground(rl.RayWhite)
        rl.DrawText("Hello, Raylib in Go!", 190, 200, 20, rl.DarkGray)
        rl.EndDrawing()
    }
}
