package b2

import (
	"slices"
	"unsafe"
)

type DrawPolygon func(vertices []Vec2, color HexColor)
type DrawSolidPolygon func(transform Transform, vertices []Vec2, radius float32, color HexColor)
type DrawCircle func(center Vec2, radius float32, color HexColor)
type DrawSolidCircle func(transform Transform, radius float32, color HexColor)
type DrawSolidCapsule func(p1 Vec2, p2 Vec2, radius float32, color HexColor)
type DrawSegment func(p1 Vec2, p2 Vec2, color HexColor)
type DrawTransform func(transform Transform)
type DrawPoint func(p Vec2, size float32, color HexColor)
type DrawString func(p Vec2, s string, color HexColor)

type DebugDraw struct {
	DrawPolygon      DrawPolygon
	DrawSolidPolygon DrawSolidPolygon
	DrawCircle       DrawCircle
	DrawSolidCircle  DrawSolidCircle
	DrawSolidCapsule DrawSolidCapsule
	DrawSegment      DrawSegment
	DrawTransform    DrawTransform
	DrawPoint        DrawPoint
	DrawString       DrawString

	DrawingBounds    AABB
	UseDrawingBounds bool

	DrawJoints           bool
	DrawShapes           bool
	DrawJointExtras      bool
	DrawBounds           bool
	DrawMass             bool
	DrawBodyNames        bool
	DrawContacts         bool
	DrawGraphColors      bool
	DrawContactNormals   bool
	DrawContactImpulses  bool
	DrawContactFeatures  bool
	DrawFrictionImpulses bool
	DrawIslands          bool
}

func (b World) Draw(draw DebugDraw) {
	// move the DebugDraw object temporarily onto the heap
	context := theStack.RegisterObject(&draw)
	defer theStack.ForgetObject(context)

	var dd b2DebugDraw
	dd.DrawPolygonFcn = __ccgo_fp(drawPolygonFcn(cbDrawPolygon))
	dd.DrawSolidPolygonFcn = __ccgo_fp(drawSolidPolygonFcn(cbDrawSolidPolygon))
	dd.DrawCircleFcn = __ccgo_fp(drawCircleFcn(cbDrawCircle))
	dd.DrawSolidCircleFcn = __ccgo_fp(drawSolidCircleFcn(cbDrawSolidCircle))
	dd.DrawSolidCapsuleFcn = __ccgo_fp(drawSolidCapsuleFcn(cbDrawSolidCapsule))
	dd.DrawSegmentFcn = __ccgo_fp(drawSegmentFcn(cbDrawSegment))
	dd.DrawTransformFcn = __ccgo_fp(drawTransformFcn(cbDrawTransform))
	dd.DrawPointFcn = __ccgo_fp(drawPointFcn(cbDrawPoint))
	dd.DrawStringFcn = __ccgo_fp(drawStringFcn(cbDrawString))

	dd.DrawingBounds = draw.DrawingBounds
	dd.UseDrawingBounds = boolUint8(draw.UseDrawingBounds)

	dd.DrawShapes = boolUint8(draw.DrawShapes)
	dd.DrawJoints = boolUint8(draw.DrawJoints)
	dd.DrawJointExtras = boolUint8(draw.DrawJointExtras)
	dd.DrawBounds = boolUint8(draw.DrawBounds)
	dd.DrawMass = boolUint8(draw.DrawMass)
	dd.DrawBodyNames = boolUint8(draw.DrawBodyNames)
	dd.DrawContacts = boolUint8(draw.DrawContacts)
	dd.DrawGraphColors = boolUint8(draw.DrawGraphColors)
	dd.DrawContactNormals = boolUint8(draw.DrawContactNormals)
	dd.DrawContactImpulses = boolUint8(draw.DrawContactImpulses)
	dd.DrawContactFeatures = boolUint8(draw.DrawContactFeatures)
	dd.DrawFrictionImpulses = boolUint8(draw.DrawFrictionImpulses)
	dd.DrawIslands = boolUint8(draw.DrawIslands)

	dd.Context = context

	ddStack := copyToStack(theStack, dd)
	defer ddStack.Free()

	b2World_Draw(theStack, b.Id, ddStack.Addr)
}

func cbDrawPolygon(tls *_Stack, vertices uintptr, vertexCount int32, color HexColor, context uintptr) {
	ctx := tls.GetObject(context).(*DebugDraw)
	if ctx.DrawPolygon != nil {
		vertices := toSlice[Vec2](vertices, vertexCount)
		ctx.DrawPolygon(vertices, color)
	}
}
func cbDrawSolidPolygon(tls *_Stack, transform Transform, vertices uintptr, vertexCount int32, radius float32, color HexColor, context uintptr) {
	ctx := tls.GetObject(context).(*DebugDraw)
	if ctx.DrawSolidPolygon != nil {
		vertices := toSlice[Vec2](vertices, vertexCount)
		ctx.DrawSolidPolygon(transform, vertices, radius, color)
	}
}
func cbDrawCircle(tls *_Stack, center Vec2, radius float32, color HexColor, context uintptr) {
	ctx := tls.GetObject(context).(*DebugDraw)
	if ctx.DrawCircle != nil {
		ctx.DrawCircle(center, radius, color)
	}
}
func cbDrawSolidCircle(tls *_Stack, transform Transform, radius float32, color HexColor, context uintptr) {
	ctx := tls.GetObject(context).(*DebugDraw)
	if ctx.DrawSolidCircle != nil {
		ctx.DrawSolidCircle(transform, radius, color)
	}
}
func cbDrawSolidCapsule(tls *_Stack, p1 Vec2, p2 Vec2, radius float32, color HexColor, context uintptr) {
	ctx := tls.GetObject(context).(*DebugDraw)
	if ctx.DrawSolidCapsule != nil {
		ctx.DrawSolidCapsule(p1, p2, radius, color)
	}
}
func cbDrawSegment(tls *_Stack, p1 Vec2, p2 Vec2, color HexColor, context uintptr) {
	ctx := tls.GetObject(context).(*DebugDraw)
	if ctx.DrawSegment != nil {
		ctx.DrawSegment(p1, p2, color)
	}
}
func cbDrawTransform(tls *_Stack, transform Transform, context uintptr) {
	ctx := tls.GetObject(context).(*DebugDraw)
	if ctx.DrawTransform != nil {
		ctx.DrawTransform(transform)
	}
}
func cbDrawPoint(tls *_Stack, p Vec2, size float32, color HexColor, context uintptr) {
	ctx := tls.GetObject(context).(*DebugDraw)
	if ctx.DrawPoint != nil {
		ctx.DrawPoint(p, size, color)
	}
}
func cbDrawString(tls *_Stack, p Vec2, s uintptr, color HexColor, context uintptr) {
	ctx := tls.GetObject(context).(*DebugDraw)
	if ctx.DrawString != nil {
		ctx.DrawString(p, toString(s), color)
	}
}

func toSlice[T any](ptr uintptr, count int32) []T {
	return slices.Clone(unsafe.Slice((*T)(unsafe.Pointer(ptr)), count))
}
