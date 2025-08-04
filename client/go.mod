module client

go 1.24.2

require (
	github.com/gen2brain/raylib-go/raylib v0.55.1
	github.com/or-n/util-go v0.1.3
)

require (
	github.com/ebitengine/purego v0.7.1 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
	golang.org/x/sys v0.20.0 // indirect
)

// replace github.com/gen2brain/raylib-go/raylib => ./Raylib-Go-Wasm/raylib

require (
	github.com/coder/websocket v1.8.13
	shared v0.0.0
)

replace shared => ../shared
