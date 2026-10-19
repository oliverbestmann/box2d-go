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
	var C, transform, v1, v15, v2 Transform
	var aa, dx, dy, invLength, length, radiusA, radiusB, x, y, v44, v68, v72, v73 float32
	var d, n, nonUnitNormal, normal, v12, v13, v7, v8, v9 Vec2
	var duplicate, v70 uint8
	var i, i1, i2, iteration, maxIterations, saveCount, simplexIndex int32
	var output DistanceOutput
	var proxyA, vertex uintptr
	var qr, v11, v3, v4, v5 Rot
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
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, aa, d, duplicate, dx, dy, i, i1, i2, invLength, iteration, length, maxIterations, n, nonUnitNormal, normal, output, proxyA, qr, radiusA, radiusB, saveA, saveB, saveCount, simplexIndex, transform, vertex, vertices, x, y, v1, v11, v12, v13, v15, v2, v3, v4, v44, v5, v68, v7, v70, v72, v73, v8, v9
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
	(*(*ShapeProxy)(unsafe.Add(nil, bp))).Count = (*DistanceInput)(unsafe.Pointer(input)).ProxyB.Count
	(*(*ShapeProxy)(unsafe.Add(nil, bp))).Radius = (*DistanceInput)(unsafe.Pointer(input)).ProxyB.Radius
	i = 0
	for {
		if !(i < (*(*ShapeProxy)(unsafe.Add(nil, bp))).Count) {
			break
		}
		v1 = transform
		v7 = *(*Vec2)(unsafe.Add(nil, input+72+uintptr(i)*8))
		x = float32(v1.Q.C*v7.X) - float32(v1.Q.S*v7.Y) + v1.P.X
		y = float32(v1.Q.S*v7.X) + float32(v1.Q.C*v7.Y) + v1.P.Y
		v8 = Vec2{
			X: x,
			Y: y,
		}
		goto _21
	_21:
		*(*Vec2)(unsafe.Add(nil, bp+uintptr(i)*8)) = v8
		goto _17
	_17:
		;
		i = i + 1
	}
	// Initialize the simplex.
	*(*Simplex)(unsafe.Add(nil, bp+72)) = b2MakeSimplexFromCache(tls, cache, proxyA, bp)
	simplexIndex = 0
	if simplexes != uintptrFromInt32(0) && simplexIndex < simplexCapacity {
		*(*Simplex)(unsafe.Add(nil, simplexes+uintptr(simplexIndex)*112)) = *(*Simplex)(unsafe.Add(nil, bp+72))
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
		saveCount = (*(*Simplex)(unsafe.Add(nil, bp+72))).Count
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
		switch (*(*Simplex)(unsafe.Add(nil, bp+72))).Count {
		case int32(1):
			goto _23
		case int32(2):
			goto _24
		case int32(3):
			goto _25
		default:
			goto _26
		}
	_23:
		;
		v7 = (*(*Simplex)(unsafe.Add(nil, bp+72))).V1.W
		v8 = Vec2{
			X: -v7.X,
			Y: -v7.Y,
		}
		goto _30
	_30:
		d = v8
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
		if (*(*Simplex)(unsafe.Add(nil, bp+72))).Count == int32(3) {
			b2ComputeSimplexWitnessPoints(tls, bp+184, bp+192, bp+72)
			v1 = (*DistanceInput)(unsafe.Pointer(input)).TransformA
			v7 = *(*Vec2)(unsafe.Add(nil, bp+184))
			x = float32(v1.Q.C*v7.X) - float32(v1.Q.S*v7.Y) + v1.P.X
			y = float32(v1.Q.S*v7.X) + float32(v1.Q.C*v7.Y) + v1.P.Y
			v8 = Vec2{
				X: x,
				Y: y,
			}
			goto _37
		_37:
			output.PointA = v8
			v1 = (*DistanceInput)(unsafe.Pointer(input)).TransformA
			v7 = *(*Vec2)(unsafe.Add(nil, bp+192))
			x = float32(v1.Q.C*v7.X) - float32(v1.Q.S*v7.Y) + v1.P.X
			y = float32(v1.Q.S*v7.X) + float32(v1.Q.C*v7.Y) + v1.P.Y
			v8 = Vec2{
				X: x,
				Y: y,
			}
			goto _41
		_41:
			output.PointB = v8
			return output
		}
		if simplexes != uintptrFromInt32(0) && simplexIndex < simplexCapacity {
			*(*Simplex)(unsafe.Add(nil, simplexes+uintptr(simplexIndex)*112)) = *(*Simplex)(unsafe.Add(nil, bp+72))
			simplexIndex = simplexIndex + int32(1)
		}
		// Ensure the search direction is numerically fit.
		v7 = d
		v8 = d
		v44 = float32(v7.X*v8.X) + float32(v7.Y*v8.Y)
		goto _45
	_45:
		if v44 < float32(float32FromFloat32(1.1920928955078125e-07)*float32FromFloat32(1.1920928955078125e-07)) {
			b2ComputeSimplexWitnessPoints(tls, bp+200, bp+208, bp+72)
			v1 = (*DistanceInput)(unsafe.Pointer(input)).TransformA
			v7 = *(*Vec2)(unsafe.Add(nil, bp+200))
			x = float32(v1.Q.C*v7.X) - float32(v1.Q.S*v7.Y) + v1.P.X
			y = float32(v1.Q.S*v7.X) + float32(v1.Q.C*v7.Y) + v1.P.Y
			v8 = Vec2{
				X: x,
				Y: y,
			}
			goto _49
		_49:
			output.PointA = v8
			v1 = (*DistanceInput)(unsafe.Pointer(input)).TransformA
			v7 = *(*Vec2)(unsafe.Add(nil, bp+208))
			x = float32(v1.Q.C*v7.X) - float32(v1.Q.S*v7.Y) + v1.P.X
			y = float32(v1.Q.S*v7.X) + float32(v1.Q.C*v7.Y) + v1.P.Y
			v8 = Vec2{
				X: x,
				Y: y,
			}
			goto _53
		_53:
			output.PointB = v8
			return output
		}
		// Save the normal
		nonUnitNormal = d
		// Compute a tentative new simplex vertex using support points.
		// support = support(a, d) - support(b, -d)
		vertex = vertices[(*(*Simplex)(unsafe.Add(nil, bp+72))).Count]
		(*SimplexVertex)(unsafe.Pointer(vertex)).IndexA = b2FindSupport(tls, proxyA, d)
		(*SimplexVertex)(unsafe.Pointer(vertex)).WA = *(*Vec2)(unsafe.Add(nil, proxyA+uintptr((*SimplexVertex)(unsafe.Pointer(vertex)).IndexA)*8))
		v7 = d
		v8 = Vec2{
			X: -v7.X,
			Y: -v7.Y,
		}
		goto _56
	_56:
		(*SimplexVertex)(unsafe.Pointer(vertex)).IndexB = b2FindSupport(tls, bp, v8)
		(*SimplexVertex)(unsafe.Pointer(vertex)).WB = *(*Vec2)(unsafe.Add(nil, bp+uintptr((*SimplexVertex)(unsafe.Pointer(vertex)).IndexB)*8))
		v7 = (*SimplexVertex)(unsafe.Pointer(vertex)).WA
		v8 = (*SimplexVertex)(unsafe.Pointer(vertex)).WB
		v9 = Vec2{
			X: v7.X - v8.X,
			Y: v7.Y - v8.Y,
		}
		goto _60
	_60:
		(*SimplexVertex)(unsafe.Pointer(vertex)).W = v9
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
		(*(*Simplex)(unsafe.Add(nil, bp+72))).Count += int32(1)
	}
	if simplexes != uintptrFromInt32(0) && simplexIndex < simplexCapacity {
		*(*Simplex)(unsafe.Add(nil, simplexes+uintptr(simplexIndex)*112)) = *(*Simplex)(unsafe.Add(nil, bp+72))
		simplexIndex = simplexIndex + int32(1)
	}
	v7 = nonUnitNormal
	length = sqrtf(tls, float32(v7.X*v7.X)+float32(v7.Y*v7.Y))
	if length < float32FromFloat32(1.1920928955078125e-07) {
		v8 = Vec2{}
		goto _64
	}
	invLength = float32FromFloat32(1) / length
	n = Vec2{
		X: float32(invLength * v7.X),
		Y: float32(invLength * v7.Y),
	}
	v8 = n
	goto _64
_64:
	// Prepare output
	normal = v8
	v9 = normal
	v12 = v9
	v13 = v9
	v44 = float32(v12.X*v13.X) + float32(v12.Y*v13.Y)
	goto _69
_69:
	aa = v44
	v68 = float32FromFloat32(1) - aa
	if v68 < float32FromInt32(0) {
		v73 = -v68
	} else {
		v73 = v68
	}
	v72 = v73
	goto _74
_74:
	v70 = boolUint8(v72 < float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07)))
	goto _71
_71:
	;
	if !(v70 != 0) && b2InternalAssertFcn(tls, __ccgo_ts+4238, __ccgo_ts+4097, int32FromInt32(577)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*DistanceInput)(unsafe.Pointer(input)).TransformA.Q
	v7 = normal
	v8 = Vec2{
		X: float32(v3.C*v7.X) - float32(v3.S*v7.Y),
		Y: float32(v3.S*v7.X) + float32(v3.C*v7.Y),
	}
	goto _79
_79:
	normal = v8
	b2ComputeSimplexWitnessPoints(tls, bp+216, bp+224, bp+72)
	output.Normal = normal
	v7 = *(*Vec2)(unsafe.Add(nil, bp+216))
	v8 = *(*Vec2)(unsafe.Add(nil, bp+224))
	dx = v8.X - v7.X
	dy = v8.Y - v7.Y
	v44 = sqrtf(tls, float32(dx*dx)+float32(dy*dy))
	goto _83
_83:
	output.Distance = v44
	v1 = (*DistanceInput)(unsafe.Pointer(input)).TransformA
	v7 = *(*Vec2)(unsafe.Add(nil, bp+216))
	x = float32(v1.Q.C*v7.X) - float32(v1.Q.S*v7.Y) + v1.P.X
	y = float32(v1.Q.S*v7.X) + float32(v1.Q.C*v7.Y) + v1.P.Y
	v8 = Vec2{
		X: x,
		Y: y,
	}
	goto _87
_87:
	output.PointA = v8
	v1 = (*DistanceInput)(unsafe.Pointer(input)).TransformA
	v7 = *(*Vec2)(unsafe.Add(nil, bp+224))
	x = float32(v1.Q.C*v7.X) - float32(v1.Q.S*v7.Y) + v1.P.X
	y = float32(v1.Q.S*v7.X) + float32(v1.Q.C*v7.Y) + v1.P.Y
	v8 = Vec2{
		X: x,
		Y: y,
	}
	goto _91
_91:
	output.PointB = v8
	output.Iterations = iteration
	output.SimplexCount = simplexIndex
	// Cache the simplex
	b2MakeSimplexCache(tls, cache, bp+72)
	// Apply radii if requested
	if (*DistanceInput)(unsafe.Pointer(input)).UseRadii != 0 && output.Distance > float32(float32FromFloat32(0.1)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
		radiusA = (*DistanceInput)(unsafe.Pointer(input)).ProxyA.Radius
		radiusB = (*DistanceInput)(unsafe.Pointer(input)).ProxyB.Radius
		v44 = float32FromFloat32(0)
		v68 = output.Distance - radiusA - radiusB
		if v44 > v68 {
			v73 = v44
		} else {
			v73 = v68
		}
		v72 = v73
		goto _95
	_95:
		output.Distance = v72
		// Keep closest points on perimeter even if overlapped, this way the points move smoothly.
		v7 = output.PointA
		v44 = radiusA
		v8 = normal
		v9 = Vec2{
			X: v7.X + float32(v44*v8.X),
			Y: v7.Y + float32(v44*v8.Y),
		}
		goto _101
	_101:
		output.PointA = v9
		v7 = output.PointB
		v44 = radiusB
		v8 = normal
		v9 = Vec2{
			X: v7.X - float32(v44*v8.X),
			Y: v7.Y - float32(v44*v8.Y),
		}
		goto _106
	_106:
		output.PointB = v9
	}
	return output
}

// C documentation
//
//	// Shape cast using conservative advancement
func b2ShapeCast(tls *_Stack, input uintptr) (r CastOutput) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var aa, denominator, fraction, linearSlop, target, tolerance, totalRadius, v1, v2, v3, v5 float32
	var c1, c2, delta2, v10, v12, v14, v15, v17, v18, v20, v7, v9 Vec2
	var distanceOutput DistanceOutput
	var iteration, maxIterations int32
	var output CastOutput
	var v27 uint8
	var v33 bool
	var _ /* cache at bp+0 */ SimplexCache
	var _ /* distanceInput at bp+8 */ DistanceInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aa, c1, c2, delta2, denominator, distanceOutput, fraction, iteration, linearSlop, maxIterations, output, target, tolerance, totalRadius, v1, v10, v12, v14, v15, v17, v18, v2, v20, v27, v3, v33, v5, v7, v9
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
	*(*SimplexCache)(unsafe.Add(nil, bp)) = SimplexCache{}
	fraction = float32FromFloat32(0)
	*(*DistanceInput)(unsafe.Add(nil, bp+8)) = DistanceInput{}
	(*(*DistanceInput)(unsafe.Add(nil, bp+8))).ProxyA = (*ShapeCastPairInput)(unsafe.Pointer(input)).ProxyA
	(*(*DistanceInput)(unsafe.Add(nil, bp+8))).ProxyB = (*ShapeCastPairInput)(unsafe.Pointer(input)).ProxyB
	(*(*DistanceInput)(unsafe.Add(nil, bp+8))).TransformA = (*ShapeCastPairInput)(unsafe.Pointer(input)).TransformA
	(*(*DistanceInput)(unsafe.Add(nil, bp+8))).TransformB = (*ShapeCastPairInput)(unsafe.Pointer(input)).TransformB
	(*(*DistanceInput)(unsafe.Add(nil, bp+8))).UseRadii = boolUint8(false1 != 0)
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
					v1 = (*ShapeCastPairInput)(unsafe.Pointer(input)).ProxyA.Radius
					v9 = distanceOutput.Normal
					v10 = Vec2{
						X: v7.X + float32(v1*v9.X),
						Y: v7.Y + float32(v1*v9.Y),
					}
					goto _11
				_11:
					// Compute a common point
					c1 = v10
					v12 = distanceOutput.PointB
					v2 = -(*ShapeCastPairInput)(unsafe.Pointer(input)).ProxyB.Radius
					v14 = distanceOutput.Normal
					v15 = Vec2{
						X: v12.X + float32(v2*v14.X),
						Y: v12.Y + float32(v2*v14.Y),
					}
					goto _16
				_16:
					c2 = v15
					v17 = c1
					v18 = c2
					v3 = float32FromFloat32(0.5)
					v20 = Vec2{
						X: float32((float32FromFloat32(1)-v3)*v17.X) + float32(v3*v18.X),
						Y: float32((float32FromFloat32(1)-v3)*v17.Y) + float32(v3*v18.Y),
					}
					goto _21
				_21:
					output.Point = v20
					return output
				}
			} else {
				// Regular hit
				if v33 = distanceOutput.Distance > float32FromFloat32(0); v33 {
					v7 = distanceOutput.Normal
					v9 = v7
					v10 = v7
					v1 = float32(v9.X*v10.X) + float32(v9.Y*v10.Y)
					goto _26
				_26:
					aa = v1
					v2 = float32FromFloat32(1) - aa
					if v2 < float32FromInt32(0) {
						v5 = -v2
					} else {
						v5 = v2
					}
					v3 = v5
					goto _31
				_31:
					v27 = boolUint8(v3 < float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07)))
					goto _28
				_28:
				}
				if !(v33 && v27 != 0) && b2InternalAssertFcn(tls, __ccgo_ts+4282, __ccgo_ts+4097, int32FromInt32(664)) != 0 {
					__builtin_trap(tls)
				}
				output.Fraction = fraction
				v7 = distanceOutput.PointA
				v1 = (*ShapeCastPairInput)(unsafe.Pointer(input)).ProxyA.Radius
				v9 = distanceOutput.Normal
				v10 = Vec2{
					X: v7.X + float32(v1*v9.X),
					Y: v7.Y + float32(v1*v9.Y),
				}
				goto _38
			_38:
				output.Point = v10
				output.Normal = distanceOutput.Normal
				output.Hit = boolUint8(true1 != 0)
				return output
			}
		}
		if !(distanceOutput.Distance > float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+4356, __ccgo_ts+4097, int32FromInt32(673)) != 0 {
			__builtin_trap(tls)
		}
		v7 = distanceOutput.Normal
		v9 = v7
		v10 = v7
		v1 = float32(v9.X*v10.X) + float32(v9.Y*v10.Y)
		goto _43
	_43:
		aa = v1
		v2 = float32FromFloat32(1) - aa
		if v2 < float32FromInt32(0) {
			v5 = -v2
		} else {
			v5 = v2
		}
		v3 = v5
		goto _48
	_48:
		v27 = boolUint8(v3 < float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07)))
		goto _45
	_45:
		;
		if !(v27 != 0) && b2InternalAssertFcn(tls, __ccgo_ts+4387, __ccgo_ts+4097, int32FromInt32(674)) != 0 {
			__builtin_trap(tls)
		}
		v7 = delta2
		v9 = distanceOutput.Normal
		v1 = float32(v7.X*v9.X) + float32(v7.Y*v9.Y)
		goto _53
	_53:
		// Check if shapes are approaching each other
		denominator = v1
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
		v7 = (*ShapeCastPairInput)(unsafe.Pointer(input)).TransformB.P
		v1 = fraction
		v9 = delta2
		v10 = Vec2{
			X: v7.X + float32(v1*v9.X),
			Y: v7.Y + float32(v1*v9.Y),
		}
		goto _58
	_58:
		(*(*DistanceInput)(unsafe.Add(nil, bp+8))).TransformB.P = v10
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
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).ProxyA = b2MakeProxy(tls, shape, int32(1), (*Circle)(unsafe.Pointer(shape)).Radius)
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).ProxyB = (*ShapeCastInput)(unsafe.Pointer(input)).Proxy
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).TransformA = b2Transform_identity
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).TransformB = b2Transform_identity
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).TranslationB = (*ShapeCastInput)(unsafe.Pointer(input)).Translation
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).MaxFraction = (*ShapeCastInput)(unsafe.Pointer(input)).MaxFraction
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).CanEncroach = (*ShapeCastInput)(unsafe.Pointer(input)).CanEncroach
	output = b2ShapeCast(tls, bp)
	return output
}

func b2ShapeCastCapsule(tls *_Stack, input uintptr, shape uintptr) (r CastOutput) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var output CastOutput
	var _ /* pairInput at bp+0 */ ShapeCastPairInput
	_ = output
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).ProxyA = b2MakeProxy(tls, shape, int32(2), (*Capsule)(unsafe.Pointer(shape)).Radius)
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).ProxyB = (*ShapeCastInput)(unsafe.Pointer(input)).Proxy
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).TransformA = b2Transform_identity
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).TransformB = b2Transform_identity
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).TranslationB = (*ShapeCastInput)(unsafe.Pointer(input)).Translation
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).MaxFraction = (*ShapeCastInput)(unsafe.Pointer(input)).MaxFraction
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).CanEncroach = (*ShapeCastInput)(unsafe.Pointer(input)).CanEncroach
	output = b2ShapeCast(tls, bp)
	return output
}

func b2ShapeCastSegment(tls *_Stack, input uintptr, shape uintptr) (r CastOutput) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var output CastOutput
	var _ /* pairInput at bp+0 */ ShapeCastPairInput
	_ = output
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).ProxyA = b2MakeProxy(tls, shape, int32(2), float32FromFloat32(0))
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).ProxyB = (*ShapeCastInput)(unsafe.Pointer(input)).Proxy
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).TransformA = b2Transform_identity
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).TransformB = b2Transform_identity
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).TranslationB = (*ShapeCastInput)(unsafe.Pointer(input)).Translation
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).MaxFraction = (*ShapeCastInput)(unsafe.Pointer(input)).MaxFraction
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).CanEncroach = (*ShapeCastInput)(unsafe.Pointer(input)).CanEncroach
	output = b2ShapeCast(tls, bp)
	return output
}

func b2ShapeCastPolygon(tls *_Stack, input uintptr, shape uintptr) (r CastOutput) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var output CastOutput
	var _ /* pairInput at bp+0 */ ShapeCastPairInput
	_ = output
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).ProxyA = b2MakeProxy(tls, shape, (*Polygon)(unsafe.Pointer(shape)).Count, (*Polygon)(unsafe.Pointer(shape)).Radius)
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).ProxyB = (*ShapeCastInput)(unsafe.Pointer(input)).Proxy
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).TransformA = b2Transform_identity
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).TransformB = b2Transform_identity
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).TranslationB = (*ShapeCastInput)(unsafe.Pointer(input)).Translation
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).MaxFraction = (*ShapeCastInput)(unsafe.Pointer(input)).MaxFraction
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).CanEncroach = (*ShapeCastInput)(unsafe.Pointer(input)).CanEncroach
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
	var headShape, shape, v1, v3 uintptr
	var newCapacity, newCapacity1, shapeId, v2 int32
	var proxyType b2BodyType1
	var sensor b2Sensor
	_, _, _, _, _, _, _, _, _, _ = headShape, newCapacity, newCapacity1, proxyType, sensor, shape, shapeId, v1, v2, v3
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
		*(*b2Shape)(unsafe.Add(nil, (*b2ShapeArray)(unsafe.Pointer(v1)).Data+uintptr((*b2ShapeArray)(unsafe.Pointer(v1)).Count)*288)) = b2Shape{}
		*(*int32)(unsafe.Add(nil, v1+8)) += int32(1)
	} else {
		if !((*(*b2Shape)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Shapes.Data+uintptr(shapeId)*288))).Id == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+11119, __ccgo_ts+11037, int32FromInt32(70)) != 0 {
			__builtin_trap(tls)
		}
	}
	v1 = world + 1248
	v2 = shapeId
	if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
	goto _6
_6:
	shape = v3
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
_7:
	;
	(*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo19_132.Capsule = *(*Capsule)(unsafe.Add(nil, geometry))
	goto _13
_8:
	;
	*(*Circle)(unsafe.Add(unsafe.Pointer(shape), 132)) = *(*Circle)(unsafe.Add(nil, geometry))
	goto _13
_9:
	;
	*(*Polygon)(unsafe.Add(unsafe.Pointer(shape), 132)) = *(*Polygon)(unsafe.Add(nil, geometry))
	goto _13
_10:
	;
	*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 132)) = *(*Segment)(unsafe.Add(nil, geometry))
	goto _13
_11:
	;
	*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 132)) = *(*ChainSegment)(unsafe.Add(nil, geometry))
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
	v1 = shape + 276
	*(*uint16_t)(unsafe.Pointer(v1)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v1))) + int32FromInt32(1))
	if (*b2Body)(unsafe.Pointer(body)).SetIndex != int32(b2_disabledSet) {
		proxyType = (*b2Body)(unsafe.Pointer(body)).Type1
		b2CreateShapeProxy(tls, shape, world+40, proxyType, transform, boolUint8((*ShapeDef)(unsafe.Pointer(def)).InvokeContactCreation != 0 || (*ShapeDef)(unsafe.Pointer(def)).IsSensor != 0))
	}
	// Add to shape doubly linked list
	if (*b2Body)(unsafe.Pointer(body)).HeadShapeId != -int32(1) {
		v1 = world + 1248
		v2 = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
		if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
		goto _21
	_21:
		headShape = v3
		(*b2Shape)(unsafe.Pointer(headShape)).PrevShapeId = shapeId
	}
	(*b2Shape)(unsafe.Pointer(shape)).PrevShapeId = -int32(1)
	(*b2Shape)(unsafe.Pointer(shape)).NextShapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	(*b2Body)(unsafe.Pointer(body)).HeadShapeId = shapeId
	*(*int32)(unsafe.Add(nil, body+60)) += int32(1)
	if (*ShapeDef)(unsafe.Pointer(def)).IsSensor != 0 {
		(*b2Shape)(unsafe.Pointer(shape)).SensorIndex = (*b2World)(unsafe.Pointer(world)).Sensors.Count
		sensor = b2Sensor{
			Overlaps1: b2ShapeRefArray_Create(tls, int32(16)),
			Overlaps2: b2ShapeRefArray_Create(tls, int32(16)),
			ShapeId:   shapeId,
		}
		v1 = world + 1280
		if (*b2SensorArray)(unsafe.Pointer(v1)).Count == (*b2SensorArray)(unsafe.Pointer(v1)).Capacity {
			if (*b2SensorArray)(unsafe.Pointer(v1)).Capacity < int32(2) {
				v2 = int32(2)
			} else {
				v2 = (*b2SensorArray)(unsafe.Pointer(v1)).Capacity + (*b2SensorArray)(unsafe.Pointer(v1)).Capacity>>int32(1)
			}
			newCapacity1 = v2
			b2SensorArray_Reserve(tls, v1, newCapacity1)
		}
		*(*b2Sensor)(unsafe.Add(nil, (*b2SensorArray)(unsafe.Pointer(v1)).Data+uintptr((*b2SensorArray)(unsafe.Pointer(v1)).Count)*40)) = sensor
		*(*int32)(unsafe.Add(nil, v1+8)) += int32(1)
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
		*(*Circle)(unsafe.Add(nil, bp)) = Circle{
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
	var contact, movedSensor, nextShape, otherSensorShape, prevShape, ref, sensor, v1, v3, v5, v7, v9 uintptr
	var contactId, contactKey, edgeIndex, i, movedIndex, movedIndex1, newCapacity, shapeId, v10, v14, v2, v6 int32
	var event SensorEndTouchEvent
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = contact, contactId, contactKey, edgeIndex, event, i, movedIndex, movedIndex1, movedSensor, newCapacity, nextShape, otherSensorShape, prevShape, ref, sensor, shapeId, v1, v10, v14, v2, v3, v5, v6, v7, v9
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
		v1 = world + 1248
		v2 = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
		if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
		goto _8
	_8:
		nextShape = v3
		(*b2Shape)(unsafe.Pointer(nextShape)).PrevShapeId = (*b2Shape)(unsafe.Pointer(shape)).PrevShapeId
	}
	if shapeId == (*b2Body)(unsafe.Pointer(body)).HeadShapeId {
		(*b2Body)(unsafe.Pointer(body)).HeadShapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
	}
	*(*int32)(unsafe.Add(nil, body+60)) -= int32(1)
	// Remove from broad-phase.
	b2DestroyShapeProxy(tls, shape, world+40)
	// Destroy any contacts associated with the shape.
	contactKey = (*b2Body)(unsafe.Pointer(body)).HeadContactKey
	for contactKey != -int32(1) {
		contactId = contactKey >> int32(1)
		edgeIndex = contactKey & int32(1)
		v1 = world + 1144
		v2 = contactId
		if !(0 <= v2 && v2 < (*b2ContactArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2ContactArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*68
		goto _12
	_12:
		contact = v3
		contactKey = (*(*b2ContactEdge)(unsafe.Add(nil, contact+12+uintptr(edgeIndex)*12))).NextKey
		if (*b2Contact)(unsafe.Pointer(contact)).ShapeIdA == shapeId || (*b2Contact)(unsafe.Pointer(contact)).ShapeIdB == shapeId {
			b2DestroyContact(tls, world, contact, wakeBodies)
		}
	}
	if (*b2Shape)(unsafe.Pointer(shape)).SensorIndex != -int32(1) {
		v1 = world + 1280
		v2 = (*b2Shape)(unsafe.Pointer(shape)).SensorIndex
		if !(0 <= v2 && v2 < (*b2SensorArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+10835, int32FromInt32(35)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2SensorArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*40
		goto _16
	_16:
		sensor = v3
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
			v5 = world + 1376 + uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16
			if (*b2SensorEndTouchEventArray)(unsafe.Pointer(v5)).Count == (*b2SensorEndTouchEventArray)(unsafe.Pointer(v5)).Capacity {
				if (*b2SensorEndTouchEventArray)(unsafe.Pointer(v5)).Capacity < int32(2) {
					v6 = int32(2)
				} else {
					v6 = (*b2SensorEndTouchEventArray)(unsafe.Pointer(v5)).Capacity + (*b2SensorEndTouchEventArray)(unsafe.Pointer(v5)).Capacity>>int32(1)
				}
				newCapacity = v6
				b2SensorEndTouchEventArray_Reserve(tls, v5, newCapacity)
			}
			*(*SensorEndTouchEvent)(unsafe.Add(nil, (*b2SensorEndTouchEventArray)(unsafe.Pointer(v5)).Data+uintptr((*b2SensorEndTouchEventArray)(unsafe.Pointer(v5)).Count)*16)) = event
			*(*int32)(unsafe.Add(nil, v5+8)) += int32(1)
			goto _17
		_17:
			;
			i = i + 1
		}
		// Destroy sensor
		b2ShapeRefArray_Destroy(tls, sensor)
		b2ShapeRefArray_Destroy(tls, sensor+16)
		v1 = world + 1280
		v2 = (*b2Shape)(unsafe.Pointer(shape)).SensorIndex
		if !(0 <= v2 && v2 < (*b2SensorArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+10835, int32FromInt32(35)) != 0 {
			__builtin_trap(tls)
		}
		movedIndex = -int32(1)
		if v2 != (*b2SensorArray)(unsafe.Pointer(v1)).Count-int32FromInt32(1) {
			movedIndex = (*b2SensorArray)(unsafe.Pointer(v1)).Count - int32(1)
			*(*b2Sensor)(unsafe.Add(nil, (*b2SensorArray)(unsafe.Pointer(v1)).Data+uintptr(v2)*40)) = *(*b2Sensor)(unsafe.Add(nil, (*b2SensorArray)(unsafe.Pointer(v1)).Data+uintptr(movedIndex)*40))
		}
		*(*int32)(unsafe.Add(nil, v1+8)) -= int32(1)
		v6 = movedIndex
		goto _23
	_23:
		movedIndex1 = v6
		if movedIndex1 != -int32(1) {
			v3 = world + 1280
			v10 = (*b2Shape)(unsafe.Pointer(shape)).SensorIndex
			if !(0 <= v10 && v10 < (*b2SensorArray)(unsafe.Pointer(v3)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+10835, int32FromInt32(35)) != 0 {
				__builtin_trap(tls)
			}
			v5 = (*b2SensorArray)(unsafe.Pointer(v3)).Data + uintptr(v10)*40
			goto _27
		_27:
			// Fixup moved sensor
			movedSensor = v5
			v7 = world + 1248
			v14 = (*b2Sensor)(unsafe.Pointer(movedSensor)).ShapeId
			if !(0 <= v14 && v14 < (*b2ShapeArray)(unsafe.Pointer(v7)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
				__builtin_trap(tls)
			}
			v9 = (*b2ShapeArray)(unsafe.Pointer(v7)).Data + uintptr(v14)*288
			goto _31
		_31:
			otherSensorShape = v9
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
}

func b2GetShapeCentroid(tls *_Stack, shape uintptr) (r Vec2) {
	var v1, v2, v4 Vec2
	var v3 float32
	_, _, _, _ = v1, v2, v3, v4
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
		v1 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 132))).Point1
		v2 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 132))).Point2
		v3 = float32FromFloat32(0.5)
		v4 = Vec2{
			X: float32((float32FromFloat32(1)-v3)*v1.X) + float32(v3*v2.X),
			Y: float32((float32FromFloat32(1)-v3)*v1.Y) + float32(v3*v2.Y),
		}
		goto _10
	_10:
		return v4
	case int32(b2_chainSegmentShape):
		v1 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 132))).Segment.Point1
		v2 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 132))).Segment.Point2
		v3 = float32FromFloat32(0.5)
		v4 = Vec2{
			X: float32((float32FromFloat32(1)-v3)*v1.X) + float32(v3*v2.X),
			Y: float32((float32FromFloat32(1)-v3)*v1.Y) + float32(v3*v2.Y),
		}
		goto _15
	_15:
		return v4
	default:
		return b2Vec2_zero
	}
}

// C documentation
//
//	// todo_erin maybe compute this on shape creation
func b2GetShapePerimeter(tls *_Stack, shape uintptr) (r float32) {
	var count, i int32
	var next, prev, v1, v2, v3, v5 Vec2
	var perimeter, v6 float32
	var points uintptr
	_, _, _, _, _, _, _, _, _, _, _ = count, i, next, perimeter, points, prev, v1, v2, v3, v5, v6
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
		prev = *(*Vec2)(unsafe.Add(nil, points+uintptr(count-int32(1))*8))
		i = 0
		for {
			if !(i < count) {
				break
			}
			next = *(*Vec2)(unsafe.Add(nil, points+uintptr(i)*8))
			v1 = next
			v2 = prev
			v3 = Vec2{
				X: v1.X - v2.X,
				Y: v1.Y - v2.Y,
			}
			goto _12
		_12:
			v5 = v3
			v6 = sqrtf(tls, float32(v5.X*v5.X)+float32(v5.Y*v5.Y))
			goto _15
		_15:
			perimeter = perimeter + v6
			prev = next
			goto _8
		_8:
			;
			i = i + 1
		}
		return perimeter
	case int32(b2_segmentShape):
		v1 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 132))).Point1
		v2 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 132))).Point2
		v3 = Vec2{
			X: v1.X - v2.X,
			Y: v1.Y - v2.Y,
		}
		goto _19
	_19:
		v5 = v3
		v6 = sqrtf(tls, float32(v5.X*v5.X)+float32(v5.Y*v5.Y))
		goto _22
	_22:
		return float32(float32FromFloat32(2) * v6)
	case int32(b2_chainSegmentShape):
		v1 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 132))).Segment.Point1
		v2 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 132))).Segment.Point2
		v3 = Vec2{
			X: v1.X - v2.X,
			Y: v1.Y - v2.Y,
		}
		goto _26
	_26:
		v5 = v3
		v6 = sqrtf(tls, float32(v5.X*v5.X)+float32(v5.Y*v5.Y))
		goto _29
	_29:
		return float32(float32FromFloat32(2) * v6)
	default:
		return float32FromFloat32(0)
	}
}

// C documentation
//
//	// This projects the shape perimeter onto an infinite line
func b2GetShapeProjectedPerimeter(tls *_Stack, shape uintptr, line Vec2) (r float32) {
	var axis, v1, v2, v3, v5, v6 Vec2
	var count, i int32
	var lower, projectedLength, upper, value, value1, value11, value2, value21, v10, v12, v15, v7, v9 float32
	var points uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axis, count, i, lower, points, projectedLength, upper, value, value1, value11, value2, value21, v1, v10, v12, v15, v2, v3, v5, v6, v7, v9
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
		v1 = *(*Vec2)(unsafe.Add(nil, points))
		v2 = line
		v7 = float32(v1.X*v2.X) + float32(v1.Y*v2.Y)
		goto _16
	_16:
		value = v7
		lower = value
		upper = value
		i = int32(1)
		for {
			if !(i < count) {
				break
			}
			v3 = *(*Vec2)(unsafe.Add(nil, points+uintptr(i)*8))
			v5 = line
			v9 = float32(v3.X*v5.X) + float32(v3.Y*v5.Y)
			goto _21
		_21:
			value = v9
			v7 = lower
			v9 = value
			if v7 < v9 {
				v12 = v7
			} else {
				v12 = v9
			}
			v10 = v12
			goto _25
		_25:
			lower = v10
			v7 = upper
			v9 = value
			if v7 > v9 {
				v12 = v7
			} else {
				v12 = v9
			}
			v10 = v12
			goto _30
		_30:
			upper = v10
			goto _17
		_17:
			;
			i = i + 1
		}
		return upper - lower + float32(float32FromFloat32(2)*(*(*Polygon)(unsafe.Add(unsafe.Pointer(shape), 132))).Radius)
	case int32(b2_segmentShape):
		v1 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 132))).Point1
		v2 = line
		v7 = float32(v1.X*v2.X) + float32(v1.Y*v2.Y)
		goto _35
	_35:
		value1 = v7
		v3 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 132))).Point2
		v5 = line
		v9 = float32(v3.X*v5.X) + float32(v3.Y*v5.Y)
		goto _39
	_39:
		value2 = v9
		v10 = value2 - value1
		if v10 < float32FromInt32(0) {
			v15 = -v10
		} else {
			v15 = v10
		}
		v12 = v15
		goto _42
	_42:
		return v12
	case int32(b2_chainSegmentShape):
		v1 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 132))).Segment.Point1
		v2 = line
		v7 = float32(v1.X*v2.X) + float32(v1.Y*v2.Y)
		goto _47
	_47:
		value11 = v7
		v3 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 132))).Segment.Point2
		v5 = line
		v9 = float32(v3.X*v5.X) + float32(v3.Y*v5.Y)
		goto _51
	_51:
		value21 = v9
		v10 = value21 - value11
		if v10 < float32FromInt32(0) {
			v15 = -v10
		} else {
			v15 = v10
		}
		v12 = v15
		goto _54
	_54:
		return v12
	default:
		return float32FromFloat32(0)
	}
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
}

func b2ComputeShapeExtent(tls *_Stack, shape uintptr, localCenter Vec2) (r b2ShapeExtent) {
	var c1, c11, c12, c2, c21, c22, v2, v1, v12, v21, v3, v5, v6, v7, v9 Vec2
	var count, i int32
	var distanceSqr, maxExtentSqr, minExtent, planeOffset, radius, radius1, v10, v13, v15, v16, v17, v19 float32
	var extent b2ShapeExtent
	var poly uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c1, c11, c12, c2, c21, c22, count, distanceSqr, extent, i, maxExtentSqr, minExtent, planeOffset, poly, radius, radius1, v2, v1, v10, v12, v13, v15, v16, v17, v19, v21, v3, v5, v6, v7, v9
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
		v1 = (*(*Circle)(unsafe.Add(unsafe.Pointer(shape), 132))).Center
		v21 = localCenter
		v3 = Vec2{
			X: v1.X - v21.X,
			Y: v1.Y - v21.Y,
		}
		goto _23
	_23:
		v5 = v3
		v10 = sqrtf(tls, float32(v5.X*v5.X)+float32(v5.Y*v5.Y))
		goto _26
	_26:
		extent.MaxExtent = v10 + radius1
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
			v2 = *(*Vec2)(unsafe.Add(nil, poly+uintptr(i)*8))
			v1 = v2
			v21 = (*Polygon)(unsafe.Pointer(poly)).Centroid
			v3 = Vec2{
				X: v1.X - v21.X,
				Y: v1.Y - v21.Y,
			}
			goto _31
		_31:
			v5 = *(*Vec2)(unsafe.Add(nil, poly+64+uintptr(i)*8))
			v6 = v3
			v10 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y)
			goto _35
		_35:
			planeOffset = v10
			v13 = minExtent
			v15 = planeOffset
			if v13 < v15 {
				v17 = v13
			} else {
				v17 = v15
			}
			v16 = v17
			goto _39
		_39:
			minExtent = v16
			v1 = v2
			v21 = localCenter
			v3 = Vec2{
				X: v1.X - v21.X,
				Y: v1.Y - v21.Y,
			}
			goto _44
		_44:
			v5 = v3
			v10 = float32(v5.X*v5.X) + float32(v5.Y*v5.Y)
			goto _47
		_47:
			distanceSqr = v10
			v13 = maxExtentSqr
			v15 = distanceSqr
			if v13 > v15 {
				v17 = v13
			} else {
				v17 = v15
			}
			v16 = v17
			goto _51
		_51:
			maxExtentSqr = v16
			goto _27
		_27:
			;
			i = i + 1
		}
		extent.MinExtent = minExtent + (*Polygon)(unsafe.Pointer(poly)).Radius
		extent.MaxExtent = sqrtf(tls, maxExtentSqr) + (*Polygon)(unsafe.Pointer(poly)).Radius
	case int32(b2_segmentShape):
		extent.MinExtent = float32FromFloat32(0)
		v1 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 132))).Point1
		v21 = localCenter
		v3 = Vec2{
			X: v1.X - v21.X,
			Y: v1.Y - v21.Y,
		}
		goto _56
	_56:
		c11 = v3
		v5 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 132))).Point2
		v6 = localCenter
		v7 = Vec2{
			X: v5.X - v6.X,
			Y: v5.Y - v6.Y,
		}
		goto _60
	_60:
		c21 = v7
		v9 = c11
		v10 = float32(v9.X*v9.X) + float32(v9.Y*v9.Y)
		goto _63
	_63:
		v12 = c21
		v13 = float32(v12.X*v12.X) + float32(v12.Y*v12.Y)
		goto _66
	_66:
		v15 = v10
		v16 = v13
		if v15 > v16 {
			v19 = v15
		} else {
			v19 = v16
		}
		v17 = v19
		goto _70
	_70:
		extent.MaxExtent = sqrtf(tls, v17)
	case int32(b2_chainSegmentShape):
		extent.MinExtent = float32FromFloat32(0)
		v1 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 132))).Segment.Point1
		v21 = localCenter
		v3 = Vec2{
			X: v1.X - v21.X,
			Y: v1.Y - v21.Y,
		}
		goto _75
	_75:
		c12 = v3
		v5 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 132))).Segment.Point2
		v6 = localCenter
		v7 = Vec2{
			X: v5.X - v6.X,
			Y: v5.Y - v6.Y,
		}
		goto _79
	_79:
		c22 = v7
		v9 = c12
		v10 = float32(v9.X*v9.X) + float32(v9.Y*v9.Y)
		goto _82
	_82:
		v12 = c22
		v13 = float32(v12.X*v12.X) + float32(v12.Y*v12.Y)
		goto _85
	_85:
		v15 = v10
		v16 = v13
		if v15 > v16 {
			v19 = v15
		} else {
			v19 = v16
		}
		v17 = v19
		goto _89
	_89:
		extent.MaxExtent = sqrtf(tls, v17)
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
	var v1 Transform
	var v2, v3 Vec2
	var v5 Rot
	var _ /* localInput at bp+0 */ RayCastInput
	_, _, _, _, _, _, _, _, _ = output, vx, vy, x, y, v1, v2, v3, v5
	*(*RayCastInput)(unsafe.Add(nil, bp)) = *(*RayCastInput)(unsafe.Add(nil, input))
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
	(*(*RayCastInput)(unsafe.Add(nil, bp))).Origin = v3
	v5 = transform.Q
	v2 = (*RayCastInput)(unsafe.Pointer(input)).Translation
	v3 = Vec2{
		X: float32(v5.C*v2.X) + float32(v5.S*v2.Y),
		Y: float32(-v5.S*v2.X) + float32(v5.C*v2.Y),
	}
	goto _8
_8:
	(*(*RayCastInput)(unsafe.Add(nil, bp))).Translation = v3
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
	v1 = transform
	v2 = output.Point
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _12
_12:
	output.Point = v3
	v5 = transform.Q
	v2 = output.Normal
	v3 = Vec2{
		X: float32(v5.C*v2.X) - float32(v5.S*v2.Y),
		Y: float32(v5.S*v2.X) + float32(v5.C*v2.Y),
	}
	goto _16
_16:
	output.Normal = v3
	return output
}

func b2ShapeCastShape(tls *_Stack, input uintptr, shape uintptr, transform Transform) (r CastOutput) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var i int32
	var output CastOutput
	var vx, vy, x, y float32
	var v2 Transform
	var v3, v4 Vec2
	var v6 Rot
	var _ /* localInput at bp+0 */ ShapeCastInput
	_, _, _, _, _, _, _, _, _, _ = i, output, vx, vy, x, y, v2, v3, v4, v6
	*(*ShapeCastInput)(unsafe.Add(nil, bp)) = *(*ShapeCastInput)(unsafe.Add(nil, input))
	i = 0
	for {
		if !(i < (*(*ShapeCastInput)(unsafe.Add(nil, bp))).Proxy.Count) {
			break
		}
		v2 = transform
		v3 = *(*Vec2)(unsafe.Add(nil, input+uintptr(i)*8))
		vx = v3.X - v2.P.X
		vy = v3.Y - v2.P.Y
		v4 = Vec2{
			X: float32(v2.Q.C*vx) + float32(v2.Q.S*vy),
			Y: float32(-v2.Q.S*vx) + float32(v2.Q.C*vy),
		}
		goto _5
	_5:
		*(*Vec2)(unsafe.Add(nil, bp+uintptr(i)*8)) = v4
		goto _1
	_1:
		;
		i = i + 1
	}
	v6 = transform.Q
	v3 = (*ShapeCastInput)(unsafe.Pointer(input)).Translation
	v4 = Vec2{
		X: float32(v6.C*v3.X) + float32(v6.S*v3.Y),
		Y: float32(-v6.S*v3.X) + float32(v6.C*v3.Y),
	}
	goto _9
_9:
	(*(*ShapeCastInput)(unsafe.Add(nil, bp))).Translation = v4
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
	v2 = transform
	v3 = output.Point
	x = float32(v2.Q.C*v3.X) - float32(v2.Q.S*v3.Y) + v2.P.X
	y = float32(v2.Q.S*v3.X) + float32(v2.Q.C*v3.Y) + v2.P.Y
	v4 = Vec2{
		X: x,
		Y: y,
	}
	goto _13
_13:
	output.Point = v4
	v6 = transform.Q
	v3 = output.Normal
	v4 = Vec2{
		X: float32(v6.C*v3.X) - float32(v6.S*v3.Y),
		Y: float32(v6.S*v3.X) + float32(v6.C*v3.Y),
	}
	goto _17
_17:
	output.Normal = v4
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
}

// C documentation
//
//	// todo_erin untested
func b2Shape_RayCast(tls *_Stack, shapeId ShapeId, input uintptr) (r CastOutput) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var output CastOutput
	var shape, world uintptr
	var transform, v1 Transform
	var vx, vy, x, y float32
	var v5 Rot
	var v2, v3 Vec2
	var _ /* localInput at bp+0 */ RayCastInput
	_, _, _, _, _, _, _, _, _, _, _, _ = output, shape, transform, vx, vy, world, x, y, v1, v2, v3, v5
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
	(*(*RayCastInput)(unsafe.Add(nil, bp))).Origin = v3
	v5 = transform.Q
	v2 = (*RayCastInput)(unsafe.Pointer(input)).Translation
	v3 = Vec2{
		X: float32(v5.C*v2.X) + float32(v5.S*v2.Y),
		Y: float32(-v5.S*v2.X) + float32(v5.C*v2.Y),
	}
	goto _8
_8:
	(*(*RayCastInput)(unsafe.Add(nil, bp))).Translation = v3
	(*(*RayCastInput)(unsafe.Add(nil, bp))).MaxFraction = (*RayCastInput)(unsafe.Pointer(input)).MaxFraction
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
		v5 = transform.Q
		v2 = output.Normal
		v3 = Vec2{
			X: float32(v5.C*v2.X) - float32(v5.S*v2.Y),
			Y: float32(v5.S*v2.X) + float32(v5.C*v2.Y),
		}
		goto _22
	_22:
		output.Normal = v3
		v1 = transform
		v2 = output.Point
		x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
		y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
		v3 = Vec2{
			X: x,
			Y: y,
		}
		goto _26
	_26:
		output.Point = v3
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
	*(*Circle)(unsafe.Add(unsafe.Pointer(shape), 132)) = *(*Circle)(unsafe.Add(nil, circle))
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
	(*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo19_132.Capsule = *(*Capsule)(unsafe.Add(nil, capsule))
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
	*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 132)) = *(*Segment)(unsafe.Add(nil, segment))
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
	*(*Polygon)(unsafe.Add(unsafe.Pointer(shape), 132)) = *(*Polygon)(unsafe.Add(nil, polygon))
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
			(*(*ContactData)(unsafe.Add(nil, contactData+uintptr(index2)*128))).ShapeIdA = ShapeId{
				Index1:     (*b2Shape)(unsafe.Pointer(shapeA)).Id + int32(1),
				World0:     shapeId.World0,
				Generation: (*b2Shape)(unsafe.Pointer(shapeA)).Generation,
			}
			(*(*ContactData)(unsafe.Add(nil, contactData+uintptr(index2)*128))).ShapeIdB = ShapeId{
				Index1:     (*b2Shape)(unsafe.Pointer(shapeB)).Id + int32(1),
				World0:     shapeId.World0,
				Generation: (*b2Shape)(unsafe.Pointer(shapeB)).Generation,
			}
			contactSim = b2GetContactSim(tls, world, contact)
			(*(*ContactData)(unsafe.Add(nil, contactData+uintptr(index2)*128))).Manifold = (*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold
			index2 = index2 + int32(1)
		}
		contactKey = (*(*b2ContactEdge)(unsafe.Add(nil, contact+12+uintptr(edgeIndex)*12))).NextKey
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
		*(*ShapeId)(unsafe.Add(nil, overlaps+uintptr(i)*8)) = ShapeId{
			Index1:     (*(*b2ShapeRef)(unsafe.Add(nil, refs+uintptr(i)*8))).ShapeId + int32(1),
			World0:     shapeId.World0,
			Generation: (*(*b2ShapeRef)(unsafe.Add(nil, refs+uintptr(i)*8))).Generation,
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
	(*(*DistanceInput)(unsafe.Add(nil, bp+8))).ProxyA = b2MakeShapeDistanceProxy(tls, shape)
	(*(*DistanceInput)(unsafe.Add(nil, bp+8))).ProxyB = b2MakeProxy(tls, bp, int32(1), float32FromFloat32(0))
	(*(*DistanceInput)(unsafe.Add(nil, bp+8))).TransformA = transform
	(*(*DistanceInput)(unsafe.Add(nil, bp+8))).TransformB = b2Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp+8))).UseRadii = boolUint8(true1 != 0)
	*(*SimplexCache)(unsafe.Add(nil, bp+192)) = SimplexCache{}
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
	var p1, p11, p12, p2, p21, p22, v2, v3, v6, v7 Vec2
	var x, y, v31 float32
	var v1, v5 Transform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = capsule, circle, p1, p11, p12, p2, p21, p22, poly, segment, segment1, x, y, v1, v2, v3, v31, v5, v6, v7
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
		v1 = xf
		v2 = (*Circle)(unsafe.Pointer(circle)).Center
		x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
		y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
		v3 = Vec2{
			X: x,
			Y: y,
		}
		goto _12
	_12:
		xf.P = v3
		(*(*func(*_Stack, Transform, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSolidCircleFcn})))(tls, xf, (*Circle)(unsafe.Pointer(circle)).Radius, color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	case int32(b2_polygonShape):
		poly = shape + 132
		(*(*func(*_Stack, Transform, uintptr, int32, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSolidPolygonFcn})))(tls, xf, poly, (*Polygon)(unsafe.Pointer(poly)).Count, (*Polygon)(unsafe.Pointer(poly)).Radius, color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	case int32(b2_segmentShape):
		segment = shape + 132
		v1 = xf
		v2 = (*Segment)(unsafe.Pointer(segment)).Point1
		x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
		y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
		v3 = Vec2{
			X: x,
			Y: y,
		}
		goto _16
	_16:
		p11 = v3
		v5 = xf
		v6 = (*Segment)(unsafe.Pointer(segment)).Point2
		x = float32(v5.Q.C*v6.X) - float32(v5.Q.S*v6.Y) + v5.P.X
		y = float32(v5.Q.S*v6.X) + float32(v5.Q.C*v6.Y) + v5.P.Y
		v7 = Vec2{
			X: x,
			Y: y,
		}
		goto _20
	_20:
		p21 = v7
		(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, p11, p21, color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	case int32(b2_chainSegmentShape):
		segment1 = shape + 132 + 8
		v1 = xf
		v2 = (*Segment)(unsafe.Pointer(segment1)).Point1
		x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
		y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
		v3 = Vec2{
			X: x,
			Y: y,
		}
		goto _24
	_24:
		p12 = v3
		v5 = xf
		v6 = (*Segment)(unsafe.Pointer(segment1)).Point2
		x = float32(v5.Q.C*v6.X) - float32(v5.Q.S*v6.Y) + v5.P.X
		y = float32(v5.Q.S*v6.X) + float32(v5.Q.C*v6.Y) + v5.P.Y
		v7 = Vec2{
			X: x,
			Y: y,
		}
		goto _28
	_28:
		p22 = v7
		(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, p12, p22, color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		(*(*func(*_Stack, Vec2, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, p22, float32FromFloat32(4), color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		v2 = p12
		v3 = p22
		v31 = float32FromFloat32(0.1)
		v6 = Vec2{
			X: float32((float32FromFloat32(1)-v31)*v2.X) + float32(v31*v3.X),
			Y: float32((float32FromFloat32(1)-v31)*v2.Y) + float32(v31*v3.Y),
		}
		goto _33
	_33:
		(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, p12, v6, int32(b2_colorPaleGreen), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
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
