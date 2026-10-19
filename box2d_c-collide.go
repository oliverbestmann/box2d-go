package b2

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

func b2CollideMoverAndCircle(tls *_Stack, shape uintptr, mover uintptr) (r PlaneResult) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var distanceOutput DistanceOutput
	var plane Plane
	var totalRadius float32
	var _ /* cache at bp+184 */ SimplexCache
	var _ /* distanceInput at bp+0 */ DistanceInput
	_, _, _ = distanceOutput, plane, totalRadius
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyA = b2MakeProxy(tls, shape, int32(1), float32FromFloat32(0))
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyB = b2MakeProxy(tls, mover, int32(2), (*Capsule)(unsafe.Pointer(mover)).Radius)
	(*(*DistanceInput)(unsafe.Add(nil, bp))).TransformA = b2Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp))).TransformB = b2Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp))).UseRadii = boolUint8(false1 != 0)
	totalRadius = (*Capsule)(unsafe.Pointer(mover)).Radius + (*Circle)(unsafe.Pointer(shape)).Radius
	*(*SimplexCache)(unsafe.Add(nil, bp+184)) = SimplexCache{}
	distanceOutput = b2ShapeDistance(tls, bp, bp+184, uintptrFromInt32(0), 0)
	if distanceOutput.Distance <= totalRadius {
		plane = Plane{
			Normal: distanceOutput.Normal,
			Offset: totalRadius - distanceOutput.Distance,
		}
		return PlaneResult{
			Plane: plane,
			Point: distanceOutput.PointA,
			Hit:   boolUint8(true1 != 0),
		}
	}
	return PlaneResult{}
}

func b2CollideMoverAndCapsule(tls *_Stack, shape uintptr, mover uintptr) (r PlaneResult) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var distanceOutput DistanceOutput
	var plane Plane
	var totalRadius float32
	var _ /* cache at bp+184 */ SimplexCache
	var _ /* distanceInput at bp+0 */ DistanceInput
	_, _, _ = distanceOutput, plane, totalRadius
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyA = b2MakeProxy(tls, shape, int32(2), float32FromFloat32(0))
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyB = b2MakeProxy(tls, mover, int32(2), (*Capsule)(unsafe.Pointer(mover)).Radius)
	(*(*DistanceInput)(unsafe.Add(nil, bp))).TransformA = b2Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp))).TransformB = b2Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp))).UseRadii = boolUint8(false1 != 0)
	totalRadius = (*Capsule)(unsafe.Pointer(mover)).Radius + (*Capsule)(unsafe.Pointer(shape)).Radius
	*(*SimplexCache)(unsafe.Add(nil, bp+184)) = SimplexCache{}
	distanceOutput = b2ShapeDistance(tls, bp, bp+184, uintptrFromInt32(0), 0)
	if distanceOutput.Distance <= totalRadius {
		plane = Plane{
			Normal: distanceOutput.Normal,
			Offset: totalRadius - distanceOutput.Distance,
		}
		return PlaneResult{
			Plane: plane,
			Point: distanceOutput.PointA,
			Hit:   boolUint8(true1 != 0),
		}
	}
	return PlaneResult{}
}

func b2CollideMoverAndPolygon(tls *_Stack, shape uintptr, mover uintptr) (r PlaneResult) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var distanceOutput DistanceOutput
	var plane Plane
	var totalRadius float32
	var _ /* cache at bp+184 */ SimplexCache
	var _ /* distanceInput at bp+0 */ DistanceInput
	_, _, _ = distanceOutput, plane, totalRadius
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyA = b2MakeProxy(tls, shape, (*Polygon)(unsafe.Pointer(shape)).Count, (*Polygon)(unsafe.Pointer(shape)).Radius)
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyB = b2MakeProxy(tls, mover, int32(2), (*Capsule)(unsafe.Pointer(mover)).Radius)
	(*(*DistanceInput)(unsafe.Add(nil, bp))).TransformA = b2Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp))).TransformB = b2Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp))).UseRadii = boolUint8(false1 != 0)
	totalRadius = (*Capsule)(unsafe.Pointer(mover)).Radius + (*Polygon)(unsafe.Pointer(shape)).Radius
	*(*SimplexCache)(unsafe.Add(nil, bp+184)) = SimplexCache{}
	distanceOutput = b2ShapeDistance(tls, bp, bp+184, uintptrFromInt32(0), 0)
	if distanceOutput.Distance <= totalRadius {
		plane = Plane{
			Normal: distanceOutput.Normal,
			Offset: totalRadius - distanceOutput.Distance,
		}
		return PlaneResult{
			Plane: plane,
			Point: distanceOutput.PointA,
			Hit:   boolUint8(true1 != 0),
		}
	}
	return PlaneResult{}
}

func b2CollideMoverAndSegment(tls *_Stack, shape uintptr, mover uintptr) (r PlaneResult) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var distanceOutput DistanceOutput
	var plane Plane
	var totalRadius float32
	var _ /* cache at bp+184 */ SimplexCache
	var _ /* distanceInput at bp+0 */ DistanceInput
	_, _, _ = distanceOutput, plane, totalRadius
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyA = b2MakeProxy(tls, shape, int32(2), float32FromFloat32(0))
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyB = b2MakeProxy(tls, mover, int32(2), (*Capsule)(unsafe.Pointer(mover)).Radius)
	(*(*DistanceInput)(unsafe.Add(nil, bp))).TransformA = b2Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp))).TransformB = b2Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp))).UseRadii = boolUint8(false1 != 0)
	totalRadius = (*Capsule)(unsafe.Pointer(mover)).Radius
	*(*SimplexCache)(unsafe.Add(nil, bp+184)) = SimplexCache{}
	distanceOutput = b2ShapeDistance(tls, bp, bp+184, uintptrFromInt32(0), 0)
	if distanceOutput.Distance <= totalRadius {
		plane = Plane{
			Normal: distanceOutput.Normal,
			Offset: totalRadius - distanceOutput.Distance,
		}
		return PlaneResult{
			Plane: plane,
			Point: distanceOutput.PointA,
			Hit:   boolUint8(true1 != 0),
		}
	}
	return PlaneResult{}
}

// C documentation
//
//	// point = qA * localAnchorA + pA
//	// localAnchorB = qBc * (point - pB)
//	// anchorB = point - pB = qA * localAnchorA + pA - pB
//	//         = anchorA + (pA - pB)
func b2CollideCircles(tls *_Stack, circleA uintptr, xfA Transform, circleB uintptr, xfB Transform) (r Manifold) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var C, xf, v1, v15, v17, v2 Transform
	var cA, cB, contactPointA, n, normal, pointA, pointB, v12, v13, v18, v19, v21, v22, v23, v26, v27, v7, v8, v9 Vec2
	var invLength, radiusA, radiusB, separation, x, y, v30, v35, v41 float32
	var mp, v25 uintptr
	var qr, v11, v3, v4, v5 Rot
	var _ /* distance at bp+112 */ float32
	var _ /* manifold at bp+0 */ Manifold
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, cA, cB, contactPointA, invLength, mp, n, normal, pointA, pointB, qr, radiusA, radiusB, separation, x, xf, y, v1, v11, v12, v13, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v3, v30, v35, v4, v41, v5, v7, v8, v9
	*(*Manifold)(unsafe.Add(nil, bp)) = Manifold{}
	v1 = xfA
	v2 = xfB
	v3 = v1.Q
	v4 = v2.Q
	qr.S = float32(v3.C*v4.S) - float32(v3.S*v4.C)
	qr.C = float32(v3.C*v4.C) + float32(v3.S*v4.S)
	v5 = qr
	goto _6
_6:
	C.Q = v5
	v7 = v2.P
	v8 = v1.P
	v9 = Vec2{
		X: v7.X - v8.X,
		Y: v7.Y - v8.Y,
	}
	goto _10
_10:
	v11 = v1.Q
	v12 = v9
	v13 = Vec2{
		X: float32(v11.C*v12.X) + float32(v11.S*v12.Y),
		Y: float32(-v11.S*v12.X) + float32(v11.C*v12.Y),
	}
	goto _14
_14:
	C.P = v13
	v15 = C
	goto _16
_16:
	xf = v15
	pointA = (*Circle)(unsafe.Pointer(circleA)).Center
	v17 = xf
	v18 = (*Circle)(unsafe.Pointer(circleB)).Center
	x = float32(v17.Q.C*v18.X) - float32(v17.Q.S*v18.Y) + v17.P.X
	y = float32(v17.Q.S*v18.X) + float32(v17.Q.C*v18.Y) + v17.P.Y
	v19 = Vec2{
		X: x,
		Y: y,
	}
	goto _20
_20:
	pointB = v19
	v21 = pointB
	v22 = pointA
	v23 = Vec2{
		X: v21.X - v22.X,
		Y: v21.Y - v22.Y,
	}
	goto _24
_24:
	v25 = bp + 112
	v26 = v23
	*(*float32)(unsafe.Add(nil, v25)) = sqrtf(tls, float32(v26.X*v26.X)+float32(v26.Y*v26.Y))
	if *(*float32)(unsafe.Add(nil, v25)) < float32FromFloat32(1.1920928955078125e-07) {
		v27 = Vec2{}
		goto _28
	}
	invLength = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v25))
	n = Vec2{
		X: float32(invLength * v26.X),
		Y: float32(invLength * v26.Y),
	}
	v27 = n
	goto _28
_28:
	normal = v27
	radiusA = (*Circle)(unsafe.Pointer(circleA)).Radius
	radiusB = (*Circle)(unsafe.Pointer(circleB)).Radius
	separation = *(*float32)(unsafe.Add(nil, bp+112)) - radiusA - radiusB
	if separation > float32(float32FromFloat32(4)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
		return *(*Manifold)(unsafe.Add(nil, bp))
	}
	v7 = pointA
	v30 = radiusA
	v8 = normal
	v9 = Vec2{
		X: v7.X + float32(v30*v8.X),
		Y: v7.Y + float32(v30*v8.Y),
	}
	goto _33
_33:
	cA = v9
	v12 = pointB
	v35 = -radiusB
	v13 = normal
	v18 = Vec2{
		X: v12.X + float32(v35*v13.X),
		Y: v12.Y + float32(v35*v13.Y),
	}
	goto _38
_38:
	cB = v18
	v19 = cA
	v21 = cB
	v41 = float32FromFloat32(0.5)
	v22 = Vec2{
		X: float32((float32FromFloat32(1)-v41)*v19.X) + float32(v41*v21.X),
		Y: float32((float32FromFloat32(1)-v41)*v19.Y) + float32(v41*v21.Y),
	}
	goto _43
_43:
	contactPointA = v22
	v3 = xfA.Q
	v23 = normal
	v26 = Vec2{
		X: float32(v3.C*v23.X) - float32(v3.S*v23.Y),
		Y: float32(v3.S*v23.X) + float32(v3.C*v23.Y),
	}
	goto _47
_47:
	(*(*Manifold)(unsafe.Add(nil, bp))).Normal = v26
	mp = bp + 12 + uintptr(0)*48
	v3 = xfA.Q
	v7 = contactPointA
	v8 = Vec2{
		X: float32(v3.C*v7.X) - float32(v3.S*v7.Y),
		Y: float32(v3.S*v7.X) + float32(v3.C*v7.Y),
	}
	goto _51
_51:
	(*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA = v8
	v7 = xfA.P
	v8 = xfB.P
	v9 = Vec2{
		X: v7.X - v8.X,
		Y: v7.Y - v8.Y,
	}
	goto _55
_55:
	v12 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
	v13 = v9
	v18 = Vec2{
		X: v12.X + v13.X,
		Y: v12.Y + v13.Y,
	}
	goto _59
_59:
	(*ManifoldPoint)(unsafe.Pointer(mp)).AnchorB = v18
	v7 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
	v8 = xfA.P
	v9 = Vec2{
		X: v7.X + v8.X,
		Y: v7.Y + v8.Y,
	}
	goto _63
_63:
	(*ManifoldPoint)(unsafe.Pointer(mp)).Point = v9
	(*ManifoldPoint)(unsafe.Pointer(mp)).Separation = separation
	(*ManifoldPoint)(unsafe.Pointer(mp)).Id = uint16(0)
	(*(*Manifold)(unsafe.Add(nil, bp))).PointCount = int32(1)
	return *(*Manifold)(unsafe.Add(nil, bp))
}

// C documentation
//
//	/// Compute the collision manifold between a capsule and circle
func b2CollideCapsuleAndCircle(tls *_Stack, capsuleA uintptr, xfA Transform, circleB uintptr, xfB Transform) (r Manifold) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var C, xf, v1, v15, v17, v2 Transform
	var cA, cB, contactPointA, e, n, normal, p1, p2, pA, pB, v12, v13, v18, v19, v21, v22, v23, v25, v26, v27, v29, v30, v33, v34, v35, v37, v38, v7, v8, v9 Vec2
	var invLength, radiusA, radiusB, s1, s11, s2, separation, x, y, v31, v39, v43 float32
	var mp, v54 uintptr
	var qr, v11, v3, v4, v5 Rot
	var _ /* distance at bp+112 */ float32
	var _ /* manifold at bp+0 */ Manifold
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, cA, cB, contactPointA, e, invLength, mp, n, normal, p1, p2, pA, pB, qr, radiusA, radiusB, s1, s11, s2, separation, x, xf, y, v1, v11, v12, v13, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v29, v3, v30, v31, v33, v34, v35, v37, v38, v39, v4, v43, v5, v54, v7, v8, v9
	*(*Manifold)(unsafe.Add(nil, bp)) = Manifold{}
	v1 = xfA
	v2 = xfB
	v3 = v1.Q
	v4 = v2.Q
	qr.S = float32(v3.C*v4.S) - float32(v3.S*v4.C)
	qr.C = float32(v3.C*v4.C) + float32(v3.S*v4.S)
	v5 = qr
	goto _6
_6:
	C.Q = v5
	v7 = v2.P
	v8 = v1.P
	v9 = Vec2{
		X: v7.X - v8.X,
		Y: v7.Y - v8.Y,
	}
	goto _10
_10:
	v11 = v1.Q
	v12 = v9
	v13 = Vec2{
		X: float32(v11.C*v12.X) + float32(v11.S*v12.Y),
		Y: float32(-v11.S*v12.X) + float32(v11.C*v12.Y),
	}
	goto _14
_14:
	C.P = v13
	v15 = C
	goto _16
_16:
	xf = v15
	v17 = xf
	v18 = (*Circle)(unsafe.Pointer(circleB)).Center
	x = float32(v17.Q.C*v18.X) - float32(v17.Q.S*v18.Y) + v17.P.X
	y = float32(v17.Q.S*v18.X) + float32(v17.Q.C*v18.Y) + v17.P.Y
	v19 = Vec2{
		X: x,
		Y: y,
	}
	goto _20
_20:
	// Compute circle position in the frame of the capsule.
	pB = v19
	// Compute closest point
	p1 = (*Capsule)(unsafe.Pointer(capsuleA)).Center1
	p2 = (*Capsule)(unsafe.Pointer(capsuleA)).Center2
	v21 = p2
	v22 = p1
	v23 = Vec2{
		X: v21.X - v22.X,
		Y: v21.Y - v22.Y,
	}
	goto _24
_24:
	e = v23
	v25 = pB
	v26 = p1
	v27 = Vec2{
		X: v25.X - v26.X,
		Y: v25.Y - v26.Y,
	}
	goto _28
_28:
	v29 = v27
	v30 = e
	v31 = float32(v29.X*v30.X) + float32(v29.Y*v30.Y)
	goto _32
_32:
	s11 = v31
	v33 = p2
	v34 = pB
	v35 = Vec2{
		X: v33.X - v34.X,
		Y: v33.Y - v34.Y,
	}
	goto _36
_36:
	v37 = v35
	v38 = e
	v39 = float32(v37.X*v38.X) + float32(v37.Y*v38.Y)
	goto _40
_40:
	s2 = v39
	if s11 < float32FromFloat32(0) {
		// p1 region
		pA = p1
	} else {
		if s2 < float32FromFloat32(0) {
			// p2 region
			pA = p2
		} else {
			v7 = e
			v8 = e
			v31 = float32(v7.X*v8.X) + float32(v7.Y*v8.Y)
			goto _44
		_44:
			// circle colliding with segment interior
			s1 = s11 / v31
			v9 = p1
			v39 = s1
			v12 = e
			v13 = Vec2{
				X: v9.X + float32(v39*v12.X),
				Y: v9.Y + float32(v39*v12.Y),
			}
			goto _49
		_49:
			pA = v13
		}
	}
	v7 = pB
	v8 = pA
	v9 = Vec2{
		X: v7.X - v8.X,
		Y: v7.Y - v8.Y,
	}
	goto _53
_53:
	v54 = bp + 112
	v12 = v9
	*(*float32)(unsafe.Add(nil, v54)) = sqrtf(tls, float32(v12.X*v12.X)+float32(v12.Y*v12.Y))
	if *(*float32)(unsafe.Add(nil, v54)) < float32FromFloat32(1.1920928955078125e-07) {
		v13 = Vec2{}
		goto _57
	}
	invLength = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v54))
	n = Vec2{
		X: float32(invLength * v12.X),
		Y: float32(invLength * v12.Y),
	}
	v13 = n
	goto _57
_57:
	normal = v13
	radiusA = (*Capsule)(unsafe.Pointer(capsuleA)).Radius
	radiusB = (*Circle)(unsafe.Pointer(circleB)).Radius
	separation = *(*float32)(unsafe.Add(nil, bp+112)) - radiusA - radiusB
	if separation > float32(float32FromFloat32(4)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
		return *(*Manifold)(unsafe.Add(nil, bp))
	}
	v7 = pA
	v31 = radiusA
	v8 = normal
	v9 = Vec2{
		X: v7.X + float32(v31*v8.X),
		Y: v7.Y + float32(v31*v8.Y),
	}
	goto _62
_62:
	cA = v9
	v12 = pB
	v39 = -radiusB
	v13 = normal
	v18 = Vec2{
		X: v12.X + float32(v39*v13.X),
		Y: v12.Y + float32(v39*v13.Y),
	}
	goto _67
_67:
	cB = v18
	v19 = cA
	v21 = cB
	v43 = float32FromFloat32(0.5)
	v22 = Vec2{
		X: float32((float32FromFloat32(1)-v43)*v19.X) + float32(v43*v21.X),
		Y: float32((float32FromFloat32(1)-v43)*v19.Y) + float32(v43*v21.Y),
	}
	goto _72
_72:
	contactPointA = v22
	v3 = xfA.Q
	v23 = normal
	v25 = Vec2{
		X: float32(v3.C*v23.X) - float32(v3.S*v23.Y),
		Y: float32(v3.S*v23.X) + float32(v3.C*v23.Y),
	}
	goto _76
_76:
	(*(*Manifold)(unsafe.Add(nil, bp))).Normal = v25
	mp = bp + 12 + uintptr(0)*48
	v3 = xfA.Q
	v7 = contactPointA
	v8 = Vec2{
		X: float32(v3.C*v7.X) - float32(v3.S*v7.Y),
		Y: float32(v3.S*v7.X) + float32(v3.C*v7.Y),
	}
	goto _80
_80:
	(*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA = v8
	v7 = xfA.P
	v8 = xfB.P
	v9 = Vec2{
		X: v7.X - v8.X,
		Y: v7.Y - v8.Y,
	}
	goto _84
_84:
	v12 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
	v13 = v9
	v18 = Vec2{
		X: v12.X + v13.X,
		Y: v12.Y + v13.Y,
	}
	goto _88
_88:
	(*ManifoldPoint)(unsafe.Pointer(mp)).AnchorB = v18
	v7 = xfA.P
	v8 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
	v9 = Vec2{
		X: v7.X + v8.X,
		Y: v7.Y + v8.Y,
	}
	goto _92
_92:
	(*ManifoldPoint)(unsafe.Pointer(mp)).Point = v9
	(*ManifoldPoint)(unsafe.Pointer(mp)).Separation = separation
	(*ManifoldPoint)(unsafe.Pointer(mp)).Id = uint16(0)
	(*(*Manifold)(unsafe.Add(nil, bp))).PointCount = int32(1)
	return *(*Manifold)(unsafe.Add(nil, bp))
}

func b2CollidePolygonAndCircle(tls *_Stack, polygonA uintptr, xfA Transform, circleB uintptr, xfB Transform) (r Manifold) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var C, xf, v1, v15, v17, v2 Transform
	var cA, cA1, cA2, cB, cB1, cB2, center, contactPointA, contactPointA1, contactPointA2, n, normal, normal1, normal2, v11, v21, v12, v13, v18, v19, v22, v23, v24, v26, v27, v31, v32, v33, v35, v36, v37, v39, v40, v43, v44, v45, v47, v48, v49, v7, v8, v9 Vec2
	var i, normalIndex, vertIndex1, vertIndex2, vertexCount, v30 int32
	var invLength, length, radius, radiusA, radiusB, s2, separation, speculativeDistance, u1, u2, x, y, v28, v41, v53, v68 float32
	var mp, mp1, mp2, normals, vertices uintptr
	var qr, v111, v3, v4, v5 Rot
	var _ /* manifold at bp+0 */ Manifold
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, cA, cA1, cA2, cB, cB1, cB2, center, contactPointA, contactPointA1, contactPointA2, i, invLength, length, mp, mp1, mp2, n, normal, normal1, normal2, normalIndex, normals, qr, radius, radiusA, radiusB, s2, separation, speculativeDistance, u1, u2, v11, v21, vertIndex1, vertIndex2, vertexCount, vertices, x, xf, y, v1, v111, v12, v13, v15, v17, v18, v19, v2, v22, v23, v24, v26, v27, v28, v3, v30, v31, v32, v33, v35, v36, v37, v39, v4, v40, v41, v43, v44, v45, v47, v48, v49, v5, v53, v68, v7, v8, v9
	*(*Manifold)(unsafe.Add(nil, bp)) = Manifold{}
	speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter))
	v1 = xfA
	v2 = xfB
	v3 = v1.Q
	v4 = v2.Q
	qr.S = float32(v3.C*v4.S) - float32(v3.S*v4.C)
	qr.C = float32(v3.C*v4.C) + float32(v3.S*v4.S)
	v5 = qr
	goto _6
_6:
	C.Q = v5
	v7 = v2.P
	v8 = v1.P
	v9 = Vec2{
		X: v7.X - v8.X,
		Y: v7.Y - v8.Y,
	}
	goto _10
_10:
	v111 = v1.Q
	v12 = v9
	v13 = Vec2{
		X: float32(v111.C*v12.X) + float32(v111.S*v12.Y),
		Y: float32(-v111.S*v12.X) + float32(v111.C*v12.Y),
	}
	goto _14
_14:
	C.P = v13
	v15 = C
	goto _16
_16:
	xf = v15
	v17 = xf
	v18 = (*Circle)(unsafe.Pointer(circleB)).Center
	x = float32(v17.Q.C*v18.X) - float32(v17.Q.S*v18.Y) + v17.P.X
	y = float32(v17.Q.S*v18.X) + float32(v17.Q.C*v18.Y) + v17.P.Y
	v19 = Vec2{
		X: x,
		Y: y,
	}
	goto _20
_20:
	// Compute circle position in the frame of the polygon.
	center = v19
	radiusA = (*Polygon)(unsafe.Pointer(polygonA)).Radius
	radiusB = (*Circle)(unsafe.Pointer(circleB)).Radius
	radius = radiusA + radiusB
	// Find the min separating edge.
	normalIndex = 0
	separation = -float32FromFloat32(3.4028234663852886e+38)
	vertexCount = (*Polygon)(unsafe.Pointer(polygonA)).Count
	vertices = polygonA
	normals = polygonA + 64
	i = 0
	for {
		if !(i < vertexCount) {
			break
		}
		v22 = center
		v23 = *(*Vec2)(unsafe.Add(nil, vertices+uintptr(i)*8))
		v24 = Vec2{
			X: v22.X - v23.X,
			Y: v22.Y - v23.Y,
		}
		goto _25
	_25:
		v26 = *(*Vec2)(unsafe.Add(nil, normals+uintptr(i)*8))
		v27 = v24
		v28 = float32(v26.X*v27.X) + float32(v26.Y*v27.Y)
		goto _29
	_29:
		s2 = v28
		if s2 > separation {
			separation = s2
			normalIndex = i
		}
		goto _21
	_21:
		;
		i = i + 1
	}
	if separation > radius+speculativeDistance {
		return *(*Manifold)(unsafe.Add(nil, bp))
	}
	// Vertices of the reference edge.
	vertIndex1 = normalIndex
	if vertIndex1+int32(1) < vertexCount {
		v30 = vertIndex1 + int32(1)
	} else {
		v30 = 0
	}
	vertIndex2 = v30
	v11 = *(*Vec2)(unsafe.Add(nil, vertices+uintptr(vertIndex1)*8))
	v21 = *(*Vec2)(unsafe.Add(nil, vertices+uintptr(vertIndex2)*8))
	v7 = center
	v8 = v11
	v9 = Vec2{
		X: v7.X - v8.X,
		Y: v7.Y - v8.Y,
	}
	goto _34
_34:
	v12 = v21
	v13 = v11
	v18 = Vec2{
		X: v12.X - v13.X,
		Y: v12.Y - v13.Y,
	}
	goto _38
_38:
	v19 = v9
	v22 = v18
	v28 = float32(v19.X*v22.X) + float32(v19.Y*v22.Y)
	goto _42
_42:
	// Compute barycentric coordinates
	u1 = v28
	v23 = center
	v24 = v21
	v26 = Vec2{
		X: v23.X - v24.X,
		Y: v23.Y - v24.Y,
	}
	goto _46
_46:
	v27 = v11
	v31 = v21
	v32 = Vec2{
		X: v27.X - v31.X,
		Y: v27.Y - v31.Y,
	}
	goto _50
_50:
	v33 = v26
	v35 = v32
	v41 = float32(v33.X*v35.X) + float32(v33.Y*v35.Y)
	goto _54
_54:
	u2 = v41
	if u1 < float32FromFloat32(0) && separation > float32FromFloat32(1.1920928955078125e-07) {
		v36 = center
		v37 = v11
		v39 = Vec2{
			X: v36.X - v37.X,
			Y: v36.Y - v37.Y,
		}
		goto _58
	_58:
		v40 = v39
		length = sqrtf(tls, float32(v40.X*v40.X)+float32(v40.Y*v40.Y))
		if length < float32FromFloat32(1.1920928955078125e-07) {
			v43 = Vec2{}
			goto _61
		}
		invLength = float32FromFloat32(1) / length
		n = Vec2{
			X: float32(invLength * v40.X),
			Y: float32(invLength * v40.Y),
		}
		v43 = n
		goto _61
	_61:
		// Circle center is closest to v1 and safely outside the polygon
		normal = v43
		v44 = center
		v45 = v11
		v47 = Vec2{
			X: v44.X - v45.X,
			Y: v44.Y - v45.Y,
		}
		goto _65
	_65:
		v48 = v47
		v49 = normal
		v53 = float32(v48.X*v49.X) + float32(v48.Y*v49.Y)
		goto _69
	_69:
		separation = v53
		if separation > radius+speculativeDistance {
			return *(*Manifold)(unsafe.Add(nil, bp))
		}
		v7 = v11
		v28 = radiusA
		v8 = normal
		v9 = Vec2{
			X: v7.X + float32(v28*v8.X),
			Y: v7.Y + float32(v28*v8.Y),
		}
		goto _74
	_74:
		cA = v9
		v12 = center
		v41 = radiusB
		v13 = normal
		v18 = Vec2{
			X: v12.X - float32(v41*v13.X),
			Y: v12.Y - float32(v41*v13.Y),
		}
		goto _79
	_79:
		cB = v18
		v19 = cA
		v22 = cB
		v53 = float32FromFloat32(0.5)
		v23 = Vec2{
			X: float32((float32FromFloat32(1)-v53)*v19.X) + float32(v53*v22.X),
			Y: float32((float32FromFloat32(1)-v53)*v19.Y) + float32(v53*v22.Y),
		}
		goto _84
	_84:
		contactPointA = v23
		v3 = xfA.Q
		v24 = normal
		v26 = Vec2{
			X: float32(v3.C*v24.X) - float32(v3.S*v24.Y),
			Y: float32(v3.S*v24.X) + float32(v3.C*v24.Y),
		}
		goto _88
	_88:
		(*(*Manifold)(unsafe.Add(nil, bp))).Normal = v26
		mp = bp + 12 + uintptr(0)*48
		v3 = xfA.Q
		v7 = contactPointA
		v8 = Vec2{
			X: float32(v3.C*v7.X) - float32(v3.S*v7.Y),
			Y: float32(v3.S*v7.X) + float32(v3.C*v7.Y),
		}
		goto _92
	_92:
		(*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA = v8
		v7 = xfA.P
		v8 = xfB.P
		v9 = Vec2{
			X: v7.X - v8.X,
			Y: v7.Y - v8.Y,
		}
		goto _96
	_96:
		v12 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
		v13 = v9
		v18 = Vec2{
			X: v12.X + v13.X,
			Y: v12.Y + v13.Y,
		}
		goto _100
	_100:
		(*ManifoldPoint)(unsafe.Pointer(mp)).AnchorB = v18
		v7 = xfA.P
		v8 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
		v9 = Vec2{
			X: v7.X + v8.X,
			Y: v7.Y + v8.Y,
		}
		goto _104
	_104:
		(*ManifoldPoint)(unsafe.Pointer(mp)).Point = v9
		v7 = cB
		v8 = cA
		v9 = Vec2{
			X: v7.X - v8.X,
			Y: v7.Y - v8.Y,
		}
		goto _108
	_108:
		v12 = v9
		v13 = normal
		v28 = float32(v12.X*v13.X) + float32(v12.Y*v13.Y)
		goto _112
	_112:
		(*ManifoldPoint)(unsafe.Pointer(mp)).Separation = v28
		(*ManifoldPoint)(unsafe.Pointer(mp)).Id = uint16(0)
		(*(*Manifold)(unsafe.Add(nil, bp))).PointCount = int32(1)
	} else {
		if u2 < float32FromFloat32(0) && separation > float32FromFloat32(1.1920928955078125e-07) {
			v7 = center
			v8 = v21
			v9 = Vec2{
				X: v7.X - v8.X,
				Y: v7.Y - v8.Y,
			}
			goto _116
		_116:
			v12 = v9
			length = sqrtf(tls, float32(v12.X*v12.X)+float32(v12.Y*v12.Y))
			if length < float32FromFloat32(1.1920928955078125e-07) {
				v13 = Vec2{}
				goto _119
			}
			invLength = float32FromFloat32(1) / length
			n = Vec2{
				X: float32(invLength * v12.X),
				Y: float32(invLength * v12.Y),
			}
			v13 = n
			goto _119
		_119:
			// Circle center is closest to v2 and safely outside the polygon
			normal1 = v13
			v18 = center
			v19 = v21
			v22 = Vec2{
				X: v18.X - v19.X,
				Y: v18.Y - v19.Y,
			}
			goto _123
		_123:
			v23 = v22
			v24 = normal1
			v28 = float32(v23.X*v24.X) + float32(v23.Y*v24.Y)
			goto _127
		_127:
			separation = v28
			if separation > radius+speculativeDistance {
				return *(*Manifold)(unsafe.Add(nil, bp))
			}
			v7 = v21
			v28 = radiusA
			v8 = normal1
			v9 = Vec2{
				X: v7.X + float32(v28*v8.X),
				Y: v7.Y + float32(v28*v8.Y),
			}
			goto _132
		_132:
			cA1 = v9
			v12 = center
			v41 = radiusB
			v13 = normal1
			v18 = Vec2{
				X: v12.X - float32(v41*v13.X),
				Y: v12.Y - float32(v41*v13.Y),
			}
			goto _137
		_137:
			cB1 = v18
			v19 = cA1
			v22 = cB1
			v53 = float32FromFloat32(0.5)
			v23 = Vec2{
				X: float32((float32FromFloat32(1)-v53)*v19.X) + float32(v53*v22.X),
				Y: float32((float32FromFloat32(1)-v53)*v19.Y) + float32(v53*v22.Y),
			}
			goto _142
		_142:
			contactPointA1 = v23
			v3 = xfA.Q
			v24 = normal1
			v26 = Vec2{
				X: float32(v3.C*v24.X) - float32(v3.S*v24.Y),
				Y: float32(v3.S*v24.X) + float32(v3.C*v24.Y),
			}
			goto _146
		_146:
			(*(*Manifold)(unsafe.Add(nil, bp))).Normal = v26
			mp1 = bp + 12 + uintptr(0)*48
			v3 = xfA.Q
			v7 = contactPointA1
			v8 = Vec2{
				X: float32(v3.C*v7.X) - float32(v3.S*v7.Y),
				Y: float32(v3.S*v7.X) + float32(v3.C*v7.Y),
			}
			goto _150
		_150:
			(*ManifoldPoint)(unsafe.Pointer(mp1)).AnchorA = v8
			v7 = xfA.P
			v8 = xfB.P
			v9 = Vec2{
				X: v7.X - v8.X,
				Y: v7.Y - v8.Y,
			}
			goto _154
		_154:
			v12 = (*ManifoldPoint)(unsafe.Pointer(mp1)).AnchorA
			v13 = v9
			v18 = Vec2{
				X: v12.X + v13.X,
				Y: v12.Y + v13.Y,
			}
			goto _158
		_158:
			(*ManifoldPoint)(unsafe.Pointer(mp1)).AnchorB = v18
			v7 = xfA.P
			v8 = (*ManifoldPoint)(unsafe.Pointer(mp1)).AnchorA
			v9 = Vec2{
				X: v7.X + v8.X,
				Y: v7.Y + v8.Y,
			}
			goto _162
		_162:
			(*ManifoldPoint)(unsafe.Pointer(mp1)).Point = v9
			v7 = cB1
			v8 = cA1
			v9 = Vec2{
				X: v7.X - v8.X,
				Y: v7.Y - v8.Y,
			}
			goto _166
		_166:
			v12 = v9
			v13 = normal1
			v28 = float32(v12.X*v13.X) + float32(v12.Y*v13.Y)
			goto _170
		_170:
			(*ManifoldPoint)(unsafe.Pointer(mp1)).Separation = v28
			(*ManifoldPoint)(unsafe.Pointer(mp1)).Id = uint16(0)
			(*(*Manifold)(unsafe.Add(nil, bp))).PointCount = int32(1)
		} else {
			// Circle center is between v1 and v2. Center may be inside polygon
			normal2 = *(*Vec2)(unsafe.Add(nil, normals+uintptr(normalIndex)*8))
			v3 = xfA.Q
			v7 = normal2
			v8 = Vec2{
				X: float32(v3.C*v7.X) - float32(v3.S*v7.Y),
				Y: float32(v3.S*v7.X) + float32(v3.C*v7.Y),
			}
			goto _174
		_174:
			(*(*Manifold)(unsafe.Add(nil, bp))).Normal = v8
			v7 = center
			v8 = v11
			v9 = Vec2{
				X: v7.X - v8.X,
				Y: v7.Y - v8.Y,
			}
			goto _178
		_178:
			v12 = v9
			v13 = normal2
			v28 = float32(v12.X*v13.X) + float32(v12.Y*v13.Y)
			goto _182
		_182:
			v18 = center
			v41 = radiusA - v28
			v19 = normal2
			v22 = Vec2{
				X: v18.X + float32(v41*v19.X),
				Y: v18.Y + float32(v41*v19.Y),
			}
			goto _187
		_187:
			// cA is the projection of the circle center onto to the reference edge
			cA2 = v22
			v23 = center
			v53 = radiusB
			v24 = normal2
			v26 = Vec2{
				X: v23.X - float32(v53*v24.X),
				Y: v23.Y - float32(v53*v24.Y),
			}
			goto _192
		_192:
			// cB is the deepest point on the circle with respect to the reference edge
			cB2 = v26
			v27 = cA2
			v31 = cB2
			v68 = float32FromFloat32(0.5)
			v32 = Vec2{
				X: float32((float32FromFloat32(1)-v68)*v27.X) + float32(v68*v31.X),
				Y: float32((float32FromFloat32(1)-v68)*v27.Y) + float32(v68*v31.Y),
			}
			goto _197
		_197:
			contactPointA2 = v32
			// The contact point is the midpoint in world space
			mp2 = bp + 12 + uintptr(0)*48
			v3 = xfA.Q
			v33 = contactPointA2
			v35 = Vec2{
				X: float32(v3.C*v33.X) - float32(v3.S*v33.Y),
				Y: float32(v3.S*v33.X) + float32(v3.C*v33.Y),
			}
			goto _201
		_201:
			(*ManifoldPoint)(unsafe.Pointer(mp2)).AnchorA = v35
			v7 = xfA.P
			v8 = xfB.P
			v9 = Vec2{
				X: v7.X - v8.X,
				Y: v7.Y - v8.Y,
			}
			goto _205
		_205:
			v12 = (*ManifoldPoint)(unsafe.Pointer(mp2)).AnchorA
			v13 = v9
			v18 = Vec2{
				X: v12.X + v13.X,
				Y: v12.Y + v13.Y,
			}
			goto _209
		_209:
			(*ManifoldPoint)(unsafe.Pointer(mp2)).AnchorB = v18
			v7 = xfA.P
			v8 = (*ManifoldPoint)(unsafe.Pointer(mp2)).AnchorA
			v9 = Vec2{
				X: v7.X + v8.X,
				Y: v7.Y + v8.Y,
			}
			goto _213
		_213:
			(*ManifoldPoint)(unsafe.Pointer(mp2)).Point = v9
			(*ManifoldPoint)(unsafe.Pointer(mp2)).Separation = separation - radius
			(*ManifoldPoint)(unsafe.Pointer(mp2)).Id = uint16(0)
			(*(*Manifold)(unsafe.Add(nil, bp))).PointCount = int32(1)
		}
	}
	return *(*Manifold)(unsafe.Add(nil, bp))
}

// C documentation
//
//	// Follows Ericson 5.1.9 Closest Points of Two Line Segments
//	// Adds some logic to support clipping to get two contact points
func b2CollideCapsules(tls *_Stack, capsuleA uintptr, xfA Transform, capsuleB uintptr, xfB Transform) (r Manifold) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var C, sfA, xf, v10, v23, v29, v33, v9 Transform
	var c, c1, c2, closest1, closest2, cp, cp1, cq, cq1, d1, d2, n, n1, normal, normalA, normalB, origin, p1, p2, q11, q21, r1, u1, u2, v15, v16, v17, v2, v20, v21, v25, v26, v27, v3, v30, v31, v34, v35, v37, v38, v39, v41, v42, v43, v45, v46, v49, v5, v50, v6, v7 Vec2
	var d12, dd1, dd2, denom, distance, distanceSquared, epsSqr, f1, f2, fp1, fp2, fq1, fq2, invLength, invLength1, length, maxDistance, radius, radiusA, radiusB, rd1, rd2, s1n, s1n1, s1p, s1p1, separationA, separationB, sp, sp1, sq, sq1, ss1, ss11, ss2, ss21, x, y, v47, v51, v59, v63, v67, v69, v70, v71, v72 float32
	var i, i1, i2, v311, v312 int32
	var mp, mp1, mp2, v104, v108 uintptr
	var outsideA, outsideB uint8
	var qr, v1, v11, v12, v13, v19 Rot
	var _ /* length1 at bp+112 */ float32
	var _ /* length2 at bp+116 */ float32
	var _ /* manifold at bp+0 */ Manifold
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, c, c1, c2, closest1, closest2, cp, cp1, cq, cq1, d1, d12, d2, dd1, dd2, denom, distance, distanceSquared, epsSqr, f1, f2, fp1, fp2, fq1, fq2, i, i1, i2, invLength, invLength1, length, maxDistance, mp, mp1, mp2, n, n1, normal, normalA, normalB, origin, outsideA, outsideB, p1, p2, q11, q21, qr, r1, radius, radiusA, radiusB, rd1, rd2, s1n, s1n1, s1p, s1p1, separationA, separationB, sfA, sp, sp1, sq, sq1, ss1, ss11, ss2, ss21, u1, u2, x, xf, y, v1, v10, v104, v108, v11, v12, v13, v15, v16, v17, v19, v2, v20, v21, v23, v25, v26, v27, v29, v3, v30, v31, v311, v312, v33, v34, v35, v37, v38, v39, v41, v42, v43, v45, v46, v47, v49, v5, v50, v51, v59, v6, v63, v67, v69, v7, v70, v71, v72, v9
	origin = (*Capsule)(unsafe.Pointer(capsuleA)).Center1
	v1 = xfA.Q
	v2 = origin
	v3 = Vec2{
		X: float32(v1.C*v2.X) - float32(v1.S*v2.Y),
		Y: float32(v1.S*v2.X) + float32(v1.C*v2.Y),
	}
	goto _4
_4:
	v5 = xfA.P
	v6 = v3
	v7 = Vec2{
		X: v5.X + v6.X,
		Y: v5.Y + v6.Y,
	}
	goto _8
_8:
	// Shift polyA to origin
	// pw = q * pb + p
	// pw = q * (pbs + origin) + p
	// pw = q * pbs + (p + q * origin)
	sfA = Transform{
		P: v7,
		Q: xfA.Q,
	}
	v9 = sfA
	v10 = xfB
	v11 = v9.Q
	v12 = v10.Q
	qr.S = float32(v11.C*v12.S) - float32(v11.S*v12.C)
	qr.C = float32(v11.C*v12.C) + float32(v11.S*v12.S)
	v13 = qr
	goto _14
_14:
	C.Q = v13
	v15 = v10.P
	v16 = v9.P
	v17 = Vec2{
		X: v15.X - v16.X,
		Y: v15.Y - v16.Y,
	}
	goto _18
_18:
	v19 = v9.Q
	v20 = v17
	v21 = Vec2{
		X: float32(v19.C*v20.X) + float32(v19.S*v20.Y),
		Y: float32(-v19.S*v20.X) + float32(v19.C*v20.Y),
	}
	goto _22
_22:
	C.P = v21
	v23 = C
	goto _24
_24:
	xf = v23
	p1 = b2Vec2_zero
	v25 = (*Capsule)(unsafe.Pointer(capsuleA)).Center2
	v26 = origin
	v27 = Vec2{
		X: v25.X - v26.X,
		Y: v25.Y - v26.Y,
	}
	goto _28
_28:
	q11 = v27
	v29 = xf
	v30 = (*Capsule)(unsafe.Pointer(capsuleB)).Center1
	x = float32(v29.Q.C*v30.X) - float32(v29.Q.S*v30.Y) + v29.P.X
	y = float32(v29.Q.S*v30.X) + float32(v29.Q.C*v30.Y) + v29.P.Y
	v31 = Vec2{
		X: x,
		Y: y,
	}
	goto _32
_32:
	p2 = v31
	v33 = xf
	v34 = (*Capsule)(unsafe.Pointer(capsuleB)).Center2
	x = float32(v33.Q.C*v34.X) - float32(v33.Q.S*v34.Y) + v33.P.X
	y = float32(v33.Q.S*v34.X) + float32(v33.Q.C*v34.Y) + v33.P.Y
	v35 = Vec2{
		X: x,
		Y: y,
	}
	goto _36
_36:
	q21 = v35
	v37 = q11
	v38 = p1
	v39 = Vec2{
		X: v37.X - v38.X,
		Y: v37.Y - v38.Y,
	}
	goto _40
_40:
	d1 = v39
	v41 = q21
	v42 = p2
	v43 = Vec2{
		X: v41.X - v42.X,
		Y: v41.Y - v42.Y,
	}
	goto _44
_44:
	d2 = v43
	v45 = d1
	v46 = d1
	v47 = float32(v45.X*v46.X) + float32(v45.Y*v46.Y)
	goto _48
_48:
	dd1 = v47
	v49 = d2
	v50 = d2
	v51 = float32(v49.X*v50.X) + float32(v49.Y*v50.Y)
	goto _52
_52:
	dd2 = v51
	epsSqr = float32(float32FromFloat32(1.1920928955078125e-07) * float32FromFloat32(1.1920928955078125e-07))
	if !(dd1 > epsSqr && dd2 > epsSqr) && b2InternalAssertFcn(tls, __ccgo_ts+9962, __ccgo_ts+9936, int32FromInt32(285)) != 0 {
		__builtin_trap(tls)
	}
	v2 = p1
	v3 = p2
	v5 = Vec2{
		X: v2.X - v3.X,
		Y: v2.Y - v3.Y,
	}
	goto _56
_56:
	r1 = v5
	v6 = r1
	v7 = d1
	v47 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y)
	goto _60
_60:
	rd1 = v47
	v15 = r1
	v16 = d2
	v51 = float32(v15.X*v16.X) + float32(v15.Y*v16.Y)
	goto _64
_64:
	rd2 = v51
	v17 = d1
	v20 = d2
	v59 = float32(v17.X*v20.X) + float32(v17.Y*v20.Y)
	goto _68
_68:
	d12 = v59
	denom = float32(dd1*dd2) - float32(d12*d12)
	// Fraction on segment 1
	f1 = float32FromFloat32(0)
	if denom != float32FromFloat32(0) {
		// not parallel
		v63 = (float32(d12*rd2) - float32(rd1*dd2)) / denom
		v67 = float32FromFloat32(0)
		v69 = float32FromFloat32(1)
		if v63 < v67 {
			v71 = v67
		} else {
			if v63 > v69 {
				v72 = v69
			} else {
				v72 = v63
			}
			v71 = v72
		}
		v70 = v71
		goto _73
	_73:
		f1 = v70
	}
	// Compute point on segment 2 closest to p1 + f1 * d1
	f2 = (float32(d12*f1) + rd2) / dd2
	// Clamping of segment 2 requires a do over on segment 1
	if f2 < float32FromFloat32(0) {
		f2 = float32FromFloat32(0)
		v47 = -rd1 / dd1
		v51 = float32FromFloat32(0)
		v59 = float32FromFloat32(1)
		if v47 < v51 {
			v67 = v51
		} else {
			if v47 > v59 {
				v69 = v59
			} else {
				v69 = v47
			}
			v67 = v69
		}
		v63 = v67
		goto _80
	_80:
		f1 = v63
	} else {
		if f2 > float32FromFloat32(1) {
			f2 = float32FromFloat32(1)
			v47 = (d12 - rd1) / dd1
			v51 = float32FromFloat32(0)
			v59 = float32FromFloat32(1)
			if v47 < v51 {
				v67 = v51
			} else {
				if v47 > v59 {
					v69 = v59
				} else {
					v69 = v47
				}
				v67 = v69
			}
			v63 = v67
			goto _87
		_87:
			f1 = v63
		}
	}
	v2 = p1
	v47 = f1
	v3 = d1
	v5 = Vec2{
		X: v2.X + float32(v47*v3.X),
		Y: v2.Y + float32(v47*v3.Y),
	}
	goto _94
_94:
	closest1 = v5
	v6 = p2
	v51 = f2
	v7 = d2
	v15 = Vec2{
		X: v6.X + float32(v51*v7.X),
		Y: v6.Y + float32(v51*v7.Y),
	}
	goto _99
_99:
	closest2 = v15
	v16 = closest1
	v17 = closest2
	c = Vec2{
		X: v17.X - v16.X,
		Y: v17.Y - v16.Y,
	}
	v59 = float32(c.X*c.X) + float32(c.Y*c.Y)
	goto _103
_103:
	distanceSquared = v59
	*(*Manifold)(unsafe.Add(nil, bp)) = Manifold{}
	radiusA = (*Capsule)(unsafe.Pointer(capsuleA)).Radius
	radiusB = (*Capsule)(unsafe.Pointer(capsuleB)).Radius
	radius = radiusA + radiusB
	maxDistance = radius + float32(float32FromFloat32(4)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter))
	if distanceSquared > float32(maxDistance*maxDistance) {
		return *(*Manifold)(unsafe.Add(nil, bp))
	}
	distance = sqrtf(tls, distanceSquared)
	v104 = bp + 112
	v2 = d1
	*(*float32)(unsafe.Add(nil, v104)) = sqrtf(tls, float32(v2.X*v2.X)+float32(v2.Y*v2.Y))
	if *(*float32)(unsafe.Add(nil, v104)) < float32FromFloat32(1.1920928955078125e-07) {
		v3 = Vec2{}
		goto _107
	}
	invLength1 = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v104))
	n1 = Vec2{
		X: float32(invLength1 * v2.X),
		Y: float32(invLength1 * v2.Y),
	}
	v3 = n1
	goto _107
_107:
	u1 = v3
	v108 = bp + 116
	v5 = d2
	*(*float32)(unsafe.Add(nil, v108)) = sqrtf(tls, float32(v5.X*v5.X)+float32(v5.Y*v5.Y))
	if *(*float32)(unsafe.Add(nil, v108)) < float32FromFloat32(1.1920928955078125e-07) {
		v6 = Vec2{}
		goto _111
	}
	invLength1 = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v108))
	n1 = Vec2{
		X: float32(invLength1 * v5.X),
		Y: float32(invLength1 * v5.Y),
	}
	v6 = n1
	goto _111
_111:
	u2 = v6
	v7 = p2
	v15 = p1
	v16 = Vec2{
		X: v7.X - v15.X,
		Y: v7.Y - v15.Y,
	}
	goto _115
_115:
	v17 = v16
	v20 = u1
	v47 = float32(v17.X*v20.X) + float32(v17.Y*v20.Y)
	goto _119
_119:
	// Does segment B project outside segment A?
	fp2 = v47
	v21 = q21
	v25 = p1
	v26 = Vec2{
		X: v21.X - v25.X,
		Y: v21.Y - v25.Y,
	}
	goto _123
_123:
	v27 = v26
	v30 = u1
	v51 = float32(v27.X*v30.X) + float32(v27.Y*v30.Y)
	goto _127
_127:
	fq2 = v51
	outsideA = boolUint8(fp2 <= float32FromFloat32(0) && fq2 <= float32FromFloat32(0) || fp2 >= *(*float32)(unsafe.Add(nil, bp+112)) && fq2 >= *(*float32)(unsafe.Add(nil, bp+112)))
	v31 = p1
	v34 = p2
	v35 = Vec2{
		X: v31.X - v34.X,
		Y: v31.Y - v34.Y,
	}
	goto _131
_131:
	v37 = v35
	v38 = u2
	v59 = float32(v37.X*v38.X) + float32(v37.Y*v38.Y)
	goto _135
_135:
	// Does segment A project outside segment B?
	fp1 = v59
	v39 = q11
	v41 = p2
	v42 = Vec2{
		X: v39.X - v41.X,
		Y: v39.Y - v41.Y,
	}
	goto _139
_139:
	v43 = v42
	v45 = u2
	v63 = float32(v43.X*v45.X) + float32(v43.Y*v45.Y)
	goto _143
_143:
	fq1 = v63
	outsideB = boolUint8(fp1 <= float32FromFloat32(0) && fq1 <= float32FromFloat32(0) || fp1 >= *(*float32)(unsafe.Add(nil, bp+116)) && fq1 >= *(*float32)(unsafe.Add(nil, bp+116)))
	if int32FromUint8(outsideA) == false1 && int32FromUint8(outsideB) == false1 {
		v46 = u1
		v49 = Vec2{
			X: -v46.Y,
			Y: v46.X,
		}
		goto _146
	_146:
		normalA = v49
		v2 = p2
		v3 = p1
		v5 = Vec2{
			X: v2.X - v3.X,
			Y: v2.Y - v3.Y,
		}
		goto _150
	_150:
		v6 = v5
		v7 = normalA
		v47 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y)
		goto _154
	_154:
		ss1 = v47
		v15 = q21
		v16 = p1
		v17 = Vec2{
			X: v15.X - v16.X,
			Y: v15.Y - v16.Y,
		}
		goto _158
	_158:
		v20 = v17
		v21 = normalA
		v51 = float32(v20.X*v21.X) + float32(v20.Y*v21.Y)
		goto _162
	_162:
		ss2 = v51
		if ss1 < ss2 {
			v59 = ss1
		} else {
			v59 = ss2
		}
		s1p = v59
		if -ss1 < -ss2 {
			v63 = -ss1
		} else {
			v63 = -ss2
		}
		s1n = v63
		if s1p > s1n {
			separationA = s1p
		} else {
			separationA = s1n
			v2 = normalA
			v3 = Vec2{
				X: -v2.X,
				Y: -v2.Y,
			}
			goto _167
		_167:
			normalA = v3
		}
		v2 = u2
		v3 = Vec2{
			X: -v2.Y,
			Y: v2.X,
		}
		goto _170
	_170:
		normalB = v3
		v2 = p1
		v3 = p2
		v5 = Vec2{
			X: v2.X - v3.X,
			Y: v2.Y - v3.Y,
		}
		goto _174
	_174:
		v6 = v5
		v7 = normalB
		v47 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y)
		goto _178
	_178:
		ss11 = v47
		v15 = q11
		v16 = p2
		v17 = Vec2{
			X: v15.X - v16.X,
			Y: v15.Y - v16.Y,
		}
		goto _182
	_182:
		v20 = v17
		v21 = normalB
		v51 = float32(v20.X*v21.X) + float32(v20.Y*v21.Y)
		goto _186
	_186:
		ss21 = v51
		if ss11 < ss21 {
			v59 = ss11
		} else {
			v59 = ss21
		}
		s1p1 = v59
		if -ss11 < -ss21 {
			v63 = -ss11
		} else {
			v63 = -ss21
		}
		s1n1 = v63
		if s1p1 > s1n1 {
			separationB = s1p1
		} else {
			separationB = s1n1
			v2 = normalB
			v3 = Vec2{
				X: -v2.X,
				Y: -v2.Y,
			}
			goto _191
		_191:
			normalB = v3
		}
		// biased to avoid feature flip-flop
		// todo more testing?
		if separationA+float32(float32FromFloat32(0.1)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) >= separationB {
			(*(*Manifold)(unsafe.Add(nil, bp))).Normal = normalA
			cp = p2
			cq = q21
			// clip to p1
			if fp2 < float32FromFloat32(0) && fq2 > float32FromFloat32(0) {
				v2 = p2
				v3 = q21
				v47 = (float32FromFloat32(0) - fp2) / (fq2 - fp2)
				v5 = Vec2{
					X: float32((float32FromFloat32(1)-v47)*v2.X) + float32(v47*v3.X),
					Y: float32((float32FromFloat32(1)-v47)*v2.Y) + float32(v47*v3.Y),
				}
				goto _196
			_196:
				cp = v5
			} else {
				if fq2 < float32FromFloat32(0) && fp2 > float32FromFloat32(0) {
					v2 = q21
					v3 = p2
					v47 = (float32FromFloat32(0) - fq2) / (fp2 - fq2)
					v5 = Vec2{
						X: float32((float32FromFloat32(1)-v47)*v2.X) + float32(v47*v3.X),
						Y: float32((float32FromFloat32(1)-v47)*v2.Y) + float32(v47*v3.Y),
					}
					goto _201
				_201:
					cq = v5
				}
			}
			// clip to q1
			if fp2 > *(*float32)(unsafe.Add(nil, bp+112)) && fq2 < *(*float32)(unsafe.Add(nil, bp+112)) {
				v2 = p2
				v3 = q21
				v47 = (fp2 - *(*float32)(unsafe.Add(nil, bp+112))) / (fp2 - fq2)
				v5 = Vec2{
					X: float32((float32FromFloat32(1)-v47)*v2.X) + float32(v47*v3.X),
					Y: float32((float32FromFloat32(1)-v47)*v2.Y) + float32(v47*v3.Y),
				}
				goto _206
			_206:
				cp = v5
			} else {
				if fq2 > *(*float32)(unsafe.Add(nil, bp+112)) && fp2 < *(*float32)(unsafe.Add(nil, bp+112)) {
					v2 = q21
					v3 = p2
					v47 = (fq2 - *(*float32)(unsafe.Add(nil, bp+112))) / (fq2 - fp2)
					v5 = Vec2{
						X: float32((float32FromFloat32(1)-v47)*v2.X) + float32(v47*v3.X),
						Y: float32((float32FromFloat32(1)-v47)*v2.Y) + float32(v47*v3.Y),
					}
					goto _211
				_211:
					cq = v5
				}
			}
			v2 = cp
			v3 = p1
			v5 = Vec2{
				X: v2.X - v3.X,
				Y: v2.Y - v3.Y,
			}
			goto _215
		_215:
			v6 = v5
			v7 = normalA
			v47 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y)
			goto _219
		_219:
			sp = v47
			v15 = cq
			v16 = p1
			v17 = Vec2{
				X: v15.X - v16.X,
				Y: v15.Y - v16.Y,
			}
			goto _223
		_223:
			v20 = v17
			v21 = normalA
			v51 = float32(v20.X*v21.X) + float32(v20.Y*v21.Y)
			goto _227
		_227:
			sq = v51
			if sp <= distance+float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter) || sq <= distance+float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter) {
				mp = bp + 12 + uintptr(0)*48
				v2 = cp
				v47 = float32(float32FromFloat32(0.5) * (radiusA - radiusB - sp))
				v3 = normalA
				v5 = Vec2{
					X: v2.X + float32(v47*v3.X),
					Y: v2.Y + float32(v47*v3.Y),
				}
				goto _232
			_232:
				(*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA = v5
				(*ManifoldPoint)(unsafe.Pointer(mp)).Separation = sp - radius
				(*ManifoldPoint)(unsafe.Pointer(mp)).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(int32FromInt32(0)))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(int32FromInt32(0))))
				mp = bp + 12 + uintptr(1)*48
				v2 = cq
				v47 = float32(float32FromFloat32(0.5) * (radiusA - radiusB - sq))
				v3 = normalA
				v5 = Vec2{
					X: v2.X + float32(v47*v3.X),
					Y: v2.Y + float32(v47*v3.Y),
				}
				goto _237
			_237:
				(*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA = v5
				(*ManifoldPoint)(unsafe.Pointer(mp)).Separation = sq - radius
				(*ManifoldPoint)(unsafe.Pointer(mp)).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(int32FromInt32(0)))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(int32FromInt32(1))))
				(*(*Manifold)(unsafe.Add(nil, bp))).PointCount = int32(2)
			}
		} else {
			// normal always points from A to B
			v2 = normalB
			v3 = Vec2{
				X: -v2.X,
				Y: -v2.Y,
			}
			goto _240
		_240:
			(*(*Manifold)(unsafe.Add(nil, bp))).Normal = v3
			cp1 = p1
			cq1 = q11
			// clip to p2
			if fp1 < float32FromFloat32(0) && fq1 > float32FromFloat32(0) {
				v2 = p1
				v3 = q11
				v47 = (float32FromFloat32(0) - fp1) / (fq1 - fp1)
				v5 = Vec2{
					X: float32((float32FromFloat32(1)-v47)*v2.X) + float32(v47*v3.X),
					Y: float32((float32FromFloat32(1)-v47)*v2.Y) + float32(v47*v3.Y),
				}
				goto _245
			_245:
				cp1 = v5
			} else {
				if fq1 < float32FromFloat32(0) && fp1 > float32FromFloat32(0) {
					v2 = q11
					v3 = p1
					v47 = (float32FromFloat32(0) - fq1) / (fp1 - fq1)
					v5 = Vec2{
						X: float32((float32FromFloat32(1)-v47)*v2.X) + float32(v47*v3.X),
						Y: float32((float32FromFloat32(1)-v47)*v2.Y) + float32(v47*v3.Y),
					}
					goto _250
				_250:
					cq1 = v5
				}
			}
			// clip to q2
			if fp1 > *(*float32)(unsafe.Add(nil, bp+116)) && fq1 < *(*float32)(unsafe.Add(nil, bp+116)) {
				v2 = p1
				v3 = q11
				v47 = (fp1 - *(*float32)(unsafe.Add(nil, bp+116))) / (fp1 - fq1)
				v5 = Vec2{
					X: float32((float32FromFloat32(1)-v47)*v2.X) + float32(v47*v3.X),
					Y: float32((float32FromFloat32(1)-v47)*v2.Y) + float32(v47*v3.Y),
				}
				goto _255
			_255:
				cp1 = v5
			} else {
				if fq1 > *(*float32)(unsafe.Add(nil, bp+116)) && fp1 < *(*float32)(unsafe.Add(nil, bp+116)) {
					v2 = q11
					v3 = p1
					v47 = (fq1 - *(*float32)(unsafe.Add(nil, bp+116))) / (fq1 - fp1)
					v5 = Vec2{
						X: float32((float32FromFloat32(1)-v47)*v2.X) + float32(v47*v3.X),
						Y: float32((float32FromFloat32(1)-v47)*v2.Y) + float32(v47*v3.Y),
					}
					goto _260
				_260:
					cq1 = v5
				}
			}
			v2 = cp1
			v3 = p2
			v5 = Vec2{
				X: v2.X - v3.X,
				Y: v2.Y - v3.Y,
			}
			goto _264
		_264:
			v6 = v5
			v7 = normalB
			v47 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y)
			goto _268
		_268:
			sp1 = v47
			v15 = cq1
			v16 = p2
			v17 = Vec2{
				X: v15.X - v16.X,
				Y: v15.Y - v16.Y,
			}
			goto _272
		_272:
			v20 = v17
			v21 = normalB
			v51 = float32(v20.X*v21.X) + float32(v20.Y*v21.Y)
			goto _276
		_276:
			sq1 = v51
			if sp1 <= distance+float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter) || sq1 <= distance+float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter) {
				mp1 = bp + 12 + uintptr(0)*48
				v2 = cp1
				v47 = float32(float32FromFloat32(0.5) * (radiusB - radiusA - sp1))
				v3 = normalB
				v5 = Vec2{
					X: v2.X + float32(v47*v3.X),
					Y: v2.Y + float32(v47*v3.Y),
				}
				goto _281
			_281:
				(*ManifoldPoint)(unsafe.Pointer(mp1)).AnchorA = v5
				(*ManifoldPoint)(unsafe.Pointer(mp1)).Separation = sp1 - radius
				(*ManifoldPoint)(unsafe.Pointer(mp1)).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(int32FromInt32(0)))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(int32FromInt32(0))))
				mp1 = bp + 12 + uintptr(1)*48
				v2 = cq1
				v47 = float32(float32FromFloat32(0.5) * (radiusB - radiusA - sq1))
				v3 = normalB
				v5 = Vec2{
					X: v2.X + float32(v47*v3.X),
					Y: v2.Y + float32(v47*v3.Y),
				}
				goto _286
			_286:
				(*ManifoldPoint)(unsafe.Pointer(mp1)).AnchorA = v5
				(*ManifoldPoint)(unsafe.Pointer(mp1)).Separation = sq1 - radius
				(*ManifoldPoint)(unsafe.Pointer(mp1)).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(int32FromInt32(1)))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(int32FromInt32(0))))
				(*(*Manifold)(unsafe.Add(nil, bp))).PointCount = int32(2)
			}
		}
	}
	if (*(*Manifold)(unsafe.Add(nil, bp))).PointCount == 0 {
		v2 = closest2
		v3 = closest1
		v5 = Vec2{
			X: v2.X - v3.X,
			Y: v2.Y - v3.Y,
		}
		goto _290
	_290:
		// single point collision
		normal = v5
		v6 = normal
		v7 = normal
		v47 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y)
		goto _294
	_294:
		if v47 > epsSqr {
			v15 = normal
			length = sqrtf(tls, float32(v15.X*v15.X)+float32(v15.Y*v15.Y))
			if length < float32FromFloat32(1.1920928955078125e-07) {
				v16 = Vec2{}
				goto _297
			}
			invLength = float32FromFloat32(1) / length
			n = Vec2{
				X: float32(invLength * v15.X),
				Y: float32(invLength * v15.Y),
			}
			v16 = n
			goto _297
		_297:
			normal = v16
		} else {
			v2 = u1
			v3 = Vec2{
				X: -v2.Y,
				Y: v2.X,
			}
			goto _300
		_300:
			normal = v3
		}
		v2 = closest1
		v47 = radiusA
		v3 = normal
		v5 = Vec2{
			X: v2.X + float32(v47*v3.X),
			Y: v2.Y + float32(v47*v3.Y),
		}
		goto _305
	_305:
		c1 = v5
		v6 = closest2
		v51 = -radiusB
		v7 = normal
		v15 = Vec2{
			X: v6.X + float32(v51*v7.X),
			Y: v6.Y + float32(v51*v7.Y),
		}
		goto _310
	_310:
		c2 = v15
		if f1 == float32FromFloat32(0) {
			v311 = 0
		} else {
			v311 = int32(1)
		}
		i1 = v311
		if f2 == float32FromFloat32(0) {
			v312 = 0
		} else {
			v312 = int32(1)
		}
		i2 = v312
		(*(*Manifold)(unsafe.Add(nil, bp))).Normal = normal
		v2 = c1
		v3 = c2
		v47 = float32FromFloat32(0.5)
		v5 = Vec2{
			X: float32((float32FromFloat32(1)-v47)*v2.X) + float32(v47*v3.X),
			Y: float32((float32FromFloat32(1)-v47)*v2.Y) + float32(v47*v3.Y),
		}
		goto _317
	_317:
		(*(*ManifoldPoint)(unsafe.Add(nil, bp+12))).AnchorA = v5
		(*(*ManifoldPoint)(unsafe.Add(nil, bp+12))).Separation = sqrtf(tls, distanceSquared) - radius
		(*(*ManifoldPoint)(unsafe.Add(nil, bp+12))).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(i1))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(i2)))
		(*(*Manifold)(unsafe.Add(nil, bp))).PointCount = int32(1)
	}
	// Convert manifold to world space
	v1 = xfA.Q
	v2 = (*(*Manifold)(unsafe.Add(nil, bp))).Normal
	v3 = Vec2{
		X: float32(v1.C*v2.X) - float32(v1.S*v2.Y),
		Y: float32(v1.S*v2.X) + float32(v1.C*v2.Y),
	}
	goto _321
_321:
	(*(*Manifold)(unsafe.Add(nil, bp))).Normal = v3
	i = 0
	for {
		if !(i < (*(*Manifold)(unsafe.Add(nil, bp))).PointCount) {
			break
		}
		mp2 = bp + 12 + uintptr(i)*48
		// anchor points relative to shape origin in world space
		v2 = (*ManifoldPoint)(unsafe.Pointer(mp2)).AnchorA
		v3 = origin
		v5 = Vec2{
			X: v2.X + v3.X,
			Y: v2.Y + v3.Y,
		}
		goto _326
	_326:
		v1 = xfA.Q
		v6 = v5
		v7 = Vec2{
			X: float32(v1.C*v6.X) - float32(v1.S*v6.Y),
			Y: float32(v1.S*v6.X) + float32(v1.C*v6.Y),
		}
		goto _330
	_330:
		(*ManifoldPoint)(unsafe.Pointer(mp2)).AnchorA = v7
		v2 = xfA.P
		v3 = xfB.P
		v5 = Vec2{
			X: v2.X - v3.X,
			Y: v2.Y - v3.Y,
		}
		goto _334
	_334:
		v6 = (*ManifoldPoint)(unsafe.Pointer(mp2)).AnchorA
		v7 = v5
		v15 = Vec2{
			X: v6.X + v7.X,
			Y: v6.Y + v7.Y,
		}
		goto _338
	_338:
		(*ManifoldPoint)(unsafe.Pointer(mp2)).AnchorB = v15
		v2 = xfA.P
		v3 = (*ManifoldPoint)(unsafe.Pointer(mp2)).AnchorA
		v5 = Vec2{
			X: v2.X + v3.X,
			Y: v2.Y + v3.Y,
		}
		goto _342
	_342:
		(*ManifoldPoint)(unsafe.Pointer(mp2)).Point = v5
		goto _322
	_322:
		;
		i = i + 1
	}
	return *(*Manifold)(unsafe.Add(nil, bp))
}

func b2CollideSegmentAndCapsule(tls *_Stack, segmentA uintptr, xfA Transform, capsuleB uintptr, xfB Transform) (r Manifold) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* capsuleA at bp+0 */ Capsule
	*(*Capsule)(unsafe.Add(nil, bp)) = Capsule{
		Center1: (*Segment)(unsafe.Pointer(segmentA)).Point1,
		Center2: (*Segment)(unsafe.Pointer(segmentA)).Point2,
	}
	return b2CollideCapsules(tls, bp, xfA, capsuleB, xfB)
}

func b2CollidePolygonAndCapsule(tls *_Stack, polygonA uintptr, xfA Transform, capsuleB uintptr, xfB Transform) (r Manifold) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var _ /* polyB at bp+0 */ Polygon
	*(*Polygon)(unsafe.Add(nil, bp)) = b2MakeCapsule(tls, (*Capsule)(unsafe.Pointer(capsuleB)).Center1, (*Capsule)(unsafe.Pointer(capsuleB)).Center2, (*Capsule)(unsafe.Pointer(capsuleB)).Radius)
	return b2CollidePolygons(tls, polygonA, xfA, bp, xfB)
}

// Due to speculation, every polygon is rounded
// Algorithm:
//
// compute edge separation using the separating axis test (SAT)
// if (separation > speculation_distance)
//   return
// find reference and incident edge
// if separation >= 0.1f * B2_LINEAR_SLOP
//   compute closest points between reference and incident edge
//   if vertices are closest
//      single vertex-vertex contact
//   else
//      clip edges
//   end
// else
//   clip edges
// end

func b2CollidePolygons(tls *_Stack, polygonA uintptr, xfA Transform, polygonB uintptr, xfB Transform) (r Manifold) {
	bp := tls.Alloc(416)
	defer tls.Free(416)
	var C, sfA, xf, v10, v23, v9 Transform
	var c1, c11, c12, c13, c2, c21, c22, c23, normal, normal1, normal2, normal3, origin, searchDirection, searchDirection1, v11, v12, v21, v22, v15, v16, v17, v2, v20, v211, v3, v5, v6, v7 Vec2
	var count, count1, i, i1, i11, i12, i2, i21, i22, i3, i4, i5, v49, v50 int32
	var distance, dot, dot1, invDistance, invDistance1, invDistance2, invDistance3, linearSlop, minDot, minDot1, minSeparation, radius, separation, separationA, separationB, speculativeDistance, x, y, v42, v47, v52, v53 float32
	var flip uint8
	var mp, normals, normals1 uintptr
	var qr, v1, v111, v121, v13, v19 Rot
	var result SegmentDistanceResult
	var _ /* edgeA at bp+288 */ int32
	var _ /* edgeB at bp+292 */ int32
	var _ /* localPolyA at bp+0 */ Polygon
	var _ /* localPolyB at bp+144 */ Polygon
	var _ /* manifold at bp+296 */ Manifold
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, c1, c11, c12, c13, c2, c21, c22, c23, count, count1, distance, dot, dot1, flip, i, i1, i11, i12, i2, i21, i22, i3, i4, i5, invDistance, invDistance1, invDistance2, invDistance3, linearSlop, minDot, minDot1, minSeparation, mp, normal, normal1, normal2, normal3, normals, normals1, origin, qr, radius, result, searchDirection, searchDirection1, separation, separationA, separationB, sfA, speculativeDistance, v11, v12, v21, v22, x, xf, y, v1, v10, v111, v121, v13, v15, v16, v17, v19, v2, v20, v211, v23, v3, v42, v47, v49, v5, v50, v52, v53, v6, v7, v9
	origin = *(*Vec2)(unsafe.Add(nil, polygonA))
	linearSlop = float32(float32FromFloat32(0.005) * b2_lengthUnitsPerMeter)
	speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter))
	v1 = xfA.Q
	v2 = origin
	v3 = Vec2{
		X: float32(v1.C*v2.X) - float32(v1.S*v2.Y),
		Y: float32(v1.S*v2.X) + float32(v1.C*v2.Y),
	}
	goto _4
_4:
	v5 = xfA.P
	v6 = v3
	v7 = Vec2{
		X: v5.X + v6.X,
		Y: v5.Y + v6.Y,
	}
	goto _8
_8:
	// Shift polyA to origin
	// pw = q * pb + p
	// pw = q * (pbs + origin) + p
	// pw = q * pbs + (p + q * origin)
	sfA = Transform{
		P: v7,
		Q: xfA.Q,
	}
	v9 = sfA
	v10 = xfB
	v111 = v9.Q
	v121 = v10.Q
	qr.S = float32(v111.C*v121.S) - float32(v111.S*v121.C)
	qr.C = float32(v111.C*v121.C) + float32(v111.S*v121.S)
	v13 = qr
	goto _14
_14:
	C.Q = v13
	v15 = v10.P
	v16 = v9.P
	v17 = Vec2{
		X: v15.X - v16.X,
		Y: v15.Y - v16.Y,
	}
	goto _18
_18:
	v19 = v9.Q
	v20 = v17
	v211 = Vec2{
		X: float32(v19.C*v20.X) + float32(v19.S*v20.Y),
		Y: float32(-v19.S*v20.X) + float32(v19.C*v20.Y),
	}
	goto _22
_22:
	C.P = v211
	v23 = C
	goto _24
_24:
	xf = v23
	(*(*Polygon)(unsafe.Add(nil, bp))).Count = (*Polygon)(unsafe.Pointer(polygonA)).Count
	(*(*Polygon)(unsafe.Add(nil, bp))).Radius = (*Polygon)(unsafe.Pointer(polygonA)).Radius
	*(*Vec2)(unsafe.Add(nil, bp)) = b2Vec2_zero
	*(*Vec2)(unsafe.Add(nil, bp+64)) = *(*Vec2)(unsafe.Add(nil, polygonA+64))
	i = int32(1)
	for {
		if !(i < (*(*Polygon)(unsafe.Add(nil, bp))).Count) {
			break
		}
		v2 = *(*Vec2)(unsafe.Add(nil, polygonA+uintptr(i)*8))
		v3 = origin
		v5 = Vec2{
			X: v2.X - v3.X,
			Y: v2.Y - v3.Y,
		}
		goto _29
	_29:
		*(*Vec2)(unsafe.Add(nil, bp+uintptr(i)*8)) = v5
		*(*Vec2)(unsafe.Add(nil, bp+64+uintptr(i)*8)) = *(*Vec2)(unsafe.Add(nil, polygonA+64+uintptr(i)*8))
		goto _25
	_25:
		;
		i = i + 1
	}
	(*(*Polygon)(unsafe.Add(nil, bp+144))).Count = (*Polygon)(unsafe.Pointer(polygonB)).Count
	(*(*Polygon)(unsafe.Add(nil, bp+144))).Radius = (*Polygon)(unsafe.Pointer(polygonB)).Radius
	i1 = 0
	for {
		if !(i1 < (*(*Polygon)(unsafe.Add(nil, bp+144))).Count) {
			break
		}
		v9 = xf
		v2 = *(*Vec2)(unsafe.Add(nil, polygonB+uintptr(i1)*8))
		x = float32(v9.Q.C*v2.X) - float32(v9.Q.S*v2.Y) + v9.P.X
		y = float32(v9.Q.S*v2.X) + float32(v9.Q.C*v2.Y) + v9.P.Y
		v3 = Vec2{
			X: x,
			Y: y,
		}
		goto _34
	_34:
		*(*Vec2)(unsafe.Add(nil, bp+144+uintptr(i1)*8)) = v3
		v1 = xf.Q
		v2 = *(*Vec2)(unsafe.Add(nil, polygonB+64+uintptr(i1)*8))
		v3 = Vec2{
			X: float32(v1.C*v2.X) - float32(v1.S*v2.Y),
			Y: float32(v1.S*v2.X) + float32(v1.C*v2.Y),
		}
		goto _38
	_38:
		*(*Vec2)(unsafe.Add(nil, bp+144+64+uintptr(i1)*8)) = v3
		goto _30
	_30:
		;
		i1 = i1 + 1
	}
	*(*int32)(unsafe.Add(nil, bp+288)) = 0
	separationA = b2FindMaxSeparation(tls, bp+288, bp, bp+144)
	*(*int32)(unsafe.Add(nil, bp+292)) = 0
	separationB = b2FindMaxSeparation(tls, bp+292, bp+144, bp)
	radius = (*(*Polygon)(unsafe.Add(nil, bp))).Radius + (*(*Polygon)(unsafe.Add(nil, bp+144))).Radius
	if separationA > speculativeDistance+radius || separationB > speculativeDistance+radius {
		return Manifold{}
	}
	if separationA >= separationB {
		flip = boolUint8(false1 != 0)
		searchDirection = *(*Vec2)(unsafe.Add(nil, bp+64+uintptr(*(*int32)(unsafe.Add(nil, bp+288)))*8))
		// Find the incident edge on polyB
		count = (*(*Polygon)(unsafe.Add(nil, bp+144))).Count
		normals = bp + 144 + 64
		*(*int32)(unsafe.Add(nil, bp+292)) = 0
		minDot = float32FromFloat32(3.4028234663852886e+38)
		i2 = 0
		for {
			if !(i2 < count) {
				break
			}
			v2 = searchDirection
			v3 = *(*Vec2)(unsafe.Add(nil, normals+uintptr(i2)*8))
			v42 = float32(v2.X*v3.X) + float32(v2.Y*v3.Y)
			goto _43
		_43:
			dot = v42
			if dot < minDot {
				minDot = dot
				*(*int32)(unsafe.Add(nil, bp+292)) = i2
			}
			goto _39
		_39:
			;
			i2 = i2 + 1
		}
	} else {
		flip = boolUint8(true1 != 0)
		searchDirection1 = *(*Vec2)(unsafe.Add(nil, bp+144+64+uintptr(*(*int32)(unsafe.Add(nil, bp+292)))*8))
		// Find the incident edge on polyA
		count1 = (*(*Polygon)(unsafe.Add(nil, bp))).Count
		normals1 = bp + 64
		*(*int32)(unsafe.Add(nil, bp+288)) = 0
		minDot1 = float32FromFloat32(3.4028234663852886e+38)
		i3 = 0
		for {
			if !(i3 < count1) {
				break
			}
			v2 = searchDirection1
			v3 = *(*Vec2)(unsafe.Add(nil, normals1+uintptr(i3)*8))
			v42 = float32(v2.X*v3.X) + float32(v2.Y*v3.Y)
			goto _48
		_48:
			dot1 = v42
			if dot1 < minDot1 {
				minDot1 = dot1
				*(*int32)(unsafe.Add(nil, bp+288)) = i3
			}
			goto _44
		_44:
			;
			i3 = i3 + 1
		}
	}
	*(*Manifold)(unsafe.Add(nil, bp+296)) = Manifold{}
	// Using slop here to ensure vertex-vertex normal vectors can be safely normalized
	// todo this means edge clipping needs to handle slightly non-overlapping edges.
	if separationA > float32(float32FromFloat32(0.1)*linearSlop) || separationB > float32(float32FromFloat32(0.1)*linearSlop) {
		// Edges are disjoint. Find closest points between reference edge and incident edge
		// Reference edge on polygon A
		i11 = *(*int32)(unsafe.Add(nil, bp+288))
		if *(*int32)(unsafe.Add(nil, bp+288))+int32(1) < (*(*Polygon)(unsafe.Add(nil, bp))).Count {
			v49 = *(*int32)(unsafe.Add(nil, bp+288)) + int32(1)
		} else {
			v49 = 0
		}
		i12 = v49
		i21 = *(*int32)(unsafe.Add(nil, bp+292))
		if *(*int32)(unsafe.Add(nil, bp+292))+int32(1) < (*(*Polygon)(unsafe.Add(nil, bp+144))).Count {
			v50 = *(*int32)(unsafe.Add(nil, bp+292)) + int32(1)
		} else {
			v50 = 0
		}
		i22 = v50
		v11 = *(*Vec2)(unsafe.Add(nil, bp+uintptr(i11)*8))
		v12 = *(*Vec2)(unsafe.Add(nil, bp+uintptr(i12)*8))
		v21 = *(*Vec2)(unsafe.Add(nil, bp+144+uintptr(i21)*8))
		v22 = *(*Vec2)(unsafe.Add(nil, bp+144+uintptr(i22)*8))
		result = b2SegmentDistance(tls, v11, v12, v21, v22)
		if !(result.DistanceSquared > float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+9991, __ccgo_ts+9936, int32FromInt32(848)) != 0 {
			__builtin_trap(tls)
		}
		distance = sqrtf(tls, result.DistanceSquared)
		separation = distance - radius
		if distance-radius > speculativeDistance {
			// This can happen in the vertex-vertex case
			return *(*Manifold)(unsafe.Add(nil, bp+296))
		}
		// Attempt to clip edges
		*(*Manifold)(unsafe.Add(nil, bp+296)) = b2ClipPolygons(tls, bp, bp+144, *(*int32)(unsafe.Add(nil, bp+288)), *(*int32)(unsafe.Add(nil, bp+292)), flip)
		minSeparation = float32FromFloat32(3.4028234663852886e+38)
		i4 = 0
		for {
			if !(i4 < (*(*Manifold)(unsafe.Add(nil, bp+296))).PointCount) {
				break
			}
			v42 = minSeparation
			v47 = (*(*ManifoldPoint)(unsafe.Add(nil, bp+296+12+uintptr(i4)*48))).Separation
			if v42 < v47 {
				v53 = v42
			} else {
				v53 = v47
			}
			v52 = v53
			goto _55
		_55:
			minSeparation = v52
			goto _51
		_51:
			;
			i4 = i4 + 1
		}
		// Does vertex-vertex have substantially larger separation?
		if separation+float32(float32FromFloat32(0.1)*linearSlop) < minSeparation {
			if result.Fraction1 == float32FromFloat32(0) && result.Fraction2 == float32FromFloat32(0) {
				v2 = v21
				v3 = v11
				v5 = Vec2{
					X: v2.X - v3.X,
					Y: v2.Y - v3.Y,
				}
				goto _60
			_60:
				// v11 - v21
				normal = v5
				invDistance = float32FromFloat32(1) / distance
				normal.X = float32(normal.X * invDistance)
				normal.Y = float32(normal.Y * invDistance)
				v2 = v11
				v42 = (*(*Polygon)(unsafe.Add(nil, bp))).Radius
				v3 = normal
				v5 = Vec2{
					X: v2.X + float32(v42*v3.X),
					Y: v2.Y + float32(v42*v3.Y),
				}
				goto _65
			_65:
				c1 = v5
				v6 = v21
				v47 = -(*(*Polygon)(unsafe.Add(nil, bp+144))).Radius
				v7 = normal
				v15 = Vec2{
					X: v6.X + float32(v47*v7.X),
					Y: v6.Y + float32(v47*v7.Y),
				}
				goto _70
			_70:
				c2 = v15
				(*(*Manifold)(unsafe.Add(nil, bp+296))).Normal = normal
				v2 = c1
				v3 = c2
				v42 = float32FromFloat32(0.5)
				v5 = Vec2{
					X: float32((float32FromFloat32(1)-v42)*v2.X) + float32(v42*v3.X),
					Y: float32((float32FromFloat32(1)-v42)*v2.Y) + float32(v42*v3.Y),
				}
				goto _75
			_75:
				(*(*ManifoldPoint)(unsafe.Add(nil, bp+296+12))).AnchorA = v5
				(*(*ManifoldPoint)(unsafe.Add(nil, bp+296+12))).Separation = distance - radius
				(*(*ManifoldPoint)(unsafe.Add(nil, bp+296+12))).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(i11))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(i21)))
				(*(*Manifold)(unsafe.Add(nil, bp+296))).PointCount = int32(1)
			} else {
				if result.Fraction1 == float32FromFloat32(0) && result.Fraction2 == float32FromFloat32(1) {
					v2 = v22
					v3 = v11
					v5 = Vec2{
						X: v2.X - v3.X,
						Y: v2.Y - v3.Y,
					}
					goto _79
				_79:
					// v11 - v22
					normal1 = v5
					invDistance1 = float32FromFloat32(1) / distance
					normal1.X = float32(normal1.X * invDistance1)
					normal1.Y = float32(normal1.Y * invDistance1)
					v2 = v11
					v42 = (*(*Polygon)(unsafe.Add(nil, bp))).Radius
					v3 = normal1
					v5 = Vec2{
						X: v2.X + float32(v42*v3.X),
						Y: v2.Y + float32(v42*v3.Y),
					}
					goto _84
				_84:
					c11 = v5
					v6 = v22
					v47 = -(*(*Polygon)(unsafe.Add(nil, bp+144))).Radius
					v7 = normal1
					v15 = Vec2{
						X: v6.X + float32(v47*v7.X),
						Y: v6.Y + float32(v47*v7.Y),
					}
					goto _89
				_89:
					c21 = v15
					(*(*Manifold)(unsafe.Add(nil, bp+296))).Normal = normal1
					v2 = c11
					v3 = c21
					v42 = float32FromFloat32(0.5)
					v5 = Vec2{
						X: float32((float32FromFloat32(1)-v42)*v2.X) + float32(v42*v3.X),
						Y: float32((float32FromFloat32(1)-v42)*v2.Y) + float32(v42*v3.Y),
					}
					goto _94
				_94:
					(*(*ManifoldPoint)(unsafe.Add(nil, bp+296+12))).AnchorA = v5
					(*(*ManifoldPoint)(unsafe.Add(nil, bp+296+12))).Separation = distance - radius
					(*(*ManifoldPoint)(unsafe.Add(nil, bp+296+12))).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(i11))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(i22)))
					(*(*Manifold)(unsafe.Add(nil, bp+296))).PointCount = int32(1)
				} else {
					if result.Fraction1 == float32FromFloat32(1) && result.Fraction2 == float32FromFloat32(0) {
						v2 = v21
						v3 = v12
						v5 = Vec2{
							X: v2.X - v3.X,
							Y: v2.Y - v3.Y,
						}
						goto _98
					_98:
						// v12 - v21
						normal2 = v5
						invDistance2 = float32FromFloat32(1) / distance
						normal2.X = float32(normal2.X * invDistance2)
						normal2.Y = float32(normal2.Y * invDistance2)
						v2 = v12
						v42 = (*(*Polygon)(unsafe.Add(nil, bp))).Radius
						v3 = normal2
						v5 = Vec2{
							X: v2.X + float32(v42*v3.X),
							Y: v2.Y + float32(v42*v3.Y),
						}
						goto _103
					_103:
						c12 = v5
						v6 = v21
						v47 = -(*(*Polygon)(unsafe.Add(nil, bp+144))).Radius
						v7 = normal2
						v15 = Vec2{
							X: v6.X + float32(v47*v7.X),
							Y: v6.Y + float32(v47*v7.Y),
						}
						goto _108
					_108:
						c22 = v15
						(*(*Manifold)(unsafe.Add(nil, bp+296))).Normal = normal2
						v2 = c12
						v3 = c22
						v42 = float32FromFloat32(0.5)
						v5 = Vec2{
							X: float32((float32FromFloat32(1)-v42)*v2.X) + float32(v42*v3.X),
							Y: float32((float32FromFloat32(1)-v42)*v2.Y) + float32(v42*v3.Y),
						}
						goto _113
					_113:
						(*(*ManifoldPoint)(unsafe.Add(nil, bp+296+12))).AnchorA = v5
						(*(*ManifoldPoint)(unsafe.Add(nil, bp+296+12))).Separation = distance - radius
						(*(*ManifoldPoint)(unsafe.Add(nil, bp+296+12))).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(i12))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(i21)))
						(*(*Manifold)(unsafe.Add(nil, bp+296))).PointCount = int32(1)
					} else {
						if result.Fraction1 == float32FromFloat32(1) && result.Fraction2 == float32FromFloat32(1) {
							v2 = v22
							v3 = v12
							v5 = Vec2{
								X: v2.X - v3.X,
								Y: v2.Y - v3.Y,
							}
							goto _117
						_117:
							// v12 - v22
							normal3 = v5
							invDistance3 = float32FromFloat32(1) / distance
							normal3.X = float32(normal3.X * invDistance3)
							normal3.Y = float32(normal3.Y * invDistance3)
							v2 = v12
							v42 = (*(*Polygon)(unsafe.Add(nil, bp))).Radius
							v3 = normal3
							v5 = Vec2{
								X: v2.X + float32(v42*v3.X),
								Y: v2.Y + float32(v42*v3.Y),
							}
							goto _122
						_122:
							c13 = v5
							v6 = v22
							v47 = -(*(*Polygon)(unsafe.Add(nil, bp+144))).Radius
							v7 = normal3
							v15 = Vec2{
								X: v6.X + float32(v47*v7.X),
								Y: v6.Y + float32(v47*v7.Y),
							}
							goto _127
						_127:
							c23 = v15
							(*(*Manifold)(unsafe.Add(nil, bp+296))).Normal = normal3
							v2 = c13
							v3 = c23
							v42 = float32FromFloat32(0.5)
							v5 = Vec2{
								X: float32((float32FromFloat32(1)-v42)*v2.X) + float32(v42*v3.X),
								Y: float32((float32FromFloat32(1)-v42)*v2.Y) + float32(v42*v3.Y),
							}
							goto _132
						_132:
							(*(*ManifoldPoint)(unsafe.Add(nil, bp+296+12))).AnchorA = v5
							(*(*ManifoldPoint)(unsafe.Add(nil, bp+296+12))).Separation = distance - radius
							(*(*ManifoldPoint)(unsafe.Add(nil, bp+296+12))).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(i12))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(i22)))
							(*(*Manifold)(unsafe.Add(nil, bp+296))).PointCount = int32(1)
						}
					}
				}
			}
		}
	} else {
		// Polygons overlap
		*(*Manifold)(unsafe.Add(nil, bp+296)) = b2ClipPolygons(tls, bp, bp+144, *(*int32)(unsafe.Add(nil, bp+288)), *(*int32)(unsafe.Add(nil, bp+292)), flip)
	}
	// Convert manifold to world space
	if (*(*Manifold)(unsafe.Add(nil, bp+296))).PointCount > 0 {
		v1 = xfA.Q
		v2 = (*(*Manifold)(unsafe.Add(nil, bp+296))).Normal
		v3 = Vec2{
			X: float32(v1.C*v2.X) - float32(v1.S*v2.Y),
			Y: float32(v1.S*v2.X) + float32(v1.C*v2.Y),
		}
		goto _136
	_136:
		(*(*Manifold)(unsafe.Add(nil, bp+296))).Normal = v3
		i5 = 0
		for {
			if !(i5 < (*(*Manifold)(unsafe.Add(nil, bp+296))).PointCount) {
				break
			}
			mp = bp + 296 + 12 + uintptr(i5)*48
			// anchor points relative to shape origin in world space
			v2 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
			v3 = origin
			v5 = Vec2{
				X: v2.X + v3.X,
				Y: v2.Y + v3.Y,
			}
			goto _141
		_141:
			v1 = xfA.Q
			v6 = v5
			v7 = Vec2{
				X: float32(v1.C*v6.X) - float32(v1.S*v6.Y),
				Y: float32(v1.S*v6.X) + float32(v1.C*v6.Y),
			}
			goto _145
		_145:
			(*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA = v7
			v2 = xfA.P
			v3 = xfB.P
			v5 = Vec2{
				X: v2.X - v3.X,
				Y: v2.Y - v3.Y,
			}
			goto _149
		_149:
			v6 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
			v7 = v5
			v15 = Vec2{
				X: v6.X + v7.X,
				Y: v6.Y + v7.Y,
			}
			goto _153
		_153:
			(*ManifoldPoint)(unsafe.Pointer(mp)).AnchorB = v15
			v2 = xfA.P
			v3 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
			v5 = Vec2{
				X: v2.X + v3.X,
				Y: v2.Y + v3.Y,
			}
			goto _157
		_157:
			(*ManifoldPoint)(unsafe.Pointer(mp)).Point = v5
			goto _137
		_137:
			;
			i5 = i5 + 1
		}
	}
	return *(*Manifold)(unsafe.Add(nil, bp+296))
}

func b2CollideSegmentAndCircle(tls *_Stack, segmentA uintptr, xfA Transform, circleB uintptr, xfB Transform) (r Manifold) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* capsuleA at bp+0 */ Capsule
	*(*Capsule)(unsafe.Add(nil, bp)) = Capsule{
		Center1: (*Segment)(unsafe.Pointer(segmentA)).Point1,
		Center2: (*Segment)(unsafe.Pointer(segmentA)).Point2,
	}
	return b2CollideCapsuleAndCircle(tls, bp, xfA, circleB, xfB)
}

func b2CollideSegmentAndPolygon(tls *_Stack, segmentA uintptr, xfA Transform, polygonB uintptr, xfB Transform) (r Manifold) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var _ /* polygonA at bp+0 */ Polygon
	*(*Polygon)(unsafe.Add(nil, bp)) = b2MakeCapsule(tls, (*Segment)(unsafe.Pointer(segmentA)).Point1, (*Segment)(unsafe.Pointer(segmentA)).Point2, float32FromFloat32(0))
	return b2CollidePolygons(tls, bp, xfA, polygonB, xfB)
}

func b2CollideChainSegmentAndCircle(tls *_Stack, segmentA uintptr, xfA Transform, circleB uintptr, xfB Transform) (r Manifold) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var C, xf, v1, v15, v17, v2 Transform
	var cA, cB, contactPointA, e, n, nextEdge, normal, p1, p2, pA, pB, prevEdge, v12, v13, v18, v19, v21, v22, v23, v25, v26, v28, v29, v30, v32, v33, v36, v7, v8, v9 Vec2
	var ee, invLength, offset, radius, separation, u, uPrev, v5, vNext, x, y, v34, v42, v50 float32
	var mp, v89 uintptr
	var qr, v11, v3, v4, v51 Rot
	var _ /* distance at bp+112 */ float32
	var _ /* manifold at bp+0 */ Manifold
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, cA, cB, contactPointA, e, ee, invLength, mp, n, nextEdge, normal, offset, p1, p2, pA, pB, prevEdge, qr, radius, separation, u, uPrev, v5, vNext, x, xf, y, v1, v11, v12, v13, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v28, v29, v3, v30, v32, v33, v34, v36, v4, v42, v51, v50, v7, v8, v89, v9
	*(*Manifold)(unsafe.Add(nil, bp)) = Manifold{}
	v1 = xfA
	v2 = xfB
	v3 = v1.Q
	v4 = v2.Q
	qr.S = float32(v3.C*v4.S) - float32(v3.S*v4.C)
	qr.C = float32(v3.C*v4.C) + float32(v3.S*v4.S)
	v51 = qr
	goto _6
_6:
	C.Q = v51
	v7 = v2.P
	v8 = v1.P
	v9 = Vec2{
		X: v7.X - v8.X,
		Y: v7.Y - v8.Y,
	}
	goto _10
_10:
	v11 = v1.Q
	v12 = v9
	v13 = Vec2{
		X: float32(v11.C*v12.X) + float32(v11.S*v12.Y),
		Y: float32(-v11.S*v12.X) + float32(v11.C*v12.Y),
	}
	goto _14
_14:
	C.P = v13
	v15 = C
	goto _16
_16:
	xf = v15
	v17 = xf
	v18 = (*Circle)(unsafe.Pointer(circleB)).Center
	x = float32(v17.Q.C*v18.X) - float32(v17.Q.S*v18.Y) + v17.P.X
	y = float32(v17.Q.S*v18.X) + float32(v17.Q.C*v18.Y) + v17.P.Y
	v19 = Vec2{
		X: x,
		Y: y,
	}
	goto _20
_20:
	// Compute circle in frame of segment
	pB = v19
	p1 = (*ChainSegment)(unsafe.Pointer(segmentA)).Segment.Point1
	p2 = (*ChainSegment)(unsafe.Pointer(segmentA)).Segment.Point2
	v21 = p2
	v22 = p1
	v23 = Vec2{
		X: v21.X - v22.X,
		Y: v21.Y - v22.Y,
	}
	goto _24
_24:
	e = v23
	v25 = e
	v26 = Vec2{
		X: v25.Y,
		Y: -v25.X,
	}
	goto _27
_27:
	v28 = pB
	v29 = p1
	v30 = Vec2{
		X: v28.X - v29.X,
		Y: v28.Y - v29.Y,
	}
	goto _31
_31:
	v32 = v26
	v33 = v30
	v34 = float32(v32.X*v33.X) + float32(v32.Y*v33.Y)
	goto _35
_35:
	// Normal points to the right
	offset = v34
	if offset < float32FromFloat32(0) {
		// collision is one-sided
		return *(*Manifold)(unsafe.Add(nil, bp))
	}
	v7 = p2
	v8 = pB
	v9 = Vec2{
		X: v7.X - v8.X,
		Y: v7.Y - v8.Y,
	}
	goto _39
_39:
	v12 = e
	v13 = v9
	v34 = float32(v12.X*v13.X) + float32(v12.Y*v13.Y)
	goto _43
_43:
	// Barycentric coordinates
	u = v34
	v18 = pB
	v19 = p1
	v21 = Vec2{
		X: v18.X - v19.X,
		Y: v18.Y - v19.Y,
	}
	goto _47
_47:
	v22 = e
	v23 = v21
	v42 = float32(v22.X*v23.X) + float32(v22.Y*v23.Y)
	goto _51
_51:
	v5 = v42
	if v5 <= float32FromFloat32(0) {
		v25 = p1
		v26 = (*ChainSegment)(unsafe.Pointer(segmentA)).Ghost1
		v28 = Vec2{
			X: v25.X - v26.X,
			Y: v25.Y - v26.Y,
		}
		goto _55
	_55:
		// Behind point1?
		// Is pB in the Voronoi region of the previous edge?
		prevEdge = v28
		v29 = pB
		v30 = p1
		v32 = Vec2{
			X: v29.X - v30.X,
			Y: v29.Y - v30.Y,
		}
		goto _59
	_59:
		v33 = prevEdge
		v36 = v32
		v50 = float32(v33.X*v36.X) + float32(v33.Y*v36.Y)
		goto _63
	_63:
		uPrev = v50
		if uPrev <= float32FromFloat32(0) {
			return *(*Manifold)(unsafe.Add(nil, bp))
		}
		pA = p1
	} else {
		if u <= float32FromFloat32(0) {
			v7 = (*ChainSegment)(unsafe.Pointer(segmentA)).Ghost2
			v8 = p2
			v9 = Vec2{
				X: v7.X - v8.X,
				Y: v7.Y - v8.Y,
			}
			goto _67
		_67:
			// Ahead of point2?
			nextEdge = v9
			v12 = pB
			v13 = p2
			v18 = Vec2{
				X: v12.X - v13.X,
				Y: v12.Y - v13.Y,
			}
			goto _71
		_71:
			v19 = nextEdge
			v21 = v18
			v34 = float32(v19.X*v21.X) + float32(v19.Y*v21.Y)
			goto _75
		_75:
			vNext = v34
			// Is pB in the Voronoi region of the next edge?
			if vNext > float32FromFloat32(0) {
				return *(*Manifold)(unsafe.Add(nil, bp))
			}
			pA = p2
		} else {
			v7 = e
			v8 = e
			v34 = float32(v7.X*v8.X) + float32(v7.Y*v8.Y)
			goto _79
		_79:
			ee = v34
			pA = Vec2{
				X: float32(u*p1.X) + float32(v5*p2.X),
				Y: float32(u*p1.Y) + float32(v5*p2.Y),
			}
			if ee > float32FromFloat32(0) {
				v34 = float32FromFloat32(1) / ee
				v8 = pA
				v9 = Vec2{
					X: float32(v34 * v8.X),
					Y: float32(v34 * v8.Y),
				}
				goto _84
			_84:
				v7 = v9
			} else {
				v7 = p1
			}
			pA = v7
		}
	}
	v7 = pB
	v8 = pA
	v9 = Vec2{
		X: v7.X - v8.X,
		Y: v7.Y - v8.Y,
	}
	goto _88
_88:
	v89 = bp + 112
	v12 = v9
	*(*float32)(unsafe.Add(nil, v89)) = sqrtf(tls, float32(v12.X*v12.X)+float32(v12.Y*v12.Y))
	if *(*float32)(unsafe.Add(nil, v89)) < float32FromFloat32(1.1920928955078125e-07) {
		v13 = Vec2{}
		goto _92
	}
	invLength = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v89))
	n = Vec2{
		X: float32(invLength * v12.X),
		Y: float32(invLength * v12.Y),
	}
	v13 = n
	goto _92
_92:
	normal = v13
	radius = (*Circle)(unsafe.Pointer(circleB)).Radius
	separation = *(*float32)(unsafe.Add(nil, bp+112)) - radius
	if separation > float32(float32FromFloat32(4)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
		return *(*Manifold)(unsafe.Add(nil, bp))
	}
	cA = pA
	v7 = pB
	v34 = -radius
	v8 = normal
	v9 = Vec2{
		X: v7.X + float32(v34*v8.X),
		Y: v7.Y + float32(v34*v8.Y),
	}
	goto _97
_97:
	cB = v9
	v12 = cA
	v13 = cB
	v42 = float32FromFloat32(0.5)
	v18 = Vec2{
		X: float32((float32FromFloat32(1)-v42)*v12.X) + float32(v42*v13.X),
		Y: float32((float32FromFloat32(1)-v42)*v12.Y) + float32(v42*v13.Y),
	}
	goto _102
_102:
	contactPointA = v18
	v3 = xfA.Q
	v19 = normal
	v21 = Vec2{
		X: float32(v3.C*v19.X) - float32(v3.S*v19.Y),
		Y: float32(v3.S*v19.X) + float32(v3.C*v19.Y),
	}
	goto _106
_106:
	(*(*Manifold)(unsafe.Add(nil, bp))).Normal = v21
	mp = bp + 12 + uintptr(0)*48
	v3 = xfA.Q
	v7 = contactPointA
	v8 = Vec2{
		X: float32(v3.C*v7.X) - float32(v3.S*v7.Y),
		Y: float32(v3.S*v7.X) + float32(v3.C*v7.Y),
	}
	goto _110
_110:
	(*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA = v8
	v7 = xfA.P
	v8 = xfB.P
	v9 = Vec2{
		X: v7.X - v8.X,
		Y: v7.Y - v8.Y,
	}
	goto _114
_114:
	v12 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
	v13 = v9
	v18 = Vec2{
		X: v12.X + v13.X,
		Y: v12.Y + v13.Y,
	}
	goto _118
_118:
	(*ManifoldPoint)(unsafe.Pointer(mp)).AnchorB = v18
	v7 = xfA.P
	v8 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
	v9 = Vec2{
		X: v7.X + v8.X,
		Y: v7.Y + v8.Y,
	}
	goto _122
_122:
	(*ManifoldPoint)(unsafe.Pointer(mp)).Point = v9
	(*ManifoldPoint)(unsafe.Pointer(mp)).Separation = separation
	(*ManifoldPoint)(unsafe.Pointer(mp)).Id = uint16(0)
	(*(*Manifold)(unsafe.Add(nil, bp))).PointCount = int32(1)
	return *(*Manifold)(unsafe.Add(nil, bp))
}

func b2CollideChainSegmentAndCapsule(tls *_Stack, segmentA uintptr, xfA Transform, capsuleB uintptr, xfB Transform, cache uintptr) (r Manifold) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var _ /* polyB at bp+0 */ Polygon
	*(*Polygon)(unsafe.Add(nil, bp)) = b2MakeCapsule(tls, (*Capsule)(unsafe.Pointer(capsuleB)).Center1, (*Capsule)(unsafe.Pointer(capsuleB)).Center2, (*Capsule)(unsafe.Pointer(capsuleB)).Radius)
	return b2CollideChainSegmentAndPolygon(tls, segmentA, xfA, bp, xfB, cache)
}

func b2CollideChainSegmentAndPolygon(tls *_Stack, segmentA uintptr, xfA Transform, polygonB uintptr, xfB Transform, cache uintptr) (r Manifold) {
	bp := tls.Alloc(368)
	defer tls.Free(368)
	var C, xf, v1, v15, v17, v2 Transform
	var a11, a21, b11, b12, b21, b22, centroidB, edge0, edge1, edge2, n, n0, n1, n2, n21, normal, normal1, normalB, p1, p11, p2, pA, pAB, pAB1, pAB2, pB, v12, v13, v18, v19, v21, v22, v23, v25, v26, v28, v29, v7, v8, v9 Vec2
	var behind0, behind1, behind2 uint8
	var convexTol, d1, d2, dot1, dot11, dot12, dot2, dot21, dot22, edgeSeparation, invLength, length, polygonSeparation, radiusB, s, s0, s1, s2, s21, s3, x, y, v127, v40, v54, v65, v73, v81 float32
	var count, i, i1, i11, i2, i21, i3, i4, ia1, ia11, ia2, ia21, ib, ib1, ib11, ib2, ib21, incidentIndex, incidentNormal, referenceIndex, v133 int32
	var cp uintptr
	var normals [8]Vec2
	var output DistanceOutput
	var qr, v11, v3, v4, v5 Rot
	var smoothParams b2ChainSegmentParams
	var type1, type11, type2 b2NormalType
	var _ /* input at bp+176 */ DistanceInput
	var _ /* manifold at bp+0 */ Manifold
	var _ /* vertices at bp+112 */ [8]Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, a11, a21, b11, b12, b21, b22, behind0, behind1, behind2, centroidB, convexTol, count, cp, d1, d2, dot1, dot11, dot12, dot2, dot21, dot22, edge0, edge1, edge2, edgeSeparation, i, i1, i11, i2, i21, i3, i4, ia1, ia11, ia2, ia21, ib, ib1, ib11, ib2, ib21, incidentIndex, incidentNormal, invLength, length, n, n0, n1, n2, n21, normal, normal1, normalB, normals, output, p1, p11, p2, pA, pAB, pAB1, pAB2, pB, polygonSeparation, qr, radiusB, referenceIndex, s, s0, s1, s2, s21, s3, smoothParams, type1, type11, type2, x, xf, y, v1, v11, v12, v127, v13, v133, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v28, v29, v3, v4, v40, v5, v54, v65, v7, v73, v8, v81, v9
	*(*Manifold)(unsafe.Add(nil, bp)) = Manifold{}
	v1 = xfA
	v2 = xfB
	v3 = v1.Q
	v4 = v2.Q
	qr.S = float32(v3.C*v4.S) - float32(v3.S*v4.C)
	qr.C = float32(v3.C*v4.C) + float32(v3.S*v4.S)
	v5 = qr
	goto _6
_6:
	C.Q = v5
	v7 = v2.P
	v8 = v1.P
	v9 = Vec2{
		X: v7.X - v8.X,
		Y: v7.Y - v8.Y,
	}
	goto _10
_10:
	v11 = v1.Q
	v12 = v9
	v13 = Vec2{
		X: float32(v11.C*v12.X) + float32(v11.S*v12.Y),
		Y: float32(-v11.S*v12.X) + float32(v11.C*v12.Y),
	}
	goto _14
_14:
	C.P = v13
	v15 = C
	goto _16
_16:
	xf = v15
	v17 = xf
	v18 = (*Polygon)(unsafe.Pointer(polygonB)).Centroid
	x = float32(v17.Q.C*v18.X) - float32(v17.Q.S*v18.Y) + v17.P.X
	y = float32(v17.Q.S*v18.X) + float32(v17.Q.C*v18.Y) + v17.P.Y
	v19 = Vec2{
		X: x,
		Y: y,
	}
	goto _20
_20:
	centroidB = v19
	radiusB = (*Polygon)(unsafe.Pointer(polygonB)).Radius
	p11 = (*ChainSegment)(unsafe.Pointer(segmentA)).Segment.Point1
	p2 = (*ChainSegment)(unsafe.Pointer(segmentA)).Segment.Point2
	v21 = p2
	v22 = p11
	v23 = Vec2{
		X: v21.X - v22.X,
		Y: v21.Y - v22.Y,
	}
	goto _24
_24:
	v25 = v23
	length = sqrtf(tls, float32(v25.X*v25.X)+float32(v25.Y*v25.Y))
	if length < float32FromFloat32(1.1920928955078125e-07) {
		v26 = Vec2{}
		goto _27
	}
	invLength = float32FromFloat32(1) / length
	n = Vec2{
		X: float32(invLength * v25.X),
		Y: float32(invLength * v25.Y),
	}
	v26 = n
	goto _27
_27:
	edge1 = v26
	smoothParams = b2ChainSegmentParams{}
	smoothParams.Edge1 = edge1
	convexTol = float32FromFloat32(0.01)
	v7 = p11
	v8 = (*ChainSegment)(unsafe.Pointer(segmentA)).Ghost1
	v9 = Vec2{
		X: v7.X - v8.X,
		Y: v7.Y - v8.Y,
	}
	goto _31
_31:
	v12 = v9
	length = sqrtf(tls, float32(v12.X*v12.X)+float32(v12.Y*v12.Y))
	if length < float32FromFloat32(1.1920928955078125e-07) {
		v13 = Vec2{}
		goto _34
	}
	invLength = float32FromFloat32(1) / length
	n = Vec2{
		X: float32(invLength * v12.X),
		Y: float32(invLength * v12.Y),
	}
	v13 = n
	goto _34
_34:
	edge0 = v13
	v18 = edge0
	v19 = Vec2{
		X: v18.Y,
		Y: -v18.X,
	}
	goto _37
_37:
	smoothParams.Normal0 = v19
	v7 = edge0
	v8 = edge1
	v40 = float32(v7.X*v8.Y) - float32(v7.Y*v8.X)
	goto _41
_41:
	smoothParams.Convex1 = boolUint8(v40 >= convexTol)
	v7 = (*ChainSegment)(unsafe.Pointer(segmentA)).Ghost2
	v8 = p2
	v9 = Vec2{
		X: v7.X - v8.X,
		Y: v7.Y - v8.Y,
	}
	goto _45
_45:
	v12 = v9
	length = sqrtf(tls, float32(v12.X*v12.X)+float32(v12.Y*v12.Y))
	if length < float32FromFloat32(1.1920928955078125e-07) {
		v13 = Vec2{}
		goto _48
	}
	invLength = float32FromFloat32(1) / length
	n = Vec2{
		X: float32(invLength * v12.X),
		Y: float32(invLength * v12.Y),
	}
	v13 = n
	goto _48
_48:
	edge2 = v13
	v18 = edge2
	v19 = Vec2{
		X: v18.Y,
		Y: -v18.X,
	}
	goto _51
_51:
	smoothParams.Normal2 = v19
	v7 = edge1
	v8 = edge2
	v40 = float32(v7.X*v8.Y) - float32(v7.Y*v8.X)
	goto _55
_55:
	smoothParams.Convex2 = boolUint8(v40 >= convexTol)
	v7 = edge1
	v8 = Vec2{
		X: v7.Y,
		Y: -v7.X,
	}
	goto _58
_58:
	// Normal points to the right
	normal1 = v8
	v9 = centroidB
	v12 = p11
	v13 = Vec2{
		X: v9.X - v12.X,
		Y: v9.Y - v12.Y,
	}
	goto _62
_62:
	v18 = normal1
	v19 = v13
	v40 = float32(v18.X*v19.X) + float32(v18.Y*v19.Y)
	goto _66
_66:
	behind1 = boolUint8(v40 < float32FromFloat32(0))
	behind0 = boolUint8(true1 != 0)
	behind2 = boolUint8(true1 != 0)
	if smoothParams.Convex1 != 0 {
		v21 = centroidB
		v22 = p11
		v23 = Vec2{
			X: v21.X - v22.X,
			Y: v21.Y - v22.Y,
		}
		goto _70
	_70:
		v25 = smoothParams.Normal0
		v26 = v23
		v54 = float32(v25.X*v26.X) + float32(v25.Y*v26.Y)
		goto _74
	_74:
		behind0 = boolUint8(v54 < float32FromFloat32(0))
	}
	if smoothParams.Convex2 != 0 {
		v7 = centroidB
		v8 = p2
		v9 = Vec2{
			X: v7.X - v8.X,
			Y: v7.Y - v8.Y,
		}
		goto _78
	_78:
		v12 = smoothParams.Normal2
		v13 = v9
		v40 = float32(v12.X*v13.X) + float32(v12.Y*v13.Y)
		goto _82
	_82:
		behind2 = boolUint8(v40 < float32FromFloat32(0))
	}
	if behind1 != 0 && behind0 != 0 && behind2 != 0 {
		// one-sided collision
		return *(*Manifold)(unsafe.Add(nil, bp))
	}
	// Get polygonB in frameA
	count = (*Polygon)(unsafe.Pointer(polygonB)).Count
	i = 0
	for {
		if !(i < count) {
			break
		}
		v1 = xf
		v7 = *(*Vec2)(unsafe.Add(nil, polygonB+uintptr(i)*8))
		x = float32(v1.Q.C*v7.X) - float32(v1.Q.S*v7.Y) + v1.P.X
		y = float32(v1.Q.S*v7.X) + float32(v1.Q.C*v7.Y) + v1.P.Y
		v8 = Vec2{
			X: x,
			Y: y,
		}
		goto _87
	_87:
		(*(*[8]Vec2)(unsafe.Add(nil, bp+112)))[i] = v8
		v3 = xf.Q
		v7 = *(*Vec2)(unsafe.Add(nil, polygonB+64+uintptr(i)*8))
		v8 = Vec2{
			X: float32(v3.C*v7.X) - float32(v3.S*v7.Y),
			Y: float32(v3.S*v7.X) + float32(v3.C*v7.Y),
		}
		goto _91
	_91:
		normals[i] = v8
		goto _83
	_83:
		;
		i = i + 1
	}
	(*(*DistanceInput)(unsafe.Add(nil, bp+176))).ProxyA = b2MakeProxy(tls, segmentA+8, int32(2), float32FromFloat32(0))
	(*(*DistanceInput)(unsafe.Add(nil, bp+176))).ProxyB = b2MakeProxy(tls, bp+112, count, float32FromFloat32(0))
	(*(*DistanceInput)(unsafe.Add(nil, bp+176))).TransformA = b2Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp+176))).TransformB = b2Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp+176))).UseRadii = boolUint8(false1 != 0)
	output = b2ShapeDistance(tls, bp+176, cache, uintptrFromInt32(0), 0)
	if output.Distance > radiusB+float32(float32FromFloat32(4)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
		return *(*Manifold)(unsafe.Add(nil, bp))
	}
	if smoothParams.Convex1 != 0 {
		v7 = smoothParams.Normal0
	} else {
		v7 = normal1
	}
	// Snap concave normals for partial polygon
	n0 = v7
	if smoothParams.Convex2 != 0 {
		v8 = smoothParams.Normal2
	} else {
		v8 = normal1
	}
	n21 = v8
	// Index of incident vertex on polygon
	incidentIndex = -int32(1)
	incidentNormal = -int32(1)
	if int32FromUint8(behind1) == false1 && output.Distance > float32(float32FromFloat32(0.1)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
		// The closest features may be two vertices or an edge and a vertex even when there should
		// be face contact
		if int32FromUint16((*SimplexCache)(unsafe.Pointer(cache)).Count) == int32(1) {
			// vertex-vertex collision
			pA = output.PointA
			pB = output.PointB
			v9 = pB
			v12 = pA
			v13 = Vec2{
				X: v9.X - v12.X,
				Y: v9.Y - v12.Y,
			}
			goto _97
		_97:
			v18 = v13
			length = sqrtf(tls, float32(v18.X*v18.X)+float32(v18.Y*v18.Y))
			if length < float32FromFloat32(1.1920928955078125e-07) {
				v19 = Vec2{}
				goto _100
			}
			invLength = float32FromFloat32(1) / length
			n = Vec2{
				X: float32(invLength * v18.X),
				Y: float32(invLength * v18.Y),
			}
			v19 = n
			goto _100
		_100:
			normal = v19
			type1 = b2ClassifyNormal(tls, smoothParams, normal)
			if type1 == int32(b2_normalSkip) {
				return *(*Manifold)(unsafe.Add(nil, bp))
			}
			if type1 == int32(b2_normalAdmit) {
				v3 = xfA.Q
				v7 = normal
				v8 = Vec2{
					X: float32(v3.C*v7.X) - float32(v3.S*v7.Y),
					Y: float32(v3.S*v7.X) + float32(v3.C*v7.Y),
				}
				goto _104
			_104:
				(*(*Manifold)(unsafe.Add(nil, bp))).Normal = v8
				cp = bp + 12 + uintptr(0)*48
				v3 = xfA.Q
				v7 = pA
				v8 = Vec2{
					X: float32(v3.C*v7.X) - float32(v3.S*v7.Y),
					Y: float32(v3.S*v7.X) + float32(v3.C*v7.Y),
				}
				goto _108
			_108:
				(*ManifoldPoint)(unsafe.Pointer(cp)).AnchorA = v8
				v7 = xfA.P
				v8 = xfB.P
				v9 = Vec2{
					X: v7.X - v8.X,
					Y: v7.Y - v8.Y,
				}
				goto _112
			_112:
				v12 = (*ManifoldPoint)(unsafe.Pointer(cp)).AnchorA
				v13 = v9
				v18 = Vec2{
					X: v12.X + v13.X,
					Y: v12.Y + v13.Y,
				}
				goto _116
			_116:
				(*ManifoldPoint)(unsafe.Pointer(cp)).AnchorB = v18
				v7 = xfA.P
				v8 = (*ManifoldPoint)(unsafe.Pointer(cp)).AnchorA
				v9 = Vec2{
					X: v7.X + v8.X,
					Y: v7.Y + v8.Y,
				}
				goto _120
			_120:
				(*ManifoldPoint)(unsafe.Pointer(cp)).Point = v9
				(*ManifoldPoint)(unsafe.Pointer(cp)).Separation = output.Distance - radiusB
				(*ManifoldPoint)(unsafe.Pointer(cp)).Id = uint16FromInt32(int32FromUint8(*(*uint8_t)(unsafe.Add(nil, cache+2)))<<int32FromInt32(8) | int32FromUint8(*(*uint8_t)(unsafe.Add(nil, cache+5))))
				(*(*Manifold)(unsafe.Add(nil, bp))).PointCount = int32(1)
				return *(*Manifold)(unsafe.Add(nil, bp))
			}
			// fall through b2_normalSnap
			incidentIndex = int32FromUint8(*(*uint8_t)(unsafe.Add(nil, cache+5)))
		} else {
			// vertex-edge collision
			if !(int32FromUint16((*SimplexCache)(unsafe.Pointer(cache)).Count) == int32FromInt32(2)) && b2InternalAssertFcn(tls, __ccgo_ts+10021, __ccgo_ts+9936, int32FromInt32(1438)) != 0 {
				__builtin_trap(tls)
			}
			ia1 = int32FromUint8(*(*uint8_t)(unsafe.Add(nil, cache+2)))
			ia2 = int32FromUint8(*(*uint8_t)(unsafe.Add(nil, cache+2+1)))
			ib1 = int32FromUint8(*(*uint8_t)(unsafe.Add(nil, cache+5)))
			ib2 = int32FromUint8(*(*uint8_t)(unsafe.Add(nil, cache+5+1)))
			if ia1 == ia2 {
				// 1 point on A, expect 2 points on B
				if !(ib1 != ib2) && b2InternalAssertFcn(tls, __ccgo_ts+10039, __ccgo_ts+9936, int32FromInt32(1448)) != 0 {
					__builtin_trap(tls)
				}
				v7 = output.PointA
				v8 = output.PointB
				v9 = Vec2{
					X: v7.X - v8.X,
					Y: v7.Y - v8.Y,
				}
				goto _124
			_124:
				// Find polygon normal most aligned with vector between closest points.
				// This effectively sorts ib1 and ib2
				normalB = v9
				v12 = normalB
				v13 = normals[ib1]
				v40 = float32(v12.X*v13.X) + float32(v12.Y*v13.Y)
				goto _128
			_128:
				dot1 = v40
				v18 = normalB
				v19 = normals[ib2]
				v54 = float32(v18.X*v19.X) + float32(v18.Y*v19.Y)
				goto _132
			_132:
				dot2 = v54
				if dot1 > dot2 {
					v133 = ib1
				} else {
					v133 = ib2
				}
				ib = v133
				// Use accurate normal
				normalB = normals[ib]
				v7 = normalB
				v8 = Vec2{
					X: -v7.X,
					Y: -v7.Y,
				}
				goto _136
			_136:
				type11 = b2ClassifyNormal(tls, smoothParams, v8)
				if type11 == int32(b2_normalSkip) {
					return *(*Manifold)(unsafe.Add(nil, bp))
				}
				if type11 == int32(b2_normalAdmit) {
					// Get polygon edge associated with normal
					ib1 = ib
					if ib < count-int32(1) {
						v133 = ib + int32(1)
					} else {
						v133 = 0
					}
					ib2 = v133
					b11 = (*(*[8]Vec2)(unsafe.Add(nil, bp+112)))[ib1]
					b21 = (*(*[8]Vec2)(unsafe.Add(nil, bp+112)))[ib2]
					// Find incident segment vertex
					v7 = p11
					v8 = b11
					v9 = Vec2{
						X: v7.X - v8.X,
						Y: v7.Y - v8.Y,
					}
					goto _141
				_141:
					v12 = normalB
					v13 = v9
					v40 = float32(v12.X*v13.X) + float32(v12.Y*v13.Y)
					goto _145
				_145:
					dot1 = v40
					v7 = p2
					v8 = b11
					v9 = Vec2{
						X: v7.X - v8.X,
						Y: v7.Y - v8.Y,
					}
					goto _149
				_149:
					v12 = normalB
					v13 = v9
					v40 = float32(v12.X*v13.X) + float32(v12.Y*v13.Y)
					goto _153
				_153:
					dot2 = v40
					if dot1 < dot2 {
						v7 = n0
						v8 = normalB
						v40 = float32(v7.X*v8.X) + float32(v7.Y*v8.Y)
						goto _157
					_157:
						v9 = normal1
						v12 = normalB
						v54 = float32(v9.X*v12.X) + float32(v9.Y*v12.Y)
						goto _161
					_161:
						if v40 < v54 {
							// Neighbor is incident
							return *(*Manifold)(unsafe.Add(nil, bp))
						}
					} else {
						v7 = n21
						v8 = normalB
						v40 = float32(v7.X*v8.X) + float32(v7.Y*v8.Y)
						goto _165
					_165:
						v9 = normal1
						v12 = normalB
						v54 = float32(v9.X*v12.X) + float32(v9.Y*v12.Y)
						goto _169
					_169:
						if v40 < v54 {
							// Neighbor is incident
							return *(*Manifold)(unsafe.Add(nil, bp))
						}
					}
					*(*Manifold)(unsafe.Add(nil, bp)) = b2ClipSegments(tls, b11, b21, p11, p2, normalB, radiusB, float32FromFloat32(0), uint16FromInt32(int32FromUint8(uint8FromInt32(ib1))<<int32FromInt32(8)|int32FromUint8(uint8FromInt32(int32FromInt32(1)))), uint16FromInt32(int32FromUint8(uint8FromInt32(ib2))<<int32FromInt32(8)|int32FromUint8(uint8FromInt32(int32FromInt32(0)))))
					if !((*(*Manifold)(unsafe.Add(nil, bp))).PointCount == 0 || (*(*Manifold)(unsafe.Add(nil, bp))).PointCount == int32(2)) && b2InternalAssertFcn(tls, __ccgo_ts+10050, __ccgo_ts+9936, int32FromInt32(1499)) != 0 {
						__builtin_trap(tls)
					}
					if (*(*Manifold)(unsafe.Add(nil, bp))).PointCount == int32(2) {
						v7 = normalB
						v8 = Vec2{
							X: -v7.X,
							Y: -v7.Y,
						}
						goto _172
					_172:
						v3 = xfA.Q
						v9 = v8
						v12 = Vec2{
							X: float32(v3.C*v9.X) - float32(v3.S*v9.Y),
							Y: float32(v3.S*v9.X) + float32(v3.C*v9.Y),
						}
						goto _176
					_176:
						(*(*Manifold)(unsafe.Add(nil, bp))).Normal = v12
						v3 = xfA.Q
						v7 = (*(*ManifoldPoint)(unsafe.Add(nil, bp+12))).AnchorA
						v8 = Vec2{
							X: float32(v3.C*v7.X) - float32(v3.S*v7.Y),
							Y: float32(v3.S*v7.X) + float32(v3.C*v7.Y),
						}
						goto _180
					_180:
						(*(*ManifoldPoint)(unsafe.Add(nil, bp+12))).AnchorA = v8
						v3 = xfA.Q
						v7 = (*(*ManifoldPoint)(unsafe.Add(nil, bp+12+1*48))).AnchorA
						v8 = Vec2{
							X: float32(v3.C*v7.X) - float32(v3.S*v7.Y),
							Y: float32(v3.S*v7.X) + float32(v3.C*v7.Y),
						}
						goto _184
					_184:
						(*(*ManifoldPoint)(unsafe.Add(nil, bp+12+1*48))).AnchorA = v8
						v7 = xfA.P
						v8 = xfB.P
						v9 = Vec2{
							X: v7.X - v8.X,
							Y: v7.Y - v8.Y,
						}
						goto _188
					_188:
						pAB = v9
						v12 = (*(*ManifoldPoint)(unsafe.Add(nil, bp+12))).AnchorA
						v13 = pAB
						v18 = Vec2{
							X: v12.X + v13.X,
							Y: v12.Y + v13.Y,
						}
						goto _192
					_192:
						(*(*ManifoldPoint)(unsafe.Add(nil, bp+12))).AnchorB = v18
						v7 = (*(*ManifoldPoint)(unsafe.Add(nil, bp+12+1*48))).AnchorA
						v8 = pAB
						v9 = Vec2{
							X: v7.X + v8.X,
							Y: v7.Y + v8.Y,
						}
						goto _196
					_196:
						(*(*ManifoldPoint)(unsafe.Add(nil, bp+12+1*48))).AnchorB = v9
						v7 = xfA.P
						v8 = (*(*ManifoldPoint)(unsafe.Add(nil, bp+12))).AnchorA
						v9 = Vec2{
							X: v7.X + v8.X,
							Y: v7.Y + v8.Y,
						}
						goto _200
					_200:
						(*(*ManifoldPoint)(unsafe.Add(nil, bp+12))).Point = v9
						v7 = xfA.P
						v8 = (*(*ManifoldPoint)(unsafe.Add(nil, bp+12+1*48))).AnchorA
						v9 = Vec2{
							X: v7.X + v8.X,
							Y: v7.Y + v8.Y,
						}
						goto _204
					_204:
						(*(*ManifoldPoint)(unsafe.Add(nil, bp+12+1*48))).Point = v9
					}
					return *(*Manifold)(unsafe.Add(nil, bp))
				}
				// fall through b2_normalSnap
				incidentNormal = ib
			} else {
				v7 = (*(*[8]Vec2)(unsafe.Add(nil, bp+112)))[ib1]
				v8 = p11
				v9 = Vec2{
					X: v7.X - v8.X,
					Y: v7.Y - v8.Y,
				}
				goto _208
			_208:
				v12 = normal1
				v13 = v9
				v40 = float32(v12.X*v13.X) + float32(v12.Y*v13.Y)
				goto _212
			_212:
				// Get index of incident polygonB vertex
				dot11 = v40
				v18 = (*(*[8]Vec2)(unsafe.Add(nil, bp+112)))[ib2]
				v19 = p2
				v21 = Vec2{
					X: v18.X - v19.X,
					Y: v18.Y - v19.Y,
				}
				goto _216
			_216:
				v22 = normal1
				v23 = v21
				v54 = float32(v22.X*v23.X) + float32(v22.Y*v23.Y)
				goto _220
			_220:
				dot21 = v54
				if dot11 < dot21 {
					v133 = ib1
				} else {
					v133 = ib2
				}
				incidentIndex = v133
			}
		}
	} else {
		// SAT edge normal
		edgeSeparation = float32FromFloat32(3.4028234663852886e+38)
		i1 = 0
		for {
			if !(i1 < count) {
				break
			}
			v7 = (*(*[8]Vec2)(unsafe.Add(nil, bp+112)))[i1]
			v8 = p11
			v9 = Vec2{
				X: v7.X - v8.X,
				Y: v7.Y - v8.Y,
			}
			goto _226
		_226:
			v12 = normal1
			v13 = v9
			v40 = float32(v12.X*v13.X) + float32(v12.Y*v13.Y)
			goto _230
		_230:
			s = v40
			if s < edgeSeparation {
				edgeSeparation = s
				incidentIndex = i1
			}
			goto _222
		_222:
			;
			i1 = i1 + 1
		}
		// Check convex neighbor for edge separation
		if smoothParams.Convex1 != 0 {
			s0 = float32FromFloat32(3.4028234663852886e+38)
			i2 = 0
			for {
				if !(i2 < count) {
					break
				}
				v7 = (*(*[8]Vec2)(unsafe.Add(nil, bp+112)))[i2]
				v8 = p11
				v9 = Vec2{
					X: v7.X - v8.X,
					Y: v7.Y - v8.Y,
				}
				goto _235
			_235:
				v12 = smoothParams.Normal0
				v13 = v9
				v40 = float32(v12.X*v13.X) + float32(v12.Y*v13.Y)
				goto _239
			_239:
				s1 = v40
				if s1 < s0 {
					s0 = s1
				}
				goto _231
			_231:
				;
				i2 = i2 + 1
			}
			if s0 > edgeSeparation {
				edgeSeparation = s0
				// Indicate neighbor owns edge separation
				incidentIndex = -int32(1)
			}
		}
		// Check convex neighbor for edge separation
		if smoothParams.Convex2 != 0 {
			s21 = float32FromFloat32(3.4028234663852886e+38)
			i3 = 0
			for {
				if !(i3 < count) {
					break
				}
				v7 = (*(*[8]Vec2)(unsafe.Add(nil, bp+112)))[i3]
				v8 = p2
				v9 = Vec2{
					X: v7.X - v8.X,
					Y: v7.Y - v8.Y,
				}
				goto _244
			_244:
				v12 = smoothParams.Normal2
				v13 = v9
				v40 = float32(v12.X*v13.X) + float32(v12.Y*v13.Y)
				goto _248
			_248:
				s2 = v40
				if s2 < s21 {
					s21 = s2
				}
				goto _240
			_240:
				;
				i3 = i3 + 1
			}
			if s21 > edgeSeparation {
				edgeSeparation = s21
				// Indicate neighbor owns edge separation
				incidentIndex = -int32(1)
			}
		}
		// SAT polygon normals
		polygonSeparation = -float32FromFloat32(3.4028234663852886e+38)
		referenceIndex = -int32(1)
		i4 = 0
		for {
			if !(i4 < count) {
				break
			}
			n1 = normals[i4]
			v7 = n1
			v8 = Vec2{
				X: -v7.X,
				Y: -v7.Y,
			}
			goto _252
		_252:
			type2 = b2ClassifyNormal(tls, smoothParams, v8)
			if type2 != int32(b2_normalAdmit) {
				goto _249
			}
			// Check the infinite sides of the partial polygon
			// if ((smoothParams.convex1 && b2Cross(n0, n) > 0.0f) || (smoothParams.convex2 && b2Cross(n, n2) > 0.0f))
			//{
			//	continue;
			//}
			p1 = (*(*[8]Vec2)(unsafe.Add(nil, bp+112)))[i4]
			v7 = p2
			v8 = p1
			v9 = Vec2{
				X: v7.X - v8.X,
				Y: v7.Y - v8.Y,
			}
			goto _256
		_256:
			v12 = n1
			v13 = v9
			v40 = float32(v12.X*v13.X) + float32(v12.Y*v13.Y)
			goto _260
		_260:
			v18 = p11
			v19 = p1
			v21 = Vec2{
				X: v18.X - v19.X,
				Y: v18.Y - v19.Y,
			}
			goto _264
		_264:
			v22 = n1
			v23 = v21
			v54 = float32(v22.X*v23.X) + float32(v22.Y*v23.Y)
			goto _268
		_268:
			v65 = v40
			v73 = v54
			if v65 < v73 {
				v127 = v65
			} else {
				v127 = v73
			}
			v81 = v127
			goto _272
		_272:
			s3 = v81
			if s3 > polygonSeparation {
				polygonSeparation = s3
				referenceIndex = i4
			}
			goto _249
		_249:
			;
			i4 = i4 + 1
		}
		if polygonSeparation > edgeSeparation {
			ia11 = referenceIndex
			if ia11 < count-int32(1) {
				v133 = ia11 + int32(1)
			} else {
				v133 = 0
			}
			ia21 = v133
			a11 = (*(*[8]Vec2)(unsafe.Add(nil, bp+112)))[ia11]
			a21 = (*(*[8]Vec2)(unsafe.Add(nil, bp+112)))[ia21]
			n2 = normals[ia11]
			v7 = p11
			v8 = a11
			v9 = Vec2{
				X: v7.X - v8.X,
				Y: v7.Y - v8.Y,
			}
			goto _278
		_278:
			v12 = n2
			v13 = v9
			v40 = float32(v12.X*v13.X) + float32(v12.Y*v13.Y)
			goto _282
		_282:
			dot12 = v40
			v18 = p2
			v19 = a11
			v21 = Vec2{
				X: v18.X - v19.X,
				Y: v18.Y - v19.Y,
			}
			goto _286
		_286:
			v22 = n2
			v23 = v21
			v54 = float32(v22.X*v23.X) + float32(v22.Y*v23.Y)
			goto _290
		_290:
			dot22 = v54
			if dot12 < dot22 {
				v25 = n0
				v26 = n2
				v65 = float32(v25.X*v26.X) + float32(v25.Y*v26.Y)
				goto _294
			_294:
				v28 = normal1
				v29 = n2
				v73 = float32(v28.X*v29.X) + float32(v28.Y*v29.Y)
				goto _298
			_298:
				if v65 < v73 {
					// Neighbor is incident
					return *(*Manifold)(unsafe.Add(nil, bp))
				}
			} else {
				v7 = n21
				v8 = n2
				v40 = float32(v7.X*v8.X) + float32(v7.Y*v8.Y)
				goto _302
			_302:
				v9 = normal1
				v12 = n2
				v54 = float32(v9.X*v12.X) + float32(v9.Y*v12.Y)
				goto _306
			_306:
				if v40 < v54 {
					// Neighbor is incident
					return *(*Manifold)(unsafe.Add(nil, bp))
				}
			}
			*(*Manifold)(unsafe.Add(nil, bp)) = b2ClipSegments(tls, a11, a21, p11, p2, normals[ia11], radiusB, float32FromFloat32(0), uint16FromInt32(int32FromUint8(uint8FromInt32(ia11))<<int32FromInt32(8)|int32FromUint8(uint8FromInt32(int32FromInt32(1)))), uint16FromInt32(int32FromUint8(uint8FromInt32(ia21))<<int32FromInt32(8)|int32FromUint8(uint8FromInt32(int32FromInt32(0)))))
			if !((*(*Manifold)(unsafe.Add(nil, bp))).PointCount == 0 || (*(*Manifold)(unsafe.Add(nil, bp))).PointCount == int32(2)) && b2InternalAssertFcn(tls, __ccgo_ts+10050, __ccgo_ts+9936, int32FromInt32(1648)) != 0 {
				__builtin_trap(tls)
			}
			if (*(*Manifold)(unsafe.Add(nil, bp))).PointCount == int32(2) {
				v7 = normals[ia11]
				v8 = Vec2{
					X: -v7.X,
					Y: -v7.Y,
				}
				goto _309
			_309:
				v3 = xfA.Q
				v9 = v8
				v12 = Vec2{
					X: float32(v3.C*v9.X) - float32(v3.S*v9.Y),
					Y: float32(v3.S*v9.X) + float32(v3.C*v9.Y),
				}
				goto _313
			_313:
				(*(*Manifold)(unsafe.Add(nil, bp))).Normal = v12
				v3 = xfA.Q
				v7 = (*(*ManifoldPoint)(unsafe.Add(nil, bp+12))).AnchorA
				v8 = Vec2{
					X: float32(v3.C*v7.X) - float32(v3.S*v7.Y),
					Y: float32(v3.S*v7.X) + float32(v3.C*v7.Y),
				}
				goto _317
			_317:
				(*(*ManifoldPoint)(unsafe.Add(nil, bp+12))).AnchorA = v8
				v3 = xfA.Q
				v7 = (*(*ManifoldPoint)(unsafe.Add(nil, bp+12+1*48))).AnchorA
				v8 = Vec2{
					X: float32(v3.C*v7.X) - float32(v3.S*v7.Y),
					Y: float32(v3.S*v7.X) + float32(v3.C*v7.Y),
				}
				goto _321
			_321:
				(*(*ManifoldPoint)(unsafe.Add(nil, bp+12+1*48))).AnchorA = v8
				v7 = xfA.P
				v8 = xfB.P
				v9 = Vec2{
					X: v7.X - v8.X,
					Y: v7.Y - v8.Y,
				}
				goto _325
			_325:
				pAB1 = v9
				v12 = (*(*ManifoldPoint)(unsafe.Add(nil, bp+12))).AnchorA
				v13 = pAB1
				v18 = Vec2{
					X: v12.X + v13.X,
					Y: v12.Y + v13.Y,
				}
				goto _329
			_329:
				(*(*ManifoldPoint)(unsafe.Add(nil, bp+12))).AnchorB = v18
				v7 = (*(*ManifoldPoint)(unsafe.Add(nil, bp+12+1*48))).AnchorA
				v8 = pAB1
				v9 = Vec2{
					X: v7.X + v8.X,
					Y: v7.Y + v8.Y,
				}
				goto _333
			_333:
				(*(*ManifoldPoint)(unsafe.Add(nil, bp+12+1*48))).AnchorB = v9
				v7 = xfA.P
				v8 = (*(*ManifoldPoint)(unsafe.Add(nil, bp+12))).AnchorA
				v9 = Vec2{
					X: v7.X + v8.X,
					Y: v7.Y + v8.Y,
				}
				goto _337
			_337:
				(*(*ManifoldPoint)(unsafe.Add(nil, bp+12))).Point = v9
				v7 = xfA.P
				v8 = (*(*ManifoldPoint)(unsafe.Add(nil, bp+12+1*48))).AnchorA
				v9 = Vec2{
					X: v7.X + v8.X,
					Y: v7.Y + v8.Y,
				}
				goto _341
			_341:
				(*(*ManifoldPoint)(unsafe.Add(nil, bp+12+1*48))).Point = v9
			}
			return *(*Manifold)(unsafe.Add(nil, bp))
		}
		if incidentIndex == -int32(1) {
			// neighboring segment is the separating axis
			return *(*Manifold)(unsafe.Add(nil, bp))
		}
		// fall through segment normal axis
	}
	if !(incidentNormal != -int32(1) || incidentIndex != -int32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+10103, __ccgo_ts+9936, int32FromInt32(1674)) != 0 {
		__builtin_trap(tls)
	}
	if incidentNormal != -int32(1) {
		ib11 = incidentNormal
		if ib11 < count-int32(1) {
			v133 = ib11 + int32(1)
		} else {
			v133 = 0
		}
		ib21 = v133
		b12 = (*(*[8]Vec2)(unsafe.Add(nil, bp+112)))[ib11]
		b22 = (*(*[8]Vec2)(unsafe.Add(nil, bp+112)))[ib21]
	} else {
		i21 = incidentIndex
		if i21 > 0 {
			v133 = i21 - int32(1)
		} else {
			v133 = count - int32(1)
		}
		i11 = v133
		v7 = normal1
		v8 = normals[i11]
		v40 = float32(v7.X*v8.X) + float32(v7.Y*v8.Y)
		goto _347
	_347:
		d1 = v40
		v9 = normal1
		v12 = normals[i21]
		v54 = float32(v9.X*v12.X) + float32(v9.Y*v12.Y)
		goto _351
	_351:
		d2 = v54
		if d1 < d2 {
			ib11 = i11
			ib21 = i21
			b12 = (*(*[8]Vec2)(unsafe.Add(nil, bp+112)))[ib11]
			b22 = (*(*[8]Vec2)(unsafe.Add(nil, bp+112)))[ib21]
		} else {
			ib11 = i21
			if i21 < count-int32(1) {
				v133 = i21 + int32(1)
			} else {
				v133 = 0
			}
			ib21 = v133
			b12 = (*(*[8]Vec2)(unsafe.Add(nil, bp+112)))[ib11]
			b22 = (*(*[8]Vec2)(unsafe.Add(nil, bp+112)))[ib21]
		}
	}
	*(*Manifold)(unsafe.Add(nil, bp)) = b2ClipSegments(tls, p11, p2, b12, b22, normal1, float32FromFloat32(0), radiusB, uint16FromInt32(int32FromUint8(uint8FromInt32(int32FromInt32(0)))<<int32FromInt32(8)|int32FromUint8(uint8FromInt32(ib21))), uint16FromInt32(int32FromUint8(uint8FromInt32(int32FromInt32(1)))<<int32FromInt32(8)|int32FromUint8(uint8FromInt32(ib11))))
	if !((*(*Manifold)(unsafe.Add(nil, bp))).PointCount == 0 || (*(*Manifold)(unsafe.Add(nil, bp))).PointCount == int32(2)) && b2InternalAssertFcn(tls, __ccgo_ts+10050, __ccgo_ts+9936, int32FromInt32(1711)) != 0 {
		__builtin_trap(tls)
	}
	if (*(*Manifold)(unsafe.Add(nil, bp))).PointCount == int32(2) {
		// There may be no points c
		v3 = xfA.Q
		v7 = (*(*Manifold)(unsafe.Add(nil, bp))).Normal
		v8 = Vec2{
			X: float32(v3.C*v7.X) - float32(v3.S*v7.Y),
			Y: float32(v3.S*v7.X) + float32(v3.C*v7.Y),
		}
		goto _356
	_356:
		(*(*Manifold)(unsafe.Add(nil, bp))).Normal = v8
		v3 = xfA.Q
		v7 = (*(*ManifoldPoint)(unsafe.Add(nil, bp+12))).AnchorA
		v8 = Vec2{
			X: float32(v3.C*v7.X) - float32(v3.S*v7.Y),
			Y: float32(v3.S*v7.X) + float32(v3.C*v7.Y),
		}
		goto _360
	_360:
		(*(*ManifoldPoint)(unsafe.Add(nil, bp+12))).AnchorA = v8
		v3 = xfA.Q
		v7 = (*(*ManifoldPoint)(unsafe.Add(nil, bp+12+1*48))).AnchorA
		v8 = Vec2{
			X: float32(v3.C*v7.X) - float32(v3.S*v7.Y),
			Y: float32(v3.S*v7.X) + float32(v3.C*v7.Y),
		}
		goto _364
	_364:
		(*(*ManifoldPoint)(unsafe.Add(nil, bp+12+1*48))).AnchorA = v8
		v7 = xfA.P
		v8 = xfB.P
		v9 = Vec2{
			X: v7.X - v8.X,
			Y: v7.Y - v8.Y,
		}
		goto _368
	_368:
		pAB2 = v9
		v12 = (*(*ManifoldPoint)(unsafe.Add(nil, bp+12))).AnchorA
		v13 = pAB2
		v18 = Vec2{
			X: v12.X + v13.X,
			Y: v12.Y + v13.Y,
		}
		goto _372
	_372:
		(*(*ManifoldPoint)(unsafe.Add(nil, bp+12))).AnchorB = v18
		v7 = (*(*ManifoldPoint)(unsafe.Add(nil, bp+12+1*48))).AnchorA
		v8 = pAB2
		v9 = Vec2{
			X: v7.X + v8.X,
			Y: v7.Y + v8.Y,
		}
		goto _376
	_376:
		(*(*ManifoldPoint)(unsafe.Add(nil, bp+12+1*48))).AnchorB = v9
		v7 = xfA.P
		v8 = (*(*ManifoldPoint)(unsafe.Add(nil, bp+12))).AnchorA
		v9 = Vec2{
			X: v7.X + v8.X,
			Y: v7.Y + v8.Y,
		}
		goto _380
	_380:
		(*(*ManifoldPoint)(unsafe.Add(nil, bp+12))).Point = v9
		v7 = xfA.P
		v8 = (*(*ManifoldPoint)(unsafe.Add(nil, bp+12+1*48))).AnchorA
		v9 = Vec2{
			X: v7.X + v8.X,
			Y: v7.Y + v8.Y,
		}
		goto _384
	_384:
		(*(*ManifoldPoint)(unsafe.Add(nil, bp+12+1*48))).Point = v9
	}
	return *(*Manifold)(unsafe.Add(nil, bp))
}

func b2CollideMover(tls *_Stack, shape uintptr, transform Transform, mover uintptr) (r PlaneResult) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var result PlaneResult
	var vx, vy float32
	var v1 Transform
	var v2, v3 Vec2
	var v9 Rot
	var _ /* localMover at bp+0 */ Capsule
	_, _, _, _, _, _, _ = result, vx, vy, v1, v2, v3, v9
	v1 = transform
	v2 = (*Capsule)(unsafe.Pointer(mover)).Center1
	vx = v2.X - v1.P.X
	vy = v2.Y - v1.P.Y
	v3 = Vec2{
		X: float32(v1.Q.C*vx) + float32(v1.Q.S*vy),
		Y: float32(-v1.Q.S*vx) + float32(v1.Q.C*vy),
	}
	goto _4
_4:
	(*(*Capsule)(unsafe.Add(nil, bp))).Center1 = v3
	v1 = transform
	v2 = (*Capsule)(unsafe.Pointer(mover)).Center2
	vx = v2.X - v1.P.X
	vy = v2.Y - v1.P.Y
	v3 = Vec2{
		X: float32(v1.Q.C*vx) + float32(v1.Q.S*vy),
		Y: float32(-v1.Q.S*vx) + float32(v1.Q.C*vy),
	}
	goto _8
_8:
	(*(*Capsule)(unsafe.Add(nil, bp))).Center2 = v3
	(*(*Capsule)(unsafe.Add(nil, bp))).Radius = (*Capsule)(unsafe.Pointer(mover)).Radius
	result = PlaneResult{}
	switch (*b2Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b2_capsuleShape):
		result = b2CollideMoverAndCapsule(tls, shape+132, bp)
	case int32(b2_circleShape):
		result = b2CollideMoverAndCircle(tls, shape+132, bp)
	case int32(b2_polygonShape):
		result = b2CollideMoverAndPolygon(tls, shape+132, bp)
	case int32(b2_segmentShape):
		result = b2CollideMoverAndSegment(tls, shape+132, bp)
	case int32(b2_chainSegmentShape):
		result = b2CollideMoverAndSegment(tls, shape+132+8, bp)
	default:
		return result
	}
	if int32FromUint8(result.Hit) == false1 {
		return result
	}
	v9 = transform.Q
	v2 = result.Plane.Normal
	v3 = Vec2{
		X: float32(v9.C*v2.X) - float32(v9.S*v2.Y),
		Y: float32(v9.S*v2.X) + float32(v9.C*v2.Y),
	}
	goto _12
_12:
	result.Plane.Normal = v3
	return result
}

func b2CollideTask(tls *_Stack, startIndex int32, endIndex int32, threadIndex uint32, context uintptr) {
	var blockIndex, v7 uint32_t
	var bodies, bodyA, bodyB, bodySimA, bodySimB, contactSim, contactSims, shapeA, shapeB, shapes, stepContext, taskContext, world, v6 uintptr
	var centerOffsetA, centerOffsetB, v11, v12, v15, v16 Vec2
	var contactId, contactIndex, v8 int32
	var overlap, touching, wasTouching, v4 uint8
	var transformA, transformB Transform
	var v10, v14 Rot
	var v2, v3 AABB
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = blockIndex, bodies, bodyA, bodyB, bodySimA, bodySimB, centerOffsetA, centerOffsetB, contactId, contactIndex, contactSim, contactSims, overlap, shapeA, shapeB, shapes, stepContext, taskContext, touching, transformA, transformB, wasTouching, world, v10, v11, v12, v14, v15, v16, v2, v3, v4, v6, v7, v8
	stepContext = context
	world = (*b2StepContext)(unsafe.Pointer(stepContext)).World
	if !(int32FromUint32(threadIndex) < (*b2World)(unsafe.Pointer(world)).WorkerCount) && b2InternalAssertFcn(tls, __ccgo_ts+10751, __ccgo_ts+15342, int32FromInt32(363)) != 0 {
		__builtin_trap(tls)
	}
	taskContext = (*b2World)(unsafe.Pointer(world)).TaskContexts.Data + uintptr(threadIndex)*56
	contactSims = (*b2StepContext)(unsafe.Pointer(stepContext)).Contacts
	shapes = (*b2World)(unsafe.Pointer(world)).Shapes.Data
	bodies = (*b2World)(unsafe.Pointer(world)).Bodies.Data
	if !(startIndex < endIndex) && b2InternalAssertFcn(tls, __ccgo_ts+10813, __ccgo_ts+15342, int32FromInt32(369)) != 0 {
		__builtin_trap(tls)
	}
	contactIndex = startIndex
	for {
		if !(contactIndex < endIndex) {
			break
		}
		contactSim = *(*uintptr)(unsafe.Add(nil, contactSims+uintptr(contactIndex)*8))
		contactId = (*b2ContactSim)(unsafe.Pointer(contactSim)).ContactId
		shapeA = shapes + uintptr((*b2ContactSim)(unsafe.Pointer(contactSim)).ShapeIdA)*288
		shapeB = shapes + uintptr((*b2ContactSim)(unsafe.Pointer(contactSim)).ShapeIdB)*288
		v2 = (*b2Shape)(unsafe.Pointer(shapeA)).FatAABB
		v3 = (*b2Shape)(unsafe.Pointer(shapeB)).FatAABB
		v4 = boolUint8(!(v3.LowerBound.X > v2.UpperBound.X || v3.LowerBound.Y > v2.UpperBound.Y || v2.LowerBound.X > v3.UpperBound.X || v2.LowerBound.Y > v3.UpperBound.Y))
		goto _5
	_5:
		// Do proxies still overlap?
		overlap = v4
		if int32FromUint8(overlap) == false1 {
			*(*uint32_t)(unsafe.Add(nil, contactSim+164)) |= uint32(b2_simDisjoint)
			*(*uint32_t)(unsafe.Add(nil, contactSim+164)) &= uint32FromInt32(^int32(b2_simTouchingFlag))
			v6 = taskContext
			v7 = uint32FromInt32(contactId)
			blockIndex = v7 / uint32(64)
			if !(blockIndex < (*b2BitSet)(unsafe.Pointer(v6)).BlockCount) && b2InternalAssertFcn(tls, __ccgo_ts+10859, __ccgo_ts+10891, int32FromInt32(28)) != 0 {
				__builtin_trap(tls)
			}
			*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v6)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v7 % uint32FromInt32(64))
		} else {
			wasTouching = uint8(boolUint32((*b2ContactSim)(unsafe.Pointer(contactSim)).SimFlags&uint32(b2_simTouchingFlag) != 0))
			// Update contact respecting shape/body order (A,B)
			bodyA = bodies + uintptr((*b2Shape)(unsafe.Pointer(shapeA)).BodyId)*128
			bodyB = bodies + uintptr((*b2Shape)(unsafe.Pointer(shapeB)).BodyId)*128
			bodySimA = b2GetBodySim(tls, world, bodyA)
			bodySimB = b2GetBodySim(tls, world, bodyB)
			// avoid cache misses in b2PrepareContactsTask
			if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) {
				v8 = (*b2Body)(unsafe.Pointer(bodyA)).LocalIndex
			} else {
				v8 = -int32(1)
			}
			(*b2ContactSim)(unsafe.Pointer(contactSim)).BodySimIndexA = v8
			(*b2ContactSim)(unsafe.Pointer(contactSim)).InvMassA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvMass
			(*b2ContactSim)(unsafe.Pointer(contactSim)).InvIA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvInertia
			if (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet) {
				v8 = (*b2Body)(unsafe.Pointer(bodyB)).LocalIndex
			} else {
				v8 = -int32(1)
			}
			(*b2ContactSim)(unsafe.Pointer(contactSim)).BodySimIndexB = v8
			(*b2ContactSim)(unsafe.Pointer(contactSim)).InvMassB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvMass
			(*b2ContactSim)(unsafe.Pointer(contactSim)).InvIB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvInertia
			transformA = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform
			transformB = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform
			v10 = transformA.Q
			v11 = (*b2BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
			v12 = Vec2{
				X: float32(v10.C*v11.X) - float32(v10.S*v11.Y),
				Y: float32(v10.S*v11.X) + float32(v10.C*v11.Y),
			}
			goto _13
		_13:
			centerOffsetA = v12
			v14 = transformB.Q
			v15 = (*b2BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
			v16 = Vec2{
				X: float32(v14.C*v15.X) - float32(v14.S*v15.Y),
				Y: float32(v14.S*v15.X) + float32(v14.C*v15.Y),
			}
			goto _17
		_17:
			centerOffsetB = v16
			// This updates solid contacts
			touching = b2UpdateContact(tls, world, contactSim, shapeA, transformA, centerOffsetA, shapeB, transformB, centerOffsetB)
			// State changes that affect island connectivity. Also affects contact events.
			if int32FromUint8(touching) == int32(true1) && int32FromUint8(wasTouching) == false1 {
				*(*uint32_t)(unsafe.Add(nil, contactSim+164)) |= uint32(b2_simStartedTouching)
				v6 = taskContext
				v7 = uint32FromInt32(contactId)
				blockIndex = v7 / uint32(64)
				if !(blockIndex < (*b2BitSet)(unsafe.Pointer(v6)).BlockCount) && b2InternalAssertFcn(tls, __ccgo_ts+10859, __ccgo_ts+10891, int32FromInt32(28)) != 0 {
					__builtin_trap(tls)
				}
				*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v6)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v7 % uint32FromInt32(64))
			} else {
				if int32FromUint8(touching) == false1 && int32FromUint8(wasTouching) == int32(true1) {
					*(*uint32_t)(unsafe.Add(nil, contactSim+164)) |= uint32(b2_simStoppedTouching)
					v6 = taskContext
					v7 = uint32FromInt32(contactId)
					blockIndex = v7 / uint32(64)
					if !(blockIndex < (*b2BitSet)(unsafe.Pointer(v6)).BlockCount) && b2InternalAssertFcn(tls, __ccgo_ts+10859, __ccgo_ts+10891, int32FromInt32(28)) != 0 {
						__builtin_trap(tls)
					}
					*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v6)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v7 % uint32FromInt32(64))
				}
			}
			// To make this work, the time of impact code needs to adjust the target
			// distance based on the number of TOI events for a body.
			// if (touching && bodySimB->isFast)
			//{
			//	b2Manifold* manifold = &contactSim->manifold;
			//	int pointCount = manifold->pointCount;
			//	for (int i = 0; i < pointCount; ++i)
			//	{
			//		// trick the solver into pushing the fast shapes apart
			//		manifold->points[i].separation -= 0.25f * B2_SPECULATIVE_DISTANCE;
			//	}
			//}
		}
		goto _1
	_1:
		;
		contactIndex = contactIndex + 1
	}
}

// C documentation
//
//	// Narrow-phase collision
func b2Collide(tls *_Stack, context uintptr) {
	var awakeSet, base, base1, bitSet, color, color1, contact, contactSim, contactSims, graphColors, shapeA, shapeB, shapes, userCollideTask, world, v10, v12, v17, v19 uintptr
	var bits uint64_t
	var bodyIdA, bodyIdB, colorIndex, contactCount, contactId, contactIdCapacity, contactIndex, count, endEventArrayIndex, i, i1, i2, i3, i4, j, localIndex, minRange, newCapacity, newCapacity1, nonTouchingCount, v1, v6 int32
	var ctz, flags, k, simFlags, v15 uint32_t
	var event ContactBeginTouchEvent
	var event1 ContactEndTouchEvent
	var shapeIdA, shapeIdB ShapeId
	var worldId uint16_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = awakeSet, base, base1, bitSet, bits, bodyIdA, bodyIdB, color, color1, colorIndex, contact, contactCount, contactId, contactIdCapacity, contactIndex, contactSim, contactSims, count, ctz, endEventArrayIndex, event, event1, flags, graphColors, i, i1, i2, i3, i4, j, k, localIndex, minRange, newCapacity, newCapacity1, nonTouchingCount, shapeA, shapeB, shapeIdA, shapeIdB, shapes, simFlags, userCollideTask, world, worldId, v1, v10, v12, v15, v17, v19, v6
	world = (*b2StepContext)(unsafe.Pointer(context)).World
	if !((*b2World)(unsafe.Pointer(world)).WorkerCount > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+10955, __ccgo_ts+15342, int32FromInt32(492)) != 0 {
		__builtin_trap(tls)
	}
	// Task that can be done in parallel with the narrow-phase
	// - rebuild the collision tree for dynamic and kinematic bodies to keep their query performance good
	// todo_erin move this to start when contacts are being created
	(*b2World)(unsafe.Pointer(world)).UserTreeTask = (*(*func(*_Stack, uintptr, int32, int32, uintptr, uintptr) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).EnqueueTaskFcn})))(tls, __ccgo_fp(b2UpdateTreesTask), int32(1), int32(1), world, (*b2World)(unsafe.Pointer(world)).UserTaskContext)
	*(*int32)(unsafe.Add(nil, world+1776)) += int32(1)
	if (*b2World)(unsafe.Pointer(world)).UserTreeTask == uintptrFromInt32(0) {
		v1 = 0
	} else {
		v1 = int32(1)
	}
	*(*int32)(unsafe.Add(nil, world+1772)) += v1
	// gather contacts into a single array for easier parallel-for
	contactCount = 0
	graphColors = world + 328
	i = 0
	for {
		if !(i < int32(_B2_GRAPH_COLOR_COUNT)) {
			break
		}
		contactCount = contactCount + (*(*b2GraphColor)(unsafe.Add(nil, graphColors+uintptr(i)*56))).ContactSims.Count
		goto _2
	_2:
		;
		i = i + 1
	}
	nonTouchingCount = (*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_awakeSet)*88))).ContactSims.Count
	contactCount = contactCount + nonTouchingCount
	if contactCount == 0 {
		return
	}
	contactSims = b2AllocateArenaItem(tls, world, int32FromUint64(uint64FromInt32(contactCount)*uint64(8)), __ccgo_ts+15810)
	contactIndex = 0
	i1 = 0
	for {
		if !(i1 < int32(_B2_GRAPH_COLOR_COUNT)) {
			break
		}
		color = graphColors + uintptr(i1)*56
		count = (*b2GraphColor)(unsafe.Pointer(color)).ContactSims.Count
		base = (*b2GraphColor)(unsafe.Pointer(color)).ContactSims.Data
		j = 0
		for {
			if !(j < count) {
				break
			}
			*(*uintptr)(unsafe.Add(nil, contactSims+uintptr(contactIndex)*8)) = base + uintptr(j)*176
			contactIndex = contactIndex + int32(1)
			goto _4
		_4:
			;
			j = j + 1
		}
		goto _3
	_3:
		;
		i1 = i1 + 1
	}
	base1 = (*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_awakeSet)*88))).ContactSims.Data
	i2 = 0
	for {
		if !(i2 < nonTouchingCount) {
			break
		}
		*(*uintptr)(unsafe.Add(nil, contactSims+uintptr(contactIndex)*8)) = base1 + uintptr(i2)*176
		contactIndex = contactIndex + int32(1)
		goto _5
	_5:
		;
		i2 = i2 + 1
	}
	if !(contactIndex == contactCount) && b2InternalAssertFcn(tls, __ccgo_ts+15819, __ccgo_ts+15342, int32FromInt32(544)) != 0 {
		__builtin_trap(tls)
	}
	(*b2StepContext)(unsafe.Pointer(context)).Contacts = contactSims
	v1 = (*b2IdPool)(unsafe.Pointer(world + 1120)).NextIndex
	goto _7
_7:
	// Contact bit set on ids because contact pointers are unstable as they move between touching and not touching.
	contactIdCapacity = v1
	i3 = 0
	for {
		if !(i3 < (*b2World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		b2SetBitCountAndClear(tls, (*b2World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i3)*56, uint32FromInt32(contactIdCapacity))
		goto _8
	_8:
		;
		i3 = i3 + 1
	}
	// Task should take at least 40us on a 4GHz CPU (10K cycles)
	minRange = int32(64)
	userCollideTask = (*(*func(*_Stack, uintptr, int32, int32, uintptr, uintptr) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).EnqueueTaskFcn})))(tls, __ccgo_fp(b2CollideTask), contactCount, minRange, context, (*b2World)(unsafe.Pointer(world)).UserTaskContext)
	*(*int32)(unsafe.Add(nil, world+1776)) += int32(1)
	if userCollideTask != uintptrFromInt32(0) {
		(*(*func(*_Stack, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).FinishTaskFcn})))(tls, userCollideTask, (*b2World)(unsafe.Pointer(world)).UserTaskContext)
	}
	b2FreeArenaItem(tls, world, contactSims)
	(*b2StepContext)(unsafe.Pointer(context)).Contacts = uintptrFromInt32(0)
	contactSims = uintptrFromInt32(0)
	// Serially update contact state
	// todo_erin bring this zone together with island merge
	// Bitwise OR all contact bits
	bitSet = (*b2World)(unsafe.Pointer(world)).TaskContexts.Data
	i4 = int32(1)
	for {
		if !(i4 < (*b2World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		b2InPlaceUnion(tls, bitSet, (*b2World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i4)*56)
		goto _9
	_9:
		;
		i4 = i4 + 1
	}
	v10 = world + 1064
	v1 = int32(b2_awakeSet)
	if !(0 <= v1 && v1 < (*b2SolverSetArray)(unsafe.Pointer(v10)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v12 = (*b2SolverSetArray)(unsafe.Pointer(v10)).Data + uintptr(v1)*88
	goto _13
_13:
	awakeSet = v12
	endEventArrayIndex = (*b2World)(unsafe.Pointer(world)).EndEventArrayIndex
	shapes = (*b2World)(unsafe.Pointer(world)).Shapes.Data
	worldId = (*b2World)(unsafe.Pointer(world)).WorldId
	// Process contact state changes. Iterate over set bits
	k = uint32(0)
	for {
		if !(k < (*b2BitSet)(unsafe.Pointer(bitSet)).BlockCount) {
			break
		}
		bits = *(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(bitSet)).Bits+uintptr(k)*8))
		for bits != uint64(0) {
			v15 = uint32FromInt32(__builtin_ctzll(tls, bits))
			goto _16
		_16:
			ctz = v15
			contactId = int32FromUint32(uint32FromInt32(64)*k + ctz)
			v17 = world + 1144
			v6 = contactId
			if !(0 <= v6 && v6 < (*b2ContactArray)(unsafe.Pointer(v17)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
				__builtin_trap(tls)
			}
			v19 = (*b2ContactArray)(unsafe.Pointer(v17)).Data + uintptr(v6)*68
			goto _20
		_20:
			contact = v19
			if !((*b2Contact)(unsafe.Pointer(contact)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+3637, __ccgo_ts+15342, int32FromInt32(596)) != 0 {
				__builtin_trap(tls)
			}
			colorIndex = (*b2Contact)(unsafe.Pointer(contact)).ColorIndex
			localIndex = (*b2Contact)(unsafe.Pointer(contact)).LocalIndex
			contactSim = uintptrFromInt32(0)
			if colorIndex != -int32(1) {
				// contact lives in constraint graph
				if !(0 <= colorIndex && colorIndex < int32(_B2_GRAPH_COLOR_COUNT)) && b2InternalAssertFcn(tls, __ccgo_ts+3038, __ccgo_ts+15342, int32FromInt32(605)) != 0 {
					__builtin_trap(tls)
				}
				color1 = graphColors + uintptr(colorIndex)*56
				v10 = color1 + 16
				v1 = localIndex
				if !(0 <= v1 && v1 < (*b2ContactSimArray)(unsafe.Pointer(v10)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(147)) != 0 {
					__builtin_trap(tls)
				}
				v12 = (*b2ContactSimArray)(unsafe.Pointer(v10)).Data + uintptr(v1)*176
				goto _24
			_24:
				contactSim = v12
			} else {
				v10 = awakeSet + 48
				v1 = localIndex
				if !(0 <= v1 && v1 < (*b2ContactSimArray)(unsafe.Pointer(v10)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(147)) != 0 {
					__builtin_trap(tls)
				}
				v12 = (*b2ContactSimArray)(unsafe.Pointer(v10)).Data + uintptr(v1)*176
				goto _28
			_28:
				contactSim = v12
			}
			shapeA = shapes + uintptr((*b2Contact)(unsafe.Pointer(contact)).ShapeIdA)*288
			shapeB = shapes + uintptr((*b2Contact)(unsafe.Pointer(contact)).ShapeIdB)*288
			shapeIdA = ShapeId{
				Index1:     (*b2Shape)(unsafe.Pointer(shapeA)).Id + int32(1),
				World0:     worldId,
				Generation: (*b2Shape)(unsafe.Pointer(shapeA)).Generation,
			}
			shapeIdB = ShapeId{
				Index1:     (*b2Shape)(unsafe.Pointer(shapeB)).Id + int32(1),
				World0:     worldId,
				Generation: (*b2Shape)(unsafe.Pointer(shapeB)).Generation,
			}
			flags = (*b2Contact)(unsafe.Pointer(contact)).Flags
			simFlags = (*b2ContactSim)(unsafe.Pointer(contactSim)).SimFlags
			if simFlags&uint32(b2_simDisjoint) != 0 {
				// Bounding boxes no longer overlap
				b2DestroyContact(tls, world, contact, boolUint8(false1 != 0))
				contact = uintptrFromInt32(0)
				contactSim = uintptrFromInt32(0)
			} else {
				if simFlags&uint32(b2_simStartedTouching) != 0 {
					if !((*b2Contact)(unsafe.Pointer(contact)).IslandId == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+7197, __ccgo_ts+15342, int32FromInt32(630)) != 0 {
						__builtin_trap(tls)
					}
					if flags&uint32(b2_contactEnableContactEvents) != 0 {
						event = ContactBeginTouchEvent{
							ShapeIdA: shapeIdA,
							ShapeIdB: shapeIdB,
							Manifold: (*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold,
						}
						v10 = world + 1360
						if (*b2ContactBeginTouchEventArray)(unsafe.Pointer(v10)).Count == (*b2ContactBeginTouchEventArray)(unsafe.Pointer(v10)).Capacity {
							if (*b2ContactBeginTouchEventArray)(unsafe.Pointer(v10)).Capacity < int32(2) {
								v1 = int32(2)
							} else {
								v1 = (*b2ContactBeginTouchEventArray)(unsafe.Pointer(v10)).Capacity + (*b2ContactBeginTouchEventArray)(unsafe.Pointer(v10)).Capacity>>int32(1)
							}
							newCapacity = v1
							b2ContactBeginTouchEventArray_Reserve(tls, v10, newCapacity)
						}
						*(*ContactBeginTouchEvent)(unsafe.Add(nil, (*b2ContactBeginTouchEventArray)(unsafe.Pointer(v10)).Data+uintptr((*b2ContactBeginTouchEventArray)(unsafe.Pointer(v10)).Count)*128)) = event
						*(*int32)(unsafe.Add(nil, v10+8)) += int32(1)
					}
					if !((*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.PointCount > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+2821, __ccgo_ts+15342, int32FromInt32(638)) != 0 {
						__builtin_trap(tls)
					}
					if !((*b2Contact)(unsafe.Pointer(contact)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+3637, __ccgo_ts+15342, int32FromInt32(639)) != 0 {
						__builtin_trap(tls)
					}
					// Link first because this wakes colliding bodies and ensures the body sims
					// are in the correct place.
					*(*uint32_t)(unsafe.Add(nil, contact+60)) |= uint32(b2_contactTouchingFlag)
					b2LinkContact(tls, world, contact)
					// Make sure these didn't change
					if !((*b2Contact)(unsafe.Pointer(contact)).ColorIndex == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+15848, __ccgo_ts+15342, int32FromInt32(647)) != 0 {
						__builtin_trap(tls)
					}
					if !((*b2Contact)(unsafe.Pointer(contact)).LocalIndex == localIndex) && b2InternalAssertFcn(tls, __ccgo_ts+15885, __ccgo_ts+15342, int32FromInt32(648)) != 0 {
						__builtin_trap(tls)
					}
					// Contact sim pointer may have become orphaned due to awake set growth,
					// so I just need to refresh it.
					v10 = awakeSet + 48
					v1 = localIndex
					if !(0 <= v1 && v1 < (*b2ContactSimArray)(unsafe.Pointer(v10)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(147)) != 0 {
						__builtin_trap(tls)
					}
					v12 = (*b2ContactSimArray)(unsafe.Pointer(v10)).Data + uintptr(v1)*176
					goto _34
				_34:
					contactSim = v12
					*(*uint32_t)(unsafe.Add(nil, contactSim+164)) &= uint32FromInt32(^int32(b2_simStartedTouching))
					b2AddContactToGraph(tls, world, contactSim, contact)
					b2RemoveNonTouchingContact(tls, world, int32(b2_awakeSet), localIndex)
					contactSim = uintptrFromInt32(0)
				} else {
					if simFlags&uint32(b2_simStoppedTouching) != 0 {
						*(*uint32_t)(unsafe.Add(nil, contactSim+164)) &= uint32FromInt32(^int32(b2_simStoppedTouching))
						*(*uint32_t)(unsafe.Add(nil, contact+60)) &= uint32FromInt32(^int32(b2_contactTouchingFlag))
						if (*b2Contact)(unsafe.Pointer(contact)).Flags&uint32(b2_contactEnableContactEvents) != 0 {
							event1 = ContactEndTouchEvent{
								ShapeIdA: shapeIdA,
								ShapeIdB: shapeIdB,
							}
							v10 = world + 1408 + uintptr(endEventArrayIndex)*16
							if (*b2ContactEndTouchEventArray)(unsafe.Pointer(v10)).Count == (*b2ContactEndTouchEventArray)(unsafe.Pointer(v10)).Capacity {
								if (*b2ContactEndTouchEventArray)(unsafe.Pointer(v10)).Capacity < int32(2) {
									v1 = int32(2)
								} else {
									v1 = (*b2ContactEndTouchEventArray)(unsafe.Pointer(v10)).Capacity + (*b2ContactEndTouchEventArray)(unsafe.Pointer(v10)).Capacity>>int32(1)
								}
								newCapacity1 = v1
								b2ContactEndTouchEventArray_Reserve(tls, v10, newCapacity1)
							}
							*(*ContactEndTouchEvent)(unsafe.Add(nil, (*b2ContactEndTouchEventArray)(unsafe.Pointer(v10)).Data+uintptr((*b2ContactEndTouchEventArray)(unsafe.Pointer(v10)).Count)*16)) = event1
							*(*int32)(unsafe.Add(nil, v10+8)) += int32(1)
						}
						if !((*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.PointCount == int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+14641, __ccgo_ts+15342, int32FromInt32(671)) != 0 {
							__builtin_trap(tls)
						}
						b2UnlinkContact(tls, world, contact)
						bodyIdA = (*(*b2ContactEdge)(unsafe.Add(nil, contact+12))).BodyId
						bodyIdB = (*(*b2ContactEdge)(unsafe.Add(nil, contact+12+1*12))).BodyId
						b2AddNonTouchingContact(tls, world, contact, contactSim)
						b2RemoveContactFromGraph(tls, world, bodyIdA, bodyIdB, colorIndex, localIndex)
						contact = uintptrFromInt32(0)
						contactSim = uintptrFromInt32(0)
					}
				}
			}
			// Clear the smallest set bit
			bits = bits & (bits - uint64(1))
		}
		goto _14
	_14:
		;
		k = k + 1
	}
	b2ValidateSolverSets(tls, world)
	b2ValidateContacts(tls, world)
}
