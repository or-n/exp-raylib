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
	ScreenSize = NewVector2(1920, 1080)
)

func ScreenPosition(p Vector2) Vector2 {
	return Vector2Multiply(p, ScreenSize)
}

func WorldPosition(p Vector2) Vector2 {
	return Vector2Divide(p, ScreenSize)
}

func main() {
	InitWindow(int32(ScreenSize.X), int32(ScreenSize.Y), "")
	defer CloseWindow()
	SetTargetFPS(600)
	ToggleFullscreen()
	MainFont := LoadFontEx(fontFile, 20, codepoints, int32(len(codepoints)))
	Bg := LoadTexture("asset/bg.jpg")
	for !WindowShouldClose() {
		CircuitUpdate()
		BeginDrawing()
		DrawTexture(Bg, 0, 0, White)
		DrawTextEx(MainFont, "", NewVector2(20, 20), 20, 2, White)
		CircuitDraw()
		DrawFPS(30, 30)
		EndDrawing()
	}
}
