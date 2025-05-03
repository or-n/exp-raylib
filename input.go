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

func KeyString(x int32) string {
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
	Input = map[Action]int32{
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
	keys = []int32{
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

func boolI32(x bool) int32 {
	var result int32
	if x {
		result = 1
	}
	return result
}

func InputAxisX() int32 {
	return boolI32(IsKeyDown(Input[ActionX])) - boolI32(IsKeyDown(Input[ActionNegX]))
}

func InputDraw() {
	textY := int32(15)
	textSpacing := float32(2)
	fontSize := int32(20)
	spacing := int32(20)
	buttonSize := NewVector2(300, 50)
	n := int32(7)
	size := Vector2{X: buttonSize.X, Y: buttonSize.Y * float32(n)}
	size.Y += float32(spacing * (n + 1))
	start := Vector2{X: (WindowSize.X - size.X) * 0.5, Y: (WindowSize.Y - size.Y) * 0.5}
	title := "Configure Keys (Press to Change)"
	titleSize := MeasureTextEx(GetFontDefault(), title, float32(fontSize), textSpacing)
	startTitle := (WindowSize.X - titleSize.X) * 0.5
	DrawText(title, int32(startTitle), int32(start.Y)+textY, fontSize, White)
	pad := int32(40)
	for action := range n {
		y := (action + 1) * (int32(buttonSize.Y) + spacing)
		p := NewVector2(start.X, start.Y+float32(y))
		DrawRectangleV(p, buttonSize, Blue)
		text := actionString(Action(action))
		DrawText(text, int32(start.X)+pad, int32(start.Y)+y+textY, fontSize, White)
		key := Input[Action(action)]
		var name string
		if changeAction != nil && *changeAction == Action(action) {
			name = "_"
		} else {
			name = KeyString(key)
		}
		nameSize := MeasureTextEx(GetFontDefault(), name, float32(fontSize), 1)
		position := NewVector2(start.X+buttonSize.X-nameSize.X-float32(pad), p.Y+float32(textY))
		DrawTextEx(GetFontDefault(), name, position, float32(fontSize), 1, White)
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
