package main

import . "github.com/gen2brain/raylib-go/raylib"

var (
	PlayerPosition Vector2
	PlayerSize Vector2
	PlayerSpeed Vector2
	PlayerTexture Texture2D
)

func PlayerInit() {
    PlayerPosition = NewVector2(0, -100)
    PlayerSize = NewVector2(16, 16)
    PlayerTexture = LoadTexture("asset/dirt.png")
}

func PlayerUpdate() {
	var speed_x int32
	if IsKeyDown(InputSprint) {
		speed_x = 400
	} else {
		speed_x = 200
	}
	PlayerSpeed.X = float32(InputAxisX() * speed_x)
	PlayerSpeed = Vector2Scale(PlayerSpeed, GetFrameTime())
	PlayerPosition = Vector2Add(PlayerPosition, PlayerSpeed)
}

func PlayerDraw() {
    DrawTextureV(PlayerTexture, PlayerPosition, White)
}
