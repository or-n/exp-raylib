package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

var (
	en         = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	pl         = []rune("ąćęłńóśźżĄĆĘŁŃÓŚŹŻ")
	symbols    = []rune(" .,!?:;_()/")
	codepoints = append(append(en, pl...), symbols...)
	MainFont   Font
	fontFile   = "asset/FiraCode-Bold.ttf"
)

func main() {
	screenWidth := int32(1920) / 2
	screenHeight := int32(1080) / 2

	InitWindow(screenWidth, screenHeight, "raylib-go")
	defer CloseWindow()

	SetTargetFPS(60)
	ToggleFullscreen()

	MainFont = LoadFontEx(fontFile, 20, codepoints, int32(len(codepoints)))

	margin := float32(20) / float32(1080)
	w := float32(screenWidth)
	h := float32(screenHeight)
	m := h * margin

	var (
		panelLocalX       = m
		panelLocalY       = m
		panelLocalWidth   = w/3 - m - m/2
		panelLocalHeight  = h/2 - m - m/2
		panelServerX      = w/3 + m/2
		panelServerY      = m
		panelServerWidth  = w/3 - m
		panelServerHeight = h/2 - m - m/2
		panelCmdX         = w*2/3 + m/2
		panelCmdY         = m
		panelCmdWidth     = w/3 - m/2 - m
		panelCmdHeight    = h/4 - m - m/2
		panelInfoX        = w*2/3 + m/2
		panelInfoY        = h/4 + m/2
		panelInfoWidth    = w/3 - m/2 - m
		panelInfoHeight   = h/4 - m
		panelLogX         = m
		panelLogY         = h/2 + m/2
		panelLogWidth     = w - m*2
		panelLogHeight    = h/2 - m/2 - m
	)

	for !WindowShouldClose() {
		BeginDrawing()
		ClearBackground(RayWhite)

		DrawRectangle(
			int32(panelLocalX), int32(panelLocalY),
			int32(panelLocalWidth), int32(panelLocalHeight),
			Color{R: 0, G: 120, B: 255, A: 255},
		)
		DrawTextEx(MainFont, "pliki lokalne",
			Vector2{X: panelLocalX + m, Y: panelLocalY + m}, 20, 2, White)

		DrawRectangle(
			int32(panelServerX), int32(panelServerY),
			int32(panelServerWidth), int32(panelServerHeight),
			Color{R: 255, G: 0, B: 0, A: 255},
		)
		DrawTextEx(MainFont, "pliki serwerowe",
			Vector2{X: panelServerX + m, Y: panelServerY + m}, 20, 2, White)

		DrawRectangle(
			int32(panelCmdX), int32(panelCmdY),
			int32(panelCmdWidth), int32(panelCmdHeight),
			Black,
		)
		DrawTextEx(MainFont, "lista podstawowych komend",
			Vector2{X: panelCmdX + m, Y: panelCmdY + m}, 20, 2, White)

		DrawRectangle(
			int32(panelInfoX), int32(panelInfoY),
			int32(panelInfoWidth), int32(panelInfoHeight),
			Color{R: 128, G: 128, B: 128, A: 255},
		)

		DrawTextEx(MainFont, "podstawowe informacje:",
			Vector2{X: panelInfoX + m, Y: panelInfoY + m}, 20, 2, White)
		DrawTextEx(MainFont, "stan serwera: wł/wył",
			Vector2{X: panelInfoX + m, Y: panelInfoY + m + 20}, 20, 2, White)
		DrawTextEx(MainFont, "ilość osób",
			Vector2{X: panelInfoX + m, Y: panelInfoY + m + 20*2}, 20, 2, White)
		DrawTextEx(MainFont, "itp.",
			Vector2{X: panelInfoX + m, Y: panelInfoY + m + 20*3}, 20, 2, White)

		DrawRectangle(
			int32(panelLogX), int32(panelLogY),
			int32(panelLogWidth), int32(panelLogHeight),
			Color{R: 255, G: 215, B: 0, A: 255},
		)
		DrawTextEx(MainFont, "logi",
			Vector2{X: panelLogX + m, Y: panelLogY + m}, 20, 2, Black)

		EndDrawing()
	}

}
