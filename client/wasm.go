package main

import (
	"embed"
	// . "github.com/gen2brain/raylib-go/raylib"
)

//go:embed asset
var ASSETS embed.FS

func init() {
	// AddFileSystem(ASSETS)
}
