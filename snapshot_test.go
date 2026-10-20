package b2

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Ported from SnapshotTest in box2d test/test_snapshot.c, using body positions and snapshot
// images instead of the internal world hash.

func buildSnapshotScene(workerCount int32) (World, []Body) {
	def := DefaultWorldDef()
	def.WorkerCount = workerCount
	world := CreateWorld(def)

	groundDef := DefaultBodyDef()
	groundDef.Position = Pos{Y: -1}
	world.CreateBody(groundDef).CreatePolygonShape(DefaultShapeDef(), MakeBox(40, 1))

	var bodies []Body
	for i := range 8 {
		bodyDef := DefaultBodyDef()
		bodyDef.Type1 = DynamicBody
		bodyDef.Position = Pos{Y: 0.5 + float32(i)*1.1}
		body := world.CreateBody(bodyDef)
		body.CreatePolygonShape(DefaultShapeDef(), MakeBox(0.5, 0.5))
		bodies = append(bodies, body)
	}

	return world, bodies
}

func stepUntilSleep(t *testing.T, world World) {
	for range 500 {
		world.Step(1.0/60.0, 4)
		if world.GetAwakeBodyCount() == 0 {
			return
		}
	}

	t.Fatal("world did not fall asleep")
}

func positionsOf(bodies []Body) []Pos {
	var positions []Pos
	for _, body := range bodies {
		positions = append(positions, body.GetPosition())
	}

	return positions
}

func TestSnapshotRestore(t *testing.T) {
	world, bodies := buildSnapshotScene(0)
	defer world.DestroyWorld()

	stepUntilSleep(t, world)

	image := world.GetSnapshot(nil)
	require.NotEmpty(t, image)
	positions := positionsOf(bodies)

	// diverge from the snapshot
	top := bodies[len(bodies)-1]
	top.SetLinearVelocity(Vec2{X: 3, Y: 6})
	top.SetAwake(1)

	postDef := DefaultBodyDef()
	postDef.Type1 = DynamicBody
	postDef.Position = Pos{X: 20, Y: 20}
	postBody := world.CreateBody(postDef)

	for range 30 {
		world.Step(1.0/60.0, 4)
	}

	require.NotEqual(t, positions, positionsOf(bodies))

	require.True(t, world.RestoreFromSnapshot(image))
	assert.Equal(t, positions, positionsOf(bodies))

	// ids from before the snapshot stay valid, newer ones are rejected
	assert.EqualValues(t, 1, top.IsValid())
	assert.EqualValues(t, 0, postBody.IsValid())

	// a rejected image leaves the world untouched
	assert.False(t, world.RestoreFromSnapshot(nil))

	corrupt := slices.Clone(image)
	corrupt[0] ^= 0xff
	assert.False(t, world.RestoreFromSnapshot(corrupt))
	assert.Equal(t, positions, positionsOf(bodies))
}

func TestCreateWorldFromSnapshot(t *testing.T) {
	world, bodies := buildSnapshotScene(0)
	defer world.DestroyWorld()

	stepUntilSleep(t, world)

	// wake the stack so the worlds have something to simulate
	top := bodies[len(bodies)-1]
	top.SetLinearVelocity(Vec2{X: 3, Y: 6})
	top.SetAwake(1)

	image := world.GetSnapshot(nil)
	positions := positionsOf(bodies)

	// the result must not depend on the worker count
	for _, workerCount := range []int32{0, 4} {
		copied, ok := CreateWorldFromSnapshot(image, workerCount)
		require.True(t, ok)

		// the same bodies in the copied world
		var copiedBodies []Body
		for _, body := range bodies {
			id := body.Id
			id.World0 = copied.Id.Index1 - 1
			copiedBodies = append(copiedBodies, Body{Id: id})
		}

		for range 120 {
			world.Step(1.0/60.0, 4)
			copied.Step(1.0/60.0, 4)
		}

		require.NotEqual(t, positions, positionsOf(bodies))
		assert.Equal(t, positionsOf(bodies), positionsOf(copiedBodies))

		copied.DestroyWorld()
		require.True(t, world.RestoreFromSnapshot(image))
	}

	_, ok := CreateWorldFromSnapshot([]byte{1, 2, 3}, 0)
	assert.False(t, ok)
}
