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
	jointDef.BodyIdA = ground.Id
	jointDef.BodyIdB = ball.Id
	jointDef.LocalAnchorA = Vec2{X: 0, Y: 1}
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
		solidPolygons []solidPolygon
		circles       []Vec2
		solidCircles  []float32
		capsules      [][2]Vec2
		transforms    []Transform
		points        []Vec2
	)

	world.Draw(DebugDraw{
		DrawPolygon: func(vertices []Vec2, color HexColor) {
			polygons = append(polygons, vertices)
		},
		DrawSolidPolygon: func(transform Transform, vertices []Vec2, radius float32, color HexColor) {
			solidPolygons = append(solidPolygons, solidPolygon{transform, vertices, radius})
		},
		DrawCircle: func(center Vec2, radius float32, color HexColor) {
			circles = append(circles, center)
		},
		DrawSolidCircle: func(transform Transform, radius float32, color HexColor) {
			solidCircles = append(solidCircles, radius)
		},
		DrawSolidCapsule: func(p1 Vec2, p2 Vec2, radius float32, color HexColor) {
			assert.InDelta(t, 0.25, radius, 1e-6)
			capsules = append(capsules, [2]Vec2{p1, p2})
		},
		DrawTransform: func(transform Transform) {
			transforms = append(transforms, transform)
		},
		DrawPoint: func(p Vec2, size float32, color HexColor) {
			assert.Positive(t, size)
			points = append(points, p)
		},

		DrawShapes:   true,
		DrawJoints:   true,
		DrawBounds:   true,
		DrawMass:     true,
		DrawContacts: true,
	})

	require.Len(t, solidPolygons, 1)
	assert.Len(t, solidPolygons[0].vertices, 4)
	assert.Zero(t, solidPolygons[0].radius)
	assert.Equal(t, ground.GetTransform(), solidPolygons[0].transform)

	assert.Equal(t, []float32{0.5}, solidCircles)

	require.Len(t, capsules, 1)
	assert.InDelta(t, 2.5, capsules[0][0].X, 0.01)
	assert.InDelta(t, 3.5, capsules[0][1].X, 0.01)

	// the revolute joint draws a circle around its anchor
	require.NotEmpty(t, circles)
	assert.InDelta(t, 1, circles[0].Y, 0.01)

	// one bounding box per shape
	require.Len(t, polygons, 3)
	for _, vertices := range polygons {
		assert.Len(t, vertices, 4)
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
				def.BodyIdA, def.BodyIdB = ground.Id, body.Id
				def.LocalAnchorB = Vec2{X: -1}
				return world.CreateRevoluteJoint(def).Joint
			},
			check: func(t *testing.T, joint Joint, body Body) {
				// a pendulum around the origin, swinging down
				p := body.GetPosition()
				assert.InDelta(t, 1, math.Hypot(float64(p.X), float64(p.Y)), 0.01)
				assert.Negative(t, joint.AsRevoluteJoint().RevoluteJoint_GetAngle())
			},
		},
		{
			name:      "distance",
			jointType: JointTypeDistance,
			create: func(world World, ground, body Body) Joint {
				def := DefaultDistanceJointDef()
				def.BodyIdA, def.BodyIdB = ground.Id, body.Id
				def.Length = 1
				return world.CreateDistanceJoint(def).Joint
			},
			check: func(t *testing.T, joint Joint, body Body) {
				assert.InDelta(t, 1, joint.AsDistanceJoint().DistanceJoint_GetCurrentLength(), 0.01)
				assert.Less(t, body.GetPosition().Y, float32(-0.1))
			},
		},
		{
			name:      "prismatic",
			jointType: JointTypePrismatic,
			create: func(world World, ground, body Body) Joint {
				def := DefaultPrismaticJointDef()
				def.BodyIdA, def.BodyIdB = ground.Id, body.Id
				def.LocalAnchorA = Vec2{X: 1}
				def.LocalAxisA = Vec2{Y: 1}
				return world.CreatePrismaticJoint(def).Joint
			},
			check: func(t *testing.T, joint Joint, body Body) {
				// falls freely along the vertical axis, 0.5 * g * t^2
				assert.InDelta(t, -1.25, joint.AsPrismaticJoint().PrismaticJoint_GetTranslation(), 0.1)
				assert.InDelta(t, 1, body.GetPosition().X, 0.01)
			},
		},
		{
			name:      "weld",
			jointType: JointTypeWeld,
			create: func(world World, ground, body Body) Joint {
				def := DefaultWeldJointDef()
				def.BodyIdA, def.BodyIdB = ground.Id, body.Id
				def.LocalAnchorA = Vec2{X: 1}
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
				def.BodyIdA, def.BodyIdB = ground.Id, body.Id
				def.LocalAnchorA = Vec2{X: 1}
				def.LocalAxisA = Vec2{Y: 1}
				return world.CreateWheelJoint(def).Joint
			},
			check: func(t *testing.T, joint Joint, body Body) {
				// hangs on the suspension spring
				assert.True(t, joint.AsWheelJoint().WheelJoint_IsSpringEnabled() != 0)
				assert.InDelta(t, 1, body.GetPosition().X, 0.01)
				assert.Less(t, body.GetPosition().Y, float32(-0.05))
			},
		},
		{
			name:      "motor",
			jointType: JointTypeMotor,
			create: func(world World, ground, body Body) Joint {
				def := DefaultMotorJointDef()
				def.BodyIdA, def.BodyIdB = ground.Id, body.Id
				def.LinearOffset = Vec2{X: 1}
				def.MaxForce = 1000
				def.MaxTorque = 1000
				return world.CreateMotorJoint(def).Joint
			},
			check: func(t *testing.T, joint Joint, body Body) {
				assert.Equal(t, Vec2{X: 1}, joint.AsMotorJoint().MotorJoint_GetLinearOffset())
				assert.InDelta(t, 1, body.GetPosition().X, 0.05)
				assert.InDelta(t, 0, body.GetPosition().Y, 0.05)
			},
		},
		{
			name:      "mouse",
			jointType: JointTypeMouse,
			create: func(world World, ground, body Body) Joint {
				def := DefaultMouseJointDef()
				def.BodyIdA, def.BodyIdB = ground.Id, body.Id
				// the joint grabs the body at the initial target, then drags it to the new one
				def.Target = body.GetPosition()
				def.Hertz = 5
				def.MaxForce = 1000
				joint := world.CreateMouseJoint(def)
				joint.MouseJoint_SetTarget(Vec2{X: 1, Y: 1})
				return joint.Joint
			},
			check: func(t *testing.T, joint Joint, body Body) {
				assert.Equal(t, Vec2{X: 1, Y: 1}, joint.AsMouseJoint().MouseJoint_GetTarget())
				assert.InDelta(t, 1, body.GetPosition().Y, 0.1)
			},
		},
		{
			name:      "filter",
			jointType: JointTypeFilter,
			create: func(world World, ground, body Body) Joint {
				// the body would rest on this box without the filter joint
				ground.CreatePolygonShape(DefaultShapeDef(), MakeBox(5, 0.5))

				def := DefaultFilterJointDef()
				def.BodyIdA, def.BodyIdB = ground.Id, body.Id
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
