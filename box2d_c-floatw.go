//go:build !(goexperiment.simd && go1.27 && (arm64 || (amd64 && amd64.v3) || wasm))

package b2

import "unsafe"

var _ unsafe.Pointer

// C documentation
//
//	// scalar math
type b2FloatW4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

func b2ZeroW(tls *_Stack) (r b2FloatW4) {
	return b2FloatW4{}
}

func b2SplatW(tls *_Stack, scalar float32) (r b2FloatW4) {
	return b2FloatW4{
		X: scalar,
		Y: scalar,
		Z: scalar,
		W: scalar,
	}
}

func b2SetW(tls *_Stack, a float32, b float32, c float32, d float32) (r b2FloatW4) {
	return b2FloatW4{
		X: a,
		Y: b,
		Z: c,
		W: d,
	}
}

func b2NegW(tls *_Stack, a b2FloatW4) (r b2FloatW4) {
	return b2FloatW4{
		X: -a.X,
		Y: -a.Y,
		Z: -a.Z,
		W: -a.W,
	}
}

func b2AddW(tls *_Stack, a b2FloatW4, b b2FloatW4) (r b2FloatW4) {
	return b2FloatW4{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
		W: a.W + b.W,
	}
}

func b2SubW(tls *_Stack, a b2FloatW4, b b2FloatW4) (r b2FloatW4) {
	return b2FloatW4{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
		W: a.W - b.W,
	}
}

func b2MulW(tls *_Stack, a b2FloatW4, b b2FloatW4) (r b2FloatW4) {
	return b2FloatW4{
		X: float32(a.X * b.X),
		Y: float32(a.Y * b.Y),
		Z: float32(a.Z * b.Z),
		W: float32(a.W * b.W),
	}
}

func b2DivW(tls *_Stack, a b2FloatW4, b b2FloatW4) (r b2FloatW4) {
	return b2FloatW4{
		X: a.X / b.X,
		Y: a.Y / b.Y,
		Z: a.Z / b.Z,
		W: a.W / b.W,
	}
}

func b2MulAddW(tls *_Stack, a b2FloatW4, b b2FloatW4, c b2FloatW4) (r b2FloatW4) {
	return b2FloatW4{
		X: a.X + float32(b.X*c.X),
		Y: a.Y + float32(b.Y*c.Y),
		Z: a.Z + float32(b.Z*c.Z),
		W: a.W + float32(b.W*c.W),
	}
}

func b2MulSubW(tls *_Stack, a b2FloatW4, b b2FloatW4, c b2FloatW4) (r b2FloatW4) {
	return b2FloatW4{
		X: a.X - float32(b.X*c.X),
		Y: a.Y - float32(b.Y*c.Y),
		Z: a.Z - float32(b.Z*c.Z),
		W: a.W - float32(b.W*c.W),
	}
}

func b2MinW(tls *_Stack, a b2FloatW4, b b2FloatW4) (r1 b2FloatW4) {
	var r b2FloatW4
	var v1 float32
	_, _ = r, v1
	if a.X <= b.X {
		v1 = a.X
	} else {
		v1 = b.X
	}
	r.X = v1
	if a.Y <= b.Y {
		v1 = a.Y
	} else {
		v1 = b.Y
	}
	r.Y = v1
	if a.Z <= b.Z {
		v1 = a.Z
	} else {
		v1 = b.Z
	}
	r.Z = v1
	if a.W <= b.W {
		v1 = a.W
	} else {
		v1 = b.W
	}
	r.W = v1
	return r
}

func b2MaxW(tls *_Stack, a b2FloatW4, b b2FloatW4) (r1 b2FloatW4) {
	var r b2FloatW4
	var v1 float32
	_, _ = r, v1
	if a.X >= b.X {
		v1 = a.X
	} else {
		v1 = b.X
	}
	r.X = v1
	if a.Y >= b.Y {
		v1 = a.Y
	} else {
		v1 = b.Y
	}
	r.Y = v1
	if a.Z >= b.Z {
		v1 = a.Z
	} else {
		v1 = b.Z
	}
	r.Z = v1
	if a.W >= b.W {
		v1 = a.W
	} else {
		v1 = b.W
	}
	r.W = v1
	return r
}

// C documentation
//
//	// a = clamp(a, -b, b)
func b2SymClampW(tls *_Stack, a1 b2FloatW4, b b2FloatW4) (r1 b2FloatW4) {
	var r b2FloatW4
	var v1, v2, v3, v4, v6, v7 float32
	_, _, _, _, _, _, _ = r, v1, v2, v3, v4, v6, v7
	v1 = a1.X
	v2 = -b.X
	v3 = b.X
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
	r.X = v4
	v1 = a1.Y
	v2 = -b.Y
	v3 = b.Y
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
	goto _12
_12:
	r.Y = v4
	v1 = a1.Z
	v2 = -b.Z
	v3 = b.Z
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
	goto _19
_19:
	r.Z = v4
	v1 = a1.W
	v2 = -b.W
	v3 = b.W
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
	goto _26
_26:
	r.W = v4
	return r
}

func b2AndW(tls *_Stack, a b2FloatW4, b b2FloatW4) (r1 b2FloatW4) {
	var r b2FloatW4
	var v1 float32
	_, _ = r, v1
	if a.X != float32FromFloat32(0) && b.X != float32FromFloat32(0) {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.X = v1
	if a.Y != float32FromFloat32(0) && b.Y != float32FromFloat32(0) {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Y = v1
	if a.Z != float32FromFloat32(0) && b.Z != float32FromFloat32(0) {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Z = v1
	if a.W != float32FromFloat32(0) && b.W != float32FromFloat32(0) {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.W = v1
	return r
}

func b2GreaterThanW(tls *_Stack, a b2FloatW4, b b2FloatW4) (r1 b2FloatW4) {
	var r b2FloatW4
	var v1 float32
	_, _ = r, v1
	if a.X > b.X {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.X = v1
	if a.Y > b.Y {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Y = v1
	if a.Z > b.Z {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Z = v1
	if a.W > b.W {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.W = v1
	return r
}

func b2LessThanW(tls *_Stack, a b2FloatW4, b b2FloatW4) (r1 b2FloatW4) {
	var r b2FloatW4
	var v1 float32
	_, _ = r, v1
	if a.X < b.X {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.X = v1
	if a.Y < b.Y {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Y = v1
	if a.Z < b.Z {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Z = v1
	if a.W < b.W {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.W = v1
	return r
}

func b2AnyTrueW(tls *_Stack, mask b2FloatW4) (r uint8) {
	return boolUint8(mask.X != float32FromFloat32(0) || mask.Y != float32FromFloat32(0) || mask.Z != float32FromFloat32(0) || mask.W != float32FromFloat32(0))
}

func b2AllZeroW(tls *_Stack, a b2FloatW4) (r uint8) {
	return boolUint8(a.X == float32FromFloat32(0) && a.Y == float32FromFloat32(0) && a.Z == float32FromFloat32(0) && a.W == float32FromFloat32(0))
}

func b2LoadU(tls *_Stack, data uintptr) (r b2FloatW4) {
	return b2FloatW4{
		X: *(*float32)(unsafe.Add(nil, data)),
		Y: *(*float32)(unsafe.Add(nil, data+1*4)),
		Z: *(*float32)(unsafe.Add(nil, data+2*4)),
		W: *(*float32)(unsafe.Add(nil, data+3*4)),
	}
}

func b2StoreW(tls *_Stack, data uintptr, a b2FloatW4) {
	*(*float32)(unsafe.Add(nil, data)) = a.X
	*(*float32)(unsafe.Add(nil, data+1*4)) = a.Y
	*(*float32)(unsafe.Add(nil, data+2*4)) = a.Z
	*(*float32)(unsafe.Add(nil, data+3*4)) = a.W
}

// C documentation
//
//	// component-wise returns mask ? b : a
func b2BlendW(tls *_Stack, a b2FloatW4, b b2FloatW4, mask b2FloatW4) (r1 b2FloatW4) {
	var r b2FloatW4
	var v1 float32
	_, _ = r, v1
	if mask.X != float32FromFloat32(0) {
		v1 = b.X
	} else {
		v1 = a.X
	}
	r.X = v1
	if mask.Y != float32FromFloat32(0) {
		v1 = b.Y
	} else {
		v1 = a.Y
	}
	r.Y = v1
	if mask.Z != float32FromFloat32(0) {
		v1 = b.Z
	} else {
		v1 = a.Z
	}
	r.Z = v1
	if mask.W != float32FromFloat32(0) {
		v1 = b.W
	} else {
		v1 = a.W
	}
	r.W = v1
	return r
}

func b2UnpackLoW(tls *_Stack, a b2FloatW4, b b2FloatW4) (r b2FloatW4) {
	return b2FloatW4{
		X: a.X,
		Y: b.X,
		Z: a.Y,
		W: b.Y,
	}
}

func b2UnpackHiW(tls *_Stack, a b2FloatW4, b b2FloatW4) (r b2FloatW4) {
	return b2FloatW4{
		X: a.Z,
		Y: b.Z,
		Z: a.W,
		W: b.W,
	}
}

func b2SoftMaskW(tls *_Stack, indexA uintptr, indexB uintptr) (r1 b2FloatW4) {
	var r b2FloatW4
	var v1 float32
	_, _ = r, v1
	if *(*int32)(unsafe.Add(nil, indexA)) == 0 || *(*int32)(unsafe.Add(nil, indexB)) == 0 {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.X = v1
	if *(*int32)(unsafe.Add(nil, indexA+1*4)) == 0 || *(*int32)(unsafe.Add(nil, indexB+1*4)) == 0 {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Y = v1
	if *(*int32)(unsafe.Add(nil, indexA+2*4)) == 0 || *(*int32)(unsafe.Add(nil, indexB+2*4)) == 0 {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Z = v1
	if *(*int32)(unsafe.Add(nil, indexA+3*4)) == 0 || *(*int32)(unsafe.Add(nil, indexB+3*4)) == 0 {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.W = v1
	return r
}

// Custom gather/scatter for each SIMD type.

// C documentation
//
//	// This is a load and transpose
func b2GatherBodies(tls *_Stack, states uintptr, indices uintptr) (r b2BodyStateW4) {
	var i1, i2, i3, i4 int32
	var identity, s1, s2, s3, s4, v1, v2, v3, v4 b2BodyState
	var simdBody b2BodyStateW4
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = i1, i2, i3, i4, identity, s1, s2, s3, s4, simdBody, v1, v2, v3, v4
	identity = b2_identityBodyState5
	// zero means null
	i1 = *(*int32)(unsafe.Add(nil, indices)) - int32(1)
	i2 = *(*int32)(unsafe.Add(nil, indices+1*4)) - int32(1)
	i3 = *(*int32)(unsafe.Add(nil, indices+2*4)) - int32(1)
	i4 = *(*int32)(unsafe.Add(nil, indices+3*4)) - int32(1)
	if i1 == -int32(1) {
		v1 = identity
	} else {
		v1 = *(*b2BodyState)(unsafe.Add(nil, states+uintptr(i1)*32))
	}
	s1 = v1
	if i2 == -int32(1) {
		v2 = identity
	} else {
		v2 = *(*b2BodyState)(unsafe.Add(nil, states+uintptr(i2)*32))
	}
	s2 = v2
	if i3 == -int32(1) {
		v3 = identity
	} else {
		v3 = *(*b2BodyState)(unsafe.Add(nil, states+uintptr(i3)*32))
	}
	s3 = v3
	if i4 == -int32(1) {
		v4 = identity
	} else {
		v4 = *(*b2BodyState)(unsafe.Add(nil, states+uintptr(i4)*32))
	}
	s4 = v4
	simdBody.V.X = b2FloatW4{
		X: s1.LinearVelocity.X,
		Y: s2.LinearVelocity.X,
		Z: s3.LinearVelocity.X,
		W: s4.LinearVelocity.X,
	}
	simdBody.V.Y = b2FloatW4{
		X: s1.LinearVelocity.Y,
		Y: s2.LinearVelocity.Y,
		Z: s3.LinearVelocity.Y,
		W: s4.LinearVelocity.Y,
	}
	simdBody.W = b2FloatW4{
		X: s1.AngularVelocity,
		Y: s2.AngularVelocity,
		Z: s3.AngularVelocity,
		W: s4.AngularVelocity,
	}
	simdBody.Flags = b2FloatW4{
		X: float32(s1.Flags),
		Y: float32(s2.Flags),
		Z: float32(s3.Flags),
		W: float32(s4.Flags),
	}
	simdBody.Dp.X = b2FloatW4{
		X: s1.DeltaPosition.X,
		Y: s2.DeltaPosition.X,
		Z: s3.DeltaPosition.X,
		W: s4.DeltaPosition.X,
	}
	simdBody.Dp.Y = b2FloatW4{
		X: s1.DeltaPosition.Y,
		Y: s2.DeltaPosition.Y,
		Z: s3.DeltaPosition.Y,
		W: s4.DeltaPosition.Y,
	}
	simdBody.Dq.C = b2FloatW4{
		X: s1.DeltaRotation.C,
		Y: s2.DeltaRotation.C,
		Z: s3.DeltaRotation.C,
		W: s4.DeltaRotation.C,
	}
	simdBody.Dq.S = b2FloatW4{
		X: s1.DeltaRotation.S,
		Y: s2.DeltaRotation.S,
		Z: s3.DeltaRotation.S,
		W: s4.DeltaRotation.S,
	}
	return simdBody
}

// C documentation
//
//	// This writes only the velocities back to the solver bodies
func b2ScatterBodies(tls *_Stack, states uintptr, indices uintptr, simdBody uintptr) {
	var i1, i2, i3, i4 int32
	var state, state1, state2, state3 uintptr
	_, _, _, _, _, _, _, _ = i1, i2, i3, i4, state, state1, state2, state3
	// zero means null
	i1 = *(*int32)(unsafe.Add(nil, indices)) - int32(1)
	i2 = *(*int32)(unsafe.Add(nil, indices+1*4)) - int32(1)
	i3 = *(*int32)(unsafe.Add(nil, indices+2*4)) - int32(1)
	i4 = *(*int32)(unsafe.Add(nil, indices+3*4)) - int32(1)
	if i1 != -int32(1) && (*(*b2BodyState)(unsafe.Add(nil, states+uintptr(i1)*32))).Flags&uint32(b2_dynamicFlag) != uint32(0) {
		state = states + uintptr(i1)*32
		(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity.X = (*b2BodyStateW4)(unsafe.Pointer(simdBody)).V.X.X
		(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity.Y = (*b2BodyStateW4)(unsafe.Pointer(simdBody)).V.Y.X
		(*b2BodyState)(unsafe.Pointer(state)).AngularVelocity = (*b2BodyStateW4)(unsafe.Pointer(simdBody)).W.X
	}
	if i2 != -int32(1) && (*(*b2BodyState)(unsafe.Add(nil, states+uintptr(i2)*32))).Flags&uint32(b2_dynamicFlag) != uint32(0) {
		state1 = states + uintptr(i2)*32
		(*b2BodyState)(unsafe.Pointer(state1)).LinearVelocity.X = (*b2BodyStateW4)(unsafe.Pointer(simdBody)).V.X.Y
		(*b2BodyState)(unsafe.Pointer(state1)).LinearVelocity.Y = (*b2BodyStateW4)(unsafe.Pointer(simdBody)).V.Y.Y
		(*b2BodyState)(unsafe.Pointer(state1)).AngularVelocity = (*b2BodyStateW4)(unsafe.Pointer(simdBody)).W.Y
	}
	if i3 != -int32(1) && (*(*b2BodyState)(unsafe.Add(nil, states+uintptr(i3)*32))).Flags&uint32(b2_dynamicFlag) != uint32(0) {
		state2 = states + uintptr(i3)*32
		(*b2BodyState)(unsafe.Pointer(state2)).LinearVelocity.X = (*b2BodyStateW4)(unsafe.Pointer(simdBody)).V.X.Z
		(*b2BodyState)(unsafe.Pointer(state2)).LinearVelocity.Y = (*b2BodyStateW4)(unsafe.Pointer(simdBody)).V.Y.Z
		(*b2BodyState)(unsafe.Pointer(state2)).AngularVelocity = (*b2BodyStateW4)(unsafe.Pointer(simdBody)).W.Z
	}
	if i4 != -int32(1) && (*(*b2BodyState)(unsafe.Add(nil, states+uintptr(i4)*32))).Flags&uint32(b2_dynamicFlag) != uint32(0) {
		state3 = states + uintptr(i4)*32
		(*b2BodyState)(unsafe.Pointer(state3)).LinearVelocity.X = (*b2BodyStateW4)(unsafe.Pointer(simdBody)).V.X.W
		(*b2BodyState)(unsafe.Pointer(state3)).LinearVelocity.Y = (*b2BodyStateW4)(unsafe.Pointer(simdBody)).V.Y.W
		(*b2BodyState)(unsafe.Pointer(state3)).AngularVelocity = (*b2BodyStateW4)(unsafe.Pointer(simdBody)).W.W
	}
}
