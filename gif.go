package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
	"image/color"
	"unsafe"
)

var (
	frames     i32
	frame      i32
	gif        *Image
	GifTexture Texture2D
	frameSize  int
	width      int
	height     int
	timer      f32
	frameTimer f32
	Gif        bool
)

func GifInit() {
	gif = LoadImageAnim("asset/smell-money.gif", &frames)
	width = int(gif.Width)
	height = int(gif.Height)
	frameSize = int(gif.Width) * int(gif.Height) * 4
	GifTexture = LoadTextureFromImage(gif)
}

func GifUpdate() {
	if !Gif {
		timer += GetFrameTime()
		if timer >= 1 {
			timer = 0
			if GetRandomValue(0, 599) == 0 {
				Gif = true
			}
		}
		return
	}
	frameTimer += GetFrameTime()
	data := unsafe.Slice((*byte)(gif.Data), int(frames)*frameSize)
	offset := int(frame) * frameSize
	frameBytes := data[offset : offset+frameSize]
	colors := unsafe.Slice((*color.RGBA)(unsafe.Pointer(&frameBytes[0])), width*height)
	UpdateTexture(GifTexture, colors)
	if frameTimer >= 0.1 {
		frameTimer = 0
		frame++
		if frame >= frames {
			frame = 0
			Gif = false
		}
	}
}
