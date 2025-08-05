package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
	. "github.com/or-n/util-go"
	"log"
	// "math"
	. "shared"
	"strconv"
)

var (
	PlayerTexture Texture2D
	PlayerName    string
	Players       = make(map[string]Player)
)

func PlayerInit() {
	PlayerName = "rep"
	Players[PlayerName] = Player{}
	PlayerTexture = LoadTexture("asset/nwm.png")
}

func PlayerRealPosition(player *Player) Vector2 {
	return player.Position
	// return Vector2Add(player.Position, NewVector2(1, 2))
}

func PlayerRealSize() Vector2 {
	return PlayerSize
	// return Vector2Subtract(PlayerSize, NewVector2(2, 2))
}

func PlayerCenter(player *Player) Vector2 {
	return Vector2Add(PlayerRealPosition(player), Vector2Scale(PlayerRealSize(), 0.5))
}

func PlayerFeet(player *Player) Vector2 {
	p := PlayerRealPosition(player)
	p.X += PlayerSize.X / 2
	p.Y += PlayerSize.Y * 0.9
	return p
}

func PlayerPositionUpdate(player *Player) {
	dt := GetFrameTime()
	if player.JumpTo != nil {
		if player.Position.Y < *player.JumpTo {
			player.JumpTo = nil
			player.PlaceBelow = nil
		} else {
			if player.PlaceBelow != nil {
				rect := PlayerGetRect(player.Position)
				x := player.PlaceBelow.X
				y := player.PlaceBelow.Y
				r := MapRect(x, y)
				if r.Y < player.Position.Y {

				} else if !CheckCollisionRecs(rect, r) {
					change := ChangeBlockData{X: x, Y: y, Block: Dirt}
					data := PlayerChangeBlock{Name: PlayerName, Change: change}
					Outgoing <- Message{Type: ClientChangeBlock, Data: data}
				}
			}
		}
		change := NewVector2(0, -50*dt)
		positionUp := Vector2Add(player.Position, change)
		rect := PlayerGetRect(positionUp)
		if MapCollide(&rect) {
			player.JumpTo = nil
			player.PlaceBelow = nil
		} else {
			player.Position = positionUp
		}
	}
	if player.JumpTo == nil {
		positionWithGravity := Vector2Add(player.Position, NewVector2(0, 100*dt))
		rect := PlayerGetRect(positionWithGravity)
		if MapCollide(&rect) {
			player.Grounded = true
			bottomEdge := player.Position.Y + PlayerSize.Y
			snapped := RoundF32(bottomEdge/16) * 16
			player.Position.Y = snapped - PlayerSize.Y
		} else {
			player.Grounded = false
			player.Position = positionWithGravity
		}
	}
	if player.Grounded && IsKeyDown(Input[ActionJump]) {
		value := player.Position.Y - 1.25*f32(TextureY)
		player.JumpTo = new(f32)
		*player.JumpTo = value
		player.Grounded = false
	}
	var speedX i32
	if IsKeyDown(Input[ActionSneak]) {
		speedX = 25
	} else if IsKeyDown(Input[ActionSprint]) {
		speedX = 100
	} else {
		speedX = 50
	}
	deltaX := f32(InputAxisX() * speedX)
	positionMove := Vector2Add(player.Position, NewVector2(deltaX*dt, 0))
	rect := PlayerGetRect(positionMove)
	if !MapCollide(&rect) {
		player.Position = positionMove
	} else {
		if deltaX > 0 {
			rightEdge := player.Position.X + PlayerSize.X
			snapped := RoundF32(rightEdge/16) * 16
			player.Position.X = snapped - PlayerSize.X
		} else if deltaX < 0 {
			leftEdge := player.Position.X
			snapped := RoundF32(leftEdge/16) * 16
			player.Position.X = snapped
		}
	}
}

func PlayerUpdate(player *Player) {
	if !MapLoaded {
		return
	}
	old := *player
	PlayerPositionUpdate(player)
	if old != *player {
		Outgoing <- Message{Type: ClientChangePlayer, Data: PlayerData{Name: PlayerName, Player: *player}}
	}
	p := CursorPosition()
	x, y := MapIndex(p)
	if MapInsideX(x) && MapInsideY(y) {
		if IsMouseButtonDown(MouseButtonLeft) && Map[y][x] != Empty {
			change := ChangeBlockData{X: x, Y: y, Block: Empty}
			data := PlayerChangeBlock{Name: PlayerName, Change: change}
			Outgoing <- Message{Type: ClientChangeBlock, Data: data}
		}
		if IsMouseButtonDown(MouseButtonRight) && Map[y][x] == Empty && player.Inventory > 0 {
			change := ChangeBlockData{X: x, Y: y, Block: Dirt}
			data := PlayerChangeBlock{Name: PlayerName, Change: change}
			Outgoing <- Message{Type: ClientChangeBlock, Data: data}
			log.Println("RPM", x, y)
		}
	}
}

func PlayerOverlayDraw(player *Player) {
	inventory := strconv.Itoa(player.Inventory)
	DrawText(inventory, 30, 100, 20, White)
	p := PlayerCenter(player)
	x, y := MapIndex(p)
	X := strconv.Itoa(x)
	Y := strconv.Itoa(y)
	DrawText(X, 200, 30, 20, White)
	DrawText(Y, 250, 30, 20, White)
}

func PlayerDraw(player *Player) {
	rect := PlayerGetRect(NewVector2(0, 0))
	// r := PlayerGetRect(player.Position)
	// DrawRectangleLinesEx(r, 1, Green)
	// x, y := MapIndex(PlayerFeet(player))
	// r2 := MapRect(x, y)
	// DrawRectangleLinesEx(r2, 1, Red)
	DrawTextureRec(PlayerTexture, rect, PlayerRealPosition(player), White)
	cursor := CursorPosition()
	cx, cy := MapIndex(cursor)
	cr := MapRect(cx, cy)
	DrawRectangleLinesEx(cr, 1, Blue)
}
