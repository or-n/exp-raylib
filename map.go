package main

import (
	. "exp-raylib/shared"
	. "github.com/gen2brain/raylib-go/raylib"
)

var (
	MapLoaded   bool
	Map         [MaxY][MaxX]Block
	dirtTexture Texture2D
	texture_x   = i32(16)
	texture_y   = i32(16)
	offset_x    = -i32(MaxX) * texture_x / 2
	offset_y    = i32(0)
)

func MapInit() {
	dirtTexture = LoadTexture("asset/dirt.png")
}

func MapCollide(rec *Rectangle) bool {
	t := RectangleInt32{}
	t.Width = texture_x
	t.Height = texture_y
	center := RectCenter(*rec)
	cx, cy := MapIndex(center)
	n := 9
	for y := range MaxY {
		iy := y - n/2 + cy
		if !MapInsideY(iy) {
			continue
		}
		position_y := i32(iy)*texture_y + offset_y
		t.Y = position_y
		for x := range n {
			ix := x - n/2 + cx
			if !MapInsideX(ix) {
				continue
			}
			position_x := i32(ix)*texture_x + offset_x
			t.X = position_x
			if Map[iy][ix] == Dirt {
				tile := t.ToFloat32()
				if CheckCollisionRecs(tile, *rec) {
					return true
				}
			}
		}
	}
	return false
}

func RectCenter(r Rectangle) Vector2 {
	return NewVector2(r.X+r.Width*0.5, r.Y+r.Height*0.5)
}

func MapDraw() {
	cameraRect := CameraRect(0)
	rect := Rectangle{}
	rect.Width = f32(texture_x)
	rect.Height = f32(texture_y)
	center := RectCenter(cameraRect)
	cx, cy := MapIndex(center)
	n := 61
	for y := range n {
		iy := y - n/2 + cy
		if !MapInsideY(iy) {
			continue
		}
		position_y := i32(iy)*texture_y + offset_y
		rect.Y = f32(position_y)
		for x := range n {
			ix := x - n/2 + cx
			if !MapInsideX(ix) {
				continue
			}
			position_x := i32(ix)*texture_x + offset_x
			rect.X = f32(position_x)
			if Map[iy][ix] == Dirt {
				if CheckCollisionRecs(rect, cameraRect) {
					DrawTexture(dirtTexture, position_x, position_y, White)
				}
			}
		}
	}
}

func MapIndex(position Vector2) (int, int) {
	x := (i32(position.X) - offset_x) / texture_x
	y := (i32(position.Y) - offset_y) / texture_y
	return int(x), int(y)
}

func MapInsideX(x int) bool {
	return x >= 0 && x < MaxX
}

func MapInsideY(y int) bool {
	return y >= 0 && y < MaxY
}

func MapRect(x, y int) Rectangle {
	r := Rectangle{}
	r.X = f32(i32(x)*texture_x + offset_x)
	r.Y = f32(i32(y)*texture_y + offset_y)
	r.Width = f32(texture_x)
	r.Height = f32(texture_y)
	return r
}
