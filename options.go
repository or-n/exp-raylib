package main

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	. "github.com/gen2brain/raylib-go/raylib"
)

var (
	changeAction *Action
	MusicVolume  f32
)

func OptionsUpdate() {
	if changeAction == nil {
		return
	}
	for _, key := range keys {
		if IsKeyPressed(key) {
			Input[*changeAction] = key
			changeAction = nil
		}
	}
}

func OptionsDraw() {
	InputOptionsDraw()
	MusicOptionsDraw()
}

func InputOptionsDraw() {
	textY := i32(15)
	textSpacing := f32(2)
	fontSize := i32(20)
	spacing := i32(20)
	buttonSize := NewVector2(300, 50)
	n := i32(7)
	size := NewVector2(buttonSize.X, buttonSize.Y*f32(n))
	size.Y += f32(spacing * (n + 1))
	start := NewVector2((WindowSize.X-size.X)*0.5, (WindowSize.Y-size.Y)*0.5)
	title := Lang[ConfigureKeys]
	titleSize := MeasureTextEx(GetFontDefault(), title, f32(fontSize), textSpacing)
	startTitle := (WindowSize.X - titleSize.X) * 0.5
	DrawText(title, i32(startTitle), i32(start.Y)+textY, fontSize, White)
	pad := i32(40)
	for action := range n {
		y := (action + 1) * (i32(buttonSize.Y) + spacing)
		p := NewVector2(start.X, start.Y+f32(y))
		DrawRectangleV(p, buttonSize, Blue)
		text := actionString(Action(action))
		DrawText(text, i32(p.X)+pad, i32(p.Y)+textY, fontSize, White)
		key := Input[Action(action)]
		var name string
		if changeAction != nil && *changeAction == Action(action) {
			name = "_"
		} else {
			name = KeyString(key)
		}
		nameSize := MeasureTextEx(GetFontDefault(), name, f32(fontSize), 1)
		position := NewVector2(p.X+buttonSize.X-nameSize.X-f32(pad), p.Y+f32(textY))
		DrawTextEx(GetFontDefault(), name, position, f32(fontSize), 1, White)
		rect := Rectangle{}
		rect.X = p.X
		rect.Y = p.Y
		rect.Width = buttonSize.X
		rect.Height = buttonSize.Y
		if CheckCollisionPointRec(GetMousePosition(), rect) {
			if IsMouseButtonPressed(MouseButtonLeft) {
				changeAction = new(Action)
				*changeAction = Action(action)
			}
		}
	}
}

func MusicOptionsDraw() {
	rect := NewRectangle(WindowSize.X*0.5, 200, 400, 50)
	MusicVolume = gui.Slider(rect, Lang[Volume], "", MusicVolume, 0, 1)
}
