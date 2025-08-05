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
	chunks      [][]RenderTexture2D
	chunkX      = i32(16)
	chunkY      = i32(16)
	chunkW      = MaxX / chunkX
	chunkH      = MaxY / chunkY
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
	chunk_x := i32(data.X) / chunkX
	chunk_y := i32(data.Y) / chunkY
	BeginTextureMode(chunks[chunk_y][chunk_x])
	x := (i32(data.X) % chunkX) * TextureX
	y := (i32(data.Y) % chunkY) * TextureY
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
		// rect := CameraRect(0.25)
		rect := CameraRect(0)
		blockWidth := chunkX * TextureX
		blockHeight := chunkY * TextureY
		viewMinX := i32((rect.X - f32(OffsetX)) / f32(blockWidth))
		viewMaxX := i32((rect.X + rect.Width - f32(OffsetX)) / f32(blockWidth))
		viewMinY := i32((rect.Y - f32(OffsetY)) / f32(blockHeight))
		viewMaxY := i32((rect.Y + rect.Height - f32(OffsetY)) / f32(blockHeight))
		for chunk_y := max(0, viewMinY); chunk_y <= min(chunkH-1, viewMaxY); chunk_y++ {
			for chunk_x := max(0, viewMinX); chunk_x <= min(chunkW-1, viewMaxX); chunk_x++ {
				start_x := chunk_x * chunkX
				start_y := chunk_y * chunkY
				texture := chunks[chunk_y][chunk_x].Texture
				src := Rectangle{}
				src.Width = f32(texture.Width)
				src.Height = -f32(texture.Height)
				dst := Rectangle{}
				dst.X = f32(OffsetX + start_x*TextureX)
				dst.Y = f32(OffsetY + start_y*TextureY)
				dst.Width = f32(texture.Width)
				dst.Height = f32(texture.Height)
				origin := NewVector2(0, 0)
				DrawTexturePro(texture, src, dst, origin, 0, White)
			}
		}
		// DrawRectangleLinesEx(rect, 10.0/MainCamera.Zoom, Violet)
	}
}

func MapTextureInit() {
	chunks = make([][]RenderTexture2D, chunkH)
	for chunk_y := range chunkH {
		chunks[chunk_y] = make([]RenderTexture2D, chunkW)
		for chunk_x := range chunkW {
			chunks[chunk_y][chunk_x] = LoadRenderTexture(chunkX*TextureX, chunkY*TextureY)
			start_x := chunk_x * chunkX
			start_y := chunk_y * chunkY
			BeginTextureMode(chunks[chunk_y][chunk_x])
			ClearBackground(WindowBg)
			for y := range chunkY {
				pos_y := i32(y) * TextureY
				for x := range chunkX {
					pos_x := i32(x) * TextureX
					if Map[y+start_y][x+start_x] == Dirt {
						DrawTexture(dirtTexture, pos_x, pos_y, White)
					}
				}
			}
			EndTextureMode()
		}
	}
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
