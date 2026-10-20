package b2

import (
	"math"
	"slices"
	"unsafe"
)

type (
	DrawPolygon      func(transform Transform, vertices []Vec2, color HexColor)
	DrawSolidPolygon func(transform Transform, vertices []Vec2, radius float32, color HexColor)
	DrawCircle       func(center Vec2, radius float32, color HexColor)
	DrawSolidCircle  func(transform Transform, center Vec2, radius float32, color HexColor)
	DrawSolidCapsule func(p1 Vec2, p2 Vec2, radius float32, color HexColor)
	DrawLine         func(p1 Vec2, p2 Vec2, color HexColor)
	DrawTransform    func(transform Transform)
	DrawPoint        func(p Vec2, size float32, color HexColor)
	DrawString       func(p Vec2, s string, color HexColor)
	DrawBounds       func(aabb AABB, color HexColor)
)

// DebugDraw holds the callbacks to draw a world. Start with DefaultDebugDraw,
// box2d draws nothing outside of DrawingBounds.
type DebugDraw struct {
	DrawPolygon      DrawPolygon
	DrawSolidPolygon DrawSolidPolygon
	DrawCircle       DrawCircle
	DrawSolidCircle  DrawSolidCircle
	DrawSolidCapsule DrawSolidCapsule
	DrawLine         DrawLine
	DrawTransform    DrawTransform
	DrawPoint        DrawPoint
	DrawString       DrawString
	DrawBounds       DrawBounds

	DrawingBounds AABB

	// ForceScale scales the drawn forces, JointScale the drawn joints.
	ForceScale float32
	JointScale float32

	DrawContacts        bool
	DrawAnchorA         bool
	DrawShapes          bool
	DrawChainNormals    bool
	DrawJoints          bool
	DrawJointExtras     bool
	DrawShapeBounds     bool
	DrawMass            bool
	DrawBodyNames       bool
	DrawGraphColors     bool
	DrawContactFeatures bool
	DrawContactNormals  bool
	DrawContactForces   bool
	DrawFrictionForces  bool
	DrawIslands         bool
}

// DefaultDebugDraw draws the shapes of the whole world, like b2DefaultDebugDraw.
func DefaultDebugDraw() DebugDraw {
	return DebugDraw{
		DrawingBounds: AABB{
			LowerBound: Vec2{X: -math.MaxFloat32, Y: -math.MaxFloat32},
			UpperBound: Vec2{X: math.MaxFloat32, Y: math.MaxFloat32},
		},
		ForceScale: 1,
		JointScale: 1,
		DrawShapes: true,
	}
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
	dd.DrawLineFcn = __ccgo_fp(drawLineFcn(cbDrawLine))
	dd.DrawTransformFcn = __ccgo_fp(drawTransformFcn(cbDrawTransform))
	dd.DrawPointFcn = __ccgo_fp(drawPointFcn(cbDrawPoint))
	dd.DrawStringFcn = __ccgo_fp(drawStringFcn(cbDrawString))
	dd.DrawBoundsFcn = __ccgo_fp(drawBoundsFcn(cbDrawBounds))

	dd.DrawingBounds = draw.DrawingBounds
	dd.ForceScale = draw.ForceScale
	dd.JointScale = draw.JointScale

	dd.DrawContacts = boolUint8(draw.DrawContacts)
	dd.DrawAnchorA = boolUint8(draw.DrawAnchorA)
	dd.DrawShapes = boolUint8(draw.DrawShapes)
	dd.DrawChainNormals = boolUint8(draw.DrawChainNormals)
	dd.DrawJoints = boolUint8(draw.DrawJoints)
	dd.DrawJointExtras = boolUint8(draw.DrawJointExtras)
	dd.DrawBounds = boolUint8(draw.DrawShapeBounds)
	dd.DrawMass = boolUint8(draw.DrawMass)
	dd.DrawBodyNames = boolUint8(draw.DrawBodyNames)
	dd.DrawGraphColors = boolUint8(draw.DrawGraphColors)
	dd.DrawContactFeatures = boolUint8(draw.DrawContactFeatures)
	dd.DrawContactNormals = boolUint8(draw.DrawContactNormals)
	dd.DrawContactForces = boolUint8(draw.DrawContactForces)
	dd.DrawFrictionForces = boolUint8(draw.DrawFrictionForces)
	dd.DrawIslands = boolUint8(draw.DrawIslands)

	dd.Context = context

	ddStack := copyToStack(theStack, dd)
	defer ddStack.Free()

	b2World_Draw(theStack, b.Id, ddStack.Addr)
}

func cbDrawPolygon(tls *_Stack, transform Transform, vertices uintptr, vertexCount int32, color HexColor, context uintptr) {
	ctx := tls.GetObject(context).(*DebugDraw)
	if ctx.DrawPolygon != nil {
		vertices := toSlice[Vec2](vertices, vertexCount)
		ctx.DrawPolygon(transform, vertices, color)
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

func cbDrawSolidCircle(tls *_Stack, transform Transform, center Vec2, radius float32, color HexColor, context uintptr) {
	ctx := tls.GetObject(context).(*DebugDraw)
	if ctx.DrawSolidCircle != nil {
		ctx.DrawSolidCircle(transform, center, radius, color)
	}
}

func cbDrawSolidCapsule(tls *_Stack, p1 Vec2, p2 Vec2, radius float32, color HexColor, context uintptr) {
	ctx := tls.GetObject(context).(*DebugDraw)
	if ctx.DrawSolidCapsule != nil {
		ctx.DrawSolidCapsule(p1, p2, radius, color)
	}
}

func cbDrawLine(tls *_Stack, p1 Vec2, p2 Vec2, color HexColor, context uintptr) {
	ctx := tls.GetObject(context).(*DebugDraw)
	if ctx.DrawLine != nil {
		ctx.DrawLine(p1, p2, color)
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

func cbDrawBounds(tls *_Stack, aabb AABB, color HexColor, context uintptr) {
	ctx := tls.GetObject(context).(*DebugDraw)
	if ctx.DrawBounds != nil {
		ctx.DrawBounds(aabb, color)
	}
}

func toSlice[T any](ptr uintptr, count int32) []T {
	return slices.Clone(unsafe.Slice((*T)(unsafe.Pointer(ptr)), count))
}
