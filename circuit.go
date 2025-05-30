package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

var (
	Position   = make(map[int]Vector2)
	in         = make(map[int]map[int]struct{})
	out        = make(map[int]map[int]struct{})
	Active     = make([]bool, 0)
	Active2    = make([]bool, 0)
	Drag       = make(map[int]Vector2)
	radius     = float32(20)
	thick      = float32(10)
	colorOn    = NewColor(255, 127, 127, 255)
	colorOff   = NewColor(255, 0, 0, 255)
	lastChange float64
	choice     *int
	new_id     = 0
)

func create_node(position Vector2) {
	Position[new_id] = WorldPosition(Clip(ScreenPosition(position)))
	Active = append(Active, false)
	Active2 = append(Active2, false)
	new_id += 1
}

func create_edge(from, to int) {
	if out[from] == nil {
		out[from] = make(map[int]struct{})
	}
	out[from][to] = struct{}{}
	if in[to] == nil {
		in[to] = make(map[int]struct{})
	}
	in[to][from] = struct{}{}
}

func delete_edge(from, to int) {
	if tos, ok := out[from]; ok {
		delete(tos, to)
		if len(tos) == 0 {
			delete(out, from)
		}
	}
	if froms, ok := in[to]; ok {
		delete(froms, from)
		if len(froms) == 0 {
			delete(in, to)
		}
	}
}

func delete_node(i int) {
	if tos, ok := out[i]; ok {
		for to := range tos {
			delete_edge(i, to)
		}
	}
	if froms, ok := in[i]; ok {
		for from := range froms {
			delete_edge(from, i)
		}
	}
	delete(Position, i)
	delete(Drag, i)
	Active[i] = false
	Active2[i] = false
	if choice != nil && *choice == i {
		choice = nil
	}
}

func CircuitInit() {
	positions := []Vector2{
		NewVector2(0.25, 0.75),
		NewVector2(0.75, 0.75),
		NewVector2(0.5, 0.25),
		NewVector2(0.5, 0.75),
	}
	for _, p := range positions {
		create_node(p)
	}
	create_edge(0, 2)
	create_edge(1, 2)
	create_edge(2, 3)
	create_edge(3, 1)
}

func CircuitUpdate() {
	cursor := GetMousePosition()
	var hit_exists bool
	for i := range Position {
		if offset, exists := Drag[i]; exists {
			Position[i] = WorldPosition(Clip(Vector2Add(cursor, offset)))
		}
		position := ScreenPosition(Position[i])
		if CheckCollisionPointCircle(cursor, position, radius) {
			if IsMouseButtonPressed(MouseButtonRight) {
				delete_node(i)
				continue
			}
			if IsMouseButtonDown(MouseButtonLeft) {
				Active[i] = true
				Active2[i] = true
			}
			if IsMouseButtonPressed(MouseButtonLeft) {
				hit_exists = true
				Drag[i] = Vector2Subtract(position, cursor)
				if choice == nil {
					choice = new(int)
					*choice = i
				} else {
					if IsKeyDown(KeyLeftControl) {
						create_edge(*choice, i)
					}
					*choice = i
				}
			}
		}
	}
	if !hit_exists && IsMouseButtonPressed(MouseButtonLeft) {
		if IsKeyDown(KeyLeftControl) {
			create_node(WorldPosition(cursor))
		} else {
			choice = nil
		}
	}
	if IsMouseButtonReleased(MouseButtonLeft) {
		for k := range Drag {
			delete(Drag, k)
		}
	}
	if GetTime() > lastChange+0.1 {
		lastChange = GetTime()
		for from, active := range Active {
			if !active {
				continue
			}
			if tos, ok := out[from]; ok {
				for to := range tos {
					Active2[to] = true
				}
			}
		}
		Active, Active2 = Active2, Active
		for i := range Active2 {
			Active2[i] = false
		}
	}
}

func color(i int) Color {
	if Active[i] {
		return colorOn
	}
	return colorOff
}

func CircuitDraw() {
	for from, tos := range out {
		for to := range tos {
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
