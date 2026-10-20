package main

import (
	"image/color"
	"math"

	b2 "github.com/oliverbestmann/box2d-go/v3"
)

// The world is 40x22.5 meters, y points up, the origin is the bottom left corner.
const (
	worldWidth  = 40
	worldHeight = 22.5
	timeStep    = 1.0 / tps
	tps         = 120
)

// collision categories
const (
	categoryWorld      = 1
	categoryPlayer     = 2
	categoryProjectile = 4
)

// shape user data of platforms the player can jump through from below
const userDataOneWay = 1

var (
	colorStatic     = color.RGBA{0x4a, 0x4e, 0x69, 0xff}
	colorPlatform   = color.RGBA{0x9a, 0x8c, 0x98, 0xff}
	colorCrate      = color.RGBA{0xc9, 0x8b, 0x4a, 0xff}
	colorPlank      = color.RGBA{0xa0, 0x6b, 0x3c, 0xff}
	colorBrick      = color.RGBA{0xb5, 0x4a, 0x3c, 0xff}
	colorBall       = color.RGBA{0x5c, 0xa8, 0xd6, 0xff}
	colorProjectile = color.RGBA{0xf2, 0xe9, 0x4e, 0xff}
)

type thing struct {
	body  b2.Body
	color color.RGBA
}

type projectile struct {
	thing
	expires float64
}

type weld struct {
	a, b b2.Body
}

// platform is a kinematic body that moves between two points.
type platform struct {
	body       b2.Body
	from, to   b2.Pos
	period     float64
	phaseShift float64
}

type Scene struct {
	World       b2.World
	Player      *Mover
	Things      []thing
	Projectiles []projectile
	Welds       map[b2.JointId]weld
	platforms   []platform
	time        float64
	lastShot    float64
}

func NewScene() *Scene {
	def := b2.DefaultWorldDef()
	def.Gravity = b2.Vec2{Y: -10}

	s := &Scene{World: b2.CreateWorld(def), Welds: map[b2.JointId]weld{}}

	// frame: ground, walls and ceiling
	s.staticBox(0, 0, worldWidth, 1)
	s.staticBox(-1, 0, 0, worldHeight)
	s.staticBox(worldWidth, 0, worldWidth+1, worldHeight)
	s.staticBox(0, worldHeight, worldWidth, worldHeight+1)

	// ledges
	s.staticBox(5.8, 11.6, 13, 12).SetUserData(userDataOneWay)
	s.staticBox(29, 1, 32, 3.5)
	s.staticBox(33, 5.6, 40, 6).SetUserData(userDataOneWay)

	// elevator next to the left wall, up to the high ledge
	s.addPlatform(b2.Pos{X: 4, Y: 1.2}, b2.Pos{X: 4, Y: 12 - 0.2}, 1.5, 8, 0)
	// shuttle in the middle
	s.addPlatform(b2.Pos{X: 15, Y: 8.5}, b2.Pos{X: 27, Y: 8.5}, 2, 7, 0)

	// a slowly turning bar below the ceiling
	{
		bodyDef := b2.DefaultBodyDef()
		bodyDef.Type1 = b2.KinematicBody
		bodyDef.Position = b2.Pos{X: 21, Y: 16}
		bodyDef.AngularVelocity = 0.6
		body := s.World.CreateBody(bodyDef)
		body.CreatePolygonShape(b2.DefaultShapeDef(), b2.MakeBox(4, 0.2))
		s.Things = append(s.Things, thing{body, colorPlatform})
	}

	// a breakable tower of welded bricks in the middle
	s.weldedTower(20, 1, 3, 8, 0.4)

	// loose towers to knock over
	s.looseTower(11, 1, 4)
	s.looseTower(36, 6, 3)
	s.pyramid(9.75, 12, 5, 0.3)

	// some balls to kick around
	for i := range 4 {
		s.ball(b2.Pos{X: 26 + float32(i)*0.7, Y: 1.5}, 0.3, 1, colorBall)
	}

	s.Player = NewMover(s.World, b2.Pos{X: 2, Y: 3}, b2.Filter{
		CategoryBits: categoryPlayer,
		MaskBits:     ^uint64(categoryPlayer | categoryProjectile),
	})

	s.World.SetPreSolveCallback(s.preSolve, s.preContinuous)

	return s
}

func (s *Scene) preSolve(shapeA, shapeB b2.Shape, manifold *b2.Manifold) {
	if !s.collideOneWay(shapeA, shapeB, manifold.Normal) {
		manifold.PointCount = 0
	}
}

func (s *Scene) preContinuous(shapeA, shapeB b2.Shape, _ b2.Pos, normal b2.Vec2) bool {
	return s.collideOneWay(shapeA, shapeB, normal)
}

// collideOneWay returns false if the player touches a one-way platform from below or the side.
// normal points from shapeA to shapeB.
func (s *Scene) collideOneWay(shapeA, shapeB b2.Shape, normal b2.Vec2) bool {
	player := s.Player.Body.Id

	switch {
	case shapeA.GetUserData() == userDataOneWay && shapeB.GetBody().Id == player:
		return normal.Y > 0.7
	case shapeB.GetUserData() == userDataOneWay && shapeA.GetBody().Id == player:
		return normal.Y < -0.7
	default:
		return true
	}
}

// Step advances the scene by one time step with the given player input. jumpHeld
// is true while the jump button is down, releasing it early makes a shorter jump.
func (s *Scene) Step(throttle float32, jump, jumpHeld bool) {
	s.time += timeStep

	for _, p := range s.platforms {
		t := 0.5 - 0.5*math.Cos(2*math.Pi*(s.time+p.phaseShift)/p.period)
		target := b2.Pos{
			X: p.from.X + float32(t)*(p.to.X-p.from.X),
			Y: p.from.Y + float32(t)*(p.to.Y-p.from.Y),
		}
		p.body.SetTargetTransform(b2.WorldTransform{P: target, Q: b2.IdentityRot}, timeStep, 1)
	}

	if jump {
		s.Player.Jump()
	}

	s.Player.Update(timeStep, throttle, jumpHeld)

	s.World.Step(timeStep, 4)

	// break welds that took too much force
	for _, event := range s.World.GetJointEvents().JointEvents {
		if _, ok := s.Welds[event.JointId]; ok {
			delete(s.Welds, event.JointId)
			b2.Joint{Id: event.JointId}.DestroyJoint()
		}
	}

	// remove old projectiles
	alive := s.Projectiles[:0]
	for _, p := range s.Projectiles {
		if p.expires < s.time {
			p.body.DestroyBody()
			continue
		}
		alive = append(alive, p)
	}
	s.Projectiles = alive
}

// Shoot fires a projectile from the player towards target, limited to a few shots per second.
func (s *Scene) Shoot(target b2.Pos) {
	if s.time-s.lastShot < 0.15 {
		return
	}

	origin := s.Player.Body.GetPosition()
	dir := b2.Vec2{X: target.X - origin.X, Y: target.Y - origin.Y}
	l := length(dir)
	if l < 0.01 {
		return
	}

	dir.X /= l
	dir.Y /= l

	s.lastShot = s.time

	bodyDef := b2.DefaultBodyDef()
	bodyDef.Type1 = b2.DynamicBody
	bodyDef.Position = b2.Pos{X: origin.X + 0.6*dir.X, Y: origin.Y + 0.6*dir.Y}
	bodyDef.LinearVelocity = b2.Vec2{X: 30 * dir.X, Y: 30 * dir.Y}
	bodyDef.IsBullet = 1
	body := s.World.CreateBody(bodyDef)

	shapeDef := b2.DefaultShapeDef()
	shapeDef.Density = 20
	shapeDef.Material.Restitution = 0.3
	shapeDef.Filter = b2.Filter{CategoryBits: categoryProjectile, MaskBits: ^uint64(categoryPlayer)}
	body.CreateCircleShape(shapeDef, b2.Circle{Radius: 0.15})

	s.Projectiles = append(s.Projectiles, projectile{thing{body, colorProjectile}, s.time + 5})
}

func (s *Scene) Destroy() {
	s.World.DestroyWorld()
}

func (s *Scene) staticBox(x0, y0, x1, y1 float32) b2.Shape {
	bodyDef := b2.DefaultBodyDef()
	bodyDef.Position = b2.Pos{X: (x0 + x1) / 2, Y: (y0 + y1) / 2}
	body := s.World.CreateBody(bodyDef)
	shape := body.CreatePolygonShape(b2.DefaultShapeDef(), b2.MakeBox((x1-x0)/2, (y1-y0)/2))
	s.Things = append(s.Things, thing{body, colorStatic})
	return shape
}

func (s *Scene) addPlatform(from, to b2.Pos, halfWidth float32, period, phaseShift float64) {
	bodyDef := b2.DefaultBodyDef()
	bodyDef.Type1 = b2.KinematicBody
	bodyDef.Position = from
	body := s.World.CreateBody(bodyDef)
	body.CreatePolygonShape(b2.DefaultShapeDef(), b2.MakeBox(halfWidth, 0.2)).SetUserData(userDataOneWay)

	s.Things = append(s.Things, thing{body, colorPlatform})
	s.platforms = append(s.platforms, platform{body, from, to, period, phaseShift})
}

func (s *Scene) box(center b2.Pos, halfWidth, halfHeight, density float32, c color.RGBA) b2.Body {
	bodyDef := b2.DefaultBodyDef()
	bodyDef.Type1 = b2.DynamicBody
	bodyDef.Position = center
	body := s.World.CreateBody(bodyDef)

	shapeDef := b2.DefaultShapeDef()
	shapeDef.Density = density
	shapeDef.Material.Friction = 0.7
	body.CreatePolygonShape(shapeDef, b2.MakeBox(halfWidth, halfHeight))

	s.Things = append(s.Things, thing{body, c})
	return body
}

func (s *Scene) ball(center b2.Pos, radius, density float32, c color.RGBA) {
	bodyDef := b2.DefaultBodyDef()
	bodyDef.Type1 = b2.DynamicBody
	bodyDef.Position = center
	body := s.World.CreateBody(bodyDef)

	shapeDef := b2.DefaultShapeDef()
	shapeDef.Density = density
	shapeDef.Material.Restitution = 0.5
	shapeDef.Material.RollingResistance = 0.1
	body.CreateCircleShape(shapeDef, b2.Circle{Radius: radius})

	s.Things = append(s.Things, thing{body, c})
}

// weldedTower builds cols x rows bricks standing on y, welded to their neighbors.
// The welds break when the force on them gets too large.
func (s *Scene) weldedTower(centerX, y float32, cols, rows int, half float32) {
	bricks := make([][]b2.Body, cols)
	for c := range cols {
		bricks[c] = make([]b2.Body, rows)
		for r := range rows {
			center := b2.Pos{
				X: centerX + (float32(c)-float32(cols-1)/2)*2*half,
				Y: y + half + float32(r)*2*half,
			}
			bricks[c][r] = s.box(center, half, half, 1, colorBrick)
		}
	}

	for c := range cols {
		for r := range rows {
			if c+1 < cols {
				s.weld(bricks[c][r], bricks[c+1][r])
			}
			if r+1 < rows {
				s.weld(bricks[c][r], bricks[c][r+1])
			}
		}
	}
}

func (s *Scene) weld(a, b b2.Body) {
	pa, pb := a.GetPosition(), b.GetPosition()
	mid := b2.Pos{X: (pa.X + pb.X) / 2, Y: (pa.Y + pb.Y) / 2}

	def := b2.DefaultWeldJointDef()
	def.Base.BodyIdA = a.Id
	def.Base.BodyIdB = b.Id
	def.Base.LocalFrameA.P = a.GetLocalPoint(mid)
	def.Base.LocalFrameB.P = b.GetLocalPoint(mid)
	def.Base.ForceThreshold = 400

	joint := s.World.CreateWeldJoint(def)
	s.Welds[joint.Id] = weld{a, b}
}

// looseTower stacks floors of two columns with a plank on top, standing on y.
func (s *Scene) looseTower(centerX, y float32, floors int) {
	const columnHalfHeight, plankHalfHeight = 0.6, 0.12

	for range floors {
		s.box(b2.Pos{X: centerX - 0.8, Y: y + columnHalfHeight}, 0.12, columnHalfHeight, 1, colorPlank)
		s.box(b2.Pos{X: centerX + 0.8, Y: y + columnHalfHeight}, 0.12, columnHalfHeight, 1, colorPlank)
		y += 2 * columnHalfHeight

		s.box(b2.Pos{X: centerX, Y: y + plankHalfHeight}, 1.1, plankHalfHeight, 1, colorPlank)
		y += 2 * plankHalfHeight
	}

	s.box(b2.Pos{X: centerX, Y: y + 0.35}, 0.35, 0.35, 1, colorCrate)
}

// pyramid stacks crates with the given base count, standing on y.
func (s *Scene) pyramid(centerX, y float32, base int, half float32) {
	for row := range base {
		count := base - row
		for i := range count {
			x := centerX + (float32(i)-float32(count-1)/2)*2*half
			s.box(b2.Pos{X: x, Y: y + half + float32(row)*2*half}, half, half, 0.5, colorCrate)
		}
	}
}
