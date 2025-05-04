package main

import . "github.com/gen2brain/raylib-go/raylib"

type Action int

const (
	ActionNegY Action = iota
	ActionNegX
	ActionY
	ActionX
	ActionJump
	ActionSneak
	ActionSprint
)

func actionString(a Action) string {
	switch a {
	case ActionNegY:
		return "neg_y"
	case ActionNegX:
		return "neg_x"
	case ActionY:
		return "y"
	case ActionX:
		return "x"
	case ActionJump:
		return "jump"
	case ActionSneak:
		return "sneak"
	case ActionSprint:
		return "sprint"
	}
	return ""
}

func KeyString(x i32) string {
	switch x {
	case KeyW:
		return "W"
	case KeyA:
		return "A"
	case KeyS:
		return "S"
	case KeyD:
		return "D"
	case KeyLeftShift:
		return "Left Shift"
	case KeyLeftControl:
		return "Left Control"
	case KeySpace:
		return "Space"
	}
	return ""
}

var (
	Input = map[Action]i32{
		ActionNegY:   KeyW,
		ActionNegX:   KeyA,
		ActionY:      KeyS,
		ActionX:      KeyD,
		ActionJump:   KeyW,
		ActionSneak:  KeyLeftShift,
		ActionSprint: KeyLeftControl,
	}
	actions = []Action{
		ActionNegY,
		ActionNegX,
		ActionY,
		ActionX,
		ActionJump,
		ActionSneak,
		ActionSprint,
	}
	keys = []i32{
		KeyW,
		KeyA,
		KeyS,
		KeyD,
		KeyLeftShift,
		KeyLeftControl,
		KeySpace,
	}
	changeAction *Action
)

func InputUpdate() {
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

func boolI32(x bool) i32 {
	var result i32
	if x {
		result = 1
	}
	return result
}

func InputAxisX() i32 {
	return boolI32(IsKeyDown(Input[ActionX])) - boolI32(IsKeyDown(Input[ActionNegX]))
}

func InputDraw() {
	textY := i32(15)
	textSpacing := f32(2)
	fontSize := i32(20)
	spacing := i32(20)
	buttonSize := NewVector2(300, 50)
	n := i32(len(actions))
	size := NewVector2(buttonSize.X, buttonSize.Y*f32(n))
	size.Y += f32(spacing * (n + 1))
	start := NewVector2((WindowSize.X-size.X)*0.5, (WindowSize.Y-size.Y)*0.5)
	title := "Configure Keys (Press to Change)"
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
