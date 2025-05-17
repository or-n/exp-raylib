package shared

import (
	"fmt"
	. "github.com/or-n/util-go"
	"math/rand"
)

const (
	MaxX = 10000
	MaxY = 256
)

type Block int

const (
	Empty Block = iota
	Dirt
)

func MapGen(Map *[MaxY][MaxX]Block) {
	var noise_scale = 0.01
	for y := range MaxY {
		for x := range MaxX {
			n := OctaveNoise(f64(x)*noise_scale, 0, 8, 0.5)
			var block Block
			if f64(MaxY-1-y) <= f64(MaxY/2)+n*64 {
				r := rand.Intn(8)
				if r > 0 {
					block = Dirt
				}
			}
			Map[y][x] = Block(block)
		}
	}
}

func MapLoad(filepath string, Map *[MaxY][MaxX]Block) {
	if err := Load(filepath, Map); err != nil {
		fmt.Println("Error loading map:", err)
		MapGen(Map)
		MapSave(filepath, Map)
	}
}

func MapSave(filepath string, Map *[MaxY][MaxX]Block) {
	if err := Save(filepath, *Map); err != nil {
		fmt.Println("Failed to save map:", err)
	}
}
