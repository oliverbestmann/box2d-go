package b2

import (
	"strings"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// stackedBoxes creates a static ground at y=0 and a box resting on it at y=1.
func stackedBoxes(t *testing.T) (World, Shape, Body, Shape) {
	world := CreateWorld(DefaultWorldDef())
	t.Cleanup(world.DestroyWorld)

	ground := world.CreateBody(DefaultBodyDef())
	groundShape := ground.CreatePolygonShape(DefaultShapeDef(), MakeBox(10, 0.5))

	bodyDef := DefaultBodyDef()
	bodyDef.Type1 = DynamicBody
	bodyDef.Position = Pos{Y: 1}
	body := world.CreateBody(bodyDef)
	bodyShape := body.CreatePolygonShape(DefaultShapeDef(), MakeBox(0.5, 0.5))

	return world, groundShape, body, bodyShape
}

func TestBodyName(t *testing.T) {
	_, _, body, _ := stackedBoxes(t)

	assert.Equal(t, "", body.GetName())

	body.SetName("crate")
	assert.Equal(t, "crate", body.GetName())

	body.SetName(strings.Repeat("x", 40))
	assert.Equal(t, strings.Repeat("x", 10), body.GetName())
}

func TestCastRay(t *testing.T) {
	world, groundShape, _, bodyShape := stackedBoxes(t)

	var hits []Shape
	world.CastRay(Pos{Y: 10}, Vec2{Y: -20}, DefaultQueryFilter(), func(shape Shape, point Pos, normal Vec2, fraction float32) float32 {
		hits = append(hits, shape)
		assert.InDelta(t, 1, normal.Y, 1e-5)
		return 1
	})
	assert.ElementsMatch(t, []Shape{groundShape, bodyShape}, hits)

	hits = nil
	world.CastRay(Pos{Y: 10}, Vec2{Y: -20}, DefaultQueryFilter(), func(shape Shape, point Pos, normal Vec2, fraction float32) float32 {
		hits = append(hits, shape)
		return 0
	})
	assert.Len(t, hits, 1, "returning 0 stops the cast")
}

func TestCastShape(t *testing.T) {
	world, _, _, bodyShape := stackedBoxes(t)

	proxy := MakeProxy([]Vec2{{X: 0, Y: 10}}, 0.25)

	var hits []Shape
	world.CastShape(Pos{}, proxy, Vec2{Y: -9}, DefaultQueryFilter(), func(shape Shape, point Pos, normal Vec2, fraction float32) float32 {
		hits = append(hits, shape)
		return 1
	})

	// the circle ends at y=1 with radius 0.25, inside the box (top 1.5) but above the ground (top 0.5)
	assert.Equal(t, []Shape{bodyShape}, hits)
}

func TestGetContactData(t *testing.T) {
	world, groundShape, body, bodyShape := stackedBoxes(t)

	for range 30 {
		world.Step(1.0/60.0, 4)
	}

	contacts := body.GetContactData(nil)
	require.Len(t, contacts, 1)
	assert.ElementsMatch(t,
		[]Shape{groundShape, bodyShape},
		[]Shape{{Id: contacts[0].ShapeIdA}, {Id: contacts[0].ShapeIdB}})
	assert.EqualValues(t, 2, contacts[0].Manifold.PointCount)

	reuse := make([]ContactData, 0, 8)
	contacts = bodyShape.GetContactData(reuse)
	require.Len(t, contacts, 1)
	assert.Same(t, &reuse[:1][0], &contacts[0], "reuse is used when large enough")
}

func TestSolvePlanes(t *testing.T) {
	planes := []CollisionPlane{{
		Plane:     Plane{Normal: Vec2{Y: 1}},
		PushLimit: 100,
	}}

	// moving into the ground is clipped to the linear slop
	result := SolvePlanes(Vec2{X: 1, Y: -1}, planes)
	assert.InDelta(t, 1, result.Delta.X, 1e-5)
	assert.InDelta(t, -0.005, result.Delta.Y, 1e-5)
	assert.InDelta(t, 0.995, planes[0].Push, 1e-5)

	result = SolvePlanes(Vec2{X: 1, Y: 2}, nil)
	assert.Equal(t, Vec2{X: 1, Y: 2}, result.Delta)
}

func TestAtomicCompareExchange(t *testing.T) {
	value, expected := int32(5), int32(3)

	ok := __builtin___atomic_compare_exchange_n(nil, uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Pointer(&expected)), 7, 0, 0, 0)
	assert.EqualValues(t, false1, ok)
	assert.EqualValues(t, 5, value)
	assert.EqualValues(t, 5, expected, "expected receives the current value on failure")

	ok = __builtin___atomic_compare_exchange_n(nil, uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Pointer(&expected)), 7, 0, 0, 0)
	assert.EqualValues(t, true1, ok)
	assert.EqualValues(t, 7, value)
}
