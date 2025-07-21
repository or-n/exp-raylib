module client

go 1.24.2

require (
	github.com/gen2brain/raylib-go/raylib v0.55.1
	github.com/or-n/util-go v0.1.3
)

require github.com/BrownNPC/wasm-ffi-go v1.1.0 // indirect

replace github.com/gen2brain/raylib-go/raylib => ./Raylib-Go-Wasm/raylib

require (
	github.com/coder/websocket v1.8.13
	shared v0.0.0
)

replace shared => ../shared
