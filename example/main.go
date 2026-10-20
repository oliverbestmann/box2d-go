package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// sizes multiplies the number of squares in the scene, cycled with the s key.
var sizes = []float64{1, 5, 10, 20, 50}

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
	makePhysics func(gravity float64, concurrent bool) Physics
	physics     Physics
	bodies      []Body
	physicsTime float64
	paused      bool
	slowmo      bool
	timeAcc     float64
	size        int
	layers      int
	concurrent  bool
}

func (v *viz) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyB) {
		v.makePhysics = b2New
		v.initialize()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		v.makePhysics = cpNew
		v.initialize()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		v.size = (v.size + 1) % len(sizes)
		v.initialize()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		v.concurrent = !v.concurrent
		v.initialize()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		v.paused = !v.paused
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
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

const keyHelp = `b: box2d
c: jakecoffman/cp
s: scene size
p: concurrency
space: pause
m: slow motion`

func (v *viz) Draw(screen *ebiten.Image) {
	// the debug font is 6 pixels wide
	width := 0
	for line := range strings.Lines(keyHelp) {
		width = max(width, 6*len(strings.TrimSpace(line)))
	}

	// deferred so it is drawn on top of the scene
	defer ebitenutil.DebugPrintAt(screen, keyHelp, screen.Bounds().Dx()-16-width, 16)

	if v.physics == nil {
		ebitenutil.DebugPrintAt(screen, "press b for box2d, c for jakecoffman/cp", 16, 16)
		return
	}

	w, h := float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy())

	var toScreen ebiten.GeoM

	scale := w / (5 * float64(v.layers))

	toScreen.Scale(scale, -scale)
	toScreen.Translate(w/2, h-50)

	v.physics.Draw(screen, toScreen)

	text := fmt.Sprintf("physics %T: %1.2fms\nsquares: %dx, concurrency: %v",
		v.physics, v.physicsTime*1000.0, int(sizes[v.size]), v.concurrent)
	ebitenutil.DebugPrintAt(screen, text, 16, 16)

	if v.paused {
		ebitenutil.DebugPrintAt(screen, "paused", 16, 48)
	}
}

func (v *viz) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

// initialize (re)creates the scene with the current engine, size and concurrency.
func (v *viz) initialize() {
	if v.makePhysics == nil {
		return
	}

	if v.physics != nil {
		v.physics.Destroy()
	}

	ph := v.makePhysics(-10, v.concurrent)

	// a pyramid with l layers has l*(l-1)/2 squares, 190 at the base size
	layers := int(math.Round(20 * math.Sqrt(sizes[v.size])))
	L := float64(layers)

	var bodies []Body

	floor := BodyDef{Elasticity: 0.1, Friction: 0.9, Density: 1}
	box := BodyDef{Elasticity: 0.25, Friction: 0.5, Density: 1}
	big := BodyDef{Density: 10, Elasticity: 0.1, Friction: 0.9}

	bodies = append(bodies, ph.CreateStaticLine(-2*L, 0, 2*L, 0, floor))
	bodies = append(bodies, ph.CreateStaticLine(-2*L, 0, -2*L, 5*L, floor))
	bodies = append(bodies, ph.CreateStaticLine(2*L, 0, 2*L, 5*L, floor))

	// create layers of boxes
	for l := range layers {
		for i := range l {
			centerX := float64(i) + 0.5 - float64(l)/2
			centerY := L - float64(l) - 0.5

			bodies = append(bodies, ph.CreateSquare(0.5, centerX, centerY, box))
		}
	}

	// create a fast moving box
	b := ph.CreateSquare(2, -1.8*L, 2.6*L, big)
	b.SetVelocity(75, -100)
	bodies = append(bodies, b)

	v.physics = ph
	v.bodies = bodies
	v.layers = layers
	v.physicsTime = 0
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
	Destroy()
	CreateSquare(halfSize, centerX, centerY float64, def BodyDef) Body
	CreateStaticLine(x0, y0, x1, y1 float64, def BodyDef) Body
}
