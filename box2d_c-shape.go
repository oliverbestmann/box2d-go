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
func b2ComputeShapeDistance(tls *_Stack, input uintptr, cache uintptr, simplexes uintptr, simplexCapacity int32) (r DistanceOutput) {
	bp := tls.Alloc(240)
	defer tls.Free(240)
	var aa, dx, dy, lengthSquared, radiusA, radiusB, s2, x, y, v16, v32, v36, v37 float32
	var d, nonUnitNormal, normal, u, v10, v11, v14, v6, v7 Vec2
	var duplicate, v34 uint8
	var i, i1, i2, iteration, maxIterations, saveCount, simplexIndex int32
	var output DistanceOutput
	var proxyA, vertex uintptr
	var saveA, saveB [3]int32
	var vertices [3]uintptr
	var v1 bool
	var v5 Transform
	var _ /* localPointA at bp+184 */ Vec2
	var _ /* localPointA at bp+200 */ Vec2
	var _ /* localPointA at bp+216 */ Vec2
	var _ /* localPointB at bp+192 */ Vec2
	var _ /* localPointB at bp+208 */ Vec2
	var _ /* localPointB at bp+224 */ Vec2
	var _ /* localProxyB at bp+0 */ ShapeProxy
	var _ /* simplex at bp+72 */ Simplex
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aa, d, duplicate, dx, dy, i, i1, i2, iteration, lengthSquared, maxIterations, nonUnitNormal, normal, output, proxyA, radiusA, radiusB, s2, saveA, saveB, saveCount, simplexIndex, u, vertex, vertices, x, y, v1, v10, v11, v14, v16, v32, v34, v36, v37, v5, v6, v7
	_ = uint64FromInt64(4)
	if v1 = !!((*DistanceInput)(unsafe.Pointer(input)).ProxyA.Count > 0 && (*DistanceInput)(unsafe.Pointer(input)).ProxyB.Count > 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+7721, __ccgo_ts+7689, int32FromInt32(427))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!((*DistanceInput)(unsafe.Pointer(input)).ProxyA.Radius >= float32FromFloat32(0)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+7772, __ccgo_ts+7689, int32FromInt32(428))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!((*DistanceInput)(unsafe.Pointer(input)).ProxyB.Radius >= float32FromFloat32(0)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+7801, __ccgo_ts+7689, int32FromInt32(429))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	output = DistanceOutput{}
	proxyA = input
	(*(*ShapeProxy)(unsafe.Add(nil, bp))).Count = (*DistanceInput)(unsafe.Pointer(input)).ProxyB.Count
	(*(*ShapeProxy)(unsafe.Add(nil, bp))).Radius = (*DistanceInput)(unsafe.Pointer(input)).ProxyB.Radius
	i = 0
	for {
		if !(i < (*(*ShapeProxy)(unsafe.Add(nil, bp))).Count) {
			break
		}
		v5 = (*DistanceInput)(unsafe.Pointer(input)).Transform
		v6 = *(*Vec2)(unsafe.Add(nil, input+72+uintptr(i)*8))
		x = float32(v5.Q.C*v6.X) - float32(v5.Q.S*v6.Y) + v5.P.X
		y = float32(v5.Q.S*v6.X) + float32(v5.Q.C*v6.Y) + v5.P.Y
		v7 = Vec2{
			X: x,
			Y: y,
		}
		goto _8
	_8:
		*(*Vec2)(unsafe.Add(nil, bp+uintptr(i)*8)) = v7
		goto _4
	_4:
		;
		i = i + 1
	}
	// Initialize the simplex.
	*(*Simplex)(unsafe.Add(nil, bp+72)) = b2MakeSimplexFromCache(tls, *(*SimplexCache)(unsafe.Add(nil, cache)), proxyA, bp)
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
			goto _9
		_9:
			;
			i1 = i1 + 1
		}
		d = Vec2{}
		switch (*(*Simplex)(unsafe.Add(nil, bp+72))).Count {
		case int32(1):
			v6 = (*(*Simplex)(unsafe.Add(nil, bp+72))).V1.W
			v7 = Vec2{
				X: -v6.X,
				Y: -v6.Y,
			}
			goto _12
		_12:
			d = v7
		case int32(2):
			d = b2SolveSimplex2(tls, bp+72)
		case int32(3):
			d = b2SolveSimplex3(tls, bp+72)
		default:
			if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
				b2InternalAssert(tls, __ccgo_ts+7715, __ccgo_ts+7689, int32FromInt32(494))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
		}
		// If we have 3 points, then the origin is in the corresponding triangle.
		if (*(*Simplex)(unsafe.Add(nil, bp+72))).Count == int32(3) {
			// Overlap
			*(*SimplexCache)(unsafe.Add(nil, cache)) = b2MakeSimplexCache(tls, bp+72)
			b2ComputeWitnessPoints(tls, bp+72, bp+184, bp+192)
			output.Distance = float32FromFloat32(0)
			output.PointA = *(*Vec2)(unsafe.Add(nil, bp+184))
			output.PointB = *(*Vec2)(unsafe.Add(nil, bp+192))
			return output
		}
		if simplexes != uintptrFromInt32(0) && simplexIndex < simplexCapacity {
			*(*Simplex)(unsafe.Add(nil, simplexes+uintptr(simplexIndex)*112)) = *(*Simplex)(unsafe.Add(nil, bp+72))
			simplexIndex = simplexIndex + int32(1)
		}
		// Ensure the search direction is numerically fit.
		v6 = d
		v7 = d
		v16 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y)
		goto _17
	_17:
		if v16 < float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			// This is unlikely but could lead to bad cycling.
			// The branch predictor seems to make this check have low cost.
			// The origin is probably contained by a line segment
			// or triangle. Thus the shapes are overlapped.
			*(*SimplexCache)(unsafe.Add(nil, cache)) = b2MakeSimplexCache(tls, bp+72)
			b2ComputeWitnessPoints(tls, bp+72, bp+200, bp+208)
			output.Distance = float32FromFloat32(0)
			output.PointA = *(*Vec2)(unsafe.Add(nil, bp+200))
			output.PointB = *(*Vec2)(unsafe.Add(nil, bp+208))
			return output
		}
		// Save the normal
		nonUnitNormal = d
		// Compute a tentative new simplex vertex using support points.
		// support = support(a, d) - support(b, -d)
		vertex = vertices[(*(*Simplex)(unsafe.Add(nil, bp+72))).Count]
		(*SimplexVertex)(unsafe.Pointer(vertex)).IndexA = b2FindSupport(tls, proxyA, d)
		(*SimplexVertex)(unsafe.Pointer(vertex)).WA = *(*Vec2)(unsafe.Add(nil, proxyA+uintptr((*SimplexVertex)(unsafe.Pointer(vertex)).IndexA)*8))
		v6 = d
		v7 = Vec2{
			X: -v6.X,
			Y: -v6.Y,
		}
		goto _20
	_20:
		(*SimplexVertex)(unsafe.Pointer(vertex)).IndexB = b2FindSupport(tls, bp, v7)
		(*SimplexVertex)(unsafe.Pointer(vertex)).WB = *(*Vec2)(unsafe.Add(nil, bp+uintptr((*SimplexVertex)(unsafe.Pointer(vertex)).IndexB)*8))
		v6 = (*SimplexVertex)(unsafe.Pointer(vertex)).WA
		v7 = (*SimplexVertex)(unsafe.Pointer(vertex)).WB
		v10 = Vec2{
			X: v6.X - v7.X,
			Y: v6.Y - v7.Y,
		}
		goto _24
	_24:
		(*SimplexVertex)(unsafe.Pointer(vertex)).W = v10
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
			goto _25
		_25:
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
	v6 = nonUnitNormal
	lengthSquared = float32(v6.X*v6.X) + float32(v6.Y*v6.Y)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s2 = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
		u = Vec2{
			X: float32(s2 * v6.X),
			Y: float32(s2 * v6.Y),
		}
		v7 = u
		goto _28
	}
	v7 = Vec2{}
	goto _28
_28:
	// Prepare output in frame A
	normal = v7
	v10 = normal
	v11 = v10
	v14 = v10
	v16 = float32(v11.X*v14.X) + float32(v11.Y*v14.Y)
	goto _33
_33:
	aa = v16
	v32 = float32FromFloat32(1) - aa
	if v32 < float32FromInt32(0) {
		v37 = -v32
	} else {
		v37 = v32
	}
	v36 = v37
	goto _38
_38:
	v34 = boolUint8(v36 < float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07)))
	goto _35
_35:
	;
	if v1 = !!(v34 != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+7830, __ccgo_ts+7689, int32FromInt32(584))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	b2ComputeWitnessPoints(tls, bp+72, bp+216, bp+224)
	output.Normal = normal
	v6 = *(*Vec2)(unsafe.Add(nil, bp+216))
	v7 = *(*Vec2)(unsafe.Add(nil, bp+224))
	dx = v7.X - v6.X
	dy = v7.Y - v6.Y
	v16 = sqrtf(tls, float32(dx*dx)+float32(dy*dy))
	goto _44
_44:
	output.Distance = v16
	output.PointA = *(*Vec2)(unsafe.Add(nil, bp+216))
	output.PointB = *(*Vec2)(unsafe.Add(nil, bp+224))
	output.Iterations = iteration
	output.SimplexCount = simplexIndex
	// Cache the simplex
	*(*SimplexCache)(unsafe.Add(nil, cache)) = b2MakeSimplexCache(tls, bp+72)
	// Apply radii if requested
	if (*DistanceInput)(unsafe.Pointer(input)).UseRadii != 0 {
		radiusA = (*DistanceInput)(unsafe.Pointer(input)).ProxyA.Radius
		radiusB = (*DistanceInput)(unsafe.Pointer(input)).ProxyB.Radius
		v16 = float32FromFloat32(0)
		v32 = output.Distance - radiusA - radiusB
		if v16 > v32 {
			v37 = v16
		} else {
			v37 = v32
		}
		v36 = v37
		goto _48
	_48:
		output.Distance = v36
		// Keep closest points on perimeter even if overlapped, this way the points move smoothly.
		v6 = output.PointA
		v16 = radiusA
		v7 = normal
		v10 = Vec2{
			X: v6.X + float32(v16*v7.X),
			Y: v6.Y + float32(v16*v7.Y),
		}
		goto _54
	_54:
		output.PointA = v10
		v6 = output.PointB
		v16 = radiusB
		v7 = normal
		v10 = Vec2{
			X: v6.X - float32(v16*v7.X),
			Y: v6.Y - float32(v16*v7.Y),
		}
		goto _59
	_59:
		output.PointB = v10
	}
	return output
}

// C documentation
//
//	// Shape cast using conservative advancement
func b2CastShape(tls *_Stack, input uintptr) (r CastOutput) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var aa, denominator, fraction, linearSlop, target, tolerance, totalRadius, v1, v2, v3, v5 float32
	var c1, c2, delta2, v10, v11, v13, v15, v16, v18, v19, v21, v8 Vec2
	var distanceOutput DistanceOutput
	var iteration, maxIterations int32
	var output CastOutput
	var v28 uint8
	var v34, v6 bool
	var _ /* cache at bp+0 */ SimplexCache
	var _ /* distanceInput at bp+8 */ DistanceInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aa, c1, c2, delta2, denominator, distanceOutput, fraction, iteration, linearSlop, maxIterations, output, target, tolerance, totalRadius, v1, v10, v11, v13, v15, v16, v18, v19, v2, v21, v28, v3, v34, v5, v6, v8
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
	if v6 = !!(target > tolerance); !v6 {
		b2InternalAssert(tls, __ccgo_ts+7855, __ccgo_ts+7689, int32FromInt32(622))
	}
	_ = v6 || bool(int32FromInt32(0) != 0)
	// Prepare input for distance query
	*(*SimplexCache)(unsafe.Add(nil, bp)) = SimplexCache{}
	fraction = float32FromFloat32(0)
	*(*DistanceInput)(unsafe.Add(nil, bp+8)) = DistanceInput{}
	(*(*DistanceInput)(unsafe.Add(nil, bp+8))).ProxyA = (*ShapeCastPairInput)(unsafe.Pointer(input)).ProxyA
	(*(*DistanceInput)(unsafe.Add(nil, bp+8))).ProxyB = (*ShapeCastPairInput)(unsafe.Pointer(input)).ProxyB
	(*(*DistanceInput)(unsafe.Add(nil, bp+8))).UseRadii = boolUint8(false1 != 0)
	// The whole cast runs in frame A. Advance the relative pose of B in float each iteration,
	// which keeps the math near the local origin and avoids re-relativizing world poses.
	(*(*DistanceInput)(unsafe.Add(nil, bp+8))).Transform = (*ShapeCastPairInput)(unsafe.Pointer(input)).Transform
	delta2 = (*ShapeCastPairInput)(unsafe.Pointer(input)).TranslationB
	output = CastOutput{}
	iteration = 0
	maxIterations = int32(20)
	for {
		if !(iteration < maxIterations) {
			break
		}
		output.Iterations += int32(1)
		distanceOutput = b2ComputeShapeDistance(tls, bp+8, bp, uintptrFromInt32(0), 0)
		if distanceOutput.Distance < target+tolerance {
			if iteration == 0 {
				if (*ShapeCastPairInput)(unsafe.Pointer(input)).CanEncroach != 0 && distanceOutput.Distance > float32(float32FromFloat32(2)*linearSlop) {
					target = distanceOutput.Distance - linearSlop
				} else {
					// Initial overlap
					output.Hit = boolUint8(true1 != 0)
					v8 = distanceOutput.PointA
					v1 = (*ShapeCastPairInput)(unsafe.Pointer(input)).ProxyA.Radius
					v10 = distanceOutput.Normal
					v11 = Vec2{
						X: v8.X + float32(v1*v10.X),
						Y: v8.Y + float32(v1*v10.Y),
					}
					goto _12
				_12:
					// Compute a common point
					c1 = v11
					v13 = distanceOutput.PointB
					v2 = -(*ShapeCastPairInput)(unsafe.Pointer(input)).ProxyB.Radius
					v15 = distanceOutput.Normal
					v16 = Vec2{
						X: v13.X + float32(v2*v15.X),
						Y: v13.Y + float32(v2*v15.Y),
					}
					goto _17
				_17:
					c2 = v16
					v18 = c1
					v19 = c2
					v3 = float32FromFloat32(0.5)
					v21 = Vec2{
						X: float32((float32FromFloat32(1)-v3)*v18.X) + float32(v3*v19.X),
						Y: float32((float32FromFloat32(1)-v3)*v18.Y) + float32(v3*v19.Y),
					}
					goto _22
				_22:
					output.Point = v21
					return output
				}
			} else {
				// Regular hit
				if v6 = distanceOutput.Distance > float32FromFloat32(0); v6 {
					v8 = distanceOutput.Normal
					v10 = v8
					v11 = v8
					v1 = float32(v10.X*v11.X) + float32(v10.Y*v11.Y)
					goto _27
				_27:
					aa = v1
					v2 = float32FromFloat32(1) - aa
					if v2 < float32FromInt32(0) {
						v5 = -v2
					} else {
						v5 = v2
					}
					v3 = v5
					goto _32
				_32:
					v28 = boolUint8(v3 < float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07)))
					goto _29
				_29:
				}
				if v34 = !!(v6 && v28 != 0); !v34 {
					b2InternalAssert(tls, __ccgo_ts+7874, __ccgo_ts+7689, int32FromInt32(673))
				}
				_ = v34 || bool(int32FromInt32(0) != 0)
				output.Fraction = fraction
				v8 = distanceOutput.PointA
				v1 = (*ShapeCastPairInput)(unsafe.Pointer(input)).ProxyA.Radius
				v10 = distanceOutput.Normal
				v11 = Vec2{
					X: v8.X + float32(v1*v10.X),
					Y: v8.Y + float32(v1*v10.Y),
				}
				goto _40
			_40:
				output.Point = v11
				output.Normal = distanceOutput.Normal
				output.Hit = boolUint8(true1 != 0)
				return output
			}
		}
		if v6 = !!(distanceOutput.Distance > float32FromFloat32(0)); !v6 {
			b2InternalAssert(tls, __ccgo_ts+7948, __ccgo_ts+7689, int32FromInt32(682))
		}
		_ = v6 || bool(int32FromInt32(0) != 0)
		v8 = distanceOutput.Normal
		v10 = v8
		v11 = v8
		v1 = float32(v10.X*v11.X) + float32(v10.Y*v11.Y)
		goto _46
	_46:
		aa = v1
		v2 = float32FromFloat32(1) - aa
		if v2 < float32FromInt32(0) {
			v5 = -v2
		} else {
			v5 = v2
		}
		v3 = v5
		goto _51
	_51:
		v28 = boolUint8(v3 < float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07)))
		goto _48
	_48:
		;
		if v6 = !!(v28 != 0); !v6 {
			b2InternalAssert(tls, __ccgo_ts+7979, __ccgo_ts+7689, int32FromInt32(683))
		}
		_ = v6 || bool(int32FromInt32(0) != 0)
		v8 = delta2
		v10 = distanceOutput.Normal
		v1 = float32(v8.X*v10.X) + float32(v8.Y*v10.Y)
		goto _57
	_57:
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
		v8 = (*ShapeCastPairInput)(unsafe.Pointer(input)).Transform.P
		v1 = fraction
		v10 = delta2
		v11 = Vec2{
			X: v8.X + float32(v1*v10.X),
			Y: v8.Y + float32(v1*v10.Y),
		}
		goto _62
	_62:
		(*(*DistanceInput)(unsafe.Add(nil, bp+8))).Transform.P = v11
		goto _7
	_7:
		;
		iteration = iteration + 1
	}
	// Failure!
	return output
}

func b2ComputeFatShapeAABB(tls *_Stack, shape uintptr, xf WorldTransform, extra float32) (r AABB) {
	var empty AABB
	var v1 bool
	_, _ = empty, v1
	switch (*b2Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b2_capsuleShape):
		return b2ComputeCapsuleFatAABB(tls, shape+128, xf, extra)
	case int32(b2_circleShape):
		return b2ComputeCircleFatAABB(tls, shape+128, xf, extra)
	case int32(b2_polygonShape):
		return b2ComputePolygonFatAABB(tls, shape+128, xf, extra)
	case int32(b2_segmentShape):
		return b2ComputeSegmentFatAABB(tls, shape+128, xf, extra)
	case int32(b2_chainSegmentShape):
		return b2ComputeSegmentFatAABB(tls, shape+128+8, xf, extra)
	default:
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+7715, __ccgo_ts+9594, int32FromInt32(501))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		empty = AABB{}
		return empty
	}
}

func b2CastShapeCircle(tls *_Stack, shape uintptr, input uintptr) (r CastOutput) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var _ /* pairInput at bp+0 */ ShapeCastPairInput
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).ProxyA = b2MakeProxy(tls, shape, int32(1), (*Circle)(unsafe.Pointer(shape)).Radius)
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).ProxyB = (*ShapeCastInput)(unsafe.Pointer(input)).Proxy
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).Transform = b2Transform_identity
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).TranslationB = (*ShapeCastInput)(unsafe.Pointer(input)).Translation
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).MaxFraction = (*ShapeCastInput)(unsafe.Pointer(input)).MaxFraction
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).CanEncroach = (*ShapeCastInput)(unsafe.Pointer(input)).CanEncroach
	return b2CastShape(tls, bp)
}

func b2CastShapeCapsule(tls *_Stack, shape uintptr, input uintptr) (r CastOutput) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var _ /* pairInput at bp+0 */ ShapeCastPairInput
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).ProxyA = b2MakeProxy(tls, shape, int32(2), (*Capsule)(unsafe.Pointer(shape)).Radius)
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).ProxyB = (*ShapeCastInput)(unsafe.Pointer(input)).Proxy
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).Transform = b2Transform_identity
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).TranslationB = (*ShapeCastInput)(unsafe.Pointer(input)).Translation
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).MaxFraction = (*ShapeCastInput)(unsafe.Pointer(input)).MaxFraction
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).CanEncroach = (*ShapeCastInput)(unsafe.Pointer(input)).CanEncroach
	return b2CastShape(tls, bp)
}

func b2CastShapeSegment(tls *_Stack, shape uintptr, input uintptr) (r CastOutput) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var _ /* pairInput at bp+0 */ ShapeCastPairInput
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).ProxyA = b2MakeProxy(tls, shape, int32(2), float32FromFloat32(0))
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).ProxyB = (*ShapeCastInput)(unsafe.Pointer(input)).Proxy
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).Transform = b2Transform_identity
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).TranslationB = (*ShapeCastInput)(unsafe.Pointer(input)).Translation
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).MaxFraction = (*ShapeCastInput)(unsafe.Pointer(input)).MaxFraction
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).CanEncroach = (*ShapeCastInput)(unsafe.Pointer(input)).CanEncroach
	return b2CastShape(tls, bp)
}

func b2CastShapePolygon(tls *_Stack, shape uintptr, input uintptr) (r CastOutput) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var _ /* pairInput at bp+0 */ ShapeCastPairInput
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).ProxyA = b2MakeProxy(tls, shape, (*Polygon)(unsafe.Pointer(shape)).Count, (*Polygon)(unsafe.Pointer(shape)).Radius)
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).ProxyB = (*ShapeCastInput)(unsafe.Pointer(input)).Proxy
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).Transform = b2Transform_identity
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).TranslationB = (*ShapeCastInput)(unsafe.Pointer(input)).Translation
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).MaxFraction = (*ShapeCastInput)(unsafe.Pointer(input)).MaxFraction
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).CanEncroach = (*ShapeCastInput)(unsafe.Pointer(input)).CanEncroach
	return b2CastShape(tls, bp)
}

func b2DrawShape(tls *_Stack, draw uintptr, shape uintptr, transform WorldTransform, color HexColor, drawChainNormals uint8) {
	var L, lengthSquared, rx, ry, s1, v27, v40 float32
	var c, p11, p12, p13, p2, p21, p22, v11, v15, v19, v23, v25, v3, v7 Pos
	var capsule, circle, poly, segment, segment1 uintptr
	var e, n, u, v10, v14, v18, v2, v22, v32, v34, v6 Vec2
	var v1, v5 WorldTransform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = L, c, capsule, circle, e, lengthSquared, n, p11, p12, p13, p2, p21, p22, poly, rx, ry, s1, segment, segment1, u, v1, v10, v11, v14, v15, v18, v19, v2, v22, v23, v25, v27, v3, v32, v34, v40, v5, v6, v7
	switch (*b2Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b2_capsuleShape):
		capsule = shape + 128
		v1 = transform
		v2 = (*Capsule)(unsafe.Pointer(capsule)).Center1
		rx = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y)
		ry = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y)
		v3 = Pos{
			X: v1.P.X + rx,
			Y: v1.P.Y + ry,
		}
		goto _4
	_4:
		p11 = v3
		v5 = transform
		v6 = (*Capsule)(unsafe.Pointer(capsule)).Center2
		rx = float32(v5.Q.C*v6.X) - float32(v5.Q.S*v6.Y)
		ry = float32(v5.Q.S*v6.X) + float32(v5.Q.C*v6.Y)
		v7 = Pos{
			X: v5.P.X + rx,
			Y: v5.P.Y + ry,
		}
		goto _8
	_8:
		p2 = v7
		(*(*func(*_Stack, Pos, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSolidCapsuleFcn})))(tls, p11, p2, (*Capsule)(unsafe.Pointer(capsule)).Radius, color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	case int32(b2_circleShape):
		circle = shape + 128
		(*(*func(*_Stack, WorldTransform, Vec2, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSolidCircleFcn})))(tls, transform, (*Circle)(unsafe.Pointer(circle)).Center, (*Circle)(unsafe.Pointer(circle)).Radius, color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	case int32(b2_polygonShape):
		poly = shape + 128
		(*(*func(*_Stack, WorldTransform, uintptr, int32, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSolidPolygonFcn})))(tls, transform, poly, (*Polygon)(unsafe.Pointer(poly)).Count, (*Polygon)(unsafe.Pointer(poly)).Radius, color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	case int32(b2_segmentShape):
		segment = shape + 128
		v1 = transform
		v2 = (*Segment)(unsafe.Pointer(segment)).Point1
		rx = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y)
		ry = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y)
		v3 = Pos{
			X: v1.P.X + rx,
			Y: v1.P.Y + ry,
		}
		goto _12
	_12:
		p12 = v3
		v5 = transform
		v6 = (*Segment)(unsafe.Pointer(segment)).Point2
		rx = float32(v5.Q.C*v6.X) - float32(v5.Q.S*v6.Y)
		ry = float32(v5.Q.S*v6.X) + float32(v5.Q.C*v6.Y)
		v7 = Pos{
			X: v5.P.X + rx,
			Y: v5.P.Y + ry,
		}
		goto _16
	_16:
		p21 = v7
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, p12, p21, color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	case int32(b2_chainSegmentShape):
		segment1 = shape + 128 + 8
		v1 = transform
		v2 = (*Segment)(unsafe.Pointer(segment1)).Point1
		rx = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y)
		ry = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y)
		v3 = Pos{
			X: v1.P.X + rx,
			Y: v1.P.Y + ry,
		}
		goto _20
	_20:
		p13 = v3
		v5 = transform
		v6 = (*Segment)(unsafe.Pointer(segment1)).Point2
		rx = float32(v5.Q.C*v6.X) - float32(v5.Q.S*v6.Y)
		ry = float32(v5.Q.S*v6.X) + float32(v5.Q.C*v6.Y)
		v7 = Pos{
			X: v5.P.X + rx,
			Y: v5.P.Y + ry,
		}
		goto _24
	_24:
		p22 = v7
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, p13, p22, color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, p22, float32FromFloat32(4), color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		if drawChainNormals != 0 {
			v3 = p13
			v7 = p22
			v27 = float32FromFloat32(0.5)
			v11 = Pos{
				X: float32((float32FromFloat32(1)-v27)*v3.X) + float32(v27*v7.X),
				Y: float32((float32FromFloat32(1)-v27)*v3.Y) + float32(v27*v7.Y),
			}
			goto _29
		_29:
			c = v11
			v15 = p22
			v19 = p13
			v2 = Vec2{
				X: v15.X - v19.X,
				Y: v15.Y - v19.Y,
			}
			goto _33
		_33:
			v6 = v2
			lengthSquared = float32(v6.X*v6.X) + float32(v6.Y*v6.Y)
			if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
				s1 = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
				u = Vec2{
					X: float32(s1 * v6.X),
					Y: float32(s1 * v6.Y),
				}
				v10 = u
				goto _36
			}
			v10 = Vec2{}
			goto _36
		_36:
			e = v10
			v14 = e
			v18 = Vec2{
				X: v14.Y,
				Y: -v14.X,
			}
			goto _39
		_39:
			n = v18
			L = float32(float32FromFloat32(0.2) * b2GetLengthUnitsPerMeter(tls))
			v40 = L
			v22 = n
			v32 = Vec2{
				X: float32(v40 * v22.X),
				Y: float32(v40 * v22.Y),
			}
			goto _43
		_43:
			v23 = c
			v34 = v32
			v25 = Pos{
				X: v23.X + v34.X,
				Y: v23.Y + v34.Y,
			}
			goto _47
		_47:
			(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, c, v25, int32(b2_colorPaleGreen), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		}
	default:
		break
	}
}

func b2Shape_IsValid(tls *_Stack, id ShapeId) (r uint8) {
	var shape, world uintptr
	var shapeId int32
	var v1 bool
	_, _, _, _ = shape, shapeId, world, v1
	if int32(_B2_MAX_WORLDS) <= int32FromUint16(id.World0) {
		return boolUint8(false1 != 0)
	}
	world = uintptr(unsafe.Pointer(&b2_worlds)) + uintptr(id.World0)*2792
	if int32FromUint16((*b2World)(unsafe.Pointer(world)).WorldId) != int32FromUint16(id.World0) {
		// world is free
		return boolUint8(false1 != 0)
	}
	shapeId = id.Index1 - int32(1)
	if shapeId < 0 || (*b2World)(unsafe.Pointer(world)).Shapes.Count <= shapeId {
		return boolUint8(false1 != 0)
	}
	shape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*280
	if (*b2Shape)(unsafe.Pointer(shape)).Id == -int32(1) {
		// shape is free
		return boolUint8(false1 != 0)
	}
	if v1 = !!((*b2Shape)(unsafe.Pointer(shape)).Id == shapeId); !v1 {
		b2InternalAssert(tls, __ccgo_ts+17954, __ccgo_ts+16660, int32FromInt32(1644))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return boolUint8(int32FromUint16(id.Generation) == int32FromUint16((*b2Shape)(unsafe.Pointer(shape)).Generation))
}

func b2RecMakeShapeId(tls *_Stack, rdr uintptr, recorded ShapeId) (r ShapeId) {
	var id ShapeId
	_ = id
	id.Index1 = recorded.Index1
	id.World0 = uint16(uint32((*b2RecReader)(unsafe.Pointer(rdr)).ReplayWorldId.Index1) - uint32FromUint32(1))
	id.Generation = recorded.Generation
	return id
}

func b2RecCheckShapeId(tls *_Stack, rdr uintptr, got ShapeId, rec ShapeId) {
	b2RecCheckId(tls, rdr, __ccgo_ts+20833, got.Index1, uint32(got.Generation), rec.Index1, uint32(rec.Generation))
}

/** @} */

func b2GetShape(tls *_Stack, world uintptr, shapeId ShapeId) (r uintptr) {
	var id int32
	var shape uintptr
	var v1, v2 bool
	_, _, _, _ = id, shape, v1, v2
	id = shapeId.Index1 - int32(1)
	if v1 = !!(0 <= id && id < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+21947, __ccgo_ts+21995, int32FromInt32(23))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	shape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr(id)*280
	if v2 = !!((*b2Shape)(unsafe.Pointer(shape)).Id == id && int32FromUint16((*b2Shape)(unsafe.Pointer(shape)).Generation) == int32FromUint16(shapeId.Generation)); !v2 {
		b2InternalAssert(tls, __ccgo_ts+22018, __ccgo_ts+21995, int32FromInt32(24))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	return shape
}

func b2GetChainShape(tls *_Stack, world uintptr, chainId ChainId) (r uintptr) {
	var chain uintptr
	var id int32
	var v1, v2 bool
	_, _, _, _ = chain, id, v1, v2
	id = chainId.Index1 - int32(1)
	if v1 = !!(0 <= id && id < (*b2World)(unsafe.Pointer(world)).ChainShapes.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+22077, __ccgo_ts+21995, int32FromInt32(31))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	chain = (*b2World)(unsafe.Pointer(world)).ChainShapes.Data + uintptr(id)*32
	if v2 = !!((*b2ChainShape)(unsafe.Pointer(chain)).Id == id && int32FromUint16((*b2ChainShape)(unsafe.Pointer(chain)).Generation) == int32FromUint16(chainId.Generation)); !v2 {
		b2InternalAssert(tls, __ccgo_ts+22130, __ccgo_ts+21995, int32FromInt32(32))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	return chain
}

func b2ComputeShapeMargin(tls *_Stack, shape uintptr) (r float32) {
	var c, v1, v2 Vec2
	var count, i int32
	var distanceSqr, dx, dy, margin, maxExtentSqr, v10, v11, v12, v3, v8 float32
	var poly uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, count, distanceSqr, dx, dy, i, margin, maxExtentSqr, poly, v1, v10, v11, v12, v2, v3, v8
	margin = float32FromFloat32(0)
	switch (*b2Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b2_capsuleShape):
		v1 = (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo14_128.Capsule.Center2
		v2 = (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo14_128.Capsule.Center1
		dx = v2.X - v1.X
		dy = v2.Y - v1.Y
		v3 = sqrtf(tls, float32(dx*dx)+float32(dy*dy))
		goto _4
	_4:
		margin = float32(float32FromFloat32(0.5)*v3) + (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo14_128.Capsule.Radius
	case int32(b2_circleShape):
		margin = (*(*Circle)(unsafe.Add(unsafe.Pointer(shape), 128))).Radius
	case int32(b2_polygonShape):
		poly = shape + 128
		maxExtentSqr = float32FromFloat32(0)
		count = (*Polygon)(unsafe.Pointer(poly)).Count
		i = 0
		for {
			if !(i < count) {
				break
			}
			v1 = *(*Vec2)(unsafe.Add(nil, poly+uintptr(i)*8))
			v2 = (*Polygon)(unsafe.Pointer(poly)).Centroid
			c = Vec2{
				X: v2.X - v1.X,
				Y: v2.Y - v1.Y,
			}
			v3 = float32(c.X*c.X) + float32(c.Y*c.Y)
			goto _9
		_9:
			distanceSqr = v3
			v8 = maxExtentSqr
			v10 = distanceSqr
			if v8 > v10 {
				v12 = v8
			} else {
				v12 = v10
			}
			v11 = v12
			goto _13
		_13:
			maxExtentSqr = v11
			goto _5
		_5:
			;
			i = i + 1
		}
		margin = sqrtf(tls, maxExtentSqr)
	case int32(b2_segmentShape):
		v1 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 128))).Point1
		v2 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 128))).Point2
		dx = v2.X - v1.X
		dy = v2.Y - v1.Y
		v3 = sqrtf(tls, float32(dx*dx)+float32(dy*dy))
		goto _18
	_18:
		margin = float32(float32FromFloat32(0.5) * v3)
	case int32(b2_chainSegmentShape):
		v1 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 128))).Segment.Point1
		v2 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 128))).Segment.Point2
		dx = v2.X - v1.X
		dy = v2.Y - v1.Y
		v3 = sqrtf(tls, float32(dx*dx)+float32(dy*dy))
		goto _22
	_22:
		margin = float32(float32FromFloat32(0.5) * v3)
	default:
		return float32(float32FromFloat32(0.05) * b2_lengthUnitsPerMeter)
	}
	v3 = float32(float32FromFloat32(0.05) * b2_lengthUnitsPerMeter)
	v8 = float32(float32FromFloat32(0.125) * margin)
	if v3 < v8 {
		v11 = v3
	} else {
		v11 = v8
	}
	v10 = v11
	goto _26
_26:
	return v10
}

func b2UpdateShapeAABBs(tls *_Stack, shape uintptr, fatAABB uintptr, transform WorldTransform, proxyType BodyType) {
	var aabb AABB
	var aabbMargin, margin, speculativeDistance, v1 float32
	_, _, _, _, _ = aabb, aabbMargin, margin, speculativeDistance, v1
	// Compute a bounding box with a speculative margin
	speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter))
	aabbMargin = (*b2Shape)(unsafe.Pointer(shape)).AabbMargin
	aabb = b2ComputeFatShapeAABB(tls, shape, transform, speculativeDistance)
	(*b2Shape)(unsafe.Pointer(shape)).Aabb = aabb
	if proxyType == int32(b2_staticBody) {
		v1 = speculativeDistance
	} else {
		v1 = aabbMargin
	}
	// Smaller margin for static bodies. Cannot be zero due to TOI tolerance.
	margin = v1
	(*AABB)(unsafe.Pointer(fatAABB)).LowerBound.X = aabb.LowerBound.X - margin
	(*AABB)(unsafe.Pointer(fatAABB)).LowerBound.Y = aabb.LowerBound.Y - margin
	(*AABB)(unsafe.Pointer(fatAABB)).UpperBound.X = aabb.UpperBound.X + margin
	(*AABB)(unsafe.Pointer(fatAABB)).UpperBound.Y = aabb.UpperBound.Y + margin
}

func b2CreateShapeInternal(tls *_Stack, world uintptr, body uintptr, transform WorldTransform, def uintptr, geometry uintptr, shapeType ShapeType) (r uintptr) {
	var elementSize, elementSize1, elementSize2, newCapacity, newCapacity1, newCapacity2, newSize, newSize1, newSize2, oldSize, oldSize1, oldSize2, shapeId, v1 int32
	var headShape, shape, v7 uintptr
	var proxyType b2BodyType1
	var sensor b2Sensor
	var v3 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = elementSize, elementSize1, elementSize2, headShape, newCapacity, newCapacity1, newCapacity2, newSize, newSize1, newSize2, oldSize, oldSize1, oldSize2, proxyType, sensor, shape, shapeId, v1, v3, v7
	shapeId = b2AllocId(tls, world+2112)
	if shapeId == (*b2World)(unsafe.Pointer(world)).Shapes.Count {
		elementSize = int32(280)
		if (*b2World)(unsafe.Pointer(world)).Shapes.Count >= (*b2World)(unsafe.Pointer(world)).Shapes.Capacity {
			oldSize = (*b2World)(unsafe.Pointer(world)).Shapes.Capacity * elementSize
			if (*b2World)(unsafe.Pointer(world)).Shapes.Capacity == 0 {
				v1 = int32(8)
			} else {
				v1 = int32(2) * (*b2World)(unsafe.Pointer(world)).Shapes.Capacity
			}
			newCapacity = v1
			newSize = newCapacity * elementSize
			(*b2World)(unsafe.Pointer(world)).Shapes.Data = b2GrowAlloc(tls, (*b2World)(unsafe.Pointer(world)).Shapes.Data, uint64FromInt32(oldSize), uint64FromInt32(newSize))
			(*b2World)(unsafe.Pointer(world)).Shapes.Capacity = newCapacity
		}
		*(*b2Shape)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Shapes.Data+uintptr((*b2World)(unsafe.Pointer(world)).Shapes.Count)*280)) = b2Shape{}
		(*b2World)(unsafe.Pointer(world)).Shapes.Count += int32(1)
		elementSize1 = int32(16)
		if (*b2World)(unsafe.Pointer(world)).FatAABBs.Count >= (*b2World)(unsafe.Pointer(world)).FatAABBs.Capacity {
			oldSize1 = (*b2World)(unsafe.Pointer(world)).FatAABBs.Capacity * elementSize1
			if (*b2World)(unsafe.Pointer(world)).FatAABBs.Capacity == 0 {
				v1 = int32(8)
			} else {
				v1 = int32(2) * (*b2World)(unsafe.Pointer(world)).FatAABBs.Capacity
			}
			newCapacity1 = v1
			newSize1 = newCapacity1 * elementSize1
			(*b2World)(unsafe.Pointer(world)).FatAABBs.Data = b2GrowAlloc(tls, (*b2World)(unsafe.Pointer(world)).FatAABBs.Data, uint64FromInt32(oldSize1), uint64FromInt32(newSize1))
			(*b2World)(unsafe.Pointer(world)).FatAABBs.Capacity = newCapacity1
		}
		*(*AABB)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).FatAABBs.Data+uintptr((*b2World)(unsafe.Pointer(world)).FatAABBs.Count)*16)) = AABB{}
		(*b2World)(unsafe.Pointer(world)).FatAABBs.Count += int32(1)
	} else {
		if v3 = !!((*(*b2Shape)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Shapes.Data+uintptr(shapeId)*280))).Id == -int32FromInt32(1)); !v3 {
			b2InternalAssert(tls, __ccgo_ts+22189, __ccgo_ts+21995, int32FromInt32(118))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
	}
	if v3 = !!((*b2World)(unsafe.Pointer(world)).FatAABBs.Count == (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+22237, __ccgo_ts+21995, int32FromInt32(121))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	if v3 = !!(0 <= shapeId && shapeId < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+1908, __ccgo_ts+21995, int32FromInt32(123))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	shape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*280
	switch shapeType {
	case int32(b2_capsuleShape):
		(*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo14_128.Capsule = *(*Capsule)(unsafe.Add(nil, geometry))
	case int32(b2_circleShape):
		*(*Circle)(unsafe.Add(unsafe.Pointer(shape), 128)) = *(*Circle)(unsafe.Add(nil, geometry))
	case int32(b2_polygonShape):
		*(*Polygon)(unsafe.Add(unsafe.Pointer(shape), 128)) = *(*Polygon)(unsafe.Add(nil, geometry))
	case int32(b2_segmentShape):
		*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 128)) = *(*Segment)(unsafe.Add(nil, geometry))
	case int32(b2_chainSegmentShape):
		*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 128)) = *(*ChainSegment)(unsafe.Add(nil, geometry))
	default:
		if v3 = bool(!!(int32FromInt32(false1) != 0)); !v3 {
			b2InternalAssert(tls, __ccgo_ts+7715, __ccgo_ts+21995, int32FromInt32(148))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		break
	}
	(*b2Shape)(unsafe.Pointer(shape)).Id = shapeId
	(*b2Shape)(unsafe.Pointer(shape)).BodyId = (*b2Body)(unsafe.Pointer(body)).Id
	(*b2Shape)(unsafe.Pointer(shape)).Type1 = shapeType
	(*b2Shape)(unsafe.Pointer(shape)).Density = (*ShapeDef)(unsafe.Pointer(def)).Density
	(*b2Shape)(unsafe.Pointer(shape)).Material = (*ShapeDef)(unsafe.Pointer(def)).Material
	(*b2Shape)(unsafe.Pointer(shape)).Filter = (*ShapeDef)(unsafe.Pointer(def)).Filter
	(*b2Shape)(unsafe.Pointer(shape)).UserData = (*ShapeDef)(unsafe.Pointer(def)).UserData
	(*b2Shape)(unsafe.Pointer(shape)).EnableSensorEvents = (*ShapeDef)(unsafe.Pointer(def)).EnableSensorEvents
	(*b2Shape)(unsafe.Pointer(shape)).EnableContactEvents = (*ShapeDef)(unsafe.Pointer(def)).EnableContactEvents
	(*b2Shape)(unsafe.Pointer(shape)).EnableCustomFiltering = (*ShapeDef)(unsafe.Pointer(def)).EnableCustomFiltering
	(*b2Shape)(unsafe.Pointer(shape)).EnableHitEvents = (*ShapeDef)(unsafe.Pointer(def)).EnableHitEvents
	(*b2Shape)(unsafe.Pointer(shape)).EnablePreSolveEvents = (*ShapeDef)(unsafe.Pointer(def)).EnablePreSolveEvents
	(*b2Shape)(unsafe.Pointer(shape)).ProxyKey = -int32(1)
	(*b2Shape)(unsafe.Pointer(shape)).LocalCentroid = b2GetShapeCentroid(tls, shape)
	(*b2Shape)(unsafe.Pointer(shape)).AabbMargin = b2ComputeShapeMargin(tls, shape)
	(*b2Shape)(unsafe.Pointer(shape)).Aabb = AABB{
		LowerBound: b2Vec2_zero,
		UpperBound: b2Vec2_zero,
	}
	*(*AABB)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).FatAABBs.Data+uintptr(shapeId)*16)) = AABB{
		LowerBound: b2Vec2_zero,
		UpperBound: b2Vec2_zero,
	}
	v7 = shape + 272
	*(*uint16_t)(unsafe.Pointer(v7)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v7))) + int32FromInt32(1))
	if (*b2Body)(unsafe.Pointer(body)).SetIndex != int32(b2_disabledSet) {
		proxyType = (*b2Body)(unsafe.Pointer(body)).Type1
		b2CreateShapeProxy(tls, world, shape, proxyType, transform, boolUint8((*ShapeDef)(unsafe.Pointer(def)).InvokeContactCreation != 0 || (*ShapeDef)(unsafe.Pointer(def)).IsSensor != 0))
	}
	// Add to shape doubly linked list
	if (*b2Body)(unsafe.Pointer(body)).HeadShapeId != -int32(1) {
		if v3 = !!(0 <= (*b2Body)(unsafe.Pointer(body)).HeadShapeId && (*b2Body)(unsafe.Pointer(body)).HeadShapeId < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v3 {
			b2InternalAssert(tls, __ccgo_ts+2316, __ccgo_ts+21995, int32FromInt32(180))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		headShape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b2Body)(unsafe.Pointer(body)).HeadShapeId)*280
		(*b2Shape)(unsafe.Pointer(headShape)).PrevShapeId = shapeId
	}
	(*b2Shape)(unsafe.Pointer(shape)).PrevShapeId = -int32(1)
	(*b2Shape)(unsafe.Pointer(shape)).NextShapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	(*b2Body)(unsafe.Pointer(body)).HeadShapeId = shapeId
	*(*int32)(unsafe.Add(nil, body+28)) += int32(1)
	if (*ShapeDef)(unsafe.Pointer(def)).IsSensor != 0 {
		(*b2Shape)(unsafe.Pointer(shape)).SensorIndex = (*b2World)(unsafe.Pointer(world)).Sensors.Count
		sensor = b2Sensor{}
		sensor.ShapeId = shapeId
		sensor.Hits.Data = b2GrowAlloc(tls, uintptrFromInt32(0), uint64(0), uint64FromInt32(int32FromInt32(4))*uint64FromInt64(8))
		sensor.Hits.Count = 0
		sensor.Hits.Capacity = int32(4)
		sensor.Overlaps1.Data = b2GrowAlloc(tls, uintptrFromInt32(0), uint64(0), uint64FromInt32(int32FromInt32(16))*uint64FromInt64(8))
		sensor.Overlaps1.Count = 0
		sensor.Overlaps1.Capacity = int32(16)
		sensor.Overlaps2.Data = b2GrowAlloc(tls, uintptrFromInt32(0), uint64(0), uint64FromInt32(int32FromInt32(16))*uint64FromInt64(8))
		sensor.Overlaps2.Count = 0
		sensor.Overlaps2.Capacity = int32(16)
		elementSize2 = int32(56)
		if (*b2World)(unsafe.Pointer(world)).Sensors.Count >= (*b2World)(unsafe.Pointer(world)).Sensors.Capacity {
			oldSize2 = (*b2World)(unsafe.Pointer(world)).Sensors.Capacity * elementSize2
			if (*b2World)(unsafe.Pointer(world)).Sensors.Capacity == 0 {
				v1 = int32(8)
			} else {
				v1 = int32(2) * (*b2World)(unsafe.Pointer(world)).Sensors.Capacity
			}
			newCapacity2 = v1
			newSize2 = newCapacity2 * elementSize2
			(*b2World)(unsafe.Pointer(world)).Sensors.Data = b2GrowAlloc(tls, (*b2World)(unsafe.Pointer(world)).Sensors.Data, uint64FromInt32(oldSize2), uint64FromInt32(newSize2))
			(*b2World)(unsafe.Pointer(world)).Sensors.Capacity = newCapacity2
		}
		*(*b2Sensor)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Sensors.Data+uintptr((*b2World)(unsafe.Pointer(world)).Sensors.Count)*56)) = sensor
		(*b2World)(unsafe.Pointer(world)).Sensors.Count += int32(1)
	} else {
		(*b2Shape)(unsafe.Pointer(shape)).SensorIndex = -int32(1)
	}
	b2ValidateSolverSets(tls, world)
	return shape
}

func b2CreateShape(tls *_Stack, bodyId BodyId, def uintptr, geometry uintptr, shapeType ShapeType) (r ShapeId) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var body, shape, world uintptr
	var id ShapeId
	var transform WorldTransform
	var v1 bool
	_, _, _, _, _, _ = body, id, shape, transform, world, v1
	if v1 = !!((*ShapeDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B2_SECRET_COOKIE)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1077, __ccgo_ts+21995, int32FromInt32(268))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if boolInt32(b2IsValidFloat(tls, (*ShapeDef)(unsafe.Pointer(def)).Density) != 0 && (*ShapeDef)(unsafe.Pointer(def)).Density >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+8, __ccgo_ts+22282, uintptr(unsafe.Pointer(&__func__65))))
		if v1 = !!(b2IsValidFloat(tls, (*ShapeDef)(unsafe.Pointer(def)).Density) != 0 && (*ShapeDef)(unsafe.Pointer(def)).Density >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+22282, __ccgo_ts+21995, int32FromInt32(269))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return ShapeId{}
	}
	if int32FromUint8(b2IsValidMaterial(tls, def+8)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+8, __ccgo_ts+22337, uintptr(unsafe.Pointer(&__func__65))))
		if v1 = !!(b2IsValidMaterial(tls, def+8) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+22337, __ccgo_ts+21995, int32FromInt32(270))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return ShapeId{}
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
	} else {
		if (*b2Body)(unsafe.Pointer(body)).Flags&uint32(b2_dirtyMass) == uint32(0) {
			*(*uint32_t)(unsafe.Add(nil, body+80)) |= uint32(b2_dirtyMass)
			b2SyncBodyFlags(tls, world, body)
		}
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
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var id ShapeId
	var world uintptr
	var v1 bool
	var _ /* _ca at bp+0 */ b2RecArgs_CreateCircleShape
	_, _, _ = id, world, v1
	if int32FromUint8(b2IsValidCircle(tls, circle)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+120, __ccgo_ts+22373, uintptr(unsafe.Pointer(&__func__66))))
		if v1 = !!(b2IsValidCircle(tls, circle) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+22373, __ccgo_ts+21995, int32FromInt32(301))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return b2_nullShapeId
	}
	id = b2CreateShape(tls, bodyId, def, circle, int32(b2_circleShape))
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_CreateCircleShape)(unsafe.Add(nil, bp)) = b2RecArgs_CreateCircleShape{
			Body:   bodyId,
			Def:    *(*ShapeDef)(unsafe.Add(nil, def)),
			Circle: *(*Circle)(unsafe.Add(nil, circle)),
		}
		b2RecWriteRet_CreateCircleShape(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp, id)
	}
	return id
}

func b2CreateCapsuleShape(tls *_Stack, bodyId BodyId, def uintptr, capsule uintptr) (r ShapeId) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var c, v2, v3 Vec2
	var id ShapeId
	var lengthSqr, v4 float32
	var world uintptr
	var v1 bool
	var _ /* _ca at bp+0 */ b2RecArgs_CreateCapsuleShape
	_, _, _, _, _, _, _, _ = c, id, lengthSqr, world, v1, v2, v3, v4
	if int32FromUint8(b2IsValidCapsule(tls, capsule)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+128, __ccgo_ts+22399, uintptr(unsafe.Pointer(&__func__117))))
		if v1 = !!(b2IsValidCapsule(tls, capsule) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+22399, __ccgo_ts+21995, int32FromInt32(313))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return b2_nullShapeId
	}
	v2 = (*Capsule)(unsafe.Pointer(capsule)).Center1
	v3 = (*Capsule)(unsafe.Pointer(capsule)).Center2
	c = Vec2{
		X: v3.X - v2.X,
		Y: v3.Y - v2.Y,
	}
	v4 = float32(c.X*c.X) + float32(c.Y*c.Y)
	goto _5
_5:
	lengthSqr = v4
	if lengthSqr <= float32(float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
		return b2_nullShapeId
	}
	id = b2CreateShape(tls, bodyId, def, capsule, int32(b2_capsuleShape))
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_CreateCapsuleShape)(unsafe.Add(nil, bp)) = b2RecArgs_CreateCapsuleShape{
			Body:    bodyId,
			Def:     *(*ShapeDef)(unsafe.Add(nil, def)),
			Capsule: *(*Capsule)(unsafe.Add(nil, capsule)),
		}
		b2RecWriteRet_CreateCapsuleShape(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp, id)
	}
	return id
}

func b2CreatePolygonShape(tls *_Stack, bodyId BodyId, def uintptr, polygon uintptr) (r ShapeId) {
	bp := tls.Alloc(272)
	defer tls.Free(272)
	var id ShapeId
	var world uintptr
	var v1 bool
	var _ /* _ca at bp+0 */ b2RecArgs_CreatePolygonShape
	_, _, _ = id, world, v1
	if int32FromUint8(b2IsValidPolygon(tls, polygon)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+248, __ccgo_ts+22427, uintptr(unsafe.Pointer(&__func__67))))
		if v1 = !!(b2IsValidPolygon(tls, polygon) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+22427, __ccgo_ts+21995, int32FromInt32(331))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return b2_nullShapeId
	}
	id = b2CreateShape(tls, bodyId, def, polygon, int32(b2_polygonShape))
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_CreatePolygonShape)(unsafe.Add(nil, bp)) = b2RecArgs_CreatePolygonShape{
			Body:    bodyId,
			Def:     *(*ShapeDef)(unsafe.Add(nil, def)),
			Polygon: *(*Polygon)(unsafe.Add(nil, polygon)),
		}
		b2RecWriteRet_CreatePolygonShape(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp, id)
	}
	return id
}

func b2CreateSegmentShape(tls *_Stack, bodyId BodyId, def uintptr, segment uintptr) (r ShapeId) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var c, v2, v3 Vec2
	var id ShapeId
	var lengthSqr, v4 float32
	var world uintptr
	var v1 bool
	var _ /* _ca at bp+0 */ b2RecArgs_CreateSegmentShape
	_, _, _, _, _, _, _, _ = c, id, lengthSqr, world, v1, v2, v3, v4
	if int32FromUint8(b2IsValidSegment(tls, segment)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+120, __ccgo_ts+22455, uintptr(unsafe.Pointer(&__func__118))))
		if v1 = !!(b2IsValidSegment(tls, segment) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+22455, __ccgo_ts+21995, int32FromInt32(343))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return b2_nullShapeId
	}
	v2 = (*Segment)(unsafe.Pointer(segment)).Point1
	v3 = (*Segment)(unsafe.Pointer(segment)).Point2
	c = Vec2{
		X: v3.X - v2.X,
		Y: v3.Y - v2.Y,
	}
	v4 = float32(c.X*c.X) + float32(c.Y*c.Y)
	goto _5
_5:
	lengthSqr = v4
	if lengthSqr <= float32(float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+7715, __ccgo_ts+21995, int32FromInt32(348))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return b2_nullShapeId
	}
	id = b2CreateShape(tls, bodyId, def, segment, int32(b2_segmentShape))
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_CreateSegmentShape)(unsafe.Add(nil, bp)) = b2RecArgs_CreateSegmentShape{
			Body:    bodyId,
			Def:     *(*ShapeDef)(unsafe.Add(nil, def)),
			Segment: *(*Segment)(unsafe.Add(nil, segment)),
		}
		b2RecWriteRet_CreateSegmentShape(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp, id)
	}
	return id
}

func b2CreateChainSegmentShape(tls *_Stack, bodyId BodyId, def uintptr, chainSegment uintptr) (r ShapeId) {
	bp := tls.Alloc(208)
	defer tls.Free(208)
	var c, v2, v3 Vec2
	var id ShapeId
	var lengthSqr, v4 float32
	var world uintptr
	var v1 bool
	var _ /* _ca at bp+40 */ b2RecArgs_CreateChainSegmentShape
	var _ /* local at bp+0 */ ChainSegment
	_, _, _, _, _, _, _, _ = c, id, lengthSqr, world, v1, v2, v3, v4
	if int32FromUint8(b2IsValidChainSegment(tls, chainSegment)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+184, __ccgo_ts+22483, uintptr(unsafe.Pointer(&__func__119))))
		if v1 = !!(b2IsValidChainSegment(tls, chainSegment) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+22483, __ccgo_ts+21995, int32FromInt32(362))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return b2_nullShapeId
	}
	v2 = (*ChainSegment)(unsafe.Pointer(chainSegment)).Segment.Point1
	v3 = (*ChainSegment)(unsafe.Pointer(chainSegment)).Segment.Point2
	c = Vec2{
		X: v3.X - v2.X,
		Y: v3.Y - v2.Y,
	}
	v4 = float32(c.X*c.X) + float32(c.Y*c.Y)
	goto _5
_5:
	lengthSqr = v4
	if lengthSqr <= float32(float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+7715, __ccgo_ts+21995, int32FromInt32(367))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return b2_nullShapeId
	}
	// No parent chain shape
	*(*ChainSegment)(unsafe.Add(nil, bp)) = *(*ChainSegment)(unsafe.Add(nil, chainSegment))
	(*(*ChainSegment)(unsafe.Add(nil, bp))).ChainId = -int32(1)
	id = b2CreateShape(tls, bodyId, def, bp, int32(b2_chainSegmentShape))
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_CreateChainSegmentShape)(unsafe.Add(nil, bp+40)) = b2RecArgs_CreateChainSegmentShape{
			Body:         bodyId,
			Def:          *(*ShapeDef)(unsafe.Add(nil, def)),
			ChainSegment: *(*ChainSegment)(unsafe.Add(nil, bp)),
		}
		b2RecWriteRet_CreateChainSegmentShape(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp+40, id)
	}
	return id
}

// C documentation
//
//	// Destroy a shape on a body. This doesn't need to be called when destroying a body.
func b2DestroyShapeInternal(tls *_Stack, world uintptr, shape uintptr, body uintptr) {
	var contact, movedSensor, nextShape, otherSensorShape, prevShape, ref, sensor, v7, v8 uintptr
	var contactId, contactKey, edgeIndex, elementSize1, i, movedIndex, newCapacity, newSize, oldSize, shapeId, v10, v6, v9 int32
	var event SensorEndTouchEvent
	var v1, v2, v3 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = contact, contactId, contactKey, edgeIndex, elementSize1, event, i, movedIndex, movedSensor, newCapacity, newSize, nextShape, oldSize, otherSensorShape, prevShape, ref, sensor, shapeId, v1, v10, v2, v3, v6, v7, v8, v9
	shapeId = (*b2Shape)(unsafe.Pointer(shape)).Id
	// Remove the shape from the body's doubly linked list.
	if (*b2Shape)(unsafe.Pointer(shape)).PrevShapeId != -int32(1) {
		if v1 = !!(0 <= (*b2Shape)(unsafe.Pointer(shape)).PrevShapeId && (*b2Shape)(unsafe.Pointer(shape)).PrevShapeId < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+22521, __ccgo_ts+21995, int32FromInt32(391))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		prevShape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b2Shape)(unsafe.Pointer(shape)).PrevShapeId)*280
		(*b2Shape)(unsafe.Pointer(prevShape)).NextShapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
	}
	if (*b2Shape)(unsafe.Pointer(shape)).NextShapeId != -int32(1) {
		if v1 = !!(0 <= (*b2Shape)(unsafe.Pointer(shape)).NextShapeId && (*b2Shape)(unsafe.Pointer(shape)).NextShapeId < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+2394, __ccgo_ts+21995, int32FromInt32(397))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		nextShape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b2Shape)(unsafe.Pointer(shape)).NextShapeId)*280
		(*b2Shape)(unsafe.Pointer(nextShape)).PrevShapeId = (*b2Shape)(unsafe.Pointer(shape)).PrevShapeId
	}
	if shapeId == (*b2Body)(unsafe.Pointer(body)).HeadShapeId {
		(*b2Body)(unsafe.Pointer(body)).HeadShapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
	}
	*(*int32)(unsafe.Add(nil, body+28)) -= int32(1)
	// Remove from broad-phase.
	b2DestroyShapeProxy(tls, shape, world+40)
	// Destroy any contacts associated with the shape.
	contactKey = (*b2Body)(unsafe.Pointer(body)).HeadContactKey
	for contactKey != -int32(1) {
		contactId = contactKey >> int32(1)
		edgeIndex = contactKey & int32(1)
		if v1 = !!(0 <= contactId && contactId < (*b2World)(unsafe.Pointer(world)).Contacts.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+301, __ccgo_ts+21995, int32FromInt32(418))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		contact = (*b2World)(unsafe.Pointer(world)).Contacts.Data + uintptr(contactId)*64
		contactKey = (*(*b2ContactEdge)(unsafe.Add(nil, contact+uintptr(edgeIndex)*12))).NextKey
		if (*b2Contact)(unsafe.Pointer(contact)).ShapeIdA == shapeId || (*b2Contact)(unsafe.Pointer(contact)).ShapeIdB == shapeId {
			b2DestroyContact(tls, world, contact)
		}
	}
	if (*b2Shape)(unsafe.Pointer(shape)).SensorIndex != -int32(1) {
		if v1 = !!(0 <= (*b2Shape)(unsafe.Pointer(shape)).SensorIndex && (*b2Shape)(unsafe.Pointer(shape)).SensorIndex < (*b2World)(unsafe.Pointer(world)).Sensors.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+22601, __ccgo_ts+21995, int32FromInt32(429))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		sensor = (*b2World)(unsafe.Pointer(world)).Sensors.Data + uintptr((*b2Shape)(unsafe.Pointer(shape)).SensorIndex)*56
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
					Index1:     (*b2Visitor)(unsafe.Pointer(ref)).ShapeId + int32(1),
					World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
					Generation: (*b2Visitor)(unsafe.Pointer(ref)).Generation,
				},
			}
			elementSize1 = int32(16)
			if (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Count >= (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity {
				oldSize = (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity * elementSize1
				if (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity == 0 {
					v6 = int32(8)
				} else {
					v6 = int32(2) * (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity
				}
				newCapacity = v6
				newSize = newCapacity * elementSize1
				(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data = b2GrowAlloc(tls, (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data, uint64FromInt32(oldSize), uint64FromInt32(newSize))
				(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity = newCapacity
			}
			*(*SensorEndTouchEvent)(unsafe.Add(nil, (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data+uintptr((*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Count)*16)) = event
			(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Count += int32(1)
			goto _5
		_5:
			;
			i = i + 1
		}
		// Destroy sensor
		b2Free(tls, (*b2Sensor)(unsafe.Pointer(sensor)).Hits.Data, uint64FromInt32((*b2Sensor)(unsafe.Pointer(sensor)).Hits.Capacity)*uint64(8))
		(*b2Sensor)(unsafe.Pointer(sensor)).Hits.Data = uintptrFromInt32(0)
		(*b2Sensor)(unsafe.Pointer(sensor)).Hits.Count = 0
		(*b2Sensor)(unsafe.Pointer(sensor)).Hits.Capacity = 0
		b2Free(tls, (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1.Data, uint64FromInt32((*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1.Capacity)*uint64(8))
		(*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1.Data = uintptrFromInt32(0)
		(*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1.Count = 0
		(*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1.Capacity = 0
		b2Free(tls, (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data, uint64FromInt32((*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Capacity)*uint64(8))
		(*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data = uintptrFromInt32(0)
		(*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count = 0
		(*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Capacity = 0
		v7 = (*b2World)(unsafe.Pointer(world)).Sensors.Data
		v8 = world + 2208 + 8
		v6 = (*b2Shape)(unsafe.Pointer(shape)).SensorIndex
		v9 = int32(56)
		if v1 = !!(0 <= v6 && v6 < *(*int32)(unsafe.Add(nil, v8)) && bool(__ccgo_ts+2029 != 0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+2055, __ccgo_ts+2115, int32FromInt32(153))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		*(*int32)(unsafe.Add(nil, v8)) = *(*int32)(unsafe.Add(nil, v8)) - 1
		if v6 != *(*int32)(unsafe.Add(nil, v8)) {
			memcpy(tls, v7+uintptr(v6*v9), v7+uintptr(*(*int32)(unsafe.Add(nil, v8))*v9), uint64FromInt32(v9))
			v10 = *(*int32)(unsafe.Add(nil, v8))
			goto _13
		}
		v10 = -int32(1)
		goto _13
	_13:
		movedIndex = v10
		if movedIndex != -int32(1) {
			if v2 = !!(0 <= (*b2Shape)(unsafe.Pointer(shape)).SensorIndex && (*b2Shape)(unsafe.Pointer(shape)).SensorIndex < (*b2World)(unsafe.Pointer(world)).Sensors.Count); !v2 {
				b2InternalAssert(tls, __ccgo_ts+22601, __ccgo_ts+21995, int32FromInt32(460))
			}
			_ = v2 || bool(int32FromInt32(0) != 0)
			// Fixup moved sensor
			movedSensor = (*b2World)(unsafe.Pointer(world)).Sensors.Data + uintptr((*b2Shape)(unsafe.Pointer(shape)).SensorIndex)*56
			if v3 = !!(0 <= (*b2Sensor)(unsafe.Pointer(movedSensor)).ShapeId && (*b2Sensor)(unsafe.Pointer(movedSensor)).ShapeId < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v3 {
				b2InternalAssert(tls, __ccgo_ts+21863, __ccgo_ts+21995, int32FromInt32(461))
			}
			_ = v3 || bool(int32FromInt32(0) != 0)
			otherSensorShape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b2Sensor)(unsafe.Pointer(movedSensor)).ShapeId)*280
			(*b2Shape)(unsafe.Pointer(otherSensorShape)).SensorIndex = (*b2Shape)(unsafe.Pointer(shape)).SensorIndex
		}
	}
	// Return shape to free list.
	b2FreeId(tls, world+2112, shapeId)
	(*b2Shape)(unsafe.Pointer(shape)).Id = -int32(1)
	b2ValidateSolverSets(tls, world)
}

func b2DestroyShape(tls *_Stack, shapeId ShapeId, updateBodyMass uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, shape, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_DestroyShape
	_, _, _, _ = body, shape, world, v1
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	shape = b2GetShape(tls, world, shapeId)
	// Cannot destroy a chain segment that has a parent chain shape
	if (*b2Shape)(unsafe.Pointer(shape)).Type1 == int32(b2_chainSegmentShape) && (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 128))).ChainId != -int32(1) {
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+7715, __ccgo_ts+21995, int32FromInt32(486))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_DestroyShape)(unsafe.Add(nil, bp)) = b2RecArgs_DestroyShape{
			Shape:          shapeId,
			UpdateBodyMass: updateBodyMass,
		}
		b2RecWrite_DestroyShape(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	if v1 = !!(0 <= (*b2Shape)(unsafe.Pointer(shape)).BodyId && (*b2Shape)(unsafe.Pointer(shape)).BodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+17975, __ccgo_ts+21995, int32FromInt32(492))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	body = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b2Shape)(unsafe.Pointer(shape)).BodyId)*104
	b2DestroyShapeInternal(tls, world, shape, body)
	if int32FromUint8(updateBodyMass) == int32(true1) {
		b2UpdateBodyMassData(tls, world, body)
	}
}

func b2ComputeShapeAABB(tls *_Stack, shape uintptr, xf WorldTransform) (r AABB) {
	var empty AABB
	var v1 bool
	_, _ = empty, v1
	switch (*b2Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b2_capsuleShape):
		return b2ComputeCapsuleBounds(tls, shape+128, xf)
	case int32(b2_circleShape):
		return b2ComputeCircleBounds(tls, shape+128, xf)
	case int32(b2_polygonShape):
		return b2ComputePolygonBounds(tls, shape+128, xf)
	case int32(b2_segmentShape):
		return b2ComputeSegmentBounds(tls, shape+128, xf)
	case int32(b2_chainSegmentShape):
		return b2ComputeSegmentBounds(tls, shape+128+8, xf)
	default:
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+7715, __ccgo_ts+21995, int32FromInt32(795))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		empty = AABB{}
		return empty
	}
}

func b2GetShapeCentroid(tls *_Stack, shape uintptr) (r Vec2) {
	var v1, v2, v4 Vec2
	var v3 float32
	_, _, _, _ = v1, v2, v3, v4
	switch (*b2Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b2_capsuleShape):
		v1 = (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo14_128.Capsule.Center1
		v2 = (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo14_128.Capsule.Center2
		v3 = float32FromFloat32(0.5)
		v4 = Vec2{
			X: float32((float32FromFloat32(1)-v3)*v1.X) + float32(v3*v2.X),
			Y: float32((float32FromFloat32(1)-v3)*v1.Y) + float32(v3*v2.Y),
		}
		goto _5
	_5:
		return v4
	case int32(b2_circleShape):
		return (*(*Circle)(unsafe.Add(unsafe.Pointer(shape), 128))).Center
	case int32(b2_polygonShape):
		return (*(*Polygon)(unsafe.Add(unsafe.Pointer(shape), 128))).Centroid
	case int32(b2_segmentShape):
		v1 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 128))).Point1
		v2 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 128))).Point2
		v3 = float32FromFloat32(0.5)
		v4 = Vec2{
			X: float32((float32FromFloat32(1)-v3)*v1.X) + float32(v3*v2.X),
			Y: float32((float32FromFloat32(1)-v3)*v1.Y) + float32(v3*v2.Y),
		}
		goto _10
	_10:
		return v4
	case int32(b2_chainSegmentShape):
		v1 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 128))).Segment.Point1
		v2 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 128))).Segment.Point2
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
	var v8 bool
	_, _, _, _, _, _, _, _, _, _, _, _ = count, i, next, perimeter, points, prev, v1, v2, v3, v5, v6, v8
	switch (*b2Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b2_capsuleShape):
		v1 = (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo14_128.Capsule.Center1
		v2 = (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo14_128.Capsule.Center2
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
		return float32(float32FromFloat32(2)*v6) + float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359))*(*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo14_128.Capsule.Radius)
	case int32(b2_circleShape):
		return float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * (*(*Circle)(unsafe.Add(unsafe.Pointer(shape), 128))).Radius)
	case int32(b2_polygonShape):
		points = shape + 128
		count = (*(*Polygon)(unsafe.Add(unsafe.Pointer(shape), 128))).Count
		perimeter = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * (*(*Polygon)(unsafe.Add(unsafe.Pointer(shape), 128))).Radius)
		if v8 = !!(count > int32FromInt32(0)); !v8 {
			b2InternalAssert(tls, __ccgo_ts+9458, __ccgo_ts+21995, int32FromInt32(836))
		}
		_ = v8 || bool(int32FromInt32(0) != 0)
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
			goto _13
		_13:
			v5 = v3
			v6 = sqrtf(tls, float32(v5.X*v5.X)+float32(v5.Y*v5.Y))
			goto _16
		_16:
			perimeter = perimeter + v6
			prev = next
			goto _9
		_9:
			;
			i = i + 1
		}
		return perimeter
	case int32(b2_segmentShape):
		v1 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 128))).Point1
		v2 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 128))).Point2
		v3 = Vec2{
			X: v1.X - v2.X,
			Y: v1.Y - v2.Y,
		}
		goto _20
	_20:
		v5 = v3
		v6 = sqrtf(tls, float32(v5.X*v5.X)+float32(v5.Y*v5.Y))
		goto _23
	_23:
		return float32(float32FromFloat32(2) * v6)
	case int32(b2_chainSegmentShape):
		v1 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 128))).Segment.Point1
		v2 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 128))).Segment.Point2
		v3 = Vec2{
			X: v1.X - v2.X,
			Y: v1.Y - v2.Y,
		}
		goto _27
	_27:
		v5 = v3
		v6 = sqrtf(tls, float32(v5.X*v5.X)+float32(v5.Y*v5.Y))
		goto _30
	_30:
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
	var lower, projectedLength, upper, value, value1, value11, value2, value21, v10, v12, v16, v7, v9 float32
	var points uintptr
	var v13 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axis, count, i, lower, points, projectedLength, upper, value, value1, value11, value2, value21, v1, v10, v12, v13, v16, v2, v3, v5, v6, v7, v9
	switch (*b2Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b2_capsuleShape):
		v1 = (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo14_128.Capsule.Center2
		v2 = (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo14_128.Capsule.Center1
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
		return projectedLength + float32(float32FromFloat32(2)*(*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo14_128.Capsule.Radius)
	case int32(b2_circleShape):
		return float32(float32FromFloat32(2) * (*(*Circle)(unsafe.Add(unsafe.Pointer(shape), 128))).Radius)
	case int32(b2_polygonShape):
		points = shape + 128
		count = (*(*Polygon)(unsafe.Add(unsafe.Pointer(shape), 128))).Count
		if v13 = !!(count > int32FromInt32(0)); !v13 {
			b2InternalAssert(tls, __ccgo_ts+9458, __ccgo_ts+21995, int32FromInt32(875))
		}
		_ = v13 || bool(int32FromInt32(0) != 0)
		v1 = *(*Vec2)(unsafe.Add(nil, points))
		v2 = line
		v7 = float32(v1.X*v2.X) + float32(v1.Y*v2.Y)
		goto _17
	_17:
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
			goto _22
		_22:
			value = v9
			v7 = lower
			v9 = value
			if v7 < v9 {
				v12 = v7
			} else {
				v12 = v9
			}
			v10 = v12
			goto _26
		_26:
			lower = v10
			v7 = upper
			v9 = value
			if v7 > v9 {
				v12 = v7
			} else {
				v12 = v9
			}
			v10 = v12
			goto _31
		_31:
			upper = v10
			goto _18
		_18:
			;
			i = i + 1
		}
		return upper - lower + float32(float32FromFloat32(2)*(*(*Polygon)(unsafe.Add(unsafe.Pointer(shape), 128))).Radius)
	case int32(b2_segmentShape):
		v1 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 128))).Point1
		v2 = line
		v7 = float32(v1.X*v2.X) + float32(v1.Y*v2.Y)
		goto _36
	_36:
		value1 = v7
		v3 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 128))).Point2
		v5 = line
		v9 = float32(v3.X*v5.X) + float32(v3.Y*v5.Y)
		goto _40
	_40:
		value2 = v9
		v10 = value2 - value1
		if v10 < float32FromInt32(0) {
			v16 = -v10
		} else {
			v16 = v10
		}
		v12 = v16
		goto _43
	_43:
		return v12
	case int32(b2_chainSegmentShape):
		v1 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 128))).Segment.Point1
		v2 = line
		v7 = float32(v1.X*v2.X) + float32(v1.Y*v2.Y)
		goto _48
	_48:
		value11 = v7
		v3 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 128))).Segment.Point2
		v5 = line
		v9 = float32(v3.X*v5.X) + float32(v3.Y*v5.Y)
		goto _52
	_52:
		value21 = v9
		v10 = value21 - value11
		if v10 < float32FromInt32(0) {
			v16 = -v10
		} else {
			v16 = v10
		}
		v12 = v16
		goto _55
	_55:
		return v12
	default:
		return float32FromFloat32(0)
	}
}

func b2ComputeShapeMass(tls *_Stack, shape uintptr) (r MassData) {
	switch (*b2Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b2_capsuleShape):
		return b2ComputeCapsuleMass(tls, shape+128, (*b2Shape)(unsafe.Pointer(shape)).Density)
	case int32(b2_circleShape):
		return b2ComputeCircleMass(tls, shape+128, (*b2Shape)(unsafe.Pointer(shape)).Density)
	case int32(b2_polygonShape):
		return b2ComputePolygonMass(tls, shape+128, (*b2Shape)(unsafe.Pointer(shape)).Density)
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
		radius = (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo14_128.Capsule.Radius
		extent.MinExtent = radius
		v1 = (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo14_128.Capsule.Center1
		v21 = localCenter
		v3 = Vec2{
			X: v1.X - v21.X,
			Y: v1.Y - v21.Y,
		}
		goto _4
	_4:
		c1 = v3
		v5 = (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo14_128.Capsule.Center2
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
		radius1 = (*(*Circle)(unsafe.Add(unsafe.Pointer(shape), 128))).Radius
		extent.MinExtent = radius1
		v1 = (*(*Circle)(unsafe.Add(unsafe.Pointer(shape), 128))).Center
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
		poly = shape + 128
		minExtent = float32(float32(float32FromFloat32(100)*float32FromFloat32(1000)) * b2_lengthUnitsPerMeter)
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
		v1 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 128))).Point1
		v21 = localCenter
		v3 = Vec2{
			X: v1.X - v21.X,
			Y: v1.Y - v21.Y,
		}
		goto _56
	_56:
		c11 = v3
		v5 = (*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 128))).Point2
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
		v1 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 128))).Segment.Point1
		v21 = localCenter
		v3 = Vec2{
			X: v1.X - v21.X,
			Y: v1.Y - v21.Y,
		}
		goto _75
	_75:
		c12 = v3
		v5 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 128))).Segment.Point2
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
		output = b2CastRayCapsule(tls, shape+128, bp)
	case int32(b2_circleShape):
		output = b2CastRayCircle(tls, shape+128, bp)
	case int32(b2_polygonShape):
		output = b2CastRayPolygon(tls, shape+128, bp)
	case int32(b2_segmentShape):
		output = b2CastRaySegment(tls, shape+128, bp, boolUint8(false1 != 0))
	case int32(b2_chainSegmentShape):
		output = b2CastRaySegment(tls, shape+128+8, bp, boolUint8(true1 != 0))
	default:
		return output
	}
	// The output point stays in the frame of the input transform, a caller chosen frame that is
	// typically re-centered near the origin.
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

func b2ShapeCastShape(tls *_Stack, input uintptr, shape uintptr, transform Transform) (r1 CastOutput) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var approximateCentroid, edge, r, v11, v12, v13, v16, v3, v4, v7, v8 Vec2
	var i, i1 int32
	var output CastOutput
	var vx, vy, x, y, v15 float32
	var v2 Transform
	var v6 Rot
	var _ /* localInput at bp+0 */ ShapeCastInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = approximateCentroid, edge, i, i1, output, r, vx, vy, x, y, v11, v12, v13, v15, v16, v2, v3, v4, v6, v7, v8
	output = CastOutput{}
	if (*ShapeCastInput)(unsafe.Pointer(input)).Proxy.Count == 0 {
		return output
	}
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
	switch (*b2Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b2_capsuleShape):
		output = b2CastShapeCapsule(tls, shape+128, bp)
	case int32(b2_circleShape):
		output = b2CastShapeCircle(tls, shape+128, bp)
	case int32(b2_polygonShape):
		output = b2CastShapePolygon(tls, shape+128, bp)
	case int32(b2_segmentShape):
		output = b2CastShapeSegment(tls, shape+128, bp)
	case int32(b2_chainSegmentShape):
		// Check for back side collision
		approximateCentroid = *(*Vec2)(unsafe.Add(nil, bp))
		i1 = int32(1)
		for {
			if !(i1 < (*(*ShapeCastInput)(unsafe.Add(nil, bp))).Proxy.Count) {
				break
			}
			v3 = approximateCentroid
			v4 = *(*Vec2)(unsafe.Add(nil, bp+uintptr(i1)*8))
			v7 = Vec2{
				X: v3.X + v4.X,
				Y: v3.Y + v4.Y,
			}
			goto _14
		_14:
			approximateCentroid = v7
			goto _10
		_10:
			;
			i1 = i1 + 1
		}
		v15 = float32FromFloat32(1) / float32((*(*ShapeCastInput)(unsafe.Add(nil, bp))).Proxy.Count)
		v3 = approximateCentroid
		v4 = Vec2{
			X: float32(v15 * v3.X),
			Y: float32(v15 * v3.Y),
		}
		goto _18
	_18:
		approximateCentroid = v4
		v3 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 128))).Segment.Point2
		v4 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 128))).Segment.Point1
		v7 = Vec2{
			X: v3.X - v4.X,
			Y: v3.Y - v4.Y,
		}
		goto _22
	_22:
		edge = v7
		v8 = approximateCentroid
		v11 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 128))).Segment.Point1
		v12 = Vec2{
			X: v8.X - v11.X,
			Y: v8.Y - v11.Y,
		}
		goto _26
	_26:
		r = v12
		v13 = r
		v16 = edge
		v15 = float32(v13.X*v16.Y) - float32(v13.Y*v16.X)
		goto _30
	_30:
		if v15 < float32FromFloat32(0) {
			// Shape cast starts behind
			return output
		}
		output = b2CastShapeSegment(tls, shape+128+8, bp)
	default:
		return output
	}
	// Same frame contract as b2RayCastShape, the point stays in the input transform frame
	v2 = transform
	v3 = output.Point
	x = float32(v2.Q.C*v3.X) - float32(v2.Q.S*v3.Y) + v2.P.X
	y = float32(v2.Q.S*v3.X) + float32(v2.Q.C*v3.Y) + v2.P.Y
	v4 = Vec2{
		X: x,
		Y: y,
	}
	goto _34
_34:
	output.Point = v4
	v6 = transform.Q
	v3 = output.Normal
	v4 = Vec2{
		X: float32(v6.C*v3.X) - float32(v6.S*v3.Y),
		Y: float32(v6.S*v3.X) + float32(v6.C*v3.Y),
	}
	goto _38
_38:
	output.Normal = v4
	return output
}

func b2CreateShapeProxy(tls *_Stack, world uintptr, shape uintptr, type1 BodyType, transform WorldTransform, forcePairCreation uint8) {
	var fatAABB uintptr
	var v1 bool
	_, _ = fatAABB, v1
	if v1 = !!((*b2Shape)(unsafe.Pointer(shape)).ProxyKey == -int32FromInt32(1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+23429, __ccgo_ts+21995, int32FromInt32(1133))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	fatAABB = (*b2World)(unsafe.Pointer(world)).FatAABBs.Data + uintptr((*b2Shape)(unsafe.Pointer(shape)).Id)*16
	b2UpdateShapeAABBs(tls, shape, fatAABB, transform, type1)
	// Create proxies in the broad-phase.
	(*b2Shape)(unsafe.Pointer(shape)).ProxyKey = b2BroadPhase_CreateProxy(tls, world+40, type1, *(*AABB)(unsafe.Add(nil, fatAABB)), (*b2Shape)(unsafe.Pointer(shape)).Filter.CategoryBits, (*b2Shape)(unsafe.Pointer(shape)).Id, forcePairCreation)
	if v1 = !!((*b2Shape)(unsafe.Pointer(shape)).ProxyKey&int32FromInt32(3) < int32(b2_bodyTypeCount)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+23462, __ccgo_ts+21995, int32FromInt32(1141))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
}

func b2DestroyShapeProxy(tls *_Stack, shape uintptr, bp uintptr) {
	if (*b2Shape)(unsafe.Pointer(shape)).ProxyKey != -int32(1) {
		b2BroadPhase_DestroyProxy(tls, bp, (*b2Shape)(unsafe.Pointer(shape)).ProxyKey)
		(*b2Shape)(unsafe.Pointer(shape)).ProxyKey = -int32(1)
	}
}

func b2MakeShapeDistanceProxy(tls *_Stack, shape uintptr) (r ShapeProxy) {
	var empty ShapeProxy
	var v1 bool
	_, _ = empty, v1
	switch (*b2Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b2_capsuleShape):
		return b2MakeProxy(tls, shape+128, int32(2), (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo14_128.Capsule.Radius)
	case int32(b2_circleShape):
		return b2MakeProxy(tls, shape+128, int32(1), (*(*Circle)(unsafe.Add(unsafe.Pointer(shape), 128))).Radius)
	case int32(b2_polygonShape):
		return b2MakeProxy(tls, shape+128, (*(*Polygon)(unsafe.Add(unsafe.Pointer(shape), 128))).Count, (*(*Polygon)(unsafe.Add(unsafe.Pointer(shape), 128))).Radius)
	case int32(b2_segmentShape):
		return b2MakeProxy(tls, shape+128, int32(2), float32FromFloat32(0))
	case int32(b2_chainSegmentShape):
		return b2MakeProxy(tls, shape+128+8, int32(2), float32FromFloat32(0))
	default:
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+7715, __ccgo_ts+21995, int32FromInt32(1169))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
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

func b2Shape_TestPoint(tls *_Stack, shapeId ShapeId, point Pos) (r uint8) {
	var localPoint, v3 Vec2
	var result uint8
	var shape, world uintptr
	var transform, v1 WorldTransform
	var vx, vy float32
	var v2 Pos
	_, _, _, _, _, _, _, _, _, _ = localPoint, result, shape, transform, vx, vy, world, v1, v2, v3
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
		result = b2PointInCapsule(tls, shape+128, localPoint)
	case int32(b2_circleShape):
		result = b2PointInCircle(tls, shape+128, localPoint)
	case int32(b2_polygonShape):
		result = b2PointInPolygon(tls, shape+128, localPoint)
	default:
		result = boolUint8(false1 != 0)
		break
	}
	return result
}

func b2Shape_CastRay(tls *_Stack, shapeId ShapeId, origin Pos, translation Vec2) (r1 WorldCastOutput) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var local CastOutput
	var output WorldCastOutput
	var r, transform, v5 Transform
	var shape, world uintptr
	var v1 bool
	var v3 WorldTransform
	var v4, v7 Pos
	var v8 Vec2
	var _ /* input at bp+0 */ RayCastInput
	_, _, _, _, _, _, _, _, _, _, _, _ = local, output, r, shape, transform, world, v1, v3, v4, v5, v7, v8
	if v1 = !!(b2IsValidPos(tls, origin) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+19491, __ccgo_ts+21995, int32FromInt32(1243))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b2IsValidVec2(tls, translation) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+19578, __ccgo_ts+21995, int32FromInt32(1244))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	v3 = b2GetBodyTransform(tls, world, (*b2Shape)(unsafe.Pointer(shape)).BodyId)
	v4 = origin
	r.Q = v3.Q
	r.P = Vec2{
		X: v3.P.X - v4.X,
		Y: v3.P.Y - v4.Y,
	}
	v5 = r
	goto _6
_6:
	// Re-center on the origin so the cast runs in float precision
	transform = v5
	// The ray starts at the origin, so its origin in the re-centered frame is zero
	*(*RayCastInput)(unsafe.Add(nil, bp)) = RayCastInput{
		Origin:      b2Vec2_zero,
		Translation: translation,
		MaxFraction: float32FromFloat32(1),
	}
	// Lift the re-centered float result back to a world position
	local = b2RayCastShape(tls, bp, shape, transform)
	output.Normal = local.Normal
	v4 = origin
	v8 = local.Point
	v7 = Pos{
		X: v4.X + v8.X,
		Y: v4.Y + v8.Y,
	}
	goto _10
_10:
	output.Point = v7
	output.Fraction = local.Fraction
	output.Iterations = local.Iterations
	output.Hit = local.Hit
	return output
}

func b2Shape_SetDensity(tls *_Stack, shapeId ShapeId, density float32, updateBodyMass uint8) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var body, shape, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_ShapeSetDensity
	_, _, _, _ = body, shape, world, v1
	if boolInt32(b2IsValidFloat(tls, density) != 0 && density >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+23514, uintptr(unsafe.Pointer(&__func__70))))
		if v1 = !!(b2IsValidFloat(tls, density) != 0 && density >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+23514, __ccgo_ts+21995, int32FromInt32(1269))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_ShapeSetDensity)(unsafe.Add(nil, bp)) = b2RecArgs_ShapeSetDensity{
			Shape:          shapeId,
			Density:        density,
			UpdateBodyMass: updateBodyMass,
		}
		b2RecWrite_ShapeSetDensity(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	shape = b2GetShape(tls, world, shapeId)
	if density == (*b2Shape)(unsafe.Pointer(shape)).Density {
		// early return to avoid expensive function
		return
	}
	(*b2Shape)(unsafe.Pointer(shape)).Density = density
	if int32FromUint8(updateBodyMass) == int32(true1) {
		if v1 = !!(0 <= (*b2Shape)(unsafe.Pointer(shape)).BodyId && (*b2Shape)(unsafe.Pointer(shape)).BodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+17975, __ccgo_ts+21995, int32FromInt32(1290))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		body = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b2Shape)(unsafe.Pointer(shape)).BodyId)*104
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
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var shape, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_ShapeSetFriction
	_, _, _ = shape, world, v1
	if boolInt32(b2IsValidFloat(tls, friction) != 0 && friction >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+23559, uintptr(unsafe.Pointer(&__func__71))))
		if v1 = !!(b2IsValidFloat(tls, friction) != 0 && friction >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+23559, __ccgo_ts+21995, int32FromInt32(1304))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+21995, int32FromInt32(1307))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_ShapeSetFriction)(unsafe.Add(nil, bp)) = b2RecArgs_ShapeSetFriction{
			Shape:    shapeId,
			Friction: friction,
		}
		b2RecWrite_ShapeSetFriction(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	shape = b2GetShape(tls, world, shapeId)
	(*b2Shape)(unsafe.Pointer(shape)).Material.Friction = friction
}

func b2Shape_GetFriction(tls *_Stack, shapeId ShapeId) (r float32) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	return (*b2Shape)(unsafe.Pointer(shape)).Material.Friction
}

func b2Shape_SetRestitution(tls *_Stack, shapeId ShapeId, restitution float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var shape, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_ShapeSetRestitution
	_, _, _ = shape, world, v1
	if boolInt32(b2IsValidFloat(tls, restitution) != 0 && restitution >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+23606, uintptr(unsafe.Pointer(&__func__72))))
		if v1 = !!(b2IsValidFloat(tls, restitution) != 0 && restitution >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+23606, __ccgo_ts+21995, int32FromInt32(1328))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+21995, int32FromInt32(1331))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_ShapeSetRestitution)(unsafe.Add(nil, bp)) = b2RecArgs_ShapeSetRestitution{
			Shape:       shapeId,
			Restitution: restitution,
		}
		b2RecWrite_ShapeSetRestitution(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	shape = b2GetShape(tls, world, shapeId)
	(*b2Shape)(unsafe.Pointer(shape)).Material.Restitution = restitution
}

func b2Shape_GetRestitution(tls *_Stack, shapeId ShapeId) (r float32) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	return (*b2Shape)(unsafe.Pointer(shape)).Material.Restitution
}

func b2Shape_SetUserMaterial(tls *_Stack, shapeId ShapeId, material uint64) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var shape, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_ShapeSetUserMaterial
	_, _, _ = shape, world, v1
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+21995, int32FromInt32(1353))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_ShapeSetUserMaterial)(unsafe.Add(nil, bp)) = b2RecArgs_ShapeSetUserMaterial{
			Shape:    shapeId,
			Material: material,
		}
		b2RecWrite_ShapeSetUserMaterial(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	shape = b2GetShape(tls, world, shapeId)
	(*b2Shape)(unsafe.Pointer(shape)).Material.UserMaterialId = material
}

func b2Shape_GetUserMaterial(tls *_Stack, shapeId ShapeId) (r uint64) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	return (*b2Shape)(unsafe.Pointer(shape)).Material.UserMaterialId
}

func b2Shape_GetSurfaceMaterial(tls *_Stack, shapeId ShapeId) (r SurfaceMaterial) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	return (*b2Shape)(unsafe.Pointer(shape)).Material
}

func b2Shape_SetSurfaceMaterial(tls *_Stack, shapeId ShapeId, surfaceMaterial uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var shape, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_ShapeSetSurfaceMaterial
	_, _, _ = shape, world, v1
	if int32FromUint8(b2IsValidMaterial(tls, surfaceMaterial)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+48, __ccgo_ts+23659, uintptr(unsafe.Pointer(&__func__73))))
		if v1 = !!(b2IsValidMaterial(tls, surfaceMaterial) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+23659, __ccgo_ts+21995, int32FromInt32(1381))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_ShapeSetSurfaceMaterial)(unsafe.Add(nil, bp)) = b2RecArgs_ShapeSetSurfaceMaterial{
			Shape:    shapeId,
			Material: *(*SurfaceMaterial)(unsafe.Add(nil, surfaceMaterial)),
		}
		b2RecWrite_ShapeSetSurfaceMaterial(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	shape = b2GetShape(tls, world, shapeId)
	(*b2Shape)(unsafe.Pointer(shape)).Material = *(*SurfaceMaterial)(unsafe.Add(nil, surfaceMaterial))
}

func b2Shape_GetFilter(tls *_Stack, shapeId ShapeId) (r Filter) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	return (*b2Shape)(unsafe.Pointer(shape)).Filter
}

func b2Shape_SetFilter(tls *_Stack, shapeId ShapeId, filter Filter) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var destroyProxy uint8
	var shape, world uintptr
	var _ /* _a at bp+0 */ b2RecArgs_ShapeSetFilter
	_, _, _ = destroyProxy, shape, world
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_ShapeSetFilter)(unsafe.Add(nil, bp)) = b2RecArgs_ShapeSetFilter{
			Shape:  shapeId,
			Filter: filter,
		}
		b2RecWrite_ShapeSetFilter(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	shape = b2GetShape(tls, world, shapeId)
	if filter.MaskBits == (*b2Shape)(unsafe.Pointer(shape)).Filter.MaskBits && filter.CategoryBits == (*b2Shape)(unsafe.Pointer(shape)).Filter.CategoryBits && filter.GroupIndex == (*b2Shape)(unsafe.Pointer(shape)).Filter.GroupIndex {
		return
	}
	// If the category bits change, I need to destroy the proxy because it affects the tree sorting.
	destroyProxy = boolUint8(filter.CategoryBits != (*b2Shape)(unsafe.Pointer(shape)).Filter.CategoryBits)
	(*b2Shape)(unsafe.Pointer(shape)).Filter = filter
	b2ResetProxy(tls, world, shape, destroyProxy)
	// note: this does not immediately update sensor overlaps. Instead sensor
	// overlaps are updated the next time step
}

func b2Shape_EnableSensorEvents(tls *_Stack, shapeId ShapeId, flag uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var shape, world uintptr
	var _ /* _a at bp+0 */ b2RecArgs_ShapeEnableSensorEvents
	_, _ = shape, world
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_ShapeEnableSensorEvents)(unsafe.Add(nil, bp)) = b2RecArgs_ShapeEnableSensorEvents{
			Shape: shapeId,
			Flag:  flag,
		}
		b2RecWrite_ShapeEnableSensorEvents(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
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
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var shape, world uintptr
	var _ /* _a at bp+0 */ b2RecArgs_ShapeEnableContactEvents
	_, _ = shape, world
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_ShapeEnableContactEvents)(unsafe.Add(nil, bp)) = b2RecArgs_ShapeEnableContactEvents{
			Shape: shapeId,
			Flag:  flag,
		}
		b2RecWrite_ShapeEnableContactEvents(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
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
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var shape, world uintptr
	var _ /* _a at bp+0 */ b2RecArgs_ShapeEnablePreSolveEvents
	_, _ = shape, world
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_ShapeEnablePreSolveEvents)(unsafe.Add(nil, bp)) = b2RecArgs_ShapeEnablePreSolveEvents{
			Shape: shapeId,
			Flag:  flag,
		}
		b2RecWrite_ShapeEnablePreSolveEvents(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
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
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var shape, world uintptr
	var _ /* _a at bp+0 */ b2RecArgs_ShapeEnableHitEvents
	_, _ = shape, world
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_ShapeEnableHitEvents)(unsafe.Add(nil, bp)) = b2RecArgs_ShapeEnableHitEvents{
			Shape: shapeId,
			Flag:  flag,
		}
		b2RecWrite_ShapeEnableHitEvents(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
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
	var v1 bool
	_, _, _ = shape, world, v1
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	if v1 = !!((*b2Shape)(unsafe.Pointer(shape)).Type1 == int32(b2_circleShape)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+23696, __ccgo_ts+21995, int32FromInt32(1570))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return *(*Circle)(unsafe.Add(unsafe.Pointer(shape), 128))
}

func b2Shape_GetSegment(tls *_Stack, shapeId ShapeId) (r Segment) {
	var shape, world uintptr
	var v1 bool
	_, _, _ = shape, world, v1
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	if v1 = !!((*b2Shape)(unsafe.Pointer(shape)).Type1 == int32(b2_segmentShape)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+23726, __ccgo_ts+21995, int32FromInt32(1578))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return *(*Segment)(unsafe.Add(unsafe.Pointer(shape), 128))
}

func b2Shape_GetChainSegment(tls *_Stack, shapeId ShapeId) (r ChainSegment) {
	var shape, world uintptr
	var v1 bool
	_, _, _ = shape, world, v1
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	if v1 = !!((*b2Shape)(unsafe.Pointer(shape)).Type1 == int32(b2_chainSegmentShape)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+23757, __ccgo_ts+21995, int32FromInt32(1586))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return *(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 128))
}

func b2Shape_GetCapsule(tls *_Stack, shapeId ShapeId) (r Capsule) {
	var shape, world uintptr
	var v1 bool
	_, _, _ = shape, world, v1
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	if v1 = !!((*b2Shape)(unsafe.Pointer(shape)).Type1 == int32(b2_capsuleShape)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+23793, __ccgo_ts+21995, int32FromInt32(1594))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo14_128.Capsule
}

func b2Shape_GetPolygon(tls *_Stack, shapeId ShapeId) (r Polygon) {
	var shape, world uintptr
	var v1 bool
	_, _, _ = shape, world, v1
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	if v1 = !!((*b2Shape)(unsafe.Pointer(shape)).Type1 == int32(b2_polygonShape)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+23824, __ccgo_ts+21995, int32FromInt32(1602))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return *(*Polygon)(unsafe.Add(unsafe.Pointer(shape), 128))
}

func b2Shape_SetCircle(tls *_Stack, shapeId ShapeId, circle uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var destroyProxy uint8
	var shape, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_ShapeSetCircle
	_, _, _, _ = destroyProxy, shape, world, v1
	if int32FromUint8(b2IsValidCircle(tls, circle)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+32, __ccgo_ts+22373, uintptr(unsafe.Pointer(&__func__74))))
		if v1 = !!(b2IsValidCircle(tls, circle) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+22373, __ccgo_ts+21995, int32FromInt32(1608))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_ShapeSetCircle)(unsafe.Add(nil, bp)) = b2RecArgs_ShapeSetCircle{
			Shape:  shapeId,
			Circle: *(*Circle)(unsafe.Add(nil, circle)),
		}
		b2RecWrite_ShapeSetCircle(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	shape = b2GetShape(tls, world, shapeId)
	*(*Circle)(unsafe.Add(unsafe.Pointer(shape), 128)) = *(*Circle)(unsafe.Add(nil, circle))
	(*b2Shape)(unsafe.Pointer(shape)).Type1 = int32(b2_circleShape)
	(*b2Shape)(unsafe.Pointer(shape)).AabbMargin = b2ComputeShapeMargin(tls, shape)
	destroyProxy = boolUint8(true1 != 0)
	b2ResetProxy(tls, world, shape, destroyProxy)
}

func b2Shape_SetCapsule(tls *_Stack, shapeId ShapeId, capsule uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var c, v2, v3 Vec2
	var destroyProxy uint8
	var lengthSqr, v4 float32
	var shape, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_ShapeSetCapsule
	_, _, _, _, _, _, _, _, _ = c, destroyProxy, lengthSqr, shape, world, v1, v2, v3, v4
	if int32FromUint8(b2IsValidCapsule(tls, capsule)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+40, __ccgo_ts+22399, uintptr(unsafe.Pointer(&__func__121))))
		if v1 = !!(b2IsValidCapsule(tls, capsule) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+22399, __ccgo_ts+21995, int32FromInt32(1629))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	v2 = (*Capsule)(unsafe.Pointer(capsule)).Center1
	v3 = (*Capsule)(unsafe.Pointer(capsule)).Center2
	c = Vec2{
		X: v3.X - v2.X,
		Y: v3.Y - v2.Y,
	}
	v4 = float32(c.X*c.X) + float32(c.Y*c.Y)
	goto _5
_5:
	lengthSqr = v4
	if lengthSqr <= float32(float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_ShapeSetCapsule)(unsafe.Add(nil, bp)) = b2RecArgs_ShapeSetCapsule{
			Shape:   shapeId,
			Capsule: *(*Capsule)(unsafe.Add(nil, capsule)),
		}
		b2RecWrite_ShapeSetCapsule(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	shape = b2GetShape(tls, world, shapeId)
	(*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo14_128.Capsule = *(*Capsule)(unsafe.Add(nil, capsule))
	(*b2Shape)(unsafe.Pointer(shape)).Type1 = int32(b2_capsuleShape)
	(*b2Shape)(unsafe.Pointer(shape)).AabbMargin = b2ComputeShapeMargin(tls, shape)
	destroyProxy = boolUint8(true1 != 0)
	b2ResetProxy(tls, world, shape, destroyProxy)
}

func b2Shape_SetSegment(tls *_Stack, shapeId ShapeId, segment uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var destroyProxy uint8
	var shape, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_ShapeSetSegment
	_, _, _, _ = destroyProxy, shape, world, v1
	if int32FromUint8(b2IsValidSegment(tls, segment)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+32, __ccgo_ts+22455, uintptr(unsafe.Pointer(&__func__75))))
		if v1 = !!(b2IsValidSegment(tls, segment) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+22455, __ccgo_ts+21995, int32FromInt32(1656))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_ShapeSetSegment)(unsafe.Add(nil, bp)) = b2RecArgs_ShapeSetSegment{
			Shape:   shapeId,
			Segment: *(*Segment)(unsafe.Add(nil, segment)),
		}
		b2RecWrite_ShapeSetSegment(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	shape = b2GetShape(tls, world, shapeId)
	*(*Segment)(unsafe.Add(unsafe.Pointer(shape), 128)) = *(*Segment)(unsafe.Add(nil, segment))
	(*b2Shape)(unsafe.Pointer(shape)).Type1 = int32(b2_segmentShape)
	(*b2Shape)(unsafe.Pointer(shape)).AabbMargin = b2ComputeShapeMargin(tls, shape)
	destroyProxy = boolUint8(true1 != 0)
	b2ResetProxy(tls, world, shape, destroyProxy)
}

func b2Shape_SetPolygon(tls *_Stack, shapeId ShapeId, polygon uintptr) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var destroyProxy uint8
	var shape, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_ShapeSetPolygon
	_, _, _, _ = destroyProxy, shape, world, v1
	if int32FromUint8(b2IsValidPolygon(tls, polygon)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+160, __ccgo_ts+22427, uintptr(unsafe.Pointer(&__func__76))))
		if v1 = !!(b2IsValidPolygon(tls, polygon) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+22427, __ccgo_ts+21995, int32FromInt32(1677))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_ShapeSetPolygon)(unsafe.Add(nil, bp)) = b2RecArgs_ShapeSetPolygon{
			Shape:   shapeId,
			Polygon: *(*Polygon)(unsafe.Add(nil, polygon)),
		}
		b2RecWrite_ShapeSetPolygon(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	shape = b2GetShape(tls, world, shapeId)
	*(*Polygon)(unsafe.Add(unsafe.Pointer(shape), 128)) = *(*Polygon)(unsafe.Add(nil, polygon))
	(*b2Shape)(unsafe.Pointer(shape)).Type1 = int32(b2_polygonShape)
	(*b2Shape)(unsafe.Pointer(shape)).AabbMargin = b2ComputeShapeMargin(tls, shape)
	destroyProxy = boolUint8(true1 != 0)
	b2ResetProxy(tls, world, shape, destroyProxy)
}

func b2Shape_SetChainSegment(tls *_Stack, shapeId ShapeId, chainSegment uintptr) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var c, v3, v4 Vec2
	var destroyProxy uint8
	var lengthSqr, v5 float32
	var shape, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_ShapeSetChainSegment
	_, _, _, _, _, _, _, _, _ = c, destroyProxy, lengthSqr, shape, world, v1, v3, v4, v5
	if int32FromUint8(b2IsValidChainSegment(tls, chainSegment)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+56, __ccgo_ts+22483, uintptr(unsafe.Pointer(&__func__122))))
		if v1 = !!(b2IsValidChainSegment(tls, chainSegment) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+22483, __ccgo_ts+21995, int32FromInt32(1698))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	shape = b2GetShape(tls, world, shapeId)
	// Cannot modify a chain segment that has a parent chain shape
	if (*b2Shape)(unsafe.Pointer(shape)).Type1 == int32(b2_chainSegmentShape) && (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 128))).ChainId != -int32(1) {
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+7715, __ccgo_ts+21995, int32FromInt32(1711))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	v3 = (*ChainSegment)(unsafe.Pointer(chainSegment)).Segment.Point1
	v4 = (*ChainSegment)(unsafe.Pointer(chainSegment)).Segment.Point2
	c = Vec2{
		X: v4.X - v3.X,
		Y: v4.Y - v3.Y,
	}
	v5 = float32(c.X*c.X) + float32(c.Y*c.Y)
	goto _6
_6:
	lengthSqr = v5
	if lengthSqr <= float32(float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_ShapeSetChainSegment)(unsafe.Add(nil, bp)) = b2RecArgs_ShapeSetChainSegment{
			Shape:        shapeId,
			ChainSegment: *(*ChainSegment)(unsafe.Add(nil, chainSegment)),
		}
		b2RecWrite_ShapeSetChainSegment(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 128)) = *(*ChainSegment)(unsafe.Add(nil, chainSegment))
	(*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 128))).ChainId = -int32(1)
	(*b2Shape)(unsafe.Pointer(shape)).Type1 = int32(b2_chainSegmentShape)
	(*b2Shape)(unsafe.Pointer(shape)).AabbMargin = b2ComputeShapeMargin(tls, shape)
	destroyProxy = boolUint8(true1 != 0)
	b2ResetProxy(tls, world, shape, destroyProxy)
}

func b2Shape_GetParentChain(tls *_Stack, shapeId ShapeId) (r ChainId) {
	var chain, shape, world uintptr
	var chainId int32
	var id ChainId
	var v1 bool
	_, _, _, _, _, _ = chain, chainId, id, shape, world, v1
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	shape = b2GetShape(tls, world, shapeId)
	if (*b2Shape)(unsafe.Pointer(shape)).Type1 == int32(b2_chainSegmentShape) {
		chainId = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 128))).ChainId
		if chainId != -int32(1) {
			if v1 = !!(0 <= chainId && chainId < (*b2World)(unsafe.Pointer(world)).ChainShapes.Count); !v1 {
				b2InternalAssert(tls, __ccgo_ts+1966, __ccgo_ts+21995, int32FromInt32(1741))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			chain = (*b2World)(unsafe.Pointer(world)).ChainShapes.Data + uintptr(chainId)*32
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
	var body, shape, world uintptr
	var v1 bool
	_, _, _, _ = body, shape, world, v1
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return 0
	}
	shape = b2GetShape(tls, world, shapeId)
	if (*b2Shape)(unsafe.Pointer(shape)).SensorIndex != -int32(1) {
		return 0
	}
	if v1 = !!(0 <= (*b2Shape)(unsafe.Pointer(shape)).BodyId && (*b2Shape)(unsafe.Pointer(shape)).BodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+17975, __ccgo_ts+21995, int32FromInt32(1764))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	body = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b2Shape)(unsafe.Pointer(shape)).BodyId)*104
	// Conservative and fast
	return (*b2Body)(unsafe.Pointer(body)).ContactCount
}

func b2Shape_GetContactData(tls *_Stack, shapeId ShapeId, contactData uintptr, capacity int32) (r int32) {
	var body, contact, contactSim, shape, shapeA, shapeB, world uintptr
	var contactId, contactKey, edgeIndex, index int32
	var v1, v2 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _ = body, contact, contactId, contactKey, contactSim, edgeIndex, index, shape, shapeA, shapeB, world, v1, v2
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return 0
	}
	shape = b2GetShape(tls, world, shapeId)
	if (*b2Shape)(unsafe.Pointer(shape)).SensorIndex != -int32(1) {
		return 0
	}
	if v1 = !!(0 <= (*b2Shape)(unsafe.Pointer(shape)).BodyId && (*b2Shape)(unsafe.Pointer(shape)).BodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+17975, __ccgo_ts+21995, int32FromInt32(1784))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	body = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b2Shape)(unsafe.Pointer(shape)).BodyId)*104
	contactKey = (*b2Body)(unsafe.Pointer(body)).HeadContactKey
	index = 0
	for contactKey != -int32(1) && index < capacity {
		contactId = contactKey >> int32(1)
		edgeIndex = contactKey & int32(1)
		if v2 = !!(0 <= contactId && contactId < (*b2World)(unsafe.Pointer(world)).Contacts.Count); !v2 {
			b2InternalAssert(tls, __ccgo_ts+301, __ccgo_ts+21995, int32FromInt32(1792))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		contact = (*b2World)(unsafe.Pointer(world)).Contacts.Data + uintptr(contactId)*64
		// Does contact involve this shape and is it touching?
		if ((*b2Contact)(unsafe.Pointer(contact)).ShapeIdA == shapeId.Index1-int32(1) || (*b2Contact)(unsafe.Pointer(contact)).ShapeIdB == shapeId.Index1-int32(1)) && (*b2Contact)(unsafe.Pointer(contact)).Flags&uint32(b2_contactTouchingFlag) != uint32(0) {
			shapeA = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b2Contact)(unsafe.Pointer(contact)).ShapeIdA)*280
			shapeB = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b2Contact)(unsafe.Pointer(contact)).ShapeIdB)*280
			(*(*ContactData)(unsafe.Add(nil, contactData+uintptr(index)*132))).ContactId = ContactId{
				Index1:     (*b2Contact)(unsafe.Pointer(contact)).ContactId + int32(1),
				World0:     shapeId.World0,
				Generation: (*b2Contact)(unsafe.Pointer(contact)).Generation,
			}
			(*(*ContactData)(unsafe.Add(nil, contactData+uintptr(index)*132))).ShapeIdA = ShapeId{
				Index1:     (*b2Shape)(unsafe.Pointer(shapeA)).Id + int32(1),
				World0:     shapeId.World0,
				Generation: (*b2Shape)(unsafe.Pointer(shapeA)).Generation,
			}
			(*(*ContactData)(unsafe.Add(nil, contactData+uintptr(index)*132))).ShapeIdB = ShapeId{
				Index1:     (*b2Shape)(unsafe.Pointer(shapeB)).Id + int32(1),
				World0:     shapeId.World0,
				Generation: (*b2Shape)(unsafe.Pointer(shapeB)).Generation,
			}
			contactSim = b2GetContactSim(tls, world, contact)
			(*(*ContactData)(unsafe.Add(nil, contactData+uintptr(index)*132))).Manifold = (*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold
			index = index + int32(1)
		}
		contactKey = (*(*b2ContactEdge)(unsafe.Add(nil, contact+uintptr(edgeIndex)*12))).NextKey
	}
	if v1 = !!(index <= capacity); !v1 {
		b2InternalAssert(tls, __ccgo_ts+2298, __ccgo_ts+21995, int32FromInt32(1813))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return index
}

func b2Shape_GetSensorCapacity(tls *_Stack, shapeId ShapeId) (r int32) {
	var sensor, shape, world uintptr
	var v1 bool
	_, _, _, _ = sensor, shape, world, v1
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return 0
	}
	shape = b2GetShape(tls, world, shapeId)
	if (*b2Shape)(unsafe.Pointer(shape)).SensorIndex == -int32(1) {
		return 0
	}
	if v1 = !!(0 <= (*b2Shape)(unsafe.Pointer(shape)).SensorIndex && (*b2Shape)(unsafe.Pointer(shape)).SensorIndex < (*b2World)(unsafe.Pointer(world)).Sensors.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+22601, __ccgo_ts+21995, int32FromInt32(1832))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	sensor = (*b2World)(unsafe.Pointer(world)).Sensors.Data + uintptr((*b2Shape)(unsafe.Pointer(shape)).SensorIndex)*56
	return (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count
}

func b2Shape_GetSensorData(tls *_Stack, shapeId ShapeId, visitorIds uintptr, capacity int32) (r int32) {
	var count, i, v2, v3, v4, v6 int32
	var refs, sensor, shape, world uintptr
	var visitorId ShapeId
	var v1 bool
	_, _, _, _, _, _, _, _, _, _, _, _ = count, i, refs, sensor, shape, visitorId, world, v1, v2, v3, v4, v6
	world = b2GetWorldLocked(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return 0
	}
	shape = b2GetShape(tls, world, shapeId)
	if (*b2Shape)(unsafe.Pointer(shape)).SensorIndex == -int32(1) {
		return 0
	}
	if v1 = !!(0 <= (*b2Shape)(unsafe.Pointer(shape)).SensorIndex && (*b2Shape)(unsafe.Pointer(shape)).SensorIndex < (*b2World)(unsafe.Pointer(world)).Sensors.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+22601, __ccgo_ts+21995, int32FromInt32(1850))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	sensor = (*b2World)(unsafe.Pointer(world)).Sensors.Data + uintptr((*b2Shape)(unsafe.Pointer(shape)).SensorIndex)*56
	v2 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count
	v3 = capacity
	if v2 < v3 {
		v6 = v2
	} else {
		v6 = v3
	}
	v4 = v6
	goto _5
_5:
	count = v4
	refs = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data
	i = 0
	for {
		if !(i < count) {
			break
		}
		visitorId = ShapeId{
			Index1:     (*(*b2Visitor)(unsafe.Add(nil, refs+uintptr(i)*8))).ShapeId + int32(1),
			World0:     shapeId.World0,
			Generation: (*(*b2Visitor)(unsafe.Add(nil, refs+uintptr(i)*8))).Generation,
		}
		*(*ShapeId)(unsafe.Add(nil, visitorIds+uintptr(i)*8)) = visitorId
		goto _7
	_7:
		;
		i = i + 1
	}
	return count
}

func b2Shape_GetBounds(tls *_Stack, shapeId ShapeId) (r AABB) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return AABB{}
	}
	shape = b2GetShape(tls, world, shapeId)
	return (*b2Shape)(unsafe.Pointer(shape)).Aabb
}

func b2Shape_ComputeMassData(tls *_Stack, shapeId ShapeId) (r MassData) {
	var shape, world uintptr
	_, _ = shape, world
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return MassData{}
	}
	shape = b2GetShape(tls, world, shapeId)
	return b2ComputeShapeMass(tls, shape)
}

func b2Shape_GetClosestPoint(tls *_Stack, shapeId ShapeId, target Pos) (r1 Pos) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var C, v12 Transform
	var body, shape, world uintptr
	var d, v10, v9 Vec2
	var output DistanceOutput
	var r, v4, v5, v6, v8 Rot
	var rx, ry float32
	var targetTransform, transform, v2, v3 WorldTransform
	var v1 bool
	var v16 Pos
	var _ /* cache at bp+176 */ SimplexCache
	var _ /* input at bp+8 */ DistanceInput
	var _ /* zero at bp+0 */ Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, body, d, output, r, rx, ry, shape, targetTransform, transform, world, v1, v10, v12, v16, v2, v3, v4, v5, v6, v8, v9
	world = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	if world == uintptrFromInt32(0) {
		return b2Pos_zero
	}
	shape = b2GetShape(tls, world, shapeId)
	if v1 = !!(0 <= (*b2Shape)(unsafe.Pointer(shape)).BodyId && (*b2Shape)(unsafe.Pointer(shape)).BodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+17975, __ccgo_ts+21995, int32FromInt32(1901))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	body = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b2Shape)(unsafe.Pointer(shape)).BodyId)*104
	transform = b2GetBodyTransformQuick(tls, world, body)
	// The target rides in as the frame of proxy B, so the relative pose is differenced in double
	// and the result stays exact far from the origin
	*(*Vec2)(unsafe.Add(nil, bp)) = b2Vec2_zero
	targetTransform = WorldTransform{
		P: target,
		Q: b2Rot_identity,
	}
	(*(*DistanceInput)(unsafe.Add(nil, bp+8))).ProxyA = b2MakeShapeDistanceProxy(tls, shape)
	(*(*DistanceInput)(unsafe.Add(nil, bp+8))).ProxyB = b2MakeProxy(tls, bp, int32(1), float32FromFloat32(0))
	v2 = transform
	v3 = targetTransform
	v4 = v2.Q
	v5 = v3.Q
	r.S = float32(v4.C*v5.S) - float32(v4.S*v5.C)
	r.C = float32(v4.C*v5.C) + float32(v4.S*v5.S)
	v6 = r
	goto _7
_7:
	C.Q = v6
	d = Vec2{
		X: v3.P.X - v2.P.X,
		Y: v3.P.Y - v2.P.Y,
	}
	v8 = v2.Q
	v9 = d
	v10 = Vec2{
		X: float32(v8.C*v9.X) + float32(v8.S*v9.Y),
		Y: float32(-v8.S*v9.X) + float32(v8.C*v9.Y),
	}
	goto _11
_11:
	C.P = v10
	v12 = C
	goto _13
_13:
	(*(*DistanceInput)(unsafe.Add(nil, bp+8))).Transform = v12
	(*(*DistanceInput)(unsafe.Add(nil, bp+8))).UseRadii = boolUint8(true1 != 0)
	*(*SimplexCache)(unsafe.Add(nil, bp+176)) = SimplexCache{}
	output = b2ComputeShapeDistance(tls, bp+8, bp+176, uintptrFromInt32(0), 0)
	v2 = transform
	v9 = output.PointA
	rx = float32(v2.Q.C*v9.X) - float32(v2.Q.S*v9.Y)
	ry = float32(v2.Q.S*v9.X) + float32(v2.Q.C*v9.Y)
	v16 = Pos{
		X: v2.P.X + rx,
		Y: v2.P.Y + ry,
	}
	goto _17
_17:
	return v16
}

// C documentation
//
//	// https://en.wikipedia.org/wiki/Density_of_air
//	// https://www.engineeringtoolbox.com/wind-load-d_1775.html
//	// force = 0.5 * air_density * velocity^2 * area
//	// https://en.wikipedia.org/wiki/Lift_(force)
func b2Shape_ApplyWind(tls *_Stack, shapeId ShapeId, wind Vec2, drag float32, lift float32, wake uint8) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var airDensity, forceMagnitude, forceMagnitude1, lengthSquared, lengthSquared1, lengthUnits, projectedArea, projectedArea1, projectedArea2, radius, radius1, s4, s5, torque, volumeUnits, v26, v35, v43, v49, v59 float32
	var body2, bodySim, set, set1, shape, sim, state, vertices, world2, v12, v5, v6, v9 uintptr
	var centroid, centroid1, centroid2, d, d1, direction, direction1, direction2, edgeCenter, edgeLever, edgeLever1, f, force, lever, lever1, lever2, liftDirection, liftDirection1, normal, normal1, relativeVelocity, relativeVelocity1, relativeVelocity2, shapeVelocity, shapeVelocity1, shapeVelocity2, u, u1, v11, v21, v18, v19, v20, v23, v24, v27, v28, v30, v31, v32, v34, v36, v37, v40, v41, v44, v45, v47, v48, v51, v52 Vec2
	var count, i int32
	var shapeType b2ShapeType1
	var transform WorldTransform
	var v1, v2 bool
	var v22, v55 Rot
	var _ /* _a at bp+0 */ b2RecArgs_ShapeApplyWind
	var _ /* speed at bp+28 */ float32
	var _ /* speed at bp+32 */ float32
	var _ /* speed at bp+36 */ float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = airDensity, body2, bodySim, centroid, centroid1, centroid2, count, d, d1, direction, direction1, direction2, edgeCenter, edgeLever, edgeLever1, f, force, forceMagnitude, forceMagnitude1, i, lengthSquared, lengthSquared1, lengthUnits, lever, lever1, lever2, liftDirection, liftDirection1, normal, normal1, projectedArea, projectedArea1, projectedArea2, radius, radius1, relativeVelocity, relativeVelocity1, relativeVelocity2, s4, s5, set, set1, shape, shapeType, shapeVelocity, shapeVelocity1, shapeVelocity2, sim, state, torque, transform, u, u1, v11, v21, vertices, volumeUnits, world2, v1, v12, v18, v19, v2, v20, v22, v23, v24, v26, v27, v28, v30, v31, v32, v34, v35, v36, v37, v40, v41, v43, v44, v45, v47, v48, v49, v5, v51, v52, v55, v59, v6, v9
	if int32FromUint8(b2IsValidVec2(tls, wind)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+48, __ccgo_ts+23855, uintptr(unsafe.Pointer(&__func__181))))
		if v1 = !!(b2IsValidVec2(tls, wind) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+23855, __ccgo_ts+21995, int32FromInt32(1927))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	if int32FromUint8(b2IsValidFloat(tls, drag)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+48, __ccgo_ts+23877, uintptr(unsafe.Pointer(&__func__181))))
		if v1 = !!(b2IsValidFloat(tls, drag) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+23877, __ccgo_ts+21995, int32FromInt32(1928))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	if int32FromUint8(b2IsValidFloat(tls, lift)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+48, __ccgo_ts+23900, uintptr(unsafe.Pointer(&__func__181))))
		if v1 = !!(b2IsValidFloat(tls, lift) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+23900, __ccgo_ts+21995, int32FromInt32(1929))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world2 = b2GetWorld(tls, int32FromUint16(shapeId.World0))
	if world2 == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world2)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_ShapeApplyWind)(unsafe.Add(nil, bp)) = b2RecArgs_ShapeApplyWind{
			Shape: shapeId,
			Wind:  wind,
			Drag:  drag,
			Lift:  lift,
			Wake:  wake,
		}
		b2RecWrite_ShapeApplyWind(tls, (*b2World)(unsafe.Pointer(world2)).Recording, bp)
	}
	shape = b2GetShape(tls, world2, shapeId)
	shapeType = (*b2Shape)(unsafe.Pointer(shape)).Type1
	if shapeType != int32(b2_circleShape) && shapeType != int32(b2_capsuleShape) && shapeType != int32(b2_polygonShape) {
		return
	}
	if v1 = !!(0 <= (*b2Shape)(unsafe.Pointer(shape)).BodyId && (*b2Shape)(unsafe.Pointer(shape)).BodyId < (*b2World)(unsafe.Pointer(world2)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+17975, __ccgo_ts+21995, int32FromInt32(1947))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	body2 = (*b2World)(unsafe.Pointer(world2)).Bodies.Data + uintptr((*b2Shape)(unsafe.Pointer(shape)).BodyId)*104
	if (*b2Body)(unsafe.Pointer(body2)).Type1 != int32(b2_dynamicBody) {
		return
	}
	if (*b2Body)(unsafe.Pointer(body2)).SetIndex >= int32(b2_firstSleepingSet) && int32FromUint8(wake) == false1 {
		return
	}
	v5 = world2
	v6 = body2
	if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(v6)).SetIndex && (*b2Body)(unsafe.Pointer(v6)).SetIndex < (*b2World)(unsafe.Pointer(v5)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(v5)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v6)).SetIndex)*88
	if v2 = !!(0 <= (*b2Body)(unsafe.Pointer(v6)).LocalIndex && (*b2Body)(unsafe.Pointer(v6)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v6)).LocalIndex)*96
	v9 = bodySim
	goto _10
_10:
	sim = v9
	if (*b2Body)(unsafe.Pointer(body2)).SetIndex != int32(b2_awakeSet) {
		// Must wake for state to exist
		b2WakeBody(tls, world2, body2)
	}
	if v1 = !!((*b2Body)(unsafe.Pointer(body2)).SetIndex == int32(b2_awakeSet)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+2683, __ccgo_ts+21995, int32FromInt32(1967))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	v5 = world2
	v6 = body2
	if (*b2Body)(unsafe.Pointer(v6)).SetIndex == int32(b2_awakeSet) {
		if v1 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(v5)).SolverSets.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+698, int32FromInt32(254))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		set1 = (*b2World)(unsafe.Pointer(v5)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
		if v2 = !!(0 <= (*b2Body)(unsafe.Pointer(v6)).LocalIndex && (*b2Body)(unsafe.Pointer(v6)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set1)).BodyStates.Count); !v2 {
			b2InternalAssert(tls, __ccgo_ts+790, __ccgo_ts+698, int32FromInt32(255))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		v9 = (*b2SolverSet)(unsafe.Pointer(set1)).BodyStates.Data + uintptr((*b2Body)(unsafe.Pointer(v6)).LocalIndex)*32
		goto _16
	}
	v9 = uintptrFromInt32(0)
	goto _16
_16:
	state = v9
	transform = (*b2BodySim)(unsafe.Pointer(sim)).Transform
	lengthUnits = b2GetLengthUnitsPerMeter(tls)
	volumeUnits = float32(float32(lengthUnits*lengthUnits) * lengthUnits)
	// In 2D I'm assuming unit depth
	airDensity = float32FromFloat32(1.225) / volumeUnits
	force = Vec2{}
	torque = float32FromFloat32(0)
	switch (*b2Shape)(unsafe.Pointer(shape)).Type1 {
	case int32(b2_circleShape):
		radius = (*(*Circle)(unsafe.Add(unsafe.Pointer(shape), 128))).Radius
		centroid = (*b2Shape)(unsafe.Pointer(shape)).LocalCentroid
		v18 = centroid
		v19 = (*b2BodySim)(unsafe.Pointer(sim)).LocalCenter
		v20 = Vec2{
			X: v18.X - v19.X,
			Y: v18.Y - v19.Y,
		}
		goto _21
	_21:
		v22 = transform.Q
		v23 = v20
		v24 = Vec2{
			X: float32(v22.C*v23.X) - float32(v22.S*v23.Y),
			Y: float32(v22.S*v23.X) + float32(v22.C*v23.Y),
		}
		goto _25
	_25:
		lever = v24
		v26 = (*b2BodyState)(unsafe.Pointer(state)).AngularVelocity
		v27 = lever
		v28 = Vec2{
			X: float32(-v26 * v27.Y),
			Y: float32(v26 * v27.X),
		}
		goto _29
	_29:
		v30 = (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
		v31 = v28
		v32 = Vec2{
			X: v30.X + v31.X,
			Y: v30.Y + v31.Y,
		}
		goto _33
	_33:
		shapeVelocity = v32
		v34 = wind
		v35 = drag
		v36 = shapeVelocity
		v37 = Vec2{
			X: v34.X - float32(v35*v36.X),
			Y: v34.Y - float32(v35*v36.Y),
		}
		goto _38
	_38:
		relativeVelocity = v37
		v12 = bp + 28
		v40 = relativeVelocity
		lengthSquared1 = float32(v40.X*v40.X) + float32(v40.Y*v40.Y)
		if lengthSquared1 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			*(*float32)(unsafe.Add(nil, v12)) = sqrtf(tls, lengthSquared1)
			s5 = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v12))
			u1 = Vec2{
				X: float32(s5 * v40.X),
				Y: float32(s5 * v40.Y),
			}
			v41 = u1
			goto _42
		}
		*(*float32)(unsafe.Add(nil, v12)) = float32FromFloat32(0)
		v41 = Vec2{}
		goto _42
	_42:
		direction = v41
		projectedArea = float32(float32FromFloat32(2) * radius)
		v43 = float32(float32(float32(float32(float32FromFloat32(0.5)*airDensity)*projectedArea)**(*float32)(unsafe.Add(nil, bp+28))) * *(*float32)(unsafe.Add(nil, bp+28)))
		v44 = direction
		v45 = Vec2{
			X: float32(v43 * v44.X),
			Y: float32(v43 * v44.Y),
		}
		goto _46
	_46:
		force = v45
		v18 = lever
		v19 = force
		v26 = float32(v18.X*v19.Y) - float32(v18.Y*v19.X)
		goto _50
	_50:
		torque = v26
	case int32(b2_capsuleShape):
		centroid1 = (*b2Shape)(unsafe.Pointer(shape)).LocalCentroid
		v18 = centroid1
		v19 = (*b2BodySim)(unsafe.Pointer(sim)).LocalCenter
		v20 = Vec2{
			X: v18.X - v19.X,
			Y: v18.Y - v19.Y,
		}
		goto _54
	_54:
		v22 = transform.Q
		v23 = v20
		v24 = Vec2{
			X: float32(v22.C*v23.X) - float32(v22.S*v23.Y),
			Y: float32(v22.S*v23.X) + float32(v22.C*v23.Y),
		}
		goto _58
	_58:
		lever1 = v24
		v26 = (*b2BodyState)(unsafe.Pointer(state)).AngularVelocity
		v27 = lever1
		v28 = Vec2{
			X: float32(-v26 * v27.Y),
			Y: float32(v26 * v27.X),
		}
		goto _62
	_62:
		v30 = (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
		v31 = v28
		v32 = Vec2{
			X: v30.X + v31.X,
			Y: v30.Y + v31.Y,
		}
		goto _66
	_66:
		shapeVelocity1 = v32
		v34 = wind
		v35 = drag
		v36 = shapeVelocity1
		v37 = Vec2{
			X: v34.X - float32(v35*v36.X),
			Y: v34.Y - float32(v35*v36.Y),
		}
		goto _71
	_71:
		relativeVelocity1 = v37
		v5 = bp + 32
		v40 = relativeVelocity1
		lengthSquared1 = float32(v40.X*v40.X) + float32(v40.Y*v40.Y)
		if lengthSquared1 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			*(*float32)(unsafe.Add(nil, v5)) = sqrtf(tls, lengthSquared1)
			s5 = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v5))
			u1 = Vec2{
				X: float32(s5 * v40.X),
				Y: float32(s5 * v40.Y),
			}
			v41 = u1
			goto _75
		}
		*(*float32)(unsafe.Add(nil, v5)) = float32FromFloat32(0)
		v41 = Vec2{}
		goto _75
	_75:
		direction1 = v41
		v44 = (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo14_128.Capsule.Center2
		v45 = (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo14_128.Capsule.Center1
		v47 = Vec2{
			X: v44.X - v45.X,
			Y: v44.Y - v45.Y,
		}
		goto _79
	_79:
		d = v47
		v55 = transform.Q
		v48 = d
		v51 = Vec2{
			X: float32(v55.C*v48.X) - float32(v55.S*v48.Y),
			Y: float32(v55.S*v48.X) + float32(v55.C*v48.Y),
		}
		goto _83
	_83:
		d = v51
		radius1 = (*b2Shape)(unsafe.Pointer(shape)).ｆ__ccgo14_128.Capsule.Radius
		v18 = d
		v19 = direction1
		v26 = float32(v18.X*v19.Y) - float32(v18.Y*v19.X)
		goto _87
	_87:
		v35 = v26
		if v35 < float32FromInt32(0) {
			v49 = -v35
		} else {
			v49 = v35
		}
		v43 = v49
		goto _90
	_90:
		projectedArea1 = float32(float32FromFloat32(2)*radius1) + v43
		v20 = d
		lengthSquared = float32(v20.X*v20.X) + float32(v20.Y*v20.Y)
		if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			s4 = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
			u = Vec2{
				X: float32(s4 * v20.X),
				Y: float32(s4 * v20.Y),
			}
			v23 = u
			goto _94
		}
		v23 = Vec2{}
		goto _94
	_94:
		v24 = v23
		v27 = Vec2{
			X: -v24.Y,
			Y: v24.X,
		}
		goto _97
	_97:
		// Normal that opposes the wind
		normal = v27
		v30 = normal
		v31 = direction1
		v59 = float32(v30.X*v31.X) + float32(v30.Y*v31.Y)
		goto _102
	_102:
		if v59 > float32FromFloat32(0) {
			v32 = normal
			v34 = Vec2{
				X: -v32.X,
				Y: -v32.Y,
			}
			goto _105
		_105:
			v28 = v34
		} else {
			v28 = normal
		}
		normal = v28
		v18 = normal
		v19 = direction1
		v26 = float32(v18.X*v19.Y) - float32(v18.Y*v19.X)
		goto _109
	_109:
		v35 = v26
		v20 = direction1
		v23 = Vec2{
			X: float32(-v35 * v20.Y),
			Y: float32(v35 * v20.X),
		}
		goto _113
	_113:
		// portion of wind that is perpendicular to surface
		liftDirection = v23
		forceMagnitude = float32(float32(float32(float32(float32FromFloat32(0.5)*airDensity)*projectedArea1)**(*float32)(unsafe.Add(nil, bp+32))) * *(*float32)(unsafe.Add(nil, bp+32)))
		v24 = direction1
		v43 = lift
		v27 = liftDirection
		v28 = Vec2{
			X: v24.X + float32(v43*v27.X),
			Y: v24.Y + float32(v43*v27.Y),
		}
		goto _118
	_118:
		v49 = forceMagnitude
		v30 = v28
		v31 = Vec2{
			X: float32(v49 * v30.X),
			Y: float32(v49 * v30.Y),
		}
		goto _122
	_122:
		force = v31
		v18 = lever1
		v26 = radius1
		v19 = normal
		v20 = Vec2{
			X: v18.X + float32(v26*v19.X),
			Y: v18.Y + float32(v26*v19.Y),
		}
		goto _127
	_127:
		edgeLever = v20
		v23 = edgeLever
		v24 = force
		v35 = float32(v23.X*v24.Y) - float32(v23.Y*v24.X)
		goto _131
	_131:
		torque = v35
	case int32(b2_polygonShape):
		centroid2 = (*b2Shape)(unsafe.Pointer(shape)).LocalCentroid
		v18 = centroid2
		v19 = (*b2BodySim)(unsafe.Pointer(sim)).LocalCenter
		v20 = Vec2{
			X: v18.X - v19.X,
			Y: v18.Y - v19.Y,
		}
		goto _135
	_135:
		v22 = transform.Q
		v23 = v20
		v24 = Vec2{
			X: float32(v22.C*v23.X) - float32(v22.S*v23.Y),
			Y: float32(v22.S*v23.X) + float32(v22.C*v23.Y),
		}
		goto _139
	_139:
		lever2 = v24
		v26 = (*b2BodyState)(unsafe.Pointer(state)).AngularVelocity
		v27 = lever2
		v28 = Vec2{
			X: float32(-v26 * v27.Y),
			Y: float32(v26 * v27.X),
		}
		goto _143
	_143:
		v30 = (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
		v31 = v28
		v32 = Vec2{
			X: v30.X + v31.X,
			Y: v30.Y + v31.Y,
		}
		goto _147
	_147:
		shapeVelocity2 = v32
		v34 = wind
		v35 = drag
		v36 = shapeVelocity2
		v37 = Vec2{
			X: v34.X - float32(v35*v36.X),
			Y: v34.Y - float32(v35*v36.Y),
		}
		goto _152
	_152:
		relativeVelocity2 = v37
		v5 = bp + 36
		v40 = relativeVelocity2
		lengthSquared1 = float32(v40.X*v40.X) + float32(v40.Y*v40.Y)
		if lengthSquared1 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			*(*float32)(unsafe.Add(nil, v5)) = sqrtf(tls, lengthSquared1)
			s5 = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v5))
			u1 = Vec2{
				X: float32(s5 * v40.X),
				Y: float32(s5 * v40.Y),
			}
			v41 = u1
			goto _156
		}
		*(*float32)(unsafe.Add(nil, v5)) = float32FromFloat32(0)
		v41 = Vec2{}
		goto _156
	_156:
		direction2 = v41
		// polygon radius is ignored for simplicity
		count = (*(*Polygon)(unsafe.Add(unsafe.Pointer(shape), 128))).Count
		vertices = shape + 128
		v11 = *(*Vec2)(unsafe.Add(nil, vertices+uintptr(count-int32(1))*8))
		i = 0
		for {
			if !(i < count) {
				break
			}
			v21 = *(*Vec2)(unsafe.Add(nil, vertices+uintptr(i)*8))
			v44 = v21
			v45 = v11
			v47 = Vec2{
				X: v44.X - v45.X,
				Y: v44.Y - v45.Y,
			}
			goto _161
		_161:
			d1 = v47
			v48 = v11
			v51 = v21
			v43 = float32FromFloat32(0.5)
			v52 = Vec2{
				X: float32((float32FromFloat32(1)-v43)*v48.X) + float32(v43*v51.X),
				Y: float32((float32FromFloat32(1)-v43)*v48.Y) + float32(v43*v51.Y),
			}
			goto _166
		_166:
			edgeCenter = v52
			v11 = v21
			v22 = transform.Q
			v18 = d1
			v19 = Vec2{
				X: float32(v22.C*v18.X) - float32(v22.S*v18.Y),
				Y: float32(v22.S*v18.X) + float32(v22.C*v18.Y),
			}
			goto _170
		_170:
			d1 = v19
			v18 = d1
			v19 = direction2
			v26 = float32(v18.X*v19.Y) - float32(v18.Y*v19.X)
			goto _174
		_174:
			projectedArea2 = v26
			if projectedArea2 <= float32FromFloat32(0) {
				// back facing
				goto _157
			}
			v18 = d1
			lengthSquared = float32(v18.X*v18.X) + float32(v18.Y*v18.Y)
			if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
				s4 = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
				u = Vec2{
					X: float32(s4 * v18.X),
					Y: float32(s4 * v18.Y),
				}
				v19 = u
				goto _177
			}
			v19 = Vec2{}
			goto _177
		_177:
			v20 = v19
			v23 = Vec2{
				X: v20.Y,
				Y: -v20.X,
			}
			goto _180
		_180:
			normal1 = v23
			v24 = normal1
			v27 = direction2
			v26 = float32(v24.X*v27.Y) - float32(v24.Y*v27.X)
			goto _184
		_184:
			v35 = v26
			v28 = direction2
			v30 = Vec2{
				X: float32(-v35 * v28.Y),
				Y: float32(v35 * v28.X),
			}
			goto _188
		_188:
			// portion of wind that is perpendicular to surface
			liftDirection1 = v30
			forceMagnitude1 = float32(float32(float32(float32(float32FromFloat32(0.5)*airDensity)*projectedArea2)**(*float32)(unsafe.Add(nil, bp+36))) * *(*float32)(unsafe.Add(nil, bp+36)))
			v31 = direction2
			v43 = lift
			v32 = liftDirection1
			v34 = Vec2{
				X: v31.X + float32(v43*v32.X),
				Y: v31.Y + float32(v43*v32.Y),
			}
			goto _193
		_193:
			v49 = forceMagnitude1
			v36 = v34
			v37 = Vec2{
				X: float32(v49 * v36.X),
				Y: float32(v49 * v36.Y),
			}
			goto _197
		_197:
			f = v37
			v40 = edgeCenter
			v41 = (*b2BodySim)(unsafe.Pointer(sim)).LocalCenter
			v44 = Vec2{
				X: v40.X - v41.X,
				Y: v40.Y - v41.Y,
			}
			goto _201
		_201:
			v22 = transform.Q
			v45 = v44
			v47 = Vec2{
				X: float32(v22.C*v45.X) - float32(v22.S*v45.Y),
				Y: float32(v22.S*v45.X) + float32(v22.C*v45.Y),
			}
			goto _205
		_205:
			edgeLever1 = v47
			v48 = force
			v51 = f
			v52 = Vec2{
				X: v48.X + v51.X,
				Y: v48.Y + v51.Y,
			}
			goto _209
		_209:
			force = v52
			v18 = edgeLever1
			v19 = f
			v26 = float32(v18.X*v19.Y) - float32(v18.Y*v19.X)
			goto _213
		_213:
			torque = torque + v26
			goto _157
		_157:
			;
			i = i + 1
		}
	default:
		break
	}
	v18 = (*b2BodySim)(unsafe.Pointer(sim)).Force
	v19 = force
	v20 = Vec2{
		X: v18.X + v19.X,
		Y: v18.Y + v19.Y,
	}
	goto _217
_217:
	(*b2BodySim)(unsafe.Pointer(sim)).Force = v20
	*(*float32)(unsafe.Add(nil, sim+56)) += torque
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
