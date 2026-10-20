package docs

import (
	"fmt"
	"testing"

	b2 "github.com/oliverbestmann/box2d-go/v3"
)

func TestAPI(t *testing.T) {
	// b2WorldDef worldDef = b2DefaultWorldDef();
	// worldDef.gravity = (b2Vec2){0.0f, -10.0f};
	// b2WorldId worldId = b2CreateWorld(&worldDef);
	worldDef := b2.DefaultWorldDef()
	worldDef.Gravity = b2.Vec2{X: 0, Y: -10}
	world := b2.CreateWorld(worldDef)

	// b2BodyDef groundBodyDef = b2DefaultBodyDef();
	// groundBodyDef.position = (b2Vec2){0.0f, -10.0f};
	// b2BodyId groundId = b2CreateBody(worldId, &groundBodyDef);
	groundBodyDef := b2.DefaultBodyDef()
	groundBodyDef.Position = b2.Vec2{X: 0, Y: -10}
	ground := world.CreateBody(groundBodyDef)

	// b2Polygon groundBox = b2MakeBox(50.0f, 10.0f);
	groundBox := b2.MakeBox(50, 10)

	// b2ShapeDef groundShapeDef = b2DefaultShapeDef();
	// b2CreatePolygonShape(groundId, &groundShapeDef, &groundBox);
	groundShapeDef := b2.DefaultShapeDef()
	ground.CreatePolygonShape(groundShapeDef, groundBox)

	// b2BodyDef bodyDef = b2DefaultBodyDef();
	// bodyDef.type = b2_dynamicBody;
	// bodyDef.position = (b2Vec2){0.0f, 4.0f};
	// b2BodyId bodyId = b2CreateBody(worldId, &bodyDef);
	bodyDef := b2.DefaultBodyDef()
	bodyDef.Type1 = b2.DynamicBody
	bodyDef.Position = b2.Vec2{Y: 4}
	body := world.CreateBody(bodyDef)

	// b2Polygon dynamicBox = b2MakeBox(1.0f, 1.0f);
	dynamicBox := b2.MakeBox(1, 1)

	// b2ShapeDef shapeDef = b2DefaultShapeDef();
	// shapeDef.density = 1.0f;
	// shapeDef.material.friction = 0.3f;
	shapeDef := b2.DefaultShapeDef()
	shapeDef.Density = 1
	shapeDef.Material.Friction = 0.3

	// b2CreatePolygonShape(bodyId, &shapeDef, &dynamicBox);
	body.CreatePolygonShape(shapeDef, dynamicBox)

	// float timeStep = 1.0f / 60.0f;
	// int subStepCount = 4;
	timeStep := float32(1 / 60.0)
	subStepCount := int32(4)

	// for (int i = 0; i < 90; ++i) {
	// 	 2World_Step(worldId, timeStep, subStepCount);
	// 	 2Vec2 position = b2Body_GetPosition(bodyId);
	// 	 2Rot rotation = b2Body_GetRotation(bodyId);
	// 	 rintf("%4.2f %4.2f %4.2f\n", position.x, position.y, b2Rot_GetAngle(rotation));
	// }
	for range 90 {
		world.Step(timeStep, subStepCount)
		position := body.GetPosition()
		rotation := body.GetRotation()
		fmt.Printf("%4.2f %4.2f %4.2f\n", position.X, position.Y, rotation.Angle())
	}

	// b2DestroyWorld(worldId);
	world.DestroyWorld()
}
