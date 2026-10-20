[![Go Reference](https://pkg.go.dev/badge/github.com/oliverbestmann/box2d-go?utm_source=godoc.svg)](https://pkg.go.dev/github.com/oliverbestmann/box2d-go?utm_source=godoc)

# box2d go

This is a mainly automated Golang port of Box2D v3.2. The original C code is transpiled into Go using modernc.org/ccgo/v4.
Some post-processing creates a user-friendly Go API, concealing the low-level details of the transpiled C code.

## API

The API largely mirrors the original C API. The `b2` prefix in the C API is incorporated into the package name, allowing
`b2CreateWorld` to be accessed as `b2.CreateWorld`.

Functions akin to classes or namespaces that operate on instances, such as `b2World_Draw(worldId, ...)`, are contained
within a `b2.World` instance, which is a wrapper for `b2.WorldId`. The aforementioned method is accessible via
`myWorld.Draw(...)`. This format is also applied to `Body`, `Shape`, and `Joint` methods.

`Joint` instances provide methods to convert them into specific joint types, like `AsMotorJoint()` or
`AsPrismaticJoint()`, enabling access to methods specific to those joint types.

## Example

This example is directly taken from the official [Box2D documentation](https://box2d.org/documentation/hello.html)
and adapted to Go. The original C code is provided for comparison.

```go
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
```

## Multithreading

box2d has a built-in task scheduler that runs on `WorldDef.WorkerCount` threads. In Go, these threads are
goroutines. Use `b2.EnableConcurrency` to pick a worker count for a new `World`, or set `WorkerCount` yourself.

**Caution**: The rest of the API is not thread safe. The `b2` package must not be
accessed from multiple threads, even in read-only mode.

## Use of AI

AI was used solely to refine the phrasing in this README. No AI-generated content is present in the code.

## Development

Run `go vet` with the `unsafeptr` check disabled:

```sh
go vet -unsafeptr=false ./...
```

The code transpiled by ccgo stores all C memory as `uintptr` and converts it back to `unsafe.Pointer`,
which `unsafeptr` reports everywhere. This cannot be fixed without changing how ccgo generates code.

The tests use `-race` only with checkptr disabled for the same reason:
`go test -race -gcflags=all=-d=checkptr=0 ./...`.

### Updating box2d

- Callbacks: box2d only sees a `uintptr` for Go callbacks, so a changed signature would corrupt memory at runtime.
  `transpile/csrc/callbacks.c` repeats every callback signature the Go code implements, and `./transpile.sh`
  checks it against the box2d headers with the host C compiler. `callbacks.go` checks the Go closures against
  the transpiled result. Update both when the build fails there.
- Determinism: take the expected hash and sleep step from box2d's `shared/determinism.h` and update
  `determinism_test.go`.
