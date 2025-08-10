package main

import (
	"fmt"
	. "github.com/gen2brain/raylib-go/raylib"
)

var (
	Actions    = 5
	Rewards    [][]float64
	Balance1   = float64(0)
	Balance2   = float64(0)
	History1   = []int32{}
	History2   = []int32{}
	scale      = float32(100)
	boxSize    = NewVector2(scale, scale)
	size       = Vector2Scale(boxSize, float32(Actions))
	maxHistory = 32
	maxReward  = int32(4)
	pad        = int32(1)
)

func GameInit() {
	Rewards = make([][]float64, Actions)
	for i := range Actions {
		Rewards[i] = make([]float64, Actions)
		for j := range Actions {
			Rewards[i][j] = float64(GetRandomValue(-maxReward, maxReward))
		}
	}
}

func GameUpdate() {
	start := Vector2Scale(Vector2Subtract(WindowSize, size), 0.5)
	if IsMouseButtonPressed(MouseButtonLeft) {
		cursor := GetMousePosition()
		rec := Rectangle{}
		rec.X = start.X
		rec.Y = start.Y
		rec.Width = size.X
		rec.Height = size.Y
		if CheckCollisionPointRec(cursor, rec) {
			delta := Vector2Subtract(cursor, start)
			i := GetRandomValue(0, int32(Actions)-1)
			j := int32(delta.X / scale)
			Balance1 += Rewards[i][j]
			Balance2 += Rewards[j][i]
			History1 = append(History1, j)
			History2 = append(History2, i)
			if len(History1) > maxHistory {
				History1 = History1[len(History1)-maxHistory:]
				History2 = History2[len(History2)-maxHistory:]
			}
		}
	}
}

func GameDraw() {
	start := Vector2Scale(Vector2Subtract(WindowSize, size), 0.5)
	cursor := GetMousePosition()
	rec := Rectangle{}
	rec.Width = boxSize.X
	rec.Height = boxSize.Y
	for i := range Actions {
		for j := range Actions {
			text := fmt.Sprintf("%.0f", Rewards[i][j])
			var color Color
			if Rewards[i][j] > 0 {
				v := Rewards[i][j] / float64(maxReward)
				color = ColorLerp(Yellow, Green, float32(v))
			} else {
				v := -Rewards[i][j] / float64(maxReward)
				color = ColorLerp(Yellow, Red, float32(v))
			}
			textSize := MeasureTextEx(GetFontDefault(), text, 20, 2)
			startBox := Vector2Scale(Vector2Subtract(boxSize, textSize), 0.5)
			rec.X = float32(j)*scale + start.X
			rec.Y = float32(i)*scale + start.Y
			p := pad
			if CheckCollisionPointRec(cursor, rec) {
				p = -1
			}
			s := int32(scale) - 2*p
			DrawRectangle(int32(rec.X)+p, int32(rec.Y)+p, s, s, color)
			DrawText(text, int32(rec.X+startBox.X), int32(rec.Y+startBox.Y), 20, White)
		}
	}
	text := fmt.Sprintf("%.0f", Balance1)
	textSize := MeasureTextEx(GetFontDefault(), text, 20, 2)
	start2 := Vector2Scale(Vector2Subtract(WindowSize, textSize), 0.5)
	DrawText(text, int32(start2.X), 100, 20, White)
	text = fmt.Sprintf("%.0f", Balance2)
	textSize = MeasureTextEx(GetFontDefault(), text, 20, 2)
	start2 = Vector2Scale(Vector2Subtract(WindowSize, textSize), 0.5)
	DrawText(text, int32(start2.X), 200, 20, White)
	for i := range History1 {
		action := History1[len(History1)-1-i]
		text := fmt.Sprintf("%d", action)
		DrawText(text, 100, 100+20*int32(i), 20, White)
	}
	for i := range History2 {
		action := History2[len(History2)-1-i]
		text := fmt.Sprintf("%d", action)
		DrawText(text, 200, 100+20*int32(i), 20, White)
	}
}
