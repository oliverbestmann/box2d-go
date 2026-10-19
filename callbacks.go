package b2

// Go signatures of the box2d callbacks, as ccgo translates them from transpile/csrc/callbacks.c.
// box2d only sees a uintptr, so every Go function passed to it is converted to one of these
// types. If box2d changes a callback, the build breaks instead of memory getting corrupted.
type (
	taskCallback        = func(*_Stack, int32, int32, uint32, uintptr)
	enqueueTaskCallback = func(*_Stack, uintptr, int32, int32, uintptr, uintptr) uintptr
	finishTaskCallback  = func(*_Stack, uintptr, uintptr)
	frictionCallback    = func(*_Stack, float32, int32, float32, int32) float32
	restitutionCallback = func(*_Stack, float32, int32, float32, int32) float32
	customFilterFcn     = func(*_Stack, ShapeId, ShapeId, uintptr) uint8
	preSolveFcn         = func(*_Stack, ShapeId, ShapeId, uintptr, uintptr) uint8
	overlapResultFcn    = func(*_Stack, ShapeId, uintptr) uint8
	castResultFcn       = func(*_Stack, ShapeId, Vec2, Vec2, float32, uintptr) float32
	assertFcn           = func(*_Stack, uintptr, uintptr, int32) int32
	drawPolygonFcn      = func(*_Stack, uintptr, int32, HexColor, uintptr)
	drawSolidPolygonFcn = func(*_Stack, Transform, uintptr, int32, float32, HexColor, uintptr)
	drawCircleFcn       = func(*_Stack, Vec2, float32, HexColor, uintptr)
	drawSolidCircleFcn  = func(*_Stack, Transform, float32, HexColor, uintptr)
	drawSolidCapsuleFcn = func(*_Stack, Vec2, Vec2, float32, HexColor, uintptr)
	drawSegmentFcn      = func(*_Stack, Vec2, Vec2, HexColor, uintptr)
	drawTransformFcn    = func(*_Stack, Transform, uintptr)
	drawPointFcn        = func(*_Stack, Vec2, float32, HexColor, uintptr)
	drawStringFcn       = func(*_Stack, Vec2, uintptr, HexColor, uintptr)
)

var (
	_ taskCallback        = b2go_TaskCallback
	_ enqueueTaskCallback = b2go_EnqueueTaskCallback
	_ finishTaskCallback  = b2go_FinishTaskCallback
	_ frictionCallback    = b2go_FrictionCallback
	_ restitutionCallback = b2go_RestitutionCallback
	_ customFilterFcn     = b2go_CustomFilterFcn
	_ preSolveFcn         = b2go_PreSolveFcn
	_ overlapResultFcn    = b2go_OverlapResultFcn
	_ castResultFcn       = b2go_CastResultFcn
	_ assertFcn           = b2go_AssertFcn
	_ drawPolygonFcn      = b2go_DrawPolygonFcn
	_ drawSolidPolygonFcn = b2go_DrawSolidPolygonFcn
	_ drawCircleFcn       = b2go_DrawCircleFcn
	_ drawSolidCircleFcn  = b2go_DrawSolidCircleFcn
	_ drawSolidCapsuleFcn = b2go_DrawSolidCapsuleFcn
	_ drawSegmentFcn      = b2go_DrawSegmentFcn
	_ drawTransformFcn    = b2go_DrawTransformFcn
	_ drawPointFcn        = b2go_DrawPointFcn
	_ drawStringFcn       = b2go_DrawStringFcn
)
