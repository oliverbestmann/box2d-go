package b2

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

func b2ComputeManifold(tls *_Stack, shapeA uintptr, transformA Transform, shapeB uintptr, transformB Transform) (r Manifold) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var fcn uintptr
	var _ /* cache at bp+0 */ SimplexCache
	_ = fcn
	fcn = (*(*b2ContactRegister)(unsafe.Pointer(uintptr(unsafe.Pointer(&s_registers)) + uintptr((*b2Shape)(unsafe.Pointer(shapeA)).Type1)*80 + uintptr((*b2Shape)(unsafe.Pointer(shapeB)).Type1)*16))).Fcn
	*(*SimplexCache)(unsafe.Pointer(bp)) = SimplexCache{}
	return (*(*func(*_Stack, uintptr, Transform, uintptr, Transform, uintptr) Manifold)(unsafe.Pointer(&struct{ uintptr }{fcn})))(tls, shapeA, transformA, shapeB, transformB, bp)
}

func b2ComputeSimplexWitnessPoints(tls *_Stack, a uintptr, b uintptr, s uintptr) {
	switch (*Simplex)(unsafe.Pointer(s)).Count {
	case int32(1):
		*(*Vec2)(unsafe.Pointer(a)) = (*Simplex)(unsafe.Pointer(s)).V1.WA
		*(*Vec2)(unsafe.Pointer(b)) = (*Simplex)(unsafe.Pointer(s)).V1.WB
	case int32(2):
		*(*Vec2)(unsafe.Pointer(a)) = b2Weight2(tls, (*Simplex)(unsafe.Pointer(s)).V1.A, (*Simplex)(unsafe.Pointer(s)).V1.WA, (*Simplex)(unsafe.Pointer(s)).V2.A, (*Simplex)(unsafe.Pointer(s)).V2.WA)
		*(*Vec2)(unsafe.Pointer(b)) = b2Weight2(tls, (*Simplex)(unsafe.Pointer(s)).V1.A, (*Simplex)(unsafe.Pointer(s)).V1.WB, (*Simplex)(unsafe.Pointer(s)).V2.A, (*Simplex)(unsafe.Pointer(s)).V2.WB)
	case int32(3):
		*(*Vec2)(unsafe.Pointer(a)) = b2Weight3(tls, (*Simplex)(unsafe.Pointer(s)).V1.A, (*Simplex)(unsafe.Pointer(s)).V1.WA, (*Simplex)(unsafe.Pointer(s)).V2.A, (*Simplex)(unsafe.Pointer(s)).V2.WA, (*Simplex)(unsafe.Pointer(s)).V3.A, (*Simplex)(unsafe.Pointer(s)).V3.WA)
		// todo why are these not equal?
		//*b = b2Weight3(s->v1.a, s->v1.wB, s->v2.a, s->v2.wB, s->v3.a, s->v3.wB);
		*(*Vec2)(unsafe.Pointer(b)) = *(*Vec2)(unsafe.Pointer(a))
	default:
		*(*Vec2)(unsafe.Pointer(a)) = b2Vec2_zero
		*(*Vec2)(unsafe.Pointer(b)) = b2Vec2_zero
		if bool(!(int32FromInt32(false1) != 0)) && b2InternalAssertFcn(tls, __ccgo_ts+4123, __ccgo_ts+4097, int32FromInt32(241)) != 0 {
			__builtin_trap(tls)
		}
		break
	}
}

func b2ComputePolygonCentroid(tls *_Stack, vertices uintptr, count int32) (r Vec2) {
	var a4, area, inv3, invArea, v12, v19 float32
	var center, e1, e2, origin, v10, v11, v14, v15, v16, v18, v2, v20, v21, v23, v24, v25, v3, v4, v6, v7, v8 Vec2
	var i int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a4, area, center, e1, e2, i, inv3, invArea, origin, v10, v11, v12, v14, v15, v16, v18, v19, v2, v20, v21, v23, v24, v25, v3, v4, v6, v7, v8
	center = Vec2{}
	area = float32FromFloat32(0)
	// Get a reference point for forming triangles.
	// Use the first vertex to reduce round-off errors.
	origin = *(*Vec2)(unsafe.Pointer(vertices))
	inv3 = float32FromFloat32(1) / float32FromFloat32(3)
	i = int32(1)
	for {
		if !(i < count-int32(1)) {
			break
		}
		v2 = *(*Vec2)(unsafe.Pointer(vertices + uintptr(i)*8))
		v3 = origin
		v4 = Vec2{
			X: v2.X - v3.X,
			Y: v2.Y - v3.Y,
		}
		goto _5
	_5:
		// Triangle edges
		e1 = v4
		v6 = *(*Vec2)(unsafe.Pointer(vertices + uintptr(i+int32(1))*8))
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
	if !(area > float32FromFloat32(1.1920928955078125e-07)) && b2InternalAssertFcn(tls, __ccgo_ts+6505, __ccgo_ts+6524, int32FromInt32(45)) != 0 {
		__builtin_trap(tls)
	}
	invArea = float32FromFloat32(1) / area
	center.X *= invArea
	center.Y *= invArea
	// Restore offset
	v23 = origin
	v24 = center
	v25 = Vec2{
		X: v23.X + v24.X,
		Y: v23.Y + v24.Y,
	}
	goto _26
_26:
	center = v25
	return center
}

func b2ComputeCircleMass(tls *_Stack, shape uintptr, density float32) (r MassData) {
	var massData MassData
	var rr, v3 float32
	var v1, v2 Vec2
	_, _, _, _, _ = massData, rr, v1, v2, v3
	rr = float32((*Circle)(unsafe.Pointer(shape)).Radius * (*Circle)(unsafe.Pointer(shape)).Radius)
	massData.Mass = float32(float32(density*float32FromFloat32(3.14159265359)) * rr)
	massData.Center = (*Circle)(unsafe.Pointer(shape)).Center
	// inertia about the local origin
	v1 = (*Circle)(unsafe.Pointer(shape)).Center
	v2 = (*Circle)(unsafe.Pointer(shape)).Center
	v3 = float32(v1.X*v2.X) + float32(v1.Y*v2.Y)
	goto _4
_4:
	massData.RotationalInertia = float32(massData.Mass * (float32(float32FromFloat32(0.5)*rr) + v3))
	return massData
}

func b2ComputeCapsuleMass(tls *_Stack, shape uintptr, density float32) (r MassData) {
	var boxInertia, boxMass, circleInertia, circleMass, h, lc, length, ll, radius, rr, v10, v6 float32
	var massData MassData
	var p1, p2, v1, v2, v3, v5, v8, v9 Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = boxInertia, boxMass, circleInertia, circleMass, h, lc, length, ll, massData, p1, p2, radius, rr, v1, v10, v2, v3, v5, v6, v8, v9
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
	// inertia about the local origin
	v8 = massData.Center
	v9 = massData.Center
	v10 = float32(v8.X*v9.X) + float32(v8.Y*v9.Y)
	goto _11
_11:
	massData.RotationalInertia += float32(massData.Mass * v10)
	return massData
}

func b2ComputePolygonMass(tls *_Stack, shape uintptr, density float32) (r1 MassData) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var D, area, ex1, ex2, ey1, ey2, intx2, inty2, inv3, invArea, invLength, length, radius, rotationalInertia, sqrt2, triangleArea, v11, v27, v34, v44, v48 float32
	var center, e1, e2, mid, n, n1, n2, r, v10, v12, v13, v17, v18, v19, v21, v22, v23, v25, v26, v29, v3, v30, v31, v33, v35, v36, v38, v39, v4, v40, v42, v43, v46, v47, v5, v7, v8 Vec2
	var count, i, i1, i2, j, v2 int32
	var massData MassData
	var vertices [8]Vec2
	var _ /* capsule at bp+12 */ Capsule
	var _ /* circle at bp+0 */ Circle
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = D, area, center, count, e1, e2, ex1, ex2, ey1, ey2, i, i1, i2, intx2, inty2, inv3, invArea, invLength, j, length, massData, mid, n, n1, n2, r, radius, rotationalInertia, sqrt2, triangleArea, vertices, v10, v11, v12, v13, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v29, v3, v30, v31, v33, v34, v35, v36, v38, v39, v4, v40, v42, v43, v44, v46, v47, v48, v5, v7, v8
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
	if !((*Polygon)(unsafe.Pointer(shape)).Count > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+6762, __ccgo_ts+6524, int32FromInt32(302)) != 0 {
		__builtin_trap(tls)
	}
	if (*Polygon)(unsafe.Pointer(shape)).Count == int32(1) {
		(*(*Circle)(unsafe.Pointer(bp))).Center = *(*Vec2)(unsafe.Pointer(shape))
		(*(*Circle)(unsafe.Pointer(bp))).Radius = (*Polygon)(unsafe.Pointer(shape)).Radius
		return b2ComputeCircleMass(tls, bp, density)
	}
	if (*Polygon)(unsafe.Pointer(shape)).Count == int32(2) {
		(*(*Capsule)(unsafe.Pointer(bp + 12))).Center1 = *(*Vec2)(unsafe.Pointer(shape))
		(*(*Capsule)(unsafe.Pointer(bp + 12))).Center2 = *(*Vec2)(unsafe.Pointer(shape + 1*8))
		(*(*Capsule)(unsafe.Pointer(bp + 12))).Radius = (*Polygon)(unsafe.Pointer(shape)).Radius
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
				v2 = count - int32(1)
			} else {
				v2 = i - int32(1)
			}
			j = v2
			n1 = *(*Vec2)(unsafe.Pointer(shape + 64 + uintptr(j)*8))
			n2 = *(*Vec2)(unsafe.Pointer(shape + 64 + uintptr(i)*8))
			v3 = n1
			v4 = n2
			v5 = Vec2{
				X: v3.X + v4.X,
				Y: v3.Y + v4.Y,
			}
			goto _6
		_6:
			v7 = v5
			length = sqrtf(tls, float32(v7.X*v7.X)+float32(v7.Y*v7.Y))
			if length < float32FromFloat32(1.1920928955078125e-07) {
				v8 = Vec2{}
				goto _9
			}
			invLength = float32FromFloat32(1) / length
			n = Vec2{
				X: float32(invLength * v7.X),
				Y: float32(invLength * v7.Y),
			}
			v8 = n
			goto _9
		_9:
			mid = v8
			v10 = *(*Vec2)(unsafe.Pointer(shape + uintptr(i)*8))
			v11 = float32(sqrt2 * radius)
			v12 = mid
			v13 = Vec2{
				X: v10.X + float32(v11*v12.X),
				Y: v10.Y + float32(v11*v12.Y),
			}
			goto _14
		_14:
			vertices[i] = v13
			goto _1
		_1:
			;
			i = i + 1
		}
	} else {
		i1 = 0
		for {
			if !(i1 < count) {
				break
			}
			vertices[i1] = *(*Vec2)(unsafe.Pointer(shape + uintptr(i1)*8))
			goto _15
		_15:
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
		v17 = vertices[i2]
		v18 = r
		v19 = Vec2{
			X: v17.X - v18.X,
			Y: v17.Y - v18.Y,
		}
		goto _20
	_20:
		// Triangle edges
		e1 = v19
		v21 = vertices[i2+int32(1)]
		v22 = r
		v23 = Vec2{
			X: v21.X - v22.X,
			Y: v21.Y - v22.Y,
		}
		goto _24
	_24:
		e2 = v23
		v25 = e1
		v26 = e2
		v27 = float32(v25.X*v26.Y) - float32(v25.Y*v26.X)
		goto _28
	_28:
		D = v27
		triangleArea = float32(float32FromFloat32(0.5) * D)
		area = area + triangleArea
		// Area weighted centroid, r at origin
		v29 = e1
		v30 = e2
		v31 = Vec2{
			X: v29.X + v30.X,
			Y: v29.Y + v30.Y,
		}
		goto _32
	_32:
		v33 = center
		v34 = float32(triangleArea * inv3)
		v35 = v31
		v36 = Vec2{
			X: v33.X + float32(v34*v35.X),
			Y: v33.Y + float32(v34*v35.Y),
		}
		goto _37
	_37:
		center = v36
		ex1 = e1.X
		ey1 = e1.Y
		ex2 = e2.X
		ey2 = e2.Y
		intx2 = float32(ex1*ex1) + float32(ex2*ex1) + float32(ex2*ex2)
		inty2 = float32(ey1*ey1) + float32(ey2*ey1) + float32(ey2*ey2)
		rotationalInertia = rotationalInertia + float32(float32(float32(float32FromFloat32(0.25)*inv3)*D)*(intx2+inty2))
		goto _16
	_16:
		;
		i2 = i2 + 1
	}
	// Total mass
	massData.Mass = float32(density * area)
	// Center of mass, shift back from origin at r
	if !(area > float32FromFloat32(1.1920928955078125e-07)) && b2InternalAssertFcn(tls, __ccgo_ts+6505, __ccgo_ts+6524, int32FromInt32(386)) != 0 {
		__builtin_trap(tls)
	}
	invArea = float32FromFloat32(1) / area
	center.X *= invArea
	center.Y *= invArea
	v38 = r
	v39 = center
	v40 = Vec2{
		X: v38.X + v39.X,
		Y: v38.Y + v39.Y,
	}
	goto _41
_41:
	massData.Center = v40
	// Inertia tensor relative to the local origin (point s).
	massData.RotationalInertia = float32(density * rotationalInertia)
	// Shift to center of mass then to original body origin.
	v42 = massData.Center
	v43 = massData.Center
	v44 = float32(v42.X*v43.X) + float32(v42.Y*v43.Y)
	goto _45
_45:
	v46 = center
	v47 = center
	v48 = float32(v46.X*v47.X) + float32(v46.Y*v47.Y)
	goto _49
_49:
	massData.RotationalInertia += float32(massData.Mass * (v44 - v48))
	return massData
}

func b2ComputeCircleAABB(tls *_Stack, shape uintptr, xf Transform) (r1 AABB) {
	var aabb AABB
	var p1, v2, v3 Vec2
	var r, x, y float32
	var v1 Transform
	_, _, _, _, _, _, _, _ = aabb, p1, r, x, y, v1, v2, v3
	v1 = xf
	v2 = (*Circle)(unsafe.Pointer(shape)).Center
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	p1 = v3
	r = (*Circle)(unsafe.Pointer(shape)).Radius
	aabb = AABB{
		LowerBound: Vec2{
			X: p1.X - r,
			Y: p1.Y - r,
		},
		UpperBound: Vec2{
			X: p1.X + r,
			Y: p1.Y + r,
		},
	}
	return aabb
}

func b2ComputeCapsuleAABB(tls *_Stack, shape uintptr, xf Transform) (r1 AABB) {
	var aabb AABB
	var c, c1, lower, r, upper, v1, v2, v10, v21, v211, v23, v24, v25, v27, v28, v3, v39, v41, v42, v43, v6, v7, v9 Vec2
	var x, y, v111, v12, v13, v15, v16, v17, v18, v20, v29, v30, v31, v33, v34, v35, v36, v38 float32
	var v11, v5 Transform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, c, c1, lower, r, upper, v1, v2, x, y, v11, v10, v111, v12, v13, v15, v16, v17, v18, v21, v20, v211, v23, v24, v25, v27, v28, v29, v3, v30, v31, v33, v34, v35, v36, v38, v39, v41, v42, v43, v5, v6, v7, v9
	v11 = xf
	v21 = (*Capsule)(unsafe.Pointer(shape)).Center1
	x = float32(v11.Q.C*v21.X) - float32(v11.Q.S*v21.Y) + v11.P.X
	y = float32(v11.Q.S*v21.X) + float32(v11.Q.C*v21.Y) + v11.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	v1 = v3
	v5 = xf
	v6 = (*Capsule)(unsafe.Pointer(shape)).Center2
	x = float32(v5.Q.C*v6.X) - float32(v5.Q.S*v6.Y) + v5.P.X
	y = float32(v5.Q.S*v6.X) + float32(v5.Q.C*v6.Y) + v5.P.Y
	v7 = Vec2{
		X: x,
		Y: y,
	}
	goto _8
_8:
	v2 = v7
	r = Vec2{
		X: (*Capsule)(unsafe.Pointer(shape)).Radius,
		Y: (*Capsule)(unsafe.Pointer(shape)).Radius,
	}
	v9 = v1
	v10 = v2
	v111 = v9.X
	v12 = v10.X
	if v111 < v12 {
		v15 = v111
	} else {
		v15 = v12
	}
	v13 = v15
	goto _14
_14:
	c.X = v13
	v16 = v9.Y
	v17 = v10.Y
	if v16 < v17 {
		v20 = v16
	} else {
		v20 = v17
	}
	v18 = v20
	goto _19
_19:
	c.Y = v18
	v211 = c
	goto _22
_22:
	v23 = v211
	v24 = r
	v25 = Vec2{
		X: v23.X - v24.X,
		Y: v23.Y - v24.Y,
	}
	goto _26
_26:
	lower = v25
	v27 = v1
	v28 = v2
	v29 = v27.X
	v30 = v28.X
	if v29 > v30 {
		v33 = v29
	} else {
		v33 = v30
	}
	v31 = v33
	goto _32
_32:
	c1.X = v31
	v34 = v27.Y
	v35 = v28.Y
	if v34 > v35 {
		v38 = v34
	} else {
		v38 = v35
	}
	v36 = v38
	goto _37
_37:
	c1.Y = v36
	v39 = c1
	goto _40
_40:
	v41 = v39
	v42 = r
	v43 = Vec2{
		X: v41.X + v42.X,
		Y: v41.Y + v42.Y,
	}
	goto _44
_44:
	upper = v43
	aabb = AABB{
		LowerBound: lower,
		UpperBound: upper,
	}
	return aabb
}

func b2ComputePolygonAABB(tls *_Stack, shape uintptr, xf Transform) (r1 AABB) {
	var aabb AABB
	var c, c1, lower, r, upper, v, v10, v11, v2, v22, v24, v25, v3, v36, v38, v39, v40, v42, v43, v44, v7, v8 Vec2
	var i int32
	var x, y, v12, v13, v14, v16, v17, v18, v19, v21, v26, v27, v28, v30, v31, v32, v33, v35 float32
	var v1, v6 Transform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, c, c1, i, lower, r, upper, v, x, y, v1, v10, v11, v12, v13, v14, v16, v17, v18, v19, v2, v21, v22, v24, v25, v26, v27, v28, v3, v30, v31, v32, v33, v35, v36, v38, v39, v40, v42, v43, v44, v6, v7, v8
	if !((*Polygon)(unsafe.Pointer(shape)).Count > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+6762, __ccgo_ts+6524, int32FromInt32(425)) != 0 {
		__builtin_trap(tls)
	}
	v1 = xf
	v2 = *(*Vec2)(unsafe.Pointer(shape))
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	lower = v3
	upper = lower
	i = int32(1)
	for {
		if !(i < (*Polygon)(unsafe.Pointer(shape)).Count) {
			break
		}
		v6 = xf
		v7 = *(*Vec2)(unsafe.Pointer(shape + uintptr(i)*8))
		x = float32(v6.Q.C*v7.X) - float32(v6.Q.S*v7.Y) + v6.P.X
		y = float32(v6.Q.S*v7.X) + float32(v6.Q.C*v7.Y) + v6.P.Y
		v8 = Vec2{
			X: x,
			Y: y,
		}
		goto _9
	_9:
		v = v8
		v10 = lower
		v11 = v
		v12 = v10.X
		v13 = v11.X
		if v12 < v13 {
			v16 = v12
		} else {
			v16 = v13
		}
		v14 = v16
		goto _15
	_15:
		c.X = v14
		v17 = v10.Y
		v18 = v11.Y
		if v17 < v18 {
			v21 = v17
		} else {
			v21 = v18
		}
		v19 = v21
		goto _20
	_20:
		c.Y = v19
		v22 = c
		goto _23
	_23:
		lower = v22
		v24 = upper
		v25 = v
		v26 = v24.X
		v27 = v25.X
		if v26 > v27 {
			v30 = v26
		} else {
			v30 = v27
		}
		v28 = v30
		goto _29
	_29:
		c1.X = v28
		v31 = v24.Y
		v32 = v25.Y
		if v31 > v32 {
			v35 = v31
		} else {
			v35 = v32
		}
		v33 = v35
		goto _34
	_34:
		c1.Y = v33
		v36 = c1
		goto _37
	_37:
		upper = v36
		goto _5
	_5:
		;
		i = i + 1
	}
	r = Vec2{
		X: (*Polygon)(unsafe.Pointer(shape)).Radius,
		Y: (*Polygon)(unsafe.Pointer(shape)).Radius,
	}
	v38 = lower
	v39 = r
	v40 = Vec2{
		X: v38.X - v39.X,
		Y: v38.Y - v39.Y,
	}
	goto _41
_41:
	lower = v40
	v42 = upper
	v43 = r
	v44 = Vec2{
		X: v42.X + v43.X,
		Y: v42.Y + v43.Y,
	}
	goto _45
_45:
	upper = v44
	aabb = AABB{
		LowerBound: lower,
		UpperBound: upper,
	}
	return aabb
}

func b2ComputeSegmentAABB(tls *_Stack, shape uintptr, xf Transform) (r AABB) {
	var aabb AABB
	var c, c1, lower, upper, v1, v2, v10, v21, v211, v23, v24, v3, v35, v6, v7, v9 Vec2
	var x, y, v111, v12, v13, v15, v16, v17, v18, v20, v25, v26, v27, v29, v30, v31, v32, v34 float32
	var v11, v5 Transform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, c, c1, lower, upper, v1, v2, x, y, v11, v10, v111, v12, v13, v15, v16, v17, v18, v21, v20, v211, v23, v24, v25, v26, v27, v29, v3, v30, v31, v32, v34, v35, v5, v6, v7, v9
	v11 = xf
	v21 = (*Segment)(unsafe.Pointer(shape)).Point1
	x = float32(v11.Q.C*v21.X) - float32(v11.Q.S*v21.Y) + v11.P.X
	y = float32(v11.Q.S*v21.X) + float32(v11.Q.C*v21.Y) + v11.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	v1 = v3
	v5 = xf
	v6 = (*Segment)(unsafe.Pointer(shape)).Point2
	x = float32(v5.Q.C*v6.X) - float32(v5.Q.S*v6.Y) + v5.P.X
	y = float32(v5.Q.S*v6.X) + float32(v5.Q.C*v6.Y) + v5.P.Y
	v7 = Vec2{
		X: x,
		Y: y,
	}
	goto _8
_8:
	v2 = v7
	v9 = v1
	v10 = v2
	v111 = v9.X
	v12 = v10.X
	if v111 < v12 {
		v15 = v111
	} else {
		v15 = v12
	}
	v13 = v15
	goto _14
_14:
	c.X = v13
	v16 = v9.Y
	v17 = v10.Y
	if v16 < v17 {
		v20 = v16
	} else {
		v20 = v17
	}
	v18 = v20
	goto _19
_19:
	c.Y = v18
	v211 = c
	goto _22
_22:
	lower = v211
	v23 = v1
	v24 = v2
	v25 = v23.X
	v26 = v24.X
	if v25 > v26 {
		v29 = v25
	} else {
		v29 = v26
	}
	v27 = v29
	goto _28
_28:
	c1.X = v27
	v30 = v23.Y
	v31 = v24.Y
	if v30 > v31 {
		v34 = v30
	} else {
		v34 = v31
	}
	v32 = v34
	goto _33
_33:
	c1.Y = v32
	v35 = c1
	goto _36
_36:
	upper = v35
	aabb = AABB{
		LowerBound: lower,
		UpperBound: upper,
	}
	return aabb
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
	var b8, c, c1, c2, c3, e, n, p1, p2, r, s1, s2, s3, vi, vj, v100, v102, v103, v19, v21, v22, v33, v36, v37, v42, v44, v45, v49, v50, v53, v54, v58, v59, v62, v63, v64, v66, v67, v7, v70, v71, v72, v74, v75, v8, v91, v92, v93, v95, v96, v98, v99 Vec2
	var d, distSqr, distance, dsq, dsq1, dsq11, dsq2, invLength, length, linearSlop, tolSqr, v10, v104, v11, v13, v14, v15, v16, v18, v23, v24, v25, v27, v28, v29, v30, v32, v38, v46, v51, v55, v60, v76, v9 float32
	var f1, f2, i, i1, i11, i2, i21, i3, i31, i4, i5, i6, j, j1, leftCount, n1, rightCount, v1, v2, v3, v40, v5, v78, v79, v80, v83, v85, v88 int32
	var ps [8]Vec2
	var searching, unique uint8
	var v81, v84, v86, v89 uintptr
	var _ /* hull at bp+0 */ Hull
	var _ /* hull1 at bp+164 */ Hull
	var _ /* hull2 at bp+232 */ Hull
	var _ /* leftPoints at bp+116 */ [6]Vec2
	var _ /* rightPoints at bp+68 */ [6]Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, b8, c, c1, c2, c3, d, distSqr, distance, dsq, dsq1, dsq11, dsq2, e, f1, f2, i, i1, i11, i2, i21, i3, i31, i4, i5, i6, invLength, j, j1, leftCount, length, linearSlop, n, n1, p1, p2, ps, r, rightCount, s1, s2, s3, searching, tolSqr, unique, vi, vj, v1, v10, v100, v102, v103, v104, v11, v13, v14, v15, v16, v18, v19, v2, v21, v22, v23, v24, v25, v27, v28, v29, v3, v30, v32, v33, v36, v37, v38, v40, v41, v42, v44, v45, v46, v49, v5, v50, v51, v53, v54, v55, v58, v59, v60, v62, v63, v64, v66, v67, v7, v70, v71, v72, v74, v75, v76, v78, v79, v8, v80, v81, v83, v84, v85, v86, v88, v89, v9, v91, v92, v93, v95, v96, v98, v99
	(*(*Hull)(unsafe.Pointer(bp))).Count = 0
	if count < int32(3) || count > int32(_B2_MAX_POLYGON_VERTICES) {
		// check your data
		return *(*Hull)(unsafe.Pointer(bp))
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
	n1 = 0
	linearSlop = float32(float32FromFloat32(0.005) * b2_lengthUnitsPerMeter)
	tolSqr = float32(float32(float32FromFloat32(16)*linearSlop) * linearSlop)
	i = 0
	for {
		if !(i < count) {
			break
		}
		v7 = aabb.LowerBound
		v8 = *(*Vec2)(unsafe.Pointer(points + uintptr(i)*8))
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
		v21 = aabb.UpperBound
		v22 = *(*Vec2)(unsafe.Pointer(points + uintptr(i)*8))
		v23 = v21.X
		v24 = v22.X
		if v23 > v24 {
			v27 = v23
		} else {
			v27 = v24
		}
		v25 = v27
		goto _26
	_26:
		c1.X = v25
		v28 = v21.Y
		v29 = v22.Y
		if v28 > v29 {
			v32 = v28
		} else {
			v32 = v29
		}
		v30 = v32
		goto _31
	_31:
		c1.Y = v30
		v33 = c1
		goto _34
	_34:
		aabb.UpperBound = v33
		vi = *(*Vec2)(unsafe.Pointer(points + uintptr(i)*8))
		unique = boolUint8(true1 != 0)
		j = 0
		for {
			if !(j < i) {
				break
			}
			vj = *(*Vec2)(unsafe.Pointer(points + uintptr(j)*8))
			v36 = vi
			v37 = vj
			c2 = Vec2{
				X: v37.X - v36.X,
				Y: v37.Y - v36.Y,
			}
			v38 = float32(c2.X*c2.X) + float32(c2.Y*c2.Y)
			goto _39
		_39:
			distSqr = v38
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
			v40 = n1
			n1 = n1 + 1
			ps[v40] = vi
		}
		goto _6
	_6:
		;
		i = i + 1
	}
	if n1 < int32(3) {
		// all points very close together, check your data and check your scale
		return *(*Hull)(unsafe.Pointer(bp))
	}
	v41 = aabb
	b8 = Vec2{
		X: float32(float32FromFloat32(0.5) * (v41.LowerBound.X + v41.UpperBound.X)),
		Y: float32(float32FromFloat32(0.5) * (v41.LowerBound.Y + v41.UpperBound.Y)),
	}
	v42 = b8
	goto _43
_43:
	// Find an extreme point as the first point on the hull
	c3 = v42
	f1 = 0
	v44 = c3
	v45 = ps[f1]
	c2 = Vec2{
		X: v45.X - v44.X,
		Y: v45.Y - v44.Y,
	}
	v46 = float32(c2.X*c2.X) + float32(c2.Y*c2.Y)
	goto _47
_47:
	dsq11 = v46
	i1 = int32(1)
	for {
		if !(i1 < n1) {
			break
		}
		v49 = c3
		v50 = ps[i1]
		c2 = Vec2{
			X: v50.X - v49.X,
			Y: v50.Y - v49.Y,
		}
		v51 = float32(c2.X*c2.X) + float32(c2.Y*c2.Y)
		goto _52
	_52:
		dsq = v51
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
	ps[f1] = ps[n1-int32(1)]
	n1 = n1 - int32(1)
	f2 = 0
	v53 = p1
	v54 = ps[f2]
	c2 = Vec2{
		X: v54.X - v53.X,
		Y: v54.Y - v53.Y,
	}
	v55 = float32(c2.X*c2.X) + float32(c2.Y*c2.Y)
	goto _56
_56:
	dsq2 = v55
	i2 = int32(1)
	for {
		if !(i2 < n1) {
			break
		}
		v58 = p1
		v59 = ps[i2]
		c2 = Vec2{
			X: v59.X - v58.X,
			Y: v59.Y - v58.Y,
		}
		v60 = float32(c2.X*c2.X) + float32(c2.Y*c2.Y)
		goto _61
	_61:
		dsq1 = v60
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
	ps[f2] = ps[n1-int32(1)]
	n1 = n1 - int32(1)
	rightCount = 0
	leftCount = 0
	v62 = p2
	v63 = p1
	v64 = Vec2{
		X: v62.X - v63.X,
		Y: v62.Y - v63.Y,
	}
	goto _65
_65:
	v66 = v64
	length = sqrtf(tls, float32(v66.X*v66.X)+float32(v66.Y*v66.Y))
	if length < float32FromFloat32(1.1920928955078125e-07) {
		v67 = Vec2{}
		goto _68
	}
	invLength = float32FromFloat32(1) / length
	n = Vec2{
		X: float32(invLength * v66.X),
		Y: float32(invLength * v66.Y),
	}
	v67 = n
	goto _68
_68:
	e = v67
	i3 = 0
	for {
		if !(i3 < n1) {
			break
		}
		v70 = ps[i3]
		v71 = p1
		v72 = Vec2{
			X: v70.X - v71.X,
			Y: v70.Y - v71.Y,
		}
		goto _73
	_73:
		v74 = v72
		v75 = e
		v76 = float32(v74.X*v75.Y) - float32(v74.Y*v75.X)
		goto _77
	_77:
		d = v76
		// slop used here to skip points that are very close to the line p1-p2
		if d >= float32(float32FromFloat32(2)*linearSlop) {
			v78 = rightCount
			rightCount = rightCount + 1
			(*(*[6]Vec2)(unsafe.Pointer(bp + 68)))[v78] = ps[i3]
		} else {
			if d <= float32(-float32FromFloat32(2)*linearSlop) {
				v79 = leftCount
				leftCount = leftCount + 1
				(*(*[6]Vec2)(unsafe.Pointer(bp + 116)))[v79] = ps[i3]
			}
		}
		goto _69
	_69:
		;
		i3 = i3 + 1
	}
	// compute hulls on right and left
	*(*Hull)(unsafe.Pointer(bp + 164)) = b2RecurseHull(tls, p1, p2, bp+68, rightCount)
	*(*Hull)(unsafe.Pointer(bp + 232)) = b2RecurseHull(tls, p2, p1, bp+116, leftCount)
	if (*(*Hull)(unsafe.Pointer(bp + 164))).Count == 0 && (*(*Hull)(unsafe.Pointer(bp + 232))).Count == 0 {
		// all points collinear
		return *(*Hull)(unsafe.Pointer(bp))
	}
	// stitch hulls together, preserving CCW winding order
	v81 = bp + 64
	v80 = *(*int32)(unsafe.Pointer(v81))
	*(*int32)(unsafe.Pointer(v81)) = *(*int32)(unsafe.Pointer(v81)) + 1
	*(*Vec2)(unsafe.Pointer(bp + uintptr(v80)*8)) = p1
	i4 = 0
	for {
		if !(i4 < (*(*Hull)(unsafe.Pointer(bp + 164))).Count) {
			break
		}
		v84 = bp + 64
		v83 = *(*int32)(unsafe.Pointer(v84))
		*(*int32)(unsafe.Pointer(v84)) = *(*int32)(unsafe.Pointer(v84)) + 1
		*(*Vec2)(unsafe.Pointer(bp + uintptr(v83)*8)) = *(*Vec2)(unsafe.Pointer(bp + 164 + uintptr(i4)*8))
		goto _82
	_82:
		;
		i4 = i4 + 1
	}
	v86 = bp + 64
	v85 = *(*int32)(unsafe.Pointer(v86))
	*(*int32)(unsafe.Pointer(v86)) = *(*int32)(unsafe.Pointer(v86)) + 1
	*(*Vec2)(unsafe.Pointer(bp + uintptr(v85)*8)) = p2
	i5 = 0
	for {
		if !(i5 < (*(*Hull)(unsafe.Pointer(bp + 232))).Count) {
			break
		}
		v89 = bp + 64
		v88 = *(*int32)(unsafe.Pointer(v89))
		*(*int32)(unsafe.Pointer(v89)) = *(*int32)(unsafe.Pointer(v89)) + 1
		*(*Vec2)(unsafe.Pointer(bp + uintptr(v88)*8)) = *(*Vec2)(unsafe.Pointer(bp + 232 + uintptr(i5)*8))
		goto _87
	_87:
		;
		i5 = i5 + 1
	}
	if !((*(*Hull)(unsafe.Pointer(bp))).Count <= int32FromInt32(_B2_MAX_POLYGON_VERTICES)) && b2InternalAssertFcn(tls, __ccgo_ts+6905, __ccgo_ts+6883, int32FromInt32(225)) != 0 {
		__builtin_trap(tls)
	}
	// merge collinear
	searching = boolUint8(true1 != 0)
	for searching != 0 && (*(*Hull)(unsafe.Pointer(bp))).Count > int32(2) {
		searching = boolUint8(false1 != 0)
		i6 = 0
		for {
			if !(i6 < (*(*Hull)(unsafe.Pointer(bp))).Count) {
				break
			}
			i11 = i6
			i21 = (i6 + int32(1)) % (*(*Hull)(unsafe.Pointer(bp))).Count
			i31 = (i6 + int32(2)) % (*(*Hull)(unsafe.Pointer(bp))).Count
			s1 = *(*Vec2)(unsafe.Pointer(bp + uintptr(i11)*8))
			s2 = *(*Vec2)(unsafe.Pointer(bp + uintptr(i21)*8))
			s3 = *(*Vec2)(unsafe.Pointer(bp + uintptr(i31)*8))
			v91 = s3
			v92 = s1
			v93 = Vec2{
				X: v91.X - v92.X,
				Y: v91.Y - v92.Y,
			}
			goto _94
		_94:
			v95 = v93
			length = sqrtf(tls, float32(v95.X*v95.X)+float32(v95.Y*v95.Y))
			if length < float32FromFloat32(1.1920928955078125e-07) {
				v96 = Vec2{}
				goto _97
			}
			invLength = float32FromFloat32(1) / length
			n = Vec2{
				X: float32(invLength * v95.X),
				Y: float32(invLength * v95.Y),
			}
			v96 = n
			goto _97
		_97:
			// unit edge vector for s1-s3
			r = v96
			v98 = s2
			v99 = s1
			v100 = Vec2{
				X: v98.X - v99.X,
				Y: v98.Y - v99.Y,
			}
			goto _101
		_101:
			v102 = v100
			v103 = r
			v104 = float32(v102.X*v103.Y) - float32(v102.Y*v103.X)
			goto _105
		_105:
			distance = v104
			if distance <= float32(float32FromFloat32(2)*linearSlop) {
				// remove midpoint from hull
				j1 = i21
				for {
					if !(j1 < (*(*Hull)(unsafe.Pointer(bp))).Count-int32(1)) {
						break
					}
					*(*Vec2)(unsafe.Pointer(bp + uintptr(j1)*8)) = *(*Vec2)(unsafe.Pointer(bp + uintptr(j1+int32(1))*8))
					goto _106
				_106:
					;
					j1 = j1 + 1
				}
				(*(*Hull)(unsafe.Pointer(bp))).Count -= int32(1)
				// continue searching for collinear points
				searching = boolUint8(true1 != 0)
				break
			}
			goto _90
		_90:
			;
			i6 = i6 + 1
		}
	}
	if (*(*Hull)(unsafe.Pointer(bp))).Count < int32(3) {
		// all points collinear, shouldn't be reached since this was validated above
		(*(*Hull)(unsafe.Pointer(bp))).Count = 0
	}
	return *(*Hull)(unsafe.Pointer(bp))
}

// C documentation
//
//	// Approximate cosine and sine for determinism. In my testing cosf and sinf produced
//	// the same results on x64 and ARM using MSVC, GCC, and Clang. However, I don't trust
//	// this result.
//	// https://en.wikipedia.org/wiki/Bh%C4%81skara_I%27s_sine_approximation_formula
func b2ComputeCosSin(tls *_Stack, radians1 float32) (r CosSin) {
	var c, invMag, mag, pi2, s, x, y, y1, y2, y21, y22, y23, v1, v3 float32
	var cs CosSin
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, cs, invMag, mag, pi2, s, x, y, y1, y2, y21, y22, y23, v1, v3
	v1 = __builtin_remainderf(tls, radians1, float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)))
	goto _2
_2:
	x = v1
	pi2 = float32(float32FromFloat32(3.14159265359) * float32FromFloat32(3.14159265359))
	if x < float32(-float32FromFloat32(0.5)*float32FromFloat32(3.14159265359)) {
		y = x + float32FromFloat32(3.14159265359)
		y21 = float32(y * y)
		c = -(pi2 - float32(float32FromFloat32(4)*y21)) / (pi2 + y21)
	} else {
		if x > float32(float32FromFloat32(0.5)*float32FromFloat32(3.14159265359)) {
			y1 = x - float32FromFloat32(3.14159265359)
			y22 = float32(y1 * y1)
			c = -(pi2 - float32(float32FromFloat32(4)*y22)) / (pi2 + y22)
		} else {
			y23 = float32(x * x)
			c = (pi2 - float32(float32FromFloat32(4)*y23)) / (pi2 + y23)
		}
	}
	if x < float32FromFloat32(0) {
		y2 = x + float32FromFloat32(3.14159265359)
		s = float32(float32(-float32FromFloat32(16)*y2)*(float32FromFloat32(3.14159265359)-y2)) / (float32(float32FromFloat32(5)*pi2) - float32(float32(float32FromFloat32(4)*y2)*(float32FromFloat32(3.14159265359)-y2)))
	} else {
		s = float32(float32(float32FromFloat32(16)*x)*(float32FromFloat32(3.14159265359)-x)) / (float32(float32FromFloat32(5)*pi2) - float32(float32(float32FromFloat32(4)*x)*(float32FromFloat32(3.14159265359)-x)))
	}
	mag = sqrtf(tls, float32(s*s)+float32(c*c))
	if float64(mag) > float64(0) {
		v3 = float32FromFloat32(1) / mag
	} else {
		v3 = float32FromFloat32(0)
	}
	invMag = v3
	cs = CosSin{
		Cosine: float32(c * invMag),
		Sine:   float32(s * invMag),
	}
	return cs
}

func b2ComputeRotationBetweenUnitVectors(tls *_Stack, v1 Vec2, v2 Vec2) (r Rot) {
	var invMag, mag, v111, v12, v14, v17, v21, v211, v24, v4, v5, v7, v9 float32
	var qn, rot, v23, v25 Rot
	var v11, v15, v16, v19, v20, v8 Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = invMag, mag, qn, rot, v11, v111, v12, v14, v15, v16, v17, v19, v21, v20, v211, v23, v24, v25, v4, v5, v7, v8, v9
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
	if !(v5 < float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07))) && b2InternalAssertFcn(tls, __ccgo_ts+10147, __ccgo_ts+10206, int32FromInt32(152)) != 0 {
		__builtin_trap(tls)
	}
	v8 = v2
	v9 = sqrtf(tls, float32(v8.X*v8.X)+float32(v8.Y*v8.Y))
	goto _10
_10:
	v111 = float32FromFloat32(1) - v9
	if v111 < float32FromInt32(0) {
		v14 = -v111
	} else {
		v14 = v111
	}
	v12 = v14
	goto _13
_13:
	;
	if !(v12 < float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07))) && b2InternalAssertFcn(tls, __ccgo_ts+10238, __ccgo_ts+10206, int32FromInt32(153)) != 0 {
		__builtin_trap(tls)
	}
	v15 = v1
	v16 = v2
	v17 = float32(v15.X*v16.X) + float32(v15.Y*v16.Y)
	goto _18
_18:
	rot.C = v17
	v19 = v1
	v20 = v2
	v211 = float32(v19.X*v20.Y) - float32(v19.Y*v20.X)
	goto _22
_22:
	rot.S = v211
	v23 = rot
	mag = sqrtf(tls, float32(v23.S*v23.S)+float32(v23.C*v23.C))
	if float64(mag) > float64(0) {
		v24 = float32FromFloat32(1) / mag
	} else {
		v24 = float32FromFloat32(0)
	}
	invMag = v24
	qn = Rot{
		C: float32(v23.C * invMag),
		S: float32(v23.S * invMag),
	}
	v25 = qn
	goto _26
_26:
	return v25
}
