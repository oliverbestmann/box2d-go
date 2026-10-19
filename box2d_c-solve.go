package b2

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

func b2SolveOverflowContacts(tls *_Stack, context uintptr, useBias uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var P, P1, dp, ds, normal, rA, rA1, rB, rB1, tangent, vA, vB, vrA, vrA1, vrB, vrB1, v10, v12, v13, v17, v18, v21, v22, v24, v25, v26, v28, v29, v30, v32, v33, v8, v9 Vec2
	var awakeSet, color, constraint, constraints, cp, cp1, graph, stateA, stateB, states, world, v1, v3, v6, v7 uintptr
	var contactCount, i, j, j1, pointCount, v2 int32
	var deltaLambda, friction, iA, iB, impulse, impulse1, impulseScale, inv_h, lambda, mA, mB, massScale, maxFriction, maxLambda, newImpulse, newImpulse1, pushout, s4, totalNormalImpulse, velocityBias, vn, vt, wA, wB, v34, v36, v37, v38, v40, v41, v49, v63, v65 float32
	var dqA, dqB, v16, v20 Rot
	var softness b2Softness
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = P, P1, awakeSet, color, constraint, constraints, contactCount, cp, cp1, deltaLambda, dp, dqA, dqB, ds, friction, graph, i, iA, iB, impulse, impulse1, impulseScale, inv_h, j, j1, lambda, mA, mB, massScale, maxFriction, maxLambda, newImpulse, newImpulse1, normal, pointCount, pushout, rA, rA1, rB, rB1, s4, softness, stateA, stateB, states, tangent, totalNormalImpulse, vA, vB, velocityBias, vn, vrA, vrA1, vrB, vrB1, vt, wA, wB, world, v1, v10, v12, v13, v16, v17, v18, v2, v20, v21, v22, v24, v25, v26, v28, v29, v3, v30, v32, v33, v34, v36, v37, v38, v40, v41, v49, v6, v63, v65, v7, v8, v9
	graph = (*b2StepContext)(unsafe.Pointer(context)).Graph
	color = graph + uintptr(int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1))*56
	constraints = *(*uintptr)(unsafe.Add(unsafe.Pointer(color), 48))
	contactCount = (*b2GraphColor)(unsafe.Pointer(color)).ContactSims.Count
	world = (*b2StepContext)(unsafe.Pointer(context)).World
	v1 = world + 1064
	v2 = int32(b2_awakeSet)
	if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*88
	goto _4
_4:
	awakeSet = v3
	states = (*b2SolverSet)(unsafe.Pointer(awakeSet)).BodyStates.Data
	inv_h = (*b2StepContext)(unsafe.Pointer(context)).Inv_h
	pushout = (*b2World)(unsafe.Pointer((*b2StepContext)(unsafe.Pointer(context)).World)).MaxContactPushSpeed
	// This is a dummy body to represent a static body since static bodies don't have a solver body.
	*(*b2BodyState)(unsafe.Add(nil, bp)) = b2_identityBodyState4
	i = 0
	for {
		if !(i < contactCount) {
			break
		}
		constraint = constraints + uintptr(i)*160
		mA = (*b2ContactConstraint)(unsafe.Pointer(constraint)).InvMassA
		iA = (*b2ContactConstraint)(unsafe.Pointer(constraint)).InvIA
		mB = (*b2ContactConstraint)(unsafe.Pointer(constraint)).InvMassB
		iB = (*b2ContactConstraint)(unsafe.Pointer(constraint)).InvIB
		if (*b2ContactConstraint)(unsafe.Pointer(constraint)).IndexA == -int32(1) {
			v6 = bp
		} else {
			v6 = states + uintptr((*b2ContactConstraint)(unsafe.Pointer(constraint)).IndexA)*32
		}
		stateA = v6
		vA = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
		wA = (*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity
		dqA = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
		if (*b2ContactConstraint)(unsafe.Pointer(constraint)).IndexB == -int32(1) {
			v7 = bp
		} else {
			v7 = states + uintptr((*b2ContactConstraint)(unsafe.Pointer(constraint)).IndexB)*32
		}
		stateB = v7
		vB = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
		wB = (*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity
		dqB = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
		v8 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaPosition
		v9 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaPosition
		v10 = Vec2{
			X: v8.X - v9.X,
			Y: v8.Y - v9.Y,
		}
		goto _11
	_11:
		dp = v10
		normal = (*b2ContactConstraint)(unsafe.Pointer(constraint)).Normal
		v12 = normal
		v13 = Vec2{
			X: v12.Y,
			Y: -v12.X,
		}
		goto _14
	_14:
		tangent = v13
		friction = (*b2ContactConstraint)(unsafe.Pointer(constraint)).Friction
		softness = (*b2ContactConstraint)(unsafe.Pointer(constraint)).Softness
		pointCount = (*b2ContactConstraint)(unsafe.Pointer(constraint)).PointCount
		totalNormalImpulse = float32FromFloat32(0)
		// Non-penetration
		j = 0
		for {
			if !(j < pointCount) {
				break
			}
			cp = constraint + 8 + uintptr(j)*44
			// fixed anchor points
			rA = (*b2ContactConstraintPoint)(unsafe.Pointer(cp)).AnchorA
			rB = (*b2ContactConstraintPoint)(unsafe.Pointer(cp)).AnchorB
			v16 = dqB
			v17 = rB
			v18 = Vec2{
				X: float32(v16.C*v17.X) - float32(v16.S*v17.Y),
				Y: float32(v16.S*v17.X) + float32(v16.C*v17.Y),
			}
			goto _19
		_19:
			v20 = dqA
			v21 = rA
			v22 = Vec2{
				X: float32(v20.C*v21.X) - float32(v20.S*v21.Y),
				Y: float32(v20.S*v21.X) + float32(v20.C*v21.Y),
			}
			goto _23
		_23:
			v24 = v18
			v25 = v22
			v26 = Vec2{
				X: v24.X - v25.X,
				Y: v24.Y - v25.Y,
			}
			goto _27
		_27:
			v28 = dp
			v29 = v26
			v30 = Vec2{
				X: v28.X + v29.X,
				Y: v28.Y + v29.Y,
			}
			goto _31
		_31:
			// compute current separation
			// this is subject to round-off error if the anchor is far from the body center of mass
			ds = v30
			v32 = ds
			v33 = normal
			v34 = float32(v32.X*v33.X) + float32(v32.Y*v33.Y)
			goto _35
		_35:
			s4 = (*b2ContactConstraintPoint)(unsafe.Pointer(cp)).BaseSeparation + v34
			velocityBias = float32FromFloat32(0)
			massScale = float32FromFloat32(1)
			impulseScale = float32FromFloat32(0)
			if s4 > float32FromFloat32(0) {
				// speculative bias
				velocityBias = float32(s4 * inv_h)
			} else {
				if useBias != 0 {
					v34 = float32(softness.BiasRate * s4)
					v36 = -pushout
					if v34 > v36 {
						v38 = v34
					} else {
						v38 = v36
					}
					v37 = v38
					goto _39
				_39:
					velocityBias = v37
					massScale = softness.MassScale
					impulseScale = softness.ImpulseScale
				}
			}
			v34 = wA
			v8 = rA
			v9 = Vec2{
				X: float32(-v34 * v8.Y),
				Y: float32(v34 * v8.X),
			}
			goto _44
		_44:
			v10 = vA
			v12 = v9
			v13 = Vec2{
				X: v10.X + v12.X,
				Y: v10.Y + v12.Y,
			}
			goto _48
		_48:
			// relative normal velocity at contact
			vrA = v13
			v36 = wB
			v17 = rB
			v18 = Vec2{
				X: float32(-v36 * v17.Y),
				Y: float32(v36 * v17.X),
			}
			goto _52
		_52:
			v21 = vB
			v22 = v18
			v24 = Vec2{
				X: v21.X + v22.X,
				Y: v21.Y + v22.Y,
			}
			goto _56
		_56:
			vrB = v24
			v25 = vrB
			v26 = vrA
			v28 = Vec2{
				X: v25.X - v26.X,
				Y: v25.Y - v26.Y,
			}
			goto _60
		_60:
			v29 = v28
			v30 = normal
			v37 = float32(v29.X*v30.X) + float32(v29.Y*v30.Y)
			goto _64
		_64:
			vn = v37
			// incremental normal impulse
			impulse = float32(float32(-(*b2ContactConstraintPoint)(unsafe.Pointer(cp)).NormalMass*massScale)*(vn+velocityBias)) - float32(impulseScale*(*b2ContactConstraintPoint)(unsafe.Pointer(cp)).NormalImpulse)
			v38 = (*b2ContactConstraintPoint)(unsafe.Pointer(cp)).NormalImpulse + impulse
			v40 = float32FromFloat32(0)
			if v38 > v40 {
				v49 = v38
			} else {
				v49 = v40
			}
			v41 = v49
			goto _68
		_68:
			// clamp the accumulated impulse
			newImpulse = v41
			impulse = newImpulse - (*b2ContactConstraintPoint)(unsafe.Pointer(cp)).NormalImpulse
			(*b2ContactConstraintPoint)(unsafe.Pointer(cp)).NormalImpulse = newImpulse
			*(*float32)(unsafe.Add(nil, cp+32)) += newImpulse
			totalNormalImpulse = totalNormalImpulse + newImpulse
			v34 = impulse
			v8 = normal
			v9 = Vec2{
				X: float32(v34 * v8.X),
				Y: float32(v34 * v8.Y),
			}
			goto _73
		_73:
			// apply normal impulse
			P = v9
			v10 = vA
			v36 = mA
			v12 = P
			v13 = Vec2{
				X: v10.X - float32(v36*v12.X),
				Y: v10.Y - float32(v36*v12.Y),
			}
			goto _78
		_78:
			vA = v13
			v8 = rA
			v9 = P
			v34 = float32(v8.X*v9.Y) - float32(v8.Y*v9.X)
			goto _82
		_82:
			wA = wA - float32(iA*v34)
			v8 = vB
			v34 = mB
			v9 = P
			v10 = Vec2{
				X: v8.X + float32(v34*v9.X),
				Y: v8.Y + float32(v34*v9.Y),
			}
			goto _87
		_87:
			vB = v10
			v8 = rB
			v9 = P
			v34 = float32(v8.X*v9.Y) - float32(v8.Y*v9.X)
			goto _91
		_91:
			wB = wB + float32(iB*v34)
			goto _15
		_15:
			;
			j = j + 1
		}
		// Friction
		j1 = 0
		for {
			if !(j1 < pointCount) {
				break
			}
			cp1 = constraint + 8 + uintptr(j1)*44
			// fixed anchor points
			rA1 = (*b2ContactConstraintPoint)(unsafe.Pointer(cp1)).AnchorA
			rB1 = (*b2ContactConstraintPoint)(unsafe.Pointer(cp1)).AnchorB
			v34 = wB
			v8 = rB1
			v9 = Vec2{
				X: float32(-v34 * v8.Y),
				Y: float32(v34 * v8.X),
			}
			goto _96
		_96:
			v10 = vB
			v12 = v9
			v13 = Vec2{
				X: v10.X + v12.X,
				Y: v10.Y + v12.Y,
			}
			goto _100
		_100:
			// relative tangent velocity at contact
			vrB1 = v13
			v36 = wA
			v17 = rA1
			v18 = Vec2{
				X: float32(-v36 * v17.Y),
				Y: float32(v36 * v17.X),
			}
			goto _104
		_104:
			v21 = vA
			v22 = v18
			v24 = Vec2{
				X: v21.X + v22.X,
				Y: v21.Y + v22.Y,
			}
			goto _108
		_108:
			vrA1 = v24
			v25 = vrB1
			v26 = vrA1
			v28 = Vec2{
				X: v25.X - v26.X,
				Y: v25.Y - v26.Y,
			}
			goto _112
		_112:
			v29 = v28
			v30 = tangent
			v37 = float32(v29.X*v30.X) + float32(v29.Y*v30.Y)
			goto _116
		_116:
			// vt = dot(vrB - sB * tangent - (vrA + sA * tangent), tangent)
			//    = dot(vrB - vrA, tangent) - (sA + sB)
			vt = v37 - (*b2ContactConstraint)(unsafe.Pointer(constraint)).TangentSpeed
			// incremental tangent impulse
			impulse1 = float32((*b2ContactConstraintPoint)(unsafe.Pointer(cp1)).TangentMass * -vt)
			// clamp the accumulated force
			maxFriction = float32(friction * (*b2ContactConstraintPoint)(unsafe.Pointer(cp1)).NormalImpulse)
			v38 = (*b2ContactConstraintPoint)(unsafe.Pointer(cp1)).TangentImpulse + impulse1
			v40 = -maxFriction
			v41 = maxFriction
			if v38 < v40 {
				v63 = v40
			} else {
				if v38 > v41 {
					v65 = v41
				} else {
					v65 = v38
				}
				v63 = v65
			}
			v49 = v63
			goto _121
		_121:
			newImpulse1 = v49
			impulse1 = newImpulse1 - (*b2ContactConstraintPoint)(unsafe.Pointer(cp1)).TangentImpulse
			(*b2ContactConstraintPoint)(unsafe.Pointer(cp1)).TangentImpulse = newImpulse1
			v34 = impulse1
			v8 = tangent
			v9 = Vec2{
				X: float32(v34 * v8.X),
				Y: float32(v34 * v8.Y),
			}
			goto _127
		_127:
			// apply tangent impulse
			P1 = v9
			v10 = vA
			v36 = mA
			v12 = P1
			v13 = Vec2{
				X: v10.X - float32(v36*v12.X),
				Y: v10.Y - float32(v36*v12.Y),
			}
			goto _132
		_132:
			vA = v13
			v8 = rA1
			v9 = P1
			v34 = float32(v8.X*v9.Y) - float32(v8.Y*v9.X)
			goto _136
		_136:
			wA = wA - float32(iA*v34)
			v8 = vB
			v34 = mB
			v9 = P1
			v10 = Vec2{
				X: v8.X + float32(v34*v9.X),
				Y: v8.Y + float32(v34*v9.Y),
			}
			goto _141
		_141:
			vB = v10
			v8 = rB1
			v9 = P1
			v34 = float32(v8.X*v9.Y) - float32(v8.Y*v9.X)
			goto _145
		_145:
			wB = wB + float32(iB*v34)
			goto _92
		_92:
			;
			j1 = j1 + 1
		}
		// Rolling resistance
		deltaLambda = float32(-(*b2ContactConstraint)(unsafe.Pointer(constraint)).RollingMass * (wB - wA))
		lambda = (*b2ContactConstraint)(unsafe.Pointer(constraint)).RollingImpulse
		maxLambda = float32((*b2ContactConstraint)(unsafe.Pointer(constraint)).RollingResistance * totalNormalImpulse)
		v34 = lambda + deltaLambda
		v36 = -maxLambda
		v37 = maxLambda
		if v34 < v36 {
			v40 = v36
		} else {
			if v34 > v37 {
				v41 = v37
			} else {
				v41 = v34
			}
			v40 = v41
		}
		v38 = v40
		goto _150
	_150:
		(*b2ContactConstraint)(unsafe.Pointer(constraint)).RollingImpulse = v38
		deltaLambda = (*b2ContactConstraint)(unsafe.Pointer(constraint)).RollingImpulse - lambda
		wA = wA - float32(iA*deltaLambda)
		wB = wB + float32(iB*deltaLambda)
		(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
		(*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
		(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
		(*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
		goto _5
	_5:
		;
		i = i + 1
	}
}

func b2SolveContactsTask(tls *_Stack, startIndex int32, endIndex int32, context uintptr, colorIndex int32, useBias uint8) {
	bp := tls.Alloc(256)
	defer tls.Free(256)
	var Px, Px1, Px2, Px3, Py, Py1, Py2, Py3, bias, bias1, biasRate, deltaLambda, dvx, dvx1, dvx2, dvx3, dvy, dvy1, dvy2, dvy3, impulse, impulse1, impulse2, impulse3, impulseScale, inv_h, lambda, mask, mask1, massScale, maxFriction, maxFriction1, maxLambda, minBiasVel, negImpulse, negImpulse1, negImpulse2, negImpulse3, newImpulse, newImpulse1, newImpulse2, newImpulse3, oneW, pointImpulseScale, pointImpulseScale1, pointMassScale, pointMassScale1, s, s1, softBias, softBias1, specBias, specBias1, tangentX, tangentY, totalNormalImpulse, vn, vn1, vt, vt1 b2FloatW
	var c, constraints, states uintptr
	var dp, ds, ds1, rA, rA1, rA2, rA3, rB, rB1, rB2, rB3, rsA, rsA1, rsB, rsB1 b2Vec2W
	var i int32
	var _ /* bA at bp+0 */ b2BodyStateW
	var _ /* bB at bp+128 */ b2BodyStateW
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = Px, Px1, Px2, Px3, Py, Py1, Py2, Py3, bias, bias1, biasRate, c, constraints, deltaLambda, dp, ds, ds1, dvx, dvx1, dvx2, dvx3, dvy, dvy1, dvy2, dvy3, i, impulse, impulse1, impulse2, impulse3, impulseScale, inv_h, lambda, mask, mask1, massScale, maxFriction, maxFriction1, maxLambda, minBiasVel, negImpulse, negImpulse1, negImpulse2, negImpulse3, newImpulse, newImpulse1, newImpulse2, newImpulse3, oneW, pointImpulseScale, pointImpulseScale1, pointMassScale, pointMassScale1, rA, rA1, rA2, rA3, rB, rB1, rB2, rB3, rsA, rsA1, rsB, rsB1, s, s1, softBias, softBias1, specBias, specBias1, states, tangentX, tangentY, totalNormalImpulse, vn, vn1, vt, vt1
	states = (*b2StepContext)(unsafe.Pointer(context)).States
	constraints = (*(*b2GraphColor)(unsafe.Add(nil, (*b2StepContext)(unsafe.Pointer(context)).Graph+uintptr(colorIndex)*56))).ｆ__ccgo3_48.SimdConstraints
	inv_h = b2SplatW(tls, (*b2StepContext)(unsafe.Pointer(context)).Inv_h)
	minBiasVel = b2SplatW(tls, -(*b2World)(unsafe.Pointer((*b2StepContext)(unsafe.Pointer(context)).World)).ContactSpeed)
	oneW = b2SplatW(tls, float32FromFloat32(1))
	i = startIndex
	for {
		if !(i < endIndex) {
			break
		}
		c = constraints + uintptr(i)*624
		*(*b2BodyStateW)(unsafe.Add(nil, bp)) = b2GatherBodies(tls, states, c)
		*(*b2BodyStateW)(unsafe.Add(nil, bp+128)) = b2GatherBodies(tls, states, c+16)
		if useBias != 0 {
			biasRate = (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).BiasRate
			massScale = (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).MassScale
			impulseScale = (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).ImpulseScale
		} else {
			biasRate = b2ZeroW(tls)
			massScale = oneW
			impulseScale = b2ZeroW(tls)
		}
		totalNormalImpulse = b2ZeroW(tls)
		dp = b2Vec2W{
			X: b2SubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).Dp.X, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).Dp.X),
			Y: b2SubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).Dp.Y, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).Dp.Y),
		}
		// point1 non-penetration constraint
		// Fixed anchors for impulses
		rA = (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).AnchorA1
		rB = (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).AnchorB1
		// Moving anchors for current separation
		rsA = b2RotateVectorW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).Dq, rA)
		rsB = b2RotateVectorW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).Dq, rB)
		// compute current separation
		// this is subject to round-off error if the anchor is far from the body center of mass
		ds = b2Vec2W{
			X: b2AddW(tls, dp.X, b2SubW(tls, rsB.X, rsA.X)),
			Y: b2AddW(tls, dp.Y, b2SubW(tls, rsB.Y, rsA.Y)),
		}
		s = b2AddW(tls, b2DotW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Normal, ds), (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).BaseSeparation1)
		// Apply speculative bias if separation is greater than zero, otherwise apply soft constraint bias
		// The minBiasVel is meant to limit stiffness, not increase it.
		mask = b2GreaterThanW(tls, s, b2ZeroW(tls))
		specBias = b2MulW(tls, s, inv_h)
		softBias = b2MaxW(tls, b2MulW(tls, biasRate, s), minBiasVel)
		// todo try b2MaxW(softBias, specBias);
		bias = b2BlendW(tls, softBias, specBias, mask)
		pointMassScale = b2BlendW(tls, massScale, oneW, mask)
		pointImpulseScale = b2BlendW(tls, impulseScale, b2ZeroW(tls), mask)
		// Relative velocity at contact
		dvx = b2SubW(tls, b2SubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.X, b2MulW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W, rB.Y)), b2SubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.X, b2MulW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).W, rA.Y)))
		dvy = b2SubW(tls, b2AddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.Y, b2MulW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W, rB.X)), b2AddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.Y, b2MulW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).W, rA.X)))
		vn = b2AddW(tls, b2MulW(tls, dvx, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Normal.X), b2MulW(tls, dvy, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Normal.Y))
		// Compute normal impulse
		negImpulse = b2AddW(tls, b2MulW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).NormalMass1, b2MulW(tls, pointMassScale, b2AddW(tls, vn, bias))), b2MulW(tls, pointImpulseScale, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).NormalImpulse1))
		// Clamp the accumulated impulse
		newImpulse = b2MaxW(tls, b2SubW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).NormalImpulse1, negImpulse), b2ZeroW(tls))
		impulse = b2SubW(tls, newImpulse, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).NormalImpulse1)
		(*b2ContactConstraintSIMD)(unsafe.Pointer(c)).NormalImpulse1 = newImpulse
		(*b2ContactConstraintSIMD)(unsafe.Pointer(c)).TotalNormalImpulse1 = b2AddW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).TotalNormalImpulse1, newImpulse)
		totalNormalImpulse = b2AddW(tls, totalNormalImpulse, newImpulse)
		// Apply contact impulse
		Px = b2MulW(tls, impulse, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Normal.X)
		Py = b2MulW(tls, impulse, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Normal.Y)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.X = b2MulSubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.X, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassA, Px)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.Y = b2MulSubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.Y, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassA, Py)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp))).W = b2MulSubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).W, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvIA, b2SubW(tls, b2MulW(tls, rA.X, Py), b2MulW(tls, rA.Y, Px)))
		(*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.X = b2MulAddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.X, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassB, Px)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.Y = b2MulAddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.Y, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassB, Py)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W = b2MulAddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvIB, b2SubW(tls, b2MulW(tls, rB.X, Py), b2MulW(tls, rB.Y, Px)))
		// second point non-penetration constraint
		// moving anchors for current separation
		rsA1 = b2RotateVectorW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).Dq, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).AnchorA2)
		rsB1 = b2RotateVectorW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).Dq, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).AnchorB2)
		// compute current separation
		ds1 = b2Vec2W{
			X: b2AddW(tls, dp.X, b2SubW(tls, rsB1.X, rsA1.X)),
			Y: b2AddW(tls, dp.Y, b2SubW(tls, rsB1.Y, rsA1.Y)),
		}
		s1 = b2AddW(tls, b2DotW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Normal, ds1), (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).BaseSeparation2)
		mask1 = b2GreaterThanW(tls, s1, b2ZeroW(tls))
		specBias1 = b2MulW(tls, s1, inv_h)
		softBias1 = b2MaxW(tls, b2MulW(tls, biasRate, s1), minBiasVel)
		bias1 = b2BlendW(tls, softBias1, specBias1, mask1)
		pointMassScale1 = b2BlendW(tls, massScale, oneW, mask1)
		pointImpulseScale1 = b2BlendW(tls, impulseScale, b2ZeroW(tls), mask1)
		// fixed anchors for Jacobians
		rA1 = (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).AnchorA2
		rB1 = (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).AnchorB2
		// Relative velocity at contact
		dvx1 = b2SubW(tls, b2SubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.X, b2MulW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W, rB1.Y)), b2SubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.X, b2MulW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).W, rA1.Y)))
		dvy1 = b2SubW(tls, b2AddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.Y, b2MulW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W, rB1.X)), b2AddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.Y, b2MulW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).W, rA1.X)))
		vn1 = b2AddW(tls, b2MulW(tls, dvx1, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Normal.X), b2MulW(tls, dvy1, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Normal.Y))
		// Compute normal impulse
		negImpulse1 = b2AddW(tls, b2MulW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).NormalMass2, b2MulW(tls, pointMassScale1, b2AddW(tls, vn1, bias1))), b2MulW(tls, pointImpulseScale1, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).NormalImpulse2))
		// Clamp the accumulated impulse
		newImpulse1 = b2MaxW(tls, b2SubW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).NormalImpulse2, negImpulse1), b2ZeroW(tls))
		impulse1 = b2SubW(tls, newImpulse1, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).NormalImpulse2)
		(*b2ContactConstraintSIMD)(unsafe.Pointer(c)).NormalImpulse2 = newImpulse1
		(*b2ContactConstraintSIMD)(unsafe.Pointer(c)).TotalNormalImpulse2 = b2AddW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).TotalNormalImpulse2, newImpulse1)
		totalNormalImpulse = b2AddW(tls, totalNormalImpulse, newImpulse1)
		// Apply contact impulse
		Px1 = b2MulW(tls, impulse1, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Normal.X)
		Py1 = b2MulW(tls, impulse1, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Normal.Y)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.X = b2MulSubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.X, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassA, Px1)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.Y = b2MulSubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.Y, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassA, Py1)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp))).W = b2MulSubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).W, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvIA, b2SubW(tls, b2MulW(tls, rA1.X, Py1), b2MulW(tls, rA1.Y, Px1)))
		(*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.X = b2MulAddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.X, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassB, Px1)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.Y = b2MulAddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.Y, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassB, Py1)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W = b2MulAddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvIB, b2SubW(tls, b2MulW(tls, rB1.X, Py1), b2MulW(tls, rB1.Y, Px1)))
		tangentX = (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Normal.Y
		tangentY = b2SubW(tls, b2ZeroW(tls), (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Normal.X)
		// point 1 friction constraint
		// fixed anchors for Jacobians
		rA2 = (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).AnchorA1
		rB2 = (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).AnchorB1
		// Relative velocity at contact
		dvx2 = b2SubW(tls, b2SubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.X, b2MulW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W, rB2.Y)), b2SubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.X, b2MulW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).W, rA2.Y)))
		dvy2 = b2SubW(tls, b2AddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.Y, b2MulW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W, rB2.X)), b2AddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.Y, b2MulW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).W, rA2.X)))
		vt = b2AddW(tls, b2MulW(tls, dvx2, tangentX), b2MulW(tls, dvy2, tangentY))
		// Tangent speed (conveyor belt)
		vt = b2SubW(tls, vt, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).TangentSpeed)
		// Compute tangent force
		negImpulse2 = b2MulW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).TangentMass1, vt)
		// Clamp the accumulated force
		maxFriction = b2MulW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Friction, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).NormalImpulse1)
		newImpulse2 = b2SubW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).TangentImpulse1, negImpulse2)
		newImpulse2 = b2MaxW(tls, b2SubW(tls, b2ZeroW(tls), maxFriction), b2MinW(tls, newImpulse2, maxFriction))
		impulse2 = b2SubW(tls, newImpulse2, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).TangentImpulse1)
		(*b2ContactConstraintSIMD)(unsafe.Pointer(c)).TangentImpulse1 = newImpulse2
		// Apply contact impulse
		Px2 = b2MulW(tls, impulse2, tangentX)
		Py2 = b2MulW(tls, impulse2, tangentY)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.X = b2MulSubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.X, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassA, Px2)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.Y = b2MulSubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.Y, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassA, Py2)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp))).W = b2MulSubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).W, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvIA, b2SubW(tls, b2MulW(tls, rA2.X, Py2), b2MulW(tls, rA2.Y, Px2)))
		(*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.X = b2MulAddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.X, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassB, Px2)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.Y = b2MulAddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.Y, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassB, Py2)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W = b2MulAddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvIB, b2SubW(tls, b2MulW(tls, rB2.X, Py2), b2MulW(tls, rB2.Y, Px2)))
		// second point friction constraint
		// fixed anchors for Jacobians
		rA3 = (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).AnchorA2
		rB3 = (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).AnchorB2
		// Relative velocity at contact
		dvx3 = b2SubW(tls, b2SubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.X, b2MulW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W, rB3.Y)), b2SubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.X, b2MulW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).W, rA3.Y)))
		dvy3 = b2SubW(tls, b2AddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.Y, b2MulW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W, rB3.X)), b2AddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.Y, b2MulW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).W, rA3.X)))
		vt1 = b2AddW(tls, b2MulW(tls, dvx3, tangentX), b2MulW(tls, dvy3, tangentY))
		// Tangent speed (conveyor belt)
		vt1 = b2SubW(tls, vt1, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).TangentSpeed)
		// Compute tangent force
		negImpulse3 = b2MulW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).TangentMass2, vt1)
		// Clamp the accumulated force
		maxFriction1 = b2MulW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Friction, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).NormalImpulse2)
		newImpulse3 = b2SubW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).TangentImpulse2, negImpulse3)
		newImpulse3 = b2MaxW(tls, b2SubW(tls, b2ZeroW(tls), maxFriction1), b2MinW(tls, newImpulse3, maxFriction1))
		impulse3 = b2SubW(tls, newImpulse3, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).TangentImpulse2)
		(*b2ContactConstraintSIMD)(unsafe.Pointer(c)).TangentImpulse2 = newImpulse3
		// Apply contact impulse
		Px3 = b2MulW(tls, impulse3, tangentX)
		Py3 = b2MulW(tls, impulse3, tangentY)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.X = b2MulSubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.X, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassA, Px3)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.Y = b2MulSubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.Y, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassA, Py3)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp))).W = b2MulSubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).W, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvIA, b2SubW(tls, b2MulW(tls, rA3.X, Py3), b2MulW(tls, rA3.Y, Px3)))
		(*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.X = b2MulAddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.X, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassB, Px3)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.Y = b2MulAddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.Y, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassB, Py3)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W = b2MulAddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvIB, b2SubW(tls, b2MulW(tls, rB3.X, Py3), b2MulW(tls, rB3.Y, Px3)))
		// Rolling resistance
		deltaLambda = b2MulW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).RollingMass, b2SubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).W, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W))
		lambda = (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).RollingImpulse
		maxLambda = b2MulW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).RollingResistance, totalNormalImpulse)
		(*b2ContactConstraintSIMD)(unsafe.Pointer(c)).RollingImpulse = b2SymClampW(tls, b2AddW(tls, lambda, deltaLambda), maxLambda)
		deltaLambda = b2SubW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).RollingImpulse, lambda)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp))).W = b2MulSubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).W, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvIA, deltaLambda)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W = b2MulAddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvIB, deltaLambda)
		b2ScatterBodies(tls, states, c, bp)
		b2ScatterBodies(tls, states, c+16, bp+128)
		goto _1
	_1:
		;
		i = i + 1
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
		Translation:    delta,
		IterationCount: iteration,
	}
}

// C documentation
//
//	// Continuous collision of dynamic versus static
func b2SolveContinuous(tls *_Stack, world uintptr, bodySimIndex int32) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var aabb, box, box1, box2, c, fatAABB, fatAABB1, v44, v45, v66 AABB
	var aabbMargin, invMag, mag, omt, speculativeDistance, x, y, v46, v47, v48, v50, v51, v52, v53, v55, v56, v57, v58, v60, v61, v62, v63, v65 float32
	var awakeSet, dynamicTree, event, fastBody, fastBodySim, fastShape, kinematicTree, shape, shape1, staticTree, v1, v3, v5, v7 uintptr
	var c1, origin, v13, v14, v16, v17, v18, v21, v22, v24 Vec2
	var isBullet, s, v97 uint8
	var q, q2, qn, v12, v20, v68, v69 Rot
	var s1, sweep, v10 Sweep
	var shapeId, v2, v6 int32
	var transform, xf1, xf2, v36 Transform
	var _ /* context at bp+0 */ b2ContinuousContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, aabbMargin, awakeSet, box, box1, box2, c, c1, dynamicTree, event, fastBody, fastBodySim, fastShape, fatAABB, fatAABB1, invMag, isBullet, kinematicTree, mag, omt, origin, q, q2, qn, s, s1, shape, shape1, shapeId, speculativeDistance, staticTree, sweep, transform, x, xf1, xf2, y, v1, v10, v12, v13, v14, v16, v17, v18, v2, v20, v21, v22, v24, v3, v36, v44, v45, v46, v47, v48, v5, v50, v51, v52, v53, v55, v56, v57, v58, v6, v60, v61, v62, v63, v65, v66, v68, v69, v7, v97
	v1 = world + 1064
	v2 = int32(b2_awakeSet)
	if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*88
	goto _4
_4:
	awakeSet = v3
	v5 = awakeSet
	v6 = bodySimIndex
	if !(0 <= v6 && v6 < (*b2BodySimArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2BodySimArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*100
	goto _8
_8:
	fastBodySim = v7
	if !((*b2BodySim)(unsafe.Pointer(fastBodySim)).IsFast != 0) && b2InternalAssertFcn(tls, __ccgo_ts+12621, __ccgo_ts+12460, int32FromInt32(403)) != 0 {
		__builtin_trap(tls)
	}
	v1 = fastBodySim
	s1.C1 = (*b2BodySim)(unsafe.Pointer(v1)).Center0
	s1.C2 = (*b2BodySim)(unsafe.Pointer(v1)).Center
	s1.Q1 = (*b2BodySim)(unsafe.Pointer(v1)).Rotation0
	s1.Q2 = (*b2BodySim)(unsafe.Pointer(v1)).Transform.Q
	s1.LocalCenter = (*b2BodySim)(unsafe.Pointer(v1)).LocalCenter
	v10 = s1
	goto _11
_11:
	sweep = v10
	xf1.Q = sweep.Q1
	v12 = sweep.Q1
	v13 = sweep.LocalCenter
	v14 = Vec2{
		X: float32(v12.C*v13.X) - float32(v12.S*v13.Y),
		Y: float32(v12.S*v13.X) + float32(v12.C*v13.Y),
	}
	goto _15
_15:
	v16 = sweep.C1
	v17 = v14
	v18 = Vec2{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
	}
	goto _19
_19:
	xf1.P = v18
	xf2.Q = sweep.Q2
	v12 = sweep.Q2
	v13 = sweep.LocalCenter
	v14 = Vec2{
		X: float32(v12.C*v13.X) - float32(v12.S*v13.Y),
		Y: float32(v12.S*v13.X) + float32(v12.C*v13.Y),
	}
	goto _23
_23:
	v16 = sweep.C2
	v17 = v14
	v18 = Vec2{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
	}
	goto _27
_27:
	xf2.P = v18
	staticTree = world + 40 + uintptr(b2_staticBody)*72
	kinematicTree = world + 40 + uintptr(b2_kinematicBody)*72
	dynamicTree = world + 40 + uintptr(b2_dynamicBody)*72
	v1 = world + 1024
	v2 = (*b2BodySim)(unsafe.Pointer(fastBodySim)).BodyId
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _31
_31:
	fastBody = v3
	(*(*b2ContinuousContext)(unsafe.Add(nil, bp))).World = world
	(*(*b2ContinuousContext)(unsafe.Add(nil, bp))).Sweep = sweep
	(*(*b2ContinuousContext)(unsafe.Add(nil, bp))).FastBodySim = fastBodySim
	(*(*b2ContinuousContext)(unsafe.Add(nil, bp))).Fraction = float32FromFloat32(1)
	isBullet = (*b2BodySim)(unsafe.Pointer(fastBodySim)).IsBullet
	shapeId = (*b2Body)(unsafe.Pointer(fastBody)).HeadShapeId
	for shapeId != -int32(1) {
		v1 = world + 1248
		v2 = shapeId
		if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
		goto _35
	_35:
		fastShape = v3
		shapeId = (*b2Shape)(unsafe.Pointer(fastShape)).NextShapeId
		(*(*b2ContinuousContext)(unsafe.Add(nil, bp))).FastShape = fastShape
		v36 = xf1
		v13 = (*b2Shape)(unsafe.Pointer(fastShape)).LocalCentroid
		x = float32(v36.Q.C*v13.X) - float32(v36.Q.S*v13.Y) + v36.P.X
		y = float32(v36.Q.S*v13.X) + float32(v36.Q.C*v13.Y) + v36.P.Y
		v14 = Vec2{
			X: x,
			Y: y,
		}
		goto _39
	_39:
		(*(*b2ContinuousContext)(unsafe.Add(nil, bp))).Centroid1 = v14
		v36 = xf2
		v13 = (*b2Shape)(unsafe.Pointer(fastShape)).LocalCentroid
		x = float32(v36.Q.C*v13.X) - float32(v36.Q.S*v13.Y) + v36.P.X
		y = float32(v36.Q.S*v13.X) + float32(v36.Q.C*v13.Y) + v36.P.Y
		v14 = Vec2{
			X: x,
			Y: y,
		}
		goto _43
	_43:
		(*(*b2ContinuousContext)(unsafe.Add(nil, bp))).Centroid2 = v14
		box1 = (*b2Shape)(unsafe.Pointer(fastShape)).Aabb
		box2 = b2ComputeShapeAABB(tls, fastShape, xf2)
		v44 = box1
		v45 = box2
		v46 = v44.LowerBound.X
		v47 = v45.LowerBound.X
		if v46 < v47 {
			v50 = v46
		} else {
			v50 = v47
		}
		v48 = v50
		goto _49
	_49:
		c.LowerBound.X = v48
		v51 = v44.LowerBound.Y
		v52 = v45.LowerBound.Y
		if v51 < v52 {
			v55 = v51
		} else {
			v55 = v52
		}
		v53 = v55
		goto _54
	_54:
		c.LowerBound.Y = v53
		v56 = v44.UpperBound.X
		v57 = v45.UpperBound.X
		if v56 > v57 {
			v60 = v56
		} else {
			v60 = v57
		}
		v58 = v60
		goto _59
	_59:
		c.UpperBound.X = v58
		v61 = v44.UpperBound.Y
		v62 = v45.UpperBound.Y
		if v61 > v62 {
			v65 = v61
		} else {
			v65 = v62
		}
		v63 = v65
		goto _64
	_64:
		c.UpperBound.Y = v63
		v66 = c
		goto _67
	_67:
		box = v66
		// Store this to avoid double computation in the case there is no impact event
		(*b2Shape)(unsafe.Pointer(fastShape)).Aabb = box2
		// No continuous collision for sensors (but still need the updated bounds)
		if (*b2Shape)(unsafe.Pointer(fastShape)).SensorIndex != -int32(1) {
			continue
		}
		b2DynamicTree_Query(tls, staticTree, box, uint64FromUint64(0xffffffffffffffff), __ccgo_fp(b2ContinuousQueryCallback), bp)
		if isBullet != 0 {
			b2DynamicTree_Query(tls, kinematicTree, box, uint64FromUint64(0xffffffffffffffff), __ccgo_fp(b2ContinuousQueryCallback), bp)
			b2DynamicTree_Query(tls, dynamicTree, box, uint64FromUint64(0xffffffffffffffff), __ccgo_fp(b2ContinuousQueryCallback), bp)
		}
	}
	speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter))
	aabbMargin = float32(float32FromFloat32(0.05) * b2_lengthUnitsPerMeter)
	if (*(*b2ContinuousContext)(unsafe.Add(nil, bp))).Fraction < float32FromFloat32(1) {
		v12 = sweep.Q1
		v20 = sweep.Q2
		v46 = (*(*b2ContinuousContext)(unsafe.Add(nil, bp))).Fraction
		omt = float32FromFloat32(1) - v46
		q = Rot{
			C: float32(omt*v12.C) + float32(v46*v20.C),
			S: float32(omt*v12.S) + float32(v46*v20.S),
		}
		mag = sqrtf(tls, float32(q.S*q.S)+float32(q.C*q.C))
		if float64(mag) > float64(0) {
			v47 = float32FromFloat32(1) / mag
		} else {
			v47 = float32FromFloat32(0)
		}
		invMag = v47
		qn = Rot{
			C: float32(q.C * invMag),
			S: float32(q.S * invMag),
		}
		v68 = qn
		goto _73
	_73:
		// Handle time of impact event
		q2 = v68
		v13 = sweep.C1
		v14 = sweep.C2
		v48 = (*(*b2ContinuousContext)(unsafe.Add(nil, bp))).Fraction
		v16 = Vec2{
			X: float32((float32FromFloat32(1)-v48)*v13.X) + float32(v48*v14.X),
			Y: float32((float32FromFloat32(1)-v48)*v13.Y) + float32(v48*v14.Y),
		}
		goto _78
	_78:
		c1 = v16
		v69 = q2
		v17 = sweep.LocalCenter
		v18 = Vec2{
			X: float32(v69.C*v17.X) - float32(v69.S*v17.Y),
			Y: float32(v69.S*v17.X) + float32(v69.C*v17.Y),
		}
		goto _82
	_82:
		v21 = c1
		v22 = v18
		v24 = Vec2{
			X: v21.X - v22.X,
			Y: v21.Y - v22.Y,
		}
		goto _86
	_86:
		origin = v24
		// Advance body
		transform = Transform{
			P: origin,
			Q: q2,
		}
		(*b2BodySim)(unsafe.Pointer(fastBodySim)).Transform = transform
		(*b2BodySim)(unsafe.Pointer(fastBodySim)).Center = c1
		(*b2BodySim)(unsafe.Pointer(fastBodySim)).Rotation0 = q2
		(*b2BodySim)(unsafe.Pointer(fastBodySim)).Center0 = c1
		v1 = world + 1328
		v2 = bodySimIndex
		if !(0 <= v2 && v2 < (*b2BodyMoveEventArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+12641, int32FromInt32(185)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2BodyMoveEventArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*40
		goto _90
	_90:
		// Update body move event
		event = v3
		(*BodyMoveEvent)(unsafe.Pointer(event)).Transform = transform
		// Prepare AABBs for broad-phase.
		// Even though a body is fast, it may not move much. So the
		// AABB may not need enlargement.
		shapeId = (*b2Body)(unsafe.Pointer(fastBody)).HeadShapeId
		for shapeId != -int32(1) {
			v1 = world + 1248
			v2 = shapeId
			if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
				__builtin_trap(tls)
			}
			v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
			goto _94
		_94:
			shape = v3
			// Must recompute aabb at the interpolated transform
			aabb = b2ComputeShapeAABB(tls, shape, transform)
			aabb.LowerBound.X -= speculativeDistance
			aabb.LowerBound.Y -= speculativeDistance
			aabb.UpperBound.X += speculativeDistance
			aabb.UpperBound.Y += speculativeDistance
			(*b2Shape)(unsafe.Pointer(shape)).Aabb = aabb
			v44 = (*b2Shape)(unsafe.Pointer(shape)).FatAABB
			v45 = aabb
			s = boolUint8(true1 != 0)
			s = boolUint8(s != 0 && v44.LowerBound.X <= v45.LowerBound.X)
			s = boolUint8(s != 0 && v44.LowerBound.Y <= v45.LowerBound.Y)
			s = boolUint8(s != 0 && v45.UpperBound.X <= v44.UpperBound.X)
			s = boolUint8(s != 0 && v45.UpperBound.Y <= v44.UpperBound.Y)
			v97 = s
			goto _98
		_98:
			if int32FromUint8(v97) == false1 {
				fatAABB.LowerBound.X = aabb.LowerBound.X - aabbMargin
				fatAABB.LowerBound.Y = aabb.LowerBound.Y - aabbMargin
				fatAABB.UpperBound.X = aabb.UpperBound.X + aabbMargin
				fatAABB.UpperBound.Y = aabb.UpperBound.Y + aabbMargin
				(*b2Shape)(unsafe.Pointer(shape)).FatAABB = fatAABB
				(*b2Shape)(unsafe.Pointer(shape)).EnlargedAABB = boolUint8(true1 != 0)
				(*b2BodySim)(unsafe.Pointer(fastBodySim)).EnlargeAABB = boolUint8(true1 != 0)
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
			v1 = world + 1248
			v2 = shapeId
			if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
				__builtin_trap(tls)
			}
			v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
			goto _102
		_102:
			shape1 = v3
			// shape->aabb is still valid from above
			v44 = (*b2Shape)(unsafe.Pointer(shape1)).FatAABB
			v45 = (*b2Shape)(unsafe.Pointer(shape1)).Aabb
			s = boolUint8(true1 != 0)
			s = boolUint8(s != 0 && v44.LowerBound.X <= v45.LowerBound.X)
			s = boolUint8(s != 0 && v44.LowerBound.Y <= v45.LowerBound.Y)
			s = boolUint8(s != 0 && v45.UpperBound.X <= v44.UpperBound.X)
			s = boolUint8(s != 0 && v45.UpperBound.Y <= v44.UpperBound.Y)
			v97 = s
			goto _106
		_106:
			if int32FromUint8(v97) == false1 {
				fatAABB1.LowerBound.X = (*b2Shape)(unsafe.Pointer(shape1)).Aabb.LowerBound.X - aabbMargin
				fatAABB1.LowerBound.Y = (*b2Shape)(unsafe.Pointer(shape1)).Aabb.LowerBound.Y - aabbMargin
				fatAABB1.UpperBound.X = (*b2Shape)(unsafe.Pointer(shape1)).Aabb.UpperBound.X + aabbMargin
				fatAABB1.UpperBound.Y = (*b2Shape)(unsafe.Pointer(shape1)).Aabb.UpperBound.Y + aabbMargin
				(*b2Shape)(unsafe.Pointer(shape1)).FatAABB = fatAABB1
				(*b2Shape)(unsafe.Pointer(shape1)).EnlargedAABB = boolUint8(true1 != 0)
				(*b2BodySim)(unsafe.Pointer(fastBodySim)).EnlargeAABB = boolUint8(true1 != 0)
			}
			shapeId = (*b2Shape)(unsafe.Pointer(shape1)).NextShapeId
		}
	}
}

// C documentation
//
//	// This should not use the thread index because thread 0 can be called twice by enkiTS.
func b2SolverTask(tls *_Stack, startIndex int32, endIndex int32, threadIndexIgnore uint32, taskContext uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var activeColorCount, bodySyncIndex, colorIndex, colorIndex1, colorIndex2, colorIndex3, graphSyncIndex, i, iterStageIndex, iterStageIndex1, j, j1, previousSyncIndex, spinCount, stageIndex, stageIndex1, subStepCount, syncIndex, workerIndex int32
	var contactSyncIndex, jointSyncIndex, lastSyncBits, syncBits, syncBits1, v8, v9 uint32_t
	var context, profile, stage, stages, workerContext uintptr
	var useBias uint8
	var _ /* ticks at bp+0 */ uint64_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = activeColorCount, bodySyncIndex, colorIndex, colorIndex1, colorIndex2, colorIndex3, contactSyncIndex, context, graphSyncIndex, i, iterStageIndex, iterStageIndex1, j, j1, jointSyncIndex, lastSyncBits, previousSyncIndex, profile, spinCount, stage, stageIndex, stageIndex1, stages, subStepCount, syncBits, syncBits1, syncIndex, useBias, workerContext, workerIndex, v8, v9
	_ = uint64FromInt64(4)
	workerContext = taskContext
	workerIndex = (*b2WorkerContext)(unsafe.Pointer(workerContext)).WorkerIndex
	context = (*b2WorkerContext)(unsafe.Pointer(workerContext)).Context
	activeColorCount = (*b2StepContext)(unsafe.Pointer(context)).ActiveColorCount
	stages = (*b2StepContext)(unsafe.Pointer(context)).Stages
	profile = (*b2StepContext)(unsafe.Pointer(context)).World + 1596
	if workerIndex == 0 {
		// Main thread synchronizes the workers and does work itself.
		//
		// Stages are re-used by loops so that I don't need more stages for large iteration counts.
		// The sync indices grow monotonically for the body/graph/constraint groupings because they share solver blocks.
		// The stage index and sync indices are combined in to sync bits for atomic synchronization.
		// The workers need to compute the previous sync index for a given stage so that CAS works correctly. This
		// setup makes this easy to do.
		/*
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
		*(*uint64_t)(unsafe.Add(nil, bp)) = b2GetTicks(tls)
		bodySyncIndex = int32(1)
		stageIndex = 0
		// This stage loops over all awake joints
		jointSyncIndex = uint32(1)
		syncBits = jointSyncIndex<<int32FromInt32(16) | uint32FromInt32(stageIndex)
		if !((*(*b2SolverStage)(unsafe.Add(nil, stages+uintptr(stageIndex)*32))).Type1 == int32(b2_stagePrepareJoints)) && b2InternalAssertFcn(tls, __ccgo_ts+12913, __ccgo_ts+12460, int32FromInt32(977)) != 0 {
			__builtin_trap(tls)
		}
		b2ExecuteMainStage(tls, stages+uintptr(stageIndex)*32, context, syncBits)
		stageIndex = stageIndex + int32(1)
		jointSyncIndex = jointSyncIndex + uint32(1)
		// This stage loops over all contact constraints
		contactSyncIndex = uint32(1)
		syncBits = contactSyncIndex<<int32FromInt32(16) | uint32FromInt32(stageIndex)
		if !((*(*b2SolverStage)(unsafe.Add(nil, stages+uintptr(stageIndex)*32))).Type1 == int32(b2_stagePrepareContacts)) && b2InternalAssertFcn(tls, __ccgo_ts+12962, __ccgo_ts+12460, int32FromInt32(985)) != 0 {
			__builtin_trap(tls)
		}
		b2ExecuteMainStage(tls, stages+uintptr(stageIndex)*32, context, syncBits)
		stageIndex = stageIndex + int32(1)
		contactSyncIndex = contactSyncIndex + uint32(1)
		graphSyncIndex = int32(1)
		// Single-threaded overflow work. These constraints don't fit in the graph coloring.
		b2PrepareOverflowJoints(tls, context)
		b2PrepareOverflowContacts(tls, context)
		*(*float32)(unsafe.Add(nil, profile+28)) += b2GetMillisecondsAndReset(tls, bp)
		subStepCount = (*b2StepContext)(unsafe.Pointer(context)).SubStepCount
		i = 0
		for {
			if !(i < subStepCount) {
				break
			}
			// stage index restarted each iteration
			// syncBits still increases monotonically because the upper bits increase each iteration
			iterStageIndex = stageIndex
			// integrate velocities
			syncBits = uint32FromInt32(bodySyncIndex<<int32(16) | iterStageIndex)
			if !((*(*b2SolverStage)(unsafe.Add(nil, stages+uintptr(iterStageIndex)*32))).Type1 == int32(b2_stageIntegrateVelocities)) && b2InternalAssertFcn(tls, __ccgo_ts+13013, __ccgo_ts+12460, int32FromInt32(1007)) != 0 {
				__builtin_trap(tls)
			}
			b2ExecuteMainStage(tls, stages+uintptr(iterStageIndex)*32, context, syncBits)
			iterStageIndex = iterStageIndex + int32(1)
			bodySyncIndex = bodySyncIndex + int32(1)
			*(*float32)(unsafe.Add(nil, profile+32)) += b2GetMillisecondsAndReset(tls, bp)
			// warm start constraints
			b2WarmStartOverflowJoints(tls, context)
			b2WarmStartOverflowContacts(tls, context)
			colorIndex = 0
			for {
				if !(colorIndex < activeColorCount) {
					break
				}
				syncBits = uint32FromInt32(graphSyncIndex<<int32(16) | iterStageIndex)
				if !((*(*b2SolverStage)(unsafe.Add(nil, stages+uintptr(iterStageIndex)*32))).Type1 == int32(b2_stageWarmStart)) && b2InternalAssertFcn(tls, __ccgo_ts+13072, __ccgo_ts+12460, int32FromInt32(1021)) != 0 {
					__builtin_trap(tls)
				}
				b2ExecuteMainStage(tls, stages+uintptr(iterStageIndex)*32, context, syncBits)
				iterStageIndex = iterStageIndex + int32(1)
				goto _2
			_2:
				;
				colorIndex = colorIndex + 1
			}
			graphSyncIndex = graphSyncIndex + int32(1)
			*(*float32)(unsafe.Add(nil, profile+36)) += b2GetMillisecondsAndReset(tls, bp)
			// solve constraints
			useBias = boolUint8(true1 != 0)
			j = 0
			for {
				if !(j < int32(_ITERATIONS)) {
					break
				}
				b2SolveOverflowJoints(tls, context, useBias)
				b2SolveOverflowContacts(tls, context, useBias)
				colorIndex1 = 0
				for {
					if !(colorIndex1 < activeColorCount) {
						break
					}
					syncBits = uint32FromInt32(graphSyncIndex<<int32(16) | iterStageIndex)
					if !((*(*b2SolverStage)(unsafe.Add(nil, stages+uintptr(iterStageIndex)*32))).Type1 == int32(b2_stageSolve)) && b2InternalAssertFcn(tls, __ccgo_ts+13121, __ccgo_ts+12460, int32FromInt32(1040)) != 0 {
						__builtin_trap(tls)
					}
					b2ExecuteMainStage(tls, stages+uintptr(iterStageIndex)*32, context, syncBits)
					iterStageIndex = iterStageIndex + int32(1)
					goto _4
				_4:
					;
					colorIndex1 = colorIndex1 + 1
				}
				graphSyncIndex = graphSyncIndex + int32(1)
				goto _3
			_3:
				;
				j = j + 1
			}
			*(*float32)(unsafe.Add(nil, profile+40)) += b2GetMillisecondsAndReset(tls, bp)
			// integrate positions
			if !((*(*b2SolverStage)(unsafe.Add(nil, stages+uintptr(iterStageIndex)*32))).Type1 == int32(b2_stageIntegratePositions)) && b2InternalAssertFcn(tls, __ccgo_ts+13166, __ccgo_ts+12460, int32FromInt32(1050)) != 0 {
				__builtin_trap(tls)
			}
			syncBits = uint32FromInt32(bodySyncIndex<<int32(16) | iterStageIndex)
			b2ExecuteMainStage(tls, stages+uintptr(iterStageIndex)*32, context, syncBits)
			iterStageIndex = iterStageIndex + int32(1)
			bodySyncIndex = bodySyncIndex + int32(1)
			*(*float32)(unsafe.Add(nil, profile+44)) += b2GetMillisecondsAndReset(tls, bp)
			// relax constraints
			useBias = boolUint8(false1 != 0)
			j1 = 0
			for {
				if !(j1 < int32(_RELAX_ITERATIONS)) {
					break
				}
				b2SolveOverflowJoints(tls, context, useBias)
				b2SolveOverflowContacts(tls, context, useBias)
				colorIndex2 = 0
				for {
					if !(colorIndex2 < activeColorCount) {
						break
					}
					syncBits = uint32FromInt32(graphSyncIndex<<int32(16) | iterStageIndex)
					if !((*(*b2SolverStage)(unsafe.Add(nil, stages+uintptr(iterStageIndex)*32))).Type1 == int32(b2_stageRelax)) && b2InternalAssertFcn(tls, __ccgo_ts+13224, __ccgo_ts+12460, int32FromInt32(1068)) != 0 {
						__builtin_trap(tls)
					}
					b2ExecuteMainStage(tls, stages+uintptr(iterStageIndex)*32, context, syncBits)
					iterStageIndex = iterStageIndex + int32(1)
					goto _6
				_6:
					;
					colorIndex2 = colorIndex2 + 1
				}
				graphSyncIndex = graphSyncIndex + int32(1)
				goto _5
			_5:
				;
				j1 = j1 + 1
			}
			*(*float32)(unsafe.Add(nil, profile+48)) += b2GetMillisecondsAndReset(tls, bp)
			goto _1
		_1:
			;
			i = i + 1
		}
		// advance the stage according to the sub-stepping tasks just completed
		// integrate velocities / warm start / solve / integrate positions / relax
		stageIndex = stageIndex + (int32(1) + activeColorCount + int32(_ITERATIONS)*activeColorCount + int32(1) + int32(_RELAX_ITERATIONS)*activeColorCount)
		// Restitution
		b2ApplyOverflowRestitution(tls, context)
		iterStageIndex1 = stageIndex
		colorIndex3 = 0
		for {
			if !(colorIndex3 < activeColorCount) {
				break
			}
			syncBits = uint32FromInt32(graphSyncIndex<<int32(16) | iterStageIndex1)
			if !((*(*b2SolverStage)(unsafe.Add(nil, stages+uintptr(iterStageIndex1)*32))).Type1 == int32(b2_stageRestitution)) && b2InternalAssertFcn(tls, __ccgo_ts+13269, __ccgo_ts+12460, int32FromInt32(1090)) != 0 {
				__builtin_trap(tls)
			}
			b2ExecuteMainStage(tls, stages+uintptr(iterStageIndex1)*32, context, syncBits)
			iterStageIndex1 = iterStageIndex1 + int32(1)
			goto _7
		_7:
			;
			colorIndex3 = colorIndex3 + 1
		}
		// graphSyncIndex += 1;
		stageIndex = stageIndex + activeColorCount
		*(*float32)(unsafe.Add(nil, profile+52)) += b2GetMillisecondsAndReset(tls, bp)
		b2StoreOverflowImpulses(tls, context)
		syncBits = contactSyncIndex<<int32FromInt32(16) | uint32FromInt32(stageIndex)
		if !((*(*b2SolverStage)(unsafe.Add(nil, stages+uintptr(stageIndex)*32))).Type1 == int32(b2_stageStoreImpulses)) && b2InternalAssertFcn(tls, __ccgo_ts+13320, __ccgo_ts+12460, int32FromInt32(1103)) != 0 {
			__builtin_trap(tls)
		}
		b2ExecuteMainStage(tls, stages+uintptr(stageIndex)*32, context, syncBits)
		*(*float32)(unsafe.Add(nil, profile+56)) += b2GetMillisecondsAndReset(tls, bp)
		// Signal workers to finish
		atomicStoreNUint32(context+232, uint32(0xffffffff), int32FromInt32(__ATOMIC_SEQ_CST))
		if !(stageIndex+int32FromInt32(1) == (*b2StepContext)(unsafe.Pointer(context)).StageCount) && b2InternalAssertFcn(tls, __ccgo_ts+13369, __ccgo_ts+12460, int32FromInt32(1111)) != 0 {
			__builtin_trap(tls)
		}
		return
	}
	// Worker spins and waits for work
	lastSyncBits = uint32(0)
	// uint64_t maxSpinTime = 10;
	for int32(true1) != 0 {
		spinCount = 0
		for {
			v9 = atomicLoadNUint32(context+232, int32FromInt32(__ATOMIC_SEQ_CST))
			goto _10
		_10:
			v8 = v9
			syncBits1 = v8
			if !(v8 == lastSyncBits) {
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
				//	b2Pause();
				//}
				// while ((__rdtsc() - prev) < maxSpinTime);
				// maxSpinTime += 10;
				b2Pause(tls)
				b2Pause(tls)
				spinCount = spinCount + int32(1)
			}
		}
		if syncBits1 == uint32(0xffffffff) {
			// sentinel hit
			break
		}
		stageIndex1 = int32FromUint32(syncBits1 & uint32(0xFFFF))
		if !(stageIndex1 < (*b2StepContext)(unsafe.Pointer(context)).StageCount) && b2InternalAssertFcn(tls, __ccgo_ts+13407, __ccgo_ts+12460, int32FromInt32(1155)) != 0 {
			__builtin_trap(tls)
		}
		syncIndex = int32FromUint32(syncBits1 >> int32FromInt32(16) & uint32(0xFFFF))
		if !(syncIndex > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+12899, __ccgo_ts+12460, int32FromInt32(1158)) != 0 {
			__builtin_trap(tls)
		}
		previousSyncIndex = syncIndex - int32(1)
		stage = stages + uintptr(stageIndex1)*32
		b2ExecuteStage(tls, stage, context, previousSyncIndex, syncIndex, workerIndex)
		lastSyncBits = syncBits1
	}
}

// C documentation
//
//	// Solve with graph coloring
func b2Solve(tls *_Stack, world uintptr, stepContext uintptr) {
	bp := tls.Alloc(1552)
	defer tls.Free(1552)
	var activeColorCount, awakeBodyCount, awakeIslandCount, awakeJointCount, blocksPerWorker, bodyBlockCount, bodyBlockSize, bodyId, bulletBodyCount, c, colorContactBlockCount, colorContactBlockSize, colorContactCount, colorContactCount1, colorContactCountSIMD, colorContactCountSIMD1, colorJointBlockCount, colorJointBlockSize, colorJointCount, colorJointCount1, contactBase, contactBlockCount, contactBlockSize, contactCount, count1, graphBlockCount, i, i1, i10, i11, i12, i13, i14, i15, i16, i17, i18, i2, i3, i4, i5, i6, i7, i8, i9, islandId, islandIndex, j, j1, j2, j3, j4, j5, jointBase, jointBlockCount, jointBlockSize, k, k1, k2, k3, maxBlockCount, minRange, newCapacity, newCapacity1, occupancyCount, overflowContactCount, perColorContactCount, perColorJointCount, pointCount, proxyId, proxyKey, shapeId, shapeId1, simdConstraintSize, simdContactCount, stageCount, workerCount, v2, v6 int32
	var activeColorIndices, colorContactBlockCounts, colorContactBlockSizes, colorContactCounts, colorJointBlockCounts, colorJointBlockSizes, colorJointCounts [12]int32
	var alreadyAdded, hit, v63 uint8
	var approachSpeed, splitSleepTimer, threshold float32
	var awakeIslandBitSet, awakeSet, baseGraphBlock, bits, block1, block2, block3, block4, block5, body, bodyArray, bodyArray1, bodyBlocks, bodySim, bodySimArray, bodySimArray1, broadPhase, broadPhase1, bulletBody, bulletBodySim, bulletBodySimIndices, color, color1, colors, colors1, contactBlocks, contactSim, contactSims, contacts, dynamicTree, enlargedBodyBitSet, finalizeBodiesTask, graph, graphBlocks, island, islands, jointBlocks, joints, mp, overflowContactConstraints, shape, shape1, shape2, shapeA, shapeArray, shapeArray1, shapeB, simdContactConstraints, splitIslandTask, stage, stages, taskContext, taskContext1, userBulletBodyTask, v1, v3, v37, v7 uintptr
	var blockIndex, bodySimIndex, ctz, k4, wordCount, v49 uint32_t
	var bulletTicks, hitTicks, mergeTicks, refitTicks, sleepTicks, transformTicks, word uint64_t
	var event ContactHitEvent
	var graphColorBlocks [12]uintptr
	var _ /* constraintTicks at bp+1544 */ uint64_t
	var _ /* prepareTicks at bp+0 */ uint64_t
	var _ /* workerContext at bp+8 */ [64]b2WorkerContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = activeColorCount, activeColorIndices, alreadyAdded, approachSpeed, awakeBodyCount, awakeIslandBitSet, awakeIslandCount, awakeJointCount, awakeSet, baseGraphBlock, bits, block1, block2, block3, block4, block5, blockIndex, blocksPerWorker, body, bodyArray, bodyArray1, bodyBlockCount, bodyBlockSize, bodyBlocks, bodyId, bodySim, bodySimArray, bodySimArray1, bodySimIndex, broadPhase, broadPhase1, bulletBody, bulletBodyCount, bulletBodySim, bulletBodySimIndices, bulletTicks, c, color, color1, colorContactBlockCount, colorContactBlockCounts, colorContactBlockSize, colorContactBlockSizes, colorContactCount, colorContactCount1, colorContactCountSIMD, colorContactCountSIMD1, colorContactCounts, colorJointBlockCount, colorJointBlockCounts, colorJointBlockSize, colorJointBlockSizes, colorJointCount, colorJointCount1, colorJointCounts, colors, colors1, contactBase, contactBlockCount, contactBlockSize, contactBlocks, contactCount, contactSim, contactSims, contacts, count1, ctz, dynamicTree, enlargedBodyBitSet, event, finalizeBodiesTask, graph, graphBlockCount, graphBlocks, graphColorBlocks, hit, hitTicks, i, i1, i10, i11, i12, i13, i14, i15, i16, i17, i18, i2, i3, i4, i5, i6, i7, i8, i9, island, islandId, islandIndex, islands, j, j1, j2, j3, j4, j5, jointBase, jointBlockCount, jointBlockSize, jointBlocks, joints, k, k1, k2, k3, k4, maxBlockCount, mergeTicks, minRange, mp, newCapacity, newCapacity1, occupancyCount, overflowContactConstraints, overflowContactCount, perColorContactCount, perColorJointCount, pointCount, proxyId, proxyKey, refitTicks, shape, shape1, shape2, shapeA, shapeArray, shapeArray1, shapeB, shapeId, shapeId1, simdConstraintSize, simdContactConstraints, simdContactCount, sleepTicks, splitIslandTask, splitSleepTimer, stage, stageCount, stages, taskContext, taskContext1, threshold, transformTicks, userBulletBodyTask, word, wordCount, workerCount, v1, v2, v3, v37, v49, v6, v63, v7
	*(*uint64_t)(unsafe.Add(nil, world+1528)) += uint64(1)
	// Merge islands
	mergeTicks = b2GetTicks(tls)
	b2MergeAwakeIslands(tls, world)
	(*b2World)(unsafe.Pointer(world)).Profile.MergeIslands = b2GetMilliseconds(tls, mergeTicks)
	v1 = world + 1064
	v2 = int32(b2_awakeSet)
	if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*88
	goto _4
_4:
	// Are there any awake bodies? This scenario should not be important for profiling.
	awakeSet = v3
	awakeBodyCount = (*b2SolverSet)(unsafe.Pointer(awakeSet)).BodySims.Count
	if awakeBodyCount == 0 {
		// Nothing to simulate, however the tree rebuild must be finished.
		if (*b2World)(unsafe.Pointer(world)).UserTreeTask != uintptrFromInt32(0) {
			(*(*func(*_Stack, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).FinishTaskFcn})))(tls, (*b2World)(unsafe.Pointer(world)).UserTreeTask, (*b2World)(unsafe.Pointer(world)).UserTaskContext)
			(*b2World)(unsafe.Pointer(world)).UserTreeTask = uintptrFromInt32(0)
			*(*int32)(unsafe.Add(nil, world+1772)) -= int32(1)
		}
		b2ValidateNoEnlarged(tls, world+40)
		return
	}
	// Solve constraints using graph coloring
	// Prepare buffers for bullets
	atomicStoreNInt32(stepContext+112, 0, int32FromInt32(__ATOMIC_SEQ_CST))
	(*b2StepContext)(unsafe.Pointer(stepContext)).BulletBodies = b2AllocateArenaItem(tls, world, int32FromUint64(uint64FromInt32(awakeBodyCount)*uint64(4)), __ccgo_ts+13440)
	*(*uint64_t)(unsafe.Add(nil, bp)) = b2GetTicks(tls)
	graph = world + 328
	colors = graph
	(*b2StepContext)(unsafe.Pointer(stepContext)).Sims = (*b2SolverSet)(unsafe.Pointer(awakeSet)).BodySims.Data
	(*b2StepContext)(unsafe.Pointer(stepContext)).States = (*b2SolverSet)(unsafe.Pointer(awakeSet)).BodyStates.Data
	// count contacts, joints, and colors
	awakeJointCount = 0
	activeColorCount = 0
	i = 0
	for {
		if !(i < int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1)) {
			break
		}
		perColorContactCount = (*(*b2GraphColor)(unsafe.Add(nil, colors+uintptr(i)*56))).ContactSims.Count
		perColorJointCount = (*(*b2GraphColor)(unsafe.Add(nil, colors+uintptr(i)*56))).JointSims.Count
		occupancyCount = perColorContactCount + perColorJointCount
		if occupancyCount > 0 {
			v2 = int32(1)
		} else {
			v2 = 0
		}
		activeColorCount = activeColorCount + v2
		awakeJointCount = awakeJointCount + perColorJointCount
		goto _5
	_5:
		;
		i = i + 1
	}
	// prepare for move events
	v1 = world + 1328
	v2 = awakeBodyCount
	b2BodyMoveEventArray_Reserve(tls, v1, v2)
	(*b2BodyMoveEventArray)(unsafe.Pointer(v1)).Count = v2
	// Each worker receives at most M blocks of work. The workers may receive less blocks if there is not sufficient work.
	// Each block of work has a minimum number of elements (block size). This in turn may limit the number of blocks.
	// If there are many elements then the block size is increased so there are still at most M blocks of work per worker.
	// M is a tunable number that has two goals:
	// 1. keep M small to reduce overhead
	// 2. keep M large enough for other workers to be able to steal work
	// The block size is a power of two to make math efficient.
	workerCount = (*b2World)(unsafe.Pointer(world)).WorkerCount
	blocksPerWorker = int32(4)
	maxBlockCount = blocksPerWorker * workerCount
	// Configure blocks for tasks that parallel-for bodies
	bodyBlockSize = int32FromInt32(1) << int32FromInt32(5)
	if awakeBodyCount > bodyBlockSize*maxBlockCount {
		// Too many blocks, increase block size
		bodyBlockSize = awakeBodyCount / maxBlockCount
		bodyBlockCount = maxBlockCount
	} else {
		bodyBlockCount = (awakeBodyCount-int32(1))>>int32(5) + int32(1)
	}
	graphBlockCount = 0
	// c is the active color index
	simdContactCount = 0
	c = 0
	i1 = 0
	for {
		if !(i1 < int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1)) {
			break
		}
		colorContactCount = (*(*b2GraphColor)(unsafe.Add(nil, colors+uintptr(i1)*56))).ContactSims.Count
		colorJointCount = (*(*b2GraphColor)(unsafe.Add(nil, colors+uintptr(i1)*56))).JointSims.Count
		if colorContactCount+colorJointCount > 0 {
			activeColorIndices[c] = i1
			if colorContactCount > 0 {
				v2 = (colorContactCount-int32(1))>>int32(_B2_SIMD_SHIFT) + int32(1)
			} else {
				v2 = 0
			}
			// 4/8-way SIMD
			colorContactCountSIMD = v2
			colorContactCounts[c] = colorContactCountSIMD
			// determine the number of contact work blocks for this color
			if colorContactCountSIMD > blocksPerWorker*maxBlockCount {
				// too many contact blocks
				colorContactBlockSizes[c] = colorContactCountSIMD / maxBlockCount
				colorContactBlockCounts[c] = maxBlockCount
			} else {
				if colorContactCountSIMD > 0 {
					// dividing by blocksPerWorker (4)
					colorContactBlockSizes[c] = blocksPerWorker
					colorContactBlockCounts[c] = (colorContactCountSIMD-int32(1))>>int32(2) + int32(1)
				} else {
					// no contacts in this color
					colorContactBlockSizes[c] = 0
					colorContactBlockCounts[c] = 0
				}
			}
			colorJointCounts[c] = colorJointCount
			// determine number of joint work blocks for this color
			if colorJointCount > blocksPerWorker*maxBlockCount {
				// too many joint blocks
				colorJointBlockSizes[c] = colorJointCount / maxBlockCount
				colorJointBlockCounts[c] = maxBlockCount
			} else {
				if colorJointCount > 0 {
					// dividing by blocksPerWorker (4)
					colorJointBlockSizes[c] = blocksPerWorker
					colorJointBlockCounts[c] = (colorJointCount-int32(1))>>int32(2) + int32(1)
				} else {
					colorJointBlockSizes[c] = 0
					colorJointBlockCounts[c] = 0
				}
			}
			graphBlockCount = graphBlockCount + (colorContactBlockCounts[c] + colorJointBlockCounts[c])
			simdContactCount = simdContactCount + colorContactCountSIMD
			c = c + int32(1)
		}
		goto _9
	_9:
		;
		i1 = i1 + 1
	}
	activeColorCount = c
	// Gather contact pointers for easy parallel-for traversal. Some may be NULL due to SIMD remainders.
	contacts = b2AllocateArenaItem(tls, world, int32FromUint64(uint64FromInt32(int32(_B2_SIMD_WIDTH)*simdContactCount)*uint64(8)), __ccgo_ts+13454)
	// Gather joint pointers for easy parallel-for traversal.
	joints = b2AllocateArenaItem(tls, world, int32FromUint64(uint64FromInt32(awakeJointCount)*uint64(8)), __ccgo_ts+13471)
	simdConstraintSize = b2GetContactConstraintSIMDByteCount(tls)
	simdContactConstraints = b2AllocateArenaItem(tls, world, simdContactCount*simdConstraintSize, __ccgo_ts+13486)
	overflowContactCount = (*(*b2GraphColor)(unsafe.Add(nil, colors+uintptr(int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1))*56))).ContactSims.Count
	overflowContactConstraints = b2AllocateArenaItem(tls, world, int32FromUint64(uint64FromInt32(overflowContactCount)*uint64(160)), __ccgo_ts+13505)
	*(*uintptr)(unsafe.Add(unsafe.Pointer(&*(*b2GraphColor)(unsafe.Add(nil, graph+uintptr(int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1))*56))), 48)) = overflowContactConstraints
	// Distribute transient constraints to each graph color and build flat arrays of contact and joint pointers
	contactBase = 0
	jointBase = 0
	i2 = 0
	for {
		if !(i2 < activeColorCount) {
			break
		}
		j = activeColorIndices[i2]
		color = colors + uintptr(j)*56
		colorContactCount1 = (*b2GraphColor)(unsafe.Pointer(color)).ContactSims.Count
		if colorContactCount1 == 0 {
			(*b2GraphColor)(unsafe.Pointer(color)).ｆ__ccgo3_48.SimdConstraints = uintptrFromInt32(0)
		} else {
			(*b2GraphColor)(unsafe.Pointer(color)).ｆ__ccgo3_48.SimdConstraints = simdContactConstraints + uintptr(contactBase*simdConstraintSize)
			k = 0
			for {
				if !(k < colorContactCount1) {
					break
				}
				*(*uintptr)(unsafe.Add(nil, contacts+uintptr(int32(_B2_SIMD_WIDTH)*contactBase+k)*8)) = (*b2GraphColor)(unsafe.Pointer(color)).ContactSims.Data + uintptr(k)*176
				goto _12
			_12:
				;
				k = k + 1
			}
			// remainder
			colorContactCountSIMD1 = (colorContactCount1-int32(1))>>int32(_B2_SIMD_SHIFT) + int32(1)
			k1 = colorContactCount1
			for {
				if !(k1 < int32(_B2_SIMD_WIDTH)*colorContactCountSIMD1) {
					break
				}
				*(*uintptr)(unsafe.Add(nil, contacts+uintptr(int32(_B2_SIMD_WIDTH)*contactBase+k1)*8)) = uintptrFromInt32(0)
				goto _13
			_13:
				;
				k1 = k1 + 1
			}
			contactBase = contactBase + colorContactCountSIMD1
		}
		colorJointCount1 = (*b2GraphColor)(unsafe.Pointer(color)).JointSims.Count
		k2 = 0
		for {
			if !(k2 < colorJointCount1) {
				break
			}
			*(*uintptr)(unsafe.Add(nil, joints+uintptr(jointBase+k2)*8)) = (*b2GraphColor)(unsafe.Pointer(color)).JointSims.Data + uintptr(k2)*196
			goto _14
		_14:
			;
			k2 = k2 + 1
		}
		jointBase = jointBase + colorJointCount1
		goto _11
	_11:
		;
		i2 = i2 + 1
	}
	if !(contactBase == simdContactCount) && b2InternalAssertFcn(tls, __ccgo_ts+13533, __ccgo_ts+12460, int32FromInt32(1425)) != 0 {
		__builtin_trap(tls)
	}
	if !(jointBase == awakeJointCount) && b2InternalAssertFcn(tls, __ccgo_ts+13565, __ccgo_ts+12460, int32FromInt32(1426)) != 0 {
		__builtin_trap(tls)
	}
	// Define work blocks for preparing contacts and storing contact impulses
	contactBlockSize = blocksPerWorker
	if simdContactCount > 0 {
		v2 = (simdContactCount-int32(1))>>int32(2) + int32(1)
	} else {
		v2 = 0
	}
	contactBlockCount = v2
	if simdContactCount > contactBlockSize*maxBlockCount {
		// Too many blocks, increase block size
		contactBlockSize = simdContactCount / maxBlockCount
		contactBlockCount = maxBlockCount
	}
	// Define work blocks for preparing joints
	jointBlockSize = blocksPerWorker
	if awakeJointCount > 0 {
		v2 = (awakeJointCount-int32(1))>>int32(2) + int32(1)
	} else {
		v2 = 0
	}
	jointBlockCount = v2
	if awakeJointCount > jointBlockSize*maxBlockCount {
		// Too many blocks, increase block size
		jointBlockSize = awakeJointCount / maxBlockCount
		jointBlockCount = maxBlockCount
	}
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
	stages = b2AllocateArenaItem(tls, world, int32FromUint64(uint64FromInt32(stageCount)*uint64(32)), __ccgo_ts+13594)
	bodyBlocks = b2AllocateArenaItem(tls, world, int32FromUint64(uint64FromInt32(bodyBlockCount)*uint64(12)), __ccgo_ts+13601)
	contactBlocks = b2AllocateArenaItem(tls, world, int32FromUint64(uint64FromInt32(contactBlockCount)*uint64(12)), __ccgo_ts+13613)
	jointBlocks = b2AllocateArenaItem(tls, world, int32FromUint64(uint64FromInt32(jointBlockCount)*uint64(12)), __ccgo_ts+13628)
	graphBlocks = b2AllocateArenaItem(tls, world, int32FromUint64(uint64FromInt32(graphBlockCount)*uint64(12)), __ccgo_ts+13641)
	// Split an awake island. This modifies:
	// - stack allocator
	// - world island array and solver set
	// - island indices on bodies, contacts, and joints
	// I'm squeezing this task in here because it may be expensive and this is a safe place to put it.
	// Note: cannot split islands in parallel with FinalizeBodies
	splitIslandTask = uintptrFromInt32(0)
	if (*b2World)(unsafe.Pointer(world)).SplitIslandId != -int32(1) {
		splitIslandTask = (*(*func(*_Stack, uintptr, int32, int32, uintptr, uintptr) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).EnqueueTaskFcn})))(tls, __ccgo_fp(b2SplitIslandTask), int32(1), int32(1), world, (*b2World)(unsafe.Pointer(world)).UserTaskContext)
		*(*int32)(unsafe.Add(nil, world+1776)) += int32(1)
		if splitIslandTask == uintptrFromInt32(0) {
			v2 = 0
		} else {
			v2 = int32(1)
		}
		*(*int32)(unsafe.Add(nil, world+1772)) += v2
	}
	// Prepare body work blocks
	i3 = 0
	for {
		if !(i3 < bodyBlockCount) {
			break
		}
		block1 = bodyBlocks + uintptr(i3)*12
		(*b2SolverBlock)(unsafe.Pointer(block1)).StartIndex = i3 * bodyBlockSize
		(*b2SolverBlock)(unsafe.Pointer(block1)).Count = int16(bodyBlockSize)
		(*b2SolverBlock)(unsafe.Pointer(block1)).BlockType = int16(b2_bodyBlock)
		atomicStoreNInt32(block1+8, 0, int32FromInt32(__ATOMIC_SEQ_CST))
		goto _18
	_18:
		;
		i3 = i3 + 1
	}
	(*(*b2SolverBlock)(unsafe.Add(nil, bodyBlocks+uintptr(bodyBlockCount-int32(1))*12))).Count = int16(awakeBodyCount - (bodyBlockCount-int32FromInt32(1))*bodyBlockSize)
	// Prepare joint work blocks
	i4 = 0
	for {
		if !(i4 < jointBlockCount) {
			break
		}
		block2 = jointBlocks + uintptr(i4)*12
		(*b2SolverBlock)(unsafe.Pointer(block2)).StartIndex = i4 * jointBlockSize
		(*b2SolverBlock)(unsafe.Pointer(block2)).Count = int16(jointBlockSize)
		(*b2SolverBlock)(unsafe.Pointer(block2)).BlockType = int16(b2_jointBlock)
		atomicStoreNInt32(block2+8, 0, int32FromInt32(__ATOMIC_SEQ_CST))
		goto _19
	_19:
		;
		i4 = i4 + 1
	}
	if jointBlockCount > 0 {
		(*(*b2SolverBlock)(unsafe.Add(nil, jointBlocks+uintptr(jointBlockCount-int32(1))*12))).Count = int16(awakeJointCount - (jointBlockCount-int32FromInt32(1))*jointBlockSize)
	}
	// Prepare contact work blocks
	i5 = 0
	for {
		if !(i5 < contactBlockCount) {
			break
		}
		block3 = contactBlocks + uintptr(i5)*12
		(*b2SolverBlock)(unsafe.Pointer(block3)).StartIndex = i5 * contactBlockSize
		(*b2SolverBlock)(unsafe.Pointer(block3)).Count = int16(contactBlockSize)
		(*b2SolverBlock)(unsafe.Pointer(block3)).BlockType = int16(b2_contactBlock)
		atomicStoreNInt32(block3+8, 0, int32FromInt32(__ATOMIC_SEQ_CST))
		goto _20
	_20:
		;
		i5 = i5 + 1
	}
	if contactBlockCount > 0 {
		(*(*b2SolverBlock)(unsafe.Add(nil, contactBlocks+uintptr(contactBlockCount-int32(1))*12))).Count = int16(simdContactCount - (contactBlockCount-int32FromInt32(1))*contactBlockSize)
	}
	baseGraphBlock = graphBlocks
	i6 = 0
	for {
		if !(i6 < activeColorCount) {
			break
		}
		graphColorBlocks[i6] = baseGraphBlock
		colorJointBlockCount = colorJointBlockCounts[i6]
		colorJointBlockSize = colorJointBlockSizes[i6]
		j1 = 0
		for {
			if !(j1 < colorJointBlockCount) {
				break
			}
			block4 = baseGraphBlock + uintptr(j1)*12
			(*b2SolverBlock)(unsafe.Pointer(block4)).StartIndex = j1 * colorJointBlockSize
			(*b2SolverBlock)(unsafe.Pointer(block4)).Count = int16(colorJointBlockSize)
			(*b2SolverBlock)(unsafe.Pointer(block4)).BlockType = int16(b2_graphJointBlock)
			atomicStoreNInt32(block4+8, 0, int32FromInt32(__ATOMIC_SEQ_CST))
			goto _22
		_22:
			;
			j1 = j1 + 1
		}
		if colorJointBlockCount > 0 {
			(*(*b2SolverBlock)(unsafe.Add(nil, baseGraphBlock+uintptr(colorJointBlockCount-int32(1))*12))).Count = int16(colorJointCounts[i6] - (colorJointBlockCount-int32FromInt32(1))*colorJointBlockSize)
			baseGraphBlock = baseGraphBlock + uintptr(colorJointBlockCount)*12
		}
		colorContactBlockCount = colorContactBlockCounts[i6]
		colorContactBlockSize = colorContactBlockSizes[i6]
		j2 = 0
		for {
			if !(j2 < colorContactBlockCount) {
				break
			}
			block5 = baseGraphBlock + uintptr(j2)*12
			(*b2SolverBlock)(unsafe.Pointer(block5)).StartIndex = j2 * colorContactBlockSize
			(*b2SolverBlock)(unsafe.Pointer(block5)).Count = int16(colorContactBlockSize)
			(*b2SolverBlock)(unsafe.Pointer(block5)).BlockType = int16(b2_graphContactBlock)
			atomicStoreNInt32(block5+8, 0, int32FromInt32(__ATOMIC_SEQ_CST))
			goto _23
		_23:
			;
			j2 = j2 + 1
		}
		if colorContactBlockCount > 0 {
			(*(*b2SolverBlock)(unsafe.Add(nil, baseGraphBlock+uintptr(colorContactBlockCount-int32(1))*12))).Count = int16(colorContactCounts[i6] - (colorContactBlockCount-int32FromInt32(1))*colorContactBlockSize)
			baseGraphBlock = baseGraphBlock + uintptr(colorContactBlockCount)*12
		}
		goto _21
	_21:
		;
		i6 = i6 + 1
	}
	if !((int64(baseGraphBlock)-int64(graphBlocks))/12 == int64(graphBlockCount)) && b2InternalAssertFcn(tls, __ccgo_ts+13654, __ccgo_ts+12460, int32FromInt32(1581)) != 0 {
		__builtin_trap(tls)
	}
	stage = stages
	// Prepare joints
	(*b2SolverStage)(unsafe.Pointer(stage)).Type1 = int32(b2_stagePrepareJoints)
	(*b2SolverStage)(unsafe.Pointer(stage)).Blocks = jointBlocks
	(*b2SolverStage)(unsafe.Pointer(stage)).BlockCount = jointBlockCount
	(*b2SolverStage)(unsafe.Pointer(stage)).ColorIndex = -int32(1)
	atomicStoreNInt32(stage+24, 0, int32FromInt32(__ATOMIC_SEQ_CST))
	stage = stage + uintptr(1)*32
	// Prepare contacts
	(*b2SolverStage)(unsafe.Pointer(stage)).Type1 = int32(b2_stagePrepareContacts)
	(*b2SolverStage)(unsafe.Pointer(stage)).Blocks = contactBlocks
	(*b2SolverStage)(unsafe.Pointer(stage)).BlockCount = contactBlockCount
	(*b2SolverStage)(unsafe.Pointer(stage)).ColorIndex = -int32(1)
	atomicStoreNInt32(stage+24, 0, int32FromInt32(__ATOMIC_SEQ_CST))
	stage = stage + uintptr(1)*32
	// Integrate velocities
	(*b2SolverStage)(unsafe.Pointer(stage)).Type1 = int32(b2_stageIntegrateVelocities)
	(*b2SolverStage)(unsafe.Pointer(stage)).Blocks = bodyBlocks
	(*b2SolverStage)(unsafe.Pointer(stage)).BlockCount = bodyBlockCount
	(*b2SolverStage)(unsafe.Pointer(stage)).ColorIndex = -int32(1)
	atomicStoreNInt32(stage+24, 0, int32FromInt32(__ATOMIC_SEQ_CST))
	stage = stage + uintptr(1)*32
	// Warm start
	i7 = 0
	for {
		if !(i7 < activeColorCount) {
			break
		}
		(*b2SolverStage)(unsafe.Pointer(stage)).Type1 = int32(b2_stageWarmStart)
		(*b2SolverStage)(unsafe.Pointer(stage)).Blocks = graphColorBlocks[i7]
		(*b2SolverStage)(unsafe.Pointer(stage)).BlockCount = colorJointBlockCounts[i7] + colorContactBlockCounts[i7]
		(*b2SolverStage)(unsafe.Pointer(stage)).ColorIndex = activeColorIndices[i7]
		atomicStoreNInt32(stage+24, 0, int32FromInt32(__ATOMIC_SEQ_CST))
		stage = stage + uintptr(1)*32
		goto _24
	_24:
		;
		i7 = i7 + 1
	}
	// Solve graph
	j3 = 0
	for {
		if !(j3 < int32(_ITERATIONS)) {
			break
		}
		i8 = 0
		for {
			if !(i8 < activeColorCount) {
				break
			}
			(*b2SolverStage)(unsafe.Pointer(stage)).Type1 = int32(b2_stageSolve)
			(*b2SolverStage)(unsafe.Pointer(stage)).Blocks = graphColorBlocks[i8]
			(*b2SolverStage)(unsafe.Pointer(stage)).BlockCount = colorJointBlockCounts[i8] + colorContactBlockCounts[i8]
			(*b2SolverStage)(unsafe.Pointer(stage)).ColorIndex = activeColorIndices[i8]
			atomicStoreNInt32(stage+24, 0, int32FromInt32(__ATOMIC_SEQ_CST))
			stage = stage + uintptr(1)*32
			goto _26
		_26:
			;
			i8 = i8 + 1
		}
		goto _25
	_25:
		;
		j3 = j3 + 1
	}
	// Integrate positions
	(*b2SolverStage)(unsafe.Pointer(stage)).Type1 = int32(b2_stageIntegratePositions)
	(*b2SolverStage)(unsafe.Pointer(stage)).Blocks = bodyBlocks
	(*b2SolverStage)(unsafe.Pointer(stage)).BlockCount = bodyBlockCount
	(*b2SolverStage)(unsafe.Pointer(stage)).ColorIndex = -int32(1)
	atomicStoreNInt32(stage+24, 0, int32FromInt32(__ATOMIC_SEQ_CST))
	stage = stage + uintptr(1)*32
	// Relax constraints
	j4 = 0
	for {
		if !(j4 < int32(_RELAX_ITERATIONS)) {
			break
		}
		i9 = 0
		for {
			if !(i9 < activeColorCount) {
				break
			}
			(*b2SolverStage)(unsafe.Pointer(stage)).Type1 = int32(b2_stageRelax)
			(*b2SolverStage)(unsafe.Pointer(stage)).Blocks = graphColorBlocks[i9]
			(*b2SolverStage)(unsafe.Pointer(stage)).BlockCount = colorJointBlockCounts[i9] + colorContactBlockCounts[i9]
			(*b2SolverStage)(unsafe.Pointer(stage)).ColorIndex = activeColorIndices[i9]
			atomicStoreNInt32(stage+24, 0, int32FromInt32(__ATOMIC_SEQ_CST))
			stage = stage + uintptr(1)*32
			goto _28
		_28:
			;
			i9 = i9 + 1
		}
		goto _27
	_27:
		;
		j4 = j4 + 1
	}
	// Restitution
	// Note: joint blocks mixed in, could have joint limit restitution
	i10 = 0
	for {
		if !(i10 < activeColorCount) {
			break
		}
		(*b2SolverStage)(unsafe.Pointer(stage)).Type1 = int32(b2_stageRestitution)
		(*b2SolverStage)(unsafe.Pointer(stage)).Blocks = graphColorBlocks[i10]
		(*b2SolverStage)(unsafe.Pointer(stage)).BlockCount = colorJointBlockCounts[i10] + colorContactBlockCounts[i10]
		(*b2SolverStage)(unsafe.Pointer(stage)).ColorIndex = activeColorIndices[i10]
		atomicStoreNInt32(stage+24, 0, int32FromInt32(__ATOMIC_SEQ_CST))
		stage = stage + uintptr(1)*32
		goto _29
	_29:
		;
		i10 = i10 + 1
	}
	// Store impulses
	(*b2SolverStage)(unsafe.Pointer(stage)).Type1 = int32(b2_stageStoreImpulses)
	(*b2SolverStage)(unsafe.Pointer(stage)).Blocks = contactBlocks
	(*b2SolverStage)(unsafe.Pointer(stage)).BlockCount = contactBlockCount
	(*b2SolverStage)(unsafe.Pointer(stage)).ColorIndex = -int32(1)
	atomicStoreNInt32(stage+24, 0, int32FromInt32(__ATOMIC_SEQ_CST))
	stage = stage + uintptr(1)*32
	if !(int32((int64(stage)-int64(stages))/32) == stageCount) && b2InternalAssertFcn(tls, __ccgo_ts+13715, __ccgo_ts+12460, int32FromInt32(1676)) != 0 {
		__builtin_trap(tls)
	}
	if !(workerCount <= int32FromInt32(_B2_MAX_WORKERS)) && b2InternalAssertFcn(tls, __ccgo_ts+13753, __ccgo_ts+12460, int32FromInt32(1678)) != 0 {
		__builtin_trap(tls)
	}
	(*b2StepContext)(unsafe.Pointer(stepContext)).Graph = graph
	(*b2StepContext)(unsafe.Pointer(stepContext)).Joints = joints
	(*b2StepContext)(unsafe.Pointer(stepContext)).Contacts = contacts
	(*b2StepContext)(unsafe.Pointer(stepContext)).SimdContactConstraints = simdContactConstraints
	(*b2StepContext)(unsafe.Pointer(stepContext)).ActiveColorCount = activeColorCount
	(*b2StepContext)(unsafe.Pointer(stepContext)).WorkerCount = workerCount
	(*b2StepContext)(unsafe.Pointer(stepContext)).StageCount = stageCount
	(*b2StepContext)(unsafe.Pointer(stepContext)).Stages = stages
	atomicStoreNUint32(stepContext+232, uint32(0), int32FromInt32(__ATOMIC_SEQ_CST))
	(*b2World)(unsafe.Pointer(world)).Profile.PrepareStages = b2GetMillisecondsAndReset(tls, bp)
	*(*uint64_t)(unsafe.Add(nil, bp+1544)) = b2GetTicks(tls)
	// Must use worker index because thread 0 can be assigned multiple tasks by enkiTS
	i11 = 0
	for {
		if !(i11 < workerCount) {
			break
		}
		(*(*[64]b2WorkerContext)(unsafe.Add(nil, bp+8)))[i11].Context = stepContext
		(*(*[64]b2WorkerContext)(unsafe.Add(nil, bp+8)))[i11].WorkerIndex = i11
		(*(*[64]b2WorkerContext)(unsafe.Add(nil, bp+8)))[i11].UserTask = (*(*func(*_Stack, uintptr, int32, int32, uintptr, uintptr) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).EnqueueTaskFcn})))(tls, __ccgo_fp(b2SolverTask), int32(1), int32(1), bp+8+uintptr(i11)*24, (*b2World)(unsafe.Pointer(world)).UserTaskContext)
		*(*int32)(unsafe.Add(nil, world+1776)) += int32(1)
		if (*(*[64]b2WorkerContext)(unsafe.Add(nil, bp+8)))[i11].UserTask == uintptrFromInt32(0) {
			v2 = 0
		} else {
			v2 = int32(1)
		}
		*(*int32)(unsafe.Add(nil, world+1772)) += v2
		goto _30
	_30:
		;
		i11 = i11 + 1
	}
	// Finish island split
	if splitIslandTask != uintptrFromInt32(0) {
		(*(*func(*_Stack, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).FinishTaskFcn})))(tls, splitIslandTask, (*b2World)(unsafe.Pointer(world)).UserTaskContext)
		*(*int32)(unsafe.Add(nil, world+1772)) -= int32(1)
	}
	(*b2World)(unsafe.Pointer(world)).SplitIslandId = -int32(1)
	// Finish constraint solve
	i12 = 0
	for {
		if !(i12 < workerCount) {
			break
		}
		if (*(*[64]b2WorkerContext)(unsafe.Add(nil, bp+8)))[i12].UserTask != uintptrFromInt32(0) {
			(*(*func(*_Stack, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).FinishTaskFcn})))(tls, (*(*[64]b2WorkerContext)(unsafe.Add(nil, bp+8)))[i12].UserTask, (*b2World)(unsafe.Pointer(world)).UserTaskContext)
			*(*int32)(unsafe.Add(nil, world+1772)) -= int32(1)
		}
		goto _32
	_32:
		;
		i12 = i12 + 1
	}
	(*b2World)(unsafe.Pointer(world)).Profile.SolveConstraints = b2GetMillisecondsAndReset(tls, bp+1544)
	transformTicks = b2GetTicks(tls)
	// Prepare contact, enlarged body, and island bit sets used in body finalization.
	awakeIslandCount = (*b2SolverSet)(unsafe.Pointer(awakeSet)).IslandSims.Count
	i13 = 0
	for {
		if !(i13 < (*b2World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		taskContext = (*b2World)(unsafe.Pointer(world)).TaskContexts.Data + uintptr(i13)*56
		b2SetBitCountAndClear(tls, taskContext+16, uint32FromInt32(awakeBodyCount))
		b2SetBitCountAndClear(tls, taskContext+32, uint32FromInt32(awakeIslandCount))
		(*b2TaskContext)(unsafe.Pointer(taskContext)).SplitIslandId = -int32(1)
		(*b2TaskContext)(unsafe.Pointer(taskContext)).SplitSleepTime = float32FromFloat32(0)
		goto _33
	_33:
		;
		i13 = i13 + 1
	}
	// Finalize bodies. Must happen after the constraint solver and after island splitting.
	finalizeBodiesTask = (*(*func(*_Stack, uintptr, int32, int32, uintptr, uintptr) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).EnqueueTaskFcn})))(tls, __ccgo_fp(b2FinalizeBodiesTask), awakeBodyCount, int32(64), stepContext, (*b2World)(unsafe.Pointer(world)).UserTaskContext)
	*(*int32)(unsafe.Add(nil, world+1776)) += int32(1)
	if finalizeBodiesTask != uintptrFromInt32(0) {
		(*(*func(*_Stack, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).FinishTaskFcn})))(tls, finalizeBodiesTask, (*b2World)(unsafe.Pointer(world)).UserTaskContext)
	}
	b2FreeArenaItem(tls, world, graphBlocks)
	b2FreeArenaItem(tls, world, jointBlocks)
	b2FreeArenaItem(tls, world, contactBlocks)
	b2FreeArenaItem(tls, world, bodyBlocks)
	b2FreeArenaItem(tls, world, stages)
	b2FreeArenaItem(tls, world, overflowContactConstraints)
	b2FreeArenaItem(tls, world, simdContactConstraints)
	b2FreeArenaItem(tls, world, joints)
	b2FreeArenaItem(tls, world, contacts)
	(*b2World)(unsafe.Pointer(world)).Profile.Transforms = b2GetMilliseconds(tls, transformTicks)
	// Report hit events
	// todo_erin perhaps optimize this with a bitset
	// todo_erin perhaps do this in parallel with other work below
	hitTicks = b2GetTicks(tls)
	if !((*b2World)(unsafe.Pointer(world)).ContactHitEvents.Count == int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+13783, __ccgo_ts+12460, int32FromInt32(1772)) != 0 {
		__builtin_trap(tls)
	}
	threshold = (*b2World)(unsafe.Pointer(world)).HitEventThreshold
	colors1 = world + 328
	i14 = 0
	for {
		if !(i14 < int32(_B2_GRAPH_COLOR_COUNT)) {
			break
		}
		color1 = colors1 + uintptr(i14)*56
		contactCount = (*b2GraphColor)(unsafe.Pointer(color1)).ContactSims.Count
		contactSims = (*b2GraphColor)(unsafe.Pointer(color1)).ContactSims.Data
		j5 = 0
		for {
			if !(j5 < contactCount) {
				break
			}
			contactSim = contactSims + uintptr(j5)*176
			if (*b2ContactSim)(unsafe.Pointer(contactSim)).SimFlags&uint32(b2_simEnableHitEvent) == uint32(0) {
				goto _35
			}
			event = ContactHitEvent{}
			event.ApproachSpeed = threshold
			hit = boolUint8(false1 != 0)
			pointCount = (*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.PointCount
			k3 = 0
			for {
				if !(k3 < pointCount) {
					break
				}
				mp = contactSim + 36 + 12 + uintptr(k3)*48
				approachSpeed = -(*ManifoldPoint)(unsafe.Pointer(mp)).NormalVelocity
				// Need to check total impulse because the point may be speculative and not colliding
				if approachSpeed > event.ApproachSpeed && (*ManifoldPoint)(unsafe.Pointer(mp)).TotalNormalImpulse > float32FromFloat32(0) {
					event.ApproachSpeed = approachSpeed
					event.Point = (*ManifoldPoint)(unsafe.Pointer(mp)).Point
					hit = boolUint8(true1 != 0)
				}
				goto _36
			_36:
				;
				k3 = k3 + 1
			}
			if int32FromUint8(hit) == int32(true1) {
				event.Normal = (*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.Normal
				v1 = world + 1248
				v2 = (*b2ContactSim)(unsafe.Pointer(contactSim)).ShapeIdA
				if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
					__builtin_trap(tls)
				}
				v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
				goto _40
			_40:
				shapeA = v3
				v7 = world + 1248
				v6 = (*b2ContactSim)(unsafe.Pointer(contactSim)).ShapeIdB
				if !(0 <= v6 && v6 < (*b2ShapeArray)(unsafe.Pointer(v7)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
					__builtin_trap(tls)
				}
				v37 = (*b2ShapeArray)(unsafe.Pointer(v7)).Data + uintptr(v6)*288
				goto _44
			_44:
				shapeB = v37
				event.ShapeIdA = ShapeId{
					Index1:     (*b2Shape)(unsafe.Pointer(shapeA)).Id + int32(1),
					World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
					Generation: (*b2Shape)(unsafe.Pointer(shapeA)).Generation,
				}
				event.ShapeIdB = ShapeId{
					Index1:     (*b2Shape)(unsafe.Pointer(shapeB)).Id + int32(1),
					World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
					Generation: (*b2Shape)(unsafe.Pointer(shapeB)).Generation,
				}
				v1 = world + 1448
				if (*b2ContactHitEventArray)(unsafe.Pointer(v1)).Count == (*b2ContactHitEventArray)(unsafe.Pointer(v1)).Capacity {
					if (*b2ContactHitEventArray)(unsafe.Pointer(v1)).Capacity < int32(2) {
						v2 = int32(2)
					} else {
						v2 = (*b2ContactHitEventArray)(unsafe.Pointer(v1)).Capacity + (*b2ContactHitEventArray)(unsafe.Pointer(v1)).Capacity>>int32(1)
					}
					newCapacity1 = v2
					b2ContactHitEventArray_Reserve(tls, v1, newCapacity1)
				}
				*(*ContactHitEvent)(unsafe.Add(nil, (*b2ContactHitEventArray)(unsafe.Pointer(v1)).Data+uintptr((*b2ContactHitEventArray)(unsafe.Pointer(v1)).Count)*36)) = event
				*(*int32)(unsafe.Add(nil, v1+8)) += int32(1)
			}
			goto _35
		_35:
			;
			j5 = j5 + 1
		}
		goto _34
	_34:
		;
		i14 = i14 + 1
	}
	(*b2World)(unsafe.Pointer(world)).Profile.HitEvents = b2GetMilliseconds(tls, hitTicks)
	refitTicks = b2GetTicks(tls)
	// Finish the user tree task that was queued earlier in the time step. This must be complete before touching the
	// broad-phase.
	if (*b2World)(unsafe.Pointer(world)).UserTreeTask != uintptrFromInt32(0) {
		(*(*func(*_Stack, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).FinishTaskFcn})))(tls, (*b2World)(unsafe.Pointer(world)).UserTreeTask, (*b2World)(unsafe.Pointer(world)).UserTaskContext)
		(*b2World)(unsafe.Pointer(world)).UserTreeTask = uintptrFromInt32(0)
		*(*int32)(unsafe.Add(nil, world+1772)) -= int32(1)
	}
	b2ValidateNoEnlarged(tls, world+40)
	// Gather bits for all sim bodies that have enlarged AABBs
	enlargedBodyBitSet = (*b2World)(unsafe.Pointer(world)).TaskContexts.Data + 16
	i15 = int32(1)
	for {
		if !(i15 < (*b2World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		b2InPlaceUnion(tls, enlargedBodyBitSet, (*b2World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i15)*56+16)
		goto _47
	_47:
		;
		i15 = i15 + 1
	}
	// Enlarge broad-phase proxies and build move array
	// Apply shape AABB changes to broad-phase. This also create the move array which must be
	// in deterministic order. I'm tracking sim bodies because the number of shape ids can be huge.
	// This has to happen before bullets are processed.
	broadPhase = world + 40
	wordCount = (*b2BitSet)(unsafe.Pointer(enlargedBodyBitSet)).BlockCount
	bits = (*b2BitSet)(unsafe.Pointer(enlargedBodyBitSet)).Bits
	// Fast array access is important here
	bodyArray = (*b2World)(unsafe.Pointer(world)).Bodies.Data
	bodySimArray = (*b2SolverSet)(unsafe.Pointer(awakeSet)).BodySims.Data
	shapeArray = (*b2World)(unsafe.Pointer(world)).Shapes.Data
	k4 = uint32(0)
	for {
		if !(k4 < wordCount) {
			break
		}
		word = *(*uint64_t)(unsafe.Add(nil, bits+uintptr(k4)*8))
		for word != uint64(0) {
			v49 = uint32FromInt32(__builtin_ctzll(tls, word))
			goto _50
		_50:
			ctz = v49
			bodySimIndex = uint32(64)*k4 + ctz
			bodySim = bodySimArray + uintptr(bodySimIndex)*100
			body = bodyArray + uintptr((*b2BodySim)(unsafe.Pointer(bodySim)).BodyId)*128
			shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
			if (*b2BodySim)(unsafe.Pointer(bodySim)).IsBullet != 0 && (*b2BodySim)(unsafe.Pointer(bodySim)).IsFast != 0 {
				// Fast bullet bodies don't have their final AABB yet
				for shapeId != -int32(1) {
					shape = shapeArray + uintptr(shapeId)*288
					// Shape is fast. It's aabb will be enlarged in continuous collision.
					// Update the move array here for determinism because bullets are processed
					// below in non-deterministic order.
					v1 = broadPhase
					v2 = (*b2Shape)(unsafe.Pointer(shape)).ProxyKey
					alreadyAdded = b2AddKey(tls, v1+216, uint64FromInt32(v2+int32(1)))
					if int32FromUint8(alreadyAdded) == int32FromInt32(false1) {
						v3 = v1 + 232
						if (*b2IntArray)(unsafe.Pointer(v3)).Count == (*b2IntArray)(unsafe.Pointer(v3)).Capacity {
							if (*b2IntArray)(unsafe.Pointer(v3)).Capacity < int32(2) {
								v6 = int32(2)
							} else {
								v6 = (*b2IntArray)(unsafe.Pointer(v3)).Capacity + (*b2IntArray)(unsafe.Pointer(v3)).Capacity>>int32(1)
							}
							newCapacity = v6
							b2IntArray_Reserve(tls, v3, newCapacity)
						}
						*(*int32)(unsafe.Add(nil, (*b2IntArray)(unsafe.Pointer(v3)).Data+uintptr((*b2IntArray)(unsafe.Pointer(v3)).Count)*4)) = v2
						*(*int32)(unsafe.Add(nil, v3+8)) += int32(1)
					}
					shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
				}
			} else {
				for shapeId != -int32(1) {
					shape1 = shapeArray + uintptr(shapeId)*288
					// The AABB may not have been enlarged, despite the body being flagged as enlarged.
					// For example, a body with multiple shapes may have not have all shapes enlarged.
					// A fast body may have been flagged as enlarged despite having no shapes enlarged.
					if (*b2Shape)(unsafe.Pointer(shape1)).EnlargedAABB != 0 {
						b2BroadPhase_EnlargeProxy(tls, broadPhase, (*b2Shape)(unsafe.Pointer(shape1)).ProxyKey, (*b2Shape)(unsafe.Pointer(shape1)).FatAABB)
						(*b2Shape)(unsafe.Pointer(shape1)).EnlargedAABB = boolUint8(false1 != 0)
					}
					shapeId = (*b2Shape)(unsafe.Pointer(shape1)).NextShapeId
				}
			}
			// Clear the smallest set bit
			word = word & (word - uint64(1))
		}
		goto _48
	_48:
		;
		k4 = k4 + 1
	}
	b2ValidateBroadphase(tls, world+40)
	(*b2World)(unsafe.Pointer(world)).Profile.Refit = b2GetMilliseconds(tls, refitTicks)
	v2 = atomicLoadNInt32(stepContext+112, int32FromInt32(__ATOMIC_SEQ_CST))
	goto _56
_56:
	bulletBodyCount = v2
	if bulletBodyCount > 0 {
		bulletTicks = b2GetTicks(tls)
		// Fast bullet bodies
		// Note: a bullet body may be moving slow
		minRange = int32(8)
		userBulletBodyTask = (*(*func(*_Stack, uintptr, int32, int32, uintptr, uintptr) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).EnqueueTaskFcn})))(tls, __ccgo_fp(b2BulletBodyTask), bulletBodyCount, minRange, stepContext, (*b2World)(unsafe.Pointer(world)).UserTaskContext)
		*(*int32)(unsafe.Add(nil, world+1776)) += int32(1)
		if userBulletBodyTask != uintptrFromInt32(0) {
			(*(*func(*_Stack, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).FinishTaskFcn})))(tls, userBulletBodyTask, (*b2World)(unsafe.Pointer(world)).UserTaskContext)
		}
		// Serially enlarge broad-phase proxies for bullet shapes
		broadPhase1 = world + 40
		dynamicTree = broadPhase1 + uintptr(b2_dynamicBody)*72
		// Fast array access is important here
		bodyArray1 = (*b2World)(unsafe.Pointer(world)).Bodies.Data
		bodySimArray1 = (*b2SolverSet)(unsafe.Pointer(awakeSet)).BodySims.Data
		shapeArray1 = (*b2World)(unsafe.Pointer(world)).Shapes.Data
		// Serially enlarge broad-phase proxies for bullet shapes
		bulletBodySimIndices = (*b2StepContext)(unsafe.Pointer(stepContext)).BulletBodies
		// This loop has non-deterministic order but it shouldn't affect the result
		i16 = 0
		for {
			if !(i16 < bulletBodyCount) {
				break
			}
			bulletBodySim = bodySimArray1 + uintptr(*(*int32)(unsafe.Add(nil, bulletBodySimIndices+uintptr(i16)*4)))*100
			if int32FromUint8((*b2BodySim)(unsafe.Pointer(bulletBodySim)).EnlargeAABB) == false1 {
				goto _57
			}
			// clear flag
			(*b2BodySim)(unsafe.Pointer(bulletBodySim)).EnlargeAABB = boolUint8(false1 != 0)
			bodyId = (*b2BodySim)(unsafe.Pointer(bulletBodySim)).BodyId
			if !(0 <= bodyId && bodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count) && b2InternalAssertFcn(tls, __ccgo_ts+13818, __ccgo_ts+12460, int32FromInt32(1964)) != 0 {
				__builtin_trap(tls)
			}
			bulletBody = bodyArray1 + uintptr(bodyId)*128
			shapeId1 = (*b2Body)(unsafe.Pointer(bulletBody)).HeadShapeId
			for shapeId1 != -int32(1) {
				shape2 = shapeArray1 + uintptr(shapeId1)*288
				if int32FromUint8((*b2Shape)(unsafe.Pointer(shape2)).EnlargedAABB) == false1 {
					shapeId1 = (*b2Shape)(unsafe.Pointer(shape2)).NextShapeId
					continue
				}
				// clear flag
				(*b2Shape)(unsafe.Pointer(shape2)).EnlargedAABB = boolUint8(false1 != 0)
				proxyKey = (*b2Shape)(unsafe.Pointer(shape2)).ProxyKey
				proxyId = proxyKey >> int32(2)
				if !(proxyKey&int32FromInt32(3) == int32(b2_dynamicBody)) && b2InternalAssertFcn(tls, __ccgo_ts+13862, __ccgo_ts+12460, int32FromInt32(1982)) != 0 {
					__builtin_trap(tls)
				}
				// all fast bullet shapes should already be in the move buffer
				if !(b2ContainsKey(tls, broadPhase1+216, uint64FromInt32(proxyKey+int32(1))) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+13906, __ccgo_ts+12460, int32FromInt32(1985)) != 0 {
					__builtin_trap(tls)
				}
				b2DynamicTree_EnlargeProxy(tls, dynamicTree, proxyId, (*b2Shape)(unsafe.Pointer(shape2)).FatAABB)
				shapeId1 = (*b2Shape)(unsafe.Pointer(shape2)).NextShapeId
			}
			goto _57
		_57:
			;
			i16 = i16 + 1
		}
		(*b2World)(unsafe.Pointer(world)).Profile.Bullets = b2GetMilliseconds(tls, bulletTicks)
	}
	// Need to free this even if no bullets got processed.
	b2FreeArenaItem(tls, world, (*b2StepContext)(unsafe.Pointer(stepContext)).BulletBodies)
	(*b2StepContext)(unsafe.Pointer(stepContext)).BulletBodies = uintptrFromInt32(0)
	atomicStoreNInt32(stepContext+112, 0, int32FromInt32(__ATOMIC_SEQ_CST))
	// Island sleeping
	// This must be done last because putting islands to sleep invalidates the enlarged body bits.
	// todo_erin figure out how to do this in parallel with tree refit
	if int32FromUint8((*b2World)(unsafe.Pointer(world)).EnableSleep) == int32(true1) {
		sleepTicks = b2GetTicks(tls)
		// Collect split island candidate for the next time step. No need to split if sleeping is disabled.
		if !((*b2World)(unsafe.Pointer(world)).SplitIslandId == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+13958, __ccgo_ts+12460, int32FromInt32(2011)) != 0 {
			__builtin_trap(tls)
		}
		splitSleepTimer = float32FromFloat32(0)
		i17 = 0
		for {
			if !(i17 < (*b2World)(unsafe.Pointer(world)).WorkerCount) {
				break
			}
			taskContext1 = (*b2World)(unsafe.Pointer(world)).TaskContexts.Data + uintptr(i17)*56
			if (*b2TaskContext)(unsafe.Pointer(taskContext1)).SplitIslandId != -int32(1) && (*b2TaskContext)(unsafe.Pointer(taskContext1)).SplitSleepTime >= splitSleepTimer {
				if !((*b2TaskContext)(unsafe.Pointer(taskContext1)).SplitSleepTime > float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+13996, __ccgo_ts+12460, int32FromInt32(2018)) != 0 {
					__builtin_trap(tls)
				}
				// Tie breaking for determinism. Largest island id wins. Needed due to work stealing.
				if (*b2TaskContext)(unsafe.Pointer(taskContext1)).SplitSleepTime == splitSleepTimer && (*b2TaskContext)(unsafe.Pointer(taskContext1)).SplitIslandId < (*b2World)(unsafe.Pointer(world)).SplitIslandId {
					goto _58
				}
				(*b2World)(unsafe.Pointer(world)).SplitIslandId = (*b2TaskContext)(unsafe.Pointer(taskContext1)).SplitIslandId
				splitSleepTimer = (*b2TaskContext)(unsafe.Pointer(taskContext1)).SplitSleepTime
			}
			goto _58
		_58:
			;
			i17 = i17 + 1
		}
		awakeIslandBitSet = (*b2World)(unsafe.Pointer(world)).TaskContexts.Data + 32
		i18 = int32(1)
		for {
			if !(i18 < (*b2World)(unsafe.Pointer(world)).WorkerCount) {
				break
			}
			b2InPlaceUnion(tls, awakeIslandBitSet, (*b2World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i18)*56+32)
			goto _59
		_59:
			;
			i18 = i18 + 1
		}
		// Need to process in reverse because this moves islands to sleeping solver sets.
		islands = (*b2SolverSet)(unsafe.Pointer(awakeSet)).IslandSims.Data
		count1 = (*b2SolverSet)(unsafe.Pointer(awakeSet)).IslandSims.Count
		islandIndex = count1 - int32(1)
		for {
			if !(islandIndex >= 0) {
				break
			}
			v1 = awakeIslandBitSet
			v49 = uint32FromInt32(islandIndex)
			blockIndex = v49 / uint32(64)
			if blockIndex >= (*b2BitSet)(unsafe.Pointer(v1)).BlockCount {
				v63 = boolUint8(false1 != 0)
				goto _64
			}
			v63 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v1)).Bits+uintptr(blockIndex)*8))&(uint64FromInt32(1)<<(v49%uint32FromInt32(64))) != uint64(0))
			goto _64
		_64:
			if int32FromUint8(v63) == int32(true1) {
				// this island is still awake
				goto _60
			}
			island = islands + uintptr(islandIndex)*4
			islandId = (*b2IslandSim)(unsafe.Pointer(island)).IslandId
			b2TrySleepIsland(tls, world, islandId)
			goto _60
		_60:
			;
			islandIndex = islandIndex - int32(1)
		}
		b2ValidateSolverSets(tls, world)
		(*b2World)(unsafe.Pointer(world)).Profile.SleepIslands = b2GetMilliseconds(tls, sleepTicks)
	}
}

func b2SolverSetArray_Create(tls *_Stack, capacity int32) (r b2SolverSetArray) {
	var a b2SolverSetArray
	_ = a
	a = b2SolverSetArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(88)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2SolverSetArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2SolverSetArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2SolverSetArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2SolverSetArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2SolverSetArray)(unsafe.Pointer(a)).Capacity)*uint64(88)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(88)))
	(*b2SolverSetArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2SolverSetArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2SolverSetArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2SolverSetArray)(unsafe.Pointer(a)).Capacity)*uint64(88)))
	(*b2SolverSetArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2SolverSetArray)(unsafe.Pointer(a)).Count = 0
	(*b2SolverSetArray)(unsafe.Pointer(a)).Capacity = 0
}

func b2DestroySolverSet(tls *_Stack, world uintptr, setIndex int32) {
	var set, v1, v3 uintptr
	var v2 int32
	_, _, _, _ = set, v1, v2, v3
	v1 = world + 1064
	v2 = setIndex
	if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*88
	goto _4
_4:
	set = v3
	b2BodySimArray_Destroy(tls, set)
	b2BodyStateArray_Destroy(tls, set+16)
	b2ContactSimArray_Destroy(tls, set+48)
	b2JointSimArray_Destroy(tls, set+32)
	b2IslandSimArray_Destroy(tls, set+64)
	b2FreeId(tls, world+1040, setIndex)
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
	var awakeContactSim, awakeSet, bodies, body, contact, contact1, contactSim, contactSim1, disabledSet, island, islandDst, islandSrc, joint, jointSim, movedContact, movedContactSim, set, simDst, simSrc, state, v1, v11, v3, v5, v7, v9 uintptr
	var bodyCount, contactCount, contactId, contactKey, edgeIndex, i, i1, i2, i3, islandCount, jointCount, localIndex, movedIndex, movedLocalIndex, newCapacity, newCapacity1, newCapacity2, newCapacity3, v10, v2, v6 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = awakeContactSim, awakeSet, bodies, body, bodyCount, contact, contact1, contactCount, contactId, contactKey, contactSim, contactSim1, disabledSet, edgeIndex, i, i1, i2, i3, island, islandCount, islandDst, islandSrc, joint, jointCount, jointSim, localIndex, movedContact, movedContactSim, movedIndex, movedLocalIndex, newCapacity, newCapacity1, newCapacity2, newCapacity3, set, simDst, simSrc, state, v1, v10, v11, v2, v3, v5, v6, v7, v9
	if !(setIndex >= int32(b2_firstSleepingSet)) && b2InternalAssertFcn(tls, __ccgo_ts+14031, __ccgo_ts+14063, int32FromInt32(39)) != 0 {
		__builtin_trap(tls)
	}
	v1 = world + 1064
	v2 = setIndex
	if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*88
	goto _4
_4:
	set = v3
	v5 = world + 1064
	v6 = int32(b2_awakeSet)
	if !(0 <= v6 && v6 < (*b2SolverSetArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2SolverSetArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*88
	goto _8
_8:
	awakeSet = v7
	v9 = world + 1064
	v10 = int32(b2_disabledSet)
	if !(0 <= v10 && v10 < (*b2SolverSetArray)(unsafe.Pointer(v9)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v11 = (*b2SolverSetArray)(unsafe.Pointer(v9)).Data + uintptr(v10)*88
	goto _12
_12:
	disabledSet = v11
	bodies = (*b2World)(unsafe.Pointer(world)).Bodies.Data
	bodyCount = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count
	i = 0
	for {
		if !(i < bodyCount) {
			break
		}
		simSrc = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr(i)*100
		body = bodies + uintptr((*b2BodySim)(unsafe.Pointer(simSrc)).BodyId)*128
		if !((*b2Body)(unsafe.Pointer(body)).SetIndex == setIndex) && b2InternalAssertFcn(tls, __ccgo_ts+14091, __ccgo_ts+14063, int32FromInt32(52)) != 0 {
			__builtin_trap(tls)
		}
		(*b2Body)(unsafe.Pointer(body)).SetIndex = int32(b2_awakeSet)
		(*b2Body)(unsafe.Pointer(body)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(awakeSet)).BodySims.Count
		// Reset sleep timer
		(*b2Body)(unsafe.Pointer(body)).SleepTime = float32FromFloat32(0)
		v1 = awakeSet
		if (*b2BodySimArray)(unsafe.Pointer(v1)).Count == (*b2BodySimArray)(unsafe.Pointer(v1)).Capacity {
			if (*b2BodySimArray)(unsafe.Pointer(v1)).Capacity < int32(2) {
				v2 = int32(2)
			} else {
				v2 = (*b2BodySimArray)(unsafe.Pointer(v1)).Capacity + (*b2BodySimArray)(unsafe.Pointer(v1)).Capacity>>int32(1)
			}
			newCapacity = v2
			b2BodySimArray_Reserve(tls, v1, newCapacity)
		}
		*(*int32)(unsafe.Add(nil, v1+8)) += int32(1)
		v3 = (*b2BodySimArray)(unsafe.Pointer(v1)).Data + uintptr((*b2BodySimArray)(unsafe.Pointer(v1)).Count-int32FromInt32(1))*100
		goto _17
	_17:
		simDst = v3
		memcpy(tls, simDst, simSrc, uint64(100))
		v1 = awakeSet + 16
		if (*b2BodyStateArray)(unsafe.Pointer(v1)).Count == (*b2BodyStateArray)(unsafe.Pointer(v1)).Capacity {
			if (*b2BodyStateArray)(unsafe.Pointer(v1)).Capacity < int32(2) {
				v2 = int32(2)
			} else {
				v2 = (*b2BodyStateArray)(unsafe.Pointer(v1)).Capacity + (*b2BodyStateArray)(unsafe.Pointer(v1)).Capacity>>int32(1)
			}
			newCapacity1 = v2
			b2BodyStateArray_Reserve(tls, v1, newCapacity1)
		}
		*(*int32)(unsafe.Add(nil, v1+8)) += int32(1)
		v3 = (*b2BodyStateArray)(unsafe.Pointer(v1)).Data + uintptr((*b2BodyStateArray)(unsafe.Pointer(v1)).Count-int32FromInt32(1))*32
		goto _21
	_21:
		state = v3
		*(*b2BodyState)(unsafe.Add(nil, state)) = b2_identityBodyState15
		// move non-touching contacts from disabled set to awake set
		contactKey = (*b2Body)(unsafe.Pointer(body)).HeadContactKey
		for contactKey != -int32(1) {
			edgeIndex = contactKey & int32(1)
			contactId = contactKey >> int32(1)
			v1 = world + 1144
			v2 = contactId
			if !(0 <= v2 && v2 < (*b2ContactArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
				__builtin_trap(tls)
			}
			v3 = (*b2ContactArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*68
			goto _25
		_25:
			contact = v3
			contactKey = (*(*b2ContactEdge)(unsafe.Add(nil, contact+12+uintptr(edgeIndex)*12))).NextKey
			if (*b2Contact)(unsafe.Pointer(contact)).SetIndex != int32(b2_disabledSet) {
				if !((*b2Contact)(unsafe.Pointer(contact)).SetIndex == int32(b2_awakeSet) || (*b2Contact)(unsafe.Pointer(contact)).SetIndex == setIndex) && b2InternalAssertFcn(tls, __ccgo_ts+14118, __ccgo_ts+14063, int32FromInt32(78)) != 0 {
					__builtin_trap(tls)
				}
				continue
			}
			localIndex = (*b2Contact)(unsafe.Pointer(contact)).LocalIndex
			v1 = disabledSet + 48
			v2 = localIndex
			if !(0 <= v2 && v2 < (*b2ContactSimArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(147)) != 0 {
				__builtin_trap(tls)
			}
			v3 = (*b2ContactSimArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*176
			goto _29
		_29:
			contactSim = v3
			if !((*b2Contact)(unsafe.Pointer(contact)).Flags&uint32(b2_contactTouchingFlag) == uint32(0) && (*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.PointCount == 0) && b2InternalAssertFcn(tls, __ccgo_ts+14184, __ccgo_ts+14063, int32FromInt32(85)) != 0 {
				__builtin_trap(tls)
			}
			(*b2Contact)(unsafe.Pointer(contact)).SetIndex = int32(b2_awakeSet)
			(*b2Contact)(unsafe.Pointer(contact)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(awakeSet)).ContactSims.Count
			v1 = awakeSet + 48
			if (*b2ContactSimArray)(unsafe.Pointer(v1)).Count == (*b2ContactSimArray)(unsafe.Pointer(v1)).Capacity {
				if (*b2ContactSimArray)(unsafe.Pointer(v1)).Capacity < int32(2) {
					v2 = int32(2)
				} else {
					v2 = (*b2ContactSimArray)(unsafe.Pointer(v1)).Capacity + (*b2ContactSimArray)(unsafe.Pointer(v1)).Capacity>>int32(1)
				}
				newCapacity2 = v2
				b2ContactSimArray_Reserve(tls, v1, newCapacity2)
			}
			*(*int32)(unsafe.Add(nil, v1+8)) += int32(1)
			v3 = (*b2ContactSimArray)(unsafe.Pointer(v1)).Data + uintptr((*b2ContactSimArray)(unsafe.Pointer(v1)).Count-int32FromInt32(1))*176
			goto _33
		_33:
			awakeContactSim = v3
			memcpy(tls, awakeContactSim, contactSim, uint64(176))
			v1 = disabledSet + 48
			v2 = localIndex
			if !(0 <= v2 && v2 < (*b2ContactSimArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(147)) != 0 {
				__builtin_trap(tls)
			}
			movedIndex = -int32(1)
			if v2 != (*b2ContactSimArray)(unsafe.Pointer(v1)).Count-int32FromInt32(1) {
				movedIndex = (*b2ContactSimArray)(unsafe.Pointer(v1)).Count - int32(1)
				*(*b2ContactSim)(unsafe.Add(nil, (*b2ContactSimArray)(unsafe.Pointer(v1)).Data+uintptr(v2)*176)) = *(*b2ContactSim)(unsafe.Add(nil, (*b2ContactSimArray)(unsafe.Pointer(v1)).Data+uintptr(movedIndex)*176))
			}
			*(*int32)(unsafe.Add(nil, v1+8)) -= int32(1)
			v6 = movedIndex
			goto _37
		_37:
			movedLocalIndex = v6
			if movedLocalIndex != -int32(1) {
				// fix moved element
				movedContactSim = (*b2SolverSet)(unsafe.Pointer(disabledSet)).ContactSims.Data + uintptr(localIndex)*176
				v3 = world + 1144
				v10 = (*b2ContactSim)(unsafe.Pointer(movedContactSim)).ContactId
				if !(0 <= v10 && v10 < (*b2ContactArray)(unsafe.Pointer(v3)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
					__builtin_trap(tls)
				}
				v5 = (*b2ContactArray)(unsafe.Pointer(v3)).Data + uintptr(v10)*68
				goto _41
			_41:
				movedContact = v5
				if !((*b2Contact)(unsafe.Pointer(movedContact)).LocalIndex == movedLocalIndex) && b2InternalAssertFcn(tls, __ccgo_ts+14273, __ccgo_ts+14063, int32FromInt32(98)) != 0 {
					__builtin_trap(tls)
				}
				(*b2Contact)(unsafe.Pointer(movedContact)).LocalIndex = localIndex
			}
		}
		goto _13
	_13:
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
		contactSim1 = (*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Data + uintptr(i1)*176
		v1 = world + 1144
		v2 = (*b2ContactSim)(unsafe.Pointer(contactSim1)).ContactId
		if !(0 <= v2 && v2 < (*b2ContactArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2ContactArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*68
		goto _46
	_46:
		contact1 = v3
		if !((*b2Contact)(unsafe.Pointer(contact1)).Flags&uint32(b2_contactTouchingFlag) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+2899, __ccgo_ts+14063, int32FromInt32(111)) != 0 {
			__builtin_trap(tls)
		}
		if !((*b2ContactSim)(unsafe.Pointer(contactSim1)).SimFlags&uint32(b2_simTouchingFlag) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+2857, __ccgo_ts+14063, int32FromInt32(112)) != 0 {
			__builtin_trap(tls)
		}
		if !((*b2ContactSim)(unsafe.Pointer(contactSim1)).Manifold.PointCount > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+2821, __ccgo_ts+14063, int32FromInt32(113)) != 0 {
			__builtin_trap(tls)
		}
		if !((*b2Contact)(unsafe.Pointer(contact1)).SetIndex == setIndex) && b2InternalAssertFcn(tls, __ccgo_ts+14317, __ccgo_ts+14063, int32FromInt32(114)) != 0 {
			__builtin_trap(tls)
		}
		b2AddContactToGraph(tls, world, contactSim1, contact1)
		(*b2Contact)(unsafe.Pointer(contact1)).SetIndex = int32(b2_awakeSet)
		goto _42
	_42:
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
		jointSim = (*b2SolverSet)(unsafe.Pointer(set)).JointSims.Data + uintptr(i2)*196
		v1 = world + 1104
		v2 = (*b2JointSim)(unsafe.Pointer(jointSim)).JointId
		if !(0 <= v2 && v2 < (*b2JointArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2JointArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*72
		goto _51
	_51:
		joint = v3
		if !((*b2Joint)(unsafe.Pointer(joint)).SetIndex == setIndex) && b2InternalAssertFcn(tls, __ccgo_ts+9243, __ccgo_ts+14063, int32FromInt32(127)) != 0 {
			__builtin_trap(tls)
		}
		b2AddJointToGraph(tls, world, jointSim, joint)
		(*b2Joint)(unsafe.Pointer(joint)).SetIndex = int32(b2_awakeSet)
		goto _47
	_47:
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
		v1 = world + 1184
		v2 = (*b2IslandSim)(unsafe.Pointer(islandSrc)).IslandId
		if !(0 <= v2 && v2 < (*b2IslandArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2IslandArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*56
		goto _56
	_56:
		island = v3
		(*b2Island)(unsafe.Pointer(island)).SetIndex = int32(b2_awakeSet)
		(*b2Island)(unsafe.Pointer(island)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(awakeSet)).IslandSims.Count
		v1 = awakeSet + 64
		if (*b2IslandSimArray)(unsafe.Pointer(v1)).Count == (*b2IslandSimArray)(unsafe.Pointer(v1)).Capacity {
			if (*b2IslandSimArray)(unsafe.Pointer(v1)).Capacity < int32(2) {
				v2 = int32(2)
			} else {
				v2 = (*b2IslandSimArray)(unsafe.Pointer(v1)).Capacity + (*b2IslandSimArray)(unsafe.Pointer(v1)).Capacity>>int32(1)
			}
			newCapacity3 = v2
			b2IslandSimArray_Reserve(tls, v1, newCapacity3)
		}
		*(*int32)(unsafe.Add(nil, v1+8)) += int32(1)
		v3 = (*b2IslandSimArray)(unsafe.Pointer(v1)).Data + uintptr((*b2IslandSimArray)(unsafe.Pointer(v1)).Count-int32FromInt32(1))*4
		goto _60
	_60:
		islandDst = v3
		memcpy(tls, islandDst, islandSrc, uint64(4))
		goto _52
	_52:
		;
		i3 = i3 + 1
	}
	// destroy the sleeping set
	b2DestroySolverSet(tls, world, setIndex)
	b2ValidateSolverSets(tls, world)
}

// C documentation
//
//	// This is called when joints are created between sets. I want to allow the sets
//	// to continue sleeping if both are asleep. Otherwise one set is waked.
//	// Islands will get merge when the set is waked.
func b2MergeSolverSets(tls *_Stack, world uintptr, setId1 int32, setId2 int32) {
	var bodies, body, contact, contactDst, contactSrc, island, islandDst, islandSrc, joint, jointDst, jointSrc, set1, set2, simDst, simSrc, tempSet, v1, v3, v5, v7 uintptr
	var bodyCount, contactCount, i, i1, i2, i3, islandCount, islandId, jointCount, newCapacity, newCapacity1, newCapacity2, newCapacity3, tempId, v2, v6 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodies, body, bodyCount, contact, contactCount, contactDst, contactSrc, i, i1, i2, i3, island, islandCount, islandDst, islandId, islandSrc, joint, jointCount, jointDst, jointSrc, newCapacity, newCapacity1, newCapacity2, newCapacity3, set1, set2, simDst, simSrc, tempId, tempSet, v1, v2, v3, v5, v6, v7
	if !(setId1 >= int32(b2_firstSleepingSet)) && b2InternalAssertFcn(tls, __ccgo_ts+14829, __ccgo_ts+14063, int32FromInt32(429)) != 0 {
		__builtin_trap(tls)
	}
	if !(setId2 >= int32(b2_firstSleepingSet)) && b2InternalAssertFcn(tls, __ccgo_ts+14859, __ccgo_ts+14063, int32FromInt32(430)) != 0 {
		__builtin_trap(tls)
	}
	v1 = world + 1064
	v2 = setId1
	if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*88
	goto _4
_4:
	set1 = v3
	v5 = world + 1064
	v6 = setId2
	if !(0 <= v6 && v6 < (*b2SolverSetArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2SolverSetArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*88
	goto _8
_8:
	set2 = v7
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
		simSrc = (*b2SolverSet)(unsafe.Pointer(set2)).BodySims.Data + uintptr(i)*100
		body = bodies + uintptr((*b2BodySim)(unsafe.Pointer(simSrc)).BodyId)*128
		if !((*b2Body)(unsafe.Pointer(body)).SetIndex == setId2) && b2InternalAssertFcn(tls, __ccgo_ts+14889, __ccgo_ts+14063, int32FromInt32(455)) != 0 {
			__builtin_trap(tls)
		}
		(*b2Body)(unsafe.Pointer(body)).SetIndex = setId1
		(*b2Body)(unsafe.Pointer(body)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(set1)).BodySims.Count
		v1 = set1
		if (*b2BodySimArray)(unsafe.Pointer(v1)).Count == (*b2BodySimArray)(unsafe.Pointer(v1)).Capacity {
			if (*b2BodySimArray)(unsafe.Pointer(v1)).Capacity < int32(2) {
				v2 = int32(2)
			} else {
				v2 = (*b2BodySimArray)(unsafe.Pointer(v1)).Capacity + (*b2BodySimArray)(unsafe.Pointer(v1)).Capacity>>int32(1)
			}
			newCapacity = v2
			b2BodySimArray_Reserve(tls, v1, newCapacity)
		}
		*(*int32)(unsafe.Add(nil, v1+8)) += int32(1)
		v3 = (*b2BodySimArray)(unsafe.Pointer(v1)).Data + uintptr((*b2BodySimArray)(unsafe.Pointer(v1)).Count-int32FromInt32(1))*100
		goto _13
	_13:
		simDst = v3
		memcpy(tls, simDst, simSrc, uint64(100))
		goto _9
	_9:
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
		contactSrc = (*b2SolverSet)(unsafe.Pointer(set2)).ContactSims.Data + uintptr(i1)*176
		v1 = world + 1144
		v2 = (*b2ContactSim)(unsafe.Pointer(contactSrc)).ContactId
		if !(0 <= v2 && v2 < (*b2ContactArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2ContactArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*68
		goto _18
	_18:
		contact = v3
		if !((*b2Contact)(unsafe.Pointer(contact)).SetIndex == setId2) && b2InternalAssertFcn(tls, __ccgo_ts+14914, __ccgo_ts+14063, int32FromInt32(472)) != 0 {
			__builtin_trap(tls)
		}
		(*b2Contact)(unsafe.Pointer(contact)).SetIndex = setId1
		(*b2Contact)(unsafe.Pointer(contact)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(set1)).ContactSims.Count
		v1 = set1 + 48
		if (*b2ContactSimArray)(unsafe.Pointer(v1)).Count == (*b2ContactSimArray)(unsafe.Pointer(v1)).Capacity {
			if (*b2ContactSimArray)(unsafe.Pointer(v1)).Capacity < int32(2) {
				v2 = int32(2)
			} else {
				v2 = (*b2ContactSimArray)(unsafe.Pointer(v1)).Capacity + (*b2ContactSimArray)(unsafe.Pointer(v1)).Capacity>>int32(1)
			}
			newCapacity1 = v2
			b2ContactSimArray_Reserve(tls, v1, newCapacity1)
		}
		*(*int32)(unsafe.Add(nil, v1+8)) += int32(1)
		v3 = (*b2ContactSimArray)(unsafe.Pointer(v1)).Data + uintptr((*b2ContactSimArray)(unsafe.Pointer(v1)).Count-int32FromInt32(1))*176
		goto _22
	_22:
		contactDst = v3
		memcpy(tls, contactDst, contactSrc, uint64(176))
		goto _14
	_14:
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
		jointSrc = (*b2SolverSet)(unsafe.Pointer(set2)).JointSims.Data + uintptr(i2)*196
		v1 = world + 1104
		v2 = (*b2JointSim)(unsafe.Pointer(jointSrc)).JointId
		if !(0 <= v2 && v2 < (*b2JointArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2JointArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*72
		goto _27
	_27:
		joint = v3
		if !((*b2Joint)(unsafe.Pointer(joint)).SetIndex == setId2) && b2InternalAssertFcn(tls, __ccgo_ts+14942, __ccgo_ts+14063, int32FromInt32(489)) != 0 {
			__builtin_trap(tls)
		}
		(*b2Joint)(unsafe.Pointer(joint)).SetIndex = setId1
		(*b2Joint)(unsafe.Pointer(joint)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(set1)).JointSims.Count
		v1 = set1 + 32
		if (*b2JointSimArray)(unsafe.Pointer(v1)).Count == (*b2JointSimArray)(unsafe.Pointer(v1)).Capacity {
			if (*b2JointSimArray)(unsafe.Pointer(v1)).Capacity < int32(2) {
				v2 = int32(2)
			} else {
				v2 = (*b2JointSimArray)(unsafe.Pointer(v1)).Capacity + (*b2JointSimArray)(unsafe.Pointer(v1)).Capacity>>int32(1)
			}
			newCapacity3 = v2
			b2JointSimArray_Reserve(tls, v1, newCapacity3)
		}
		*(*int32)(unsafe.Add(nil, v1+8)) += int32(1)
		v3 = (*b2JointSimArray)(unsafe.Pointer(v1)).Data + uintptr((*b2JointSimArray)(unsafe.Pointer(v1)).Count-int32FromInt32(1))*196
		goto _31
	_31:
		jointDst = v3
		memcpy(tls, jointDst, jointSrc, uint64(196))
		goto _23
	_23:
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
		v1 = world + 1184
		v2 = islandId
		if !(0 <= v2 && v2 < (*b2IslandArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2IslandArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*56
		goto _36
	_36:
		island = v3
		(*b2Island)(unsafe.Pointer(island)).SetIndex = setId1
		(*b2Island)(unsafe.Pointer(island)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(set1)).IslandSims.Count
		v1 = set1 + 64
		if (*b2IslandSimArray)(unsafe.Pointer(v1)).Count == (*b2IslandSimArray)(unsafe.Pointer(v1)).Capacity {
			if (*b2IslandSimArray)(unsafe.Pointer(v1)).Capacity < int32(2) {
				v2 = int32(2)
			} else {
				v2 = (*b2IslandSimArray)(unsafe.Pointer(v1)).Capacity + (*b2IslandSimArray)(unsafe.Pointer(v1)).Capacity>>int32(1)
			}
			newCapacity2 = v2
			b2IslandSimArray_Reserve(tls, v1, newCapacity2)
		}
		*(*int32)(unsafe.Add(nil, v1+8)) += int32(1)
		v3 = (*b2IslandSimArray)(unsafe.Pointer(v1)).Data + uintptr((*b2IslandSimArray)(unsafe.Pointer(v1)).Count-int32FromInt32(1))*4
		goto _40
	_40:
		islandDst = v3
		memcpy(tls, islandDst, islandSrc, uint64(4))
		goto _32
	_32:
		;
		i3 = i3 + 1
	}
	// destroy the merged set
	b2DestroySolverSet(tls, world, setId2)
	b2ValidateSolverSets(tls, world)
}

func b2ValidateSolverSets(tls *_Stack, world uintptr) {
	_ = uint64FromInt64(4)
}
