package b2

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestWorld_DrawAllCallbacks triggers every debug draw callback, so a changed
// callback signature shows up as wrong values instead of going unnoticed.
func TestWorld_DrawAllCallbacks(t *testing.T) {
	world := CreateWorld(DefaultWorldDef())
	defer world.DestroyWorld()

	ground := world.CreateBody(DefaultBodyDef())
	ground.CreatePolygonShape(DefaultShapeDef(), MakeBox(5, 0.5))

	dynamicDef := DefaultBodyDef()
	dynamicDef.Type1 = DynamicBody

	dynamicDef.Position = Vec2{X: 0, Y: 1}
	ball := world.CreateBody(dynamicDef)
	ball.CreateCircleShape(DefaultShapeDef(), Circle{Radius: 0.5})

	dynamicDef.Position = Vec2{X: 3, Y: 0.75}
	capsule := world.CreateBody(dynamicDef)
	capsule.CreateCapsuleShape(DefaultShapeDef(), Capsule{Center1: Vec2{X: -0.5}, Center2: Vec2{X: 0.5}, Radius: 0.25})

	jointDef := DefaultRevoluteJointDef()
	jointDef.Base.BodyIdA = ground.Id
	jointDef.Base.BodyIdB = ball.Id
	jointDef.Base.LocalFrameA.P = Vec2{X: 0, Y: 1}
	world.CreateRevoluteJoint(jointDef)

	// create contacts
	world.Step(1.0/60.0, 4)

	type solidPolygon struct {
		transform Transform
		vertices  []Vec2
		radius    float32
	}

	var (
		polygons      [][]Vec2
		bounds        []AABB
		solidPolygons []solidPolygon
		circles       []Vec2
		solidCircles  []Vec2
		capsules      [][2]Vec2
		transforms    []Transform
		points        []Vec2
	)

	draw := DefaultDebugDraw()
	draw.DrawPolygon = func(transform Transform, vertices []Vec2, color HexColor) {
		polygons = append(polygons, vertices)
	}
	draw.DrawSolidPolygon = func(transform Transform, vertices []Vec2, radius float32, color HexColor) {
		solidPolygons = append(solidPolygons, solidPolygon{transform, vertices, radius})
	}
	draw.DrawCircle = func(center Vec2, radius float32, color HexColor) {
		circles = append(circles, center)
	}
	draw.DrawSolidCircle = func(transform Transform, center Vec2, radius float32, color HexColor) {
		assert.InDelta(t, 0.5, radius, 1e-6)
		assert.Equal(t, Vec2{}, center, "local center of the circle")
		solidCircles = append(solidCircles, transform.P)
	}
	draw.DrawSolidCapsule = func(p1 Vec2, p2 Vec2, radius float32, color HexColor) {
		assert.InDelta(t, 0.25, radius, 1e-6)
		capsules = append(capsules, [2]Vec2{p1, p2})
	}
	draw.DrawTransform = func(transform Transform) {
		transforms = append(transforms, transform)
	}
	draw.DrawPoint = func(p Vec2, size float32, color HexColor) {
		assert.Positive(t, size)
		points = append(points, p)
	}
	draw.DrawBounds = func(aabb AABB, color HexColor) {
		bounds = append(bounds, aabb)
	}

	draw.DrawJoints = true
	draw.DrawShapeBounds = true
	draw.DrawMass = true
	draw.DrawContacts = true

	world.Draw(draw)

	require.Len(t, solidPolygons, 1)
	assert.Len(t, solidPolygons[0].vertices, 4)
	assert.Zero(t, solidPolygons[0].radius)
	assert.Equal(t, ground.GetTransform(), solidPolygons[0].transform)

	require.Len(t, solidCircles, 1)
	assert.Equal(t, ball.GetPosition(), solidCircles[0])

	require.Len(t, capsules, 1)
	assert.InDelta(t, 2.5, capsules[0][0].X, 0.01)
	assert.InDelta(t, 3.5, capsules[0][1].X, 0.01)

	// the revolute joint draws a circle around its anchor
	require.NotEmpty(t, circles)
	assert.InDelta(t, 1, circles[0].Y, 0.01)

	// one bounding box per shape
	require.Len(t, bounds, 3)
	for _, aabb := range bounds {
		assert.True(t, aabb.LowerBound.X < aabb.UpperBound.X && aabb.LowerBound.Y < aabb.UpperBound.Y)
	}

	assert.Contains(t, transforms, ball.GetTransform())

	// contact points between ball, capsule and ground are at the top of the ground
	require.NotEmpty(t, points)
	for _, p := range points {
		assert.InDelta(t, 0.5, p.Y, 0.05)
	}
}

func TestJoints(t *testing.T) {
	tests := []struct {
		name      string
		jointType JointType
		create    func(world World, ground, body Body) Joint
		check     func(t *testing.T, joint Joint, body Body)
	}{
		{
			name:      "revolute",
			jointType: JointTypeRevolute,
			create: func(world World, ground, body Body) Joint {
				def := DefaultRevoluteJointDef()
				def.Base.BodyIdA, def.Base.BodyIdB = ground.Id, body.Id
				def.Base.LocalFrameB.P = Vec2{X: -1}
				return world.CreateRevoluteJoint(def).Joint
			},
			check: func(t *testing.T, joint Joint, body Body) {
				// a pendulum around the origin, swinging down
				p := body.GetPosition()
				assert.InDelta(t, 1, math.Hypot(float64(p.X), float64(p.Y)), 0.01)
				assert.Negative(t, joint.AsRevoluteJoint().GetCurrentAngle())
			},
		},
		{
			name:      "distance",
			jointType: JointTypeDistance,
			create: func(world World, ground, body Body) Joint {
				def := DefaultDistanceJointDef()
				def.Base.BodyIdA, def.Base.BodyIdB = ground.Id, body.Id
				def.Length = 1
				return world.CreateDistanceJoint(def).Joint
			},
			check: func(t *testing.T, joint Joint, body Body) {
				assert.InDelta(t, 1, joint.AsDistanceJoint().GetCurrentLength(), 0.01)
				assert.Less(t, body.GetPosition().Y, float32(-0.1))
			},
		},
		{
			name:      "prismatic",
			jointType: JointTypePrismatic,
			create: func(world World, ground, body Body) Joint {
				def := DefaultPrismaticJointDef()
				def.Base.BodyIdA, def.Base.BodyIdB = ground.Id, body.Id
				// the axis is the x-axis of frame A, rotate both frames to keep the body upright
				def.Base.LocalFrameA = Transform{P: Vec2{X: 1}, Q: Rot{C: 0, S: 1}}
				def.Base.LocalFrameB.Q = Rot{C: 0, S: 1}
				return world.CreatePrismaticJoint(def).Joint
			},
			check: func(t *testing.T, joint Joint, body Body) {
				// falls freely along the vertical axis, 0.5 * g * t^2
				assert.InDelta(t, -1.25, joint.AsPrismaticJoint().GetCurrentTranslation(), 0.1)
				assert.InDelta(t, 1, body.GetPosition().X, 0.01)
			},
		},
		{
			name:      "weld",
			jointType: JointTypeWeld,
			create: func(world World, ground, body Body) Joint {
				def := DefaultWeldJointDef()
				def.Base.BodyIdA, def.Base.BodyIdB = ground.Id, body.Id
				def.Base.LocalFrameA.P = Vec2{X: 1}
				return world.CreateWeldJoint(def).Joint
			},
			check: func(t *testing.T, joint Joint, body Body) {
				joint.AsWeldJoint()
				assert.InDelta(t, 1, body.GetPosition().X, 0.01)
				assert.InDelta(t, 0, body.GetPosition().Y, 0.01)
			},
		},
		{
			name:      "wheel",
			jointType: JointTypeWheel,
			create: func(world World, ground, body Body) Joint {
				def := DefaultWheelJointDef()
				def.Base.BodyIdA, def.Base.BodyIdB = ground.Id, body.Id
				def.Base.LocalFrameA = Transform{P: Vec2{X: 1}, Q: Rot{C: 0, S: 1}}
				def.Base.LocalFrameB.Q = Rot{C: 0, S: 1}
				return world.CreateWheelJoint(def).Joint
			},
			check: func(t *testing.T, joint Joint, body Body) {
				// hangs on the suspension spring
				assert.True(t, joint.AsWheelJoint().IsSpringEnabled() != 0)
				assert.InDelta(t, 1, body.GetPosition().X, 0.01)
				assert.Less(t, body.GetPosition().Y, float32(-0.05))
			},
		},
		{
			name:      "motor",
			jointType: JointTypeMotor,
			create: func(world World, ground, body Body) Joint {
				// a spring pulls frame B of the body to frame A
				def := DefaultMotorJointDef()
				def.Base.BodyIdA, def.Base.BodyIdB = ground.Id, body.Id
				def.Base.LocalFrameA.P = Vec2{X: 1}
				def.LinearHertz = 5
				def.LinearDampingRatio = 1
				def.MaxSpringForce = 1000
				def.AngularHertz = 5
				def.AngularDampingRatio = 1
				def.MaxSpringTorque = 1000
				return world.CreateMotorJoint(def).Joint
			},
			check: func(t *testing.T, joint Joint, body Body) {
				assert.Equal(t, float32(5), joint.AsMotorJoint().GetLinearHertz())
				assert.InDelta(t, 1, body.GetPosition().X, 0.05)
				assert.InDelta(t, 0, body.GetPosition().Y, 0.05)
			},
		},
		{
			name:      "mover",
			jointType: JointTypeMover,
			create: func(world World, ground, body Body) Joint {
				// drives the body with a constant velocity, holding it against gravity
				def := DefaultMoverJointDef()
				def.Base.BodyIdA, def.Base.BodyIdB = ground.Id, body.Id
				def.LinearVelocity = Vec2{X: 1}
				def.MaxVelocityForce = Vec2{X: 1000, Y: 1000}
				return world.CreateMoverJoint(def).Joint
			},
			check: func(t *testing.T, joint Joint, body Body) {
				assert.Equal(t, Vec2{X: 1}, joint.AsMoverJoint().GetLinearVelocity())
				assert.InDelta(t, 1.5, body.GetPosition().X, 0.05)
				assert.InDelta(t, 0, body.GetPosition().Y, 0.05)
			},
		},
		{
			name:      "filter",
			jointType: JointTypeFilter,
			create: func(world World, ground, body Body) Joint {
				// the body would rest on this box without the filter joint
				ground.CreatePolygonShape(DefaultShapeDef(), MakeBox(5, 0.5))

				def := DefaultFilterJointDef()
				def.Base.BodyIdA, def.Base.BodyIdB = ground.Id, body.Id
				return world.CreateFilterJoint(def).Joint
			},
			check: func(t *testing.T, joint Joint, body Body) {
				joint.AsFilterJoint()
				assert.Less(t, body.GetPosition().Y, float32(-1))
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			world := CreateWorld(DefaultWorldDef())
			defer world.DestroyWorld()

			ground := world.CreateBody(DefaultBodyDef())

			bodyDef := DefaultBodyDef()
			bodyDef.Type1 = DynamicBody
			bodyDef.Position = Vec2{X: 1}
			body := world.CreateBody(bodyDef)

			shapeDef := DefaultShapeDef()
			body.CreatePolygonShape(shapeDef, MakeBox(0.25, 0.25))

			joint := tc.create(world, ground, body)
			require.Equal(t, tc.jointType, joint.GetType())
			assert.Equal(t, ground, joint.GetBodyA())
			assert.Equal(t, body, joint.GetBodyB())

			for range 30 {
				world.Step(1.0/60.0, 4)
			}

			tc.check(t, joint, body)

			if tc.jointType != JointTypeWeld {
				assert.Panics(t, func() { joint.AsWeldJoint() }, "wrong joint type")
			}
		})
	}
}
