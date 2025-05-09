package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

var (
	texture Texture2D
)

func NoiseGenerate(width, height int, scale f64) {
	image := GenImagePerlinNoise(int(width), int(height), 0, 0, f32(scale))
	texture = LoadTextureFromImage(image)
}

func NoiseDraw() {
	DrawTexture(texture, 0, 0, White)
}
