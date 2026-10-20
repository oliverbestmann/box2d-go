package b2

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

// C documentation
//
//	/// Follows Ericson 5.1.9 Closest Points of Two Line Segments
func b2ComputeSegmentDistance(tls *_Stack, p1 Vec2, q1 Vec2, p2 Vec2, q2 Vec2) (r1 SegmentDistanceResult) {
	var c, d1, d2, r, v1, v10, v11, v13, v14, v17, v18, v2, v21, v22, v25, v26, v3, v5, v6, v7, v9 Vec2
	var d12, dd1, dd2, denominator, epsSqr, f1, f2, rd1, rd2, v15, v19, v23, v27, v29, v30, v31, v32, v34, v35 float32
	var result SegmentDistanceResult
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, d1, d12, d2, dd1, dd2, denominator, epsSqr, f1, f2, r, rd1, rd2, result, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v29, v3, v30, v31, v32, v34, v35, v5, v6, v7, v9
	result = SegmentDistanceResult{}
	v1 = q1
	v2 = p1
	v3 = Vec2{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
	goto _4
_4:
	d1 = v3
	v5 = q2
	v6 = p2
	v7 = Vec2{
		X: v5.X - v6.X,
		Y: v5.Y - v6.Y,
	}
	goto _8
_8:
	d2 = v7
	v9 = p1
	v10 = p2
	v11 = Vec2{
		X: v9.X - v10.X,
		Y: v9.Y - v10.Y,
	}
	goto _12
_12:
	r = v11
	v13 = d1
	v14 = d1
	v15 = float32(v13.X*v14.X) + float32(v13.Y*v14.Y)
	goto _16
_16:
	dd1 = v15
	v17 = d2
	v18 = d2
	v19 = float32(v17.X*v18.X) + float32(v17.Y*v18.Y)
	goto _20
_20:
	dd2 = v19
	v21 = r
	v22 = d1
	v23 = float32(v21.X*v22.X) + float32(v21.Y*v22.Y)
	goto _24
_24:
	rd1 = v23
	v25 = r
	v26 = d2
	v27 = float32(v25.X*v26.X) + float32(v25.Y*v26.Y)
	goto _28
_28:
	rd2 = v27
	epsSqr = float32(float32FromFloat32(1.1920928955078125e-07) * float32FromFloat32(1.1920928955078125e-07))
	if dd1 < epsSqr || dd2 < epsSqr {
		// Handle all degeneracies
		if dd1 >= epsSqr {
			// Segment 2 is degenerate
			v29 = -rd1 / dd1
			v30 = float32FromFloat32(0)
			v31 = float32FromFloat32(1)
			if v29 < v30 {
				v34 = v30
			} else {
				if v29 > v31 {
					v35 = v31
				} else {
					v35 = v29
				}
				v34 = v35
			}
			v32 = v34
			goto _33
		_33:
			result.Fraction1 = v32
			result.Fraction2 = float32FromFloat32(0)
		} else {
			if dd2 >= epsSqr {
				// Segment 1 is degenerate
				result.Fraction1 = float32FromFloat32(0)
				v15 = rd2 / dd2
				v19 = float32FromFloat32(0)
				v23 = float32FromFloat32(1)
				if v15 < v19 {
					v29 = v19
				} else {
					if v15 > v23 {
						v30 = v23
					} else {
						v30 = v15
					}
					v29 = v30
				}
				v27 = v29
				goto _40
			_40:
				result.Fraction2 = v27
			} else {
				result.Fraction1 = float32FromFloat32(0)
				result.Fraction2 = float32FromFloat32(0)
			}
		}
	} else {
		v1 = d1
		v2 = d2
		v15 = float32(v1.X*v2.X) + float32(v1.Y*v2.Y)
		goto _46
	_46:
		// Non-degenerate segments
		d12 = v15
		denominator = float32(dd1*dd2) - float32(d12*d12)
		// Fraction on segment 1
		f1 = float32FromFloat32(0)
		if denominator != float32FromFloat32(0) {
			// not parallel
			v19 = (float32(d12*rd2) - float32(rd1*dd2)) / denominator
			v23 = float32FromFloat32(0)
			v27 = float32FromFloat32(1)
			if v19 < v23 {
				v30 = v23
			} else {
				if v19 > v27 {
					v31 = v27
				} else {
					v31 = v19
				}
				v30 = v31
			}
			v29 = v30
			goto _51
		_51:
			f1 = v29
		}
		// Compute point on segment 2 closest to p1 + f1 * d1
		f2 = (float32(d12*f1) + rd2) / dd2
		// Clamping of segment 2 requires a do over on segment 1
		if f2 < float32FromFloat32(0) {
			f2 = float32FromFloat32(0)
			v15 = -rd1 / dd1
			v19 = float32FromFloat32(0)
			v23 = float32FromFloat32(1)
			if v15 < v19 {
				v29 = v19
			} else {
				if v15 > v23 {
					v30 = v23
				} else {
					v30 = v15
				}
				v29 = v30
			}
			v27 = v29
			goto _58
		_58:
			f1 = v27
		} else {
			if f2 > float32FromFloat32(1) {
				f2 = float32FromFloat32(1)
				v15 = (d12 - rd1) / dd1
				v19 = float32FromFloat32(0)
				v23 = float32FromFloat32(1)
				if v15 < v19 {
					v29 = v19
				} else {
					if v15 > v23 {
						v30 = v23
					} else {
						v30 = v15
					}
					v29 = v30
				}
				v27 = v29
				goto _65
			_65:
				f1 = v27
			}
		}
		result.Fraction1 = f1
		result.Fraction2 = f2
	}
	v1 = p1
	v15 = result.Fraction1
	v2 = d1
	v3 = Vec2{
		X: v1.X + float32(v15*v2.X),
		Y: v1.Y + float32(v15*v2.Y),
	}
	goto _72
_72:
	result.Closest1 = v3
	v1 = p2
	v15 = result.Fraction2
	v2 = d2
	v3 = Vec2{
		X: v1.X + float32(v15*v2.X),
		Y: v1.Y + float32(v15*v2.Y),
	}
	goto _77
_77:
	result.Closest2 = v3
	v1 = result.Closest1
	v2 = result.Closest2
	c = Vec2{
		X: v2.X - v1.X,
		Y: v2.Y - v1.Y,
	}
	v15 = float32(c.X*c.X) + float32(c.Y*c.Y)
	goto _81
_81:
	result.DistanceSquared = v15
	return result
}

func b2ComputeWitnessPoints(tls *_Stack, s uintptr, a uintptr, b uintptr) {
	var v1 bool
	_ = v1
	switch (*Simplex)(unsafe.Pointer(s)).Count {
	case int32(1):
		*(*Vec2)(unsafe.Add(nil, a)) = (*Simplex)(unsafe.Pointer(s)).V1.WA
		*(*Vec2)(unsafe.Add(nil, b)) = (*Simplex)(unsafe.Pointer(s)).V1.WB
	case int32(2):
		*(*Vec2)(unsafe.Add(nil, a)) = b2Weight2(tls, (*Simplex)(unsafe.Pointer(s)).V1.A, (*Simplex)(unsafe.Pointer(s)).V1.WA, (*Simplex)(unsafe.Pointer(s)).V2.A, (*Simplex)(unsafe.Pointer(s)).V2.WA)
		*(*Vec2)(unsafe.Add(nil, b)) = b2Weight2(tls, (*Simplex)(unsafe.Pointer(s)).V1.A, (*Simplex)(unsafe.Pointer(s)).V1.WB, (*Simplex)(unsafe.Pointer(s)).V2.A, (*Simplex)(unsafe.Pointer(s)).V2.WB)
	case int32(3):
		*(*Vec2)(unsafe.Add(nil, a)) = b2Weight3(tls, (*Simplex)(unsafe.Pointer(s)).V1.A, (*Simplex)(unsafe.Pointer(s)).V1.WA, (*Simplex)(unsafe.Pointer(s)).V2.A, (*Simplex)(unsafe.Pointer(s)).V2.WA, (*Simplex)(unsafe.Pointer(s)).V3.A, (*Simplex)(unsafe.Pointer(s)).V3.WA)
		// todo why are these not equal?
		//*b = b2Weight3(s->v1.a, s->v1.wB, s->v2.a, s->v2.wB, s->v3.a, s->v3.wB);
		*(*Vec2)(unsafe.Add(nil, b)) = *(*Vec2)(unsafe.Add(nil, a))
	default:
		*(*Vec2)(unsafe.Add(nil, a)) = b2Vec2_zero
		*(*Vec2)(unsafe.Add(nil, b)) = b2Vec2_zero
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+7715, __ccgo_ts+7689, int32FromInt32(244))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		break
	}
}

// C documentation
//
//	// CCD via the local separating axis method. This seeks progression
//	// by computing the largest time at which separation is maintained.
func b2ComputeTimeOfImpact(tls *_Stack, input uintptr) (r1 TOIOutput) {
	bp := tls.Alloc(384)
	defer tls.Free(384)
	var C, xfA, xfB, v24, v25, v38 Transform
	var a11, a21, qq, s1, s11, s2, t11, t2, t21, tMax, target, tolerance, totalRadius, x, y, v17, v18, v19, v21 float32
	var distanceIterations, k_maxIterations, pushBackIterations, rootIterationCount int32
	var distanceOutput DistanceOutput
	var done, v2, v5 uint8
	var output TOIOutput
	var pA, pA1, pA2, pB, pB1, pB2, worldNormal, worldPointA, worldPointB, v30, v31, v32, v35, v36, v41, v42, v45, v46 Vec2
	var proxyA, proxyB uintptr
	var r, v1, v12, v4, v9 Rot
	var v7, v8 bool
	var _ /* cache at bp+80 */ SimplexCache
	var _ /* distanceInput at bp+88 */ DistanceInput
	var _ /* fcn at bp+256 */ b2SeparationFunction
	var _ /* indexA at bp+376 */ int32
	var _ /* indexB at bp+380 */ int32
	var _ /* sweepA at bp+0 */ Sweep
	var _ /* sweepB at bp+40 */ Sweep
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, a11, a21, distanceIterations, distanceOutput, done, k_maxIterations, output, pA, pA1, pA2, pB, pB1, pB2, proxyA, proxyB, pushBackIterations, qq, r, rootIterationCount, s1, s11, s2, t11, t2, t21, tMax, target, tolerance, totalRadius, worldNormal, worldPointA, worldPointB, x, xfA, xfB, y, v1, v12, v17, v18, v19, v2, v21, v24, v25, v30, v31, v32, v35, v36, v38, v4, v41, v42, v45, v46, v5, v7, v8, v9
	output = TOIOutput{}
	output.State = int32(b2_toiStateUnknown)
	output.Fraction = (*TOIInput)(unsafe.Pointer(input)).MaxFraction
	*(*Sweep)(unsafe.Add(nil, bp)) = (*TOIInput)(unsafe.Pointer(input)).SweepA
	*(*Sweep)(unsafe.Add(nil, bp+40)) = (*TOIInput)(unsafe.Pointer(input)).SweepB
	v1 = (*(*Sweep)(unsafe.Add(nil, bp))).Q1
	qq = float32(v1.S*v1.S) + float32(v1.C*v1.C)
	v2 = boolUint8(float32FromFloat32(1)-float32FromFloat32(0.0006) < qq && qq < float32FromFloat32(1)+float32FromFloat32(0.0006))
	goto _3
_3:
	;
	if v7 = v2 != 0; v7 {
		v4 = (*(*Sweep)(unsafe.Add(nil, bp))).Q2
		qq = float32(v4.S*v4.S) + float32(v4.C*v4.C)
		v5 = boolUint8(float32FromFloat32(1)-float32FromFloat32(0.0006) < qq && qq < float32FromFloat32(1)+float32FromFloat32(0.0006))
		goto _6
	_6:
	}
	if v8 = !!(v7 && v5 != 0); !v8 {
		b2InternalAssert(tls, __ccgo_ts+8042, __ccgo_ts+7689, int32FromInt32(1161))
	}
	_ = v8 || bool(int32FromInt32(0) != 0)
	v1 = (*(*Sweep)(unsafe.Add(nil, bp+40))).Q1
	qq = float32(v1.S*v1.S) + float32(v1.C*v1.C)
	v2 = boolUint8(float32FromFloat32(1)-float32FromFloat32(0.0006) < qq && qq < float32FromFloat32(1)+float32FromFloat32(0.0006))
	goto _11
_11:
	;
	if v7 = v2 != 0; v7 {
		v4 = (*(*Sweep)(unsafe.Add(nil, bp+40))).Q2
		qq = float32(v4.S*v4.S) + float32(v4.C*v4.C)
		v5 = boolUint8(float32FromFloat32(1)-float32FromFloat32(0.0006) < qq && qq < float32FromFloat32(1)+float32FromFloat32(0.0006))
		goto _14
	_14:
	}
	if v8 = !!(v7 && v5 != 0); !v8 {
		b2InternalAssert(tls, __ccgo_ts+8107, __ccgo_ts+7689, int32FromInt32(1162))
	}
	_ = v8 || bool(int32FromInt32(0) != 0)
	// todo_erin
	// c1 can be at the origin yet the points are far away
	// b2Vec2 origin = b2Add(sweepA.c1, input->proxyA.points[0]);
	proxyA = input
	proxyB = input + 72
	tMax = (*TOIInput)(unsafe.Pointer(input)).MaxFraction
	// Setup target distance and tolerance
	totalRadius = (*ShapeProxy)(unsafe.Pointer(proxyA)).Radius + (*ShapeProxy)(unsafe.Pointer(proxyB)).Radius
	v17 = float32(float32FromFloat32(0.005) * b2_lengthUnitsPerMeter)
	v18 = totalRadius - float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)
	if v17 > v18 {
		v21 = v17
	} else {
		v21 = v18
	}
	v19 = v21
	goto _20
_20:
	target = v19
	tolerance = float32(float32FromFloat32(0.25) * float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter))
	if v7 = !!(target > tolerance); !v7 {
		b2InternalAssert(tls, __ccgo_ts+7855, __ccgo_ts+7689, int32FromInt32(1177))
	}
	_ = v7 || bool(int32FromInt32(0) != 0)
	t11 = float32FromFloat32(0)
	k_maxIterations = int32(20)
	distanceIterations = 0
	// Prepare input for distance query.
	*(*SimplexCache)(unsafe.Add(nil, bp+80)) = SimplexCache{}
	*(*DistanceInput)(unsafe.Add(nil, bp+88)) = DistanceInput{}
	(*(*DistanceInput)(unsafe.Add(nil, bp+88))).ProxyA = (*TOIInput)(unsafe.Pointer(input)).ProxyA
	(*(*DistanceInput)(unsafe.Add(nil, bp+88))).ProxyB = (*TOIInput)(unsafe.Pointer(input)).ProxyB
	(*(*DistanceInput)(unsafe.Add(nil, bp+88))).UseRadii = boolUint8(false1 != 0)
	// The outer loop progressively attempts to compute new separating axes.
	// This loop terminates when an axis is repeated (no progress is made).
	for {
		// Get the distance between shapes. We can also use the results to get a separating axis.
		xfA = b2GetSweepTransform(tls, bp, t11)
		xfB = b2GetSweepTransform(tls, bp+40, t11)
		v24 = xfA
		v25 = xfB
		v1 = v24.Q
		v4 = v25.Q
		r.S = float32(v1.C*v4.S) - float32(v1.S*v4.C)
		r.C = float32(v1.C*v4.C) + float32(v1.S*v4.S)
		v9 = r
		goto _29
	_29:
		C.Q = v9
		v30 = v25.P
		v31 = v24.P
		v32 = Vec2{
			X: v30.X - v31.X,
			Y: v30.Y - v31.Y,
		}
		goto _33
	_33:
		v12 = v24.Q
		v35 = v32
		v36 = Vec2{
			X: float32(v12.C*v35.X) + float32(v12.S*v35.Y),
			Y: float32(-v12.S*v35.X) + float32(v12.C*v35.Y),
		}
		goto _37
	_37:
		C.P = v36
		v38 = C
		goto _39
	_39:
		(*(*DistanceInput)(unsafe.Add(nil, bp+88))).Transform = v38
		distanceOutput = b2ComputeShapeDistance(tls, bp+88, bp+80, uintptrFromInt32(0), 0)
		v1 = xfA.Q
		v30 = distanceOutput.Normal
		v31 = Vec2{
			X: float32(v1.C*v30.X) - float32(v1.S*v30.Y),
			Y: float32(v1.S*v30.X) + float32(v1.C*v30.Y),
		}
		goto _43
	_43:
		// The distance query runs in frame A, project the witness data back to world
		worldNormal = v31
		v24 = xfA
		v32 = distanceOutput.PointA
		x = float32(v24.Q.C*v32.X) - float32(v24.Q.S*v32.Y) + v24.P.X
		y = float32(v24.Q.S*v32.X) + float32(v24.Q.C*v32.Y) + v24.P.Y
		v35 = Vec2{
			X: x,
			Y: y,
		}
		goto _47
	_47:
		worldPointA = v35
		v25 = xfA
		v36 = distanceOutput.PointB
		x = float32(v25.Q.C*v36.X) - float32(v25.Q.S*v36.Y) + v25.P.X
		y = float32(v25.Q.S*v36.X) + float32(v25.Q.C*v36.Y) + v25.P.Y
		v41 = Vec2{
			X: x,
			Y: y,
		}
		goto _51
	_51:
		worldPointB = v41
		// Progressive time of impact. This handles slender geometry well but introduces
		// significant time loss.
		// if (distanceIterations == 0)
		//{
		//	if ( distanceOutput.distance > totalRadius + B2_SPECULATIVE_DISTANCE )
		//	{
		//		target = totalRadius + B2_SPECULATIVE_DISTANCE - tolerance;
		//	}
		//	else
		//	{
		//		target = distanceOutput.distance - 1.5f * tolerance;
		//		target = b2MaxFloat( target, 2.0f * tolerance );
		//	}
		//}
		distanceIterations = distanceIterations + int32(1)
		// If the shapes are overlapped, we give up on continuous collision.
		if distanceOutput.Distance <= float32FromFloat32(0) {
			// Failure!
			output.State = int32(b2_toiStateOverlapped)
			output.Fraction = float32FromFloat32(0)
			break
		}
		if distanceOutput.Distance <= target+tolerance {
			// Success!
			output.State = int32(b2_toiStateHit)
			v30 = worldPointA
			v17 = (*ShapeProxy)(unsafe.Pointer(proxyA)).Radius
			v31 = worldNormal
			v32 = Vec2{
				X: v30.X + float32(v17*v31.X),
				Y: v30.Y + float32(v17*v31.Y),
			}
			goto _56
		_56:
			// Averaged hit point
			pA = v32
			v35 = worldPointB
			v18 = -(*ShapeProxy)(unsafe.Pointer(proxyB)).Radius
			v36 = worldNormal
			v41 = Vec2{
				X: v35.X + float32(v18*v36.X),
				Y: v35.Y + float32(v18*v36.Y),
			}
			goto _61
		_61:
			pB = v41
			v42 = pA
			v45 = pB
			v19 = float32FromFloat32(0.5)
			v46 = Vec2{
				X: float32((float32FromFloat32(1)-v19)*v42.X) + float32(v19*v45.X),
				Y: float32((float32FromFloat32(1)-v19)*v42.Y) + float32(v19*v45.Y),
			}
			goto _66
		_66:
			output.Point = v46
			output.Normal = worldNormal
			output.Fraction = t11
			break
		}
		// Initialize the separating axis.
		*(*b2SeparationFunction)(unsafe.Add(nil, bp+256)) = b2MakeSeparationFunction(tls, *(*SimplexCache)(unsafe.Add(nil, bp+80)), proxyA, bp, proxyB, bp+40, t11)
		// Compute the TOI on the separating axis. We do this by successively
		// resolving the deepest point. This loop is bounded by the number of vertices.
		done = boolUint8(false1 != 0)
		t21 = tMax
		pushBackIterations = 0
		for {
			s2 = b2FindMinSeparation(tls, bp+256, bp+376, bp+380, t21)
			// Is the final configuration separated?
			if s2 > target+tolerance {
				// Victory!
				output.State = int32(b2_toiStateSeparated)
				output.Fraction = tMax
				done = boolUint8(true1 != 0)
				break
			}
			// Has the separation reached tolerance?
			if s2 > target-tolerance {
				// Advance the sweeps
				t11 = t21
				break
			}
			// Compute the initial separation of the witness points.
			s11 = b2EvaluateSeparation(tls, bp+256, *(*int32)(unsafe.Add(nil, bp+376)), *(*int32)(unsafe.Add(nil, bp+380)), t11)
			// Check for initial overlap. This might happen if the root finder runs out of iterations.
			if s11 < target-tolerance {
				output.State = int32(b2_toiStateFailed)
				output.Fraction = t11
				done = boolUint8(true1 != 0)
				break
			}
			// Check for touching
			if s11 <= target+tolerance {
				// Success! t1 should hold the TOI (could be 0.0).
				output.State = int32(b2_toiStateHit)
				v30 = worldPointA
				v17 = (*ShapeProxy)(unsafe.Pointer(proxyA)).Radius
				v31 = worldNormal
				v32 = Vec2{
					X: v30.X + float32(v17*v31.X),
					Y: v30.Y + float32(v17*v31.Y),
				}
				goto _72
			_72:
				// Averaged hit point
				pA1 = v32
				v35 = worldPointB
				v18 = -(*ShapeProxy)(unsafe.Pointer(proxyB)).Radius
				v36 = worldNormal
				v41 = Vec2{
					X: v35.X + float32(v18*v36.X),
					Y: v35.Y + float32(v18*v36.Y),
				}
				goto _77
			_77:
				pB1 = v41
				v42 = pA1
				v45 = pB1
				v19 = float32FromFloat32(0.5)
				v46 = Vec2{
					X: float32((float32FromFloat32(1)-v19)*v42.X) + float32(v19*v45.X),
					Y: float32((float32FromFloat32(1)-v19)*v42.Y) + float32(v19*v45.Y),
				}
				goto _82
			_82:
				output.Point = v46
				output.Normal = worldNormal
				output.Fraction = t11
				done = boolUint8(true1 != 0)
				break
			}
			// Compute 1D root of: f(x) - target = 0
			rootIterationCount = 0
			a11 = t11
			a21 = t21
			for {
				if rootIterationCount&int32(1) != 0 {
					// False position to improve convergence.
					t2 = a11 + float32((target-s11)*(a21-a11))/(s2-s11)
				} else {
					// Bisection to guarantee progress.
					t2 = float32(float32FromFloat32(0.5) * (a11 + a21))
				}
				rootIterationCount = rootIterationCount + int32(1)
				s1 = b2EvaluateSeparation(tls, bp+256, *(*int32)(unsafe.Add(nil, bp+376)), *(*int32)(unsafe.Add(nil, bp+380)), t2)
				// Has the separation reached tolerance?
				v17 = s1 - target
				if v17 < float32FromInt32(0) {
					v19 = -v17
				} else {
					v19 = v17
				}
				v18 = v19
				goto _86
			_86:
				if v18 < tolerance {
					// t2 holds a tentative value for t1
					t21 = t2
					break
				}
				// Ensure we continue to bracket the root.
				if s1 > target {
					a11 = t2
					s11 = s1
				} else {
					a21 = t2
					s2 = s1
				}
				if rootIterationCount == int32(50) {
					break
				}
				goto _83
			_83:
			}
			pushBackIterations = pushBackIterations + int32(1)
			if pushBackIterations == int32(_B2_MAX_POLYGON_VERTICES) {
				break
			}
			goto _67
		_67:
		}
		if done != 0 {
			break
		}
		if distanceIterations == k_maxIterations {
			// Root finder got stuck. Semi-victory.
			output.State = int32(b2_toiStateFailed)
			v30 = worldPointA
			v17 = (*ShapeProxy)(unsafe.Pointer(proxyA)).Radius
			v31 = worldNormal
			v32 = Vec2{
				X: v30.X + float32(v17*v31.X),
				Y: v30.Y + float32(v17*v31.Y),
			}
			goto _92
		_92:
			// Averaged hit point
			pA2 = v32
			v35 = worldPointB
			v18 = -(*ShapeProxy)(unsafe.Pointer(proxyB)).Radius
			v36 = worldNormal
			v41 = Vec2{
				X: v35.X + float32(v18*v36.X),
				Y: v35.Y + float32(v18*v36.Y),
			}
			goto _97
		_97:
			pB2 = v41
			v42 = pA2
			v45 = pB2
			v19 = float32FromFloat32(0.5)
			v46 = Vec2{
				X: float32((float32FromFloat32(1)-v19)*v42.X) + float32(v19*v45.X),
				Y: float32((float32FromFloat32(1)-v19)*v42.Y) + float32(v19*v45.Y),
			}
			goto _102
		_102:
			output.Point = v46
			output.Normal = worldNormal
			output.Fraction = t11
			break
		}
		goto _23
	_23:
	}
	return output
}

func b2ComputePolygonCentroid(tls *_Stack, vertices uintptr, count int32) (r Vec2) {
	var a4, area, inv3, invArea, v12, v19 float32
	var center, e1, e2, origin, v10, v11, v14, v15, v16, v18, v2, v20, v21, v3, v4, v6, v7, v8 Vec2
	var i int32
	var v23 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a4, area, center, e1, e2, i, inv3, invArea, origin, v10, v11, v12, v14, v15, v16, v18, v19, v2, v20, v21, v23, v3, v4, v6, v7, v8
	center = Vec2{}
	area = float32FromFloat32(0)
	// Get a reference point for forming triangles.
	// Use the first vertex to reduce round-off errors.
	origin = *(*Vec2)(unsafe.Add(nil, vertices))
	inv3 = float32FromFloat32(1) / float32FromFloat32(3)
	i = int32(1)
	for {
		if !(i < count-int32(1)) {
			break
		}
		v2 = *(*Vec2)(unsafe.Add(nil, vertices+uintptr(i)*8))
		v3 = origin
		v4 = Vec2{
			X: v2.X - v3.X,
			Y: v2.Y - v3.Y,
		}
		goto _5
	_5:
		// Triangle edges
		e1 = v4
		v6 = *(*Vec2)(unsafe.Add(nil, vertices+uintptr(i+int32(1))*8))
		v7 = origin
		v8 = Vec2{
			X: v6.X - v7.X,
			Y: v6.Y - v7.Y,
		}
		goto _9
	_9:
		e2 = v8
		v10 = e1
		v11 = e2
		v12 = float32(v10.X*v11.Y) - float32(v10.Y*v11.X)
		goto _13
	_13:
		a4 = float32(float32FromFloat32(0.5) * v12)
		// Area weighted centroid
		v14 = e1
		v15 = e2
		v16 = Vec2{
			X: v14.X + v15.X,
			Y: v14.Y + v15.Y,
		}
		goto _17
	_17:
		v18 = center
		v19 = float32(a4 * inv3)
		v20 = v16
		v21 = Vec2{
			X: v18.X + float32(v19*v20.X),
			Y: v18.Y + float32(v19*v20.Y),
		}
		goto _22
	_22:
		center = v21
		area = area + a4
		goto _1
	_1:
		;
		i = i + 1
	}
	if v23 = !!(area > float32FromFloat32(1.1920928955078125e-07)); !v23 {
		b2InternalAssert(tls, __ccgo_ts+9575, __ccgo_ts+9594, int32FromInt32(46))
	}
	_ = v23 || bool(int32FromInt32(0) != 0)
	invArea = float32FromFloat32(1) / area
	center.X = float32(center.X * invArea)
	center.Y = float32(center.Y * invArea)
	// Restore offset
	v2 = origin
	v3 = center
	v4 = Vec2{
		X: v2.X + v3.X,
		Y: v2.Y + v3.Y,
	}
	goto _27
_27:
	center = v4
	return center
}

func b2ComputeCircleMass(tls *_Stack, shape uintptr, density float32) (r MassData) {
	var massData MassData
	var rr float32
	_, _ = massData, rr
	rr = float32((*Circle)(unsafe.Pointer(shape)).Radius * (*Circle)(unsafe.Pointer(shape)).Radius)
	massData.Mass = float32(float32(density*float32FromFloat32(3.14159265359)) * rr)
	massData.Center = (*Circle)(unsafe.Pointer(shape)).Center
	// inertia about the center of mass
	massData.RotationalInertia = float32(float32(massData.Mass*float32FromFloat32(0.5)) * rr)
	return massData
}

func b2ComputeCapsuleMass(tls *_Stack, shape uintptr, density float32) (r MassData) {
	var boxInertia, boxMass, circleInertia, circleMass, h, lc, length, ll, radius, rr, v6 float32
	var massData MassData
	var p1, p2, v1, v2, v3, v5 Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = boxInertia, boxMass, circleInertia, circleMass, h, lc, length, ll, massData, p1, p2, radius, rr, v1, v2, v3, v5, v6
	radius = (*Capsule)(unsafe.Pointer(shape)).Radius
	rr = float32(radius * radius)
	p1 = (*Capsule)(unsafe.Pointer(shape)).Center1
	p2 = (*Capsule)(unsafe.Pointer(shape)).Center2
	v1 = p2
	v2 = p1
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
	length = v6
	ll = float32(length * length)
	circleMass = float32(density * float32(float32(float32FromFloat32(3.14159265359)*radius)*radius))
	boxMass = float32(density * float32(float32(float32FromFloat32(2)*radius)*length))
	massData.Mass = circleMass + boxMass
	massData.Center.X = float32(float32FromFloat32(0.5) * (p1.X + p2.X))
	massData.Center.Y = float32(float32FromFloat32(0.5) * (p1.Y + p2.Y))
	// two offset half circles, both halves add up to full circle and each half is offset by half length
	// semi-circle centroid = 4 r / 3 pi
	// Need to apply parallel-axis theorem twice:
	// 1. shift semi-circle centroid to origin
	// 2. shift semi-circle to box end
	// m * ((h + lc)^2 - lc^2) = m * (h^2 + 2 * h * lc)
	// See: https://en.wikipedia.org/wiki/Parallel_axis_theorem
	// I verified this formula by computing the convex hull of a 128 vertex capsule
	// half circle centroid
	lc = float32(float32FromFloat32(4)*radius) / float32(float32FromFloat32(3)*float32FromFloat32(3.14159265359))
	// half length of rectangular portion of capsule
	h = float32(float32FromFloat32(0.5) * length)
	circleInertia = float32(circleMass * (float32(float32FromFloat32(0.5)*rr) + float32(h*h) + float32(float32(float32FromFloat32(2)*h)*lc)))
	boxInertia = float32(boxMass*(float32(float32FromFloat32(4)*rr)+ll)) / float32FromFloat32(12)
	massData.RotationalInertia = circleInertia + boxInertia
	return massData
}

func b2ComputePolygonMass(tls *_Stack, shape uintptr, density float32) (r1 MassData) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var D, area, ex1, ex2, ey1, ey2, intx2, inty2, inv3, invArea, lengthSquared, radius, rotationalInertia, s1, sqrt2, triangleArea, v12 float32
	var center, e1, e2, mid, n1, n2, r, u, v11, v13, v14, v4, v5, v6, v8, v9 Vec2
	var count, i, i1, i2, j, v3 int32
	var massData MassData
	var vertices [8]Vec2
	var v1 bool
	var _ /* capsule at bp+12 */ Capsule
	var _ /* circle at bp+0 */ Circle
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = D, area, center, count, e1, e2, ex1, ex2, ey1, ey2, i, i1, i2, intx2, inty2, inv3, invArea, j, lengthSquared, massData, mid, n1, n2, r, radius, rotationalInertia, s1, sqrt2, triangleArea, u, vertices, v1, v11, v12, v13, v14, v3, v4, v5, v6, v8, v9
	// Polygon mass, centroid, and inertia.
	// Let rho be the polygon density in mass per unit area.
	// Then:
	// mass = rho * int(dA)
	// centroid.x = (1/mass) * rho * int(x * dA)
	// centroid.y = (1/mass) * rho * int(y * dA)
	// I = rho * int((x*x + y*y) * dA)
	//
	// We can compute these integrals by summing all the integrals
	// for each triangle of the polygon. To evaluate the integral
	// for a single triangle, we make a change of variables to
	// the (u,v) coordinates of the triangle:
	// x = x0 + e1x * u + e2x * v
	// y = y0 + e1y * u + e2y * v
	// where 0 <= u && 0 <= v && u + v <= 1.
	//
	// We integrate u from [0,1-v] and then v from [0,1].
	// We also need to use the Jacobian of the transformation:
	// D = cross(e1, e2)
	//
	// Simplification: triangle centroid = (1/3) * (p1 + p2 + p3)
	//
	// The rest of the derivation is handled by computer algebra.
	if v1 = !!((*Polygon)(unsafe.Pointer(shape)).Count > int32FromInt32(0)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+9832, __ccgo_ts+9594, int32FromInt32(300))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*Polygon)(unsafe.Pointer(shape)).Count == int32(1) {
		(*(*Circle)(unsafe.Add(nil, bp))).Center = *(*Vec2)(unsafe.Add(nil, shape))
		(*(*Circle)(unsafe.Add(nil, bp))).Radius = (*Polygon)(unsafe.Pointer(shape)).Radius
		return b2ComputeCircleMass(tls, bp, density)
	}
	if (*Polygon)(unsafe.Pointer(shape)).Count == int32(2) {
		(*(*Capsule)(unsafe.Add(nil, bp+12))).Center1 = *(*Vec2)(unsafe.Add(nil, shape))
		(*(*Capsule)(unsafe.Add(nil, bp+12))).Center2 = *(*Vec2)(unsafe.Add(nil, shape+1*8))
		(*(*Capsule)(unsafe.Add(nil, bp+12))).Radius = (*Polygon)(unsafe.Pointer(shape)).Radius
		return b2ComputeCapsuleMass(tls, bp+12, density)
	}
	vertices = [8]Vec2{}
	count = (*Polygon)(unsafe.Pointer(shape)).Count
	radius = (*Polygon)(unsafe.Pointer(shape)).Radius
	if radius > float32FromFloat32(0) {
		// Approximate mass of rounded polygons by pushing out the vertices.
		sqrt2 = float32FromFloat32(1.412)
		i = 0
		for {
			if !(i < count) {
				break
			}
			if i == 0 {
				v3 = count - int32(1)
			} else {
				v3 = i - int32(1)
			}
			j = v3
			n1 = *(*Vec2)(unsafe.Add(nil, shape+64+uintptr(j)*8))
			n2 = *(*Vec2)(unsafe.Add(nil, shape+64+uintptr(i)*8))
			v4 = n1
			v5 = n2
			v6 = Vec2{
				X: v4.X + v5.X,
				Y: v4.Y + v5.Y,
			}
			goto _7
		_7:
			v8 = v6
			lengthSquared = float32(v8.X*v8.X) + float32(v8.Y*v8.Y)
			if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
				s1 = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
				u = Vec2{
					X: float32(s1 * v8.X),
					Y: float32(s1 * v8.Y),
				}
				v9 = u
				goto _10
			}
			v9 = Vec2{}
			goto _10
		_10:
			mid = v9
			v11 = *(*Vec2)(unsafe.Add(nil, shape+uintptr(i)*8))
			v12 = float32(sqrt2 * radius)
			v13 = mid
			v14 = Vec2{
				X: v11.X + float32(v12*v13.X),
				Y: v11.Y + float32(v12*v13.Y),
			}
			goto _15
		_15:
			vertices[i] = v14
			goto _2
		_2:
			;
			i = i + 1
		}
	} else {
		i1 = 0
		for {
			if !(i1 < count) {
				break
			}
			vertices[i1] = *(*Vec2)(unsafe.Add(nil, shape+uintptr(i1)*8))
			goto _16
		_16:
			;
			i1 = i1 + 1
		}
	}
	center = Vec2{}
	area = float32FromFloat32(0)
	rotationalInertia = float32FromFloat32(0)
	// Get a reference point for forming triangles.
	// Use the first vertex to reduce round-off errors.
	r = vertices[0]
	inv3 = float32FromFloat32(1) / float32FromFloat32(3)
	i2 = int32(1)
	for {
		if !(i2 < count-int32(1)) {
			break
		}
		v4 = vertices[i2]
		v5 = r
		v6 = Vec2{
			X: v4.X - v5.X,
			Y: v4.Y - v5.Y,
		}
		goto _21
	_21:
		// Triangle edges
		e1 = v6
		v8 = vertices[i2+int32(1)]
		v9 = r
		v11 = Vec2{
			X: v8.X - v9.X,
			Y: v8.Y - v9.Y,
		}
		goto _25
	_25:
		e2 = v11
		v13 = e1
		v14 = e2
		v12 = float32(v13.X*v14.Y) - float32(v13.Y*v14.X)
		goto _29
	_29:
		D = v12
		triangleArea = float32(float32FromFloat32(0.5) * D)
		area = area + triangleArea
		// Area weighted centroid, r at origin
		v4 = e1
		v5 = e2
		v6 = Vec2{
			X: v4.X + v5.X,
			Y: v4.Y + v5.Y,
		}
		goto _33
	_33:
		v8 = center
		v12 = float32(triangleArea * inv3)
		v9 = v6
		v11 = Vec2{
			X: v8.X + float32(v12*v9.X),
			Y: v8.Y + float32(v12*v9.Y),
		}
		goto _38
	_38:
		center = v11
		ex1 = e1.X
		ey1 = e1.Y
		ex2 = e2.X
		ey2 = e2.Y
		intx2 = float32(ex1*ex1) + float32(ex2*ex1) + float32(ex2*ex2)
		inty2 = float32(ey1*ey1) + float32(ey2*ey1) + float32(ey2*ey2)
		rotationalInertia = rotationalInertia + float32(float32(float32(float32FromFloat32(0.25)*inv3)*D)*(intx2+inty2))
		goto _17
	_17:
		;
		i2 = i2 + 1
	}
	// Total mass
	massData.Mass = float32(density * area)
	// Center of mass, shift back from origin at r
	if v1 = !!(area > float32FromFloat32(1.1920928955078125e-07)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+9575, __ccgo_ts+9594, int32FromInt32(384))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	invArea = float32FromFloat32(1) / area
	center.X = float32(center.X * invArea)
	center.Y = float32(center.Y * invArea)
	v4 = r
	v5 = center
	v6 = Vec2{
		X: v4.X + v5.X,
		Y: v4.Y + v5.Y,
	}
	goto _43
_43:
	massData.Center = v6
	// Inertia tensor relative to the local origin (point s).
	massData.RotationalInertia = float32(density * rotationalInertia)
	// Shift inertia to center of mass
	v4 = center
	v5 = center
	v12 = float32(v4.X*v5.X) + float32(v4.Y*v5.Y)
	goto _47
_47:
	massData.RotationalInertia -= float32(massData.Mass * v12)
	// If this goes negative we are hosed
	if v1 = !!(massData.RotationalInertia >= float32FromFloat32(0)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+9849, __ccgo_ts+9594, int32FromInt32(397))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return massData
}

// AABBs are built in double and narrowed to float once. In large world mode the narrowing rounds
// outward so the box always contains the shape, and the inflation (speculative margin) folds into
// the double step, otherwise it vanishes into a float ULP far from the origin. In float mode the
// rounding helpers are plain casts.

func b2ComputeCircleFatAABB(tls *_Stack, shape uintptr, xf WorldTransform, extra float32) (r1 AABB) {
	var aabb AABB
	var c, v3 Pos
	var r float64
	var rx, ry, v11, v5, v7, v9 float32
	var v1 WorldTransform
	var v2 Vec2
	_, _, _, _, _, _, _, _, _, _, _, _ = aabb, c, r, rx, ry, v1, v11, v2, v3, v5, v7, v9
	v1 = xf
	v2 = (*Circle)(unsafe.Pointer(shape)).Center
	rx = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y)
	ry = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y)
	v3 = Pos{
		X: v1.P.X + rx,
		Y: v1.P.Y + ry,
	}
	goto _4
_4:
	c = v3
	r = float64((*Circle)(unsafe.Pointer(shape)).Radius) + float64(extra)
	v5 = float32(float64(c.X) - r)
	goto _6
_6:
	v7 = float32(float64(c.Y) - r)
	goto _8
_8:
	v9 = float32(float64(c.X) + r)
	goto _10
_10:
	v11 = float32(float64(c.Y) + r)
	goto _12
_12:
	aabb = AABB{
		LowerBound: Vec2{
			X: v5,
			Y: v7,
		},
		UpperBound: Vec2{
			X: v9,
			Y: v11,
		},
	}
	return aabb
}

func b2ComputeCapsuleFatAABB(tls *_Stack, shape uintptr, xf WorldTransform, extra float32) (r1 AABB) {
	var aabb AABB
	var r float64
	var rx, ry, v10, v12, v13, v15, v16, v18, v19, v9 float32
	var v1, v2, v3, v7 Pos
	var v11, v5 WorldTransform
	var v21, v6 Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, r, rx, ry, v1, v2, v11, v10, v12, v13, v15, v16, v18, v19, v21, v3, v5, v6, v7, v9
	v11 = xf
	v21 = (*Capsule)(unsafe.Pointer(shape)).Center1
	rx = float32(v11.Q.C*v21.X) - float32(v11.Q.S*v21.Y)
	ry = float32(v11.Q.S*v21.X) + float32(v11.Q.C*v21.Y)
	v3 = Pos{
		X: v11.P.X + rx,
		Y: v11.P.Y + ry,
	}
	goto _4
_4:
	v1 = v3
	v5 = xf
	v6 = (*Capsule)(unsafe.Pointer(shape)).Center2
	rx = float32(v5.Q.C*v6.X) - float32(v5.Q.S*v6.Y)
	ry = float32(v5.Q.S*v6.X) + float32(v5.Q.C*v6.Y)
	v7 = Pos{
		X: v5.P.X + rx,
		Y: v5.P.Y + ry,
	}
	goto _8
_8:
	v2 = v7
	r = float64((*Capsule)(unsafe.Pointer(shape)).Radius) + float64(extra)
	if v1.X < v2.X {
		v9 = v1.X
	} else {
		v9 = v2.X
	}
	v10 = float32(float64(v9) - r)
	goto _11
_11:
	if v1.Y < v2.Y {
		v12 = v1.Y
	} else {
		v12 = v2.Y
	}
	v13 = float32(float64(v12) - r)
	goto _14
_14:
	if v1.X > v2.X {
		v15 = v1.X
	} else {
		v15 = v2.X
	}
	v16 = float32(float64(v15) + r)
	goto _17
_17:
	if v1.Y > v2.Y {
		v18 = v1.Y
	} else {
		v18 = v2.Y
	}
	v19 = float32(float64(v18) + r)
	goto _20
_20:
	aabb = AABB{
		LowerBound: Vec2{
			X: v10,
			Y: v13,
		},
		UpperBound: Vec2{
			X: v16,
			Y: v19,
		},
	}
	return aabb
}

func b2ComputePolygonFatAABB(tls *_Stack, shape uintptr, xf WorldTransform, extra float32) (r1 AABB) {
	var aabb AABB
	var i int32
	var lx, ly, r, ux, uy, v11 float64
	var rx, ry, v15, v17, v19, v21 float32
	var v, v4, v9 Pos
	var v1 bool
	var v2, v7 WorldTransform
	var v3, v8 Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, i, lx, ly, r, rx, ry, ux, uy, v, v1, v11, v15, v17, v19, v2, v21, v3, v4, v7, v8, v9
	if v1 = !!((*Polygon)(unsafe.Pointer(shape)).Count > int32FromInt32(0)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+9832, __ccgo_ts+9594, int32FromInt32(432))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	v2 = xf
	v3 = *(*Vec2)(unsafe.Add(nil, shape))
	rx = float32(v2.Q.C*v3.X) - float32(v2.Q.S*v3.Y)
	ry = float32(v2.Q.S*v3.X) + float32(v2.Q.C*v3.Y)
	v4 = Pos{
		X: v2.P.X + rx,
		Y: v2.P.Y + ry,
	}
	goto _5
_5:
	v = v4
	lx = float64(v.X)
	ly = float64(v.Y)
	ux = float64(v.X)
	uy = float64(v.Y)
	i = int32(1)
	for {
		if !(i < (*Polygon)(unsafe.Pointer(shape)).Count) {
			break
		}
		v7 = xf
		v8 = *(*Vec2)(unsafe.Add(nil, shape+uintptr(i)*8))
		rx = float32(v7.Q.C*v8.X) - float32(v7.Q.S*v8.Y)
		ry = float32(v7.Q.S*v8.X) + float32(v7.Q.C*v8.Y)
		v9 = Pos{
			X: v7.P.X + rx,
			Y: v7.P.Y + ry,
		}
		goto _10
	_10:
		v = v9
		if float64(v.X) < lx {
			v11 = float64(v.X)
		} else {
			v11 = lx
		}
		lx = v11
		if float64(v.Y) < ly {
			v11 = float64(v.Y)
		} else {
			v11 = ly
		}
		ly = v11
		if float64(v.X) > ux {
			v11 = float64(v.X)
		} else {
			v11 = ux
		}
		ux = v11
		if float64(v.Y) > uy {
			v11 = float64(v.Y)
		} else {
			v11 = uy
		}
		uy = v11
		goto _6
	_6:
		;
		i = i + 1
	}
	r = float64((*Polygon)(unsafe.Pointer(shape)).Radius) + float64(extra)
	v15 = float32(lx - r)
	goto _16
_16:
	v17 = float32(ly - r)
	goto _18
_18:
	v19 = float32(ux + r)
	goto _20
_20:
	v21 = float32(uy + r)
	goto _22
_22:
	aabb = AABB{
		LowerBound: Vec2{
			X: v15,
			Y: v17,
		},
		UpperBound: Vec2{
			X: v19,
			Y: v21,
		},
	}
	return aabb
}

func b2ComputeSegmentFatAABB(tls *_Stack, shape uintptr, xf WorldTransform, extra float32) (r AABB) {
	var aabb AABB
	var e float64
	var rx, ry, v10, v12, v13, v15, v16, v18, v19, v9 float32
	var v1, v2, v3, v7 Pos
	var v11, v5 WorldTransform
	var v21, v6 Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, e, rx, ry, v1, v2, v11, v10, v12, v13, v15, v16, v18, v19, v21, v3, v5, v6, v7, v9
	v11 = xf
	v21 = (*Segment)(unsafe.Pointer(shape)).Point1
	rx = float32(v11.Q.C*v21.X) - float32(v11.Q.S*v21.Y)
	ry = float32(v11.Q.S*v21.X) + float32(v11.Q.C*v21.Y)
	v3 = Pos{
		X: v11.P.X + rx,
		Y: v11.P.Y + ry,
	}
	goto _4
_4:
	v1 = v3
	v5 = xf
	v6 = (*Segment)(unsafe.Pointer(shape)).Point2
	rx = float32(v5.Q.C*v6.X) - float32(v5.Q.S*v6.Y)
	ry = float32(v5.Q.S*v6.X) + float32(v5.Q.C*v6.Y)
	v7 = Pos{
		X: v5.P.X + rx,
		Y: v5.P.Y + ry,
	}
	goto _8
_8:
	v2 = v7
	e = float64(extra)
	if v1.X < v2.X {
		v9 = v1.X
	} else {
		v9 = v2.X
	}
	v10 = float32(float64(v9) - e)
	goto _11
_11:
	if v1.Y < v2.Y {
		v12 = v1.Y
	} else {
		v12 = v2.Y
	}
	v13 = float32(float64(v12) - e)
	goto _14
_14:
	if v1.X > v2.X {
		v15 = v1.X
	} else {
		v15 = v2.X
	}
	v16 = float32(float64(v15) + e)
	goto _17
_17:
	if v1.Y > v2.Y {
		v18 = v1.Y
	} else {
		v18 = v2.Y
	}
	v19 = float32(float64(v18) + e)
	goto _20
_20:
	aabb = AABB{
		LowerBound: Vec2{
			X: v10,
			Y: v13,
		},
		UpperBound: Vec2{
			X: v16,
			Y: v19,
		},
	}
	return aabb
}

func b2ComputeCircleBounds(tls *_Stack, shape uintptr, xf WorldTransform) (r AABB) {
	return b2ComputeCircleFatAABB(tls, shape, xf, float32FromFloat32(0))
}

func b2ComputeCapsuleBounds(tls *_Stack, shape uintptr, xf WorldTransform) (r AABB) {
	return b2ComputeCapsuleFatAABB(tls, shape, xf, float32FromFloat32(0))
}

func b2ComputePolygonBounds(tls *_Stack, shape uintptr, xf WorldTransform) (r AABB) {
	return b2ComputePolygonFatAABB(tls, shape, xf, float32FromFloat32(0))
}

func b2ComputeSegmentBounds(tls *_Stack, shape uintptr, xf WorldTransform) (r AABB) {
	return b2ComputeSegmentFatAABB(tls, shape, xf, float32FromFloat32(0))
}

// C documentation
//
//	// quickhull algorithm
//	// - merges vertices based on B2_LINEAR_SLOP
//	// - removes collinear points using B2_LINEAR_SLOP
//	// - returns an empty hull if it fails
func b2ComputeHull(tls *_Stack, points uintptr, count int32) (r1 Hull) {
	bp := tls.Alloc(304)
	defer tls.Free(304)
	var aabb, v41 AABB
	var b8, c, c1, c2, c3, e, p1, p2, r, s1, s2, s3, u, vi, vj, v19, v21, v22, v33, v36, v37, v42, v44, v7, v8 Vec2
	var d, distSqr, distance, dsq, dsq1, dsq11, dsq2, lengthSquared, linearSlop, s, tolSqr, v10, v11, v13, v14, v15, v16, v18, v9 float32
	var f1, f2, i, i1, i11, i2, i21, i3, i31, i4, i5, i6, j, j1, leftCount, n, rightCount, v1, v2, v3, v5 int32
	var ps [8]Vec2
	var searching, unique uint8
	var v81 uintptr
	var v90 bool
	var _ /* hull at bp+0 */ Hull
	var _ /* hull1 at bp+164 */ Hull
	var _ /* hull2 at bp+232 */ Hull
	var _ /* leftPoints at bp+116 */ [6]Vec2
	var _ /* rightPoints at bp+68 */ [6]Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, b8, c, c1, c2, c3, d, distSqr, distance, dsq, dsq1, dsq11, dsq2, e, f1, f2, i, i1, i11, i2, i21, i3, i31, i4, i5, i6, j, j1, leftCount, lengthSquared, linearSlop, n, p1, p2, ps, r, rightCount, s, s1, s2, s3, searching, tolSqr, u, unique, vi, vj, v1, v10, v11, v13, v14, v15, v16, v18, v19, v2, v21, v22, v3, v33, v36, v37, v41, v42, v44, v5, v7, v8, v81, v9, v90
	*(*Hull)(unsafe.Add(nil, bp)) = Hull{}
	(*(*Hull)(unsafe.Add(nil, bp))).Count = 0
	if count < int32(3) || count > int32(_B2_MAX_POLYGON_VERTICES) {
		// check your data
		return *(*Hull)(unsafe.Add(nil, bp))
	}
	v1 = count
	v2 = int32(_B2_MAX_POLYGON_VERTICES)
	if v1 < v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _4
_4:
	count = v3
	aabb = AABB{
		LowerBound: Vec2{
			X: float32FromFloat32(3.4028234663852886e+38),
			Y: float32FromFloat32(3.4028234663852886e+38),
		},
		UpperBound: Vec2{
			X: -float32FromFloat32(3.4028234663852886e+38),
			Y: -float32FromFloat32(3.4028234663852886e+38),
		},
	}
	n = 0
	linearSlop = float32(float32FromFloat32(0.005) * b2_lengthUnitsPerMeter)
	tolSqr = float32(float32(float32FromFloat32(0.25)*linearSlop) * linearSlop)
	i = 0
	for {
		if !(i < count) {
			break
		}
		v7 = aabb.LowerBound
		v8 = *(*Vec2)(unsafe.Add(nil, points+uintptr(i)*8))
		v9 = v7.X
		v10 = v8.X
		if v9 < v10 {
			v13 = v9
		} else {
			v13 = v10
		}
		v11 = v13
		goto _12
	_12:
		c.X = v11
		v14 = v7.Y
		v15 = v8.Y
		if v14 < v15 {
			v18 = v14
		} else {
			v18 = v15
		}
		v16 = v18
		goto _17
	_17:
		c.Y = v16
		v19 = c
		goto _20
	_20:
		aabb.LowerBound = v19
		v7 = aabb.UpperBound
		v8 = *(*Vec2)(unsafe.Add(nil, points+uintptr(i)*8))
		v9 = v7.X
		v10 = v8.X
		if v9 > v10 {
			v13 = v9
		} else {
			v13 = v10
		}
		v11 = v13
		goto _26
	_26:
		c1.X = v11
		v14 = v7.Y
		v15 = v8.Y
		if v14 > v15 {
			v18 = v14
		} else {
			v18 = v15
		}
		v16 = v18
		goto _31
	_31:
		c1.Y = v16
		v19 = c1
		goto _34
	_34:
		aabb.UpperBound = v19
		vi = *(*Vec2)(unsafe.Add(nil, points+uintptr(i)*8))
		unique = boolUint8(true1 != 0)
		j = 0
		for {
			if !(j < n) {
				break
			}
			vj = ps[j]
			v7 = vi
			v8 = vj
			c2 = Vec2{
				X: v8.X - v7.X,
				Y: v8.Y - v7.Y,
			}
			v9 = float32(c2.X*c2.X) + float32(c2.Y*c2.Y)
			goto _39
		_39:
			distSqr = v9
			if distSqr < tolSqr {
				unique = boolUint8(false1 != 0)
				break
			}
			goto _35
		_35:
			;
			j = j + 1
		}
		if unique != 0 {
			v1 = n
			n = n + 1
			ps[v1] = vi
		}
		goto _6
	_6:
		;
		i = i + 1
	}
	if n < int32(3) {
		// all points very close together, check your data and check your scale
		return *(*Hull)(unsafe.Add(nil, bp))
	}
	v41 = aabb
	b8 = Vec2{
		X: float32(float32FromFloat32(0.5) * (v41.LowerBound.X + v41.UpperBound.X)),
		Y: float32(float32FromFloat32(0.5) * (v41.LowerBound.Y + v41.UpperBound.Y)),
	}
	v7 = b8
	goto _43
_43:
	// Find an extreme point as the first point on the hull
	c3 = v7
	f1 = 0
	v8 = c3
	v19 = ps[f1]
	c2 = Vec2{
		X: v19.X - v8.X,
		Y: v19.Y - v8.Y,
	}
	v9 = float32(c2.X*c2.X) + float32(c2.Y*c2.Y)
	goto _47
_47:
	dsq11 = v9
	i1 = int32(1)
	for {
		if !(i1 < n) {
			break
		}
		v21 = c3
		v22 = ps[i1]
		c2 = Vec2{
			X: v22.X - v21.X,
			Y: v22.Y - v21.Y,
		}
		v10 = float32(c2.X*c2.X) + float32(c2.Y*c2.Y)
		goto _52
	_52:
		dsq = v10
		if dsq > dsq11 {
			f1 = i1
			dsq11 = dsq
		}
		goto _48
	_48:
		;
		i1 = i1 + 1
	}
	// remove p1 from working set
	p1 = ps[f1]
	ps[f1] = ps[n-int32(1)]
	n = n - int32(1)
	f2 = 0
	v7 = p1
	v8 = ps[f2]
	c2 = Vec2{
		X: v8.X - v7.X,
		Y: v8.Y - v7.Y,
	}
	v9 = float32(c2.X*c2.X) + float32(c2.Y*c2.Y)
	goto _56
_56:
	dsq2 = v9
	i2 = int32(1)
	for {
		if !(i2 < n) {
			break
		}
		v19 = p1
		v21 = ps[i2]
		c2 = Vec2{
			X: v21.X - v19.X,
			Y: v21.Y - v19.Y,
		}
		v10 = float32(c2.X*c2.X) + float32(c2.Y*c2.Y)
		goto _61
	_61:
		dsq1 = v10
		if dsq1 > dsq2 {
			f2 = i2
			dsq2 = dsq1
		}
		goto _57
	_57:
		;
		i2 = i2 + 1
	}
	// remove p2 from working set
	p2 = ps[f2]
	ps[f2] = ps[n-int32(1)]
	n = n - int32(1)
	rightCount = 0
	leftCount = 0
	v7 = p2
	v8 = p1
	v19 = Vec2{
		X: v7.X - v8.X,
		Y: v7.Y - v8.Y,
	}
	goto _65
_65:
	v21 = v19
	lengthSquared = float32(v21.X*v21.X) + float32(v21.Y*v21.Y)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
		u = Vec2{
			X: float32(s * v21.X),
			Y: float32(s * v21.Y),
		}
		v22 = u
		goto _68
	}
	v22 = Vec2{}
	goto _68
_68:
	e = v22
	i3 = 0
	for {
		if !(i3 < n) {
			break
		}
		v33 = ps[i3]
		v36 = p1
		v37 = Vec2{
			X: v33.X - v36.X,
			Y: v33.Y - v36.Y,
		}
		goto _73
	_73:
		v42 = v37
		v44 = e
		v9 = float32(v42.X*v44.Y) - float32(v42.Y*v44.X)
		goto _77
	_77:
		d = v9
		// slop used here to skip points that are very close to the line p1-p2
		if d >= float32(float32FromFloat32(2)*linearSlop) {
			v1 = rightCount
			rightCount = rightCount + 1
			(*(*[6]Vec2)(unsafe.Add(nil, bp+68)))[v1] = ps[i3]
		} else {
			if d <= float32(-float32FromFloat32(2)*linearSlop) {
				v1 = leftCount
				leftCount = leftCount + 1
				(*(*[6]Vec2)(unsafe.Add(nil, bp+116)))[v1] = ps[i3]
			}
		}
		goto _69
	_69:
		;
		i3 = i3 + 1
	}
	// compute hulls on right and left
	*(*Hull)(unsafe.Add(nil, bp+164)) = b2RecurseHull(tls, p1, p2, bp+68, rightCount)
	*(*Hull)(unsafe.Add(nil, bp+232)) = b2RecurseHull(tls, p2, p1, bp+116, leftCount)
	if (*(*Hull)(unsafe.Add(nil, bp+164))).Count == 0 && (*(*Hull)(unsafe.Add(nil, bp+232))).Count == 0 {
		// all points collinear
		return *(*Hull)(unsafe.Add(nil, bp))
	}
	// stitch hulls together, preserving CCW winding order
	v81 = bp + 64
	v1 = *(*int32)(unsafe.Pointer(v81))
	*(*int32)(unsafe.Pointer(v81)) = *(*int32)(unsafe.Pointer(v81)) + 1
	*(*Vec2)(unsafe.Add(nil, bp+uintptr(v1)*8)) = p1
	i4 = 0
	for {
		if !(i4 < (*(*Hull)(unsafe.Add(nil, bp+164))).Count) {
			break
		}
		v81 = bp + 64
		v1 = *(*int32)(unsafe.Pointer(v81))
		*(*int32)(unsafe.Pointer(v81)) = *(*int32)(unsafe.Pointer(v81)) + 1
		*(*Vec2)(unsafe.Add(nil, bp+uintptr(v1)*8)) = *(*Vec2)(unsafe.Add(nil, bp+164+uintptr(i4)*8))
		goto _82
	_82:
		;
		i4 = i4 + 1
	}
	v81 = bp + 64
	v1 = *(*int32)(unsafe.Pointer(v81))
	*(*int32)(unsafe.Pointer(v81)) = *(*int32)(unsafe.Pointer(v81)) + 1
	*(*Vec2)(unsafe.Add(nil, bp+uintptr(v1)*8)) = p2
	i5 = 0
	for {
		if !(i5 < (*(*Hull)(unsafe.Add(nil, bp+232))).Count) {
			break
		}
		v81 = bp + 64
		v1 = *(*int32)(unsafe.Pointer(v81))
		*(*int32)(unsafe.Pointer(v81)) = *(*int32)(unsafe.Pointer(v81)) + 1
		*(*Vec2)(unsafe.Add(nil, bp+uintptr(v1)*8)) = *(*Vec2)(unsafe.Add(nil, bp+232+uintptr(i5)*8))
		goto _87
	_87:
		;
		i5 = i5 + 1
	}
	if v90 = !!((*(*Hull)(unsafe.Add(nil, bp))).Count <= int32FromInt32(_B2_MAX_POLYGON_VERTICES)); !v90 {
		b2InternalAssert(tls, __ccgo_ts+10010, __ccgo_ts+9988, int32FromInt32(225))
	}
	_ = v90 || bool(int32FromInt32(0) != 0)
	// merge collinear
	searching = boolUint8(true1 != 0)
	for searching != 0 && (*(*Hull)(unsafe.Add(nil, bp))).Count > int32(2) {
		searching = boolUint8(false1 != 0)
		i6 = 0
		for {
			if !(i6 < (*(*Hull)(unsafe.Add(nil, bp))).Count) {
				break
			}
			i11 = i6
			i21 = (i6 + int32(1)) % (*(*Hull)(unsafe.Add(nil, bp))).Count
			i31 = (i6 + int32(2)) % (*(*Hull)(unsafe.Add(nil, bp))).Count
			s1 = *(*Vec2)(unsafe.Add(nil, bp+uintptr(i11)*8))
			s2 = *(*Vec2)(unsafe.Add(nil, bp+uintptr(i21)*8))
			s3 = *(*Vec2)(unsafe.Add(nil, bp+uintptr(i31)*8))
			v7 = s3
			v8 = s1
			v19 = Vec2{
				X: v7.X - v8.X,
				Y: v7.Y - v8.Y,
			}
			goto _95
		_95:
			v21 = v19
			lengthSquared = float32(v21.X*v21.X) + float32(v21.Y*v21.Y)
			if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
				s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
				u = Vec2{
					X: float32(s * v21.X),
					Y: float32(s * v21.Y),
				}
				v22 = u
				goto _98
			}
			v22 = Vec2{}
			goto _98
		_98:
			// unit edge vector for s1-s3
			r = v22
			v33 = s2
			v36 = s1
			v37 = Vec2{
				X: v33.X - v36.X,
				Y: v33.Y - v36.Y,
			}
			goto _102
		_102:
			v42 = v37
			v44 = r
			v9 = float32(v42.X*v44.Y) - float32(v42.Y*v44.X)
			goto _106
		_106:
			distance = v9
			if distance <= float32(float32FromFloat32(2)*linearSlop) {
				// remove midpoint from hull
				j1 = i21
				for {
					if !(j1 < (*(*Hull)(unsafe.Add(nil, bp))).Count-int32(1)) {
						break
					}
					*(*Vec2)(unsafe.Add(nil, bp+uintptr(j1)*8)) = *(*Vec2)(unsafe.Add(nil, bp+uintptr(j1+int32(1))*8))
					goto _107
				_107:
					;
					j1 = j1 + 1
				}
				(*(*Hull)(unsafe.Add(nil, bp))).Count -= int32(1)
				// continue searching for collinear points
				searching = boolUint8(true1 != 0)
				break
			}
			goto _91
		_91:
			;
			i6 = i6 + 1
		}
	}
	if (*(*Hull)(unsafe.Add(nil, bp))).Count < int32(3) {
		// all points collinear, shouldn't be reached since this was validated above
		(*(*Hull)(unsafe.Add(nil, bp))).Count = 0
	}
	return *(*Hull)(unsafe.Add(nil, bp))
}

// C documentation
//
//	// Approximate cosine and sine for determinism. In my testing cosf and sinf produced
//	// the same results on x64 and ARM using MSVC, GCC, and Clang. However, I don't trust
//	// this result.
//	// https://en.wikipedia.org/wiki/Bh%C4%81skara_I%27s_sine_approximation_formula
func b2ComputeCosSin(tls *_Stack, radians1 float32) (r CosSin) {
	var a1, k, roundToNearest, twoPi float64
	var c, invMag, mag, pi2, s, x, x1, y, y1, y2, y21, y22, y23, v1, v2, v3, v4, v6, v7, v8 float32
	var cs CosSin
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a1, c, cs, invMag, k, mag, pi2, roundToNearest, s, twoPi, x, x1, y, y1, y2, y21, y22, y23, v1, v2, v3, v4, v6, v7, v8
	v1 = radians1
	v2 = -float32FromFloat32(1e+06)
	v3 = float32FromFloat32(1e+06)
	if v1 < v2 {
		v6 = v2
	} else {
		if v1 > v3 {
			v7 = v3
		} else {
			v7 = v1
		}
		v6 = v7
	}
	v4 = v6
	goto _5
_5:
	x = v4
	twoPi = float64(float32FromFloat32(2) * float32FromFloat32(3.14159265359))
	roundToNearest = float64(6.755399441055744e+15)
	a1 = float64(x)
	k = a1/twoPi + roundToNearest - roundToNearest
	v8 = float32(a1 - float64(k*twoPi))
	goto _9
_9:
	x1 = v8
	pi2 = float32(float32FromFloat32(3.14159265359) * float32FromFloat32(3.14159265359))
	if x1 < float32(-float32FromFloat32(0.5)*float32FromFloat32(3.14159265359)) {
		y = x1 + float32FromFloat32(3.14159265359)
		y21 = float32(y * y)
		c = -(pi2 - float32(float32FromFloat32(4)*y21)) / (pi2 + y21)
	} else {
		if x1 > float32(float32FromFloat32(0.5)*float32FromFloat32(3.14159265359)) {
			y1 = x1 - float32FromFloat32(3.14159265359)
			y22 = float32(y1 * y1)
			c = -(pi2 - float32(float32FromFloat32(4)*y22)) / (pi2 + y22)
		} else {
			y23 = float32(x1 * x1)
			c = (pi2 - float32(float32FromFloat32(4)*y23)) / (pi2 + y23)
		}
	}
	if x1 < float32FromFloat32(0) {
		y2 = x1 + float32FromFloat32(3.14159265359)
		s = float32(float32(-float32FromFloat32(16)*y2)*(float32FromFloat32(3.14159265359)-y2)) / (float32(float32FromFloat32(5)*pi2) - float32(float32(float32FromFloat32(4)*y2)*(float32FromFloat32(3.14159265359)-y2)))
	} else {
		s = float32(float32(float32FromFloat32(16)*x1)*(float32FromFloat32(3.14159265359)-x1)) / (float32(float32FromFloat32(5)*pi2) - float32(float32(float32FromFloat32(4)*x1)*(float32FromFloat32(3.14159265359)-x1)))
	}
	mag = sqrtf(tls, float32(s*s)+float32(c*c))
	if mag > float32FromFloat32(0) {
		v1 = float32FromFloat32(1) / mag
	} else {
		v1 = float32FromFloat32(0)
	}
	invMag = v1
	cs = CosSin{
		Cosine: float32(c * invMag),
		Sine:   float32(s * invMag),
	}
	return cs
}

func b2ComputeRotationBetweenUnitVectors(tls *_Stack, v1 Vec2, v2 Vec2) (r Rot) {
	var invMag, mag, v21, v4, v5, v7 float32
	var qn, rot, v25, v27 Rot
	var v11, v9 Vec2
	var v8 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _ = invMag, mag, qn, rot, v11, v21, v25, v27, v4, v5, v7, v8, v9
	v11 = v1
	v21 = sqrtf(tls, float32(v11.X*v11.X)+float32(v11.Y*v11.Y))
	goto _3
_3:
	v4 = float32FromFloat32(1) - v21
	if v4 < float32FromInt32(0) {
		v7 = -v4
	} else {
		v7 = v4
	}
	v5 = v7
	goto _6
_6:
	;
	if v8 = !!(v5 < float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07))); !v8 {
		b2InternalAssert(tls, __ccgo_ts+16047, __ccgo_ts+16106, int32FromInt32(187))
	}
	_ = v8 || bool(int32FromInt32(0) != 0)
	v11 = v2
	v21 = sqrtf(tls, float32(v11.X*v11.X)+float32(v11.Y*v11.Y))
	goto _11
_11:
	v4 = float32FromFloat32(1) - v21
	if v4 < float32FromInt32(0) {
		v7 = -v4
	} else {
		v7 = v4
	}
	v5 = v7
	goto _14
_14:
	;
	if v8 = !!(v5 < float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07))); !v8 {
		b2InternalAssert(tls, __ccgo_ts+16138, __ccgo_ts+16106, int32FromInt32(188))
	}
	_ = v8 || bool(int32FromInt32(0) != 0)
	v11 = v1
	v9 = v2
	v21 = float32(v11.X*v9.X) + float32(v11.Y*v9.Y)
	goto _20
_20:
	rot.C = v21
	v11 = v1
	v9 = v2
	v21 = float32(v11.X*v9.Y) - float32(v11.Y*v9.X)
	goto _24
_24:
	rot.S = v21
	v25 = rot
	mag = sqrtf(tls, float32(v25.S*v25.S)+float32(v25.C*v25.C))
	if mag > float32FromFloat32(0) {
		v21 = float32FromFloat32(1) / mag
	} else {
		v21 = float32FromFloat32(0)
	}
	invMag = v21
	qn = Rot{
		C: float32(v25.C * invMag),
		S: float32(v25.S * invMag),
	}
	v27 = qn
	goto _28
_28:
	return v27
}

// Snapshot image magic and version

// Bump this if any of the data structures below get modified. The layout hash only catches
// size changes, a same-size reinterpretation like the contact cache reshape needs this bump.

// Header flag bits

// C documentation
//
//	// Layout hash seeds from all structs we memcpy, plus key constants.
//	// Changing any struct or constant updates the hash. Re-purposing or swapping
//	// fields might not change the hash
func b2ComputeLayoutHash(tls *_Stack) (r uint32) {
	var h uint32_t
	_ = h
	// FNV-1a hash
	h = uint32(2166136261)
	h = h ^ uint32(uint64FromInt64(104))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(96))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(32))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(280))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(32))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(64))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(200))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(72))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(252))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(64))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(4))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(12))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(12))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(56))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(8))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(88))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(64))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(104))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(32))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(24))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(8))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(24))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(32))
	h = h * uint32(16777619)
	h = h ^ uint32FromInt32(int32FromInt32(_B2_GRAPH_COLOR_COUNT))
	h = h * uint32(16777619)
	h = h ^ uint32FromInt32(int32(b2_bodyTypeCount))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(8))
	h = h * uint32(16777619)
	return h
}

// C documentation
//
//	// A block is a range of tasks, a start index and count as a sub-array. Each worker receives at
//	// most M blocks of work. The workers may receive less blocks if there is not sufficient work.
//	// Each block of work has a minimum number of elements (block size). This in turn may limit the
//	// number of blocks. If there are many elements then the block size is increased so there are
//	// still at most M blocks of work per worker. M is a tunable number that has two goals:
//	// 1. keep M small to reduce overhead
//	// 2. keep M large enough for other workers to be able to steal work
//	// The block size is a power of two to make math efficient.
func b2ComputeBlockCount(tls *_Stack, itemCount int32, minSize int32, maxBlockCount int32) (r b2BlockDim) {
	var dim b2BlockDim
	var v1 bool
	_, _ = dim, v1
	dim = b2BlockDim{}
	if itemCount == 0 {
		return dim
	}
	if itemCount <= minSize*maxBlockCount {
		dim.Size = minSize
	} else {
		dim.Size = (itemCount + maxBlockCount - int32(1)) / maxBlockCount
	}
	dim.Count = (itemCount + dim.Size - int32(1)) / dim.Size
	if v1 = !!(dim.Count >= int32FromInt32(1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+24810, __ccgo_ts+24283, int32FromInt32(789))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(dim.Size*dim.Count >= itemCount); !v1 {
		b2InternalAssert(tls, __ccgo_ts+24825, __ccgo_ts+24283, int32FromInt32(790))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return dim
}
