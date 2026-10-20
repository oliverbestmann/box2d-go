package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	b2 "github.com/oliverbestmann/box2d-go/v3"
)

func b2New(gravity float64, concurrent bool) Physics {
	def := b2.DefaultWorldDef()
	def.Gravity = b2.Vec2{Y: float32(gravity)}

	if concurrent {
		b2.EnableConcurrency(&def)
	}

	return &Box2D{
		World: b2.CreateWorld(def),
	}
}

type Box2D struct {
	World b2.World
}

func (ph Box2D) Step(dt float64, subSteps int) {
	ph.World.Step(float32(dt), int32(subSteps))
}

func (ph Box2D) Destroy() {
	ph.World.DestroyWorld()
}

func (ph Box2D) Draw(screen *ebiten.Image, toScreen ebiten.GeoM) {
	draw := b2.DefaultDebugDraw()

	draw.DrawJoints = true
	draw.DrawShapes = true
	draw.DrawJointExtras = false
	draw.DrawShapeBounds = false
	draw.DrawMass = false
	draw.DrawBodyNames = false
	draw.DrawGraphColors = false
	draw.DrawContacts = false
	draw.DrawContactNormals = true
	draw.DrawContactForces = true
	draw.DrawContactFeatures = false
	draw.DrawFrictionForces = false
	draw.DrawIslands = false

	draw.DrawLine = func(p1 b2.Vec2, p2 b2.Vec2, color b2.HexColor) {
		x1, y1 := toScreen.Apply(float64(p1.X), float64(p1.Y))
		x2, y2 := toScreen.Apply(float64(p2.X), float64(p2.Y))

		var p vector.Path
		p.MoveTo(float32(x1), float32(y1))
		p.LineTo(float32(x2), float32(y2))

		dop := vector.DrawPathOptions{ColorScale: toColorScale(color)}
		vector.StrokePath(screen, &p, &vector.StrokeOptions{Width: 1}, &dop)
	}

	draw.DrawPolygon = func(tr b2.Transform, vertices []b2.Vec2, color b2.HexColor) {
		var g ebiten.GeoM
		g.Rotate(float64(tr.Q.Angle()))
		g.Translate(float64(tr.P.X), float64(tr.P.Y))
		g.Concat(toScreen)

		var p vector.Path

		for _, v := range vertices {
			x, y := g.Apply(float64(v.X), float64(v.Y))
			p.LineTo(float32(x), float32(y))
		}
		p.Close()

		dop := vector.DrawPathOptions{ColorScale: toColorScale(color)}
		vector.StrokePath(screen, &p, &vector.StrokeOptions{Width: 1}, &dop)
	}

	draw.DrawSolidPolygon = func(tr b2.Transform, vertices []b2.Vec2, radius float32, color b2.HexColor) {
		var g ebiten.GeoM
		g.Rotate(float64(tr.Q.Angle()))
		g.Translate(float64(tr.P.X), float64(tr.P.Y))
		g.Concat(toScreen)

		var p vector.Path

		for _, v := range vertices {
			x, y := g.Apply(float64(v.X), float64(v.Y))
			p.LineTo(float32(x), float32(y))
		}

		p.Close()

		dop := vector.DrawPathOptions{ColorScale: toColorScale(color)}
		vector.FillPath(screen, &p, nil, &dop)
	}

	draw.DrawCircle = func(center b2.Vec2, radius float32, color b2.HexColor) {
		x, y := toScreen.Apply(float64(center.X), float64(center.Y))
		r := screenLength(toScreen, float64(radius))

		var p vector.Path
		p.Arc(float32(x), float32(y), float32(r), 0, 2*math.Pi, vector.Clockwise)

		dop := vector.DrawPathOptions{ColorScale: toColorScale(color)}
		vector.StrokePath(screen, &p, &vector.StrokeOptions{Width: 1}, &dop)
	}

	draw.DrawSolidCircle = func(tr b2.Transform, center b2.Vec2, radius float32, color b2.HexColor) {
		var g ebiten.GeoM
		g.Rotate(float64(tr.Q.Angle()))
		g.Translate(float64(tr.P.X), float64(tr.P.Y))
		g.Concat(toScreen)

		x, y := g.Apply(float64(center.X), float64(center.Y))
		r := screenLength(toScreen, float64(radius))

		var p vector.Path
		p.Arc(float32(x), float32(y), float32(r), 0, 2*math.Pi, vector.Clockwise)

		dop := vector.DrawPathOptions{ColorScale: toColorScale(color)}
		vector.FillPath(screen, &p, nil, &dop)
	}

	draw.DrawSolidCapsule = func(p1 b2.Vec2, p2 b2.Vec2, radius float32, color b2.HexColor) {
		// TODO
	}

	draw.DrawTransform = func(tr b2.Transform) {
		// axis length in world units, like the box2d samples
		const axisScale = 0.2
		p := tr.P
		draw.DrawLine(p, b2.Vec2{X: p.X + axisScale*tr.Q.C, Y: p.Y + axisScale*tr.Q.S}, b2.ColorRed)
		draw.DrawLine(p, b2.Vec2{X: p.X - axisScale*tr.Q.S, Y: p.Y + axisScale*tr.Q.C}, b2.ColorGreen)
	}

	draw.DrawPoint = func(c b2.Vec2, size float32, color b2.HexColor) {
		x, y := toScreen.Apply(float64(c.X), float64(c.Y))

		var p vector.Path
		// size is the point diameter in pixels
		p.Arc(float32(x), float32(y), size/2, 0, 2*math.Pi, vector.Clockwise)

		dop := vector.DrawPathOptions{ColorScale: toColorScale(color)}
		vector.FillPath(screen, &p, nil, &dop)
	}

	draw.DrawString = func(p b2.Vec2, s string, color b2.HexColor) {
		x, y := toScreen.Apply(float64(p.X), float64(p.Y))
		ebitenutil.DebugPrintAt(screen, s, int(x), int(y))
	}

	ph.World.Draw(draw)
}

func toColorScale(h b2.HexColor) ebiten.ColorScale {
	r := float32((h>>16)&0xff) / 255.0
	g := float32((h>>8)&0xff) / 255.0
	b := float32((h>>0)&0xff) / 255.0

	var c ebiten.ColorScale
	c.Scale(r, g, b, 1)
	return c
}

func (ph Box2D) CreateSquare(halfSize, centerX, centerY float64, d BodyDef) Body {
	var tr b2.Transform
	tr.P.X = float32(centerX)
	tr.P.Y = float32(centerY)
	tr.Q.C = 1

	def := b2.DefaultBodyDef()
	def.Type1 = b2.DynamicBody
	body := ph.World.CreateBody(def)
	body.SetTransform(tr.P, tr.Q)

	shape := b2.DefaultShapeDef()
	shape.Density = float32(d.Density)
	shape.Material.Restitution = float32(d.Elasticity)
	shape.Material.Friction = float32(d.Friction)
	body.CreatePolygonShape(shape, b2.MakeSquare(float32(halfSize)))

	return &b2Body{body: body}
}

func (ph Box2D) CreateStaticLine(x0, y0, x1, y1 float64, d BodyDef) Body {
	def := b2.DefaultBodyDef()
	def.Type1 = b2.StaticBody

	shape := b2.DefaultShapeDef()
	shape.Density = float32(d.Density)
	shape.Material.Restitution = float32(d.Elasticity)
	shape.Material.Friction = float32(d.Friction)

	body := ph.World.CreateBody(def)
	body.CreateSegmentShape(shape, b2.Segment{
		Point1: b2.Vec2{X: float32(x0), Y: float32(y0)},
		Point2: b2.Vec2{X: float32(x1), Y: float32(y1)},
	})

	return &b2Body{
		body: body,
	}
}

type b2Body struct {
	body b2.Body
}

func (b *b2Body) Position() (float64, float64) {
	pos := b.body.GetPosition()
	return float64(pos.X), float64(pos.Y)
}

func (b *b2Body) Rotation() float64 {
	rot := b.body.GetRotation()
	return math.Atan2(float64(rot.S), float64(rot.C))
}

func (b *b2Body) SetVelocity(x, y float64) {
	b.body.SetLinearVelocity(b2.Vec2{X: float32(x), Y: float32(y)})
}
