package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
	"math"
)

var (
	PlayerPosition Vector2
	PlayerSize Vector2
	PlayerTexture Texture2D
	PlayerGrounded bool
	PlayerJumpTo *float32
)

func PlayerInit() {
    PlayerPosition = NewVector2(0, -100)
    PlayerSize = NewVector2(16, 32)
    PlayerTexture = LoadTexture("asset/player.png")
}

func PlayerGetRect(position Vector2) Rectangle {
	rec := Rectangle{}
	rec.X = position.X + 1
	rec.Y = position.Y + 2
	rec.Width = PlayerSize.X - 2
	rec.Height = PlayerSize.Y - 2
	return rec
}

func PlayerUpdate() {
	dt := GetFrameTime()
	if PlayerJumpTo != nil && PlayerPosition.Y < *PlayerJumpTo {
		PlayerJumpTo = nil
	}
	if PlayerJumpTo != nil {
		positionUp := Vector2Add(PlayerPosition, NewVector2(0, -100 * dt))
		rect := PlayerGetRect(positionUp)
		if MapCollide(&rect) {
			PlayerJumpTo = nil
		} else {
			PlayerPosition = positionUp
		}
	}
	if PlayerJumpTo == nil {
		positionWithGravity := Vector2Add(PlayerPosition, NewVector2(0, 250 * dt))
		rect := PlayerGetRect(positionWithGravity)
		if MapCollide(&rect) {
			PlayerGrounded = true
			PlayerPosition.Y = float32(math.Round(float64(PlayerPosition.Y / 16)) * 16)
		} else {
			PlayerGrounded = false
			PlayerPosition = positionWithGravity
		}
	}
	if PlayerGrounded && IsKeyDown(InputJump) {
		value := PlayerPosition.Y - 1.25 * 16
		PlayerJumpTo = new(float32)
		*PlayerJumpTo = value
		PlayerGrounded = false
	}
	var speedX int32
	if IsKeyDown(InputSprint) {
		speedX = 400
	} else {
		speedX = 200
	}
	deltaX := float32(InputAxisX() * speedX)
	positionMove := Vector2Add(PlayerPosition, NewVector2(deltaX * dt, 0))
	rect := PlayerGetRect(positionMove)
	if !MapCollide(&rect) {
		PlayerPosition = positionMove
	}
}

func PlayerDraw() {
	color := White
	if PlayerJumpTo != nil {
		color = Red
	} else if PlayerGrounded {
		color = Green
	}
    DrawTextureV(PlayerTexture, PlayerPosition, color)
}
