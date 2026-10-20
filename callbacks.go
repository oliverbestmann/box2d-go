package b2

// Go signatures of the box2d callbacks, as ccgo translates them from transpile/csrc/callbacks.c.
// box2d only sees a uintptr, so every Go function passed to it is converted to one of these
// types. If box2d changes a callback, the build breaks instead of memory getting corrupted.
type (
	threadFunction        = func(*_Stack, uintptr)
	frictionCallback      = func(*_Stack, float32, uint64, float32, uint64) float32
	restitutionCallback   = func(*_Stack, float32, uint64, float32, uint64) float32
	customFilterCallback  = func(*_Stack, ShapeId, ShapeId, uintptr) uint8
	preSolveCallback      = func(*_Stack, ShapeId, ShapeId, uintptr, uintptr)
	preContinuousCallback = func(*_Stack, ShapeId, ShapeId, Pos, Vec2, uintptr) uint8
	overlapResultCallback = func(*_Stack, ShapeId, uintptr) uint8
	castResultCallback    = func(*_Stack, ShapeId, Pos, Vec2, float32, uintptr) float32
	planeResultCallback   = func(*_Stack, ShapeId, uintptr, uintptr) uint8
	assertCallback        = func(*_Stack, uintptr, uintptr, int32) int32
	logCallback           = func(*_Stack, uintptr)
	drawPolygonFcn        = func(*_Stack, WorldTransform, uintptr, int32, HexColor, uintptr)
	drawSolidPolygonFcn   = func(*_Stack, WorldTransform, uintptr, int32, float32, HexColor, uintptr)
	drawCircleFcn         = func(*_Stack, Pos, float32, HexColor, uintptr)
	drawSolidCircleFcn    = func(*_Stack, WorldTransform, Vec2, float32, HexColor, uintptr)
	drawSolidCapsuleFcn   = func(*_Stack, Pos, Pos, float32, HexColor, uintptr)
	drawLineFcn           = func(*_Stack, Pos, Pos, HexColor, uintptr)
	drawTransformFcn      = func(*_Stack, WorldTransform, uintptr)
	drawPointFcn          = func(*_Stack, Pos, float32, HexColor, uintptr)
	drawStringFcn         = func(*_Stack, Pos, uintptr, HexColor, uintptr)
	drawBoundsFcn         = func(*_Stack, AABB, HexColor, uintptr)
)

var (
	_ threadFunction        = b2go_ThreadFunction
	_ frictionCallback      = b2go_FrictionCallbackFcn
	_ restitutionCallback   = b2go_RestitutionCallbackFcn
	_ customFilterCallback  = b2go_CustomFilterCallbackFcn
	_ preSolveCallback      = b2go_PreSolveCallbackFcn
	_ preContinuousCallback = b2go_PreContinuousCallbackFcn
	_ overlapResultCallback = b2go_OverlapResultCallbackFcn
	_ castResultCallback    = b2go_CastResultCallbackFcn
	_ planeResultCallback   = b2go_PlaneResultCallbackFcn
	_ assertCallback        = b2go_AssertCallbackFcn
	_ logCallback           = b2go_LogCallbackFcn
	_ drawPolygonFcn        = b2go_DrawPolygonFcn
	_ drawSolidPolygonFcn   = b2go_DrawSolidPolygonFcn
	_ drawCircleFcn         = b2go_DrawCircleFcn
	_ drawSolidCircleFcn    = b2go_DrawSolidCircleFcn
	_ drawSolidCapsuleFcn   = b2go_DrawSolidCapsuleFcn
	_ drawLineFcn           = b2go_DrawLineFcn
	_ drawTransformFcn      = b2go_DrawTransformFcn
	_ drawPointFcn          = b2go_DrawPointFcn
	_ drawStringFcn         = b2go_DrawStringFcn
	_ drawBoundsFcn         = b2go_DrawBoundsFcn
)
