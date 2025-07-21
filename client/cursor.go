package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

var (
	cursorTexture Texture2D
	mouseMoved    bool
)

func CursorInit() {
	cursorTexture = LoadTexture("asset/dirt.png")
}

func CursorPosition() Vector2 {
	return GetScreenToWorld2D(GetMousePosition(), MainCamera)
}

func CursorDraw() {
	cursorPosition := CursorPosition()
	if !mouseMoved {
		if cursorPosition.X > 0 || cursorPosition.Y > 0 {
			mouseMoved = true
		} else {
			return
		}
	}
	size := NewVector2(f32(cursorTexture.Width), f32(cursorTexture.Height))
	scale := f32(0.25)
	offset := Vector2Scale(size, 0.5*scale)
	p := Vector2Subtract(cursorPosition, offset)
	DrawTextureEx(cursorTexture, p, 0, scale, White)
}
