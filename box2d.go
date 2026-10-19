package b2

import (
	"fmt"
	"runtime"
	"slices"
	"unsafe"
)

type RestitutionCallback = func(restitutionA float32, userMaterialIdA int32, restitutionB float32, userMaterialIdB int32) float32
type FrictionCallback = func(frictionA float32, userMaterialIdA int32, frictionB float32, userMaterialIdB int32) float32

var ZeroVec2 = Vec2{}
var IdentityRot = Rot{C: 1.0, S: 0.0}
var IdentityTransform = Transform{Q: IdentityRot}
var ZeroMat = Mat22{}

func init() {
	b2SetAssertFcn(theStack, __ccgo_fp(b2DefaultAssertFcnGo))
}

func ComputeHull(points []Vec2) (Hull, bool) {
	if len(points) > _B2_MAX_POLYGON_VERTICES {
		return Hull{}, false
	}

	// copy to array
	var arr [_B2_MAX_POLYGON_VERTICES]Vec2
	for idx := range min(len(points), len(arr)) {
		arr[idx] = points[idx]
	}

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

func (b Body) GetShapes(reuse []Shape) []Shape {
	itemCount := b.GetShapeCount()

	// allocate memory on the stack to read the ids
	n := int(itemCount) * int(unsafe.Sizeof(ShapeId{}))
	ptr := theStack.Alloc(n)
	defer theStack.Free(n)

	// fill the data
	_ = b2Body_GetShapes(theStack, b.Id, ptr, itemCount)

	// allocate memory for the result
	if cap(reuse) < int(itemCount) {
		// need to allocate more space
		reuse = make([]Shape, itemCount)
	}

	// ensure the data is on the heap
	items := reuse[:0]

	for _, itemId := range unsafe.Slice((*ShapeId)(unsafe.Pointer(ptr)), itemCount) {
		items = append(items, Shape{Id: itemId})
	}

	return items
}

func (b Body) GetJoints(reuse []Joint) []Joint {
	itemCount := b.GetJointCount()

	// allocate memory on the stack to read the ids
	n := int(itemCount) * int(unsafe.Sizeof(JointId{}))
	ptr := theStack.Alloc(n)
	defer theStack.Free(n)

	// fill the data
	_ = b2Body_GetJoints(theStack, b.Id, ptr, itemCount)

	// allocate memory for the result
	if cap(reuse) < int(itemCount) {
		// need to allocate more space
		reuse = make([]Joint, itemCount)
	}

	items := reuse[:0]

	// copy the items
	for _, itemId := range unsafe.Slice((*JointId)(unsafe.Pointer(ptr)), itemCount) {
		items = append(items, Joint{Id: itemId})
	}

	return items
}

func (b Chain) GetSegments(reuse []Shape) []Shape {
	itemCount := b.GetSegmentCount()

	// allocate memory on the stack to read the ids
	n := int(itemCount) * int(unsafe.Sizeof(ShapeId{}))
	ptr := theStack.Alloc(n)
	defer theStack.Free(n)

	// fill the data
	_ = b2Chain_GetSegments(theStack, b.Id, ptr, itemCount)

	// allocate memory for the result
	if cap(reuse) < int(itemCount) {
		// need to allocate more space
		reuse = make([]Shape, itemCount)
	}

	items := reuse[:0]

	// copy the items
	for _, itemId := range unsafe.Slice((*ShapeId)(unsafe.Pointer(ptr)), itemCount) {
		items = append(items, Shape{Id: itemId})
	}

	return items
}

func (b Shape) GetSensorOverlaps(reuse []Shape) []Shape {
	itemsCap := b.GetSensorCapacity()

	// allocate memory on the stack to read the ids
	n := int(itemsCap) * int(unsafe.Sizeof(ShapeId{}))
	ptr := theStack.Alloc(n)
	defer theStack.Free(n)

	// fill the data
	itemCount := b2Shape_GetSensorOverlaps(theStack, b.Id, ptr, itemsCap)

	// allocate memory for the result
	if cap(reuse) < int(itemCount) {
		// need to allocate more space
		reuse = make([]Shape, itemCount)
	}

	items := reuse[:0]

	// copy the items
	for _, itemId := range unsafe.Slice((*ShapeId)(unsafe.Pointer(ptr)), itemCount) {
		items = append(items, Shape{Id: itemId})
	}

	return items
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
		fn = b.keepCallback("restitution", func(_ *_Stack, restitutionA float32, userMaterialIdA int32, restitutionB float32, userMaterialIdB int32) float32 {
			return callback(restitutionA, userMaterialIdA, restitutionB, userMaterialIdB)
		})
	}

	b2World_SetRestitutionCallback(theStack, b.Id, fn)
}

// SetFrictionCallback sets the worlds FrictionCallback. Use nil to restore the default.
// The callback may be called from multiple goroutines when concurrency is enabled.
func (b World) SetFrictionCallback(callback FrictionCallback) {
	var fn uintptr
	if callback != nil {
		fn = b.keepCallback("friction", func(_ *_Stack, frictionA float32, userMaterialIdA int32, frictionB float32, userMaterialIdB int32) float32 {
			return callback(frictionA, userMaterialIdA, frictionB, userMaterialIdB)
		})
	}

	b2World_SetFrictionCallback(theStack, b.Id, fn)
}

type OverlapResult func(shapeId Shape) bool

// OverlapAABB calls fcn for every shape that overlaps the aabb. Return false to stop the query.
func (b World) OverlapAABB(aabb AABB, filter QueryFilter, fcn OverlapResult) TreeStats {
	fn := worldOverlapCallback(fcn)
	defer runtime.KeepAlive(fn)

	return b2World_OverlapAABB(theStack, b.Id, aabb, filter, __ccgo_fp(fn), 0)
}

// OverlapShape calls fcn for every shape that overlaps the proxy. Return false to stop the query.
func (b World) OverlapShape(proxy ShapeProxy, filter QueryFilter, fcn OverlapResult) (r TreeStats) {
	fn := worldOverlapCallback(fcn)
	defer runtime.KeepAlive(fn)

	proxyStack := copyToStack(theStack, proxy)
	defer proxyStack.Free()

	r = b2World_OverlapShape(theStack, b.Id, proxyStack.Addr, filter, __ccgo_fp(fn), 0)

	return
}

func worldOverlapCallback(fcn OverlapResult) func(*_Stack, ShapeId, uintptr) uint8 {
	fn := func(_ *_Stack, shapeId ShapeId, _ uintptr) uint8 {
		return boolUint8(fcn(Shape{Id: shapeId}))
	}

	// box2d only sees a uintptr, the closure must not live on the goroutine stack, which can move
	escapes(fn)

	return fn
}

// CustomFilter returns false to prevent the two shapes from colliding.
type CustomFilter func(shapeIdA, shapeIdB Shape) bool

// SetCustomFilterCallback sets the worlds CustomFilter. Use nil to remove it.
// The callback may be called from multiple goroutines when concurrency is enabled.
func (b World) SetCustomFilterCallback(fcn CustomFilter) {
	var fn uintptr
	if fcn != nil {
		fn = b.keepCallback("customFilter", func(_ *_Stack, shapeIdA, shapeIdB ShapeId, _ uintptr) uint8 {
			return boolUint8(fcn(Shape{Id: shapeIdA}, Shape{Id: shapeIdB}))
		})
	}

	b2World_SetCustomFilterCallback(theStack, b.Id, fn, 0)
}

// PreSolve is called for touching shapes with pre-solve events enabled. Return false to disable the contact for this step.
type PreSolve = func(shapeIdA, shapeIdB Shape, manifold Manifold) bool

// SetPreSolveCallback sets the worlds PreSolve callback. Use nil to remove it.
// The callback may be called from multiple goroutines when concurrency is enabled.
func (b World) SetPreSolveCallback(fcn PreSolve) {
	var fn uintptr
	if fcn != nil {
		fn = b.keepCallback("preSolve", func(_ *_Stack, shapeIdA, shapeIdB ShapeId, manifold uintptr, _ uintptr) uint8 {
			return boolUint8(fcn(Shape{Id: shapeIdA}, Shape{Id: shapeIdB}, derefToValue[Manifold](manifold)))
		})
	}

	b2World_SetPreSolveCallback(theStack, b.Id, fn, 0)
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

func copySlice[T any](data uintptr, n int32) []T {
	event := (*T)(unsafe.Pointer(data))
	bodyMoveEvents := unsafe.Slice(event, n)
	return slices.Clone(bodyMoveEvents)
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
func (b Joint) AsMouseJoint() MouseJoint {
	if b.GetType() != JointTypeMouse {
		panic("joint has wrong type")
	}

	return MouseJoint{b}
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
