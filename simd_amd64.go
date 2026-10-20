//go:build goexperiment.simd && go1.27 && amd64.v3

package b2

// [a0 b0 a1 b1]
func b2UnpackLoW(tls *_Stack, a, b b2FloatW4) b2FloatW4 {
	return a.ConcatPermuteScalars(0, 4, 1, 5, b)
}

// [a2 b2 a3 b3]
func b2UnpackHiW(tls *_Stack, a, b b2FloatW4) b2FloatW4 {
	return a.ConcatPermuteScalars(2, 6, 3, 7, b)
}
