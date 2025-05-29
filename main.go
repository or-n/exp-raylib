package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
	"math"
	"strconv"
)

var (
	ScreenSize    = NewVector2(1920, 1080)
	distribution1 = []float32{0.25, 0.5}
	distribution2 = []float32{1 / 3., 2 / 3.}
	offset1       = make([]float32, 2)
	offset2       = make([]float32, 2)
	change1       = make([]bool, 2)
	change2       = make([]bool, 2)
	m             = float32(20)
	epsilon       = float32(1.1920929e-07)
)

func similarity(distribution1, distribution2 []float32) float32 {
	var s float32
	for i := range distribution1 {
		d1 := distribution1[i] + epsilon
		d2 := distribution2[i] + epsilon
		ratio := float64(d1 / d2)
		s += distribution1[i] * float32(math.Log2(ratio))
	}
	return s
}

func symmetric_similarity() float32 {
	mix := make([]float32, len(distribution1))
	for i := range distribution1 {
		mix[i] = (distribution1[i] + distribution2[i]) * 0.5
	}
	return (similarity(distribution1, mix) + similarity(distribution2, mix)) * 0.5
}

func color(i int) Color {
	switch i {
	case 0:
		return Red
	case 1:
		return Green
	case 2:
		return Blue
	default:
		return White
	}
}

func draw_distribution(distribution, offset []float32, change []bool, y float32) {
	cursor := GetMousePosition()
	r := Rectangle{}
	r.Y = y
	r.Height = 100
	for i := range distribution {
		start := float32(0)
		if i > 0 {
			start = distribution[i-1]
		}
		r.X = start * ScreenSize.X
		end := distribution[i]
		r.Width = (end - start) * ScreenSize.X
		DrawRectangleRec(r, color(i))
		r.X = distribution[i]*ScreenSize.X - m
		r.Width = 2 * m
		DrawRectangleRec(r, White)
		if CheckCollisionPointRec(cursor, r) && IsMouseButtonPressed(MouseButtonLeft) {
			change[i] = true
			offset[i] = distribution[i] - cursor.X/ScreenSize.X
		}
	}
	last := len(distribution) - 1
	start := distribution[last]
	r.X = start * ScreenSize.X
	r.Width = (1 - start) * ScreenSize.X
	DrawRectangleRec(r, color(len(distribution)))
}

func main() {
	InitWindow(int32(ScreenSize.X), int32(ScreenSize.Y), "")
	defer CloseWindow()
	SetTargetFPS(600)
	ToggleFullscreen()
	for !WindowShouldClose() {
		cursor := GetMousePosition()
		if IsMouseButtonDown(MouseButtonLeft) {
			x := cursor.X / ScreenSize.X
			for i := range change1 {
				if change1[i] {
					distribution1[i] = x + offset1[i]
				}
				if change2[i] {
					distribution2[i] = x + offset2[i]
				}
			}
		} else {
			change1 = make([]bool, len(distribution1))
			change2 = make([]bool, len(distribution2))
		}
		BeginDrawing()
		ClearBackground(Blank)
		draw_distribution(distribution1, offset1, change1, 0)
		draw_distribution(distribution2, offset2, change2, 200)
		result := symmetric_similarity()
		text := strconv.FormatFloat(float64(result), 'f', -1, 32)
		DrawText(text, int32(ScreenSize.X)/2, int32(ScreenSize.Y)/2, 20, White)
		DrawFPS(30, 30)
		EndDrawing()
	}
}
