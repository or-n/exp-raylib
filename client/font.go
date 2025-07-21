package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

var (
	en         = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	pl         = []rune("ąćęłńóśźżĄĆĘŁŃÓŚŹŻ")
	symbols    = []rune(" .,!?:;_()")
	codepoints = append(append(en, pl...), symbols...)
	MainFont   Font
	fontFile   = "asset/FiraCode-Bold.ttf"
)

func FontInit() {
	MainFont = GetFontDefault()
	// MainFont = LoadFontEx(fontFile, 32, codepoints, i32(len(codepoints)))
}
