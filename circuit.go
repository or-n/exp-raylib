package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

var (
	Position   []Vector2
	Active     = make([]bool, 4, 4)
	Visited    = make([]bool, 4, 4)
	Signal     = make(map[int][]int)
	radius     = float32(20)
	thick      = float32(10)
	colorOn    = NewColor(255, 127, 127, 255)
	colorOff   = NewColor(255, 0, 0, 255)
	lastChange float64
)

func CircuitInit() {
	Position = append(Position, NewVector2(0.25, 0.75))
	Position = append(Position, NewVector2(0.75, 0.75))
	Position = append(Position, NewVector2(0.5, 0.25))
	Position = append(Position, NewVector2(0.5, 0.75))
	Signal[0] = append(Signal[0], 2)
	Signal[1] = append(Signal[1], 2)
	Signal[2] = append(Signal[2], 3)
	lastChange = GetTime()
}

func activate(from int) {
	if Visited[from] {
		return
	}
	Visited[from] = true
	if !Active[from] {
		return
	}
	for _, to := range Signal[from] {
		Active[to] = true
		activate(to)
	}
}

func CircuitUpdate() {
	cursor := GetMousePosition()
	for i, p := range Position {
		position := ScreenPosition(p)
		Active[i] = false
		Visited[i] = false
		if CheckCollisionPointCircle(cursor, position, radius) {
			Active[i] = IsMouseButtonDown(MouseButtonLeft)
		}
	}
	if GetTime() > lastChange+1 {
		for from := range len(Position) {
			activate(from)
		}
		lastChange = GetTime()
	}
}

func color(x bool) Color {
	if x {
		return colorOn
	}
	return colorOff
}

func CircuitDraw() {
	for from, targets := range Signal {
		for _, to := range targets {
			start := ScreenPosition(Position[from])
			end := ScreenPosition(Position[to])
			DrawLineEx(start, end, thick, color(Active[from]))
		}
	}
	for i, p := range Position {
		DrawCircleV(ScreenPosition(p), radius, color(Active[i]))
	}
}
