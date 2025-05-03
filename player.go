package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
	"math"
)

var (
	PlayerPosition Vector2
	PlayerSize Vector2
	PlayerTexture Texture2D
	grounded bool
	jumpTo *float32
)

func PlayerInit() {
    PlayerPosition = NewVector2(0, -100)
    PlayerSize = NewVector2(16, 32)
    PlayerTexture = LoadTexture("asset/nwm.png")
}

func PlayerGetRect(position Vector2) Rectangle {
	rec := Rectangle{}
	rec.X = position.X + 1
	rec.Y = position.Y + 2
	rec.Width = PlayerSize.X - 2
	rec.Height = PlayerSize.Y - 2
	return rec
}

func PlayerRealPosition() Vector2 {
	return Vector2Add(PlayerPosition, NewVector2(1, 2))
}

func PlayerRealSize() Vector2 {
	return Vector2Subtract(PlayerSize, NewVector2(2, 2))
}

func PlayerCenter() Vector2 {
	return Vector2Add(PlayerRealPosition(), Vector2Scale(PlayerRealSize(), 0.5))
}

func Round(x float32) float32 {
	return float32(math.Round(float64(x)))
}

func PlayerUpdate() {
	dt := GetFrameTime()
	if jumpTo != nil && PlayerPosition.Y < *jumpTo {
		jumpTo = nil
	}
	if jumpTo != nil {
		positionUp := Vector2Add(PlayerPosition, NewVector2(0, -100 * dt))
		rect := PlayerGetRect(positionUp)
		if MapCollide(&rect) {
			jumpTo = nil
			PlayerPosition.Y = Round(PlayerPosition.Y)
		} else {
			PlayerPosition = positionUp
		}
	}
	if jumpTo == nil {
		positionWithGravity := Vector2Add(PlayerPosition, NewVector2(0, 200 * dt))
		rect := PlayerGetRect(positionWithGravity)
		if MapCollide(&rect) {
			grounded = true
			PlayerPosition.Y = Round(PlayerPosition.Y)
		} else {
			grounded = false
			PlayerPosition = positionWithGravity
		}
	}
	if grounded && IsKeyPressed(InputJump) {
		value := PlayerPosition.Y - 1.25 * 16
		jumpTo = new(float32)
		*jumpTo = value
		grounded = false
	}
	var speedX int32
	if IsKeyDown(InputSneak) {
		speedX = 25
	} else {
		speedX = 200
	}
	deltaX := float32(InputAxisX() * speedX)
	positionMove := Vector2Add(PlayerPosition, NewVector2(deltaX * dt, 0))
	rect := PlayerGetRect(positionMove)
	if !MapCollide(&rect) {
		PlayerPosition = positionMove
	} else {
		PlayerPosition.X = Round(PlayerPosition.X)
	}
}

func PlayerDraw() {
	rect := PlayerGetRect(NewVector2(0, 0))
	DrawTextureRec(PlayerTexture, rect, PlayerRealPosition(), White)
}
