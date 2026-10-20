package b2

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

func b2CollideProxyAndSubtree(tls *_Stack, proxy uintptr, nodes uintptr, pair1 int32, context uintptr) {
	var boxv, v1, v11, v14 b2AABBV
	var bv, children, node4, v12 uintptr
	var i, mask, mask1, movedMask, shapeId, stackCount, v3, v4, v5 int32
	var proxyMark, v23 uint32_t
	var stack [512]int32
	var v6, v9 uint8
	var v30 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = boxv, bv, children, i, mask, mask1, movedMask, node4, proxyMark, shapeId, stack, stackCount, v1, v11, v12, v14, v23, v3, v30, v4, v5, v6, v9
	proxyMark = (*b2TreeNode)(unsafe.Pointer(proxy)).FlagIndex & (uint32FromUint32(1) << int32FromInt32(30))
	v1 = *(*AABB)(unsafe.Add(nil, proxy))
	goto _2
_2:
	boxv = v1
	shapeId = *(*int32_t)(unsafe.Add(unsafe.Pointer(proxy), 28))
	stackCount = 0
	v3 = stackCount
	stackCount = stackCount + 1
	stack[v3] = pair1
	for stackCount > 0 {
		stackCount = stackCount - 1
		v3 = stackCount
		pair1 = stack[v3]
		children = nodes + uintptr(pair1)*32
		movedMask = int32(3)
		if proxyMark == uint32(0) {
			v6 = boolUint8((*b2TreeNode)(unsafe.Pointer(children+uintptr(0)*32)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
			goto _7
		_7:
			if v6 != 0 {
				v3 = int32(1)
			} else {
				v3 = 0
			}
			v9 = boolUint8((*b2TreeNode)(unsafe.Pointer(children+uintptr(1)*32)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
			goto _10
		_10:
			if v9 != 0 {
				v4 = int32(2)
			} else {
				v4 = 0
			}
			movedMask = v3 | v4
		}
		v1 = boxv
		v12 = children
		mask = 0
		v11 = v1
		bv = v12 + uintptr(0)*32
		v6 = boolUint8(v11.LowerBound.X <= (*AABB)(unsafe.Pointer(bv)).UpperBound.X && v11.LowerBound.Y <= (*AABB)(unsafe.Pointer(bv)).UpperBound.Y && (*AABB)(unsafe.Pointer(bv)).LowerBound.X <= v11.UpperBound.X && (*AABB)(unsafe.Pointer(bv)).LowerBound.Y <= v11.UpperBound.Y)
		goto _16
	_16:
		if v6 != 0 {
			v3 = int32(1)
		} else {
			v3 = 0
		}
		mask = mask | v3
		v14 = v1
		bv = v12 + uintptr(1)*32
		v9 = boolUint8(v14.LowerBound.X <= (*AABB)(unsafe.Pointer(bv)).UpperBound.X && v14.LowerBound.Y <= (*AABB)(unsafe.Pointer(bv)).UpperBound.Y && (*AABB)(unsafe.Pointer(bv)).LowerBound.X <= v14.UpperBound.X && (*AABB)(unsafe.Pointer(bv)).LowerBound.Y <= v14.UpperBound.Y)
		goto _20
	_20:
		if v9 != 0 {
			v4 = int32(2)
		} else {
			v4 = 0
		}
		mask = mask | v4
		v5 = mask
		goto _22
	_22:
		mask1 = v5 & movedMask
		for mask1 != 0 {
			v23 = uint32FromInt32(__builtin_ctz(tls, uint32FromInt32(mask1)))
			goto _24
		_24:
			i = int32FromUint32(v23)
			mask1 = mask1 & (mask1 - int32(1))
			node4 = children + uintptr(i)*32
			v6 = boolUint8((*b2TreeNode)(unsafe.Pointer(node4)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
			goto _26
		_26:
			if v6 != 0 {
				b2AddCandidatePair(tls, shapeId, *(*int32_t)(unsafe.Add(unsafe.Pointer(node4), 28)), context)
			} else {
				if stackCount < int32(_B2_TREE_STACK_SIZE) {
					v3 = stackCount
					stackCount = stackCount + 1
					v4 = int32FromUint32((*b2TreeNode)(unsafe.Pointer(node4)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
					goto _29
				_29:
					stack[v3] = v4
				} else {
					if v30 = !!(stackCount < int32FromInt32(_B2_TREE_STACK_SIZE)); !v30 {
						b2InternalAssert(tls, __ccgo_ts+4795, __ccgo_ts+4530, int32FromInt32(311))
					}
					_ = v30 || bool(int32FromInt32(0) != 0)
				}
			}
		}
	}
}

// C documentation
//
//	// This collides two sub-trees against each other. They can live in the same dynamic tree.
//	// This can only generate pairs cross sub-tree, but not within a sub-tree. This fact means
//	// this does not generate duplicate pairs.
//	// For example consider the full binary tree A (B (D  E) C (F G))
//	// Colliding children of A (B and C) can give pairs (D,F) (D,G) (E,F) and (E,G).
//	// Then colliding children of B can give the pair (D,E) and for C (F,G).
//	// So no duplicates even when used for self-collision.
//	// Whenever a proxy is moved, the flag is propagated up the hierarchy to the root. So
//	// self collision gathers all those moved internal nodes and collides their subtrees together.
//	// See Real-time collision detection section 6.3.2. This is faster than querying every moved
//	// proxy against the whole tree. Scaling is linear instead of linear * log.
//	// Many other physics engines do this (Bepu, Rapier, etc). So nothing new here.
func b2CollideCrossPairs(tls *_Stack, arrayA uintptr, arrayB uintptr, subtreeA uintptr, subtreeB uintptr, context uintptr) {
	bp := tls.Alloc(4112)
	defer tls.Free(4112)
	var lane, mask, mask1, movedA, movedB, v1, v11, v16, v2, v21, v26, v31, v36, v5, v8 int32
	var pair b2IndexPair
	var pairA1, pairB1, v14, v15, v17, v18, v22, v23, v27, v28, v32, v33 uintptr
	var v12, v19, v24, v29, v3, v34, v6, v9 uint8
	var v38 uint32_t
	var _ /* stack at bp+0 */ [512]b2IndexPair
	var _ /* stackCount at bp+4096 */ int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = lane, mask, mask1, movedA, movedB, pair, pairA1, pairB1, v1, v11, v12, v14, v15, v16, v17, v18, v19, v2, v21, v22, v23, v24, v26, v27, v28, v29, v3, v31, v32, v33, v34, v36, v38, v5, v6, v8, v9
	*(*int32)(unsafe.Add(nil, bp+4096)) = 0
	// Seed the stack.
	if b2TestPair(tls, subtreeA, subtreeB) != 0 {
		b2ExpandPair(tls, arrayA, arrayB, subtreeA, subtreeB, bp, bp+4096, context)
	}
	for *(*int32)(unsafe.Add(nil, bp+4096)) > 0 {
		*(*int32)(unsafe.Add(nil, bp+4096)) = *(*int32)(unsafe.Add(nil, bp+4096)) - 1
		v1 = *(*int32)(unsafe.Add(nil, bp+4096))
		pair = (*(*[512]b2IndexPair)(unsafe.Add(nil, bp)))[v1]
		pairA1 = arrayA + uintptr(pair.A)*32
		pairB1 = arrayB + uintptr(pair.B)*32
		v3 = boolUint8((*b2TreeNode)(unsafe.Pointer(pairA1+uintptr(0)*32)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
		goto _4
	_4:
		if v3 != 0 {
			v2 = int32(0x3)
		} else {
			v2 = 0
		}
		v6 = boolUint8((*b2TreeNode)(unsafe.Pointer(pairA1+uintptr(1)*32)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
		goto _7
	_7:
		if v6 != 0 {
			v5 = int32(0xC)
		} else {
			v5 = 0
		}
		movedA = v2 | v5
		v9 = boolUint8((*b2TreeNode)(unsafe.Pointer(pairB1+uintptr(0)*32)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
		goto _10
	_10:
		if v9 != 0 {
			v8 = int32(0x5)
		} else {
			v8 = 0
		}
		v12 = boolUint8((*b2TreeNode)(unsafe.Pointer(pairB1+uintptr(1)*32)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
		goto _13
	_13:
		if v12 != 0 {
			v11 = int32(0xA)
		} else {
			v11 = 0
		}
		movedB = v8 | v11
		v14 = pairA1
		v15 = pairB1
		mask = 0
		v17 = v14
		v18 = v15
		v19 = boolUint8((*AABB)(unsafe.Pointer(v17)).LowerBound.X <= (*AABB)(unsafe.Pointer(v18)).UpperBound.X && (*AABB)(unsafe.Pointer(v17)).LowerBound.Y <= (*AABB)(unsafe.Pointer(v18)).UpperBound.Y && (*AABB)(unsafe.Pointer(v18)).LowerBound.X <= (*AABB)(unsafe.Pointer(v17)).UpperBound.X && (*AABB)(unsafe.Pointer(v18)).LowerBound.Y <= (*AABB)(unsafe.Pointer(v17)).UpperBound.Y)
		goto _20
	_20:
		if v19 != 0 {
			v16 = int32(1)
		} else {
			v16 = 0
		}
		mask = mask | v16
		v22 = v14
		v23 = v15 + 1*32
		v24 = boolUint8((*AABB)(unsafe.Pointer(v22)).LowerBound.X <= (*AABB)(unsafe.Pointer(v23)).UpperBound.X && (*AABB)(unsafe.Pointer(v22)).LowerBound.Y <= (*AABB)(unsafe.Pointer(v23)).UpperBound.Y && (*AABB)(unsafe.Pointer(v23)).LowerBound.X <= (*AABB)(unsafe.Pointer(v22)).UpperBound.X && (*AABB)(unsafe.Pointer(v23)).LowerBound.Y <= (*AABB)(unsafe.Pointer(v22)).UpperBound.Y)
		goto _25
	_25:
		if v24 != 0 {
			v21 = int32(2)
		} else {
			v21 = 0
		}
		mask = mask | v21
		v27 = v14 + 1*32
		v28 = v15
		v29 = boolUint8((*AABB)(unsafe.Pointer(v27)).LowerBound.X <= (*AABB)(unsafe.Pointer(v28)).UpperBound.X && (*AABB)(unsafe.Pointer(v27)).LowerBound.Y <= (*AABB)(unsafe.Pointer(v28)).UpperBound.Y && (*AABB)(unsafe.Pointer(v28)).LowerBound.X <= (*AABB)(unsafe.Pointer(v27)).UpperBound.X && (*AABB)(unsafe.Pointer(v28)).LowerBound.Y <= (*AABB)(unsafe.Pointer(v27)).UpperBound.Y)
		goto _30
	_30:
		if v29 != 0 {
			v26 = int32(4)
		} else {
			v26 = 0
		}
		mask = mask | v26
		v32 = v14 + 1*32
		v33 = v15 + 1*32
		v34 = boolUint8((*AABB)(unsafe.Pointer(v32)).LowerBound.X <= (*AABB)(unsafe.Pointer(v33)).UpperBound.X && (*AABB)(unsafe.Pointer(v32)).LowerBound.Y <= (*AABB)(unsafe.Pointer(v33)).UpperBound.Y && (*AABB)(unsafe.Pointer(v33)).LowerBound.X <= (*AABB)(unsafe.Pointer(v32)).UpperBound.X && (*AABB)(unsafe.Pointer(v33)).LowerBound.Y <= (*AABB)(unsafe.Pointer(v32)).UpperBound.Y)
		goto _35
	_35:
		if v34 != 0 {
			v31 = int32(8)
		} else {
			v31 = 0
		}
		mask = mask | v31
		v36 = mask
		goto _37
	_37:
		mask1 = v36 & (movedA | movedB)
		for mask1 != 0 {
			v38 = uint32FromInt32(__builtin_ctz(tls, uint32FromInt32(mask1)))
			goto _39
		_39:
			lane = int32FromUint32(v38)
			mask1 = mask1 & (mask1 - int32(1))
			b2ExpandPair(tls, arrayA, arrayB, pairA1+uintptr(lane>>int32FromInt32(1))*32, pairB1+uintptr(lane&int32FromInt32(1))*32, bp, bp+4096, context)
		}
	}
}

func b2CollideMoverAndCircle(tls *_Stack, mover uintptr, shape uintptr) (r PlaneResult) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var distanceOutput DistanceOutput
	var plane Plane
	var totalRadius float32
	var _ /* cache at bp+168 */ SimplexCache
	var _ /* distanceInput at bp+0 */ DistanceInput
	_, _, _ = distanceOutput, plane, totalRadius
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyA = b2MakeProxy(tls, shape, int32(1), float32FromFloat32(0))
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyB = b2MakeProxy(tls, mover, int32(2), (*Capsule)(unsafe.Pointer(mover)).Radius)
	(*(*DistanceInput)(unsafe.Add(nil, bp))).Transform = b2Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp))).UseRadii = boolUint8(false1 != 0)
	totalRadius = (*Capsule)(unsafe.Pointer(mover)).Radius + (*Circle)(unsafe.Pointer(shape)).Radius
	*(*SimplexCache)(unsafe.Add(nil, bp+168)) = SimplexCache{}
	distanceOutput = b2ComputeShapeDistance(tls, bp, bp+168, uintptrFromInt32(0), 0)
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

func b2CollideMoverAndCapsule(tls *_Stack, mover uintptr, shape uintptr) (r PlaneResult) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var distanceOutput DistanceOutput
	var plane Plane
	var totalRadius float32
	var _ /* cache at bp+168 */ SimplexCache
	var _ /* distanceInput at bp+0 */ DistanceInput
	_, _, _ = distanceOutput, plane, totalRadius
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyA = b2MakeProxy(tls, shape, int32(2), float32FromFloat32(0))
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyB = b2MakeProxy(tls, mover, int32(2), (*Capsule)(unsafe.Pointer(mover)).Radius)
	(*(*DistanceInput)(unsafe.Add(nil, bp))).Transform = b2Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp))).UseRadii = boolUint8(false1 != 0)
	totalRadius = (*Capsule)(unsafe.Pointer(mover)).Radius + (*Capsule)(unsafe.Pointer(shape)).Radius
	*(*SimplexCache)(unsafe.Add(nil, bp+168)) = SimplexCache{}
	distanceOutput = b2ComputeShapeDistance(tls, bp, bp+168, uintptrFromInt32(0), 0)
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

func b2CollideMoverAndPolygon(tls *_Stack, mover uintptr, shape uintptr) (r PlaneResult) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var distanceOutput DistanceOutput
	var plane Plane
	var totalRadius float32
	var _ /* cache at bp+168 */ SimplexCache
	var _ /* distanceInput at bp+0 */ DistanceInput
	_, _, _ = distanceOutput, plane, totalRadius
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyA = b2MakeProxy(tls, shape, (*Polygon)(unsafe.Pointer(shape)).Count, (*Polygon)(unsafe.Pointer(shape)).Radius)
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyB = b2MakeProxy(tls, mover, int32(2), (*Capsule)(unsafe.Pointer(mover)).Radius)
	(*(*DistanceInput)(unsafe.Add(nil, bp))).Transform = b2Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp))).UseRadii = boolUint8(false1 != 0)
	totalRadius = (*Capsule)(unsafe.Pointer(mover)).Radius + (*Polygon)(unsafe.Pointer(shape)).Radius
	*(*SimplexCache)(unsafe.Add(nil, bp+168)) = SimplexCache{}
	distanceOutput = b2ComputeShapeDistance(tls, bp, bp+168, uintptrFromInt32(0), 0)
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

func b2CollideMoverAndSegment(tls *_Stack, mover uintptr, shape uintptr) (r PlaneResult) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var distanceOutput DistanceOutput
	var plane Plane
	var totalRadius float32
	var _ /* cache at bp+168 */ SimplexCache
	var _ /* distanceInput at bp+0 */ DistanceInput
	_, _, _ = distanceOutput, plane, totalRadius
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyA = b2MakeProxy(tls, shape, int32(2), float32FromFloat32(0))
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyB = b2MakeProxy(tls, mover, int32(2), (*Capsule)(unsafe.Pointer(mover)).Radius)
	(*(*DistanceInput)(unsafe.Add(nil, bp))).Transform = b2Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp))).UseRadii = boolUint8(false1 != 0)
	totalRadius = (*Capsule)(unsafe.Pointer(mover)).Radius
	*(*SimplexCache)(unsafe.Add(nil, bp+168)) = SimplexCache{}
	distanceOutput = b2ComputeShapeDistance(tls, bp, bp+168, uintptrFromInt32(0), 0)
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

func b2CollideCircles(tls *_Stack, circleA uintptr, circleB uintptr, xf Transform) (r LocalManifold) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var cA, cB, normal, pointA, pointB, u, v10, v11, v2, v3, v5, v6, v7 Vec2
	var lengthSquared, radiusA, radiusB, s1, separation, x, y, v14, v19 float32
	var mp, v9 uintptr
	var v1 Transform
	var _ /* distance at bp+44 */ float32
	var _ /* manifold at bp+0 */ LocalManifold
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = cA, cB, lengthSquared, mp, normal, pointA, pointB, radiusA, radiusB, s1, separation, u, x, y, v1, v10, v11, v14, v19, v2, v3, v5, v6, v7, v9
	*(*LocalManifold)(unsafe.Add(nil, bp)) = LocalManifold{}
	pointA = (*Circle)(unsafe.Pointer(circleA)).Center
	v1 = xf
	v2 = (*Circle)(unsafe.Pointer(circleB)).Center
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	pointB = v3
	v5 = pointB
	v6 = pointA
	v7 = Vec2{
		X: v5.X - v6.X,
		Y: v5.Y - v6.Y,
	}
	goto _8
_8:
	v9 = bp + 44
	v10 = v7
	lengthSquared = float32(v10.X*v10.X) + float32(v10.Y*v10.Y)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		*(*float32)(unsafe.Add(nil, v9)) = sqrtf(tls, lengthSquared)
		s1 = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v9))
		u = Vec2{
			X: float32(s1 * v10.X),
			Y: float32(s1 * v10.Y),
		}
		v11 = u
		goto _12
	}
	*(*float32)(unsafe.Add(nil, v9)) = float32FromFloat32(0)
	v11 = Vec2{}
	goto _12
_12:
	normal = v11
	radiusA = (*Circle)(unsafe.Pointer(circleA)).Radius
	radiusB = (*Circle)(unsafe.Pointer(circleB)).Radius
	separation = *(*float32)(unsafe.Add(nil, bp+44)) - radiusA - radiusB
	if separation > float32(float32FromFloat32(4)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
		return *(*LocalManifold)(unsafe.Add(nil, bp))
	}
	v2 = pointA
	v14 = radiusA
	v3 = normal
	v5 = Vec2{
		X: v2.X + float32(v14*v3.X),
		Y: v2.Y + float32(v14*v3.Y),
	}
	goto _17
_17:
	cA = v5
	v6 = pointB
	v19 = -radiusB
	v7 = normal
	v10 = Vec2{
		X: v6.X + float32(v19*v7.X),
		Y: v6.Y + float32(v19*v7.Y),
	}
	goto _22
_22:
	cB = v10
	(*(*LocalManifold)(unsafe.Add(nil, bp))).Normal = normal
	mp = bp + 8 + uintptr(0)*16
	v2 = cA
	v3 = cB
	v14 = float32FromFloat32(0.5)
	v5 = Vec2{
		X: float32((float32FromFloat32(1)-v14)*v2.X) + float32(v14*v3.X),
		Y: float32((float32FromFloat32(1)-v14)*v2.Y) + float32(v14*v3.Y),
	}
	goto _27
_27:
	(*LocalManifoldPoint)(unsafe.Pointer(mp)).Point = v5
	(*LocalManifoldPoint)(unsafe.Pointer(mp)).Separation = separation
	(*LocalManifoldPoint)(unsafe.Pointer(mp)).Id = uint16(0)
	(*(*LocalManifold)(unsafe.Add(nil, bp))).PointCount = int32(1)
	return *(*LocalManifold)(unsafe.Add(nil, bp))
}

// C documentation
//
//	/// Compute the collision manifold between a capsule and circle
func b2CollideCapsuleAndCircle(tls *_Stack, capsuleA uintptr, circleB uintptr, xf Transform) (r LocalManifold) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var cA, cB, e, normal, p1, p2, pA, pB, u, v10, v11, v13, v14, v17, v18, v19, v2, v21, v22, v3, v5, v6, v7, v9 Vec2
	var lengthSquared, radiusA, radiusB, s1, s11, s2, s21, separation, x, y, v15, v23 float32
	var mp, v38 uintptr
	var v1 Transform
	var _ /* distance at bp+44 */ float32
	var _ /* manifold at bp+0 */ LocalManifold
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = cA, cB, e, lengthSquared, mp, normal, p1, p2, pA, pB, radiusA, radiusB, s1, s11, s2, s21, separation, u, x, y, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v21, v22, v23, v3, v38, v5, v6, v7, v9
	*(*LocalManifold)(unsafe.Add(nil, bp)) = LocalManifold{}
	v1 = xf
	v2 = (*Circle)(unsafe.Pointer(circleB)).Center
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	// Compute circle position in the frame of the capsule.
	pB = v3
	// Compute closest point
	p1 = (*Capsule)(unsafe.Pointer(capsuleA)).Center1
	p2 = (*Capsule)(unsafe.Pointer(capsuleA)).Center2
	v5 = p2
	v6 = p1
	v7 = Vec2{
		X: v5.X - v6.X,
		Y: v5.Y - v6.Y,
	}
	goto _8
_8:
	e = v7
	v9 = pB
	v10 = p1
	v11 = Vec2{
		X: v9.X - v10.X,
		Y: v9.Y - v10.Y,
	}
	goto _12
_12:
	v13 = v11
	v14 = e
	v15 = float32(v13.X*v14.X) + float32(v13.Y*v14.Y)
	goto _16
_16:
	s11 = v15
	v17 = p2
	v18 = pB
	v19 = Vec2{
		X: v17.X - v18.X,
		Y: v17.Y - v18.Y,
	}
	goto _20
_20:
	v21 = v19
	v22 = e
	v23 = float32(v21.X*v22.X) + float32(v21.Y*v22.Y)
	goto _24
_24:
	s21 = v23
	if s11 < float32FromFloat32(0) {
		// p1 region
		pA = p1
	} else {
		if s21 < float32FromFloat32(0) {
			// p2 region
			pA = p2
		} else {
			v2 = e
			v3 = e
			v15 = float32(v2.X*v3.X) + float32(v2.Y*v3.Y)
			goto _28
		_28:
			// circle colliding with segment interior
			s2 = s11 / v15
			v5 = p1
			v23 = s2
			v6 = e
			v7 = Vec2{
				X: v5.X + float32(v23*v6.X),
				Y: v5.Y + float32(v23*v6.Y),
			}
			goto _33
		_33:
			pA = v7
		}
	}
	v2 = pB
	v3 = pA
	v5 = Vec2{
		X: v2.X - v3.X,
		Y: v2.Y - v3.Y,
	}
	goto _37
_37:
	v38 = bp + 44
	v6 = v5
	lengthSquared = float32(v6.X*v6.X) + float32(v6.Y*v6.Y)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		*(*float32)(unsafe.Add(nil, v38)) = sqrtf(tls, lengthSquared)
		s1 = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v38))
		u = Vec2{
			X: float32(s1 * v6.X),
			Y: float32(s1 * v6.Y),
		}
		v7 = u
		goto _41
	}
	*(*float32)(unsafe.Add(nil, v38)) = float32FromFloat32(0)
	v7 = Vec2{}
	goto _41
_41:
	normal = v7
	radiusA = (*Capsule)(unsafe.Pointer(capsuleA)).Radius
	radiusB = (*Circle)(unsafe.Pointer(circleB)).Radius
	separation = *(*float32)(unsafe.Add(nil, bp+44)) - radiusA - radiusB
	if separation > float32(float32FromFloat32(4)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
		return *(*LocalManifold)(unsafe.Add(nil, bp))
	}
	v2 = pA
	v15 = radiusA
	v3 = normal
	v5 = Vec2{
		X: v2.X + float32(v15*v3.X),
		Y: v2.Y + float32(v15*v3.Y),
	}
	goto _46
_46:
	cA = v5
	v6 = pB
	v23 = -radiusB
	v7 = normal
	v9 = Vec2{
		X: v6.X + float32(v23*v7.X),
		Y: v6.Y + float32(v23*v7.Y),
	}
	goto _51
_51:
	cB = v9
	(*(*LocalManifold)(unsafe.Add(nil, bp))).Normal = normal
	mp = bp + 8 + uintptr(0)*16
	v2 = cA
	v3 = cB
	v15 = float32FromFloat32(0.5)
	v5 = Vec2{
		X: float32((float32FromFloat32(1)-v15)*v2.X) + float32(v15*v3.X),
		Y: float32((float32FromFloat32(1)-v15)*v2.Y) + float32(v15*v3.Y),
	}
	goto _56
_56:
	(*LocalManifoldPoint)(unsafe.Pointer(mp)).Point = v5
	(*LocalManifoldPoint)(unsafe.Pointer(mp)).Separation = separation
	(*LocalManifoldPoint)(unsafe.Pointer(mp)).Id = uint16(0)
	(*(*LocalManifold)(unsafe.Add(nil, bp))).PointCount = int32(1)
	return *(*LocalManifold)(unsafe.Add(nil, bp))
}

func b2CollidePolygonAndCircle(tls *_Stack, polygonA uintptr, circleB uintptr, xf Transform) (r LocalManifold) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var cA, cA1, cA2, cB, cB1, cB2, center, contactPointA, contactPointA1, normal, normal1, normal2, u, v1, v2, v10, v111, v15, v16, v17, v19, v21, v20, v211, v23, v24, v27, v28, v29, v3, v31, v32, v33, v35, v36, v39, v40, v41, v6, v7, v8 Vec2
	var i, normalIndex, vertIndex1, vertIndex2, vertexCount, v14 int32
	var lengthSquared, radius, radiusA, radiusB, s2, s3, separation, speculativeDistance, u1, u2, x, y, v12, v25, v37, v52 float32
	var mp, mp1, mp2, normals, vertices uintptr
	var v11 Transform
	var _ /* manifold at bp+0 */ LocalManifold
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = cA, cA1, cA2, cB, cB1, cB2, center, contactPointA, contactPointA1, i, lengthSquared, mp, mp1, mp2, normal, normal1, normal2, normalIndex, normals, radius, radiusA, radiusB, s2, s3, separation, speculativeDistance, u, u1, u2, v1, v2, vertIndex1, vertIndex2, vertexCount, vertices, x, y, v11, v10, v111, v12, v14, v15, v16, v17, v19, v21, v20, v211, v23, v24, v25, v27, v28, v29, v3, v31, v32, v33, v35, v36, v37, v39, v40, v41, v52, v6, v7, v8
	*(*LocalManifold)(unsafe.Add(nil, bp)) = LocalManifold{}
	speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter))
	v11 = xf
	v21 = (*Circle)(unsafe.Pointer(circleB)).Center
	x = float32(v11.Q.C*v21.X) - float32(v11.Q.S*v21.Y) + v11.P.X
	y = float32(v11.Q.S*v21.X) + float32(v11.Q.C*v21.Y) + v11.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	// Compute circle position in the frame of the polygon.
	center = v3
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
		v6 = center
		v7 = *(*Vec2)(unsafe.Add(nil, vertices+uintptr(i)*8))
		v8 = Vec2{
			X: v6.X - v7.X,
			Y: v6.Y - v7.Y,
		}
		goto _9
	_9:
		v10 = *(*Vec2)(unsafe.Add(nil, normals+uintptr(i)*8))
		v111 = v8
		v12 = float32(v10.X*v111.X) + float32(v10.Y*v111.Y)
		goto _13
	_13:
		s3 = v12
		if s3 > separation {
			separation = s3
			normalIndex = i
		}
		goto _5
	_5:
		;
		i = i + 1
	}
	if separation > radius+speculativeDistance {
		return *(*LocalManifold)(unsafe.Add(nil, bp))
	}
	// Vertices of the reference edge.
	vertIndex1 = normalIndex
	if vertIndex1+int32(1) < vertexCount {
		v14 = vertIndex1 + int32(1)
	} else {
		v14 = 0
	}
	vertIndex2 = v14
	v1 = *(*Vec2)(unsafe.Add(nil, vertices+uintptr(vertIndex1)*8))
	v2 = *(*Vec2)(unsafe.Add(nil, vertices+uintptr(vertIndex2)*8))
	v21 = center
	v3 = v1
	v6 = Vec2{
		X: v21.X - v3.X,
		Y: v21.Y - v3.Y,
	}
	goto _18
_18:
	v7 = v2
	v8 = v1
	v10 = Vec2{
		X: v7.X - v8.X,
		Y: v7.Y - v8.Y,
	}
	goto _22
_22:
	v111 = v6
	v15 = v10
	v12 = float32(v111.X*v15.X) + float32(v111.Y*v15.Y)
	goto _26
_26:
	// Compute barycentric coordinates
	u1 = v12
	v16 = center
	v17 = v2
	v19 = Vec2{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
	}
	goto _30
_30:
	v20 = v1
	v211 = v2
	v23 = Vec2{
		X: v20.X - v211.X,
		Y: v20.Y - v211.Y,
	}
	goto _34
_34:
	v24 = v19
	v27 = v23
	v25 = float32(v24.X*v27.X) + float32(v24.Y*v27.Y)
	goto _38
_38:
	u2 = v25
	if u1 < float32FromFloat32(0) && separation > float32FromFloat32(1.1920928955078125e-07) {
		v28 = center
		v29 = v1
		v31 = Vec2{
			X: v28.X - v29.X,
			Y: v28.Y - v29.Y,
		}
		goto _42
	_42:
		v32 = v31
		lengthSquared = float32(v32.X*v32.X) + float32(v32.Y*v32.Y)
		if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			s2 = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
			u = Vec2{
				X: float32(s2 * v32.X),
				Y: float32(s2 * v32.Y),
			}
			v33 = u
			goto _45
		}
		v33 = Vec2{}
		goto _45
	_45:
		// Circle center is closest to v1 and safely outside the polygon
		normal = v33
		v35 = center
		v36 = v1
		v39 = Vec2{
			X: v35.X - v36.X,
			Y: v35.Y - v36.Y,
		}
		goto _49
	_49:
		v40 = v39
		v41 = normal
		v37 = float32(v40.X*v41.X) + float32(v40.Y*v41.Y)
		goto _53
	_53:
		separation = v37
		if separation > radius+speculativeDistance {
			return *(*LocalManifold)(unsafe.Add(nil, bp))
		}
		v21 = v1
		v12 = radiusA
		v3 = normal
		v6 = Vec2{
			X: v21.X + float32(v12*v3.X),
			Y: v21.Y + float32(v12*v3.Y),
		}
		goto _58
	_58:
		cA = v6
		v7 = center
		v25 = radiusB
		v8 = normal
		v10 = Vec2{
			X: v7.X - float32(v25*v8.X),
			Y: v7.Y - float32(v25*v8.Y),
		}
		goto _63
	_63:
		cB = v10
		v111 = cA
		v15 = cB
		v37 = float32FromFloat32(0.5)
		v16 = Vec2{
			X: float32((float32FromFloat32(1)-v37)*v111.X) + float32(v37*v15.X),
			Y: float32((float32FromFloat32(1)-v37)*v111.Y) + float32(v37*v15.Y),
		}
		goto _68
	_68:
		contactPointA = v16
		(*(*LocalManifold)(unsafe.Add(nil, bp))).Normal = normal
		mp = bp + 8 + uintptr(0)*16
		(*LocalManifoldPoint)(unsafe.Pointer(mp)).Point = contactPointA
		v21 = cB
		v3 = cA
		v6 = Vec2{
			X: v21.X - v3.X,
			Y: v21.Y - v3.Y,
		}
		goto _72
	_72:
		v7 = v6
		v8 = normal
		v12 = float32(v7.X*v8.X) + float32(v7.Y*v8.Y)
		goto _76
	_76:
		(*LocalManifoldPoint)(unsafe.Pointer(mp)).Separation = v12
		(*LocalManifoldPoint)(unsafe.Pointer(mp)).Id = uint16(0)
		(*(*LocalManifold)(unsafe.Add(nil, bp))).PointCount = int32(1)
	} else {
		if u2 < float32FromFloat32(0) && separation > float32FromFloat32(1.1920928955078125e-07) {
			v21 = center
			v3 = v2
			v6 = Vec2{
				X: v21.X - v3.X,
				Y: v21.Y - v3.Y,
			}
			goto _80
		_80:
			v7 = v6
			lengthSquared = float32(v7.X*v7.X) + float32(v7.Y*v7.Y)
			if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
				s2 = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
				u = Vec2{
					X: float32(s2 * v7.X),
					Y: float32(s2 * v7.Y),
				}
				v8 = u
				goto _83
			}
			v8 = Vec2{}
			goto _83
		_83:
			// Circle center is closest to v2 and safely outside the polygon
			normal1 = v8
			v10 = center
			v111 = v2
			v15 = Vec2{
				X: v10.X - v111.X,
				Y: v10.Y - v111.Y,
			}
			goto _87
		_87:
			v16 = v15
			v17 = normal1
			v12 = float32(v16.X*v17.X) + float32(v16.Y*v17.Y)
			goto _91
		_91:
			separation = v12
			if separation > radius+speculativeDistance {
				return *(*LocalManifold)(unsafe.Add(nil, bp))
			}
			v21 = v2
			v12 = radiusA
			v3 = normal1
			v6 = Vec2{
				X: v21.X + float32(v12*v3.X),
				Y: v21.Y + float32(v12*v3.Y),
			}
			goto _96
		_96:
			cA1 = v6
			v7 = center
			v25 = radiusB
			v8 = normal1
			v10 = Vec2{
				X: v7.X - float32(v25*v8.X),
				Y: v7.Y - float32(v25*v8.Y),
			}
			goto _101
		_101:
			cB1 = v10
			v111 = cA1
			v15 = cB1
			v37 = float32FromFloat32(0.5)
			v16 = Vec2{
				X: float32((float32FromFloat32(1)-v37)*v111.X) + float32(v37*v15.X),
				Y: float32((float32FromFloat32(1)-v37)*v111.Y) + float32(v37*v15.Y),
			}
			goto _106
		_106:
			contactPointA1 = v16
			(*(*LocalManifold)(unsafe.Add(nil, bp))).Normal = normal1
			mp1 = bp + 8 + uintptr(0)*16
			(*LocalManifoldPoint)(unsafe.Pointer(mp1)).Point = contactPointA1
			v21 = cB1
			v3 = cA1
			v6 = Vec2{
				X: v21.X - v3.X,
				Y: v21.Y - v3.Y,
			}
			goto _110
		_110:
			v7 = v6
			v8 = normal1
			v12 = float32(v7.X*v8.X) + float32(v7.Y*v8.Y)
			goto _114
		_114:
			(*LocalManifoldPoint)(unsafe.Pointer(mp1)).Separation = v12
			(*LocalManifoldPoint)(unsafe.Pointer(mp1)).Id = uint16(0)
			(*(*LocalManifold)(unsafe.Add(nil, bp))).PointCount = int32(1)
		} else {
			// Circle center is between v1 and v2. Center may be inside polygon
			normal2 = *(*Vec2)(unsafe.Add(nil, normals+uintptr(normalIndex)*8))
			(*(*LocalManifold)(unsafe.Add(nil, bp))).Normal = normal2
			v21 = center
			v3 = v1
			v6 = Vec2{
				X: v21.X - v3.X,
				Y: v21.Y - v3.Y,
			}
			goto _118
		_118:
			v7 = v6
			v8 = normal2
			v12 = float32(v7.X*v8.X) + float32(v7.Y*v8.Y)
			goto _122
		_122:
			v10 = center
			v25 = radiusA - v12
			v111 = normal2
			v15 = Vec2{
				X: v10.X + float32(v25*v111.X),
				Y: v10.Y + float32(v25*v111.Y),
			}
			goto _127
		_127:
			// cA is the projection of the circle center onto to the reference edge
			cA2 = v15
			v16 = center
			v37 = radiusB
			v17 = normal2
			v19 = Vec2{
				X: v16.X - float32(v37*v17.X),
				Y: v16.Y - float32(v37*v17.Y),
			}
			goto _132
		_132:
			// cB is the deepest point on the circle with respect to the reference edge
			cB2 = v19
			// The contact point is the midpoint
			mp2 = bp + 8 + uintptr(0)*16
			v20 = cA2
			v211 = cB2
			v52 = float32FromFloat32(0.5)
			v23 = Vec2{
				X: float32((float32FromFloat32(1)-v52)*v20.X) + float32(v52*v211.X),
				Y: float32((float32FromFloat32(1)-v52)*v20.Y) + float32(v52*v211.Y),
			}
			goto _137
		_137:
			(*LocalManifoldPoint)(unsafe.Pointer(mp2)).Point = v23
			(*LocalManifoldPoint)(unsafe.Pointer(mp2)).Separation = separation - radius
			(*LocalManifoldPoint)(unsafe.Pointer(mp2)).Id = uint16(0)
			(*(*LocalManifold)(unsafe.Add(nil, bp))).PointCount = int32(1)
		}
	}
	return *(*LocalManifold)(unsafe.Add(nil, bp))
}

// C documentation
//
//	// Follows Ericson 5.1.9 Closest Points of Two Line Segments
//	// Adds some logic to support clipping to get two contact points
func b2CollideCapsules(tls *_Stack, capsuleA uintptr, capsuleB uintptr, xf Transform) (r1 LocalManifold) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var c, c1, c2, closest1, closest2, cp, cp1, cq, cq1, d1, d2, normal, normalA, normalB, origin, p1, p2, q1, q2, r, u, u1, u11, u2, v1, v10, v11, v14, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v29, v3, v30, v34, v35, v36, v38, v39, v42, v5, v6, v7 Vec2
	var d12, dd1, dd2, denom, distance, distanceSquared, epsSqr, f1, f2, fp1, fp2, fq1, fq2, lengthSquared, lengthSquared1, maxDistance, radius, radiusA, radiusB, rd1, rd2, s1, s1n, s1n1, s1p, s1p1, s2, separationA, separationB, sp, sp1, sq, sq1, ss1, ss11, ss2, ss21, x, y, v27, v31, v40, v44, v48, v50, v51, v52, v53 float32
	var i, i1, i2, v292, v293 int32
	var mp, mp1, v85, v89 uintptr
	var outsideA, outsideB uint8
	var xfs, v13, v9 Transform
	var v33 bool
	var _ /* length1 at bp+44 */ float32
	var _ /* length2 at bp+48 */ float32
	var _ /* manifold at bp+0 */ LocalManifold
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, c1, c2, closest1, closest2, cp, cp1, cq, cq1, d1, d12, d2, dd1, dd2, denom, distance, distanceSquared, epsSqr, f1, f2, fp1, fp2, fq1, fq2, i, i1, i2, lengthSquared, lengthSquared1, maxDistance, mp, mp1, normal, normalA, normalB, origin, outsideA, outsideB, p1, p2, q1, q2, r, radius, radiusA, radiusB, rd1, rd2, s1, s1n, s1n1, s1p, s1p1, s2, separationA, separationB, sp, sp1, sq, sq1, ss1, ss11, ss2, ss21, u, u1, u11, u2, x, xfs, y, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v29, v292, v293, v3, v30, v31, v33, v34, v35, v36, v38, v39, v40, v42, v44, v48, v5, v50, v51, v52, v53, v6, v7, v85, v89, v9
	origin = (*Capsule)(unsafe.Pointer(capsuleA)).Center1
	v1 = xf.P
	v2 = origin
	v3 = Vec2{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
	goto _4
_4:
	// Shift to the origin in frame A for round-off, a pure translation in A's frame
	xfs = Transform{
		P: v3,
		Q: xf.Q,
	}
	p1 = b2Vec2_zero
	v5 = (*Capsule)(unsafe.Pointer(capsuleA)).Center2
	v6 = origin
	v7 = Vec2{
		X: v5.X - v6.X,
		Y: v5.Y - v6.Y,
	}
	goto _8
_8:
	q1 = v7
	v9 = xfs
	v10 = (*Capsule)(unsafe.Pointer(capsuleB)).Center1
	x = float32(v9.Q.C*v10.X) - float32(v9.Q.S*v10.Y) + v9.P.X
	y = float32(v9.Q.S*v10.X) + float32(v9.Q.C*v10.Y) + v9.P.Y
	v11 = Vec2{
		X: x,
		Y: y,
	}
	goto _12
_12:
	p2 = v11
	v13 = xfs
	v14 = (*Capsule)(unsafe.Pointer(capsuleB)).Center2
	x = float32(v13.Q.C*v14.X) - float32(v13.Q.S*v14.Y) + v13.P.X
	y = float32(v13.Q.S*v14.X) + float32(v13.Q.C*v14.Y) + v13.P.Y
	v15 = Vec2{
		X: x,
		Y: y,
	}
	goto _16
_16:
	q2 = v15
	v17 = q1
	v18 = p1
	v19 = Vec2{
		X: v17.X - v18.X,
		Y: v17.Y - v18.Y,
	}
	goto _20
_20:
	d1 = v19
	v21 = q2
	v22 = p2
	v23 = Vec2{
		X: v21.X - v22.X,
		Y: v21.Y - v22.Y,
	}
	goto _24
_24:
	d2 = v23
	v25 = d1
	v26 = d1
	v27 = float32(v25.X*v26.X) + float32(v25.Y*v26.Y)
	goto _28
_28:
	dd1 = v27
	v29 = d2
	v30 = d2
	v31 = float32(v29.X*v30.X) + float32(v29.Y*v30.Y)
	goto _32
_32:
	dd2 = v31
	epsSqr = float32(float32FromFloat32(1.1920928955078125e-07) * float32FromFloat32(1.1920928955078125e-07))
	if v33 = !!(dd1 > epsSqr && dd2 > epsSqr); !v33 {
		b2InternalAssert(tls, __ccgo_ts+15862, __ccgo_ts+15836, int32FromInt32(261))
	}
	_ = v33 || bool(int32FromInt32(0) != 0)
	v1 = p1
	v2 = p2
	v3 = Vec2{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
	goto _37
_37:
	r = v3
	v5 = r
	v6 = d1
	v27 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y)
	goto _41
_41:
	rd1 = v27
	v7 = r
	v10 = d2
	v31 = float32(v7.X*v10.X) + float32(v7.Y*v10.Y)
	goto _45
_45:
	rd2 = v31
	v11 = d1
	v14 = d2
	v40 = float32(v11.X*v14.X) + float32(v11.Y*v14.Y)
	goto _49
_49:
	d12 = v40
	denom = float32(dd1*dd2) - float32(d12*d12)
	// Fraction on segment 1
	f1 = float32FromFloat32(0)
	if denom != float32FromFloat32(0) {
		// not parallel
		v44 = (float32(d12*rd2) - float32(rd1*dd2)) / denom
		v48 = float32FromFloat32(0)
		v50 = float32FromFloat32(1)
		if v44 < v48 {
			v52 = v48
		} else {
			if v44 > v50 {
				v53 = v50
			} else {
				v53 = v44
			}
			v52 = v53
		}
		v51 = v52
		goto _54
	_54:
		f1 = v51
	}
	// Compute point on segment 2 closest to p1 + f1 * d1
	f2 = (float32(d12*f1) + rd2) / dd2
	// Clamping of segment 2 requires a do over on segment 1
	if f2 < float32FromFloat32(0) {
		f2 = float32FromFloat32(0)
		v27 = -rd1 / dd1
		v31 = float32FromFloat32(0)
		v40 = float32FromFloat32(1)
		if v27 < v31 {
			v48 = v31
		} else {
			if v27 > v40 {
				v50 = v40
			} else {
				v50 = v27
			}
			v48 = v50
		}
		v44 = v48
		goto _61
	_61:
		f1 = v44
	} else {
		if f2 > float32FromFloat32(1) {
			f2 = float32FromFloat32(1)
			v27 = (d12 - rd1) / dd1
			v31 = float32FromFloat32(0)
			v40 = float32FromFloat32(1)
			if v27 < v31 {
				v48 = v31
			} else {
				if v27 > v40 {
					v50 = v40
				} else {
					v50 = v27
				}
				v48 = v50
			}
			v44 = v48
			goto _68
		_68:
			f1 = v44
		}
	}
	v1 = p1
	v27 = f1
	v2 = d1
	v3 = Vec2{
		X: v1.X + float32(v27*v2.X),
		Y: v1.Y + float32(v27*v2.Y),
	}
	goto _75
_75:
	closest1 = v3
	v5 = p2
	v31 = f2
	v6 = d2
	v7 = Vec2{
		X: v5.X + float32(v31*v6.X),
		Y: v5.Y + float32(v31*v6.Y),
	}
	goto _80
_80:
	closest2 = v7
	v10 = closest1
	v11 = closest2
	c = Vec2{
		X: v11.X - v10.X,
		Y: v11.Y - v10.Y,
	}
	v40 = float32(c.X*c.X) + float32(c.Y*c.Y)
	goto _84
_84:
	distanceSquared = v40
	*(*LocalManifold)(unsafe.Add(nil, bp)) = LocalManifold{}
	radiusA = (*Capsule)(unsafe.Pointer(capsuleA)).Radius
	radiusB = (*Capsule)(unsafe.Pointer(capsuleB)).Radius
	radius = radiusA + radiusB
	maxDistance = radius + float32(float32FromFloat32(4)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter))
	if distanceSquared > float32(maxDistance*maxDistance) {
		return *(*LocalManifold)(unsafe.Add(nil, bp))
	}
	distance = sqrtf(tls, distanceSquared)
	v85 = bp + 44
	v1 = d1
	lengthSquared1 = float32(v1.X*v1.X) + float32(v1.Y*v1.Y)
	if lengthSquared1 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		*(*float32)(unsafe.Add(nil, v85)) = sqrtf(tls, lengthSquared1)
		s2 = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v85))
		u1 = Vec2{
			X: float32(s2 * v1.X),
			Y: float32(s2 * v1.Y),
		}
		v2 = u1
		goto _88
	}
	*(*float32)(unsafe.Add(nil, v85)) = float32FromFloat32(0)
	v2 = Vec2{}
	goto _88
_88:
	u11 = v2
	v89 = bp + 48
	v3 = d2
	lengthSquared1 = float32(v3.X*v3.X) + float32(v3.Y*v3.Y)
	if lengthSquared1 > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		*(*float32)(unsafe.Add(nil, v89)) = sqrtf(tls, lengthSquared1)
		s2 = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v89))
		u1 = Vec2{
			X: float32(s2 * v3.X),
			Y: float32(s2 * v3.Y),
		}
		v5 = u1
		goto _92
	}
	*(*float32)(unsafe.Add(nil, v89)) = float32FromFloat32(0)
	v5 = Vec2{}
	goto _92
_92:
	u2 = v5
	v6 = p2
	v7 = p1
	v10 = Vec2{
		X: v6.X - v7.X,
		Y: v6.Y - v7.Y,
	}
	goto _96
_96:
	v11 = v10
	v14 = u11
	v27 = float32(v11.X*v14.X) + float32(v11.Y*v14.Y)
	goto _100
_100:
	// Does segment B project outside segment A?
	fp2 = v27
	v15 = q2
	v17 = p1
	v18 = Vec2{
		X: v15.X - v17.X,
		Y: v15.Y - v17.Y,
	}
	goto _104
_104:
	v19 = v18
	v21 = u11
	v31 = float32(v19.X*v21.X) + float32(v19.Y*v21.Y)
	goto _108
_108:
	fq2 = v31
	outsideA = boolUint8(fp2 <= float32FromFloat32(0) && fq2 <= float32FromFloat32(0) || fp2 >= *(*float32)(unsafe.Add(nil, bp+44)) && fq2 >= *(*float32)(unsafe.Add(nil, bp+44)))
	v22 = p1
	v23 = p2
	v25 = Vec2{
		X: v22.X - v23.X,
		Y: v22.Y - v23.Y,
	}
	goto _112
_112:
	v26 = v25
	v29 = u2
	v40 = float32(v26.X*v29.X) + float32(v26.Y*v29.Y)
	goto _116
_116:
	// Does segment A project outside segment B?
	fp1 = v40
	v30 = q1
	v34 = p2
	v35 = Vec2{
		X: v30.X - v34.X,
		Y: v30.Y - v34.Y,
	}
	goto _120
_120:
	v36 = v35
	v38 = u2
	v44 = float32(v36.X*v38.X) + float32(v36.Y*v38.Y)
	goto _124
_124:
	fq1 = v44
	outsideB = boolUint8(fp1 <= float32FromFloat32(0) && fq1 <= float32FromFloat32(0) || fp1 >= *(*float32)(unsafe.Add(nil, bp+48)) && fq1 >= *(*float32)(unsafe.Add(nil, bp+48)))
	if int32FromUint8(outsideA) == false1 && int32FromUint8(outsideB) == false1 {
		v39 = u11
		v42 = Vec2{
			X: -v39.Y,
			Y: v39.X,
		}
		goto _127
	_127:
		normalA = v42
		v1 = p2
		v2 = p1
		v3 = Vec2{
			X: v1.X - v2.X,
			Y: v1.Y - v2.Y,
		}
		goto _131
	_131:
		v5 = v3
		v6 = normalA
		v27 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y)
		goto _135
	_135:
		ss1 = v27
		v7 = q2
		v10 = p1
		v11 = Vec2{
			X: v7.X - v10.X,
			Y: v7.Y - v10.Y,
		}
		goto _139
	_139:
		v14 = v11
		v15 = normalA
		v31 = float32(v14.X*v15.X) + float32(v14.Y*v15.Y)
		goto _143
	_143:
		ss2 = v31
		if ss1 < ss2 {
			v40 = ss1
		} else {
			v40 = ss2
		}
		s1p = v40
		if -ss1 < -ss2 {
			v44 = -ss1
		} else {
			v44 = -ss2
		}
		s1n = v44
		if s1p > s1n {
			separationA = s1p
		} else {
			separationA = s1n
			v1 = normalA
			v2 = Vec2{
				X: -v1.X,
				Y: -v1.Y,
			}
			goto _148
		_148:
			normalA = v2
		}
		v1 = u2
		v2 = Vec2{
			X: -v1.Y,
			Y: v1.X,
		}
		goto _151
	_151:
		normalB = v2
		v1 = p1
		v2 = p2
		v3 = Vec2{
			X: v1.X - v2.X,
			Y: v1.Y - v2.Y,
		}
		goto _155
	_155:
		v5 = v3
		v6 = normalB
		v27 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y)
		goto _159
	_159:
		ss11 = v27
		v7 = q1
		v10 = p2
		v11 = Vec2{
			X: v7.X - v10.X,
			Y: v7.Y - v10.Y,
		}
		goto _163
	_163:
		v14 = v11
		v15 = normalB
		v31 = float32(v14.X*v15.X) + float32(v14.Y*v15.Y)
		goto _167
	_167:
		ss21 = v31
		if ss11 < ss21 {
			v40 = ss11
		} else {
			v40 = ss21
		}
		s1p1 = v40
		if -ss11 < -ss21 {
			v44 = -ss11
		} else {
			v44 = -ss21
		}
		s1n1 = v44
		if s1p1 > s1n1 {
			separationB = s1p1
		} else {
			separationB = s1n1
			v1 = normalB
			v2 = Vec2{
				X: -v1.X,
				Y: -v1.Y,
			}
			goto _172
		_172:
			normalB = v2
		}
		// biased to avoid feature flip-flop
		// todo more testing?
		if separationA+float32(float32FromFloat32(0.1)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) >= separationB {
			(*(*LocalManifold)(unsafe.Add(nil, bp))).Normal = normalA
			cp = p2
			cq = q2
			// clip to p1
			if fp2 < float32FromFloat32(0) && fq2 > float32FromFloat32(0) {
				v1 = p2
				v2 = q2
				v27 = (float32FromFloat32(0) - fp2) / (fq2 - fp2)
				v3 = Vec2{
					X: float32((float32FromFloat32(1)-v27)*v1.X) + float32(v27*v2.X),
					Y: float32((float32FromFloat32(1)-v27)*v1.Y) + float32(v27*v2.Y),
				}
				goto _177
			_177:
				cp = v3
			} else {
				if fq2 < float32FromFloat32(0) && fp2 > float32FromFloat32(0) {
					v1 = q2
					v2 = p2
					v27 = (float32FromFloat32(0) - fq2) / (fp2 - fq2)
					v3 = Vec2{
						X: float32((float32FromFloat32(1)-v27)*v1.X) + float32(v27*v2.X),
						Y: float32((float32FromFloat32(1)-v27)*v1.Y) + float32(v27*v2.Y),
					}
					goto _182
				_182:
					cq = v3
				}
			}
			// clip to q1
			if fp2 > *(*float32)(unsafe.Add(nil, bp+44)) && fq2 < *(*float32)(unsafe.Add(nil, bp+44)) {
				v1 = p2
				v2 = q2
				v27 = (fp2 - *(*float32)(unsafe.Add(nil, bp+44))) / (fp2 - fq2)
				v3 = Vec2{
					X: float32((float32FromFloat32(1)-v27)*v1.X) + float32(v27*v2.X),
					Y: float32((float32FromFloat32(1)-v27)*v1.Y) + float32(v27*v2.Y),
				}
				goto _187
			_187:
				cp = v3
			} else {
				if fq2 > *(*float32)(unsafe.Add(nil, bp+44)) && fp2 < *(*float32)(unsafe.Add(nil, bp+44)) {
					v1 = q2
					v2 = p2
					v27 = (fq2 - *(*float32)(unsafe.Add(nil, bp+44))) / (fq2 - fp2)
					v3 = Vec2{
						X: float32((float32FromFloat32(1)-v27)*v1.X) + float32(v27*v2.X),
						Y: float32((float32FromFloat32(1)-v27)*v1.Y) + float32(v27*v2.Y),
					}
					goto _192
				_192:
					cq = v3
				}
			}
			v1 = cp
			v2 = p1
			v3 = Vec2{
				X: v1.X - v2.X,
				Y: v1.Y - v2.Y,
			}
			goto _196
		_196:
			v5 = v3
			v6 = normalA
			v27 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y)
			goto _200
		_200:
			sp = v27
			v7 = cq
			v10 = p1
			v11 = Vec2{
				X: v7.X - v10.X,
				Y: v7.Y - v10.Y,
			}
			goto _204
		_204:
			v14 = v11
			v15 = normalA
			v31 = float32(v14.X*v15.X) + float32(v14.Y*v15.Y)
			goto _208
		_208:
			sq = v31
			if sp <= distance+float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter) || sq <= distance+float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter) {
				mp = bp + 8 + uintptr(0)*16
				v1 = cp
				v27 = float32(float32FromFloat32(0.5) * (radiusA - radiusB - sp))
				v2 = normalA
				v3 = Vec2{
					X: v1.X + float32(v27*v2.X),
					Y: v1.Y + float32(v27*v2.Y),
				}
				goto _213
			_213:
				(*LocalManifoldPoint)(unsafe.Pointer(mp)).Point = v3
				(*LocalManifoldPoint)(unsafe.Pointer(mp)).Separation = sp - radius
				(*LocalManifoldPoint)(unsafe.Pointer(mp)).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(int32FromInt32(0)))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(int32FromInt32(0))))
				mp = bp + 8 + uintptr(1)*16
				v1 = cq
				v27 = float32(float32FromFloat32(0.5) * (radiusA - radiusB - sq))
				v2 = normalA
				v3 = Vec2{
					X: v1.X + float32(v27*v2.X),
					Y: v1.Y + float32(v27*v2.Y),
				}
				goto _218
			_218:
				(*LocalManifoldPoint)(unsafe.Pointer(mp)).Point = v3
				(*LocalManifoldPoint)(unsafe.Pointer(mp)).Separation = sq - radius
				(*LocalManifoldPoint)(unsafe.Pointer(mp)).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(int32FromInt32(0)))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(int32FromInt32(1))))
				(*(*LocalManifold)(unsafe.Add(nil, bp))).PointCount = int32(2)
			}
		} else {
			// normal always points from A to B
			v1 = normalB
			v2 = Vec2{
				X: -v1.X,
				Y: -v1.Y,
			}
			goto _221
		_221:
			(*(*LocalManifold)(unsafe.Add(nil, bp))).Normal = v2
			cp1 = p1
			cq1 = q1
			// clip to p2
			if fp1 < float32FromFloat32(0) && fq1 > float32FromFloat32(0) {
				v1 = p1
				v2 = q1
				v27 = (float32FromFloat32(0) - fp1) / (fq1 - fp1)
				v3 = Vec2{
					X: float32((float32FromFloat32(1)-v27)*v1.X) + float32(v27*v2.X),
					Y: float32((float32FromFloat32(1)-v27)*v1.Y) + float32(v27*v2.Y),
				}
				goto _226
			_226:
				cp1 = v3
			} else {
				if fq1 < float32FromFloat32(0) && fp1 > float32FromFloat32(0) {
					v1 = q1
					v2 = p1
					v27 = (float32FromFloat32(0) - fq1) / (fp1 - fq1)
					v3 = Vec2{
						X: float32((float32FromFloat32(1)-v27)*v1.X) + float32(v27*v2.X),
						Y: float32((float32FromFloat32(1)-v27)*v1.Y) + float32(v27*v2.Y),
					}
					goto _231
				_231:
					cq1 = v3
				}
			}
			// clip to q2
			if fp1 > *(*float32)(unsafe.Add(nil, bp+48)) && fq1 < *(*float32)(unsafe.Add(nil, bp+48)) {
				v1 = p1
				v2 = q1
				v27 = (fp1 - *(*float32)(unsafe.Add(nil, bp+48))) / (fp1 - fq1)
				v3 = Vec2{
					X: float32((float32FromFloat32(1)-v27)*v1.X) + float32(v27*v2.X),
					Y: float32((float32FromFloat32(1)-v27)*v1.Y) + float32(v27*v2.Y),
				}
				goto _236
			_236:
				cp1 = v3
			} else {
				if fq1 > *(*float32)(unsafe.Add(nil, bp+48)) && fp1 < *(*float32)(unsafe.Add(nil, bp+48)) {
					v1 = q1
					v2 = p1
					v27 = (fq1 - *(*float32)(unsafe.Add(nil, bp+48))) / (fq1 - fp1)
					v3 = Vec2{
						X: float32((float32FromFloat32(1)-v27)*v1.X) + float32(v27*v2.X),
						Y: float32((float32FromFloat32(1)-v27)*v1.Y) + float32(v27*v2.Y),
					}
					goto _241
				_241:
					cq1 = v3
				}
			}
			v1 = cp1
			v2 = p2
			v3 = Vec2{
				X: v1.X - v2.X,
				Y: v1.Y - v2.Y,
			}
			goto _245
		_245:
			v5 = v3
			v6 = normalB
			v27 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y)
			goto _249
		_249:
			sp1 = v27
			v7 = cq1
			v10 = p2
			v11 = Vec2{
				X: v7.X - v10.X,
				Y: v7.Y - v10.Y,
			}
			goto _253
		_253:
			v14 = v11
			v15 = normalB
			v31 = float32(v14.X*v15.X) + float32(v14.Y*v15.Y)
			goto _257
		_257:
			sq1 = v31
			if sp1 <= distance+float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter) || sq1 <= distance+float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter) {
				mp1 = bp + 8 + uintptr(0)*16
				v1 = cp1
				v27 = float32(float32FromFloat32(0.5) * (radiusB - radiusA - sp1))
				v2 = normalB
				v3 = Vec2{
					X: v1.X + float32(v27*v2.X),
					Y: v1.Y + float32(v27*v2.Y),
				}
				goto _262
			_262:
				(*LocalManifoldPoint)(unsafe.Pointer(mp1)).Point = v3
				(*LocalManifoldPoint)(unsafe.Pointer(mp1)).Separation = sp1 - radius
				(*LocalManifoldPoint)(unsafe.Pointer(mp1)).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(int32FromInt32(0)))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(int32FromInt32(0))))
				mp1 = bp + 8 + uintptr(1)*16
				v1 = cq1
				v27 = float32(float32FromFloat32(0.5) * (radiusB - radiusA - sq1))
				v2 = normalB
				v3 = Vec2{
					X: v1.X + float32(v27*v2.X),
					Y: v1.Y + float32(v27*v2.Y),
				}
				goto _267
			_267:
				(*LocalManifoldPoint)(unsafe.Pointer(mp1)).Point = v3
				(*LocalManifoldPoint)(unsafe.Pointer(mp1)).Separation = sq1 - radius
				(*LocalManifoldPoint)(unsafe.Pointer(mp1)).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(int32FromInt32(1)))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(int32FromInt32(0))))
				(*(*LocalManifold)(unsafe.Add(nil, bp))).PointCount = int32(2)
			}
		}
	}
	if (*(*LocalManifold)(unsafe.Add(nil, bp))).PointCount == 0 {
		v1 = closest2
		v2 = closest1
		v3 = Vec2{
			X: v1.X - v2.X,
			Y: v1.Y - v2.Y,
		}
		goto _271
	_271:
		// single point collision
		normal = v3
		v5 = normal
		v6 = normal
		v27 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y)
		goto _275
	_275:
		if v27 > epsSqr {
			v7 = normal
			lengthSquared = float32(v7.X*v7.X) + float32(v7.Y*v7.Y)
			if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
				s1 = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
				u = Vec2{
					X: float32(s1 * v7.X),
					Y: float32(s1 * v7.Y),
				}
				v10 = u
				goto _278
			}
			v10 = Vec2{}
			goto _278
		_278:
			normal = v10
		} else {
			v1 = u11
			v2 = Vec2{
				X: -v1.Y,
				Y: v1.X,
			}
			goto _281
		_281:
			normal = v2
		}
		v1 = closest1
		v27 = radiusA
		v2 = normal
		v3 = Vec2{
			X: v1.X + float32(v27*v2.X),
			Y: v1.Y + float32(v27*v2.Y),
		}
		goto _286
	_286:
		c1 = v3
		v5 = closest2
		v31 = -radiusB
		v6 = normal
		v7 = Vec2{
			X: v5.X + float32(v31*v6.X),
			Y: v5.Y + float32(v31*v6.Y),
		}
		goto _291
	_291:
		c2 = v7
		if f1 == float32FromFloat32(0) {
			v292 = 0
		} else {
			v292 = int32(1)
		}
		i1 = v292
		if f2 == float32FromFloat32(0) {
			v293 = 0
		} else {
			v293 = int32(1)
		}
		i2 = v293
		(*(*LocalManifold)(unsafe.Add(nil, bp))).Normal = normal
		v1 = c1
		v2 = c2
		v27 = float32FromFloat32(0.5)
		v3 = Vec2{
			X: float32((float32FromFloat32(1)-v27)*v1.X) + float32(v27*v2.X),
			Y: float32((float32FromFloat32(1)-v27)*v1.Y) + float32(v27*v2.Y),
		}
		goto _298
	_298:
		(*(*LocalManifoldPoint)(unsafe.Add(nil, bp+8))).Point = v3
		(*(*LocalManifoldPoint)(unsafe.Add(nil, bp+8))).Separation = sqrtf(tls, distanceSquared) - radius
		(*(*LocalManifoldPoint)(unsafe.Add(nil, bp+8))).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(i1))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(i2)))
		(*(*LocalManifold)(unsafe.Add(nil, bp))).PointCount = int32(1)
	}
	// Undo the origin shift so points are in frame A
	i = 0
	for {
		if !(i < (*(*LocalManifold)(unsafe.Add(nil, bp))).PointCount) {
			break
		}
		v1 = (*(*LocalManifoldPoint)(unsafe.Add(nil, bp+8+uintptr(i)*16))).Point
		v2 = origin
		v3 = Vec2{
			X: v1.X + v2.X,
			Y: v1.Y + v2.Y,
		}
		goto _303
	_303:
		(*(*LocalManifoldPoint)(unsafe.Add(nil, bp+8+uintptr(i)*16))).Point = v3
		goto _299
	_299:
		;
		i = i + 1
	}
	return *(*LocalManifold)(unsafe.Add(nil, bp))
}

func b2CollideSegmentAndCapsule(tls *_Stack, segmentA uintptr, capsuleB uintptr, xf Transform) (r LocalManifold) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* capsuleA at bp+0 */ Capsule
	*(*Capsule)(unsafe.Add(nil, bp)) = Capsule{
		Center1: (*Segment)(unsafe.Pointer(segmentA)).Point1,
		Center2: (*Segment)(unsafe.Pointer(segmentA)).Point2,
	}
	return b2CollideCapsules(tls, bp, capsuleB, xf)
}

func b2CollidePolygonAndCapsule(tls *_Stack, polygonA uintptr, capsuleB uintptr, xf Transform) (r LocalManifold) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var _ /* polyB at bp+0 */ Polygon
	*(*Polygon)(unsafe.Add(nil, bp)) = b2MakeCapsule(tls, (*Capsule)(unsafe.Pointer(capsuleB)).Center1, (*Capsule)(unsafe.Pointer(capsuleB)).Center2, (*Capsule)(unsafe.Pointer(capsuleB)).Radius)
	return b2CollidePolygons(tls, polygonA, bp, xf)
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

func b2CollidePolygons(tls *_Stack, polygonA uintptr, polygonB uintptr, xf Transform) (r LocalManifold) {
	bp := tls.Alloc(352)
	defer tls.Free(352)
	var c1, c11, c12, c13, c2, c21, c22, c23, normal, normal1, normal2, normal3, origin, searchDirection, searchDirection1, v11, v12, v21, v22, v1, v2, v3, v6, v7, v8 Vec2
	var count, count1, i, i1, i11, i12, i2, i21, i22, i3, i4, i5, v29, v30 int32
	var distance, dot, dot1, invDistance, invDistance1, invDistance2, invDistance3, linearSlop, minDot, minDot1, minSeparation, radius, separation, separationA, separationB, speculativeDistance, x, y, v221, v27, v33, v34 float32
	var flip uint8
	var normals, normals1 uintptr
	var result SegmentDistanceResult
	var xfs, v111 Transform
	var v15 Rot
	var v31 bool
	var _ /* edgeA at bp+288 */ int32
	var _ /* edgeB at bp+292 */ int32
	var _ /* localPolyA at bp+0 */ Polygon
	var _ /* localPolyB at bp+144 */ Polygon
	var _ /* manifold at bp+296 */ LocalManifold
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c1, c11, c12, c13, c2, c21, c22, c23, count, count1, distance, dot, dot1, flip, i, i1, i11, i12, i2, i21, i22, i3, i4, i5, invDistance, invDistance1, invDistance2, invDistance3, linearSlop, minDot, minDot1, minSeparation, normal, normal1, normal2, normal3, normals, normals1, origin, radius, result, searchDirection, searchDirection1, separation, separationA, separationB, speculativeDistance, v11, v12, v21, v22, x, xfs, y, v1, v111, v15, v2, v221, v27, v29, v3, v30, v31, v33, v34, v6, v7, v8
	origin = *(*Vec2)(unsafe.Add(nil, polygonA))
	linearSlop = float32(float32FromFloat32(0.005) * b2_lengthUnitsPerMeter)
	speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter))
	v1 = xf.P
	v2 = origin
	v3 = Vec2{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
	goto _4
_4:
	// Shift to the origin in frame A for round-off, a pure translation in A's frame
	xfs = Transform{
		P: v3,
		Q: xf.Q,
	}
	(*(*Polygon)(unsafe.Add(nil, bp))).Count = (*Polygon)(unsafe.Pointer(polygonA)).Count
	(*(*Polygon)(unsafe.Add(nil, bp))).Radius = (*Polygon)(unsafe.Pointer(polygonA)).Radius
	*(*Vec2)(unsafe.Add(nil, bp)) = b2Vec2_zero
	*(*Vec2)(unsafe.Add(nil, bp+64)) = *(*Vec2)(unsafe.Add(nil, polygonA+64))
	i = int32(1)
	for {
		if !(i < (*(*Polygon)(unsafe.Add(nil, bp))).Count) {
			break
		}
		v1 = *(*Vec2)(unsafe.Add(nil, polygonA+uintptr(i)*8))
		v2 = origin
		v3 = Vec2{
			X: v1.X - v2.X,
			Y: v1.Y - v2.Y,
		}
		goto _9
	_9:
		*(*Vec2)(unsafe.Add(nil, bp+uintptr(i)*8)) = v3
		*(*Vec2)(unsafe.Add(nil, bp+64+uintptr(i)*8)) = *(*Vec2)(unsafe.Add(nil, polygonA+64+uintptr(i)*8))
		goto _5
	_5:
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
		v111 = xfs
		v1 = *(*Vec2)(unsafe.Add(nil, polygonB+uintptr(i1)*8))
		x = float32(v111.Q.C*v1.X) - float32(v111.Q.S*v1.Y) + v111.P.X
		y = float32(v111.Q.S*v1.X) + float32(v111.Q.C*v1.Y) + v111.P.Y
		v2 = Vec2{
			X: x,
			Y: y,
		}
		goto _14
	_14:
		*(*Vec2)(unsafe.Add(nil, bp+144+uintptr(i1)*8)) = v2
		v15 = xfs.Q
		v1 = *(*Vec2)(unsafe.Add(nil, polygonB+64+uintptr(i1)*8))
		v2 = Vec2{
			X: float32(v15.C*v1.X) - float32(v15.S*v1.Y),
			Y: float32(v15.S*v1.X) + float32(v15.C*v1.Y),
		}
		goto _18
	_18:
		*(*Vec2)(unsafe.Add(nil, bp+144+64+uintptr(i1)*8)) = v2
		goto _10
	_10:
		;
		i1 = i1 + 1
	}
	radius = (*(*Polygon)(unsafe.Add(nil, bp))).Radius + (*(*Polygon)(unsafe.Add(nil, bp+144))).Radius
	*(*int32)(unsafe.Add(nil, bp+288)) = 0
	separationA = b2FindMaxSeparation(tls, bp+288, bp, bp+144)
	if separationA > speculativeDistance+radius {
		return LocalManifold{}
	}
	*(*int32)(unsafe.Add(nil, bp+292)) = 0
	separationB = b2FindMaxSeparation(tls, bp+292, bp+144, bp)
	if separationB > speculativeDistance+radius {
		return LocalManifold{}
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
			v1 = searchDirection
			v2 = *(*Vec2)(unsafe.Add(nil, normals+uintptr(i2)*8))
			v221 = float32(v1.X*v2.X) + float32(v1.Y*v2.Y)
			goto _23
		_23:
			dot = v221
			if dot < minDot {
				minDot = dot
				*(*int32)(unsafe.Add(nil, bp+292)) = i2
			}
			goto _19
		_19:
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
			v1 = searchDirection1
			v2 = *(*Vec2)(unsafe.Add(nil, normals1+uintptr(i3)*8))
			v221 = float32(v1.X*v2.X) + float32(v1.Y*v2.Y)
			goto _28
		_28:
			dot1 = v221
			if dot1 < minDot1 {
				minDot1 = dot1
				*(*int32)(unsafe.Add(nil, bp+288)) = i3
			}
			goto _24
		_24:
			;
			i3 = i3 + 1
		}
	}
	*(*LocalManifold)(unsafe.Add(nil, bp+296)) = LocalManifold{}
	// Using slop here to ensure vertex-vertex normal vectors can be safely normalized
	// todo this means edge clipping needs to handle slightly non-overlapping edges.
	if separationA > float32(float32FromFloat32(0.1)*linearSlop) || separationB > float32(float32FromFloat32(0.1)*linearSlop) {
		// Edges are disjoint. Find closest points between reference edge and incident edge
		// Reference edge on polygon A
		i11 = *(*int32)(unsafe.Add(nil, bp+288))
		if *(*int32)(unsafe.Add(nil, bp+288))+int32(1) < (*(*Polygon)(unsafe.Add(nil, bp))).Count {
			v29 = *(*int32)(unsafe.Add(nil, bp+288)) + int32(1)
		} else {
			v29 = 0
		}
		i12 = v29
		i21 = *(*int32)(unsafe.Add(nil, bp+292))
		if *(*int32)(unsafe.Add(nil, bp+292))+int32(1) < (*(*Polygon)(unsafe.Add(nil, bp+144))).Count {
			v30 = *(*int32)(unsafe.Add(nil, bp+292)) + int32(1)
		} else {
			v30 = 0
		}
		i22 = v30
		v11 = *(*Vec2)(unsafe.Add(nil, bp+uintptr(i11)*8))
		v12 = *(*Vec2)(unsafe.Add(nil, bp+uintptr(i12)*8))
		v21 = *(*Vec2)(unsafe.Add(nil, bp+144+uintptr(i21)*8))
		v22 = *(*Vec2)(unsafe.Add(nil, bp+144+uintptr(i22)*8))
		result = b2ComputeSegmentDistance(tls, v11, v12, v21, v22)
		if v31 = !!(result.DistanceSquared > float32FromFloat32(0)); !v31 {
			b2InternalAssert(tls, __ccgo_ts+15891, __ccgo_ts+15836, int32FromInt32(880))
		}
		_ = v31 || bool(int32FromInt32(0) != 0)
		distance = sqrtf(tls, result.DistanceSquared)
		separation = distance - radius
		if distance-radius > speculativeDistance {
			// This can happen in the vertex-vertex case
			return *(*LocalManifold)(unsafe.Add(nil, bp+296))
		}
		// Attempt to clip edges
		*(*LocalManifold)(unsafe.Add(nil, bp+296)) = b2ClipPolygons(tls, bp, bp+144, *(*int32)(unsafe.Add(nil, bp+288)), *(*int32)(unsafe.Add(nil, bp+292)), flip)
		minSeparation = float32FromFloat32(3.4028234663852886e+38)
		i4 = 0
		for {
			if !(i4 < (*(*LocalManifold)(unsafe.Add(nil, bp+296))).PointCount) {
				break
			}
			v221 = minSeparation
			v27 = (*(*LocalManifoldPoint)(unsafe.Add(nil, bp+296+8+uintptr(i4)*16))).Separation
			if v221 < v27 {
				v34 = v221
			} else {
				v34 = v27
			}
			v33 = v34
			goto _36
		_36:
			minSeparation = v33
			goto _32
		_32:
			;
			i4 = i4 + 1
		}
		// Does vertex-vertex have substantially larger separation?
		if separation+float32(float32FromFloat32(0.1)*linearSlop) < minSeparation {
			if result.Fraction1 == float32FromFloat32(0) && result.Fraction2 == float32FromFloat32(0) {
				v1 = v21
				v2 = v11
				v3 = Vec2{
					X: v1.X - v2.X,
					Y: v1.Y - v2.Y,
				}
				goto _41
			_41:
				// v11 - v21
				normal = v3
				invDistance = float32FromFloat32(1) / distance
				normal.X = float32(normal.X * invDistance)
				normal.Y = float32(normal.Y * invDistance)
				v1 = v11
				v221 = (*(*Polygon)(unsafe.Add(nil, bp))).Radius
				v2 = normal
				v3 = Vec2{
					X: v1.X + float32(v221*v2.X),
					Y: v1.Y + float32(v221*v2.Y),
				}
				goto _46
			_46:
				c1 = v3
				v6 = v21
				v27 = -(*(*Polygon)(unsafe.Add(nil, bp+144))).Radius
				v7 = normal
				v8 = Vec2{
					X: v6.X + float32(v27*v7.X),
					Y: v6.Y + float32(v27*v7.Y),
				}
				goto _51
			_51:
				c2 = v8
				(*(*LocalManifold)(unsafe.Add(nil, bp+296))).Normal = normal
				v1 = c1
				v2 = c2
				v221 = float32FromFloat32(0.5)
				v3 = Vec2{
					X: float32((float32FromFloat32(1)-v221)*v1.X) + float32(v221*v2.X),
					Y: float32((float32FromFloat32(1)-v221)*v1.Y) + float32(v221*v2.Y),
				}
				goto _56
			_56:
				(*(*LocalManifoldPoint)(unsafe.Add(nil, bp+296+8))).Point = v3
				(*(*LocalManifoldPoint)(unsafe.Add(nil, bp+296+8))).Separation = distance - radius
				(*(*LocalManifoldPoint)(unsafe.Add(nil, bp+296+8))).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(i11))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(i21)))
				(*(*LocalManifold)(unsafe.Add(nil, bp+296))).PointCount = int32(1)
			} else {
				if result.Fraction1 == float32FromFloat32(0) && result.Fraction2 == float32FromFloat32(1) {
					v1 = v22
					v2 = v11
					v3 = Vec2{
						X: v1.X - v2.X,
						Y: v1.Y - v2.Y,
					}
					goto _60
				_60:
					// v11 - v22
					normal1 = v3
					invDistance1 = float32FromFloat32(1) / distance
					normal1.X = float32(normal1.X * invDistance1)
					normal1.Y = float32(normal1.Y * invDistance1)
					v1 = v11
					v221 = (*(*Polygon)(unsafe.Add(nil, bp))).Radius
					v2 = normal1
					v3 = Vec2{
						X: v1.X + float32(v221*v2.X),
						Y: v1.Y + float32(v221*v2.Y),
					}
					goto _65
				_65:
					c11 = v3
					v6 = v22
					v27 = -(*(*Polygon)(unsafe.Add(nil, bp+144))).Radius
					v7 = normal1
					v8 = Vec2{
						X: v6.X + float32(v27*v7.X),
						Y: v6.Y + float32(v27*v7.Y),
					}
					goto _70
				_70:
					c21 = v8
					(*(*LocalManifold)(unsafe.Add(nil, bp+296))).Normal = normal1
					v1 = c11
					v2 = c21
					v221 = float32FromFloat32(0.5)
					v3 = Vec2{
						X: float32((float32FromFloat32(1)-v221)*v1.X) + float32(v221*v2.X),
						Y: float32((float32FromFloat32(1)-v221)*v1.Y) + float32(v221*v2.Y),
					}
					goto _75
				_75:
					(*(*LocalManifoldPoint)(unsafe.Add(nil, bp+296+8))).Point = v3
					(*(*LocalManifoldPoint)(unsafe.Add(nil, bp+296+8))).Separation = distance - radius
					(*(*LocalManifoldPoint)(unsafe.Add(nil, bp+296+8))).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(i11))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(i22)))
					(*(*LocalManifold)(unsafe.Add(nil, bp+296))).PointCount = int32(1)
				} else {
					if result.Fraction1 == float32FromFloat32(1) && result.Fraction2 == float32FromFloat32(0) {
						v1 = v21
						v2 = v12
						v3 = Vec2{
							X: v1.X - v2.X,
							Y: v1.Y - v2.Y,
						}
						goto _79
					_79:
						// v12 - v21
						normal2 = v3
						invDistance2 = float32FromFloat32(1) / distance
						normal2.X = float32(normal2.X * invDistance2)
						normal2.Y = float32(normal2.Y * invDistance2)
						v1 = v12
						v221 = (*(*Polygon)(unsafe.Add(nil, bp))).Radius
						v2 = normal2
						v3 = Vec2{
							X: v1.X + float32(v221*v2.X),
							Y: v1.Y + float32(v221*v2.Y),
						}
						goto _84
					_84:
						c12 = v3
						v6 = v21
						v27 = -(*(*Polygon)(unsafe.Add(nil, bp+144))).Radius
						v7 = normal2
						v8 = Vec2{
							X: v6.X + float32(v27*v7.X),
							Y: v6.Y + float32(v27*v7.Y),
						}
						goto _89
					_89:
						c22 = v8
						(*(*LocalManifold)(unsafe.Add(nil, bp+296))).Normal = normal2
						v1 = c12
						v2 = c22
						v221 = float32FromFloat32(0.5)
						v3 = Vec2{
							X: float32((float32FromFloat32(1)-v221)*v1.X) + float32(v221*v2.X),
							Y: float32((float32FromFloat32(1)-v221)*v1.Y) + float32(v221*v2.Y),
						}
						goto _94
					_94:
						(*(*LocalManifoldPoint)(unsafe.Add(nil, bp+296+8))).Point = v3
						(*(*LocalManifoldPoint)(unsafe.Add(nil, bp+296+8))).Separation = distance - radius
						(*(*LocalManifoldPoint)(unsafe.Add(nil, bp+296+8))).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(i12))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(i21)))
						(*(*LocalManifold)(unsafe.Add(nil, bp+296))).PointCount = int32(1)
					} else {
						if result.Fraction1 == float32FromFloat32(1) && result.Fraction2 == float32FromFloat32(1) {
							v1 = v22
							v2 = v12
							v3 = Vec2{
								X: v1.X - v2.X,
								Y: v1.Y - v2.Y,
							}
							goto _98
						_98:
							// v12 - v22
							normal3 = v3
							invDistance3 = float32FromFloat32(1) / distance
							normal3.X = float32(normal3.X * invDistance3)
							normal3.Y = float32(normal3.Y * invDistance3)
							v1 = v12
							v221 = (*(*Polygon)(unsafe.Add(nil, bp))).Radius
							v2 = normal3
							v3 = Vec2{
								X: v1.X + float32(v221*v2.X),
								Y: v1.Y + float32(v221*v2.Y),
							}
							goto _103
						_103:
							c13 = v3
							v6 = v22
							v27 = -(*(*Polygon)(unsafe.Add(nil, bp+144))).Radius
							v7 = normal3
							v8 = Vec2{
								X: v6.X + float32(v27*v7.X),
								Y: v6.Y + float32(v27*v7.Y),
							}
							goto _108
						_108:
							c23 = v8
							(*(*LocalManifold)(unsafe.Add(nil, bp+296))).Normal = normal3
							v1 = c13
							v2 = c23
							v221 = float32FromFloat32(0.5)
							v3 = Vec2{
								X: float32((float32FromFloat32(1)-v221)*v1.X) + float32(v221*v2.X),
								Y: float32((float32FromFloat32(1)-v221)*v1.Y) + float32(v221*v2.Y),
							}
							goto _113
						_113:
							(*(*LocalManifoldPoint)(unsafe.Add(nil, bp+296+8))).Point = v3
							(*(*LocalManifoldPoint)(unsafe.Add(nil, bp+296+8))).Separation = distance - radius
							(*(*LocalManifoldPoint)(unsafe.Add(nil, bp+296+8))).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(i12))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(i22)))
							(*(*LocalManifold)(unsafe.Add(nil, bp+296))).PointCount = int32(1)
						}
					}
				}
			}
		}
	} else {
		// Polygons overlap
		*(*LocalManifold)(unsafe.Add(nil, bp+296)) = b2ClipPolygons(tls, bp, bp+144, *(*int32)(unsafe.Add(nil, bp+288)), *(*int32)(unsafe.Add(nil, bp+292)), flip)
	}
	// Undo the origin shift so points are in frame A
	i5 = 0
	for {
		if !(i5 < (*(*LocalManifold)(unsafe.Add(nil, bp+296))).PointCount) {
			break
		}
		v1 = (*(*LocalManifoldPoint)(unsafe.Add(nil, bp+296+8+uintptr(i5)*16))).Point
		v2 = origin
		v3 = Vec2{
			X: v1.X + v2.X,
			Y: v1.Y + v2.Y,
		}
		goto _118
	_118:
		(*(*LocalManifoldPoint)(unsafe.Add(nil, bp+296+8+uintptr(i5)*16))).Point = v3
		goto _114
	_114:
		;
		i5 = i5 + 1
	}
	return *(*LocalManifold)(unsafe.Add(nil, bp+296))
}

func b2CollideSegmentAndCircle(tls *_Stack, segmentA uintptr, circleB uintptr, xf Transform) (r LocalManifold) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* capsuleA at bp+0 */ Capsule
	*(*Capsule)(unsafe.Add(nil, bp)) = Capsule{
		Center1: (*Segment)(unsafe.Pointer(segmentA)).Point1,
		Center2: (*Segment)(unsafe.Pointer(segmentA)).Point2,
	}
	return b2CollideCapsuleAndCircle(tls, bp, circleB, xf)
}

func b2CollideSegmentAndPolygon(tls *_Stack, segmentA uintptr, polygonB uintptr, xf Transform) (r LocalManifold) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var _ /* polygonA at bp+0 */ Polygon
	*(*Polygon)(unsafe.Add(nil, bp)) = b2MakeCapsule(tls, (*Segment)(unsafe.Pointer(segmentA)).Point1, (*Segment)(unsafe.Pointer(segmentA)).Point2, float32FromFloat32(0))
	return b2CollidePolygons(tls, bp, polygonB, xf)
}

func b2CollideChainSegmentAndCircle(tls *_Stack, segmentA uintptr, circleB uintptr, xf Transform) (r LocalManifold) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var cA, cB, e, nextEdge, normal, p1, p2, pA, pB, prevEdge, u, v10, v12, v13, v14, v16, v17, v21, v20, v211, v22, v24, v25, v28, v3, v5, v6, v7, v9 Vec2
	var ee, lengthSquared, offset, radius, s2, separation, u1, uPrev, v2, vNext, x, y, v18, v26, v34 float32
	var mp, v73 uintptr
	var v1 Transform
	var _ /* distance at bp+44 */ float32
	var _ /* manifold at bp+0 */ LocalManifold
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = cA, cB, e, ee, lengthSquared, mp, nextEdge, normal, offset, p1, p2, pA, pB, prevEdge, radius, s2, separation, u, u1, uPrev, v2, vNext, x, y, v1, v10, v12, v13, v14, v16, v17, v18, v21, v20, v211, v22, v24, v25, v26, v28, v3, v34, v5, v6, v7, v73, v9
	*(*LocalManifold)(unsafe.Add(nil, bp)) = LocalManifold{}
	v1 = xf
	v21 = (*Circle)(unsafe.Pointer(circleB)).Center
	x = float32(v1.Q.C*v21.X) - float32(v1.Q.S*v21.Y) + v1.P.X
	y = float32(v1.Q.S*v21.X) + float32(v1.Q.C*v21.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	// Compute circle in frame of segment
	pB = v3
	p1 = (*ChainSegment)(unsafe.Pointer(segmentA)).Segment.Point1
	p2 = (*ChainSegment)(unsafe.Pointer(segmentA)).Segment.Point2
	v5 = p2
	v6 = p1
	v7 = Vec2{
		X: v5.X - v6.X,
		Y: v5.Y - v6.Y,
	}
	goto _8
_8:
	e = v7
	v9 = e
	v10 = Vec2{
		X: v9.Y,
		Y: -v9.X,
	}
	goto _11
_11:
	v12 = pB
	v13 = p1
	v14 = Vec2{
		X: v12.X - v13.X,
		Y: v12.Y - v13.Y,
	}
	goto _15
_15:
	v16 = v10
	v17 = v14
	v18 = float32(v16.X*v17.X) + float32(v16.Y*v17.Y)
	goto _19
_19:
	// Normal points to the right
	offset = v18
	if offset < float32FromFloat32(0) {
		// collision is one-sided
		return *(*LocalManifold)(unsafe.Add(nil, bp))
	}
	v21 = p2
	v3 = pB
	v5 = Vec2{
		X: v21.X - v3.X,
		Y: v21.Y - v3.Y,
	}
	goto _23
_23:
	v6 = e
	v7 = v5
	v18 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y)
	goto _27
_27:
	// Barycentric coordinates
	u1 = v18
	v9 = pB
	v10 = p1
	v12 = Vec2{
		X: v9.X - v10.X,
		Y: v9.Y - v10.Y,
	}
	goto _31
_31:
	v13 = e
	v14 = v12
	v26 = float32(v13.X*v14.X) + float32(v13.Y*v14.Y)
	goto _35
_35:
	v2 = v26
	if v2 <= float32FromFloat32(0) {
		v16 = p1
		v17 = (*ChainSegment)(unsafe.Pointer(segmentA)).Ghost1
		v20 = Vec2{
			X: v16.X - v17.X,
			Y: v16.Y - v17.Y,
		}
		goto _39
	_39:
		// Behind point1?
		// Is pB in the Voronoi region of the previous edge?
		prevEdge = v20
		v211 = pB
		v22 = p1
		v24 = Vec2{
			X: v211.X - v22.X,
			Y: v211.Y - v22.Y,
		}
		goto _43
	_43:
		v25 = prevEdge
		v28 = v24
		v34 = float32(v25.X*v28.X) + float32(v25.Y*v28.Y)
		goto _47
	_47:
		uPrev = v34
		if uPrev <= float32FromFloat32(0) {
			return *(*LocalManifold)(unsafe.Add(nil, bp))
		}
		pA = p1
	} else {
		if u1 <= float32FromFloat32(0) {
			v21 = (*ChainSegment)(unsafe.Pointer(segmentA)).Ghost2
			v3 = p2
			v5 = Vec2{
				X: v21.X - v3.X,
				Y: v21.Y - v3.Y,
			}
			goto _51
		_51:
			// Ahead of point2?
			nextEdge = v5
			v6 = pB
			v7 = p2
			v9 = Vec2{
				X: v6.X - v7.X,
				Y: v6.Y - v7.Y,
			}
			goto _55
		_55:
			v10 = nextEdge
			v12 = v9
			v18 = float32(v10.X*v12.X) + float32(v10.Y*v12.Y)
			goto _59
		_59:
			vNext = v18
			// Is pB in the Voronoi region of the next edge?
			if vNext > float32FromFloat32(0) {
				return *(*LocalManifold)(unsafe.Add(nil, bp))
			}
			pA = p2
		} else {
			v21 = e
			v3 = e
			v18 = float32(v21.X*v3.X) + float32(v21.Y*v3.Y)
			goto _63
		_63:
			ee = v18
			pA = Vec2{
				X: float32(u1*p1.X) + float32(v2*p2.X),
				Y: float32(u1*p1.Y) + float32(v2*p2.Y),
			}
			if ee > float32FromFloat32(0) {
				v18 = float32FromFloat32(1) / ee
				v3 = pA
				v5 = Vec2{
					X: float32(v18 * v3.X),
					Y: float32(v18 * v3.Y),
				}
				goto _68
			_68:
				v21 = v5
			} else {
				v21 = p1
			}
			pA = v21
		}
	}
	v21 = pB
	v3 = pA
	v5 = Vec2{
		X: v21.X - v3.X,
		Y: v21.Y - v3.Y,
	}
	goto _72
_72:
	v73 = bp + 44
	v6 = v5
	lengthSquared = float32(v6.X*v6.X) + float32(v6.Y*v6.Y)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		*(*float32)(unsafe.Add(nil, v73)) = sqrtf(tls, lengthSquared)
		s2 = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v73))
		u = Vec2{
			X: float32(s2 * v6.X),
			Y: float32(s2 * v6.Y),
		}
		v7 = u
		goto _76
	}
	*(*float32)(unsafe.Add(nil, v73)) = float32FromFloat32(0)
	v7 = Vec2{}
	goto _76
_76:
	normal = v7
	radius = (*Circle)(unsafe.Pointer(circleB)).Radius
	separation = *(*float32)(unsafe.Add(nil, bp+44)) - radius
	if separation > float32(float32FromFloat32(4)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
		return *(*LocalManifold)(unsafe.Add(nil, bp))
	}
	cA = pA
	v21 = pB
	v18 = -radius
	v3 = normal
	v5 = Vec2{
		X: v21.X + float32(v18*v3.X),
		Y: v21.Y + float32(v18*v3.Y),
	}
	goto _81
_81:
	cB = v5
	(*(*LocalManifold)(unsafe.Add(nil, bp))).Normal = normal
	mp = bp + 8 + uintptr(0)*16
	v21 = cA
	v3 = cB
	v18 = float32FromFloat32(0.5)
	v5 = Vec2{
		X: float32((float32FromFloat32(1)-v18)*v21.X) + float32(v18*v3.X),
		Y: float32((float32FromFloat32(1)-v18)*v21.Y) + float32(v18*v3.Y),
	}
	goto _86
_86:
	(*LocalManifoldPoint)(unsafe.Pointer(mp)).Point = v5
	(*LocalManifoldPoint)(unsafe.Pointer(mp)).Separation = separation
	(*LocalManifoldPoint)(unsafe.Pointer(mp)).Id = uint16(0)
	(*(*LocalManifold)(unsafe.Add(nil, bp))).PointCount = int32(1)
	return *(*LocalManifold)(unsafe.Add(nil, bp))
}

func b2CollideChainSegmentAndCapsule(tls *_Stack, segmentA uintptr, capsuleB uintptr, xf Transform, cache uintptr) (r LocalManifold) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var _ /* polyB at bp+0 */ Polygon
	*(*Polygon)(unsafe.Add(nil, bp)) = b2MakeCapsule(tls, (*Capsule)(unsafe.Pointer(capsuleB)).Center1, (*Capsule)(unsafe.Pointer(capsuleB)).Center2, (*Capsule)(unsafe.Pointer(capsuleB)).Radius)
	return b2CollideChainSegmentAndPolygon(tls, segmentA, bp, xf, cache)
}

func b2CollideChainSegmentAndPolygon(tls *_Stack, segmentA uintptr, polygonB uintptr, xf Transform, cache uintptr) (r LocalManifold) {
	bp := tls.Alloc(272)
	defer tls.Free(272)
	var a11, a21, b11, b12, b21, b22, centroidB, edge0, edge1, edge2, n, n0, n1, n2, normal, normal1, normalB, p1, p11, p2, pA, pB, u, v10, v12, v13, v14, v16, v17, v19, v2, v20, v3, v5, v6, v7, v9 Vec2
	var behind0, behind1, behind2 uint8
	var convexTol, d1, d2, dot1, dot11, dot12, dot2, dot21, dot22, edgeSeparation, lengthSquared, polygonSeparation, radiusB, s, s0, s1, s2, s21, s3, s4, x, y, v24, v38, v49, v57, v65, v93 float32
	var count, i, i1, i11, i2, i21, i3, i4, ia1, ia11, ia2, ia21, ib, ib1, ib11, ib2, ib21, incidentIndex, incidentNormal, referenceIndex, v99 int32
	var cp uintptr
	var normals [8]Vec2
	var output DistanceOutput
	var smoothParams b2ChainSegmentParams
	var type1, type11, type2 b2NormalType
	var v1 Transform
	var v85 bool
	var v72 Rot
	var _ /* input at bp+108 */ DistanceInput
	var _ /* manifold at bp+0 */ LocalManifold
	var _ /* vertices at bp+44 */ [8]Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a11, a21, b11, b12, b21, b22, behind0, behind1, behind2, centroidB, convexTol, count, cp, d1, d2, dot1, dot11, dot12, dot2, dot21, dot22, edge0, edge1, edge2, edgeSeparation, i, i1, i11, i2, i21, i3, i4, ia1, ia11, ia2, ia21, ib, ib1, ib11, ib2, ib21, incidentIndex, incidentNormal, lengthSquared, n, n0, n1, n2, normal, normal1, normalB, normals, output, p1, p11, p2, pA, pB, polygonSeparation, radiusB, referenceIndex, s, s0, s1, s2, s21, s3, s4, smoothParams, type1, type11, type2, u, x, y, v1, v10, v12, v13, v14, v16, v17, v19, v2, v20, v24, v3, v38, v49, v5, v57, v6, v65, v7, v72, v85, v9, v93, v99
	*(*LocalManifold)(unsafe.Add(nil, bp)) = LocalManifold{}
	v1 = xf
	v2 = (*Polygon)(unsafe.Pointer(polygonB)).Centroid
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	centroidB = v3
	radiusB = (*Polygon)(unsafe.Pointer(polygonB)).Radius
	p11 = (*ChainSegment)(unsafe.Pointer(segmentA)).Segment.Point1
	p2 = (*ChainSegment)(unsafe.Pointer(segmentA)).Segment.Point2
	v5 = p2
	v6 = p11
	v7 = Vec2{
		X: v5.X - v6.X,
		Y: v5.Y - v6.Y,
	}
	goto _8
_8:
	v9 = v7
	lengthSquared = float32(v9.X*v9.X) + float32(v9.Y*v9.Y)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
		u = Vec2{
			X: float32(s * v9.X),
			Y: float32(s * v9.Y),
		}
		v10 = u
		goto _11
	}
	v10 = Vec2{}
	goto _11
_11:
	edge1 = v10
	smoothParams = b2ChainSegmentParams{}
	smoothParams.Edge1 = edge1
	convexTol = float32FromFloat32(0.01)
	v2 = p11
	v3 = (*ChainSegment)(unsafe.Pointer(segmentA)).Ghost1
	v5 = Vec2{
		X: v2.X - v3.X,
		Y: v2.Y - v3.Y,
	}
	goto _15
_15:
	v6 = v5
	lengthSquared = float32(v6.X*v6.X) + float32(v6.Y*v6.Y)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
		u = Vec2{
			X: float32(s * v6.X),
			Y: float32(s * v6.Y),
		}
		v7 = u
		goto _18
	}
	v7 = Vec2{}
	goto _18
_18:
	edge0 = v7
	v9 = edge0
	v10 = Vec2{
		X: v9.Y,
		Y: -v9.X,
	}
	goto _21
_21:
	smoothParams.Normal0 = v10
	v2 = edge0
	v3 = edge1
	v24 = float32(v2.X*v3.Y) - float32(v2.Y*v3.X)
	goto _25
_25:
	smoothParams.Convex1 = boolUint8(v24 >= convexTol)
	v2 = (*ChainSegment)(unsafe.Pointer(segmentA)).Ghost2
	v3 = p2
	v5 = Vec2{
		X: v2.X - v3.X,
		Y: v2.Y - v3.Y,
	}
	goto _29
_29:
	v6 = v5
	lengthSquared = float32(v6.X*v6.X) + float32(v6.Y*v6.Y)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
		u = Vec2{
			X: float32(s * v6.X),
			Y: float32(s * v6.Y),
		}
		v7 = u
		goto _32
	}
	v7 = Vec2{}
	goto _32
_32:
	edge2 = v7
	v9 = edge2
	v10 = Vec2{
		X: v9.Y,
		Y: -v9.X,
	}
	goto _35
_35:
	smoothParams.Normal2 = v10
	v2 = edge1
	v3 = edge2
	v24 = float32(v2.X*v3.Y) - float32(v2.Y*v3.X)
	goto _39
_39:
	smoothParams.Convex2 = boolUint8(v24 >= convexTol)
	v2 = edge1
	v3 = Vec2{
		X: v2.Y,
		Y: -v2.X,
	}
	goto _42
_42:
	// Normal points to the right
	normal1 = v3
	v5 = centroidB
	v6 = p11
	v7 = Vec2{
		X: v5.X - v6.X,
		Y: v5.Y - v6.Y,
	}
	goto _46
_46:
	v9 = normal1
	v10 = v7
	v24 = float32(v9.X*v10.X) + float32(v9.Y*v10.Y)
	goto _50
_50:
	behind1 = boolUint8(v24 < float32FromFloat32(0))
	behind0 = boolUint8(true1 != 0)
	behind2 = boolUint8(true1 != 0)
	if smoothParams.Convex1 != 0 {
		v12 = centroidB
		v13 = p11
		v14 = Vec2{
			X: v12.X - v13.X,
			Y: v12.Y - v13.Y,
		}
		goto _54
	_54:
		v16 = smoothParams.Normal0
		v17 = v14
		v38 = float32(v16.X*v17.X) + float32(v16.Y*v17.Y)
		goto _58
	_58:
		behind0 = boolUint8(v38 < float32FromFloat32(0))
	}
	if smoothParams.Convex2 != 0 {
		v2 = centroidB
		v3 = p2
		v5 = Vec2{
			X: v2.X - v3.X,
			Y: v2.Y - v3.Y,
		}
		goto _62
	_62:
		v6 = smoothParams.Normal2
		v7 = v5
		v24 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y)
		goto _66
	_66:
		behind2 = boolUint8(v24 < float32FromFloat32(0))
	}
	if behind1 != 0 && behind0 != 0 && behind2 != 0 {
		// one-sided collision
		return *(*LocalManifold)(unsafe.Add(nil, bp))
	}
	// Get polygonB in frameA
	count = (*Polygon)(unsafe.Pointer(polygonB)).Count
	i = 0
	for {
		if !(i < count) {
			break
		}
		v1 = xf
		v2 = *(*Vec2)(unsafe.Add(nil, polygonB+uintptr(i)*8))
		x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
		y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
		v3 = Vec2{
			X: x,
			Y: y,
		}
		goto _71
	_71:
		(*(*[8]Vec2)(unsafe.Add(nil, bp+44)))[i] = v3
		v72 = xf.Q
		v2 = *(*Vec2)(unsafe.Add(nil, polygonB+64+uintptr(i)*8))
		v3 = Vec2{
			X: float32(v72.C*v2.X) - float32(v72.S*v2.Y),
			Y: float32(v72.S*v2.X) + float32(v72.C*v2.Y),
		}
		goto _75
	_75:
		normals[i] = v3
		goto _67
	_67:
		;
		i = i + 1
	}
	(*(*DistanceInput)(unsafe.Add(nil, bp+108))).ProxyA = b2MakeProxy(tls, segmentA+8, int32(2), float32FromFloat32(0))
	(*(*DistanceInput)(unsafe.Add(nil, bp+108))).ProxyB = b2MakeProxy(tls, bp+44, count, float32FromFloat32(0))
	(*(*DistanceInput)(unsafe.Add(nil, bp+108))).Transform = b2Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp+108))).UseRadii = boolUint8(false1 != 0)
	output = b2ComputeShapeDistance(tls, bp+108, cache, uintptrFromInt32(0), 0)
	if output.Distance > radiusB+float32(float32FromFloat32(4)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
		return *(*LocalManifold)(unsafe.Add(nil, bp))
	}
	if smoothParams.Convex1 != 0 {
		v2 = smoothParams.Normal0
	} else {
		v2 = normal1
	}
	// Snap concave normals for partial polygon
	n0 = v2
	if smoothParams.Convex2 != 0 {
		v3 = smoothParams.Normal2
	} else {
		v3 = normal1
	}
	n2 = v3
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
			v5 = pB
			v6 = pA
			v7 = Vec2{
				X: v5.X - v6.X,
				Y: v5.Y - v6.Y,
			}
			goto _81
		_81:
			v9 = v7
			lengthSquared = float32(v9.X*v9.X) + float32(v9.Y*v9.Y)
			if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
				s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
				u = Vec2{
					X: float32(s * v9.X),
					Y: float32(s * v9.Y),
				}
				v10 = u
				goto _84
			}
			v10 = Vec2{}
			goto _84
		_84:
			normal = v10
			type1 = b2ClassifyNormal(tls, smoothParams, normal)
			if type1 == int32(b2_normalSkip) {
				return *(*LocalManifold)(unsafe.Add(nil, bp))
			}
			if type1 == int32(b2_normalAdmit) {
				(*(*LocalManifold)(unsafe.Add(nil, bp))).Normal = normal
				cp = bp + 8 + uintptr(0)*16
				(*LocalManifoldPoint)(unsafe.Pointer(cp)).Point = pA
				(*LocalManifoldPoint)(unsafe.Pointer(cp)).Separation = output.Distance - radiusB
				(*LocalManifoldPoint)(unsafe.Pointer(cp)).Id = uint16FromInt32(int32FromUint8(*(*uint8_t)(unsafe.Add(nil, cache+2)))<<int32FromInt32(8) | int32FromUint8(*(*uint8_t)(unsafe.Add(nil, cache+5))))
				(*(*LocalManifold)(unsafe.Add(nil, bp))).PointCount = int32(1)
				return *(*LocalManifold)(unsafe.Add(nil, bp))
			}
			// fall through b2_normalSnap
			incidentIndex = int32FromUint8(*(*uint8_t)(unsafe.Add(nil, cache+5)))
		} else {
			// vertex-edge collision
			if v85 = !!(int32FromUint16((*SimplexCache)(unsafe.Pointer(cache)).Count) == int32FromInt32(2)); !v85 {
				b2InternalAssert(tls, __ccgo_ts+15921, __ccgo_ts+15836, int32FromInt32(1450))
			}
			_ = v85 || bool(int32FromInt32(0) != 0)
			ia1 = int32FromUint8(*(*uint8_t)(unsafe.Add(nil, cache+2)))
			ia2 = int32FromUint8(*(*uint8_t)(unsafe.Add(nil, cache+2+1)))
			ib1 = int32FromUint8(*(*uint8_t)(unsafe.Add(nil, cache+5)))
			ib2 = int32FromUint8(*(*uint8_t)(unsafe.Add(nil, cache+5+1)))
			if ia1 == ia2 {
				// 1 point on A, expect 2 points on B
				if v85 = !!(ib1 != ib2); !v85 {
					b2InternalAssert(tls, __ccgo_ts+15939, __ccgo_ts+15836, int32FromInt32(1460))
				}
				_ = v85 || bool(int32FromInt32(0) != 0)
				v2 = output.PointA
				v3 = output.PointB
				v5 = Vec2{
					X: v2.X - v3.X,
					Y: v2.Y - v3.Y,
				}
				goto _90
			_90:
				// Find polygon normal most aligned with vector between closest points.
				// This effectively sorts ib1 and ib2
				normalB = v5
				v6 = normalB
				v7 = normals[ib1]
				v24 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y)
				goto _94
			_94:
				dot1 = v24
				v9 = normalB
				v10 = normals[ib2]
				v38 = float32(v9.X*v10.X) + float32(v9.Y*v10.Y)
				goto _98
			_98:
				dot2 = v38
				if dot1 > dot2 {
					v99 = ib1
				} else {
					v99 = ib2
				}
				ib = v99
				// Use accurate normal
				normalB = normals[ib]
				v2 = normalB
				v3 = Vec2{
					X: -v2.X,
					Y: -v2.Y,
				}
				goto _102
			_102:
				type11 = b2ClassifyNormal(tls, smoothParams, v3)
				if type11 == int32(b2_normalSkip) {
					return *(*LocalManifold)(unsafe.Add(nil, bp))
				}
				if type11 == int32(b2_normalAdmit) {
					// Get polygon edge associated with normal
					ib1 = ib
					if ib < count-int32(1) {
						v99 = ib + int32(1)
					} else {
						v99 = 0
					}
					ib2 = v99
					b11 = (*(*[8]Vec2)(unsafe.Add(nil, bp+44)))[ib1]
					b21 = (*(*[8]Vec2)(unsafe.Add(nil, bp+44)))[ib2]
					// Find incident segment vertex
					v2 = p11
					v3 = b11
					v5 = Vec2{
						X: v2.X - v3.X,
						Y: v2.Y - v3.Y,
					}
					goto _107
				_107:
					v6 = normalB
					v7 = v5
					v24 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y)
					goto _111
				_111:
					dot1 = v24
					v2 = p2
					v3 = b11
					v5 = Vec2{
						X: v2.X - v3.X,
						Y: v2.Y - v3.Y,
					}
					goto _115
				_115:
					v6 = normalB
					v7 = v5
					v24 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y)
					goto _119
				_119:
					dot2 = v24
					if dot1 < dot2 {
						v2 = n0
						v3 = normalB
						v24 = float32(v2.X*v3.X) + float32(v2.Y*v3.Y)
						goto _123
					_123:
						v5 = normal1
						v6 = normalB
						v38 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y)
						goto _127
					_127:
						if v24 < v38 {
							// Neighbor is incident
							return *(*LocalManifold)(unsafe.Add(nil, bp))
						}
					} else {
						v2 = n2
						v3 = normalB
						v24 = float32(v2.X*v3.X) + float32(v2.Y*v3.Y)
						goto _131
					_131:
						v5 = normal1
						v6 = normalB
						v38 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y)
						goto _135
					_135:
						if v24 < v38 {
							// Neighbor is incident
							return *(*LocalManifold)(unsafe.Add(nil, bp))
						}
					}
					*(*LocalManifold)(unsafe.Add(nil, bp)) = b2ClipSegments(tls, b11, b21, p11, p2, normalB, radiusB, float32FromFloat32(0), uint16FromInt32(int32FromUint8(uint8FromInt32(ib1))<<int32FromInt32(8)|int32FromUint8(uint8FromInt32(int32FromInt32(1)))), uint16FromInt32(int32FromUint8(uint8FromInt32(ib2))<<int32FromInt32(8)|int32FromUint8(uint8FromInt32(int32FromInt32(0)))))
					if v85 = !!((*(*LocalManifold)(unsafe.Add(nil, bp))).PointCount == 0 || (*(*LocalManifold)(unsafe.Add(nil, bp))).PointCount == int32(2)); !v85 {
						b2InternalAssert(tls, __ccgo_ts+15950, __ccgo_ts+15836, int32FromInt32(1511))
					}
					_ = v85 || bool(int32FromInt32(0) != 0)
					if (*(*LocalManifold)(unsafe.Add(nil, bp))).PointCount == int32(2) {
						v2 = normalB
						v3 = Vec2{
							X: -v2.X,
							Y: -v2.Y,
						}
						goto _139
					_139:
						(*(*LocalManifold)(unsafe.Add(nil, bp))).Normal = v3
					}
					return *(*LocalManifold)(unsafe.Add(nil, bp))
				}
				// fall through b2_normalSnap
				incidentNormal = ib
			} else {
				v2 = (*(*[8]Vec2)(unsafe.Add(nil, bp+44)))[ib1]
				v3 = p11
				v5 = Vec2{
					X: v2.X - v3.X,
					Y: v2.Y - v3.Y,
				}
				goto _143
			_143:
				v6 = normal1
				v7 = v5
				v24 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y)
				goto _147
			_147:
				// Get index of incident polygonB vertex
				dot11 = v24
				v9 = (*(*[8]Vec2)(unsafe.Add(nil, bp+44)))[ib2]
				v10 = p2
				v12 = Vec2{
					X: v9.X - v10.X,
					Y: v9.Y - v10.Y,
				}
				goto _151
			_151:
				v13 = normal1
				v14 = v12
				v38 = float32(v13.X*v14.X) + float32(v13.Y*v14.Y)
				goto _155
			_155:
				dot21 = v38
				if dot11 < dot21 {
					v99 = ib1
				} else {
					v99 = ib2
				}
				incidentIndex = v99
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
			v2 = (*(*[8]Vec2)(unsafe.Add(nil, bp+44)))[i1]
			v3 = p11
			v5 = Vec2{
				X: v2.X - v3.X,
				Y: v2.Y - v3.Y,
			}
			goto _161
		_161:
			v6 = normal1
			v7 = v5
			v24 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y)
			goto _165
		_165:
			s1 = v24
			if s1 < edgeSeparation {
				edgeSeparation = s1
				incidentIndex = i1
			}
			goto _157
		_157:
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
				v2 = (*(*[8]Vec2)(unsafe.Add(nil, bp+44)))[i2]
				v3 = p11
				v5 = Vec2{
					X: v2.X - v3.X,
					Y: v2.Y - v3.Y,
				}
				goto _170
			_170:
				v6 = smoothParams.Normal0
				v7 = v5
				v24 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y)
				goto _174
			_174:
				s2 = v24
				if s2 < s0 {
					s0 = s2
				}
				goto _166
			_166:
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
				v2 = (*(*[8]Vec2)(unsafe.Add(nil, bp+44)))[i3]
				v3 = p2
				v5 = Vec2{
					X: v2.X - v3.X,
					Y: v2.Y - v3.Y,
				}
				goto _179
			_179:
				v6 = smoothParams.Normal2
				v7 = v5
				v24 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y)
				goto _183
			_183:
				s3 = v24
				if s3 < s21 {
					s21 = s3
				}
				goto _175
			_175:
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
			n = normals[i4]
			v2 = n
			v3 = Vec2{
				X: -v2.X,
				Y: -v2.Y,
			}
			goto _187
		_187:
			type2 = b2ClassifyNormal(tls, smoothParams, v3)
			if type2 != int32(b2_normalAdmit) {
				goto _184
			}
			// Check the infinite sides of the partial polygon
			// if ((smoothParams.convex1 && b2Cross(n0, n) > 0.0f) || (smoothParams.convex2 && b2Cross(n, n2) > 0.0f))
			//{
			//	continue;
			//}
			p1 = (*(*[8]Vec2)(unsafe.Add(nil, bp+44)))[i4]
			v2 = p2
			v3 = p1
			v5 = Vec2{
				X: v2.X - v3.X,
				Y: v2.Y - v3.Y,
			}
			goto _191
		_191:
			v6 = n
			v7 = v5
			v24 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y)
			goto _195
		_195:
			v9 = p11
			v10 = p1
			v12 = Vec2{
				X: v9.X - v10.X,
				Y: v9.Y - v10.Y,
			}
			goto _199
		_199:
			v13 = n
			v14 = v12
			v38 = float32(v13.X*v14.X) + float32(v13.Y*v14.Y)
			goto _203
		_203:
			v49 = v24
			v57 = v38
			if v49 < v57 {
				v93 = v49
			} else {
				v93 = v57
			}
			v65 = v93
			goto _207
		_207:
			s4 = v65
			if s4 > polygonSeparation {
				polygonSeparation = s4
				referenceIndex = i4
			}
			goto _184
		_184:
			;
			i4 = i4 + 1
		}
		if polygonSeparation > edgeSeparation {
			ia11 = referenceIndex
			if ia11 < count-int32(1) {
				v99 = ia11 + int32(1)
			} else {
				v99 = 0
			}
			ia21 = v99
			a11 = (*(*[8]Vec2)(unsafe.Add(nil, bp+44)))[ia11]
			a21 = (*(*[8]Vec2)(unsafe.Add(nil, bp+44)))[ia21]
			n1 = normals[ia11]
			v2 = p11
			v3 = a11
			v5 = Vec2{
				X: v2.X - v3.X,
				Y: v2.Y - v3.Y,
			}
			goto _213
		_213:
			v6 = n1
			v7 = v5
			v24 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y)
			goto _217
		_217:
			dot12 = v24
			v9 = p2
			v10 = a11
			v12 = Vec2{
				X: v9.X - v10.X,
				Y: v9.Y - v10.Y,
			}
			goto _221
		_221:
			v13 = n1
			v14 = v12
			v38 = float32(v13.X*v14.X) + float32(v13.Y*v14.Y)
			goto _225
		_225:
			dot22 = v38
			if dot12 < dot22 {
				v16 = n0
				v17 = n1
				v49 = float32(v16.X*v17.X) + float32(v16.Y*v17.Y)
				goto _229
			_229:
				v19 = normal1
				v20 = n1
				v57 = float32(v19.X*v20.X) + float32(v19.Y*v20.Y)
				goto _233
			_233:
				if v49 < v57 {
					// Neighbor is incident
					return *(*LocalManifold)(unsafe.Add(nil, bp))
				}
			} else {
				v2 = n2
				v3 = n1
				v24 = float32(v2.X*v3.X) + float32(v2.Y*v3.Y)
				goto _237
			_237:
				v5 = normal1
				v6 = n1
				v38 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y)
				goto _241
			_241:
				if v24 < v38 {
					// Neighbor is incident
					return *(*LocalManifold)(unsafe.Add(nil, bp))
				}
			}
			*(*LocalManifold)(unsafe.Add(nil, bp)) = b2ClipSegments(tls, a11, a21, p11, p2, normals[ia11], radiusB, float32FromFloat32(0), uint16FromInt32(int32FromUint8(uint8FromInt32(ia11))<<int32FromInt32(8)|int32FromUint8(uint8FromInt32(int32FromInt32(1)))), uint16FromInt32(int32FromUint8(uint8FromInt32(ia21))<<int32FromInt32(8)|int32FromUint8(uint8FromInt32(int32FromInt32(0)))))
			if v85 = !!((*(*LocalManifold)(unsafe.Add(nil, bp))).PointCount == 0 || (*(*LocalManifold)(unsafe.Add(nil, bp))).PointCount == int32(2)); !v85 {
				b2InternalAssert(tls, __ccgo_ts+15950, __ccgo_ts+15836, int32FromInt32(1653))
			}
			_ = v85 || bool(int32FromInt32(0) != 0)
			if (*(*LocalManifold)(unsafe.Add(nil, bp))).PointCount == int32(2) {
				v2 = normals[ia11]
				v3 = Vec2{
					X: -v2.X,
					Y: -v2.Y,
				}
				goto _245
			_245:
				(*(*LocalManifold)(unsafe.Add(nil, bp))).Normal = v3
			}
			return *(*LocalManifold)(unsafe.Add(nil, bp))
		}
		if incidentIndex == -int32(1) {
			// neighboring segment is the separating axis
			return *(*LocalManifold)(unsafe.Add(nil, bp))
		}
		// fall through segment normal axis
	}
	if v85 = !!(incidentNormal != -int32(1) || incidentIndex != -int32(1)); !v85 {
		b2InternalAssert(tls, __ccgo_ts+16003, __ccgo_ts+15836, int32FromInt32(1671))
	}
	_ = v85 || bool(int32FromInt32(0) != 0)
	if incidentNormal != -int32(1) {
		ib11 = incidentNormal
		if ib11 < count-int32(1) {
			v99 = ib11 + int32(1)
		} else {
			v99 = 0
		}
		ib21 = v99
		b12 = (*(*[8]Vec2)(unsafe.Add(nil, bp+44)))[ib11]
		b22 = (*(*[8]Vec2)(unsafe.Add(nil, bp+44)))[ib21]
	} else {
		i21 = incidentIndex
		if i21 > 0 {
			v99 = i21 - int32(1)
		} else {
			v99 = count - int32(1)
		}
		i11 = v99
		v2 = normal1
		v3 = normals[i11]
		v24 = float32(v2.X*v3.X) + float32(v2.Y*v3.Y)
		goto _252
	_252:
		d1 = v24
		v5 = normal1
		v6 = normals[i21]
		v38 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y)
		goto _256
	_256:
		d2 = v38
		if d1 < d2 {
			ib11 = i11
			ib21 = i21
			b12 = (*(*[8]Vec2)(unsafe.Add(nil, bp+44)))[ib11]
			b22 = (*(*[8]Vec2)(unsafe.Add(nil, bp+44)))[ib21]
		} else {
			ib11 = i21
			if i21 < count-int32(1) {
				v99 = i21 + int32(1)
			} else {
				v99 = 0
			}
			ib21 = v99
			b12 = (*(*[8]Vec2)(unsafe.Add(nil, bp+44)))[ib11]
			b22 = (*(*[8]Vec2)(unsafe.Add(nil, bp+44)))[ib21]
		}
	}
	*(*LocalManifold)(unsafe.Add(nil, bp)) = b2ClipSegments(tls, p11, p2, b12, b22, normal1, float32FromFloat32(0), radiusB, uint16FromInt32(int32FromUint8(uint8FromInt32(int32FromInt32(0)))<<int32FromInt32(8)|int32FromUint8(uint8FromInt32(ib21))), uint16FromInt32(int32FromUint8(uint8FromInt32(int32FromInt32(1)))<<int32FromInt32(8)|int32FromUint8(uint8FromInt32(ib11))))
	if v85 = !!((*(*LocalManifold)(unsafe.Add(nil, bp))).PointCount == 0 || (*(*LocalManifold)(unsafe.Add(nil, bp))).PointCount == int32(2)); !v85 {
		b2InternalAssert(tls, __ccgo_ts+15950, __ccgo_ts+15836, int32FromInt32(1708))
	}
	_ = v85 || bool(int32FromInt32(0) != 0)
	return *(*LocalManifold)(unsafe.Add(nil, bp))
}

func b2CollideTask(tls *_Stack, startIndex int32, endIndex int32, workerIndex int32, context uintptr) {
	var C, xf, xfc, v25 Transform
	var awakeSims, bodySimA, bodySimB, contactSim, fatAABBs, mp, mp1, shapeA, shapeB, shapes, solverSets, spanBase, spans, staticSims, stepContext, taskContext, world, v8, v9 uintptr
	var blockIndex, v13 uint32_t
	var cachedQA, cachedQB, dqA, dqB, qr, qr1, r1, v17, v18, v19, v21, v47, v48, v49, v57, v58, v60, v61, v62, v64, v65, v67, v68, v69, v76, v80 Rot
	var centerOffsetA, centerOffsetB, d, dc, dp, normal, rA, rB, v22, v23, v43, v44, v73, v77, v78, v81, v82, v84, v85, v86, v88, v89, v90, v92, v93 Vec2
	var contactId, contactIndex, encodedA, encodedB, i, i1, shapeIdA, shapeIdB, spanEnd, spanIndex, spanStart int32
	var cosA, cosB, distance, dx, dy, maxExtent, maxExtentA, maxExtentB, minCos, recycleDistance, recycleDistanceNonTouching, speculativeDistance, tolerance, v2, v27, v28, v29, v3, v31, v32, v35, v38, v39, v4, v40, v42, v45, v51, v6 float32
	var overlap, touching, wasTouching, v10, v33 uint8
	var transformA, transformB, v15, v16 WorldTransform
	var v1 bool
	var v71, v72 Pos
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, awakeSims, blockIndex, bodySimA, bodySimB, cachedQA, cachedQB, centerOffsetA, centerOffsetB, contactId, contactIndex, contactSim, cosA, cosB, d, dc, distance, dp, dqA, dqB, dx, dy, encodedA, encodedB, fatAABBs, i, i1, maxExtent, maxExtentA, maxExtentB, minCos, mp, mp1, normal, overlap, qr, qr1, r1, rA, rB, recycleDistance, recycleDistanceNonTouching, shapeA, shapeB, shapeIdA, shapeIdB, shapes, solverSets, spanBase, spanEnd, spanIndex, spanStart, spans, speculativeDistance, staticSims, stepContext, taskContext, tolerance, touching, transformA, transformB, wasTouching, world, xf, xfc, v1, v10, v13, v15, v16, v17, v18, v19, v2, v21, v22, v23, v25, v27, v28, v29, v3, v31, v32, v33, v35, v38, v39, v4, v40, v42, v43, v44, v45, v47, v48, v49, v51, v57, v58, v6, v60, v61, v62, v64, v65, v67, v68, v69, v71, v72, v73, v76, v77, v78, v8, v80, v81, v82, v84, v85, v86, v88, v89, v9, v90, v92, v93
	stepContext = context
	world = (*b2StepContext)(unsafe.Pointer(stepContext)).World
	taskContext = (*b2World)(unsafe.Pointer(world)).TaskContexts.Data + uintptr(workerIndex)*136
	spans = (*b2StepContext)(unsafe.Pointer(stepContext)).CollideSpans
	shapes = (*b2World)(unsafe.Pointer(world)).Shapes.Data
	fatAABBs = (*b2World)(unsafe.Pointer(world)).FatAABBs.Data
	solverSets = (*b2World)(unsafe.Pointer(world)).SolverSets.Data
	awakeSims = (*(*b2SolverSet)(unsafe.Add(nil, solverSets+uintptr(b2_awakeSet)*88))).BodySims.Data
	staticSims = (*(*b2SolverSet)(unsafe.Add(nil, solverSets+uintptr(b2_staticSet)*88))).BodySims.Data
	if v1 = !!(startIndex < endIndex); !v1 {
		b2InternalAssert(tls, __ccgo_ts+17428, __ccgo_ts+16660, int32FromInt32(467))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	spanIndex = 0
	for (*(*b2ContactCollideSpan)(unsafe.Add(nil, spans+uintptr(spanIndex+int32(1))*16))).Start <= startIndex {
		spanIndex = spanIndex + int32(1)
	}
	spanStart = (*(*b2ContactCollideSpan)(unsafe.Add(nil, spans+uintptr(spanIndex)*16))).Start
	spanEnd = (*(*b2ContactCollideSpan)(unsafe.Add(nil, spans+uintptr(spanIndex+int32(1))*16))).Start
	spanBase = (*(*b2ContactCollideSpan)(unsafe.Add(nil, spans+uintptr(spanIndex)*16))).Contacts
	recycleDistance = (*b2World)(unsafe.Pointer(world)).ContactRecycleDistance
	speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter))
	v2 = recycleDistance
	v3 = speculativeDistance
	if v2 < v3 {
		v6 = v2
	} else {
		v6 = v3
	}
	v4 = v6
	goto _5
_5:
	recycleDistanceNonTouching = v4
	contactIndex = startIndex
	for {
		if !(contactIndex < endIndex) {
			break
		}
		if contactIndex == spanEnd {
			spanIndex = spanIndex + int32(1)
			spanStart = (*(*b2ContactCollideSpan)(unsafe.Add(nil, spans+uintptr(spanIndex)*16))).Start
			spanEnd = (*(*b2ContactCollideSpan)(unsafe.Add(nil, spans+uintptr(spanIndex+int32(1))*16))).Start
			spanBase = (*(*b2ContactCollideSpan)(unsafe.Add(nil, spans+uintptr(spanIndex)*16))).Contacts
		}
		contactSim = spanBase + uintptr(contactIndex-spanStart)*200
		contactId = (*b2ContactSim)(unsafe.Pointer(contactSim)).ContactId
		shapeIdA = (*b2ContactSim)(unsafe.Pointer(contactSim)).ShapeIdA
		shapeIdB = (*b2ContactSim)(unsafe.Pointer(contactSim)).ShapeIdB
		v8 = fatAABBs + uintptr(shapeIdA)*16
		v9 = fatAABBs + uintptr(shapeIdB)*16
		v10 = boolUint8((*AABB)(unsafe.Pointer(v8)).LowerBound.X <= (*AABB)(unsafe.Pointer(v9)).UpperBound.X && (*AABB)(unsafe.Pointer(v8)).LowerBound.Y <= (*AABB)(unsafe.Pointer(v9)).UpperBound.Y && (*AABB)(unsafe.Pointer(v9)).LowerBound.X <= (*AABB)(unsafe.Pointer(v8)).UpperBound.X && (*AABB)(unsafe.Pointer(v9)).LowerBound.Y <= (*AABB)(unsafe.Pointer(v8)).UpperBound.Y)
		goto _11
	_11:
		// Do proxies still overlap?
		overlap = v10
		if int32FromUint8(overlap) == false1 {
			*(*uint32_t)(unsafe.Add(nil, contactSim+44)) |= uint32(b2_simDisjoint)
			*(*uint32_t)(unsafe.Add(nil, contactSim+44)) &= uint32FromInt32(^int32(b2_simTouchingFlag))
			v8 = taskContext + 32
			v13 = uint32FromInt32(contactId)
			blockIndex = v13 / uint32(64)
			if v1 = !!(blockIndex < (*b2BitSet)(unsafe.Pointer(v8)).BlockCount); !v1 {
				b2InternalAssert(tls, __ccgo_ts+7356, __ccgo_ts+7388, int32FromInt32(29))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v8)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v13 % uint32FromInt32(64))
		} else {
			shapeA = shapes + uintptr(shapeIdA)*280
			shapeB = shapes + uintptr(shapeIdB)*280
			wasTouching = uint8(boolUint32((*b2ContactSim)(unsafe.Pointer(contactSim)).SimFlags&uint32(b2_simTouchingFlag) != 0))
			// Update contact respecting shape/body order (A,B)
			encodedA = (*b2ContactSim)(unsafe.Pointer(contactSim)).EncodedBodySimA
			encodedB = (*b2ContactSim)(unsafe.Pointer(contactSim)).EncodedBodySimB
			bodySimA = b2ResolveContactBodySim(tls, world, awakeSims, staticSims, encodedA, (*b2Shape)(unsafe.Pointer(shapeA)).BodyId)
			bodySimB = b2ResolveContactBodySim(tls, world, awakeSims, staticSims, encodedB, (*b2Shape)(unsafe.Pointer(shapeB)).BodyId)
			transformA = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform
			transformB = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform
			// These may not be skipped by relative transform check below
			(*b2ContactSim)(unsafe.Pointer(contactSim)).InvMassA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvMass
			(*b2ContactSim)(unsafe.Pointer(contactSim)).InvIA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvInertia
			(*b2ContactSim)(unsafe.Pointer(contactSim)).InvMassB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvMass
			(*b2ContactSim)(unsafe.Pointer(contactSim)).InvIB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvInertia
			// Contact recycling optimization. Please cite this code if you use this optimization.
			// This is inspired by persistent contact manifolds used in some physics engines, such as PhysX.
			// However, this allows larger relative motion and has fewer tuning parameters (just one).
			if recycleDistance > float32FromFloat32(0) && (*b2ContactSim)(unsafe.Pointer(contactSim)).SimFlags&uint32(b2_simRelativeTransformValid) != 0 && (*b2ContactSim)(unsafe.Pointer(contactSim)).SimFlags&uint32(b2_contactRecycleFlag) != 0 {
				cachedQA = (*b2ContactSim)(unsafe.Pointer(contactSim)).CachedRotationA
				cachedQB = (*b2ContactSim)(unsafe.Pointer(contactSim)).CachedRotationB
				xfc = (*b2ContactSim)(unsafe.Pointer(contactSim)).CachedRelativePose
				v15 = transformA
				v16 = transformB
				v17 = v15.Q
				v18 = v16.Q
				r1.S = float32(v17.C*v18.S) - float32(v17.S*v18.C)
				r1.C = float32(v17.C*v18.C) + float32(v17.S*v18.S)
				v19 = r1
				goto _20
			_20:
				C.Q = v19
				d = Vec2{
					X: v16.P.X - v15.P.X,
					Y: v16.P.Y - v15.P.Y,
				}
				v21 = v15.Q
				v22 = d
				v23 = Vec2{
					X: float32(v21.C*v22.X) + float32(v21.S*v22.Y),
					Y: float32(-v21.S*v22.X) + float32(v21.C*v22.Y),
				}
				goto _24
			_24:
				C.P = v23
				v25 = C
				goto _26
			_26:
				xf = v25
				cosA = b2RelativeCos(tls, transformA.Q, cachedQA)
				cosB = b2RelativeCos(tls, transformB.Q, cachedQB)
				v2 = cosA
				v3 = cosB
				if v2 < v3 {
					v6 = v2
				} else {
					v6 = v3
				}
				v4 = v6
				goto _30
			_30:
				minCos = v4
				v10 = boolUint8(encodedA < -int32(1))
				goto _34
			_34:
				if v10 != 0 {
					v27 = float32FromFloat32(0)
				} else {
					v27 = (*b2BodySim)(unsafe.Pointer(bodySimA)).MaxExtent
				}
				maxExtentA = v27
				v33 = boolUint8(encodedB < -int32(1))
				goto _37
			_37:
				if v33 != 0 {
					v28 = float32FromFloat32(0)
				} else {
					v28 = (*b2BodySim)(unsafe.Pointer(bodySimB)).MaxExtent
				}
				maxExtentB = v28
				v29 = maxExtentA
				v31 = maxExtentB
				if v29 > v31 {
					v35 = v29
				} else {
					v35 = v31
				}
				v32 = v35
				goto _41
			_41:
				maxExtent = v32
				v43 = xf.P
				v44 = xfc.P
				dx = v44.X - v43.X
				dy = v44.Y - v43.Y
				v38 = sqrtf(tls, float32(dx*dx)+float32(dy*dy))
				goto _46
			_46:
				distance = v38
				v47 = xf.Q
				v48 = xfc.Q
				r1.S = float32(v47.C*v48.S) - float32(v47.S*v48.C)
				r1.C = float32(v47.C*v48.C) + float32(v47.S*v48.S)
				v49 = r1
				goto _50
			_50:
				qr1 = v49
				if wasTouching != 0 {
					v39 = recycleDistance
				} else {
					v39 = recycleDistanceNonTouching
				}
				// This metric is used for fast bodies and sleeping. It comes from conservative advancement.
				// Note that qr.s == sin(theta) ~= theta for small angles.
				// Need a tighter tolerance for non-touching shapes so that contacts are not missed.
				tolerance = v39
				if v1 = minCos > float32FromFloat32(0.98); v1 {
					v40 = qr1.S
					if v40 < float32FromInt32(0) {
						v45 = -v40
					} else {
						v45 = v40
					}
					v42 = v45
					goto _54
				_54:
				}
				if v1 && distance+float32(maxExtent*v42) < tolerance {
					v57 = cachedQA
					v58 = Rot{
						C: v57.C,
						S: -v57.S,
					}
					goto _59
				_59:
					v60 = transformA.Q
					v61 = v58
					qr.S = float32(v60.S*v61.C) + float32(v60.C*v61.S)
					qr.C = float32(v60.C*v61.C) - float32(v60.S*v61.S)
					v62 = qr
					goto _63
				_63:
					dqA = v62
					v64 = cachedQB
					v65 = Rot{
						C: v64.C,
						S: -v64.S,
					}
					goto _66
				_66:
					v67 = transformB.Q
					v68 = v65
					qr.S = float32(v67.S*v68.C) + float32(v67.C*v68.S)
					qr.C = float32(v67.C*v68.C) - float32(v67.S*v68.S)
					v69 = qr
					goto _70
				_70:
					dqB = v69
					normal = (*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.Normal
					v71 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Center
					v72 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Center
					v73 = Vec2{
						X: v71.X - v72.X,
						Y: v71.Y - v72.Y,
					}
					goto _74
				_74:
					// Minimize round-off
					dc = v73
					i = 0
					for {
						if !(i < (*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.PointCount) {
							break
						}
						// Keep anchors but update separation, same as sub-stepping. This eliminates jitter.
						mp = contactSim + 72 + 12 + uintptr(i)*44
						v76 = dqA
						v77 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
						v78 = Vec2{
							X: float32(v76.C*v77.X) - float32(v76.S*v77.Y),
							Y: float32(v76.S*v77.X) + float32(v76.C*v77.Y),
						}
						goto _79
					_79:
						rA = v78
						v80 = dqB
						v81 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorB
						v82 = Vec2{
							X: float32(v80.C*v81.X) - float32(v80.S*v81.Y),
							Y: float32(v80.S*v81.X) + float32(v80.C*v81.Y),
						}
						goto _83
					_83:
						rB = v82
						v84 = rB
						v85 = rA
						v86 = Vec2{
							X: v84.X - v85.X,
							Y: v84.Y - v85.Y,
						}
						goto _87
					_87:
						v88 = dc
						v89 = v86
						v90 = Vec2{
							X: v88.X + v89.X,
							Y: v88.Y + v89.Y,
						}
						goto _91
					_91:
						dp = v90
						v92 = dp
						v93 = normal
						v51 = float32(v92.X*v93.X) + float32(v92.Y*v93.Y)
						goto _95
					_95:
						(*ManifoldPoint)(unsafe.Pointer(mp)).Separation = (*ManifoldPoint)(unsafe.Pointer(mp)).BaseSeparation + v51
						(*ManifoldPoint)(unsafe.Pointer(mp)).NormalVelocity = float32FromFloat32(0)
						(*ManifoldPoint)(unsafe.Pointer(mp)).Persisted = boolUint8(true1 != 0)
						goto _75
					_75:
						;
						i = i + 1
					}
					*(*int32)(unsafe.Add(nil, taskContext+128)) += int32(1)
					// Contact is recycled. This also skips updating other aspects of the contact
					// such as material parameters.
					goto _7
				}
			}
			// Caching for contact recycling.
			(*b2ContactSim)(unsafe.Pointer(contactSim)).CachedRotationA = transformA.Q
			(*b2ContactSim)(unsafe.Pointer(contactSim)).CachedRotationB = transformB.Q
			v15 = transformA
			v16 = transformB
			v17 = v15.Q
			v18 = v16.Q
			r1.S = float32(v17.C*v18.S) - float32(v17.S*v18.C)
			r1.C = float32(v17.C*v18.C) + float32(v17.S*v18.S)
			v19 = r1
			goto _101
		_101:
			C.Q = v19
			d = Vec2{
				X: v16.P.X - v15.P.X,
				Y: v16.P.Y - v15.P.Y,
			}
			v21 = v15.Q
			v22 = d
			v23 = Vec2{
				X: float32(v21.C*v22.X) + float32(v21.S*v22.Y),
				Y: float32(-v21.S*v22.X) + float32(v21.C*v22.Y),
			}
			goto _105
		_105:
			C.P = v23
			v25 = C
			goto _107
		_107:
			(*b2ContactSim)(unsafe.Pointer(contactSim)).CachedRelativePose = v25
			*(*uint32_t)(unsafe.Add(nil, contactSim+44)) |= uint32(b2_simRelativeTransformValid)
			v17 = transformA.Q
			v22 = (*b2BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
			v23 = Vec2{
				X: float32(v17.C*v22.X) - float32(v17.S*v22.Y),
				Y: float32(v17.S*v22.X) + float32(v17.C*v22.Y),
			}
			goto _111
		_111:
			centerOffsetA = v23
			v18 = transformB.Q
			v43 = (*b2BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
			v44 = Vec2{
				X: float32(v18.C*v43.X) - float32(v18.S*v43.Y),
				Y: float32(v18.S*v43.X) + float32(v18.C*v43.Y),
			}
			goto _115
		_115:
			centerOffsetB = v44
			// This updates solid contacts
			touching = b2UpdateContact(tls, world, contactSim, shapeA, transformA, centerOffsetA, shapeB, transformB, centerOffsetB)
			// State changes that affect island connectivity. Also affects contact events.
			if int32FromUint8(touching) == int32(true1) && int32FromUint8(wasTouching) == false1 {
				*(*uint32_t)(unsafe.Add(nil, contactSim+44)) |= uint32(b2_simStartedTouching)
				v8 = taskContext + 32
				v13 = uint32FromInt32(contactId)
				blockIndex = v13 / uint32(64)
				if v1 = !!(blockIndex < (*b2BitSet)(unsafe.Pointer(v8)).BlockCount); !v1 {
					b2InternalAssert(tls, __ccgo_ts+7356, __ccgo_ts+7388, int32FromInt32(29))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v8)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v13 % uint32FromInt32(64))
			} else {
				if int32FromUint8(touching) == false1 && int32FromUint8(wasTouching) == int32(true1) {
					*(*uint32_t)(unsafe.Add(nil, contactSim+44)) |= uint32(b2_simStoppedTouching)
					v8 = taskContext + 32
					v13 = uint32FromInt32(contactId)
					blockIndex = v13 / uint32(64)
					if v1 = !!(blockIndex < (*b2BitSet)(unsafe.Pointer(v8)).BlockCount); !v1 {
						b2InternalAssert(tls, __ccgo_ts+7356, __ccgo_ts+7388, int32FromInt32(29))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v8)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v13 % uint32FromInt32(64))
				}
			}
			if touching != 0 {
				i1 = 0
				for {
					if !(i1 < (*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.PointCount) {
						break
					}
					// Cache separation
					mp1 = contactSim + 72 + 12 + uintptr(i1)*44
					(*ManifoldPoint)(unsafe.Pointer(mp1)).BaseSeparation = (*ManifoldPoint)(unsafe.Pointer(mp1)).Separation
					goto _122
				_122:
					;
					i1 = i1 + 1
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
		goto _7
	_7:
		;
		contactIndex = contactIndex + 1
	}
}

// C documentation
//
//	// Narrow-phase collision
func b2Collide(tls *_Stack, context uintptr) {
	bp := tls.Alloc(416)
	defer tls.Free(416)
	var awakeSet, bitSet, color, color1, contact, contactSim, graphColors, shapeA, shapeA1, shapeB, shapeB1, shapes, world uintptr
	var bits uint64_t
	var bodyIdA, bodyIdB, colorIndex, contactCount, contactId, contactIdCapacity, contactIndex, count, elementSize, elementSize1, endEventArrayIndex, i, i1, i2, i3, localIndex, minRange, newCapacity, newCapacity1, newSize, newSize1, nonTouchingCount, oldSize, oldSize1, spanCount, v6 int32
	var contactFullId, contactFullId1 ContactId
	var ctz, flags, k, simFlags, v12 uint32_t
	var event ContactBeginTouchEvent
	var event1 ContactEndTouchEvent
	var shapeIdA, shapeIdA1, shapeIdB, shapeIdB1 ShapeId
	var worldId uint16_t
	var v1, v4, v5 bool
	var _ /* collideSpans at bp+0 */ [26]b2ContactCollideSpan
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = awakeSet, bitSet, bits, bodyIdA, bodyIdB, color, color1, colorIndex, contact, contactCount, contactFullId, contactFullId1, contactId, contactIdCapacity, contactIndex, contactSim, count, ctz, elementSize, elementSize1, endEventArrayIndex, event, event1, flags, graphColors, i, i1, i2, i3, k, localIndex, minRange, newCapacity, newCapacity1, newSize, newSize1, nonTouchingCount, oldSize, oldSize1, shapeA, shapeA1, shapeB, shapeB1, shapeIdA, shapeIdA1, shapeIdB, shapeIdB1, shapes, simFlags, spanCount, world, worldId, v1, v12, v4, v5, v6
	world = (*b2StepContext)(unsafe.Pointer(context)).World
	if v1 = !!((*b2World)(unsafe.Pointer(world)).WorkerCount > int32FromInt32(0)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+17527, __ccgo_ts+16660, int32FromInt32(668))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// gather contacts into a single array for easier parallel-for
	contactCount = 0
	graphColors = world + 376
	i = 0
	for {
		if !(i < int32(_B2_GRAPH_COLOR_COUNT)) {
			break
		}
		contactCount = contactCount + (*(*b2GraphColor)(unsafe.Add(nil, graphColors+uintptr(i)*64))).ContactSims.Count
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
	spanCount = 0
	contactIndex = 0
	i1 = 0
	for {
		if !(i1 < int32(_B2_GRAPH_COLOR_COUNT)) {
			break
		}
		color = graphColors + uintptr(i1)*64
		count = (*b2GraphColor)(unsafe.Pointer(color)).ContactSims.Count
		if count > 0 {
			(*(*[26]b2ContactCollideSpan)(unsafe.Add(nil, bp)))[spanCount].Start = contactIndex
			(*(*[26]b2ContactCollideSpan)(unsafe.Add(nil, bp)))[spanCount].Contacts = (*b2GraphColor)(unsafe.Pointer(color)).ContactSims.Data
			spanCount = spanCount + int32(1)
			contactIndex = contactIndex + count
		}
		goto _3
	_3:
		;
		i1 = i1 + 1
	}
	if nonTouchingCount > 0 {
		(*(*[26]b2ContactCollideSpan)(unsafe.Add(nil, bp)))[spanCount].Start = contactIndex
		(*(*[26]b2ContactCollideSpan)(unsafe.Add(nil, bp)))[spanCount].Contacts = (*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_awakeSet)*88))).ContactSims.Data
		spanCount = spanCount + int32(1)
		contactIndex = contactIndex + nonTouchingCount
	}
	if v1 = !!(contactIndex == contactCount); !v1 {
		b2InternalAssert(tls, __ccgo_ts+17550, __ccgo_ts+16660, int32FromInt32(713))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(spanCount <= int32FromInt32(_B2_GRAPH_COLOR_COUNT)+int32FromInt32(1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+17579, __ccgo_ts+16660, int32FromInt32(714))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	(*(*[26]b2ContactCollideSpan)(unsafe.Add(nil, bp)))[spanCount].Start = contactCount
	(*(*[26]b2ContactCollideSpan)(unsafe.Add(nil, bp)))[spanCount].Contacts = uintptrFromInt32(0)
	(*b2StepContext)(unsafe.Pointer(context)).CollideSpans = bp
	v6 = (*b2IdPool)(unsafe.Pointer(world + 2032)).NextIndex
	goto _7
_7:
	// Contact bit set on ids because contact pointers are unstable as they move between touching and not touching.
	contactIdCapacity = v6
	i2 = 0
	for {
		if !(i2 < (*b2World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		b2SetBitCountAndClear(tls, (*b2World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i2)*136+32, uint32FromInt32(contactIdCapacity))
		(*(*b2TaskContext)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i2)*136))).RecycledContactCount = 0
		goto _8
	_8:
		;
		i2 = i2 + 1
	}
	// Task should take at least 40us on a 4GHz CPU (10K cycles)
	minRange = int32(64)
	b2ParallelFor(tls, world, __ccgo_fp(b2CollideTask), contactCount, minRange, context)
	(*b2StepContext)(unsafe.Pointer(context)).CollideSpans = uintptrFromInt32(0)
	// Serially update contact state
	// Bitwise OR all contact bits
	bitSet = (*b2World)(unsafe.Pointer(world)).TaskContexts.Data + 32
	i3 = int32(1)
	for {
		if !(i3 < (*b2World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		b2InPlaceUnion(tls, bitSet, (*b2World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i3)*136+32)
		goto _9
	_9:
		;
		i3 = i3 + 1
	}
	if v1 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+16660, int32FromInt32(745))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	awakeSet = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
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
			v12 = uint32FromInt32(__builtin_ctzll(tls, bits))
			goto _13
		_13:
			ctz = v12
			contactId = int32FromUint32(uint32FromInt32(64)*k + ctz)
			if v4 = !!(0 <= contactId && contactId < (*b2World)(unsafe.Pointer(world)).Contacts.Count); !v4 {
				b2InternalAssert(tls, __ccgo_ts+301, __ccgo_ts+16660, int32FromInt32(761))
			}
			_ = v4 || bool(int32FromInt32(0) != 0)
			contact = (*b2World)(unsafe.Pointer(world)).Contacts.Data + uintptr(contactId)*64
			if v5 = !!((*b2Contact)(unsafe.Pointer(contact)).SetIndex == int32(b2_awakeSet)); !v5 {
				b2InternalAssert(tls, __ccgo_ts+6749, __ccgo_ts+16660, int32FromInt32(762))
			}
			_ = v5 || bool(int32FromInt32(0) != 0)
			colorIndex = (*b2Contact)(unsafe.Pointer(contact)).ColorIndex
			localIndex = (*b2Contact)(unsafe.Pointer(contact)).LocalIndex
			contactSim = uintptrFromInt32(0)
			if colorIndex != -int32(1) {
				// contact lives in constraint graph
				if v1 = !!(0 <= colorIndex && colorIndex < int32(_B2_GRAPH_COLOR_COUNT)); !v1 {
					b2InternalAssert(tls, __ccgo_ts+5383, __ccgo_ts+16660, int32FromInt32(771))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				color1 = graphColors + uintptr(colorIndex)*64
				if v1 = !!(0 <= localIndex && localIndex < (*b2GraphColor)(unsafe.Pointer(color1)).ContactSims.Count); !v1 {
					b2InternalAssert(tls, __ccgo_ts+17617, __ccgo_ts+16660, int32FromInt32(773))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				contactSim = (*b2GraphColor)(unsafe.Pointer(color1)).ContactSims.Data + uintptr(localIndex)*200
			} else {
				if v1 = !!(0 <= localIndex && localIndex < (*b2SolverSet)(unsafe.Pointer(awakeSet)).ContactSims.Count); !v1 {
					b2InternalAssert(tls, __ccgo_ts+17686, __ccgo_ts+16660, int32FromInt32(777))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				contactSim = (*b2SolverSet)(unsafe.Pointer(awakeSet)).ContactSims.Data + uintptr(localIndex)*200
			}
			flags = (*b2Contact)(unsafe.Pointer(contact)).Flags
			simFlags = (*b2ContactSim)(unsafe.Pointer(contactSim)).SimFlags
			if simFlags&uint32(b2_simDisjoint) != 0 {
				// Bounding boxes no longer overlap
				b2DestroyContact(tls, world, contact)
				contact = uintptrFromInt32(0)
				contactSim = uintptrFromInt32(0)
			} else {
				if simFlags&uint32(b2_simStartedTouching) != 0 {
					if v1 = !!((*b2Contact)(unsafe.Pointer(contact)).IslandId == -int32FromInt32(1)); !v1 {
						b2InternalAssert(tls, __ccgo_ts+10743, __ccgo_ts+16660, int32FromInt32(792))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					if flags&uint32(b2_contactEnableContactEvents) != 0 {
						shapeA = shapes + uintptr((*b2Contact)(unsafe.Pointer(contact)).ShapeIdA)*280
						shapeB = shapes + uintptr((*b2Contact)(unsafe.Pointer(contact)).ShapeIdB)*280
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
						contactFullId = ContactId{
							Index1:     contactId + int32(1),
							World0:     worldId,
							Generation: (*b2Contact)(unsafe.Pointer(contact)).Generation,
						}
						event = ContactBeginTouchEvent{
							ShapeIdA:  shapeIdA,
							ShapeIdB:  shapeIdB,
							ContactId: contactFullId,
						}
						elementSize = int32(28)
						if (*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Count >= (*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Capacity {
							oldSize = (*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Capacity * elementSize
							if (*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Capacity == 0 {
								v6 = int32(8)
							} else {
								v6 = int32(2) * (*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Capacity
							}
							newCapacity = v6
							newSize = newCapacity * elementSize
							(*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Data = b2GrowAlloc(tls, (*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Data, uint64FromInt32(oldSize), uint64FromInt32(newSize))
							(*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Capacity = newCapacity
						}
						*(*ContactBeginTouchEvent)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Data+uintptr((*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Count)*28)) = event
						(*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Count += int32(1)
					}
					if v1 = !!((*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.PointCount > int32FromInt32(0)); !v1 {
						b2InternalAssert(tls, __ccgo_ts+5083, __ccgo_ts+16660, int32FromInt32(811))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					if v1 = !!((*b2Contact)(unsafe.Pointer(contact)).SetIndex == int32(b2_awakeSet)); !v1 {
						b2InternalAssert(tls, __ccgo_ts+6749, __ccgo_ts+16660, int32FromInt32(812))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					// Link first because this wakes colliding bodies and ensures the body sims
					// are in the correct place.
					*(*uint32_t)(unsafe.Add(nil, contact+56)) |= uint32(b2_contactTouchingFlag)
					b2LinkContact(tls, world, contact)
					// Make sure these didn't change
					if v1 = !!((*b2Contact)(unsafe.Pointer(contact)).ColorIndex == -int32FromInt32(1)); !v1 {
						b2InternalAssert(tls, __ccgo_ts+17758, __ccgo_ts+16660, int32FromInt32(820))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					if v1 = !!((*b2Contact)(unsafe.Pointer(contact)).LocalIndex == localIndex); !v1 {
						b2InternalAssert(tls, __ccgo_ts+17795, __ccgo_ts+16660, int32FromInt32(821))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					// Contact sim pointer may have become orphaned due to awake set growth,
					// so I just need to refresh it.
					if v1 = !!(0 <= localIndex && localIndex < (*b2SolverSet)(unsafe.Pointer(awakeSet)).ContactSims.Count); !v1 {
						b2InternalAssert(tls, __ccgo_ts+17686, __ccgo_ts+16660, int32FromInt32(825))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					contactSim = (*b2SolverSet)(unsafe.Pointer(awakeSet)).ContactSims.Data + uintptr(localIndex)*200
					*(*uint32_t)(unsafe.Add(nil, contactSim+44)) &= uint32FromInt32(^int32(b2_simStartedTouching))
					// Add first for memcpy
					b2AddContactToGraph(tls, world, contactSim, contact)
					// This destroys the contact sim
					b2RemoveNonTouchingContact(tls, world, int32(b2_awakeSet), localIndex)
					contactSim = uintptrFromInt32(0)
				} else {
					if simFlags&uint32(b2_simStoppedTouching) != 0 {
						*(*uint32_t)(unsafe.Add(nil, contactSim+44)) &= uint32FromInt32(^int32(b2_simStoppedTouching))
						*(*uint32_t)(unsafe.Add(nil, contact+56)) &= uint32FromInt32(^int32(b2_contactTouchingFlag))
						if (*b2Contact)(unsafe.Pointer(contact)).Flags&uint32(b2_contactEnableContactEvents) != 0 {
							shapeA1 = shapes + uintptr((*b2Contact)(unsafe.Pointer(contact)).ShapeIdA)*280
							shapeB1 = shapes + uintptr((*b2Contact)(unsafe.Pointer(contact)).ShapeIdB)*280
							shapeIdA1 = ShapeId{
								Index1:     (*b2Shape)(unsafe.Pointer(shapeA1)).Id + int32(1),
								World0:     worldId,
								Generation: (*b2Shape)(unsafe.Pointer(shapeA1)).Generation,
							}
							shapeIdB1 = ShapeId{
								Index1:     (*b2Shape)(unsafe.Pointer(shapeB1)).Id + int32(1),
								World0:     worldId,
								Generation: (*b2Shape)(unsafe.Pointer(shapeB1)).Generation,
							}
							contactFullId1 = ContactId{
								Index1:     contactId + int32(1),
								World0:     worldId,
								Generation: (*b2Contact)(unsafe.Pointer(contact)).Generation,
							}
							event1 = ContactEndTouchEvent{
								ShapeIdA:  shapeIdA1,
								ShapeIdB:  shapeIdB1,
								ContactId: contactFullId1,
							}
							elementSize1 = int32(28)
							if (*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336+uintptr(endEventArrayIndex)*16))).Count >= (*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336+uintptr(endEventArrayIndex)*16))).Capacity {
								oldSize1 = (*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336+uintptr(endEventArrayIndex)*16))).Capacity * elementSize1
								if (*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336+uintptr(endEventArrayIndex)*16))).Capacity == 0 {
									v6 = int32(8)
								} else {
									v6 = int32(2) * (*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336+uintptr(endEventArrayIndex)*16))).Capacity
								}
								newCapacity1 = v6
								newSize1 = newCapacity1 * elementSize1
								(*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336+uintptr(endEventArrayIndex)*16))).Data = b2GrowAlloc(tls, (*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336+uintptr(endEventArrayIndex)*16))).Data, uint64FromInt32(oldSize1), uint64FromInt32(newSize1))
								(*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336+uintptr(endEventArrayIndex)*16))).Capacity = newCapacity1
							}
							*(*ContactEndTouchEvent)(unsafe.Add(nil, (*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336+uintptr(endEventArrayIndex)*16))).Data+uintptr((*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336+uintptr(endEventArrayIndex)*16))).Count)*28)) = event1
							(*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336+uintptr(endEventArrayIndex)*16))).Count += int32(1)
						}
						if v1 = !!((*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.PointCount == int32FromInt32(0)); !v1 {
							b2InternalAssert(tls, __ccgo_ts+17829, __ccgo_ts+16660, int32FromInt32(859))
						}
						_ = v1 || bool(int32FromInt32(0) != 0)
						b2UnlinkContact(tls, world, contact)
						bodyIdA = (*(*b2ContactEdge)(unsafe.Add(nil, contact))).BodyId
						bodyIdB = (*(*b2ContactEdge)(unsafe.Add(nil, contact+1*12))).BodyId
						// Add first for memcpy
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
		goto _11
	_11:
		;
		k = k + 1
	}
	b2ValidateSolverSets(tls, world)
	b2ValidateContacts(tls, world)
}

func b2CollideMover(tls *_Stack, mover uintptr, shape uintptr, transform Transform) (r PlaneResult) {
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
		result = b2CollideMoverAndCapsule(tls, bp, shape+128)
	case int32(b2_circleShape):
		result = b2CollideMoverAndCircle(tls, bp, shape+128)
	case int32(b2_polygonShape):
		result = b2CollideMoverAndPolygon(tls, bp, shape+128)
	case int32(b2_segmentShape):
		result = b2CollideMoverAndSegment(tls, bp, shape+128)
	case int32(b2_chainSegmentShape):
		result = b2CollideMoverAndSegment(tls, bp, shape+128+8)
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
