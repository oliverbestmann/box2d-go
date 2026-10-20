[![Go Reference](https://pkg.go.dev/badge/github.com/oliverbestmann/box2d-go/v3?utm_source=godoc.svg)](https://pkg.go.dev/github.com/oliverbestmann/box2d-go/v3?utm_source=godoc)

# box2d go

This is a Go port of Box2D v3.2. Most of it is generated: [ccgo](https://gitlab.com/cznic/ccgo) translates the C code
to Go, and a generator in `transpile/` adds a Go API on top that hides the translated C code.

```sh
go get github.com/oliverbestmann/box2d-go/v3
```

## Versions

The major and minor version follow Box2D: `v3.2.x` is a port of Box2D v3.2. The patch version counts releases of this
port and does not match the Box2D patch version. The table shows which Box2D version each release is built from.

| box2d-go | Box2D  |
|----------|--------|
| v3.2.0   | v3.2.0 |

## API

The API follows the C API. The `b2` prefix becomes the package name, so `b2CreateWorld` is `b2.CreateWorld`.

C functions that take an id as the first argument are methods on a wrapper type. For example, `b2.World` wraps
`b2.WorldId`, so `b2World_Draw(worldId, ...)` becomes `world.Draw(...)`. `Body`, `Shape`, `Joint`, `Chain`, and
`Contact` work the same way.

To call the methods of a specific joint type, convert the `Joint` first, for example with `AsMotorJoint()` or
`AsPrismaticJoint()`.

C pointers stay `uintptr` fields in the definitions. Box2D never reads the `UserData` fields, so any value works there.
For the other pointer fields, use the Go API instead:

- `ChainDef.Points` and `ChainDef.Materials`: pass slices to `Body.CreateChain`.
- `BodyDef.Name`: call `Body.SetName` after creating the body.
- `WorldDef.FrictionCallback` and `WorldDef.RestitutionCallback`: call `World.SetFrictionCallback` and
  `World.SetRestitutionCallback` after creating the world.
- `WorldDef.EnqueueTask` and `WorldDef.FinishTask`: not supported, the built-in scheduler runs the workers, see below.

## Example

This is the hello world example from the [Box2D documentation](https://box2d.org/documentation/hello.html), with the
original C code in comments.

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
//     b2World_Step(worldId, timeStep, subStepCount);
//     b2Vec2 position = b2Body_GetPosition(bodyId);
//     b2Rot rotation = b2Body_GetRotation(bodyId);
//     printf("%4.2f %4.2f %4.2f\n", position.x, position.y, b2Rot_GetAngle(rotation));
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

Box2D splits each step across `WorldDef.WorkerCount` threads. This port runs them as goroutines.
`b2.EnableConcurrency(&worldDef)` sets the worker count from the number of CPUs, up to 8. You can also set
`WorkerCount` yourself.

The rest of the API is not thread safe. Don't call into the `b2` package from more than one goroutine at a time, not
even for reads.

## SIMD

The contact solver can use SIMD through Go's experimental `simd/archsimd` package. This needs Go 1.27 and
`GOEXPERIMENT=simd`, and works on arm64, wasm, and amd64 with `GOAMD64=v3` (AVX2):

```sh
GOEXPERIMENT=simd go build .
GOEXPERIMENT=simd GOAMD64=v3 go build .
```

Other builds use the scalar version. Both give the same results. The SIMD build is about 2x faster on arm64
and 3x faster on wasm when many bodies touch. The wasm build then needs a runtime with wasm SIMD support, which all
current browsers have.

ccgo can't translate C vector types, so the port uses Box2D's scalar code and the generator moves the scalar wide
float math into `box2d_c-floatw.go`. `simd.go` and `simd_<arch>.go` implement the same functions with `archsimd`.

## Development

Run `go vet` with the `unsafeptr` check disabled:

```sh
go vet -unsafeptr=false ./...
```

The code from ccgo keeps C pointers as `uintptr` and converts them back to `unsafe.Pointer`. `unsafeptr` reports
each of these conversions, and only a change to ccgo could fix that.

For the same reason, run the race tests with checkptr disabled:

```sh
go test -race -gcflags=all=-d=checkptr=0 ./...
```

### Updating Box2D

- **Callbacks:** Box2D sees Go callbacks only as a `uintptr`, so a changed signature corrupts memory at runtime.
  `transpile/csrc/callbacks.c` repeats every callback signature the Go code implements, and `./transpile.sh`
  checks it against the Box2D headers with the host C compiler. `callbacks.go` checks the Go closures against
  the transpiled code. If the build fails there, update both files.
- **Determinism:** Copy the expected hash and sleep step from Box2D's `shared/determinism.h` into
  `determinism_test.go`.
- **SIMD:** If Box2D changes the scalar wide float math in `src/contact_solver_wide.inl`, update `simd.go`.
  `./transpile.sh` fails if a function listed in `floatWDecls` in `transpile/main.go` no longer exists.
