//go:build goexperiment.simd && go1.27

package b2

import "simd/archsimd"

// wasm has no shuffle of two vectors. Swizzle the bytes of each input into its lanes,
// out of range indices (-1) give zero, and combine both.
var (
	unpackLoA = archsimd.LoadInt8x16Array(&[16]int8{0, 1, 2, 3, -1, -1, -1, -1, 4, 5, 6, 7, -1, -1, -1, -1})
	unpackLoB = archsimd.LoadInt8x16Array(&[16]int8{-1, -1, -1, -1, 0, 1, 2, 3, -1, -1, -1, -1, 4, 5, 6, 7})
	unpackHiA = archsimd.LoadInt8x16Array(&[16]int8{8, 9, 10, 11, -1, -1, -1, -1, 12, 13, 14, 15, -1, -1, -1, -1})
	unpackHiB = archsimd.LoadInt8x16Array(&[16]int8{-1, -1, -1, -1, 8, 9, 10, 11, -1, -1, -1, -1, 12, 13, 14, 15})
)

func swizzle(a b2FloatW4, indices archsimd.Int8x16) archsimd.Uint32x4 {
	return a.ToBits().ReshapeToUint8s().BitsToInt8().LookupOrZero(indices).ToBits().ReshapeToUint32s()
}

// [a0 b0 a1 b1]
func b2UnpackLoW(tls *_Stack, a, b b2FloatW4) b2FloatW4 {
	return swizzle(a, unpackLoA).Or(swizzle(b, unpackLoB)).BitsToFloat32()
}

// [a2 b2 a3 b3]
func b2UnpackHiW(tls *_Stack, a, b b2FloatW4) b2FloatW4 {
	return swizzle(a, unpackHiA).Or(swizzle(b, unpackHiB)).BitsToFloat32()
}
