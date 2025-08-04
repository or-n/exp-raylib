package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
	. "github.com/or-n/util-go"
	. "shared"
	"strconv"
)

var (
	PlayerSize    Vector2
	PlayerTexture Texture2D
	MainPlayer    Player
	PlayerName    string
	PlayerFile    = "data/player.gob"
)

func PlayerInit() {
	PlayerName = "rep"
	PlayerLoad(PlayerFile, &MainPlayer)
	PlayerSize = NewVector2(8, 12)
	PlayerTexture = LoadTexture("asset/nwm.png")
}

func PlayerRestart() {
	MainPlayer.Position = NewVector2(0, f32(100*TextureY))
	MainPlayer.Grounded = false
	MainPlayer.JumpTo = nil
}

func PlayerGetRect(position Vector2) Rectangle {
	rec := Rectangle{}
	rec.X = position.X + 1
	rec.Y = position.Y + 2
	rec.Width = PlayerSize.X - 2
	rec.Height = PlayerSize.Y - 2
	return rec
}

func PlayerRealPosition(player *Player) Vector2 {
	return Vector2Add(player.Position, NewVector2(1, 2))
}

func PlayerRealSize() Vector2 {
	return Vector2Subtract(PlayerSize, NewVector2(2, 2))
}

func PlayerCenter(player *Player) Vector2 {
	return Vector2Add(PlayerRealPosition(player), Vector2Scale(PlayerRealSize(), 0.5))
}

func PlayerPositionUpdate(player *Player) {
	dt := GetFrameTime()
	if player.JumpTo != nil && player.Position.Y < *player.JumpTo {
		player.JumpTo = nil
	}
	if player.JumpTo != nil {
		positionUp := Vector2Add(player.Position, NewVector2(0, -100*dt))
		rect := PlayerGetRect(positionUp)
		if MapCollide(&rect) {
			player.JumpTo = nil
			player.Position.Y = RoundF32(player.Position.Y)
		} else {
			player.Position = positionUp
		}
	}
	if player.JumpTo == nil {
		positionWithGravity := Vector2Add(player.Position, NewVector2(0, 200*dt))
		rect := PlayerGetRect(positionWithGravity)
		if MapCollide(&rect) {
			player.Grounded = true
			player.Position.Y = RoundF32(player.Position.Y)
		} else {
			player.Grounded = false
			player.Position = positionWithGravity
		}
	}
	if player.Grounded && IsKeyDown(Input[ActionJump]) {
		value := player.Position.Y - 1.25*16
		player.JumpTo = new(f32)
		*player.JumpTo = value
		player.Grounded = false
	}
	var speedX i32
	if IsKeyDown(Input[ActionSneak]) {
		speedX = 25
	} else if IsKeyDown(Input[ActionSprint]) {
		speedX = 400
	} else {
		speedX = 200
	}
	deltaX := f32(InputAxisX() * speedX)
	positionMove := Vector2Add(player.Position, NewVector2(deltaX*dt, 0))
	rect := PlayerGetRect(positionMove)
	if !MapCollide(&rect) {
		player.Position = positionMove
	} else {
		player.Position.X = RoundF32(player.Position.X)
	}
}

func PlayerUpdate(player *Player) {
	if !MapLoaded {
		return
	}
	PlayerPositionUpdate(player)
	p := CursorPosition()
	x, y := MapIndex(p)
	if MapInsideX(x) && MapInsideY(y) {
		r := MapRect(x, y)
		if IsMouseButtonDown(MouseButtonLeft) && Map[y][x] != Empty {
			Outgoing <- Message{Type: ClientChangeBlock, Data: ChangeBlockData{
				X: x, Y: y, Block: Empty,
			}}
			player.Inventory += 1
		}
		if IsMouseButtonDown(MouseButtonRight) && Map[y][x] == Empty && player.Inventory > 0 {
			p := PlayerGetRect(player.Position)
			if CheckCollisionRecs(p, r) {
				return
			}
			Outgoing <- Message{Type: ClientChangeBlock, Data: ChangeBlockData{
				X: x, Y: y, Block: Dirt,
			}}
			player.Inventory -= 1
		}
	}
}

func PlayerOverlayDraw(player *Player) {
	inventory := strconv.Itoa(player.Inventory)
	DrawText(inventory, 30, 100, 20, White)
	p := PlayerCenter(player)
	x := strconv.Itoa(int(p.X / f32(TextureX)))
	y := strconv.Itoa(int(p.Y / f32(TextureY)))
	DrawText(x, 200, 30, 20, White)
	DrawText(y, 250, 30, 20, White)
}

func PlayerDraw(player *Player) {
	rect := PlayerGetRect(NewVector2(0, 0))
	DrawTextureRec(PlayerTexture, rect, PlayerRealPosition(player), White)
}
