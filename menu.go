package main

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	. "github.com/gen2brain/raylib-go/raylib"
)

type State int

const (
	StateMenu State = iota
	StateGame
	StateOptions
	StateExit
)

var (
	SimulationState State
	button          = NewVector2(200, 100)
)

func MenuInit() {
	SimulationState = StateGame
	gui.SetStyle(gui.DEFAULT, gui.TEXT_SIZE, 30)
	gui.SetFont(MainFont)
}

func MenuDraw() {
	x := (WindowSize.X - button.X) * 0.5
	y := (WindowSize.Y - button.Y*4) * 0.5
	rect := NewRectangle(x, y, button.X, button.Y)
	if gui.Button(rect, Lang[Play]) {
		SimulationState = StateGame
	}
	rect.Y += button.Y
	if gui.Button(rect, Lang[Restart]) {
		SimulationState = StateGame
		MainAccount.Balance = 1000
		MainAccount.decided = false
	}
	rect.Y += button.Y
	if gui.Button(rect, Lang[Options]) {
		SimulationState = StateOptions
	}
	rect.Y += button.Y
	if gui.Button(rect, Lang[Exit]) {
		SimulationState = StateExit
	}
}
