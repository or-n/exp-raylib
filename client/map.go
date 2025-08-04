package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
	. "shared"
)

var (
	MapLoaded   bool
	MapRendered bool
	Map         [MaxY][MaxX]Block
	dirtTexture Texture2D
	mapTexture  RenderTexture2D
)

func MapInit() {
	dirtTexture = LoadTexture("asset/dirt.png")
}

func MapCollide(rec *Rectangle) bool {
	t := RectangleInt32{}
	t.Width = TextureX
	t.Height = TextureY
	center := RectCenter(*rec)
	cx, cy := MapIndex(center)
	n := 9
	for y := range MaxY {
		iy := y - n/2 + cy
		if !MapInsideY(iy) {
			continue
		}
		position_y := i32(iy)*TextureY + OffsetY
		t.Y = position_y
		for x := range n {
			ix := x - n/2 + cx
			if !MapInsideX(ix) {
				continue
			}
			position_x := i32(ix)*TextureX + OffsetX
			t.X = position_x
			if Map[iy][ix] == Dirt {
				tile := t.ToFloat32()
				if CheckCollisionRecs(tile, *rec) {
					return true
				}
			}
		}
	}
	return false
}

func RectCenter(r Rectangle) Vector2 {
	return NewVector2(r.X+r.Width*0.5, r.Y+r.Height*0.5)
}

func changeBlock(data ChangeBlockData) {
	Map[data.Y][data.X] = data.Block
	BeginTextureMode(mapTexture)
	x := i32(data.X) * TextureX
	y := i32(data.Y) * TextureY
	if data.Block == Dirt {
		DrawTexture(dirtTexture, x, y, White)
	} else {
		DrawRectangle(x, y, TextureX, TextureY, WindowBg)
	}
	EndTextureMode()
}

func MapDraw() {
	if MapLoaded {
		if !MapRendered {
			MapTextureInit()
			MapRendered = true
		}
		texture := mapTexture.Texture
		src := Rectangle{}
		src.Width = f32(texture.Width)
		src.Height = -f32(texture.Height)
		dst := Rectangle{}
		dst.X = f32(OffsetX)
		dst.Y = f32(OffsetY)
		dst.Width = f32(texture.Width)
		dst.Height = f32(texture.Height)
		origin := NewVector2(0, 0)
		DrawTexturePro(texture, src, dst, origin, 0, White)
	}
}

func MapTextureInit() {
	mapTexture = LoadRenderTexture(MaxX*TextureX, MaxY*TextureY)
	BeginTextureMode(mapTexture)
	for y := range MaxY {
		pos_y := i32(y) * TextureY
		for x := range MaxX {
			pos_x := i32(x) * TextureX
			if Map[y][x] == Dirt {
				DrawTexture(dirtTexture, pos_x, pos_y, White)
			}
		}
	}
	EndTextureMode()
}

func MapIndex(position Vector2) (int, int) {
	x := (i32(position.X) - OffsetX) / TextureX
	y := (i32(position.Y) - OffsetY) / TextureY
	return int(x), int(y)
}

func MapInsideX(x int) bool {
	return x >= 0 && x < MaxX
}

func MapInsideY(y int) bool {
	return y >= 0 && y < MaxY
}
