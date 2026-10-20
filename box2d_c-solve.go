package b2

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

// C documentation
//
//	// Solve contacts: normal, friction and rolling resistance.
func b2SolveContacts_Overflow(tls *_Stack, context uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var Cdot, angularA, angularB, friction, iA, iB, impulse, impulse1, inv_h, lambda, mA, mB, maxFriction, maxLambda, normalImpulse, rollingImpulse, s3, tangentImpulse, totalNormalImpulse, velocityBias, vn, vnc, vt, vtc, wA, wB, v18, v22, v27, v31, v35, v39 float32
	var P, dp, ds, dv, normal, rA, rB, tangent, vA, vB, v10, v12, v13, v14, v16, v17, v20, v21, v25, v26, v5, v6, v7, v9 Vec2
	var awakeSet, color, constraint, constraints, cp, cp1, cp2, graph, stateA, stateB, states, world, v3, v4 uintptr
	var contactCount, i, indexA, indexB, j, j1, j2, j3, j4, k, k1, k2, pointCount int32
	var dqA, dqB, v42, v46 Rot
	var normalImpulses, rnA, rnB, rtA, rtB, tangentImpulses [2]float32
	var v1 bool
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = Cdot, P, angularA, angularB, awakeSet, color, constraint, constraints, contactCount, cp, cp1, cp2, dp, dqA, dqB, ds, dv, friction, graph, i, iA, iB, impulse, impulse1, indexA, indexB, inv_h, j, j1, j2, j3, j4, k, k1, k2, lambda, mA, mB, maxFriction, maxLambda, normal, normalImpulse, normalImpulses, pointCount, rA, rB, rnA, rnB, rollingImpulse, rtA, rtB, s3, stateA, stateB, states, tangent, tangentImpulse, tangentImpulses, totalNormalImpulse, vA, vB, velocityBias, vn, vnc, vt, vtc, wA, wB, world, v1, v10, v12, v13, v14, v16, v17, v18, v20, v21, v22, v25, v26, v27, v3, v31, v35, v39, v4, v42, v46, v5, v6, v7, v9
	graph = (*b2StepContext)(unsafe.Pointer(context)).Graph
	color = graph + uintptr(int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1))*64
	constraints = *(*uintptr)(unsafe.Add(unsafe.Pointer(color), 48))
	contactCount = (*b2GraphColor)(unsafe.Pointer(color)).ContactSims.Count
	world = (*b2StepContext)(unsafe.Pointer(context)).World
	if v1 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+7324, int32FromInt32(406))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	awakeSet = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
	states = (*b2SolverSet)(unsafe.Pointer(awakeSet)).BodyStates.Data
	inv_h = (*b2StepContext)(unsafe.Pointer(context)).Inv_h
	// This is a dummy body to represent a static body since static bodies don't have a solver body.
	*(*b2BodyState)(unsafe.Add(nil, bp)) = b2_identityBodyState4
	i = 0
	for {
		if !(i < contactCount) {
			break
		}
		constraint = constraints + uintptr(i)*168
		mA = (*b2ContactConstraint)(unsafe.Pointer(constraint)).InvMassA
		iA = (*b2ContactConstraint)(unsafe.Pointer(constraint)).InvIA
		mB = (*b2ContactConstraint)(unsafe.Pointer(constraint)).InvMassB
		iB = (*b2ContactConstraint)(unsafe.Pointer(constraint)).InvIB
		indexA = (*b2ContactConstraint)(unsafe.Pointer(constraint)).IndexA - int32(1)
		indexB = (*b2ContactConstraint)(unsafe.Pointer(constraint)).IndexB - int32(1)
		if indexA == -int32(1) {
			v3 = bp
		} else {
			v3 = states + uintptr(indexA)*32
		}
		stateA = v3
		vA = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
		wA = (*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity
		dqA = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
		if indexB == -int32(1) {
			v4 = bp
		} else {
			v4 = states + uintptr(indexB)*32
		}
		stateB = v4
		vB = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
		wB = (*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity
		dqB = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
		v5 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaPosition
		v6 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaPosition
		v7 = Vec2{
			X: v5.X - v6.X,
			Y: v5.Y - v6.Y,
		}
		goto _8
	_8:
		dp = v7
		normal = (*b2ContactConstraint)(unsafe.Pointer(constraint)).Normal
		v9 = normal
		v10 = Vec2{
			X: v9.Y,
			Y: -v9.X,
		}
		goto _11
	_11:
		tangent = v10
		friction = (*b2ContactConstraint)(unsafe.Pointer(constraint)).Friction
		pointCount = (*b2ContactConstraint)(unsafe.Pointer(constraint)).PointCount
		totalNormalImpulse = float32FromFloat32(0)
		v12 = vB
		v13 = vA
		v14 = Vec2{
			X: v12.X - v13.X,
			Y: v12.Y - v13.Y,
		}
		goto _15
	_15:
		dv = v14
		v16 = dv
		v17 = normal
		v18 = float32(v16.X*v17.X) + float32(v16.Y*v17.Y)
		goto _19
	_19:
		vnc = v18
		v20 = dv
		v21 = tangent
		v22 = float32(v20.X*v21.X) + float32(v20.Y*v21.Y)
		goto _23
	_23:
		// vt = dot(vrB - sB * tangent - (vrA + sA * tangent), tangent)
		//    = dot(vrB - vrA, tangent) - (sA + sB)
		vtc = v22 - (*b2ContactConstraint)(unsafe.Pointer(constraint)).TangentSpeed
		normalImpulses = [2]float32{}
		tangentImpulses = [2]float32{}
		j = 0
		for {
			if !(j < pointCount) {
				break
			}
			cp = constraint + 8 + uintptr(j)*48
			v25 = (*b2ContactConstraintPoint)(unsafe.Pointer(cp)).AnchorA
			v26 = normal
			v27 = float32(v25.X*v26.Y) - float32(v25.Y*v26.X)
			goto _28
		_28:
			rnA[j] = v27
			v5 = (*b2ContactConstraintPoint)(unsafe.Pointer(cp)).AnchorB
			v6 = normal
			v18 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
			goto _32
		_32:
			rnB[j] = v18
			v5 = (*b2ContactConstraintPoint)(unsafe.Pointer(cp)).AnchorA
			v6 = tangent
			v18 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
			goto _36
		_36:
			rtA[j] = v18
			v5 = (*b2ContactConstraintPoint)(unsafe.Pointer(cp)).AnchorB
			v6 = tangent
			v18 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
			goto _40
		_40:
			rtB[j] = v18
			goto _24
		_24:
			;
			j = j + 1
		}
		j1 = 0
		for {
			if !(j1 < pointCount) {
				break
			}
			cp1 = constraint + 8 + uintptr(j1)*48
			// Fixed anchor points.
			rA = (*b2ContactConstraintPoint)(unsafe.Pointer(cp1)).AnchorA
			rB = (*b2ContactConstraintPoint)(unsafe.Pointer(cp1)).AnchorB
			v42 = dqB
			v5 = rB
			v6 = Vec2{
				X: float32(v42.C*v5.X) - float32(v42.S*v5.Y),
				Y: float32(v42.S*v5.X) + float32(v42.C*v5.Y),
			}
			goto _45
		_45:
			v46 = dqA
			v7 = rA
			v9 = Vec2{
				X: float32(v46.C*v7.X) - float32(v46.S*v7.Y),
				Y: float32(v46.S*v7.X) + float32(v46.C*v7.Y),
			}
			goto _49
		_49:
			v10 = v6
			v12 = v9
			v13 = Vec2{
				X: v10.X - v12.X,
				Y: v10.Y - v12.Y,
			}
			goto _53
		_53:
			v14 = dp
			v16 = v13
			v17 = Vec2{
				X: v14.X + v16.X,
				Y: v14.Y + v16.Y,
			}
			goto _57
		_57:
			// Compute current separation.
			// This is subject to round-off error if the anchor is far from the body center of mass.
			ds = v17
			v20 = ds
			v21 = normal
			v18 = float32(v20.X*v21.X) + float32(v20.Y*v21.Y)
			goto _61
		_61:
			s3 = (*b2ContactConstraintPoint)(unsafe.Pointer(cp1)).BaseSeparation + v18
			if s3 > float32FromFloat32(0) {
				v22 = float32(s3 * inv_h)
			} else {
				v22 = float32FromFloat32(0)
			}
			// Speculative bias, zero when overlapped.
			velocityBias = v22
			// Relative normal velocity at contact.
			vn = vnc + float32(wB*rnB[j1]) - float32(wA*rnA[j1])
			k = 0
			for {
				if !(k < j1) {
					break
				}
				// This allows getting point 2 normal velocity with only one dependency.
				vn = vn + float32((mA+mB+float32(float32(iA*rnA[k])*rnA[j1])+float32(float32(iB*rnB[k])*rnB[j1]))*normalImpulses[k])
				goto _63
			_63:
				;
				k = k + 1
			}
			v18 = float32(-(*b2ContactConstraintPoint)(unsafe.Pointer(cp1)).NormalMass * (vn + velocityBias))
			v22 = -(*b2ContactConstraintPoint)(unsafe.Pointer(cp1)).NormalImpulse
			if v18 > v22 {
				v31 = v18
			} else {
				v31 = v22
			}
			v27 = v31
			goto _67
		_67:
			impulse = v27
			*(*float32)(unsafe.Add(nil, cp1+24)) += impulse
			*(*float32)(unsafe.Add(nil, cp1+32)) += impulse
			normalImpulses[j1] = impulse
			// This is used for rolling resistance
			totalNormalImpulse = totalNormalImpulse + (*b2ContactConstraintPoint)(unsafe.Pointer(cp1)).NormalImpulse
			goto _41
		_41:
			;
			j1 = j1 + 1
		}
		Cdot = wB - wA
		j2 = 0
		for {
			if !(j2 < pointCount) {
				break
			}
			Cdot = Cdot + float32((float32(iA*rnA[j2])+float32(iB*rnB[j2]))*normalImpulses[j2])
			goto _69
		_69:
			;
			j2 = j2 + 1
		}
		lambda = (*b2ContactConstraint)(unsafe.Pointer(constraint)).RollingImpulse
		maxLambda = float32((*b2ContactConstraint)(unsafe.Pointer(constraint)).RollingResistance * totalNormalImpulse)
		v18 = lambda - float32((*b2ContactConstraint)(unsafe.Pointer(constraint)).RollingMass*Cdot)
		v22 = -maxLambda
		v27 = maxLambda
		if v18 < v22 {
			v35 = v22
		} else {
			if v18 > v27 {
				v39 = v27
			} else {
				v39 = v18
			}
			v35 = v39
		}
		v31 = v35
		goto _74
	_74:
		(*b2ContactConstraint)(unsafe.Pointer(constraint)).RollingImpulse = v31
		rollingImpulse = (*b2ContactConstraint)(unsafe.Pointer(constraint)).RollingImpulse - lambda
		// Friction
		j3 = 0
		for {
			if !(j3 < pointCount) {
				break
			}
			cp2 = constraint + 8 + uintptr(j3)*48
			// relative tangent velocity at contact
			vt = vtc + float32(wB*rtB[j3]) - float32(wA*rtA[j3]) + float32((float32(iA*rtA[j3])+float32(iB*rtB[j3]))*rollingImpulse)
			k1 = 0
			for {
				if !(k1 < pointCount) {
					break
				}
				vt = vt + float32((float32(float32(iA*rtA[j3])*rnA[k1])+float32(float32(iB*rtB[j3])*rnB[k1]))*normalImpulses[k1])
				goto _78
			_78:
				;
				k1 = k1 + 1
			}
			k2 = 0
			for {
				if !(k2 < j3) {
					break
				}
				vt = vt + float32((mA+mB+float32(float32(iA*rtA[k2])*rtA[j3])+float32(float32(iB*rtB[k2])*rtB[j3]))*tangentImpulses[k2])
				goto _79
			_79:
				;
				k2 = k2 + 1
			}
			// clamp the accumulated force
			maxFriction = float32(friction * (*b2ContactConstraintPoint)(unsafe.Pointer(cp2)).NormalImpulse)
			v18 = float32(-(*b2ContactConstraintPoint)(unsafe.Pointer(cp2)).TangentMass * vt)
			v22 = -maxFriction - (*b2ContactConstraintPoint)(unsafe.Pointer(cp2)).TangentImpulse
			v27 = maxFriction - (*b2ContactConstraintPoint)(unsafe.Pointer(cp2)).TangentImpulse
			if v18 < v22 {
				v35 = v22
			} else {
				if v18 > v27 {
					v39 = v27
				} else {
					v39 = v18
				}
				v35 = v39
			}
			v31 = v35
			goto _84
		_84:
			impulse1 = v31
			*(*float32)(unsafe.Add(nil, cp2+28)) += impulse1
			tangentImpulses[j3] = impulse1
			goto _77
		_77:
			;
			j3 = j3 + 1
		}
		normalImpulse = float32FromFloat32(0)
		tangentImpulse = float32FromFloat32(0)
		angularA = rollingImpulse
		angularB = rollingImpulse
		j4 = 0
		for {
			if !(j4 < pointCount) {
				break
			}
			normalImpulse = normalImpulse + normalImpulses[j4]
			tangentImpulse = tangentImpulse + tangentImpulses[j4]
			angularA = angularA + (float32(normalImpulses[j4]*rnA[j4]) + float32(tangentImpulses[j4]*rtA[j4]))
			angularB = angularB + (float32(normalImpulses[j4]*rnB[j4]) + float32(tangentImpulses[j4]*rtB[j4]))
			goto _87
		_87:
			;
			j4 = j4 + 1
		}
		v18 = normalImpulse
		v5 = normal
		v6 = Vec2{
			X: float32(v18 * v5.X),
			Y: float32(v18 * v5.Y),
		}
		goto _91
	_91:
		v22 = tangentImpulse
		v7 = tangent
		v9 = Vec2{
			X: float32(v22 * v7.X),
			Y: float32(v22 * v7.Y),
		}
		goto _95
	_95:
		v10 = v6
		v12 = v9
		v13 = Vec2{
			X: v10.X + v12.X,
			Y: v10.Y + v12.Y,
		}
		goto _99
	_99:
		P = v13
		v14 = vA
		v27 = mA
		v16 = P
		v17 = Vec2{
			X: v14.X - float32(v27*v16.X),
			Y: v14.Y - float32(v27*v16.Y),
		}
		goto _104
	_104:
		vA = v17
		wA = wA - float32(iA*angularA)
		v5 = vB
		v18 = mB
		v6 = P
		v7 = Vec2{
			X: v5.X + float32(v18*v6.X),
			Y: v5.Y + float32(v18*v6.Y),
		}
		goto _109
	_109:
		vB = v7
		wB = wB + float32(iB*angularB)
		if (*b2BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b2_dynamicFlag) != 0 {
			(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
			(*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
		}
		if (*b2BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b2_dynamicFlag) != 0 {
			(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
			(*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
		}
		goto _2
	_2:
		;
		i = i + 1
	}
}

func b2SolveContacts_Wide(tls *_Stack, block b2SolverBlock, context uintptr) {
	b2SolveContacts_WideW4(tls, block, context)
}

// C documentation
//
//	// Solve the normal constraint, friction, and rolling resistance.
func b2SolveContacts_WideW4(tls *_Stack, block b2SolverBlock, context uintptr) {
	bp := tls.Alloc(256)
	defer tls.Free(256)
	var Px, Py, angularA, angularB, deltaLambda, dn, dt, invMass, inv_h, k, k21, k31, k32, k41, k42, k43, lambda, mask, maxFriction1, maxFriction2, maxLambda, normalImpulse, normalImpulse1, normalImpulse2, relativeW, rnA1, rnA2, rnB1, rnB2, rtA1, rtA2, rtB1, rtB2, s, s1, tangentImpulse, tangentImpulse1, tangentImpulse2, velocityBias1, velocityBias2, vn1, vn2, vt1, vt2, wnA1, wnA2, wnB1, wnB2, wtA1, wtA2, wtB1, wtB2, zeroW b2FloatW4
	var c, color, constraints, states uintptr
	var dp, ds, ds1, dv, normal, rsA, rsA1, rsB, rsB1, tangent b2Vec2W4
	var wideIndex int32
	var _ /* bA at bp+0 */ b2BodyStateW4
	var _ /* bB at bp+128 */ b2BodyStateW4
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = Px, Py, angularA, angularB, c, color, constraints, deltaLambda, dn, dp, ds, ds1, dt, dv, invMass, inv_h, k, k21, k31, k32, k41, k42, k43, lambda, mask, maxFriction1, maxFriction2, maxLambda, normal, normalImpulse, normalImpulse1, normalImpulse2, relativeW, rnA1, rnA2, rnB1, rnB2, rsA, rsA1, rsB, rsB1, rtA1, rtA2, rtB1, rtB2, s, s1, states, tangent, tangentImpulse, tangentImpulse1, tangentImpulse2, velocityBias1, velocityBias2, vn1, vn2, vt1, vt2, wideIndex, wnA1, wnA2, wnB1, wnB2, wtA1, wtA2, wtB1, wtB2, zeroW
	states = (*b2StepContext)(unsafe.Pointer(context)).States
	color = (*b2StepContext)(unsafe.Pointer(context)).Graph + uintptr(block.ColorIndex)*64
	constraints = (*b2GraphColor)(unsafe.Pointer(color)).ｆ__ccgo3_48.WideConstraints
	inv_h = b2SplatW(tls, (*b2StepContext)(unsafe.Pointer(context)).Inv_h)
	zeroW = b2ZeroW(tls)
	wideIndex = block.StartIndex
	for {
		if !(wideIndex < block.StartIndex+int32FromUint16(block.Count)) {
			break
		}
		c = constraints + uintptr(wideIndex)*592
		*(*b2BodyStateW4)(unsafe.Add(nil, bp)) = b2GatherBodies(tls, states, c)
		*(*b2BodyStateW4)(unsafe.Add(nil, bp+128)) = b2GatherBodies(tls, states, c+16)
		normal = (*b2ContactConstraintW4)(unsafe.Pointer(c)).Normal
		tangent = b2RightPerpW(tls, normal)
		dp = b2Vec2W4{
			X: b2SubW(tls, (*(*b2BodyStateW4)(unsafe.Add(nil, bp+128))).Dp.X, (*(*b2BodyStateW4)(unsafe.Add(nil, bp))).Dp.X),
			Y: b2SubW(tls, (*(*b2BodyStateW4)(unsafe.Add(nil, bp+128))).Dp.Y, (*(*b2BodyStateW4)(unsafe.Add(nil, bp))).Dp.Y),
		}
		// Moving anchors for current separation
		rsA = b2RotateVectorW(tls, (*(*b2BodyStateW4)(unsafe.Add(nil, bp))).Dq, (*b2ContactConstraintW4)(unsafe.Pointer(c)).AnchorA1)
		rsB = b2RotateVectorW(tls, (*(*b2BodyStateW4)(unsafe.Add(nil, bp+128))).Dq, (*b2ContactConstraintW4)(unsafe.Pointer(c)).AnchorB1)
		// compute current separation
		// this is subject to round-off error if the anchor is far from the body center of mass
		ds = b2Vec2W4{
			X: b2AddW(tls, dp.X, b2SubW(tls, rsB.X, rsA.X)),
			Y: b2AddW(tls, dp.Y, b2SubW(tls, rsB.Y, rsA.Y)),
		}
		s = b2AddW(tls, b2DotW(tls, normal, ds), (*b2ContactConstraintW4)(unsafe.Pointer(c)).BaseSeparation1)
		// Speculative bias, positive if separated and zero if overlapped
		velocityBias1 = b2MaxW(tls, zeroW, b2MulW(tls, s, inv_h))
		rsA1 = b2RotateVectorW(tls, (*(*b2BodyStateW4)(unsafe.Add(nil, bp))).Dq, (*b2ContactConstraintW4)(unsafe.Pointer(c)).AnchorA2)
		rsB1 = b2RotateVectorW(tls, (*(*b2BodyStateW4)(unsafe.Add(nil, bp+128))).Dq, (*b2ContactConstraintW4)(unsafe.Pointer(c)).AnchorB2)
		ds1 = b2Vec2W4{
			X: b2AddW(tls, dp.X, b2SubW(tls, rsB1.X, rsA1.X)),
			Y: b2AddW(tls, dp.Y, b2SubW(tls, rsB1.Y, rsA1.Y)),
		}
		s1 = b2AddW(tls, b2DotW(tls, normal, ds1), (*b2ContactConstraintW4)(unsafe.Pointer(c)).BaseSeparation2)
		velocityBias2 = b2MaxW(tls, zeroW, b2MulW(tls, s1, inv_h))
		rnA1 = b2CrossW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).AnchorA1, normal)
		rnB1 = b2CrossW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).AnchorB1, normal)
		rnA2 = b2CrossW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).AnchorA2, normal)
		rnB2 = b2CrossW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).AnchorB2, normal)
		rtA1 = b2CrossW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).AnchorA1, tangent)
		rtB1 = b2CrossW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).AnchorB1, tangent)
		rtA2 = b2CrossW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).AnchorA2, tangent)
		rtB2 = b2CrossW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).AnchorB2, tangent)
		wnA1 = b2MulW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).InvIA, rnA1)
		wnB1 = b2MulW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).InvIB, rnB1)
		wnA2 = b2MulW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).InvIA, rnA2)
		wnB2 = b2MulW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).InvIB, rnB2)
		wtA1 = b2MulW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).InvIA, rtA1)
		wtB1 = b2MulW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).InvIB, rtB1)
		invMass = b2AddW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).InvMassA, (*b2ContactConstraintW4)(unsafe.Pointer(c)).InvMassB)
		k21 = b2AddW(tls, invMass, b2AddW(tls, b2MulW(tls, wnA1, rnA2), b2MulW(tls, wnB1, rnB2)))
		k31 = b2AddW(tls, b2MulW(tls, wnA1, rtA1), b2MulW(tls, wnB1, rtB1))
		k32 = b2AddW(tls, b2MulW(tls, wnA2, rtA1), b2MulW(tls, wnB2, rtB1))
		k41 = b2AddW(tls, b2MulW(tls, wnA1, rtA2), b2MulW(tls, wnB1, rtB2))
		k42 = b2AddW(tls, b2MulW(tls, wnA2, rtA2), b2MulW(tls, wnB2, rtB2))
		k43 = b2AddW(tls, invMass, b2AddW(tls, b2MulW(tls, wtA1, rtA2), b2MulW(tls, wtB1, rtB2)))
		dv = b2SubVW(tls, (*(*b2BodyStateW4)(unsafe.Add(nil, bp+128))).V, (*(*b2BodyStateW4)(unsafe.Add(nil, bp))).V)
		dn = b2DotW(tls, dv, normal)
		dt = b2SubW(tls, b2DotW(tls, dv, tangent), (*b2ContactConstraintW4)(unsafe.Pointer(c)).TangentSpeed)
		vn1 = b2AddW(tls, dn, b2SubW(tls, b2MulW(tls, (*(*b2BodyStateW4)(unsafe.Add(nil, bp+128))).W, rnB1), b2MulW(tls, (*(*b2BodyStateW4)(unsafe.Add(nil, bp))).W, rnA1)))
		vn2 = b2AddW(tls, dn, b2SubW(tls, b2MulW(tls, (*(*b2BodyStateW4)(unsafe.Add(nil, bp+128))).W, rnB2), b2MulW(tls, (*(*b2BodyStateW4)(unsafe.Add(nil, bp))).W, rnA2)))
		vt1 = b2AddW(tls, dt, b2SubW(tls, b2MulW(tls, (*(*b2BodyStateW4)(unsafe.Add(nil, bp+128))).W, rtB1), b2MulW(tls, (*(*b2BodyStateW4)(unsafe.Add(nil, bp))).W, rtA1)))
		vt2 = b2AddW(tls, dt, b2SubW(tls, b2MulW(tls, (*(*b2BodyStateW4)(unsafe.Add(nil, bp+128))).W, rtB2), b2MulW(tls, (*(*b2BodyStateW4)(unsafe.Add(nil, bp))).W, rtA2)))
		normalImpulse1 = b2MaxW(tls, b2NegW(tls, b2MulW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).NormalMass1, b2AddW(tls, vn1, velocityBias1))), b2NegW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).NormalImpulse1))
		(*b2ContactConstraintW4)(unsafe.Pointer(c)).NormalImpulse1 = b2AddW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).NormalImpulse1, normalImpulse1)
		(*b2ContactConstraintW4)(unsafe.Pointer(c)).TotalNormalImpulse1 = b2AddW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).TotalNormalImpulse1, normalImpulse1)
		vn2 = b2MulAddW(tls, vn2, k21, normalImpulse1)
		normalImpulse2 = b2MaxW(tls, b2NegW(tls, b2MulW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).NormalMass2, b2AddW(tls, vn2, velocityBias2))), b2NegW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).NormalImpulse2))
		(*b2ContactConstraintW4)(unsafe.Pointer(c)).NormalImpulse2 = b2AddW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).NormalImpulse2, normalImpulse2)
		(*b2ContactConstraintW4)(unsafe.Pointer(c)).TotalNormalImpulse2 = b2AddW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).TotalNormalImpulse2, normalImpulse2)
		vt1 = b2MulAddW(tls, b2MulAddW(tls, vt1, k31, normalImpulse1), k32, normalImpulse2)
		vt2 = b2MulAddW(tls, b2MulAddW(tls, vt2, k41, normalImpulse1), k42, normalImpulse2)
		angularA = b2AddW(tls, b2MulW(tls, normalImpulse1, rnA1), b2MulW(tls, normalImpulse2, rnA2))
		angularB = b2AddW(tls, b2MulW(tls, normalImpulse1, rnB1), b2MulW(tls, normalImpulse2, rnB2))
		// Rolling resistance
		if int32FromUint8(b2AllZeroW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).RollingResistance)) == false1 {
			relativeW = b2SubW(tls, b2SubW(tls, (*(*b2BodyStateW4)(unsafe.Add(nil, bp))).W, (*(*b2BodyStateW4)(unsafe.Add(nil, bp+128))).W), b2AddW(tls, b2MulW(tls, b2AddW(tls, wnA1, wnB1), normalImpulse1), b2MulW(tls, b2AddW(tls, wnA2, wnB2), normalImpulse2)))
			k = b2AddW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).InvIA, (*b2ContactConstraintW4)(unsafe.Pointer(c)).InvIB)
			mask = b2GreaterThanW(tls, k, zeroW)
			deltaLambda = b2BlendW(tls, zeroW, b2DivW(tls, relativeW, k), mask)
			lambda = (*b2ContactConstraintW4)(unsafe.Pointer(c)).RollingImpulse
			maxLambda = b2MulW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).RollingResistance, b2AddW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).NormalImpulse1, (*b2ContactConstraintW4)(unsafe.Pointer(c)).NormalImpulse2))
			(*b2ContactConstraintW4)(unsafe.Pointer(c)).RollingImpulse = b2SymClampW(tls, b2AddW(tls, lambda, deltaLambda), maxLambda)
			deltaLambda = b2SubW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).RollingImpulse, lambda)
			wtA2 = b2MulW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).InvIA, rtA2)
			wtB2 = b2MulW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).InvIB, rtB2)
			vt1 = b2MulAddW(tls, vt1, b2AddW(tls, wtA1, wtB1), deltaLambda)
			vt2 = b2MulAddW(tls, vt2, b2AddW(tls, wtA2, wtB2), deltaLambda)
			angularA = b2AddW(tls, angularA, deltaLambda)
			angularB = b2AddW(tls, angularB, deltaLambda)
		}
		maxFriction1 = b2MulW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).Friction, (*b2ContactConstraintW4)(unsafe.Pointer(c)).NormalImpulse1)
		tangentImpulse1 = b2NegW(tls, b2MulW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).TangentMass1, vt1))
		tangentImpulse1 = b2MaxW(tls, b2SubW(tls, b2NegW(tls, maxFriction1), (*b2ContactConstraintW4)(unsafe.Pointer(c)).TangentImpulse1), b2MinW(tls, tangentImpulse1, b2SubW(tls, maxFriction1, (*b2ContactConstraintW4)(unsafe.Pointer(c)).TangentImpulse1)))
		(*b2ContactConstraintW4)(unsafe.Pointer(c)).TangentImpulse1 = b2AddW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).TangentImpulse1, tangentImpulse1)
		vt2 = b2MulAddW(tls, vt2, k43, tangentImpulse1)
		maxFriction2 = b2MulW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).Friction, (*b2ContactConstraintW4)(unsafe.Pointer(c)).NormalImpulse2)
		tangentImpulse2 = b2NegW(tls, b2MulW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).TangentMass2, vt2))
		tangentImpulse2 = b2MaxW(tls, b2SubW(tls, b2NegW(tls, maxFriction2), (*b2ContactConstraintW4)(unsafe.Pointer(c)).TangentImpulse2), b2MinW(tls, tangentImpulse2, b2SubW(tls, maxFriction2, (*b2ContactConstraintW4)(unsafe.Pointer(c)).TangentImpulse2)))
		(*b2ContactConstraintW4)(unsafe.Pointer(c)).TangentImpulse2 = b2AddW(tls, (*b2ContactConstraintW4)(unsafe.Pointer(c)).TangentImpulse2, tangentImpulse2)
		angularA = b2AddW(tls, angularA, b2AddW(tls, b2MulW(tls, tangentImpulse1, rtA1), b2MulW(tls, tangentImpulse2, rtA2)))
		angularB = b2AddW(tls, angularB, b2AddW(tls, b2MulW(tls, tangentImpulse1, rtB1), b2MulW(tls, tangentImpulse2, rtB2)))
		normalImpulse = b2AddW(tls, normalImpulse1, normalImpulse2)
		tangentImpulse = b2AddW(tls, tangentImpulse1, tangentImpulse2)
		Px = b2AddW(tls, b2MulW(tls, normalImpulse, normal.X), b2MulW(tls, tangentImpulse, tangent.X))
		Py = b2AddW(tls, b2MulW(tls, normalImpulse, normal.Y), b2MulW(tls, tangentImpulse, tangent.Y))
		(*(*b2BodyStateW4)(unsafe.Add(nil, bp))).V.X = b2MulSubW(tls, (*(*b2BodyStateW4)(unsafe.Add(nil, bp))).V.X, (*b2ContactConstraintW4)(unsafe.Pointer(c)).InvMassA, Px)
		(*(*b2BodyStateW4)(unsafe.Add(nil, bp))).V.Y = b2MulSubW(tls, (*(*b2BodyStateW4)(unsafe.Add(nil, bp))).V.Y, (*b2ContactConstraintW4)(unsafe.Pointer(c)).InvMassA, Py)
		(*(*b2BodyStateW4)(unsafe.Add(nil, bp))).W = b2MulSubW(tls, (*(*b2BodyStateW4)(unsafe.Add(nil, bp))).W, (*b2ContactConstraintW4)(unsafe.Pointer(c)).InvIA, angularA)
		(*(*b2BodyStateW4)(unsafe.Add(nil, bp+128))).V.X = b2MulAddW(tls, (*(*b2BodyStateW4)(unsafe.Add(nil, bp+128))).V.X, (*b2ContactConstraintW4)(unsafe.Pointer(c)).InvMassB, Px)
		(*(*b2BodyStateW4)(unsafe.Add(nil, bp+128))).V.Y = b2MulAddW(tls, (*(*b2BodyStateW4)(unsafe.Add(nil, bp+128))).V.Y, (*b2ContactConstraintW4)(unsafe.Pointer(c)).InvMassB, Py)
		(*(*b2BodyStateW4)(unsafe.Add(nil, bp+128))).W = b2MulAddW(tls, (*(*b2BodyStateW4)(unsafe.Add(nil, bp+128))).W, (*b2ContactConstraintW4)(unsafe.Pointer(c)).InvIB, angularB)
		b2ScatterBodies(tls, states, c, bp)
		b2ScatterBodies(tls, states, c+16, bp+128)
		goto _1
	_1:
		;
		wideIndex = wideIndex + 1
	}
}

// C documentation
//
//	// Solve a line segment using barycentric coordinates.
//	//
//	// p = a1 * w1 + a2 * w2
//	// a1 + a2 = 1
//	//
//	// The vector from the origin to the closest point on the line is
//	// perpendicular to the line.
//	// e12 = w2 - w1
//	// dot(p, e) = 0
//	// a1 * dot(w1, e) + a2 * dot(w2, e) = 0
//	//
//	// 2-by-2 linear system
//	// [1      1     ][a1] = [1]
//	// [w1.e12 w2.e12][a2] = [0]
//	//
//	// Define
//	// d12_1 =  dot(w2, e12)
//	// d12_2 = -dot(w1, e12)
//	// d12 = d12_1 + d12_2
//	//
//	// Solution
//	// a1 = d12_1 / d12
//	// a2 = d12_2 / d12
//	//
//	// returns a vector that points towards the origin
func b2SolveSimplex2(tls *_Stack, s1 uintptr) (r Vec2) {
	var d12_1, d12_2, inv_d12, v14, v7 float32
	var e12, w1, w2, v1, v10, v2, v3, v5, v6, v9 Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = d12_1, d12_2, e12, inv_d12, w1, w2, v1, v10, v14, v2, v3, v5, v6, v7, v9
	w1 = (*Simplex)(unsafe.Pointer(s1)).V1.W
	w2 = (*Simplex)(unsafe.Pointer(s1)).V2.W
	v1 = w2
	v2 = w1
	v3 = Vec2{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
	goto _4
_4:
	e12 = v3
	v5 = w1
	v6 = e12
	v7 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y)
	goto _8
_8:
	// w1 region
	d12_2 = -v7
	if d12_2 <= float32FromFloat32(0) {
		// a2 <= 0, so we clamp it to 0
		(*Simplex)(unsafe.Pointer(s1)).V1.A = float32FromFloat32(1)
		(*Simplex)(unsafe.Pointer(s1)).Count = int32(1)
		v1 = w1
		v2 = Vec2{
			X: -v1.X,
			Y: -v1.Y,
		}
		goto _11
	_11:
		return v2
	}
	v1 = w2
	v2 = e12
	v7 = float32(v1.X*v2.X) + float32(v1.Y*v2.Y)
	goto _15
_15:
	// w2 region
	d12_1 = v7
	if d12_1 <= float32FromFloat32(0) {
		// a1 <= 0, so we clamp it to 0
		(*Simplex)(unsafe.Pointer(s1)).V2.A = float32FromFloat32(1)
		(*Simplex)(unsafe.Pointer(s1)).Count = int32(1)
		(*Simplex)(unsafe.Pointer(s1)).V1 = (*Simplex)(unsafe.Pointer(s1)).V2
		v1 = w2
		v2 = Vec2{
			X: -v1.X,
			Y: -v1.Y,
		}
		goto _18
	_18:
		return v2
	}
	// Must be in e12 region.
	inv_d12 = float32FromFloat32(1) / (d12_1 + d12_2)
	(*Simplex)(unsafe.Pointer(s1)).V1.A = float32(d12_1 * inv_d12)
	(*Simplex)(unsafe.Pointer(s1)).V2.A = float32(d12_2 * inv_d12)
	(*Simplex)(unsafe.Pointer(s1)).Count = int32(2)
	v1 = w1
	v2 = w2
	v3 = Vec2{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
	goto _22
_22:
	v5 = v3
	v6 = e12
	v7 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	goto _26
_26:
	v14 = v7
	v9 = e12
	v10 = Vec2{
		X: float32(-v14 * v9.Y),
		Y: float32(v14 * v9.X),
	}
	goto _30
_30:
	return v10
}

func b2SolveSimplex3(tls *_Stack, s1 uintptr) (r Vec2) {
	var d123_1, d123_2, d123_3, d12_1, d12_2, d13_1, d13_2, d23_1, d23_2, inv_d12, inv_d123, inv_d13, inv_d23, n123, w1e12, w1e13, w2e12, w2e23, w3e13, w3e23, v11, v19, v23, v31, v35, v39, v43, v47, v51, v7 float32
	var e12, e13, e23, w1, w2, w3, v1, v10, v13, v14, v15, v17, v18, v2, v21, v22, v25, v26, v27, v29, v3, v30, v33, v34, v37, v38, v41, v42, v45, v46, v49, v5, v50, v6, v9 Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = d123_1, d123_2, d123_3, d12_1, d12_2, d13_1, d13_2, d23_1, d23_2, e12, e13, e23, inv_d12, inv_d123, inv_d13, inv_d23, n123, w1, w1e12, w1e13, w2, w2e12, w2e23, w3, w3e13, w3e23, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v29, v3, v30, v31, v33, v34, v35, v37, v38, v39, v41, v42, v43, v45, v46, v47, v49, v5, v50, v51, v6, v7, v9
	w1 = (*Simplex)(unsafe.Pointer(s1)).V1.W
	w2 = (*Simplex)(unsafe.Pointer(s1)).V2.W
	w3 = (*Simplex)(unsafe.Pointer(s1)).V3.W
	v1 = w2
	v2 = w1
	v3 = Vec2{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
	goto _4
_4:
	// Edge12
	// [1      1     ][a1] = [1]
	// [w1.e12 w2.e12][a2] = [0]
	// a3 = 0
	e12 = v3
	v5 = w1
	v6 = e12
	v7 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y)
	goto _8
_8:
	w1e12 = v7
	v9 = w2
	v10 = e12
	v11 = float32(v9.X*v10.X) + float32(v9.Y*v10.Y)
	goto _12
_12:
	w2e12 = v11
	d12_1 = w2e12
	d12_2 = -w1e12
	v13 = w3
	v14 = w1
	v15 = Vec2{
		X: v13.X - v14.X,
		Y: v13.Y - v14.Y,
	}
	goto _16
_16:
	// Edge13
	// [1      1     ][a1] = [1]
	// [w1.e13 w3.e13][a3] = [0]
	// a2 = 0
	e13 = v15
	v17 = w1
	v18 = e13
	v19 = float32(v17.X*v18.X) + float32(v17.Y*v18.Y)
	goto _20
_20:
	w1e13 = v19
	v21 = w3
	v22 = e13
	v23 = float32(v21.X*v22.X) + float32(v21.Y*v22.Y)
	goto _24
_24:
	w3e13 = v23
	d13_1 = w3e13
	d13_2 = -w1e13
	v25 = w3
	v26 = w2
	v27 = Vec2{
		X: v25.X - v26.X,
		Y: v25.Y - v26.Y,
	}
	goto _28
_28:
	// Edge23
	// [1      1     ][a2] = [1]
	// [w2.e23 w3.e23][a3] = [0]
	// a1 = 0
	e23 = v27
	v29 = w2
	v30 = e23
	v31 = float32(v29.X*v30.X) + float32(v29.Y*v30.Y)
	goto _32
_32:
	w2e23 = v31
	v33 = w3
	v34 = e23
	v35 = float32(v33.X*v34.X) + float32(v33.Y*v34.Y)
	goto _36
_36:
	w3e23 = v35
	d23_1 = w3e23
	d23_2 = -w2e23
	v37 = e12
	v38 = e13
	v39 = float32(v37.X*v38.Y) - float32(v37.Y*v38.X)
	goto _40
_40:
	// Triangle123
	n123 = v39
	v41 = w2
	v42 = w3
	v43 = float32(v41.X*v42.Y) - float32(v41.Y*v42.X)
	goto _44
_44:
	d123_1 = float32(n123 * v43)
	v45 = w3
	v46 = w1
	v47 = float32(v45.X*v46.Y) - float32(v45.Y*v46.X)
	goto _48
_48:
	d123_2 = float32(n123 * v47)
	v49 = w1
	v50 = w2
	v51 = float32(v49.X*v50.Y) - float32(v49.Y*v50.X)
	goto _52
_52:
	d123_3 = float32(n123 * v51)
	// w1 region
	if d12_2 <= float32FromFloat32(0) && d13_2 <= float32FromFloat32(0) {
		(*Simplex)(unsafe.Pointer(s1)).V1.A = float32FromFloat32(1)
		(*Simplex)(unsafe.Pointer(s1)).Count = int32(1)
		v1 = w1
		v2 = Vec2{
			X: -v1.X,
			Y: -v1.Y,
		}
		goto _55
	_55:
		return v2
	}
	// e12
	if d12_1 > float32FromFloat32(0) && d12_2 > float32FromFloat32(0) && d123_3 <= float32FromFloat32(0) {
		inv_d12 = float32FromFloat32(1) / (d12_1 + d12_2)
		(*Simplex)(unsafe.Pointer(s1)).V1.A = float32(d12_1 * inv_d12)
		(*Simplex)(unsafe.Pointer(s1)).V2.A = float32(d12_2 * inv_d12)
		(*Simplex)(unsafe.Pointer(s1)).Count = int32(2)
		v1 = w1
		v2 = w2
		v3 = Vec2{
			X: v1.X + v2.X,
			Y: v1.Y + v2.Y,
		}
		goto _59
	_59:
		v5 = v3
		v6 = e12
		v7 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
		goto _63
	_63:
		v11 = v7
		v9 = e12
		v10 = Vec2{
			X: float32(-v11 * v9.Y),
			Y: float32(v11 * v9.X),
		}
		goto _67
	_67:
		return v10
	}
	// e13
	if d13_1 > float32FromFloat32(0) && d13_2 > float32FromFloat32(0) && d123_2 <= float32FromFloat32(0) {
		inv_d13 = float32FromFloat32(1) / (d13_1 + d13_2)
		(*Simplex)(unsafe.Pointer(s1)).V1.A = float32(d13_1 * inv_d13)
		(*Simplex)(unsafe.Pointer(s1)).V3.A = float32(d13_2 * inv_d13)
		(*Simplex)(unsafe.Pointer(s1)).Count = int32(2)
		(*Simplex)(unsafe.Pointer(s1)).V2 = (*Simplex)(unsafe.Pointer(s1)).V3
		v1 = w1
		v2 = w3
		v3 = Vec2{
			X: v1.X + v2.X,
			Y: v1.Y + v2.Y,
		}
		goto _71
	_71:
		v5 = v3
		v6 = e13
		v7 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
		goto _75
	_75:
		v11 = v7
		v9 = e13
		v10 = Vec2{
			X: float32(-v11 * v9.Y),
			Y: float32(v11 * v9.X),
		}
		goto _79
	_79:
		return v10
	}
	// w2 region
	if d12_1 <= float32FromFloat32(0) && d23_2 <= float32FromFloat32(0) {
		(*Simplex)(unsafe.Pointer(s1)).V2.A = float32FromFloat32(1)
		(*Simplex)(unsafe.Pointer(s1)).Count = int32(1)
		(*Simplex)(unsafe.Pointer(s1)).V1 = (*Simplex)(unsafe.Pointer(s1)).V2
		v1 = w2
		v2 = Vec2{
			X: -v1.X,
			Y: -v1.Y,
		}
		goto _82
	_82:
		return v2
	}
	// w3 region
	if d13_1 <= float32FromFloat32(0) && d23_1 <= float32FromFloat32(0) {
		(*Simplex)(unsafe.Pointer(s1)).V3.A = float32FromFloat32(1)
		(*Simplex)(unsafe.Pointer(s1)).Count = int32(1)
		(*Simplex)(unsafe.Pointer(s1)).V1 = (*Simplex)(unsafe.Pointer(s1)).V3
		v1 = w3
		v2 = Vec2{
			X: -v1.X,
			Y: -v1.Y,
		}
		goto _85
	_85:
		return v2
	}
	// e23
	if d23_1 > float32FromFloat32(0) && d23_2 > float32FromFloat32(0) && d123_1 <= float32FromFloat32(0) {
		inv_d23 = float32FromFloat32(1) / (d23_1 + d23_2)
		(*Simplex)(unsafe.Pointer(s1)).V2.A = float32(d23_1 * inv_d23)
		(*Simplex)(unsafe.Pointer(s1)).V3.A = float32(d23_2 * inv_d23)
		(*Simplex)(unsafe.Pointer(s1)).Count = int32(2)
		(*Simplex)(unsafe.Pointer(s1)).V1 = (*Simplex)(unsafe.Pointer(s1)).V3
		v1 = w2
		v2 = w3
		v3 = Vec2{
			X: v1.X + v2.X,
			Y: v1.Y + v2.Y,
		}
		goto _89
	_89:
		v5 = v3
		v6 = e23
		v7 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
		goto _93
	_93:
		v11 = v7
		v9 = e23
		v10 = Vec2{
			X: float32(-v11 * v9.Y),
			Y: float32(v11 * v9.X),
		}
		goto _97
	_97:
		return v10
	}
	// Must be in triangle123
	inv_d123 = float32FromFloat32(1) / (d123_1 + d123_2 + d123_3)
	(*Simplex)(unsafe.Pointer(s1)).V1.A = float32(d123_1 * inv_d123)
	(*Simplex)(unsafe.Pointer(s1)).V2.A = float32(d123_2 * inv_d123)
	(*Simplex)(unsafe.Pointer(s1)).V3.A = float32(d123_3 * inv_d123)
	(*Simplex)(unsafe.Pointer(s1)).Count = int32(3)
	// No search direction
	return b2Vec2_zero
}

/**@}*/

/**
 * @defgroup constants Constants
 * Constants. Some can be configured. I try to include most internal constants here for convenience
 * and clarity. This is why they are not all configurable.
 * @{
 */

/// Used to detect bad values. Positions greater than about 16km will have precision problems in
/// single precision, so 100km as a limit should be fine in all cases. In large world mode the
/// broad-phase starts to have excessive padding at 10,000km, so a value beyond 100,000km is
/// extreme.

/// Maximum parallel workers. Used for some fixed size arrays.

/// Maximum number of tasks queued per world step. b2EnqueueTaskCallbackFcn will never be called
/// more than this per world step. This is related to B2_MAX_WORKERS. With 32 workers, the maximum
/// observed task count is 130. This allows an external task system to use a fixed size array for
/// Box2D tasks, which may help with creating stable user task pointers.

/// Maximum number of colors in the constraint graph. Constraints that cannot find a color are added
/// to the overflow set, which is solved single-threaded. The compound barrel benchmark has minor
/// overflow with 24 colors

/// A small length used as a collision and constraint tolerance. Usually it is chosen to be
/// numerically significant, but visually insignificant. In meters. Normally this is 0.5cm.
/// @warning modifying this can have a significant impact on stability

/// Maximum number of simultaneous worlds that can be allocated

/// Maximum length of the body name. Can be 0 if you don't need names.

/// The maximum rotation of a body per time step. This limit is very large and is used to prevent
/// numerical problems. You shouldn't need to adjust this.
/// @warning increasing this to 0.5f * B2_PI or greater will break continuous collision.

/// Box2D uses limited speculative collision. This reduces jitter. Normally this is 2cm.
/// @warning modifying this can have a significant impact on performance and stability

/// The default contact recycling distance.

/// The default contact recycling world angle threshold. 0.98 ~= 11.5 degrees

/// This is used to fatten AABBs in the dynamic tree. This allows proxies to move by a small amount
/// without triggering a tree adjustment. This is in meters. Normally this is 5cm. This has been
/// tuned on benchmarks and I found that keeping the margin small improves performance by reducing
/// the number of contacts. I used to keep this larger in the past to avoid tree insert and remove
/// churn. But the broad-phase no longer works that way and now uses a two-phase refit and rebuild
/// setup similar to Jolt.
/// @warning modifying this can have a significant impact on performance

/// For small objects the margin is limited to this fraction times the maximum extent

/// The time that a body must be still before it will go to sleep. In seconds.

/// The maximum number of restitution iterations. Needed to avoid a solver overflow.

/** @} */

func b2SolvePlanes(tls *_Stack, targetDelta Vec2, planes uintptr, count int32) (r PlaneSolverResult) {
	var accumulatedPush, push, separation, tolerance, totalPush, v11, v12, v13, v14, v16, v17, v5, v9 float32
	var delta, v18, v7, v8 Vec2
	var i, iteration, planeIndex int32
	var plane1 uintptr
	var v4 Plane
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = accumulatedPush, delta, i, iteration, plane1, planeIndex, push, separation, tolerance, totalPush, v11, v12, v13, v14, v16, v17, v18, v4, v5, v7, v8, v9
	i = 0
	for {
		if !(i < count) {
			break
		}
		(*(*CollisionPlane)(unsafe.Add(nil, planes+uintptr(i)*24))).Push = float32FromFloat32(0)
		goto _1
	_1:
		;
		i = i + 1
	}
	delta = targetDelta
	tolerance = float32(float32FromFloat32(0.005) * b2_lengthUnitsPerMeter)
	iteration = 0
	for {
		if !(iteration < int32(20)) {
			break
		}
		totalPush = float32FromFloat32(0)
		planeIndex = 0
		for {
			if !(planeIndex < count) {
				break
			}
			plane1 = planes + uintptr(planeIndex)*24
			v4 = (*CollisionPlane)(unsafe.Pointer(plane1)).Plane
			v7 = v4.Normal
			v8 = delta
			v9 = float32(v7.X*v8.X) + float32(v7.Y*v8.Y)
			goto _10
		_10:
			v5 = v9 - v4.Offset
			goto _6
		_6:
			// Add slop to prevent jitter
			separation = v5 + float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)
			// if (separation > 0.0f)
			//{
			//	continue;
			// }
			push = -separation
			// Clamp accumulated push
			accumulatedPush = (*CollisionPlane)(unsafe.Pointer(plane1)).Push
			v11 = (*CollisionPlane)(unsafe.Pointer(plane1)).Push + push
			v12 = float32FromFloat32(0)
			v13 = (*CollisionPlane)(unsafe.Pointer(plane1)).PushLimit
			if v11 < v12 {
				v16 = v12
			} else {
				if v11 > v13 {
					v17 = v13
				} else {
					v17 = v11
				}
				v16 = v17
			}
			v14 = v16
			goto _15
		_15:
			(*CollisionPlane)(unsafe.Pointer(plane1)).Push = v14
			push = (*CollisionPlane)(unsafe.Pointer(plane1)).Push - accumulatedPush
			v7 = delta
			v5 = push
			v8 = (*CollisionPlane)(unsafe.Pointer(plane1)).Plane.Normal
			v18 = Vec2{
				X: v7.X + float32(v5*v8.X),
				Y: v7.Y + float32(v5*v8.Y),
			}
			goto _22
		_22:
			delta = v18
			// Track maximum push for convergence
			v5 = push
			if v5 < float32FromInt32(0) {
				v11 = -v5
			} else {
				v11 = v5
			}
			v9 = v11
			goto _25
		_25:
			totalPush = totalPush + v9
			goto _3
		_3:
			;
			planeIndex = planeIndex + 1
		}
		if totalPush < tolerance {
			break
		}
		goto _2
	_2:
		;
		iteration = iteration + 1
	}
	return PlaneSolverResult{
		Delta:          delta,
		IterationCount: iteration,
	}
}

func b2ValidateSolverSets(tls *_Stack, world uintptr) {
	_ = uint64FromInt64(4)
}

// C documentation
//
//	// Solver set: setIndex + 5 POD arrays
func b2SerSolverSet(tls *_Stack, buf uintptr, set uintptr) {
	b2SnapW_I32(tls, buf, (*b2SolverSet)(unsafe.Pointer(set)).SetIndex)
	b2SnapW_I32(tls, buf, (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count)
	if (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count > 0 {
		b2SnapW_Bytes(tls, buf, (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data, (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count*int32FromInt64(96))
	}
	b2SnapW_I32(tls, buf, (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Count)
	if (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Count > 0 {
		b2SnapW_Bytes(tls, buf, (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Data, (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Count*int32FromInt64(32))
	}
	b2SnapW_I32(tls, buf, (*b2SolverSet)(unsafe.Pointer(set)).JointSims.Count)
	if (*b2SolverSet)(unsafe.Pointer(set)).JointSims.Count > 0 {
		b2SnapW_Bytes(tls, buf, (*b2SolverSet)(unsafe.Pointer(set)).JointSims.Data, (*b2SolverSet)(unsafe.Pointer(set)).JointSims.Count*int32FromInt64(252))
	}
	b2SnapW_I32(tls, buf, (*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Count)
	if (*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Count > 0 {
		b2SnapW_Bytes(tls, buf, (*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Data, (*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Count*int32FromInt64(200))
	}
	b2SnapW_I32(tls, buf, (*b2SolverSet)(unsafe.Pointer(set)).IslandSims.Count)
	if (*b2SolverSet)(unsafe.Pointer(set)).IslandSims.Count > 0 {
		b2SnapW_Bytes(tls, buf, (*b2SolverSet)(unsafe.Pointer(set)).IslandSims.Data, (*b2SolverSet)(unsafe.Pointer(set)).IslandSims.Count*int32FromInt64(4))
	}
}

func b2DesSolverSet(tls *_Stack, r uintptr, set uintptr) {
	var cnt, cnt1, cnt2, cnt3, cnt4, elemSize, elemSize1, elemSize2, elemSize3, elemSize4, newSize, newSize1, newSize2, newSize3, newSize4, oldSize, oldSize1, oldSize2, oldSize3, oldSize4 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = cnt, cnt1, cnt2, cnt3, cnt4, elemSize, elemSize1, elemSize2, elemSize3, elemSize4, newSize, newSize1, newSize2, newSize3, newSize4, oldSize, oldSize1, oldSize2, oldSize3, oldSize4
	(*b2SolverSet)(unsafe.Pointer(set)).SetIndex = b2SnapR_I32(tls, r)
	cnt = b2SnapR_I32(tls, r)
	elemSize = int32FromInt64(96)
	if (*b2SnapReader)(unsafe.Pointer(r)).Ok != 0 && int32FromUint8(b2SnapCheckCount(tls, r, cnt, elemSize, elemSize)) == false1 {
		(*b2SnapReader)(unsafe.Pointer(r)).Ok = boolUint8(false1 != 0)
	}
	if (*b2SnapReader)(unsafe.Pointer(r)).Ok != 0 && cnt > 0 {
		if (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Capacity < cnt {
			oldSize = int32FromUint64(uint64FromInt32((*b2SolverSet)(unsafe.Pointer(set)).BodySims.Capacity) * uint64(96))
			newSize = int32FromUint64(uint64FromInt32(cnt) * uint64(96))
			(*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data = b2GrowAlloc(tls, (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data, uint64FromInt32(oldSize), uint64FromInt32(newSize))
			(*b2SolverSet)(unsafe.Pointer(set)).BodySims.Capacity = cnt
		}
		(*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count = cnt
		b2SnapR_Bytes(tls, r, (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data, cnt*elemSize)
	} else {
		if (*b2SnapReader)(unsafe.Pointer(r)).Ok != 0 {
			(*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count = 0
		}
	}
	cnt1 = b2SnapR_I32(tls, r)
	elemSize1 = int32FromInt64(32)
	if (*b2SnapReader)(unsafe.Pointer(r)).Ok != 0 && int32FromUint8(b2SnapCheckCount(tls, r, cnt1, elemSize1, elemSize1)) == false1 {
		(*b2SnapReader)(unsafe.Pointer(r)).Ok = boolUint8(false1 != 0)
	}
	if (*b2SnapReader)(unsafe.Pointer(r)).Ok != 0 && cnt1 > 0 {
		if (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Capacity < cnt1 {
			oldSize1 = int32FromUint64(uint64FromInt32((*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Capacity) * uint64(32))
			newSize1 = int32FromUint64(uint64FromInt32(cnt1) * uint64(32))
			(*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Data = b2GrowAlloc(tls, (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Data, uint64FromInt32(oldSize1), uint64FromInt32(newSize1))
			(*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Capacity = cnt1
		}
		(*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Count = cnt1
		b2SnapR_Bytes(tls, r, (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Data, cnt1*elemSize1)
	} else {
		if (*b2SnapReader)(unsafe.Pointer(r)).Ok != 0 {
			(*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Count = 0
		}
	}
	cnt2 = b2SnapR_I32(tls, r)
	elemSize2 = int32FromInt64(252)
	if (*b2SnapReader)(unsafe.Pointer(r)).Ok != 0 && int32FromUint8(b2SnapCheckCount(tls, r, cnt2, elemSize2, elemSize2)) == false1 {
		(*b2SnapReader)(unsafe.Pointer(r)).Ok = boolUint8(false1 != 0)
	}
	if (*b2SnapReader)(unsafe.Pointer(r)).Ok != 0 && cnt2 > 0 {
		if (*b2SolverSet)(unsafe.Pointer(set)).JointSims.Capacity < cnt2 {
			oldSize2 = int32FromUint64(uint64FromInt32((*b2SolverSet)(unsafe.Pointer(set)).JointSims.Capacity) * uint64(252))
			newSize2 = int32FromUint64(uint64FromInt32(cnt2) * uint64(252))
			(*b2SolverSet)(unsafe.Pointer(set)).JointSims.Data = b2GrowAlloc(tls, (*b2SolverSet)(unsafe.Pointer(set)).JointSims.Data, uint64FromInt32(oldSize2), uint64FromInt32(newSize2))
			(*b2SolverSet)(unsafe.Pointer(set)).JointSims.Capacity = cnt2
		}
		(*b2SolverSet)(unsafe.Pointer(set)).JointSims.Count = cnt2
		b2SnapR_Bytes(tls, r, (*b2SolverSet)(unsafe.Pointer(set)).JointSims.Data, cnt2*elemSize2)
	} else {
		if (*b2SnapReader)(unsafe.Pointer(r)).Ok != 0 {
			(*b2SolverSet)(unsafe.Pointer(set)).JointSims.Count = 0
		}
	}
	cnt3 = b2SnapR_I32(tls, r)
	elemSize3 = int32FromInt64(200)
	if (*b2SnapReader)(unsafe.Pointer(r)).Ok != 0 && int32FromUint8(b2SnapCheckCount(tls, r, cnt3, elemSize3, elemSize3)) == false1 {
		(*b2SnapReader)(unsafe.Pointer(r)).Ok = boolUint8(false1 != 0)
	}
	if (*b2SnapReader)(unsafe.Pointer(r)).Ok != 0 && cnt3 > 0 {
		if (*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Capacity < cnt3 {
			oldSize3 = int32FromUint64(uint64FromInt32((*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Capacity) * uint64(200))
			newSize3 = int32FromUint64(uint64FromInt32(cnt3) * uint64(200))
			(*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Data = b2GrowAlloc(tls, (*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Data, uint64FromInt32(oldSize3), uint64FromInt32(newSize3))
			(*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Capacity = cnt3
		}
		(*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Count = cnt3
		b2SnapR_Bytes(tls, r, (*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Data, cnt3*elemSize3)
	} else {
		if (*b2SnapReader)(unsafe.Pointer(r)).Ok != 0 {
			(*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Count = 0
		}
	}
	cnt4 = b2SnapR_I32(tls, r)
	elemSize4 = int32FromInt64(4)
	if (*b2SnapReader)(unsafe.Pointer(r)).Ok != 0 && int32FromUint8(b2SnapCheckCount(tls, r, cnt4, elemSize4, elemSize4)) == false1 {
		(*b2SnapReader)(unsafe.Pointer(r)).Ok = boolUint8(false1 != 0)
	}
	if (*b2SnapReader)(unsafe.Pointer(r)).Ok != 0 && cnt4 > 0 {
		if (*b2SolverSet)(unsafe.Pointer(set)).IslandSims.Capacity < cnt4 {
			oldSize4 = int32FromUint64(uint64FromInt32((*b2SolverSet)(unsafe.Pointer(set)).IslandSims.Capacity) * uint64(4))
			newSize4 = int32FromUint64(uint64FromInt32(cnt4) * uint64(4))
			(*b2SolverSet)(unsafe.Pointer(set)).IslandSims.Data = b2GrowAlloc(tls, (*b2SolverSet)(unsafe.Pointer(set)).IslandSims.Data, uint64FromInt32(oldSize4), uint64FromInt32(newSize4))
			(*b2SolverSet)(unsafe.Pointer(set)).IslandSims.Capacity = cnt4
		}
		(*b2SolverSet)(unsafe.Pointer(set)).IslandSims.Count = cnt4
		b2SnapR_Bytes(tls, r, (*b2SolverSet)(unsafe.Pointer(set)).IslandSims.Data, cnt4*elemSize4)
	} else {
		if (*b2SnapReader)(unsafe.Pointer(r)).Ok != 0 {
			(*b2SolverSet)(unsafe.Pointer(set)).IslandSims.Count = 0
		}
	}
}

// C documentation
//
//	// Continuous collision of dynamic versus static
func b2SolveContinuous(tls *_Stack, world uintptr, bodySimIndex int32, taskContext uintptr, dt float32) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var aabb1, box1, box2, c, fatAABB, fatAABB1, result, sweptBox, v47, v57, v59 AABB
	var awakeSet, dynamicTree, event, fastBody, fastBodySim, fastBodyState, fastShape, kinematicTree, shape, shape1, shapeFatAABB, shapeFatAABB1, staticTree, v31 uintptr
	var base1, v10, v4, v5, v6, v9 Pos
	var c1, dv, origin1, v3, v11, v16, v17, v19, v20, v21, v24, v7 Vec2
	var elementSize, i, newCapacity, newSize, oldSize, proxyId, shapeId, v127 int32
	var invMag, mag, margin, margin1, omt, speculativeDistance, timeLoss, x, y, v49, v51, v53, v55, v61, v62, v63, v65, v66, v67, v68, v70, v71, v72, v73, v75 float32
	var isBullet, s1, v125 uint8
	var proxyType b2BodyType1
	var q, q2, qn, v15, v23, v83, v84 Rot
	var s2, sweep, v13 Sweep
	var w, v45 WorldTransform
	var xf1, xf2, v33 Transform
	var v1, v2 bool
	var _ /* context at bp+0 */ b2ContinuousContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb1, awakeSet, base1, box1, box2, c, c1, dv, dynamicTree, elementSize, event, fastBody, fastBodySim, fastBodyState, fastShape, fatAABB, fatAABB1, i, invMag, isBullet, kinematicTree, mag, margin, margin1, newCapacity, newSize, oldSize, omt, origin1, proxyId, proxyType, q, q2, qn, result, s1, s2, shape, shape1, shapeFatAABB, shapeFatAABB1, shapeId, speculativeDistance, staticTree, sweep, sweptBox, timeLoss, v3, w, x, xf1, xf2, y, v1, v10, v11, v125, v127, v13, v15, v16, v17, v19, v2, v20, v21, v23, v24, v31, v33, v4, v45, v47, v49, v5, v51, v53, v55, v57, v59, v6, v61, v62, v63, v65, v66, v67, v68, v7, v70, v71, v72, v73, v75, v83, v84, v9
	if v1 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+24283, int32FromInt32(363))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	awakeSet = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
	if v2 = !!(0 <= bodySimIndex && bodySimIndex < (*b2SolverSet)(unsafe.Pointer(awakeSet)).BodySims.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+24457, __ccgo_ts+24283, int32FromInt32(364))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	fastBodySim = (*b2SolverSet)(unsafe.Pointer(awakeSet)).BodySims.Data + uintptr(bodySimIndex)*96
	// Re-center the sweep on the fast body so the TOI and the swept query stay in float precision
	base1 = (*b2BodySim)(unsafe.Pointer(fastBodySim)).Center0
	v31 = fastBodySim
	v4 = base1
	v5 = (*b2BodySim)(unsafe.Pointer(v31)).Center0
	v6 = v4
	v7 = Vec2{
		X: v5.X - v6.X,
		Y: v5.Y - v6.Y,
	}
	goto _8
_8:
	s2.C1 = v7
	v9 = (*b2BodySim)(unsafe.Pointer(v31)).Center
	v10 = v4
	v11 = Vec2{
		X: v9.X - v10.X,
		Y: v9.Y - v10.Y,
	}
	goto _12
_12:
	s2.C2 = v11
	s2.Q1 = (*b2BodySim)(unsafe.Pointer(v31)).Rotation0
	s2.Q2 = (*b2BodySim)(unsafe.Pointer(v31)).Transform.Q
	s2.LocalCenter = (*b2BodySim)(unsafe.Pointer(v31)).LocalCenter
	v13 = s2
	goto _14
_14:
	sweep = v13
	xf1.Q = sweep.Q1
	v15 = sweep.Q1
	v7 = sweep.LocalCenter
	v11 = Vec2{
		X: float32(v15.C*v7.X) - float32(v15.S*v7.Y),
		Y: float32(v15.S*v7.X) + float32(v15.C*v7.Y),
	}
	goto _18
_18:
	v16 = sweep.C1
	v17 = v11
	v19 = Vec2{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
	}
	goto _22
_22:
	xf1.P = v19
	xf2.Q = sweep.Q2
	v15 = sweep.Q2
	v7 = sweep.LocalCenter
	v11 = Vec2{
		X: float32(v15.C*v7.X) - float32(v15.S*v7.Y),
		Y: float32(v15.S*v7.X) + float32(v15.C*v7.Y),
	}
	goto _26
_26:
	v16 = sweep.C2
	v17 = v11
	v19 = Vec2{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
	}
	goto _30
_30:
	xf2.P = v19
	staticTree = world + 40 + uintptr(b2_staticBody)*104
	kinematicTree = world + 40 + uintptr(b2_kinematicBody)*104
	dynamicTree = world + 40 + uintptr(b2_dynamicBody)*104
	if v1 = !!(0 <= (*b2BodySim)(unsafe.Pointer(fastBodySim)).BodyId && (*b2BodySim)(unsafe.Pointer(fastBodySim)).BodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+24375, __ccgo_ts+24283, int32FromInt32(382))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	fastBody = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b2BodySim)(unsafe.Pointer(fastBodySim)).BodyId)*104
	*(*b2ContinuousContext)(unsafe.Add(nil, bp)) = b2ContinuousContext{}
	(*(*b2ContinuousContext)(unsafe.Add(nil, bp))).World = world
	(*(*b2ContinuousContext)(unsafe.Add(nil, bp))).Sweep = sweep
	(*(*b2ContinuousContext)(unsafe.Add(nil, bp))).Base = base1
	(*(*b2ContinuousContext)(unsafe.Add(nil, bp))).FastBodySim = fastBodySim
	(*(*b2ContinuousContext)(unsafe.Add(nil, bp))).Fraction = float32FromFloat32(1)
	isBullet = boolUint8((*b2BodySim)(unsafe.Pointer(fastBodySim)).Flags&uint32(b2_isBullet) != uint32(0))
	shapeId = (*b2Body)(unsafe.Pointer(fastBody)).HeadShapeId
	for shapeId != -int32(1) {
		if v1 = !!(0 <= shapeId && shapeId < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1908, __ccgo_ts+24283, int32FromInt32(396))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		fastShape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*280
		shapeId = (*b2Shape)(unsafe.Pointer(fastShape)).NextShapeId
		(*(*b2ContinuousContext)(unsafe.Add(nil, bp))).FastShape = fastShape
		v33 = xf1
		v7 = (*b2Shape)(unsafe.Pointer(fastShape)).LocalCentroid
		x = float32(v33.Q.C*v7.X) - float32(v33.Q.S*v7.Y) + v33.P.X
		y = float32(v33.Q.S*v7.X) + float32(v33.Q.C*v7.Y) + v33.P.Y
		v11 = Vec2{
			X: x,
			Y: y,
		}
		goto _36
	_36:
		(*(*b2ContinuousContext)(unsafe.Add(nil, bp))).Centroid1 = v11
		v33 = xf2
		v7 = (*b2Shape)(unsafe.Pointer(fastShape)).LocalCentroid
		x = float32(v33.Q.C*v7.X) - float32(v33.Q.S*v7.Y) + v33.P.X
		y = float32(v33.Q.S*v7.X) + float32(v33.Q.C*v7.Y) + v33.P.Y
		v11 = Vec2{
			X: x,
			Y: y,
		}
		goto _40
	_40:
		(*(*b2ContinuousContext)(unsafe.Add(nil, bp))).Centroid2 = v11
		box1 = (*b2Shape)(unsafe.Pointer(fastShape)).Aabb
		v33 = xf2
		v7 = v33.P
		v4 = Pos{
			X: v7.X,
			Y: v7.Y,
		}
		goto _44
	_44:
		w.P = v4
		w.Q = v33.Q
		v45 = w
		goto _46
	_46:
		v47 = b2ComputeShapeAABB(tls, fastShape, v45)
		v5 = base1
		v49 = float32(float64(v5.X) + float64(v47.LowerBound.X))
		goto _50
	_50:
		result.LowerBound.X = v49
		v51 = float32(float64(v5.Y) + float64(v47.LowerBound.Y))
		goto _52
	_52:
		result.LowerBound.Y = v51
		v53 = float32(float64(v5.X) + float64(v47.UpperBound.X))
		goto _54
	_54:
		result.UpperBound.X = v53
		v55 = float32(float64(v5.Y) + float64(v47.UpperBound.Y))
		goto _56
	_56:
		result.UpperBound.Y = v55
		v57 = result
		goto _58
	_58:
		// xf2 is in the base frame, compute the tight box near the origin then lift to world
		box2 = v57
		// Store this to avoid double computation in the case there is no impact event
		(*b2Shape)(unsafe.Pointer(fastShape)).Aabb = box2
		// No continuous collision for sensors (but still need the updated bounds)
		if (*b2Shape)(unsafe.Pointer(fastShape)).SensorIndex != -int32(1) {
			continue
		}
		v47 = box1
		v57 = box2
		v49 = v47.LowerBound.X
		v51 = v57.LowerBound.X
		if v49 < v51 {
			v55 = v49
		} else {
			v55 = v51
		}
		v53 = v55
		goto _64
	_64:
		c.LowerBound.X = v53
		v61 = v47.LowerBound.Y
		v62 = v57.LowerBound.Y
		if v61 < v62 {
			v65 = v61
		} else {
			v65 = v62
		}
		v63 = v65
		goto _69
	_69:
		c.LowerBound.Y = v63
		v66 = v47.UpperBound.X
		v67 = v57.UpperBound.X
		if v66 > v67 {
			v70 = v66
		} else {
			v70 = v67
		}
		v68 = v70
		goto _74
	_74:
		c.UpperBound.X = v68
		v71 = v47.UpperBound.Y
		v72 = v57.UpperBound.Y
		if v71 > v72 {
			v75 = v71
		} else {
			v75 = v72
		}
		v73 = v75
		goto _79
	_79:
		c.UpperBound.Y = v73
		v59 = c
		goto _82
	_82:
		sweptBox = v59
		b2DynamicTree_OverlapAABB(tls, staticTree, sweptBox, uint64FromUint64(0xffffffffffffffff), __ccgo_fp(b2ContinuousQueryCallback), bp)
		if isBullet != 0 {
			b2DynamicTree_OverlapAABB(tls, kinematicTree, sweptBox, uint64FromUint64(0xffffffffffffffff), __ccgo_fp(b2ContinuousQueryCallback), bp)
			b2DynamicTree_OverlapAABB(tls, dynamicTree, sweptBox, uint64FromUint64(0xffffffffffffffff), __ccgo_fp(b2ContinuousQueryCallback), bp)
		}
	}
	speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter))
	if (*(*b2ContinuousContext)(unsafe.Add(nil, bp))).Fraction < float32FromFloat32(1) {
		v15 = sweep.Q1
		v23 = sweep.Q2
		v49 = (*(*b2ContinuousContext)(unsafe.Add(nil, bp))).Fraction
		omt = float32FromFloat32(1) - v49
		q = Rot{
			C: float32(omt*v15.C) + float32(v49*v23.C),
			S: float32(omt*v15.S) + float32(v49*v23.S),
		}
		mag = sqrtf(tls, float32(q.S*q.S)+float32(q.C*q.C))
		if mag > float32FromFloat32(0) {
			v51 = float32FromFloat32(1) / mag
		} else {
			v51 = float32FromFloat32(0)
		}
		invMag = v51
		qn = Rot{
			C: float32(q.C * invMag),
			S: float32(q.S * invMag),
		}
		v83 = qn
		goto _88
	_88:
		// Handle time of impact event
		q2 = v83
		v7 = sweep.C1
		v11 = sweep.C2
		v53 = (*(*b2ContinuousContext)(unsafe.Add(nil, bp))).Fraction
		v16 = Vec2{
			X: float32((float32FromFloat32(1)-v53)*v7.X) + float32(v53*v11.X),
			Y: float32((float32FromFloat32(1)-v53)*v7.Y) + float32(v53*v11.Y),
		}
		goto _93
	_93:
		c1 = v16
		v84 = q2
		v17 = sweep.LocalCenter
		v19 = Vec2{
			X: float32(v84.C*v17.X) - float32(v84.S*v17.Y),
			Y: float32(v84.S*v17.X) + float32(v84.C*v17.Y),
		}
		goto _97
	_97:
		v20 = c1
		v21 = v19
		v24 = Vec2{
			X: v20.X - v21.X,
			Y: v20.Y - v21.Y,
		}
		goto _101
	_101:
		origin1 = v24
		// Advance body, lifting the base frame result back to world
		(*b2BodySim)(unsafe.Pointer(fastBodySim)).Transform.Q = q2
		v4 = base1
		v7 = origin1
		v5 = Pos{
			X: v4.X + v7.X,
			Y: v4.Y + v7.Y,
		}
		goto _105
	_105:
		(*b2BodySim)(unsafe.Pointer(fastBodySim)).Transform.P = v5
		v4 = base1
		v7 = c1
		v5 = Pos{
			X: v4.X + v7.X,
			Y: v4.Y + v7.Y,
		}
		goto _109
	_109:
		(*b2BodySim)(unsafe.Pointer(fastBodySim)).Center = v5
		(*b2BodySim)(unsafe.Pointer(fastBodySim)).Rotation0 = q2
		(*b2BodySim)(unsafe.Pointer(fastBodySim)).Center0 = (*b2BodySim)(unsafe.Pointer(fastBodySim)).Center
		// Warning: writing to the body sim flags means we should not read from other body sim flags in this function.
		*(*uint32_t)(unsafe.Add(nil, fastBodySim+92)) |= uint32(b2_hadTimeOfImpact)
		if v1 = !!(0 <= bodySimIndex && bodySimIndex < (*b2SolverSet)(unsafe.Pointer(awakeSet)).BodyStates.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+24530, __ccgo_ts+24283, int32FromInt32(448))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		// Timeloss means there is a lost gravity contribution. Other forces and torques are ignored for now.
		fastBodyState = (*b2SolverSet)(unsafe.Pointer(awakeSet)).BodyStates.Data + uintptr(bodySimIndex)*32
		v3 = (*b2BodyState)(unsafe.Pointer(fastBodyState)).LinearVelocity
		timeLoss = float32((float32FromFloat32(1) - (*(*b2ContinuousContext)(unsafe.Add(nil, bp))).Fraction) * dt)
		v49 = float32(-timeLoss * (*b2BodySim)(unsafe.Pointer(fastBodySim)).GravityScale)
		v7 = (*b2World)(unsafe.Pointer(world)).Gravity
		v11 = Vec2{
			X: float32(v49 * v7.X),
			Y: float32(v49 * v7.Y),
		}
		goto _114
	_114:
		dv = v11
		if (*b2BodyState)(unsafe.Pointer(fastBodyState)).Flags&uint32(b2_lockLinearX) != 0 {
			v51 = float32FromFloat32(0)
		} else {
			v51 = dv.X
		}
		dv.X = v51
		if (*b2BodyState)(unsafe.Pointer(fastBodyState)).Flags&uint32(b2_lockLinearY) != 0 {
			v49 = float32FromFloat32(0)
		} else {
			v49 = dv.Y
		}
		dv.Y = v49
		v7 = v3
		v11 = dv
		v16 = Vec2{
			X: v7.X + v11.X,
			Y: v7.Y + v11.Y,
		}
		goto _120
	_120:
		(*b2BodyState)(unsafe.Pointer(fastBodyState)).LinearVelocity = v16
		if v1 = !!(0 <= bodySimIndex && bodySimIndex < (*b2World)(unsafe.Pointer(world)).BodyMoveEvents.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+24605, __ccgo_ts+24283, int32FromInt32(457))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		// Update body move event
		event = (*b2World)(unsafe.Pointer(world)).BodyMoveEvents.Data + uintptr(bodySimIndex)*40
		(*BodyMoveEvent)(unsafe.Pointer(event)).Transform = (*b2BodySim)(unsafe.Pointer(fastBodySim)).Transform
		// Prepare AABBs for broad-phase.
		// Even though a body is fast, it may not move much. So the AABB may not need enlargement.
		shapeId = (*b2Body)(unsafe.Pointer(fastBody)).HeadShapeId
		for shapeId != -int32(1) {
			if v1 = !!(0 <= shapeId && shapeId < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
				b2InternalAssert(tls, __ccgo_ts+1908, __ccgo_ts+24283, int32FromInt32(466))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			shape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*280
			// Must recompute aabb at the interpolated transform
			aabb1 = b2ComputeFatShapeAABB(tls, shape, (*b2BodySim)(unsafe.Pointer(fastBodySim)).Transform, speculativeDistance)
			(*b2Shape)(unsafe.Pointer(shape)).Aabb = aabb1
			shapeFatAABB = (*b2World)(unsafe.Pointer(world)).FatAABBs.Data + uintptr(shapeId)*16
			v47 = *(*AABB)(unsafe.Add(nil, shapeFatAABB))
			v57 = aabb1
			s1 = boolUint8(true1 != 0)
			s1 = boolUint8(s1 != 0 && v47.LowerBound.X <= v57.LowerBound.X)
			s1 = boolUint8(s1 != 0 && v47.LowerBound.Y <= v57.LowerBound.Y)
			s1 = boolUint8(s1 != 0 && v57.UpperBound.X <= v47.UpperBound.X)
			s1 = boolUint8(s1 != 0 && v57.UpperBound.Y <= v47.UpperBound.Y)
			v125 = s1
			goto _126
		_126:
			if int32FromUint8(v125) == false1 {
				margin = (*b2Shape)(unsafe.Pointer(shape)).AabbMargin
				fatAABB.LowerBound.X = aabb1.LowerBound.X - margin
				fatAABB.LowerBound.Y = aabb1.LowerBound.Y - margin
				fatAABB.UpperBound.X = aabb1.UpperBound.X + margin
				fatAABB.UpperBound.Y = aabb1.UpperBound.Y + margin
				*(*AABB)(unsafe.Add(nil, shapeFatAABB)) = fatAABB
				// Regular bodies mark the hierarchy as enlarged using atomic operations.
				// Bullets are handled separately at a later stage.
				if int32FromUint8(isBullet) == false1 {
					v127 = (*b2Shape)(unsafe.Pointer(shape)).ProxyKey
					proxyType = v127 & int32FromInt32(3)
					proxyId = v127 >> int32(2)
					b2DynamicTree_MarkProxyMoved(tls, world+40+uintptr(proxyType)*104, proxyId, fatAABB)
				} else {
					*(*uint32_t)(unsafe.Add(nil, fastBodySim+92)) |= uint32(b2_enlargeBulletBounds)
				}
			}
			shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
		}
	} else {
		// No time of impact event
		// Advance body
		(*b2BodySim)(unsafe.Pointer(fastBodySim)).Rotation0 = (*b2BodySim)(unsafe.Pointer(fastBodySim)).Transform.Q
		(*b2BodySim)(unsafe.Pointer(fastBodySim)).Center0 = (*b2BodySim)(unsafe.Pointer(fastBodySim)).Center
		// Prepare AABBs for broad-phase
		shapeId = (*b2Body)(unsafe.Pointer(fastBody)).HeadShapeId
		for shapeId != -int32(1) {
			if v1 = !!(0 <= shapeId && shapeId < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
				b2InternalAssert(tls, __ccgo_ts+1908, __ccgo_ts+24283, int32FromInt32(510))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			shape1 = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*280
			// shape->aabb is still valid from above
			shapeFatAABB1 = (*b2World)(unsafe.Pointer(world)).FatAABBs.Data + uintptr(shapeId)*16
			v47 = *(*AABB)(unsafe.Add(nil, shapeFatAABB1))
			v57 = (*b2Shape)(unsafe.Pointer(shape1)).Aabb
			s1 = boolUint8(true1 != 0)
			s1 = boolUint8(s1 != 0 && v47.LowerBound.X <= v57.LowerBound.X)
			s1 = boolUint8(s1 != 0 && v47.LowerBound.Y <= v57.LowerBound.Y)
			s1 = boolUint8(s1 != 0 && v57.UpperBound.X <= v47.UpperBound.X)
			s1 = boolUint8(s1 != 0 && v57.UpperBound.Y <= v47.UpperBound.Y)
			v125 = s1
			goto _132
		_132:
			if int32FromUint8(v125) == false1 {
				margin1 = (*b2Shape)(unsafe.Pointer(shape1)).AabbMargin
				fatAABB1.LowerBound.X = (*b2Shape)(unsafe.Pointer(shape1)).Aabb.LowerBound.X - margin1
				fatAABB1.LowerBound.Y = (*b2Shape)(unsafe.Pointer(shape1)).Aabb.LowerBound.Y - margin1
				fatAABB1.UpperBound.X = (*b2Shape)(unsafe.Pointer(shape1)).Aabb.UpperBound.X + margin1
				fatAABB1.UpperBound.Y = (*b2Shape)(unsafe.Pointer(shape1)).Aabb.UpperBound.Y + margin1
				*(*AABB)(unsafe.Add(nil, shapeFatAABB1)) = fatAABB1
				if int32FromUint8(isBullet) == false1 {
					v127 = (*b2Shape)(unsafe.Pointer(shape1)).ProxyKey
					proxyType = v127 & int32FromInt32(3)
					proxyId = v127 >> int32(2)
					b2DynamicTree_MarkProxyMoved(tls, world+40+uintptr(proxyType)*104, proxyId, fatAABB1)
				} else {
					*(*uint32_t)(unsafe.Add(nil, fastBodySim+92)) |= uint32(b2_enlargeBulletBounds)
				}
			}
			shapeId = (*b2Shape)(unsafe.Pointer(shape1)).NextShapeId
		}
	}
	// Push sensor hits on the the task context for serial processing.
	i = 0
	for {
		if !(i < (*(*b2ContinuousContext)(unsafe.Add(nil, bp))).SensorCount) {
			break
		}
		// Skip any sensor hits that occurred after a solid hit
		if *(*float32)(unsafe.Add(nil, bp+156+uintptr(i)*4)) < (*(*b2ContinuousContext)(unsafe.Add(nil, bp))).Fraction {
			elementSize = int32(8)
			if (*b2TaskContext)(unsafe.Pointer(taskContext)).SensorHits.Count >= (*b2TaskContext)(unsafe.Pointer(taskContext)).SensorHits.Capacity {
				oldSize = (*b2TaskContext)(unsafe.Pointer(taskContext)).SensorHits.Capacity * elementSize
				if (*b2TaskContext)(unsafe.Pointer(taskContext)).SensorHits.Capacity == 0 {
					v127 = int32(8)
				} else {
					v127 = int32(2) * (*b2TaskContext)(unsafe.Pointer(taskContext)).SensorHits.Capacity
				}
				newCapacity = v127
				newSize = newCapacity * elementSize
				(*b2TaskContext)(unsafe.Pointer(taskContext)).SensorHits.Data = b2GrowAlloc(tls, (*b2TaskContext)(unsafe.Pointer(taskContext)).SensorHits.Data, uint64FromInt32(oldSize), uint64FromInt32(newSize))
				(*b2TaskContext)(unsafe.Pointer(taskContext)).SensorHits.Capacity = newCapacity
			}
			*(*b2SensorHit)(unsafe.Add(nil, (*b2TaskContext)(unsafe.Pointer(taskContext)).SensorHits.Data+uintptr((*b2TaskContext)(unsafe.Pointer(taskContext)).SensorHits.Count)*8)) = *(*b2SensorHit)(unsafe.Add(nil, bp+92+uintptr(i)*8))
			(*b2TaskContext)(unsafe.Pointer(taskContext)).SensorHits.Count += int32(1)
		}
		goto _134
	_134:
		;
		i = i + 1
	}
}

// C documentation
//
//	// Parallel solver task
func b2SolverTask(tls *_Stack, taskContext uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var activeColorCount, bodySyncIndex, colorIndex, colorIndex1, colorIndex2, colorIndex3, graphSyncIndex, iterationStageIndex, iterationStageIndex1, j, j1, j2, previousSyncIndex, restitutionIterations, spinCount, stageIndex, stageIndex1, subStepCount, subStepIndex, syncIndex, workerIndex, v1 int32
	var contactSyncIndex, jointSyncIndex, lastSyncBits, syncBits, syncBits1, v25, v26 uint32_t
	var context, profile, stage, stages, workerContext uintptr
	var useBias, v2 uint8
	var v4 bool
	var _ /* ticks at bp+8 */ uint64_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = activeColorCount, bodySyncIndex, colorIndex, colorIndex1, colorIndex2, colorIndex3, contactSyncIndex, context, graphSyncIndex, iterationStageIndex, iterationStageIndex1, j, j1, j2, jointSyncIndex, lastSyncBits, previousSyncIndex, profile, restitutionIterations, spinCount, stage, stageIndex, stageIndex1, stages, subStepCount, subStepIndex, syncBits, syncBits1, syncIndex, useBias, workerContext, workerIndex, v1, v2, v25, v26, v4
	workerContext = taskContext
	workerIndex = (*b2WorkerContext)(unsafe.Pointer(workerContext)).WorkerIndex
	context = (*b2WorkerContext)(unsafe.Pointer(workerContext)).Context
	activeColorCount = (*b2StepContext)(unsafe.Pointer(context)).ActiveColorCount
	stages = (*b2StepContext)(unsafe.Pointer(context)).Stages
	profile = (*b2StepContext)(unsafe.Pointer(context)).World + 2548
	if workerIndex == 0 {
		// The orchestrator slot is a race. The calling thread of b2World_Step also enters here
		// as worker 0, so progress is guaranteed even if the user's task system schedules tasks
		// out of order, has fewer threads than workerCount, or runs the task synchronously
		// inside enqueueTaskFcn. Whoever wins the CAS becomes the orchestrator; the loser
		// returns and lets the spinner-only path handle workers >0.
		v1 = 0
		*(*int32)(unsafe.Pointer(bp)) = v1
		v2 = __builtin___atomic_compare_exchange_n(tls, context+324, bp, int32(1), boolUint8(false1 != 0), int32(__ATOMIC_SEQ_CST), int32(__ATOMIC_SEQ_CST))
		goto _3
	_3:
		if int32FromUint8(v2) == false1 {
			return
		}
		// Main thread synchronizes the workers and does work itself.
		//
		// This single task is able to fully complete all work even if all other workers are
		// blocked, so a fully serial task system still drives the simulation forward.
		// Stages are re-used by loops so that I don't need more stages for large substep counts.
		// The sync indices grow monotonically for the body/graph/constraint groupings because they share solver blocks.
		// The stage index and sync indices are combined in to sync bits for atomic synchronization.
		// The workers need to compute the previous sync index for a given stage so that CAS works correctly. This
		// setup makes this easy to do.
		/*
			Stage sequence
			b2_stagePrepareJoints,
			b2_stagePrepareContacts,
			b2_stageIntegrateVelocities,
			b2_stageWarmStart,
			b2_stageSolve,
			b2_stageIntegratePositions,
			b2_stageRelax,
			b2_stageRestitution,
			b2_stageStoreImpulses
		*/
		*(*uint64_t)(unsafe.Add(nil, bp+8)) = b2GetTicks(tls)
		bodySyncIndex = int32(1)
		stageIndex = 0
		// Prepare joint constraints
		jointSyncIndex = uint32(1)
		syncBits = jointSyncIndex<<int32FromInt32(16) | uint32FromInt32(stageIndex)
		if v4 = !!((*(*b2SolverStage)(unsafe.Add(nil, stages+uintptr(stageIndex)*152))).Type1 == int32(b2_stagePrepareJoints)); !v4 {
			b2InternalAssert(tls, __ccgo_ts+25047, __ccgo_ts+24283, int32FromInt32(1071))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		b2ExecuteMainStage(tls, stages+uintptr(stageIndex)*152, context, syncBits)
		stageIndex = stageIndex + int32(1)
		jointSyncIndex = jointSyncIndex + uint32(1)
		// Prepare contact constraints
		contactSyncIndex = uint32(1)
		syncBits = contactSyncIndex<<int32FromInt32(16) | uint32FromInt32(stageIndex)
		if v4 = !!((*(*b2SolverStage)(unsafe.Add(nil, stages+uintptr(stageIndex)*152))).Type1 == int32(b2_stagePrepareContacts)); !v4 {
			b2InternalAssert(tls, __ccgo_ts+25096, __ccgo_ts+24283, int32FromInt32(1079))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		b2ExecuteMainStage(tls, stages+uintptr(stageIndex)*152, context, syncBits)
		stageIndex = stageIndex + int32(1)
		contactSyncIndex = contactSyncIndex + uint32(1)
		// Single-threaded overflow work. These constraints don't fit in the graph coloring.
		b2PrepareJoints_Overflow(tls, context)
		b2PrepareContacts_Overflow(tls, context)
		*(*float32)(unsafe.Add(nil, profile+24)) += b2GetMillisecondsAndReset(tls, bp+8)
		graphSyncIndex = int32(1)
		subStepCount = (*b2StepContext)(unsafe.Pointer(context)).SubStepCount
		subStepIndex = 0
		for {
			if !(subStepIndex < subStepCount) {
				break
			}
			// stage index restarted each iteration
			// syncBits still increases monotonically because the upper bits increase each iteration
			iterationStageIndex = stageIndex
			// Integrate velocities
			syncBits = uint32FromInt32(bodySyncIndex<<int32(16) | iterationStageIndex)
			if v4 = !!((*(*b2SolverStage)(unsafe.Add(nil, stages+uintptr(iterationStageIndex)*152))).Type1 == int32(b2_stageIntegrateVelocities)); !v4 {
				b2InternalAssert(tls, __ccgo_ts+25147, __ccgo_ts+24283, int32FromInt32(1100))
			}
			_ = v4 || bool(int32FromInt32(0) != 0)
			b2ExecuteMainStage(tls, stages+uintptr(iterationStageIndex)*152, context, syncBits)
			iterationStageIndex = iterationStageIndex + int32(1)
			bodySyncIndex = bodySyncIndex + int32(1)
			*(*float32)(unsafe.Add(nil, profile+28)) += b2GetMillisecondsAndReset(tls, bp+8)
			// Warm start constraints
			b2WarmStartJoints_Overflow(tls, context)
			b2WarmStartContacts_Overflow(tls, context)
			colorIndex = 0
			for {
				if !(colorIndex < activeColorCount) {
					break
				}
				syncBits = uint32FromInt32(graphSyncIndex<<int32(16) | iterationStageIndex)
				if v4 = !!((*(*b2SolverStage)(unsafe.Add(nil, stages+uintptr(iterationStageIndex)*152))).Type1 == int32(b2_stageWarmStart)); !v4 {
					b2InternalAssert(tls, __ccgo_ts+25211, __ccgo_ts+24283, int32FromInt32(1114))
				}
				_ = v4 || bool(int32FromInt32(0) != 0)
				b2ExecuteMainStage(tls, stages+uintptr(iterationStageIndex)*152, context, syncBits)
				iterationStageIndex = iterationStageIndex + int32(1)
				goto _8
			_8:
				;
				colorIndex = colorIndex + 1
			}
			graphSyncIndex = graphSyncIndex + int32(1)
			*(*float32)(unsafe.Add(nil, profile+32)) += b2GetMillisecondsAndReset(tls, bp+8)
			// Solve constraints
			useBias = boolUint8(true1 != 0)
			j = 0
			for {
				if !(j < int32(_ITERATIONS)) {
					break
				}
				// Overflow constraints have lower priority. Typically these are dynamic-vs-dynamic.
				b2SolveJoints_Overflow(tls, context, useBias)
				b2PushContacts_Overflow(tls, context)
				colorIndex1 = 0
				for {
					if !(colorIndex1 < activeColorCount) {
						break
					}
					syncBits = uint32FromInt32(graphSyncIndex<<int32(16) | iterationStageIndex)
					if v4 = !!((*(*b2SolverStage)(unsafe.Add(nil, stages+uintptr(iterationStageIndex)*152))).Type1 == int32(b2_stageSolve)); !v4 {
						b2InternalAssert(tls, __ccgo_ts+25265, __ccgo_ts+24283, int32FromInt32(1133))
					}
					_ = v4 || bool(int32FromInt32(0) != 0)
					b2ExecuteMainStage(tls, stages+uintptr(iterationStageIndex)*152, context, syncBits)
					iterationStageIndex = iterationStageIndex + int32(1)
					goto _11
				_11:
					;
					colorIndex1 = colorIndex1 + 1
				}
				graphSyncIndex = graphSyncIndex + int32(1)
				goto _10
			_10:
				;
				j = j + 1
			}
			*(*float32)(unsafe.Add(nil, profile+36)) += b2GetMillisecondsAndReset(tls, bp+8)
			// Integrate positions
			if v4 = !!((*(*b2SolverStage)(unsafe.Add(nil, stages+uintptr(iterationStageIndex)*152))).Type1 == int32(b2_stageIntegratePositions)); !v4 {
				b2InternalAssert(tls, __ccgo_ts+25315, __ccgo_ts+24283, int32FromInt32(1143))
			}
			_ = v4 || bool(int32FromInt32(0) != 0)
			syncBits = uint32FromInt32(bodySyncIndex<<int32(16) | iterationStageIndex)
			b2ExecuteMainStage(tls, stages+uintptr(iterationStageIndex)*152, context, syncBits)
			iterationStageIndex = iterationStageIndex + int32(1)
			bodySyncIndex = bodySyncIndex + int32(1)
			*(*float32)(unsafe.Add(nil, profile+40)) += b2GetMillisecondsAndReset(tls, bp+8)
			// Relax constraints
			useBias = boolUint8(false1 != 0)
			j1 = 0
			for {
				if !(j1 < int32(_RELAX_ITERATIONS)) {
					break
				}
				b2SolveJoints_Overflow(tls, context, useBias)
				b2SolveContacts_Overflow(tls, context)
				colorIndex2 = 0
				for {
					if !(colorIndex2 < activeColorCount) {
						break
					}
					syncBits = uint32FromInt32(graphSyncIndex<<int32(16) | iterationStageIndex)
					if v4 = !!((*(*b2SolverStage)(unsafe.Add(nil, stages+uintptr(iterationStageIndex)*152))).Type1 == int32(b2_stageRelax)); !v4 {
						b2InternalAssert(tls, __ccgo_ts+25378, __ccgo_ts+24283, int32FromInt32(1160))
					}
					_ = v4 || bool(int32FromInt32(0) != 0)
					b2ExecuteMainStage(tls, stages+uintptr(iterationStageIndex)*152, context, syncBits)
					iterationStageIndex = iterationStageIndex + int32(1)
					goto _15
				_15:
					;
					colorIndex2 = colorIndex2 + 1
				}
				graphSyncIndex = graphSyncIndex + int32(1)
				goto _14
			_14:
				;
				j1 = j1 + 1
			}
			*(*float32)(unsafe.Add(nil, profile+44)) += b2GetMillisecondsAndReset(tls, bp+8)
			goto _6
		_6:
			;
			subStepIndex = subStepIndex + 1
		}
		// Advance the stage according to the sub-stepping tasks just completed
		// integrate velocities / warm start / solve / integrate positions / relax
		stageIndex = stageIndex + (int32(1) + activeColorCount + int32(_ITERATIONS)*activeColorCount + int32(1) + int32(_RELAX_ITERATIONS)*activeColorCount)
		restitutionIterations = (*b2World)(unsafe.Pointer((*b2StepContext)(unsafe.Pointer(context)).World)).RestitutionIterations
		if v4 = restitutionIterations > 0; v4 {
			v1 = atomicLoadNInt32(context+116, int32FromInt32(__ATOMIC_SEQ_CST))
			goto _18
		_18:
		}
		if v4 && v1 != 0 {
			j2 = 0
			for {
				if !(j2 < restitutionIterations) {
					break
				}
				b2ApplyRestitution_Overflow(tls, context)
				iterationStageIndex1 = stageIndex
				colorIndex3 = 0
				for {
					if !(colorIndex3 < activeColorCount) {
						break
					}
					syncBits = uint32FromInt32(graphSyncIndex<<int32(16) | iterationStageIndex1)
					if v4 = !!((*(*b2SolverStage)(unsafe.Add(nil, stages+uintptr(iterationStageIndex1)*152))).Type1 == int32(b2_stageRestitution)); !v4 {
						b2InternalAssert(tls, __ccgo_ts+25428, __ccgo_ts+24283, int32FromInt32(1185))
					}
					_ = v4 || bool(int32FromInt32(0) != 0)
					b2ExecuteMainStage(tls, stages+uintptr(iterationStageIndex1)*152, context, syncBits)
					iterationStageIndex1 = iterationStageIndex1 + int32(1)
					goto _21
				_21:
					;
					colorIndex3 = colorIndex3 + 1
				}
				graphSyncIndex = graphSyncIndex + int32(1)
				goto _20
			_20:
				;
				j2 = j2 + 1
			}
			*(*float32)(unsafe.Add(nil, profile+48)) += b2GetMillisecondsAndReset(tls, bp+8)
		}
		stageIndex = stageIndex + activeColorCount
		// Store impulses
		b2StoreImpulses_Overflow(tls, context)
		syncBits = contactSyncIndex<<int32FromInt32(16) | uint32FromInt32(stageIndex)
		if v4 = !!((*(*b2SolverStage)(unsafe.Add(nil, stages+uintptr(stageIndex)*152))).Type1 == int32(b2_stageStoreImpulses)); !v4 {
			b2InternalAssert(tls, __ccgo_ts+25484, __ccgo_ts+24283, int32FromInt32(1201))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		b2ExecuteMainStage(tls, stages+uintptr(stageIndex)*152, context, syncBits)
		*(*float32)(unsafe.Add(nil, profile+52)) += b2GetMillisecondsAndReset(tls, bp+8)
		// Signal workers to finish
		atomicStoreNUint32(context+256, uint32(0xffffffff), int32FromInt32(__ATOMIC_SEQ_CST))
		if v4 = !!(stageIndex+int32FromInt32(1) == (*b2StepContext)(unsafe.Pointer(context)).StageCount); !v4 {
			b2InternalAssert(tls, __ccgo_ts+25533, __ccgo_ts+24283, int32FromInt32(1209))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		return
	}
	// Worker spins and waits for work
	lastSyncBits = uint32(0)
	// uint64_t maxSpinTime = 10;
	for int32(true1) != 0 {
		spinCount = 0
		for {
			v26 = atomicLoadNUint32(context+256, int32FromInt32(__ATOMIC_SEQ_CST))
			goto _27
		_27:
			v25 = v26
			syncBits1 = v25
			if !(v25 == lastSyncBits) {
				break
			}
			if spinCount > int32(5) {
				b2Yield(tls)
				spinCount = 0
			} else {
				// Using the cycle counter helps to account for variation in mm_pause timing across different
				// CPUs. However, this is X64 only.
				// uint64_t prev = __rdtsc();
				// do
				//{
				//	_mm_pause();
				//}
				// while ((__rdtsc() - prev) < maxSpinTime);
				// maxSpinTime += 10;
				__builtin_Gosched(tls)
				__builtin_Gosched(tls)
				spinCount = spinCount + int32(1)
			}
		}
		if syncBits1 == uint32(0xffffffff) {
			// sentinel hit
			break
		}
		stageIndex1 = int32FromUint32(syncBits1 & uint32(0xFFFF))
		if v4 = !!(stageIndex1 < (*b2StepContext)(unsafe.Pointer(context)).StageCount); !v4 {
			b2InternalAssert(tls, __ccgo_ts+25571, __ccgo_ts+24283, int32FromInt32(1255))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		syncIndex = int32FromUint32(syncBits1 >> int32FromInt32(16) & uint32(0xFFFF))
		if v4 = !!(syncIndex > int32FromInt32(0)); !v4 {
			b2InternalAssert(tls, __ccgo_ts+25033, __ccgo_ts+24283, int32FromInt32(1258))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		previousSyncIndex = syncIndex - int32(1)
		stage = stages + uintptr(stageIndex1)*152
		b2ExecuteStage(tls, stage, context, previousSyncIndex, syncIndex, workerIndex)
		lastSyncBits = syncBits1
	}
}

// C documentation
//
//	// Solve with graph coloring
func b2Solve(tls *_Stack, world1 uintptr, stepContext uintptr) {
	bp1 := tls.Alloc(2000)
	defer tls.Free(2000)
	var activeColorCount, awakeBodyCount, awakeIslandCount, bodyId, bulletBodyCount, c, colorContactCount, colorContactCount1, colorContactCountW, colorContactCountW1, colorJointCount, contactId, contactIdCapacity, count, elementSize, elementSize1, elementSize2, graphBlockCount, hitCount, i, i1, i10, i11, i12, i13, i2, i3, i4, i5, i6, i7, i8, i9, islandId, islandIndex, j, j1, jointBase, jointCount, jointId, jointIdCapacity, maxBlockCount, minBodiesPerBlock, minContactsPerBlock, minJointsPerBlock, minRange, newCapacity, newCapacity1, newCapacity2, newSize, newSize1, newSize2, newSize3, occupancyCount, oldSize, oldSize1, oldSize2, oldSize3, overflowCount, p1, perColorContactCount, perColorJointCount, pointCount, proxyId, proxyKey, shapeId, simdShift, stageCount, wideBase, wideContactConstraintByteCount, wideContactCount, workerCount, workerCount1, v1, v4 int32
	var anyHitEvents, s, v68 uint8
	var approachSpeed, splitSleepTimer, threshold float32
	var awakeIslandBitSet, awakeSet, baseGraphBlock, bestPoint, bits, bits1, bodyA, bodyArray, bodyB, bodyBlocks, bodySim, bodySimA, bodySimArray, bodySimB, bp, broadPhase, bulletBody, bulletBodySim, bulletBodySimIndices, color, color1, colors, colors1, contact, contactArray, contactBlocks, contactSim, dynamicTree, graph, graphBlocks, hitEventBitSet, hits, island, islands, joint, jointArray, jointBlocks, jointStateBitSet, mp, overflow, overflowContacts, sensor, sensorShape, set, shape, shapeA, shapeArray, shapeArray1, shapeB, splitIslandTask, stage, stages, taskContext, taskContext1, taskContext2, taskContext3, visitor, wideContactConstraints, v41, v42, v45 uintptr
	var blockIndex, ctz, ctz1, k, k1, wordCount, wordCount1, v26 uint32_t
	var bodyDim, contactPrepareDim, jointPrepareDim b2BlockDim
	var bulletTicks, hitTicks, jointEventTicks, refitTicks, sensorHitTicks, sleepTicks, transformTicks, word, word1 uint64_t
	var colorContactCounts, colorJointCounts [24]int32
	var colorIndex uint8_t
	var event JointEvent
	var event1 ContactHitEvent
	var graphContactDims, graphJointDims [24]b2BlockDim
	var hit b2SensorHit
	var shapeFatAABB, treeAABB, v66, v67 AABB
	var shapeRef b2Visitor
	var worldId, worldIndex0 uint16_t
	var v12, v2, v8, v9 bool
	var v47, v49 Pos
	var v48 Vec2
	var _ /* activeColorIndices at bp+8 */ [24]int32
	var _ /* callerContext at bp+1968 */ b2WorkerContext
	var _ /* constraintTicks at bp+1960 */ uint64_t
	var _ /* contactPrepareSpans at bp+104 */ [25]b2ContactPrepareSpan
	var _ /* graphBlockCounts at bp+1096 */ [24]int32
	var _ /* graphColorBlocks at bp+904 */ [24]uintptr
	var _ /* jointPrepareSpans at bp+504 */ [25]b2JointPrepareSpan
	var _ /* setupTicks at bp+0 */ uint64_t
	var _ /* workerContext at bp+1192 */ [32]b2WorkerContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = activeColorCount, anyHitEvents, approachSpeed, awakeBodyCount, awakeIslandBitSet, awakeIslandCount, awakeSet, baseGraphBlock, bestPoint, bits, bits1, blockIndex, bodyA, bodyArray, bodyB, bodyBlocks, bodyDim, bodyId, bodySim, bodySimA, bodySimArray, bodySimB, bp, broadPhase, bulletBody, bulletBodyCount, bulletBodySim, bulletBodySimIndices, bulletTicks, c, color, color1, colorContactCount, colorContactCount1, colorContactCountW, colorContactCountW1, colorContactCounts, colorIndex, colorJointCount, colorJointCounts, colors, colors1, contact, contactArray, contactBlocks, contactId, contactIdCapacity, contactPrepareDim, contactSim, count, ctz, ctz1, dynamicTree, elementSize, elementSize1, elementSize2, event, event1, graph, graphBlockCount, graphBlocks, graphContactDims, graphJointDims, hit, hitCount, hitEventBitSet, hitTicks, hits, i, i1, i10, i11, i12, i13, i2, i3, i4, i5, i6, i7, i8, i9, island, islandId, islandIndex, islands, j, j1, joint, jointArray, jointBase, jointBlocks, jointCount, jointEventTicks, jointId, jointIdCapacity, jointPrepareDim, jointStateBitSet, k, k1, maxBlockCount, minBodiesPerBlock, minContactsPerBlock, minJointsPerBlock, minRange, mp, newCapacity, newCapacity1, newCapacity2, newSize, newSize1, newSize2, newSize3, occupancyCount, oldSize, oldSize1, oldSize2, oldSize3, overflow, overflowContacts, overflowCount, p1, perColorContactCount, perColorJointCount, pointCount, proxyId, proxyKey, refitTicks, s, sensor, sensorHitTicks, sensorShape, set, shape, shapeA, shapeArray, shapeArray1, shapeB, shapeFatAABB, shapeId, shapeRef, simdShift, sleepTicks, splitIslandTask, splitSleepTimer, stage, stageCount, stages, taskContext, taskContext1, taskContext2, taskContext3, threshold, transformTicks, treeAABB, visitor, wideBase, wideContactConstraintByteCount, wideContactConstraints, wideContactCount, word, word1, wordCount, wordCount1, workerCount, workerCount1, worldId, worldIndex0, v1, v12, v2, v26, v4, v41, v42, v45, v47, v48, v49, v66, v67, v68, v8, v9
	if (*b2World)(unsafe.Pointer(world1)).SimdWidth == int32(8) {
		v1 = int32(3)
	} else {
		v1 = int32(2)
	}
	simdShift = v1
	// Only count steps that advance the simulation
	*(*uint64_t)(unsafe.Add(nil, world1+2472)) += uint64(1)
	if v2 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(world1)).SolverSets.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+24283, int32FromInt32(1296))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	// Are there any awake bodies? This scenario should not be important for profiling.
	awakeSet = (*b2World)(unsafe.Pointer(world1)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
	awakeBodyCount = (*b2SolverSet)(unsafe.Pointer(awakeSet)).BodySims.Count
	if awakeBodyCount == 0 {
		return
	}
	// Solve constraints using graph coloring
	*(*uint64_t)(unsafe.Add(nil, bp1)) = b2GetTicks(tls)
	// Prepare buffers for bullets
	atomicStoreNInt32(stepContext+112, 0, int32FromInt32(__ATOMIC_SEQ_CST))
	(*b2StepContext)(unsafe.Pointer(stepContext)).BulletBodies = b2StackAlloc(tls, world1, int32FromUint64(uint64FromInt32(awakeBodyCount)*uint64(4)), __ccgo_ts+25627)
	graph = world1 + 376
	colors = graph
	(*b2StepContext)(unsafe.Pointer(stepContext)).Sims = (*b2SolverSet)(unsafe.Pointer(awakeSet)).BodySims.Data
	(*b2StepContext)(unsafe.Pointer(stepContext)).States = (*b2SolverSet)(unsafe.Pointer(awakeSet)).BodyStates.Data
	// count contacts, joints, and colors
	activeColorCount = 0
	i = 0
	for {
		if !(i < int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1)) {
			break
		}
		perColorContactCount = (*(*b2GraphColor)(unsafe.Add(nil, colors+uintptr(i)*64))).ContactSims.Count
		perColorJointCount = (*(*b2GraphColor)(unsafe.Add(nil, colors+uintptr(i)*64))).JointSims.Count
		occupancyCount = perColorContactCount + perColorJointCount
		if occupancyCount > 0 {
			v1 = int32(1)
		} else {
			v1 = 0
		}
		activeColorCount = activeColorCount + v1
		goto _3
	_3:
		;
		i = i + 1
	}
	// prepare for move events
	if (*b2World)(unsafe.Pointer(world1)).BodyMoveEvents.Capacity < awakeBodyCount {
		oldSize = int32FromUint64(uint64FromInt32((*b2World)(unsafe.Pointer(world1)).BodyMoveEvents.Capacity) * uint64(40))
		newSize = int32FromUint64(uint64FromInt32(awakeBodyCount) * uint64(40))
		(*b2World)(unsafe.Pointer(world1)).BodyMoveEvents.Data = b2GrowAlloc(tls, (*b2World)(unsafe.Pointer(world1)).BodyMoveEvents.Data, uint64FromInt32(oldSize), uint64FromInt32(newSize))
		(*b2World)(unsafe.Pointer(world1)).BodyMoveEvents.Capacity = awakeBodyCount
	}
	(*b2World)(unsafe.Pointer(world1)).BodyMoveEvents.Count = awakeBodyCount
	workerCount = (*b2World)(unsafe.Pointer(world1)).WorkerCount
	// Target 4 blocks per worker to allow work stealing
	maxBlockCount = int32(4) * workerCount
	// Body blocks are for parallel iteration over bodies directly (integration, update transforms)
	minBodiesPerBlock = int32(32)
	bodyDim = b2ComputeBlockCount(tls, awakeBodyCount, minBodiesPerBlock, maxBlockCount)
	minContactsPerBlock = int32(4)
	minJointsPerBlock = int32(4)
	graphBlockCount = 0
	// c is the active color index
	wideContactCount = 0
	jointCount = 0
	c = 0
	i1 = 0
	for {
		if !(i1 < int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1)) {
			break
		}
		colorContactCount = (*(*b2GraphColor)(unsafe.Add(nil, colors+uintptr(i1)*64))).ContactSims.Count
		colorJointCount = (*(*b2GraphColor)(unsafe.Add(nil, colors+uintptr(i1)*64))).JointSims.Count
		if colorContactCount+colorJointCount == 0 {
			goto _5
		}
		(*(*[24]int32)(unsafe.Add(nil, bp1+8)))[c] = i1
		if colorContactCount > 0 {
			v1 = (colorContactCount-int32(1))>>simdShift + int32(1)
		} else {
			v1 = 0
		}
		// Ceiling for wide constraint count
		colorContactCountW = v1
		wideContactCount = wideContactCount + colorContactCountW
		colorContactCounts[c] = colorContactCountW
		colorJointCounts[c] = colorJointCount
		jointCount = jointCount + colorJointCount
		// Graph solver block dimensions
		graphContactDims[c] = b2ComputeBlockCount(tls, colorContactCountW, minContactsPerBlock, maxBlockCount)
		graphJointDims[c] = b2ComputeBlockCount(tls, colorJointCount, minJointsPerBlock, maxBlockCount)
		graphBlockCount = graphBlockCount + (graphContactDims[c].Count + graphJointDims[c].Count)
		c = c + int32(1)
		goto _5
	_5:
		;
		i1 = i1 + 1
	}
	activeColorCount = c
	// Prepare and store run as one flat parallel-for over the entire wide constraint range,
	// partitioned into uniformly sized blocks. Color info is consulted inside the task via
	// a small span array, so blocks do not need to honor color boundaries here.
	contactPrepareDim = b2ComputeBlockCount(tls, wideContactCount, minContactsPerBlock, maxBlockCount)
	jointPrepareDim = b2ComputeBlockCount(tls, jointCount, minJointsPerBlock, maxBlockCount)
	wideContactConstraintByteCount = b2GetWideContactConstraintByteCount(tls, (*b2World)(unsafe.Pointer(world1)).SimdWidth)
	wideContactConstraints = b2StackAlloc(tls, world1, wideContactCount*wideContactConstraintByteCount, __ccgo_ts+25641)
	overflow = colors + uintptr(int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1))*64
	overflowCount = (*b2GraphColor)(unsafe.Pointer(overflow)).ContactSims.Count
	overflowContacts = b2StackAlloc(tls, world1, int32FromUint64(uint64FromInt32(overflowCount)*uint64(168)), __ccgo_ts+25660)
	*(*uintptr)(unsafe.Add(unsafe.Pointer(overflow), 48)) = overflowContacts
	// Distribute transient constraints to each graph color and prepare spans
	wideBase = 0
	jointBase = 0
	i2 = 0
	for {
		if !(i2 < activeColorCount) {
			break
		}
		j = (*(*[24]int32)(unsafe.Add(nil, bp1+8)))[i2]
		color = colors + uintptr(j)*64
		colorContactCount1 = (*b2GraphColor)(unsafe.Pointer(color)).ContactSims.Count
		(*(*[25]b2ContactPrepareSpan)(unsafe.Add(nil, bp1+104)))[i2].Start = wideBase
		(*(*[25]b2ContactPrepareSpan)(unsafe.Add(nil, bp1+104)))[i2].Count = colorContactCount1
		(*(*[25]b2ContactPrepareSpan)(unsafe.Add(nil, bp1+104)))[i2].Contacts = (*b2GraphColor)(unsafe.Pointer(color)).ContactSims.Data
		if colorContactCount1 == 0 {
			(*b2GraphColor)(unsafe.Pointer(color)).ｆ__ccgo3_48.WideConstraints = uintptrFromInt32(0)
			(*b2GraphColor)(unsafe.Pointer(color)).WideConstraintCount = 0
		} else {
			(*b2GraphColor)(unsafe.Pointer(color)).ｆ__ccgo3_48.WideConstraints = wideContactConstraints + uintptr(wideBase*wideContactConstraintByteCount)
			colorContactCountW1 = (colorContactCount1-int32(1))>>simdShift + int32(1)
			(*b2GraphColor)(unsafe.Pointer(color)).WideConstraintCount = colorContactCountW1
			// Zero remainder lanes in the tail wide slot so prepare workers don't need to
			// initialize them.
			if colorContactCount1&((*b2World)(unsafe.Pointer(world1)).SimdWidth-int32(1)) != 0 {
				memset(tls, (*b2GraphColor)(unsafe.Pointer(color)).ｆ__ccgo3_48.WideConstraints+uintptr((colorContactCountW1-int32(1))*wideContactConstraintByteCount), 0, uint64FromInt32(wideContactConstraintByteCount))
			}
			wideBase = wideBase + colorContactCountW1
		}
		(*(*[25]b2JointPrepareSpan)(unsafe.Add(nil, bp1+504)))[i2].Start = jointBase
		(*(*[25]b2JointPrepareSpan)(unsafe.Add(nil, bp1+504)))[i2].Count = (*b2GraphColor)(unsafe.Pointer(color)).JointSims.Count
		(*(*[25]b2JointPrepareSpan)(unsafe.Add(nil, bp1+504)))[i2].Joints = (*b2GraphColor)(unsafe.Pointer(color)).JointSims.Data
		jointBase = jointBase + (*b2GraphColor)(unsafe.Pointer(color)).JointSims.Count
		goto _7
	_7:
		;
		i2 = i2 + 1
	}
	// Sentinel
	(*(*[25]b2ContactPrepareSpan)(unsafe.Add(nil, bp1+104)))[activeColorCount].Start = wideContactCount
	(*(*[25]b2ContactPrepareSpan)(unsafe.Add(nil, bp1+104)))[activeColorCount].Count = 0
	(*(*[25]b2ContactPrepareSpan)(unsafe.Add(nil, bp1+104)))[activeColorCount].Contacts = uintptrFromInt32(0)
	if v2 = !!(wideBase == wideContactCount); !v2 {
		b2InternalAssert(tls, __ccgo_ts+25688, __ccgo_ts+24283, int32FromInt32(1454))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	(*(*[25]b2JointPrepareSpan)(unsafe.Add(nil, bp1+504)))[activeColorCount].Start = jointCount
	(*(*[25]b2JointPrepareSpan)(unsafe.Add(nil, bp1+504)))[activeColorCount].Count = 0
	(*(*[25]b2JointPrepareSpan)(unsafe.Add(nil, bp1+504)))[activeColorCount].Joints = uintptrFromInt32(0)
	if v2 = !!(jointBase == jointCount); !v2 {
		b2InternalAssert(tls, __ccgo_ts+25717, __ccgo_ts+24283, int32FromInt32(1459))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	stageCount = 0
	// b2_stagePrepareJoints
	stageCount = stageCount + int32(1)
	// b2_stagePrepareContacts
	stageCount = stageCount + int32(1)
	// b2_stageIntegrateVelocities
	stageCount = stageCount + int32(1)
	// b2_stageWarmStart
	stageCount = stageCount + activeColorCount
	// b2_stageSolve
	stageCount = stageCount + int32(_ITERATIONS)*activeColorCount
	// b2_stageIntegratePositions
	stageCount = stageCount + int32(1)
	// b2_stageRelax
	stageCount = stageCount + int32(_RELAX_ITERATIONS)*activeColorCount
	// b2_stageRestitution
	stageCount = stageCount + activeColorCount
	// b2_stageStoreImpulses
	stageCount = stageCount + int32(1)
	stages = b2StackAlloc(tls, world1, int32FromUint64(uint64FromInt32(stageCount)*uint64(152)), __ccgo_ts+25741)
	bodyBlocks = b2StackAlloc(tls, world1, int32FromUint64(uint64FromInt32(bodyDim.Count)*uint64(64)), __ccgo_ts+25748)
	contactBlocks = b2StackAlloc(tls, world1, int32FromUint64(uint64FromInt32(contactPrepareDim.Count)*uint64(64)), __ccgo_ts+25760)
	jointBlocks = b2StackAlloc(tls, world1, int32FromUint64(uint64FromInt32(jointPrepareDim.Count)*uint64(64)), __ccgo_ts+25775)
	graphBlocks = b2StackAlloc(tls, world1, int32FromUint64(uint64FromInt32(graphBlockCount)*uint64(64)), __ccgo_ts+25788)
	// Split an awake island. This modifies:
	// - stack allocator
	// - world island array and solver set
	// - island indices on bodies, contacts, and joints
	// I'm squeezing this task in here because it may be expensive and this is a safe place to put it.
	// Note: cannot split islands in parallel with FinalizeBodies
	splitIslandTask = uintptrFromInt32(0)
	if (*b2World)(unsafe.Pointer(world1)).SplitIslandId != -int32(1) {
		if (*b2World)(unsafe.Pointer(world1)).TaskCount < int32(_B2_MAX_TASKS) {
			splitIslandTask = (*(*func(*_Stack, uintptr, uintptr, uintptr) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world1)).EnqueueTaskFcn})))(tls, __ccgo_fp(b2SplitIslandTask), world1, (*b2World)(unsafe.Pointer(world1)).UserTaskContext)
			*(*int32)(unsafe.Add(nil, world1+2780)) += int32(1)
			if splitIslandTask == uintptrFromInt32(0) {
				v1 = 0
			} else {
				v1 = int32(1)
			}
			*(*int32)(unsafe.Add(nil, world1+2776)) += v1
		} else {
			b2SplitIslandTask(tls, world1)
		}
	}
	// Prepare body blocks
	b2InitBlocks(tls, bodyBlocks, bodyDim, awakeBodyCount, uint8(b2_bodyBlock), uint8FromInt32(int32FromInt32(_UINT8_MAX1)))
	// Prepare blocks as a single flat parallel-for over the whole constraint range.
	// The task walks spans to decode flat slot indices back to per-color arrays.
	b2InitBlocks(tls, contactBlocks, contactPrepareDim, wideContactCount, uint8(b2_contactBlock), uint8FromInt32(int32FromInt32(_UINT8_MAX1)))
	b2InitBlocks(tls, jointBlocks, jointPrepareDim, jointCount, uint8(b2_jointBlock), uint8FromInt32(int32FromInt32(_UINT8_MAX1)))
	// Prepare graph work blocks. Each color gets joint blocks followed by contact blocks.
	*(*[24]uintptr)(unsafe.Add(nil, bp1+904)) = [24]uintptr{}
	baseGraphBlock = graphBlocks
	*(*[24]int32)(unsafe.Add(nil, bp1+1096)) = [24]int32{}
	i3 = 0
	for {
		if !(i3 < activeColorCount) {
			break
		}
		(*(*[24]uintptr)(unsafe.Add(nil, bp1+904)))[i3] = baseGraphBlock
		colorIndex = uint8FromInt32((*(*[24]int32)(unsafe.Add(nil, bp1+8)))[i3])
		b2InitBlocks(tls, baseGraphBlock, graphJointDims[i3], colorJointCounts[i3], uint8(b2_graphJointBlock), colorIndex)
		baseGraphBlock = baseGraphBlock + uintptr(graphJointDims[i3].Count)*64
		b2InitBlocks(tls, baseGraphBlock, graphContactDims[i3], colorContactCounts[i3], uint8(b2_graphContactBlock), colorIndex)
		baseGraphBlock = baseGraphBlock + uintptr(graphContactDims[i3].Count)*64
		(*(*[24]int32)(unsafe.Add(nil, bp1+1096)))[i3] = graphJointDims[i3].Count + graphContactDims[i3].Count
		goto _11
	_11:
		;
		i3 = i3 + 1
	}
	if v2 = !!((int64(baseGraphBlock)-int64(graphBlocks))/64 == int64(graphBlockCount)); !v2 {
		b2InternalAssert(tls, __ccgo_ts+25801, __ccgo_ts+24283, int32FromInt32(1537))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	stage = stages
	stage = b2InitStage(tls, stage, int32(b2_stagePrepareJoints), jointBlocks, jointPrepareDim.Count, uint8FromInt32(int32FromInt32(_UINT8_MAX1)))
	stage = b2InitStage(tls, stage, int32(b2_stagePrepareContacts), contactBlocks, contactPrepareDim.Count, uint8FromInt32(int32FromInt32(_UINT8_MAX1)))
	stage = b2InitStage(tls, stage, int32(b2_stageIntegrateVelocities), bodyBlocks, bodyDim.Count, uint8FromInt32(int32FromInt32(_UINT8_MAX1)))
	stage = b2InitColorStages(tls, stage, int32(b2_stageWarmStart), int32(1), activeColorCount, bp1+904, bp1+1096, bp1+8)
	stage = b2InitColorStages(tls, stage, int32(b2_stageSolve), int32(_ITERATIONS), activeColorCount, bp1+904, bp1+1096, bp1+8)
	stage = b2InitStage(tls, stage, int32(b2_stageIntegratePositions), bodyBlocks, bodyDim.Count, uint8FromInt32(int32FromInt32(_UINT8_MAX1)))
	stage = b2InitColorStages(tls, stage, int32(b2_stageRelax), int32(_RELAX_ITERATIONS), activeColorCount, bp1+904, bp1+1096, bp1+8)
	stage = b2InitColorStages(tls, stage, int32(b2_stageRestitution), int32(1), activeColorCount, bp1+904, bp1+1096, bp1+8)
	stage = b2InitStage(tls, stage, int32(b2_stageStoreImpulses), contactBlocks, contactPrepareDim.Count, uint8FromInt32(int32FromInt32(_UINT8_MAX1)))
	if v2 = !!(int32((int64(stage)-int64(stages))/152) == stageCount); !v2 {
		b2InternalAssert(tls, __ccgo_ts+25864, __ccgo_ts+24283, int32FromInt32(1554))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	if v2 = !!(workerCount <= int32FromInt32(_B2_MAX_WORKERS)); !v2 {
		b2InternalAssert(tls, __ccgo_ts+25902, __ccgo_ts+24283, int32FromInt32(1556))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	(*b2StepContext)(unsafe.Pointer(stepContext)).Graph = graph
	(*b2StepContext)(unsafe.Pointer(stepContext)).ActiveColorCount = activeColorCount
	(*b2StepContext)(unsafe.Pointer(stepContext)).WorkerCount = workerCount
	(*b2StepContext)(unsafe.Pointer(stepContext)).StageCount = stageCount
	(*b2StepContext)(unsafe.Pointer(stepContext)).Stages = stages
	(*b2StepContext)(unsafe.Pointer(stepContext)).WideContactConstraints = wideContactConstraints
	(*b2StepContext)(unsafe.Pointer(stepContext)).ContactPrepareSpans = bp1 + 104
	(*b2StepContext)(unsafe.Pointer(stepContext)).WideContactCount = wideContactCount
	(*b2StepContext)(unsafe.Pointer(stepContext)).JointPrepareSpans = bp1 + 504
	atomicStoreNUint32(stepContext+256, uint32(0), int32FromInt32(__ATOMIC_SEQ_CST))
	atomicStoreNInt32(stepContext+324, 0, int32FromInt32(__ATOMIC_SEQ_CST))
	(*b2World)(unsafe.Pointer(world1)).Profile.SolverSetup = b2GetMillisecondsAndReset(tls, bp1)
	*(*uint64_t)(unsafe.Add(nil, bp1+1960)) = b2GetTicks(tls)
	v1 = (*b2IdPool)(unsafe.Pointer(world1 + 1992)).NextIndex
	goto _16
_16:
	jointIdCapacity = v1
	v4 = (*b2IdPool)(unsafe.Pointer(world1 + 2032)).NextIndex
	goto _18
_18:
	contactIdCapacity = v4
	i4 = 0
	for {
		if !(i4 < workerCount) {
			break
		}
		if v2 = !!(0 <= i4 && i4 < (*b2World)(unsafe.Pointer(world1)).TaskContexts.Count); !v2 {
			b2InternalAssert(tls, __ccgo_ts+25932, __ccgo_ts+24283, int32FromInt32(1581))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		taskContext = (*b2World)(unsafe.Pointer(world1)).TaskContexts.Data + uintptr(i4)*136
		b2SetBitCountAndClear(tls, taskContext+72, uint32FromInt32(jointIdCapacity))
		b2SetBitCountAndClear(tls, taskContext+48, uint32FromInt32(contactIdCapacity))
		(*b2TaskContext)(unsafe.Pointer(taskContext)).HasHitEvents = boolUint8(false1 != 0)
		(*(*[32]b2WorkerContext)(unsafe.Add(nil, bp1+1192)))[i4].Context = stepContext
		(*(*[32]b2WorkerContext)(unsafe.Add(nil, bp1+1192)))[i4].WorkerIndex = i4
		if (*b2World)(unsafe.Pointer(world1)).TaskCount < int32(_B2_MAX_TASKS) {
			(*(*[32]b2WorkerContext)(unsafe.Add(nil, bp1+1192)))[i4].UserTask = (*(*func(*_Stack, uintptr, uintptr, uintptr) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world1)).EnqueueTaskFcn})))(tls, __ccgo_fp(b2SolverTask), bp1+1192+uintptr(i4)*24, (*b2World)(unsafe.Pointer(world1)).UserTaskContext)
			*(*int32)(unsafe.Add(nil, world1+2780)) += int32(1)
			if (*(*[32]b2WorkerContext)(unsafe.Add(nil, bp1+1192)))[i4].UserTask == uintptrFromInt32(0) {
				v1 = 0
			} else {
				v1 = int32(1)
			}
			*(*int32)(unsafe.Add(nil, world1+2776)) += v1
		} else {
			(*(*[32]b2WorkerContext)(unsafe.Add(nil, bp1+1192)))[i4].UserTask = uintptrFromInt32(0)
			b2SolverTask(tls, bp1+1192+uintptr(i4)*24)
		}
		goto _19
	_19:
		;
		i4 = i4 + 1
	}
	// The calling thread of b2World_Step also enters b2SolverTask as worker 0 and races for the
	// orchestrator slot via the CAS inside. This guarantees progress even when the user's task
	// system can't run the queued worker 0 promptly: it might schedule out of order, have fewer
	// threads than workerCount, or invert priority by parking the calling thread in finishTaskFcn.
	// Whoever wins the CAS becomes the orchestrator; the loser returns and lets the spinner-only
	// path handle workers >0.
	*(*b2WorkerContext)(unsafe.Add(nil, bp1+1968)) = b2WorkerContext{
		Context: stepContext,
	}
	b2SolverTask(tls, bp1+1968)
	// Finish constraint solve
	i5 = 0
	for {
		if !(i5 < workerCount) {
			break
		}
		if (*(*[32]b2WorkerContext)(unsafe.Add(nil, bp1+1192)))[i5].UserTask != uintptrFromInt32(0) {
			(*(*func(*_Stack, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world1)).FinishTaskFcn})))(tls, (*(*[32]b2WorkerContext)(unsafe.Add(nil, bp1+1192)))[i5].UserTask, (*b2World)(unsafe.Pointer(world1)).UserTaskContext)
			*(*int32)(unsafe.Add(nil, world1+2776)) -= int32(1)
		}
		goto _22
	_22:
		;
		i5 = i5 + 1
	}
	// Finish island split
	if splitIslandTask != uintptrFromInt32(0) {
		(*(*func(*_Stack, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world1)).FinishTaskFcn})))(tls, splitIslandTask, (*b2World)(unsafe.Pointer(world1)).UserTaskContext)
		*(*int32)(unsafe.Add(nil, world1+2776)) -= int32(1)
	}
	(*b2World)(unsafe.Pointer(world1)).SplitIslandId = -int32(1)
	(*b2World)(unsafe.Pointer(world1)).Profile.Constraints = b2GetMillisecondsAndReset(tls, bp1+1960)
	transformTicks = b2GetTicks(tls)
	// Prepare contact, enlarged body, and island bit sets used in body finalization.
	awakeIslandCount = (*b2SolverSet)(unsafe.Pointer(awakeSet)).IslandSims.Count
	i6 = 0
	for {
		if !(i6 < (*b2World)(unsafe.Pointer(world1)).WorkerCount) {
			break
		}
		taskContext1 = (*b2World)(unsafe.Pointer(world1)).TaskContexts.Data + uintptr(i6)*136
		(*b2TaskContext)(unsafe.Pointer(taskContext1)).SensorHits.Count = 0
		b2SetBitCountAndClear(tls, taskContext1+88, uint32FromInt32(awakeBodyCount))
		b2SetBitCountAndClear(tls, taskContext1+104, uint32FromInt32(awakeIslandCount))
		(*b2TaskContext)(unsafe.Pointer(taskContext1)).SplitIslandId = -int32(1)
		(*b2TaskContext)(unsafe.Pointer(taskContext1)).SplitSleepTime = float32FromFloat32(0)
		goto _23
	_23:
		;
		i6 = i6 + 1
	}
	// Finish the user tree task that was queued earlier in the time step. This must be complete before touching the
	// broad-phase.
	if (*b2World)(unsafe.Pointer(world1)).UserTreeTask != uintptrFromInt32(0) {
		(*(*func(*_Stack, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world1)).FinishTaskFcn})))(tls, (*b2World)(unsafe.Pointer(world1)).UserTreeTask, (*b2World)(unsafe.Pointer(world1)).UserTaskContext)
		(*b2World)(unsafe.Pointer(world1)).UserTreeTask = uintptrFromInt32(0)
		*(*int32)(unsafe.Add(nil, world1+2776)) -= int32(1)
	}
	b2ValidateNoEnlarged(tls, world1+40)
	// Finalize bodies. Must happen after the constraint solver and after island splitting.
	b2ParallelFor(tls, world1, __ccgo_fp(b2FinalizeBodiesTask), awakeBodyCount, int32(64), stepContext)
	b2StackFree(tls, world1, graphBlocks)
	b2StackFree(tls, world1, jointBlocks)
	b2StackFree(tls, world1, contactBlocks)
	b2StackFree(tls, world1, bodyBlocks)
	b2StackFree(tls, world1, stages)
	b2StackFree(tls, world1, overflowContacts)
	b2StackFree(tls, world1, wideContactConstraints)
	(*b2World)(unsafe.Pointer(world1)).Profile.Transforms = b2GetMilliseconds(tls, transformTicks)
	// Report joint events
	jointEventTicks = b2GetTicks(tls)
	// Gather bits for all joints that have force/torque events
	jointStateBitSet = (*b2World)(unsafe.Pointer(world1)).TaskContexts.Data + 72
	i7 = int32(1)
	for {
		if !(i7 < (*b2World)(unsafe.Pointer(world1)).WorkerCount) {
			break
		}
		b2InPlaceUnion(tls, jointStateBitSet, (*b2World)(unsafe.Pointer(world1)).TaskContexts.Data+uintptr(i7)*136+72)
		goto _24
	_24:
		;
		i7 = i7 + 1
	}
	wordCount = (*b2BitSet)(unsafe.Pointer(jointStateBitSet)).BlockCount
	bits = (*b2BitSet)(unsafe.Pointer(jointStateBitSet)).Bits
	jointArray = (*b2World)(unsafe.Pointer(world1)).Joints.Data
	worldIndex0 = (*b2World)(unsafe.Pointer(world1)).WorldId
	k = uint32(0)
	for {
		if !(k < wordCount) {
			break
		}
		word = *(*uint64_t)(unsafe.Add(nil, bits+uintptr(k)*8))
		for word != uint64(0) {
			v26 = uint32FromInt32(__builtin_ctzll(tls, word))
			goto _27
		_27:
			ctz = v26
			jointId = int32FromUint32(uint32FromInt32(64)*k + ctz)
			if v2 = !!(jointId < (*b2World)(unsafe.Pointer(world1)).Joints.Capacity); !v2 {
				b2InternalAssert(tls, __ccgo_ts+25984, __ccgo_ts+24283, int32FromInt32(1700))
			}
			_ = v2 || bool(int32FromInt32(0) != 0)
			joint = jointArray + uintptr(jointId)*72
			if v2 = !!((*b2Joint)(unsafe.Pointer(joint)).SetIndex == int32(b2_awakeSet)); !v2 {
				b2InternalAssert(tls, __ccgo_ts+26017, __ccgo_ts+24283, int32FromInt32(1704))
			}
			_ = v2 || bool(int32FromInt32(0) != 0)
			event = JointEvent{
				JointId: JointId{
					Index1:     jointId + int32(1),
					World0:     worldIndex0,
					Generation: (*b2Joint)(unsafe.Pointer(joint)).Generation,
				},
				UserData: (*b2Joint)(unsafe.Pointer(joint)).UserData,
			}
			elementSize = int32(16)
			if (*b2World)(unsafe.Pointer(world1)).JointEvents.Count >= (*b2World)(unsafe.Pointer(world1)).JointEvents.Capacity {
				oldSize1 = (*b2World)(unsafe.Pointer(world1)).JointEvents.Capacity * elementSize
				if (*b2World)(unsafe.Pointer(world1)).JointEvents.Capacity == 0 {
					v1 = int32(8)
				} else {
					v1 = int32(2) * (*b2World)(unsafe.Pointer(world1)).JointEvents.Capacity
				}
				newCapacity = v1
				newSize1 = newCapacity * elementSize
				(*b2World)(unsafe.Pointer(world1)).JointEvents.Data = b2GrowAlloc(tls, (*b2World)(unsafe.Pointer(world1)).JointEvents.Data, uint64FromInt32(oldSize1), uint64FromInt32(newSize1))
				(*b2World)(unsafe.Pointer(world1)).JointEvents.Capacity = newCapacity
			}
			*(*JointEvent)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world1)).JointEvents.Data+uintptr((*b2World)(unsafe.Pointer(world1)).JointEvents.Count)*16)) = event
			(*b2World)(unsafe.Pointer(world1)).JointEvents.Count += int32(1)
			// Clear the smallest set bit
			word = word & (word - uint64(1))
		}
		goto _25
	_25:
		;
		k = k + 1
	}
	(*b2World)(unsafe.Pointer(world1)).Profile.JointEvents = b2GetMilliseconds(tls, jointEventTicks)
	// Report hit events
	hitTicks = b2GetTicks(tls)
	if v2 = !!((*b2World)(unsafe.Pointer(world1)).ContactHitEvents.Count == int32FromInt32(0)); !v2 {
		b2InternalAssert(tls, __ccgo_ts+26048, __ccgo_ts+24283, int32FromInt32(1733))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	// Fast path: skip if no worker flagged any hit-event candidates during impulse storing.
	anyHitEvents = boolUint8(false1 != 0)
	i8 = 0
	for {
		if !(i8 < (*b2World)(unsafe.Pointer(world1)).WorkerCount) {
			break
		}
		if (*(*b2TaskContext)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world1)).TaskContexts.Data+uintptr(i8)*136))).HasHitEvents != 0 {
			anyHitEvents = boolUint8(true1 != 0)
			break
		}
		goto _32
	_32:
		;
		i8 = i8 + 1
	}
	if anyHitEvents != 0 {
		// Union per-worker bits into worker 0's bit set.
		hitEventBitSet = (*b2World)(unsafe.Pointer(world1)).TaskContexts.Data + 48
		i9 = int32(1)
		for {
			if !(i9 < (*b2World)(unsafe.Pointer(world1)).WorkerCount) {
				break
			}
			if (*(*b2TaskContext)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world1)).TaskContexts.Data+uintptr(i9)*136))).HasHitEvents != 0 {
				b2InPlaceUnion(tls, hitEventBitSet, (*b2World)(unsafe.Pointer(world1)).TaskContexts.Data+uintptr(i9)*136+48)
			}
			goto _33
		_33:
			;
			i9 = i9 + 1
		}
		threshold = (*b2World)(unsafe.Pointer(world1)).HitEventThreshold
		colors1 = world1 + 376
		contactArray = (*b2World)(unsafe.Pointer(world1)).Contacts.Data
		shapeArray = (*b2World)(unsafe.Pointer(world1)).Shapes.Data
		worldId = (*b2World)(unsafe.Pointer(world1)).WorldId
		wordCount1 = (*b2BitSet)(unsafe.Pointer(hitEventBitSet)).BlockCount
		bits1 = (*b2BitSet)(unsafe.Pointer(hitEventBitSet)).Bits
		k1 = uint32(0)
		for {
			if !(k1 < wordCount1) {
				break
			}
			word1 = *(*uint64_t)(unsafe.Add(nil, bits1+uintptr(k1)*8))
			for word1 != uint64(0) {
				v26 = uint32FromInt32(__builtin_ctzll(tls, word1))
				goto _36
			_36:
				ctz1 = v26
				contactId = int32FromUint32(uint32FromInt32(64)*k1 + ctz1)
				contact = contactArray + uintptr(contactId)*64
				if v2 = !!((*b2Contact)(unsafe.Pointer(contact)).SetIndex == int32(b2_awakeSet) && (*b2Contact)(unsafe.Pointer(contact)).ColorIndex != -int32(1)); !v2 {
					b2InternalAssert(tls, __ccgo_ts+26083, __ccgo_ts+24283, int32FromInt32(1775))
				}
				_ = v2 || bool(int32FromInt32(0) != 0)
				color1 = colors1 + uintptr((*b2Contact)(unsafe.Pointer(contact)).ColorIndex)*64
				contactSim = (*b2GraphColor)(unsafe.Pointer(color1)).ContactSims.Data + uintptr((*b2Contact)(unsafe.Pointer(contact)).LocalIndex)*200
				event1 = ContactHitEvent{}
				event1.ApproachSpeed = threshold
				bestPoint = uintptrFromInt32(0)
				pointCount = (*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.PointCount
				p1 = 0
				for {
					if !(p1 < pointCount) {
						break
					}
					mp = contactSim + 72 + 12 + uintptr(p1)*44
					approachSpeed = -(*ManifoldPoint)(unsafe.Pointer(mp)).NormalVelocity
					// Need to check total impulse because the point may be speculative and not colliding
					if approachSpeed > event1.ApproachSpeed && (*ManifoldPoint)(unsafe.Pointer(mp)).TotalNormalImpulse > float32FromFloat32(0) {
						event1.ApproachSpeed = approachSpeed
						bestPoint = mp
					}
					goto _38
				_38:
					;
					p1 = p1 + 1
				}
				if bestPoint != uintptrFromInt32(0) {
					event1.Normal = (*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.Normal
					shapeA = shapeArray + uintptr((*b2ContactSim)(unsafe.Pointer(contactSim)).ShapeIdA)*280
					shapeB = shapeArray + uintptr((*b2ContactSim)(unsafe.Pointer(contactSim)).ShapeIdB)*280
					if v2 = !!(0 <= (*b2Shape)(unsafe.Pointer(shapeA)).BodyId && (*b2Shape)(unsafe.Pointer(shapeA)).BodyId < (*b2World)(unsafe.Pointer(world1)).Bodies.Count); !v2 {
						b2InternalAssert(tls, __ccgo_ts+3985, __ccgo_ts+24283, int32FromInt32(1811))
					}
					_ = v2 || bool(int32FromInt32(0) != 0)
					// World contact point reconstructed from a body center of mass and the matching
					// anchor. The anchors were built with the manifold, so a body that has moved since
					// drags the point with it. A static body has not moved, prefer one so the common
					// case of a fast body striking the world stays exact.
					bodyA = (*b2World)(unsafe.Pointer(world1)).Bodies.Data + uintptr((*b2Shape)(unsafe.Pointer(shapeA)).BodyId)*104
					if v8 = !!(0 <= (*b2Shape)(unsafe.Pointer(shapeB)).BodyId && (*b2Shape)(unsafe.Pointer(shapeB)).BodyId < (*b2World)(unsafe.Pointer(world1)).Bodies.Count); !v8 {
						b2InternalAssert(tls, __ccgo_ts+3913, __ccgo_ts+24283, int32FromInt32(1812))
					}
					_ = v8 || bool(int32FromInt32(0) != 0)
					bodyB = (*b2World)(unsafe.Pointer(world1)).Bodies.Data + uintptr((*b2Shape)(unsafe.Pointer(shapeB)).BodyId)*104
					if (*b2Body)(unsafe.Pointer(bodyA)).Type1 != int32(b2_staticBody) && (*b2Body)(unsafe.Pointer(bodyB)).Type1 == int32(b2_staticBody) {
						v41 = world1
						v42 = bodyB
						if v9 = !!(0 <= (*b2Body)(unsafe.Pointer(v42)).SetIndex && (*b2Body)(unsafe.Pointer(v42)).SetIndex < (*b2World)(unsafe.Pointer(v41)).SolverSets.Count); !v9 {
							b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
						}
						_ = v9 || bool(int32FromInt32(0) != 0)
						set = (*b2World)(unsafe.Pointer(v41)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v42)).SetIndex)*88
						if v12 = !!(0 <= (*b2Body)(unsafe.Pointer(v42)).LocalIndex && (*b2Body)(unsafe.Pointer(v42)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v12 {
							b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
						}
						_ = v12 || bool(int32FromInt32(0) != 0)
						bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v42)).LocalIndex)*96
						v45 = bodySim
						goto _46
					_46:
						bodySimB = v45
						v47 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Center
						v48 = (*ManifoldPoint)(unsafe.Pointer(bestPoint)).AnchorB
						v49 = Pos{
							X: v47.X + v48.X,
							Y: v47.Y + v48.Y,
						}
						goto _50
					_50:
						event1.Point = v49
					} else {
						v41 = world1
						v42 = bodyA
						if v2 = !!(0 <= (*b2Body)(unsafe.Pointer(v42)).SetIndex && (*b2Body)(unsafe.Pointer(v42)).SetIndex < (*b2World)(unsafe.Pointer(v41)).SolverSets.Count); !v2 {
							b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
						}
						_ = v2 || bool(int32FromInt32(0) != 0)
						set = (*b2World)(unsafe.Pointer(v41)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v42)).SetIndex)*88
						if v8 = !!(0 <= (*b2Body)(unsafe.Pointer(v42)).LocalIndex && (*b2Body)(unsafe.Pointer(v42)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v8 {
							b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
						}
						_ = v8 || bool(int32FromInt32(0) != 0)
						bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v42)).LocalIndex)*96
						v45 = bodySim
						goto _56
					_56:
						bodySimA = v45
						v47 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Center
						v48 = (*ManifoldPoint)(unsafe.Pointer(bestPoint)).AnchorA
						v49 = Pos{
							X: v47.X + v48.X,
							Y: v47.Y + v48.Y,
						}
						goto _60
					_60:
						event1.Point = v49
					}
					event1.ShapeIdA = ShapeId{
						Index1:     (*b2Shape)(unsafe.Pointer(shapeA)).Id + int32(1),
						World0:     worldId,
						Generation: (*b2Shape)(unsafe.Pointer(shapeA)).Generation,
					}
					event1.ShapeIdB = ShapeId{
						Index1:     (*b2Shape)(unsafe.Pointer(shapeB)).Id + int32(1),
						World0:     worldId,
						Generation: (*b2Shape)(unsafe.Pointer(shapeB)).Generation,
					}
					event1.ContactId = ContactId{
						Index1:     (*b2Contact)(unsafe.Pointer(contact)).ContactId + int32(1),
						World0:     worldId,
						Generation: (*b2Contact)(unsafe.Pointer(contact)).Generation,
					}
					elementSize1 = int32(48)
					if (*b2World)(unsafe.Pointer(world1)).ContactHitEvents.Count >= (*b2World)(unsafe.Pointer(world1)).ContactHitEvents.Capacity {
						oldSize2 = (*b2World)(unsafe.Pointer(world1)).ContactHitEvents.Capacity * elementSize1
						if (*b2World)(unsafe.Pointer(world1)).ContactHitEvents.Capacity == 0 {
							v1 = int32(8)
						} else {
							v1 = int32(2) * (*b2World)(unsafe.Pointer(world1)).ContactHitEvents.Capacity
						}
						newCapacity1 = v1
						newSize2 = newCapacity1 * elementSize1
						(*b2World)(unsafe.Pointer(world1)).ContactHitEvents.Data = b2GrowAlloc(tls, (*b2World)(unsafe.Pointer(world1)).ContactHitEvents.Data, uint64FromInt32(oldSize2), uint64FromInt32(newSize2))
						(*b2World)(unsafe.Pointer(world1)).ContactHitEvents.Capacity = newCapacity1
					}
					*(*ContactHitEvent)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world1)).ContactHitEvents.Data+uintptr((*b2World)(unsafe.Pointer(world1)).ContactHitEvents.Count)*48)) = event1
					(*b2World)(unsafe.Pointer(world1)).ContactHitEvents.Count += int32(1)
				}
				// Clear the smallest set bit
				word1 = word1 & (word1 - uint64(1))
			}
			goto _34
		_34:
			;
			k1 = k1 + 1
		}
	}
	(*b2World)(unsafe.Pointer(world1)).Profile.HitEvents = b2GetMilliseconds(tls, hitTicks)
	refitTicks = b2GetTicks(tls)
	bp = world1 + 40
	b2DynamicTree_Refit(tls, bp+uintptr(b2_kinematicBody)*104)
	b2DynamicTree_Refit(tls, bp+uintptr(b2_dynamicBody)*104)
	b2ValidateBroadphase(tls, world1+40)
	(*b2World)(unsafe.Pointer(world1)).Profile.Refit = b2GetMilliseconds(tls, refitTicks)
	v1 = atomicLoadNInt32(stepContext+112, int32FromInt32(__ATOMIC_SEQ_CST))
	goto _63
_63:
	// Bullets are processed after the broad-phase refit so they can query the
	// final non-bullet world.
	bulletBodyCount = v1
	if bulletBodyCount > 0 {
		bulletTicks = b2GetTicks(tls)
		// Fast bullet bodies
		// Note: a bullet body may be moving slow
		minRange = int32(8)
		b2ParallelFor(tls, world1, __ccgo_fp(b2BulletBodyTask), bulletBodyCount, minRange, stepContext)
		broadPhase = world1 + 40
		dynamicTree = broadPhase + uintptr(b2_dynamicBody)*104
		// Fast array access is important here
		bodyArray = (*b2World)(unsafe.Pointer(world1)).Bodies.Data
		bodySimArray = (*b2SolverSet)(unsafe.Pointer(awakeSet)).BodySims.Data
		shapeArray1 = (*b2World)(unsafe.Pointer(world1)).Shapes.Data
		bulletBodySimIndices = (*b2StepContext)(unsafe.Pointer(stepContext)).BulletBodies
		// Serially enlarge broad-phase proxies for bullet shapes.
		// This loop has non-deterministic order but it shouldn't affect the result.
		i10 = 0
		for {
			if !(i10 < bulletBodyCount) {
				break
			}
			bulletBodySim = bodySimArray + uintptr(*(*int32)(unsafe.Add(nil, bulletBodySimIndices+uintptr(i10)*4)))*96
			// It is worth tracking this flag because bullets may be moving slowly.
			if (*b2BodySim)(unsafe.Pointer(bulletBodySim)).Flags&uint32(b2_enlargeBulletBounds) == uint32(0) {
				goto _64
			}
			// Clear flag
			*(*uint32_t)(unsafe.Add(nil, bulletBodySim+92)) &= uint32FromInt32(^int32(b2_enlargeBulletBounds))
			bodyId = (*b2BodySim)(unsafe.Pointer(bulletBodySim)).BodyId
			if v2 = !!(0 <= bodyId && bodyId < (*b2World)(unsafe.Pointer(world1)).Bodies.Count); !v2 {
				b2InternalAssert(tls, __ccgo_ts+26156, __ccgo_ts+24283, int32FromInt32(1899))
			}
			_ = v2 || bool(int32FromInt32(0) != 0)
			bulletBody = bodyArray + uintptr(bodyId)*104
			shapeId = (*b2Body)(unsafe.Pointer(bulletBody)).HeadShapeId
			for shapeId != -int32(1) {
				shape = shapeArray1 + uintptr(shapeId)*280
				proxyKey = (*b2Shape)(unsafe.Pointer(shape)).ProxyKey
				proxyId = proxyKey >> int32(2)
				treeAABB = b2DynamicTree_GetProxyBounds(tls, dynamicTree, proxyId)
				shapeFatAABB = *(*AABB)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world1)).FatAABBs.Data+uintptr(shapeId)*16))
				// Double check containment because a bullet can have multiple
				// shapes and maybe just one needs an update.
				v66 = treeAABB
				v67 = shapeFatAABB
				s = boolUint8(true1 != 0)
				s = boolUint8(s != 0 && v66.LowerBound.X <= v67.LowerBound.X)
				s = boolUint8(s != 0 && v66.LowerBound.Y <= v67.LowerBound.Y)
				s = boolUint8(s != 0 && v67.UpperBound.X <= v66.UpperBound.X)
				s = boolUint8(s != 0 && v67.UpperBound.Y <= v66.UpperBound.Y)
				v68 = s
				goto _69
			_69:
				if int32FromUint8(v68) == false1 {
					b2DynamicTree_EnlargeProxy(tls, dynamicTree, proxyId, shapeFatAABB)
				}
				shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
			}
			goto _64
		_64:
			;
			i10 = i10 + 1
		}
		(*b2World)(unsafe.Pointer(world1)).Profile.Bullets = b2GetMilliseconds(tls, bulletTicks)
	}
	// Need to free this even if no bullets got processed.
	b2StackFree(tls, world1, (*b2StepContext)(unsafe.Pointer(stepContext)).BulletBodies)
	(*b2StepContext)(unsafe.Pointer(stepContext)).BulletBodies = uintptrFromInt32(0)
	atomicStoreNInt32(stepContext+112, 0, int32FromInt32(__ATOMIC_SEQ_CST))
	// Report sensor hits. This may include bullets sensor hits.
	sensorHitTicks = b2GetTicks(tls)
	workerCount1 = (*b2World)(unsafe.Pointer(world1)).WorkerCount
	if v2 = !!(workerCount1 == (*b2World)(unsafe.Pointer(world1)).TaskContexts.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+26200, __ccgo_ts+24283, int32FromInt32(1939))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	i11 = 0
	for {
		if !(i11 < workerCount1) {
			break
		}
		taskContext2 = (*b2World)(unsafe.Pointer(world1)).TaskContexts.Data + uintptr(i11)*136
		hitCount = (*b2TaskContext)(unsafe.Pointer(taskContext2)).SensorHits.Count
		hits = (*b2TaskContext)(unsafe.Pointer(taskContext2)).SensorHits.Data
		j1 = 0
		for {
			if !(j1 < hitCount) {
				break
			}
			hit = *(*b2SensorHit)(unsafe.Add(nil, hits+uintptr(j1)*8))
			if v2 = !!(0 <= hit.SensorId && hit.SensorId < (*b2World)(unsafe.Pointer(world1)).Shapes.Count); !v2 {
				b2InternalAssert(tls, __ccgo_ts+26241, __ccgo_ts+24283, int32FromInt32(1950))
			}
			_ = v2 || bool(int32FromInt32(0) != 0)
			sensorShape = (*b2World)(unsafe.Pointer(world1)).Shapes.Data + uintptr(hit.SensorId)*280
			if v8 = !!(0 <= hit.VisitorId && hit.VisitorId < (*b2World)(unsafe.Pointer(world1)).Shapes.Count); !v8 {
				b2InternalAssert(tls, __ccgo_ts+26309, __ccgo_ts+24283, int32FromInt32(1951))
			}
			_ = v8 || bool(int32FromInt32(0) != 0)
			visitor = (*b2World)(unsafe.Pointer(world1)).Shapes.Data + uintptr(hit.VisitorId)*280
			if v9 = !!(0 <= (*b2Shape)(unsafe.Pointer(sensorShape)).SensorIndex && (*b2Shape)(unsafe.Pointer(sensorShape)).SensorIndex < (*b2World)(unsafe.Pointer(world1)).Sensors.Count); !v9 {
				b2InternalAssert(tls, __ccgo_ts+21770, __ccgo_ts+24283, int32FromInt32(1953))
			}
			_ = v9 || bool(int32FromInt32(0) != 0)
			sensor = (*b2World)(unsafe.Pointer(world1)).Sensors.Data + uintptr((*b2Shape)(unsafe.Pointer(sensorShape)).SensorIndex)*56
			shapeRef = b2Visitor{
				ShapeId:    hit.VisitorId,
				Generation: (*b2Shape)(unsafe.Pointer(visitor)).Generation,
			}
			elementSize2 = int32(8)
			if (*b2Sensor)(unsafe.Pointer(sensor)).Hits.Count >= (*b2Sensor)(unsafe.Pointer(sensor)).Hits.Capacity {
				oldSize3 = (*b2Sensor)(unsafe.Pointer(sensor)).Hits.Capacity * elementSize2
				if (*b2Sensor)(unsafe.Pointer(sensor)).Hits.Capacity == 0 {
					v1 = int32(8)
				} else {
					v1 = int32(2) * (*b2Sensor)(unsafe.Pointer(sensor)).Hits.Capacity
				}
				newCapacity2 = v1
				newSize3 = newCapacity2 * elementSize2
				(*b2Sensor)(unsafe.Pointer(sensor)).Hits.Data = b2GrowAlloc(tls, (*b2Sensor)(unsafe.Pointer(sensor)).Hits.Data, uint64FromInt32(oldSize3), uint64FromInt32(newSize3))
				(*b2Sensor)(unsafe.Pointer(sensor)).Hits.Capacity = newCapacity2
			}
			*(*b2Visitor)(unsafe.Add(nil, (*b2Sensor)(unsafe.Pointer(sensor)).Hits.Data+uintptr((*b2Sensor)(unsafe.Pointer(sensor)).Hits.Count)*8)) = shapeRef
			(*b2Sensor)(unsafe.Pointer(sensor)).Hits.Count += int32(1)
			goto _72
		_72:
			;
			j1 = j1 + 1
		}
		goto _71
	_71:
		;
		i11 = i11 + 1
	}
	(*b2World)(unsafe.Pointer(world1)).Profile.SensorHits = b2GetMilliseconds(tls, sensorHitTicks)
	// Island sleeping
	// This must be done last because putting islands to sleep invalidates the enlarged body bits.
	// todo_erin figure out how to do this in parallel with tree refit
	if int32FromUint8((*b2World)(unsafe.Pointer(world1)).EnableSleep) == int32(true1) {
		sleepTicks = b2GetTicks(tls)
		// Collect split island candidate for the next time step. No need to split if sleeping is disabled.
		if v2 = !!((*b2World)(unsafe.Pointer(world1)).SplitIslandId == -int32FromInt32(1)); !v2 {
			b2InternalAssert(tls, __ccgo_ts+26379, __ccgo_ts+24283, int32FromInt32(1975))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		splitSleepTimer = float32FromFloat32(0)
		i12 = 0
		for {
			if !(i12 < (*b2World)(unsafe.Pointer(world1)).WorkerCount) {
				break
			}
			taskContext3 = (*b2World)(unsafe.Pointer(world1)).TaskContexts.Data + uintptr(i12)*136
			if (*b2TaskContext)(unsafe.Pointer(taskContext3)).SplitIslandId != -int32(1) && (*b2TaskContext)(unsafe.Pointer(taskContext3)).SplitSleepTime >= splitSleepTimer {
				if v2 = !!((*b2TaskContext)(unsafe.Pointer(taskContext3)).SplitSleepTime > float32FromFloat32(0)); !v2 {
					b2InternalAssert(tls, __ccgo_ts+26417, __ccgo_ts+24283, int32FromInt32(1982))
				}
				_ = v2 || bool(int32FromInt32(0) != 0)
				// Tie breaking for determinism. Largest island id wins. Needed due to work stealing.
				if (*b2TaskContext)(unsafe.Pointer(taskContext3)).SplitSleepTime == splitSleepTimer && (*b2TaskContext)(unsafe.Pointer(taskContext3)).SplitIslandId < (*b2World)(unsafe.Pointer(world1)).SplitIslandId {
					goto _78
				}
				(*b2World)(unsafe.Pointer(world1)).SplitIslandId = (*b2TaskContext)(unsafe.Pointer(taskContext3)).SplitIslandId
				splitSleepTimer = (*b2TaskContext)(unsafe.Pointer(taskContext3)).SplitSleepTime
			}
			goto _78
		_78:
			;
			i12 = i12 + 1
		}
		awakeIslandBitSet = (*b2World)(unsafe.Pointer(world1)).TaskContexts.Data + 104
		i13 = int32(1)
		for {
			if !(i13 < (*b2World)(unsafe.Pointer(world1)).WorkerCount) {
				break
			}
			b2InPlaceUnion(tls, awakeIslandBitSet, (*b2World)(unsafe.Pointer(world1)).TaskContexts.Data+uintptr(i13)*136+104)
			goto _80
		_80:
			;
			i13 = i13 + 1
		}
		// Need to process in reverse because this moves islands to sleeping solver sets.
		islands = (*b2SolverSet)(unsafe.Pointer(awakeSet)).IslandSims.Data
		count = (*b2SolverSet)(unsafe.Pointer(awakeSet)).IslandSims.Count
		islandIndex = count - int32(1)
		for {
			if !(islandIndex >= 0) {
				break
			}
			v41 = awakeIslandBitSet
			v26 = uint32FromInt32(islandIndex)
			blockIndex = v26 / uint32(64)
			if blockIndex >= (*b2BitSet)(unsafe.Pointer(v41)).BlockCount {
				v68 = boolUint8(false1 != 0)
				goto _85
			}
			v68 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v41)).Bits+uintptr(blockIndex)*8))&(uint64FromInt32(1)<<(v26%uint32FromInt32(64))) != uint64(0))
			goto _85
		_85:
			if int32FromUint8(v68) == int32(true1) {
				// this island is still awake
				goto _81
			}
			island = islands + uintptr(islandIndex)*4
			islandId = (*b2IslandSim)(unsafe.Pointer(island)).IslandId
			b2TrySleepIsland(tls, world1, islandId)
			goto _81
		_81:
			;
			islandIndex = islandIndex - int32(1)
		}
		b2ValidateSolverSets(tls, world1)
		(*b2World)(unsafe.Pointer(world1)).Profile.SleepIslands = b2GetMilliseconds(tls, sleepTicks)
	}
}

func b2DestroySolverSet(tls *_Stack, world uintptr, setIndex int32) {
	var set uintptr
	var v1 bool
	_, _ = set, v1
	if v1 = !!(0 <= setIndex && setIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+6188, __ccgo_ts+26452, int32FromInt32(18))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(setIndex)*88
	b2Free(tls, (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data, uint64FromInt32((*b2SolverSet)(unsafe.Pointer(set)).BodySims.Capacity)*uint64(96))
	(*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data = uintptrFromInt32(0)
	(*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count = 0
	(*b2SolverSet)(unsafe.Pointer(set)).BodySims.Capacity = 0
	b2Free(tls, (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Data, uint64FromInt32((*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Capacity)*uint64(32))
	(*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Data = uintptrFromInt32(0)
	(*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Count = 0
	(*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Capacity = 0
	b2Free(tls, (*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Data, uint64FromInt32((*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Capacity)*uint64(200))
	(*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Data = uintptrFromInt32(0)
	(*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Count = 0
	(*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Capacity = 0
	b2Free(tls, (*b2SolverSet)(unsafe.Pointer(set)).JointSims.Data, uint64FromInt32((*b2SolverSet)(unsafe.Pointer(set)).JointSims.Capacity)*uint64(252))
	(*b2SolverSet)(unsafe.Pointer(set)).JointSims.Data = uintptrFromInt32(0)
	(*b2SolverSet)(unsafe.Pointer(set)).JointSims.Count = 0
	(*b2SolverSet)(unsafe.Pointer(set)).JointSims.Capacity = 0
	b2Free(tls, (*b2SolverSet)(unsafe.Pointer(set)).IslandSims.Data, uint64FromInt32((*b2SolverSet)(unsafe.Pointer(set)).IslandSims.Capacity)*uint64(4))
	(*b2SolverSet)(unsafe.Pointer(set)).IslandSims.Data = uintptrFromInt32(0)
	(*b2SolverSet)(unsafe.Pointer(set)).IslandSims.Count = 0
	(*b2SolverSet)(unsafe.Pointer(set)).IslandSims.Capacity = 0
	b2FreeId(tls, world+1952, setIndex)
	*(*b2SolverSet)(unsafe.Add(nil, set)) = b2SolverSet{}
	(*b2SolverSet)(unsafe.Pointer(set)).SetIndex = -int32(1)
}

// C documentation
//
//	// Wake a solver set. Does not merge islands.
//	// Contacts can be in several places:
//	// 1. non-touching contacts in the disabled set
//	// 2. non-touching contacts already in the awake set
//	// 3. touching contacts in the sleeping set
//	// This handles contact types 1 and 3. Type 2 doesn't need any action.
func b2WakeSolverSet(tls *_Stack, world uintptr, setIndex int32) {
	var awakeContactSim, awakeContactSim1, awakeSet, bodies, body, contact, contact1, contactSim, contactSim1, disabledSet, island, islandDst, islandSrc, joint, jointSim, movedContact, movedContactSim, set, simDst, simSrc, state, v12, v15, v7, v8, v9 uintptr
	var bodyCount, contactCount, contactId, contactKey, edgeIndex, encodedBodySimIndex, i, i1, i2, i3, islandCount, jointCount, localIndex, movedLocalIndex, newCapacity, newSize, oldCapacity, oldSize, v10, v11, v14 int32
	var v1, v2, v3, v4 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = awakeContactSim, awakeContactSim1, awakeSet, bodies, body, bodyCount, contact, contact1, contactCount, contactId, contactKey, contactSim, contactSim1, disabledSet, edgeIndex, encodedBodySimIndex, i, i1, i2, i3, island, islandCount, islandDst, islandSrc, joint, jointCount, jointSim, localIndex, movedContact, movedContactSim, movedLocalIndex, newCapacity, newSize, oldCapacity, oldSize, set, simDst, simSrc, state, v1, v10, v11, v12, v14, v15, v2, v3, v4, v7, v8, v9
	if v1 = !!(setIndex >= int32(b2_firstSleepingSet)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+26480, __ccgo_ts+26452, int32FromInt32(37))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(0 <= setIndex && setIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+6188, __ccgo_ts+26452, int32FromInt32(38))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(setIndex)*88
	if v2 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+26452, int32FromInt32(39))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	awakeSet = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
	if v3 = !!(bool(0 <= int32(b2_disabledSet)) && int32(b2_disabledSet) < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+4171, __ccgo_ts+26452, int32FromInt32(40))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	disabledSet = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b2_disabledSet))*88
	bodies = (*b2World)(unsafe.Pointer(world)).Bodies.Data
	bodyCount = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count
	i = 0
	for {
		if !(i < bodyCount) {
			break
		}
		simSrc = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr(i)*96
		body = bodies + uintptr((*b2BodySim)(unsafe.Pointer(simSrc)).BodyId)*104
		if v4 = !!((*b2Body)(unsafe.Pointer(body)).SetIndex == setIndex); !v4 {
			b2InternalAssert(tls, __ccgo_ts+26512, __ccgo_ts+26452, int32FromInt32(50))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		(*b2Body)(unsafe.Pointer(body)).SetIndex = int32(b2_awakeSet)
		(*b2Body)(unsafe.Pointer(body)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(awakeSet)).BodySims.Count
		// Reset sleep timer
		(*b2Body)(unsafe.Pointer(body)).SleepTime = float32FromFloat32(0)
		v7 = awakeSet
		v8 = awakeSet + 8
		v9 = awakeSet + 12
		v10 = int32(96)
		if *(*int32)(unsafe.Add(nil, v8)) >= *(*int32)(unsafe.Add(nil, v9)) {
			oldCapacity = *(*int32)(unsafe.Add(nil, v9))
			oldSize = oldCapacity * v10
			if oldCapacity == 0 {
				v11 = int32(8)
			} else {
				v11 = int32(2) * oldCapacity
			}
			newCapacity = v11
			newSize = newCapacity * v10
			*(*uintptr)(unsafe.Add(nil, v7)) = b2GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v7)), uint64FromInt32(oldSize), uint64FromInt32(newSize))
			*(*int32)(unsafe.Add(nil, v9)) = newCapacity
		}
		v15 = v8
		v14 = *(*int32)(unsafe.Pointer(v15))
		*(*int32)(unsafe.Pointer(v15)) = *(*int32)(unsafe.Pointer(v15)) + 1
		v12 = *(*uintptr)(unsafe.Add(nil, v7)) + uintptr(v14*v10)
		goto _13
	_13:
		simDst = v12
		memcpy(tls, simDst, simSrc, uint64(96))
		v7 = awakeSet + 16
		v8 = awakeSet + 16 + 8
		v9 = awakeSet + 16 + 12
		v10 = int32(32)
		if *(*int32)(unsafe.Add(nil, v8)) >= *(*int32)(unsafe.Add(nil, v9)) {
			oldCapacity = *(*int32)(unsafe.Add(nil, v9))
			oldSize = oldCapacity * v10
			if oldCapacity == 0 {
				v11 = int32(8)
			} else {
				v11 = int32(2) * oldCapacity
			}
			newCapacity = v11
			newSize = newCapacity * v10
			*(*uintptr)(unsafe.Add(nil, v7)) = b2GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v7)), uint64FromInt32(oldSize), uint64FromInt32(newSize))
			*(*int32)(unsafe.Add(nil, v9)) = newCapacity
		}
		v15 = v8
		v14 = *(*int32)(unsafe.Pointer(v15))
		*(*int32)(unsafe.Pointer(v15)) = *(*int32)(unsafe.Pointer(v15)) + 1
		v12 = *(*uintptr)(unsafe.Add(nil, v7)) + uintptr(v14*v10)
		goto _22
	_22:
		state = v12
		*(*b2BodyState)(unsafe.Add(nil, state)) = b2_identityBodyState21
		(*b2BodyState)(unsafe.Pointer(state)).Flags = (*b2Body)(unsafe.Pointer(body)).Flags
		// Move non-touching contacts from disabled set to awake set.
		encodedBodySimIndex = (*b2Body)(unsafe.Pointer(body)).LocalIndex
		contactKey = (*b2Body)(unsafe.Pointer(body)).HeadContactKey
		for contactKey != -int32(1) {
			edgeIndex = contactKey & int32(1)
			contactId = contactKey >> int32(1)
			if v1 = !!(0 <= contactId && contactId < (*b2World)(unsafe.Pointer(world)).Contacts.Count); !v1 {
				b2InternalAssert(tls, __ccgo_ts+301, __ccgo_ts+26452, int32FromInt32(72))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			contact = (*b2World)(unsafe.Pointer(world)).Contacts.Data + uintptr(contactId)*64
			contactKey = (*(*b2ContactEdge)(unsafe.Add(nil, contact+uintptr(edgeIndex)*12))).NextKey
			if (*b2Contact)(unsafe.Pointer(contact)).SetIndex != int32(b2_disabledSet) {
				if v1 = !!((*b2Contact)(unsafe.Pointer(contact)).SetIndex == int32(b2_awakeSet) || (*b2Contact)(unsafe.Pointer(contact)).SetIndex == setIndex); !v1 {
					b2InternalAssert(tls, __ccgo_ts+26539, __ccgo_ts+26452, int32FromInt32(78))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				// If this contact is awake but not touching I need to fixup the body sim index here.
				if (*b2Contact)(unsafe.Pointer(contact)).SetIndex == int32(b2_awakeSet) {
					if v1 = !!((*b2Contact)(unsafe.Pointer(contact)).ColorIndex == -int32FromInt32(1)); !v1 {
						b2InternalAssert(tls, __ccgo_ts+17758, __ccgo_ts+26452, int32FromInt32(83))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					if v1 = !!(0 <= (*b2Contact)(unsafe.Pointer(contact)).LocalIndex && (*b2Contact)(unsafe.Pointer(contact)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(awakeSet)).ContactSims.Count); !v1 {
						b2InternalAssert(tls, __ccgo_ts+26605, __ccgo_ts+26452, int32FromInt32(84))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					awakeContactSim = (*b2SolverSet)(unsafe.Pointer(awakeSet)).ContactSims.Data + uintptr((*b2Contact)(unsafe.Pointer(contact)).LocalIndex)*200
					if edgeIndex == 0 {
						(*b2ContactSim)(unsafe.Pointer(awakeContactSim)).EncodedBodySimA = encodedBodySimIndex
					} else {
						(*b2ContactSim)(unsafe.Pointer(awakeContactSim)).EncodedBodySimB = encodedBodySimIndex
					}
				}
				continue
			}
			localIndex = (*b2Contact)(unsafe.Pointer(contact)).LocalIndex
			if v1 = !!(0 <= localIndex && localIndex < (*b2SolverSet)(unsafe.Pointer(disabledSet)).ContactSims.Count); !v1 {
				b2InternalAssert(tls, __ccgo_ts+26695, __ccgo_ts+26452, int32FromInt32(99))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			contactSim = (*b2SolverSet)(unsafe.Pointer(disabledSet)).ContactSims.Data + uintptr(localIndex)*200
			if v2 = !!((*b2Contact)(unsafe.Pointer(contact)).Flags&uint32(b2_contactTouchingFlag) == uint32(0) && (*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.PointCount == 0); !v2 {
				b2InternalAssert(tls, __ccgo_ts+26770, __ccgo_ts+26452, int32FromInt32(101))
			}
			_ = v2 || bool(int32FromInt32(0) != 0)
			(*b2Contact)(unsafe.Pointer(contact)).SetIndex = int32(b2_awakeSet)
			(*b2Contact)(unsafe.Pointer(contact)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(awakeSet)).ContactSims.Count
			v7 = awakeSet + 48
			v8 = awakeSet + 48 + 8
			v9 = awakeSet + 48 + 12
			v10 = int32(200)
			if *(*int32)(unsafe.Add(nil, v8)) >= *(*int32)(unsafe.Add(nil, v9)) {
				oldCapacity = *(*int32)(unsafe.Add(nil, v9))
				oldSize = oldCapacity * v10
				if oldCapacity == 0 {
					v11 = int32(8)
				} else {
					v11 = int32(2) * oldCapacity
				}
				newCapacity = v11
				newSize = newCapacity * v10
				*(*uintptr)(unsafe.Add(nil, v7)) = b2GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v7)), uint64FromInt32(oldSize), uint64FromInt32(newSize))
				*(*int32)(unsafe.Add(nil, v9)) = newCapacity
			}
			v15 = v8
			v14 = *(*int32)(unsafe.Pointer(v15))
			*(*int32)(unsafe.Pointer(v15)) = *(*int32)(unsafe.Pointer(v15)) + 1
			v12 = *(*uintptr)(unsafe.Add(nil, v7)) + uintptr(v14*v10)
			goto _37
		_37:
			awakeContactSim1 = v12
			memcpy(tls, awakeContactSim1, contactSim, uint64(200))
			if edgeIndex == 0 {
				(*b2ContactSim)(unsafe.Pointer(awakeContactSim1)).EncodedBodySimA = encodedBodySimIndex
			} else {
				(*b2ContactSim)(unsafe.Pointer(awakeContactSim1)).EncodedBodySimB = encodedBodySimIndex
			}
			v7 = (*b2SolverSet)(unsafe.Pointer(disabledSet)).ContactSims.Data
			v8 = disabledSet + 48 + 8
			v10 = localIndex
			v11 = int32(200)
			if v1 = !!(0 <= v10 && v10 < *(*int32)(unsafe.Add(nil, v8)) && bool(__ccgo_ts+2029 != 0)); !v1 {
				b2InternalAssert(tls, __ccgo_ts+2055, __ccgo_ts+2115, int32FromInt32(153))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			*(*int32)(unsafe.Add(nil, v8)) = *(*int32)(unsafe.Add(nil, v8)) - 1
			if v10 != *(*int32)(unsafe.Add(nil, v8)) {
				memcpy(tls, v7+uintptr(v10*v11), v7+uintptr(*(*int32)(unsafe.Add(nil, v8))*v11), uint64FromInt32(v11))
				v14 = *(*int32)(unsafe.Add(nil, v8))
				goto _46
			}
			v14 = -int32(1)
			goto _46
		_46:
			movedLocalIndex = v14
			if movedLocalIndex != -int32(1) {
				// fix moved element
				movedContactSim = (*b2SolverSet)(unsafe.Pointer(disabledSet)).ContactSims.Data + uintptr(localIndex)*200
				if v2 = !!(0 <= (*b2ContactSim)(unsafe.Pointer(movedContactSim)).ContactId && (*b2ContactSim)(unsafe.Pointer(movedContactSim)).ContactId < (*b2World)(unsafe.Pointer(world)).Contacts.Count); !v2 {
					b2InternalAssert(tls, __ccgo_ts+6949, __ccgo_ts+26452, int32FromInt32(122))
				}
				_ = v2 || bool(int32FromInt32(0) != 0)
				movedContact = (*b2World)(unsafe.Pointer(world)).Contacts.Data + uintptr((*b2ContactSim)(unsafe.Pointer(movedContactSim)).ContactId)*64
				if v3 = !!((*b2Contact)(unsafe.Pointer(movedContact)).LocalIndex == movedLocalIndex); !v3 {
					b2InternalAssert(tls, __ccgo_ts+26859, __ccgo_ts+26452, int32FromInt32(123))
				}
				_ = v3 || bool(int32FromInt32(0) != 0)
				(*b2Contact)(unsafe.Pointer(movedContact)).LocalIndex = localIndex
			}
		}
		goto _5
	_5:
		;
		i = i + 1
	}
	// transfer touching contacts from sleeping set to contact graph
	contactCount = (*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Count
	i1 = 0
	for {
		if !(i1 < contactCount) {
			break
		}
		contactSim1 = (*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Data + uintptr(i1)*200
		if v1 = !!(0 <= (*b2ContactSim)(unsafe.Pointer(contactSim1)).ContactId && (*b2ContactSim)(unsafe.Pointer(contactSim1)).ContactId < (*b2World)(unsafe.Pointer(world)).Contacts.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+26903, __ccgo_ts+26452, int32FromInt32(135))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		contact1 = (*b2World)(unsafe.Pointer(world)).Contacts.Data + uintptr((*b2ContactSim)(unsafe.Pointer(contactSim1)).ContactId)*64
		if v2 = !!((*b2Contact)(unsafe.Pointer(contact1)).Flags&uint32(b2_contactTouchingFlag) != 0); !v2 {
			b2InternalAssert(tls, __ccgo_ts+5161, __ccgo_ts+26452, int32FromInt32(136))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		if v1 = !!((*b2ContactSim)(unsafe.Pointer(contactSim1)).SimFlags&uint32(b2_simTouchingFlag) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+5119, __ccgo_ts+26452, int32FromInt32(137))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		if v1 = !!((*b2ContactSim)(unsafe.Pointer(contactSim1)).Manifold.PointCount > int32FromInt32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+5083, __ccgo_ts+26452, int32FromInt32(138))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		if v1 = !!((*b2Contact)(unsafe.Pointer(contact1)).SetIndex == setIndex); !v1 {
			b2InternalAssert(tls, __ccgo_ts+26991, __ccgo_ts+26452, int32FromInt32(139))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		b2AddContactToGraph(tls, world, contactSim1, contact1)
		(*b2Contact)(unsafe.Pointer(contact1)).SetIndex = int32(b2_awakeSet)
		goto _49
	_49:
		;
		i1 = i1 + 1
	}
	// transfer joints from sleeping set to awake set
	jointCount = (*b2SolverSet)(unsafe.Pointer(set)).JointSims.Count
	i2 = 0
	for {
		if !(i2 < jointCount) {
			break
		}
		jointSim = (*b2SolverSet)(unsafe.Pointer(set)).JointSims.Data + uintptr(i2)*252
		if v1 = !!(0 <= (*b2JointSim)(unsafe.Pointer(jointSim)).JointId && (*b2JointSim)(unsafe.Pointer(jointSim)).JointId < (*b2World)(unsafe.Pointer(world)).Joints.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+27021, __ccgo_ts+26452, int32FromInt32(151))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		joint = (*b2World)(unsafe.Pointer(world)).Joints.Data + uintptr((*b2JointSim)(unsafe.Pointer(jointSim)).JointId)*72
		if v2 = !!((*b2Joint)(unsafe.Pointer(joint)).SetIndex == setIndex); !v2 {
			b2InternalAssert(tls, __ccgo_ts+13269, __ccgo_ts+26452, int32FromInt32(152))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		b2AddJointToGraph(tls, world, jointSim, joint)
		(*b2Joint)(unsafe.Pointer(joint)).SetIndex = int32(b2_awakeSet)
		goto _55
	_55:
		;
		i2 = i2 + 1
	}
	// transfer island from sleeping set to awake set
	// Usually a sleeping set has only one island, but it is possible
	// that joints are created between sleeping islands and they
	// are moved to the same sleeping set.
	islandCount = (*b2SolverSet)(unsafe.Pointer(set)).IslandSims.Count
	i3 = 0
	for {
		if !(i3 < islandCount) {
			break
		}
		islandSrc = (*b2SolverSet)(unsafe.Pointer(set)).IslandSims.Data + uintptr(i3)*4
		if v1 = !!(0 <= (*b2IslandSim)(unsafe.Pointer(islandSrc)).IslandId && (*b2IslandSim)(unsafe.Pointer(islandSrc)).IslandId < (*b2World)(unsafe.Pointer(world)).Islands.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+27099, __ccgo_ts+26452, int32FromInt32(167))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		island = (*b2World)(unsafe.Pointer(world)).Islands.Data + uintptr((*b2IslandSim)(unsafe.Pointer(islandSrc)).IslandId)*64
		(*b2Island)(unsafe.Pointer(island)).SetIndex = int32(b2_awakeSet)
		(*b2Island)(unsafe.Pointer(island)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(awakeSet)).IslandSims.Count
		v7 = awakeSet + 64
		v8 = awakeSet + 64 + 8
		v9 = awakeSet + 64 + 12
		v10 = int32(4)
		if *(*int32)(unsafe.Add(nil, v8)) >= *(*int32)(unsafe.Add(nil, v9)) {
			oldCapacity = *(*int32)(unsafe.Add(nil, v9))
			oldSize = oldCapacity * v10
			if oldCapacity == 0 {
				v11 = int32(8)
			} else {
				v11 = int32(2) * oldCapacity
			}
			newCapacity = v11
			newSize = newCapacity * v10
			*(*uintptr)(unsafe.Add(nil, v7)) = b2GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v7)), uint64FromInt32(oldSize), uint64FromInt32(newSize))
			*(*int32)(unsafe.Add(nil, v9)) = newCapacity
		}
		v15 = v8
		v14 = *(*int32)(unsafe.Pointer(v15))
		*(*int32)(unsafe.Pointer(v15)) = *(*int32)(unsafe.Pointer(v15)) + 1
		v12 = *(*uintptr)(unsafe.Add(nil, v7)) + uintptr(v14*v10)
		goto _66
	_66:
		islandDst = v12
		memcpy(tls, islandDst, islandSrc, uint64(4))
		goto _58
	_58:
		;
		i3 = i3 + 1
	}
	// destroy the sleeping set
	b2DestroySolverSet(tls, world, setIndex)
}

// C documentation
//
//	// This is called when joints are created between sets. I want to allow the sets
//	// to continue sleeping if both are asleep. Otherwise one set is waked.
//	// Islands will get merge when the set is waked.
func b2MergeSolverSets(tls *_Stack, world uintptr, setId1 int32, setId2 int32) {
	var bodies, body, contact, contactDst, contactSrc, island, islandDst, islandSrc, joint, jointDst, jointSrc, set1, set2, simDst, simSrc, tempSet, v12, v15, v7, v8, v9 uintptr
	var bodyCount, contactCount, i, i1, i2, i3, islandCount, islandId, jointCount, newCapacity, newSize, oldCapacity, oldSize, tempId, v10, v11, v14 int32
	var v1, v2 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodies, body, bodyCount, contact, contactCount, contactDst, contactSrc, i, i1, i2, i3, island, islandCount, islandDst, islandId, islandSrc, joint, jointCount, jointDst, jointSrc, newCapacity, newSize, oldCapacity, oldSize, set1, set2, simDst, simSrc, tempId, tempSet, v1, v10, v11, v12, v14, v15, v2, v7, v8, v9
	if v1 = !!(setId1 >= int32(b2_firstSleepingSet)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+27951, __ccgo_ts+26452, int32FromInt32(460))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(setId2 >= int32(b2_firstSleepingSet)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+27981, __ccgo_ts+26452, int32FromInt32(461))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(0 <= setId1 && setId1 < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+28011, __ccgo_ts+26452, int32FromInt32(462))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	set1 = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(setId1)*88
	if v2 = !!(0 <= setId2 && setId2 < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+28071, __ccgo_ts+26452, int32FromInt32(463))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	set2 = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(setId2)*88
	// Move the fewest number of bodies
	if (*b2SolverSet)(unsafe.Pointer(set1)).BodySims.Count < (*b2SolverSet)(unsafe.Pointer(set2)).BodySims.Count {
		tempSet = set1
		set1 = set2
		set2 = tempSet
		tempId = setId1
		setId1 = setId2
		setId2 = tempId
	}
	// transfer bodies
	bodies = (*b2World)(unsafe.Pointer(world)).Bodies.Data
	bodyCount = (*b2SolverSet)(unsafe.Pointer(set2)).BodySims.Count
	i = 0
	for {
		if !(i < bodyCount) {
			break
		}
		simSrc = (*b2SolverSet)(unsafe.Pointer(set2)).BodySims.Data + uintptr(i)*96
		body = bodies + uintptr((*b2BodySim)(unsafe.Pointer(simSrc)).BodyId)*104
		if v1 = !!((*b2Body)(unsafe.Pointer(body)).SetIndex == setId2); !v1 {
			b2InternalAssert(tls, __ccgo_ts+28131, __ccgo_ts+26452, int32FromInt32(486))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		(*b2Body)(unsafe.Pointer(body)).SetIndex = setId1
		(*b2Body)(unsafe.Pointer(body)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(set1)).BodySims.Count
		v7 = set1
		v8 = set1 + 8
		v9 = set1 + 12
		v10 = int32(96)
		if *(*int32)(unsafe.Add(nil, v8)) >= *(*int32)(unsafe.Add(nil, v9)) {
			oldCapacity = *(*int32)(unsafe.Add(nil, v9))
			oldSize = oldCapacity * v10
			if oldCapacity == 0 {
				v11 = int32(8)
			} else {
				v11 = int32(2) * oldCapacity
			}
			newCapacity = v11
			newSize = newCapacity * v10
			*(*uintptr)(unsafe.Add(nil, v7)) = b2GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v7)), uint64FromInt32(oldSize), uint64FromInt32(newSize))
			*(*int32)(unsafe.Add(nil, v9)) = newCapacity
		}
		v15 = v8
		v14 = *(*int32)(unsafe.Pointer(v15))
		*(*int32)(unsafe.Pointer(v15)) = *(*int32)(unsafe.Pointer(v15)) + 1
		v12 = *(*uintptr)(unsafe.Add(nil, v7)) + uintptr(v14*v10)
		goto _13
	_13:
		simDst = v12
		memcpy(tls, simDst, simSrc, uint64(96))
		goto _5
	_5:
		;
		i = i + 1
	}
	// transfer contacts
	contactCount = (*b2SolverSet)(unsafe.Pointer(set2)).ContactSims.Count
	i1 = 0
	for {
		if !(i1 < contactCount) {
			break
		}
		contactSrc = (*b2SolverSet)(unsafe.Pointer(set2)).ContactSims.Data + uintptr(i1)*200
		if v1 = !!(0 <= (*b2ContactSim)(unsafe.Pointer(contactSrc)).ContactId && (*b2ContactSim)(unsafe.Pointer(contactSrc)).ContactId < (*b2World)(unsafe.Pointer(world)).Contacts.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+28156, __ccgo_ts+26452, int32FromInt32(502))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		contact = (*b2World)(unsafe.Pointer(world)).Contacts.Data + uintptr((*b2ContactSim)(unsafe.Pointer(contactSrc)).ContactId)*64
		if v2 = !!((*b2Contact)(unsafe.Pointer(contact)).SetIndex == setId2); !v2 {
			b2InternalAssert(tls, __ccgo_ts+28244, __ccgo_ts+26452, int32FromInt32(503))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		(*b2Contact)(unsafe.Pointer(contact)).SetIndex = setId1
		(*b2Contact)(unsafe.Pointer(contact)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(set1)).ContactSims.Count
		v7 = set1 + 48
		v8 = set1 + 48 + 8
		v9 = set1 + 48 + 12
		v10 = int32(200)
		if *(*int32)(unsafe.Add(nil, v8)) >= *(*int32)(unsafe.Add(nil, v9)) {
			oldCapacity = *(*int32)(unsafe.Add(nil, v9))
			oldSize = oldCapacity * v10
			if oldCapacity == 0 {
				v11 = int32(8)
			} else {
				v11 = int32(2) * oldCapacity
			}
			newCapacity = v11
			newSize = newCapacity * v10
			*(*uintptr)(unsafe.Add(nil, v7)) = b2GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v7)), uint64FromInt32(oldSize), uint64FromInt32(newSize))
			*(*int32)(unsafe.Add(nil, v9)) = newCapacity
		}
		v15 = v8
		v14 = *(*int32)(unsafe.Pointer(v15))
		*(*int32)(unsafe.Pointer(v15)) = *(*int32)(unsafe.Pointer(v15)) + 1
		v12 = *(*uintptr)(unsafe.Add(nil, v7)) + uintptr(v14*v10)
		goto _25
	_25:
		contactDst = v12
		memcpy(tls, contactDst, contactSrc, uint64(200))
		goto _16
	_16:
		;
		i1 = i1 + 1
	}
	// transfer joints
	jointCount = (*b2SolverSet)(unsafe.Pointer(set2)).JointSims.Count
	i2 = 0
	for {
		if !(i2 < jointCount) {
			break
		}
		jointSrc = (*b2SolverSet)(unsafe.Pointer(set2)).JointSims.Data + uintptr(i2)*252
		if v1 = !!(0 <= (*b2JointSim)(unsafe.Pointer(jointSrc)).JointId && (*b2JointSim)(unsafe.Pointer(jointSrc)).JointId < (*b2World)(unsafe.Pointer(world)).Joints.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+28272, __ccgo_ts+26452, int32FromInt32(519))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		joint = (*b2World)(unsafe.Pointer(world)).Joints.Data + uintptr((*b2JointSim)(unsafe.Pointer(jointSrc)).JointId)*72
		if v2 = !!((*b2Joint)(unsafe.Pointer(joint)).SetIndex == setId2); !v2 {
			b2InternalAssert(tls, __ccgo_ts+28350, __ccgo_ts+26452, int32FromInt32(520))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		(*b2Joint)(unsafe.Pointer(joint)).SetIndex = setId1
		(*b2Joint)(unsafe.Pointer(joint)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(set1)).JointSims.Count
		v7 = set1 + 32
		v8 = set1 + 32 + 8
		v9 = set1 + 32 + 12
		v10 = int32(252)
		if *(*int32)(unsafe.Add(nil, v8)) >= *(*int32)(unsafe.Add(nil, v9)) {
			oldCapacity = *(*int32)(unsafe.Add(nil, v9))
			oldSize = oldCapacity * v10
			if oldCapacity == 0 {
				v11 = int32(8)
			} else {
				v11 = int32(2) * oldCapacity
			}
			newCapacity = v11
			newSize = newCapacity * v10
			*(*uintptr)(unsafe.Add(nil, v7)) = b2GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v7)), uint64FromInt32(oldSize), uint64FromInt32(newSize))
			*(*int32)(unsafe.Add(nil, v9)) = newCapacity
		}
		v15 = v8
		v14 = *(*int32)(unsafe.Pointer(v15))
		*(*int32)(unsafe.Pointer(v15)) = *(*int32)(unsafe.Pointer(v15)) + 1
		v12 = *(*uintptr)(unsafe.Add(nil, v7)) + uintptr(v14*v10)
		goto _37
	_37:
		jointDst = v12
		memcpy(tls, jointDst, jointSrc, uint64(252))
		goto _28
	_28:
		;
		i2 = i2 + 1
	}
	// transfer islands
	islandCount = (*b2SolverSet)(unsafe.Pointer(set2)).IslandSims.Count
	i3 = 0
	for {
		if !(i3 < islandCount) {
			break
		}
		islandSrc = (*b2SolverSet)(unsafe.Pointer(set2)).IslandSims.Data + uintptr(i3)*4
		islandId = (*b2IslandSim)(unsafe.Pointer(islandSrc)).IslandId
		if v1 = !!(0 <= islandId && islandId < (*b2World)(unsafe.Pointer(world)).Islands.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+962, __ccgo_ts+26452, int32FromInt32(537))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		island = (*b2World)(unsafe.Pointer(world)).Islands.Data + uintptr(islandId)*64
		(*b2Island)(unsafe.Pointer(island)).SetIndex = setId1
		(*b2Island)(unsafe.Pointer(island)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(set1)).IslandSims.Count
		v7 = set1 + 64
		v8 = set1 + 64 + 8
		v9 = set1 + 64 + 12
		v10 = int32(4)
		if *(*int32)(unsafe.Add(nil, v8)) >= *(*int32)(unsafe.Add(nil, v9)) {
			oldCapacity = *(*int32)(unsafe.Add(nil, v9))
			oldSize = oldCapacity * v10
			if oldCapacity == 0 {
				v11 = int32(8)
			} else {
				v11 = int32(2) * oldCapacity
			}
			newCapacity = v11
			newSize = newCapacity * v10
			*(*uintptr)(unsafe.Add(nil, v7)) = b2GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v7)), uint64FromInt32(oldSize), uint64FromInt32(newSize))
			*(*int32)(unsafe.Add(nil, v9)) = newCapacity
		}
		v15 = v8
		v14 = *(*int32)(unsafe.Pointer(v15))
		*(*int32)(unsafe.Pointer(v15)) = *(*int32)(unsafe.Pointer(v15)) + 1
		v12 = *(*uintptr)(unsafe.Add(nil, v7)) + uintptr(v14*v10)
		goto _48
	_48:
		islandDst = v12
		memcpy(tls, islandDst, islandSrc, uint64(4))
		goto _40
	_40:
		;
		i3 = i3 + 1
	}
	// destroy the merged set
	b2DestroySolverSet(tls, world, setId2)
	b2ValidateSolverSets(tls, world)
}
