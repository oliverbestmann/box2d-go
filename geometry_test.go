package b2

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
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
		assert.InDelta(t, 1.5*_B2_PI, md.RotationalInertia, fltEpsilon)
	})

	t.Run("capsule", func(t *testing.T) {
		radius := testCapsule.Radius
		length := float32(2)

		md := ComputeCapsuleMass(testCapsule, 1)

		// box that fully contains the capsule
		mdr := ComputePolygonMass(MakeBox(radius, radius+0.5*length), 1)

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
}

func TestShapeAABB(t *testing.T) {
	assertAABB := func(t *testing.T, expected, actual AABB) {
		assert.InDelta(t, expected.LowerBound.X, actual.LowerBound.X, fltEpsilon)
		assert.InDelta(t, expected.LowerBound.Y, actual.LowerBound.Y, fltEpsilon)
		assert.InDelta(t, expected.UpperBound.X, actual.UpperBound.X, fltEpsilon)
		assert.InDelta(t, expected.UpperBound.Y, actual.UpperBound.Y, fltEpsilon)
	}

	assertAABB(t, AABB{Vec2{0, -1}, Vec2{2, 1}}, ComputeCircleAABB(testCircle, IdentityTransform))
	assertAABB(t, AABB{Vec2{-1, -1}, Vec2{1, 1}}, ComputePolygonAABB(MakeBox(1, 1), IdentityTransform))
	assertAABB(t, AABB{Vec2{0, -1}, Vec2{0, 1}}, ComputeSegmentAABB(testSegment, IdentityTransform))
	assertAABB(t, AABB{Vec2{-2, -1}, Vec2{2, 1}}, ComputeCapsuleAABB(testCapsule, IdentityTransform))

	// rotated by 90 degrees and moved
	xf := Transform{P: Vec2{X: 5}, Q: Rot{C: 0, S: 1}}
	assertAABB(t, AABB{Vec2{4, 0}, Vec2{6, 2}}, ComputeCircleAABB(testCircle, xf))
}

func TestPointInShape(t *testing.T) {
	p1 := Vec2{X: 0.5, Y: 0.5}
	p2 := Vec2{X: 4, Y: -4}

	assert.EqualValues(t, 1, PointInCircle(p1, testCircle))
	assert.EqualValues(t, 0, PointInCircle(p2, testCircle))

	box := MakeBox(1, 1)
	assert.EqualValues(t, 1, PointInPolygon(p1, box))
	assert.EqualValues(t, 0, PointInPolygon(p2, box))
}

func TestRayCastShape(t *testing.T) {
	input := RayCastInput{Origin: Vec2{X: -4}, Translation: Vec2{X: 8}, MaxFraction: 1}

	assertHit := func(t *testing.T, output CastOutput, fraction float32) {
		assert.EqualValues(t, 1, output.Hit)
		assert.InDelta(t, -1, output.Normal.X, fltEpsilon)
		assert.InDelta(t, 0, output.Normal.Y, fltEpsilon)
		assert.InDelta(t, fraction, output.Fraction, fltEpsilon)
	}

	assertHit(t, RayCastCircle(input, testCircle), 0.5)
	assertHit(t, RayCastPolygon(input, MakeBox(1, 1)), 3.0/8.0)
	assertHit(t, RayCastSegment(input, testSegment, 1), 0.5)

	miss := RayCastInput{Origin: Vec2{X: -4, Y: 5}, Translation: Vec2{X: 8}, MaxFraction: 1}
	assert.EqualValues(t, 0, RayCastCircle(miss, testCircle).Hit)
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
	fraction := world.CastMover(mover, Vec2{Y: -2}, DefaultQueryFilter())

	assert.InDelta(t, 0.5, fraction, 0.01)
}

func TestRayCastClosest(t *testing.T) {
	world, ground := newTestWorld(t, DefaultWorldDef())

	result := world.CastRayClosest(Vec2{Y: 10}, Vec2{Y: -20}, DefaultQueryFilter())
	assert.EqualValues(t, 1, result.Hit)
	assert.Equal(t, ground.Id, result.ShapeId)
	assert.InDelta(t, 0, result.Point.Y, fltEpsilon)
	assert.InDelta(t, 0.5, result.Fraction, fltEpsilon)
}
