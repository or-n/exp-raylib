package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

var (
	MainFont Font
)

func FontInit() {
	polish := []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		' ', '.', ',', '!', '?', ':', ';', '_', '(', ')',
		260, 261, // Ą ą
		262, 263, // Ć ć
		280, 281, // Ę ę
		321, 322, // Ł ł
		323, 324, // Ń ń
		211, 243, // Ó ó
		346, 347, // Ś ś
		377, 378, // Ź ź
		379, 380, // Ż ż
	}
	MainFont = LoadFontEx("asset/FiraCode-Bold.ttf", 32, polish, i32(len(polish)))
}
