package main

import (
	"fmt"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const Layers = 20

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(1200, 800)
	ops := ebiten.RunGameOptions{}
	ops.SingleThread = true

	err := ebiten.RunGameWithOptions(&viz{}, &ops)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

type viz struct {
	physics     Physics
	bodies      []Body
	physicsTime float64
	paused      bool
	slowmo      bool
	timeAcc     float64
}

func (v *viz) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyB) {
		v.initialize(b2New)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		v.initialize(cpNew)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		v.paused = !v.paused
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		v.slowmo = !v.slowmo
	}

	if v.physics != nil && !v.paused {
		step := 1 / 60.0

		if v.slowmo {
			step *= 0.1
		}

		v.timeAcc += step

		for v.timeAcc > 1/60.0 {
			v.timeAcc -= 1 / 60.0

			startTime := time.Now()
			v.physics.Step(1/60.0, 4)

			if v.physicsTime == 0 {
				v.physicsTime = time.Since(startTime).Seconds()
			} else {
				v.physicsTime = 0.95*v.physicsTime + 0.05*time.Since(startTime).Seconds()
			}
		}
	}

	return nil
}

func (v *viz) Draw(screen *ebiten.Image) {
	if v.physics == nil {
		text := "press b for box2d, c for jakecoffman/cp"
		ebitenutil.DebugPrintAt(screen, text, 16, 16)
		return
	}

	w, h := float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy())

	var toScreen ebiten.GeoM

	// scale := (h - 50) / (Layers + 5)
	scale := w / (5 * Layers)

	toScreen.Scale(scale, -scale)
	toScreen.Translate(w/2, h-50)

	v.physics.Draw(screen, toScreen)

	text := fmt.Sprintf("physics %T: %1.2fms", v.physics, v.physicsTime*1000.0)
	ebitenutil.DebugPrintAt(screen, text, 16, 16)

	if v.paused {
		ebitenutil.DebugPrintAt(screen, "paused", 16, 32)
	}
}

func (v *viz) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (v *viz) initialize(makePhysics func(gravity float64) Physics) {
	ph := makePhysics(-10)

	var bodies []Body

	floor := BodyDef{Elasticity: 0.1, Friction: 0.9, Density: 1}
	box := BodyDef{Elasticity: 0.25, Friction: 0.5, Density: 1}
	big := BodyDef{Density: 10, Elasticity: 0.1, Friction: 0.9}

	bodies = append(bodies, ph.CreateStaticLine(-2*Layers, 0, 2*Layers, 0, floor))
	bodies = append(bodies, ph.CreateStaticLine(-2*Layers, 0, -2*Layers, 100, floor))
	bodies = append(bodies, ph.CreateStaticLine(2*Layers, 0, 2*Layers, 100, floor))

	// create layers of boxes
	for l := range Layers {
		for i := range l {
			centerX := float64(i) + 0.5 - float64(l)/2
			centerY := (Layers - float64(l) - 0.5) * 1.0

			bodies = append(bodies, ph.CreateSquare(0.5, centerX, centerY, box))
		}
	}

	// create a fast moving box
	b := ph.CreateSquare(2, -1.8*Layers, 2.6*Layers, big)
	b.SetVelocity(75, -100)
	bodies = append(bodies, b)

	v.physics = ph
	v.bodies = bodies
	v.paused = false
}

// screenLength converts a length in world units to pixels.
func screenLength(toScreen ebiten.GeoM, length float64) float64 {
	x0, y0 := toScreen.Apply(0, 0)
	x1, y1 := toScreen.Apply(length, 0)
	return math.Hypot(x1-x0, y1-y0)
}

type Body interface {
	Position() (float64, float64)
	Rotation() float64
	SetVelocity(x, y float64)
}

type BodyDef struct {
	Density    float64
	Friction   float64
	Elasticity float64
}

type Physics interface {
	Draw(screen *ebiten.Image, toScreen ebiten.GeoM)
	Step(dt float64, subSteps int)
	CreateSquare(halfSize, centerX, centerY float64, def BodyDef) Body
	CreateStaticLine(x0, y0, x1, y1 float64, def BodyDef) Body
}
