package b2

import (
	"fmt"
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
	// allow up to 1024 overlaps to be returned
	itemsCap := int32(1024)

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

// SetRestitutionCallback sets the worlds RestitutionCallback.
// CAUTION: Callback must not be a closure
func (b World) SetRestitutionCallback(callback RestitutionCallback) {
	b2World_SetRestitutionCallback(theStack, b.Id, __ccgo_fp(callback))
}

// SetFrictionCallback sets the worlds FrictionCallback.
// CAUTION: Callback must not be a closure
func (b World) SetFrictionCallback(callback FrictionCallback) {
	b2World_SetFrictionCallback(theStack, b.Id, __ccgo_fp(callback))
}

type OverlapResult func(shapeId Shape) bool

// OverlapAABB sets the worlds OverlapResultFcn
func (b World) OverlapAABB(aabb AABB, filter QueryFilter, fcn OverlapResult) TreeStats {
	fn := theStack.RegisterObject(fcn)
	defer theStack.ForgetObject(fn)

	return b2World_OverlapAABB(theStack, b.Id, aabb, filter, __ccgo_fp(worldOverlapCallback), fn)
}

func (b World) OverlapShape(proxy ShapeProxy, filter QueryFilter, fcn OverlapResult) (r TreeStats) {
	fn := theStack.RegisterObject(fcn)
	defer theStack.ForgetObject(fn)

	proxyStack := copyToStack(theStack, proxy)
	defer proxyStack.Free()

	r = b2World_OverlapShape(theStack, b.Id, proxyStack.Addr, filter, __ccgo_fp(worldOverlapCallback), fn)

	return
}

func worldOverlapCallback(tls *_Stack, shapeId ShapeId, context uintptr) bool {
	fn := tls.GetObject(context).(OverlapResult)
	return fn(Shape{Id: shapeId})
}

type CustomFilter func(shapeIdA, shapeIdB Shape) bool

func (b World) SetCustomFilterCallback(fcn CustomFilter) {
	fn := theStack.RegisterObject(fcn)
	defer theStack.ForgetObject(fn)

	b2World_SetCustomFilterCallback(theStack, b.Id, __ccgo_fp(worldCustomFilterCallback), fn)
}

func worldCustomFilterCallback(tls *_Stack, shapeIdA, shapeIdB ShapeId, context uintptr) bool {
	fn := tls.GetObject(context).(CustomFilter)
	return fn(Shape{Id: shapeIdA}, Shape{Id: shapeIdB})
}

type PreSolve = func(shapeIdA, shapeIdB Shape, manifold Manifold)

func (b World) SetPreSolveCallback(fcn PreSolve) {
	fn := theStack.RegisterObject(fcn)
	defer theStack.ForgetObject(fn)

	b2World_SetPreSolveCallback(theStack, b.Id, __ccgo_fp(worldPreSolveCallback), fn)
}

func worldPreSolveCallback(tls *_Stack, shapeIdA, shapeIdB ShapeId, manifold uintptr, context uintptr) {
	shapeA := Shape{Id: shapeIdA}
	shapeB := Shape{Id: shapeIdB}
	manifoldValue := derefToValue[Manifold](manifold)

	fn := tls.GetObject(context).(PreSolve)
	fn(shapeA, shapeB, manifoldValue)
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
