//go:build goexperiment.simd && go1.27 && (arm64 || (amd64 && amd64.v3) || wasm)

package b2

// SIMD version of box2d's 4-wide contact solver math. The transpiled scalar version is in
// box2d_c-floatw.go, see floatWDecls in transpile/main.go. Most of this follows the SSE2
// branch of box2d-c/src/contact_solver_wide.inl, b2UnpackLoW and b2UnpackHiW are in
// simd_<arch>.go.
//
// Masks are all bits set per lane, like in the SSE2 and NEON versions.

import (
	"simd/archsimd"
	"unsafe"
)

type b2FloatW4 = archsimd.Float32x4

var allBitsW = archsimd.BroadcastUint32x4(^uint32(0))

func maskToW(mask archsimd.Mask32x4) b2FloatW4 {
	return allBitsW.Masked(mask).BitsToFloat32()
}

func wToMask(w b2FloatW4) archsimd.Mask32x4 {
	return w.ToBits().NotEqual(archsimd.Uint32x4{})
}

func b2ZeroW(tls *_Stack) b2FloatW4 {
	return b2FloatW4{}
}

func b2SplatW(tls *_Stack, scalar float32) b2FloatW4 {
	return archsimd.BroadcastFloat32x4(scalar)
}

func b2SetW(tls *_Stack, a, b, c, d float32) b2FloatW4 {
	return archsimd.LoadFloat32x4Array(&[4]float32{a, b, c, d})
}

func b2NegW(tls *_Stack, a b2FloatW4) b2FloatW4 {
	return a.Neg()
}

func b2AddW(tls *_Stack, a, b b2FloatW4) b2FloatW4 {
	return a.Add(b)
}

func b2SubW(tls *_Stack, a, b b2FloatW4) b2FloatW4 {
	return a.Sub(b)
}

func b2MulW(tls *_Stack, a, b b2FloatW4) b2FloatW4 {
	return a.Mul(b)
}

func b2DivW(tls *_Stack, a, b b2FloatW4) b2FloatW4 {
	return a.Div(b)
}

// not archsimd's MulAdd, it is fused and breaks determinism
func b2MulAddW(tls *_Stack, a, b, c b2FloatW4) b2FloatW4 {
	return a.Add(b.Mul(c))
}

func b2MulSubW(tls *_Stack, a, b, c b2FloatW4) b2FloatW4 {
	return a.Sub(b.Mul(c))
}

func b2MinW(tls *_Stack, a, b b2FloatW4) b2FloatW4 {
	return a.Min(b)
}

func b2MaxW(tls *_Stack, a, b b2FloatW4) b2FloatW4 {
	return a.Max(b)
}

// a = clamp(a, -b, b)
func b2SymClampW(tls *_Stack, a, b b2FloatW4) b2FloatW4 {
	return b.Neg().Max(a.Min(b))
}

func b2AndW(tls *_Stack, a, b b2FloatW4) b2FloatW4 {
	return a.ToBits().And(b.ToBits()).BitsToFloat32()
}

func b2GreaterThanW(tls *_Stack, a, b b2FloatW4) b2FloatW4 {
	return maskToW(a.Greater(b))
}

func b2LessThanW(tls *_Stack, a, b b2FloatW4) b2FloatW4 {
	return maskToW(a.Less(b))
}

func b2AnyTrueW(tls *_Stack, mask b2FloatW4) uint8 {
	var bits [4]uint32
	mask.ToBits().StoreArray(&bits)
	return boolUint8(bits[0]|bits[1]|bits[2]|bits[3] != 0)
}

func b2AllZeroW(tls *_Stack, a b2FloatW4) uint8 {
	var v [4]float32
	a.StoreArray(&v)
	return boolUint8(v[0] == 0 && v[1] == 0 && v[2] == 0 && v[3] == 0)
}

func b2LoadU(tls *_Stack, data uintptr) b2FloatW4 {
	return archsimd.LoadFloat32x4Array((*[4]float32)(unsafe.Add(nil, data)))
}

func b2StoreW(tls *_Stack, data uintptr, a b2FloatW4) {
	a.StoreArray((*[4]float32)(unsafe.Add(nil, data)))
}

// component-wise returns mask ? b : a
func b2BlendW(tls *_Stack, a, b, mask b2FloatW4) b2FloatW4 {
	return b.IfElse(wToMask(mask), a)
}

func b2SoftMaskW(tls *_Stack, indexA, indexB uintptr) b2FloatW4 {
	var zero archsimd.Int32x4
	a := archsimd.LoadInt32x4Array((*[4]int32)(unsafe.Add(nil, indexA))).Equal(zero)
	b := archsimd.LoadInt32x4Array((*[4]int32)(unsafe.Add(nil, indexB))).Equal(zero)
	return maskToW(a.Or(b))
}

// This is a load and transpose. A body state is [vx vy w flags] [dpx dpy dqc dqs].
func b2GatherBodies(tls *_Stack, states uintptr, indices uintptr) (r b2BodyStateW4) {
	idx := (*[4]int32)(unsafe.Add(nil, indices))

	// zero means null, use the identity body state
	load := func(index int32) (a, b b2FloatW4) {
		if index == 0 {
			return b2FloatW4{}, archsimd.LoadFloat32x4Array(&[4]float32{0, 0, 1, 0})
		}

		state := (*[8]float32)(unsafe.Add(nil, states+uintptr(index-1)*unsafe.Sizeof(b2BodyState{})))
		return archsimd.LoadFloat32x4Array((*[4]float32)(state[0:4])), archsimd.LoadFloat32x4Array((*[4]float32)(state[4:8]))
	}

	b1a, b1b := load(idx[0])
	b2a, b2b := load(idx[1])
	b3a, b3b := load(idx[2])
	b4a, b4b := load(idx[3])

	// [vx1 vx3 vy1 vy3], [vx2 vx4 vy2 vy4], [w1 w3 f1 f3], [w2 w4 f2 f4]
	t1a := b2UnpackLoW(tls, b1a, b3a)
	t2a := b2UnpackLoW(tls, b2a, b4a)
	t3a := b2UnpackHiW(tls, b1a, b3a)
	t4a := b2UnpackHiW(tls, b2a, b4a)

	r.V.X = b2UnpackLoW(tls, t1a, t2a)
	r.V.Y = b2UnpackHiW(tls, t1a, t2a)
	r.W = b2UnpackLoW(tls, t3a, t4a)
	r.Flags = b2UnpackHiW(tls, t3a, t4a)

	t1b := b2UnpackLoW(tls, b1b, b3b)
	t2b := b2UnpackLoW(tls, b2b, b4b)
	t3b := b2UnpackHiW(tls, b1b, b3b)
	t4b := b2UnpackHiW(tls, b2b, b4b)

	r.Dp.X = b2UnpackLoW(tls, t1b, t2b)
	r.Dp.Y = b2UnpackHiW(tls, t1b, t2b)
	r.Dq.C = b2UnpackLoW(tls, t3b, t4b)
	r.Dq.S = b2UnpackHiW(tls, t3b, t4b)

	return r
}

// This writes only the velocities back to the solver bodies
func b2ScatterBodies(tls *_Stack, states uintptr, indices uintptr, simdBody uintptr) {
	idx := (*[4]int32)(unsafe.Add(nil, indices))
	body := (*b2BodyStateW4)(unsafe.Add(nil, simdBody))

	store := func(index int32, vx, vy, w float32) {
		// zero means null
		if index == 0 {
			return
		}

		state := (*b2BodyState)(unsafe.Add(nil, states+uintptr(index-1)*unsafe.Sizeof(b2BodyState{})))
		if state.Flags&uint32(b2_dynamicFlag) == 0 {
			return
		}

		state.LinearVelocity.X = vx
		state.LinearVelocity.Y = vy
		state.AngularVelocity = w
	}

	// GetElem needs a constant lane, otherwise it is a function call
	store(idx[0], body.V.X.GetElem(0), body.V.Y.GetElem(0), body.W.GetElem(0))
	store(idx[1], body.V.X.GetElem(1), body.V.Y.GetElem(1), body.W.GetElem(1))
	store(idx[2], body.V.X.GetElem(2), body.V.Y.GetElem(2), body.W.GetElem(2))
	store(idx[3], body.V.X.GetElem(3), body.V.Y.GetElem(3), body.W.GetElem(3))
}
