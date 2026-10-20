package b2

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Ported from box2d test/test_shape.c and test/test_collision.c

const fltEpsilon = 1.1920929e-07

var (
	testCapsule = Capsule{Center1: Vec2{X: -1}, Center2: Vec2{X: 1}, Radius: 1}
	testCircle  = Circle{Center: Vec2{X: 1}, Radius: 1}
	testSegment = Segment{Point1: Vec2{Y: 1}, Point2: Vec2{Y: -1}}
)

func TestShapeMass(t *testing.T) {
	box := MakeBox(1, 1)

	t.Run("circle", func(t *testing.T) {
		md := ComputeCircleMass(testCircle, 1)
		assert.InDelta(t, _B2_PI, md.Mass, fltEpsilon)
		assert.Equal(t, Vec2{X: 1}, md.Center)
		// about the center of mass
		assert.InDelta(t, 0.5*_B2_PI, md.RotationalInertia, fltEpsilon)
	})

	t.Run("capsule", func(t *testing.T) {
		radius := testCapsule.Radius
		length := float32(2)

		md := ComputeCapsuleMass(testCapsule, 1)

		// box that fully contains the capsule
		mdr := ComputePolygonMass(MakeBox(radius+0.5*length, radius), 1)

		// approximate the capsule using a convex hull
		const n = 4
		var points []Vec2

		d := math.Pi / (n - 1.0)
		angle := -0.5 * math.Pi
		for range n {
			points = append(points, Vec2{X: 1 + radius*float32(math.Cos(angle)), Y: radius * float32(math.Sin(angle))})
			angle += d
		}

		angle = 0.5 * math.Pi
		for range n {
			points = append(points, Vec2{X: -1 + radius*float32(math.Cos(angle)), Y: radius * float32(math.Sin(angle))})
			angle += d
		}

		hull, ok := ComputeHull(points)
		assert.True(t, ok)
		ma := ComputePolygonMass(MakePolygon(hull, 0), 1)

		assert.Less(t, ma.Mass, md.Mass)
		assert.Less(t, md.Mass, mdr.Mass)
		assert.Less(t, ma.RotationalInertia, md.RotationalInertia)
		assert.Less(t, md.RotationalInertia, mdr.RotationalInertia)
	})

	t.Run("polygon", func(t *testing.T) {
		md := ComputePolygonMass(box, 1)
		assert.InDelta(t, 4, md.Mass, fltEpsilon)
		assert.InDelta(t, 0, md.Center.X, fltEpsilon)
		assert.InDelta(t, 0, md.Center.Y, fltEpsilon)
		assert.InDelta(t, float32(8.0/3.0), md.RotationalInertia, 2*fltEpsilon)
	})

	t.Run("offset polygon", func(t *testing.T) {
		offset := Vec2{X: 0.4, Y: -0.7}
		m1 := ComputePolygonMass(MakeBox(0.25, 0.5), 1)
		m2 := ComputePolygonMass(MakeOffsetBox(0.25, 0.5, offset, IdentityRot), 1)

		assert.InDelta(t, m1.Mass, m2.Mass, fltEpsilon)
		assert.InDelta(t, m1.RotationalInertia, m2.RotationalInertia, fltEpsilon)
		assert.InDelta(t, offset.X, m2.Center.X, fltEpsilon)
		assert.InDelta(t, offset.Y, m2.Center.Y, fltEpsilon)
	})
}

func TestShapeBounds(t *testing.T) {
	assertAABB := func(t *testing.T, expected, actual AABB) {
		assert.InDelta(t, expected.LowerBound.X, actual.LowerBound.X, fltEpsilon)
		assert.InDelta(t, expected.LowerBound.Y, actual.LowerBound.Y, fltEpsilon)
		assert.InDelta(t, expected.UpperBound.X, actual.UpperBound.X, fltEpsilon)
		assert.InDelta(t, expected.UpperBound.Y, actual.UpperBound.Y, fltEpsilon)
	}

	assertAABB(t, AABB{Vec2{0, -1}, Vec2{2, 1}}, ComputeCircleBounds(testCircle, IdentityTransform))
	assertAABB(t, AABB{Vec2{-1, -1}, Vec2{1, 1}}, ComputePolygonBounds(MakeBox(1, 1), IdentityTransform))
	assertAABB(t, AABB{Vec2{0, -1}, Vec2{0, 1}}, ComputeSegmentBounds(testSegment, IdentityTransform))
	assertAABB(t, AABB{Vec2{-2, -1}, Vec2{2, 1}}, ComputeCapsuleBounds(testCapsule, IdentityTransform))

	// rotated by 90 degrees and moved
	xf := Transform{P: Vec2{X: 5}, Q: Rot{C: 0, S: 1}}
	assertAABB(t, AABB{Vec2{4, 0}, Vec2{6, 2}}, ComputeCircleBounds(testCircle, xf))
}

func TestPointInShape(t *testing.T) {
	p1 := Vec2{X: 0.5, Y: 0.5}
	p2 := Vec2{X: 4, Y: -4}

	assert.EqualValues(t, 1, PointInCircle(testCircle, p1))
	assert.EqualValues(t, 0, PointInCircle(testCircle, p2))

	box := MakeBox(1, 1)
	assert.EqualValues(t, 1, PointInPolygon(box, p1))
	assert.EqualValues(t, 0, PointInPolygon(box, p2))
}

func TestCastRayShape(t *testing.T) {
	input := RayCastInput{Origin: Vec2{X: -4}, Translation: Vec2{X: 8}, MaxFraction: 1}

	assertHit := func(t *testing.T, output CastOutput, fraction float32) {
		assert.EqualValues(t, 1, output.Hit)
		assert.InDelta(t, -1, output.Normal.X, fltEpsilon)
		assert.InDelta(t, 0, output.Normal.Y, fltEpsilon)
		assert.InDelta(t, fraction, output.Fraction, fltEpsilon)
	}

	assertHit(t, CastRayCircle(testCircle, input), 0.5)
	assertHit(t, CastRayPolygon(MakeBox(1, 1), input), 3.0/8.0)
	assertHit(t, CastRaySegment(testSegment, input, 1), 0.5)

	miss := RayCastInput{Origin: Vec2{X: -4, Y: 5}, Translation: Vec2{X: 8}, MaxFraction: 1}
	assert.EqualValues(t, 0, CastRayCircle(testCircle, miss).Hit)
}

func TestAABB(t *testing.T) {
	a := AABB{LowerBound: Vec2{X: -1, Y: -1}, UpperBound: Vec2{X: -2, Y: -2}}
	assert.EqualValues(t, 0, IsValidAABB(a))

	a.UpperBound = Vec2{X: 1, Y: 1}
	assert.EqualValues(t, 1, IsValidAABB(a))

	// b2AABB_Overlaps and b2AABB_Contains are inline C functions and not transpiled
	output := b2AABB_RayCast(theStack, a, Vec2{X: -2}, Vec2{X: 2})
	assert.EqualValues(t, 1, output.Hit)
	assert.Greater(t, output.Fraction, float32(0.1))
	assert.Less(t, output.Fraction, float32(0.9))
}

func TestCastMover(t *testing.T) {
	world, _ := newTestWorld(t, DefaultWorldDef())

	// capsule standing 1 unit above the ground, moving down 2 units
	mover := Capsule{Center1: Vec2{Y: 1.5}, Center2: Vec2{Y: 2.5}, Radius: 0.5}
	fraction := world.CastMover(Pos{}, mover, Vec2{Y: -2}, DefaultQueryFilter())

	assert.InDelta(t, 0.5, fraction, 0.01)
}

// Ported from LargeWorldOriginQueryTest in box2d test/test_large_world.c
func TestCollideMover(t *testing.T) {
	for _, base := range []Pos{{}, {X: 1000, Y: -500}} {
		world := CreateWorld(DefaultWorldDef())

		bodyDef := DefaultBodyDef()
		bodyDef.Position = base
		box := world.CreateBody(bodyDef)
		box.CreatePolygonShape(DefaultShapeDef(), MakeBox(0.5, 0.5))

		// overlaps the left face of the box, relative to base
		mover := Capsule{Center1: Vec2{X: -0.9, Y: -0.2}, Center2: Vec2{X: -0.9, Y: 0.2}, Radius: 0.5}

		var planes []CollisionPlane
		world.CollideMover(base, mover, DefaultQueryFilter(), func(shape Shape, plane PlaneResult) bool {
			assert.Equal(t, box, shape.GetBody())

			if plane.Hit != 0 {
				planes = append(planes, CollisionPlane{Plane: plane.Plane, PushLimit: 1e30, ClipVelocity: 1})
			}

			return true
		})

		require.NotEmpty(t, planes)
		assert.InDelta(t, -1, planes[0].Plane.Normal.X, 1e-3)
		assert.InDelta(t, 0, planes[0].Plane.Normal.Y, 1e-3)

		// the solver pushes the mover out of the box to the left
		result := SolvePlanes(Vec2{}, planes)
		assert.Less(t, result.Delta.X, float32(-0.05))

		world.DestroyWorld()
	}
}

func TestCastRayClosest(t *testing.T) {
	world, ground := newTestWorld(t, DefaultWorldDef())

	result := world.CastRayClosest(Pos{Y: 10}, Vec2{Y: -20}, DefaultQueryFilter())
	assert.EqualValues(t, 1, result.Hit)
	assert.Equal(t, ground.Id, result.ShapeId)
	assert.InDelta(t, 0, result.Point.Y, fltEpsilon)
	assert.InDelta(t, 0.5, result.Fraction, fltEpsilon)
}
