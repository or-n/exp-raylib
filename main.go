package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

var (
	en         = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	pl         = []rune("ąćęłńóśźżĄĆĘŁŃÓŚŹŻ")
	symbols    = []rune(" .,!?:;_()/")
	codepoints = append(append(en, pl...), symbols...)
	fontFile   = "asset/FiraCode-Bold.ttf"
)

func main() {
	screenWidth := int32(1920)
	screenHeight := int32(1080)
	InitWindow(screenWidth, screenHeight, "")
	defer CloseWindow()
	SetTargetFPS(60)
	ToggleFullscreen()
	MainFont := LoadFontEx(fontFile, 20, codepoints, int32(len(codepoints)))
	Bg := LoadTexture("asset/bg.jpg")
	w := float32(screenWidth)
	h := float32(screenHeight)
	m := h * 20 / 1080
	r := Rectangle{}
	splitX1 := w * 0.25
	splitX2 := w * 0.5
	splitY1 := h * 0.25
	splitY2 := h * 0.5
	p := func(line float32) Vector2 {
		return NewVector2(r.X+m, r.Y+m+20*line)
	}
	changeSplit := func(split *bool, r Rectangle) {
		if CheckCollisionPointRec(GetMousePosition(), r) {
			if IsMouseButtonPressed(MouseButtonLeft) {
				*split = true
			}
		}
	}
	opacity := uint8(196)
	textColor := NewColor(255, 255, 255, 255)
	changeSplitX1 := false
	changeSplitX2 := false
	changeSplitY1 := false
	changeSplitY2 := false
	for !WindowShouldClose() {
		cursor := GetMousePosition()
		if IsMouseButtonDown(MouseButtonLeft) {
			if changeSplitX1 {
				splitX1 = cursor.X
			}
			if changeSplitX2 {
				splitX2 = cursor.X
			}
			if changeSplitY1 {
				splitY1 = cursor.Y
			}
			if changeSplitY2 {
				splitY2 = cursor.Y
			}
		} else {
			changeSplitX1 = false
			changeSplitX2 = false
			changeSplitY1 = false
			changeSplitY2 = false
		}
		changeSplit(&changeSplitX1, NewRectangle(splitX1-m/2, 0, m, splitY2+m/2))
		changeSplit(&changeSplitX2, NewRectangle(splitX2-m/2, 0, m, splitY2+m/2))
		changeSplit(&changeSplitY1, NewRectangle(splitX2-m/2, splitY1-m/2, w-(splitX2-m/2), m))
		changeSplit(&changeSplitY2, NewRectangle(0, splitY2-m/2, w, m))
		BeginDrawing()
		DrawTexture(Bg, 0, 0, White)
		r.X = m
		r.Y = m
		r.Width = splitX1 - m*3/2
		r.Height = splitY2 - m*3/2
		DrawRectangleRec(r, NewColor(0, 120, 255, opacity))
		DrawTextEx(MainFont, "pliki lokalne", p(0), 20, 2, textColor)
		r.X = splitX1 + m/2
		r.Y = m
		r.Width = (splitX2 - splitX1) - m
		r.Height = splitY2 - m*3/2
		DrawRectangleRec(r, NewColor(255, 0, 0, opacity))
		DrawTextEx(MainFont, "pliki serwerowe", p(0), 20, 2, textColor)
		r.X = splitX2 + m/2
		r.Y = m
		r.Width = (w - splitX2) - m*3/2
		r.Height = splitY1 - m*3/2
		DrawRectangleRec(r, NewColor(0, 0, 0, opacity))
		DrawTextEx(MainFont, "lista podstawowych komend", p(0), 20, 2, textColor)
		r.X = splitX2 + m/2
		r.Y = splitY1 + m/2
		r.Width = (w - splitX2) - m*3/2
		r.Height = (splitY2 - splitY1) - m
		DrawRectangleRec(r, NewColor(128, 128, 128, opacity))
		DrawTextEx(MainFont, "podstawowe informacje:", p(0), 20, 2, textColor)
		DrawTextEx(MainFont, "stan serwera: wł/wył", p(1), 20, 2, textColor)
		DrawTextEx(MainFont, "ilość osób", p(2), 20, 2, textColor)
		DrawTextEx(MainFont, "itp.", p(3), 20, 2, textColor)
		r.X = m
		r.Y = splitY2 + m/2
		r.Width = w - m*2
		r.Height = (h - splitY2) - m*3/2
		DrawRectangleRec(r, NewColor(255, 215, 0, opacity))
		DrawTextEx(MainFont, "logi", p(0), 20, 2, textColor)
		EndDrawing()
	}
}
