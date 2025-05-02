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
