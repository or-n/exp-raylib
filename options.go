package main

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	. "github.com/gen2brain/raylib-go/raylib"
)

var (
	MusicVolume f32
)

func OptionsDraw() {
	rect := NewRectangle(WindowSize.X*0.5, 200, 400, 50)
	MusicVolume = gui.Slider(rect, Lang[Volume], "", MusicVolume, 0, 1)
}
