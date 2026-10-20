module github.com/oliverbestmann/box2d-go/example

go 1.26.0

require (
	github.com/hajimehoshi/ebiten/v2 v2.10.5
	github.com/jakecoffman/cp/v2 v2.4.0
	github.com/oliverbestmann/box2d-go/v3 v3.2.0
)

require (
	github.com/ebitengine/gomobile v0.0.0-20260820040257-d11f821a26a6 // indirect
	github.com/ebitengine/hideconsole v1.0.0 // indirect
	github.com/ebitengine/purego v0.11.2 // indirect
	golang.org/x/sync v0.24.0 // indirect
	golang.org/x/sys v0.49.0 // indirect
)

replace github.com/oliverbestmann/box2d-go/v3 => ../
