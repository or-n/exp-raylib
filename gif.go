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
	frameSize  i32
	timer      f32
	frameTimer f32
	Gif        bool
	GifX, GifY i32
)

func GifInit() {
	gif = LoadImageAnim("asset/smell-money.gif", &frames)
	frameSize = gif.Width * gif.Height * 4
	GifTexture = LoadTextureFromImage(gif)
}

func GifUpdate() {
	if !Gif {
		timer += GetFrameTime()
		if timer >= 1 {
			timer = 0
			if GetRandomValue(0, 599) == 0 {
				Gif = true
				GifX = (int32(WindowSize.X) - gif.Width) / 2
				GifY = (int32(WindowSize.Y) - gif.Height) / 2
			}
		}
		return
	}
	frameTimer += GetFrameTime()
	data := unsafe.Slice((*byte)(gif.Data), frames*frameSize)
	offset := frame * frameSize
	frameBytes := data[offset : offset+frameSize]
	colors := unsafe.Slice((*color.RGBA)(unsafe.Pointer(&frameBytes[0])), gif.Width*gif.Height)
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

func GifDraw() {
	if !Gif {
		return
	}
	DrawTexture(GifTexture, GifX, GifY, White)
}
