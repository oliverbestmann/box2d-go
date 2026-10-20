package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/jakecoffman/cp/v2"
)

// cpNew ignores concurrent, chipmunk is single threaded.
func cpNew(gravity float64, _ bool) Physics {
	space := cp.NewSpace()
	space.SetGravity(cp.Vector{Y: gravity})
	return Chipmunk{Space: space}
}

type Chipmunk struct {
	Space *cp.Space
}

func (ph Chipmunk) Step(dt float64, subSteps int) {
	ph.Space.Iterations = uint(subSteps)
	ph.Space.Step(dt)
}

func (ph Chipmunk) Destroy() {}

func (ph Chipmunk) Draw(screen *ebiten.Image, toScreen ebiten.GeoM) {
	cp.DrawSpace(ph.Space, drawer{screen, toScreen})
}

func (ph Chipmunk) CreateSquare(halfSize, centerX, centerY float64, d BodyDef) Body {
	body := cp.NewBody(1, 1)

	shape := cp.NewBox(body, halfSize*2, halfSize*2, 0)
	shape.SetElasticity(d.Elasticity)
	shape.SetDensity(d.Density)
	shape.SetFriction(d.Friction)
	body.AddShape(shape)
	body.SetPosition(cp.Vector{X: centerX, Y: centerY})

	ph.Space.AddShape(shape)
	ph.Space.AddBody(body)

	return &cpBody{
		body: body,
	}
}

func (ph Chipmunk) CreateStaticLine(x0, y0, x1, y1 float64, d BodyDef) {
	body := cp.NewStaticBody()

	a := cp.Vector{X: x0, Y: y0}
	b := cp.Vector{X: x1, Y: y1}
	shape := cp.NewSegment(body, a, b, 0)
	shape.SetElasticity(d.Elasticity)
	shape.SetDensity(d.Density)
	shape.SetFriction(d.Friction)
	body.AddShape(shape)

	ph.Space.AddShape(shape)
	ph.Space.AddBody(body)
}

type cpBody struct {
	body *cp.Body
}

func (b *cpBody) SetVelocity(x, y float64) {
	b.body.SetVelocity(x, y)
}

type drawer struct {
	Screen   *ebiten.Image
	toScreen ebiten.GeoM
}

func (d drawer) DrawCircle(pos cp.Vector, angle, radius float64, outline, fill cp.FColor, data any) {
	var p vector.Path

	x, y := d.toScreen.Apply(pos.X, pos.Y)
	radius = screenLength(d.toScreen, radius)
	p.Arc(float32(x), float32(y), float32(radius), 0, 2*math.Pi, vector.Clockwise)

	dop := vector.DrawPathOptions{ColorScale: toColorScaleF(outline)}
	vector.StrokePath(d.Screen, &p, &vector.StrokeOptions{Width: 1}, &dop)

	dop.ColorScale = toColorScaleF(fill)
	vector.FillPath(d.Screen, &p, nil, &dop)
}

func (d drawer) DrawSegment(a, b cp.Vector, fill cp.FColor, data any) {
	var p vector.Path

	x, y := d.toScreen.Apply(a.X, a.Y)
	p.LineTo(float32(x), float32(y))

	x, y = d.toScreen.Apply(b.X, b.Y)
	p.LineTo(float32(x), float32(y))

	dop := vector.DrawPathOptions{ColorScale: toColorScaleF(fill)}
	vector.StrokePath(d.Screen, &p, &vector.StrokeOptions{Width: 1}, &dop)
}

func (d drawer) DrawFatSegment(a, b cp.Vector, radius float64, outline, fill cp.FColor, data any) {
	var p vector.Path

	x, y := d.toScreen.Apply(a.X, a.Y)
	p.LineTo(float32(x), float32(y))

	x, y = d.toScreen.Apply(b.X, b.Y)
	p.LineTo(float32(x), float32(y))

	// TODO make this actually a capsule based on radius

	dop := vector.DrawPathOptions{ColorScale: toColorScaleF(outline)}
	vector.StrokePath(d.Screen, &p, &vector.StrokeOptions{Width: 1}, &dop)

	dop.ColorScale = toColorScaleF(fill)
	vector.FillPath(d.Screen, &p, nil, &dop)
}

func (d drawer) DrawPolygon(count int, verts []cp.Vector, radius float64, outline, fill cp.FColor, data any) {
	var p vector.Path

	for _, v := range verts[:count] {
		x, y := d.toScreen.Apply(v.X, v.Y)
		p.LineTo(float32(x), float32(y))
	}
	p.Close()

	dop := vector.DrawPathOptions{ColorScale: toColorScaleF(outline)}
	vector.StrokePath(d.Screen, &p, &vector.StrokeOptions{Width: 1}, &dop)

	dop.ColorScale = toColorScaleF(fill)
	vector.FillPath(d.Screen, &p, nil, &dop)
}

func (d drawer) DrawDot(size float64, pos cp.Vector, fill cp.FColor, data any) {
	// size is in pixels, DrawCircle takes a radius in world units
	d.DrawCircle(pos, 0, size/2/screenLength(d.toScreen, 1), cp.FColor{}, fill, data)
}

func (d drawer) Flags() uint {
	return cp.DRAW_SHAPES
}

func (d drawer) OutlineColor() cp.FColor {
	return cp.FColor{R: 1, G: 1, B: 1, A: 1}
}

func (d drawer) ShapeColor(shape *cp.Shape, data any) cp.FColor {
	return cp.FColor{R: 1, G: 1, B: 1, A: 0.1}
}

func (d drawer) ConstraintColor() cp.FColor {
	return cp.FColor{R: 1, G: 0, B: 0, A: 1}
}

func (d drawer) CollisionPointColor() cp.FColor {
	return cp.FColor{}
}

func (d drawer) Data() any {
	return nil
}

func toColorScaleF(f cp.FColor) ebiten.ColorScale {
	var c ebiten.ColorScale
	c.Scale(f.R, f.G, f.B, 1)
	c.ScaleAlpha(f.A)
	return c
}
