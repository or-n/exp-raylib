package main

import . "github.com/gen2brain/raylib-go/raylib"

var cursorTexture Texture2D

func CursorInit() {
	cursorTexture = LoadTexture("asset/player.png")
	HideCursor()
}

func CursorPosition() Vector2 {
	return GetScreenToWorld2D(GetMousePosition(), MainCamera)
}

func CursorDraw() {
	p := CursorPosition()
	size := NewVector2(float32(cursorTexture.Width), float32(cursorTexture.Height))
	offset := Vector2Scale(size, 0.5)
	DrawTextureV(cursorTexture, Vector2Subtract(p, offset), White)
}
