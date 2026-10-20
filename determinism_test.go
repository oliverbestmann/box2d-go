package b2

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Ported from box2d test/test_determinism.c and shared/determinism.c.
// The expected values are the same as in the C test, so this checks that the
// transpiled solver produces bit-identical results.
const (
	expectedSleepStep = 299
	expectedHash      = 0x465f3df4
)

func TestDeterminismCrossPlatform(t *testing.T) {
	sleepStep, hash := runFallingHinges(t, DefaultWorldDef())
	assert.Equal(t, expectedSleepStep, sleepStep)
	assert.Equal(t, uint32(expectedHash), hash, "got %#x", hash)
}

func TestDeterminismMultithreaded(t *testing.T) {
	def := DefaultWorldDef()
	EnableConcurrency(&def)

	sleepStep, hash := runFallingHinges(t, def)
	assert.Equal(t, expectedSleepStep, sleepStep)
	assert.Equal(t, uint32(expectedHash), hash, "got %#x", hash)
}

func runFallingHinges(t *testing.T, def WorldDef) (sleepStep int, hash uint32) {
	world := CreateWorld(def)
	defer world.DestroyWorld()

	groundDef := DefaultBodyDef()
	groundDef.Position = Pos{X: 0, Y: -1}
	ground := world.CreateBody(groundDef)
	ground.CreatePolygonShape(DefaultShapeDef(), MakeBox(40, 1))

	const columnCount = 4
	const rowCount = 20

	h := float32(0.25)
	box := MakeSquare(h)

	shapeDef := DefaultShapeDef()

	jointDef := DefaultRevoluteJointDef()
	jointDef.EnableLimit = 1
	jointDef.LowerAngle = -0.1 * _B2_PI
	jointDef.UpperAngle = 0.2 * _B2_PI
	jointDef.EnableSpring = 1
	jointDef.Hertz = 1
	jointDef.DampingRatio = 1
	jointDef.EnableMotor = 1
	jointDef.MaxMotorTorque = 0.25
	jointDef.Base.LocalFrameA.P = Vec2{X: -h, Y: h}
	jointDef.Base.LocalFrameB.P = Vec2{X: -h, Y: -h}
	jointDef.Base.ConstraintHertz = 60
	jointDef.Base.ConstraintDampingRatio = 0
	jointDef.Base.DrawScale = 0.5

	offset := 0.4 * h
	dx := 10.0 * h
	xBase := -0.5 * dx * (columnCount - 1.0)

	var bodies []Body

	for j := range columnCount {
		// explicit float32 conversions prevent fused multiply-add, which C does not use
		x := xBase + float32(float32(j)*dx)

		var prevBody BodyId

		for i := range rowCount {
			bodyDef := DefaultBodyDef()
			bodyDef.Type1 = DynamicBody
			bodyDef.Position.X = x + float32(offset*float32(i))
			bodyDef.Position.Y = h + float32(2.0*h*float32(i))

			// this tests the deterministic cosine and sine functions
			angle := float32(-0.1)
			if i&1 == 1 {
				angle = 0.1
			}

			cs := ComputeCosSin(angle)
			bodyDef.Rotation = Rot{C: cs.Cosine, S: cs.Sine}

			body := world.CreateBody(bodyDef)

			if i&1 == 0 {
				prevBody = body.Id
			} else {
				jointDef.Base.BodyIdA = prevBody
				jointDef.Base.BodyIdB = body.Id
				world.CreateRevoluteJoint(jointDef)
				prevBody = BodyId{}
			}

			body.CreatePolygonShape(shapeDef, box)
			bodies = append(bodies, body)
		}
	}

	for step := 0; ; step++ {
		require.Less(t, step, 500, "world never fell asleep")

		world.Step(1.0/60.0, 4)

		if len(world.GetBodyEvents().MoveEvents) != 0 {
			continue
		}

		require.EqualValues(t, 0, world.GetAwakeBodyCount())

		// djb2 over all body transforms, like b2Hash
		hash = 5381
		for _, body := range bodies {
			xf := body.GetTransform()
			for _, b := range unsafe.Slice((*byte)(unsafe.Pointer(&xf)), unsafe.Sizeof(xf)) {
				hash = hash<<5 + hash + uint32(b)
			}
		}

		return step, hash
	}
}
