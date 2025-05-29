package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

var (
	Position = []Vector2{
		NewVector2(0.25, 0.75),
		NewVector2(0.75, 0.75),
		NewVector2(0.5, 0.25),
		NewVector2(0.5, 0.75),
	}
	Signal = map[int][]int{
		0: {2},
		1: {2},
		2: {3},
		3: {1},
	}
	Active     = make([]bool, 4, 4)
	Drag       = make([]bool, 4, 4)
	Offset     = make([]Vector2, 4, 4)
	radius     = float32(20)
	thick      = float32(10)
	colorOn    = NewColor(255, 127, 127, 255)
	colorOff   = NewColor(255, 0, 0, 255)
	lastChange float64
	choice     *int
)

func CircuitUpdate() {
	cursor := GetMousePosition()
	Active2 := make([]bool, len(Position))
	for i := range Position {
		if Drag[i] {
			Position[i] = WorldPosition(Vector2Add(cursor, Offset[i]))
		}
		position := ScreenPosition(Position[i])
		if CheckCollisionPointCircle(cursor, position, radius) {
			if IsMouseButtonDown(MouseButtonLeft) {
				Active[i] = true
				Active2[i] = true
			}
			if IsMouseButtonPressed(MouseButtonLeft) {
				Drag[i] = true
				Offset[i] = Vector2Subtract(position, cursor)
				if choice == nil {
					choice = new(int)
					*choice = i
				} else {
					if IsKeyDown(KeyLeftControl) {
						Signal[*choice] = append(Signal[*choice], i)
					}
					*choice = i
				}
			}
		}
	}
	var hit *int
	for i := range Position {
		if Active2[i] {
			hit = new(int)
			*hit = i
		}
	}
	if hit == nil && IsMouseButtonPressed(MouseButtonLeft) {
		if IsKeyDown(KeyLeftControl) {
			i := len(Position)
			Position = append(Position, WorldPosition(cursor))
			Active = append(Active, false)
			Drag = append(Drag, false)
			Offset = append(Offset, NewVector2(0, 0))
			Signal[i] = []int{}
		} else {
			choice = nil
		}
	}
	if IsMouseButtonReleased(MouseButtonLeft) {
		Drag = make([]bool, len(Position))
	}
	if GetTime() > lastChange+0.1 {
		for from := range Position {
			if !Active[from] {
				continue
			}
			for _, to := range Signal[from] {
				Active2[to] = true
			}
		}
		copy(Active, Active2)
		lastChange = GetTime()
	}
}

func color(i int) Color {
	if Active[i] {
		return colorOn
	}
	return colorOff
}

func CircuitDraw() {
	for from, targets := range Signal {
		for _, to := range targets {
			start := ScreenPosition(Position[from])
			end := ScreenPosition(Position[to])
			DrawLineEx(start, end, thick, color(from))
		}
	}
	for i, p := range Position {
		c := color(i)
		if choice != nil && i == *choice {
			c = White
		}
		DrawCircleV(ScreenPosition(p), radius, c)
	}
}
