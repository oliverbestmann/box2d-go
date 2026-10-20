package main

import (
	"math"

	b2 "github.com/oliverbestmann/box2d-go/v3"
)

// Mover is a Go port of the DynamicMover from the box2d 3.2 samples (samples/dynamic_mover.cpp).
// The capsule is a dynamic body with locked rotation. A mover joint drives it to the target
// velocity and a pogo joint, rebuilt every update, holds it above the ground.
//
// Change from the sample: friction and acceleration work relative to the velocity of the
// ground, so the character rides along on horizontally moving platforms. Gravity is stronger
// while falling or after the jump button is released, for a short and snappy jump arc.
type Mover struct {
	JumpSpeed            float32
	MaxSpeed             float32
	MinSpeed             float32
	StopSpeed            float32
	Accelerate           float32
	AirSteer             float32
	Friction             float32
	GravityScale         float32
	FallGravityScale     float32
	MaxGroundForce       float32
	MaxAirForce          float32
	PogoRestLength       float32
	PogoHertz            float32
	PogoDampingRatio     float32
	PogoCompressionScale float32
	PogoTensionScale     float32
	MinGroundNormalY     float32

	Capsule b2.Capsule
	Filter  b2.Filter

	World b2.World
	Body  b2.Body

	ground  b2.Body
	joint   b2.MoverJoint
	pogo    b2.PogoJoint
	hasPogo bool

	cast moverCast

	Velocity b2.Vec2
	OnGround bool
	walkable bool
	jumping  bool
	jumpTime float32
	jumpCut  bool

	pogoImpulse  float32
	pogoVelocity float32

	// ground probe of the last update, for drawing
	pogoAnchor    b2.Vec2 // local point on cast.body
	pogoRayLength float32
}

type moverCast struct {
	point    b2.Pos
	normal   b2.Vec2
	body     b2.Body
	fraction float32
	hit      bool
}

func NewMover(world b2.World, position b2.Pos, filter b2.Filter) *Mover {
	m := &Mover{
		JumpSpeed:            13,
		MaxSpeed:             7,
		MinSpeed:             0.1,
		StopSpeed:            3,
		Accelerate:           20,
		AirSteer:             0.8,
		Friction:             8,
		GravityScale:         3,
		FallGravityScale:     5,
		MaxGroundForce:       70,
		MaxAirForce:          40,
		PogoRestLength:       0.9,
		PogoHertz:            5,
		PogoDampingRatio:     0.8,
		PogoCompressionScale: 100,
		PogoTensionScale:     100,
		MinGroundNormalY:     0.7,

		Capsule: b2.Capsule{Center1: b2.Vec2{Y: -0.5}, Center2: b2.Vec2{Y: 0.5}, Radius: 0.3},
		Filter:  filter,
		World:   world,
	}

	// the mover joint needs a reference body, it has no shapes
	m.ground = world.CreateBody(b2.DefaultBodyDef())

	bodyDef := b2.DefaultBodyDef()
	bodyDef.Type1 = b2.DynamicBody
	bodyDef.Position = position
	bodyDef.GravityScale = m.GravityScale
	bodyDef.MotionLocks.AngularZ = 1
	bodyDef.EnableSleep = 0
	m.Body = world.CreateBody(bodyDef)

	shapeDef := b2.DefaultShapeDef()
	shapeDef.Density = 1

	// the mover joint and the pogo joint handle all the movement, surface friction would fight them
	shapeDef.Material.Friction = 0
	shapeDef.Filter = filter
	shapeDef.EnablePreSolveEvents = 1
	m.Body.CreateCapsuleShape(shapeDef, m.Capsule)

	moverDef := b2.DefaultMoverJointDef()
	moverDef.MaxVelocityForce = b2.Vec2{X: m.MaxGroundForce}
	moverDef.Base.BodyIdA = m.ground.Id
	moverDef.Base.BodyIdB = m.Body.Id
	moverDef.Base.CollideConnected = 1
	m.joint = world.CreateMoverJoint(moverDef)

	return m
}

// Jump jumps if grounded and returns false when airborne.
func (m *Mover) Jump() bool {
	if !m.OnGround || !m.walkable {
		return false
	}

	var surfaceVelocity float32
	if m.cast.body.IsValid() != 0 {
		surfaceVelocity = m.cast.body.GetWorldPointVelocity(m.cast.point).Y
	}

	mass := m.Body.GetMass()
	vy := m.Body.GetLinearVelocity().Y

	// remove the pogo constraint velocity but add the surface velocity
	dv := max(0, m.JumpSpeed-vy) + surfaceVelocity
	m.Body.ApplyLinearImpulseToCenter(b2.Vec2{Y: mass * dv}, 1)

	// this removes the pogo step down on the next update
	m.OnGround = false
	m.walkable = false
	m.jumping = true
	m.jumpTime = 0
	m.jumpCut = false
	return true
}

// Update advances the mover. throttle is the horizontal input in [-1, 1], jumpHeld
// is true while the jump button is down. Movement follows the Quake ground and air model.
func (m *Mover) Update(timeStep, throttle float32, jumpHeld bool) {
	// reach further while grounded so the spring finds the ground over a step
	rayLength := m.PogoRestLength
	if m.OnGround {
		rayLength *= 2
	}

	translation := b2.Vec2{Y: -rayLength}
	m.pogoRayLength = rayLength
	position := m.Body.GetPosition()
	origin := b2.Pos{X: position.X + m.Capsule.Center1.X, Y: position.Y + m.Capsule.Center1.Y}

	queryFilter := b2.QueryFilter{CategoryBits: m.Filter.CategoryBits, MaskBits: m.Filter.MaskBits}

	m.cast = moverCast{}
	m.World.CastRay(origin, translation, queryFilter, func(shape b2.Shape, point b2.Pos, normal b2.Vec2, fraction float32) float32 {
		m.cast = moverCast{point: point, normal: normal, body: shape.GetBody(), fraction: fraction, hit: true}
		return fraction
	})

	if m.jumping {
		m.jumpTime += timeStep

		// give the jump time to leave the ground (2 ticks at 60 Hz in the sample), then end
		// it when approaching the ground
		if m.jumpTime > 2.5/60 && m.cast.hit && dot(m.Velocity, m.cast.normal) <= 0 {
			m.jumping = false
		}
	}

	m.OnGround = m.cast.hit && !m.jumping

	// the pogo weight below stays based on GravityScale
	gravityScale := m.GravityScale
	if !m.OnGround && (!jumpHeld || m.Body.GetLinearVelocity().Y < 0) {
		gravityScale = m.FallGravityScale
	}
	m.Body.SetGravityScale(gravityScale)

	// releasing jump while rising halves the upward speed once, for short hops
	if m.jumping && !m.jumpCut && !jumpHeld {
		m.jumpCut = true
		if v := m.Body.GetLinearVelocity(); v.Y > 0 {
			m.Body.SetLinearVelocity(b2.Vec2{X: v.X, Y: v.Y / 2})
		}
	}
	m.walkable = m.cast.hit && m.cast.normal.Y >= m.MinGroundNormalY

	// work relative to the ground so moving platforms carry the character
	var groundVelocity b2.Vec2
	if m.OnGround {
		groundVelocity = m.cast.body.GetWorldPointVelocity(m.cast.point)
	}

	v := m.Body.GetLinearVelocity()
	v.X -= groundVelocity.X
	v.Y -= groundVelocity.Y

	// friction
	if m.OnGround {
		speed := length(v)
		if speed < m.MinSpeed {
			v.X = 0
		} else {
			// linear damping above StopSpeed and fixed reduction below
			control := max(speed, m.StopSpeed)
			newSpeed := max(0, speed-control*m.Friction*timeStep)
			v.X *= newSpeed / speed
		}
	}

	// mover force
	maxForce := m.MaxAirForce
	if m.OnGround {
		maxForce = 0
		if m.walkable {
			maxForce = m.MaxGroundForce
		}
	}
	m.joint.SetMaxVelocityForce(b2.Vec2{X: maxForce})

	// accelerate
	throttle = min(max(throttle, -1), 1)
	if throttle != 0 {
		direction := float32(math.Copysign(1, float64(throttle)))
		desiredSpeed := m.MaxSpeed * throttle * direction // = MaxSpeed * |throttle|
		addSpeed := desiredSpeed - v.X*direction

		if addSpeed > 0 {
			steer := m.AirSteer
			if m.walkable {
				steer = 1
			}

			v.X += min(steer*m.Accelerate*m.MaxSpeed*timeStep, addSpeed) * direction
		}
	}

	v.X += groundVelocity.X
	v.Y += groundVelocity.Y
	m.Velocity = v

	// the pogo joint is rebuilt every update because it may land on a different body
	if m.hasPogo {
		m.pogoImpulse = m.pogo.GetImpulse()
		m.pogoVelocity = m.pogo.GetCurrentSpeed()
		m.pogo.DestroyJoint()
		m.hasPogo = false
	}

	if m.cast.hit {
		gravity := m.World.GetGravity()
		weight := m.GravityScale * length(gravity) * m.Body.GetMass()

		pogoDef := b2.DefaultPogoJointDef()
		pogoDef.Base.BodyIdA = m.cast.body.Id
		pogoDef.Base.BodyIdB = m.Body.Id
		m.pogoAnchor = m.cast.body.GetLocalPoint(m.cast.point)
		pogoDef.Base.LocalFrameA.P = m.pogoAnchor
		pogoDef.Base.LocalFrameB.P = m.Capsule.Center1
		pogoDef.Base.CollideConnected = 1
		pogoDef.Normal = m.cast.normal
		pogoDef.RestLength = m.PogoRestLength
		pogoDef.Hertz = m.PogoHertz
		pogoDef.DampingRatio = m.PogoDampingRatio
		pogoDef.MaxCompressionForce = m.PogoCompressionScale * weight

		// warm start from the joint that was just destroyed
		pogoDef.Impulse = m.pogoImpulse
		pogoDef.Velocity = m.pogoVelocity

		// don't let the pogo pull down while jumping
		if !m.jumping {
			pogoDef.MaxTensionForce = m.PogoTensionScale * weight
		}

		m.pogo = m.World.CreatePogoJoint(pogoDef)
		m.hasPogo = true
	} else {
		m.pogoImpulse = 0
		m.pogoVelocity = 0
	}

	m.joint.SetLinearVelocity(m.Velocity)
}

// Pogo returns the ends of the pogo: the foot on the capsule and, if the ground probe hit,
// the anchor on the ground. Without a hit, ground is the end of the probe ray.
func (m *Mover) Pogo() (foot, ground b2.Pos, hit bool) {
	foot = m.Body.GetWorldPoint(m.Capsule.Center1)
	if !m.hasPogo {
		return foot, b2.Pos{X: foot.X, Y: foot.Y - m.pogoRayLength}, false
	}

	return foot, m.cast.body.GetWorldPoint(m.pogoAnchor), true
}

func dot(a, b b2.Vec2) float32 {
	return a.X*b.X + a.Y*b.Y
}

func length(v b2.Vec2) float32 {
	return float32(math.Hypot(float64(v.X), float64(v.Y)))
}
