package b2

import (
	"unsafe"
	"runtime"
)

var _ unsafe.Pointer
var _ runtime.MemStats

type Body struct {
	Id BodyId
}
type Shape struct {
	Id ShapeId
}
type World struct {
	Id WorldId
}
type Joint struct {
	Id JointId
}
type Chain struct {
	Id ChainId
}
type Contact struct {
	Id ContactId
}
type PrismaticJoint struct {
	Joint
}
type MotorJoint struct {
	Joint
}
type RevoluteJoint struct {
	Joint
}
type DistanceJoint struct {
	Joint
}
type WheelJoint struct {
	Joint
}
type WeldJoint struct {
	Joint
}
type MoverJoint struct {
	Joint
}
type PogoJoint struct {
	Joint
}
type FilterJoint struct {
	Joint
}

func IsValidAABB(a1 AABB) (r uint8) { r = b2IsValidAABB(theStack, a1); return }
func (b World) CreateBody(def BodyDef) (r Body) {
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	r.Id = b2CreateBody(theStack, b.Id, defStack.Addr)
	return
}
func (b Body) DestroyBody()                     { b2DestroyBody(theStack, b.Id) }
func (b Body) GetContactCapacity() (r int32)    { r = b2Body_GetContactCapacity(theStack, b.Id); return }
func (b Body) ComputeBounds() (r AABB)          { r = b2Body_ComputeBounds(theStack, b.Id); return }
func (b Body) GetPosition() (r Pos)             { r = b2Body_GetPosition(theStack, b.Id); return }
func (b Body) GetRotation() (r Rot)             { r = b2Body_GetRotation(theStack, b.Id); return }
func (b Body) GetTransform() (r WorldTransform) { r = b2Body_GetTransform(theStack, b.Id); return }
func (b Body) GetLocalPoint(worldPoint Pos) (r Vec2) {
	r = b2Body_GetLocalPoint(theStack, b.Id, worldPoint)
	return
}
func (b Body) GetWorldPoint(localPoint Vec2) (r Pos) {
	r = b2Body_GetWorldPoint(theStack, b.Id, localPoint)
	return
}
func (b Body) GetLocalVector(worldVector Vec2) (r Vec2) {
	r = b2Body_GetLocalVector(theStack, b.Id, worldVector)
	return
}
func (b Body) GetWorldVector(localVector Vec2) (r Vec2) {
	r = b2Body_GetWorldVector(theStack, b.Id, localVector)
	return
}
func (b Body) SetTransform(position Pos, rotation Rot) {
	b2Body_SetTransform(theStack, b.Id, position, rotation)
}
func (b Body) GetLinearVelocity() (r Vec2)     { r = b2Body_GetLinearVelocity(theStack, b.Id); return }
func (b Body) GetAngularVelocity() (r float32) { r = b2Body_GetAngularVelocity(theStack, b.Id); return }
func (b Body) SetLinearVelocity(linearVelocity Vec2) {
	b2Body_SetLinearVelocity(theStack, b.Id, linearVelocity)
}
func (b Body) SetAngularVelocity(angularVelocity float32) {
	b2Body_SetAngularVelocity(theStack, b.Id, angularVelocity)
}
func (b Body) SetTargetTransform(target WorldTransform, timeStep float32, wake uint8) {
	b2Body_SetTargetTransform(theStack, b.Id, target, timeStep, wake)
}
func (b Body) GetLocalPointVelocity(localPoint Vec2) (r1 Vec2) {
	r1 = b2Body_GetLocalPointVelocity(theStack, b.Id, localPoint)
	return
}
func (b Body) GetWorldPointVelocity(worldPoint Pos) (r1 Vec2) {
	r1 = b2Body_GetWorldPointVelocity(theStack, b.Id, worldPoint)
	return
}
func (b Body) ApplyForce(force Vec2, point Pos, wake uint8) {
	b2Body_ApplyForce(theStack, b.Id, force, point, wake)
}
func (b Body) ApplyForceToCenter(force Vec2, wake uint8) {
	b2Body_ApplyForceToCenter(theStack, b.Id, force, wake)
}
func (b Body) ApplyTorque(torque float32, wake uint8) {
	b2Body_ApplyTorque(theStack, b.Id, torque, wake)
}
func (b Body) ClearForces() { b2Body_ClearForces(theStack, b.Id) }
func (b Body) ApplyLinearImpulse(impulse Vec2, point Pos, wake uint8) {
	b2Body_ApplyLinearImpulse(theStack, b.Id, impulse, point, wake)
}
func (b Body) ApplyLinearImpulseToCenter(impulse Vec2, wake uint8) {
	b2Body_ApplyLinearImpulseToCenter(theStack, b.Id, impulse, wake)
}
func (b Body) ApplyAngularImpulse(impulse float32, wake uint8) {
	b2Body_ApplyAngularImpulse(theStack, b.Id, impulse, wake)
}
func (b Body) GetType() (r BodyType)        { r = b2Body_GetType(theStack, b.Id); return }
func (b Body) SetType(type1 BodyType)       { b2Body_SetType(theStack, b.Id, type1) }
func (b Body) SetUserData(userData uintptr) { b2Body_SetUserData(theStack, b.Id, userData) }
func (b Body) GetUserData() (r uintptr)     { r = b2Body_GetUserData(theStack, b.Id); return }
func (b Body) GetMass() (r float32)         { r = b2Body_GetMass(theStack, b.Id); return }
func (b Body) GetRotationalInertia() (r float32) {
	r = b2Body_GetRotationalInertia(theStack, b.Id)
	return
}
func (b Body) GetLocalCenter() (r Vec2)      { r = b2Body_GetLocalCenter(theStack, b.Id); return }
func (b Body) GetWorldCenter() (r Pos)       { r = b2Body_GetWorldCenter(theStack, b.Id); return }
func (b Body) SetMassData(massData MassData) { b2Body_SetMassData(theStack, b.Id, massData) }
func (b Body) GetMassData() (r MassData)     { r = b2Body_GetMassData(theStack, b.Id); return }
func (b Body) UpdateMassFromShapes()         { b2Body_UpdateMassFromShapes(theStack, b.Id) }
func (b Body) SetLinearDamping(linearDamping float32) {
	b2Body_SetLinearDamping(theStack, b.Id, linearDamping)
}
func (b Body) GetLinearDamping() (r float32) { r = b2Body_GetLinearDamping(theStack, b.Id); return }
func (b Body) SetAngularDamping(angularDamping float32) {
	b2Body_SetAngularDamping(theStack, b.Id, angularDamping)
}
func (b Body) GetAngularDamping() (r float32) { r = b2Body_GetAngularDamping(theStack, b.Id); return }
func (b Body) SetGravityScale(gravityScale float32) {
	b2Body_SetGravityScale(theStack, b.Id, gravityScale)
}
func (b Body) GetGravityScale() (r float32) { r = b2Body_GetGravityScale(theStack, b.Id); return }
func (b Body) IsAwake() (r uint8)           { r = b2Body_IsAwake(theStack, b.Id); return }
func (b Body) SetAwake(awake uint8)         { b2Body_SetAwake(theStack, b.Id, awake) }
func (b Body) WakeTouching()                { b2Body_WakeTouching(theStack, b.Id) }
func (b Body) IsEnabled() (r uint8)         { r = b2Body_IsEnabled(theStack, b.Id); return }
func (b Body) IsSleepingEnabled() (r uint8) { r = b2Body_IsSleepingEnabled(theStack, b.Id); return }
func (b Body) SetSleepThreshold(sleepThreshold float32) {
	b2Body_SetSleepThreshold(theStack, b.Id, sleepThreshold)
}
func (b Body) GetSleepThreshold() (r float32) { r = b2Body_GetSleepThreshold(theStack, b.Id); return }
func (b Body) SetSafetyFactor(safetyFactor float32) {
	b2Body_SetSafetyFactor(theStack, b.Id, safetyFactor)
}
func (b Body) GetSafetyFactor() (r float32)      { r = b2Body_GetSafetyFactor(theStack, b.Id); return }
func (b Body) EnableSleeping(enableSleep uint8)  { b2Body_EnableSleeping(theStack, b.Id, enableSleep) }
func (b Body) Disable()                          { b2Body_Disable(theStack, b.Id) }
func (b Body) Enable()                           { b2Body_Enable(theStack, b.Id) }
func (b Body) SetMotionLocks(locks MotionLocks)  { b2Body_SetMotionLocks(theStack, b.Id, locks) }
func (b Body) GetMotionLocks() (r MotionLocks)   { r = b2Body_GetMotionLocks(theStack, b.Id); return }
func (b Body) SetBullet(flag uint8)              { b2Body_SetBullet(theStack, b.Id, flag) }
func (b Body) IsBullet() (r uint8)               { r = b2Body_IsBullet(theStack, b.Id); return }
func (b Body) EnableContactRecycling(flag uint8) { b2Body_EnableContactRecycling(theStack, b.Id, flag) }
func (b Body) IsContactRecyclingEnabled() (r uint8) {
	r = b2Body_IsContactRecyclingEnabled(theStack, b.Id)
	return
}
func (b Body) EnableContactEvents(flag uint8) { b2Body_EnableContactEvents(theStack, b.Id, flag) }
func (b Body) EnableHitEvents(flag uint8)     { b2Body_EnableHitEvents(theStack, b.Id, flag) }
func (b Body) GetWorld() (r World)            { r.Id = b2Body_GetWorld(theStack, b.Id); return }
func (b Body) GetShapeCount() (r int32)       { r = b2Body_GetShapeCount(theStack, b.Id); return }
func (b Body) GetJointCount() (r int32)       { r = b2Body_GetJointCount(theStack, b.Id); return }
func (b Body) GetMinExtent() (r float32)      { r = b2Body_GetMinExtent(theStack, b.Id); return }
func (b Body) GetMaxExtent() (r float32)      { r = b2Body_GetMaxExtent(theStack, b.Id); return }
func (b Body) GetMaxExtentFromOrigin() (r float32) {
	r = b2Body_GetMaxExtentFromOrigin(theStack, b.Id)
	return
}
func GetGraphColor(index int32) (r HexColor)     { r = b2GetGraphColor(theStack, index); return }
func (b Contact) GetData() (r ContactData)       { r = b2Contact_GetData(theStack, b.Id); return }
func SetLengthUnitsPerMeter(lengthUnits float32) { b2SetLengthUnitsPerMeter(theStack, lengthUnits) }
func GetLengthUnitsPerMeter() (r float32)        { r = b2GetLengthUnitsPerMeter(theStack); return }
func GetVersion() (r Version)                    { r = b2GetVersion(theStack); return }
func IsDoublePrecision() (r uint8)               { r = b2IsDoublePrecision(theStack); return }
func GetByteCount() (r size_t)                   { r = b2GetByteCount(theStack); return }
func GetSweepTransform(sweep Sweep, time float32) (r Transform) {
	sweepStack := copyToStack(theStack, sweep)
	defer sweepStack.Free()
	r = b2GetSweepTransform(theStack, sweepStack.Addr, time)
	return
}
func ComputeSegmentDistance(p1 Vec2, q1 Vec2, p2 Vec2, q2 Vec2) (r1 SegmentDistanceResult) {
	r1 = b2ComputeSegmentDistance(theStack, p1, q1, p2, q2)
	return
}
func CastShape(input ShapeCastPairInput) (r CastOutput) {
	inputStack := copyToStack(theStack, input)
	defer inputStack.Free()
	r = b2CastShape(theStack, inputStack.Addr)
	return
}
func ComputeTimeOfImpact(input TOIInput) (r1 TOIOutput) {
	inputStack := copyToStack(theStack, input)
	defer inputStack.Free()
	r1 = b2ComputeTimeOfImpact(theStack, inputStack.Addr)
	return
}
func (b DistanceJoint) SetRestLength(length float32) {
	b2DistanceJoint_SetRestLength(theStack, b.Id, length)
}
func (b DistanceJoint) GetRestLength() (r float32) {
	r = b2DistanceJoint_GetRestLength(theStack, b.Id)
	return
}
func (b DistanceJoint) EnableLimit(enableLimit uint8) {
	b2DistanceJoint_EnableLimit(theStack, b.Id, enableLimit)
}
func (b DistanceJoint) IsLimitEnabled() (r uint8) {
	r = b2DistanceJoint_IsLimitEnabled(theStack, b.Id)
	return
}
func (b DistanceJoint) SetLimits(minLength float32, maxLength float32) {
	b2DistanceJoint_SetLimits(theStack, b.Id, minLength, maxLength)
}
func (b DistanceJoint) GetLowerLimit() (r float32) {
	r = b2DistanceJoint_GetLowerLimit(theStack, b.Id)
	return
}
func (b DistanceJoint) GetUpperLimit() (r float32) {
	r = b2DistanceJoint_GetUpperLimit(theStack, b.Id)
	return
}
func (b DistanceJoint) GetCurrentLength() (r1 float32) {
	r1 = b2DistanceJoint_GetCurrentLength(theStack, b.Id)
	return
}
func (b DistanceJoint) EnableSpring(enableSpring uint8) {
	b2DistanceJoint_EnableSpring(theStack, b.Id, enableSpring)
}
func (b DistanceJoint) IsSpringEnabled() (r uint8) {
	r = b2DistanceJoint_IsSpringEnabled(theStack, b.Id)
	return
}
func (b DistanceJoint) SetSpringForceRange(lowerForce float32, upperForce float32) {
	b2DistanceJoint_SetSpringForceRange(theStack, b.Id, lowerForce, upperForce)
}
func (b DistanceJoint) SetSpringHertz(hertz float32) {
	b2DistanceJoint_SetSpringHertz(theStack, b.Id, hertz)
}
func (b DistanceJoint) SetSpringDampingRatio(dampingRatio float32) {
	b2DistanceJoint_SetSpringDampingRatio(theStack, b.Id, dampingRatio)
}
func (b DistanceJoint) GetSpringHertz() (r float32) {
	r = b2DistanceJoint_GetSpringHertz(theStack, b.Id)
	return
}
func (b DistanceJoint) GetSpringDampingRatio() (r float32) {
	r = b2DistanceJoint_GetSpringDampingRatio(theStack, b.Id)
	return
}
func (b DistanceJoint) GetSpringForce() (r float32) {
	r = b2DistanceJoint_GetSpringForce(theStack, b.Id)
	return
}
func (b DistanceJoint) EnableMotor(enableMotor uint8) {
	b2DistanceJoint_EnableMotor(theStack, b.Id, enableMotor)
}
func (b DistanceJoint) IsMotorEnabled() (r uint8) {
	r = b2DistanceJoint_IsMotorEnabled(theStack, b.Id)
	return
}
func (b DistanceJoint) SetTargetMotorSpeed(motorSpeed float32) {
	b2DistanceJoint_SetTargetMotorSpeed(theStack, b.Id, motorSpeed)
}
func (b DistanceJoint) GetTargetMotorSpeed() (r float32) {
	r = b2DistanceJoint_GetTargetMotorSpeed(theStack, b.Id)
	return
}
func (b DistanceJoint) GetCurrentMotorForce() (r float32) {
	r = b2DistanceJoint_GetCurrentMotorForce(theStack, b.Id)
	return
}
func (b DistanceJoint) SetMaxMotorForce(force float32) {
	b2DistanceJoint_SetMaxMotorForce(theStack, b.Id, force)
}
func (b DistanceJoint) GetMaxMotorForce() (r float32) {
	r = b2DistanceJoint_GetMaxMotorForce(theStack, b.Id)
	return
}
func IsValidRay(input RayCastInput) (r uint8) {
	inputStack := copyToStack(theStack, input)
	defer inputStack.Free()
	r = b2IsValidRay(theStack, inputStack.Addr)
	return
}
func MakePolygon(hull Hull, radius float32) (r Polygon) {
	hullStack := copyToStack(theStack, hull)
	defer hullStack.Free()
	r = b2MakePolygon(theStack, hullStack.Addr, radius)
	return
}
func MakeOffsetPolygon(hull Hull, position Vec2, rotation Rot) (r Polygon) {
	hullStack := copyToStack(theStack, hull)
	defer hullStack.Free()
	r = b2MakeOffsetPolygon(theStack, hullStack.Addr, position, rotation)
	return
}
func MakeOffsetRoundedPolygon(hull Hull, position Vec2, rotation Rot, radius float32) (r Polygon) {
	hullStack := copyToStack(theStack, hull)
	defer hullStack.Free()
	r = b2MakeOffsetRoundedPolygon(theStack, hullStack.Addr, position, rotation, radius)
	return
}
func MakeSquare(halfWidth float32) (r Polygon) { r = b2MakeSquare(theStack, halfWidth); return }
func MakeBox(halfWidth float32, halfHeight float32) (r Polygon) {
	r = b2MakeBox(theStack, halfWidth, halfHeight)
	return
}
func MakeRoundedBox(halfWidth float32, halfHeight float32, radius float32) (r Polygon) {
	r = b2MakeRoundedBox(theStack, halfWidth, halfHeight, radius)
	return
}
func MakeOffsetBox(halfWidth float32, halfHeight float32, center Vec2, rotation Rot) (r Polygon) {
	r = b2MakeOffsetBox(theStack, halfWidth, halfHeight, center, rotation)
	return
}
func MakeOffsetRoundedBox(halfWidth float32, halfHeight float32, center Vec2, rotation Rot, radius float32) (r Polygon) {
	r = b2MakeOffsetRoundedBox(theStack, halfWidth, halfHeight, center, rotation, radius)
	return
}
func TransformPolygon(transform Transform, polygon Polygon) (r Polygon) {
	polygonStack := copyToStack(theStack, polygon)
	defer polygonStack.Free()
	r = b2TransformPolygon(theStack, transform, polygonStack.Addr)
	return
}
func ComputeCircleMass(shape Circle, density float32) (r MassData) {
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	r = b2ComputeCircleMass(theStack, shapeStack.Addr, density)
	return
}
func ComputeCapsuleMass(shape Capsule, density float32) (r MassData) {
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	r = b2ComputeCapsuleMass(theStack, shapeStack.Addr, density)
	return
}
func ComputePolygonMass(shape Polygon, density float32) (r1 MassData) {
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	r1 = b2ComputePolygonMass(theStack, shapeStack.Addr, density)
	return
}
func ComputeCircleBounds(shape Circle, xf WorldTransform) (r AABB) {
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	r = b2ComputeCircleBounds(theStack, shapeStack.Addr, xf)
	return
}
func ComputeCapsuleBounds(shape Capsule, xf WorldTransform) (r AABB) {
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	r = b2ComputeCapsuleBounds(theStack, shapeStack.Addr, xf)
	return
}
func ComputePolygonBounds(shape Polygon, xf WorldTransform) (r AABB) {
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	r = b2ComputePolygonBounds(theStack, shapeStack.Addr, xf)
	return
}
func ComputeSegmentBounds(shape Segment, xf WorldTransform) (r AABB) {
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	r = b2ComputeSegmentBounds(theStack, shapeStack.Addr, xf)
	return
}
func PointInCircle(shape Circle, point Vec2) (r uint8) {
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	r = b2PointInCircle(theStack, shapeStack.Addr, point)
	return
}
func PointInCapsule(shape Capsule, point Vec2) (r uint8) {
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	r = b2PointInCapsule(theStack, shapeStack.Addr, point)
	return
}
func PointInPolygon(shape Polygon, _point Vec2) (r uint8) {
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	r = b2PointInPolygon(theStack, shapeStack.Addr, _point)
	return
}
func CastRayCircle(shape Circle, input RayCastInput) (r1 CastOutput) {
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	inputStack := copyToStack(theStack, input)
	defer inputStack.Free()
	r1 = b2CastRayCircle(theStack, shapeStack.Addr, inputStack.Addr)
	return
}
func CastRayCapsule(shape Capsule, input RayCastInput) (r CastOutput) {
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	inputStack := copyToStack(theStack, input)
	defer inputStack.Free()
	r = b2CastRayCapsule(theStack, shapeStack.Addr, inputStack.Addr)
	return
}
func CastRaySegment(shape Segment, input RayCastInput, oneSided uint8) (r CastOutput) {
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	inputStack := copyToStack(theStack, input)
	defer inputStack.Free()
	r = b2CastRaySegment(theStack, shapeStack.Addr, inputStack.Addr, oneSided)
	return
}
func CastRayPolygon(shape Polygon, input RayCastInput) (r CastOutput) {
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	inputStack := copyToStack(theStack, input)
	defer inputStack.Free()
	r = b2CastRayPolygon(theStack, shapeStack.Addr, inputStack.Addr)
	return
}
func CastShapeCircle(shape Circle, input ShapeCastInput) (r CastOutput) {
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	inputStack := copyToStack(theStack, input)
	defer inputStack.Free()
	r = b2CastShapeCircle(theStack, shapeStack.Addr, inputStack.Addr)
	return
}
func CastShapeCapsule(shape Capsule, input ShapeCastInput) (r CastOutput) {
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	inputStack := copyToStack(theStack, input)
	defer inputStack.Free()
	r = b2CastShapeCapsule(theStack, shapeStack.Addr, inputStack.Addr)
	return
}
func CastShapeSegment(shape Segment, input ShapeCastInput) (r CastOutput) {
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	inputStack := copyToStack(theStack, input)
	defer inputStack.Free()
	r = b2CastShapeSegment(theStack, shapeStack.Addr, inputStack.Addr)
	return
}
func CastShapePolygon(shape Polygon, input ShapeCastInput) (r CastOutput) {
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	inputStack := copyToStack(theStack, input)
	defer inputStack.Free()
	r = b2CastShapePolygon(theStack, shapeStack.Addr, inputStack.Addr)
	return
}
func ValidateHull(hull Hull) (r uint8) {
	hullStack := copyToStack(theStack, hull)
	defer hullStack.Free()
	r = b2ValidateHull(theStack, hullStack.Addr)
	return
}
func DefaultDistanceJointDef() (r DistanceJointDef) { r = b2DefaultDistanceJointDef(theStack); return }
func DefaultFilterJointDef() (r FilterJointDef)     { r = b2DefaultFilterJointDef(theStack); return }
func DefaultMotorJointDef() (r MotorJointDef)       { r = b2DefaultMotorJointDef(theStack); return }
func DefaultMoverJointDef() (r MoverJointDef)       { r = b2DefaultMoverJointDef(theStack); return }
func DefaultPogoJointDef() (r PogoJointDef)         { r = b2DefaultPogoJointDef(theStack); return }
func DefaultPrismaticJointDef() (r PrismaticJointDef) {
	r = b2DefaultPrismaticJointDef(theStack)
	return
}
func DefaultRevoluteJointDef() (r RevoluteJointDef) { r = b2DefaultRevoluteJointDef(theStack); return }
func DefaultWeldJointDef() (r WeldJointDef)         { r = b2DefaultWeldJointDef(theStack); return }
func DefaultWheelJointDef() (r WheelJointDef)       { r = b2DefaultWheelJointDef(theStack); return }
func DefaultExplosionDef() (r ExplosionDef)         { r = b2DefaultExplosionDef(theStack); return }
func (b World) CreateDistanceJoint(def DistanceJointDef) (r DistanceJoint) {
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	r.Id = b2CreateDistanceJoint(theStack, b.Id, defStack.Addr)
	return
}
func (b World) CreateFilterJoint(def FilterJointDef) (r FilterJoint) {
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	r.Id = b2CreateFilterJoint(theStack, b.Id, defStack.Addr)
	return
}
func (b World) CreateMotorJoint(def MotorJointDef) (r MotorJoint) {
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	r.Id = b2CreateMotorJoint(theStack, b.Id, defStack.Addr)
	return
}
func (b World) CreateMoverJoint(def MoverJointDef) (r MoverJoint) {
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	r.Id = b2CreateMoverJoint(theStack, b.Id, defStack.Addr)
	return
}
func (b World) CreatePogoJoint(def PogoJointDef) (r PogoJoint) {
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	r.Id = b2CreatePogoJoint(theStack, b.Id, defStack.Addr)
	return
}
func (b World) CreatePrismaticJoint(def PrismaticJointDef) (r PrismaticJoint) {
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	r.Id = b2CreatePrismaticJoint(theStack, b.Id, defStack.Addr)
	return
}
func (b World) CreateRevoluteJoint(def RevoluteJointDef) (r RevoluteJoint) {
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	r.Id = b2CreateRevoluteJoint(theStack, b.Id, defStack.Addr)
	return
}
func (b World) CreateWeldJoint(def WeldJointDef) (r WeldJoint) {
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	r.Id = b2CreateWeldJoint(theStack, b.Id, defStack.Addr)
	return
}
func (b World) CreateWheelJoint(def WheelJointDef) (r WheelJoint) {
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	r.Id = b2CreateWheelJoint(theStack, b.Id, defStack.Addr)
	return
}
func (b Joint) DestroyJoint()          { b2DestroyJoint(theStack, b.Id) }
func (b Joint) GetType() (r JointType) { r = b2Joint_GetType(theStack, b.Id); return }
func (b Joint) GetBodyA() (r Body)     { r.Id = b2Joint_GetBodyA(theStack, b.Id); return }
func (b Joint) GetBodyB() (r Body)     { r.Id = b2Joint_GetBodyB(theStack, b.Id); return }
func (b Joint) GetWorld() (r World)    { r.Id = b2Joint_GetWorld(theStack, b.Id); return }
func (b Joint) SetLocalFrameA(localFrame Transform) {
	b2Joint_SetLocalFrameA(theStack, b.Id, localFrame)
}
func (b Joint) GetLocalFrameA() (r Transform) { r = b2Joint_GetLocalFrameA(theStack, b.Id); return }
func (b Joint) SetLocalFrameB(localFrame Transform) {
	b2Joint_SetLocalFrameB(theStack, b.Id, localFrame)
}
func (b Joint) GetLocalFrameB() (r Transform) { r = b2Joint_GetLocalFrameB(theStack, b.Id); return }
func (b Joint) SetCollideConnected(shouldCollide uint8) {
	b2Joint_SetCollideConnected(theStack, b.Id, shouldCollide)
}
func (b Joint) GetCollideConnected() (r uint8) {
	r = b2Joint_GetCollideConnected(theStack, b.Id)
	return
}
func (b Joint) SetUserData(userData uintptr) { b2Joint_SetUserData(theStack, b.Id, userData) }
func (b Joint) GetUserData() (r uintptr)     { r = b2Joint_GetUserData(theStack, b.Id); return }
func (b Joint) WakeBodies()                  { b2Joint_WakeBodies(theStack, b.Id) }
func (b Joint) GetConstraintForce() (r Vec2) { r = b2Joint_GetConstraintForce(theStack, b.Id); return }
func (b Joint) GetConstraintTorque() (r float32) {
	r = b2Joint_GetConstraintTorque(theStack, b.Id)
	return
}
func (b Joint) GetLinearSeparation() (r1 float32) {
	r1 = b2Joint_GetLinearSeparation(theStack, b.Id)
	return
}
func (b Joint) GetAngularSeparation() (r float32) {
	r = b2Joint_GetAngularSeparation(theStack, b.Id)
	return
}
func (b Joint) SetConstraintTuning(hertz float32, dampingRatio float32) {
	b2Joint_SetConstraintTuning(theStack, b.Id, hertz, dampingRatio)
}
func (b Joint) SetForceThreshold(threshold float32) {
	b2Joint_SetForceThreshold(theStack, b.Id, threshold)
}
func (b Joint) GetForceThreshold() (r float32) { r = b2Joint_GetForceThreshold(theStack, b.Id); return }
func (b Joint) SetTorqueThreshold(threshold float32) {
	b2Joint_SetTorqueThreshold(theStack, b.Id, threshold)
}
func (b Joint) GetTorqueThreshold() (r float32) {
	r = b2Joint_GetTorqueThreshold(theStack, b.Id)
	return
}
func CollideCircles(circleA Circle, circleB Circle, xf Transform) (r LocalManifold) {
	circleAStack := copyToStack(theStack, circleA)
	defer circleAStack.Free()
	circleBStack := copyToStack(theStack, circleB)
	defer circleBStack.Free()
	r = b2CollideCircles(theStack, circleAStack.Addr, circleBStack.Addr, xf)
	return
}
func CollideCapsuleAndCircle(capsuleA Capsule, circleB Circle, xf Transform) (r LocalManifold) {
	capsuleAStack := copyToStack(theStack, capsuleA)
	defer capsuleAStack.Free()
	circleBStack := copyToStack(theStack, circleB)
	defer circleBStack.Free()
	r = b2CollideCapsuleAndCircle(theStack, capsuleAStack.Addr, circleBStack.Addr, xf)
	return
}
func CollidePolygonAndCircle(polygonA Polygon, circleB Circle, xf Transform) (r LocalManifold) {
	polygonAStack := copyToStack(theStack, polygonA)
	defer polygonAStack.Free()
	circleBStack := copyToStack(theStack, circleB)
	defer circleBStack.Free()
	r = b2CollidePolygonAndCircle(theStack, polygonAStack.Addr, circleBStack.Addr, xf)
	return
}
func CollideCapsules(capsuleA Capsule, capsuleB Capsule, xf Transform) (r1 LocalManifold) {
	capsuleAStack := copyToStack(theStack, capsuleA)
	defer capsuleAStack.Free()
	capsuleBStack := copyToStack(theStack, capsuleB)
	defer capsuleBStack.Free()
	r1 = b2CollideCapsules(theStack, capsuleAStack.Addr, capsuleBStack.Addr, xf)
	return
}
func CollideSegmentAndCapsule(segmentA Segment, capsuleB Capsule, xf Transform) (r LocalManifold) {
	segmentAStack := copyToStack(theStack, segmentA)
	defer segmentAStack.Free()
	capsuleBStack := copyToStack(theStack, capsuleB)
	defer capsuleBStack.Free()
	r = b2CollideSegmentAndCapsule(theStack, segmentAStack.Addr, capsuleBStack.Addr, xf)
	return
}
func CollidePolygonAndCapsule(polygonA Polygon, capsuleB Capsule, xf Transform) (r LocalManifold) {
	polygonAStack := copyToStack(theStack, polygonA)
	defer polygonAStack.Free()
	capsuleBStack := copyToStack(theStack, capsuleB)
	defer capsuleBStack.Free()
	r = b2CollidePolygonAndCapsule(theStack, polygonAStack.Addr, capsuleBStack.Addr, xf)
	return
}
func CollidePolygons(polygonA Polygon, polygonB Polygon, xf Transform) (r LocalManifold) {
	polygonAStack := copyToStack(theStack, polygonA)
	defer polygonAStack.Free()
	polygonBStack := copyToStack(theStack, polygonB)
	defer polygonBStack.Free()
	r = b2CollidePolygons(theStack, polygonAStack.Addr, polygonBStack.Addr, xf)
	return
}
func CollideSegmentAndCircle(segmentA Segment, circleB Circle, xf Transform) (r LocalManifold) {
	segmentAStack := copyToStack(theStack, segmentA)
	defer segmentAStack.Free()
	circleBStack := copyToStack(theStack, circleB)
	defer circleBStack.Free()
	r = b2CollideSegmentAndCircle(theStack, segmentAStack.Addr, circleBStack.Addr, xf)
	return
}
func CollideSegmentAndPolygon(segmentA Segment, polygonB Polygon, xf Transform) (r LocalManifold) {
	segmentAStack := copyToStack(theStack, segmentA)
	defer segmentAStack.Free()
	polygonBStack := copyToStack(theStack, polygonB)
	defer polygonBStack.Free()
	r = b2CollideSegmentAndPolygon(theStack, segmentAStack.Addr, polygonBStack.Addr, xf)
	return
}
func CollideChainSegmentAndCircle(segmentA ChainSegment, circleB Circle, xf Transform) (r LocalManifold) {
	segmentAStack := copyToStack(theStack, segmentA)
	defer segmentAStack.Free()
	circleBStack := copyToStack(theStack, circleB)
	defer circleBStack.Free()
	r = b2CollideChainSegmentAndCircle(theStack, segmentAStack.Addr, circleBStack.Addr, xf)
	return
}
func CollideChainSegmentAndCapsule(segmentA ChainSegment, capsuleB Capsule, xf Transform, cache *SimplexCache) (r LocalManifold) {
	defer runtime.KeepAlive(cache)
	segmentAStack := copyToStack(theStack, segmentA)
	defer segmentAStack.Free()
	capsuleBStack := copyToStack(theStack, capsuleB)
	defer capsuleBStack.Free()
	escapes(cache)
	r = b2CollideChainSegmentAndCapsule(theStack, segmentAStack.Addr, capsuleBStack.Addr, xf, uintptr(unsafe.Pointer(cache)))
	return
}
func CollideChainSegmentAndPolygon(segmentA ChainSegment, polygonB Polygon, xf Transform, cache *SimplexCache) (r LocalManifold) {
	defer runtime.KeepAlive(cache)
	segmentAStack := copyToStack(theStack, segmentA)
	defer segmentAStack.Free()
	polygonBStack := copyToStack(theStack, polygonB)
	defer polygonBStack.Free()
	escapes(cache)
	r = b2CollideChainSegmentAndPolygon(theStack, segmentAStack.Addr, polygonBStack.Addr, xf, uintptr(unsafe.Pointer(cache)))
	return
}
func IsValidFloat(a float32) (r uint8)       { r = b2IsValidFloat(theStack, a); return }
func IsValidVec2(v Vec2) (r uint8)           { r = b2IsValidVec2(theStack, v); return }
func IsValidRot(q1 Rot) (r uint8)            { r = b2IsValidRot(theStack, q1); return }
func IsValidTransform(t Transform) (r uint8) { r = b2IsValidTransform(theStack, t); return }
func IsValidPlane(a3 Plane) (r uint8)        { r = b2IsValidPlane(theStack, a3); return }
func IsValidPos(p Pos) (r uint8)             { r = b2IsValidPos(theStack, p); return }
func IsValidWorldTransform(t WorldTransform) (r uint8) {
	r = b2IsValidWorldTransform(theStack, t)
	return
}
func Atan2(y float32, x float32) (r1 float32)   { r1 = b2Atan2(theStack, y, x); return }
func ComputeCosSin(radians1 float32) (r CosSin) { r = b2ComputeCosSin(theStack, radians1); return }
func ComputeRotationBetweenUnitVectors(v1 Vec2, v2 Vec2) (r Rot) {
	r = b2ComputeRotationBetweenUnitVectors(theStack, v1, v2)
	return
}
func (b MotorJoint) SetLinearVelocity(velocity Vec2) {
	b2MotorJoint_SetLinearVelocity(theStack, b.Id, velocity)
}
func (b MotorJoint) GetLinearVelocity() (r Vec2) {
	r = b2MotorJoint_GetLinearVelocity(theStack, b.Id)
	return
}
func (b MotorJoint) SetAngularVelocity(velocity float32) {
	b2MotorJoint_SetAngularVelocity(theStack, b.Id, velocity)
}
func (b MotorJoint) GetAngularVelocity() (r float32) {
	r = b2MotorJoint_GetAngularVelocity(theStack, b.Id)
	return
}
func (b MotorJoint) SetMaxVelocityTorque(maxTorque float32) {
	b2MotorJoint_SetMaxVelocityTorque(theStack, b.Id, maxTorque)
}
func (b MotorJoint) GetMaxVelocityTorque() (r float32) {
	r = b2MotorJoint_GetMaxVelocityTorque(theStack, b.Id)
	return
}
func (b MotorJoint) SetMaxVelocityForce(maxForce float32) {
	b2MotorJoint_SetMaxVelocityForce(theStack, b.Id, maxForce)
}
func (b MotorJoint) GetMaxVelocityForce() (r float32) {
	r = b2MotorJoint_GetMaxVelocityForce(theStack, b.Id)
	return
}
func (b MotorJoint) SetLinearHertz(hertz float32) { b2MotorJoint_SetLinearHertz(theStack, b.Id, hertz) }
func (b MotorJoint) GetLinearHertz() (r float32) {
	r = b2MotorJoint_GetLinearHertz(theStack, b.Id)
	return
}
func (b MotorJoint) SetLinearDampingRatio(damping float32) {
	b2MotorJoint_SetLinearDampingRatio(theStack, b.Id, damping)
}
func (b MotorJoint) GetLinearDampingRatio() (r float32) {
	r = b2MotorJoint_GetLinearDampingRatio(theStack, b.Id)
	return
}
func (b MotorJoint) SetAngularHertz(hertz float32) {
	b2MotorJoint_SetAngularHertz(theStack, b.Id, hertz)
}
func (b MotorJoint) GetAngularHertz() (r float32) {
	r = b2MotorJoint_GetAngularHertz(theStack, b.Id)
	return
}
func (b MotorJoint) SetAngularDampingRatio(damping float32) {
	b2MotorJoint_SetAngularDampingRatio(theStack, b.Id, damping)
}
func (b MotorJoint) GetAngularDampingRatio() (r float32) {
	r = b2MotorJoint_GetAngularDampingRatio(theStack, b.Id)
	return
}
func (b MotorJoint) SetMaxSpringForce(maxForce float32) {
	b2MotorJoint_SetMaxSpringForce(theStack, b.Id, maxForce)
}
func (b MotorJoint) GetMaxSpringForce() (r float32) {
	r = b2MotorJoint_GetMaxSpringForce(theStack, b.Id)
	return
}
func (b MotorJoint) SetMaxSpringTorque(maxTorque float32) {
	b2MotorJoint_SetMaxSpringTorque(theStack, b.Id, maxTorque)
}
func (b MotorJoint) GetMaxSpringTorque() (r float32) {
	r = b2MotorJoint_GetMaxSpringTorque(theStack, b.Id)
	return
}
func (b MoverJoint) SetLinearVelocity(velocity Vec2) {
	b2MoverJoint_SetLinearVelocity(theStack, b.Id, velocity)
}
func (b MoverJoint) GetLinearVelocity() (r Vec2) {
	r = b2MoverJoint_GetLinearVelocity(theStack, b.Id)
	return
}
func (b MoverJoint) SetMaxVelocityForce(maxForce Vec2) {
	b2MoverJoint_SetMaxVelocityForce(theStack, b.Id, maxForce)
}
func (b MoverJoint) GetMaxVelocityForce() (r Vec2) {
	r = b2MoverJoint_GetMaxVelocityForce(theStack, b.Id)
	return
}
func CreateWorld(def WorldDef) (r World) {
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	r.Id = b2CreateWorld(theStack, defStack.Addr)
	return
}
func (b World) Step(timeStep float32, subStepCount int32) {
	b2World_Step(theStack, b.Id, timeStep, subStepCount)
}
func (b World) GetBounds() (r AABB)           { r = b2World_GetBounds(theStack, b.Id); return }
func (b World) IsValid() (r uint8)            { r = b2World_IsValid(theStack, b.Id); return }
func (b Body) IsValid() (r uint8)             { r = b2Body_IsValid(theStack, b.Id); return }
func (b Shape) IsValid() (r uint8)            { r = b2Shape_IsValid(theStack, b.Id); return }
func (b Chain) IsValid() (r uint8)            { r = b2Chain_IsValid(theStack, b.Id); return }
func (b Joint) IsValid() (r uint8)            { r = b2Joint_IsValid(theStack, b.Id); return }
func (b Contact) IsValid() (r uint8)          { r = b2Contact_IsValid(theStack, b.Id); return }
func (b World) EnableSleeping(flag uint8)     { b2World_EnableSleeping(theStack, b.Id, flag) }
func (b World) IsSleepingEnabled() (r uint8)  { r = b2World_IsSleepingEnabled(theStack, b.Id); return }
func (b World) EnableWarmStarting(flag uint8) { b2World_EnableWarmStarting(theStack, b.Id, flag) }
func (b World) IsWarmStartingEnabled() (r uint8) {
	r = b2World_IsWarmStartingEnabled(theStack, b.Id)
	return
}
func (b World) EnableSSE2Fallback(flag uint8) { b2World_EnableSSE2Fallback(theStack, b.Id, flag) }
func (b World) GetAwakeBodyCount() (r int32)  { r = b2World_GetAwakeBodyCount(theStack, b.Id); return }
func (b World) EnableContinuous(flag uint8)   { b2World_EnableContinuous(theStack, b.Id, flag) }
func (b World) IsContinuousEnabled() (r uint8) {
	r = b2World_IsContinuousEnabled(theStack, b.Id)
	return
}
func (b World) SetRestitutionThreshold(value float32) {
	b2World_SetRestitutionThreshold(theStack, b.Id, value)
}
func (b World) GetRestitutionThreshold() (r float32) {
	r = b2World_GetRestitutionThreshold(theStack, b.Id)
	return
}
func (b World) SetRestitutionIterations(iterations int32) {
	b2World_SetRestitutionIterations(theStack, b.Id, iterations)
}
func (b World) GetRestitutionIterations() (r int32) {
	r = b2World_GetRestitutionIterations(theStack, b.Id)
	return
}
func (b World) EnableRestitutionPropagation(flag uint8) {
	b2World_EnableRestitutionPropagation(theStack, b.Id, flag)
}
func (b World) IsRestitutionPropagationEnabled() (r uint8) {
	r = b2World_IsRestitutionPropagationEnabled(theStack, b.Id)
	return
}
func (b World) SetHitEventThreshold(value float32) {
	b2World_SetHitEventThreshold(theStack, b.Id, value)
}
func (b World) GetHitEventThreshold() (r float32) {
	r = b2World_GetHitEventThreshold(theStack, b.Id)
	return
}
func (b World) SetContactTuning(hertz float32, dampingRatio float32, pushSpeed float32) {
	b2World_SetContactTuning(theStack, b.Id, hertz, dampingRatio, pushSpeed)
}
func (b World) SetContactRecycleDistance(recycleDistance float32) {
	b2World_SetContactRecycleDistance(theStack, b.Id, recycleDistance)
}
func (b World) GetContactRecycleDistance() (r float32) {
	r = b2World_GetContactRecycleDistance(theStack, b.Id)
	return
}
func (b World) SetMaxLinearSpeed(maximumLinearSpeed float32) {
	b2World_SetMaxLinearSpeed(theStack, b.Id, maximumLinearSpeed)
}
func (b World) GetMaxLinearSpeed() (r float32) { r = b2World_GetMaxLinearSpeed(theStack, b.Id); return }
func (b World) GetProfile() (r Profile)        { r = b2World_GetProfile(theStack, b.Id); return }
func (b World) GetCounters() (r Counters)      { r = b2World_GetCounters(theStack, b.Id); return }
func (b World) GetMaxCapacity() (r Capacity)   { r = b2World_GetMaxCapacity(theStack, b.Id); return }
func (b World) SetUserData(userData uintptr)   { b2World_SetUserData(theStack, b.Id, userData) }
func (b World) GetUserData() (r uintptr)       { r = b2World_GetUserData(theStack, b.Id); return }
func (b World) SetWorkerCount(count int32)     { b2World_SetWorkerCount(theStack, b.Id, count) }
func (b World) GetWorkerCount() (r int32)      { r = b2World_GetWorkerCount(theStack, b.Id); return }
func (b World) CastRayClosest(origin Pos, translation Vec2, filter QueryFilter) (r RayResult) {
	r = b2World_CastRayClosest(theStack, b.Id, origin, translation, filter)
	return
}
func (b World) CastMover(origin1 Pos, mover Capsule, translation Vec2, filter QueryFilter) (r1 float32) {
	moverStack := copyToStack(theStack, mover)
	defer moverStack.Free()
	r1 = b2World_CastMover(theStack, b.Id, origin1, moverStack.Addr, translation, filter)
	return
}
func (b World) SetGravity(gravity Vec2) { b2World_SetGravity(theStack, b.Id, gravity) }
func (b World) GetGravity() (r Vec2)    { r = b2World_GetGravity(theStack, b.Id); return }
func (b World) Explode(explosionDef ExplosionDef) {
	explosionDefStack := copyToStack(theStack, explosionDef)
	defer explosionDefStack.Free()
	b2World_Explode(theStack, b.Id, explosionDefStack.Addr)
}
func (b World) RebuildStaticTree()               { b2World_RebuildStaticTree(theStack, b.Id) }
func (b PogoJoint) SetRestLength(length float32) { b2PogoJoint_SetRestLength(theStack, b.Id, length) }
func (b PogoJoint) GetRestLength() (r float32)   { r = b2PogoJoint_GetRestLength(theStack, b.Id); return }
func (b PogoJoint) GetSpringHertz() (r float32) {
	r = b2PogoJoint_GetSpringHertz(theStack, b.Id)
	return
}
func (b PogoJoint) SetSpringHertz(hertz float32) { b2PogoJoint_SetSpringHertz(theStack, b.Id, hertz) }
func (b PogoJoint) SetSpringDampingRatio(dampingRatio float32) {
	b2PogoJoint_SetSpringDampingRatio(theStack, b.Id, dampingRatio)
}
func (b PogoJoint) GetSpringDampingRatio() (r float32) {
	r = b2PogoJoint_GetSpringDampingRatio(theStack, b.Id)
	return
}
func (b PogoJoint) GetCurrentLength() (r1 float32) {
	r1 = b2PogoJoint_GetCurrentLength(theStack, b.Id)
	return
}
func (b PogoJoint) GetImpulse() (r float32) { r = b2PogoJoint_GetImpulse(theStack, b.Id); return }
func (b PogoJoint) GetCurrentSpeed() (r float32) {
	r = b2PogoJoint_GetCurrentSpeed(theStack, b.Id)
	return
}
func (b PrismaticJoint) EnableSpring(enableSpring uint8) {
	b2PrismaticJoint_EnableSpring(theStack, b.Id, enableSpring)
}
func (b PrismaticJoint) IsSpringEnabled() (r uint8) {
	r = b2PrismaticJoint_IsSpringEnabled(theStack, b.Id)
	return
}
func (b PrismaticJoint) SetSpringHertz(hertz float32) {
	b2PrismaticJoint_SetSpringHertz(theStack, b.Id, hertz)
}
func (b PrismaticJoint) GetSpringHertz() (r float32) {
	r = b2PrismaticJoint_GetSpringHertz(theStack, b.Id)
	return
}
func (b PrismaticJoint) SetSpringDampingRatio(dampingRatio float32) {
	b2PrismaticJoint_SetSpringDampingRatio(theStack, b.Id, dampingRatio)
}
func (b PrismaticJoint) GetSpringDampingRatio() (r float32) {
	r = b2PrismaticJoint_GetSpringDampingRatio(theStack, b.Id)
	return
}
func (b PrismaticJoint) SetTargetTranslation(translation float32) {
	b2PrismaticJoint_SetTargetTranslation(theStack, b.Id, translation)
}
func (b PrismaticJoint) GetTargetTranslation() (r float32) {
	r = b2PrismaticJoint_GetTargetTranslation(theStack, b.Id)
	return
}
func (b PrismaticJoint) EnableLimit(enableLimit uint8) {
	b2PrismaticJoint_EnableLimit(theStack, b.Id, enableLimit)
}
func (b PrismaticJoint) IsLimitEnabled() (r uint8) {
	r = b2PrismaticJoint_IsLimitEnabled(theStack, b.Id)
	return
}
func (b PrismaticJoint) GetLowerLimit() (r float32) {
	r = b2PrismaticJoint_GetLowerLimit(theStack, b.Id)
	return
}
func (b PrismaticJoint) GetUpperLimit() (r float32) {
	r = b2PrismaticJoint_GetUpperLimit(theStack, b.Id)
	return
}
func (b PrismaticJoint) SetLimits(lower float32, upper float32) {
	b2PrismaticJoint_SetLimits(theStack, b.Id, lower, upper)
}
func (b PrismaticJoint) EnableMotor(enableMotor uint8) {
	b2PrismaticJoint_EnableMotor(theStack, b.Id, enableMotor)
}
func (b PrismaticJoint) IsMotorEnabled() (r uint8) {
	r = b2PrismaticJoint_IsMotorEnabled(theStack, b.Id)
	return
}
func (b PrismaticJoint) SetTargetMotorSpeed(motorSpeed float32) {
	b2PrismaticJoint_SetTargetMotorSpeed(theStack, b.Id, motorSpeed)
}
func (b PrismaticJoint) GetTargetMotorSpeed() (r float32) {
	r = b2PrismaticJoint_GetTargetMotorSpeed(theStack, b.Id)
	return
}
func (b PrismaticJoint) GetCurrentMotorForce() (r float32) {
	r = b2PrismaticJoint_GetCurrentMotorForce(theStack, b.Id)
	return
}
func (b PrismaticJoint) SetMaxMotorForce(force float32) {
	b2PrismaticJoint_SetMaxMotorForce(theStack, b.Id, force)
}
func (b PrismaticJoint) GetMaxMotorForce() (r float32) {
	r = b2PrismaticJoint_GetMaxMotorForce(theStack, b.Id)
	return
}
func (b PrismaticJoint) GetCurrentTranslation() (r1 float32) {
	r1 = b2PrismaticJoint_GetCurrentTranslation(theStack, b.Id)
	return
}
func (b PrismaticJoint) GetCurrentSpeed() (r float32) {
	r = b2PrismaticJoint_GetCurrentSpeed(theStack, b.Id)
	return
}
func (b RevoluteJoint) EnableSpring(enableSpring uint8) {
	b2RevoluteJoint_EnableSpring(theStack, b.Id, enableSpring)
}
func (b RevoluteJoint) IsSpringEnabled() (r uint8) {
	r = b2RevoluteJoint_IsSpringEnabled(theStack, b.Id)
	return
}
func (b RevoluteJoint) SetSpringHertz(hertz float32) {
	b2RevoluteJoint_SetSpringHertz(theStack, b.Id, hertz)
}
func (b RevoluteJoint) GetSpringHertz() (r float32) {
	r = b2RevoluteJoint_GetSpringHertz(theStack, b.Id)
	return
}
func (b RevoluteJoint) SetSpringDampingRatio(dampingRatio float32) {
	b2RevoluteJoint_SetSpringDampingRatio(theStack, b.Id, dampingRatio)
}
func (b RevoluteJoint) GetSpringDampingRatio() (r float32) {
	r = b2RevoluteJoint_GetSpringDampingRatio(theStack, b.Id)
	return
}
func (b RevoluteJoint) SetTargetAngle(angle float32) {
	b2RevoluteJoint_SetTargetAngle(theStack, b.Id, angle)
}
func (b RevoluteJoint) GetTargetAngle() (r float32) {
	r = b2RevoluteJoint_GetTargetAngle(theStack, b.Id)
	return
}
func (b RevoluteJoint) EnableLimit(enableLimit uint8) {
	b2RevoluteJoint_EnableLimit(theStack, b.Id, enableLimit)
}
func (b RevoluteJoint) IsLimitEnabled() (r uint8) {
	r = b2RevoluteJoint_IsLimitEnabled(theStack, b.Id)
	return
}
func (b RevoluteJoint) GetLowerLimit() (r float32) {
	r = b2RevoluteJoint_GetLowerLimit(theStack, b.Id)
	return
}
func (b RevoluteJoint) GetUpperLimit() (r float32) {
	r = b2RevoluteJoint_GetUpperLimit(theStack, b.Id)
	return
}
func (b RevoluteJoint) SetLimits(lower1 float32, upper1 float32) {
	b2RevoluteJoint_SetLimits(theStack, b.Id, lower1, upper1)
}
func (b RevoluteJoint) EnableMotor(enableMotor uint8) {
	b2RevoluteJoint_EnableMotor(theStack, b.Id, enableMotor)
}
func (b RevoluteJoint) IsMotorEnabled() (r uint8) {
	r = b2RevoluteJoint_IsMotorEnabled(theStack, b.Id)
	return
}
func (b RevoluteJoint) SetTargetMotorSpeed(motorSpeed float32) {
	b2RevoluteJoint_SetTargetMotorSpeed(theStack, b.Id, motorSpeed)
}
func (b RevoluteJoint) GetTargetMotorSpeed() (r float32) {
	r = b2RevoluteJoint_GetTargetMotorSpeed(theStack, b.Id)
	return
}
func (b RevoluteJoint) SetMaxMotorTorque(torque float32) {
	b2RevoluteJoint_SetMaxMotorTorque(theStack, b.Id, torque)
}
func (b RevoluteJoint) GetMaxMotorTorque() (r float32) {
	r = b2RevoluteJoint_GetMaxMotorTorque(theStack, b.Id)
	return
}
func (b RevoluteJoint) GetCurrentAngle() (r float32) {
	r = b2RevoluteJoint_GetCurrentAngle(theStack, b.Id)
	return
}
func (b RevoluteJoint) GetCurrentSpeed() (r float32) {
	r = b2RevoluteJoint_GetCurrentSpeed(theStack, b.Id)
	return
}
func (b RevoluteJoint) GetCurrentMotorTorque() (r float32) {
	r = b2RevoluteJoint_GetCurrentMotorTorque(theStack, b.Id)
	return
}
func (b Body) CreateCircleShape(def ShapeDef, circle Circle) (r Shape) {
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	circleStack := copyToStack(theStack, circle)
	defer circleStack.Free()
	r.Id = b2CreateCircleShape(theStack, b.Id, defStack.Addr, circleStack.Addr)
	return
}
func (b Body) CreateCapsuleShape(def ShapeDef, capsule Capsule) (r Shape) {
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	capsuleStack := copyToStack(theStack, capsule)
	defer capsuleStack.Free()
	r.Id = b2CreateCapsuleShape(theStack, b.Id, defStack.Addr, capsuleStack.Addr)
	return
}
func (b Body) CreatePolygonShape(def ShapeDef, polygon Polygon) (r Shape) {
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	polygonStack := copyToStack(theStack, polygon)
	defer polygonStack.Free()
	r.Id = b2CreatePolygonShape(theStack, b.Id, defStack.Addr, polygonStack.Addr)
	return
}
func (b Body) CreateSegmentShape(def ShapeDef, segment Segment) (r Shape) {
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	segmentStack := copyToStack(theStack, segment)
	defer segmentStack.Free()
	r.Id = b2CreateSegmentShape(theStack, b.Id, defStack.Addr, segmentStack.Addr)
	return
}
func (b Body) CreateChainSegmentShape(def ShapeDef, chainSegment ChainSegment) (r Shape) {
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	chainSegmentStack := copyToStack(theStack, chainSegment)
	defer chainSegmentStack.Free()
	r.Id = b2CreateChainSegmentShape(theStack, b.Id, defStack.Addr, chainSegmentStack.Addr)
	return
}
func (b Shape) DestroyShape(updateBodyMass uint8) { b2DestroyShape(theStack, b.Id, updateBodyMass) }
func (b Chain) DestroyChain()                     { b2DestroyChain(theStack, b.Id) }
func (b Chain) GetWorld() (r World)               { r.Id = b2Chain_GetWorld(theStack, b.Id); return }
func (b Chain) GetSegmentCount() (r int32)        { r = b2Chain_GetSegmentCount(theStack, b.Id); return }
func (b Chain) SetSurfaceMaterial(material SurfaceMaterial, segmentIndex int32) {
	materialStack := copyToStack(theStack, material)
	defer materialStack.Free()
	b2Chain_SetSurfaceMaterial(theStack, b.Id, materialStack.Addr, segmentIndex)
}
func (b Chain) SetAllSurfaceMaterials(material SurfaceMaterial) {
	materialStack := copyToStack(theStack, material)
	defer materialStack.Free()
	b2Chain_SetAllSurfaceMaterials(theStack, b.Id, materialStack.Addr)
}
func (b Chain) GetSurfaceMaterial(segmentIndex int32) (r SurfaceMaterial) {
	r = b2Chain_GetSurfaceMaterial(theStack, b.Id, segmentIndex)
	return
}
func (b Shape) GetBody() (r Body)             { r.Id = b2Shape_GetBody(theStack, b.Id); return }
func (b Shape) GetWorld() (r World)           { r.Id = b2Shape_GetWorld(theStack, b.Id); return }
func (b Shape) SetUserData(userData uintptr)  { b2Shape_SetUserData(theStack, b.Id, userData) }
func (b Shape) GetUserData() (r uintptr)      { r = b2Shape_GetUserData(theStack, b.Id); return }
func (b Shape) IsSensor() (r uint8)           { r = b2Shape_IsSensor(theStack, b.Id); return }
func (b Shape) TestPoint(point Pos) (r uint8) { r = b2Shape_TestPoint(theStack, b.Id, point); return }
func (b Shape) CastRay(origin Pos, translation Vec2) (r1 WorldCastOutput) {
	r1 = b2Shape_CastRay(theStack, b.Id, origin, translation)
	return
}
func (b Shape) SetDensity(density float32, updateBodyMass uint8) {
	b2Shape_SetDensity(theStack, b.Id, density, updateBodyMass)
}
func (b Shape) GetDensity() (r float32)      { r = b2Shape_GetDensity(theStack, b.Id); return }
func (b Shape) SetFriction(friction float32) { b2Shape_SetFriction(theStack, b.Id, friction) }
func (b Shape) GetFriction() (r float32)     { r = b2Shape_GetFriction(theStack, b.Id); return }
func (b Shape) SetRestitution(restitution float32) {
	b2Shape_SetRestitution(theStack, b.Id, restitution)
}
func (b Shape) GetRestitution() (r float32)     { r = b2Shape_GetRestitution(theStack, b.Id); return }
func (b Shape) SetUserMaterial(material uint64) { b2Shape_SetUserMaterial(theStack, b.Id, material) }
func (b Shape) GetUserMaterial() (r uint64)     { r = b2Shape_GetUserMaterial(theStack, b.Id); return }
func (b Shape) GetSurfaceMaterial() (r SurfaceMaterial) {
	r = b2Shape_GetSurfaceMaterial(theStack, b.Id)
	return
}
func (b Shape) SetSurfaceMaterial(surfaceMaterial SurfaceMaterial) {
	surfaceMaterialStack := copyToStack(theStack, surfaceMaterial)
	defer surfaceMaterialStack.Free()
	b2Shape_SetSurfaceMaterial(theStack, b.Id, surfaceMaterialStack.Addr)
}
func (b Shape) GetFilter() (r Filter)         { r = b2Shape_GetFilter(theStack, b.Id); return }
func (b Shape) SetFilter(filter Filter)       { b2Shape_SetFilter(theStack, b.Id, filter) }
func (b Shape) EnableSensorEvents(flag uint8) { b2Shape_EnableSensorEvents(theStack, b.Id, flag) }
func (b Shape) AreSensorEventsEnabled() (r uint8) {
	r = b2Shape_AreSensorEventsEnabled(theStack, b.Id)
	return
}
func (b Shape) EnableContactEvents(flag uint8) { b2Shape_EnableContactEvents(theStack, b.Id, flag) }
func (b Shape) AreContactEventsEnabled() (r uint8) {
	r = b2Shape_AreContactEventsEnabled(theStack, b.Id)
	return
}
func (b Shape) EnablePreSolveEvents(flag uint8) { b2Shape_EnablePreSolveEvents(theStack, b.Id, flag) }
func (b Shape) ArePreSolveEventsEnabled() (r uint8) {
	r = b2Shape_ArePreSolveEventsEnabled(theStack, b.Id)
	return
}
func (b Shape) EnableHitEvents(flag uint8) { b2Shape_EnableHitEvents(theStack, b.Id, flag) }
func (b Shape) AreHitEventsEnabled() (r uint8) {
	r = b2Shape_AreHitEventsEnabled(theStack, b.Id)
	return
}
func (b Shape) GetType() (r ShapeType)  { r = b2Shape_GetType(theStack, b.Id); return }
func (b Shape) GetCircle() (r Circle)   { r = b2Shape_GetCircle(theStack, b.Id); return }
func (b Shape) GetSegment() (r Segment) { r = b2Shape_GetSegment(theStack, b.Id); return }
func (b Shape) GetChainSegment() (r ChainSegment) {
	r = b2Shape_GetChainSegment(theStack, b.Id)
	return
}
func (b Shape) GetCapsule() (r Capsule) { r = b2Shape_GetCapsule(theStack, b.Id); return }
func (b Shape) GetPolygon() (r Polygon) { r = b2Shape_GetPolygon(theStack, b.Id); return }
func (b Shape) SetCircle(circle Circle) {
	circleStack := copyToStack(theStack, circle)
	defer circleStack.Free()
	b2Shape_SetCircle(theStack, b.Id, circleStack.Addr)
}
func (b Shape) SetCapsule(capsule Capsule) {
	capsuleStack := copyToStack(theStack, capsule)
	defer capsuleStack.Free()
	b2Shape_SetCapsule(theStack, b.Id, capsuleStack.Addr)
}
func (b Shape) SetSegment(segment Segment) {
	segmentStack := copyToStack(theStack, segment)
	defer segmentStack.Free()
	b2Shape_SetSegment(theStack, b.Id, segmentStack.Addr)
}
func (b Shape) SetPolygon(polygon Polygon) {
	polygonStack := copyToStack(theStack, polygon)
	defer polygonStack.Free()
	b2Shape_SetPolygon(theStack, b.Id, polygonStack.Addr)
}
func (b Shape) SetChainSegment(chainSegment ChainSegment) {
	chainSegmentStack := copyToStack(theStack, chainSegment)
	defer chainSegmentStack.Free()
	b2Shape_SetChainSegment(theStack, b.Id, chainSegmentStack.Addr)
}
func (b Shape) GetParentChain() (r Chain)     { r.Id = b2Shape_GetParentChain(theStack, b.Id); return }
func (b Shape) GetContactCapacity() (r int32) { r = b2Shape_GetContactCapacity(theStack, b.Id); return }
func (b Shape) GetSensorCapacity() (r int32)  { r = b2Shape_GetSensorCapacity(theStack, b.Id); return }
func (b Shape) GetBounds() (r AABB)           { r = b2Shape_GetBounds(theStack, b.Id); return }
func (b Shape) ComputeMassData() (r MassData) { r = b2Shape_ComputeMassData(theStack, b.Id); return }
func (b Shape) GetClosestPoint(target Pos) (r1 Pos) {
	r1 = b2Shape_GetClosestPoint(theStack, b.Id, target)
	return
}
func (b Shape) ApplyWind(wind Vec2, drag float32, lift float32, wake uint8) {
	b2Shape_ApplyWind(theStack, b.Id, wind, drag, lift, wake)
}
func IsAVX2Available() (r uint8)                  { r = b2IsAVX2Available(theStack); return }
func (b World) GetStateHash() (r uint64)          { r = b2World_GetStateHash(theStack, b.Id); return }
func GetTicks() (r uint64)                        { r = b2GetTicks(theStack); return }
func GetMilliseconds(ticks uint64) (r float32)    { r = b2GetMilliseconds(theStack, ticks); return }
func Yield()                                      { b2Yield(theStack) }
func DefaultWorldDef() (r WorldDef)               { r = b2DefaultWorldDef(theStack); return }
func DefaultBodyDef() (r BodyDef)                 { r = b2DefaultBodyDef(theStack); return }
func DefaultFilter() (r Filter)                   { r = b2DefaultFilter(theStack); return }
func DefaultQueryFilter() (r QueryFilter)         { r = b2DefaultQueryFilter(theStack); return }
func DefaultShapeDef() (r ShapeDef)               { r = b2DefaultShapeDef(theStack); return }
func DefaultSurfaceMaterial() (r SurfaceMaterial) { r = b2DefaultSurfaceMaterial(theStack); return }
func DefaultChainDef() (r ChainDef)               { r = b2DefaultChainDef(theStack); return }
func (b WeldJoint) SetLinearHertz(hertz float32)  { b2WeldJoint_SetLinearHertz(theStack, b.Id, hertz) }
func (b WeldJoint) GetLinearHertz() (r float32) {
	r = b2WeldJoint_GetLinearHertz(theStack, b.Id)
	return
}
func (b WeldJoint) SetLinearDampingRatio(dampingRatio float32) {
	b2WeldJoint_SetLinearDampingRatio(theStack, b.Id, dampingRatio)
}
func (b WeldJoint) GetLinearDampingRatio() (r float32) {
	r = b2WeldJoint_GetLinearDampingRatio(theStack, b.Id)
	return
}
func (b WeldJoint) SetAngularHertz(hertz float32) { b2WeldJoint_SetAngularHertz(theStack, b.Id, hertz) }
func (b WeldJoint) GetAngularHertz() (r float32) {
	r = b2WeldJoint_GetAngularHertz(theStack, b.Id)
	return
}
func (b WeldJoint) SetAngularDampingRatio(dampingRatio float32) {
	b2WeldJoint_SetAngularDampingRatio(theStack, b.Id, dampingRatio)
}
func (b WeldJoint) GetAngularDampingRatio() (r float32) {
	r = b2WeldJoint_GetAngularDampingRatio(theStack, b.Id)
	return
}
func (b WheelJoint) EnableSpring(enableSpring uint8) {
	b2WheelJoint_EnableSpring(theStack, b.Id, enableSpring)
}
func (b WheelJoint) IsSpringEnabled() (r uint8) {
	r = b2WheelJoint_IsSpringEnabled(theStack, b.Id)
	return
}
func (b WheelJoint) SetSpringHertz(hertz float32) { b2WheelJoint_SetSpringHertz(theStack, b.Id, hertz) }
func (b WheelJoint) GetSpringHertz() (r float32) {
	r = b2WheelJoint_GetSpringHertz(theStack, b.Id)
	return
}
func (b WheelJoint) SetSpringDampingRatio(dampingRatio float32) {
	b2WheelJoint_SetSpringDampingRatio(theStack, b.Id, dampingRatio)
}
func (b WheelJoint) GetSpringDampingRatio() (r float32) {
	r = b2WheelJoint_GetSpringDampingRatio(theStack, b.Id)
	return
}
func (b WheelJoint) EnableLimit(enableLimit uint8) {
	b2WheelJoint_EnableLimit(theStack, b.Id, enableLimit)
}
func (b WheelJoint) IsLimitEnabled() (r uint8) {
	r = b2WheelJoint_IsLimitEnabled(theStack, b.Id)
	return
}
func (b WheelJoint) GetLowerLimit() (r float32) {
	r = b2WheelJoint_GetLowerLimit(theStack, b.Id)
	return
}
func (b WheelJoint) GetUpperLimit() (r float32) {
	r = b2WheelJoint_GetUpperLimit(theStack, b.Id)
	return
}
func (b WheelJoint) SetLimits(lower float32, upper float32) {
	b2WheelJoint_SetLimits(theStack, b.Id, lower, upper)
}
func (b WheelJoint) EnableMotor(enableMotor uint8) {
	b2WheelJoint_EnableMotor(theStack, b.Id, enableMotor)
}
func (b WheelJoint) IsMotorEnabled() (r uint8) {
	r = b2WheelJoint_IsMotorEnabled(theStack, b.Id)
	return
}
func (b WheelJoint) SetTargetMotorSpeed(motorSpeed float32) {
	b2WheelJoint_SetTargetMotorSpeed(theStack, b.Id, motorSpeed)
}
func (b WheelJoint) GetTargetMotorSpeed() (r float32) {
	r = b2WheelJoint_GetTargetMotorSpeed(theStack, b.Id)
	return
}
func (b WheelJoint) GetCurrentMotorTorque() (r float32) {
	r = b2WheelJoint_GetCurrentMotorTorque(theStack, b.Id)
	return
}
func (b WheelJoint) SetMaxMotorTorque(torque float32) {
	b2WheelJoint_SetMaxMotorTorque(theStack, b.Id, torque)
}
func (b WheelJoint) GetMaxMotorTorque() (r float32) {
	r = b2WheelJoint_GetMaxMotorTorque(theStack, b.Id)
	return
}

const ToiStateUnknown TOIState = b2_toiStateUnknown
const ToiStateFailed TOIState = b2_toiStateFailed
const ToiStateOverlapped TOIState = b2_toiStateOverlapped
const ToiStateHit TOIState = b2_toiStateHit
const ToiStateSeparated TOIState = b2_toiStateSeparated
const StaticBody BodyType = b2_staticBody
const KinematicBody BodyType = b2_kinematicBody
const DynamicBody BodyType = b2_dynamicBody
const CircleShape ShapeType = b2_circleShape
const CapsuleShape ShapeType = b2_capsuleShape
const SegmentShape ShapeType = b2_segmentShape
const PolygonShape ShapeType = b2_polygonShape
const ChainSegmentShape ShapeType = b2_chainSegmentShape
const JointTypeDistance JointType = b2_distanceJoint
const JointTypeFilter JointType = b2_filterJoint
const JointTypeMotor JointType = b2_motorJoint
const JointTypeMover JointType = b2_moverJoint
const JointTypePogo JointType = b2_pogoJoint
const JointTypePrismatic JointType = b2_prismaticJoint
const JointTypeRevolute JointType = b2_revoluteJoint
const JointTypeWeld JointType = b2_weldJoint
const JointTypeWheel JointType = b2_wheelJoint
const ColorAliceBlue HexColor = b2_colorAliceBlue
const ColorAntiqueWhite HexColor = b2_colorAntiqueWhite
const ColorAqua HexColor = b2_colorAqua
const ColorAquamarine HexColor = b2_colorAquamarine
const ColorAzure HexColor = b2_colorAzure
const ColorBeige HexColor = b2_colorBeige
const ColorBisque HexColor = b2_colorBisque
const ColorBlack HexColor = b2_colorBlack
const ColorBlanchedAlmond HexColor = b2_colorBlanchedAlmond
const ColorBlue HexColor = b2_colorBlue
const ColorBlueViolet HexColor = b2_colorBlueViolet
const ColorBrown HexColor = b2_colorBrown
const ColorBurlywood HexColor = b2_colorBurlywood
const ColorCadetBlue HexColor = b2_colorCadetBlue
const ColorChartreuse HexColor = b2_colorChartreuse
const ColorChocolate HexColor = b2_colorChocolate
const ColorCoral HexColor = b2_colorCoral
const ColorCornflowerBlue HexColor = b2_colorCornflowerBlue
const ColorCornsilk HexColor = b2_colorCornsilk
const ColorCrimson HexColor = b2_colorCrimson
const ColorCyan HexColor = b2_colorCyan
const ColorDarkBlue HexColor = b2_colorDarkBlue
const ColorDarkCyan HexColor = b2_colorDarkCyan
const ColorDarkGoldenRod HexColor = b2_colorDarkGoldenRod
const ColorDarkGray HexColor = b2_colorDarkGray
const ColorDarkGreen HexColor = b2_colorDarkGreen
const ColorDarkKhaki HexColor = b2_colorDarkKhaki
const ColorDarkMagenta HexColor = b2_colorDarkMagenta
const ColorDarkOliveGreen HexColor = b2_colorDarkOliveGreen
const ColorDarkOrange HexColor = b2_colorDarkOrange
const ColorDarkOrchid HexColor = b2_colorDarkOrchid
const ColorDarkRed HexColor = b2_colorDarkRed
const ColorDarkSalmon HexColor = b2_colorDarkSalmon
const ColorDarkSeaGreen HexColor = b2_colorDarkSeaGreen
const ColorDarkSlateBlue HexColor = b2_colorDarkSlateBlue
const ColorDarkSlateGray HexColor = b2_colorDarkSlateGray
const ColorDarkTurquoise HexColor = b2_colorDarkTurquoise
const ColorDarkViolet HexColor = b2_colorDarkViolet
const ColorDeepPink HexColor = b2_colorDeepPink
const ColorDeepSkyBlue HexColor = b2_colorDeepSkyBlue
const ColorDimGray HexColor = b2_colorDimGray
const ColorDodgerBlue HexColor = b2_colorDodgerBlue
const ColorFireBrick HexColor = b2_colorFireBrick
const ColorFloralWhite HexColor = b2_colorFloralWhite
const ColorForestGreen HexColor = b2_colorForestGreen
const ColorFuchsia HexColor = b2_colorFuchsia
const ColorGainsboro HexColor = b2_colorGainsboro
const ColorGhostWhite HexColor = b2_colorGhostWhite
const ColorGold HexColor = b2_colorGold
const ColorGoldenRod HexColor = b2_colorGoldenRod
const ColorGray HexColor = b2_colorGray
const ColorGreen HexColor = b2_colorGreen
const ColorGreenYellow HexColor = b2_colorGreenYellow
const ColorHoneyDew HexColor = b2_colorHoneyDew
const ColorHotPink HexColor = b2_colorHotPink
const ColorIndianRed HexColor = b2_colorIndianRed
const ColorIndigo HexColor = b2_colorIndigo
const ColorIvory HexColor = b2_colorIvory
const ColorKhaki HexColor = b2_colorKhaki
const ColorLavender HexColor = b2_colorLavender
const ColorLavenderBlush HexColor = b2_colorLavenderBlush
const ColorLawnGreen HexColor = b2_colorLawnGreen
const ColorLemonChiffon HexColor = b2_colorLemonChiffon
const ColorLightBlue HexColor = b2_colorLightBlue
const ColorLightCoral HexColor = b2_colorLightCoral
const ColorLightCyan HexColor = b2_colorLightCyan
const ColorLightGoldenRodYellow HexColor = b2_colorLightGoldenRodYellow
const ColorLightGray HexColor = b2_colorLightGray
const ColorLightGreen HexColor = b2_colorLightGreen
const ColorLightPink HexColor = b2_colorLightPink
const ColorLightSalmon HexColor = b2_colorLightSalmon
const ColorLightSeaGreen HexColor = b2_colorLightSeaGreen
const ColorLightSkyBlue HexColor = b2_colorLightSkyBlue
const ColorLightSlateGray HexColor = b2_colorLightSlateGray
const ColorLightSteelBlue HexColor = b2_colorLightSteelBlue
const ColorLightYellow HexColor = b2_colorLightYellow
const ColorLime HexColor = b2_colorLime
const ColorLimeGreen HexColor = b2_colorLimeGreen
const ColorLinen HexColor = b2_colorLinen
const ColorMagenta HexColor = b2_colorMagenta
const ColorMaroon HexColor = b2_colorMaroon
const ColorMediumAquaMarine HexColor = b2_colorMediumAquaMarine
const ColorMediumBlue HexColor = b2_colorMediumBlue
const ColorMediumOrchid HexColor = b2_colorMediumOrchid
const ColorMediumPurple HexColor = b2_colorMediumPurple
const ColorMediumSeaGreen HexColor = b2_colorMediumSeaGreen
const ColorMediumSlateBlue HexColor = b2_colorMediumSlateBlue
const ColorMediumSpringGreen HexColor = b2_colorMediumSpringGreen
const ColorMediumTurquoise HexColor = b2_colorMediumTurquoise
const ColorMediumVioletRed HexColor = b2_colorMediumVioletRed
const ColorMidnightBlue HexColor = b2_colorMidnightBlue
const ColorMintCream HexColor = b2_colorMintCream
const ColorMistyRose HexColor = b2_colorMistyRose
const ColorMoccasin HexColor = b2_colorMoccasin
const ColorNavajoWhite HexColor = b2_colorNavajoWhite
const ColorNavy HexColor = b2_colorNavy
const ColorOldLace HexColor = b2_colorOldLace
const ColorOlive HexColor = b2_colorOlive
const ColorOliveDrab HexColor = b2_colorOliveDrab
const ColorOrange HexColor = b2_colorOrange
const ColorOrangeRed HexColor = b2_colorOrangeRed
const ColorOrchid HexColor = b2_colorOrchid
const ColorPaleGoldenRod HexColor = b2_colorPaleGoldenRod
const ColorPaleGreen HexColor = b2_colorPaleGreen
const ColorPaleTurquoise HexColor = b2_colorPaleTurquoise
const ColorPaleVioletRed HexColor = b2_colorPaleVioletRed
const ColorPapayaWhip HexColor = b2_colorPapayaWhip
const ColorPeachPuff HexColor = b2_colorPeachPuff
const ColorPeru HexColor = b2_colorPeru
const ColorPink HexColor = b2_colorPink
const ColorPlum HexColor = b2_colorPlum
const ColorPowderBlue HexColor = b2_colorPowderBlue
const ColorPurple HexColor = b2_colorPurple
const ColorRebeccaPurple HexColor = b2_colorRebeccaPurple
const ColorRed HexColor = b2_colorRed
const ColorRosyBrown HexColor = b2_colorRosyBrown
const ColorRoyalBlue HexColor = b2_colorRoyalBlue
const ColorSaddleBrown HexColor = b2_colorSaddleBrown
const ColorSalmon HexColor = b2_colorSalmon
const ColorSandyBrown HexColor = b2_colorSandyBrown
const ColorSeaGreen HexColor = b2_colorSeaGreen
const ColorSeaShell HexColor = b2_colorSeaShell
const ColorSienna HexColor = b2_colorSienna
const ColorSilver HexColor = b2_colorSilver
const ColorSkyBlue HexColor = b2_colorSkyBlue
const ColorSlateBlue HexColor = b2_colorSlateBlue
const ColorSlateGray HexColor = b2_colorSlateGray
const ColorSnow HexColor = b2_colorSnow
const ColorSpringGreen HexColor = b2_colorSpringGreen
const ColorSteelBlue HexColor = b2_colorSteelBlue
const ColorTan HexColor = b2_colorTan
const ColorTeal HexColor = b2_colorTeal
const ColorThistle HexColor = b2_colorThistle
const ColorTomato HexColor = b2_colorTomato
const ColorTurquoise HexColor = b2_colorTurquoise
const ColorViolet HexColor = b2_colorViolet
const ColorWheat HexColor = b2_colorWheat
const ColorWhite HexColor = b2_colorWhite
const ColorWhiteSmoke HexColor = b2_colorWhiteSmoke
const ColorYellow HexColor = b2_colorYellow
const ColorYellowGreen HexColor = b2_colorYellowGreen
const ColorBox2DRed HexColor = b2_colorBox2DRed
const ColorBox2DBlue HexColor = b2_colorBox2DBlue
const ColorBox2DGreen HexColor = b2_colorBox2DGreen
const ColorBox2DYellow HexColor = b2_colorBox2DYellow
