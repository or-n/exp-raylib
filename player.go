package main

import . "github.com/gen2brain/raylib-go/raylib"

var (
	PlayerPosition Vector2
	PlayerSize Vector2
	PlayerSpeed Vector2
	PlayerTexture Texture2D
	PlayerGrounded bool
)

func PlayerInit() {
    PlayerPosition = NewVector2(0, -100)
    PlayerSize = NewVector2(16, 16)
    PlayerTexture = LoadTexture("asset/dirt.png")
}

func PlayerGetRect() Rectangle {
	rec := Rectangle{}
	rec.X = PlayerPosition.X
	rec.Y = PlayerPosition.Y
	rec.Width = PlayerSize.X
	rec.Height = PlayerSize.Y
	return rec
}

func PlayerSetRect(rec Rectangle) {
	PlayerPosition.X = rec.X
	PlayerPosition.Y = rec.Y
}

func PlayerUpdate() {
	var speed_x int32
	if IsKeyDown(InputSprint) {
		speed_x = 400
	} else {
		speed_x = 200
	}
	PlayerSpeed.X = float32(InputAxisX() * speed_x)
	if PlayerGrounded {
		PlayerSpeed.Y = 0
	} else {
		PlayerSpeed.Y = 250
	}
	delta := Vector2Scale(PlayerSpeed, GetFrameTime())
	PlayerPosition = Vector2Add(PlayerPosition, delta)
}

func PlayerDraw() {
	color := White
	if PlayerGrounded {
		color = Green
	}
    DrawTextureV(PlayerTexture, PlayerPosition, color)
    color.A = 127
    DrawRectangleRec(PlayerGetRect(), color)
    start := Vector2Add(PlayerPosition, Vector2Scale(PlayerSize, 0.5))
    end := Vector2Add(start, PlayerSpeed)
    DrawLineV(start, end, White)
}
