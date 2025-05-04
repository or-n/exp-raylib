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

const (
	start   = "start"
	restart = "restart"
	options = "options"
	exit    = "exit"
)

var (
	SimulationState State
	button          Vector2
)

func MenuInit() {
	SimulationState = StateMenu
	button = NewVector2(200, 100)
}

func MenuDraw() {
	gui.SetStyle(gui.DEFAULT, gui.TEXT_SIZE, 30)
	x := (WindowSize.X - button.X) * 0.5
	y := (WindowSize.Y - button.Y*4) * 0.5
	rect := NewRectangle(x, y, button.X, button.Y)
	if gui.Button(rect, "Start") {
		SimulationState = StateGame
	}
	rect.Y += button.Y
	if gui.Button(rect, "Restart") {
		PlayerRestart()
		CameraRestart()
		SimulationState = StateGame
	}
	rect.Y += button.Y
	if gui.Button(rect, "Options") {
		SimulationState = StateOptions
	}
	rect.Y += button.Y
	if gui.Button(rect, "Exit") {
		SimulationState = StateExit
	}
}
