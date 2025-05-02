package main

import . "github.com/gen2brain/raylib-go/raylib"
import "math/rand"

const (
	max_x = 100
	max_y = 10
)

var (
	data [max_y][max_x]int
	texture Texture2D
	texture_x int32
	texture_y int32
	offset_x int32
	offset_y int32
)

func MapInit() {
	for y := 0; y < max_y; y++ {
		for x := 0; x < max_x; x++ {
			data[y][x] = rand.Intn(2)
		}
	}
	texture = LoadTexture("asset/dirt.png")
	texture_x = 16
	texture_y = 16
	offset_x = -int32(max_x) * texture_x / 2
}

func MapCollide(rec *Rectangle, speed Vector2, up, down *bool) {
	t := RectangleInt32{}
	t.Width = texture_x
	t.Height = texture_y
	for y := 0; y < max_y; y++ {
		position_y := int32(y) * texture_y + offset_y
		t.Y = position_y
		for x := 0; x < max_x; x++ {
			if data[y][x] == 1 {
				position_x := int32(x) * texture_x + offset_x
				t.X = position_x
				tile := t.ToFloat32()
				if CheckCollisionRecs(tile, *rec) {
					if rec.X + rec.Width > tile.X + tile.Width && speed.X < 0 {
						rec.X = tile.X + tile.Width
					}
					if rec.Y + rec.Height > tile.Y + tile.Height && speed.Y < 0 {
						rec.Y = tile.Y + tile.Height
						*up = true
					}
                    if rec.X < tile.X && speed.X > 0 {
                        rec.X = tile.X - rec.Width
                    }
                    if rec.Y < tile.Y && speed.Y > 0 {
                        rec.Y = tile.Y - rec.Height
                    }
				}
				tile.Height += 0.25 * 16
				if CheckCollisionRecs(tile, *rec) {
                    if rec.Y < tile.Y && speed.Y > 0 {
                        *down = true
                    }
				}
			}
		}
	}
}

func MapDraw() {
	for y := 0; y < max_y; y++ {
		position_y := int32(y) * texture_y + offset_y
		for x := 0; x < max_x; x++ {
			if data[y][x] == 1 {
				position_x := int32(x) * texture_x + offset_x
				DrawTexture(texture, position_x, position_y, White)
			}
		}
	}
}
