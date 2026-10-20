package b2

import (
	"fmt"
	"log/slog"
	"runtime"
	"slices"
	"strings"
	"unsafe"
)

type (
	RestitutionCallback = func(restitutionA float32, userMaterialIdA uint64, restitutionB float32, userMaterialIdB uint64) float32
	FrictionCallback    = func(frictionA float32, userMaterialIdA uint64, frictionB float32, userMaterialIdB uint64) float32
)

var (
	ZeroVec2          = Vec2{}
	IdentityRot       = Rot{C: 1.0, S: 0.0}
	IdentityTransform = Transform{Q: IdentityRot}
	ZeroMat           = Mat22{}
)

func init() {
	b2SetAssertCallback(theStack, __ccgo_fp(assertCallback(b2DefaultAssertFcnGo)))
	b2SetLogCallback(theStack, __ccgo_fp(logCallback(b2LogGo)))
}

// b2LogGo logs messages of box2d, for example about invalid input.
func b2LogGo(_ *_Stack, message uintptr) {
	slog.Warn("box2d", slog.String("message", strings.TrimSpace(toString(message))))
}

func ComputeHull(points []Vec2) (Hull, bool) {
	if len(points) > _B2_MAX_POLYGON_VERTICES {
		return Hull{}, false
	}

	var arr [_B2_MAX_POLYGON_VERTICES]Vec2
	copy(arr[:], points)

	pArr := copyToStack(theStack, arr)
	defer pArr.Free()

	hull := b2ComputeHull(theStack, pArr.Addr, int32(len(points)))
	return hull, hull.Count >= 3
}

func (b Body) SetName(name string) {
	nameC := theStack.toCString(name)
	defer nameC.Free()

	b2Body_SetName(theStack, b.Id, nameC.Addr)
}

// GetName returns the name set with SetName, truncated to 10 bytes.
func (b Body) GetName() string {
	return toString(b2Body_GetName(theStack, b.Id))
}

// The wrapper types have the same memory layout as their ids, so box2d can write the ids
// directly into the slices.

func (b Body) GetShapes(reuse []Shape) []Shape {
	return fillSlice(reuse, b.GetShapeCount(), func(ptr uintptr, capacity int32) int32 {
		return b2Body_GetShapes(theStack, b.Id, ptr, capacity)
	})
}

func (b Body) GetJoints(reuse []Joint) []Joint {
	return fillSlice(reuse, b.GetJointCount(), func(ptr uintptr, capacity int32) int32 {
		return b2Body_GetJoints(theStack, b.Id, ptr, capacity)
	})
}

func (b Chain) GetSegments(reuse []Shape) []Shape {
	return fillSlice(reuse, b.GetSegmentCount(), func(ptr uintptr, capacity int32) int32 {
		return b2Chain_GetSegments(theStack, b.Id, ptr, capacity)
	})
}

// GetSensorData returns the shapes that overlap the sensor shape.
func (b Shape) GetSensorData(reuse []Shape) []Shape {
	return fillSlice(reuse, b.GetSensorCapacity(), func(ptr uintptr, capacity int32) int32 {
		return b2Shape_GetSensorData(theStack, b.Id, ptr, capacity)
	})
}

// GetContactData returns the touching contacts of the body, reusing the memory of reuse if it is large enough.
func (b Body) GetContactData(reuse []ContactData) []ContactData {
	return fillSlice(reuse, b.GetContactCapacity(), func(ptr uintptr, capacity int32) int32 {
		return b2Body_GetContactData(theStack, b.Id, ptr, capacity)
	})
}

// GetContactData returns the touching contacts of the shape, reusing the memory of reuse if it is large enough.
func (b Shape) GetContactData(reuse []ContactData) []ContactData {
	return fillSlice(reuse, b.GetContactCapacity(), func(ptr uintptr, capacity int32) int32 {
		return b2Shape_GetContactData(theStack, b.Id, ptr, capacity)
	})
}

// fillSlice lets fill write up to capacity items into reuse and returns the items written.
func fillSlice[T any](reuse []T, capacity int32, fill func(ptr uintptr, capacity int32) int32) []T {
	if capacity <= 0 {
		return reuse[:0]
	}

	if cap(reuse) < int(capacity) {
		reuse = make([]T, capacity)
	}

	items := reuse[:capacity]
	count := fill(sliceAddr(items), capacity)
	runtime.KeepAlive(items)

	return items[:count]
}

// SolvePlanes solves the position of a mover that satisfies the given collision planes.
// It updates the Push field of the planes.
func SolvePlanes(targetDelta Vec2, planes []CollisionPlane) PlaneSolverResult {
	defer runtime.KeepAlive(planes)
	return b2SolvePlanes(theStack, targetDelta, sliceAddr(planes), int32(len(planes)))
}

// ComputeShapeDistance computes the closest points between two shapes. cache speeds up repeated
// calls for the same shapes and may be nil. For debugging, the solver writes its simplexes
// into simplexes, up to its length. DistanceOutput.SimplexCount tells how many it computed.
func ComputeShapeDistance(input DistanceInput, cache *SimplexCache, simplexes []Simplex) DistanceOutput {
	if cache == nil {
		cache = &SimplexCache{}
	}

	inputStack := copyToStack(theStack, input)
	defer inputStack.Free()

	// box2d writes into cache via a uintptr, it must live on the heap
	escapes(cache)
	defer runtime.KeepAlive(cache)
	defer runtime.KeepAlive(simplexes)

	return b2ComputeShapeDistance(theStack, inputStack.Addr, uintptr(unsafe.Pointer(cache)), sliceAddr(simplexes), int32(len(simplexes)))
}

// worldCallbacks keeps the Go closures alive that box2d only stores as a uintptr.
// Keyed by world slot, so there are at most B2_MAX_WORLDS entries per kind.
var worldCallbacks = map[worldCallbackKey]any{}

type worldCallbackKey struct {
	world uint16
	kind  string
}

// keepCallback stores fn for the lifetime of the world and returns the function pointer for box2d.
// fn must have the signature that the transpiled code uses, including the leading *_Stack.
func (b World) keepCallback(kind string, fn any) uintptr {
	worldCallbacks[worldCallbackKey{world: b.Id.Index1, kind: kind}] = fn
	return __ccgo_fp(fn)
}

// SetRestitutionCallback sets the worlds RestitutionCallback. Use nil to restore the default.
// The callback may be called from multiple goroutines when concurrency is enabled.
func (b World) SetRestitutionCallback(callback RestitutionCallback) {
	var fn uintptr
	if callback != nil {
		fn = b.keepCallback("restitution", restitutionCallback(func(_ *_Stack, restitutionA float32, userMaterialIdA uint64, restitutionB float32, userMaterialIdB uint64) float32 {
			return callback(restitutionA, userMaterialIdA, restitutionB, userMaterialIdB)
		}))
	}

	b2World_SetRestitutionCallback(theStack, b.Id, fn)
}

// SetFrictionCallback sets the worlds FrictionCallback. Use nil to restore the default.
// The callback may be called from multiple goroutines when concurrency is enabled.
func (b World) SetFrictionCallback(callback FrictionCallback) {
	var fn uintptr
	if callback != nil {
		fn = b.keepCallback("friction", frictionCallback(func(_ *_Stack, frictionA float32, userMaterialIdA uint64, frictionB float32, userMaterialIdB uint64) float32 {
			return callback(frictionA, userMaterialIdA, frictionB, userMaterialIdB)
		}))
	}

	b2World_SetFrictionCallback(theStack, b.Id, fn)
}

type OverlapResult func(shapeId Shape) bool

// OverlapAABB calls fcn for every shape that overlaps the aabb. The aabb is relative to origin.
// Return false to stop the query.
func (b World) OverlapAABB(origin Pos, aabb AABB, filter QueryFilter, fcn OverlapResult) TreeStats {
	fn := worldOverlapCallback(fcn)
	defer runtime.KeepAlive(fn)

	return b2World_OverlapAABB(theStack, b.Id, origin, aabb, filter, __ccgo_fp(fn), 0)
}

// OverlapShape calls fcn for every shape that overlaps the proxy. The proxy is relative to origin.
// Return false to stop the query.
func (b World) OverlapShape(origin Pos, proxy ShapeProxy, filter QueryFilter, fcn OverlapResult) (r TreeStats) {
	fn := worldOverlapCallback(fcn)
	defer runtime.KeepAlive(fn)

	proxyStack := copyToStack(theStack, proxy)
	defer proxyStack.Free()

	r = b2World_OverlapShape(theStack, b.Id, origin, proxyStack.Addr, filter, __ccgo_fp(fn), 0)

	return
}

func worldOverlapCallback(fcn OverlapResult) overlapResultCallback {
	fn := func(_ *_Stack, shapeId ShapeId, _ uintptr) uint8 {
		return boolUint8(fcn(Shape{Id: shapeId}))
	}

	// box2d only sees a uintptr, the closure must not live on the goroutine stack, which can move
	escapes(fn)

	return fn
}

// CastResult is called for every shape hit by a cast. Return -1 to ignore the shape,
// 0 to stop, fraction to clip the cast at this hit, or 1 to continue without clipping.
type CastResult func(shape Shape, point Pos, normal Vec2, fraction float32) float32

// CastRay calls fcn for every shape hit by the ray, in no particular order.
func (b World) CastRay(origin Pos, translation Vec2, filter QueryFilter, fcn CastResult) TreeStats {
	fn := worldCastCallback(fcn)
	defer runtime.KeepAlive(fn)

	return b2World_CastRay(theStack, b.Id, origin, translation, filter, __ccgo_fp(fn), 0)
}

// CastShape calls fcn for every shape hit by the moving proxy, in no particular order.
// The proxy is relative to origin.
func (b World) CastShape(origin Pos, proxy ShapeProxy, translation Vec2, filter QueryFilter, fcn CastResult) TreeStats {
	fn := worldCastCallback(fcn)
	defer runtime.KeepAlive(fn)

	proxyStack := copyToStack(theStack, proxy)
	defer proxyStack.Free()

	return b2World_CastShape(theStack, b.Id, origin, proxyStack.Addr, translation, filter, __ccgo_fp(fn), 0)
}

func worldCastCallback(fcn CastResult) castResultCallback {
	fn := func(_ *_Stack, shapeId ShapeId, point Pos, normal Vec2, fraction float32, _ uintptr) float32 {
		return fcn(Shape{Id: shapeId}, point, normal, fraction)
	}

	// box2d only sees a uintptr, the closure must not live on the goroutine stack, which can move
	escapes(fn)

	return fn
}

// CustomFilter returns false to prevent the two shapes from colliding. It is only called
// if one of the shapes has ShapeDef.EnableCustomFiltering set.
type CustomFilter func(shapeIdA, shapeIdB Shape) bool

// SetCustomFilterCallback sets the worlds CustomFilter. Use nil to remove it.
// The callback may be called from multiple goroutines when concurrency is enabled.
func (b World) SetCustomFilterCallback(fcn CustomFilter) {
	var fn uintptr
	if fcn != nil {
		fn = b.keepCallback("customFilter", customFilterCallback(func(_ *_Stack, shapeIdA, shapeIdB ShapeId, _ uintptr) uint8 {
			return boolUint8(fcn(Shape{Id: shapeIdA}, Shape{Id: shapeIdB}))
		}))
	}

	b2World_SetCustomFilterCallback(theStack, b.Id, fn, 0)
}

// PreSolve is called for touching shapes with pre-solve events enabled, after the manifold
// is updated. The manifold may be modified, set its PointCount to zero to disable the contact.
// The manifold is only valid during the call.
type PreSolve = func(shapeA, shapeB Shape, manifold *Manifold)

// PreContinuous is called for shapes with pre-solve events enabled when continuous collision
// finds a time of impact. Return false to skip the event.
type PreContinuous = func(shapeA, shapeB Shape, point Pos, normal Vec2) bool

// SetPreSolveCallback sets the worlds PreSolve and PreContinuous callbacks. Use nil to remove them.
// The callbacks may be called from multiple goroutines when concurrency is enabled.
func (b World) SetPreSolveCallback(preSolve PreSolve, preContinuous PreContinuous) {
	var preSolveFn, preContinuousFn uintptr

	if preSolve != nil {
		preSolveFn = b.keepCallback("preSolve", preSolveCallback(func(_ *_Stack, shapeIdA, shapeIdB ShapeId, manifold uintptr, _ uintptr) {
			preSolve(Shape{Id: shapeIdA}, Shape{Id: shapeIdB}, (*Manifold)(unsafe.Pointer(manifold)))
		}))
	}

	if preContinuous != nil {
		preContinuousFn = b.keepCallback("preContinuous", preContinuousCallback(func(_ *_Stack, shapeIdA, shapeIdB ShapeId, point Pos, normal Vec2, _ uintptr) uint8 {
			return boolUint8(preContinuous(Shape{Id: shapeIdA}, Shape{Id: shapeIdB}, point, normal))
		}))
	}

	b2World_SetPreSolveCallback(theStack, b.Id, preSolveFn, preContinuousFn, 0)
}

// MakeProxy makes a proxy for use in overlap, shape cast, and related functions.
// At most B2_MAX_POLYGON_VERTICES points are used.
func MakeProxy(points []Vec2, radius float32) ShapeProxy {
	pointsStack := sliceToStack(theStack, points)
	defer pointsStack.Free()

	return b2MakeProxy(theStack, pointsStack.Addr, int32(len(points)), radius)
}

// MakeOffsetProxy makes a proxy with a transform. At most B2_MAX_POLYGON_VERTICES points are used.
func MakeOffsetProxy(points []Vec2, radius float32, position Vec2, rotation Rot) ShapeProxy {
	pointsStack := sliceToStack(theStack, points)
	defer pointsStack.Free()

	return b2MakeOffsetProxy(theStack, pointsStack.Addr, int32(len(points)), radius, position, rotation)
}

// ClipVector clips the vector against the collision planes.
func ClipVector(vector Vec2, planes []CollisionPlane) Vec2 {
	planesStack := sliceToStack(theStack, planes)
	defer planesStack.Free()

	return b2ClipVector(theStack, vector, planesStack.Addr, int32(len(planes)))
}

// PlaneResultFn is called for every shape that collides with the mover. Ignore the plane if
// PlaneResult.Hit is false. Return false to stop the query.
type PlaneResultFn func(shape Shape, plane PlaneResult) bool

// CollideMover calls fcn with the collision planes between the mover capsule and the world,
// for use with SolvePlanes. The mover and the planes are relative to origin.
func (b World) CollideMover(origin Pos, mover Capsule, filter QueryFilter, fcn PlaneResultFn) {
	fn := planeResultCallback(func(_ *_Stack, shapeId ShapeId, plane uintptr, _ uintptr) uint8 {
		return boolUint8(fcn(Shape{Id: shapeId}, *(*PlaneResult)(unsafe.Pointer(plane))))
	})

	// box2d only sees a uintptr, the closure must not live on the goroutine stack, which can move
	escapes(fn)
	defer runtime.KeepAlive(fn)

	moverStack := copyToStack(theStack, mover)
	defer moverStack.Free()

	b2World_CollideMover(theStack, b.Id, origin, moverStack.Addr, filter, __ccgo_fp(fn), 0)
}

// GetSnapshot returns the simulation state of the world, reusing the memory of reuse if it is
// large enough. Call it between steps, it returns an empty slice during a step.
func (b World) GetSnapshot(reuse []byte) []byte {
	size := b2World_GetSnapshot(theStack, b.Id, 0, 0)

	return fillSlice(reuse, size, func(ptr uintptr, capacity int32) int32 {
		return b2World_GetSnapshot(theStack, b.Id, ptr, capacity)
	})
}

// RestoreFromSnapshot resets the world to the state in image, taken with GetSnapshot. Ids of
// objects that existed at the snapshot stay valid, callbacks and threading are kept. User data
// of bodies, shapes and joints is cleared.
//
// It returns false if image is invalid. The world is unchanged if the header is invalid. If the
// data turns out to be corrupt during the restore, the world is unusable and must be destroyed.
func (b World) RestoreFromSnapshot(image []byte) bool {
	ok := b2World_RestoreFromSnapshot(theStack, b.Id, sliceAddr(image), int32(len(image)))
	runtime.KeepAlive(image)

	return ok != 0
}

// CreateWorldFromSnapshot creates a new world from image, taken with GetSnapshot. Callbacks
// are reset to the default. A workerCount of 0 runs the world on a single goroutine.
// It returns false if image is invalid.
func CreateWorldFromSnapshot(image []byte, workerCount int32) (World, bool) {
	id := b2CreateWorldFromSnapshot(theStack, sliceAddr(image), int32(len(image)), workerCount)
	runtime.KeepAlive(image)

	return World{Id: id}, id != WorldId{}
}

// sliceAddr returns the address of the first element for box2d, or 0 if s is empty.
// The caller must keep s alive until box2d is done with it.
func sliceAddr[T any](s []T) uintptr {
	if len(s) == 0 {
		return 0
	}

	// box2d accesses the slice via a uintptr, it must live on the heap
	escapes(s)

	return uintptr(unsafe.Pointer(&s[0]))
}

type BodyEvents struct {
	MoveEvents []BodyMoveEvent
}

func (b World) GetBodyEvents() BodyEvents {
	events := b2World_GetBodyEvents(theStack, b.Id)

	return BodyEvents{
		MoveEvents: copySlice[BodyMoveEvent](events.MoveEvents, events.MoveCount),
	}
}

type SensorEvents struct {
	BeginEvents []SensorBeginTouchEvent
	EndEvents   []SensorEndTouchEvent
}

func (b World) GetSensorEvents() SensorEvents {
	events := b2World_GetSensorEvents(theStack, b.Id)

	return SensorEvents{
		BeginEvents: copySlice[SensorBeginTouchEvent](events.BeginEvents, events.BeginCount),
		EndEvents:   copySlice[SensorEndTouchEvent](events.EndEvents, events.EndCount),
	}
}

type ContactEvents struct {
	BeginEvents []ContactBeginTouchEvent
	EndEvents   []ContactEndTouchEvent
	HitEvents   []ContactHitEvent
}

func (b World) GetContactEvents() ContactEvents {
	events := b2World_GetContactEvents(theStack, b.Id)

	return ContactEvents{
		BeginEvents: copySlice[ContactBeginTouchEvent](events.BeginEvents, events.BeginCount),
		EndEvents:   copySlice[ContactEndTouchEvent](events.EndEvents, events.EndCount),
		HitEvents:   copySlice[ContactHitEvent](events.HitEvents, events.HitCount),
	}
}

type JointEvents struct {
	JointEvents []JointEvent
}

func (b World) GetJointEvents() JointEvents {
	events := b2World_GetJointEvents(theStack, b.Id)

	return JointEvents{
		JointEvents: copySlice[JointEvent](events.JointEvents, events.Count),
	}
}

// copySlice copies n items at data, which box2d owns, onto the Go heap.
func copySlice[T any](data uintptr, n int32) []T {
	return slices.Clone(unsafe.Slice((*T)(unsafe.Pointer(data)), n))
}

func b2DefaultAssertFcnGo(tls *_Stack, condition uintptr, fileName uintptr, lineNumber int32) (r int32) {
	err := fmt.Errorf("%s:%d: %s", toString(fileName), lineNumber, toString(condition))
	panic(err)
}

func (b Joint) GetConstraintTuning() (hertz, dampingRatio float32) {
	buf := theStack.Alloc(8)
	defer theStack.Free(8)

	b2Joint_GetConstraintTuning(theStack, b.Id, buf, buf+4)

	hertz = *(*float32)(unsafe.Pointer(buf))
	dampingRatio = *(*float32)(unsafe.Pointer(buf + 4))
	return
}

func (b DistanceJoint) GetSpringForceRange() (lowerForce, upperForce float32) {
	buf := theStack.Alloc(8)
	defer theStack.Free(8)

	b2DistanceJoint_GetSpringForceRange(theStack, b.Id, buf, buf+4)

	lowerForce = *(*float32)(unsafe.Pointer(buf))
	upperForce = *(*float32)(unsafe.Pointer(buf + 4))
	return
}

func (b World) GetContactTuning() (hertz, dampingRatio, pushSpeed float32) {
	buf := theStack.Alloc(12)
	defer theStack.Free(12)

	b2World_GetContactTuning(theStack, b.Id, buf, buf+4, buf+8)

	hertz = *(*float32)(unsafe.Pointer(buf))
	dampingRatio = *(*float32)(unsafe.Pointer(buf + 4))
	pushSpeed = *(*float32)(unsafe.Pointer(buf + 8))
	return
}

func (b Joint) AsPrismaticJoint() PrismaticJoint {
	if b.GetType() != JointTypePrismatic {
		panic("joint has wrong type")
	}

	return PrismaticJoint{b}
}

func (b Joint) AsMotorJoint() MotorJoint {
	if b.GetType() != JointTypeMotor {
		panic("joint has wrong type")
	}

	return MotorJoint{b}
}

func (b Joint) AsRevoluteJoint() RevoluteJoint {
	if b.GetType() != JointTypeRevolute {
		panic("joint has wrong type")
	}

	return RevoluteJoint{b}
}

func (b Joint) AsDistanceJoint() DistanceJoint {
	if b.GetType() != JointTypeDistance {
		panic("joint has wrong type")
	}

	return DistanceJoint{b}
}

func (b Joint) AsWheelJoint() WheelJoint {
	if b.GetType() != JointTypeWheel {
		panic("joint has wrong type")
	}

	return WheelJoint{b}
}

func (b Joint) AsMoverJoint() MoverJoint {
	if b.GetType() != JointTypeMover {
		panic("joint has wrong type")
	}

	return MoverJoint{b}
}

func (b Joint) AsPogoJoint() PogoJoint {
	if b.GetType() != JointTypePogo {
		panic("joint has wrong type")
	}

	return PogoJoint{b}
}

func (b Joint) AsWeldJoint() WeldJoint {
	if b.GetType() != JointTypeWeld {
		panic("joint has wrong type")
	}

	return WeldJoint{b}
}

func (b Joint) AsFilterJoint() FilterJoint {
	if b.GetType() != JointTypeFilter {
		panic("joint has wrong type")
	}

	return FilterJoint{b}
}

func (r Rot) Angle() float32 {
	return b2Atan2(theStack, r.S, r.C)
}

func (r Rot) XAxis() Vec2 {
	return Vec2{X: r.C, Y: r.S}
}

func (r Rot) YAxis() Vec2 {
	return Vec2{X: -r.S, Y: r.C}
}
