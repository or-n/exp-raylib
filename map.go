package main

import . "github.com/gen2brain/raylib-go/raylib"
import "math/rand"

const (
	MaxX = 100
	MaxY = 10
)

var (
	data        [MaxY][MaxX]int
	dirtTexture Texture2D
	texture_x   int32
	texture_y   int32
	offset_x    int32
	offset_y    int32
)

func MapInit() {
	for y := range MaxY {
		for x := range MaxX {
			data[y][x] = rand.Intn(2)
		}
	}
	dirtTexture = LoadTexture("asset/dirt.png")
	texture_x = 16
	texture_y = 16
	offset_x = -int32(MaxX) * texture_x / 2
}

func MapCollide(rec *Rectangle) bool {
	t := RectangleInt32{}
	t.Width = texture_x
	t.Height = texture_y
	for y := range MaxY {
		position_y := int32(y)*texture_y + offset_y
		t.Y = position_y
		for x := range MaxX {
			if data[y][x] == 1 {
				position_x := int32(x)*texture_x + offset_x
				t.X = position_x
				tile := t.ToFloat32()
				if CheckCollisionRecs(tile, *rec) {
					return true
				}
			}
		}
	}
	return false
}

func MapDraw() {
	cameraRect := CameraRect(0)
	rect := Rectangle{}
	rect.Width = float32(texture_x)
	rect.Height = float32(texture_y)
	for y := range MaxY {
		position_y := int32(y)*texture_y + offset_y
		rect.Y = float32(position_y)
		for x := range MaxX {
			if data[y][x] == 1 {
				position_x := int32(x)*texture_x + offset_x
				rect.X = float32(position_x)
				if CheckCollisionRecs(rect, cameraRect) {
					DrawTexture(dirtTexture, position_x, position_y, White)
				}
			}
		}
	}
}

func MapIndex(position Vector2) (int, int) {
	x := (int32(position.X) - offset_x) / texture_x
	y := (int32(position.Y) - offset_y) / texture_y
	return int(x), int(y)
}
