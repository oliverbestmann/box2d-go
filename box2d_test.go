package b2

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Tests for the hand-written wrappers in box2d.go, draw.go and box2d-worker.go.

// newTestWorld creates a world with a static ground box whose top is at y = 0.
func newTestWorld(t *testing.T, def WorldDef) (World, Shape) {
	world := CreateWorld(def)
	t.Cleanup(world.DestroyWorld)

	groundDef := DefaultBodyDef()
	groundDef.Position = Vec2{Y: -1}
	ground := world.CreateBody(groundDef)
	groundShape := ground.CreatePolygonShape(DefaultShapeDef(), MakeBox(20, 1))

	return world, groundShape
}

func createBox(world World, pos Vec2, shapeDef ShapeDef) (Body, Shape) {
	bodyDef := DefaultBodyDef()
	bodyDef.Type1 = DynamicBody
	bodyDef.Position = pos
	body := world.CreateBody(bodyDef)
	return body, body.CreatePolygonShape(shapeDef, MakeBox(0.5, 0.5))
}

func TestWorld_Draw(t *testing.T) {
	world := CreateWorld(DefaultWorldDef())
	defer world.DestroyWorld()

	seg := Segment{Point1: Vec2{X: 1, Y: 1}, Point2: Vec2{X: 2, Y: 5}}

	body := world.CreateBody(DefaultBodyDef())
	body.SetName("Foobar")
	body.CreateSegmentShape(DefaultShapeDef(), seg)

	var segments [][2]Vec2
	var strings []string

	world.Draw(DebugDraw{
		DrawSegment: func(p1 Vec2, p2 Vec2, color HexColor) {
			segments = append(segments, [2]Vec2{p1, p2})
		},
		DrawString: func(p Vec2, s string, color HexColor) {
			strings = append(strings, s)
		},

		DrawBodyNames: true,
		DrawMass:      true,
		DrawShapes:    true,
	})

	assert.Equal(t, [][2]Vec2{{seg.Point1, seg.Point2}}, segments)
	assert.Contains(t, strings, "Foobar")
	assert.Contains(t, strings, "  0.00", "mass is printed via snprintf")
}

func TestComputeHull(t *testing.T) {
	points := []Vec2{{X: -1, Y: -1}, {X: 1, Y: -1}, {X: 0, Y: 0}, {X: 1, Y: 1}, {X: -1, Y: 1}}

	hull, ok := ComputeHull(points)
	require.True(t, ok)
	assert.EqualValues(t, 4, hull.Count, "interior point is removed")

	_, ok = ComputeHull(points[:2])
	assert.False(t, ok)

	_, ok = ComputeHull(make([]Vec2, _B2_MAX_POLYGON_VERTICES+1))
	assert.False(t, ok)
}

func TestMakeProxy(t *testing.T) {
	points := []Vec2{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 1, Y: 1}}

	proxy := MakeProxy(points, 0.5)
	assert.EqualValues(t, 3, proxy.Count)
	assert.Equal(t, points, proxy.Points[:proxy.Count])
	assert.Equal(t, float32(0.5), proxy.Radius)

	proxy = MakeOffsetProxy(points, 0, Vec2{X: 10}, IdentityRot)
	assert.Equal(t, []Vec2{{X: 10}, {X: 11}, {X: 11, Y: 1}}, proxy.Points[:proxy.Count])
}

func TestBody_GetShapesAndJoints(t *testing.T) {
	world := CreateWorld(DefaultWorldDef())
	defer world.DestroyWorld()

	bodyA, shapeA := createBox(world, Vec2{}, DefaultShapeDef())
	shapeB := bodyA.CreateCircleShape(DefaultShapeDef(), Circle{Radius: 1})
	bodyB, _ := createBox(world, Vec2{X: 2}, DefaultShapeDef())

	jointDef := DefaultWeldJointDef()
	jointDef.BodyIdA = bodyA.Id
	jointDef.BodyIdB = bodyB.Id
	joint := world.CreateWeldJoint(jointDef)

	assert.ElementsMatch(t, []Shape{shapeA, shapeB}, bodyA.GetShapes(nil))
	assert.Equal(t, []Joint{joint.Joint}, bodyA.GetJoints(nil))
	assert.Equal(t, []Joint{joint.Joint}, bodyB.GetJoints(nil))

	// reuse an existing slice
	reuse := make([]Shape, 0, 8)
	shapes := bodyA.GetShapes(reuse)
	assert.Len(t, shapes, 2)
	assert.Same(t, unsafe.SliceData(reuse), unsafe.SliceData(shapes))

	assert.Equal(t, JointTypeWeld, joint.GetType())
	assert.NotPanics(t, func() { joint.AsWeldJoint() })
	assert.Panics(t, func() { joint.AsMotorJoint() })

	joint.SetConstraintTuning(30, 2)
	hertz, damping := joint.GetConstraintTuning()
	assert.Equal(t, float32(30), hertz)
	assert.Equal(t, float32(2), damping)
}

func TestChain_GetSegments(t *testing.T) {
	world := CreateWorld(DefaultWorldDef())
	defer world.DestroyWorld()

	points := []Vec2{{X: -2, Y: 0}, {X: -1, Y: 0}, {X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}}
	material := DefaultSurfaceMaterial()

	chainDef := DefaultChainDef()
	chainDef.Points = uintptr(unsafe.Pointer(&points[0]))
	chainDef.Count = int32(len(points))
	chainDef.Materials = uintptr(unsafe.Pointer(&material))
	chainDef.MaterialCount = 1

	chain := world.CreateBody(DefaultBodyDef()).CreateChain(chainDef)

	// open chains drop the first and last ghost vertex
	segments := chain.GetSegments(nil)
	assert.Len(t, segments, len(points)-3)
	for _, segment := range segments {
		assert.Equal(t, ChainSegmentShape, segment.GetType())
	}
}

func TestWorld_OverlapAABB(t *testing.T) {
	world, ground := newTestWorld(t, DefaultWorldDef())
	_, boxA := createBox(world, Vec2{X: 0, Y: 5}, DefaultShapeDef())
	_, boxB := createBox(world, Vec2{X: 10, Y: 5}, DefaultShapeDef())

	var found []Shape
	world.OverlapAABB(AABB{LowerBound: Vec2{X: -1, Y: 4}, UpperBound: Vec2{X: 1, Y: 6}}, DefaultQueryFilter(), func(shape Shape) bool {
		found = append(found, shape)
		return true
	})

	assert.Equal(t, []Shape{boxA}, found)

	found = nil
	world.OverlapAABB(AABB{LowerBound: Vec2{X: -20, Y: -2}, UpperBound: Vec2{X: 20, Y: 6}}, DefaultQueryFilter(), func(shape Shape) bool {
		found = append(found, shape)
		return true
	})

	assert.ElementsMatch(t, []Shape{ground, boxA, boxB}, found)

	// returning false stops the query of the current tree, the static ground lives in another tree
	found = nil
	world.OverlapAABB(AABB{LowerBound: Vec2{X: -20, Y: -2}, UpperBound: Vec2{X: 20, Y: 6}}, DefaultQueryFilter(), func(shape Shape) bool {
		found = append(found, shape)
		return false
	})

	require.Len(t, found, 2)
	assert.Equal(t, ground, found[0])
	assert.Contains(t, []Shape{boxA, boxB}, found[1])
}

func TestWorld_OverlapShape(t *testing.T) {
	world, _ := newTestWorld(t, DefaultWorldDef())
	_, boxA := createBox(world, Vec2{X: 0, Y: 5}, DefaultShapeDef())
	createBox(world, Vec2{X: 10, Y: 5}, DefaultShapeDef())

	// a triangle around boxA that does not touch the ground or boxB
	proxy := MakeProxy([]Vec2{{X: -1, Y: 4}, {X: 1, Y: 4}, {X: 0, Y: 6}}, 0)

	var found []Shape
	world.OverlapShape(proxy, DefaultQueryFilter(), func(shape Shape) bool {
		found = append(found, shape)
		return true
	})

	assert.Equal(t, []Shape{boxA}, found)
}

func TestWorld_SetCustomFilterCallback(t *testing.T) {
	for _, concurrent := range []bool{false, true} {
		def := DefaultWorldDef()
		if concurrent {
			EnableConcurrency(&def)
		}

		world, ground := newTestWorld(t, def)
		body, box := createBox(world, Vec2{Y: 2}, DefaultShapeDef())

		var calls int
		world.SetCustomFilterCallback(func(shapeA, shapeB Shape) bool {
			calls++
			assert.ElementsMatch(t, []Shape{ground, box}, []Shape{shapeA, shapeB})
			return false
		})

		for range 120 {
			world.Step(1.0/60.0, 4)
		}

		assert.Positive(t, calls, "concurrent=%v", concurrent)
		assert.Less(t, body.GetPosition().Y, float32(-5), "box falls through the ground, concurrent=%v", concurrent)
	}
}

func TestWorld_SetPreSolveCallback(t *testing.T) {
	world, ground := newTestWorld(t, DefaultWorldDef())

	shapeDef := DefaultShapeDef()
	shapeDef.EnablePreSolveEvents = 1
	body, box := createBox(world, Vec2{Y: 2}, shapeDef)

	var calls int
	world.SetPreSolveCallback(func(shapeA, shapeB Shape, manifold Manifold) bool {
		calls++
		assert.ElementsMatch(t, []Shape{ground, box}, []Shape{shapeA, shapeB})
		assert.Positive(t, manifold.PointCount)

		// disable the contact
		return false
	})

	for range 120 {
		world.Step(1.0/60.0, 4)
	}

	assert.Positive(t, calls)
	assert.Less(t, body.GetPosition().Y, float32(-5), "box falls through the ground")
}

func TestWorld_MaterialCallbacks(t *testing.T) {
	world := CreateWorld(DefaultWorldDef())
	defer world.DestroyWorld()

	groundDef := DefaultShapeDef()
	groundDef.Material.UserMaterialId = 7
	groundDef.Material.Friction = 0.25
	groundDef.Material.Restitution = 0.5
	ground := world.CreateBody(DefaultBodyDef())
	ground.CreatePolygonShape(groundDef, MakeBox(20, 1))

	boxDef := DefaultShapeDef()
	boxDef.Material.UserMaterialId = 3
	boxDef.Material.Friction = 0.75
	createBox(world, Vec2{Y: 1.5}, boxDef)

	type call struct {
		ValueA, ValueB float32
		IdA, IdB       int32
	}

	var frictionCalls, restitutionCalls []call

	world.SetFrictionCallback(func(frictionA float32, idA int32, frictionB float32, idB int32) float32 {
		frictionCalls = append(frictionCalls, call{frictionA, frictionB, idA, idB})
		return 0.1
	})

	world.SetRestitutionCallback(func(restitutionA float32, idA int32, restitutionB float32, idB int32) float32 {
		restitutionCalls = append(restitutionCalls, call{restitutionA, restitutionB, idA, idB})
		return 0
	})

	world.Step(1.0/60.0, 4)

	require.NotEmpty(t, frictionCalls)
	require.NotEmpty(t, restitutionCalls)

	// the order of shapes in a contact is not defined
	assert.Contains(t, []call{{0.25, 0.75, 7, 3}, {0.75, 0.25, 3, 7}}, frictionCalls[0])
	assert.Contains(t, []call{{0.5, 0, 7, 3}, {0, 0.5, 3, 7}}, restitutionCalls[0])
}

func TestShape_GetSensorOverlaps(t *testing.T) {
	world, _ := newTestWorld(t, DefaultWorldDef())

	sensorDef := DefaultShapeDef()
	sensorDef.IsSensor = 1
	sensorDef.EnableSensorEvents = 1
	sensor := world.CreateBody(DefaultBodyDef()).CreatePolygonShape(sensorDef, MakeBox(5, 5))

	visitorDef := DefaultShapeDef()
	visitorDef.EnableSensorEvents = 1

	var visitors []Shape
	for idx := range 3 {
		_, shape := createBox(world, Vec2{X: float32(idx) - 1, Y: 2}, visitorDef)
		visitors = append(visitors, shape)
	}

	world.Step(1.0/60.0, 4)

	assert.ElementsMatch(t, visitors, sensor.GetSensorOverlaps(nil))

	events := world.GetSensorEvents()
	assert.Len(t, events.BeginEvents, len(visitors))
	for _, ev := range events.BeginEvents {
		assert.Equal(t, sensor.Id, ev.SensorShapeId)
	}
}

func TestGoAssert(t *testing.T) {
	// box2d asserts are turned into Go panics
	assert.PanicsWithError(t, "../box2d-c/src/geometry.c:140: b2IsValidFloat( halfWidth ) && halfWidth > 0.0f", func() {
		MakeBox(0, 1)
	})

	// and the shared stack is still usable afterwards
	theStack.CheckEmpty()
}
