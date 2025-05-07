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
)

func FontInit() {
	MainFont = LoadFontEx("asset/FiraCode-Bold.ttf", 32, codepoints, i32(len(codepoints)))
}
