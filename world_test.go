package b2

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Ported from box2d test/test_world.c

func TestHelloWorld(t *testing.T) {
	worldDef := DefaultWorldDef()
	worldDef.Gravity = Vec2{X: 0, Y: -10}
	world := CreateWorld(worldDef)
	require.EqualValues(t, 1, world.IsValid())

	groundDef := DefaultBodyDef()
	groundDef.Position = Vec2{X: 0, Y: -10}
	ground := world.CreateBody(groundDef)
	require.EqualValues(t, 1, ground.IsValid())
	ground.CreatePolygonShape(DefaultShapeDef(), MakeBox(50, 10))

	bodyDef := DefaultBodyDef()
	bodyDef.Type1 = DynamicBody
	bodyDef.Position = Vec2{X: 0, Y: 4}
	body := world.CreateBody(bodyDef)

	shapeDef := DefaultShapeDef()
	shapeDef.Density = 1
	shapeDef.Material.Friction = 0.3
	body.CreatePolygonShape(shapeDef, MakeBox(1, 1))

	for range 90 {
		world.Step(1.0/60.0, 4)
	}

	position := body.GetPosition()
	rotation := body.GetRotation()

	world.DestroyWorld()

	assert.InDelta(t, 0, position.X, 0.01)
	assert.InDelta(t, 1, position.Y, 0.01)
	assert.InDelta(t, 0, rotation.Angle(), 0.01)
}

func TestEmptyWorld(t *testing.T) {
	world := CreateWorld(DefaultWorldDef())
	require.EqualValues(t, 1, world.IsValid())

	for range 60 {
		world.Step(1.0/60.0, 1)
	}

	world.DestroyWorld()
	assert.EqualValues(t, 0, world.IsValid())
}

func TestDestroyAllBodiesWorld(t *testing.T) {
	const bodyCount = 10

	world := CreateWorld(DefaultWorldDef())
	defer world.DestroyWorld()

	bodyDef := DefaultBodyDef()
	bodyDef.Type1 = DynamicBody
	square := MakeSquare(0.5)

	var bodies []Body
	creating := true

	for range 2*bodyCount + 10 {
		switch {
		case creating && len(bodies) < bodyCount:
			body := world.CreateBody(bodyDef)
			body.CreatePolygonShape(DefaultShapeDef(), square)
			bodies = append(bodies, body)

		case creating:
			creating = false

		case len(bodies) > 0:
			bodies[len(bodies)-1].DestroyBody()
			bodies = bodies[:len(bodies)-1]
		}

		world.Step(1.0/60.0, 3)
	}

	assert.EqualValues(t, 0, world.GetCounters().BodyCount)
}

func TestIsValid(t *testing.T) {
	world := CreateWorld(DefaultWorldDef())
	require.EqualValues(t, 1, world.IsValid())

	body1 := world.CreateBody(DefaultBodyDef())
	body2 := world.CreateBody(DefaultBodyDef())
	assert.EqualValues(t, 1, body1.IsValid())
	assert.EqualValues(t, 1, body2.IsValid())

	body1.DestroyBody()
	assert.EqualValues(t, 0, body1.IsValid())
	assert.EqualValues(t, 1, body2.IsValid())

	body2.DestroyBody()
	assert.EqualValues(t, 0, body2.IsValid())

	world.DestroyWorld()
	assert.EqualValues(t, 0, world.IsValid())
	assert.EqualValues(t, 0, body1.IsValid())
	assert.EqualValues(t, 0, body2.IsValid())
}

func TestWorldRecycle(t *testing.T) {
	worlds := make([]World, _B2_MAX_WORLDS/2)

	for range 20 {
		for idx := range worlds {
			worlds[idx] = CreateWorld(DefaultWorldDef())
			require.EqualValues(t, 1, worlds[idx].IsValid())
			worlds[idx].CreateBody(DefaultBodyDef())
		}

		for _, world := range worlds {
			for range 10 {
				world.Step(1.0/60.0, 1)
			}
		}

		for idx := len(worlds) - 1; idx >= 0; idx-- {
			worlds[idx].DestroyWorld()
			require.EqualValues(t, 0, worlds[idx].IsValid())
		}
	}
}

func TestWorldCoverage(t *testing.T) {
	world := CreateWorld(DefaultWorldDef())
	defer world.DestroyWorld()

	world.EnableSleeping(1)
	world.EnableSleeping(0)
	assert.EqualValues(t, 0, world.IsSleepingEnabled())

	world.EnableContinuous(0)
	world.EnableContinuous(1)
	assert.EqualValues(t, 1, world.IsContinuousEnabled())

	world.SetRestitutionThreshold(0)
	world.SetRestitutionThreshold(2)
	assert.Equal(t, float32(2), world.GetRestitutionThreshold())

	world.SetHitEventThreshold(0)
	world.SetHitEventThreshold(100)
	assert.Equal(t, float32(100), world.GetHitEventThreshold())

	world.SetGravity(Vec2{X: 1, Y: 2})
	assert.Equal(t, Vec2{X: 1, Y: 2}, world.GetGravity())

	world.Explode(DefaultExplosionDef())
	world.SetContactTuning(10, 2, 4)

	world.SetMaxLinearSpeed(10)
	assert.Equal(t, float32(10), world.GetMaxLinearSpeed())

	world.EnableWarmStarting(1)
	assert.EqualValues(t, 1, world.IsWarmStartingEnabled())

	assert.EqualValues(t, 0, world.GetAwakeBodyCount())

	world.SetUserData(42)
	assert.EqualValues(t, 42, world.GetUserData())

	world.Step(1, 1)
}

func TestSensorEvents(t *testing.T) {
	world := CreateWorld(DefaultWorldDef())
	defer world.DestroyWorld()

	// wall from x = 1 to x = 2
	wallDef := DefaultBodyDef()
	wallDef.Position = Vec2{X: 1.5, Y: 11}
	wall := world.CreateBody(wallDef)

	wallShape := DefaultShapeDef()
	wallShape.EnableSensorEvents = 1
	wall.CreatePolygonShape(wallShape, MakeBox(0.5, 10))

	// bullet fired towards the wall
	bulletDef := DefaultBodyDef()
	bulletDef.Type1 = DynamicBody
	bulletDef.IsBullet = 1
	bulletDef.GravityScale = 0
	bulletDef.Position = Vec2{X: 7.39814, Y: 4}
	bulletDef.LinearVelocity = Vec2{X: -20}
	bullet := world.CreateBody(bulletDef)

	bulletShape := DefaultShapeDef()
	bulletShape.IsSensor = 1
	bulletShape.EnableSensorEvents = 1
	bullet.CreateCircleShape(bulletShape, Circle{Radius: 0.1})

	var beginCount, endCount int

	for bullet.GetPosition().X >= -1 {
		world.Step(1.0/60.0, 4)

		events := world.GetSensorEvents()
		if len(events.BeginEvents) > 0 {
			beginCount++
		}

		if len(events.EndEvents) > 0 {
			endCount++
		}
	}

	assert.Equal(t, 1, beginCount)
	assert.Equal(t, 1, endCount)
}

func TestBounce(t *testing.T) {
	world := CreateWorld(DefaultWorldDef())
	defer world.DestroyWorld()

	world.SetGravity(Vec2{Y: -10})

	ballDef := DefaultBodyDef()
	ballDef.Type1 = DynamicBody
	ballDef.Position.Y = 5
	ball := world.CreateBody(ballDef)

	ballShape := DefaultShapeDef()
	ballShape.Material.Restitution = 0.2
	ballShape.EnableContactEvents = 1
	ballShape.EnableHitEvents = 1
	ballShapeId := ball.CreateCircleShape(ballShape, Circle{Radius: 1})

	ground := world.CreateBody(DefaultBodyDef())
	ground.CreateSegmentShape(DefaultShapeDef(), Segment{
		Point1: Vec2{X: -20},
		Point2: Vec2{X: 20},
	})

	var moveEvents, beginEvents, hitEvents int
	var bounced bool

	for range 100 {
		world.Step(1.0/60.0, 1)

		for _, ev := range world.GetBodyEvents().MoveEvents {
			assert.Equal(t, ball.Id, ev.BodyId)
			moveEvents++
		}

		contacts := world.GetContactEvents()
		for _, ev := range contacts.BeginEvents {
			assert.Contains(t, []ShapeId{ev.ShapeIdA, ev.ShapeIdB}, ballShapeId.Id)
			beginEvents++
		}

		hitEvents += len(contacts.HitEvents)

		if ball.GetLinearVelocity().Y > 0 {
			bounced = true
		}
	}

	assert.Positive(t, moveEvents)
	assert.Positive(t, beginEvents)
	assert.Positive(t, hitEvents)
	assert.True(t, bounced, "ball should bounce back up")

	// the ball rests on the ground
	assert.InDelta(t, 1, ball.GetPosition().Y, 0.05)
}
