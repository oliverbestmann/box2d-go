//go:build goexperiment.simd && go1.27

package b2

// [a0 b0 a1 b1]
func b2UnpackLoW(tls *_Stack, a, b b2FloatW4) b2FloatW4 {
	return a.ToBits().InterleaveLo(b.ToBits()).BitsToFloat32()
}

// [a2 b2 a3 b3]
func b2UnpackHiW(tls *_Stack, a, b b2FloatW4) b2FloatW4 {
	return a.ToBits().InterleaveHi(b.ToBits()).BitsToFloat32()
}
