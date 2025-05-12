package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

const (
	MapX       = 192
	MapY       = 108
	TileWidth  = 10
	TileHeight = 10
)

type Tile int

const (
	Empty Tile = iota
	Full
	Explore
	Visited
)

var (
	Map [MapY][MapX]Tile
)

func MapRestart() {
	Map = [MapY][MapX]Tile{}
}

func MapUpdate() {
	p := GetMousePosition()
	tile_x, tile_y := GetTile(p)
	if !Between(tile_x, 0, MapX-1) {
		return
	}
	if !Between(tile_y, 0, MapY-1) {
		return
	}
	delete := IsMouseButtonDown(MouseButtonRight)
	if delete {
		Map[tile_y][tile_x] = Empty
	}
	fill := IsMouseButtonDown(MouseButtonLeft)
	if fill {
		Map[tile_y][tile_x] = Full
	}
	if IsKeyDown(KeyLeftControl) && fill {
		Map[tile_y][tile_x] = Explore
	}
	for y := range MapY {
		for x := range MapX {
			if Map[y][x] == Explore {
				v := GetRandomValue(0, 3)
				dy, dx := dir(int(v))
				MapExplore(y, x, dy, dx)
				Map[y][x] = Visited
			}
		}
	}
}

func dir(n int) (int, int) {
	switch n {
	case 0:
		return -1, 0
	case 1:
		return 1, 0
	case 2:
		return 0, -1
	case 3:
		return 0, 1
	}
	return 0, 0
}

func MapExplore(y, x, dy, dx int) bool {
	if Map[y][x] != Explore {
		return false
	}
	new_y := y + dy
	if !Between(new_y, 0, MapY-1) {
		return false
	}
	new_x := x + dx
	if !Between(new_x, 0, MapX-1) {
		return false
	}
	if Map[new_y][new_x] != Empty {
		return false
	}
	Map[new_y][new_x] = Explore
	return true
}

func MapDraw() {
	for y := range MapY {
		for x := range MapX {
			var color Color
			if Map[y][x] == Full {
				color.A = 127
			}
			if Map[y][x] == Explore {
				color.G = 255
				color.A = 255
			}
			if Map[y][x] == Visited {
				color.R = 255
				color.A = 255
			}
			DrawRectangle(i32(x)*TileWidth, i32(y)*TileHeight, TileWidth, TileHeight, color)
		}
	}
}

func GetTile(position Vector2) (int, int) {
	x := position.X / TileWidth
	y := position.Y / TileHeight
	return int(x), int(y)
}

func Between(x, start, end int) bool {
	return x >= start && x <= end
}
