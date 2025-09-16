package b2

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

// C documentation
//
//	// Uses GJK for computing the distance between convex shapes.
//	// https://box2d.org/files/ErinCatto_GJK_GDC2010.pdf
//	// I spent time optimizing this and could find no further significant gains 3/30/2025
func b2ShapeDistance(tls *_Stack, input uintptr, cache uintptr, simplexes uintptr, simplexCapacity int32) (r DistanceOutput) {
	bp := tls.Alloc(240)
	defer tls.Free(240)
	var C, transform, v1, v15, v18, v2, v34, v38, v46, v50, v84, v88 Transform
	var aa, dx, dy, invLength, length, radiusA, radiusB, x, y, v103, v44, v68, v72, v73, v75, v82, v92, v93, v94, v96, v98 float32
	var d, n, nonUnitNormal, normal, v100, v102, v104, v105, v12, v13, v19, v20, v28, v29, v35, v36, v39, v40, v42, v43, v47, v48, v51, v52, v54, v55, v57, v58, v59, v62, v63, v65, v66, v67, v7, v77, v78, v8, v80, v81, v85, v86, v89, v9, v90, v97, v99 Vec2
	var duplicate, v70 uint8
	var i, i1, i2, iteration, maxIterations, saveCount, simplexIndex int32
	var output DistanceOutput
	var proxyA, vertex uintptr
	var qr, v11, v3, v4, v5, v76 Rot
	var saveA, saveB [3]int32
	var vertices [3]uintptr
	var _ /* localPointA at bp+184 */ Vec2
	var _ /* localPointA at bp+200 */ Vec2
	var _ /* localPointA at bp+216 */ Vec2
	var _ /* localPointB at bp+192 */ Vec2
	var _ /* localPointB at bp+208 */ Vec2
	var _ /* localPointB at bp+224 */ Vec2
	var _ /* localProxyB at bp+0 */ ShapeProxy
	var _ /* simplex at bp+72 */ Simplex
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, aa, d, duplicate, dx, dy, i, i1, i2, invLength, iteration, length, maxIterations, n, nonUnitNormal, normal, output, proxyA, qr, radiusA, radiusB, saveA, saveB, saveCount, simplexIndex, transform, vertex, vertices, x, y, v1, v100, v102, v103, v104, v105, v11, v12, v13, v15, v18, v19, v2, v20, v28, v29, v3, v34, v35, v36, v38, v39, v4, v40, v42, v43, v44, v46, v47, v48, v5, v50, v51, v52, v54, v55, v57, v58, v59, v62, v63, v65, v66, v67, v68, v7, v70, v72, v73, v75, v76, v77, v78, v8, v80, v81, v82, v84, v85, v86, v88, v89, v9, v90, v92, v93, v94, v96, v97, v98, v99
	_ = uint64FromInt64(4)
	if !((*DistanceInput)(unsafe.Pointer(input)).ProxyA.Count > 0 && (*DistanceInput)(unsafe.Pointer(input)).ProxyB.Count > 0) && b2InternalAssertFcn(tls, __ccgo_ts+4129, __ccgo_ts+4097, int32FromInt32(424)) != 0 {
		__builtin_trap(tls)
	}
	if !((*DistanceInput)(unsafe.Pointer(input)).ProxyA.Radius >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+4180, __ccgo_ts+4097, int32FromInt32(425)) != 0 {
		__builtin_trap(tls)
	}
	if !((*DistanceInput)(unsafe.Pointer(input)).ProxyB.Radius >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+4209, __ccgo_ts+4097, int32FromInt32(426)) != 0 {
		__builtin_trap(tls)
	}
	output = DistanceOutput{}
	proxyA = input
	v1 = (*DistanceInput)(unsafe.Pointer(input)).TransformA
	v2 = (*DistanceInput)(unsafe.Pointer(input)).TransformB
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
	transform = v15
	(*(*ShapeProxy)(unsafe.Pointer(bp))).Count = (*DistanceInput)(unsafe.Pointer(input)).ProxyB.Count
	(*(*ShapeProxy)(unsafe.Pointer(bp))).Radius = (*DistanceInput)(unsafe.Pointer(input)).ProxyB.Radius
	i = 0
	for {
		if !(i < (*(*ShapeProxy)(unsafe.Pointer(bp))).Count) {
			break
		}
		v18 = transform
		v19 = *(*Vec2)(unsafe.Pointer(input + 72 + uintptr(i)*8))
		x = float32(v18.Q.C*v19.X) - float32(v18.Q.S*v19.Y) + v18.P.X
		y = float32(v18.Q.S*v19.X) + float32(v18.Q.C*v19.Y) + v18.P.Y
		v20 = Vec2{
			X: x,
			Y: y,
		}
		goto _21
	_21:
		*(*Vec2)(unsafe.Pointer(bp + uintptr(i)*8)) = v20
		goto _17
	_17:
		;
		i = i + 1
	}
	// Initialize the simplex.
	*(*Simplex)(unsafe.Pointer(bp + 72)) = b2MakeSimplexFromCache(tls, cache, proxyA, bp)
	simplexIndex = 0
	if simplexes != uintptrFromInt32(0) && simplexIndex < simplexCapacity {
		*(*Simplex)(unsafe.Pointer(simplexes + uintptr(simplexIndex)*112)) = *(*Simplex)(unsafe.Pointer(bp + 72))
		simplexIndex = simplexIndex + int32(1)
	}
	// Get simplex vertices as an array.
	vertices = [3]uintptr{
		0: bp + 72,
		1: bp + 72 + 36,
		2: bp + 72 + 72,
	}
	nonUnitNormal = b2Vec2_zero
	// Main iteration loop. All computations are done in frame A.
	maxIterations = int32(20)
	iteration = 0
	for iteration < maxIterations {
		// Copy simplex so we can identify duplicates.
		saveCount = (*(*Simplex)(unsafe.Pointer(bp + 72))).Count
		i1 = 0
		for {
			if !(i1 < saveCount) {
				break
			}
			saveA[i1] = (*SimplexVertex)(unsafe.Pointer(vertices[i1])).IndexA
			saveB[i1] = (*SimplexVertex)(unsafe.Pointer(vertices[i1])).IndexB
			goto _22
		_22:
			;
			i1 = i1 + 1
		}
		d = Vec2{}
		switch (*(*Simplex)(unsafe.Pointer(bp + 72))).Count {
		case int32(1):
			goto _23
		case int32(2):
			goto _24
		case int32(3):
			goto _25
		default:
			goto _26
		}
		goto _27
	_23:
		;
		v28 = (*(*Simplex)(unsafe.Pointer(bp + 72))).V1.W
		v29 = Vec2{
			X: -v28.X,
			Y: -v28.Y,
		}
		goto _30
	_30:
		d = v29
		goto _27
	_24:
		;
		d = b2SolveSimplex2(tls, bp+72)
		goto _27
	_25:
		;
		d = b2SolveSimplex3(tls, bp+72)
		goto _27
	_26:
		;
	_33:
		;
		if bool(!(int32FromInt32(false1) != 0)) && b2InternalAssertFcn(tls, __ccgo_ts+4123, __ccgo_ts+4097, int32FromInt32(492)) != 0 {
			__builtin_trap(tls)
		}
		goto _32
	_32:
		;
		if 0 != 0 {
			goto _33
		}
		goto _31
	_31:
		;
	_27:
		;
		// If we have 3 points, then the origin is in the corresponding triangle.
		if (*(*Simplex)(unsafe.Pointer(bp + 72))).Count == int32(3) {
			b2ComputeSimplexWitnessPoints(tls, bp+184, bp+192, bp+72)
			v34 = (*DistanceInput)(unsafe.Pointer(input)).TransformA
			v35 = *(*Vec2)(unsafe.Pointer(bp + 184))
			x = float32(v34.Q.C*v35.X) - float32(v34.Q.S*v35.Y) + v34.P.X
			y = float32(v34.Q.S*v35.X) + float32(v34.Q.C*v35.Y) + v34.P.Y
			v36 = Vec2{
				X: x,
				Y: y,
			}
			goto _37
		_37:
			output.PointA = v36
			v38 = (*DistanceInput)(unsafe.Pointer(input)).TransformA
			v39 = *(*Vec2)(unsafe.Pointer(bp + 192))
			x = float32(v38.Q.C*v39.X) - float32(v38.Q.S*v39.Y) + v38.P.X
			y = float32(v38.Q.S*v39.X) + float32(v38.Q.C*v39.Y) + v38.P.Y
			v40 = Vec2{
				X: x,
				Y: y,
			}
			goto _41
		_41:
			output.PointB = v40
			return output
		}
		if simplexes != uintptrFromInt32(0) && simplexIndex < simplexCapacity {
			*(*Simplex)(unsafe.Pointer(simplexes + uintptr(simplexIndex)*112)) = *(*Simplex)(unsafe.Pointer(bp + 72))
			simplexIndex = simplexIndex + int32(1)
		}
		// Ensure the search direction is numerically fit.
		v42 = d
		v43 = d
		v44 = float32(v42.X*v43.X) + float32(v42.Y*v43.Y)
		goto _45
	_45:
		if v44 < float32(float32FromFloat32(1.1920928955078125e-07)*float32FromFloat32(1.1920928955078125e-07)) {
			b2ComputeSimplexWitnessPoints(tls, bp+200, bp+208, bp+72)
			v46 = (*DistanceInput)(unsafe.Pointer(input)).TransformA
			v47 = *(*Vec2)(unsafe.Pointer(bp + 200))
			x = float32(v46.Q.C*v47.X) - float32(v46.Q.S*v47.Y) + v46.P.X
			y = float32(v46.Q.S*v47.X) + float32(v46.Q.C*v47.Y) + v46.P.Y
			v48 = Vec2{
				X: x,
				Y: y,
			}
			goto _49
		_49:
			output.PointA = v48
			v50 = (*DistanceInput)(unsafe.Pointer(input)).TransformA
			v51 = *(*Vec2)(unsafe.Pointer(bp + 208))
			x = float32(v50.Q.C*v51.X) - float32(v50.Q.S*v51.Y) + v50.P.X
			y = float32(v50.Q.S*v51.X) + float32(v50.Q.C*v51.Y) + v50.P.Y
			v52 = Vec2{
				X: x,
				Y: y,
			}
			goto _53
		_53:
			output.PointB = v52
			return output
		}
		// Save the normal
		nonUnitNormal = d
		// Compute a tentative new simplex vertex using support points.
		// support = support(a, d) - support(b, -d)
		vertex = vertices[(*(*Simplex)(unsafe.Pointer(bp + 72))).Count]
		(*SimplexVertex)(unsafe.Pointer(vertex)).IndexA = b2FindSupport(tls, proxyA, d)
		(*SimplexVertex)(unsafe.Pointer(vertex)).WA = *(*Vec2)(unsafe.Pointer(proxyA + uintptr((*SimplexVertex)(unsafe.Pointer(vertex)).IndexA)*8))
		v54 = d
		v55 = Vec2{
			X: -v54.X,
			Y: -v54.Y,
		}
		goto _56
	_56:
		(*SimplexVertex)(unsafe.Pointer(vertex)).IndexB = b2FindSupport(tls, bp, v55)
		(*SimplexVertex)(unsafe.Pointer(vertex)).WB = *(*Vec2)(unsafe.Pointer(bp + uintptr((*SimplexVertex)(unsafe.Pointer(vertex)).IndexB)*8))
		v57 = (*SimplexVertex)(unsafe.Pointer(vertex)).WA
		v58 = (*SimplexVertex)(unsafe.Pointer(vertex)).WB
		v59 = Vec2{
			X: v57.X - v58.X,
			Y: v57.Y - v58.Y,
		}
		goto _60
	_60:
		(*SimplexVertex)(unsafe.Pointer(vertex)).W = v59
		// Iteration count is equated to the number of support point calls.
		iteration = iteration + 1
		// Check for duplicate support points. This is the main termination criteria.
		duplicate = boolUint8(false1 != 0)
		i2 = 0
		for {
			if !(i2 < saveCount) {
				break
			}
			if (*SimplexVertex)(unsafe.Pointer(vertex)).IndexA == saveA[i2] && (*SimplexVertex)(unsafe.Pointer(vertex)).IndexB == saveB[i2] {
				duplicate = boolUint8(true1 != 0)
				break
			}
			goto _61
		_61:
			;
			i2 = i2 + 1
		}
		// If we found a duplicate support point we must exit to avoid cycling.
		if duplicate != 0 {
			break
		}
		// New vertex is valid and needed.
		(*(*Simplex)(unsafe.Pointer(bp + 72))).Count += int32(1)
	}
	if simplexes != uintptrFromInt32(0) && simplexIndex < simplexCapacity {
		*(*Simplex)(unsafe.Pointer(simplexes + uintptr(simplexIndex)*112)) = *(*Simplex)(unsafe.Pointer(bp + 72))
		simplexIndex = simplexIndex + int32(1)
	}
	v62 = nonUnitNormal
	length = sqrtf(tls, float32(v62.X*v62.X)+float32(v62.Y*v62.Y))
	if length < float32FromFloat32(1.1920928955078125e-07) {
		v63 = Vec2{}
		goto _64
	}
	invLength = float32FromFloat32(1) / length
	n = Vec2{
		X: float32(invLength * v62.X),
		Y: float32(invLength * v62.Y),
	}
	v63 = n
	goto _64
_64:
	// Prepare output
	normal = v63
	v65 = normal
	v66 = v65
	v67 = v65
	v68 = float32(v66.X*v67.X) + float32(v66.Y*v67.Y)
	goto _69
_69:
	aa = v68
	v72 = float32FromFloat32(1) - aa
	if v72 < float32FromInt32(0) {
		v75 = -v72
	} else {
		v75 = v72
	}
	v73 = v75
	goto _74
_74:
	v70 = boolUint8(v73 < float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07)))
	goto _71
_71:
	;
	if !(v70 != 0) && b2InternalAssertFcn(tls, __ccgo_ts+4238, __ccgo_ts+4097, int32FromInt32(577)) != 0 {
		__builtin_trap(tls)
	}
	v76 = (*DistanceInput)(unsafe.Pointer(input)).TransformA.Q
	v77 = normal
	v78 = Vec2{
		X: float32(v76.C*v77.X) - float32(v76.S*v77.Y),
		Y: float32(v76.S*v77.X) + float32(v76.C*v77.Y),
	}
	goto _79
_79:
	normal = v78
	b2ComputeSimplexWitnessPoints(tls, bp+216, bp+224, bp+72)
	output.Normal = normal
	v80 = *(*Vec2)(unsafe.Pointer(bp + 216))
	v81 = *(*Vec2)(unsafe.Pointer(bp + 224))
	dx = v81.X - v80.X
	dy = v81.Y - v80.Y
	v82 = sqrtf(tls, float32(dx*dx)+float32(dy*dy))
	goto _83
_83:
	output.Distance = v82
	v84 = (*DistanceInput)(unsafe.Pointer(input)).TransformA
	v85 = *(*Vec2)(unsafe.Pointer(bp + 216))
	x = float32(v84.Q.C*v85.X) - float32(v84.Q.S*v85.Y) + v84.P.X
	y = float32(v84.Q.S*v85.X) + float32(v84.Q.C*v85.Y) + v84.P.Y
	v86 = Vec2{
		X: x,
		Y: y,
	}
	goto _87
_87:
	output.PointA = v86
	v88 = (*DistanceInput)(unsafe.Pointer(input)).TransformA
	v89 = *(*Vec2)(unsafe.Pointer(bp + 224))
	x = float32(v88.Q.C*v89.X) - float32(v88.Q.S*v89.Y) + v88.P.X
	y = float32(v88.Q.S*v89.X) + float32(v88.Q.C*v89.Y) + v88.P.Y
	v90 = Vec2{
		X: x,
		Y: y,
	}
	goto _91
_91:
	output.PointB = v90
	output.Iterations = iteration
	output.SimplexCount = simplexIndex
	// Cache the simplex
	b2MakeSimplexCache(tls, cache, bp+72)
	// Apply radii if requested
	if (*DistanceInput)(unsafe.Pointer(input)).UseRadii != 0 && output.Distance > float32(float32FromFloat32(0.1)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
		radiusA = (*DistanceInput)(unsafe.Pointer(input)).ProxyA.Radius
		radiusB = (*DistanceInput)(unsafe.Pointer(input)).ProxyB.Radius
		v92 = float32FromFloat32(0)
		v93 = output.Distance - radiusA - radiusB
		if v92 > v93 {
			v96 = v92
		} else {
			v96 = v93
		}
		v94 = v96
		goto _95
	_95:
		output.Distance = v94
		// Keep closest points on perimeter even if overlapped, this way the points move smoothly.
		v97 = output.PointA
		v98 = radiusA
		v99 = normal
		v100 = Vec2{
			X: v97.X + float32(v98*v99.X),
			Y: v97.Y + float32(v98*v99.Y),
		}
		goto _101
	_101:
		output.PointA = v100
		v102 = output.PointB
		v103 = radiusB
		v104 = normal
		v105 = Vec2{
			X: v102.X - float32(v103*v104.X),
			Y: v102.Y - float32(v103*v104.Y),
		}
		goto _106
	_106:
		output.PointB = v105
	}
	return output
}

// C documentation
//
//	// Shape cast using conservative advancement
func b2ShapeCast(tls *_Stack, input uintptr) (r CastOutput) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var aa, denominator, fraction, linearSlop, target, tolerance, totalRadius, v1, v13, v19, v2, v25, v29, v3, v30, v32, v35, v42, v46, v47, v49, v5, v52, v55, v8 float32
	var c1, c2, delta2, v10, v12, v14, v15, v17, v18, v20, v22, v23, v24, v34, v36, v37, v39, v40, v41, v50, v51, v54, v56, v57, v7, v9 Vec2
	var distanceOutput DistanceOutput
	var iteration, maxIterations int32
	var output CastOutput
	var v27, v44 uint8
	var v33 bool
	var _ /* cache at bp+0 */ SimplexCache
	var _ /* distanceInput at bp+8 */ DistanceInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aa, c1, c2, delta2, denominator, distanceOutput, fraction, iteration, linearSlop, maxIterations, output, target, tolerance, totalRadius, v1, v10, v12, v13, v14, v15, v17, v18, v19, v2, v20, v22, v23, v24, v25, v27, v29, v3, v30, v32, v33, v34, v35, v36, v37, v39, v40, v41, v42, v44, v46, v47, v49, v5, v50, v51, v52, v54, v55, v56, v57, v7, v8, v9
	// Compute tolerance
	linearSlop = float32(float32FromFloat32(0.005) * b2_lengthUnitsPerMeter)
	totalRadius = (*ShapeCastPairInput)(unsafe.Pointer(input)).ProxyA.Radius + (*ShapeCastPairInput)(unsafe.Pointer(input)).ProxyB.Radius
	v1 = linearSlop
	v2 = totalRadius - linearSlop
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _4
_4:
	target = v3
	tolerance = float32(float32FromFloat32(0.25) * linearSlop)
	if !(target > tolerance) && b2InternalAssertFcn(tls, __ccgo_ts+4263, __ccgo_ts+4097, int32FromInt32(616)) != 0 {
		__builtin_trap(tls)
	}
	// Prepare input for distance query
	*(*SimplexCache)(unsafe.Pointer(bp)) = SimplexCache{}
	fraction = float32FromFloat32(0)
	*(*DistanceInput)(unsafe.Pointer(bp + 8)) = DistanceInput{}
	(*(*DistanceInput)(unsafe.Pointer(bp + 8))).ProxyA = (*ShapeCastPairInput)(unsafe.Pointer(input)).ProxyA
	(*(*DistanceInput)(unsafe.Pointer(bp + 8))).ProxyB = (*ShapeCastPairInput)(unsafe.Pointer(input)).ProxyB
	(*(*DistanceInput)(unsafe.Pointer(bp + 8))).TransformA = (*ShapeCastPairInput)(unsafe.Pointer(input)).TransformA
	(*(*DistanceInput)(unsafe.Pointer(bp + 8))).TransformB = (*ShapeCastPairInput)(unsafe.Pointer(input)).TransformB
	(*(*DistanceInput)(unsafe.Pointer(bp + 8))).UseRadii = boolUint8(false1 != 0)
	delta2 = (*ShapeCastPairInput)(unsafe.Pointer(input)).TranslationB
	output = CastOutput{}
	iteration = 0
	maxIterations = int32(20)
	for {
		if !(iteration < maxIterations) {
			break
		}
		output.Iterations += int32(1)
		distanceOutput = b2ShapeDistance(tls, bp+8, bp, uintptrFromInt32(0), 0)
		if distanceOutput.Distance < target+tolerance {
			if iteration == 0 {
				if (*ShapeCastPairInput)(unsafe.Pointer(input)).CanEncroach != 0 && distanceOutput.Distance > float32(float32FromFloat32(2)*linearSlop) {
					target = distanceOutput.Distance - linearSlop
				} else {
					// Initial overlap
					output.Hit = boolUint8(true1 != 0)
					v7 = distanceOutput.PointA
					v8 = (*ShapeCastPairInput)(unsafe.Pointer(input)).ProxyA.Radius
					v9 = distanceOutput.Normal
					v10 = Vec2{
						X: v7.X + float32(v8*v9.X),
						Y: v7.Y + float32(v8*v9.Y),
					}
					goto _11
				_11:
					// Compute a common point
					c1 = v10
					v12 = distanceOutput.PointB
					v13 = -(*ShapeCastPairInput)(unsafe.Pointer(input)).ProxyB.Radius
					v14 = distanceOutput.Normal
					v15 = Vec2{
						X: v12.X + float32(v13*v14.X),
						Y: v12.Y + float32(v13*v14.Y),
					}
					goto _16
				_16:
					c2 = v15
					v17 = c1
					v18 = c2
					v19 = float32FromFloat32(0.5)
					v20 = Vec2{
						X: float32((float32FromFloat32(1)-v19)*v17.X) + float32(v19*v18.X),
						Y: float32((float32FromFloat32(1)-v19)*v17.Y) + float32(v19*v18.Y),
					}
					goto _21
				_21:
					output.Point = v20
					return output
				}
			} else {
				// Regular hit
				if v33 = distanceOutput.Distance > float32FromFloat32(0); v33 {
					v22 = distanceOutput.Normal
					v23 = v22
					v24 = v22
					v25 = float32(v23.X*v24.X) + float32(v23.Y*v24.Y)
					goto _26
				_26:
					aa = v25
					v29 = float32FromFloat32(1) - aa
					if v29 < float32FromInt32(0) {
						v32 = -v29
					} else {
						v32 = v29
					}
					v30 = v32
					goto _31
				_31:
					v27 = boolUint8(v30 < float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07)))
					goto _28
				_28:
				}
				if !(v33 && v27 != 0) && b2InternalAssertFcn(tls, __ccgo_ts+4282, __ccgo_ts+4097, int32FromInt32(664)) != 0 {
					__builtin_trap(tls)
				}
				output.Fraction = fraction
				v34 = distanceOutput.PointA
				v35 = (*ShapeCastPairInput)(unsafe.Pointer(input)).ProxyA.Radius
				v36 = distanceOutput.Normal
				v37 = Vec2{
					X: v34.X + float32(v35*v36.X),
					Y: v34.Y + float32(v35*v36.Y),
				}
				goto _38
			_38:
				output.Point = v37
				output.Normal = distanceOutput.Normal
				output.Hit = boolUint8(true1 != 0)
				return output
			}
		}
		if !(distanceOutput.Distance > float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+4356, __ccgo_ts+4097, int32FromInt32(673)) != 0 {
			__builtin_trap(tls)
		}
		v39 = distanceOutput.Normal
		v40 = v39
		v41 = v39
		v42 = float32(v40.X*v41.X) + float32(v40.Y*v41.Y)
		goto _43
	_43:
		aa = v42
		v46 = float32FromFloat32(1) - aa
		if v46 < float32FromInt32(0) {
			v49 = -v46
		} else {
			v49 = v46
		}
		v47 = v49
		goto _48
	_48:
		v44 = boolUint8(v47 < float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07)))
		goto _45
	_45:
		;
		if !(v44 != 0) && b2InternalAssertFcn(tls, __ccgo_ts+4387, __ccgo_ts+4097, int32FromInt32(674)) != 0 {
			__builtin_trap(tls)
		}
		v50 = delta2
		v51 = distanceOutput.Normal
		v52 = float32(v50.X*v51.X) + float32(v50.Y*v51.Y)
		goto _53
	_53:
		// Check if shapes are approaching each other
		denominator = v52
		if denominator >= float32FromFloat32(0) {
			// Miss
			return output
		}
		// Advance sweep
		fraction = fraction + (target-distanceOutput.Distance)/denominator
		if fraction >= (*ShapeCastPairInput)(unsafe.Pointer(input)).MaxFraction {
			// Miss
			return output
		}
		v54 = (*ShapeCastPairInput)(unsafe.Pointer(input)).TransformB.P
		v55 = fraction
		v56 = delta2
		v57 = Vec2{
			X: v54.X + float32(v55*v56.X),
			Y: v54.Y + float32(v55*v56.Y),
		}
		goto _58
	_58:
		(*(*DistanceInput)(unsafe.Pointer(bp + 8))).TransformB.P = v57
		goto _6
	_6:
		;
		iteration = iteration + 1
	}
	// Failure!
	return output
}

func b2ShapeCastCircle(tls *_Stack, input uintptr, shape uintptr) (r CastOutput) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var output CastOutput
	var _ /* pairInput at bp+0 */ ShapeCastPairInput
	_ = output
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).ProxyA = b2MakeProxy(tls, shape, int32(1), (*Circle)(unsafe.Pointer(shape)).Radius)
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).ProxyB = (*ShapeCastInput)(unsafe.Pointer(input)).Proxy
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).TransformA = b2Transform_identity
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).TransformB = b2Transform_identity
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).TranslationB = (*ShapeCastInput)(unsafe.Pointer(input)).Translation
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).MaxFraction = (*ShapeCastInput)(unsafe.Pointer(input)).MaxFraction
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).CanEncroach = (*ShapeCastInput)(unsafe.Pointer(input)).CanEncroach
	output = b2ShapeCast(tls, bp)
	return output
}

func b2ShapeCastCapsule(tls *_Stack, input uintptr, shape uintptr) (r CastOutput) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var output CastOutput
	var _ /* pairInput at bp+0 */ ShapeCastPairInput
	_ = output
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).ProxyA = b2MakeProxy(tls, shape, int32(2), (*Capsule)(unsafe.Pointer(shape)).Radius)
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).ProxyB = (*ShapeCastInput)(unsafe.Pointer(input)).Proxy
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).TransformA = b2Transform_identity
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).TransformB = b2Transform_identity
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).TranslationB = (*ShapeCastInput)(unsafe.Pointer(input)).Translation
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).MaxFraction = (*ShapeCastInput)(unsafe.Pointer(input)).MaxFraction
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).CanEncroach = (*ShapeCastInput)(unsafe.Pointer(input)).CanEncroach
	output = b2ShapeCast(tls, bp)
	return output
}

func b2ShapeCastSegment(tls *_Stack, input uintptr, shape uintptr) (r CastOutput) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var output CastOutput
	var _ /* pairInput at bp+0 */ ShapeCastPairInput
	_ = output
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).ProxyA = b2MakeProxy(tls, shape, int32(2), float32FromFloat32(0))
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).ProxyB = (*ShapeCastInput)(unsafe.Pointer(input)).Proxy
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).TransformA = b2Transform_identity
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).TransformB = b2Transform_identity
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).TranslationB = (*ShapeCastInput)(unsafe.Pointer(input)).Translation
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).MaxFraction = (*ShapeCastInput)(unsafe.Pointer(input)).MaxFraction
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).CanEncroach = (*ShapeCastInput)(unsafe.Pointer(input)).CanEncroach
	output = b2ShapeCast(tls, bp)
	return output
}

func b2ShapeCastPolygon(tls *_Stack, input uintptr, shape uintptr) (r CastOutput) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var output CastOutput
	var _ /* pairInput at bp+0 */ ShapeCastPairInput
	_ = output
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).ProxyA = b2MakeProxy(tls, shape, (*Polygon)(unsafe.Pointer(shape)).Count, (*Polygon)(unsafe.Pointer(shape)).Radius)
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).ProxyB = (*ShapeCastInput)(unsafe.Pointer(input)).Proxy
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).TransformA = b2Transform_identity
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).TransformB = b2Transform_identity
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).TranslationB = (*ShapeCastInput)(unsafe.Pointer(input)).Translation
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).MaxFraction = (*ShapeCastInput)(unsafe.Pointer(input)).MaxFraction
	(*(*ShapeCastPairInput)(unsafe.Pointer(bp))).CanEncroach = (*ShapeCastInput)(unsafe.Pointer(input)).CanEncroach
	output = b2ShapeCast(tls, bp)
	return output
}

func b2ShapeRefArray_Create(tls *_Stack, capacity int32) (r b2ShapeRefArray) {
	var a b2ShapeRefArray
	_ = a
	a = b2ShapeRefArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(8)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2ShapeRefArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2ShapeRefArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2ShapeRefArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2ShapeRefArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2ShapeRefArray)(unsafe.Pointer(a)).Capacity)*uint64(8)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(8)))
	(*b2ShapeRefArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2ShapeRefArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2ShapeRefArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2ShapeRefArray)(unsafe.Pointer(a)).Capacity)*uint64(8)))
	(*b2ShapeRefArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2ShapeRefArray)(unsafe.Pointer(a)).Count = 0
	(*b2ShapeRefArray)(unsafe.Pointer(a)).Capacity = 0
}

func b2CompareShapeRefs(tls *_Stack, a uintptr, b uintptr) (r int32) {
	var sa, sb uintptr
	_, _ = sa, sb
	sa = a
	sb = b
	if (*b2ShapeRef)(unsafe.Pointer(sa)).ShapeId < (*b2ShapeRef)(unsafe.Pointer(sb)).ShapeId {
		return -int32(1)
	}
	if (*b2ShapeRef)(unsafe.Pointer(sa)).ShapeId == (*b2ShapeRef)(unsafe.Pointer(sb)).ShapeId {
		if int32FromUint16((*b2ShapeRef)(unsafe.Pointer(sa)).Generation) < int32FromUint16((*b2ShapeRef)(unsafe.Pointer(sb)).Generation) {
			return -int32(1)
		}
		if int32FromUint16((*b2ShapeRef)(unsafe.Pointer(sa)).Generation) == int32FromUint16((*b2ShapeRef)(unsafe.Pointer(sb)).Generation) {
			return 0
		}
	}
	return int32(1)
}

/**@}*/

/**@}*/

func b2ChainShapeArray_Create(tls *_Stack, capacity int32) (r b2ChainShapeArray) {
	var a b2ChainShapeArray
	_ = a
	a = b2ChainShapeArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(48)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2ChainShapeArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2ChainShapeArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2ChainShapeArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2ChainShapeArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2ChainShapeArray)(unsafe.Pointer(a)).Capacity)*uint64(48)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(48)))
	(*b2ChainShapeArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2ChainShapeArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2ChainShapeArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2ChainShapeArray)(unsafe.Pointer(a)).Capacity)*uint64(48)))
	(*b2ChainShapeArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2ChainShapeArray)(unsafe.Pointer(a)).Count = 0
	(*b2ChainShapeArray)(unsafe.Pointer(a)).Capacity = 0
}

func b2ShapeArray_Create(tls *_Stack, capacity int32) (r b2ShapeArray) {
	var a b2ShapeArray
	_ = a
	a = b2ShapeArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(288)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2ShapeArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2ShapeArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2ShapeArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2ShapeArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2ShapeArray)(unsafe.Pointer(a)).Capacity)*uint64(288)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(288)))
	(*b2ShapeArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2ShapeArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2ShapeArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2ShapeArray)(unsafe.Pointer(a)).Capacity)*uint64(288)))
	(*b2ShapeArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2ShapeArray)(unsafe.Pointer(a)).Count = 0
	(*b2ShapeArray)(unsafe.Pointer(a)).Capacity = 0
}

func b2GetShape(tls *_Stack, world uintptr, shapeId ShapeId) (r uintptr) {
	var id, v2 int32
	var shape, v1, v3 uintptr
	_, _, _, _, _ = id, shape, v1, v2, v3
	id = shapeId.Index1 - int32(1)
	v1 = world + 1248
	v2 = id
	if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
	goto _4
_4:
	shape = v3
	if !((*b2Shape)(unsafe.Pointer(shape)).Id == id && int32FromUint16((*b2Shape)(unsafe.Pointer(shape)).Generation) == int32FromUint16(shapeId.Generation)) && b2InternalAssertFcn(tls, __ccgo_ts+10978, __ccgo_ts+11037, int32FromInt32(24)) != 0 {
		__builtin_trap(tls)
	}
	return shape
}

func b2GetChainShape(tls *_Stack, world uintptr, chainId ChainId) (r uintptr) {
	var chain, v1, v3 uintptr
	var id, v2 int32
	_, _, _, _, _ = chain, id, v1, v2, v3
	id = chainId.Index1 - int32(1)
	v1 = world + 1264
	v2 = id
	if !(0 <= v2 && v2 < (*b2ChainShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(137)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2ChainShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*48
	goto _4
_4:
	chain = v3
	if !((*b2ChainShape)(unsafe.Pointer(chain)).Id == id && int32FromUint16((*b2ChainShape)(unsafe.Pointer(chain)).Generation) == int32FromUint16(chainId.Generation)) && b2InternalAssertFcn(tls, __ccgo_ts+11060, __ccgo_ts+11037, int32FromInt32(32)) != 0 {
		__builtin_trap(tls)
	}
	return chain
}

func b2UpdateShapeAABBs(tls *_Stack, shape uintptr, transform Transform, proxyType BodyType) {
	var aabb, fatAABB AABB
	var aabbMargin, margin, speculativeDistance, v1 float32
	_, _, _, _, _, _ = aabb, aabbMargin, fatAABB, margin, speculativeDistance, v1
	// Compute a bounding box with a speculative margin
	speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter))
	aabbMargin = float32(float32FromFloat32(0.05) * b2_lengthUnitsPerMeter)
	aabb = b2ComputeShapeAABB(tls, shape, transform)
	aabb.LowerBound.X -= speculativeDistance
	aabb.LowerBound.Y -= speculativeDistance
	aabb.UpperBound.X += speculativeDistance
	aabb.UpperBound.Y += speculativeDistance
	(*b2Shape)(unsafe.Pointer(shape)).Aabb = aabb
	if proxyType == int32(b2_staticBody) {
		v1 = speculativeDistance
	} else {
		v1 = aabbMargin
	}
	// Smaller margin for static bodies. Cannot be zero due to TOI tolerance.
	margin = v1
	fatAABB.LowerBound.X = aabb.LowerBound.X - margin
	fatAABB.LowerBound.Y = aabb.LowerBound.Y - margin
	fatAABB.UpperBound.X = aabb.UpperBound.X + margin
	fatAABB.UpperBound.Y = aabb.UpperBound.Y + margin
	(*b2Shape)(unsafe.Pointer(shape)).FatAABB = fatAABB
}

func b2CreateShapeInternal(tls *_Stack, world uintptr, body uintptr, transform Transform, def uintptr, geometry uintptr, shapeType ShapeType) (r uintptr) {
	var headShape, shape, v1, v18, v20, v22, v3, v5, p17 uintptr
	var newCapacity, newCapacity1, shapeId, v19, v2, v23, v4 int32
	var proxyType b2BodyType1
	var sensor b2Sensor
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = headShape, newCapacity, newCapacity1, proxyType, sensor, shape, shapeId, v1, v18, v19, v2, v20, v22, v23, v3, v4, v5, p17
	shapeId = b2AllocId(tls, world+1200)
	if shapeId == (*b2World)(unsafe.Pointer(world)).Shapes.Count {
		v1 = world + 1248
		if (*b2ShapeArray)(unsafe.Pointer(v1)).Count == (*b2ShapeArray)(unsafe.Pointer(v1)).Capacity {
			if (*b2ShapeArray)(unsafe.Pointer(v1)).Capacity < int32(2) {
				v2 = int32(2)
			} else {
				v2 = (*b2ShapeArray)(unsafe.Pointer(v1)).Capacity + (*b2ShapeArray)(unsafe.Pointer(v1)).Capacity>>int32(1)
			}
			newCapacity = v2
			b2ShapeArray_Reserve(tls, v1, newCapacity)
		}
		*(*b2Shape)(unsafe.Pointer((*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr((*b2ShapeArray)(unsafe.Pointer(v1)).Count)*288)) = b2Shape{}
		*(*int32)(unsafe.Pointer(v1 + 8)) += int32(1)
	} else {
		if !((*(*b2Shape)(unsafe.Pointer((*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*288))).Id == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+11119, __ccgo_ts+11037, int32FromInt32(70)) != 0 {
			__builtin_trap(tls)
		}
	}
	v3 = world + 1248
	v4 = shapeId
	if !(0 <= v4 && v4 < (*b2ShapeArray)(unsafe.Pointer(v3)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
		__builtin_trap(tls)
	}
	v5 = (*b2ShapeArray)(unsafe.Pointer(v3)).Data + uintptr(v4)*288
	goto _6
_6:
	shape = v5
	switch shapeType {
	case int32(b2_capsuleShape):
		goto _7
	case int32(b2_circleShape):
		goto _8
	case int32(b2_polygonShape):
		goto _9
	case int32(b2_segmentShape):
		goto _10
	case int32(b2_chainSegmentShape):
		goto _11
	default:
		goto _12
	}
	goto _13
_7:
	;
	(*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo19_132.Capsule = *(*Capsule)(unsafe.Pointer(geometry))
	goto _13
_8:
	;
	*(*Circle)(unsafe.Add(unsafe.Pointer(shape), 132)) = *(*Circle)(unsafe.Pointer(geometry))
	goto _13
_9:
	;
	*(*Polygon)(unsafe.Add(unsafe.Pointer(shape), 132)) = *(*Polygon)(unsafe.Pointer(geometry))
	goto _13
_10:
	;
	*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 132)) = *(*Segment)(unsafe.Pointer(geometry))
	goto _13
_11:
	;
	*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 132)) = *(*ChainSegment)(unsafe.Pointer(geometry))
	goto _13
_12:
	;
_16:
	;
	if bool(!(int32FromInt32(false1) != 0)) && b2InternalAssertFcn(tls, __ccgo_ts+4123, __ccgo_ts+11037, int32FromInt32(98)) != 0 {
		__builtin_trap(tls)
	}
	goto _15
_15:
	;
	if 0 != 0 {
		goto _16
	}
	goto _14
_14:
	;
	goto _13
_13:
	;
	(*b2Shape)(unsafe.Pointer(shape)).Id = shapeId
	(*b2Shape)(unsafe.Pointer(shape)).BodyId = (*b2Body)(unsafe.Pointer(body)).Id
	(*b2Shape)(unsafe.Pointer(shape)).Type1 = shapeType
	(*b2Shape)(unsafe.Pointer(shape)).Density = (*ShapeDef)(unsafe.Pointer(def)).Density
	(*b2Shape)(unsafe.Pointer(shape)).Friction = (*ShapeDef)(unsafe.Pointer(def)).Material.Friction
	(*b2Shape)(unsafe.Pointer(shape)).Restitution = (*ShapeDef)(unsafe.Pointer(def)).Material.Restitution
	(*b2Shape)(unsafe.Pointer(shape)).RollingResistance = (*ShapeDef)(unsafe.Pointer(def)).Material.RollingResistance
	(*b2Shape)(unsafe.Pointer(shape)).TangentSpeed = (*ShapeDef)(unsafe.Pointer(def)).Material.TangentSpeed
	(*b2Shape)(unsafe.Pointer(shape)).UserMaterialId = (*ShapeDef)(unsafe.Pointer(def)).Material.UserMaterialId
	(*b2Shape)(unsafe.Pointer(shape)).Filter = (*ShapeDef)(unsafe.Pointer(def)).Filter
	(*b2Shape)(unsafe.Pointer(shape)).UserData = (*ShapeDef)(unsafe.Pointer(def)).UserData
	(*b2Shape)(unsafe.Pointer(shape)).CustomColor = (*ShapeDef)(unsafe.Pointer(def)).Material.CustomColor
	(*b2Shape)(unsafe.Pointer(shape)).EnlargedAABB = boolUint8(false1 != 0)
	(*b2Shape)(unsafe.Pointer(shape)).EnableSensorEvents = (*ShapeDef)(unsafe.Pointer(def)).EnableSensorEvents
	(*b2Shape)(unsafe.Pointer(shape)).EnableContactEvents = (*ShapeDef)(unsafe.Pointer(def)).EnableContactEvents
	(*b2Shape)(unsafe.Pointer(shape)).EnableHitEvents = (*ShapeDef)(unsafe.Pointer(def)).EnableHitEvents
	(*b2Shape)(unsafe.Pointer(shape)).EnablePreSolveEvents = (*ShapeDef)(unsafe.Pointer(def)).EnablePreSolveEvents
	(*b2Shape)(unsafe.Pointer(shape)).ProxyKey = -int32(1)
	(*b2Shape)(unsafe.Pointer(shape)).LocalCentroid = b2GetShapeCentroid(tls, shape)
	(*b2Shape)(unsafe.Pointer(shape)).Aabb = AABB{
		LowerBound: b2Vec2_zero,
		UpperBound: b2Vec2_zero,
	}
	(*b2Shape)(unsafe.Pointer(shape)).FatAABB = AABB{
		LowerBound: b2Vec2_zero,
		UpperBound: b2Vec2_zero,
	}
	p17 = shape + 276
	*(*uint16_t)(unsafe.Pointer(p17)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(p17))) + int32FromInt32(1))
	if (*b2Body)(unsafe.Pointer(body)).SetIndex != int32(b2_disabledSet) {
		proxyType = (*b2Body)(unsafe.Pointer(body)).Type1
		b2CreateShapeProxy(tls, shape, world+40, proxyType, transform, boolUint8((*ShapeDef)(unsafe.Pointer(def)).InvokeContactCreation != 0 || (*ShapeDef)(unsafe.Pointer(def)).IsSensor != 0))
	}
	// Add to shape doubly linked list
	if (*b2Body)(unsafe.Pointer(body)).HeadShapeId != -int32(1) {
		v18 = world + 1248
		v19 = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
		if !(0 <= v19 && v19 < (*b2ShapeArray)(unsafe.Pointer(v18)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v20 = (*b2ShapeArray)(unsafe.Pointer(v18)).Data + uintptr(v19)*288
		goto _21
	_21:
		headShape = v20
		(*b2Shape)(unsafe.Pointer(headShape)).PrevShapeId = shapeId
	}
	(*b2Shape)(unsafe.Pointer(shape)).PrevShapeId = -int32(1)
	(*b2Shape)(unsafe.Pointer(shape)).NextShapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	(*b2Body)(unsafe.Pointer(body)).HeadShapeId = shapeId
	*(*int32)(unsafe.Pointer(body + 60)) += int32(1)
	if (*ShapeDef)(unsafe.Pointer(def)).IsSensor != 0 {
		(*b2Shape)(unsafe.Pointer(shape)).SensorIndex = (*b2World)(unsafe.Pointer(world)).Sensors.Count
		sensor = b2Sensor{
			Overlaps1: b2ShapeRefArray_Create(tls, int32(16)),
			Overlaps2: b2ShapeRefArray_Create(tls, int32(16)),
			ShapeId:   shapeId,
		}
		v22 = world + 1280
		if (*b2SensorArray)(unsafe.Pointer(v22)).Count == (*b2SensorArray)(unsafe.Pointer(v22)).Capacity {
			if (*b2SensorArray)(unsafe.Pointer(v22)).Capacity < int32(2) {
				v23 = int32(2)
			} else {
				v23 = (*b2SensorArray)(unsafe.Pointer(v22)).Capacity + (*b2SensorArray)(unsafe.Pointer(v22)).Capacity>>int32(1)
			}
			newCapacity1 = v23
			b2SensorArray_Reserve(tls, v22, newCapacity1)
		}
		*(*b2Sensor)(unsafe.Pointer((*b2SensorArray)(unsafe.Pointer(v22)).Data + uintptr((*b2SensorArray)(unsafe.Pointer(v22)).Count)*40)) = sensor
		*(*int32)(unsafe.Pointer(v22 + 8)) += int32(1)
	} else {
		(*b2Shape)(unsafe.Pointer(shape)).SensorIndex = -int32(1)
	}
	b2ValidateSolverSets(tls, world)
	return shape
}

func b2CreateShape(tls *_Stack, bodyId BodyId, def uintptr, geometry uintptr, shapeType ShapeType) (r ShapeId) {
	var body, shape, world uintptr
	var id ShapeId
	var transform Transform
	_, _, _, _, _ = body, id, shape, transform, world
	if !((*ShapeDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B2_SECRET_COOKIE)) && b2InternalAssertFcn(tls, __ccgo_ts+730, __ccgo_ts+11037, int32FromInt32(165)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2IsValidFloat(tls, (*ShapeDef)(unsafe.Pointer(def)).Density) != 0 && (*ShapeDef)(unsafe.Pointer(def)).Density >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+11167, __ccgo_ts+11037, int32FromInt32(166)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2IsValidFloat(tls, (*ShapeDef)(unsafe.Pointer(def)).Material.Friction) != 0 && (*ShapeDef)(unsafe.Pointer(def)).Material.Friction >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+11222, __ccgo_ts+11037, int32FromInt32(167)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2IsValidFloat(tls, (*ShapeDef)(unsafe.Pointer(def)).Material.Restitution) != 0 && (*ShapeDef)(unsafe.Pointer(def)).Material.Restitution >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+11297, __ccgo_ts+11037, int32FromInt32(168)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2IsValidFloat(tls, (*ShapeDef)(unsafe.Pointer(def)).Material.RollingResistance) != 0 && (*ShapeDef)(unsafe.Pointer(def)).Material.RollingResistance >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+11378, __ccgo_ts+11037, int32FromInt32(169)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2IsValidFloat(tls, (*ShapeDef)(unsafe.Pointer(def)).Material.TangentSpeed) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+11471, __ccgo_ts+11037, int32FromInt32(170)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return ShapeId{}
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	transform = b2GetBodyTransformQuick(tls, world, body)
	shape = b2CreateShapeInternal(tls, world, body, transform, def, geometry, shapeType)
	if int32FromUint8((*ShapeDef)(unsafe.Pointer(def)).UpdateBodyMass) == int32(true1) {
		b2UpdateBodyMassData(tls, world, body)
	}
	b2ValidateSolverSets(tls, world)
	id = ShapeId{
		Index1:     (*b2Shape)(unsafe.Pointer(shape)).Id + int32(1),
		World0:     bodyId.World0,
		Generation: (*b2Shape)(unsafe.Pointer(shape)).Generation,
	}
	return id
}

func b2CreateCircleShape(tls *_Stack, bodyId BodyId, def uintptr, circle uintptr) (r ShapeId) {
	return b2CreateShape(tls, bodyId, def, circle, int32(b2_circleShape))
}

func b2CreateCapsuleShape(tls *_Stack, bodyId BodyId, def uintptr, capsule uintptr) (r ShapeId) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var c, v1, v2, v5, v6, v8 Vec2
	var lengthSqr, v3, v7 float32
	var _ /* circle at bp+0 */ Circle
	_, _, _, _, _, _, _, _, _ = c, lengthSqr, v1, v2, v3, v5, v6, v7, v8
	v1 = (*Capsule)(unsafe.Pointer(capsule)).Center1
	v2 = (*Capsule)(unsafe.Pointer(capsule)).Center2
	c = Vec2{
		X: v2.X - v1.X,
		Y: v2.Y - v1.Y,
	}
	v3 = float32(c.X*c.X) + float32(c.Y*c.Y)
	goto _4
_4:
	lengthSqr = v3
	if lengthSqr <= float32(float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
		v5 = (*Capsule)(unsafe.Pointer(capsule)).Center1
		v6 = (*Capsule)(unsafe.Pointer(capsule)).Center2
		v7 = float32FromFloat32(0.5)
		v8 = Vec2{
			X: float32((float32FromFloat32(1)-v7)*v5.X) + float32(v7*v6.X),
			Y: float32((float32FromFloat32(1)-v7)*v5.Y) + float32(v7*v6.Y),
		}
		goto _9
	_9:
		*(*Circle)(unsafe.Pointer(bp)) = Circle{
			Center: v8,
			Radius: (*Capsule)(unsafe.Pointer(capsule)).Radius,
		}
		return b2CreateShape(tls, bodyId, def, bp, int32(b2_circleShape))
	}
	return b2CreateShape(tls, bodyId, def, capsule, int32(b2_capsuleShape))
}

func b2CreatePolygonShape(tls *_Stack, bodyId BodyId, def uintptr, polygon uintptr) (r ShapeId) {
	if !(b2IsValidFloat(tls, (*Polygon)(unsafe.Pointer(polygon)).Radius) != 0 && (*Polygon)(unsafe.Pointer(polygon)).Radius >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+11516, __ccgo_ts+11037, int32FromInt32(213)) != 0 {
		__builtin_trap(tls)
	}
	return b2CreateShape(tls, bodyId, def, polygon, int32(b2_polygonShape))
}

func b2CreateSegmentShape(tls *_Stack, bodyId BodyId, def uintptr, segment uintptr) (r ShapeId) {
	var c, v1, v2 Vec2
	var lengthSqr, v3 float32
	_, _, _, _, _ = c, lengthSqr, v1, v2, v3
	v1 = (*Segment)(unsafe.Pointer(segment)).Point1
	v2 = (*Segment)(unsafe.Pointer(segment)).Point2
	c = Vec2{
		X: v2.X - v1.X,
		Y: v2.Y - v1.Y,
	}
	v3 = float32(c.X*c.X) + float32(c.Y*c.Y)
	goto _4
_4:
	lengthSqr = v3
	if lengthSqr <= float32(float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
		if bool(!(int32FromInt32(false1) != 0)) && b2InternalAssertFcn(tls, __ccgo_ts+4123, __ccgo_ts+11037, int32FromInt32(222)) != 0 {
			__builtin_trap(tls)
		}
		return b2_nullShapeId
	}
	return b2CreateShape(tls, bodyId, def, segment, int32(b2_segmentShape))
}

// C documentation
//
//	// Destroy a shape on a body. This doesn't need to be called when destroying a body.
func b2DestroyShapeInternal(tls *_Stack, world uintptr, shape uintptr, body uintptr, wakeBodies uint8) {
	var contact, movedSensor, nextShape, otherSensorShape, prevShape, ref, sensor, v1, v11, v13, v15, v18, v20, v24, v26, v28, v3, v30, v5, v7, v9 uintptr
	var contactId, contactKey, edgeIndex, i, movedIndex, movedIndex1, newCapacity, shapeId, v10, v14, v19, v2, v21, v22, v25, v29, v6 int32
	var event SensorEndTouchEvent
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = contact, contactId, contactKey, edgeIndex, event, i, movedIndex, movedIndex1, movedSensor, newCapacity, nextShape, otherSensorShape, prevShape, ref, sensor, shapeId, v1, v10, v11, v13, v14, v15, v18, v19, v2, v20, v21, v22, v24, v25, v26, v28, v29, v3, v30, v5, v6, v7, v9
	shapeId = (*b2Shape)(unsafe.Pointer(shape)).Id
	// Remove the shape from the body's doubly linked list.
	if (*b2Shape)(unsafe.Pointer(shape)).PrevShapeId != -int32(1) {
		v1 = world + 1248
		v2 = (*b2Shape)(unsafe.Pointer(shape)).PrevShapeId
		if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
		goto _4
	_4:
		prevShape = v3
		(*b2Shape)(unsafe.Pointer(prevShape)).NextShapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
	}
	if (*b2Shape)(unsafe.Pointer(shape)).NextShapeId != -int32(1) {
		v5 = world + 1248
		v6 = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
		if !(0 <= v6 && v6 < (*b2ShapeArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v7 = (*b2ShapeArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*288
		goto _8
	_8:
		nextShape = v7
		(*b2Shape)(unsafe.Pointer(nextShape)).PrevShapeId = (*b2Shape)(unsafe.Pointer(shape)).PrevShapeId
	}
	if shapeId == (*b2Body)(unsafe.Pointer(body)).HeadShapeId {
		(*b2Body)(unsafe.Pointer(body)).HeadShapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
	}
	*(*int32)(unsafe.Pointer(body + 60)) -= int32(1)
	// Remove from broad-phase.
	b2DestroyShapeProxy(tls, shape, world+40)
	// Destroy any contacts associated with the shape.
	contactKey = (*b2Body)(unsafe.Pointer(body)).HeadContactKey
	for contactKey != -int32(1) {
		contactId = contactKey >> int32(1)
		edgeIndex = contactKey & int32(1)
		v9 = world + 1144
		v10 = contactId
		if !(0 <= v10 && v10 < (*b2ContactArray)(unsafe.Pointer(v9)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
			__builtin_trap(tls)
		}
		v11 = (*b2ContactArray)(unsafe.Pointer(v9)).Data + uintptr(v10)*68
		goto _12
	_12:
		contact = v11
		contactKey = (*(*b2ContactEdge)(unsafe.Pointer(contact + 12 + uintptr(edgeIndex)*12))).NextKey
		if (*b2Contact)(unsafe.Pointer(contact)).ShapeIdA == shapeId || (*b2Contact)(unsafe.Pointer(contact)).ShapeIdB == shapeId {
			b2DestroyContact(tls, world, contact, wakeBodies)
		}
	}
	if (*b2Shape)(unsafe.Pointer(shape)).SensorIndex != -int32(1) {
		v13 = world + 1280
		v14 = (*b2Shape)(unsafe.Pointer(shape)).SensorIndex
		if !(0 <= v14 && v14 < (*b2SensorArray)(unsafe.Pointer(v13)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+10835, int32FromInt32(35)) != 0 {
			__builtin_trap(tls)
		}
		v15 = (*b2SensorArray)(unsafe.Pointer(v13)).Data + uintptr(v14)*40
		goto _16
	_16:
		sensor = v15
		i = 0
		for {
			if !(i < (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count) {
				break
			}
			ref = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data + uintptr(i)*8
			event = SensorEndTouchEvent{
				SensorShapeId: ShapeId{
					Index1:     shapeId + int32(1),
					World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
					Generation: (*b2Shape)(unsafe.Pointer(shape)).Generation,
				},
				VisitorShapeId: ShapeId{
					Index1:     (*b2ShapeRef)(unsafe.Pointer(ref)).ShapeId + int32(1),
					World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
					Generation: (*b2ShapeRef)(unsafe.Pointer(ref)).Generation,
				},
			}
			v18 = world + 1376 + uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16
			if (*b2SensorEndTouchEventArray)(unsafe.Pointer(v18)).Count == (*b2SensorEndTouchEventArray)(unsafe.Pointer(v18)).Capacity {
				if (*b2SensorEndTouchEventArray)(unsafe.Pointer(v18)).Capacity < int32(2) {
					v19 = int32(2)
				} else {
					v19 = (*b2SensorEndTouchEventArray)(unsafe.Pointer(v18)).Capacity + (*b2SensorEndTouchEventArray)(unsafe.Pointer(v18)).Capacity>>int32(1)
				}
				newCapacity = v19
				b2SensorEndTouchEventArray_Reserve(tls, v18, newCapacity)
			}
			*(*SensorEndTouchEvent)(unsafe.Pointer((*b2SensorEndTouchEventArray)(unsafe.Pointer(v18)).Data + uintptr((*b2SensorEndTouchEventArray)(unsafe.Pointer(v18)).Count)*16)) = event
			*(*int32)(unsafe.Pointer(v18 + 8)) += int32(1)
			goto _17
		_17:
			;
			i = i + 1
		}
		// Destroy sensor
		b2ShapeRefArray_Destroy(tls, sensor)
		b2ShapeRefArray_Destroy(tls, sensor+16)
		v20 = world + 1280
		v21 = (*b2Shape)(unsafe.Pointer(shape)).SensorIndex
		if !(0 <= v21 && v21 < (*b2SensorArray)(unsafe.Pointer(v20)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+10835, int32FromInt32(35)) != 0 {
			__builtin_trap(tls)
		}
		movedIndex = -int32(1)
		if v21 != (*b2SensorArray)(unsafe.Pointer(v20)).Count-int32FromInt32(1) {
			movedIndex = (*b2SensorArray)(unsafe.Pointer(v20)).Count - int32(1)
			*(*b2Sensor)(unsafe.Pointer((*b2SensorArray)(unsafe.Pointer(v20)).Data + uintptr(v21)*40)) = *(*b2Sensor)(unsafe.Pointer((*b2SensorArray)(unsafe.Pointer(v20)).Data + uintptr(movedIndex)*40))
		}
		*(*int32)(unsafe.Pointer(v20 + 8)) -= int32(1)
		v22 = movedIndex
		goto _23
	_23:
		movedIndex1 = v22
		if movedIndex1 != -int32(1) {
			v24 = world + 1280
			v25 = (*b2Shape)(unsafe.Pointer(shape)).SensorIndex
			if !(0 <= v25 && v25 < (*b2SensorArray)(unsafe.Pointer(v24)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+10835, int32FromInt32(35)) != 0 {
				__builtin_trap(tls)
			}
			v26 = (*b2SensorArray)(unsafe.Pointer(v24)).Data + uintptr(v25)*40
			goto _27
		_27:
			// Fixup moved sensor
			movedSensor = v26
			v28 = world + 1248
			v29 = (*b2Sensor)(unsafe.Pointer(movedSensor)).ShapeId
			if !(0 <= v29 && v29 < (*b2ShapeArray)(unsafe.Pointer(v28)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
				__builtin_trap(tls)
			}
			v30 = (*b2ShapeArray)(unsafe.Pointer(v28)).Data + uintptr(v29)*288
			goto _31
		_31:
			otherSensorShape = v30
			(*b2Shape)(unsafe.Pointer(otherSensorShape)).SensorIndex = (*b2Shape)(unsafe.Pointer(shape)).SensorIndex
		}
	}
	// Return shape to free list.
	b2FreeId(tls, world+1200, shapeId)
	(*b2Shape)(unsafe.Pointer(shape)).Id = -int32(1)
	b2ValidateSolverSets(tls, world)
}

func b2DestroyShape(tls *_Stack, shapeId ShapeId, updateBodyMass uint8) {
	var body, shape, world, v1, v3 uintptr
	var wakeBodies uint8
	var v2 int32
	_, _, _, _, _, _, _ = body, shape, wakeBodies, world, v1, v2, v3
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	shape = b2GetShape(tls, world, shapeId)
	// need to wake bodies because this might be a static body
	wakeBodies = boolUint8(true1 != 0)
	v1 = world + 1024
	v2 = (*b2Shape)(unsafe.Pointer(shape)).BodyId
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	body = v3
	b2DestroyShapeInternal(tls, world, shape, body, wakeBodies)
	if int32FromUint8(updateBodyMass) == int32(true1) {
		b2UpdateBodyMassData(tls, world, body)
	}
}

func b2ComputeShapeAABB(tls *_Stack, shape uintptr, xf Transform) (r AABB) {
	var empty AABB
	_ = empty
	switch (*b2Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b2_capsuleShape):
		return b2ComputeCapsuleAABB(tls, shape+132, xf)
	case int32(b2_circleShape):
		return b2ComputeCircleAABB(tls, shape+132, xf)
	case int32(b2_polygonShape):
		return b2ComputePolygonAABB(tls, shape+132, xf)
	case int32(b2_segmentShape):
		return b2ComputeSegmentAABB(tls, shape+132, xf)
	case int32(b2_chainSegmentShape):
		return b2ComputeSegmentAABB(tls, shape+132+8, xf)
	default:
		if bool(!(int32FromInt32(false1) != 0)) && b2InternalAssertFcn(tls, __ccgo_ts+4123, __ccgo_ts+11037, int32FromInt32(594)) != 0 {
			__builtin_trap(tls)
		}
		empty = AABB{
			LowerBound: xf.P,
			UpperBound: xf.P,
		}
		return empty
	}
	return r
}

func b2GetShapeCentroid(tls *_Stack, shape uintptr) (r Vec2) {
	var v1, v11, v12, v14, v2, v4, v6, v7, v9 Vec2
	var v13, v3, v8 float32
	_, _, _, _, _, _, _, _, _, _, _, _ = v1, v11, v12, v13, v14, v2, v3, v4, v6, v7, v8, v9
	switch (*b2Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b2_capsuleShape):
		v1 = (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo19_132.Capsule.Center1
		v2 = (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo19_132.Capsule.Center2
		v3 = float32FromFloat32(0.5)
		v4 = Vec2{
			X: float32((float32FromFloat32(1)-v3)*v1.X) + float32(v3*v2.X),
			Y: float32((float32FromFloat32(1)-v3)*v1.Y) + float32(v3*v2.Y),
		}
		goto _5
	_5:
		return v4
	case int32(b2_circleShape):
		return (*(*Circle)(unsafe.Add(unsafe.Pointer(shape), 132))).Center
	case int32(b2_polygonShape):
		return (*(*Polygon)(unsafe.Add(unsafe.Pointer(shape), 132))).Centroid
	case int32(b2_segmentShape):
		v6 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 132))).Point1
		v7 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 132))).Point2
		v8 = float32FromFloat32(0.5)
		v9 = Vec2{
			X: float32((float32FromFloat32(1)-v8)*v6.X) + float32(v8*v7.X),
			Y: float32((float32FromFloat32(1)-v8)*v6.Y) + float32(v8*v7.Y),
		}
		goto _10
	_10:
		return v9
	case int32(b2_chainSegmentShape):
		v11 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 132))).Segment.Point1
		v12 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 132))).Segment.Point2
		v13 = float32FromFloat32(0.5)
		v14 = Vec2{
			X: float32((float32FromFloat32(1)-v13)*v11.X) + float32(v13*v12.X),
			Y: float32((float32FromFloat32(1)-v13)*v11.Y) + float32(v13*v12.Y),
		}
		goto _15
	_15:
		return v14
	default:
		return b2Vec2_zero
	}
	return r
}

// C documentation
//
//	// todo_erin maybe compute this on shape creation
func b2GetShapePerimeter(tls *_Stack, shape uintptr) (r float32) {
	var count, i int32
	var next, prev, v1, v10, v11, v13, v16, v17, v18, v2, v20, v23, v24, v25, v27, v3, v5, v9 Vec2
	var perimeter, v14, v21, v28, v6 float32
	var points uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = count, i, next, perimeter, points, prev, v1, v10, v11, v13, v14, v16, v17, v18, v2, v20, v21, v23, v24, v25, v27, v28, v3, v5, v6, v9
	switch (*b2Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b2_capsuleShape):
		v1 = (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo19_132.Capsule.Center1
		v2 = (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo19_132.Capsule.Center2
		v3 = Vec2{
			X: v1.X - v2.X,
			Y: v1.Y - v2.Y,
		}
		goto _4
	_4:
		v5 = v3
		v6 = sqrtf(tls, float32(v5.X*v5.X)+float32(v5.Y*v5.Y))
		goto _7
	_7:
		return float32(float32FromFloat32(2)*v6) + float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359))*(*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo19_132.Capsule.Radius)
	case int32(b2_circleShape):
		return float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * (*(*Circle)(unsafe.Add(unsafe.Pointer(shape), 132))).Radius)
	case int32(b2_polygonShape):
		points = shape + 132
		count = (*(*Polygon)(unsafe.Add(unsafe.Pointer(shape), 132))).Count
		perimeter = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * (*(*Polygon)(unsafe.Add(unsafe.Pointer(shape), 132))).Radius)
		if !(count > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+6294, __ccgo_ts+11037, int32FromInt32(635)) != 0 {
			__builtin_trap(tls)
		}
		prev = *(*Vec2)(unsafe.Pointer(points + uintptr(count-int32(1))*8))
		i = 0
		for {
			if !(i < count) {
				break
			}
			next = *(*Vec2)(unsafe.Pointer(points + uintptr(i)*8))
			v9 = next
			v10 = prev
			v11 = Vec2{
				X: v9.X - v10.X,
				Y: v9.Y - v10.Y,
			}
			goto _12
		_12:
			v13 = v11
			v14 = sqrtf(tls, float32(v13.X*v13.X)+float32(v13.Y*v13.Y))
			goto _15
		_15:
			perimeter = perimeter + v14
			prev = next
			goto _8
		_8:
			;
			i = i + 1
		}
		return perimeter
	case int32(b2_segmentShape):
		v16 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 132))).Point1
		v17 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 132))).Point2
		v18 = Vec2{
			X: v16.X - v17.X,
			Y: v16.Y - v17.Y,
		}
		goto _19
	_19:
		v20 = v18
		v21 = sqrtf(tls, float32(v20.X*v20.X)+float32(v20.Y*v20.Y))
		goto _22
	_22:
		return float32(float32FromFloat32(2) * v21)
	case int32(b2_chainSegmentShape):
		v23 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 132))).Segment.Point1
		v24 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 132))).Segment.Point2
		v25 = Vec2{
			X: v23.X - v24.X,
			Y: v23.Y - v24.Y,
		}
		goto _26
	_26:
		v27 = v25
		v28 = sqrtf(tls, float32(v27.X*v27.X)+float32(v27.Y*v27.Y))
		goto _29
	_29:
		return float32(float32FromFloat32(2) * v28)
	default:
		return float32FromFloat32(0)
	}
	return r
}

// C documentation
//
//	// This projects the shape perimeter onto an infinite line
func b2GetShapeProjectedPerimeter(tls *_Stack, shape uintptr, line Vec2) (r float32) {
	var axis, v1, v13, v14, v18, v19, v2, v3, v32, v33, v36, v37, v44, v45, v48, v49, v5, v6 Vec2
	var count, i int32
	var lower, projectedLength, upper, value, value1, value11, value2, value21, v10, v12, v15, v20, v22, v23, v24, v26, v27, v28, v29, v31, v34, v38, v40, v41, v43, v46, v50, v52, v53, v55, v7, v9 float32
	var points uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axis, count, i, lower, points, projectedLength, upper, value, value1, value11, value2, value21, v1, v10, v12, v13, v14, v15, v18, v19, v2, v20, v22, v23, v24, v26, v27, v28, v29, v3, v31, v32, v33, v34, v36, v37, v38, v40, v41, v43, v44, v45, v46, v48, v49, v5, v50, v52, v53, v55, v6, v7, v9
	switch (*b2Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b2_capsuleShape):
		v1 = (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo19_132.Capsule.Center2
		v2 = (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo19_132.Capsule.Center1
		v3 = Vec2{
			X: v1.X - v2.X,
			Y: v1.Y - v2.Y,
		}
		goto _4
	_4:
		axis = v3
		v5 = axis
		v6 = line
		v7 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y)
		goto _8
	_8:
		v9 = v7
		if v9 < float32FromInt32(0) {
			v12 = -v9
		} else {
			v12 = v9
		}
		v10 = v12
		goto _11
	_11:
		projectedLength = v10
		return projectedLength + float32(float32FromFloat32(2)*(*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo19_132.Capsule.Radius)
	case int32(b2_circleShape):
		return float32(float32FromFloat32(2) * (*(*Circle)(unsafe.Add(unsafe.Pointer(shape), 132))).Radius)
	case int32(b2_polygonShape):
		points = shape + 132
		count = (*(*Polygon)(unsafe.Add(unsafe.Pointer(shape), 132))).Count
		if !(count > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+6294, __ccgo_ts+11037, int32FromInt32(674)) != 0 {
			__builtin_trap(tls)
		}
		v13 = *(*Vec2)(unsafe.Pointer(points))
		v14 = line
		v15 = float32(v13.X*v14.X) + float32(v13.Y*v14.Y)
		goto _16
	_16:
		value = v15
		lower = value
		upper = value
		i = int32(1)
		for {
			if !(i < count) {
				break
			}
			v18 = *(*Vec2)(unsafe.Pointer(points + uintptr(i)*8))
			v19 = line
			v20 = float32(v18.X*v19.X) + float32(v18.Y*v19.Y)
			goto _21
		_21:
			value = v20
			v22 = lower
			v23 = value
			if v22 < v23 {
				v26 = v22
			} else {
				v26 = v23
			}
			v24 = v26
			goto _25
		_25:
			lower = v24
			v27 = upper
			v28 = value
			if v27 > v28 {
				v31 = v27
			} else {
				v31 = v28
			}
			v29 = v31
			goto _30
		_30:
			upper = v29
			goto _17
		_17:
			;
			i = i + 1
		}
		return upper - lower + float32(float32FromFloat32(2)*(*(*Polygon)(unsafe.Add(unsafe.Pointer(shape), 132))).Radius)
	case int32(b2_segmentShape):
		v32 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 132))).Point1
		v33 = line
		v34 = float32(v32.X*v33.X) + float32(v32.Y*v33.Y)
		goto _35
	_35:
		value1 = v34
		v36 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 132))).Point2
		v37 = line
		v38 = float32(v36.X*v37.X) + float32(v36.Y*v37.Y)
		goto _39
	_39:
		value2 = v38
		v40 = value2 - value1
		if v40 < float32FromInt32(0) {
			v43 = -v40
		} else {
			v43 = v40
		}
		v41 = v43
		goto _42
	_42:
		return v41
	case int32(b2_chainSegmentShape):
		v44 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 132))).Segment.Point1
		v45 = line
		v46 = float32(v44.X*v45.X) + float32(v44.Y*v45.Y)
		goto _47
	_47:
		value11 = v46
		v48 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 132))).Segment.Point2
		v49 = line
		v50 = float32(v48.X*v49.X) + float32(v48.Y*v49.Y)
		goto _51
	_51:
		value21 = v50
		v52 = value21 - value11
		if v52 < float32FromInt32(0) {
			v55 = -v52
		} else {
			v55 = v52
		}
		v53 = v55
		goto _54
	_54:
		return v53
	default:
		return float32FromFloat32(0)
	}
	return r
}

func b2ComputeShapeMass(tls *_Stack, shape uintptr) (r MassData) {
	switch (*b2Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b2_capsuleShape):
		return b2ComputeCapsuleMass(tls, shape+132, (*b2Shape)(unsafe.Pointer(shape)).Density)
	case int32(b2_circleShape):
		return b2ComputeCircleMass(tls, shape+132, (*b2Shape)(unsafe.Pointer(shape)).Density)
	case int32(b2_polygonShape):
		return b2ComputePolygonMass(tls, shape+132, (*b2Shape)(unsafe.Pointer(shape)).Density)
	default:
		return MassData{}
	}
	return r
}

func b2ComputeShapeExtent(tls *_Stack, shape uintptr, localCenter Vec2) (r b2ShapeExtent) {
	var c1, c11, c12, c2, c21, c22, v2, v1, v12, v21, v20, v211, v22, v24, v28, v29, v3, v30, v32, v33, v41, v42, v43, v45, v5, v53, v54, v55, v57, v58, v59, v6, v61, v64, v7, v72, v73, v74, v76, v77, v78, v80, v83, v9 Vec2
	var count, i int32
	var distanceSqr, maxExtentSqr, minExtent, planeOffset, radius, radius1, v10, v13, v15, v16, v17, v19, v25, v34, v36, v37, v38, v40, v46, v48, v49, v50, v52, v62, v65, v67, v68, v69, v71, v81, v84, v86, v87, v88, v90 float32
	var extent b2ShapeExtent
	var poly uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c1, c11, c12, c2, c21, c22, count, distanceSqr, extent, i, maxExtentSqr, minExtent, planeOffset, poly, radius, radius1, v2, v1, v10, v12, v13, v15, v16, v17, v19, v21, v20, v211, v22, v24, v25, v28, v29, v3, v30, v32, v33, v34, v36, v37, v38, v40, v41, v42, v43, v45, v46, v48, v49, v5, v50, v52, v53, v54, v55, v57, v58, v59, v6, v61, v62, v64, v65, v67, v68, v69, v7, v71, v72, v73, v74, v76, v77, v78, v80, v81, v83, v84, v86, v87, v88, v9, v90
	extent = b2ShapeExtent{}
	switch (*b2Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b2_capsuleShape):
		radius = (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo19_132.Capsule.Radius
		extent.MinExtent = radius
		v1 = (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo19_132.Capsule.Center1
		v21 = localCenter
		v3 = Vec2{
			X: v1.X - v21.X,
			Y: v1.Y - v21.Y,
		}
		goto _4
	_4:
		c1 = v3
		v5 = (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo19_132.Capsule.Center2
		v6 = localCenter
		v7 = Vec2{
			X: v5.X - v6.X,
			Y: v5.Y - v6.Y,
		}
		goto _8
	_8:
		c2 = v7
		v9 = c1
		v10 = float32(v9.X*v9.X) + float32(v9.Y*v9.Y)
		goto _11
	_11:
		v12 = c2
		v13 = float32(v12.X*v12.X) + float32(v12.Y*v12.Y)
		goto _14
	_14:
		v15 = v10
		v16 = v13
		if v15 > v16 {
			v19 = v15
		} else {
			v19 = v16
		}
		v17 = v19
		goto _18
	_18:
		extent.MaxExtent = sqrtf(tls, v17) + radius
	case int32(b2_circleShape):
		radius1 = (*(*Circle)(unsafe.Add(unsafe.Pointer(shape), 132))).Radius
		extent.MinExtent = radius1
		v20 = (*(*Circle)(unsafe.Add(unsafe.Pointer(shape), 132))).Center
		v211 = localCenter
		v22 = Vec2{
			X: v20.X - v211.X,
			Y: v20.Y - v211.Y,
		}
		goto _23
	_23:
		v24 = v22
		v25 = sqrtf(tls, float32(v24.X*v24.X)+float32(v24.Y*v24.Y))
		goto _26
	_26:
		extent.MaxExtent = v25 + radius1
	case int32(b2_polygonShape):
		poly = shape + 132
		minExtent = float32(float32FromFloat32(100000) * b2_lengthUnitsPerMeter)
		maxExtentSqr = float32FromFloat32(0)
		count = (*Polygon)(unsafe.Pointer(poly)).Count
		i = 0
		for {
			if !(i < count) {
				break
			}
			v2 = *(*Vec2)(unsafe.Pointer(poly + uintptr(i)*8))
			v28 = v2
			v29 = (*Polygon)(unsafe.Pointer(poly)).Centroid
			v30 = Vec2{
				X: v28.X - v29.X,
				Y: v28.Y - v29.Y,
			}
			goto _31
		_31:
			v32 = *(*Vec2)(unsafe.Pointer(poly + 64 + uintptr(i)*8))
			v33 = v30
			v34 = float32(v32.X*v33.X) + float32(v32.Y*v33.Y)
			goto _35
		_35:
			planeOffset = v34
			v36 = minExtent
			v37 = planeOffset
			if v36 < v37 {
				v40 = v36
			} else {
				v40 = v37
			}
			v38 = v40
			goto _39
		_39:
			minExtent = v38
			v41 = v2
			v42 = localCenter
			v43 = Vec2{
				X: v41.X - v42.X,
				Y: v41.Y - v42.Y,
			}
			goto _44
		_44:
			v45 = v43
			v46 = float32(v45.X*v45.X) + float32(v45.Y*v45.Y)
			goto _47
		_47:
			distanceSqr = v46
			v48 = maxExtentSqr
			v49 = distanceSqr
			if v48 > v49 {
				v52 = v48
			} else {
				v52 = v49
			}
			v50 = v52
			goto _51
		_51:
			maxExtentSqr = v50
			goto _27
		_27:
			;
			i = i + 1
		}
		extent.MinExtent = minExtent + (*Polygon)(unsafe.Pointer(poly)).Radius
		extent.MaxExtent = sqrtf(tls, maxExtentSqr) + (*Polygon)(unsafe.Pointer(poly)).Radius
	case int32(b2_segmentShape):
		extent.MinExtent = float32FromFloat32(0)
		v53 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 132))).Point1
		v54 = localCenter
		v55 = Vec2{
			X: v53.X - v54.X,
			Y: v53.Y - v54.Y,
		}
		goto _56
	_56:
		c11 = v55
		v57 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 132))).Point2
		v58 = localCenter
		v59 = Vec2{
			X: v57.X - v58.X,
			Y: v57.Y - v58.Y,
		}
		goto _60
	_60:
		c21 = v59
		v61 = c11
		v62 = float32(v61.X*v61.X) + float32(v61.Y*v61.Y)
		goto _63
	_63:
		v64 = c21
		v65 = float32(v64.X*v64.X) + float32(v64.Y*v64.Y)
		goto _66
	_66:
		v67 = v62
		v68 = v65
		if v67 > v68 {
			v71 = v67
		} else {
			v71 = v68
		}
		v69 = v71
		goto _70
	_70:
		extent.MaxExtent = sqrtf(tls, v69)
	case int32(b2_chainSegmentShape):
		extent.MinExtent = float32FromFloat32(0)
		v72 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 132))).Segment.Point1
		v73 = localCenter
		v74 = Vec2{
			X: v72.X - v73.X,
			Y: v72.Y - v73.Y,
		}
		goto _75
	_75:
		c12 = v74
		v76 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 132))).Segment.Point2
		v77 = localCenter
		v78 = Vec2{
			X: v76.X - v77.X,
			Y: v76.Y - v77.Y,
		}
		goto _79
	_79:
		c22 = v78
		v80 = c12
		v81 = float32(v80.X*v80.X) + float32(v80.Y*v80.Y)
		goto _82
	_82:
		v83 = c22
		v84 = float32(v83.X*v83.X) + float32(v83.Y*v83.Y)
		goto _85
	_85:
		v86 = v81
		v87 = v84
		if v86 > v87 {
			v90 = v86
		} else {
			v90 = v87
		}
		v88 = v90
		goto _89
	_89:
		extent.MaxExtent = sqrtf(tls, v88)
	default:
		break
	}
	return extent
}

func b2RayCastShape(tls *_Stack, input uintptr, shape uintptr, transform Transform) (r CastOutput) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var output CastOutput
	var vx, vy, x, y float32
	var v1, v9 Transform
	var v10, v11, v14, v15, v2, v3, v6, v7 Vec2
	var v13, v5 Rot
	var _ /* localInput at bp+0 */ RayCastInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = output, vx, vy, x, y, v1, v10, v11, v13, v14, v15, v2, v3, v5, v6, v7, v9
	*(*RayCastInput)(unsafe.Pointer(bp)) = *(*RayCastInput)(unsafe.Pointer(input))
	v1 = transform
	v2 = (*RayCastInput)(unsafe.Pointer(input)).Origin
	vx = v2.X - v1.P.X
	vy = v2.Y - v1.P.Y
	v3 = Vec2{
		X: float32(v1.Q.C*vx) + float32(v1.Q.S*vy),
		Y: float32(-v1.Q.S*vx) + float32(v1.Q.C*vy),
	}
	goto _4
_4:
	(*(*RayCastInput)(unsafe.Pointer(bp))).Origin = v3
	v5 = transform.Q
	v6 = (*RayCastInput)(unsafe.Pointer(input)).Translation
	v7 = Vec2{
		X: float32(v5.C*v6.X) + float32(v5.S*v6.Y),
		Y: float32(-v5.S*v6.X) + float32(v5.C*v6.Y),
	}
	goto _8
_8:
	(*(*RayCastInput)(unsafe.Pointer(bp))).Translation = v7
	output = CastOutput{}
	switch (*b2Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b2_capsuleShape):
		output = b2RayCastCapsule(tls, bp, shape+132)
	case int32(b2_circleShape):
		output = b2RayCastCircle(tls, bp, shape+132)
	case int32(b2_polygonShape):
		output = b2RayCastPolygon(tls, bp, shape+132)
	case int32(b2_segmentShape):
		output = b2RayCastSegment(tls, bp, shape+132, boolUint8(false1 != 0))
	case int32(b2_chainSegmentShape):
		output = b2RayCastSegment(tls, bp, shape+132+8, boolUint8(true1 != 0))
	default:
		return output
	}
	v9 = transform
	v10 = output.Point
	x = float32(v9.Q.C*v10.X) - float32(v9.Q.S*v10.Y) + v9.P.X
	y = float32(v9.Q.S*v10.X) + float32(v9.Q.C*v10.Y) + v9.P.Y
	v11 = Vec2{
		X: x,
		Y: y,
	}
	goto _12
_12:
	output.Point = v11
	v13 = transform.Q
	v14 = output.Normal
	v15 = Vec2{
		X: float32(v13.C*v14.X) - float32(v13.S*v14.Y),
		Y: float32(v13.S*v14.X) + float32(v13.C*v14.Y),
	}
	goto _16
_16:
	output.Normal = v15
	return output
}

func b2ShapeCastShape(tls *_Stack, input uintptr, shape uintptr, transform Transform) (r CastOutput) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var i int32
	var output CastOutput
	var vx, vy, x, y float32
	var v10, v2 Transform
	var v11, v12, v15, v16, v3, v4, v7, v8 Vec2
	var v14, v6 Rot
	var _ /* localInput at bp+0 */ ShapeCastInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = i, output, vx, vy, x, y, v10, v11, v12, v14, v15, v16, v2, v3, v4, v6, v7, v8
	*(*ShapeCastInput)(unsafe.Pointer(bp)) = *(*ShapeCastInput)(unsafe.Pointer(input))
	i = 0
	for {
		if !(i < (*(*ShapeCastInput)(unsafe.Pointer(bp))).Proxy.Count) {
			break
		}
		v2 = transform
		v3 = *(*Vec2)(unsafe.Pointer(input + uintptr(i)*8))
		vx = v3.X - v2.P.X
		vy = v3.Y - v2.P.Y
		v4 = Vec2{
			X: float32(v2.Q.C*vx) + float32(v2.Q.S*vy),
			Y: float32(-v2.Q.S*vx) + float32(v2.Q.C*vy),
		}
		goto _5
	_5:
		*(*Vec2)(unsafe.Pointer(bp + uintptr(i)*8)) = v4
		goto _1
	_1:
		;
		i = i + 1
	}
	v6 = transform.Q
	v7 = (*ShapeCastInput)(unsafe.Pointer(input)).Translation
	v8 = Vec2{
		X: float32(v6.C*v7.X) + float32(v6.S*v7.Y),
		Y: float32(-v6.S*v7.X) + float32(v6.C*v7.Y),
	}
	goto _9
_9:
	(*(*ShapeCastInput)(unsafe.Pointer(bp))).Translation = v8
	output = CastOutput{}
	switch (*b2Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b2_capsuleShape):
		output = b2ShapeCastCapsule(tls, bp, shape+132)
	case int32(b2_circleShape):
		output = b2ShapeCastCircle(tls, bp, shape+132)
	case int32(b2_polygonShape):
		output = b2ShapeCastPolygon(tls, bp, shape+132)
	case int32(b2_segmentShape):
		output = b2ShapeCastSegment(tls, bp, shape+132)
	case int32(b2_chainSegmentShape):
		output = b2ShapeCastSegment(tls, bp, shape+132+8)
	default:
		return output
	}
	v10 = transform
	v11 = output.Point
	x = float32(v10.Q.C*v11.X) - float32(v10.Q.S*v11.Y) + v10.P.X
	y = float32(v10.Q.S*v11.X) + float32(v10.Q.C*v11.Y) + v10.P.Y
	v12 = Vec2{
		X: x,
		Y: y,
	}
	goto _13
_13:
	output.Point = v12
	v14 = transform.Q
	v15 = output.Normal
	v16 = Vec2{
		X: float32(v14.C*v15.X) - float32(v14.S*v15.Y),
		Y: float32(v14.S*v15.X) + float32(v14.C*v15.Y),
	}
	goto _17
_17:
	output.Normal = v16
	return output
}

func b2CreateShapeProxy(tls *_Stack, shape uintptr, bp uintptr, type1 BodyType, transform Transform, forcePairCreation uint8) {
	if !((*b2Shape)(unsafe.Pointer(shape)).ProxyKey == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+11986, __ccgo_ts+11037, int32FromInt32(903)) != 0 {
		__builtin_trap(tls)
	}
	b2UpdateShapeAABBs(tls, shape, transform, type1)
	// Create proxies in the broad-phase.
	(*b2Shape)(unsafe.Pointer(shape)).ProxyKey = b2BroadPhase_CreateProxy(tls, bp, type1, (*b2Shape)(unsafe.Pointer(shape)).FatAABB, (*b2Shape)(unsafe.Pointer(shape)).Filter.CategoryBits, (*b2Shape)(unsafe.Pointer(shape)).Id, forcePairCreation)
	if !((*b2Shape)(unsafe.Pointer(shape)).ProxyKey&int32FromInt32(3) < int32(b2_bodyTypeCount)) && b2InternalAssertFcn(tls, __ccgo_ts+12019, __ccgo_ts+11037, int32FromInt32(910)) != 0 {
		__builtin_trap(tls)
	}
}

func b2DestroyShapeProxy(tls *_Stack, shape uintptr, bp uintptr) {
	if (*b2Shape)(unsafe.Pointer(shape)).ProxyKey != -int32(1) {
		b2BroadPhase_DestroyProxy(tls, bp, (*b2Shape)(unsafe.Pointer(shape)).ProxyKey)
		(*b2Shape)(unsafe.Pointer(shape)).ProxyKey = -int32(1)
	}
}

func b2MakeShapeDistanceProxy(tls *_Stack, shape uintptr) (r ShapeProxy) {
	var empty ShapeProxy
	_ = empty
	switch (*b2Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b2_capsuleShape):
		return b2MakeProxy(tls, shape+132, int32(2), (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo19_132.Capsule.Radius)
	case int32(b2_circleShape):
		return b2MakeProxy(tls, shape+132, int32(1), (*(*Circle)(unsafe.Add(unsafe.Pointer(shape), 132))).Radius)
	case int32(b2_polygonShape):
		return b2MakeProxy(tls, shape+132, (*(*Polygon)(unsafe.Add(unsafe.Pointer(shape), 132))).Count, (*(*Polygon)(unsafe.Add(unsafe.Pointer(shape), 132))).Radius)
	case int32(b2_segmentShape):
		return b2MakeProxy(tls, shape+132, int32(2), float32FromFloat32(0))
	case int32(b2_chainSegmentShape):
		return b2MakeProxy(tls, shape+132+8, int32(2), float32FromFloat32(0))
	default:
		if bool(!(int32FromInt32(false1) != 0)) && b2InternalAssertFcn(tls, __ccgo_ts+4123, __ccgo_ts+11037, int32FromInt32(938)) != 0 {
			__builtin_trap(tls)
		}
		empty = ShapeProxy{}
		return empty
	}
	return r
}

func b2Shape_GetBody(tls *_Stack, shapeId ShapeId) (r BodyId) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	return b2MakeBodyId(tls, world, (*b2Shape)(unsafe.Pointer(shape)).BodyId)
}

func b2Shape_GetWorld(tls *_Stack, shapeId ShapeId) (r WorldId) {
	var world uintptr
	_ = world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	return WorldId{
		Index1:     uint16FromInt32(int32FromUint16(shapeId.World0) + int32(1)),
		Generation: (*b2World)(unsafe.Pointer(world)).Generation,
	}
}

func b2Shape_SetUserData(tls *_Stack, shapeId ShapeId, userData uintptr) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	(*b2Shape)(unsafe.Pointer(shape)).UserData = userData
}

func b2Shape_GetUserData(tls *_Stack, shapeId ShapeId) (r uintptr) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	return (*b2Shape)(unsafe.Pointer(shape)).UserData
}

func b2Shape_IsSensor(tls *_Stack, shapeId ShapeId) (r uint8) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	return boolUint8((*b2Shape)(unsafe.Pointer(shape)).SensorIndex != -int32(1))
}

func b2Shape_TestPoint(tls *_Stack, shapeId ShapeId, point Vec2) (r uint8) {
	var localPoint, v2, v3 Vec2
	var shape, world uintptr
	var transform, v1 Transform
	var vx, vy float32
	_, _, _, _, _, _, _, _, _ = localPoint, shape, transform, vx, vy, world, v1, v2, v3
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	transform = b2GetBodyTransform(tls, world, (*b2Shape)(unsafe.Pointer(shape)).BodyId)
	v1 = transform
	v2 = point
	vx = v2.X - v1.P.X
	vy = v2.Y - v1.P.Y
	v3 = Vec2{
		X: float32(v1.Q.C*vx) + float32(v1.Q.S*vy),
		Y: float32(-v1.Q.S*vx) + float32(v1.Q.C*vy),
	}
	goto _4
_4:
	localPoint = v3
	switch (*b2Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b2_capsuleShape):
		return b2PointInCapsule(tls, localPoint, shape+132)
	case int32(b2_circleShape):
		return b2PointInCircle(tls, localPoint, shape+132)
	case int32(b2_polygonShape):
		return b2PointInPolygon(tls, localPoint, shape+132)
	default:
		return boolUint8(false1 != 0)
	}
	return r
}

// C documentation
//
//	// todo_erin untested
func b2Shape_RayCast(tls *_Stack, shapeId ShapeId, input uintptr) (r CastOutput) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var output CastOutput
	var shape, world uintptr
	var transform, v1, v23 Transform
	var vx, vy, x, y float32
	var v19, v5 Rot
	var v2, v20, v21, v24, v25, v3, v6, v7 Vec2
	var _ /* localInput at bp+0 */ RayCastInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = output, shape, transform, vx, vy, world, x, y, v1, v19, v2, v20, v21, v23, v24, v25, v3, v5, v6, v7
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	transform = b2GetBodyTransform(tls, world, (*b2Shape)(unsafe.Pointer(shape)).BodyId)
	v1 = transform
	v2 = (*RayCastInput)(unsafe.Pointer(input)).Origin
	vx = v2.X - v1.P.X
	vy = v2.Y - v1.P.Y
	v3 = Vec2{
		X: float32(v1.Q.C*vx) + float32(v1.Q.S*vy),
		Y: float32(-v1.Q.S*vx) + float32(v1.Q.C*vy),
	}
	goto _4
_4:
	(*(*RayCastInput)(unsafe.Pointer(bp))).Origin = v3
	v5 = transform.Q
	v6 = (*RayCastInput)(unsafe.Pointer(input)).Translation
	v7 = Vec2{
		X: float32(v5.C*v6.X) + float32(v5.S*v6.Y),
		Y: float32(-v5.S*v6.X) + float32(v5.C*v6.Y),
	}
	goto _8
_8:
	(*(*RayCastInput)(unsafe.Pointer(bp))).Translation = v7
	(*(*RayCastInput)(unsafe.Pointer(bp))).MaxFraction = (*RayCastInput)(unsafe.Pointer(input)).MaxFraction
	output = CastOutput{}
	switch (*b2Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b2_capsuleShape):
		goto _9
	case int32(b2_circleShape):
		goto _10
	case int32(b2_segmentShape):
		goto _11
	case int32(b2_polygonShape):
		goto _12
	case int32(b2_chainSegmentShape):
		goto _13
	default:
		goto _14
	}
	goto _15
_9:
	;
	output = b2RayCastCapsule(tls, bp, shape+132)
	goto _15
_10:
	;
	output = b2RayCastCircle(tls, bp, shape+132)
	goto _15
_11:
	;
	output = b2RayCastSegment(tls, bp, shape+132, boolUint8(false1 != 0))
	goto _15
_12:
	;
	output = b2RayCastPolygon(tls, bp, shape+132)
	goto _15
_13:
	;
	output = b2RayCastSegment(tls, bp, shape+132+8, boolUint8(true1 != 0))
	goto _15
_14:
	;
_18:
	;
	if bool(!(int32FromInt32(false1) != 0)) && b2InternalAssertFcn(tls, __ccgo_ts+4123, __ccgo_ts+11037, int32FromInt32(1041)) != 0 {
		__builtin_trap(tls)
	}
	goto _17
_17:
	;
	if 0 != 0 {
		goto _18
	}
	goto _16
_16:
	;
	return output
_15:
	;
	if output.Hit != 0 {
		// convert to world coordinates
		v19 = transform.Q
		v20 = output.Normal
		v21 = Vec2{
			X: float32(v19.C*v20.X) - float32(v19.S*v20.Y),
			Y: float32(v19.S*v20.X) + float32(v19.C*v20.Y),
		}
		goto _22
	_22:
		output.Normal = v21
		v23 = transform
		v24 = output.Point
		x = float32(v23.Q.C*v24.X) - float32(v23.Q.S*v24.Y) + v23.P.X
		y = float32(v23.Q.S*v24.X) + float32(v23.Q.C*v24.Y) + v23.P.Y
		v25 = Vec2{
			X: x,
			Y: y,
		}
		goto _26
	_26:
		output.Point = v25
	}
	return output
}

func b2Shape_SetDensity(tls *_Stack, shapeId ShapeId, density float32, updateBodyMass uint8) {
	var body, shape, world, v1, v3 uintptr
	var v2 int32
	_, _, _, _, _, _ = body, shape, world, v1, v2, v3
	if !(b2IsValidFloat(tls, density) != 0 && density >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+12071, __ccgo_ts+11037, int32FromInt32(1057)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	shape = b2GetShape(tls, world, shapeId)
	if density == (*b2Shape)(unsafe.Pointer(shape)).Density {
		// early return to avoid expensive function
		return
	}
	(*b2Shape)(unsafe.Pointer(shape)).Density = density
	if int32FromUint8(updateBodyMass) == int32(true1) {
		v1 = world + 1024
		v2 = (*b2Shape)(unsafe.Pointer(shape)).BodyId
		if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
		goto _4
	_4:
		body = v3
		b2UpdateBodyMassData(tls, world, body)
	}
}

func b2Shape_GetDensity(tls *_Stack, shapeId ShapeId) (r float32) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	return (*b2Shape)(unsafe.Pointer(shape)).Density
}

func b2Shape_SetFriction(tls *_Stack, shapeId ShapeId, friction float32) {
	var shape, world uintptr
	_, _ = shape, world
	if !(b2IsValidFloat(tls, friction) != 0 && friction >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+12116, __ccgo_ts+11037, int32FromInt32(1090)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	if !(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)) && b2InternalAssertFcn(tls, __ccgo_ts+1152, __ccgo_ts+11037, int32FromInt32(1093)) != 0 {
		__builtin_trap(tls)
	}
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	shape = b2GetShape(tls, world, shapeId)
	(*b2Shape)(unsafe.Pointer(shape)).Friction = friction
}

func b2Shape_GetFriction(tls *_Stack, shapeId ShapeId) (r float32) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	return (*b2Shape)(unsafe.Pointer(shape)).Friction
}

func b2Shape_SetRestitution(tls *_Stack, shapeId ShapeId, restitution float32) {
	var shape, world uintptr
	_, _ = shape, world
	if !(b2IsValidFloat(tls, restitution) != 0 && restitution >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+12163, __ccgo_ts+11037, int32FromInt32(1112)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	if !(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)) && b2InternalAssertFcn(tls, __ccgo_ts+1152, __ccgo_ts+11037, int32FromInt32(1115)) != 0 {
		__builtin_trap(tls)
	}
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	shape = b2GetShape(tls, world, shapeId)
	(*b2Shape)(unsafe.Pointer(shape)).Restitution = restitution
}

func b2Shape_GetRestitution(tls *_Stack, shapeId ShapeId) (r float32) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	return (*b2Shape)(unsafe.Pointer(shape)).Restitution
}

func b2Shape_SetMaterial(tls *_Stack, shapeId ShapeId, material int32) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	if !(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)) && b2InternalAssertFcn(tls, __ccgo_ts+1152, __ccgo_ts+11037, int32FromInt32(1135)) != 0 {
		__builtin_trap(tls)
	}
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	shape = b2GetShape(tls, world, shapeId)
	(*b2Shape)(unsafe.Pointer(shape)).UserMaterialId = material
}

func b2Shape_GetMaterial(tls *_Stack, shapeId ShapeId) (r int32) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	return (*b2Shape)(unsafe.Pointer(shape)).UserMaterialId
}

func b2Shape_GetSurfaceMaterial(tls *_Stack, shapeId ShapeId) (r SurfaceMaterial) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	return SurfaceMaterial{
		Friction:          (*b2Shape)(unsafe.Pointer(shape)).Friction,
		Restitution:       (*b2Shape)(unsafe.Pointer(shape)).Restitution,
		RollingResistance: (*b2Shape)(unsafe.Pointer(shape)).RollingResistance,
		TangentSpeed:      (*b2Shape)(unsafe.Pointer(shape)).TangentSpeed,
		UserMaterialId:    (*b2Shape)(unsafe.Pointer(shape)).UserMaterialId,
		CustomColor:       (*b2Shape)(unsafe.Pointer(shape)).CustomColor,
	}
}

func b2Shape_SetSurfaceMaterial(tls *_Stack, shapeId ShapeId, surfaceMaterial SurfaceMaterial) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	(*b2Shape)(unsafe.Pointer(shape)).Friction = surfaceMaterial.Friction
	(*b2Shape)(unsafe.Pointer(shape)).Restitution = surfaceMaterial.Restitution
	(*b2Shape)(unsafe.Pointer(shape)).RollingResistance = surfaceMaterial.RollingResistance
	(*b2Shape)(unsafe.Pointer(shape)).TangentSpeed = surfaceMaterial.TangentSpeed
	(*b2Shape)(unsafe.Pointer(shape)).UserMaterialId = surfaceMaterial.UserMaterialId
	(*b2Shape)(unsafe.Pointer(shape)).CustomColor = surfaceMaterial.CustomColor
}

func b2Shape_GetFilter(tls *_Stack, shapeId ShapeId) (r Filter) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	return (*b2Shape)(unsafe.Pointer(shape)).Filter
}

func b2Shape_SetFilter(tls *_Stack, shapeId ShapeId, filter Filter) {
	var destroyProxy, wakeBodies uint8
	var shape, world uintptr
	_, _, _, _ = destroyProxy, shape, wakeBodies, world
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	shape = b2GetShape(tls, world, shapeId)
	if filter.MaskBits == (*b2Shape)(unsafe.Pointer(shape)).Filter.MaskBits && filter.CategoryBits == (*b2Shape)(unsafe.Pointer(shape)).Filter.CategoryBits && filter.GroupIndex == (*b2Shape)(unsafe.Pointer(shape)).Filter.GroupIndex {
		return
	}
	// If the category bits change, I need to destroy the proxy because it affects the tree sorting.
	destroyProxy = boolUint8(filter.CategoryBits != (*b2Shape)(unsafe.Pointer(shape)).Filter.CategoryBits)
	(*b2Shape)(unsafe.Pointer(shape)).Filter = filter
	// need to wake bodies because a filter change may destroy contacts
	wakeBodies = boolUint8(true1 != 0)
	b2ResetProxy(tls, world, shape, wakeBodies, destroyProxy)
	// note: this does not immediately update sensor overlaps. Instead sensor
	// overlaps are updated the next time step
}

func b2Shape_EnableSensorEvents(tls *_Stack, shapeId ShapeId, flag uint8) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	shape = b2GetShape(tls, world, shapeId)
	(*b2Shape)(unsafe.Pointer(shape)).EnableSensorEvents = flag
}

func b2Shape_AreSensorEventsEnabled(tls *_Stack, shapeId ShapeId) (r uint8) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	return (*b2Shape)(unsafe.Pointer(shape)).EnableSensorEvents
}

func b2Shape_EnableContactEvents(tls *_Stack, shapeId ShapeId, flag uint8) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	shape = b2GetShape(tls, world, shapeId)
	(*b2Shape)(unsafe.Pointer(shape)).EnableContactEvents = flag
}

func b2Shape_AreContactEventsEnabled(tls *_Stack, shapeId ShapeId) (r uint8) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	return (*b2Shape)(unsafe.Pointer(shape)).EnableContactEvents
}

func b2Shape_EnablePreSolveEvents(tls *_Stack, shapeId ShapeId, flag uint8) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	shape = b2GetShape(tls, world, shapeId)
	(*b2Shape)(unsafe.Pointer(shape)).EnablePreSolveEvents = flag
}

func b2Shape_ArePreSolveEventsEnabled(tls *_Stack, shapeId ShapeId) (r uint8) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	return (*b2Shape)(unsafe.Pointer(shape)).EnablePreSolveEvents
}

func b2Shape_EnableHitEvents(tls *_Stack, shapeId ShapeId, flag uint8) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	shape = b2GetShape(tls, world, shapeId)
	(*b2Shape)(unsafe.Pointer(shape)).EnableHitEvents = flag
}

func b2Shape_AreHitEventsEnabled(tls *_Stack, shapeId ShapeId) (r uint8) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	return (*b2Shape)(unsafe.Pointer(shape)).EnableHitEvents
}

func b2Shape_GetType(tls *_Stack, shapeId ShapeId) (r ShapeType) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	return (*b2Shape)(unsafe.Pointer(shape)).Type1
}

func b2Shape_GetCircle(tls *_Stack, shapeId ShapeId) (r Circle) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	if !((*b2Shape)(unsafe.Pointer(shape)).Type1 == int32(b2_circleShape)) && b2InternalAssertFcn(tls, __ccgo_ts+12216, __ccgo_ts+11037, int32FromInt32(1350)) != 0 {
		__builtin_trap(tls)
	}
	return *(*Circle)(unsafe.Add(unsafe.Pointer(shape), 132))
}

func b2Shape_GetSegment(tls *_Stack, shapeId ShapeId) (r Segment) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	if !((*b2Shape)(unsafe.Pointer(shape)).Type1 == int32(b2_segmentShape)) && b2InternalAssertFcn(tls, __ccgo_ts+12246, __ccgo_ts+11037, int32FromInt32(1358)) != 0 {
		__builtin_trap(tls)
	}
	return *(*Segment)(unsafe.Add(unsafe.Pointer(shape), 132))
}

func b2Shape_GetChainSegment(tls *_Stack, shapeId ShapeId) (r ChainSegment) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	if !((*b2Shape)(unsafe.Pointer(shape)).Type1 == int32(b2_chainSegmentShape)) && b2InternalAssertFcn(tls, __ccgo_ts+12277, __ccgo_ts+11037, int32FromInt32(1366)) != 0 {
		__builtin_trap(tls)
	}
	return *(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 132))
}

func b2Shape_GetCapsule(tls *_Stack, shapeId ShapeId) (r Capsule) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	if !((*b2Shape)(unsafe.Pointer(shape)).Type1 == int32(b2_capsuleShape)) && b2InternalAssertFcn(tls, __ccgo_ts+12313, __ccgo_ts+11037, int32FromInt32(1374)) != 0 {
		__builtin_trap(tls)
	}
	return (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo19_132.Capsule
}

func b2Shape_GetPolygon(tls *_Stack, shapeId ShapeId) (r Polygon) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	if !((*b2Shape)(unsafe.Pointer(shape)).Type1 == int32(b2_polygonShape)) && b2InternalAssertFcn(tls, __ccgo_ts+12344, __ccgo_ts+11037, int32FromInt32(1382)) != 0 {
		__builtin_trap(tls)
	}
	return *(*Polygon)(unsafe.Add(unsafe.Pointer(shape), 132))
}

func b2Shape_SetCircle(tls *_Stack, shapeId ShapeId, circle uintptr) {
	var destroyProxy, wakeBodies uint8
	var shape, world uintptr
	_, _, _, _ = destroyProxy, shape, wakeBodies, world
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	shape = b2GetShape(tls, world, shapeId)
	*(*Circle)(unsafe.Add(unsafe.Pointer(shape), 132)) = *(*Circle)(unsafe.Pointer(circle))
	(*b2Shape)(unsafe.Pointer(shape)).Type1 = int32(b2_circleShape)
	// need to wake bodies so they can react to the shape change
	wakeBodies = boolUint8(true1 != 0)
	destroyProxy = boolUint8(true1 != 0)
	b2ResetProxy(tls, world, shape, wakeBodies, destroyProxy)
}

func b2Shape_SetCapsule(tls *_Stack, shapeId ShapeId, capsule uintptr) {
	var destroyProxy, wakeBodies uint8
	var shape, world uintptr
	_, _, _, _ = destroyProxy, shape, wakeBodies, world
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	shape = b2GetShape(tls, world, shapeId)
	(*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo19_132.Capsule = *(*Capsule)(unsafe.Pointer(capsule))
	(*b2Shape)(unsafe.Pointer(shape)).Type1 = int32(b2_capsuleShape)
	// need to wake bodies so they can react to the shape change
	wakeBodies = boolUint8(true1 != 0)
	destroyProxy = boolUint8(true1 != 0)
	b2ResetProxy(tls, world, shape, wakeBodies, destroyProxy)
}

func b2Shape_SetSegment(tls *_Stack, shapeId ShapeId, segment uintptr) {
	var destroyProxy, wakeBodies uint8
	var shape, world uintptr
	_, _, _, _ = destroyProxy, shape, wakeBodies, world
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	shape = b2GetShape(tls, world, shapeId)
	*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 132)) = *(*Segment)(unsafe.Pointer(segment))
	(*b2Shape)(unsafe.Pointer(shape)).Type1 = int32(b2_segmentShape)
	// need to wake bodies so they can react to the shape change
	wakeBodies = boolUint8(true1 != 0)
	destroyProxy = boolUint8(true1 != 0)
	b2ResetProxy(tls, world, shape, wakeBodies, destroyProxy)
}

func b2Shape_SetPolygon(tls *_Stack, shapeId ShapeId, polygon uintptr) {
	var destroyProxy, wakeBodies uint8
	var shape, world uintptr
	_, _, _, _ = destroyProxy, shape, wakeBodies, world
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	shape = b2GetShape(tls, world, shapeId)
	*(*Polygon)(unsafe.Add(unsafe.Pointer(shape), 132)) = *(*Polygon)(unsafe.Pointer(polygon))
	(*b2Shape)(unsafe.Pointer(shape)).Type1 = int32(b2_polygonShape)
	// need to wake bodies so they can react to the shape change
	wakeBodies = boolUint8(true1 != 0)
	destroyProxy = boolUint8(true1 != 0)
	b2ResetProxy(tls, world, shape, wakeBodies, destroyProxy)
}

func b2Shape_GetParentChain(tls *_Stack, shapeId ShapeId) (r ChainId) {
	var chain, shape, world, v1, v3 uintptr
	var chainId, v2 int32
	var id ChainId
	_, _, _, _, _, _, _, _ = chain, chainId, id, shape, world, v1, v2, v3
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	if (*b2Shape)(unsafe.Pointer(shape)).Type1 == int32(b2_chainSegmentShape) {
		chainId = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 132))).ChainId
		if chainId != -int32(1) {
			v1 = world + 1264
			v2 = chainId
			if !(0 <= v2 && v2 < (*b2ChainShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(137)) != 0 {
				__builtin_trap(tls)
			}
			v3 = (*b2ChainShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*48
			goto _4
		_4:
			chain = v3
			id = ChainId{
				Index1:     chainId + int32(1),
				World0:     shapeId.World0,
				Generation: (*b2ChainShape)(unsafe.Pointer(chain)).Generation,
			}
			return id
		}
	}
	return ChainId{}
}

func b2Shape_GetContactCapacity(tls *_Stack, shapeId ShapeId) (r int32) {
	var body, shape, world, v1, v3 uintptr
	var v2 int32
	_, _, _, _, _, _ = body, shape, world, v1, v2, v3
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return 0
	}
	shape = b2GetShape(tls, world, shapeId)
	if (*b2Shape)(unsafe.Pointer(shape)).SensorIndex != -int32(1) {
		return 0
	}
	v1 = world + 1024
	v2 = (*b2Shape)(unsafe.Pointer(shape)).BodyId
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	body = v3
	// Conservative and fast
	return (*b2Body)(unsafe.Pointer(body)).ContactCount
}

func b2Shape_GetContactData(tls *_Stack, shapeId ShapeId, contactData uintptr, capacity int32) (r int32) {
	var body, contact, contactSim, shape, shapeA, shapeB, world, v1, v3, v5, v7 uintptr
	var contactId, contactKey, edgeIndex, index2, v2, v6 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, contact, contactId, contactKey, contactSim, edgeIndex, index2, shape, shapeA, shapeB, world, v1, v2, v3, v5, v6, v7
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return 0
	}
	shape = b2GetShape(tls, world, shapeId)
	if (*b2Shape)(unsafe.Pointer(shape)).SensorIndex != -int32(1) {
		return 0
	}
	v1 = world + 1024
	v2 = (*b2Shape)(unsafe.Pointer(shape)).BodyId
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	body = v3
	contactKey = (*b2Body)(unsafe.Pointer(body)).HeadContactKey
	index2 = 0
	for contactKey != -int32(1) && index2 < capacity {
		contactId = contactKey >> int32(1)
		edgeIndex = contactKey & int32(1)
		v5 = world + 1144
		v6 = contactId
		if !(0 <= v6 && v6 < (*b2ContactArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
			__builtin_trap(tls)
		}
		v7 = (*b2ContactArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*68
		goto _8
	_8:
		contact = v7
		// Does contact involve this shape and is it touching?
		if ((*b2Contact)(unsafe.Pointer(contact)).ShapeIdA == shapeId.Index1-int32(1) || (*b2Contact)(unsafe.Pointer(contact)).ShapeIdB == shapeId.Index1-int32(1)) && (*b2Contact)(unsafe.Pointer(contact)).Flags&uint32(b2_contactTouchingFlag) != uint32(0) {
			shapeA = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b2Contact)(unsafe.Pointer(contact)).ShapeIdA)*288
			shapeB = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b2Contact)(unsafe.Pointer(contact)).ShapeIdB)*288
			(*(*ContactData)(unsafe.Pointer(contactData + uintptr(index2)*128))).ShapeIdA = ShapeId{
				Index1:     (*b2Shape)(unsafe.Pointer(shapeA)).Id + int32(1),
				World0:     shapeId.World0,
				Generation: (*b2Shape)(unsafe.Pointer(shapeA)).Generation,
			}
			(*(*ContactData)(unsafe.Pointer(contactData + uintptr(index2)*128))).ShapeIdB = ShapeId{
				Index1:     (*b2Shape)(unsafe.Pointer(shapeB)).Id + int32(1),
				World0:     shapeId.World0,
				Generation: (*b2Shape)(unsafe.Pointer(shapeB)).Generation,
			}
			contactSim = b2GetContactSim(tls, world, contact)
			(*(*ContactData)(unsafe.Pointer(contactData + uintptr(index2)*128))).Manifold = (*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold
			index2 = index2 + int32(1)
		}
		contactKey = (*(*b2ContactEdge)(unsafe.Pointer(contact + 12 + uintptr(edgeIndex)*12))).NextKey
	}
	if !(index2 <= capacity) && b2InternalAssertFcn(tls, __ccgo_ts+1464, __ccgo_ts+11037, int32FromInt32(1640)) != 0 {
		__builtin_trap(tls)
	}
	return index2
}

func b2Shape_GetSensorCapacity(tls *_Stack, shapeId ShapeId) (r int32) {
	var sensor, shape, world, v1, v3 uintptr
	var v2 int32
	_, _, _, _, _, _ = sensor, shape, world, v1, v2, v3
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return 0
	}
	shape = b2GetShape(tls, world, shapeId)
	if (*b2Shape)(unsafe.Pointer(shape)).SensorIndex == -int32(1) {
		return 0
	}
	v1 = world + 1280
	v2 = (*b2Shape)(unsafe.Pointer(shape)).SensorIndex
	if !(0 <= v2 && v2 < (*b2SensorArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+10835, int32FromInt32(35)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2SensorArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*40
	goto _4
_4:
	sensor = v3
	return (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count
}

func b2Shape_GetSensorOverlaps(tls *_Stack, shapeId ShapeId, overlaps uintptr, capacity int32) (r int32) {
	var count, i, v2, v5, v6, v7, v9 int32
	var refs, sensor, shape, world, v1, v3 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _ = count, i, refs, sensor, shape, world, v1, v2, v3, v5, v6, v7, v9
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return 0
	}
	shape = b2GetShape(tls, world, shapeId)
	if (*b2Shape)(unsafe.Pointer(shape)).SensorIndex == -int32(1) {
		return 0
	}
	v1 = world + 1280
	v2 = (*b2Shape)(unsafe.Pointer(shape)).SensorIndex
	if !(0 <= v2 && v2 < (*b2SensorArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+10835, int32FromInt32(35)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2SensorArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*40
	goto _4
_4:
	sensor = v3
	v5 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count
	v6 = capacity
	if v5 < v6 {
		v9 = v5
	} else {
		v9 = v6
	}
	v7 = v9
	goto _8
_8:
	count = v7
	refs = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data
	i = 0
	for {
		if !(i < count) {
			break
		}
		*(*ShapeId)(unsafe.Pointer(overlaps + uintptr(i)*8)) = ShapeId{
			Index1:     (*(*b2ShapeRef)(unsafe.Pointer(refs + uintptr(i)*8))).ShapeId + int32(1),
			World0:     shapeId.World0,
			Generation: (*(*b2ShapeRef)(unsafe.Pointer(refs + uintptr(i)*8))).Generation,
		}
		goto _10
	_10:
		;
		i = i + 1
	}
	return count
}

func b2Shape_GetAABB(tls *_Stack, shapeId ShapeId) (r AABB) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return AABB{}
	}
	shape = b2GetShape(tls, world, shapeId)
	return (*b2Shape)(unsafe.Pointer(shape)).Aabb
}

func b2Shape_GetMassData(tls *_Stack, shapeId ShapeId) (r MassData) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return MassData{}
	}
	shape = b2GetShape(tls, world, shapeId)
	return b2ComputeShapeMass(tls, shape)
}

func b2Shape_GetClosestPoint(tls *_Stack, shapeId ShapeId, _target Vec2) (r Vec2) {
	bp := tls.Alloc(208)
	defer tls.Free(208)
	*(*Vec2)(unsafe.Pointer(bp)) = _target
	var body, shape, world, v1, v3 uintptr
	var output DistanceOutput
	var transform Transform
	var v2 int32
	var _ /* cache at bp+192 */ SimplexCache
	var _ /* input at bp+8 */ DistanceInput
	_, _, _, _, _, _, _, _ = body, output, shape, transform, world, v1, v2, v3
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return Vec2{}
	}
	shape = b2GetShape(tls, world, shapeId)
	v1 = world + 1024
	v2 = (*b2Shape)(unsafe.Pointer(shape)).BodyId
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	body = v3
	transform = b2GetBodyTransformQuick(tls, world, body)
	(*(*DistanceInput)(unsafe.Pointer(bp + 8))).ProxyA = b2MakeShapeDistanceProxy(tls, shape)
	(*(*DistanceInput)(unsafe.Pointer(bp + 8))).ProxyB = b2MakeProxy(tls, bp, int32(1), float32FromFloat32(0))
	(*(*DistanceInput)(unsafe.Pointer(bp + 8))).TransformA = transform
	(*(*DistanceInput)(unsafe.Pointer(bp + 8))).TransformB = b2Transform_identity
	(*(*DistanceInput)(unsafe.Pointer(bp + 8))).UseRadii = boolUint8(true1 != 0)
	*(*SimplexCache)(unsafe.Pointer(bp + 192)) = SimplexCache{}
	output = b2ShapeDistance(tls, bp+8, bp+192, uintptrFromInt32(0), 0)
	return output.PointA
}

func b2DefaultShapeDef(tls *_Stack) (r ShapeDef) {
	var def ShapeDef
	_ = def
	def = ShapeDef{}
	def.Material.Friction = float32FromFloat32(0.6)
	def.Density = float32FromFloat32(1)
	def.Filter = b2DefaultFilter(tls)
	def.UpdateBodyMass = boolUint8(true1 != 0)
	def.InvokeContactCreation = boolUint8(true1 != 0)
	def.InternalValue = int32(_B2_SECRET_COOKIE)
	return def
}

func b2DrawShape(tls *_Stack, draw uintptr, shape uintptr, xf Transform, color HexColor) {
	var capsule, circle, poly, segment, segment1 uintptr
	var p1, p11, p12, p2, p21, p22, v10, v11, v14, v15, v18, v19, v2, v22, v23, v26, v27, v29, v3, v30, v32, v6, v7 Vec2
	var x, y, v31 float32
	var v1, v13, v17, v21, v25, v5, v9 Transform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = capsule, circle, p1, p11, p12, p2, p21, p22, poly, segment, segment1, x, y, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v29, v3, v30, v31, v32, v5, v6, v7, v9
	switch (*b2Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b2_capsuleShape):
		capsule = shape + 132
		v1 = xf
		v2 = (*Capsule)(unsafe.Pointer(capsule)).Center1
		x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
		y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
		v3 = Vec2{
			X: x,
			Y: y,
		}
		goto _4
	_4:
		p1 = v3
		v5 = xf
		v6 = (*Capsule)(unsafe.Pointer(capsule)).Center2
		x = float32(v5.Q.C*v6.X) - float32(v5.Q.S*v6.Y) + v5.P.X
		y = float32(v5.Q.S*v6.X) + float32(v5.Q.C*v6.Y) + v5.P.Y
		v7 = Vec2{
			X: x,
			Y: y,
		}
		goto _8
	_8:
		p2 = v7
		(*(*func(*_Stack, Vec2, Vec2, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSolidCapsuleFcn})))(tls, p1, p2, (*Capsule)(unsafe.Pointer(capsule)).Radius, color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	case int32(b2_circleShape):
		circle = shape + 132
		v9 = xf
		v10 = (*Circle)(unsafe.Pointer(circle)).Center
		x = float32(v9.Q.C*v10.X) - float32(v9.Q.S*v10.Y) + v9.P.X
		y = float32(v9.Q.S*v10.X) + float32(v9.Q.C*v10.Y) + v9.P.Y
		v11 = Vec2{
			X: x,
			Y: y,
		}
		goto _12
	_12:
		xf.P = v11
		(*(*func(*_Stack, Transform, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSolidCircleFcn})))(tls, xf, (*Circle)(unsafe.Pointer(circle)).Radius, color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	case int32(b2_polygonShape):
		poly = shape + 132
		(*(*func(*_Stack, Transform, uintptr, int32, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSolidPolygonFcn})))(tls, xf, poly, (*Polygon)(unsafe.Pointer(poly)).Count, (*Polygon)(unsafe.Pointer(poly)).Radius, color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	case int32(b2_segmentShape):
		segment = shape + 132
		v13 = xf
		v14 = (*Segment)(unsafe.Pointer(segment)).Point1
		x = float32(v13.Q.C*v14.X) - float32(v13.Q.S*v14.Y) + v13.P.X
		y = float32(v13.Q.S*v14.X) + float32(v13.Q.C*v14.Y) + v13.P.Y
		v15 = Vec2{
			X: x,
			Y: y,
		}
		goto _16
	_16:
		p11 = v15
		v17 = xf
		v18 = (*Segment)(unsafe.Pointer(segment)).Point2
		x = float32(v17.Q.C*v18.X) - float32(v17.Q.S*v18.Y) + v17.P.X
		y = float32(v17.Q.S*v18.X) + float32(v17.Q.C*v18.Y) + v17.P.Y
		v19 = Vec2{
			X: x,
			Y: y,
		}
		goto _20
	_20:
		p21 = v19
		(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, p11, p21, color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	case int32(b2_chainSegmentShape):
		segment1 = shape + 132 + 8
		v21 = xf
		v22 = (*Segment)(unsafe.Pointer(segment1)).Point1
		x = float32(v21.Q.C*v22.X) - float32(v21.Q.S*v22.Y) + v21.P.X
		y = float32(v21.Q.S*v22.X) + float32(v21.Q.C*v22.Y) + v21.P.Y
		v23 = Vec2{
			X: x,
			Y: y,
		}
		goto _24
	_24:
		p12 = v23
		v25 = xf
		v26 = (*Segment)(unsafe.Pointer(segment1)).Point2
		x = float32(v25.Q.C*v26.X) - float32(v25.Q.S*v26.Y) + v25.P.X
		y = float32(v25.Q.S*v26.X) + float32(v25.Q.C*v26.Y) + v25.P.Y
		v27 = Vec2{
			X: x,
			Y: y,
		}
		goto _28
	_28:
		p22 = v27
		(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, p12, p22, color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		(*(*func(*_Stack, Vec2, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, p22, float32FromFloat32(4), color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		v29 = p12
		v30 = p22
		v31 = float32FromFloat32(0.1)
		v32 = Vec2{
			X: float32((float32FromFloat32(1)-v31)*v29.X) + float32(v31*v30.X),
			Y: float32((float32FromFloat32(1)-v31)*v29.Y) + float32(v31*v30.Y),
		}
		goto _33
	_33:
		(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, p12, v32, int32(b2_colorPaleGreen), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	default:
		break
	}
}

func b2Shape_IsValid(tls *_Stack, id ShapeId) (r uint8) {
	var shape, world uintptr
	var shapeId int32
	_, _, _ = shape, shapeId, world
	if int32(_B2_MAX_WORLDS) <= int32FromUint16(id.World0) {
		return boolUint8(false1 != 0)
	}
	world = uintptr(unsafe.Pointer(&b2_worlds)) + uintptr(id.World0)*1792
	if int32FromUint16((*b2World)(unsafe.Pointer(world)).WorldId) != int32FromUint16(id.World0) {
		// world is free
		return boolUint8(false1 != 0)
	}
	shapeId = id.Index1 - int32(1)
	if shapeId < 0 || (*b2World)(unsafe.Pointer(world)).Shapes.Count <= shapeId {
		return boolUint8(false1 != 0)
	}
	shape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*288
	if (*b2Shape)(unsafe.Pointer(shape)).Id == -int32(1) {
		// shape is free
		return boolUint8(false1 != 0)
	}
	if !((*b2Shape)(unsafe.Pointer(shape)).Id == shapeId) && b2InternalAssertFcn(tls, __ccgo_ts+16034, __ccgo_ts+15342, int32FromInt32(1641)) != 0 {
		__builtin_trap(tls)
	}
	return boolUint8(int32FromUint16(id.Generation) == int32FromUint16((*b2Shape)(unsafe.Pointer(shape)).Generation))
}
