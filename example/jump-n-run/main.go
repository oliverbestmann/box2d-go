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

	if err := ebiten.RunGame(&game{scene: NewScene(), sprites: map[spriteKey]sprite{}}); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

type game struct {
	scene   *Scene
	shapes  []b2.Shape
	sprites map[spriteKey]sprite
}

// spriteKey identifies a shape by its geometry, so equal bodies share one sprite.
type spriteKey struct {
	color   color.RGBA
	polygon b2.Polygon
	circle  b2.Circle
}

// sprite is a shape drawn in body coordinates. ox, oy is the body origin in the image.
type sprite struct {
	image  *ebiten.Image
	ox, oy float64
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
	x, y := toScreen(tf.P)
	// y is flipped on screen, so a counterclockwise world rotation is a negative screen rotation
	angle := -math.Atan2(float64(tf.Q.S), float64(tf.Q.C))

	g.shapes = t.body.GetShapes(g.shapes)
	for _, shape := range g.shapes {
		s, ok := g.sprite(shape, t.color)
		if !ok {
			continue
		}

		op := &ebiten.DrawImageOptions{Filter: ebiten.FilterLinear}
		op.GeoM.Translate(-s.ox, -s.oy)
		op.GeoM.Rotate(angle)
		op.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(s.image, op)
	}
}

// sprite returns the cached image of a polygon or circle shape and draws it on first use.
func (g *game) sprite(shape b2.Shape, c color.RGBA) (sprite, bool) {
	key := spriteKey{color: c}

	var lo, hi b2.Vec2
	switch shape.GetType() {
	case b2.PolygonShape:
		key.polygon = shape.GetPolygon()
		lo, hi = key.polygon.Vertices[0], key.polygon.Vertices[0]
		for _, v := range key.polygon.Vertices[:key.polygon.Count] {
			lo = b2.Vec2{X: min(lo.X, v.X), Y: min(lo.Y, v.Y)}
			hi = b2.Vec2{X: max(hi.X, v.X), Y: max(hi.Y, v.Y)}
		}

	case b2.CircleShape:
		key.circle = shape.GetCircle()
		center, r := key.circle.Center, key.circle.Radius
		lo = b2.Vec2{X: center.X - r, Y: center.Y - r}
		hi = b2.Vec2{X: center.X + r, Y: center.Y + r}

	default:
		return sprite{}, false
	}

	if s, ok := g.sprites[key]; ok {
		return s, true
	}

	// room for the outline and anti aliasing
	const pad = 2

	ox, oy := -lo.X*pixelsPerM+pad, hi.Y*pixelsPerM+pad
	w := int(math.Ceil(float64((hi.X-lo.X)*pixelsPerM))) + 2*pad
	h := int(math.Ceil(float64((hi.Y-lo.Y)*pixelsPerM))) + 2*pad
	img := ebiten.NewImage(w, h)

	toImage := func(v b2.Vec2) (float32, float32) {
		return ox + v.X*pixelsPerM, oy - v.Y*pixelsPerM
	}

	if key.polygon.Count > 0 {
		var path vector.Path
		for _, v := range key.polygon.Vertices[:key.polygon.Count] {
			path.LineTo(toImage(v))
		}
		path.Close()

		fillPath(img, &path, c)
		strokePath(img, &path, darken(c))
	} else {
		circle := key.circle
		cx, cy := toImage(circle.Center)
		r := circle.Radius * pixelsPerM

		vector.FillCircle(img, cx, cy, r, c, true)
		vector.StrokeCircle(img, cx, cy, r, 1.5, darken(c), true)

		// a spoke to show the rotation
		ex, ey := toImage(b2.Vec2{X: circle.Center.X + circle.Radius, Y: circle.Center.Y})
		vector.StrokeLine(img, cx, cy, ex, ey, 1.5, darken(c), true)
	}

	s := sprite{image: img, ox: float64(ox), oy: float64(oy)}
	g.sprites[key] = s
	return s, true
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

func toScreen(p b2.Pos) (float32, float32) {
	return p.X * pixelsPerM, screenHeight - p.Y*pixelsPerM
}

func toWorld(x, y int) b2.Pos {
	return b2.Pos{X: float32(x) / pixelsPerM, Y: (screenHeight - float32(y)) / pixelsPerM}
}
