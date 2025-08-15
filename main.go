package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
	"time"
)

var (
	WindowSize Vector2
	g          Graph
	update     = func() {
		BeginDrawing()
		ClearBackground(Black)
		g.GraphDraw()
		EndDrawing()
	}
)

func main() {
	InitWindow(1920, 999, "dupa")
	defer CloseWindow()
	// ToggleFullscreen()
	// w, h := GetMonitorWidth(0), GetMonitorHeight(0)
	// WindowSize.X = float32(w)
	// WindowSize.Y = float32(h)
	WindowSize.X = 1920
	WindowSize.Y = 999
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	g = new(128, rng)
	SetTargetFPS(100)
	SetExitKey(0)
	SetMainLoop(update)
	for !WindowShouldClose() {
		update()
	}
}
