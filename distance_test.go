package b2

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Ported from box2d test/test_distance.c

const epsilon = 1.1920929e-07 // FLT_EPSILON

var (
	distanceBox     = []Vec2{{X: -1, Y: -1}, {X: 1, Y: -1}, {X: 1, Y: 1}, {X: -1, Y: 1}}
	distanceSegment = []Vec2{{X: 2, Y: -1}, {X: 2, Y: 1}}
)

func TestComputeSegmentDistance(t *testing.T) {
	result := ComputeSegmentDistance(Vec2{X: -1, Y: -1}, Vec2{X: -1, Y: 1}, Vec2{X: 2, Y: 0}, Vec2{X: 1, Y: 0})

	assert.InDelta(t, 0.5, result.Fraction1, epsilon)
	assert.InDelta(t, 1, result.Fraction2, epsilon)
	assert.InDelta(t, -1, result.Closest1.X, epsilon)
	assert.InDelta(t, 0, result.Closest1.Y, epsilon)
	assert.InDelta(t, 1, result.Closest2.X, epsilon)
	assert.InDelta(t, 0, result.Closest2.Y, epsilon)
	assert.InDelta(t, 4, result.DistanceSquared, epsilon)
}

func TestComputeShapeDistance(t *testing.T) {
	input := DistanceInput{
		ProxyA:    MakeProxy(distanceBox, 0),
		ProxyB:    MakeProxy(distanceSegment, 0),
		Transform: IdentityTransform,
	}

	var cache SimplexCache
	output := ComputeShapeDistance(input, &cache, nil)

	assert.InDelta(t, 1, output.Distance, epsilon)
	assert.NotZero(t, cache.Count, "cache is written")

	// simplexes are written up to the length of the slice
	simplexes := make([]Simplex, 8)
	output = ComputeShapeDistance(input, nil, simplexes)
	assert.InDelta(t, 1, output.Distance, epsilon)
	require.Positive(t, output.SimplexCount)
	assert.NotZero(t, simplexes[0].Count)
}

// A thin plank and a point a linear slop off the middle of its short end face.
// GJK must not report a false overlap for the short edge.
func TestComputeShapeDistanceShortEdge(t *testing.T) {
	const halfThickness = 0.0016
	const linearSlop = 0.005

	corners := []Vec2{{X: -0.5, Y: -halfThickness}, {X: 0.5, Y: -halfThickness}, {X: 0.5, Y: halfThickness}, {X: -0.5, Y: halfThickness}}

	output := ComputeShapeDistance(DistanceInput{
		ProxyA:    MakeProxy(corners, 0),
		ProxyB:    MakeProxy([]Vec2{{X: 0.5 + linearSlop}}, 0),
		Transform: IdentityTransform,
	}, nil, nil)

	assert.Positive(t, output.Distance)
	assert.InDelta(t, 1, output.Normal.X*output.Normal.X+output.Normal.Y*output.Normal.Y, 1e-5)
	assert.InDelta(t, linearSlop, output.Distance, 1e-6)
}

func TestCastShapePair(t *testing.T) {
	output := CastShape(ShapeCastPairInput{
		ProxyA:       MakeProxy(distanceBox, 0),
		ProxyB:       MakeProxy(distanceSegment, 0),
		Transform:    IdentityTransform,
		TranslationB: Vec2{X: -2},
		MaxFraction:  1,
	})

	assert.EqualValues(t, 1, output.Hit)
	assert.InDelta(t, 0.5, output.Fraction, 0.005)
}

func TestComputeTimeOfImpact(t *testing.T) {
	output := ComputeTimeOfImpact(TOIInput{
		ProxyA:      MakeProxy(distanceBox, 0),
		ProxyB:      MakeProxy(distanceSegment, 0),
		SweepA:      Sweep{Q1: IdentityRot, Q2: IdentityRot},
		SweepB:      Sweep{C2: Vec2{X: -2}, Q1: IdentityRot, Q2: IdentityRot},
		MaxFraction: 1,
	})

	assert.Equal(t, ToiStateHit, output.State)
	assert.InDelta(t, 0.5, output.Fraction, 0.005)
}

func TestClipVector(t *testing.T) {
	ground := CollisionPlane{Plane: Plane{Normal: Vec2{Y: 1}}, Push: 1, ClipVelocity: 1}
	wall := CollisionPlane{Plane: Plane{Normal: Vec2{X: -1}}, Push: 1, ClipVelocity: 1}

	// the velocity into the ground is removed
	assert.Equal(t, Vec2{X: 1}, ClipVector(Vec2{X: 1, Y: -1}, []CollisionPlane{ground}))

	// all planes are used, not just the first one
	assert.Equal(t, Vec2{}, ClipVector(Vec2{X: 1, Y: -1}, []CollisionPlane{ground, wall}))

	// planes without push or with ClipVelocity disabled are skipped
	ground.Push = 0
	wall.ClipVelocity = 0
	assert.Equal(t, Vec2{X: 1, Y: -1}, ClipVector(Vec2{X: 1, Y: -1}, []CollisionPlane{ground, wall}))
}
