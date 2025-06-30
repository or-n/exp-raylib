package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
	"slices"
)

var (
	points     = make([]Vector2, 50)
	attach     Vector2
	windowSize Vector2
	length     = float32(16)
	update     = func() {
		vw, vh := viewportSize()
		SetWindowSize(vw, vh)
		windowSize = NewVector2(float32(vw), float32(vh))
		attach = NewVector2(windowSize.X*0.5, 0)
		gravity()
		cursor := GetMousePosition()
		for range 500 {
			forward(cursor)
			backward(attach)
		}
		BeginDrawing()
		ClearBackground(Black)
		draw_points()
		EndDrawing()
	}
)

func init_points() {
	points[0] = attach
	for i := range points[1:] {
		points[i+1] = points[i]
		points[i+1].Y -= length
	}
}

func draw_points() {
	for _, p := range points {
		DrawCircleV(p, 2, White)
	}
	for i := range points[1:] {
		DrawLineV(points[i], points[i+1], White)
	}
}

func forward(target Vector2) {
	n := len(points)
	points[n-1] = target
	for i := range slices.Backward(points[:n-1]) {
		delta := Vector2Subtract(points[i+1], points[i])
		dir := Vector2Normalize(delta)
		scaled := Vector2Scale(dir, length)
		points[i] = Vector2Subtract(points[i+1], scaled)
	}
}

func backward(source Vector2) {
	points[0] = source
	for i := range points[1:] {
		delta := Vector2Subtract(points[i], points[i+1])
		dir := Vector2Normalize(delta)
		scaled := Vector2Scale(dir, length)
		points[i+1] = Vector2Subtract(points[i], scaled)
	}
}

func gravity() {
	for i := range points {
		points[i].Y += GetFrameTime() * 500
	}
}

func main() {
	InitWindow(0, 0, "")
	defer CloseWindow()
	init_points()
	SetTargetFPS(600)
	SetMainLoop(update)
	for !WindowShouldClose() {
		update()
	}
}
