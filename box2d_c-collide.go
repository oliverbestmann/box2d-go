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
	(*(*DistanceInput)(unsafe.Pointer(bp))).ProxyA = b2MakeProxy(tls, shape, int32(1), float32FromFloat32(0))
	(*(*DistanceInput)(unsafe.Pointer(bp))).ProxyB = b2MakeProxy(tls, mover, int32(2), (*Capsule)(unsafe.Pointer(mover)).Radius)
	(*(*DistanceInput)(unsafe.Pointer(bp))).TransformA = b2Transform_identity
	(*(*DistanceInput)(unsafe.Pointer(bp))).TransformB = b2Transform_identity
	(*(*DistanceInput)(unsafe.Pointer(bp))).UseRadii = boolUint8(false1 != 0)
	totalRadius = (*Capsule)(unsafe.Pointer(mover)).Radius + (*Circle)(unsafe.Pointer(shape)).Radius
	*(*SimplexCache)(unsafe.Pointer(bp + 184)) = SimplexCache{}
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
	(*(*DistanceInput)(unsafe.Pointer(bp))).ProxyA = b2MakeProxy(tls, shape, int32(2), float32FromFloat32(0))
	(*(*DistanceInput)(unsafe.Pointer(bp))).ProxyB = b2MakeProxy(tls, mover, int32(2), (*Capsule)(unsafe.Pointer(mover)).Radius)
	(*(*DistanceInput)(unsafe.Pointer(bp))).TransformA = b2Transform_identity
	(*(*DistanceInput)(unsafe.Pointer(bp))).TransformB = b2Transform_identity
	(*(*DistanceInput)(unsafe.Pointer(bp))).UseRadii = boolUint8(false1 != 0)
	totalRadius = (*Capsule)(unsafe.Pointer(mover)).Radius + (*Capsule)(unsafe.Pointer(shape)).Radius
	*(*SimplexCache)(unsafe.Pointer(bp + 184)) = SimplexCache{}
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
	(*(*DistanceInput)(unsafe.Pointer(bp))).ProxyA = b2MakeProxy(tls, shape, (*Polygon)(unsafe.Pointer(shape)).Count, (*Polygon)(unsafe.Pointer(shape)).Radius)
	(*(*DistanceInput)(unsafe.Pointer(bp))).ProxyB = b2MakeProxy(tls, mover, int32(2), (*Capsule)(unsafe.Pointer(mover)).Radius)
	(*(*DistanceInput)(unsafe.Pointer(bp))).TransformA = b2Transform_identity
	(*(*DistanceInput)(unsafe.Pointer(bp))).TransformB = b2Transform_identity
	(*(*DistanceInput)(unsafe.Pointer(bp))).UseRadii = boolUint8(false1 != 0)
	totalRadius = (*Capsule)(unsafe.Pointer(mover)).Radius + (*Polygon)(unsafe.Pointer(shape)).Radius
	*(*SimplexCache)(unsafe.Pointer(bp + 184)) = SimplexCache{}
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
	(*(*DistanceInput)(unsafe.Pointer(bp))).ProxyA = b2MakeProxy(tls, shape, int32(2), float32FromFloat32(0))
	(*(*DistanceInput)(unsafe.Pointer(bp))).ProxyB = b2MakeProxy(tls, mover, int32(2), (*Capsule)(unsafe.Pointer(mover)).Radius)
	(*(*DistanceInput)(unsafe.Pointer(bp))).TransformA = b2Transform_identity
	(*(*DistanceInput)(unsafe.Pointer(bp))).TransformB = b2Transform_identity
	(*(*DistanceInput)(unsafe.Pointer(bp))).UseRadii = boolUint8(false1 != 0)
	totalRadius = (*Capsule)(unsafe.Pointer(mover)).Radius
	*(*SimplexCache)(unsafe.Pointer(bp + 184)) = SimplexCache{}
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
	var cA, cB, contactPointA, n, normal, pointA, pointB, v12, v13, v18, v19, v21, v22, v23, v26, v27, v29, v31, v32, v34, v36, v37, v39, v40, v42, v45, v46, v49, v50, v52, v53, v54, v56, v57, v58, v60, v61, v62, v7, v8, v9 Vec2
	var invLength, radiusA, radiusB, separation, x, y, v30, v35, v41 float32
	var mp, v25 uintptr
	var qr, v11, v3, v4, v44, v48, v5 Rot
	var _ /* distance at bp+112 */ float32
	var _ /* manifold at bp+0 */ Manifold
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, cA, cB, contactPointA, invLength, mp, n, normal, pointA, pointB, qr, radiusA, radiusB, separation, x, xf, y, v1, v11, v12, v13, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v29, v3, v30, v31, v32, v34, v35, v36, v37, v39, v4, v40, v41, v42, v44, v45, v46, v48, v49, v5, v50, v52, v53, v54, v56, v57, v58, v60, v61, v62, v7, v8, v9
	*(*Manifold)(unsafe.Pointer(bp)) = Manifold{}
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
	*(*float32)(unsafe.Pointer(v25)) = sqrtf(tls, float32(v26.X*v26.X)+float32(v26.Y*v26.Y))
	if *(*float32)(unsafe.Pointer(v25)) < float32FromFloat32(1.1920928955078125e-07) {
		v27 = Vec2{}
		goto _28
	}
	invLength = float32FromFloat32(1) / *(*float32)(unsafe.Pointer(v25))
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
	separation = *(*float32)(unsafe.Pointer(bp + 112)) - radiusA - radiusB
	if separation > float32(float32FromFloat32(4)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
		return *(*Manifold)(unsafe.Pointer(bp))
	}
	v29 = pointA
	v30 = radiusA
	v31 = normal
	v32 = Vec2{
		X: v29.X + float32(v30*v31.X),
		Y: v29.Y + float32(v30*v31.Y),
	}
	goto _33
_33:
	cA = v32
	v34 = pointB
	v35 = -radiusB
	v36 = normal
	v37 = Vec2{
		X: v34.X + float32(v35*v36.X),
		Y: v34.Y + float32(v35*v36.Y),
	}
	goto _38
_38:
	cB = v37
	v39 = cA
	v40 = cB
	v41 = float32FromFloat32(0.5)
	v42 = Vec2{
		X: float32((float32FromFloat32(1)-v41)*v39.X) + float32(v41*v40.X),
		Y: float32((float32FromFloat32(1)-v41)*v39.Y) + float32(v41*v40.Y),
	}
	goto _43
_43:
	contactPointA = v42
	v44 = xfA.Q
	v45 = normal
	v46 = Vec2{
		X: float32(v44.C*v45.X) - float32(v44.S*v45.Y),
		Y: float32(v44.S*v45.X) + float32(v44.C*v45.Y),
	}
	goto _47
_47:
	(*(*Manifold)(unsafe.Pointer(bp))).Normal = v46
	mp = bp + 12 + uintptr(0)*48
	v48 = xfA.Q
	v49 = contactPointA
	v50 = Vec2{
		X: float32(v48.C*v49.X) - float32(v48.S*v49.Y),
		Y: float32(v48.S*v49.X) + float32(v48.C*v49.Y),
	}
	goto _51
_51:
	(*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA = v50
	v52 = xfA.P
	v53 = xfB.P
	v54 = Vec2{
		X: v52.X - v53.X,
		Y: v52.Y - v53.Y,
	}
	goto _55
_55:
	v56 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
	v57 = v54
	v58 = Vec2{
		X: v56.X + v57.X,
		Y: v56.Y + v57.Y,
	}
	goto _59
_59:
	(*ManifoldPoint)(unsafe.Pointer(mp)).AnchorB = v58
	v60 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
	v61 = xfA.P
	v62 = Vec2{
		X: v60.X + v61.X,
		Y: v60.Y + v61.Y,
	}
	goto _63
_63:
	(*ManifoldPoint)(unsafe.Pointer(mp)).Point = v62
	(*ManifoldPoint)(unsafe.Pointer(mp)).Separation = separation
	(*ManifoldPoint)(unsafe.Pointer(mp)).Id = uint16(0)
	(*(*Manifold)(unsafe.Pointer(bp))).PointCount = int32(1)
	return *(*Manifold)(unsafe.Pointer(bp))
}

// C documentation
//
//	/// Compute the collision manifold between a capsule and circle
func b2CollideCapsuleAndCircle(tls *_Stack, capsuleA uintptr, xfA Transform, circleB uintptr, xfB Transform) (r Manifold) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var C, xf, v1, v15, v17, v2 Transform
	var cA, cB, contactPointA, e, n, normal, p1, p2, pA, pB, v12, v13, v18, v19, v21, v22, v23, v25, v26, v27, v29, v30, v33, v34, v35, v37, v38, v41, v42, v45, v47, v48, v50, v51, v52, v55, v56, v58, v60, v61, v63, v65, v66, v68, v69, v7, v71, v74, v75, v78, v79, v8, v81, v82, v83, v85, v86, v87, v89, v9, v90, v91 Vec2
	var invLength, radiusA, radiusB, s1, s11, s2, separation, x, y, v31, v39, v43, v46, v59, v64, v70 float32
	var mp, v54 uintptr
	var qr, v11, v3, v4, v5, v73, v77 Rot
	var _ /* distance at bp+112 */ float32
	var _ /* manifold at bp+0 */ Manifold
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, cA, cB, contactPointA, e, invLength, mp, n, normal, p1, p2, pA, pB, qr, radiusA, radiusB, s1, s11, s2, separation, x, xf, y, v1, v11, v12, v13, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v29, v3, v30, v31, v33, v34, v35, v37, v38, v39, v4, v41, v42, v43, v45, v46, v47, v48, v5, v50, v51, v52, v54, v55, v56, v58, v59, v60, v61, v63, v64, v65, v66, v68, v69, v7, v70, v71, v73, v74, v75, v77, v78, v79, v8, v81, v82, v83, v85, v86, v87, v89, v9, v90, v91
	*(*Manifold)(unsafe.Pointer(bp)) = Manifold{}
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
			v41 = e
			v42 = e
			v43 = float32(v41.X*v42.X) + float32(v41.Y*v42.Y)
			goto _44
		_44:
			// circle colliding with segment interior
			s1 = s11 / v43
			v45 = p1
			v46 = s1
			v47 = e
			v48 = Vec2{
				X: v45.X + float32(v46*v47.X),
				Y: v45.Y + float32(v46*v47.Y),
			}
			goto _49
		_49:
			pA = v48
		}
	}
	v50 = pB
	v51 = pA
	v52 = Vec2{
		X: v50.X - v51.X,
		Y: v50.Y - v51.Y,
	}
	goto _53
_53:
	v54 = bp + 112
	v55 = v52
	*(*float32)(unsafe.Pointer(v54)) = sqrtf(tls, float32(v55.X*v55.X)+float32(v55.Y*v55.Y))
	if *(*float32)(unsafe.Pointer(v54)) < float32FromFloat32(1.1920928955078125e-07) {
		v56 = Vec2{}
		goto _57
	}
	invLength = float32FromFloat32(1) / *(*float32)(unsafe.Pointer(v54))
	n = Vec2{
		X: float32(invLength * v55.X),
		Y: float32(invLength * v55.Y),
	}
	v56 = n
	goto _57
_57:
	normal = v56
	radiusA = (*Capsule)(unsafe.Pointer(capsuleA)).Radius
	radiusB = (*Circle)(unsafe.Pointer(circleB)).Radius
	separation = *(*float32)(unsafe.Pointer(bp + 112)) - radiusA - radiusB
	if separation > float32(float32FromFloat32(4)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
		return *(*Manifold)(unsafe.Pointer(bp))
	}
	v58 = pA
	v59 = radiusA
	v60 = normal
	v61 = Vec2{
		X: v58.X + float32(v59*v60.X),
		Y: v58.Y + float32(v59*v60.Y),
	}
	goto _62
_62:
	cA = v61
	v63 = pB
	v64 = -radiusB
	v65 = normal
	v66 = Vec2{
		X: v63.X + float32(v64*v65.X),
		Y: v63.Y + float32(v64*v65.Y),
	}
	goto _67
_67:
	cB = v66
	v68 = cA
	v69 = cB
	v70 = float32FromFloat32(0.5)
	v71 = Vec2{
		X: float32((float32FromFloat32(1)-v70)*v68.X) + float32(v70*v69.X),
		Y: float32((float32FromFloat32(1)-v70)*v68.Y) + float32(v70*v69.Y),
	}
	goto _72
_72:
	contactPointA = v71
	v73 = xfA.Q
	v74 = normal
	v75 = Vec2{
		X: float32(v73.C*v74.X) - float32(v73.S*v74.Y),
		Y: float32(v73.S*v74.X) + float32(v73.C*v74.Y),
	}
	goto _76
_76:
	(*(*Manifold)(unsafe.Pointer(bp))).Normal = v75
	mp = bp + 12 + uintptr(0)*48
	v77 = xfA.Q
	v78 = contactPointA
	v79 = Vec2{
		X: float32(v77.C*v78.X) - float32(v77.S*v78.Y),
		Y: float32(v77.S*v78.X) + float32(v77.C*v78.Y),
	}
	goto _80
_80:
	(*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA = v79
	v81 = xfA.P
	v82 = xfB.P
	v83 = Vec2{
		X: v81.X - v82.X,
		Y: v81.Y - v82.Y,
	}
	goto _84
_84:
	v85 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
	v86 = v83
	v87 = Vec2{
		X: v85.X + v86.X,
		Y: v85.Y + v86.Y,
	}
	goto _88
_88:
	(*ManifoldPoint)(unsafe.Pointer(mp)).AnchorB = v87
	v89 = xfA.P
	v90 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
	v91 = Vec2{
		X: v89.X + v90.X,
		Y: v89.Y + v90.Y,
	}
	goto _92
_92:
	(*ManifoldPoint)(unsafe.Pointer(mp)).Point = v91
	(*ManifoldPoint)(unsafe.Pointer(mp)).Separation = separation
	(*ManifoldPoint)(unsafe.Pointer(mp)).Id = uint16(0)
	(*(*Manifold)(unsafe.Pointer(bp))).PointCount = int32(1)
	return *(*Manifold)(unsafe.Pointer(bp))
}

func b2CollidePolygonAndCircle(tls *_Stack, polygonA uintptr, xfA Transform, circleB uintptr, xfB Transform) (r Manifold) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var C, xf, v1, v15, v17, v2 Transform
	var cA, cA1, cA2, cB, cB1, cB2, center, contactPointA, contactPointA1, contactPointA2, n, normal, normal1, normal2, v11, v21, v101, v102, v103, v105, v106, v107, v109, v110, v113, v114, v115, v117, v118, v12, v120, v121, v122, v124, v125, v128, v13, v130, v131, v133, v135, v136, v138, v139, v141, v144, v145, v148, v149, v151, v152, v153, v155, v156, v157, v159, v160, v161, v163, v164, v165, v167, v168, v172, v173, v175, v176, v177, v179, v18, v180, v183, v185, v186, v188, v19, v190, v191, v193, v194, v196, v199, v200, v202, v203, v204, v206, v207, v208, v210, v211, v212, v22, v23, v24, v26, v27, v31, v32, v33, v35, v36, v37, v39, v40, v43, v44, v45, v47, v48, v49, v51, v52, v55, v56, v57, v59, v60, v62, v63, v64, v66, v67, v7, v70, v72, v73, v75, v77, v78, v8, v80, v81, v83, v86, v87, v9, v90, v91, v93, v94, v95, v97, v98, v99 Vec2
	var i, normalIndex, vertIndex1, vertIndex2, vertexCount, v30 int32
	var invLength, length, radius, radiusA, radiusB, s2, separation, speculativeDistance, u1, u2, x, y, v1111, v126, v129, v134, v140, v169, v181, v184, v189, v195, v28, v41, v53, v68, v71, v76, v82 float32
	var mp, mp1, mp2, normals, vertices uintptr
	var qr, v111, v143, v147, v171, v198, v3, v4, v5, v85, v89 Rot
	var _ /* manifold at bp+0 */ Manifold
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, cA, cA1, cA2, cB, cB1, cB2, center, contactPointA, contactPointA1, contactPointA2, i, invLength, length, mp, mp1, mp2, n, normal, normal1, normal2, normalIndex, normals, qr, radius, radiusA, radiusB, s2, separation, speculativeDistance, u1, u2, v11, v21, vertIndex1, vertIndex2, vertexCount, vertices, x, xf, y, v1, v101, v102, v103, v105, v106, v107, v109, v111, v110, v1111, v113, v114, v115, v117, v118, v12, v120, v121, v122, v124, v125, v126, v128, v129, v13, v130, v131, v133, v134, v135, v136, v138, v139, v140, v141, v143, v144, v145, v147, v148, v149, v15, v151, v152, v153, v155, v156, v157, v159, v160, v161, v163, v164, v165, v167, v168, v169, v17, v171, v172, v173, v175, v176, v177, v179, v18, v180, v181, v183, v184, v185, v186, v188, v189, v19, v190, v191, v193, v194, v195, v196, v198, v199, v2, v200, v202, v203, v204, v206, v207, v208, v210, v211, v212, v22, v23, v24, v26, v27, v28, v3, v30, v31, v32, v33, v35, v36, v37, v39, v4, v40, v41, v43, v44, v45, v47, v48, v49, v5, v51, v52, v53, v55, v56, v57, v59, v60, v62, v63, v64, v66, v67, v68, v7, v70, v71, v72, v73, v75, v76, v77, v78, v8, v80, v81, v82, v83, v85, v86, v87, v89, v9, v90, v91, v93, v94, v95, v97, v98, v99
	*(*Manifold)(unsafe.Pointer(bp)) = Manifold{}
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
		v23 = *(*Vec2)(unsafe.Pointer(vertices + uintptr(i)*8))
		v24 = Vec2{
			X: v22.X - v23.X,
			Y: v22.Y - v23.Y,
		}
		goto _25
	_25:
		v26 = *(*Vec2)(unsafe.Pointer(normals + uintptr(i)*8))
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
		return *(*Manifold)(unsafe.Pointer(bp))
	}
	// Vertices of the reference edge.
	vertIndex1 = normalIndex
	if vertIndex1+int32(1) < vertexCount {
		v30 = vertIndex1 + int32(1)
	} else {
		v30 = 0
	}
	vertIndex2 = v30
	v11 = *(*Vec2)(unsafe.Pointer(vertices + uintptr(vertIndex1)*8))
	v21 = *(*Vec2)(unsafe.Pointer(vertices + uintptr(vertIndex2)*8))
	v31 = center
	v32 = v11
	v33 = Vec2{
		X: v31.X - v32.X,
		Y: v31.Y - v32.Y,
	}
	goto _34
_34:
	v35 = v21
	v36 = v11
	v37 = Vec2{
		X: v35.X - v36.X,
		Y: v35.Y - v36.Y,
	}
	goto _38
_38:
	v39 = v33
	v40 = v37
	v41 = float32(v39.X*v40.X) + float32(v39.Y*v40.Y)
	goto _42
_42:
	// Compute barycentric coordinates
	u1 = v41
	v43 = center
	v44 = v21
	v45 = Vec2{
		X: v43.X - v44.X,
		Y: v43.Y - v44.Y,
	}
	goto _46
_46:
	v47 = v11
	v48 = v21
	v49 = Vec2{
		X: v47.X - v48.X,
		Y: v47.Y - v48.Y,
	}
	goto _50
_50:
	v51 = v45
	v52 = v49
	v53 = float32(v51.X*v52.X) + float32(v51.Y*v52.Y)
	goto _54
_54:
	u2 = v53
	if u1 < float32FromFloat32(0) && separation > float32FromFloat32(1.1920928955078125e-07) {
		v55 = center
		v56 = v11
		v57 = Vec2{
			X: v55.X - v56.X,
			Y: v55.Y - v56.Y,
		}
		goto _58
	_58:
		v59 = v57
		length = sqrtf(tls, float32(v59.X*v59.X)+float32(v59.Y*v59.Y))
		if length < float32FromFloat32(1.1920928955078125e-07) {
			v60 = Vec2{}
			goto _61
		}
		invLength = float32FromFloat32(1) / length
		n = Vec2{
			X: float32(invLength * v59.X),
			Y: float32(invLength * v59.Y),
		}
		v60 = n
		goto _61
	_61:
		// Circle center is closest to v1 and safely outside the polygon
		normal = v60
		v62 = center
		v63 = v11
		v64 = Vec2{
			X: v62.X - v63.X,
			Y: v62.Y - v63.Y,
		}
		goto _65
	_65:
		v66 = v64
		v67 = normal
		v68 = float32(v66.X*v67.X) + float32(v66.Y*v67.Y)
		goto _69
	_69:
		separation = v68
		if separation > radius+speculativeDistance {
			return *(*Manifold)(unsafe.Pointer(bp))
		}
		v70 = v11
		v71 = radiusA
		v72 = normal
		v73 = Vec2{
			X: v70.X + float32(v71*v72.X),
			Y: v70.Y + float32(v71*v72.Y),
		}
		goto _74
	_74:
		cA = v73
		v75 = center
		v76 = radiusB
		v77 = normal
		v78 = Vec2{
			X: v75.X - float32(v76*v77.X),
			Y: v75.Y - float32(v76*v77.Y),
		}
		goto _79
	_79:
		cB = v78
		v80 = cA
		v81 = cB
		v82 = float32FromFloat32(0.5)
		v83 = Vec2{
			X: float32((float32FromFloat32(1)-v82)*v80.X) + float32(v82*v81.X),
			Y: float32((float32FromFloat32(1)-v82)*v80.Y) + float32(v82*v81.Y),
		}
		goto _84
	_84:
		contactPointA = v83
		v85 = xfA.Q
		v86 = normal
		v87 = Vec2{
			X: float32(v85.C*v86.X) - float32(v85.S*v86.Y),
			Y: float32(v85.S*v86.X) + float32(v85.C*v86.Y),
		}
		goto _88
	_88:
		(*(*Manifold)(unsafe.Pointer(bp))).Normal = v87
		mp = bp + 12 + uintptr(0)*48
		v89 = xfA.Q
		v90 = contactPointA
		v91 = Vec2{
			X: float32(v89.C*v90.X) - float32(v89.S*v90.Y),
			Y: float32(v89.S*v90.X) + float32(v89.C*v90.Y),
		}
		goto _92
	_92:
		(*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA = v91
		v93 = xfA.P
		v94 = xfB.P
		v95 = Vec2{
			X: v93.X - v94.X,
			Y: v93.Y - v94.Y,
		}
		goto _96
	_96:
		v97 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
		v98 = v95
		v99 = Vec2{
			X: v97.X + v98.X,
			Y: v97.Y + v98.Y,
		}
		goto _100
	_100:
		(*ManifoldPoint)(unsafe.Pointer(mp)).AnchorB = v99
		v101 = xfA.P
		v102 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
		v103 = Vec2{
			X: v101.X + v102.X,
			Y: v101.Y + v102.Y,
		}
		goto _104
	_104:
		(*ManifoldPoint)(unsafe.Pointer(mp)).Point = v103
		v105 = cB
		v106 = cA
		v107 = Vec2{
			X: v105.X - v106.X,
			Y: v105.Y - v106.Y,
		}
		goto _108
	_108:
		v109 = v107
		v110 = normal
		v1111 = float32(v109.X*v110.X) + float32(v109.Y*v110.Y)
		goto _112
	_112:
		(*ManifoldPoint)(unsafe.Pointer(mp)).Separation = v1111
		(*ManifoldPoint)(unsafe.Pointer(mp)).Id = uint16(0)
		(*(*Manifold)(unsafe.Pointer(bp))).PointCount = int32(1)
	} else {
		if u2 < float32FromFloat32(0) && separation > float32FromFloat32(1.1920928955078125e-07) {
			v113 = center
			v114 = v21
			v115 = Vec2{
				X: v113.X - v114.X,
				Y: v113.Y - v114.Y,
			}
			goto _116
		_116:
			v117 = v115
			length = sqrtf(tls, float32(v117.X*v117.X)+float32(v117.Y*v117.Y))
			if length < float32FromFloat32(1.1920928955078125e-07) {
				v118 = Vec2{}
				goto _119
			}
			invLength = float32FromFloat32(1) / length
			n = Vec2{
				X: float32(invLength * v117.X),
				Y: float32(invLength * v117.Y),
			}
			v118 = n
			goto _119
		_119:
			// Circle center is closest to v2 and safely outside the polygon
			normal1 = v118
			v120 = center
			v121 = v21
			v122 = Vec2{
				X: v120.X - v121.X,
				Y: v120.Y - v121.Y,
			}
			goto _123
		_123:
			v124 = v122
			v125 = normal1
			v126 = float32(v124.X*v125.X) + float32(v124.Y*v125.Y)
			goto _127
		_127:
			separation = v126
			if separation > radius+speculativeDistance {
				return *(*Manifold)(unsafe.Pointer(bp))
			}
			v128 = v21
			v129 = radiusA
			v130 = normal1
			v131 = Vec2{
				X: v128.X + float32(v129*v130.X),
				Y: v128.Y + float32(v129*v130.Y),
			}
			goto _132
		_132:
			cA1 = v131
			v133 = center
			v134 = radiusB
			v135 = normal1
			v136 = Vec2{
				X: v133.X - float32(v134*v135.X),
				Y: v133.Y - float32(v134*v135.Y),
			}
			goto _137
		_137:
			cB1 = v136
			v138 = cA1
			v139 = cB1
			v140 = float32FromFloat32(0.5)
			v141 = Vec2{
				X: float32((float32FromFloat32(1)-v140)*v138.X) + float32(v140*v139.X),
				Y: float32((float32FromFloat32(1)-v140)*v138.Y) + float32(v140*v139.Y),
			}
			goto _142
		_142:
			contactPointA1 = v141
			v143 = xfA.Q
			v144 = normal1
			v145 = Vec2{
				X: float32(v143.C*v144.X) - float32(v143.S*v144.Y),
				Y: float32(v143.S*v144.X) + float32(v143.C*v144.Y),
			}
			goto _146
		_146:
			(*(*Manifold)(unsafe.Pointer(bp))).Normal = v145
			mp1 = bp + 12 + uintptr(0)*48
			v147 = xfA.Q
			v148 = contactPointA1
			v149 = Vec2{
				X: float32(v147.C*v148.X) - float32(v147.S*v148.Y),
				Y: float32(v147.S*v148.X) + float32(v147.C*v148.Y),
			}
			goto _150
		_150:
			(*ManifoldPoint)(unsafe.Pointer(mp1)).AnchorA = v149
			v151 = xfA.P
			v152 = xfB.P
			v153 = Vec2{
				X: v151.X - v152.X,
				Y: v151.Y - v152.Y,
			}
			goto _154
		_154:
			v155 = (*ManifoldPoint)(unsafe.Pointer(mp1)).AnchorA
			v156 = v153
			v157 = Vec2{
				X: v155.X + v156.X,
				Y: v155.Y + v156.Y,
			}
			goto _158
		_158:
			(*ManifoldPoint)(unsafe.Pointer(mp1)).AnchorB = v157
			v159 = xfA.P
			v160 = (*ManifoldPoint)(unsafe.Pointer(mp1)).AnchorA
			v161 = Vec2{
				X: v159.X + v160.X,
				Y: v159.Y + v160.Y,
			}
			goto _162
		_162:
			(*ManifoldPoint)(unsafe.Pointer(mp1)).Point = v161
			v163 = cB1
			v164 = cA1
			v165 = Vec2{
				X: v163.X - v164.X,
				Y: v163.Y - v164.Y,
			}
			goto _166
		_166:
			v167 = v165
			v168 = normal1
			v169 = float32(v167.X*v168.X) + float32(v167.Y*v168.Y)
			goto _170
		_170:
			(*ManifoldPoint)(unsafe.Pointer(mp1)).Separation = v169
			(*ManifoldPoint)(unsafe.Pointer(mp1)).Id = uint16(0)
			(*(*Manifold)(unsafe.Pointer(bp))).PointCount = int32(1)
		} else {
			// Circle center is between v1 and v2. Center may be inside polygon
			normal2 = *(*Vec2)(unsafe.Pointer(normals + uintptr(normalIndex)*8))
			v171 = xfA.Q
			v172 = normal2
			v173 = Vec2{
				X: float32(v171.C*v172.X) - float32(v171.S*v172.Y),
				Y: float32(v171.S*v172.X) + float32(v171.C*v172.Y),
			}
			goto _174
		_174:
			(*(*Manifold)(unsafe.Pointer(bp))).Normal = v173
			v175 = center
			v176 = v11
			v177 = Vec2{
				X: v175.X - v176.X,
				Y: v175.Y - v176.Y,
			}
			goto _178
		_178:
			v179 = v177
			v180 = normal2
			v181 = float32(v179.X*v180.X) + float32(v179.Y*v180.Y)
			goto _182
		_182:
			v183 = center
			v184 = radiusA - v181
			v185 = normal2
			v186 = Vec2{
				X: v183.X + float32(v184*v185.X),
				Y: v183.Y + float32(v184*v185.Y),
			}
			goto _187
		_187:
			// cA is the projection of the circle center onto to the reference edge
			cA2 = v186
			v188 = center
			v189 = radiusB
			v190 = normal2
			v191 = Vec2{
				X: v188.X - float32(v189*v190.X),
				Y: v188.Y - float32(v189*v190.Y),
			}
			goto _192
		_192:
			// cB is the deepest point on the circle with respect to the reference edge
			cB2 = v191
			v193 = cA2
			v194 = cB2
			v195 = float32FromFloat32(0.5)
			v196 = Vec2{
				X: float32((float32FromFloat32(1)-v195)*v193.X) + float32(v195*v194.X),
				Y: float32((float32FromFloat32(1)-v195)*v193.Y) + float32(v195*v194.Y),
			}
			goto _197
		_197:
			contactPointA2 = v196
			// The contact point is the midpoint in world space
			mp2 = bp + 12 + uintptr(0)*48
			v198 = xfA.Q
			v199 = contactPointA2
			v200 = Vec2{
				X: float32(v198.C*v199.X) - float32(v198.S*v199.Y),
				Y: float32(v198.S*v199.X) + float32(v198.C*v199.Y),
			}
			goto _201
		_201:
			(*ManifoldPoint)(unsafe.Pointer(mp2)).AnchorA = v200
			v202 = xfA.P
			v203 = xfB.P
			v204 = Vec2{
				X: v202.X - v203.X,
				Y: v202.Y - v203.Y,
			}
			goto _205
		_205:
			v206 = (*ManifoldPoint)(unsafe.Pointer(mp2)).AnchorA
			v207 = v204
			v208 = Vec2{
				X: v206.X + v207.X,
				Y: v206.Y + v207.Y,
			}
			goto _209
		_209:
			(*ManifoldPoint)(unsafe.Pointer(mp2)).AnchorB = v208
			v210 = xfA.P
			v211 = (*ManifoldPoint)(unsafe.Pointer(mp2)).AnchorA
			v212 = Vec2{
				X: v210.X + v211.X,
				Y: v210.Y + v211.Y,
			}
			goto _213
		_213:
			(*ManifoldPoint)(unsafe.Pointer(mp2)).Point = v212
			(*ManifoldPoint)(unsafe.Pointer(mp2)).Separation = separation - radius
			(*ManifoldPoint)(unsafe.Pointer(mp2)).Id = uint16(0)
			(*(*Manifold)(unsafe.Pointer(bp))).PointCount = int32(1)
		}
	}
	return *(*Manifold)(unsafe.Pointer(bp))
}

// C documentation
//
//	// Follows Ericson 5.1.9 Closest Points of Two Line Segments
//	// Adds some logic to support clipping to get two contact points
func b2CollideCapsules(tls *_Stack, capsuleA uintptr, xfA Transform, capsuleB uintptr, xfB Transform) (r Manifold) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var C, sfA, xf, v10, v23, v29, v33, v9 Transform
	var c, c1, c2, closest1, closest2, cp, cp1, cq, cq1, d1, d2, n, n1, normal, normalA, normalB, origin, p1, p2, q11, q21, r1, u1, u2, v100, v101, v105, v106, v109, v110, v112, v113, v114, v116, v117, v120, v121, v122, v124, v125, v128, v129, v130, v132, v133, v136, v137, v138, v140, v141, v144, v145, v147, v148, v149, v15, v151, v152, v155, v156, v157, v159, v16, v160, v165, v166, v168, v169, v17, v171, v172, v173, v175, v176, v179, v180, v181, v183, v184, v189, v190, v192, v193, v195, v197, v198, v2, v20, v200, v202, v203, v205, v207, v208, v21, v210, v212, v213, v214, v216, v217, v220, v221, v222, v224, v225, v228, v230, v231, v233, v235, v236, v238, v239, v241, v242, v244, v246, v247, v249, v25, v251, v252, v254, v256, v257, v259, v26, v261, v262, v263, v265, v266, v269, v27, v270, v271, v273, v274, v277, v279, v280, v282, v284, v285, v287, v288, v289, v291, v292, v295, v296, v298, v299, v3, v30, v301, v303, v304, v306, v308, v309, v31, v313, v314, v316, v319, v320, v323, v324, v325, v328, v329, v331, v332, v333, v335, v336, v337, v339, v34, v340, v341, v35, v37, v38, v39, v41, v42, v43, v45, v46, v49, v5, v50, v53, v54, v55, v57, v58, v6, v61, v62, v65, v66, v7, v90, v92, v93, v95, v97, v98 Vec2
	var d12, dd1, dd2, denom, distance, distanceSquared, epsSqr, f1, f2, fp1, fp2, fq1, fq2, invLength, invLength1, length, maxDistance, radius, radiusA, radiusB, rd1, rd2, s1n, s1n1, s1p, s1p1, separationA, separationB, sp, sp1, sq, sq1, ss1, ss11, ss2, ss21, x, y, v102, v118, v126, v134, v142, v153, v161, v163, v164, v177, v185, v187, v188, v194, v199, v204, v209, v218, v226, v229, v234, v243, v248, v253, v258, v267, v275, v278, v283, v293, v302, v307, v315, v47, v51, v59, v63, v67, v69, v70, v71, v72, v74, v75, v76, v77, v78, v79, v81, v82, v83, v84, v85, v86, v88, v89, v91, v96 float32
	var i, i1, i2, v311, v312 int32
	var mp, mp1, mp2, v104, v108 uintptr
	var outsideA, outsideB uint8
	var qr, v1, v11, v12, v13, v19, v318, v327 Rot
	var _ /* length1 at bp+112 */ float32
	var _ /* length2 at bp+116 */ float32
	var _ /* manifold at bp+0 */ Manifold
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, c, c1, c2, closest1, closest2, cp, cp1, cq, cq1, d1, d12, d2, dd1, dd2, denom, distance, distanceSquared, epsSqr, f1, f2, fp1, fp2, fq1, fq2, i, i1, i2, invLength, invLength1, length, maxDistance, mp, mp1, mp2, n, n1, normal, normalA, normalB, origin, outsideA, outsideB, p1, p2, q11, q21, qr, r1, radius, radiusA, radiusB, rd1, rd2, s1n, s1n1, s1p, s1p1, separationA, separationB, sfA, sp, sp1, sq, sq1, ss1, ss11, ss2, ss21, u1, u2, x, xf, y, v1, v10, v100, v101, v102, v104, v105, v106, v108, v109, v11, v110, v112, v113, v114, v116, v117, v118, v12, v120, v121, v122, v124, v125, v126, v128, v129, v13, v130, v132, v133, v134, v136, v137, v138, v140, v141, v142, v144, v145, v147, v148, v149, v15, v151, v152, v153, v155, v156, v157, v159, v16, v160, v161, v163, v164, v165, v166, v168, v169, v17, v171, v172, v173, v175, v176, v177, v179, v180, v181, v183, v184, v185, v187, v188, v189, v19, v190, v192, v193, v194, v195, v197, v198, v199, v2, v20, v200, v202, v203, v204, v205, v207, v208, v209, v21, v210, v212, v213, v214, v216, v217, v218, v220, v221, v222, v224, v225, v226, v228, v229, v23, v230, v231, v233, v234, v235, v236, v238, v239, v241, v242, v243, v244, v246, v247, v248, v249, v25, v251, v252, v253, v254, v256, v257, v258, v259, v26, v261, v262, v263, v265, v266, v267, v269, v27, v270, v271, v273, v274, v275, v277, v278, v279, v280, v282, v283, v284, v285, v287, v288, v289, v29, v291, v292, v293, v295, v296, v298, v299, v3, v30, v301, v302, v303, v304, v306, v307, v308, v309, v31, v311, v312, v313, v314, v315, v316, v318, v319, v320, v323, v324, v325, v327, v328, v329, v33, v331, v332, v333, v335, v336, v337, v339, v34, v340, v341, v35, v37, v38, v39, v41, v42, v43, v45, v46, v47, v49, v5, v50, v51, v53, v54, v55, v57, v58, v59, v6, v61, v62, v63, v65, v66, v67, v69, v7, v70, v71, v72, v74, v75, v76, v77, v78, v79, v81, v82, v83, v84, v85, v86, v88, v89, v9, v90, v91, v92, v93, v95, v96, v97, v98
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
	v53 = p1
	v54 = p2
	v55 = Vec2{
		X: v53.X - v54.X,
		Y: v53.Y - v54.Y,
	}
	goto _56
_56:
	r1 = v55
	v57 = r1
	v58 = d1
	v59 = float32(v57.X*v58.X) + float32(v57.Y*v58.Y)
	goto _60
_60:
	rd1 = v59
	v61 = r1
	v62 = d2
	v63 = float32(v61.X*v62.X) + float32(v61.Y*v62.Y)
	goto _64
_64:
	rd2 = v63
	v65 = d1
	v66 = d2
	v67 = float32(v65.X*v66.X) + float32(v65.Y*v66.Y)
	goto _68
_68:
	d12 = v67
	denom = float32(dd1*dd2) - float32(d12*d12)
	// Fraction on segment 1
	f1 = float32FromFloat32(0)
	if denom != float32FromFloat32(0) {
		// not parallel
		v69 = (float32(d12*rd2) - float32(rd1*dd2)) / denom
		v70 = float32FromFloat32(0)
		v71 = float32FromFloat32(1)
		if v69 < v70 {
			v74 = v70
		} else {
			if v69 > v71 {
				v75 = v71
			} else {
				v75 = v69
			}
			v74 = v75
		}
		v72 = v74
		goto _73
	_73:
		f1 = v72
	}
	// Compute point on segment 2 closest to p1 + f1 * d1
	f2 = (float32(d12*f1) + rd2) / dd2
	// Clamping of segment 2 requires a do over on segment 1
	if f2 < float32FromFloat32(0) {
		f2 = float32FromFloat32(0)
		v76 = -rd1 / dd1
		v77 = float32FromFloat32(0)
		v78 = float32FromFloat32(1)
		if v76 < v77 {
			v81 = v77
		} else {
			if v76 > v78 {
				v82 = v78
			} else {
				v82 = v76
			}
			v81 = v82
		}
		v79 = v81
		goto _80
	_80:
		f1 = v79
	} else {
		if f2 > float32FromFloat32(1) {
			f2 = float32FromFloat32(1)
			v83 = (d12 - rd1) / dd1
			v84 = float32FromFloat32(0)
			v85 = float32FromFloat32(1)
			if v83 < v84 {
				v88 = v84
			} else {
				if v83 > v85 {
					v89 = v85
				} else {
					v89 = v83
				}
				v88 = v89
			}
			v86 = v88
			goto _87
		_87:
			f1 = v86
		}
	}
	v90 = p1
	v91 = f1
	v92 = d1
	v93 = Vec2{
		X: v90.X + float32(v91*v92.X),
		Y: v90.Y + float32(v91*v92.Y),
	}
	goto _94
_94:
	closest1 = v93
	v95 = p2
	v96 = f2
	v97 = d2
	v98 = Vec2{
		X: v95.X + float32(v96*v97.X),
		Y: v95.Y + float32(v96*v97.Y),
	}
	goto _99
_99:
	closest2 = v98
	v100 = closest1
	v101 = closest2
	c = Vec2{
		X: v101.X - v100.X,
		Y: v101.Y - v100.Y,
	}
	v102 = float32(c.X*c.X) + float32(c.Y*c.Y)
	goto _103
_103:
	distanceSquared = v102
	*(*Manifold)(unsafe.Pointer(bp)) = Manifold{}
	radiusA = (*Capsule)(unsafe.Pointer(capsuleA)).Radius
	radiusB = (*Capsule)(unsafe.Pointer(capsuleB)).Radius
	radius = radiusA + radiusB
	maxDistance = radius + float32(float32FromFloat32(4)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter))
	if distanceSquared > float32(maxDistance*maxDistance) {
		return *(*Manifold)(unsafe.Pointer(bp))
	}
	distance = sqrtf(tls, distanceSquared)
	v104 = bp + 112
	v105 = d1
	*(*float32)(unsafe.Pointer(v104)) = sqrtf(tls, float32(v105.X*v105.X)+float32(v105.Y*v105.Y))
	if *(*float32)(unsafe.Pointer(v104)) < float32FromFloat32(1.1920928955078125e-07) {
		v106 = Vec2{}
		goto _107
	}
	invLength1 = float32FromFloat32(1) / *(*float32)(unsafe.Pointer(v104))
	n1 = Vec2{
		X: float32(invLength1 * v105.X),
		Y: float32(invLength1 * v105.Y),
	}
	v106 = n1
	goto _107
_107:
	u1 = v106
	v108 = bp + 116
	v109 = d2
	*(*float32)(unsafe.Pointer(v108)) = sqrtf(tls, float32(v109.X*v109.X)+float32(v109.Y*v109.Y))
	if *(*float32)(unsafe.Pointer(v108)) < float32FromFloat32(1.1920928955078125e-07) {
		v110 = Vec2{}
		goto _111
	}
	invLength1 = float32FromFloat32(1) / *(*float32)(unsafe.Pointer(v108))
	n1 = Vec2{
		X: float32(invLength1 * v109.X),
		Y: float32(invLength1 * v109.Y),
	}
	v110 = n1
	goto _111
_111:
	u2 = v110
	v112 = p2
	v113 = p1
	v114 = Vec2{
		X: v112.X - v113.X,
		Y: v112.Y - v113.Y,
	}
	goto _115
_115:
	v116 = v114
	v117 = u1
	v118 = float32(v116.X*v117.X) + float32(v116.Y*v117.Y)
	goto _119
_119:
	// Does segment B project outside segment A?
	fp2 = v118
	v120 = q21
	v121 = p1
	v122 = Vec2{
		X: v120.X - v121.X,
		Y: v120.Y - v121.Y,
	}
	goto _123
_123:
	v124 = v122
	v125 = u1
	v126 = float32(v124.X*v125.X) + float32(v124.Y*v125.Y)
	goto _127
_127:
	fq2 = v126
	outsideA = boolUint8(fp2 <= float32FromFloat32(0) && fq2 <= float32FromFloat32(0) || fp2 >= *(*float32)(unsafe.Pointer(bp + 112)) && fq2 >= *(*float32)(unsafe.Pointer(bp + 112)))
	v128 = p1
	v129 = p2
	v130 = Vec2{
		X: v128.X - v129.X,
		Y: v128.Y - v129.Y,
	}
	goto _131
_131:
	v132 = v130
	v133 = u2
	v134 = float32(v132.X*v133.X) + float32(v132.Y*v133.Y)
	goto _135
_135:
	// Does segment A project outside segment B?
	fp1 = v134
	v136 = q11
	v137 = p2
	v138 = Vec2{
		X: v136.X - v137.X,
		Y: v136.Y - v137.Y,
	}
	goto _139
_139:
	v140 = v138
	v141 = u2
	v142 = float32(v140.X*v141.X) + float32(v140.Y*v141.Y)
	goto _143
_143:
	fq1 = v142
	outsideB = boolUint8(fp1 <= float32FromFloat32(0) && fq1 <= float32FromFloat32(0) || fp1 >= *(*float32)(unsafe.Pointer(bp + 116)) && fq1 >= *(*float32)(unsafe.Pointer(bp + 116)))
	if int32FromUint8(outsideA) == false1 && int32FromUint8(outsideB) == false1 {
		v144 = u1
		v145 = Vec2{
			X: -v144.Y,
			Y: v144.X,
		}
		goto _146
	_146:
		normalA = v145
		v147 = p2
		v148 = p1
		v149 = Vec2{
			X: v147.X - v148.X,
			Y: v147.Y - v148.Y,
		}
		goto _150
	_150:
		v151 = v149
		v152 = normalA
		v153 = float32(v151.X*v152.X) + float32(v151.Y*v152.Y)
		goto _154
	_154:
		ss1 = v153
		v155 = q21
		v156 = p1
		v157 = Vec2{
			X: v155.X - v156.X,
			Y: v155.Y - v156.Y,
		}
		goto _158
	_158:
		v159 = v157
		v160 = normalA
		v161 = float32(v159.X*v160.X) + float32(v159.Y*v160.Y)
		goto _162
	_162:
		ss2 = v161
		if ss1 < ss2 {
			v163 = ss1
		} else {
			v163 = ss2
		}
		s1p = v163
		if -ss1 < -ss2 {
			v164 = -ss1
		} else {
			v164 = -ss2
		}
		s1n = v164
		if s1p > s1n {
			separationA = s1p
		} else {
			separationA = s1n
			v165 = normalA
			v166 = Vec2{
				X: -v165.X,
				Y: -v165.Y,
			}
			goto _167
		_167:
			normalA = v166
		}
		v168 = u2
		v169 = Vec2{
			X: -v168.Y,
			Y: v168.X,
		}
		goto _170
	_170:
		normalB = v169
		v171 = p1
		v172 = p2
		v173 = Vec2{
			X: v171.X - v172.X,
			Y: v171.Y - v172.Y,
		}
		goto _174
	_174:
		v175 = v173
		v176 = normalB
		v177 = float32(v175.X*v176.X) + float32(v175.Y*v176.Y)
		goto _178
	_178:
		ss11 = v177
		v179 = q11
		v180 = p2
		v181 = Vec2{
			X: v179.X - v180.X,
			Y: v179.Y - v180.Y,
		}
		goto _182
	_182:
		v183 = v181
		v184 = normalB
		v185 = float32(v183.X*v184.X) + float32(v183.Y*v184.Y)
		goto _186
	_186:
		ss21 = v185
		if ss11 < ss21 {
			v187 = ss11
		} else {
			v187 = ss21
		}
		s1p1 = v187
		if -ss11 < -ss21 {
			v188 = -ss11
		} else {
			v188 = -ss21
		}
		s1n1 = v188
		if s1p1 > s1n1 {
			separationB = s1p1
		} else {
			separationB = s1n1
			v189 = normalB
			v190 = Vec2{
				X: -v189.X,
				Y: -v189.Y,
			}
			goto _191
		_191:
			normalB = v190
		}
		// biased to avoid feature flip-flop
		// todo more testing?
		if separationA+float32(float32FromFloat32(0.1)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) >= separationB {
			(*(*Manifold)(unsafe.Pointer(bp))).Normal = normalA
			cp = p2
			cq = q21
			// clip to p1
			if fp2 < float32FromFloat32(0) && fq2 > float32FromFloat32(0) {
				v192 = p2
				v193 = q21
				v194 = (float32FromFloat32(0) - fp2) / (fq2 - fp2)
				v195 = Vec2{
					X: float32((float32FromFloat32(1)-v194)*v192.X) + float32(v194*v193.X),
					Y: float32((float32FromFloat32(1)-v194)*v192.Y) + float32(v194*v193.Y),
				}
				goto _196
			_196:
				cp = v195
			} else {
				if fq2 < float32FromFloat32(0) && fp2 > float32FromFloat32(0) {
					v197 = q21
					v198 = p2
					v199 = (float32FromFloat32(0) - fq2) / (fp2 - fq2)
					v200 = Vec2{
						X: float32((float32FromFloat32(1)-v199)*v197.X) + float32(v199*v198.X),
						Y: float32((float32FromFloat32(1)-v199)*v197.Y) + float32(v199*v198.Y),
					}
					goto _201
				_201:
					cq = v200
				}
			}
			// clip to q1
			if fp2 > *(*float32)(unsafe.Pointer(bp + 112)) && fq2 < *(*float32)(unsafe.Pointer(bp + 112)) {
				v202 = p2
				v203 = q21
				v204 = (fp2 - *(*float32)(unsafe.Pointer(bp + 112))) / (fp2 - fq2)
				v205 = Vec2{
					X: float32((float32FromFloat32(1)-v204)*v202.X) + float32(v204*v203.X),
					Y: float32((float32FromFloat32(1)-v204)*v202.Y) + float32(v204*v203.Y),
				}
				goto _206
			_206:
				cp = v205
			} else {
				if fq2 > *(*float32)(unsafe.Pointer(bp + 112)) && fp2 < *(*float32)(unsafe.Pointer(bp + 112)) {
					v207 = q21
					v208 = p2
					v209 = (fq2 - *(*float32)(unsafe.Pointer(bp + 112))) / (fq2 - fp2)
					v210 = Vec2{
						X: float32((float32FromFloat32(1)-v209)*v207.X) + float32(v209*v208.X),
						Y: float32((float32FromFloat32(1)-v209)*v207.Y) + float32(v209*v208.Y),
					}
					goto _211
				_211:
					cq = v210
				}
			}
			v212 = cp
			v213 = p1
			v214 = Vec2{
				X: v212.X - v213.X,
				Y: v212.Y - v213.Y,
			}
			goto _215
		_215:
			v216 = v214
			v217 = normalA
			v218 = float32(v216.X*v217.X) + float32(v216.Y*v217.Y)
			goto _219
		_219:
			sp = v218
			v220 = cq
			v221 = p1
			v222 = Vec2{
				X: v220.X - v221.X,
				Y: v220.Y - v221.Y,
			}
			goto _223
		_223:
			v224 = v222
			v225 = normalA
			v226 = float32(v224.X*v225.X) + float32(v224.Y*v225.Y)
			goto _227
		_227:
			sq = v226
			if sp <= distance+float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter) || sq <= distance+float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter) {
				mp = bp + 12 + uintptr(0)*48
				v228 = cp
				v229 = float32(float32FromFloat32(0.5) * (radiusA - radiusB - sp))
				v230 = normalA
				v231 = Vec2{
					X: v228.X + float32(v229*v230.X),
					Y: v228.Y + float32(v229*v230.Y),
				}
				goto _232
			_232:
				(*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA = v231
				(*ManifoldPoint)(unsafe.Pointer(mp)).Separation = sp - radius
				(*ManifoldPoint)(unsafe.Pointer(mp)).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(int32FromInt32(0)))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(int32FromInt32(0))))
				mp = bp + 12 + uintptr(1)*48
				v233 = cq
				v234 = float32(float32FromFloat32(0.5) * (radiusA - radiusB - sq))
				v235 = normalA
				v236 = Vec2{
					X: v233.X + float32(v234*v235.X),
					Y: v233.Y + float32(v234*v235.Y),
				}
				goto _237
			_237:
				(*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA = v236
				(*ManifoldPoint)(unsafe.Pointer(mp)).Separation = sq - radius
				(*ManifoldPoint)(unsafe.Pointer(mp)).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(int32FromInt32(0)))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(int32FromInt32(1))))
				(*(*Manifold)(unsafe.Pointer(bp))).PointCount = int32(2)
			}
		} else {
			// normal always points from A to B
			v238 = normalB
			v239 = Vec2{
				X: -v238.X,
				Y: -v238.Y,
			}
			goto _240
		_240:
			(*(*Manifold)(unsafe.Pointer(bp))).Normal = v239
			cp1 = p1
			cq1 = q11
			// clip to p2
			if fp1 < float32FromFloat32(0) && fq1 > float32FromFloat32(0) {
				v241 = p1
				v242 = q11
				v243 = (float32FromFloat32(0) - fp1) / (fq1 - fp1)
				v244 = Vec2{
					X: float32((float32FromFloat32(1)-v243)*v241.X) + float32(v243*v242.X),
					Y: float32((float32FromFloat32(1)-v243)*v241.Y) + float32(v243*v242.Y),
				}
				goto _245
			_245:
				cp1 = v244
			} else {
				if fq1 < float32FromFloat32(0) && fp1 > float32FromFloat32(0) {
					v246 = q11
					v247 = p1
					v248 = (float32FromFloat32(0) - fq1) / (fp1 - fq1)
					v249 = Vec2{
						X: float32((float32FromFloat32(1)-v248)*v246.X) + float32(v248*v247.X),
						Y: float32((float32FromFloat32(1)-v248)*v246.Y) + float32(v248*v247.Y),
					}
					goto _250
				_250:
					cq1 = v249
				}
			}
			// clip to q2
			if fp1 > *(*float32)(unsafe.Pointer(bp + 116)) && fq1 < *(*float32)(unsafe.Pointer(bp + 116)) {
				v251 = p1
				v252 = q11
				v253 = (fp1 - *(*float32)(unsafe.Pointer(bp + 116))) / (fp1 - fq1)
				v254 = Vec2{
					X: float32((float32FromFloat32(1)-v253)*v251.X) + float32(v253*v252.X),
					Y: float32((float32FromFloat32(1)-v253)*v251.Y) + float32(v253*v252.Y),
				}
				goto _255
			_255:
				cp1 = v254
			} else {
				if fq1 > *(*float32)(unsafe.Pointer(bp + 116)) && fp1 < *(*float32)(unsafe.Pointer(bp + 116)) {
					v256 = q11
					v257 = p1
					v258 = (fq1 - *(*float32)(unsafe.Pointer(bp + 116))) / (fq1 - fp1)
					v259 = Vec2{
						X: float32((float32FromFloat32(1)-v258)*v256.X) + float32(v258*v257.X),
						Y: float32((float32FromFloat32(1)-v258)*v256.Y) + float32(v258*v257.Y),
					}
					goto _260
				_260:
					cq1 = v259
				}
			}
			v261 = cp1
			v262 = p2
			v263 = Vec2{
				X: v261.X - v262.X,
				Y: v261.Y - v262.Y,
			}
			goto _264
		_264:
			v265 = v263
			v266 = normalB
			v267 = float32(v265.X*v266.X) + float32(v265.Y*v266.Y)
			goto _268
		_268:
			sp1 = v267
			v269 = cq1
			v270 = p2
			v271 = Vec2{
				X: v269.X - v270.X,
				Y: v269.Y - v270.Y,
			}
			goto _272
		_272:
			v273 = v271
			v274 = normalB
			v275 = float32(v273.X*v274.X) + float32(v273.Y*v274.Y)
			goto _276
		_276:
			sq1 = v275
			if sp1 <= distance+float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter) || sq1 <= distance+float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter) {
				mp1 = bp + 12 + uintptr(0)*48
				v277 = cp1
				v278 = float32(float32FromFloat32(0.5) * (radiusB - radiusA - sp1))
				v279 = normalB
				v280 = Vec2{
					X: v277.X + float32(v278*v279.X),
					Y: v277.Y + float32(v278*v279.Y),
				}
				goto _281
			_281:
				(*ManifoldPoint)(unsafe.Pointer(mp1)).AnchorA = v280
				(*ManifoldPoint)(unsafe.Pointer(mp1)).Separation = sp1 - radius
				(*ManifoldPoint)(unsafe.Pointer(mp1)).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(int32FromInt32(0)))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(int32FromInt32(0))))
				mp1 = bp + 12 + uintptr(1)*48
				v282 = cq1
				v283 = float32(float32FromFloat32(0.5) * (radiusB - radiusA - sq1))
				v284 = normalB
				v285 = Vec2{
					X: v282.X + float32(v283*v284.X),
					Y: v282.Y + float32(v283*v284.Y),
				}
				goto _286
			_286:
				(*ManifoldPoint)(unsafe.Pointer(mp1)).AnchorA = v285
				(*ManifoldPoint)(unsafe.Pointer(mp1)).Separation = sq1 - radius
				(*ManifoldPoint)(unsafe.Pointer(mp1)).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(int32FromInt32(1)))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(int32FromInt32(0))))
				(*(*Manifold)(unsafe.Pointer(bp))).PointCount = int32(2)
			}
		}
	}
	if (*(*Manifold)(unsafe.Pointer(bp))).PointCount == 0 {
		v287 = closest2
		v288 = closest1
		v289 = Vec2{
			X: v287.X - v288.X,
			Y: v287.Y - v288.Y,
		}
		goto _290
	_290:
		// single point collision
		normal = v289
		v291 = normal
		v292 = normal
		v293 = float32(v291.X*v292.X) + float32(v291.Y*v292.Y)
		goto _294
	_294:
		if v293 > epsSqr {
			v295 = normal
			length = sqrtf(tls, float32(v295.X*v295.X)+float32(v295.Y*v295.Y))
			if length < float32FromFloat32(1.1920928955078125e-07) {
				v296 = Vec2{}
				goto _297
			}
			invLength = float32FromFloat32(1) / length
			n = Vec2{
				X: float32(invLength * v295.X),
				Y: float32(invLength * v295.Y),
			}
			v296 = n
			goto _297
		_297:
			normal = v296
		} else {
			v298 = u1
			v299 = Vec2{
				X: -v298.Y,
				Y: v298.X,
			}
			goto _300
		_300:
			normal = v299
		}
		v301 = closest1
		v302 = radiusA
		v303 = normal
		v304 = Vec2{
			X: v301.X + float32(v302*v303.X),
			Y: v301.Y + float32(v302*v303.Y),
		}
		goto _305
	_305:
		c1 = v304
		v306 = closest2
		v307 = -radiusB
		v308 = normal
		v309 = Vec2{
			X: v306.X + float32(v307*v308.X),
			Y: v306.Y + float32(v307*v308.Y),
		}
		goto _310
	_310:
		c2 = v309
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
		(*(*Manifold)(unsafe.Pointer(bp))).Normal = normal
		v313 = c1
		v314 = c2
		v315 = float32FromFloat32(0.5)
		v316 = Vec2{
			X: float32((float32FromFloat32(1)-v315)*v313.X) + float32(v315*v314.X),
			Y: float32((float32FromFloat32(1)-v315)*v313.Y) + float32(v315*v314.Y),
		}
		goto _317
	_317:
		(*(*ManifoldPoint)(unsafe.Pointer(bp + 12))).AnchorA = v316
		(*(*ManifoldPoint)(unsafe.Pointer(bp + 12))).Separation = sqrtf(tls, distanceSquared) - radius
		(*(*ManifoldPoint)(unsafe.Pointer(bp + 12))).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(i1))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(i2)))
		(*(*Manifold)(unsafe.Pointer(bp))).PointCount = int32(1)
	}
	// Convert manifold to world space
	v318 = xfA.Q
	v319 = (*(*Manifold)(unsafe.Pointer(bp))).Normal
	v320 = Vec2{
		X: float32(v318.C*v319.X) - float32(v318.S*v319.Y),
		Y: float32(v318.S*v319.X) + float32(v318.C*v319.Y),
	}
	goto _321
_321:
	(*(*Manifold)(unsafe.Pointer(bp))).Normal = v320
	i = 0
	for {
		if !(i < (*(*Manifold)(unsafe.Pointer(bp))).PointCount) {
			break
		}
		mp2 = bp + 12 + uintptr(i)*48
		// anchor points relative to shape origin in world space
		v323 = (*ManifoldPoint)(unsafe.Pointer(mp2)).AnchorA
		v324 = origin
		v325 = Vec2{
			X: v323.X + v324.X,
			Y: v323.Y + v324.Y,
		}
		goto _326
	_326:
		v327 = xfA.Q
		v328 = v325
		v329 = Vec2{
			X: float32(v327.C*v328.X) - float32(v327.S*v328.Y),
			Y: float32(v327.S*v328.X) + float32(v327.C*v328.Y),
		}
		goto _330
	_330:
		(*ManifoldPoint)(unsafe.Pointer(mp2)).AnchorA = v329
		v331 = xfA.P
		v332 = xfB.P
		v333 = Vec2{
			X: v331.X - v332.X,
			Y: v331.Y - v332.Y,
		}
		goto _334
	_334:
		v335 = (*ManifoldPoint)(unsafe.Pointer(mp2)).AnchorA
		v336 = v333
		v337 = Vec2{
			X: v335.X + v336.X,
			Y: v335.Y + v336.Y,
		}
		goto _338
	_338:
		(*ManifoldPoint)(unsafe.Pointer(mp2)).AnchorB = v337
		v339 = xfA.P
		v340 = (*ManifoldPoint)(unsafe.Pointer(mp2)).AnchorA
		v341 = Vec2{
			X: v339.X + v340.X,
			Y: v339.Y + v340.Y,
		}
		goto _342
	_342:
		(*ManifoldPoint)(unsafe.Pointer(mp2)).Point = v341
		goto _322
	_322:
		;
		i = i + 1
	}
	return *(*Manifold)(unsafe.Pointer(bp))
}

func b2CollideSegmentAndCapsule(tls *_Stack, segmentA uintptr, xfA Transform, capsuleB uintptr, xfB Transform) (r Manifold) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* capsuleA at bp+0 */ Capsule
	*(*Capsule)(unsafe.Pointer(bp)) = Capsule{
		Center1: (*Segment)(unsafe.Pointer(segmentA)).Point1,
		Center2: (*Segment)(unsafe.Pointer(segmentA)).Point2,
	}
	return b2CollideCapsules(tls, bp, xfA, capsuleB, xfB)
}

func b2CollidePolygonAndCapsule(tls *_Stack, polygonA uintptr, xfA Transform, capsuleB uintptr, xfB Transform) (r Manifold) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var _ /* polyB at bp+0 */ Polygon
	*(*Polygon)(unsafe.Pointer(bp)) = b2MakeCapsule(tls, (*Capsule)(unsafe.Pointer(capsuleB)).Center1, (*Capsule)(unsafe.Pointer(capsuleB)).Center2, (*Capsule)(unsafe.Pointer(capsuleB)).Radius)
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
	var C, sfA, xf, v10, v23, v31, v9 Transform
	var c1, c11, c12, c13, c2, c21, c22, c23, normal, normal1, normal2, normal3, origin, searchDirection, searchDirection1, v11, v12, v21, v22, v101, v102, v104, v106, v107, v109, v110, v112, v114, v115, v116, v118, v120, v1211, v123, v125, v126, v128, v129, v131, v134, v135, v138, v139, v140, v143, v144, v146, v147, v148, v15, v150, v151, v152, v154, v155, v156, v16, v17, v2, v20, v211, v26, v27, v28, v3, v32, v33, v36, v37, v40, v41, v45, v46, v5, v57, v58, v59, v6, v61, v63, v64, v66, v68, v69, v7, v71, v72, v74, v76, v77, v78, v80, v82, v83, v85, v87, v88, v90, v91, v93, v95, v96, v97, v99 Vec2
	var count, count1, i, i1, i11, i12, i2, i21, i22, i3, i4, i5, v49, v50 int32
	var distance, dot, dot1, invDistance, invDistance1, invDistance2, invDistance3, linearSlop, minDot, minDot1, minSeparation, radius, separation, separationA, separationB, speculativeDistance, x, y, v100, v105, v1111, v119, v124, v130, v42, v47, v52, v53, v54, v56, v62, v67, v73, v81, v86, v92 float32
	var flip uint8
	var mp, normals, normals1 uintptr
	var qr, v1, v111, v121, v13, v133, v142, v19, v35 Rot
	var result SegmentDistanceResult
	var _ /* edgeA at bp+288 */ int32
	var _ /* edgeB at bp+292 */ int32
	var _ /* localPolyA at bp+0 */ Polygon
	var _ /* localPolyB at bp+144 */ Polygon
	var _ /* manifold at bp+296 */ Manifold
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, c1, c11, c12, c13, c2, c21, c22, c23, count, count1, distance, dot, dot1, flip, i, i1, i11, i12, i2, i21, i22, i3, i4, i5, invDistance, invDistance1, invDistance2, invDistance3, linearSlop, minDot, minDot1, minSeparation, mp, normal, normal1, normal2, normal3, normals, normals1, origin, qr, radius, result, searchDirection, searchDirection1, separation, separationA, separationB, sfA, speculativeDistance, v11, v12, v21, v22, x, xf, y, v1, v10, v100, v101, v102, v104, v105, v106, v107, v109, v111, v110, v1111, v112, v114, v115, v116, v118, v119, v121, v120, v1211, v123, v124, v125, v126, v128, v129, v13, v130, v131, v133, v134, v135, v138, v139, v140, v142, v143, v144, v146, v147, v148, v15, v150, v151, v152, v154, v155, v156, v16, v17, v19, v2, v20, v211, v23, v26, v27, v28, v3, v31, v32, v33, v35, v36, v37, v40, v41, v42, v45, v46, v47, v49, v5, v50, v52, v53, v54, v56, v57, v58, v59, v6, v61, v62, v63, v64, v66, v67, v68, v69, v7, v71, v72, v73, v74, v76, v77, v78, v80, v81, v82, v83, v85, v86, v87, v88, v9, v90, v91, v92, v93, v95, v96, v97, v99
	origin = *(*Vec2)(unsafe.Pointer(polygonA))
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
	(*(*Polygon)(unsafe.Pointer(bp))).Count = (*Polygon)(unsafe.Pointer(polygonA)).Count
	(*(*Polygon)(unsafe.Pointer(bp))).Radius = (*Polygon)(unsafe.Pointer(polygonA)).Radius
	*(*Vec2)(unsafe.Pointer(bp)) = b2Vec2_zero
	*(*Vec2)(unsafe.Pointer(bp + 64)) = *(*Vec2)(unsafe.Pointer(polygonA + 64))
	i = int32(1)
	for {
		if !(i < (*(*Polygon)(unsafe.Pointer(bp))).Count) {
			break
		}
		v26 = *(*Vec2)(unsafe.Pointer(polygonA + uintptr(i)*8))
		v27 = origin
		v28 = Vec2{
			X: v26.X - v27.X,
			Y: v26.Y - v27.Y,
		}
		goto _29
	_29:
		*(*Vec2)(unsafe.Pointer(bp + uintptr(i)*8)) = v28
		*(*Vec2)(unsafe.Pointer(bp + 64 + uintptr(i)*8)) = *(*Vec2)(unsafe.Pointer(polygonA + 64 + uintptr(i)*8))
		goto _25
	_25:
		;
		i = i + 1
	}
	(*(*Polygon)(unsafe.Pointer(bp + 144))).Count = (*Polygon)(unsafe.Pointer(polygonB)).Count
	(*(*Polygon)(unsafe.Pointer(bp + 144))).Radius = (*Polygon)(unsafe.Pointer(polygonB)).Radius
	i1 = 0
	for {
		if !(i1 < (*(*Polygon)(unsafe.Pointer(bp + 144))).Count) {
			break
		}
		v31 = xf
		v32 = *(*Vec2)(unsafe.Pointer(polygonB + uintptr(i1)*8))
		x = float32(v31.Q.C*v32.X) - float32(v31.Q.S*v32.Y) + v31.P.X
		y = float32(v31.Q.S*v32.X) + float32(v31.Q.C*v32.Y) + v31.P.Y
		v33 = Vec2{
			X: x,
			Y: y,
		}
		goto _34
	_34:
		*(*Vec2)(unsafe.Pointer(bp + 144 + uintptr(i1)*8)) = v33
		v35 = xf.Q
		v36 = *(*Vec2)(unsafe.Pointer(polygonB + 64 + uintptr(i1)*8))
		v37 = Vec2{
			X: float32(v35.C*v36.X) - float32(v35.S*v36.Y),
			Y: float32(v35.S*v36.X) + float32(v35.C*v36.Y),
		}
		goto _38
	_38:
		*(*Vec2)(unsafe.Pointer(bp + 144 + 64 + uintptr(i1)*8)) = v37
		goto _30
	_30:
		;
		i1 = i1 + 1
	}
	*(*int32)(unsafe.Pointer(bp + 288)) = 0
	separationA = b2FindMaxSeparation(tls, bp+288, bp, bp+144)
	*(*int32)(unsafe.Pointer(bp + 292)) = 0
	separationB = b2FindMaxSeparation(tls, bp+292, bp+144, bp)
	radius = (*(*Polygon)(unsafe.Pointer(bp))).Radius + (*(*Polygon)(unsafe.Pointer(bp + 144))).Radius
	if separationA > speculativeDistance+radius || separationB > speculativeDistance+radius {
		return Manifold{}
	}
	if separationA >= separationB {
		flip = boolUint8(false1 != 0)
		searchDirection = *(*Vec2)(unsafe.Pointer(bp + 64 + uintptr(*(*int32)(unsafe.Pointer(bp + 288)))*8))
		// Find the incident edge on polyB
		count = (*(*Polygon)(unsafe.Pointer(bp + 144))).Count
		normals = bp + 144 + 64
		*(*int32)(unsafe.Pointer(bp + 292)) = 0
		minDot = float32FromFloat32(3.4028234663852886e+38)
		i2 = 0
		for {
			if !(i2 < count) {
				break
			}
			v40 = searchDirection
			v41 = *(*Vec2)(unsafe.Pointer(normals + uintptr(i2)*8))
			v42 = float32(v40.X*v41.X) + float32(v40.Y*v41.Y)
			goto _43
		_43:
			dot = v42
			if dot < minDot {
				minDot = dot
				*(*int32)(unsafe.Pointer(bp + 292)) = i2
			}
			goto _39
		_39:
			;
			i2 = i2 + 1
		}
	} else {
		flip = boolUint8(true1 != 0)
		searchDirection1 = *(*Vec2)(unsafe.Pointer(bp + 144 + 64 + uintptr(*(*int32)(unsafe.Pointer(bp + 292)))*8))
		// Find the incident edge on polyA
		count1 = (*(*Polygon)(unsafe.Pointer(bp))).Count
		normals1 = bp + 64
		*(*int32)(unsafe.Pointer(bp + 288)) = 0
		minDot1 = float32FromFloat32(3.4028234663852886e+38)
		i3 = 0
		for {
			if !(i3 < count1) {
				break
			}
			v45 = searchDirection1
			v46 = *(*Vec2)(unsafe.Pointer(normals1 + uintptr(i3)*8))
			v47 = float32(v45.X*v46.X) + float32(v45.Y*v46.Y)
			goto _48
		_48:
			dot1 = v47
			if dot1 < minDot1 {
				minDot1 = dot1
				*(*int32)(unsafe.Pointer(bp + 288)) = i3
			}
			goto _44
		_44:
			;
			i3 = i3 + 1
		}
	}
	*(*Manifold)(unsafe.Pointer(bp + 296)) = Manifold{}
	// Using slop here to ensure vertex-vertex normal vectors can be safely normalized
	// todo this means edge clipping needs to handle slightly non-overlapping edges.
	if separationA > float32(float32FromFloat32(0.1)*linearSlop) || separationB > float32(float32FromFloat32(0.1)*linearSlop) {
		// Edges are disjoint. Find closest points between reference edge and incident edge
		// Reference edge on polygon A
		i11 = *(*int32)(unsafe.Pointer(bp + 288))
		if *(*int32)(unsafe.Pointer(bp + 288))+int32(1) < (*(*Polygon)(unsafe.Pointer(bp))).Count {
			v49 = *(*int32)(unsafe.Pointer(bp + 288)) + int32(1)
		} else {
			v49 = 0
		}
		i12 = v49
		i21 = *(*int32)(unsafe.Pointer(bp + 292))
		if *(*int32)(unsafe.Pointer(bp + 292))+int32(1) < (*(*Polygon)(unsafe.Pointer(bp + 144))).Count {
			v50 = *(*int32)(unsafe.Pointer(bp + 292)) + int32(1)
		} else {
			v50 = 0
		}
		i22 = v50
		v11 = *(*Vec2)(unsafe.Pointer(bp + uintptr(i11)*8))
		v12 = *(*Vec2)(unsafe.Pointer(bp + uintptr(i12)*8))
		v21 = *(*Vec2)(unsafe.Pointer(bp + 144 + uintptr(i21)*8))
		v22 = *(*Vec2)(unsafe.Pointer(bp + 144 + uintptr(i22)*8))
		result = b2SegmentDistance(tls, v11, v12, v21, v22)
		if !(result.DistanceSquared > float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+9991, __ccgo_ts+9936, int32FromInt32(848)) != 0 {
			__builtin_trap(tls)
		}
		distance = sqrtf(tls, result.DistanceSquared)
		separation = distance - radius
		if distance-radius > speculativeDistance {
			// This can happen in the vertex-vertex case
			return *(*Manifold)(unsafe.Pointer(bp + 296))
		}
		// Attempt to clip edges
		*(*Manifold)(unsafe.Pointer(bp + 296)) = b2ClipPolygons(tls, bp, bp+144, *(*int32)(unsafe.Pointer(bp + 288)), *(*int32)(unsafe.Pointer(bp + 292)), flip)
		minSeparation = float32FromFloat32(3.4028234663852886e+38)
		i4 = 0
		for {
			if !(i4 < (*(*Manifold)(unsafe.Pointer(bp + 296))).PointCount) {
				break
			}
			v52 = minSeparation
			v53 = (*(*ManifoldPoint)(unsafe.Pointer(bp + 296 + 12 + uintptr(i4)*48))).Separation
			if v52 < v53 {
				v56 = v52
			} else {
				v56 = v53
			}
			v54 = v56
			goto _55
		_55:
			minSeparation = v54
			goto _51
		_51:
			;
			i4 = i4 + 1
		}
		// Does vertex-vertex have substantially larger separation?
		if separation+float32(float32FromFloat32(0.1)*linearSlop) < minSeparation {
			if result.Fraction1 == float32FromFloat32(0) && result.Fraction2 == float32FromFloat32(0) {
				v57 = v21
				v58 = v11
				v59 = Vec2{
					X: v57.X - v58.X,
					Y: v57.Y - v58.Y,
				}
				goto _60
			_60:
				// v11 - v21
				normal = v59
				invDistance = float32FromFloat32(1) / distance
				normal.X *= invDistance
				normal.Y *= invDistance
				v61 = v11
				v62 = (*(*Polygon)(unsafe.Pointer(bp))).Radius
				v63 = normal
				v64 = Vec2{
					X: v61.X + float32(v62*v63.X),
					Y: v61.Y + float32(v62*v63.Y),
				}
				goto _65
			_65:
				c1 = v64
				v66 = v21
				v67 = -(*(*Polygon)(unsafe.Pointer(bp + 144))).Radius
				v68 = normal
				v69 = Vec2{
					X: v66.X + float32(v67*v68.X),
					Y: v66.Y + float32(v67*v68.Y),
				}
				goto _70
			_70:
				c2 = v69
				(*(*Manifold)(unsafe.Pointer(bp + 296))).Normal = normal
				v71 = c1
				v72 = c2
				v73 = float32FromFloat32(0.5)
				v74 = Vec2{
					X: float32((float32FromFloat32(1)-v73)*v71.X) + float32(v73*v72.X),
					Y: float32((float32FromFloat32(1)-v73)*v71.Y) + float32(v73*v72.Y),
				}
				goto _75
			_75:
				(*(*ManifoldPoint)(unsafe.Pointer(bp + 296 + 12))).AnchorA = v74
				(*(*ManifoldPoint)(unsafe.Pointer(bp + 296 + 12))).Separation = distance - radius
				(*(*ManifoldPoint)(unsafe.Pointer(bp + 296 + 12))).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(i11))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(i21)))
				(*(*Manifold)(unsafe.Pointer(bp + 296))).PointCount = int32(1)
			} else {
				if result.Fraction1 == float32FromFloat32(0) && result.Fraction2 == float32FromFloat32(1) {
					v76 = v22
					v77 = v11
					v78 = Vec2{
						X: v76.X - v77.X,
						Y: v76.Y - v77.Y,
					}
					goto _79
				_79:
					// v11 - v22
					normal1 = v78
					invDistance1 = float32FromFloat32(1) / distance
					normal1.X *= invDistance1
					normal1.Y *= invDistance1
					v80 = v11
					v81 = (*(*Polygon)(unsafe.Pointer(bp))).Radius
					v82 = normal1
					v83 = Vec2{
						X: v80.X + float32(v81*v82.X),
						Y: v80.Y + float32(v81*v82.Y),
					}
					goto _84
				_84:
					c11 = v83
					v85 = v22
					v86 = -(*(*Polygon)(unsafe.Pointer(bp + 144))).Radius
					v87 = normal1
					v88 = Vec2{
						X: v85.X + float32(v86*v87.X),
						Y: v85.Y + float32(v86*v87.Y),
					}
					goto _89
				_89:
					c21 = v88
					(*(*Manifold)(unsafe.Pointer(bp + 296))).Normal = normal1
					v90 = c11
					v91 = c21
					v92 = float32FromFloat32(0.5)
					v93 = Vec2{
						X: float32((float32FromFloat32(1)-v92)*v90.X) + float32(v92*v91.X),
						Y: float32((float32FromFloat32(1)-v92)*v90.Y) + float32(v92*v91.Y),
					}
					goto _94
				_94:
					(*(*ManifoldPoint)(unsafe.Pointer(bp + 296 + 12))).AnchorA = v93
					(*(*ManifoldPoint)(unsafe.Pointer(bp + 296 + 12))).Separation = distance - radius
					(*(*ManifoldPoint)(unsafe.Pointer(bp + 296 + 12))).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(i11))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(i22)))
					(*(*Manifold)(unsafe.Pointer(bp + 296))).PointCount = int32(1)
				} else {
					if result.Fraction1 == float32FromFloat32(1) && result.Fraction2 == float32FromFloat32(0) {
						v95 = v21
						v96 = v12
						v97 = Vec2{
							X: v95.X - v96.X,
							Y: v95.Y - v96.Y,
						}
						goto _98
					_98:
						// v12 - v21
						normal2 = v97
						invDistance2 = float32FromFloat32(1) / distance
						normal2.X *= invDistance2
						normal2.Y *= invDistance2
						v99 = v12
						v100 = (*(*Polygon)(unsafe.Pointer(bp))).Radius
						v101 = normal2
						v102 = Vec2{
							X: v99.X + float32(v100*v101.X),
							Y: v99.Y + float32(v100*v101.Y),
						}
						goto _103
					_103:
						c12 = v102
						v104 = v21
						v105 = -(*(*Polygon)(unsafe.Pointer(bp + 144))).Radius
						v106 = normal2
						v107 = Vec2{
							X: v104.X + float32(v105*v106.X),
							Y: v104.Y + float32(v105*v106.Y),
						}
						goto _108
					_108:
						c22 = v107
						(*(*Manifold)(unsafe.Pointer(bp + 296))).Normal = normal2
						v109 = c12
						v110 = c22
						v1111 = float32FromFloat32(0.5)
						v112 = Vec2{
							X: float32((float32FromFloat32(1)-v1111)*v109.X) + float32(v1111*v110.X),
							Y: float32((float32FromFloat32(1)-v1111)*v109.Y) + float32(v1111*v110.Y),
						}
						goto _113
					_113:
						(*(*ManifoldPoint)(unsafe.Pointer(bp + 296 + 12))).AnchorA = v112
						(*(*ManifoldPoint)(unsafe.Pointer(bp + 296 + 12))).Separation = distance - radius
						(*(*ManifoldPoint)(unsafe.Pointer(bp + 296 + 12))).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(i12))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(i21)))
						(*(*Manifold)(unsafe.Pointer(bp + 296))).PointCount = int32(1)
					} else {
						if result.Fraction1 == float32FromFloat32(1) && result.Fraction2 == float32FromFloat32(1) {
							v114 = v22
							v115 = v12
							v116 = Vec2{
								X: v114.X - v115.X,
								Y: v114.Y - v115.Y,
							}
							goto _117
						_117:
							// v12 - v22
							normal3 = v116
							invDistance3 = float32FromFloat32(1) / distance
							normal3.X *= invDistance3
							normal3.Y *= invDistance3
							v118 = v12
							v119 = (*(*Polygon)(unsafe.Pointer(bp))).Radius
							v120 = normal3
							v1211 = Vec2{
								X: v118.X + float32(v119*v120.X),
								Y: v118.Y + float32(v119*v120.Y),
							}
							goto _122
						_122:
							c13 = v1211
							v123 = v22
							v124 = -(*(*Polygon)(unsafe.Pointer(bp + 144))).Radius
							v125 = normal3
							v126 = Vec2{
								X: v123.X + float32(v124*v125.X),
								Y: v123.Y + float32(v124*v125.Y),
							}
							goto _127
						_127:
							c23 = v126
							(*(*Manifold)(unsafe.Pointer(bp + 296))).Normal = normal3
							v128 = c13
							v129 = c23
							v130 = float32FromFloat32(0.5)
							v131 = Vec2{
								X: float32((float32FromFloat32(1)-v130)*v128.X) + float32(v130*v129.X),
								Y: float32((float32FromFloat32(1)-v130)*v128.Y) + float32(v130*v129.Y),
							}
							goto _132
						_132:
							(*(*ManifoldPoint)(unsafe.Pointer(bp + 296 + 12))).AnchorA = v131
							(*(*ManifoldPoint)(unsafe.Pointer(bp + 296 + 12))).Separation = distance - radius
							(*(*ManifoldPoint)(unsafe.Pointer(bp + 296 + 12))).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(i12))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(i22)))
							(*(*Manifold)(unsafe.Pointer(bp + 296))).PointCount = int32(1)
						}
					}
				}
			}
		}
	} else {
		// Polygons overlap
		*(*Manifold)(unsafe.Pointer(bp + 296)) = b2ClipPolygons(tls, bp, bp+144, *(*int32)(unsafe.Pointer(bp + 288)), *(*int32)(unsafe.Pointer(bp + 292)), flip)
	}
	// Convert manifold to world space
	if (*(*Manifold)(unsafe.Pointer(bp + 296))).PointCount > 0 {
		v133 = xfA.Q
		v134 = (*(*Manifold)(unsafe.Pointer(bp + 296))).Normal
		v135 = Vec2{
			X: float32(v133.C*v134.X) - float32(v133.S*v134.Y),
			Y: float32(v133.S*v134.X) + float32(v133.C*v134.Y),
		}
		goto _136
	_136:
		(*(*Manifold)(unsafe.Pointer(bp + 296))).Normal = v135
		i5 = 0
		for {
			if !(i5 < (*(*Manifold)(unsafe.Pointer(bp + 296))).PointCount) {
				break
			}
			mp = bp + 296 + 12 + uintptr(i5)*48
			// anchor points relative to shape origin in world space
			v138 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
			v139 = origin
			v140 = Vec2{
				X: v138.X + v139.X,
				Y: v138.Y + v139.Y,
			}
			goto _141
		_141:
			v142 = xfA.Q
			v143 = v140
			v144 = Vec2{
				X: float32(v142.C*v143.X) - float32(v142.S*v143.Y),
				Y: float32(v142.S*v143.X) + float32(v142.C*v143.Y),
			}
			goto _145
		_145:
			(*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA = v144
			v146 = xfA.P
			v147 = xfB.P
			v148 = Vec2{
				X: v146.X - v147.X,
				Y: v146.Y - v147.Y,
			}
			goto _149
		_149:
			v150 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
			v151 = v148
			v152 = Vec2{
				X: v150.X + v151.X,
				Y: v150.Y + v151.Y,
			}
			goto _153
		_153:
			(*ManifoldPoint)(unsafe.Pointer(mp)).AnchorB = v152
			v154 = xfA.P
			v155 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
			v156 = Vec2{
				X: v154.X + v155.X,
				Y: v154.Y + v155.Y,
			}
			goto _157
		_157:
			(*ManifoldPoint)(unsafe.Pointer(mp)).Point = v156
			goto _137
		_137:
			;
			i5 = i5 + 1
		}
	}
	return *(*Manifold)(unsafe.Pointer(bp + 296))
}

func b2CollideSegmentAndCircle(tls *_Stack, segmentA uintptr, xfA Transform, circleB uintptr, xfB Transform) (r Manifold) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* capsuleA at bp+0 */ Capsule
	*(*Capsule)(unsafe.Pointer(bp)) = Capsule{
		Center1: (*Segment)(unsafe.Pointer(segmentA)).Point1,
		Center2: (*Segment)(unsafe.Pointer(segmentA)).Point2,
	}
	return b2CollideCapsuleAndCircle(tls, bp, xfA, circleB, xfB)
}

func b2CollideSegmentAndPolygon(tls *_Stack, segmentA uintptr, xfA Transform, polygonB uintptr, xfB Transform) (r Manifold) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var _ /* polygonA at bp+0 */ Polygon
	*(*Polygon)(unsafe.Pointer(bp)) = b2MakeCapsule(tls, (*Segment)(unsafe.Pointer(segmentA)).Point1, (*Segment)(unsafe.Pointer(segmentA)).Point2, float32FromFloat32(0))
	return b2CollidePolygons(tls, bp, xfA, polygonB, xfB)
}

func b2CollideChainSegmentAndCircle(tls *_Stack, segmentA uintptr, xfA Transform, circleB uintptr, xfB Transform) (r Manifold) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var C, xf, v1, v15, v17, v2 Transform
	var cA, cB, contactPointA, e, n, nextEdge, normal, p1, p2, pA, pB, prevEdge, v101, v104, v105, v108, v109, v111, v112, v113, v115, v116, v117, v119, v12, v120, v121, v13, v18, v19, v21, v22, v23, v25, v26, v28, v29, v30, v32, v33, v36, v37, v38, v40, v41, v44, v45, v46, v48, v49, v52, v53, v54, v56, v57, v58, v60, v61, v64, v65, v66, v68, v69, v7, v70, v72, v73, v76, v77, v8, v80, v82, v83, v85, v86, v87, v9, v90, v91, v93, v95, v96, v98, v99 Vec2
	var ee, invLength, offset, radius, separation, u, uPrev, v5, vNext, x, y, v100, v34, v42, v50, v62, v74, v78, v81, v94 float32
	var mp, v89 uintptr
	var qr, v103, v107, v11, v3, v4, v51 Rot
	var _ /* distance at bp+112 */ float32
	var _ /* manifold at bp+0 */ Manifold
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, cA, cB, contactPointA, e, ee, invLength, mp, n, nextEdge, normal, offset, p1, p2, pA, pB, prevEdge, qr, radius, separation, u, uPrev, v5, vNext, x, xf, y, v1, v100, v101, v103, v104, v105, v107, v108, v109, v11, v111, v112, v113, v115, v116, v117, v119, v12, v120, v121, v13, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v28, v29, v3, v30, v32, v33, v34, v36, v37, v38, v4, v40, v41, v42, v44, v45, v46, v48, v49, v51, v50, v52, v53, v54, v56, v57, v58, v60, v61, v62, v64, v65, v66, v68, v69, v7, v70, v72, v73, v74, v76, v77, v78, v8, v80, v81, v82, v83, v85, v86, v87, v89, v9, v90, v91, v93, v94, v95, v96, v98, v99
	*(*Manifold)(unsafe.Pointer(bp)) = Manifold{}
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
		return *(*Manifold)(unsafe.Pointer(bp))
	}
	v36 = p2
	v37 = pB
	v38 = Vec2{
		X: v36.X - v37.X,
		Y: v36.Y - v37.Y,
	}
	goto _39
_39:
	v40 = e
	v41 = v38
	v42 = float32(v40.X*v41.X) + float32(v40.Y*v41.Y)
	goto _43
_43:
	// Barycentric coordinates
	u = v42
	v44 = pB
	v45 = p1
	v46 = Vec2{
		X: v44.X - v45.X,
		Y: v44.Y - v45.Y,
	}
	goto _47
_47:
	v48 = e
	v49 = v46
	v50 = float32(v48.X*v49.X) + float32(v48.Y*v49.Y)
	goto _51
_51:
	v5 = v50
	if v5 <= float32FromFloat32(0) {
		v52 = p1
		v53 = (*ChainSegment)(unsafe.Pointer(segmentA)).Ghost1
		v54 = Vec2{
			X: v52.X - v53.X,
			Y: v52.Y - v53.Y,
		}
		goto _55
	_55:
		// Behind point1?
		// Is pB in the Voronoi region of the previous edge?
		prevEdge = v54
		v56 = pB
		v57 = p1
		v58 = Vec2{
			X: v56.X - v57.X,
			Y: v56.Y - v57.Y,
		}
		goto _59
	_59:
		v60 = prevEdge
		v61 = v58
		v62 = float32(v60.X*v61.X) + float32(v60.Y*v61.Y)
		goto _63
	_63:
		uPrev = v62
		if uPrev <= float32FromFloat32(0) {
			return *(*Manifold)(unsafe.Pointer(bp))
		}
		pA = p1
	} else {
		if u <= float32FromFloat32(0) {
			v64 = (*ChainSegment)(unsafe.Pointer(segmentA)).Ghost2
			v65 = p2
			v66 = Vec2{
				X: v64.X - v65.X,
				Y: v64.Y - v65.Y,
			}
			goto _67
		_67:
			// Ahead of point2?
			nextEdge = v66
			v68 = pB
			v69 = p2
			v70 = Vec2{
				X: v68.X - v69.X,
				Y: v68.Y - v69.Y,
			}
			goto _71
		_71:
			v72 = nextEdge
			v73 = v70
			v74 = float32(v72.X*v73.X) + float32(v72.Y*v73.Y)
			goto _75
		_75:
			vNext = v74
			// Is pB in the Voronoi region of the next edge?
			if vNext > float32FromFloat32(0) {
				return *(*Manifold)(unsafe.Pointer(bp))
			}
			pA = p2
		} else {
			v76 = e
			v77 = e
			v78 = float32(v76.X*v77.X) + float32(v76.Y*v77.Y)
			goto _79
		_79:
			ee = v78
			pA = Vec2{
				X: float32(u*p1.X) + float32(v5*p2.X),
				Y: float32(u*p1.Y) + float32(v5*p2.Y),
			}
			if ee > float32FromFloat32(0) {
				v81 = float32FromFloat32(1) / ee
				v82 = pA
				v83 = Vec2{
					X: float32(v81 * v82.X),
					Y: float32(v81 * v82.Y),
				}
				goto _84
			_84:
				v80 = v83
			} else {
				v80 = p1
			}
			pA = v80
		}
	}
	v85 = pB
	v86 = pA
	v87 = Vec2{
		X: v85.X - v86.X,
		Y: v85.Y - v86.Y,
	}
	goto _88
_88:
	v89 = bp + 112
	v90 = v87
	*(*float32)(unsafe.Pointer(v89)) = sqrtf(tls, float32(v90.X*v90.X)+float32(v90.Y*v90.Y))
	if *(*float32)(unsafe.Pointer(v89)) < float32FromFloat32(1.1920928955078125e-07) {
		v91 = Vec2{}
		goto _92
	}
	invLength = float32FromFloat32(1) / *(*float32)(unsafe.Pointer(v89))
	n = Vec2{
		X: float32(invLength * v90.X),
		Y: float32(invLength * v90.Y),
	}
	v91 = n
	goto _92
_92:
	normal = v91
	radius = (*Circle)(unsafe.Pointer(circleB)).Radius
	separation = *(*float32)(unsafe.Pointer(bp + 112)) - radius
	if separation > float32(float32FromFloat32(4)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
		return *(*Manifold)(unsafe.Pointer(bp))
	}
	cA = pA
	v93 = pB
	v94 = -radius
	v95 = normal
	v96 = Vec2{
		X: v93.X + float32(v94*v95.X),
		Y: v93.Y + float32(v94*v95.Y),
	}
	goto _97
_97:
	cB = v96
	v98 = cA
	v99 = cB
	v100 = float32FromFloat32(0.5)
	v101 = Vec2{
		X: float32((float32FromFloat32(1)-v100)*v98.X) + float32(v100*v99.X),
		Y: float32((float32FromFloat32(1)-v100)*v98.Y) + float32(v100*v99.Y),
	}
	goto _102
_102:
	contactPointA = v101
	v103 = xfA.Q
	v104 = normal
	v105 = Vec2{
		X: float32(v103.C*v104.X) - float32(v103.S*v104.Y),
		Y: float32(v103.S*v104.X) + float32(v103.C*v104.Y),
	}
	goto _106
_106:
	(*(*Manifold)(unsafe.Pointer(bp))).Normal = v105
	mp = bp + 12 + uintptr(0)*48
	v107 = xfA.Q
	v108 = contactPointA
	v109 = Vec2{
		X: float32(v107.C*v108.X) - float32(v107.S*v108.Y),
		Y: float32(v107.S*v108.X) + float32(v107.C*v108.Y),
	}
	goto _110
_110:
	(*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA = v109
	v111 = xfA.P
	v112 = xfB.P
	v113 = Vec2{
		X: v111.X - v112.X,
		Y: v111.Y - v112.Y,
	}
	goto _114
_114:
	v115 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
	v116 = v113
	v117 = Vec2{
		X: v115.X + v116.X,
		Y: v115.Y + v116.Y,
	}
	goto _118
_118:
	(*ManifoldPoint)(unsafe.Pointer(mp)).AnchorB = v117
	v119 = xfA.P
	v120 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
	v121 = Vec2{
		X: v119.X + v120.X,
		Y: v119.Y + v120.Y,
	}
	goto _122
_122:
	(*ManifoldPoint)(unsafe.Pointer(mp)).Point = v121
	(*ManifoldPoint)(unsafe.Pointer(mp)).Separation = separation
	(*ManifoldPoint)(unsafe.Pointer(mp)).Id = uint16(0)
	(*(*Manifold)(unsafe.Pointer(bp))).PointCount = int32(1)
	return *(*Manifold)(unsafe.Pointer(bp))
}

func b2CollideChainSegmentAndCapsule(tls *_Stack, segmentA uintptr, xfA Transform, capsuleB uintptr, xfB Transform, cache uintptr) (r Manifold) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var _ /* polyB at bp+0 */ Polygon
	*(*Polygon)(unsafe.Pointer(bp)) = b2MakeCapsule(tls, (*Capsule)(unsafe.Pointer(capsuleB)).Center1, (*Capsule)(unsafe.Pointer(capsuleB)).Center2, (*Capsule)(unsafe.Pointer(capsuleB)).Radius)
	return b2CollideChainSegmentAndPolygon(tls, segmentA, xfA, bp, xfB, cache)
}

func b2CollideChainSegmentAndPolygon(tls *_Stack, segmentA uintptr, xfA Transform, polygonB uintptr, xfB Transform, cache uintptr) (r Manifold) {
	bp := tls.Alloc(368)
	defer tls.Free(368)
	var C, xf, v1, v15, v17, v2, v84 Transform
	var a11, a21, b11, b12, b21, b22, centroidB, edge0, edge1, edge2, n, n0, n1, n2, n21, normal, normal1, normalB, p1, p11, p2, pA, pAB, pAB1, pAB2, pB, v102, v103, v106, v107, v109, v110, v111, v113, v114, v115, v117, v118, v119, v12, v121, v122, v123, v125, v126, v129, v13, v130, v134, v135, v138, v139, v140, v142, v143, v146, v147, v148, v150, v151, v154, v155, v158, v159, v162, v163, v166, v167, v170, v171, v174, v175, v178, v179, v18, v182, v183, v185, v186, v187, v189, v19, v190, v191, v193, v194, v195, v197, v198, v199, v201, v202, v203, v205, v206, v207, v209, v21, v210, v213, v214, v215, v217, v218, v22, v223, v224, v225, v227, v228, v23, v232, v233, v234, v236, v237, v241, v242, v243, v245, v246, v25, v250, v251, v253, v254, v255, v257, v258, v26, v261, v262, v263, v265, v266, v275, v276, v277, v279, v28, v280, v283, v284, v285, v287, v288, v29, v291, v292, v295, v296, v299, v30, v300, v303, v304, v307, v308, v311, v312, v315, v316, v319, v32, v320, v322, v323, v324, v326, v327, v328, v33, v330, v331, v332, v334, v335, v336, v338, v339, v340, v344, v345, v348, v349, v35, v354, v355, v358, v359, v36, v362, v363, v365, v366, v367, v369, v370, v371, v373, v374, v375, v377, v378, v379, v38, v381, v382, v383, v39, v42, v43, v44, v46, v47, v49, v50, v52, v53, v56, v57, v59, v60, v61, v63, v64, v67, v68, v69, v7, v71, v72, v75, v76, v77, v79, v8, v80, v85, v86, v89, v9, v90, v92, v93, v94, v95, v96, v98, v99 Vec2
	var behind0, behind1, behind2 uint8
	var convexTol, d1, d2, dot1, dot11, dot12, dot2, dot21, dot22, edgeSeparation, invLength, length, polygonSeparation, radiusB, s, s0, s1, s2, s21, s3, x, y, v127, v131, v144, v152, v156, v160, v164, v168, v211, v219, v229, v238, v247, v259, v267, v269, v270, v271, v273, v281, v289, v293, v297, v301, v305, v346, v350, v40, v54, v65, v73, v81 float32
	var count, i, i1, i11, i2, i21, i3, i4, ia1, ia11, ia2, ia21, ib, ib1, ib11, ib2, ib21, incidentIndex, incidentNormal, referenceIndex, v133, v137, v221, v274, v342, v343, v352 int32
	var cp uintptr
	var normals [8]Vec2
	var output DistanceOutput
	var qr, v101, v105, v11, v173, v177, v181, v3, v310, v314, v318, v353, v357, v361, v4, v5, v88 Rot
	var smoothParams b2ChainSegmentParams
	var type1, type11, type2 b2NormalType
	var _ /* input at bp+176 */ DistanceInput
	var _ /* manifold at bp+0 */ Manifold
	var _ /* vertices at bp+112 */ [8]Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, a11, a21, b11, b12, b21, b22, behind0, behind1, behind2, centroidB, convexTol, count, cp, d1, d2, dot1, dot11, dot12, dot2, dot21, dot22, edge0, edge1, edge2, edgeSeparation, i, i1, i11, i2, i21, i3, i4, ia1, ia11, ia2, ia21, ib, ib1, ib11, ib2, ib21, incidentIndex, incidentNormal, invLength, length, n, n0, n1, n2, n21, normal, normal1, normalB, normals, output, p1, p11, p2, pA, pAB, pAB1, pAB2, pB, polygonSeparation, qr, radiusB, referenceIndex, s, s0, s1, s2, s21, s3, smoothParams, type1, type11, type2, x, xf, y, v1, v101, v102, v103, v105, v106, v107, v109, v11, v110, v111, v113, v114, v115, v117, v118, v119, v12, v121, v122, v123, v125, v126, v127, v129, v13, v130, v131, v133, v134, v135, v137, v138, v139, v140, v142, v143, v144, v146, v147, v148, v15, v150, v151, v152, v154, v155, v156, v158, v159, v160, v162, v163, v164, v166, v167, v168, v17, v170, v171, v173, v174, v175, v177, v178, v179, v18, v181, v182, v183, v185, v186, v187, v189, v19, v190, v191, v193, v194, v195, v197, v198, v199, v2, v201, v202, v203, v205, v206, v207, v209, v21, v210, v211, v213, v214, v215, v217, v218, v219, v22, v221, v223, v224, v225, v227, v228, v229, v23, v232, v233, v234, v236, v237, v238, v241, v242, v243, v245, v246, v247, v25, v250, v251, v253, v254, v255, v257, v258, v259, v26, v261, v262, v263, v265, v266, v267, v269, v270, v271, v273, v274, v275, v276, v277, v279, v28, v280, v281, v283, v284, v285, v287, v288, v289, v29, v291, v292, v293, v295, v296, v297, v299, v3, v30, v300, v301, v303, v304, v305, v307, v308, v310, v311, v312, v314, v315, v316, v318, v319, v32, v320, v322, v323, v324, v326, v327, v328, v33, v330, v331, v332, v334, v335, v336, v338, v339, v340, v342, v343, v344, v345, v346, v348, v349, v35, v350, v352, v353, v354, v355, v357, v358, v359, v36, v361, v362, v363, v365, v366, v367, v369, v370, v371, v373, v374, v375, v377, v378, v379, v38, v381, v382, v383, v39, v4, v40, v42, v43, v44, v46, v47, v49, v5, v50, v52, v53, v54, v56, v57, v59, v60, v61, v63, v64, v65, v67, v68, v69, v7, v71, v72, v73, v75, v76, v77, v79, v8, v80, v81, v84, v85, v86, v88, v89, v9, v90, v92, v93, v94, v95, v96, v98, v99
	*(*Manifold)(unsafe.Pointer(bp)) = Manifold{}
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
	v28 = p11
	v29 = (*ChainSegment)(unsafe.Pointer(segmentA)).Ghost1
	v30 = Vec2{
		X: v28.X - v29.X,
		Y: v28.Y - v29.Y,
	}
	goto _31
_31:
	v32 = v30
	length = sqrtf(tls, float32(v32.X*v32.X)+float32(v32.Y*v32.Y))
	if length < float32FromFloat32(1.1920928955078125e-07) {
		v33 = Vec2{}
		goto _34
	}
	invLength = float32FromFloat32(1) / length
	n = Vec2{
		X: float32(invLength * v32.X),
		Y: float32(invLength * v32.Y),
	}
	v33 = n
	goto _34
_34:
	edge0 = v33
	v35 = edge0
	v36 = Vec2{
		X: v35.Y,
		Y: -v35.X,
	}
	goto _37
_37:
	smoothParams.Normal0 = v36
	v38 = edge0
	v39 = edge1
	v40 = float32(v38.X*v39.Y) - float32(v38.Y*v39.X)
	goto _41
_41:
	smoothParams.Convex1 = boolUint8(v40 >= convexTol)
	v42 = (*ChainSegment)(unsafe.Pointer(segmentA)).Ghost2
	v43 = p2
	v44 = Vec2{
		X: v42.X - v43.X,
		Y: v42.Y - v43.Y,
	}
	goto _45
_45:
	v46 = v44
	length = sqrtf(tls, float32(v46.X*v46.X)+float32(v46.Y*v46.Y))
	if length < float32FromFloat32(1.1920928955078125e-07) {
		v47 = Vec2{}
		goto _48
	}
	invLength = float32FromFloat32(1) / length
	n = Vec2{
		X: float32(invLength * v46.X),
		Y: float32(invLength * v46.Y),
	}
	v47 = n
	goto _48
_48:
	edge2 = v47
	v49 = edge2
	v50 = Vec2{
		X: v49.Y,
		Y: -v49.X,
	}
	goto _51
_51:
	smoothParams.Normal2 = v50
	v52 = edge1
	v53 = edge2
	v54 = float32(v52.X*v53.Y) - float32(v52.Y*v53.X)
	goto _55
_55:
	smoothParams.Convex2 = boolUint8(v54 >= convexTol)
	v56 = edge1
	v57 = Vec2{
		X: v56.Y,
		Y: -v56.X,
	}
	goto _58
_58:
	// Normal points to the right
	normal1 = v57
	v59 = centroidB
	v60 = p11
	v61 = Vec2{
		X: v59.X - v60.X,
		Y: v59.Y - v60.Y,
	}
	goto _62
_62:
	v63 = normal1
	v64 = v61
	v65 = float32(v63.X*v64.X) + float32(v63.Y*v64.Y)
	goto _66
_66:
	behind1 = boolUint8(v65 < float32FromFloat32(0))
	behind0 = boolUint8(true1 != 0)
	behind2 = boolUint8(true1 != 0)
	if smoothParams.Convex1 != 0 {
		v67 = centroidB
		v68 = p11
		v69 = Vec2{
			X: v67.X - v68.X,
			Y: v67.Y - v68.Y,
		}
		goto _70
	_70:
		v71 = smoothParams.Normal0
		v72 = v69
		v73 = float32(v71.X*v72.X) + float32(v71.Y*v72.Y)
		goto _74
	_74:
		behind0 = boolUint8(v73 < float32FromFloat32(0))
	}
	if smoothParams.Convex2 != 0 {
		v75 = centroidB
		v76 = p2
		v77 = Vec2{
			X: v75.X - v76.X,
			Y: v75.Y - v76.Y,
		}
		goto _78
	_78:
		v79 = smoothParams.Normal2
		v80 = v77
		v81 = float32(v79.X*v80.X) + float32(v79.Y*v80.Y)
		goto _82
	_82:
		behind2 = boolUint8(v81 < float32FromFloat32(0))
	}
	if behind1 != 0 && behind0 != 0 && behind2 != 0 {
		// one-sided collision
		return *(*Manifold)(unsafe.Pointer(bp))
	}
	// Get polygonB in frameA
	count = (*Polygon)(unsafe.Pointer(polygonB)).Count
	i = 0
	for {
		if !(i < count) {
			break
		}
		v84 = xf
		v85 = *(*Vec2)(unsafe.Pointer(polygonB + uintptr(i)*8))
		x = float32(v84.Q.C*v85.X) - float32(v84.Q.S*v85.Y) + v84.P.X
		y = float32(v84.Q.S*v85.X) + float32(v84.Q.C*v85.Y) + v84.P.Y
		v86 = Vec2{
			X: x,
			Y: y,
		}
		goto _87
	_87:
		(*(*[8]Vec2)(unsafe.Pointer(bp + 112)))[i] = v86
		v88 = xf.Q
		v89 = *(*Vec2)(unsafe.Pointer(polygonB + 64 + uintptr(i)*8))
		v90 = Vec2{
			X: float32(v88.C*v89.X) - float32(v88.S*v89.Y),
			Y: float32(v88.S*v89.X) + float32(v88.C*v89.Y),
		}
		goto _91
	_91:
		normals[i] = v90
		goto _83
	_83:
		;
		i = i + 1
	}
	(*(*DistanceInput)(unsafe.Pointer(bp + 176))).ProxyA = b2MakeProxy(tls, segmentA+8, int32(2), float32FromFloat32(0))
	(*(*DistanceInput)(unsafe.Pointer(bp + 176))).ProxyB = b2MakeProxy(tls, bp+112, count, float32FromFloat32(0))
	(*(*DistanceInput)(unsafe.Pointer(bp + 176))).TransformA = b2Transform_identity
	(*(*DistanceInput)(unsafe.Pointer(bp + 176))).TransformB = b2Transform_identity
	(*(*DistanceInput)(unsafe.Pointer(bp + 176))).UseRadii = boolUint8(false1 != 0)
	output = b2ShapeDistance(tls, bp+176, cache, uintptrFromInt32(0), 0)
	if output.Distance > radiusB+float32(float32FromFloat32(4)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
		return *(*Manifold)(unsafe.Pointer(bp))
	}
	if smoothParams.Convex1 != 0 {
		v92 = smoothParams.Normal0
	} else {
		v92 = normal1
	}
	// Snap concave normals for partial polygon
	n0 = v92
	if smoothParams.Convex2 != 0 {
		v93 = smoothParams.Normal2
	} else {
		v93 = normal1
	}
	n21 = v93
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
			v94 = pB
			v95 = pA
			v96 = Vec2{
				X: v94.X - v95.X,
				Y: v94.Y - v95.Y,
			}
			goto _97
		_97:
			v98 = v96
			length = sqrtf(tls, float32(v98.X*v98.X)+float32(v98.Y*v98.Y))
			if length < float32FromFloat32(1.1920928955078125e-07) {
				v99 = Vec2{}
				goto _100
			}
			invLength = float32FromFloat32(1) / length
			n = Vec2{
				X: float32(invLength * v98.X),
				Y: float32(invLength * v98.Y),
			}
			v99 = n
			goto _100
		_100:
			normal = v99
			type1 = b2ClassifyNormal(tls, smoothParams, normal)
			if type1 == int32(b2_normalSkip) {
				return *(*Manifold)(unsafe.Pointer(bp))
			}
			if type1 == int32(b2_normalAdmit) {
				v101 = xfA.Q
				v102 = normal
				v103 = Vec2{
					X: float32(v101.C*v102.X) - float32(v101.S*v102.Y),
					Y: float32(v101.S*v102.X) + float32(v101.C*v102.Y),
				}
				goto _104
			_104:
				(*(*Manifold)(unsafe.Pointer(bp))).Normal = v103
				cp = bp + 12 + uintptr(0)*48
				v105 = xfA.Q
				v106 = pA
				v107 = Vec2{
					X: float32(v105.C*v106.X) - float32(v105.S*v106.Y),
					Y: float32(v105.S*v106.X) + float32(v105.C*v106.Y),
				}
				goto _108
			_108:
				(*ManifoldPoint)(unsafe.Pointer(cp)).AnchorA = v107
				v109 = xfA.P
				v110 = xfB.P
				v111 = Vec2{
					X: v109.X - v110.X,
					Y: v109.Y - v110.Y,
				}
				goto _112
			_112:
				v113 = (*ManifoldPoint)(unsafe.Pointer(cp)).AnchorA
				v114 = v111
				v115 = Vec2{
					X: v113.X + v114.X,
					Y: v113.Y + v114.Y,
				}
				goto _116
			_116:
				(*ManifoldPoint)(unsafe.Pointer(cp)).AnchorB = v115
				v117 = xfA.P
				v118 = (*ManifoldPoint)(unsafe.Pointer(cp)).AnchorA
				v119 = Vec2{
					X: v117.X + v118.X,
					Y: v117.Y + v118.Y,
				}
				goto _120
			_120:
				(*ManifoldPoint)(unsafe.Pointer(cp)).Point = v119
				(*ManifoldPoint)(unsafe.Pointer(cp)).Separation = output.Distance - radiusB
				(*ManifoldPoint)(unsafe.Pointer(cp)).Id = uint16FromInt32(int32FromUint8(*(*uint8_t)(unsafe.Pointer(cache + 2)))<<int32FromInt32(8) | int32FromUint8(*(*uint8_t)(unsafe.Pointer(cache + 5))))
				(*(*Manifold)(unsafe.Pointer(bp))).PointCount = int32(1)
				return *(*Manifold)(unsafe.Pointer(bp))
			}
			// fall through b2_normalSnap
			incidentIndex = int32FromUint8(*(*uint8_t)(unsafe.Pointer(cache + 5)))
		} else {
			// vertex-edge collision
			if !(int32FromUint16((*SimplexCache)(unsafe.Pointer(cache)).Count) == int32FromInt32(2)) && b2InternalAssertFcn(tls, __ccgo_ts+10021, __ccgo_ts+9936, int32FromInt32(1438)) != 0 {
				__builtin_trap(tls)
			}
			ia1 = int32FromUint8(*(*uint8_t)(unsafe.Pointer(cache + 2)))
			ia2 = int32FromUint8(*(*uint8_t)(unsafe.Pointer(cache + 2 + 1)))
			ib1 = int32FromUint8(*(*uint8_t)(unsafe.Pointer(cache + 5)))
			ib2 = int32FromUint8(*(*uint8_t)(unsafe.Pointer(cache + 5 + 1)))
			if ia1 == ia2 {
				// 1 point on A, expect 2 points on B
				if !(ib1 != ib2) && b2InternalAssertFcn(tls, __ccgo_ts+10039, __ccgo_ts+9936, int32FromInt32(1448)) != 0 {
					__builtin_trap(tls)
				}
				v121 = output.PointA
				v122 = output.PointB
				v123 = Vec2{
					X: v121.X - v122.X,
					Y: v121.Y - v122.Y,
				}
				goto _124
			_124:
				// Find polygon normal most aligned with vector between closest points.
				// This effectively sorts ib1 and ib2
				normalB = v123
				v125 = normalB
				v126 = normals[ib1]
				v127 = float32(v125.X*v126.X) + float32(v125.Y*v126.Y)
				goto _128
			_128:
				dot1 = v127
				v129 = normalB
				v130 = normals[ib2]
				v131 = float32(v129.X*v130.X) + float32(v129.Y*v130.Y)
				goto _132
			_132:
				dot2 = v131
				if dot1 > dot2 {
					v133 = ib1
				} else {
					v133 = ib2
				}
				ib = v133
				// Use accurate normal
				normalB = normals[ib]
				v134 = normalB
				v135 = Vec2{
					X: -v134.X,
					Y: -v134.Y,
				}
				goto _136
			_136:
				type11 = b2ClassifyNormal(tls, smoothParams, v135)
				if type11 == int32(b2_normalSkip) {
					return *(*Manifold)(unsafe.Pointer(bp))
				}
				if type11 == int32(b2_normalAdmit) {
					// Get polygon edge associated with normal
					ib1 = ib
					if ib < count-int32(1) {
						v137 = ib + int32(1)
					} else {
						v137 = 0
					}
					ib2 = v137
					b11 = (*(*[8]Vec2)(unsafe.Pointer(bp + 112)))[ib1]
					b21 = (*(*[8]Vec2)(unsafe.Pointer(bp + 112)))[ib2]
					// Find incident segment vertex
					v138 = p11
					v139 = b11
					v140 = Vec2{
						X: v138.X - v139.X,
						Y: v138.Y - v139.Y,
					}
					goto _141
				_141:
					v142 = normalB
					v143 = v140
					v144 = float32(v142.X*v143.X) + float32(v142.Y*v143.Y)
					goto _145
				_145:
					dot1 = v144
					v146 = p2
					v147 = b11
					v148 = Vec2{
						X: v146.X - v147.X,
						Y: v146.Y - v147.Y,
					}
					goto _149
				_149:
					v150 = normalB
					v151 = v148
					v152 = float32(v150.X*v151.X) + float32(v150.Y*v151.Y)
					goto _153
				_153:
					dot2 = v152
					if dot1 < dot2 {
						v154 = n0
						v155 = normalB
						v156 = float32(v154.X*v155.X) + float32(v154.Y*v155.Y)
						goto _157
					_157:
						v158 = normal1
						v159 = normalB
						v160 = float32(v158.X*v159.X) + float32(v158.Y*v159.Y)
						goto _161
					_161:
						if v156 < v160 {
							// Neighbor is incident
							return *(*Manifold)(unsafe.Pointer(bp))
						}
					} else {
						v162 = n21
						v163 = normalB
						v164 = float32(v162.X*v163.X) + float32(v162.Y*v163.Y)
						goto _165
					_165:
						v166 = normal1
						v167 = normalB
						v168 = float32(v166.X*v167.X) + float32(v166.Y*v167.Y)
						goto _169
					_169:
						if v164 < v168 {
							// Neighbor is incident
							return *(*Manifold)(unsafe.Pointer(bp))
						}
					}
					*(*Manifold)(unsafe.Pointer(bp)) = b2ClipSegments(tls, b11, b21, p11, p2, normalB, radiusB, float32FromFloat32(0), uint16FromInt32(int32FromUint8(uint8FromInt32(ib1))<<int32FromInt32(8)|int32FromUint8(uint8FromInt32(int32FromInt32(1)))), uint16FromInt32(int32FromUint8(uint8FromInt32(ib2))<<int32FromInt32(8)|int32FromUint8(uint8FromInt32(int32FromInt32(0)))))
					if !((*(*Manifold)(unsafe.Pointer(bp))).PointCount == 0 || (*(*Manifold)(unsafe.Pointer(bp))).PointCount == int32(2)) && b2InternalAssertFcn(tls, __ccgo_ts+10050, __ccgo_ts+9936, int32FromInt32(1499)) != 0 {
						__builtin_trap(tls)
					}
					if (*(*Manifold)(unsafe.Pointer(bp))).PointCount == int32(2) {
						v170 = normalB
						v171 = Vec2{
							X: -v170.X,
							Y: -v170.Y,
						}
						goto _172
					_172:
						v173 = xfA.Q
						v174 = v171
						v175 = Vec2{
							X: float32(v173.C*v174.X) - float32(v173.S*v174.Y),
							Y: float32(v173.S*v174.X) + float32(v173.C*v174.Y),
						}
						goto _176
					_176:
						(*(*Manifold)(unsafe.Pointer(bp))).Normal = v175
						v177 = xfA.Q
						v178 = (*(*ManifoldPoint)(unsafe.Pointer(bp + 12))).AnchorA
						v179 = Vec2{
							X: float32(v177.C*v178.X) - float32(v177.S*v178.Y),
							Y: float32(v177.S*v178.X) + float32(v177.C*v178.Y),
						}
						goto _180
					_180:
						(*(*ManifoldPoint)(unsafe.Pointer(bp + 12))).AnchorA = v179
						v181 = xfA.Q
						v182 = (*(*ManifoldPoint)(unsafe.Pointer(bp + 12 + 1*48))).AnchorA
						v183 = Vec2{
							X: float32(v181.C*v182.X) - float32(v181.S*v182.Y),
							Y: float32(v181.S*v182.X) + float32(v181.C*v182.Y),
						}
						goto _184
					_184:
						(*(*ManifoldPoint)(unsafe.Pointer(bp + 12 + 1*48))).AnchorA = v183
						v185 = xfA.P
						v186 = xfB.P
						v187 = Vec2{
							X: v185.X - v186.X,
							Y: v185.Y - v186.Y,
						}
						goto _188
					_188:
						pAB = v187
						v189 = (*(*ManifoldPoint)(unsafe.Pointer(bp + 12))).AnchorA
						v190 = pAB
						v191 = Vec2{
							X: v189.X + v190.X,
							Y: v189.Y + v190.Y,
						}
						goto _192
					_192:
						(*(*ManifoldPoint)(unsafe.Pointer(bp + 12))).AnchorB = v191
						v193 = (*(*ManifoldPoint)(unsafe.Pointer(bp + 12 + 1*48))).AnchorA
						v194 = pAB
						v195 = Vec2{
							X: v193.X + v194.X,
							Y: v193.Y + v194.Y,
						}
						goto _196
					_196:
						(*(*ManifoldPoint)(unsafe.Pointer(bp + 12 + 1*48))).AnchorB = v195
						v197 = xfA.P
						v198 = (*(*ManifoldPoint)(unsafe.Pointer(bp + 12))).AnchorA
						v199 = Vec2{
							X: v197.X + v198.X,
							Y: v197.Y + v198.Y,
						}
						goto _200
					_200:
						(*(*ManifoldPoint)(unsafe.Pointer(bp + 12))).Point = v199
						v201 = xfA.P
						v202 = (*(*ManifoldPoint)(unsafe.Pointer(bp + 12 + 1*48))).AnchorA
						v203 = Vec2{
							X: v201.X + v202.X,
							Y: v201.Y + v202.Y,
						}
						goto _204
					_204:
						(*(*ManifoldPoint)(unsafe.Pointer(bp + 12 + 1*48))).Point = v203
					}
					return *(*Manifold)(unsafe.Pointer(bp))
				}
				// fall through b2_normalSnap
				incidentNormal = ib
			} else {
				v205 = (*(*[8]Vec2)(unsafe.Pointer(bp + 112)))[ib1]
				v206 = p11
				v207 = Vec2{
					X: v205.X - v206.X,
					Y: v205.Y - v206.Y,
				}
				goto _208
			_208:
				v209 = normal1
				v210 = v207
				v211 = float32(v209.X*v210.X) + float32(v209.Y*v210.Y)
				goto _212
			_212:
				// Get index of incident polygonB vertex
				dot11 = v211
				v213 = (*(*[8]Vec2)(unsafe.Pointer(bp + 112)))[ib2]
				v214 = p2
				v215 = Vec2{
					X: v213.X - v214.X,
					Y: v213.Y - v214.Y,
				}
				goto _216
			_216:
				v217 = normal1
				v218 = v215
				v219 = float32(v217.X*v218.X) + float32(v217.Y*v218.Y)
				goto _220
			_220:
				dot21 = v219
				if dot11 < dot21 {
					v221 = ib1
				} else {
					v221 = ib2
				}
				incidentIndex = v221
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
			v223 = (*(*[8]Vec2)(unsafe.Pointer(bp + 112)))[i1]
			v224 = p11
			v225 = Vec2{
				X: v223.X - v224.X,
				Y: v223.Y - v224.Y,
			}
			goto _226
		_226:
			v227 = normal1
			v228 = v225
			v229 = float32(v227.X*v228.X) + float32(v227.Y*v228.Y)
			goto _230
		_230:
			s = v229
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
				v232 = (*(*[8]Vec2)(unsafe.Pointer(bp + 112)))[i2]
				v233 = p11
				v234 = Vec2{
					X: v232.X - v233.X,
					Y: v232.Y - v233.Y,
				}
				goto _235
			_235:
				v236 = smoothParams.Normal0
				v237 = v234
				v238 = float32(v236.X*v237.X) + float32(v236.Y*v237.Y)
				goto _239
			_239:
				s1 = v238
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
				v241 = (*(*[8]Vec2)(unsafe.Pointer(bp + 112)))[i3]
				v242 = p2
				v243 = Vec2{
					X: v241.X - v242.X,
					Y: v241.Y - v242.Y,
				}
				goto _244
			_244:
				v245 = smoothParams.Normal2
				v246 = v243
				v247 = float32(v245.X*v246.X) + float32(v245.Y*v246.Y)
				goto _248
			_248:
				s2 = v247
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
			v250 = n1
			v251 = Vec2{
				X: -v250.X,
				Y: -v250.Y,
			}
			goto _252
		_252:
			type2 = b2ClassifyNormal(tls, smoothParams, v251)
			if type2 != int32(b2_normalAdmit) {
				goto _249
			}
			// Check the infinite sides of the partial polygon
			// if ((smoothParams.convex1 && b2Cross(n0, n) > 0.0f) || (smoothParams.convex2 && b2Cross(n, n2) > 0.0f))
			//{
			//	continue;
			//}
			p1 = (*(*[8]Vec2)(unsafe.Pointer(bp + 112)))[i4]
			v253 = p2
			v254 = p1
			v255 = Vec2{
				X: v253.X - v254.X,
				Y: v253.Y - v254.Y,
			}
			goto _256
		_256:
			v257 = n1
			v258 = v255
			v259 = float32(v257.X*v258.X) + float32(v257.Y*v258.Y)
			goto _260
		_260:
			v261 = p11
			v262 = p1
			v263 = Vec2{
				X: v261.X - v262.X,
				Y: v261.Y - v262.Y,
			}
			goto _264
		_264:
			v265 = n1
			v266 = v263
			v267 = float32(v265.X*v266.X) + float32(v265.Y*v266.Y)
			goto _268
		_268:
			v269 = v259
			v270 = v267
			if v269 < v270 {
				v273 = v269
			} else {
				v273 = v270
			}
			v271 = v273
			goto _272
		_272:
			s3 = v271
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
				v274 = ia11 + int32(1)
			} else {
				v274 = 0
			}
			ia21 = v274
			a11 = (*(*[8]Vec2)(unsafe.Pointer(bp + 112)))[ia11]
			a21 = (*(*[8]Vec2)(unsafe.Pointer(bp + 112)))[ia21]
			n2 = normals[ia11]
			v275 = p11
			v276 = a11
			v277 = Vec2{
				X: v275.X - v276.X,
				Y: v275.Y - v276.Y,
			}
			goto _278
		_278:
			v279 = n2
			v280 = v277
			v281 = float32(v279.X*v280.X) + float32(v279.Y*v280.Y)
			goto _282
		_282:
			dot12 = v281
			v283 = p2
			v284 = a11
			v285 = Vec2{
				X: v283.X - v284.X,
				Y: v283.Y - v284.Y,
			}
			goto _286
		_286:
			v287 = n2
			v288 = v285
			v289 = float32(v287.X*v288.X) + float32(v287.Y*v288.Y)
			goto _290
		_290:
			dot22 = v289
			if dot12 < dot22 {
				v291 = n0
				v292 = n2
				v293 = float32(v291.X*v292.X) + float32(v291.Y*v292.Y)
				goto _294
			_294:
				v295 = normal1
				v296 = n2
				v297 = float32(v295.X*v296.X) + float32(v295.Y*v296.Y)
				goto _298
			_298:
				if v293 < v297 {
					// Neighbor is incident
					return *(*Manifold)(unsafe.Pointer(bp))
				}
			} else {
				v299 = n21
				v300 = n2
				v301 = float32(v299.X*v300.X) + float32(v299.Y*v300.Y)
				goto _302
			_302:
				v303 = normal1
				v304 = n2
				v305 = float32(v303.X*v304.X) + float32(v303.Y*v304.Y)
				goto _306
			_306:
				if v301 < v305 {
					// Neighbor is incident
					return *(*Manifold)(unsafe.Pointer(bp))
				}
			}
			*(*Manifold)(unsafe.Pointer(bp)) = b2ClipSegments(tls, a11, a21, p11, p2, normals[ia11], radiusB, float32FromFloat32(0), uint16FromInt32(int32FromUint8(uint8FromInt32(ia11))<<int32FromInt32(8)|int32FromUint8(uint8FromInt32(int32FromInt32(1)))), uint16FromInt32(int32FromUint8(uint8FromInt32(ia21))<<int32FromInt32(8)|int32FromUint8(uint8FromInt32(int32FromInt32(0)))))
			if !((*(*Manifold)(unsafe.Pointer(bp))).PointCount == 0 || (*(*Manifold)(unsafe.Pointer(bp))).PointCount == int32(2)) && b2InternalAssertFcn(tls, __ccgo_ts+10050, __ccgo_ts+9936, int32FromInt32(1648)) != 0 {
				__builtin_trap(tls)
			}
			if (*(*Manifold)(unsafe.Pointer(bp))).PointCount == int32(2) {
				v307 = normals[ia11]
				v308 = Vec2{
					X: -v307.X,
					Y: -v307.Y,
				}
				goto _309
			_309:
				v310 = xfA.Q
				v311 = v308
				v312 = Vec2{
					X: float32(v310.C*v311.X) - float32(v310.S*v311.Y),
					Y: float32(v310.S*v311.X) + float32(v310.C*v311.Y),
				}
				goto _313
			_313:
				(*(*Manifold)(unsafe.Pointer(bp))).Normal = v312
				v314 = xfA.Q
				v315 = (*(*ManifoldPoint)(unsafe.Pointer(bp + 12))).AnchorA
				v316 = Vec2{
					X: float32(v314.C*v315.X) - float32(v314.S*v315.Y),
					Y: float32(v314.S*v315.X) + float32(v314.C*v315.Y),
				}
				goto _317
			_317:
				(*(*ManifoldPoint)(unsafe.Pointer(bp + 12))).AnchorA = v316
				v318 = xfA.Q
				v319 = (*(*ManifoldPoint)(unsafe.Pointer(bp + 12 + 1*48))).AnchorA
				v320 = Vec2{
					X: float32(v318.C*v319.X) - float32(v318.S*v319.Y),
					Y: float32(v318.S*v319.X) + float32(v318.C*v319.Y),
				}
				goto _321
			_321:
				(*(*ManifoldPoint)(unsafe.Pointer(bp + 12 + 1*48))).AnchorA = v320
				v322 = xfA.P
				v323 = xfB.P
				v324 = Vec2{
					X: v322.X - v323.X,
					Y: v322.Y - v323.Y,
				}
				goto _325
			_325:
				pAB1 = v324
				v326 = (*(*ManifoldPoint)(unsafe.Pointer(bp + 12))).AnchorA
				v327 = pAB1
				v328 = Vec2{
					X: v326.X + v327.X,
					Y: v326.Y + v327.Y,
				}
				goto _329
			_329:
				(*(*ManifoldPoint)(unsafe.Pointer(bp + 12))).AnchorB = v328
				v330 = (*(*ManifoldPoint)(unsafe.Pointer(bp + 12 + 1*48))).AnchorA
				v331 = pAB1
				v332 = Vec2{
					X: v330.X + v331.X,
					Y: v330.Y + v331.Y,
				}
				goto _333
			_333:
				(*(*ManifoldPoint)(unsafe.Pointer(bp + 12 + 1*48))).AnchorB = v332
				v334 = xfA.P
				v335 = (*(*ManifoldPoint)(unsafe.Pointer(bp + 12))).AnchorA
				v336 = Vec2{
					X: v334.X + v335.X,
					Y: v334.Y + v335.Y,
				}
				goto _337
			_337:
				(*(*ManifoldPoint)(unsafe.Pointer(bp + 12))).Point = v336
				v338 = xfA.P
				v339 = (*(*ManifoldPoint)(unsafe.Pointer(bp + 12 + 1*48))).AnchorA
				v340 = Vec2{
					X: v338.X + v339.X,
					Y: v338.Y + v339.Y,
				}
				goto _341
			_341:
				(*(*ManifoldPoint)(unsafe.Pointer(bp + 12 + 1*48))).Point = v340
			}
			return *(*Manifold)(unsafe.Pointer(bp))
		}
		if incidentIndex == -int32(1) {
			// neighboring segment is the separating axis
			return *(*Manifold)(unsafe.Pointer(bp))
		}
		// fall through segment normal axis
	}
	if !(incidentNormal != -int32(1) || incidentIndex != -int32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+10103, __ccgo_ts+9936, int32FromInt32(1674)) != 0 {
		__builtin_trap(tls)
	}
	if incidentNormal != -int32(1) {
		ib11 = incidentNormal
		if ib11 < count-int32(1) {
			v342 = ib11 + int32(1)
		} else {
			v342 = 0
		}
		ib21 = v342
		b12 = (*(*[8]Vec2)(unsafe.Pointer(bp + 112)))[ib11]
		b22 = (*(*[8]Vec2)(unsafe.Pointer(bp + 112)))[ib21]
	} else {
		i21 = incidentIndex
		if i21 > 0 {
			v343 = i21 - int32(1)
		} else {
			v343 = count - int32(1)
		}
		i11 = v343
		v344 = normal1
		v345 = normals[i11]
		v346 = float32(v344.X*v345.X) + float32(v344.Y*v345.Y)
		goto _347
	_347:
		d1 = v346
		v348 = normal1
		v349 = normals[i21]
		v350 = float32(v348.X*v349.X) + float32(v348.Y*v349.Y)
		goto _351
	_351:
		d2 = v350
		if d1 < d2 {
			ib11 = i11
			ib21 = i21
			b12 = (*(*[8]Vec2)(unsafe.Pointer(bp + 112)))[ib11]
			b22 = (*(*[8]Vec2)(unsafe.Pointer(bp + 112)))[ib21]
		} else {
			ib11 = i21
			if i21 < count-int32(1) {
				v352 = i21 + int32(1)
			} else {
				v352 = 0
			}
			ib21 = v352
			b12 = (*(*[8]Vec2)(unsafe.Pointer(bp + 112)))[ib11]
			b22 = (*(*[8]Vec2)(unsafe.Pointer(bp + 112)))[ib21]
		}
	}
	*(*Manifold)(unsafe.Pointer(bp)) = b2ClipSegments(tls, p11, p2, b12, b22, normal1, float32FromFloat32(0), radiusB, uint16FromInt32(int32FromUint8(uint8FromInt32(int32FromInt32(0)))<<int32FromInt32(8)|int32FromUint8(uint8FromInt32(ib21))), uint16FromInt32(int32FromUint8(uint8FromInt32(int32FromInt32(1)))<<int32FromInt32(8)|int32FromUint8(uint8FromInt32(ib11))))
	if !((*(*Manifold)(unsafe.Pointer(bp))).PointCount == 0 || (*(*Manifold)(unsafe.Pointer(bp))).PointCount == int32(2)) && b2InternalAssertFcn(tls, __ccgo_ts+10050, __ccgo_ts+9936, int32FromInt32(1711)) != 0 {
		__builtin_trap(tls)
	}
	if (*(*Manifold)(unsafe.Pointer(bp))).PointCount == int32(2) {
		// There may be no points c
		v353 = xfA.Q
		v354 = (*(*Manifold)(unsafe.Pointer(bp))).Normal
		v355 = Vec2{
			X: float32(v353.C*v354.X) - float32(v353.S*v354.Y),
			Y: float32(v353.S*v354.X) + float32(v353.C*v354.Y),
		}
		goto _356
	_356:
		(*(*Manifold)(unsafe.Pointer(bp))).Normal = v355
		v357 = xfA.Q
		v358 = (*(*ManifoldPoint)(unsafe.Pointer(bp + 12))).AnchorA
		v359 = Vec2{
			X: float32(v357.C*v358.X) - float32(v357.S*v358.Y),
			Y: float32(v357.S*v358.X) + float32(v357.C*v358.Y),
		}
		goto _360
	_360:
		(*(*ManifoldPoint)(unsafe.Pointer(bp + 12))).AnchorA = v359
		v361 = xfA.Q
		v362 = (*(*ManifoldPoint)(unsafe.Pointer(bp + 12 + 1*48))).AnchorA
		v363 = Vec2{
			X: float32(v361.C*v362.X) - float32(v361.S*v362.Y),
			Y: float32(v361.S*v362.X) + float32(v361.C*v362.Y),
		}
		goto _364
	_364:
		(*(*ManifoldPoint)(unsafe.Pointer(bp + 12 + 1*48))).AnchorA = v363
		v365 = xfA.P
		v366 = xfB.P
		v367 = Vec2{
			X: v365.X - v366.X,
			Y: v365.Y - v366.Y,
		}
		goto _368
	_368:
		pAB2 = v367
		v369 = (*(*ManifoldPoint)(unsafe.Pointer(bp + 12))).AnchorA
		v370 = pAB2
		v371 = Vec2{
			X: v369.X + v370.X,
			Y: v369.Y + v370.Y,
		}
		goto _372
	_372:
		(*(*ManifoldPoint)(unsafe.Pointer(bp + 12))).AnchorB = v371
		v373 = (*(*ManifoldPoint)(unsafe.Pointer(bp + 12 + 1*48))).AnchorA
		v374 = pAB2
		v375 = Vec2{
			X: v373.X + v374.X,
			Y: v373.Y + v374.Y,
		}
		goto _376
	_376:
		(*(*ManifoldPoint)(unsafe.Pointer(bp + 12 + 1*48))).AnchorB = v375
		v377 = xfA.P
		v378 = (*(*ManifoldPoint)(unsafe.Pointer(bp + 12))).AnchorA
		v379 = Vec2{
			X: v377.X + v378.X,
			Y: v377.Y + v378.Y,
		}
		goto _380
	_380:
		(*(*ManifoldPoint)(unsafe.Pointer(bp + 12))).Point = v379
		v381 = xfA.P
		v382 = (*(*ManifoldPoint)(unsafe.Pointer(bp + 12 + 1*48))).AnchorA
		v383 = Vec2{
			X: v381.X + v382.X,
			Y: v381.Y + v382.Y,
		}
		goto _384
	_384:
		(*(*ManifoldPoint)(unsafe.Pointer(bp + 12 + 1*48))).Point = v383
	}
	return *(*Manifold)(unsafe.Pointer(bp))
}

func b2CollideMover(tls *_Stack, shape uintptr, transform Transform, mover uintptr) (r PlaneResult) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var result PlaneResult
	var vx, vy float32
	var v1, v5 Transform
	var v10, v11, v2, v3, v6, v7 Vec2
	var v9 Rot
	var _ /* localMover at bp+0 */ Capsule
	_, _, _, _, _, _, _, _, _, _, _, _ = result, vx, vy, v1, v10, v11, v2, v3, v5, v6, v7, v9
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
	(*(*Capsule)(unsafe.Pointer(bp))).Center1 = v3
	v5 = transform
	v6 = (*Capsule)(unsafe.Pointer(mover)).Center2
	vx = v6.X - v5.P.X
	vy = v6.Y - v5.P.Y
	v7 = Vec2{
		X: float32(v5.Q.C*vx) + float32(v5.Q.S*vy),
		Y: float32(-v5.Q.S*vx) + float32(v5.Q.C*vy),
	}
	goto _8
_8:
	(*(*Capsule)(unsafe.Pointer(bp))).Center2 = v7
	(*(*Capsule)(unsafe.Pointer(bp))).Radius = (*Capsule)(unsafe.Pointer(mover)).Radius
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
	v10 = result.Plane.Normal
	v11 = Vec2{
		X: float32(v9.C*v10.X) - float32(v9.S*v10.Y),
		Y: float32(v9.S*v10.X) + float32(v9.C*v10.Y),
	}
	goto _12
_12:
	result.Plane.Normal = v11
	return result
}

func b2CollideTask(tls *_Stack, startIndex int32, endIndex int32, threadIndex uint32, context uintptr) {
	var blockIndex, v19, v21, v7 uint32_t
	var bodies, bodyA, bodyB, bodySimA, bodySimB, contactSim, contactSims, shapeA, shapeB, shapes, stepContext, taskContext, world, v18, v20, v6 uintptr
	var centerOffsetA, centerOffsetB, v11, v12, v15, v16 Vec2
	var contactId, contactIndex, v8, v9 int32
	var overlap, touching, wasTouching, v4 uint8
	var transformA, transformB Transform
	var v10, v14 Rot
	var v2, v3 AABB
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = blockIndex, bodies, bodyA, bodyB, bodySimA, bodySimB, centerOffsetA, centerOffsetB, contactId, contactIndex, contactSim, contactSims, overlap, shapeA, shapeB, shapes, stepContext, taskContext, touching, transformA, transformB, wasTouching, world, v10, v11, v12, v14, v15, v16, v18, v19, v2, v20, v21, v3, v4, v6, v7, v8, v9
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
		contactSim = *(*uintptr)(unsafe.Pointer(contactSims + uintptr(contactIndex)*8))
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
			*(*uint32_t)(unsafe.Pointer(contactSim + 164)) |= uint32(b2_simDisjoint)
			*(*uint32_t)(unsafe.Pointer(contactSim + 164)) &= uint32FromInt32(^int32(b2_simTouchingFlag))
			v6 = taskContext
			v7 = uint32FromInt32(contactId)
			blockIndex = v7 / uint32(64)
			if !(blockIndex < (*b2BitSet)(unsafe.Pointer(v6)).BlockCount) && b2InternalAssertFcn(tls, __ccgo_ts+10859, __ccgo_ts+10891, int32FromInt32(28)) != 0 {
				__builtin_trap(tls)
			}
			*(*uint64_t)(unsafe.Pointer((*b2BitSet)(unsafe.Pointer(v6)).Bits + uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v7 % uint32FromInt32(64))
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
				v9 = (*b2Body)(unsafe.Pointer(bodyB)).LocalIndex
			} else {
				v9 = -int32(1)
			}
			(*b2ContactSim)(unsafe.Pointer(contactSim)).BodySimIndexB = v9
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
				*(*uint32_t)(unsafe.Pointer(contactSim + 164)) |= uint32(b2_simStartedTouching)
				v18 = taskContext
				v19 = uint32FromInt32(contactId)
				blockIndex = v19 / uint32(64)
				if !(blockIndex < (*b2BitSet)(unsafe.Pointer(v18)).BlockCount) && b2InternalAssertFcn(tls, __ccgo_ts+10859, __ccgo_ts+10891, int32FromInt32(28)) != 0 {
					__builtin_trap(tls)
				}
				*(*uint64_t)(unsafe.Pointer((*b2BitSet)(unsafe.Pointer(v18)).Bits + uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v19 % uint32FromInt32(64))
			} else {
				if int32FromUint8(touching) == false1 && int32FromUint8(wasTouching) == int32(true1) {
					*(*uint32_t)(unsafe.Pointer(contactSim + 164)) |= uint32(b2_simStoppedTouching)
					v20 = taskContext
					v21 = uint32FromInt32(contactId)
					blockIndex = v21 / uint32(64)
					if !(blockIndex < (*b2BitSet)(unsafe.Pointer(v20)).BlockCount) && b2InternalAssertFcn(tls, __ccgo_ts+10859, __ccgo_ts+10891, int32FromInt32(28)) != 0 {
						__builtin_trap(tls)
					}
					*(*uint64_t)(unsafe.Pointer((*b2BitSet)(unsafe.Pointer(v20)).Bits + uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v21 % uint32FromInt32(64))
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
	var awakeSet, base, base1, bitSet, color, color1, contact, contactSim, contactSims, graphColors, shapeA, shapeB, shapes, userCollideTask, world, v10, v12, v17, v19, v21, v23, v25, v27, v29, v31, v33, v35 uintptr
	var bits uint64_t
	var bodyIdA, bodyIdB, colorIndex, contactCount, contactId, contactIdCapacity, contactIndex, count, endEventArrayIndex, i, i1, i2, i3, i4, j, localIndex, minRange, newCapacity, newCapacity1, nonTouchingCount, v1, v11, v18, v22, v26, v30, v32, v36, v6 int32
	var ctz, flags, k, simFlags, v15 uint32_t
	var event ContactBeginTouchEvent
	var event1 ContactEndTouchEvent
	var shapeIdA, shapeIdB ShapeId
	var worldId uint16_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = awakeSet, base, base1, bitSet, bits, bodyIdA, bodyIdB, color, color1, colorIndex, contact, contactCount, contactId, contactIdCapacity, contactIndex, contactSim, contactSims, count, ctz, endEventArrayIndex, event, event1, flags, graphColors, i, i1, i2, i3, i4, j, k, localIndex, minRange, newCapacity, newCapacity1, nonTouchingCount, shapeA, shapeB, shapeIdA, shapeIdB, shapes, simFlags, userCollideTask, world, worldId, v1, v10, v11, v12, v15, v17, v18, v19, v21, v22, v23, v25, v26, v27, v29, v30, v31, v32, v33, v35, v36, v6
	world = (*b2StepContext)(unsafe.Pointer(context)).World
	if !((*b2World)(unsafe.Pointer(world)).WorkerCount > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+10955, __ccgo_ts+15342, int32FromInt32(492)) != 0 {
		__builtin_trap(tls)
	}
	// Task that can be done in parallel with the narrow-phase
	// - rebuild the collision tree for dynamic and kinematic bodies to keep their query performance good
	// todo_erin move this to start when contacts are being created
	(*b2World)(unsafe.Pointer(world)).UserTreeTask = (*(*func(*_Stack, uintptr, int32, int32, uintptr, uintptr) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).EnqueueTaskFcn})))(tls, __ccgo_fp(b2UpdateTreesTask), int32(1), int32(1), world, (*b2World)(unsafe.Pointer(world)).UserTaskContext)
	*(*int32)(unsafe.Pointer(world + 1776)) += int32(1)
	if (*b2World)(unsafe.Pointer(world)).UserTreeTask == uintptrFromInt32(0) {
		v1 = 0
	} else {
		v1 = int32(1)
	}
	*(*int32)(unsafe.Pointer(world + 1772)) += v1
	// gather contacts into a single array for easier parallel-for
	contactCount = 0
	graphColors = world + 328
	i = 0
	for {
		if !(i < int32(_B2_GRAPH_COLOR_COUNT)) {
			break
		}
		contactCount = contactCount + (*(*b2GraphColor)(unsafe.Pointer(graphColors + uintptr(i)*56))).ContactSims.Count
		goto _2
	_2:
		;
		i = i + 1
	}
	nonTouchingCount = (*(*b2SolverSet)(unsafe.Pointer((*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(b2_awakeSet)*88))).ContactSims.Count
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
			*(*uintptr)(unsafe.Pointer(contactSims + uintptr(contactIndex)*8)) = base + uintptr(j)*176
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
	base1 = (*(*b2SolverSet)(unsafe.Pointer((*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(b2_awakeSet)*88))).ContactSims.Data
	i2 = 0
	for {
		if !(i2 < nonTouchingCount) {
			break
		}
		*(*uintptr)(unsafe.Pointer(contactSims + uintptr(contactIndex)*8)) = base1 + uintptr(i2)*176
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
	v6 = (*b2IdPool)(unsafe.Pointer(world + 1120)).NextIndex
	goto _7
_7:
	// Contact bit set on ids because contact pointers are unstable as they move between touching and not touching.
	contactIdCapacity = v6
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
	*(*int32)(unsafe.Pointer(world + 1776)) += int32(1)
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
	v11 = int32(b2_awakeSet)
	if !(0 <= v11 && v11 < (*b2SolverSetArray)(unsafe.Pointer(v10)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v12 = (*b2SolverSetArray)(unsafe.Pointer(v10)).Data + uintptr(v11)*88
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
		bits = *(*uint64_t)(unsafe.Pointer((*b2BitSet)(unsafe.Pointer(bitSet)).Bits + uintptr(k)*8))
		for bits != uint64(0) {
			v15 = uint32FromInt32(__builtin_ctzll(tls, bits))
			goto _16
		_16:
			ctz = v15
			contactId = int32FromUint32(uint32FromInt32(64)*k + ctz)
			v17 = world + 1144
			v18 = contactId
			if !(0 <= v18 && v18 < (*b2ContactArray)(unsafe.Pointer(v17)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
				__builtin_trap(tls)
			}
			v19 = (*b2ContactArray)(unsafe.Pointer(v17)).Data + uintptr(v18)*68
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
				v21 = color1 + 16
				v22 = localIndex
				if !(0 <= v22 && v22 < (*b2ContactSimArray)(unsafe.Pointer(v21)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(147)) != 0 {
					__builtin_trap(tls)
				}
				v23 = (*b2ContactSimArray)(unsafe.Pointer(v21)).Data + uintptr(v22)*176
				goto _24
			_24:
				contactSim = v23
			} else {
				v25 = awakeSet + 48
				v26 = localIndex
				if !(0 <= v26 && v26 < (*b2ContactSimArray)(unsafe.Pointer(v25)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(147)) != 0 {
					__builtin_trap(tls)
				}
				v27 = (*b2ContactSimArray)(unsafe.Pointer(v25)).Data + uintptr(v26)*176
				goto _28
			_28:
				contactSim = v27
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
						v29 = world + 1360
						if (*b2ContactBeginTouchEventArray)(unsafe.Pointer(v29)).Count == (*b2ContactBeginTouchEventArray)(unsafe.Pointer(v29)).Capacity {
							if (*b2ContactBeginTouchEventArray)(unsafe.Pointer(v29)).Capacity < int32(2) {
								v30 = int32(2)
							} else {
								v30 = (*b2ContactBeginTouchEventArray)(unsafe.Pointer(v29)).Capacity + (*b2ContactBeginTouchEventArray)(unsafe.Pointer(v29)).Capacity>>int32(1)
							}
							newCapacity = v30
							b2ContactBeginTouchEventArray_Reserve(tls, v29, newCapacity)
						}
						*(*ContactBeginTouchEvent)(unsafe.Pointer((*b2ContactBeginTouchEventArray)(unsafe.Pointer(v29)).Data + uintptr((*b2ContactBeginTouchEventArray)(unsafe.Pointer(v29)).Count)*128)) = event
						*(*int32)(unsafe.Pointer(v29 + 8)) += int32(1)
					}
					if !((*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.PointCount > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+2821, __ccgo_ts+15342, int32FromInt32(638)) != 0 {
						__builtin_trap(tls)
					}
					if !((*b2Contact)(unsafe.Pointer(contact)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+3637, __ccgo_ts+15342, int32FromInt32(639)) != 0 {
						__builtin_trap(tls)
					}
					// Link first because this wakes colliding bodies and ensures the body sims
					// are in the correct place.
					*(*uint32_t)(unsafe.Pointer(contact + 60)) |= uint32(b2_contactTouchingFlag)
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
					v31 = awakeSet + 48
					v32 = localIndex
					if !(0 <= v32 && v32 < (*b2ContactSimArray)(unsafe.Pointer(v31)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(147)) != 0 {
						__builtin_trap(tls)
					}
					v33 = (*b2ContactSimArray)(unsafe.Pointer(v31)).Data + uintptr(v32)*176
					goto _34
				_34:
					contactSim = v33
					*(*uint32_t)(unsafe.Pointer(contactSim + 164)) &= uint32FromInt32(^int32(b2_simStartedTouching))
					b2AddContactToGraph(tls, world, contactSim, contact)
					b2RemoveNonTouchingContact(tls, world, int32(b2_awakeSet), localIndex)
					contactSim = uintptrFromInt32(0)
				} else {
					if simFlags&uint32(b2_simStoppedTouching) != 0 {
						*(*uint32_t)(unsafe.Pointer(contactSim + 164)) &= uint32FromInt32(^int32(b2_simStoppedTouching))
						*(*uint32_t)(unsafe.Pointer(contact + 60)) &= uint32FromInt32(^int32(b2_contactTouchingFlag))
						if (*b2Contact)(unsafe.Pointer(contact)).Flags&uint32(b2_contactEnableContactEvents) != 0 {
							event1 = ContactEndTouchEvent{
								ShapeIdA: shapeIdA,
								ShapeIdB: shapeIdB,
							}
							v35 = world + 1408 + uintptr(endEventArrayIndex)*16
							if (*b2ContactEndTouchEventArray)(unsafe.Pointer(v35)).Count == (*b2ContactEndTouchEventArray)(unsafe.Pointer(v35)).Capacity {
								if (*b2ContactEndTouchEventArray)(unsafe.Pointer(v35)).Capacity < int32(2) {
									v36 = int32(2)
								} else {
									v36 = (*b2ContactEndTouchEventArray)(unsafe.Pointer(v35)).Capacity + (*b2ContactEndTouchEventArray)(unsafe.Pointer(v35)).Capacity>>int32(1)
								}
								newCapacity1 = v36
								b2ContactEndTouchEventArray_Reserve(tls, v35, newCapacity1)
							}
							*(*ContactEndTouchEvent)(unsafe.Pointer((*b2ContactEndTouchEventArray)(unsafe.Pointer(v35)).Data + uintptr((*b2ContactEndTouchEventArray)(unsafe.Pointer(v35)).Count)*16)) = event1
							*(*int32)(unsafe.Pointer(v35 + 8)) += int32(1)
						}
						if !((*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.PointCount == int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+14641, __ccgo_ts+15342, int32FromInt32(671)) != 0 {
							__builtin_trap(tls)
						}
						b2UnlinkContact(tls, world, contact)
						bodyIdA = (*(*b2ContactEdge)(unsafe.Pointer(contact + 12))).BodyId
						bodyIdB = (*(*b2ContactEdge)(unsafe.Pointer(contact + 12 + 1*12))).BodyId
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
