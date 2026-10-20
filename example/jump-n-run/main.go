// A small physics sandbox jump and run using the dynamic character mover of box2d 3.2.
package main

import (
	"fmt"
	"image/color"
	"math"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	b2 "github.com/oliverbestmann/box2d-go/v3"
)

const (
	screenWidth  = 1280
	screenHeight = 720
	pixelsPerM   = screenWidth / worldWidth
)

const help = `a/d: move  w/space: jump  mouse: shoot  r: reset`

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("box2d jump and run")

	// Draw runs once per display frame, Update (one physics step) tps times a second.
	// shortcut: fixed 120 Hz, not the monitor's real rate, ebiten does not expose it.
	ebiten.SetTPS(tps)

	if err := ebiten.RunGame(&game{scene: NewScene()}); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

type game struct {
	scene  *Scene
	shapes []b2.Shape
}

func (g *game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.scene.Destroy()
		g.scene = NewScene()
	}

	var throttle float32
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		throttle--
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		throttle++
	}

	jumpKeys := []ebiten.Key{ebiten.KeySpace, ebiten.KeyW, ebiten.KeyUp}
	jump := slices.ContainsFunc(jumpKeys, inpututil.IsKeyJustPressed)
	jumpHeld := slices.ContainsFunc(jumpKeys, ebiten.IsKeyPressed)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.scene.Shoot(toWorld(ebiten.CursorPosition()))
	}

	g.scene.Step(throttle, jump, jumpHeld)
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x1d, 0x1f, 0x2b, 0xff})

	s := g.scene

	for _, t := range s.Things {
		g.drawBody(screen, t)
	}

	for _, p := range s.Projectiles {
		g.drawBody(screen, p.thing)
	}

	// welds, so you can see where the tower cracks
	for _, w := range s.Welds {
		ax, ay := toScreen(w.a.GetPosition())
		bx, by := toScreen(w.b.GetPosition())
		vector.StrokeLine(screen, ax, ay, bx, by, 2, color.RGBA{0x6e, 0x2a, 0x22, 0xff}, true)
	}

	g.drawPlayer(screen)

	ebitenutil.DebugPrintAt(screen, help, 8, 4)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("fps %.0f  tps %.0f  bodies %d  welds %d",
		ebiten.ActualFPS(), ebiten.ActualTPS(), len(s.Things)+len(s.Projectiles)+1, len(s.Welds)), 8, 20)
}

func (g *game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func (g *game) drawBody(screen *ebiten.Image, t thing) {
	tf := t.body.GetTransform()

	g.shapes = t.body.GetShapes(g.shapes)
	for _, shape := range g.shapes {
		switch shape.GetType() {
		case b2.PolygonShape:
			poly := shape.GetPolygon()

			var path vector.Path
			for _, v := range poly.Vertices[:poly.Count] {
				path.LineTo(toScreen(transform(tf, v)))
			}
			path.Close()

			fillPath(screen, &path, t.color)
			strokePath(screen, &path, darken(t.color))

		case b2.CircleShape:
			circle := shape.GetCircle()
			cx, cy := toScreen(transform(tf, circle.Center))
			r := circle.Radius * pixelsPerM

			vector.FillCircle(screen, cx, cy, r, t.color, true)
			vector.StrokeCircle(screen, cx, cy, r, 1.5, darken(t.color), true)

			// a spoke to show the rotation
			ex, ey := toScreen(transform(tf, b2.Vec2{X: circle.Center.X + circle.Radius, Y: circle.Center.Y}))
			vector.StrokeLine(screen, cx, cy, ex, ey, 1.5, darken(t.color), true)
		}
	}
}

func (g *game) drawPlayer(screen *ebiten.Image) {
	m := g.scene.Player
	pos := m.Body.GetPosition()

	c := color.RGBA{0x6c, 0xd4, 0x7e, 0xff}
	if !m.OnGround {
		c = color.RGBA{0xf2, 0x9e, 0x4c, 0xff}
	}

	// the pogo, drawn below the capsule
	foot, ground, hit := m.Pogo()
	fx, fy := toScreen(foot)
	gx, gy := toScreen(ground)
	if hit {
		pogo := color.RGBA{0xdd, 0xa0, 0xdd, 0xff}
		vector.StrokeLine(screen, fx, fy, gx, gy, 3, pogo, true)
		vector.FillCircle(screen, gx, gy, 4, pogo, true)
	} else {
		gray := color.RGBA{0x80, 0x80, 0x80, 0xff}
		vector.StrokeLine(screen, fx, fy, gx, gy, 1, gray, true)
		vector.StrokeCircle(screen, gx, gy, 3, 1, gray, true)
	}

	// capsule as one thick round-capped line
	x1, y1 := toScreen(b2.Pos{X: pos.X + m.Capsule.Center1.X, Y: pos.Y + m.Capsule.Center1.Y})
	x2, y2 := toScreen(b2.Pos{X: pos.X + m.Capsule.Center2.X, Y: pos.Y + m.Capsule.Center2.Y})

	var path vector.Path
	path.MoveTo(x1, y1)
	path.LineTo(x2, y2)

	stroke := &vector.StrokeOptions{Width: 2 * m.Capsule.Radius * pixelsPerM, LineCap: vector.LineCapRound}
	vector.StrokePath(screen, &path, stroke, &vector.DrawPathOptions{AntiAlias: true, ColorScale: colorScale(c)})

	// eyes look at the cursor
	mx, my := ebiten.CursorPosition()
	ex, ey := float64(float32(mx)-x2), float64(float32(my)-y2)
	angle := math.Atan2(ey, ex)
	lookX, lookY := float32(1.5*math.Cos(angle)), float32(1.5*math.Sin(angle))

	for _, dx := range []float32{-4, 4} {
		vector.FillCircle(screen, x2+dx, y2, 3.5, color.White, true)
		vector.FillCircle(screen, x2+dx+lookX, y2+lookY, 1.8, color.Black, true)
	}
}

func fillPath(screen *ebiten.Image, path *vector.Path, c color.RGBA) {
	vector.FillPath(screen, path, nil, &vector.DrawPathOptions{AntiAlias: true, ColorScale: colorScale(c)})
}

func strokePath(screen *ebiten.Image, path *vector.Path, c color.RGBA) {
	stroke := &vector.StrokeOptions{Width: 1.5, LineJoin: vector.LineJoinRound}
	vector.StrokePath(screen, path, stroke, &vector.DrawPathOptions{AntiAlias: true, ColorScale: colorScale(c)})
}

func colorScale(c color.RGBA) ebiten.ColorScale {
	var cs ebiten.ColorScale
	cs.ScaleWithColor(c)
	return cs
}

func darken(c color.RGBA) color.RGBA {
	return color.RGBA{c.R / 2, c.G / 2, c.B / 2, c.A}
}

func transform(tf b2.Transform, v b2.Vec2) b2.Pos {
	return b2.Pos{
		X: tf.P.X + tf.Q.C*v.X - tf.Q.S*v.Y,
		Y: tf.P.Y + tf.Q.S*v.X + tf.Q.C*v.Y,
	}
}

func toScreen(p b2.Pos) (float32, float32) {
	return p.X * pixelsPerM, screenHeight - p.Y*pixelsPerM
}

func toWorld(x, y int) b2.Pos {
	return b2.Pos{X: float32(x) / pixelsPerM, Y: (screenHeight - float32(y)) / pixelsPerM}
}
