// Code generated for linux/arm64 by 'ccgo -verify-types --package-name b2 -D BOX2D_DISABLE_SIMD --predef \n#include <stdbool.h>\n#include <stdint.h>\nbool __builtin___atomic_compare_exchange_n(int32_t *ptr, int32_t *expected, int32_t desired, bool weak, int32_t success_memorder, int32_t failure_memorder);\n --predef \n#include <stddef.h>\n#define aligned_alloc __builtin_aligned_alloc\nextern void* __builtin_aligned_alloc(size_t align, size_t size);\n --prefix-field ｆ -o ../box2d_c.go -I ../box2d-c/include/ ../box2d-c/src/aabb.c ../box2d-c/src/arena_allocator.c ../box2d-c/src/array.c ../box2d-c/src/bitset.c ../box2d-c/src/body.c ../box2d-c/src/broad_phase.c ../box2d-c/src/constraint_graph.c ../box2d-c/src/contact.c ../box2d-c/src/contact_solver.c ../box2d-c/src/core.c ../box2d-c/src/distance.c ../box2d-c/src/distance_joint.c ../box2d-c/src/dynamic_tree.c ../box2d-c/src/geometry.c ../box2d-c/src/hull.c ../box2d-c/src/id_pool.c ../box2d-c/src/island.c ../box2d-c/src/joint.c ../box2d-c/src/manifold.c ../box2d-c/src/math_functions.c ../box2d-c/src/motor_joint.c ../box2d-c/src/mouse_joint.c ../box2d-c/src/mover.c ../box2d-c/src/prismatic_joint.c ../box2d-c/src/revolute_joint.c ../box2d-c/src/sensor.c ../box2d-c/src/shape.c ../box2d-c/src/solver.c ../box2d-c/src/solver_set.c ../box2d-c/src/statics.c ../box2d-c/src/table.c ../box2d-c/src/timer.c ../box2d-c/src/types.c ../box2d-c/src/weld_joint.c ../box2d-c/src/wheel_joint.c ../box2d-c/src/world.c', DO NOT EDIT.

package b2

import (
	"math"
	"reflect"
	"unsafe"
)

var _ = math.Pi
var _ reflect.Type
var _ unsafe.Pointer

const _B2_API = "BOX2D_EXPORT"
const _B2_DEFAULT_CATEGORY_BITS = 1
const _B2_DEFAULT_MASK_BITS = "UINT64_MAX"
const _B2_HASH_INIT = 5381
const _B2_MAX_POLYGON_VERTICES = 8
const _B2_PI = 3.14159265359
const _BOX2D_DISABLE_SIMD = 1
const _DBL_DECIMAL_DIG = 17
const _DBL_DIG = 15
const _DBL_EPSILON = 2.22044604925031308085e-16
const _DBL_HAS_SUBNORM = 1
const _DBL_MANT_DIG = 53
const _DBL_MAX = 1.79769313486231570815e+308
const _DBL_MAX_10_EXP = 308
const _DBL_MAX_EXP = 1024
const _DBL_MIN = 2.22507385850720138309e-308
const _DBL_TRUE_MIN = 4.94065645841246544177e-324
const _DECIMAL_DIG = 17
const _FLT_DECIMAL_DIG = 9
const _FLT_DIG = 6
const _FLT_EPSILON = 1.1920928955078125e-07
const _FLT_EVAL_METHOD = 0
const _FLT_HAS_SUBNORM = 1
const _FLT_MANT_DIG = 24
const _FLT_MAX = 3.4028234663852886e+38
const _FLT_MAX_10_EXP = 38
const _FLT_MAX_EXP = 128
const _FLT_MIN = 1.17549435082228750797e-38
const _FLT_RADIX = 2
const _FLT_TRUE_MIN = 1.40129846432481707092e-45
const _FP_FAST_FMA = 1
const _FP_FAST_FMAF = 1
const _FP_ILOGB0 = "FP_ILOGBNAN"
const _FP_INFINITE = 1
const _FP_NAN = 0
const _FP_NORMAL = 4
const _FP_SUBNORMAL = 3
const _FP_ZERO = 2
const _HUGE = 3.40282346638528859812e+38
const _HUGE_VALF = "INFINITY"
const _INT16_MAX = 0x7fff
const _INT32_MAX = 0x7fffffff
const _INT64_MAX = 0x7fffffffffffffff
const _INT8_MAX = 0x7f
const _INTMAX_MAX = "INT64_MAX"
const _INTMAX_MIN = "INT64_MIN"
const _INTPTR_MAX = "INT64_MAX"
const _INTPTR_MIN = "INT64_MIN"
const _INT_FAST16_MAX = "INT32_MAX"
const _INT_FAST16_MIN = "INT32_MIN"
const _INT_FAST32_MAX = "INT32_MAX"
const _INT_FAST32_MIN = "INT32_MIN"
const _INT_FAST64_MAX = "INT64_MAX"
const _INT_FAST64_MIN = "INT64_MIN"
const _INT_FAST8_MAX = "INT8_MAX"
const _INT_FAST8_MIN = "INT8_MIN"
const _INT_LEAST16_MAX = "INT16_MAX"
const _INT_LEAST16_MIN = "INT16_MIN"
const _INT_LEAST32_MAX = "INT32_MAX"
const _INT_LEAST32_MIN = "INT32_MIN"
const _INT_LEAST64_MAX = "INT64_MAX"
const _INT_LEAST64_MIN = "INT64_MIN"
const _INT_LEAST8_MAX = "INT8_MAX"
const _INT_LEAST8_MIN = "INT8_MIN"
const _LDBL_DECIMAL_DIG = "DECIMAL_DIG"
const _LDBL_DIG = 15
const _LDBL_EPSILON = 2.22044604925031308085e-16
const _LDBL_HAS_SUBNORM = 1
const _LDBL_MANT_DIG = 53
const _LDBL_MAX = 1.79769313486231570815e+308
const _LDBL_MAX_10_EXP = 308
const _LDBL_MAX_EXP = 1024
const _LDBL_MIN = 2.22507385850720138309e-308
const _LDBL_TRUE_MIN = 4.94065645841246544177e-324
const _MATH_ERREXCEPT = 2
const _MATH_ERRNO = 1
const _M_1_PI = 0.31830988618379067154
const _M_2_PI = 0.63661977236758134308
const _M_2_SQRTPI = 1.12837916709551257390
const _M_E = 2.7182818284590452354
const _M_LN10 = 2.30258509299404568402
const _M_LN2 = 0.69314718055994530942
const _M_LOG10E = 0.43429448190325182765
const _M_LOG2E = 1.4426950408889634074
const _M_PI = 3.14159265358979323846
const _M_PI_2 = 1.57079632679489661923
const _M_PI_4 = 0.78539816339744830962
const _M_SQRT1_2 = 0.70710678118654752440
const _M_SQRT2 = 1.41421356237309504880
const _PTRDIFF_MAX = "INT64_MAX"
const _PTRDIFF_MIN = "INT64_MIN"
const _SIG_ATOMIC_MAX = "INT32_MAX"
const _SIG_ATOMIC_MIN = "INT32_MIN"
const _SIZE_MAX = "UINT64_MAX"
const _UINT16_MAX = 0xffff
const _UINT32_MAX = "0xffffffffu"
const _UINT64_MAX = "0xffffffffffffffffu"
const _UINT8_MAX = 0xff
const _UINTMAX_MAX = "UINT64_MAX"
const _UINTPTR_MAX = "UINT64_MAX"
const _UINT_FAST16_MAX = "UINT32_MAX"
const _UINT_FAST32_MAX = "UINT32_MAX"
const _UINT_FAST64_MAX = "UINT64_MAX"
const _UINT_FAST8_MAX = "UINT8_MAX"
const _UINT_LEAST16_MAX = "UINT16_MAX"
const _UINT_LEAST32_MAX = "UINT32_MAX"
const _UINT_LEAST64_MAX = "UINT64_MAX"
const _UINT_LEAST8_MAX = "UINT8_MAX"
const _WINT_MAX = "UINT32_MAX"
const _WINT_MIN = 0
const _GNU_SOURCE = 1
const _LP64 = 1
const _STDC_PREDEF_H = 1
const __AARCH64EL__ = 1
const __AARCH64_CMODEL_SMALL__ = 1
const __ARM_64BIT_STATE = 1
const __ARM_ALIGN_MAX_PWR = 28
const __ARM_ALIGN_MAX_STACK_PWR = 16
const __ARM_ARCH = 8
const __ARM_ARCH_8A = 1
const __ARM_ARCH_ISA_A64 = 1
const __ARM_ARCH_PROFILE = 65
const __ARM_BUILDATTR64_FV = 65
const __ARM_FEATURE_CLZ = 1
const __ARM_FEATURE_FMA = 1
const __ARM_FEATURE_IDIV = 1
const __ARM_FEATURE_NUMERIC_MAXMIN = 1
const __ARM_FEATURE_UNALIGNED = 1
const __ARM_FP = 14
const __ARM_FP16_ARGS = 1
const __ARM_FP16_FORMAT_IEEE = 1
const __ARM_NEON = 1
const __ARM_NEON_SVE_BRIDGE = 1
const __ARM_PCS_AAPCS64 = 1
const __ARM_SIZEOF_MINIMAL_ENUM = 4
const __ARM_SIZEOF_WCHAR_T = 4
const __ARM_STATE_ZA = 1
const __ARM_STATE_ZT0 = 1
const __ATOMIC_ACQUIRE = 2
const __ATOMIC_ACQ_REL = 4
const __ATOMIC_CONSUME = 1
const __ATOMIC_RELAXED = 0
const __ATOMIC_RELEASE = 3
const __ATOMIC_SEQ_CST = 5
const __BFLT16_DECIMAL_DIG__ = 4
const __BFLT16_DENORM_MIN__ = "9.18354961579912115600575419704879436e-41B"
const __BFLT16_DIG__ = 2
const __BFLT16_EPSILON__ = "7.81250000000000000000000000000000000e-3B"
const __BFLT16_HAS_DENORM__ = 1
const __BFLT16_HAS_INFINITY__ = 1
const __BFLT16_HAS_QUIET_NAN__ = 1
const __BFLT16_IS_IEC_60559__ = 0
const __BFLT16_MANT_DIG__ = 8
const __BFLT16_MAX_10_EXP__ = 38
const __BFLT16_MAX_EXP__ = 128
const __BFLT16_MAX__ = "3.38953138925153547590470800371487867e+38B"
const __BFLT16_MIN__ = "1.17549435082228750796873653722224568e-38B"
const __BFLT16_NORM_MAX__ = "3.38953138925153547590470800371487867e+38B"
const __BIGGEST_ALIGNMENT__ = 16
const __BIG_ENDIAN = 4321
const __BITINT_MAXWIDTH__ = 65535
const __BYTE_ORDER = 1234
const __BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __CCGO__ = 1
const __CHAR_BIT__ = 8
const __CHAR_UNSIGNED__ = 1
const __DBL_DECIMAL_DIG__ = 17
const __DBL_DIG__ = 15
const __DBL_HAS_DENORM__ = 1
const __DBL_HAS_INFINITY__ = 1
const __DBL_HAS_QUIET_NAN__ = 1
const __DBL_IS_IEC_60559__ = 1
const __DBL_MANT_DIG__ = 53
const __DBL_MAX_10_EXP__ = 308
const __DBL_MAX_EXP__ = 1024
const __DEC128_EPSILON__ = 1e-33
const __DEC128_MANT_DIG__ = 34
const __DEC128_MAX_EXP__ = 6145
const __DEC128_MAX__ = "9.999999999999999999999999999999999E6144"
const __DEC128_MIN__ = 1e-6143
const __DEC128_SUBNORMAL_MIN__ = 0.000000000000000000000000000000001e-6143
const __DEC32_EPSILON__ = 1e-6
const __DEC32_MANT_DIG__ = 7
const __DEC32_MAX_EXP__ = 97
const __DEC32_MAX__ = 9.999999e96
const __DEC32_MIN__ = 1e-95
const __DEC32_SUBNORMAL_MIN__ = 0.000001e-95
const __DEC64X_EPSILON__ = "1E-33D64x"
const __DEC64X_MANT_DIG__ = 34
const __DEC64X_MAX_EXP__ = 6145
const __DEC64X_MAX__ = "9.999999999999999999999999999999999E6144D64x"
const __DEC64X_MIN__ = "1E-6143D64x"
const __DEC64X_SUBNORMAL_MIN__ = "0.000000000000000000000000000000001E-6143D64x"
const __DEC64_EPSILON__ = 1e-15
const __DEC64_MANT_DIG__ = 16
const __DEC64_MAX_EXP__ = 385
const __DEC64_MAX__ = "9.999999999999999E384"
const __DEC64_MIN__ = 1e-383
const __DEC64_SUBNORMAL_MIN__ = 0.000000000000001e-383
const __DECIMAL_BID_FORMAT__ = 1
const __DECIMAL_DIG__ = 36
const __DEC_EVAL_METHOD__ = 2
const __ELF__ = 1
const __FINITE_MATH_ONLY__ = 0
const __FLOAT_WORD_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __FLT128_DECIMAL_DIG__ = 36
const __FLT128_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const __FLT128_DIG__ = 33
const __FLT128_EPSILON__ = 1.92592994438723585305597794258492732e-34
const __FLT128_HAS_DENORM__ = 1
const __FLT128_HAS_INFINITY__ = 1
const __FLT128_HAS_QUIET_NAN__ = 1
const __FLT128_IS_IEC_60559__ = 1
const __FLT128_MANT_DIG__ = 113
const __FLT128_MAX_10_EXP__ = 4932
const __FLT128_MAX_EXP__ = 16384
const __FLT128_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT128_MIN__ = 3.36210314311209350626267781732175260e-4932
const __FLT128_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT16_DECIMAL_DIG__ = 5
const __FLT16_DENORM_MIN__ = 5.96046447753906250000000000000000000e-8
const __FLT16_DIG__ = 3
const __FLT16_EPSILON__ = 9.76562500000000000000000000000000000e-4
const __FLT16_HAS_DENORM__ = 1
const __FLT16_HAS_INFINITY__ = 1
const __FLT16_HAS_QUIET_NAN__ = 1
const __FLT16_IS_IEC_60559__ = 1
const __FLT16_MANT_DIG__ = 11
const __FLT16_MAX_10_EXP__ = 4
const __FLT16_MAX_EXP__ = 16
const __FLT16_MAX__ = 6.55040000000000000000000000000000000e+4
const __FLT16_MIN__ = 6.10351562500000000000000000000000000e-5
const __FLT16_NORM_MAX__ = 6.55040000000000000000000000000000000e+4
const __FLT32X_DECIMAL_DIG__ = 17
const __FLT32X_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __FLT32X_DIG__ = 15
const __FLT32X_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __FLT32X_HAS_DENORM__ = 1
const __FLT32X_HAS_INFINITY__ = 1
const __FLT32X_HAS_QUIET_NAN__ = 1
const __FLT32X_IS_IEC_60559__ = 1
const __FLT32X_MANT_DIG__ = 53
const __FLT32X_MAX_10_EXP__ = 308
const __FLT32X_MAX_EXP__ = 1024
const __FLT32X_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT32X_MIN__ = 2.22507385850720138309023271733240406e-308
const __FLT32X_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT32_DECIMAL_DIG__ = 9
const __FLT32_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const __FLT32_DIG__ = 6
const __FLT32_EPSILON__ = 1.19209289550781250000000000000000000e-7
const __FLT32_HAS_DENORM__ = 1
const __FLT32_HAS_INFINITY__ = 1
const __FLT32_HAS_QUIET_NAN__ = 1
const __FLT32_IS_IEC_60559__ = 1
const __FLT32_MANT_DIG__ = 24
const __FLT32_MAX_10_EXP__ = 38
const __FLT32_MAX_EXP__ = 128
const __FLT32_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT32_MIN__ = 1.17549435082228750796873653722224568e-38
const __FLT32_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT64X_DECIMAL_DIG__ = 36
const __FLT64X_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const __FLT64X_DIG__ = 33
const __FLT64X_EPSILON__ = 1.92592994438723585305597794258492732e-34
const __FLT64X_HAS_DENORM__ = 1
const __FLT64X_HAS_INFINITY__ = 1
const __FLT64X_HAS_QUIET_NAN__ = 1
const __FLT64X_IS_IEC_60559__ = 1
const __FLT64X_MANT_DIG__ = 113
const __FLT64X_MAX_10_EXP__ = 4932
const __FLT64X_MAX_EXP__ = 16384
const __FLT64X_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT64X_MIN__ = 3.36210314311209350626267781732175260e-4932
const __FLT64X_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT64_DECIMAL_DIG__ = 17
const __FLT64_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __FLT64_DIG__ = 15
const __FLT64_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __FLT64_HAS_DENORM__ = 1
const __FLT64_HAS_INFINITY__ = 1
const __FLT64_HAS_QUIET_NAN__ = 1
const __FLT64_IS_IEC_60559__ = 1
const __FLT64_MANT_DIG__ = 53
const __FLT64_MAX_10_EXP__ = 308
const __FLT64_MAX_EXP__ = 1024
const __FLT64_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT64_MIN__ = 2.22507385850720138309023271733240406e-308
const __FLT64_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT_DECIMAL_DIG__ = 9
const __FLT_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const __FLT_DIG__ = 6
const __FLT_EPSILON__ = 1.19209289550781250000000000000000000e-7
const __FLT_EVAL_METHOD_C99__ = 0
const __FLT_EVAL_METHOD_TS_18661_3__ = 0
const __FLT_EVAL_METHOD__ = 0
const __FLT_HAS_DENORM__ = 1
const __FLT_HAS_INFINITY__ = 1
const __FLT_HAS_QUIET_NAN__ = 1
const __FLT_IS_IEC_60559__ = 1
const __FLT_MANT_DIG__ = 24
const __FLT_MAX_10_EXP__ = 38
const __FLT_MAX_EXP__ = 128
const __FLT_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT_MIN__ = 1.17549435082228750796873653722224568e-38
const __FLT_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT_RADIX__ = 2
const __FP_FAST_FMA = 1
const __FP_FAST_FMAF = 1
const __FP_FAST_FMAF32 = 1
const __FP_FAST_FMAF32x = 1
const __FP_FAST_FMAF64 = 1
const __FUNCTION__ = "__func__"
const __GCC_ASM_FLAG_OUTPUTS__ = 1
const __GCC_ATOMIC_BOOL_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR16_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR32_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR8_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR_LOCK_FREE = 2
const __GCC_ATOMIC_INT_LOCK_FREE = 2
const __GCC_ATOMIC_LLONG_LOCK_FREE = 2
const __GCC_ATOMIC_LONG_LOCK_FREE = 2
const __GCC_ATOMIC_POINTER_LOCK_FREE = 2
const __GCC_ATOMIC_SHORT_LOCK_FREE = 2
const __GCC_ATOMIC_TEST_AND_SET_TRUEVAL = 1
const __GCC_ATOMIC_WCHAR_T_LOCK_FREE = 2
const __GCC_CONSTRUCTIVE_SIZE = 64
const __GCC_DESTRUCTIVE_SIZE = 256
const __GCC_HAVE_DWARF2_CFI_ASM = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_1 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_16 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const __GCC_IEC_559 = 2
const __GCC_IEC_559_COMPLEX = 2
const __GNUC_EXECUTION_CHARSET_NAME = "UTF-8"
const __GNUC_MINOR__ = 2
const __GNUC_PATCHLEVEL__ = 0
const __GNUC_STDC_INLINE__ = 1
const __GNUC_WIDE_EXECUTION_CHARSET_NAME = "UTF-32LE"
const __GNUC__ = 16
const __GXX_ABI_VERSION = 1021
const __HAVE_FUNCTION_MULTI_VERSIONING = 1
const __HAVE_SPECULATION_SAFE_VALUE = 1
const __INT16_MAX__ = 0x7fff
const __INT32_MAX__ = 0x7fffffff
const __INT32_TYPE__ = "int"
const __INT64_MAX__ = 0x7fffffffffffffff
const __INT8_MAX__ = 0x7f
const __INTMAX_MAX__ = 0x7fffffffffffffff
const __INTMAX_WIDTH__ = 64
const __INTPTR_MAX__ = 0x7fffffffffffffff
const __INTPTR_WIDTH__ = 64
const __INT_FAST16_MAX__ = 0x7fffffffffffffff
const __INT_FAST16_WIDTH__ = 64
const __INT_FAST32_MAX__ = 0x7fffffffffffffff
const __INT_FAST32_WIDTH__ = 64
const __INT_FAST64_MAX__ = 0x7fffffffffffffff
const __INT_FAST64_WIDTH__ = 64
const __INT_FAST8_MAX__ = 0x7f
const __INT_FAST8_WIDTH__ = 8
const __INT_LEAST16_MAX__ = 0x7fff
const __INT_LEAST16_WIDTH__ = 16
const __INT_LEAST32_MAX__ = 0x7fffffff
const __INT_LEAST32_TYPE__ = "int"
const __INT_LEAST32_WIDTH__ = 32
const __INT_LEAST64_MAX__ = 0x7fffffffffffffff
const __INT_LEAST64_WIDTH__ = 64
const __INT_LEAST8_MAX__ = 0x7f
const __INT_LEAST8_WIDTH__ = 8
const __INT_MAX__ = 0x7fffffff
const __INT_WIDTH__ = 32
const __LDBL_DECIMAL_DIG__ = 36
const __LDBL_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const __LDBL_DIG__ = 33
const __LDBL_EPSILON__ = 1.92592994438723585305597794258492732e-34
const __LDBL_HAS_DENORM__ = 1
const __LDBL_HAS_INFINITY__ = 1
const __LDBL_HAS_QUIET_NAN__ = 1
const __LDBL_IS_IEC_60559__ = 1
const __LDBL_MANT_DIG__ = 113
const __LDBL_MAX_10_EXP__ = 4932
const __LDBL_MAX_EXP__ = 16384
const __LDBL_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __LDBL_MIN__ = 3.36210314311209350626267781732175260e-4932
const __LDBL_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __LITTLE_ENDIAN = 1234
const __LONG_LONG_MAX__ = 0x7fffffffffffffff
const __LONG_LONG_WIDTH__ = 64
const __LONG_MAX = 0x7fffffffffffffff
const __LONG_MAX__ = 0x7fffffffffffffff
const __LONG_WIDTH__ = 64
const __LP64__ = 1
const __NO_INLINE__ = 1
const __ORDER_BIG_ENDIAN__ = 4321
const __ORDER_LITTLE_ENDIAN__ = 1234
const __ORDER_PDP_ENDIAN__ = 3412
const __PIC__ = 2
const __PIE__ = 2
const __PRAGMA_REDEFINE_EXTNAME = 1
const __PRETTY_FUNCTION__ = "__func__"
const __PTRDIFF_MAX__ = 0x7fffffffffffffff
const __PTRDIFF_WIDTH__ = 64
const __SCHAR_MAX__ = 0x7f
const __SCHAR_WIDTH__ = 8
const __SHRT_MAX__ = 0x7fff
const __SHRT_WIDTH__ = 16
const __SIG_ATOMIC_MAX__ = 0x7fffffff
const __SIG_ATOMIC_TYPE__ = "int"
const __SIG_ATOMIC_WIDTH__ = 32
const __SIZEOF_DOUBLE__ = 8
const __SIZEOF_FLOAT__ = 4
const __SIZEOF_INT128__ = 16
const __SIZEOF_INT__ = 4
const __SIZEOF_LONG_DOUBLE__ = 8
const __SIZEOF_LONG_LONG__ = 8
const __SIZEOF_LONG__ = 8
const __SIZEOF_POINTER__ = 8
const __SIZEOF_PTRDIFF_T__ = 8
const __SIZEOF_SHORT__ = 2
const __SIZEOF_SIZE_T__ = 8
const __SIZEOF_WCHAR_T__ = 4
const __SIZEOF_WINT_T__ = 4
const __SIZE_MAX__ = 0xffffffffffffffff
const __SIZE_WIDTH__ = 64
const __STDC_EMBED_EMPTY__ = 2
const __STDC_EMBED_FOUND__ = 1
const __STDC_EMBED_NOT_FOUND__ = 0
const __STDC_HOSTED__ = 1
const __STDC_IEC_559_COMPLEX__ = 1
const __STDC_IEC_559__ = 1
const __STDC_IEC_60559_BFP__ = 201404
const __STDC_IEC_60559_COMPLEX__ = 201404
const __STDC_ISO_10646__ = 201706
const __STDC_UTF_16__ = 1
const __STDC_UTF_32__ = 1
const __STDC_VERSION__ = 202311
const __STDC__ = 1
const __UINT16_MAX__ = 0xffff
const __UINT32_MAX__ = 0xffffffff
const __UINT64_MAX__ = 0xffffffffffffffff
const __UINT8_MAX__ = 0xff
const __UINTMAX_MAX__ = 0xffffffffffffffff
const __UINTPTR_MAX__ = 0xffffffffffffffff
const __UINT_FAST16_MAX__ = 0xffffffffffffffff
const __UINT_FAST32_MAX__ = 0xffffffffffffffff
const __UINT_FAST64_MAX__ = 0xffffffffffffffff
const __UINT_FAST8_MAX__ = 0xff
const __UINT_LEAST16_MAX__ = 0xffff
const __UINT_LEAST32_MAX__ = 0xffffffff
const __UINT_LEAST64_MAX__ = 0xffffffffffffffff
const __UINT_LEAST8_MAX__ = 0xff
const __USE_TIME_BITS64 = 1
const __VERSION__ = "16.2.0"
const __WCHAR_MAX__ = 0xffffffff
const __WCHAR_MIN__ = 0
const __WCHAR_WIDTH__ = 32
const __WINT_MAX__ = 0xffffffff
const __WINT_MIN__ = 0
const __WINT_WIDTH__ = 32
const __aarch64__ = 1
const __bool_true_false_are_defined = 1
const __gnu_linux__ = 1
const __inline = "inline"
const __linux = 1
const __linux__ = 1
const __pic__ = 2
const __pie__ = 2
const __restrict = "restrict"
const __restrict_arr = "restrict"
const __unix = 1
const __unix__ = 1
const aligned_alloc = "__builtin_aligned_alloc"
const bool1 = "_Bool"
const false1 = 0
const linux = 1
const math_errhandling = 2
const true1 = 1
const unix = 1

type uintptr_t = uint64

type intptr_t = int64

type int8_t = int8

type int16_t = int16

type int32_t = int32

type int64_t = int64

type intmax_t = int64

type uint8_t = uint8

type uint16_t = uint16

type uint32_t = uint32

type uint64_t = uint64

type uintmax_t = uint64

type int_fast8_t = int8

type int_fast64_t = int64

type int_least8_t = int8

type int_least16_t = int16

type int_least32_t = int32

type int_least64_t = int64

type uint_fast8_t = uint8

type uint_fast64_t = uint64

type uint_least8_t = uint8

type uint_least16_t = uint16

type uint_least32_t = uint32

type uint_least64_t = uint64

type int_fast16_t = int32

type int_fast32_t = int32

type uint_fast16_t = uint32

type uint_fast32_t = uint32

type wchar_t = uint32

type max_align_t struct {
	ｆ__ll int64
	ｆ__ld float64
}

type size_t = uint64

type ptrdiff_t = int64

/* C23 keywords, no longer defined as macros by <stdbool.h> and <stddef.h>. */

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = uint32

type __predefined_ptrdiff_t = int64

type Version struct {
	Major    int32
	Minor    int32
	Revision int32
}

type float_t = float32

type double_t = float64

type Vec2 struct {
	X float32
	Y float32
}

type CosSin struct {
	Cosine float32
	Sine   float32
}

type Rot struct {
	C float32
	S float32
}

type Transform struct {
	P Vec2
	Q Rot
}

type Mat22 struct {
	Cx Vec2
	Cy Vec2
}

type AABB struct {
	LowerBound Vec2
	UpperBound Vec2
}

type Plane struct {
	Normal Vec2
	Offset float32
}

type SimplexCache struct {
	Count  uint16
	IndexA [3]uint8_t
	IndexB [3]uint8_t
}

type Hull struct {
	Points [8]Vec2
	Count  int32
}

type RayCastInput struct {
	Origin      Vec2
	Translation Vec2
	MaxFraction float32
}

type ShapeProxy struct {
	Points [8]Vec2
	Count  int32
	Radius float32
}

type ShapeCastInput struct {
	Proxy       ShapeProxy
	Translation Vec2
	MaxFraction float32
	CanEncroach uint8
}

type CastOutput struct {
	Normal     Vec2
	Point      Vec2
	Fraction   float32
	Iterations int32
	Hit        uint8
}

type MassData struct {
	Mass              float32
	Center            Vec2
	RotationalInertia float32
}

type Circle struct {
	Center Vec2
	Radius float32
}

type Capsule struct {
	Center1 Vec2
	Center2 Vec2
	Radius  float32
}

type Polygon struct {
	Vertices [8]Vec2
	Normals  [8]Vec2
	Centroid Vec2
	Radius   float32
	Count    int32
}

type Segment struct {
	Point1 Vec2
	Point2 Vec2
}

type ChainSegment struct {
	Ghost1  Vec2
	Segment Segment
	Ghost2  Vec2
	ChainId int32
}

type SegmentDistanceResult struct {
	Closest1        Vec2
	Closest2        Vec2
	Fraction1       float32
	Fraction2       float32
	DistanceSquared float32
}

type DistanceInput struct {
	ProxyA     ShapeProxy
	ProxyB     ShapeProxy
	TransformA Transform
	TransformB Transform
	UseRadii   uint8
}

type DistanceOutput struct {
	PointA       Vec2
	PointB       Vec2
	Normal       Vec2
	Distance     float32
	Iterations   int32
	SimplexCount int32
}

type SimplexVertex struct {
	WA     Vec2
	WB     Vec2
	W      Vec2
	A      float32
	IndexA int32
	IndexB int32
}

type Simplex struct {
	V1    SimplexVertex
	V2    SimplexVertex
	V3    SimplexVertex
	Count int32
}

type ShapeCastPairInput struct {
	ProxyA       ShapeProxy
	ProxyB       ShapeProxy
	TransformA   Transform
	TransformB   Transform
	TranslationB Vec2
	MaxFraction  float32
	CanEncroach  uint8
}

type Sweep struct {
	LocalCenter Vec2
	C1          Vec2
	C2          Vec2
	Q1          Rot
	Q2          Rot
}

type TOIInput struct {
	ProxyA      ShapeProxy
	ProxyB      ShapeProxy
	SweepA      Sweep
	SweepB      Sweep
	MaxFraction float32
}

type b2TOIState1 = int32

type TOIState = int32

const b2_toiStateUnknown = 0
const b2_toiStateFailed = 1
const b2_toiStateOverlapped = 2
const b2_toiStateHit = 3
const b2_toiStateSeparated = 4

type TOIOutput struct {
	State    TOIState
	Fraction float32
}

type ManifoldPoint struct {
	Point              Vec2
	AnchorA            Vec2
	AnchorB            Vec2
	Separation         float32
	NormalImpulse      float32
	TangentImpulse     float32
	TotalNormalImpulse float32
	NormalVelocity     float32
	Id                 uint16
	Persisted          uint8
}

type Manifold struct {
	Normal         Vec2
	RollingImpulse float32
	Points         [2]ManifoldPoint
	PointCount     int32
}

type DynamicTree struct {
	Nodes           uintptr
	Root            int32
	NodeCount       int32
	NodeCapacity    int32
	FreeList        int32
	ProxyCount      int32
	LeafIndices     uintptr
	LeafBoxes       uintptr
	LeafCenters     uintptr
	BinIndices      uintptr
	RebuildCapacity int32
}

type TreeStats struct {
	NodeVisits int32
	LeafVisits int32
}

type PlaneResult struct {
	Plane Plane
	Point Vec2
	Hit   uint8
}

type CollisionPlane struct {
	Plane        Plane
	PushLimit    float32
	Push         float32
	ClipVelocity uint8
}

type PlaneSolverResult struct {
	Translation    Vec2
	IterationCount int32
}

type WorldId struct {
	Index1     uint16
	Generation uint16
}

type BodyId struct {
	Index1     int32
	World0     uint16
	Generation uint16
}

type ShapeId struct {
	Index1     int32
	World0     uint16
	Generation uint16
}

type ChainId struct {
	Index1     int32
	World0     uint16
	Generation uint16
}

type JointId struct {
	Index1     int32
	World0     uint16
	Generation uint16
}

type RayResult struct {
	ShapeId    ShapeId
	Point      Vec2
	Normal     Vec2
	Fraction   float32
	NodeVisits int32
	LeafVisits int32
	Hit        uint8
}

type WorldDef struct {
	Gravity              Vec2
	RestitutionThreshold float32
	HitEventThreshold    float32
	ContactHertz         float32
	ContactDampingRatio  float32
	MaxContactPushSpeed  float32
	MaximumLinearSpeed   float32
	FrictionCallback     uintptr
	RestitutionCallback  uintptr
	EnableSleep          uint8
	EnableContinuous     uint8
	WorkerCount          int32
	EnqueueTask          uintptr
	FinishTask           uintptr
	UserTaskContext      uintptr
	UserData             uintptr
	InternalValue        int32
}

type b2BodyType1 = int32

type BodyType = int32

const b2_staticBody = 0
const b2_kinematicBody = 1
const b2_dynamicBody = 2
const b2_bodyTypeCount = 3

type BodyDef struct {
	Type1             BodyType
	Position          Vec2
	Rotation          Rot
	LinearVelocity    Vec2
	AngularVelocity   float32
	LinearDamping     float32
	AngularDamping    float32
	GravityScale      float32
	SleepThreshold    float32
	Name              uintptr
	UserData          uintptr
	EnableSleep       uint8
	IsAwake           uint8
	FixedRotation     uint8
	IsBullet          uint8
	IsEnabled         uint8
	AllowFastRotation uint8
	InternalValue     int32
}

type Filter struct {
	CategoryBits uint64
	MaskBits     uint64
	GroupIndex   int32
}

type QueryFilter struct {
	CategoryBits uint64
	MaskBits     uint64
}

type b2ShapeType1 = int32

type ShapeType = int32

const b2_circleShape = 0
const b2_capsuleShape = 1
const b2_segmentShape = 2
const b2_polygonShape = 3
const b2_chainSegmentShape = 4
const b2_shapeTypeCount = 5

type SurfaceMaterial struct {
	Friction          float32
	Restitution       float32
	RollingResistance float32
	TangentSpeed      float32
	UserMaterialId    int32
	CustomColor       uint32
}

type ShapeDef struct {
	UserData              uintptr
	Material              SurfaceMaterial
	Density               float32
	Filter                Filter
	IsSensor              uint8
	EnableSensorEvents    uint8
	EnableContactEvents   uint8
	EnableHitEvents       uint8
	EnablePreSolveEvents  uint8
	InvokeContactCreation uint8
	UpdateBodyMass        uint8
	InternalValue         int32
}

type ChainDef struct {
	UserData           uintptr
	Points             uintptr
	Count              int32
	Materials          uintptr
	MaterialCount      int32
	Filter             Filter
	IsLoop             uint8
	EnableSensorEvents uint8
	InternalValue      int32
}

type Profile struct {
	Step                float32
	Pairs               float32
	Collide             float32
	Solve               float32
	MergeIslands        float32
	PrepareStages       float32
	SolveConstraints    float32
	PrepareConstraints  float32
	IntegrateVelocities float32
	WarmStart           float32
	SolveImpulses       float32
	IntegratePositions  float32
	RelaxImpulses       float32
	ApplyRestitution    float32
	StoreImpulses       float32
	SplitIslands        float32
	Transforms          float32
	HitEvents           float32
	Refit               float32
	Bullets             float32
	SleepIslands        float32
	Sensors             float32
}

type Counters struct {
	BodyCount        int32
	ShapeCount       int32
	ContactCount     int32
	JointCount       int32
	IslandCount      int32
	StackUsed        int32
	StaticTreeHeight int32
	TreeHeight       int32
	ByteCount        int32
	TaskCount        int32
	ColorCounts      [12]int32
}

type b2JointType1 = int32

type JointType = int32

const b2_distanceJoint = 0
const b2_filterJoint = 1
const b2_motorJoint = 2
const b2_mouseJoint = 3
const b2_prismaticJoint = 4
const b2_revoluteJoint = 5
const b2_weldJoint = 6
const b2_wheelJoint = 7

type DistanceJointDef struct {
	BodyIdA          BodyId
	BodyIdB          BodyId
	LocalAnchorA     Vec2
	LocalAnchorB     Vec2
	Length           float32
	EnableSpring     uint8
	Hertz            float32
	DampingRatio     float32
	EnableLimit      uint8
	MinLength        float32
	MaxLength        float32
	EnableMotor      uint8
	MaxMotorForce    float32
	MotorSpeed       float32
	CollideConnected uint8
	UserData         uintptr
	InternalValue    int32
}

type MotorJointDef struct {
	BodyIdA          BodyId
	BodyIdB          BodyId
	LinearOffset     Vec2
	AngularOffset    float32
	MaxForce         float32
	MaxTorque        float32
	CorrectionFactor float32
	CollideConnected uint8
	UserData         uintptr
	InternalValue    int32
}

type MouseJointDef struct {
	BodyIdA          BodyId
	BodyIdB          BodyId
	Target           Vec2
	Hertz            float32
	DampingRatio     float32
	MaxForce         float32
	CollideConnected uint8
	UserData         uintptr
	InternalValue    int32
}

type FilterJointDef struct {
	BodyIdA       BodyId
	BodyIdB       BodyId
	UserData      uintptr
	InternalValue int32
}

type PrismaticJointDef struct {
	BodyIdA           BodyId
	BodyIdB           BodyId
	LocalAnchorA      Vec2
	LocalAnchorB      Vec2
	LocalAxisA        Vec2
	ReferenceAngle    float32
	TargetTranslation float32
	EnableSpring      uint8
	Hertz             float32
	DampingRatio      float32
	EnableLimit       uint8
	LowerTranslation  float32
	UpperTranslation  float32
	EnableMotor       uint8
	MaxMotorForce     float32
	MotorSpeed        float32
	CollideConnected  uint8
	UserData          uintptr
	InternalValue     int32
}

type RevoluteJointDef struct {
	BodyIdA          BodyId
	BodyIdB          BodyId
	LocalAnchorA     Vec2
	LocalAnchorB     Vec2
	ReferenceAngle   float32
	TargetAngle      float32
	EnableSpring     uint8
	Hertz            float32
	DampingRatio     float32
	EnableLimit      uint8
	LowerAngle       float32
	UpperAngle       float32
	EnableMotor      uint8
	MaxMotorTorque   float32
	MotorSpeed       float32
	DrawSize         float32
	CollideConnected uint8
	UserData         uintptr
	InternalValue    int32
}

type WeldJointDef struct {
	BodyIdA             BodyId
	BodyIdB             BodyId
	LocalAnchorA        Vec2
	LocalAnchorB        Vec2
	ReferenceAngle      float32
	LinearHertz         float32
	AngularHertz        float32
	LinearDampingRatio  float32
	AngularDampingRatio float32
	CollideConnected    uint8
	UserData            uintptr
	InternalValue       int32
}

type WheelJointDef struct {
	BodyIdA          BodyId
	BodyIdB          BodyId
	LocalAnchorA     Vec2
	LocalAnchorB     Vec2
	LocalAxisA       Vec2
	EnableSpring     uint8
	Hertz            float32
	DampingRatio     float32
	EnableLimit      uint8
	LowerTranslation float32
	UpperTranslation float32
	EnableMotor      uint8
	MaxMotorTorque   float32
	MotorSpeed       float32
	CollideConnected uint8
	UserData         uintptr
	InternalValue    int32
}

type ExplosionDef struct {
	MaskBits         uint64
	Position         Vec2
	Radius           float32
	Falloff          float32
	ImpulsePerLength float32
}

type SensorBeginTouchEvent struct {
	SensorShapeId  ShapeId
	VisitorShapeId ShapeId
}

type SensorEndTouchEvent struct {
	SensorShapeId  ShapeId
	VisitorShapeId ShapeId
}

type b2SensorEvents struct {
	BeginEvents uintptr
	EndEvents   uintptr
	BeginCount  int32
	EndCount    int32
}

type ContactBeginTouchEvent struct {
	ShapeIdA ShapeId
	ShapeIdB ShapeId
	Manifold Manifold
}

type ContactEndTouchEvent struct {
	ShapeIdA ShapeId
	ShapeIdB ShapeId
}

type ContactHitEvent struct {
	ShapeIdA      ShapeId
	ShapeIdB      ShapeId
	Point         Vec2
	Normal        Vec2
	ApproachSpeed float32
}

type b2ContactEvents struct {
	BeginEvents uintptr
	EndEvents   uintptr
	HitEvents   uintptr
	BeginCount  int32
	EndCount    int32
	HitCount    int32
}

type BodyMoveEvent struct {
	Transform  Transform
	BodyId     BodyId
	UserData   uintptr
	FellAsleep uint8
}

type b2BodyEvents struct {
	MoveEvents uintptr
	MoveCount  int32
}

type ContactData struct {
	ShapeIdA ShapeId
	ShapeIdB ShapeId
	Manifold Manifold
}

type b2HexColor1 = int32

type HexColor = int32

const b2_colorAliceBlue = 15792383
const b2_colorAntiqueWhite = 16444375
const b2_colorAqua = 65535
const b2_colorAquamarine = 8388564
const b2_colorAzure = 15794175
const b2_colorBeige = 16119260
const b2_colorBisque = 16770244
const b2_colorBlack = 0
const b2_colorBlanchedAlmond = 16772045
const b2_colorBlue = 255
const b2_colorBlueViolet = 9055202
const b2_colorBrown = 10824234
const b2_colorBurlywood = 14596231
const b2_colorCadetBlue = 6266528
const b2_colorChartreuse = 8388352
const b2_colorChocolate = 13789470
const b2_colorCoral = 16744272
const b2_colorCornflowerBlue = 6591981
const b2_colorCornsilk = 16775388
const b2_colorCrimson = 14423100
const b2_colorCyan = 65535
const b2_colorDarkBlue = 139
const b2_colorDarkCyan = 35723
const b2_colorDarkGoldenRod = 12092939
const b2_colorDarkGray = 11119017
const b2_colorDarkGreen = 25600
const b2_colorDarkKhaki = 12433259
const b2_colorDarkMagenta = 9109643
const b2_colorDarkOliveGreen = 5597999
const b2_colorDarkOrange = 16747520
const b2_colorDarkOrchid = 10040012
const b2_colorDarkRed = 9109504
const b2_colorDarkSalmon = 15308410
const b2_colorDarkSeaGreen = 9419919
const b2_colorDarkSlateBlue = 4734347
const b2_colorDarkSlateGray = 3100495
const b2_colorDarkTurquoise = 52945
const b2_colorDarkViolet = 9699539
const b2_colorDeepPink = 16716947
const b2_colorDeepSkyBlue = 49151
const b2_colorDimGray = 6908265
const b2_colorDodgerBlue = 2003199
const b2_colorFireBrick = 11674146
const b2_colorFloralWhite = 16775920
const b2_colorForestGreen = 2263842
const b2_colorFuchsia = 16711935
const b2_colorGainsboro = 14474460
const b2_colorGhostWhite = 16316671
const b2_colorGold = 16766720
const b2_colorGoldenRod = 14329120
const b2_colorGray = 8421504
const b2_colorGreen = 32768
const b2_colorGreenYellow = 11403055
const b2_colorHoneyDew = 15794160
const b2_colorHotPink = 16738740
const b2_colorIndianRed = 13458524
const b2_colorIndigo = 4915330
const b2_colorIvory = 16777200
const b2_colorKhaki = 15787660
const b2_colorLavender = 15132410
const b2_colorLavenderBlush = 16773365
const b2_colorLawnGreen = 8190976
const b2_colorLemonChiffon = 16775885
const b2_colorLightBlue = 11393254
const b2_colorLightCoral = 15761536
const b2_colorLightCyan = 14745599
const b2_colorLightGoldenRodYellow = 16448210
const b2_colorLightGray = 13882323
const b2_colorLightGreen = 9498256
const b2_colorLightPink = 16758465
const b2_colorLightSalmon = 16752762
const b2_colorLightSeaGreen = 2142890
const b2_colorLightSkyBlue = 8900346
const b2_colorLightSlateGray = 7833753
const b2_colorLightSteelBlue = 11584734
const b2_colorLightYellow = 16777184
const b2_colorLime = 65280
const b2_colorLimeGreen = 3329330
const b2_colorLinen = 16445670
const b2_colorMagenta = 16711935
const b2_colorMaroon = 8388608
const b2_colorMediumAquaMarine = 6737322
const b2_colorMediumBlue = 205
const b2_colorMediumOrchid = 12211667
const b2_colorMediumPurple = 9662683
const b2_colorMediumSeaGreen = 3978097
const b2_colorMediumSlateBlue = 8087790
const b2_colorMediumSpringGreen = 64154
const b2_colorMediumTurquoise = 4772300
const b2_colorMediumVioletRed = 13047173
const b2_colorMidnightBlue = 1644912
const b2_colorMintCream = 16121850
const b2_colorMistyRose = 16770273
const b2_colorMoccasin = 16770229
const b2_colorNavajoWhite = 16768685
const b2_colorNavy = 128
const b2_colorOldLace = 16643558
const b2_colorOlive = 8421376
const b2_colorOliveDrab = 7048739
const b2_colorOrange = 16753920
const b2_colorOrangeRed = 16729344
const b2_colorOrchid = 14315734
const b2_colorPaleGoldenRod = 15657130
const b2_colorPaleGreen = 10025880
const b2_colorPaleTurquoise = 11529966
const b2_colorPaleVioletRed = 14381203
const b2_colorPapayaWhip = 16773077
const b2_colorPeachPuff = 16767673
const b2_colorPeru = 13468991
const b2_colorPink = 16761035
const b2_colorPlum = 14524637
const b2_colorPowderBlue = 11591910
const b2_colorPurple = 8388736
const b2_colorRebeccaPurple = 6697881
const b2_colorRed = 16711680
const b2_colorRosyBrown = 12357519
const b2_colorRoyalBlue = 4286945
const b2_colorSaddleBrown = 9127187
const b2_colorSalmon = 16416882
const b2_colorSandyBrown = 16032864
const b2_colorSeaGreen = 3050327
const b2_colorSeaShell = 16774638
const b2_colorSienna = 10506797
const b2_colorSilver = 12632256
const b2_colorSkyBlue = 8900331
const b2_colorSlateBlue = 6970061
const b2_colorSlateGray = 7372944
const b2_colorSnow = 16775930
const b2_colorSpringGreen = 65407
const b2_colorSteelBlue = 4620980
const b2_colorTan = 13808780
const b2_colorTeal = 32896
const b2_colorThistle = 14204888
const b2_colorTomato = 16737095
const b2_colorTurquoise = 4251856
const b2_colorViolet = 15631086
const b2_colorWheat = 16113331
const b2_colorWhite = 16777215
const b2_colorWhiteSmoke = 16119285
const b2_colorYellow = 16776960
const b2_colorYellowGreen = 10145074
const b2_colorBox2DRed = 14430514
const b2_colorBox2DBlue = 3190463
const b2_colorBox2DGreen = 9226532
const b2_colorBox2DYellow = 16772748

type b2DebugDraw struct {
	DrawPolygonFcn       uintptr
	DrawSolidPolygonFcn  uintptr
	DrawCircleFcn        uintptr
	DrawSolidCircleFcn   uintptr
	DrawSolidCapsuleFcn  uintptr
	DrawSegmentFcn       uintptr
	DrawTransformFcn     uintptr
	DrawPointFcn         uintptr
	DrawStringFcn        uintptr
	DrawingBounds        AABB
	UseDrawingBounds     uint8
	DrawShapes           uint8
	DrawJoints           uint8
	DrawJointExtras      uint8
	DrawBounds           uint8
	DrawMass             uint8
	DrawBodyNames        uint8
	DrawContacts         uint8
	DrawGraphColors      uint8
	DrawContactNormals   uint8
	DrawContactImpulses  uint8
	DrawContactFeatures  uint8
	DrawFrictionImpulses uint8
	DrawIslands          uint8
	Context              uintptr
}

/**@}*/

/**
 * @defgroup math_cpp C++ Math
 * @brief Math operator overloads for C++
 *
 * See math_functions.h for details.
 * @{
 */

/**@}*/
func b2IsValidAABB(tls *_Stack, a1 AABB) (r uint8) {
	var d, v1, v2, v3 Vec2
	var valid uint8
	_, _, _, _, _ = d, valid, v1, v2, v3
	v1 = a1.UpperBound
	v2 = a1.LowerBound
	v3 = Vec2{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
	goto _4
_4:
	d = v3
	valid = boolUint8(d.X >= float32FromFloat32(0) && d.Y >= float32FromFloat32(0))
	valid = boolUint8(valid != 0 && b2IsValidVec2(tls, a1.LowerBound) != 0 && b2IsValidVec2(tls, a1.UpperBound) != 0)
	return valid
}

// C documentation
//
//	// From Real-time Collision Detection, p179.
func b2AABB_RayCast(tls *_Stack, a5 AABB, p1 Vec2, p2 Vec2) (r CastOutput) {
	var absD, b3, d, normal, p, v1, v14, v2, v3, v5 Vec2
	var inv_d, inv_d1, s, s1, t1, t11, t2, t21, tmax, tmin, tmp, tmp1, v10, v11, v13, v6, v7, v9 float32
	var output CastOutput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = absD, b3, d, inv_d, inv_d1, normal, output, p, s, s1, t1, t11, t2, t21, tmax, tmin, tmp, tmp1, v1, v10, v11, v13, v14, v2, v3, v5, v6, v7, v9
	// Radius not handled
	output = CastOutput{}
	tmin = -float32FromFloat32(3.4028234663852886e+38)
	tmax = float32FromFloat32(3.4028234663852886e+38)
	p = p1
	v1 = p2
	v2 = p1
	v3 = Vec2{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
	goto _4
_4:
	d = v3
	v5 = d
	v6 = v5.X
	if v6 < float32FromInt32(0) {
		v9 = -v6
	} else {
		v9 = v6
	}
	v7 = v9
	goto _8
_8:
	b3.X = v7
	v10 = v5.Y
	if v10 < float32FromInt32(0) {
		v13 = -v10
	} else {
		v13 = v10
	}
	v11 = v13
	goto _12
_12:
	b3.Y = v11
	v14 = b3
	goto _15
_15:
	absD = v14
	normal = b2Vec2_zero
	// x-coordinate
	if absD.X < float32FromFloat32(1.1920928955078125e-07) {
		// parallel
		if p.X < a5.LowerBound.X || a5.UpperBound.X < p.X {
			return output
		}
	} else {
		inv_d = float32FromFloat32(1) / d.X
		t1 = float32((a5.LowerBound.X - p.X) * inv_d)
		t2 = float32((a5.UpperBound.X - p.X) * inv_d)
		// Sign of the normal vector.
		s = -float32FromFloat32(1)
		if t1 > t2 {
			tmp = t1
			t1 = t2
			t2 = tmp
			s = float32FromFloat32(1)
		}
		// Push the min up
		if t1 > tmin {
			normal.Y = float32FromFloat32(0)
			normal.X = s
			tmin = t1
		}
		// Pull the max down
		v6 = tmax
		v7 = t2
		if v6 < v7 {
			v10 = v6
		} else {
			v10 = v7
		}
		v9 = v10
		goto _19
	_19:
		tmax = v9
		if tmin > tmax {
			return output
		}
	}
	// y-coordinate
	if absD.Y < float32FromFloat32(1.1920928955078125e-07) {
		// parallel
		if p.Y < a5.LowerBound.Y || a5.UpperBound.Y < p.Y {
			return output
		}
	} else {
		inv_d1 = float32FromFloat32(1) / d.Y
		t11 = float32((a5.LowerBound.Y - p.Y) * inv_d1)
		t21 = float32((a5.UpperBound.Y - p.Y) * inv_d1)
		// Sign of the normal vector.
		s1 = -float32FromFloat32(1)
		if t11 > t21 {
			tmp1 = t11
			t11 = t21
			t21 = tmp1
			s1 = float32FromFloat32(1)
		}
		// Push the min up
		if t11 > tmin {
			normal.X = float32FromFloat32(0)
			normal.Y = s1
			tmin = t11
		}
		// Pull the max down
		v6 = tmax
		v7 = t21
		if v6 < v7 {
			v10 = v6
		} else {
			v10 = v7
		}
		v9 = v10
		goto _24
	_24:
		tmax = v9
		if tmin > tmax {
			return output
		}
	}
	// Does the ray start inside the box?
	// Does the ray intersect beyond the max fraction?
	if tmin < float32FromFloat32(0) || float32FromFloat32(1) < tmin {
		return output
	}
	// Intersection.
	output.Fraction = tmin
	output.Normal = normal
	v1 = p1
	v2 = p2
	v6 = tmin
	v3 = Vec2{
		X: float32((float32FromFloat32(1)-v6)*v1.X) + float32(v6*v2.X),
		Y: float32((float32FromFloat32(1)-v6)*v1.Y) + float32(v6*v2.Y),
	}
	goto _30
_30:
	output.Point = v3
	output.Hit = boolUint8(true1 != 0)
	return output
}

const _B2_DEBUG = 1
const _B2_RESTRICT = "restrict"
const _B2_SECRET_COOKIE = 1152023
const _B2_SIMD_WIDTH = 4
const _B2_SNOOP_PAIR_COUNTERS = 0
const _B2_SNOOP_TABLE_COUNTERS = 0
const _B2_SNOOP_TOI_COUNTERS = 0
const _B2_VALIDATE = 0
const _FLT_MAX1 = 3.40282346638528859812e+38

type b2AtomicInt struct {
	Value int32
}

type b2AtomicU32 struct {
	Value uint32
}

type b2IntArray struct {
	Data     uintptr
	Count    int32
	Capacity int32
}

type b2BodyArray struct {
	Data     uintptr
	Count    int32
	Capacity int32
}

type b2BodyMoveEventArray struct {
	Data     uintptr
	Count    int32
	Capacity int32
}

type b2BodySimArray struct {
	Data     uintptr
	Count    int32
	Capacity int32
}

type b2BodyStateArray struct {
	Data     uintptr
	Count    int32
	Capacity int32
}

type b2ChainShapeArray struct {
	Data     uintptr
	Count    int32
	Capacity int32
}

type b2ContactArray struct {
	Data     uintptr
	Count    int32
	Capacity int32
}

type b2ContactBeginTouchEventArray struct {
	Data     uintptr
	Count    int32
	Capacity int32
}

type b2ContactEndTouchEventArray struct {
	Data     uintptr
	Count    int32
	Capacity int32
}

type b2ContactHitEventArray struct {
	Data     uintptr
	Count    int32
	Capacity int32
}

type b2ContactSimArray struct {
	Data     uintptr
	Count    int32
	Capacity int32
}

type b2IslandArray struct {
	Data     uintptr
	Count    int32
	Capacity int32
}

type b2IslandSimArray struct {
	Data     uintptr
	Count    int32
	Capacity int32
}

type b2JointArray struct {
	Data     uintptr
	Count    int32
	Capacity int32
}

type b2JointSimArray struct {
	Data     uintptr
	Count    int32
	Capacity int32
}

type b2SensorArray struct {
	Data     uintptr
	Count    int32
	Capacity int32
}

type b2SensorBeginTouchEventArray struct {
	Data     uintptr
	Count    int32
	Capacity int32
}

type b2SensorEndTouchEventArray struct {
	Data     uintptr
	Count    int32
	Capacity int32
}

type b2SensorTaskContextArray struct {
	Data     uintptr
	Count    int32
	Capacity int32
}

type b2ShapeArray struct {
	Data     uintptr
	Count    int32
	Capacity int32
}

type b2ShapeRefArray struct {
	Data     uintptr
	Count    int32
	Capacity int32
}

type b2SolverSetArray struct {
	Data     uintptr
	Count    int32
	Capacity int32
}

type b2TaskContextArray struct {
	Data     uintptr
	Count    int32
	Capacity int32
}

type b2ArenaEntryArray struct {
	Data     uintptr
	Count    int32
	Capacity int32
}

type b2ArenaEntry struct {
	Data       uintptr
	Name       uintptr
	Size       int32
	UsedMalloc uint8
}

type b2ArenaAllocator struct {
	Data          uintptr
	Capacity      int32
	Index         int32
	Allocation    int32
	MaxAllocation int32
	Entries       b2ArenaEntryArray
}

func b2ArenaEntryArray_Create(tls *_Stack, capacity int32) (r b2ArenaEntryArray) {
	var a b2ArenaEntryArray
	_ = a
	a = b2ArenaEntryArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(24)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2ArenaEntryArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2ArenaEntryArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2ArenaEntryArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2ArenaEntryArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2ArenaEntryArray)(unsafe.Pointer(a)).Capacity)*uint64(24)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(24)))
	(*b2ArenaEntryArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2ArenaEntryArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2ArenaEntryArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2ArenaEntryArray)(unsafe.Pointer(a)).Capacity)*uint64(24)))
	(*b2ArenaEntryArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2ArenaEntryArray)(unsafe.Pointer(a)).Count = 0
	(*b2ArenaEntryArray)(unsafe.Pointer(a)).Capacity = 0
}

func b2CreateArenaAllocator(tls *_Stack, capacity int32) (r b2ArenaAllocator) {
	var allocator b2ArenaAllocator
	_ = allocator
	if !(capacity >= int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts, __ccgo_ts+14, int32FromInt32(16)) != 0 {
		__builtin_trap(tls)
	}
	allocator = b2ArenaAllocator{}
	allocator.Capacity = capacity
	allocator.Data = b2Alloc(tls, capacity)
	allocator.Allocation = 0
	allocator.MaxAllocation = 0
	allocator.Index = 0
	allocator.Entries = b2ArenaEntryArray_Create(tls, int32(32))
	return allocator
}

func b2DestroyArenaAllocator(tls *_Stack, allocator uintptr) {
	b2ArenaEntryArray_Destroy(tls, allocator+24)
	b2Free(tls, (*b2ArenaAllocator)(unsafe.Pointer(allocator)).Data, (*b2ArenaAllocator)(unsafe.Pointer(allocator)).Capacity)
}

func b2AllocateArenaItem(tls *_Stack, alloc uintptr, size int32, name uintptr) (r uintptr) {
	var entry b2ArenaEntry
	var newCapacity, size32, v2 int32
	var v1 uintptr
	_, _, _, _, _ = entry, newCapacity, size32, v1, v2
	// ensure allocation is 32 byte aligned to support 256-bit SIMD
	size32 = size - int32(1) | int32(0x1F) + int32(1)
	entry.Size = size32
	entry.Name = name
	if (*b2ArenaAllocator)(unsafe.Pointer(alloc)).Index+size32 > (*b2ArenaAllocator)(unsafe.Pointer(alloc)).Capacity {
		// fall back to the heap (undesirable)
		entry.Data = b2Alloc(tls, size32)
		entry.UsedMalloc = boolUint8(true1 != 0)
		if !(uint64(entry.Data)&uint64FromInt32(0x1F) == uint64FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+47, __ccgo_ts+14, int32FromInt32(47)) != 0 {
			__builtin_trap(tls)
		}
	} else {
		entry.Data = (*b2ArenaAllocator)(unsafe.Pointer(alloc)).Data + uintptr((*b2ArenaAllocator)(unsafe.Pointer(alloc)).Index)
		entry.UsedMalloc = boolUint8(false1 != 0)
		*(*int32)(unsafe.Add(nil, alloc+12)) += size32
		if !(uint64(entry.Data)&uint64FromInt32(0x1F) == uint64FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+47, __ccgo_ts+14, int32FromInt32(55)) != 0 {
			__builtin_trap(tls)
		}
	}
	*(*int32)(unsafe.Add(nil, alloc+16)) += size32
	if (*b2ArenaAllocator)(unsafe.Pointer(alloc)).Allocation > (*b2ArenaAllocator)(unsafe.Pointer(alloc)).MaxAllocation {
		(*b2ArenaAllocator)(unsafe.Pointer(alloc)).MaxAllocation = (*b2ArenaAllocator)(unsafe.Pointer(alloc)).Allocation
	}
	v1 = alloc + 24
	if (*b2ArenaEntryArray)(unsafe.Pointer(v1)).Count == (*b2ArenaEntryArray)(unsafe.Pointer(v1)).Capacity {
		if (*b2ArenaEntryArray)(unsafe.Pointer(v1)).Capacity < int32(2) {
			v2 = int32(2)
		} else {
			v2 = (*b2ArenaEntryArray)(unsafe.Pointer(v1)).Capacity + (*b2ArenaEntryArray)(unsafe.Pointer(v1)).Capacity>>int32(1)
		}
		newCapacity = v2
		b2ArenaEntryArray_Reserve(tls, v1, newCapacity)
	}
	*(*b2ArenaEntry)(unsafe.Add(nil, (*b2ArenaEntryArray)(unsafe.Pointer(v1)).Data+uintptr((*b2ArenaEntryArray)(unsafe.Pointer(v1)).Count)*24)) = entry
	*(*int32)(unsafe.Add(nil, v1+8)) += int32(1)
	return entry.Data
}

func b2FreeArenaItem(tls *_Stack, alloc uintptr, mem uintptr) {
	var entry, v1 uintptr
	var entryCount int32
	var value b2ArenaEntry
	_, _, _, _ = entry, entryCount, value, v1
	entryCount = (*b2ArenaAllocator)(unsafe.Pointer(alloc)).Entries.Count
	if !(entryCount > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+85, __ccgo_ts+14, int32FromInt32(71)) != 0 {
		__builtin_trap(tls)
	}
	entry = (*b2ArenaAllocator)(unsafe.Pointer(alloc)).Entries.Data + uintptr(entryCount-int32FromInt32(1))*24
	if !(mem == (*b2ArenaEntry)(unsafe.Pointer(entry)).Data) && b2InternalAssertFcn(tls, __ccgo_ts+100, __ccgo_ts+14, int32FromInt32(73)) != 0 {
		__builtin_trap(tls)
	}
	if (*b2ArenaEntry)(unsafe.Pointer(entry)).UsedMalloc != 0 {
		b2Free(tls, mem, (*b2ArenaEntry)(unsafe.Pointer(entry)).Size)
	} else {
		*(*int32)(unsafe.Add(nil, alloc+12)) -= (*b2ArenaEntry)(unsafe.Pointer(entry)).Size
	}
	*(*int32)(unsafe.Add(nil, alloc+16)) -= (*b2ArenaEntry)(unsafe.Pointer(entry)).Size
	v1 = alloc + 24
	if !((*b2ArenaEntryArray)(unsafe.Pointer(v1)).Count > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+119, __ccgo_ts+132, int32FromInt32(48)) != 0 {
		__builtin_trap(tls)
	}
	value = *(*b2ArenaEntry)(unsafe.Add(nil, (*b2ArenaEntryArray)(unsafe.Pointer(v1)).Data+uintptr((*b2ArenaEntryArray)(unsafe.Pointer(v1)).Count-int32(1))*24))
	*(*int32)(unsafe.Add(nil, v1+8)) -= int32(1)
	_ = value
	goto _2
_2:
}

func b2GrowArena(tls *_Stack, alloc uintptr) {
	// Stack must not be in use
	if !((*b2ArenaAllocator)(unsafe.Pointer(alloc)).Allocation == int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+165, __ccgo_ts+14, int32FromInt32(89)) != 0 {
		__builtin_trap(tls)
	}
	if (*b2ArenaAllocator)(unsafe.Pointer(alloc)).MaxAllocation > (*b2ArenaAllocator)(unsafe.Pointer(alloc)).Capacity {
		b2Free(tls, (*b2ArenaAllocator)(unsafe.Pointer(alloc)).Data, (*b2ArenaAllocator)(unsafe.Pointer(alloc)).Capacity)
		(*b2ArenaAllocator)(unsafe.Pointer(alloc)).Capacity = (*b2ArenaAllocator)(unsafe.Pointer(alloc)).MaxAllocation + (*b2ArenaAllocator)(unsafe.Pointer(alloc)).MaxAllocation/int32(2)
		(*b2ArenaAllocator)(unsafe.Pointer(alloc)).Data = b2Alloc(tls, (*b2ArenaAllocator)(unsafe.Pointer(alloc)).Capacity)
	}
}

func b2GetArenaCapacity(tls *_Stack, alloc uintptr) (r int32) {
	return (*b2ArenaAllocator)(unsafe.Pointer(alloc)).Capacity
}

func b2GetArenaAllocation(tls *_Stack, alloc uintptr) (r int32) {
	return (*b2ArenaAllocator)(unsafe.Pointer(alloc)).Allocation
}

func b2GetMaxArenaAllocation(tls *_Stack, alloc uintptr) (r int32) {
	return (*b2ArenaAllocator)(unsafe.Pointer(alloc)).MaxAllocation
}

func b2IntArray_Create(tls *_Stack, capacity int32) (r b2IntArray) {
	var a b2IntArray
	_ = a
	a = b2IntArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(4)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2IntArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2IntArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2IntArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2IntArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2IntArray)(unsafe.Pointer(a)).Capacity)*uint64(4)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(4)))
	(*b2IntArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2IntArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2IntArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2IntArray)(unsafe.Pointer(a)).Capacity)*uint64(4)))
	(*b2IntArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2IntArray)(unsafe.Pointer(a)).Count = 0
	(*b2IntArray)(unsafe.Pointer(a)).Capacity = 0
}

type b2BitSet struct {
	Bits          uintptr
	BlockCapacity uint32
	BlockCount    uint32
}

type locale_t = uintptr

func b2CreateBitSet(tls *_Stack, bitCapacity uint32) (r b2BitSet) {
	var bitSet b2BitSet
	_ = bitSet
	bitSet = b2BitSet{}
	bitSet.BlockCapacity = uint32((uint64(bitCapacity) + uint64FromInt64(8)*uint64FromInt32(8) - uint64(1)) / (uint64FromInt64(8) * uint64FromInt32(8)))
	bitSet.BlockCount = uint32(0)
	bitSet.Bits = b2Alloc(tls, int32FromUint64(uint64(bitSet.BlockCapacity)*uint64(8)))
	memset(tls, bitSet.Bits, 0, uint64(bitSet.BlockCapacity)*uint64(8))
	return bitSet
}

func b2DestroyBitSet(tls *_Stack, bitSet uintptr) {
	b2Free(tls, (*b2BitSet)(unsafe.Pointer(bitSet)).Bits, int32FromUint64(uint64((*b2BitSet)(unsafe.Pointer(bitSet)).BlockCapacity)*uint64(8)))
	(*b2BitSet)(unsafe.Pointer(bitSet)).BlockCapacity = uint32(0)
	(*b2BitSet)(unsafe.Pointer(bitSet)).BlockCount = uint32(0)
	(*b2BitSet)(unsafe.Pointer(bitSet)).Bits = uintptrFromInt32(0)
}

func b2SetBitCountAndClear(tls *_Stack, bitSet uintptr, bitCount uint32) {
	var blockCount, newBitCapacity uint32_t
	_, _ = blockCount, newBitCapacity
	blockCount = uint32((uint64(bitCount) + uint64FromInt64(8)*uint64FromInt32(8) - uint64(1)) / (uint64FromInt64(8) * uint64FromInt32(8)))
	if (*b2BitSet)(unsafe.Pointer(bitSet)).BlockCapacity < blockCount {
		b2DestroyBitSet(tls, bitSet)
		newBitCapacity = bitCount + bitCount>>int32FromInt32(1)
		*(*b2BitSet)(unsafe.Add(nil, bitSet)) = b2CreateBitSet(tls, newBitCapacity)
	}
	(*b2BitSet)(unsafe.Pointer(bitSet)).BlockCount = blockCount
	memset(tls, (*b2BitSet)(unsafe.Pointer(bitSet)).Bits, 0, uint64((*b2BitSet)(unsafe.Pointer(bitSet)).BlockCount)*uint64(8))
}

func b2GrowBitSet(tls *_Stack, bitSet uintptr, blockCount uint32) {
	var newBits uintptr
	var oldCapacity uint32_t
	_, _ = newBits, oldCapacity
	if !(blockCount > (*b2BitSet)(unsafe.Pointer(bitSet)).BlockCount) && b2InternalAssertFcn(tls, __ccgo_ts+188, __ccgo_ts+220, int32FromInt32(43)) != 0 {
		__builtin_trap(tls)
	}
	if blockCount > (*b2BitSet)(unsafe.Pointer(bitSet)).BlockCapacity {
		oldCapacity = (*b2BitSet)(unsafe.Pointer(bitSet)).BlockCapacity
		(*b2BitSet)(unsafe.Pointer(bitSet)).BlockCapacity = blockCount + blockCount/uint32(2)
		newBits = b2Alloc(tls, int32FromUint64(uint64((*b2BitSet)(unsafe.Pointer(bitSet)).BlockCapacity)*uint64(8)))
		memset(tls, newBits, 0, uint64((*b2BitSet)(unsafe.Pointer(bitSet)).BlockCapacity)*uint64(8))
		if !((*b2BitSet)(unsafe.Pointer(bitSet)).Bits != uintptrFromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+244, __ccgo_ts+220, int32FromInt32(50)) != 0 {
			__builtin_trap(tls)
		}
		memcpy(tls, newBits, (*b2BitSet)(unsafe.Pointer(bitSet)).Bits, uint64(oldCapacity)*uint64(8))
		b2Free(tls, (*b2BitSet)(unsafe.Pointer(bitSet)).Bits, int32FromUint64(uint64(oldCapacity)*uint64(8)))
		(*b2BitSet)(unsafe.Pointer(bitSet)).Bits = newBits
	}
	(*b2BitSet)(unsafe.Pointer(bitSet)).BlockCount = blockCount
}

func b2InPlaceUnion(tls *_Stack, setA uintptr, setB uintptr) {
	var blockCount, i uint32_t
	_, _ = blockCount, i
	if !((*b2BitSet)(unsafe.Pointer(setA)).BlockCount == (*b2BitSet)(unsafe.Pointer(setB)).BlockCount) && b2InternalAssertFcn(tls, __ccgo_ts+265, __ccgo_ts+220, int32FromInt32(61)) != 0 {
		__builtin_trap(tls)
	}
	blockCount = (*b2BitSet)(unsafe.Pointer(setA)).BlockCount
	i = uint32(0)
	for {
		if !(i < blockCount) {
			break
		}
		*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(setA)).Bits+uintptr(i)*8)) |= *(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(setB)).Bits+uintptr(i)*8))
		goto _1
	_1:
		;
		i = i + 1
	}
}

const _B2_GRAPH_COLOR_COUNT = 12
const _B2_JOINT_CONSTRAINT_DAMPING_RATIO = "2.0f"
const _B2_JOINT_CONSTRAINT_HERTZ = "60.0f"
const _B2_MAX_WORKERS = 64
const _B2_MAX_WORLDS = 128
const _B2_TIME_TO_SLEEP = "0.5f"

type b2World struct {
	Arena                b2ArenaAllocator
	BroadPhase           b2BroadPhase
	ConstraintGraph      b2ConstraintGraph
	BodyIdPool           b2IdPool
	Bodies               b2BodyArray
	SolverSetIdPool      b2IdPool
	SolverSets           b2SolverSetArray
	JointIdPool          b2IdPool
	Joints               b2JointArray
	ContactIdPool        b2IdPool
	Contacts             b2ContactArray
	IslandIdPool         b2IdPool
	Islands              b2IslandArray
	ShapeIdPool          b2IdPool
	ChainIdPool          b2IdPool
	Shapes               b2ShapeArray
	ChainShapes          b2ChainShapeArray
	Sensors              b2SensorArray
	TaskContexts         b2TaskContextArray
	SensorTaskContexts   b2SensorTaskContextArray
	BodyMoveEvents       b2BodyMoveEventArray
	SensorBeginEvents    b2SensorBeginTouchEventArray
	ContactBeginEvents   b2ContactBeginTouchEventArray
	SensorEndEvents      [2]b2SensorEndTouchEventArray
	ContactEndEvents     [2]b2ContactEndTouchEventArray
	EndEventArrayIndex   int32
	ContactHitEvents     b2ContactHitEventArray
	DebugBodySet         b2BitSet
	DebugJointSet        b2BitSet
	DebugContactSet      b2BitSet
	DebugIslandSet       b2BitSet
	StepIndex            uint64
	SplitIslandId        int32
	Gravity              Vec2
	HitEventThreshold    float32
	RestitutionThreshold float32
	MaxLinearSpeed       float32
	MaxContactPushSpeed  float32
	ContactSpeed         float32
	ContactHertz         float32
	ContactDampingRatio  float32
	FrictionCallback     uintptr
	RestitutionCallback  uintptr
	Generation           uint16
	Profile              Profile
	PreSolveFcn          uintptr
	PreSolveContext      uintptr
	CustomFilterFcn      uintptr
	CustomFilterContext  uintptr
	WorkerCount          int32
	EnqueueTaskFcn       uintptr
	FinishTaskFcn        uintptr
	UserTaskContext      uintptr
	UserTreeTask         uintptr
	UserData             uintptr
	Inv_h                float32
	ActiveTaskCount      int32
	TaskCount            int32
	WorldId              uint16
	EnableSleep          uint8
	Locked               uint8
	EnableWarmStarting   uint8
	EnableContinuous     uint8
	EnableSpeculative    uint8
	InUse                uint8
}

type b2Body struct {
	Name           [32]uint8
	UserData       uintptr
	SetIndex       int32
	LocalIndex     int32
	HeadContactKey int32
	ContactCount   int32
	HeadShapeId    int32
	ShapeCount     int32
	HeadChainId    int32
	HeadJointKey   int32
	JointCount     int32
	IslandId       int32
	IslandPrev     int32
	IslandNext     int32
	Mass           float32
	Inertia        float32
	SleepThreshold float32
	SleepTime      float32
	BodyMoveIndex  int32
	Id             int32
	Type1          BodyType
	Generation     uint16
	EnableSleep    uint8
	FixedRotation  uint8
	IsSpeedCapped  uint8
	IsMarked       uint8
}

type b2BodyState struct {
	LinearVelocity  Vec2
	AngularVelocity float32
	Flags           int32
	DeltaPosition   Vec2
	DeltaRotation   Rot
}

var b2_identityBodyState = b2BodyState{
	DeltaRotation: Rot{
		C: float32FromFloat32(1),
	},
}

type b2BodySim struct {
	Transform         Transform
	Center            Vec2
	Rotation0         Rot
	Center0           Vec2
	LocalCenter       Vec2
	Force             Vec2
	Torque            float32
	InvMass           float32
	InvInertia        float32
	MinExtent         float32
	MaxExtent         float32
	LinearDamping     float32
	AngularDamping    float32
	GravityScale      float32
	BodyId            int32
	IsFast            uint8
	IsBullet          uint8
	IsSpeedCapped     uint8
	AllowFastRotation uint8
	EnlargeAABB       uint8
}

type b2Shape struct {
	Id                int32
	BodyId            int32
	PrevShapeId       int32
	NextShapeId       int32
	SensorIndex       int32
	Type1             ShapeType
	Density           float32
	Friction          float32
	Restitution       float32
	RollingResistance float32
	TangentSpeed      float32
	UserMaterialId    int32
	Aabb              AABB
	FatAABB           AABB
	LocalCentroid     Vec2
	ProxyKey          int32
	Filter            Filter
	UserData          uintptr
	CustomColor       uint32
	ｆ__ccgo19_132     struct {
		Circle       [0]Circle
		Polygon      [0]Polygon
		Segment      [0]Segment
		ChainSegment [0]ChainSegment
		Capsule      Capsule
		ｆ__ccgo_pad5 [124]byte
	}
	Generation           uint16
	EnableSensorEvents   uint8
	EnableContactEvents  uint8
	EnableHitEvents      uint8
	EnablePreSolveEvents uint8
	EnlargedAABB         uint8
}

type b2ContactFlags = int32

const b2_contactTouchingFlag = 1
const b2_contactHitEventFlag = 2
const b2_contactEnableContactEvents = 4

type b2ContactEdge struct {
	BodyId  int32
	PrevKey int32
	NextKey int32
}

type b2Contact struct {
	SetIndex   int32
	ColorIndex int32
	LocalIndex int32
	Edges      [2]b2ContactEdge
	ShapeIdA   int32
	ShapeIdB   int32
	IslandPrev int32
	IslandNext int32
	IslandId   int32
	ContactId  int32
	Flags      uint32
	IsMarked   uint8
}

type b2ContactSimFlags = int32

const b2_simTouchingFlag = 65536
const b2_simDisjoint = 131072
const b2_simStartedTouching = 262144
const b2_simStoppedTouching = 524288
const b2_simEnableHitEvent = 1048576
const b2_simEnablePreSolveEvents = 2097152

type b2ContactSim struct {
	ContactId         int32
	BodySimIndexA     int32
	BodySimIndexB     int32
	ShapeIdA          int32
	ShapeIdB          int32
	InvMassA          float32
	InvIA             float32
	InvMassB          float32
	InvIB             float32
	Manifold          Manifold
	Friction          float32
	Restitution       float32
	RollingResistance float32
	TangentSpeed      float32
	SimFlags          uint32
	Cache             SimplexCache
}

type b2IdPool struct {
	FreeArray b2IntArray
	NextIndex int32
}

type b2Joint struct {
	UserData         uintptr
	SetIndex         int32
	ColorIndex       int32
	LocalIndex       int32
	Edges            [2]b2JointEdge
	JointId          int32
	IslandId         int32
	IslandPrev       int32
	IslandNext       int32
	DrawSize         float32
	Type1            JointType
	Generation       uint16
	IsMarked         uint8
	CollideConnected uint8
}

type b2Island struct {
	SetIndex              int32
	LocalIndex            int32
	IslandId              int32
	HeadBody              int32
	TailBody              int32
	BodyCount             int32
	HeadContact           int32
	TailContact           int32
	ContactCount          int32
	HeadJoint             int32
	TailJoint             int32
	JointCount            int32
	ParentIsland          int32
	ConstraintRemoveCount int32
}

type b2IslandSim struct {
	IslandId int32
}

type b2JointSim struct {
	JointId                int32
	BodyIdA                int32
	BodyIdB                int32
	Type1                  JointType
	LocalOriginAnchorA     Vec2
	LocalOriginAnchorB     Vec2
	InvMassA               float32
	InvMassB               float32
	InvIA                  float32
	InvIB                  float32
	ConstraintHertz        float32
	ConstraintDampingRatio float32
	ConstraintSoftness     b2Softness
	ｆ__ccgo13_68           struct {
		MotorJoint     [0]b2MotorJoint
		MouseJoint     [0]b2MouseJoint
		RevoluteJoint  [0]b2RevoluteJoint
		PrismaticJoint [0]b2PrismaticJoint
		WeldJoint      [0]b2WeldJoint
		WheelJoint     [0]b2WheelJoint
		DistanceJoint  b2DistanceJoint
		ｆ__ccgo_pad7   [32]byte
	}
}

type b2Softness struct {
	BiasRate     float32
	MassScale    float32
	ImpulseScale float32
}

type b2SolverStageType1 = int32

type b2SolverStageType = int32

const b2_stagePrepareJoints = 0
const b2_stagePrepareContacts = 1
const b2_stageIntegrateVelocities = 2
const b2_stageWarmStart = 3
const b2_stageSolve = 4
const b2_stageIntegratePositions = 5
const b2_stageRelax = 6
const b2_stageRestitution = 7
const b2_stageStoreImpulses = 8

type b2SolverBlockType1 = int32

type b2SolverBlockType = int32

const b2_bodyBlock = 0
const b2_jointBlock = 1
const b2_contactBlock = 2
const b2_graphJointBlock = 3
const b2_graphContactBlock = 4

type b2SolverBlock struct {
	StartIndex int32
	Count      int16
	BlockType  int16
	SyncIndex  b2AtomicInt
}

type b2SolverStage struct {
	Type1           b2SolverStageType1
	Blocks          uintptr
	BlockCount      int32
	ColorIndex      int32
	CompletionCount b2AtomicInt
}

type b2StepContext struct {
	Dt                     float32
	Inv_dt                 float32
	H                      float32
	Inv_h                  float32
	SubStepCount           int32
	ContactSoftness        b2Softness
	StaticSoftness         b2Softness
	RestitutionThreshold   float32
	MaxLinearVelocity      float32
	World                  uintptr
	Graph                  uintptr
	States                 uintptr
	Sims                   uintptr
	EnlargedShapes         uintptr
	EnlargedShapeCount     int32
	BulletBodies           uintptr
	BulletBodyCount        b2AtomicInt
	Joints                 uintptr
	Contacts               uintptr
	SimdContactConstraints uintptr
	ActiveColorCount       int32
	WorkerCount            int32
	Stages                 uintptr
	StageCount             int32
	EnableWarmStarting     uint8
	Dummy1                 [64]uint8
	AtomicSyncBits         b2AtomicU32
	Dummy2                 [64]uint8
}

type b2JointEdge struct {
	BodyId  int32
	PrevKey int32
	NextKey int32
}

type b2DistanceJoint struct {
	Length           float32
	Hertz            float32
	DampingRatio     float32
	MinLength        float32
	MaxLength        float32
	MaxMotorForce    float32
	MotorSpeed       float32
	Impulse          float32
	LowerImpulse     float32
	UpperImpulse     float32
	MotorImpulse     float32
	IndexA           int32
	IndexB           int32
	AnchorA          Vec2
	AnchorB          Vec2
	DeltaCenter      Vec2
	DistanceSoftness b2Softness
	AxialMass        float32
	EnableSpring     uint8
	EnableLimit      uint8
	EnableMotor      uint8
}

type b2MotorJoint struct {
	LinearOffset     Vec2
	AngularOffset    float32
	LinearImpulse    Vec2
	AngularImpulse   float32
	MaxForce         float32
	MaxTorque        float32
	CorrectionFactor float32
	IndexA           int32
	IndexB           int32
	AnchorA          Vec2
	AnchorB          Vec2
	DeltaCenter      Vec2
	DeltaAngle       float32
	LinearMass       Mat22
	AngularMass      float32
}

type b2MouseJoint struct {
	TargetA         Vec2
	Hertz           float32
	DampingRatio    float32
	MaxForce        float32
	LinearImpulse   Vec2
	AngularImpulse  float32
	LinearSoftness  b2Softness
	AngularSoftness b2Softness
	IndexB          int32
	AnchorB         Vec2
	DeltaCenter     Vec2
	LinearMass      Mat22
}

type b2PrismaticJoint struct {
	LocalAxisA        Vec2
	Impulse           Vec2
	SpringImpulse     float32
	MotorImpulse      float32
	LowerImpulse      float32
	UpperImpulse      float32
	Hertz             float32
	DampingRatio      float32
	TargetTranslation float32
	MaxMotorForce     float32
	MotorSpeed        float32
	ReferenceAngle    float32
	LowerTranslation  float32
	UpperTranslation  float32
	IndexA            int32
	IndexB            int32
	AnchorA           Vec2
	AnchorB           Vec2
	AxisA             Vec2
	DeltaCenter       Vec2
	DeltaAngle        float32
	AxialMass         float32
	SpringSoftness    b2Softness
	EnableSpring      uint8
	EnableLimit       uint8
	EnableMotor       uint8
}

type b2RevoluteJoint struct {
	LinearImpulse  Vec2
	SpringImpulse  float32
	MotorImpulse   float32
	LowerImpulse   float32
	UpperImpulse   float32
	Hertz          float32
	DampingRatio   float32
	TargetAngle    float32
	MaxMotorTorque float32
	MotorSpeed     float32
	ReferenceAngle float32
	LowerAngle     float32
	UpperAngle     float32
	IndexA         int32
	IndexB         int32
	AnchorA        Vec2
	AnchorB        Vec2
	DeltaCenter    Vec2
	DeltaAngle     float32
	AxialMass      float32
	SpringSoftness b2Softness
	EnableSpring   uint8
	EnableMotor    uint8
	EnableLimit    uint8
}

type b2WeldJoint struct {
	ReferenceAngle      float32
	LinearHertz         float32
	LinearDampingRatio  float32
	AngularHertz        float32
	AngularDampingRatio float32
	LinearSoftness      b2Softness
	AngularSoftness     b2Softness
	LinearImpulse       Vec2
	AngularImpulse      float32
	IndexA              int32
	IndexB              int32
	AnchorA             Vec2
	AnchorB             Vec2
	DeltaCenter         Vec2
	DeltaAngle          float32
	AxialMass           float32
}

type b2WheelJoint struct {
	LocalAxisA       Vec2
	PerpImpulse      float32
	MotorImpulse     float32
	SpringImpulse    float32
	LowerImpulse     float32
	UpperImpulse     float32
	MaxMotorTorque   float32
	MotorSpeed       float32
	LowerTranslation float32
	UpperTranslation float32
	Hertz            float32
	DampingRatio     float32
	IndexA           int32
	IndexB           int32
	AnchorA          Vec2
	AnchorB          Vec2
	AxisA            Vec2
	DeltaCenter      Vec2
	PerpMass         float32
	MotorMass        float32
	AxialMass        float32
	SpringSoftness   b2Softness
	EnableSpring     uint8
	EnableMotor      uint8
	EnableLimit      uint8
}

type b2ShapeRef struct {
	ShapeId    int32
	Generation uint16
}

type b2Sensor struct {
	Overlaps1 b2ShapeRefArray
	Overlaps2 b2ShapeRefArray
	ShapeId   int32
}

type b2SensorTaskContext struct {
	EventBits b2BitSet
}

type b2BroadPhase struct {
	Trees            [3]DynamicTree
	MoveSet          b2HashSet
	MoveArray        b2IntArray
	MoveResults      uintptr
	MovePairs        uintptr
	MovePairCapacity int32
	MovePairIndex    b2AtomicInt
	PairSet          b2HashSet
}

type b2ChainShape struct {
	Id            int32
	BodyId        int32
	NextChainId   int32
	Count         int32
	MaterialCount int32
	ShapeIndices  uintptr
	Materials     uintptr
	Generation    uint16
}

type b2ShapeExtent struct {
	MinExtent float32
	MaxExtent float32
}

type b2SensorOverlaps struct {
	Overlaps b2IntArray
}

type b2SolverSet struct {
	BodySims    b2BodySimArray
	BodyStates  b2BodyStateArray
	JointSims   b2JointSimArray
	ContactSims b2ContactSimArray
	IslandSims  b2IslandSimArray
	SetIndex    int32
}

type b2SetItem struct {
	Key  uint64
	Hash uint32
}

type b2HashSet struct {
	Items    uintptr
	Capacity uint32
	Count    uint32
}

type b2TreeNodeFlags = int32

const b2_allocatedNode = 1
const b2_enlargedNode = 2
const b2_leafNode = 4

type b2GraphColor struct {
	BodySet     b2BitSet
	ContactSims b2ContactSimArray
	JointSims   b2JointSimArray
	ｆ__ccgo3_48 struct {
		OverflowConstraints [0]uintptr
		SimdConstraints     uintptr
	}
}

type b2ConstraintGraph struct {
	Colors [12]b2GraphColor
}

type b2SetType = int32

const b2_staticSet = 0
const b2_disabledSet = 1
const b2_awakeSet = 2
const b2_firstSleepingSet = 3

type b2TaskContext struct {
	ContactStateBitSet b2BitSet
	EnlargedSimBitSet  b2BitSet
	AwakeIslandBitSet  b2BitSet
	SplitSleepTime     float32
	SplitIslandId      int32
}

func b2LimitVelocity(tls *_Stack, state uintptr, maxLinearSpeed float32) {
	var v2, v21, v4 float32
	var v1, v5, v6 Vec2
	_, _, _, _, _, _ = v2, v1, v21, v4, v5, v6
	v1 = (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
	v21 = float32(v1.X*v1.X) + float32(v1.Y*v1.Y)
	goto _3
_3:
	v2 = v21
	if v2 > float32(maxLinearSpeed*maxLinearSpeed) {
		v4 = maxLinearSpeed / sqrtf(tls, v2)
		v5 = (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
		v6 = Vec2{
			X: float32(v4 * v5.X),
			Y: float32(v4 * v5.Y),
		}
		goto _7
	_7:
		(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity = v6
	}
}

func b2ShouldBodiesCollide(tls *_Stack, world uintptr, bodyA uintptr, bodyB uintptr) (r uint8) {
	var edgeIndex, jointId, jointKey, otherBodyId, otherEdgeIndex, v2 int32
	var joint, v1, v3 uintptr
	_, _, _, _, _, _, _, _, _ = edgeIndex, joint, jointId, jointKey, otherBodyId, otherEdgeIndex, v1, v2, v3
	if (*b2Body)(unsafe.Pointer(bodyA)).Type1 != int32(b2_dynamicBody) && (*b2Body)(unsafe.Pointer(bodyB)).Type1 != int32(b2_dynamicBody) {
		return boolUint8(false1 != 0)
	}
	if (*b2Body)(unsafe.Pointer(bodyA)).JointCount < (*b2Body)(unsafe.Pointer(bodyB)).JointCount {
		jointKey = (*b2Body)(unsafe.Pointer(bodyA)).HeadJointKey
		otherBodyId = (*b2Body)(unsafe.Pointer(bodyB)).Id
	} else {
		jointKey = (*b2Body)(unsafe.Pointer(bodyB)).HeadJointKey
		otherBodyId = (*b2Body)(unsafe.Pointer(bodyA)).Id
	}
	for jointKey != -int32(1) {
		jointId = jointKey >> int32(1)
		edgeIndex = jointKey & int32(1)
		otherEdgeIndex = edgeIndex ^ int32(1)
		v1 = world + 1104
		v2 = jointId
		if !(0 <= v2 && v2 < (*b2JointArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2JointArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*72
		goto _4
	_4:
		joint = v3
		if int32FromUint8((*b2Joint)(unsafe.Pointer(joint)).CollideConnected) == false1 && (*(*b2JointEdge)(unsafe.Add(nil, joint+20+uintptr(otherEdgeIndex)*12))).BodyId == otherBodyId {
			return boolUint8(false1 != 0)
		}
		jointKey = (*(*b2JointEdge)(unsafe.Add(nil, joint+20+uintptr(edgeIndex)*12))).NextKey
	}
	return boolUint8(true1 != 0)
}

const _UINT64_MAX1 = 18446744073709551615

type b2MovePair struct {
	ShapeIndexA int32
	ShapeIndexB int32
	Next        uintptr
	Heap        uint8
}

type b2MoveResult struct {
	PairList uintptr
}

var b2_identityBodyState1 = b2BodyState{
	DeltaRotation: Rot{
		C: float32FromFloat32(1),
	},
}

func b2UnBufferMove(tls *_Stack, bp uintptr, proxyKey int32) {
	var count, i, movedIndex, v3 int32
	var found uint8
	var v2 uintptr
	_, _, _, _, _, _ = count, found, i, movedIndex, v2, v3
	found = b2RemoveKey(tls, bp+216, uint64FromInt32(proxyKey+int32(1)))
	if found != 0 {
		// Purge from move buffer. Linear search.
		// todo if I can iterate the move set then I don't need the moveArray
		count = (*b2BroadPhase)(unsafe.Pointer(bp)).MoveArray.Count
		i = 0
		for {
			if !(i < count) {
				break
			}
			if *(*int32)(unsafe.Add(nil, (*b2BroadPhase)(unsafe.Pointer(bp)).MoveArray.Data+uintptr(i)*4)) == proxyKey {
				v2 = bp + 232
				v3 = i
				if !(0 <= v3 && v3 < (*b2IntArray)(unsafe.Pointer(v2)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+2400, int32FromInt32(155)) != 0 {
					__builtin_trap(tls)
				}
				movedIndex = -int32(1)
				if v3 != (*b2IntArray)(unsafe.Pointer(v2)).Count-int32FromInt32(1) {
					movedIndex = (*b2IntArray)(unsafe.Pointer(v2)).Count - int32(1)
					*(*int32)(unsafe.Add(nil, (*b2IntArray)(unsafe.Pointer(v2)).Data+uintptr(v3)*4)) = *(*int32)(unsafe.Add(nil, (*b2IntArray)(unsafe.Pointer(v2)).Data+uintptr(movedIndex)*4))
				}
				*(*int32)(unsafe.Add(nil, v2+8)) -= int32(1)
				_ = movedIndex
				goto _4
			_4:
				;
				break
			}
			goto _1
		_1:
			;
			i = i + 1
		}
	}
}

type b2QueryPairContext struct {
	World           uintptr
	MoveResult      uintptr
	QueryTreeType   BodyType
	QueryProxyKey   int32
	QueryShapeIndex int32
}

// C documentation
//
//	// This is called from b2DynamicTree::Query when we are gathering pairs.
func b2PairQueryCallback(tls *_Stack, proxyId int32, userData uint64, context uintptr) (r uint8) {
	var bodyA, bodyB, broadPhase, customFilterFcn, pair, queryContext, shapeA, shapeB, world, v2, v4, v6, v8 uintptr
	var bodyIdA, bodyIdB, pairIndex, proxyKey, queryProxyKey, shapeId, shapeIdA, shapeIdB, v3, v7 int32
	var idA, idB ShapeId
	var moved, moved1, shouldCollide, v12 uint8
	var pairKey uint64_t
	var queryProxyType, treeType b2BodyType1
	var v1 uint64
	var v10, v11 Filter
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, bodyIdA, bodyIdB, broadPhase, customFilterFcn, idA, idB, moved, moved1, pair, pairIndex, pairKey, proxyKey, queryContext, queryProxyKey, queryProxyType, shapeA, shapeB, shapeId, shapeIdA, shapeIdB, shouldCollide, treeType, world, v1, v10, v11, v12, v2, v3, v4, v6, v7, v8
	shapeId = int32FromUint64(userData)
	queryContext = context
	broadPhase = (*b2QueryPairContext)(unsafe.Pointer(queryContext)).World + 40
	proxyKey = proxyId<<int32(2) | (*b2QueryPairContext)(unsafe.Pointer(queryContext)).QueryTreeType
	queryProxyKey = (*b2QueryPairContext)(unsafe.Pointer(queryContext)).QueryProxyKey
	// A proxy cannot form a pair with itself.
	if proxyKey == (*b2QueryPairContext)(unsafe.Pointer(queryContext)).QueryProxyKey {
		return boolUint8(true1 != 0)
	}
	treeType = (*b2QueryPairContext)(unsafe.Pointer(queryContext)).QueryTreeType
	queryProxyType = queryProxyKey & int32FromInt32(3)
	// De-duplication
	// It is important to prevent duplicate contacts from being created. Ideally I can prevent duplicates
	// early and in the worker. Most of the time the moveSet contains dynamic and kinematic proxies, but
	// sometimes it has static proxies.
	// I had an optimization here to skip checking the move set if this is a query into
	// the static tree. The assumption is that the static proxies are never in the move set
	// so there is no risk of duplication. However, this is not true with
	// b2ShapeDef::invokeContactCreation or when a static shape is modified.
	// There can easily be scenarios where the static proxy is in the moveSet but the dynamic proxy is not.
	// I could have some flag to indicate that there are any static bodies in the moveSet.
	// Is this proxy also moving?
	if queryProxyType == int32(b2_dynamicBody) {
		if treeType == int32(b2_dynamicBody) && proxyKey < queryProxyKey {
			moved = b2ContainsKey(tls, broadPhase+216, uint64FromInt32(proxyKey+int32(1)))
			if moved != 0 {
				// Both proxies are moving. Avoid duplicate pairs.
				return boolUint8(true1 != 0)
			}
		}
	} else {
		if !(treeType == int32(b2_dynamicBody)) && b2InternalAssertFcn(tls, __ccgo_ts+2646, __ccgo_ts+2470, int32FromInt32(206)) != 0 {
			__builtin_trap(tls)
		}
		moved1 = b2ContainsKey(tls, broadPhase+216, uint64FromInt32(proxyKey+int32(1)))
		if moved1 != 0 {
			// Both proxies are moving. Avoid duplicate pairs.
			return boolUint8(true1 != 0)
		}
	}
	if shapeId < (*b2QueryPairContext)(unsafe.Pointer(queryContext)).QueryShapeIndex {
		v1 = uint64FromInt32(shapeId)<<int32(32) | uint64FromInt32((*b2QueryPairContext)(unsafe.Pointer(queryContext)).QueryShapeIndex)
	} else {
		v1 = uint64FromInt32((*b2QueryPairContext)(unsafe.Pointer(queryContext)).QueryShapeIndex)<<int32(32) | uint64FromInt32(shapeId)
	}
	pairKey = v1
	if b2ContainsKey(tls, broadPhase+272, pairKey) != 0 {
		// contact exists
		return boolUint8(true1 != 0)
	}
	if proxyKey < queryProxyKey {
		shapeIdA = shapeId
		shapeIdB = (*b2QueryPairContext)(unsafe.Pointer(queryContext)).QueryShapeIndex
	} else {
		shapeIdA = (*b2QueryPairContext)(unsafe.Pointer(queryContext)).QueryShapeIndex
		shapeIdB = shapeId
	}
	world = (*b2QueryPairContext)(unsafe.Pointer(queryContext)).World
	v2 = world + 1248
	v3 = shapeIdA
	if !(0 <= v3 && v3 < (*b2ShapeArray)(unsafe.Pointer(v2)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
		__builtin_trap(tls)
	}
	v4 = (*b2ShapeArray)(unsafe.Pointer(v2)).Data + uintptr(v3)*288
	goto _5
_5:
	shapeA = v4
	v6 = world + 1248
	v7 = shapeIdB
	if !(0 <= v7 && v7 < (*b2ShapeArray)(unsafe.Pointer(v6)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
		__builtin_trap(tls)
	}
	v8 = (*b2ShapeArray)(unsafe.Pointer(v6)).Data + uintptr(v7)*288
	goto _9
_9:
	shapeB = v8
	bodyIdA = (*b2Shape)(unsafe.Pointer(shapeA)).BodyId
	bodyIdB = (*b2Shape)(unsafe.Pointer(shapeB)).BodyId
	// Are the shapes on the same body?
	if bodyIdA == bodyIdB {
		return boolUint8(true1 != 0)
	}
	// Sensors are handled elsewhere
	if (*b2Shape)(unsafe.Pointer(shapeA)).SensorIndex != -int32(1) || (*b2Shape)(unsafe.Pointer(shapeB)).SensorIndex != -int32(1) {
		return boolUint8(true1 != 0)
	}
	v10 = (*b2Shape)(unsafe.Pointer(shapeA)).Filter
	v11 = (*b2Shape)(unsafe.Pointer(shapeB)).Filter
	if v10.GroupIndex == v11.GroupIndex && v10.GroupIndex != 0 {
		v12 = boolUint8(v10.GroupIndex > 0)
		goto _13
	}
	v12 = boolUint8(v10.MaskBits&v11.CategoryBits != uint64(0) && v10.CategoryBits&v11.MaskBits != uint64(0))
	goto _13
_13:
	if int32FromUint8(v12) == false1 {
		return boolUint8(true1 != 0)
	}
	v2 = world + 1024
	v3 = bodyIdA
	if !(0 <= v3 && v3 < (*b2BodyArray)(unsafe.Pointer(v2)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v4 = (*b2BodyArray)(unsafe.Pointer(v2)).Data + uintptr(v3)*128
	goto _17
_17:
	// Does a joint override collision?
	bodyA = v4
	v6 = world + 1024
	v7 = bodyIdB
	if !(0 <= v7 && v7 < (*b2BodyArray)(unsafe.Pointer(v6)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v8 = (*b2BodyArray)(unsafe.Pointer(v6)).Data + uintptr(v7)*128
	goto _21
_21:
	bodyB = v8
	if int32FromUint8(b2ShouldBodiesCollide(tls, world, bodyA, bodyB)) == false1 {
		return boolUint8(true1 != 0)
	}
	// Custom user filter
	customFilterFcn = (*b2World)(unsafe.Pointer((*b2QueryPairContext)(unsafe.Pointer(queryContext)).World)).CustomFilterFcn
	if customFilterFcn != uintptrFromInt32(0) {
		idA = ShapeId{
			Index1:     shapeIdA + int32(1),
			World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
			Generation: (*b2Shape)(unsafe.Pointer(shapeA)).Generation,
		}
		idB = ShapeId{
			Index1:     shapeIdB + int32(1),
			World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
			Generation: (*b2Shape)(unsafe.Pointer(shapeB)).Generation,
		}
		shouldCollide = (*(*func(*_Stack, ShapeId, ShapeId, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{customFilterFcn})))(tls, idA, idB, (*b2World)(unsafe.Pointer((*b2QueryPairContext)(unsafe.Pointer(queryContext)).World)).CustomFilterContext)
		if int32FromUint8(shouldCollide) == false1 {
			return boolUint8(true1 != 0)
		}
	}
	v3 = __atomic_fetch_addInt32(tls, broadPhase+268, int32(1), int32FromInt32(__ATOMIC_SEQ_CST))
	goto _23
_23:
	// todo per thread to eliminate atomic?
	pairIndex = v3
	if pairIndex < (*b2BroadPhase)(unsafe.Pointer(broadPhase)).MovePairCapacity {
		pair = (*b2BroadPhase)(unsafe.Pointer(broadPhase)).MovePairs + uintptr(pairIndex)*24
		(*b2MovePair)(unsafe.Pointer(pair)).Heap = boolUint8(false1 != 0)
	} else {
		pair = b2Alloc(tls, int32(24))
		(*b2MovePair)(unsafe.Pointer(pair)).Heap = boolUint8(true1 != 0)
	}
	(*b2MovePair)(unsafe.Pointer(pair)).ShapeIndexA = shapeIdA
	(*b2MovePair)(unsafe.Pointer(pair)).ShapeIndexB = shapeIdB
	(*b2MovePair)(unsafe.Pointer(pair)).Next = (*b2MoveResult)(unsafe.Pointer((*b2QueryPairContext)(unsafe.Pointer(queryContext)).MoveResult)).PairList
	(*b2MoveResult)(unsafe.Pointer((*b2QueryPairContext)(unsafe.Pointer(queryContext)).MoveResult)).PairList = pair
	// continue the query
	return boolUint8(true1 != 0)
}

// Warning: writing to these globals significantly slows multithreading performance

func b2FindPairsTask(tls *_Stack, startIndex int32, endIndex int32, threadIndex uint32, context uintptr) {
	bp1 := tls.Alloc(32)
	defer tls.Free(32)
	var baseTree, bp, world uintptr
	var fatAABB AABB
	var i, proxyId, proxyKey int32
	var proxyType b2BodyType1
	var stats, statsDynamic, statsKinematic, statsStatic TreeStats
	var _ /* queryContext at bp+0 */ b2QueryPairContext
	_, _, _, _, _, _, _, _, _, _, _, _ = baseTree, bp, fatAABB, i, proxyId, proxyKey, proxyType, stats, statsDynamic, statsKinematic, statsStatic, world
	_ = uint64FromInt64(4)
	world = context
	bp = world + 40
	(*(*b2QueryPairContext)(unsafe.Add(nil, bp1))).World = world
	i = startIndex
	for {
		if !(i < endIndex) {
			break
		}
		// Initialize move result for this moved proxy
		(*(*b2QueryPairContext)(unsafe.Add(nil, bp1))).MoveResult = (*b2BroadPhase)(unsafe.Pointer(bp)).MoveResults + uintptr(i)*8
		(*b2MoveResult)(unsafe.Pointer((*(*b2QueryPairContext)(unsafe.Add(nil, bp1))).MoveResult)).PairList = uintptrFromInt32(0)
		proxyKey = *(*int32)(unsafe.Add(nil, (*b2BroadPhase)(unsafe.Pointer(bp)).MoveArray.Data+uintptr(i)*4))
		if proxyKey == -int32(1) {
			// proxy was destroyed after it moved
			goto _1
		}
		proxyType = proxyKey & int32FromInt32(3)
		proxyId = proxyKey >> int32(2)
		(*(*b2QueryPairContext)(unsafe.Add(nil, bp1))).QueryProxyKey = proxyKey
		baseTree = bp + uintptr(proxyType)*72
		// We have to query the tree with the fat AABB so that
		// we don't fail to create a contact that may touch later.
		fatAABB = b2DynamicTree_GetAABB(tls, baseTree, proxyId)
		(*(*b2QueryPairContext)(unsafe.Add(nil, bp1))).QueryShapeIndex = int32FromUint64(b2DynamicTree_GetUserData(tls, baseTree, proxyId))
		// Query trees. Only dynamic proxies collide with kinematic and static proxies.
		// Using B2_DEFAULT_MASK_BITS so that b2Filter::groupIndex works.
		stats = TreeStats{}
		if proxyType == int32(b2_dynamicBody) {
			// consider using bits = groupIndex > 0 ? B2_DEFAULT_MASK_BITS : maskBits
			(*(*b2QueryPairContext)(unsafe.Add(nil, bp1))).QueryTreeType = int32(b2_kinematicBody)
			statsKinematic = b2DynamicTree_Query(tls, bp+uintptr(b2_kinematicBody)*72, fatAABB, uint64FromUint64(0xffffffffffffffff), __ccgo_fp(b2PairQueryCallback), bp1)
			stats.NodeVisits += statsKinematic.NodeVisits
			stats.LeafVisits += statsKinematic.LeafVisits
			(*(*b2QueryPairContext)(unsafe.Add(nil, bp1))).QueryTreeType = int32(b2_staticBody)
			statsStatic = b2DynamicTree_Query(tls, bp+uintptr(b2_staticBody)*72, fatAABB, uint64FromUint64(0xffffffffffffffff), __ccgo_fp(b2PairQueryCallback), bp1)
			stats.NodeVisits += statsStatic.NodeVisits
			stats.LeafVisits += statsStatic.LeafVisits
		}
		// All proxies collide with dynamic proxies
		// Using B2_DEFAULT_MASK_BITS so that b2Filter::groupIndex works.
		(*(*b2QueryPairContext)(unsafe.Add(nil, bp1))).QueryTreeType = int32(b2_dynamicBody)
		statsDynamic = b2DynamicTree_Query(tls, bp+uintptr(b2_dynamicBody)*72, fatAABB, uint64FromUint64(0xffffffffffffffff), __ccgo_fp(b2PairQueryCallback), bp1)
		stats.NodeVisits += statsDynamic.NodeVisits
		stats.LeafVisits += statsDynamic.LeafVisits
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b2ValidateBroadphase(tls *_Stack, bp uintptr) {
	b2DynamicTree_Validate(tls, bp+uintptr(b2_dynamicBody)*72)
	b2DynamicTree_Validate(tls, bp+uintptr(b2_kinematicBody)*72)
	// TODO_ERIN validate every shape AABB is contained in tree AABB
}

func b2ValidateNoEnlarged(tls *_Stack, bp uintptr) {
	_ = uint64FromInt64(4)
}

const _B2_FORCE_OVERFLOW = 0
const _UINT64_MAX2 = "0xffffffffffffffffu"

var b2_identityBodyState2 = b2BodyState{
	DeltaRotation: Rot{
		C: float32FromFloat32(1),
	},
}

func b2CreateGraph(tls *_Stack, graph uintptr, bodyCapacity int32) {
	var color uintptr
	var i, v1, v2, v3, v5 int32
	_, _, _, _, _, _ = color, i, v1, v2, v3, v5
	*(*b2ConstraintGraph)(unsafe.Add(nil, graph)) = b2ConstraintGraph{}
	v1 = bodyCapacity
	v2 = int32(8)
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _4
_4:
	bodyCapacity = v3
	// Initialize graph color bit set.
	// No bitset for overflow color.
	i = 0
	for {
		if !(i < int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1)) {
			break
		}
		color = graph + uintptr(i)*56
		(*b2GraphColor)(unsafe.Pointer(color)).BodySet = b2CreateBitSet(tls, uint32FromInt32(bodyCapacity))
		b2SetBitCountAndClear(tls, color, uint32FromInt32(bodyCapacity))
		goto _6
	_6:
		;
		i = i + 1
	}
}

func b2DestroyGraph(tls *_Stack, graph uintptr) {
	var color uintptr
	var i int32
	_, _ = color, i
	i = 0
	for {
		if !(i < int32(_B2_GRAPH_COLOR_COUNT)) {
			break
		}
		color = graph + uintptr(i)*56
		// The bit set should never be used on the overflow color
		if !(i != int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1) || (*b2GraphColor)(unsafe.Pointer(color)).BodySet.Bits == uintptrFromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+2733, __ccgo_ts+2787, int32FromInt32(56)) != 0 {
			__builtin_trap(tls)
		}
		b2DestroyBitSet(tls, color)
		b2ContactSimArray_Destroy(tls, color+16)
		b2JointSimArray_Destroy(tls, color+32)
		goto _1
	_1:
		;
		i = i + 1
	}
}

// C documentation
//
//	// Contacts are always created as non-touching. They get cloned into the constraint
//	// graph once they are found to be touching.
//	// todo maybe kinematic bodies should not go into graph
func b2AddContactToGraph(tls *_Stack, world uintptr, contactSim uintptr, contact uintptr) {
	var awakeSet, awakeSet1, bodyA, bodyB, bodySimA, bodySimB, color, color1, color2, color3, graph, newContact, v1, v3, v5, v7 uintptr
	var blockIndex, blockIndex1, v11, v15 uint32_t
	var bodyIdA, bodyIdB, colorIndex, i, i1, i2, localIndex, localIndex1, newCapacity, v2, v6 int32
	var staticA, staticB, v12, v16 uint8
	var v18 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = awakeSet, awakeSet1, blockIndex, blockIndex1, bodyA, bodyB, bodyIdA, bodyIdB, bodySimA, bodySimB, color, color1, color2, color3, colorIndex, graph, i, i1, i2, localIndex, localIndex1, newCapacity, newContact, staticA, staticB, v1, v11, v12, v15, v16, v18, v2, v3, v5, v6, v7
	if !((*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.PointCount > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+2821, __ccgo_ts+2787, int32FromInt32(70)) != 0 {
		__builtin_trap(tls)
	}
	if !((*b2ContactSim)(unsafe.Pointer(contactSim)).SimFlags&uint32(b2_simTouchingFlag) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+2857, __ccgo_ts+2787, int32FromInt32(71)) != 0 {
		__builtin_trap(tls)
	}
	if !((*b2Contact)(unsafe.Pointer(contact)).Flags&uint32(b2_contactTouchingFlag) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+2899, __ccgo_ts+2787, int32FromInt32(72)) != 0 {
		__builtin_trap(tls)
	}
	graph = world + 328
	colorIndex = int32FromInt32(_B2_GRAPH_COLOR_COUNT) - int32FromInt32(1)
	bodyIdA = (*(*b2ContactEdge)(unsafe.Add(nil, contact+12))).BodyId
	bodyIdB = (*(*b2ContactEdge)(unsafe.Add(nil, contact+12+1*12))).BodyId
	v1 = world + 1024
	v2 = bodyIdA
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	bodyA = v3
	v5 = world + 1024
	v6 = bodyIdB
	if !(0 <= v6 && v6 < (*b2BodyArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2BodyArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*128
	goto _8
_8:
	bodyB = v7
	staticA = boolUint8((*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_staticSet))
	staticB = boolUint8((*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_staticSet))
	if !(int32FromUint8(staticA) == false1 || int32FromUint8(staticB) == false1) && b2InternalAssertFcn(tls, __ccgo_ts+2939, __ccgo_ts+2787, int32FromInt32(83)) != 0 {
		__builtin_trap(tls)
	}
	if int32FromUint8(staticA) == false1 && int32FromUint8(staticB) == false1 {
		i = 0
		for {
			if !(i < int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1)) {
				break
			}
			color = graph + uintptr(i)*56
			v1 = color
			v11 = uint32FromInt32(bodyIdA)
			blockIndex1 = v11 / uint32(64)
			if blockIndex1 >= (*b2BitSet)(unsafe.Pointer(v1)).BlockCount {
				v12 = boolUint8(false1 != 0)
				goto _13
			}
			v12 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v1)).Bits+uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v11%uint32FromInt32(64))) != uint64(0))
			goto _13
		_13:
			;
			if v18 = v12 != 0; !v18 {
				v3 = color
				v15 = uint32FromInt32(bodyIdB)
				blockIndex1 = v15 / uint32(64)
				if blockIndex1 >= (*b2BitSet)(unsafe.Pointer(v3)).BlockCount {
					v16 = boolUint8(false1 != 0)
					goto _17
				}
				v16 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v3)).Bits+uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v15%uint32FromInt32(64))) != uint64(0))
				goto _17
			_17:
			}
			if v18 || v16 != 0 {
				goto _9
			}
			v1 = color
			v11 = uint32FromInt32(bodyIdA)
			blockIndex = v11 / uint32(64)
			if blockIndex >= (*b2BitSet)(unsafe.Pointer(v1)).BlockCount {
				b2GrowBitSet(tls, v1, blockIndex+uint32(1))
			}
			*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v1)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v11 % uint32FromInt32(64))
			v1 = color
			v11 = uint32FromInt32(bodyIdB)
			blockIndex = v11 / uint32(64)
			if blockIndex >= (*b2BitSet)(unsafe.Pointer(v1)).BlockCount {
				b2GrowBitSet(tls, v1, blockIndex+uint32(1))
			}
			*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v1)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v11 % uint32FromInt32(64))
			colorIndex = i
			break
		_9:
			;
			i = i + 1
		}
	} else {
		if int32FromUint8(staticA) == false1 {
			// No static contacts in color 0
			i1 = int32(1)
			for {
				if !(i1 < int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1)) {
					break
				}
				color1 = graph + uintptr(i1)*56
				v1 = color1
				v11 = uint32FromInt32(bodyIdA)
				blockIndex1 = v11 / uint32(64)
				if blockIndex1 >= (*b2BitSet)(unsafe.Pointer(v1)).BlockCount {
					v12 = boolUint8(false1 != 0)
					goto _27
				}
				v12 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v1)).Bits+uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v11%uint32FromInt32(64))) != uint64(0))
				goto _27
			_27:
				if v12 != 0 {
					goto _23
				}
				v1 = color1
				v11 = uint32FromInt32(bodyIdA)
				blockIndex = v11 / uint32(64)
				if blockIndex >= (*b2BitSet)(unsafe.Pointer(v1)).BlockCount {
					b2GrowBitSet(tls, v1, blockIndex+uint32(1))
				}
				*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v1)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v11 % uint32FromInt32(64))
				colorIndex = i1
				break
			_23:
				;
				i1 = i1 + 1
			}
		} else {
			if int32FromUint8(staticB) == false1 {
				// No static contacts in color 0
				i2 = int32(1)
				for {
					if !(i2 < int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1)) {
						break
					}
					color2 = graph + uintptr(i2)*56
					v1 = color2
					v11 = uint32FromInt32(bodyIdB)
					blockIndex1 = v11 / uint32(64)
					if blockIndex1 >= (*b2BitSet)(unsafe.Pointer(v1)).BlockCount {
						v12 = boolUint8(false1 != 0)
						goto _34
					}
					v12 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v1)).Bits+uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v11%uint32FromInt32(64))) != uint64(0))
					goto _34
				_34:
					if v12 != 0 {
						goto _30
					}
					v1 = color2
					v11 = uint32FromInt32(bodyIdB)
					blockIndex = v11 / uint32(64)
					if blockIndex >= (*b2BitSet)(unsafe.Pointer(v1)).BlockCount {
						b2GrowBitSet(tls, v1, blockIndex+uint32(1))
					}
					*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v1)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v11 % uint32FromInt32(64))
					colorIndex = i2
					break
				_30:
					;
					i2 = i2 + 1
				}
			}
		}
	}
	color3 = graph + uintptr(colorIndex)*56
	(*b2Contact)(unsafe.Pointer(contact)).ColorIndex = colorIndex
	(*b2Contact)(unsafe.Pointer(contact)).LocalIndex = (*b2GraphColor)(unsafe.Pointer(color3)).ContactSims.Count
	v1 = color3 + 16
	if (*b2ContactSimArray)(unsafe.Pointer(v1)).Count == (*b2ContactSimArray)(unsafe.Pointer(v1)).Capacity {
		if (*b2ContactSimArray)(unsafe.Pointer(v1)).Capacity < int32(2) {
			v2 = int32(2)
		} else {
			v2 = (*b2ContactSimArray)(unsafe.Pointer(v1)).Capacity + (*b2ContactSimArray)(unsafe.Pointer(v1)).Capacity>>int32(1)
		}
		newCapacity = v2
		b2ContactSimArray_Reserve(tls, v1, newCapacity)
	}
	*(*int32)(unsafe.Add(nil, v1+8)) += int32(1)
	v3 = (*b2ContactSimArray)(unsafe.Pointer(v1)).Data + uintptr((*b2ContactSimArray)(unsafe.Pointer(v1)).Count-int32FromInt32(1))*176
	goto _40
_40:
	newContact = v3
	memcpy(tls, newContact, contactSim, uint64(176))
	// todo perhaps skip this if the contact is already awake
	if staticA != 0 {
		(*b2ContactSim)(unsafe.Pointer(newContact)).BodySimIndexA = -int32(1)
		(*b2ContactSim)(unsafe.Pointer(newContact)).InvMassA = float32FromFloat32(0)
		(*b2ContactSim)(unsafe.Pointer(newContact)).InvIA = float32FromFloat32(0)
	} else {
		if !((*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+2976, __ccgo_ts+2787, int32FromInt32(153)) != 0 {
			__builtin_trap(tls)
		}
		v1 = world + 1064
		v2 = int32(b2_awakeSet)
		if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*88
		goto _44
	_44:
		awakeSet = v3
		localIndex = (*b2Body)(unsafe.Pointer(bodyA)).LocalIndex
		(*b2ContactSim)(unsafe.Pointer(newContact)).BodySimIndexA = localIndex
		v1 = awakeSet
		v2 = localIndex
		if !(0 <= v2 && v2 < (*b2BodySimArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2BodySimArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*100
		goto _48
	_48:
		bodySimA = v3
		(*b2ContactSim)(unsafe.Pointer(newContact)).InvMassA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvMass
		(*b2ContactSim)(unsafe.Pointer(newContact)).InvIA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvInertia
	}
	if staticB != 0 {
		(*b2ContactSim)(unsafe.Pointer(newContact)).BodySimIndexB = -int32(1)
		(*b2ContactSim)(unsafe.Pointer(newContact)).InvMassB = float32FromFloat32(0)
		(*b2ContactSim)(unsafe.Pointer(newContact)).InvIB = float32FromFloat32(0)
	} else {
		if !((*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+3007, __ccgo_ts+2787, int32FromInt32(172)) != 0 {
			__builtin_trap(tls)
		}
		v1 = world + 1064
		v2 = int32(b2_awakeSet)
		if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*88
		goto _52
	_52:
		awakeSet1 = v3
		localIndex1 = (*b2Body)(unsafe.Pointer(bodyB)).LocalIndex
		(*b2ContactSim)(unsafe.Pointer(newContact)).BodySimIndexB = localIndex1
		v1 = awakeSet1
		v2 = localIndex1
		if !(0 <= v2 && v2 < (*b2BodySimArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2BodySimArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*100
		goto _56
	_56:
		bodySimB = v3
		(*b2ContactSim)(unsafe.Pointer(newContact)).InvMassB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvMass
		(*b2ContactSim)(unsafe.Pointer(newContact)).InvIB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvInertia
	}
}

func b2RemoveContactFromGraph(tls *_Stack, world uintptr, bodyIdA int32, bodyIdB int32, colorIndex int32, localIndex int32) {
	var blockIndex, v2 uint32_t
	var color, graph, movedContact, movedContactSim, v1, v4, v7 uintptr
	var movedId, movedIndex, movedIndex1, v12, v8, v9 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = blockIndex, color, graph, movedContact, movedContactSim, movedId, movedIndex, movedIndex1, v1, v12, v2, v4, v7, v8, v9
	graph = world + 328
	if !(0 <= colorIndex && colorIndex < int32(_B2_GRAPH_COLOR_COUNT)) && b2InternalAssertFcn(tls, __ccgo_ts+3038, __ccgo_ts+2787, int32FromInt32(188)) != 0 {
		__builtin_trap(tls)
	}
	color = graph + uintptr(colorIndex)*56
	if colorIndex != int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1) {
		// might clear a bit for a static body, but this has no effect
		v1 = color
		v2 = uint32FromInt32(bodyIdA)
		blockIndex = v2 / uint32(64)
		if blockIndex >= (*b2BitSet)(unsafe.Pointer(v1)).BlockCount {
			goto _3
		}
		*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v1)).Bits+uintptr(blockIndex)*8)) &= ^(uint64FromInt32(1) << (v2 % uint32FromInt32(64)))
	_3:
		;
		v1 = color
		v2 = uint32FromInt32(bodyIdB)
		blockIndex = v2 / uint32(64)
		if blockIndex >= (*b2BitSet)(unsafe.Pointer(v1)).BlockCount {
			goto _6
		}
		*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v1)).Bits+uintptr(blockIndex)*8)) &= ^(uint64FromInt32(1) << (v2 % uint32FromInt32(64)))
	_6:
	}
	v1 = color + 16
	v8 = localIndex
	if !(0 <= v8 && v8 < (*b2ContactSimArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(147)) != 0 {
		__builtin_trap(tls)
	}
	movedIndex = -int32(1)
	if v8 != (*b2ContactSimArray)(unsafe.Pointer(v1)).Count-int32FromInt32(1) {
		movedIndex = (*b2ContactSimArray)(unsafe.Pointer(v1)).Count - int32(1)
		*(*b2ContactSim)(unsafe.Add(nil, (*b2ContactSimArray)(unsafe.Pointer(v1)).Data+uintptr(v8)*176)) = *(*b2ContactSim)(unsafe.Add(nil, (*b2ContactSimArray)(unsafe.Pointer(v1)).Data+uintptr(movedIndex)*176))
	}
	*(*int32)(unsafe.Add(nil, v1+8)) -= int32(1)
	v9 = movedIndex
	goto _10
_10:
	movedIndex1 = v9
	if movedIndex1 != -int32(1) {
		// Fix index on swapped contact
		movedContactSim = (*b2GraphColor)(unsafe.Pointer(color)).ContactSims.Data + uintptr(localIndex)*176
		// Fix moved contact
		movedId = (*b2ContactSim)(unsafe.Pointer(movedContactSim)).ContactId
		v4 = world + 1144
		v12 = movedId
		if !(0 <= v12 && v12 < (*b2ContactArray)(unsafe.Pointer(v4)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
			__builtin_trap(tls)
		}
		v7 = (*b2ContactArray)(unsafe.Pointer(v4)).Data + uintptr(v12)*68
		goto _14
	_14:
		movedContact = v7
		if !((*b2Contact)(unsafe.Pointer(movedContact)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+3091, __ccgo_ts+2787, int32FromInt32(207)) != 0 {
			__builtin_trap(tls)
		}
		if !((*b2Contact)(unsafe.Pointer(movedContact)).ColorIndex == colorIndex) && b2InternalAssertFcn(tls, __ccgo_ts+3129, __ccgo_ts+2787, int32FromInt32(208)) != 0 {
			__builtin_trap(tls)
		}
		if !((*b2Contact)(unsafe.Pointer(movedContact)).LocalIndex == movedIndex1) && b2InternalAssertFcn(tls, __ccgo_ts+3168, __ccgo_ts+2787, int32FromInt32(209)) != 0 {
			__builtin_trap(tls)
		}
		(*b2Contact)(unsafe.Pointer(movedContact)).LocalIndex = localIndex
	}
}

var b2_identityBodyState3 = b2BodyState{
	DeltaRotation: Rot{
		C: float32FromFloat32(1),
	},
}

/**@}*/

func b2ContactArray_Create(tls *_Stack, capacity int32) (r b2ContactArray) {
	var a b2ContactArray
	_ = a
	a = b2ContactArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(68)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2ContactArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2ContactArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2ContactArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2ContactArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2ContactArray)(unsafe.Pointer(a)).Capacity)*uint64(68)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(68)))
	(*b2ContactArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2ContactArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2ContactArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2ContactArray)(unsafe.Pointer(a)).Capacity)*uint64(68)))
	(*b2ContactArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2ContactArray)(unsafe.Pointer(a)).Count = 0
	(*b2ContactArray)(unsafe.Pointer(a)).Capacity = 0
}

func b2ContactSimArray_Create(tls *_Stack, capacity int32) (r b2ContactSimArray) {
	var a b2ContactSimArray
	_ = a
	a = b2ContactSimArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(176)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2ContactSimArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2ContactSimArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2ContactSimArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2ContactSimArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2ContactSimArray)(unsafe.Pointer(a)).Capacity)*uint64(176)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(176)))
	(*b2ContactSimArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2ContactSimArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2ContactSimArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2ContactSimArray)(unsafe.Pointer(a)).Capacity)*uint64(176)))
	(*b2ContactSimArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2ContactSimArray)(unsafe.Pointer(a)).Count = 0
	(*b2ContactSimArray)(unsafe.Pointer(a)).Capacity = 0
}

type b2ContactRegister struct {
	Fcn     uintptr
	Primary uint8
}

var s_registers [5][5]b2ContactRegister
var s_initialized = boolUint8(false1 != 0)

func b2CircleManifold(tls *_Stack, shapeA uintptr, xfA Transform, shapeB uintptr, xfB Transform, cache uintptr) (r Manifold) {
	_ = uint64FromInt64(4)
	return b2CollideCircles(tls, shapeA+132, xfA, shapeB+132, xfB)
}

func b2CapsuleAndCircleManifold(tls *_Stack, shapeA uintptr, xfA Transform, shapeB uintptr, xfB Transform, cache uintptr) (r Manifold) {
	_ = uint64FromInt64(4)
	return b2CollideCapsuleAndCircle(tls, shapeA+132, xfA, shapeB+132, xfB)
}

func b2CapsuleManifold(tls *_Stack, shapeA uintptr, xfA Transform, shapeB uintptr, xfB Transform, cache uintptr) (r Manifold) {
	_ = uint64FromInt64(4)
	return b2CollideCapsules(tls, shapeA+132, xfA, shapeB+132, xfB)
}

func b2PolygonAndCircleManifold(tls *_Stack, shapeA uintptr, xfA Transform, shapeB uintptr, xfB Transform, cache uintptr) (r Manifold) {
	_ = uint64FromInt64(4)
	return b2CollidePolygonAndCircle(tls, shapeA+132, xfA, shapeB+132, xfB)
}

func b2PolygonAndCapsuleManifold(tls *_Stack, shapeA uintptr, xfA Transform, shapeB uintptr, xfB Transform, cache uintptr) (r Manifold) {
	_ = uint64FromInt64(4)
	return b2CollidePolygonAndCapsule(tls, shapeA+132, xfA, shapeB+132, xfB)
}

func b2PolygonManifold(tls *_Stack, shapeA uintptr, xfA Transform, shapeB uintptr, xfB Transform, cache uintptr) (r Manifold) {
	_ = uint64FromInt64(4)
	return b2CollidePolygons(tls, shapeA+132, xfA, shapeB+132, xfB)
}

func b2SegmentAndCircleManifold(tls *_Stack, shapeA uintptr, xfA Transform, shapeB uintptr, xfB Transform, cache uintptr) (r Manifold) {
	_ = uint64FromInt64(4)
	return b2CollideSegmentAndCircle(tls, shapeA+132, xfA, shapeB+132, xfB)
}

func b2SegmentAndCapsuleManifold(tls *_Stack, shapeA uintptr, xfA Transform, shapeB uintptr, xfB Transform, cache uintptr) (r Manifold) {
	_ = uint64FromInt64(4)
	return b2CollideSegmentAndCapsule(tls, shapeA+132, xfA, shapeB+132, xfB)
}

func b2SegmentAndPolygonManifold(tls *_Stack, shapeA uintptr, xfA Transform, shapeB uintptr, xfB Transform, cache uintptr) (r Manifold) {
	_ = uint64FromInt64(4)
	return b2CollideSegmentAndPolygon(tls, shapeA+132, xfA, shapeB+132, xfB)
}

func b2ChainSegmentAndCircleManifold(tls *_Stack, shapeA uintptr, xfA Transform, shapeB uintptr, xfB Transform, cache uintptr) (r Manifold) {
	_ = uint64FromInt64(4)
	return b2CollideChainSegmentAndCircle(tls, shapeA+132, xfA, shapeB+132, xfB)
}

func b2ChainSegmentAndCapsuleManifold(tls *_Stack, shapeA uintptr, xfA Transform, shapeB uintptr, xfB Transform, cache uintptr) (r Manifold) {
	return b2CollideChainSegmentAndCapsule(tls, shapeA+132, xfA, shapeB+132, xfB, cache)
}

func b2ChainSegmentAndPolygonManifold(tls *_Stack, shapeA uintptr, xfA Transform, shapeB uintptr, xfB Transform, cache uintptr) (r Manifold) {
	return b2CollideChainSegmentAndPolygon(tls, shapeA+132, xfA, shapeB+132, xfB, cache)
}

func b2AddType(tls *_Stack, __ccgo_fp_fcn uintptr, type1 ShapeType, type2 ShapeType) {
	if !(0 <= type1 && type1 < int32(b2_shapeTypeCount)) && b2InternalAssertFcn(tls, __ccgo_ts+3317, __ccgo_ts+3357, int32FromInt32(153)) != 0 {
		__builtin_trap(tls)
	}
	if !(0 <= type2 && type2 < int32(b2_shapeTypeCount)) && b2InternalAssertFcn(tls, __ccgo_ts+3382, __ccgo_ts+3357, int32FromInt32(154)) != 0 {
		__builtin_trap(tls)
	}
	(*(*b2ContactRegister)(unsafe.Add(nil, uintptr(unsafe.Pointer(&s_registers))+uintptr(type1)*80+uintptr(type2)*16))).Fcn = __ccgo_fp_fcn
	(*(*b2ContactRegister)(unsafe.Add(nil, uintptr(unsafe.Pointer(&s_registers))+uintptr(type1)*80+uintptr(type2)*16))).Primary = boolUint8(true1 != 0)
	if type1 != type2 {
		(*(*b2ContactRegister)(unsafe.Add(nil, uintptr(unsafe.Pointer(&s_registers))+uintptr(type2)*80+uintptr(type1)*16))).Fcn = __ccgo_fp_fcn
		(*(*b2ContactRegister)(unsafe.Add(nil, uintptr(unsafe.Pointer(&s_registers))+uintptr(type2)*80+uintptr(type1)*16))).Primary = boolUint8(false1 != 0)
	}
}

func b2InitializeContactRegisters(tls *_Stack) {
	if int32FromUint8(s_initialized) == false1 {
		b2AddType(tls, __ccgo_fp(b2CircleManifold), int32(b2_circleShape), int32(b2_circleShape))
		b2AddType(tls, __ccgo_fp(b2CapsuleAndCircleManifold), int32(b2_capsuleShape), int32(b2_circleShape))
		b2AddType(tls, __ccgo_fp(b2CapsuleManifold), int32(b2_capsuleShape), int32(b2_capsuleShape))
		b2AddType(tls, __ccgo_fp(b2PolygonAndCircleManifold), int32(b2_polygonShape), int32(b2_circleShape))
		b2AddType(tls, __ccgo_fp(b2PolygonAndCapsuleManifold), int32(b2_polygonShape), int32(b2_capsuleShape))
		b2AddType(tls, __ccgo_fp(b2PolygonManifold), int32(b2_polygonShape), int32(b2_polygonShape))
		b2AddType(tls, __ccgo_fp(b2SegmentAndCircleManifold), int32(b2_segmentShape), int32(b2_circleShape))
		b2AddType(tls, __ccgo_fp(b2SegmentAndCapsuleManifold), int32(b2_segmentShape), int32(b2_capsuleShape))
		b2AddType(tls, __ccgo_fp(b2SegmentAndPolygonManifold), int32(b2_segmentShape), int32(b2_polygonShape))
		b2AddType(tls, __ccgo_fp(b2ChainSegmentAndCircleManifold), int32(b2_chainSegmentShape), int32(b2_circleShape))
		b2AddType(tls, __ccgo_fp(b2ChainSegmentAndCapsuleManifold), int32(b2_chainSegmentShape), int32(b2_capsuleShape))
		b2AddType(tls, __ccgo_fp(b2ChainSegmentAndPolygonManifold), int32(b2_chainSegmentShape), int32(b2_polygonShape))
		s_initialized = boolUint8(true1 != 0)
	}
}

func b2CreateContact(tls *_Stack, world uintptr, shapeA uintptr, shapeB uintptr) {
	var bodyA, bodyB, contact, contactSim, headContact, headContact1, set, v1, v3, v5, v7 uintptr
	var contactId, headContactKey, headContactKey1, keyA, keyB, newCapacity, newCapacity1, setIndex, shapeIdA, shapeIdB, v2, v6 int32
	var pairKey uint64_t
	var type1, type2 b2ShapeType1
	var v27 uint64
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, contact, contactId, contactSim, headContact, headContact1, headContactKey, headContactKey1, keyA, keyB, newCapacity, newCapacity1, pairKey, set, setIndex, shapeIdA, shapeIdB, type1, type2, v1, v2, v27, v3, v5, v6, v7
	type1 = (*b2Shape)(unsafe.Pointer(shapeA)).Type1
	type2 = (*b2Shape)(unsafe.Pointer(shapeB)).Type1
	if !(0 <= type1 && type1 < int32(b2_shapeTypeCount)) && b2InternalAssertFcn(tls, __ccgo_ts+3317, __ccgo_ts+3357, int32FromInt32(191)) != 0 {
		__builtin_trap(tls)
	}
	if !(0 <= type2 && type2 < int32(b2_shapeTypeCount)) && b2InternalAssertFcn(tls, __ccgo_ts+3382, __ccgo_ts+3357, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	if (*(*b2ContactRegister)(unsafe.Add(nil, uintptr(unsafe.Pointer(&s_registers))+uintptr(type1)*80+uintptr(type2)*16))).Fcn == uintptrFromInt32(0) {
		// For example, no segment vs segment collision
		return
	}
	if int32FromUint8((*(*b2ContactRegister)(unsafe.Add(nil, uintptr(unsafe.Pointer(&s_registers))+uintptr(type1)*80+uintptr(type2)*16))).Primary) == false1 {
		// flip order
		b2CreateContact(tls, world, shapeB, shapeA)
		return
	}
	v1 = world + 1024
	v2 = (*b2Shape)(unsafe.Pointer(shapeA)).BodyId
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	bodyA = v3
	v5 = world + 1024
	v6 = (*b2Shape)(unsafe.Pointer(shapeB)).BodyId
	if !(0 <= v6 && v6 < (*b2BodyArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2BodyArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*128
	goto _8
_8:
	bodyB = v7
	if !((*b2Body)(unsafe.Pointer(bodyA)).SetIndex != int32(b2_disabledSet) && (*b2Body)(unsafe.Pointer(bodyB)).SetIndex != int32(b2_disabledSet)) && b2InternalAssertFcn(tls, __ccgo_ts+3422, __ccgo_ts+3357, int32FromInt32(210)) != 0 {
		__builtin_trap(tls)
	}
	if !((*b2Body)(unsafe.Pointer(bodyA)).SetIndex != int32(b2_staticSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex != int32(b2_staticSet)) && b2InternalAssertFcn(tls, __ccgo_ts+3493, __ccgo_ts+3357, int32FromInt32(211)) != 0 {
		__builtin_trap(tls)
	}
	if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet) {
		setIndex = int32(b2_awakeSet)
	} else {
		// sleeping and non-touching contacts live in the disabled set
		// later if this set is found to be touching then the sleeping
		// islands will be linked and the contact moved to the merged island
		setIndex = int32(b2_disabledSet)
	}
	v1 = world + 1064
	v2 = setIndex
	if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*88
	goto _12
_12:
	set = v3
	// Create contact key and contact
	contactId = b2AllocId(tls, world+1120)
	if contactId == (*b2World)(unsafe.Pointer(world)).Contacts.Count {
		v5 = world + 1144
		if (*b2ContactArray)(unsafe.Pointer(v5)).Count == (*b2ContactArray)(unsafe.Pointer(v5)).Capacity {
			if (*b2ContactArray)(unsafe.Pointer(v5)).Capacity < int32(2) {
				v6 = int32(2)
			} else {
				v6 = (*b2ContactArray)(unsafe.Pointer(v5)).Capacity + (*b2ContactArray)(unsafe.Pointer(v5)).Capacity>>int32(1)
			}
			newCapacity = v6
			b2ContactArray_Reserve(tls, v5, newCapacity)
		}
		*(*b2Contact)(unsafe.Add(nil, (*b2ContactArray)(unsafe.Pointer(v5)).Data+uintptr((*b2ContactArray)(unsafe.Pointer(v5)).Count)*68)) = b2Contact{}
		*(*int32)(unsafe.Add(nil, v5+8)) += int32(1)
	}
	shapeIdA = (*b2Shape)(unsafe.Pointer(shapeA)).Id
	shapeIdB = (*b2Shape)(unsafe.Pointer(shapeB)).Id
	v1 = world + 1144
	v2 = contactId
	if !(0 <= v2 && v2 < (*b2ContactArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2ContactArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*68
	goto _18
_18:
	contact = v3
	(*b2Contact)(unsafe.Pointer(contact)).ContactId = contactId
	(*b2Contact)(unsafe.Pointer(contact)).SetIndex = setIndex
	(*b2Contact)(unsafe.Pointer(contact)).ColorIndex = -int32(1)
	(*b2Contact)(unsafe.Pointer(contact)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Count
	(*b2Contact)(unsafe.Pointer(contact)).IslandId = -int32(1)
	(*b2Contact)(unsafe.Pointer(contact)).IslandPrev = -int32(1)
	(*b2Contact)(unsafe.Pointer(contact)).IslandNext = -int32(1)
	(*b2Contact)(unsafe.Pointer(contact)).ShapeIdA = shapeIdA
	(*b2Contact)(unsafe.Pointer(contact)).ShapeIdB = shapeIdB
	(*b2Contact)(unsafe.Pointer(contact)).IsMarked = boolUint8(false1 != 0)
	(*b2Contact)(unsafe.Pointer(contact)).Flags = uint32(0)
	if !((*b2Shape)(unsafe.Pointer(shapeA)).SensorIndex == -int32(1) && (*b2Shape)(unsafe.Pointer(shapeB)).SensorIndex == -int32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+3560, __ccgo_ts+3357, int32FromInt32(251)) != 0 {
		__builtin_trap(tls)
	}
	if (*b2Shape)(unsafe.Pointer(shapeA)).EnableContactEvents != 0 || (*b2Shape)(unsafe.Pointer(shapeB)).EnableContactEvents != 0 {
		*(*uint32_t)(unsafe.Add(nil, contact+60)) |= uint32(b2_contactEnableContactEvents)
	}
	// Connect to body A
	(*(*b2ContactEdge)(unsafe.Add(nil, contact+12))).BodyId = (*b2Shape)(unsafe.Pointer(shapeA)).BodyId
	(*(*b2ContactEdge)(unsafe.Add(nil, contact+12))).PrevKey = -int32(1)
	(*(*b2ContactEdge)(unsafe.Add(nil, contact+12))).NextKey = (*b2Body)(unsafe.Pointer(bodyA)).HeadContactKey
	keyA = contactId<<int32(1) | 0
	headContactKey = (*b2Body)(unsafe.Pointer(bodyA)).HeadContactKey
	if headContactKey != -int32(1) {
		v1 = world + 1144
		v2 = headContactKey >> int32(1)
		if !(0 <= v2 && v2 < (*b2ContactArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2ContactArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*68
		goto _22
	_22:
		headContact = v3
		(*(*b2ContactEdge)(unsafe.Add(nil, headContact+12+uintptr(headContactKey&int32(1))*12))).PrevKey = keyA
	}
	(*b2Body)(unsafe.Pointer(bodyA)).HeadContactKey = keyA
	*(*int32)(unsafe.Add(nil, bodyA+52)) += int32(1)
	// Connect to body B
	(*(*b2ContactEdge)(unsafe.Add(nil, contact+12+1*12))).BodyId = (*b2Shape)(unsafe.Pointer(shapeB)).BodyId
	(*(*b2ContactEdge)(unsafe.Add(nil, contact+12+1*12))).PrevKey = -int32(1)
	(*(*b2ContactEdge)(unsafe.Add(nil, contact+12+1*12))).NextKey = (*b2Body)(unsafe.Pointer(bodyB)).HeadContactKey
	keyB = contactId<<int32(1) | int32(1)
	headContactKey1 = (*b2Body)(unsafe.Pointer(bodyB)).HeadContactKey
	if (*b2Body)(unsafe.Pointer(bodyB)).HeadContactKey != -int32(1) {
		v1 = world + 1144
		v2 = headContactKey1 >> int32(1)
		if !(0 <= v2 && v2 < (*b2ContactArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2ContactArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*68
		goto _26
	_26:
		headContact1 = v3
		(*(*b2ContactEdge)(unsafe.Add(nil, headContact1+12+uintptr(headContactKey1&int32(1))*12))).PrevKey = keyB
	}
	(*b2Body)(unsafe.Pointer(bodyB)).HeadContactKey = keyB
	*(*int32)(unsafe.Add(nil, bodyB+52)) += int32(1)
	if shapeIdA < shapeIdB {
		v27 = uint64FromInt32(shapeIdA)<<int32(32) | uint64FromInt32(shapeIdB)
	} else {
		v27 = uint64FromInt32(shapeIdB)<<int32(32) | uint64FromInt32(shapeIdA)
	}
	// Add to pair set for fast lookup
	pairKey = v27
	b2AddKey(tls, world+40+272, pairKey)
	v1 = set + 48
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
	goto _31
_31:
	// Contacts are created as non-touching. Later if they are found to be touching
	// they will link islands and be moved into the constraint graph.
	contactSim = v3
	(*b2ContactSim)(unsafe.Pointer(contactSim)).ContactId = contactId
	(*b2ContactSim)(unsafe.Pointer(contactSim)).BodySimIndexA = -int32(1)
	(*b2ContactSim)(unsafe.Pointer(contactSim)).BodySimIndexB = -int32(1)
	(*b2ContactSim)(unsafe.Pointer(contactSim)).InvMassA = float32FromFloat32(0)
	(*b2ContactSim)(unsafe.Pointer(contactSim)).InvIA = float32FromFloat32(0)
	(*b2ContactSim)(unsafe.Pointer(contactSim)).InvMassB = float32FromFloat32(0)
	(*b2ContactSim)(unsafe.Pointer(contactSim)).InvIB = float32FromFloat32(0)
	(*b2ContactSim)(unsafe.Pointer(contactSim)).ShapeIdA = shapeIdA
	(*b2ContactSim)(unsafe.Pointer(contactSim)).ShapeIdB = shapeIdB
	(*b2ContactSim)(unsafe.Pointer(contactSim)).Cache = b2_emptySimplexCache
	(*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold = Manifold{}
	// These also get updated in the narrow phase
	(*b2ContactSim)(unsafe.Pointer(contactSim)).Friction = (*(*func(*_Stack, float32, int32, float32, int32) float32)(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).FrictionCallback})))(tls, (*b2Shape)(unsafe.Pointer(shapeA)).Friction, (*b2Shape)(unsafe.Pointer(shapeA)).UserMaterialId, (*b2Shape)(unsafe.Pointer(shapeB)).Friction, (*b2Shape)(unsafe.Pointer(shapeB)).UserMaterialId)
	(*b2ContactSim)(unsafe.Pointer(contactSim)).Restitution = (*(*func(*_Stack, float32, int32, float32, int32) float32)(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).RestitutionCallback})))(tls, (*b2Shape)(unsafe.Pointer(shapeA)).Restitution, (*b2Shape)(unsafe.Pointer(shapeA)).UserMaterialId, (*b2Shape)(unsafe.Pointer(shapeB)).Restitution, (*b2Shape)(unsafe.Pointer(shapeB)).UserMaterialId)
	(*b2ContactSim)(unsafe.Pointer(contactSim)).TangentSpeed = float32FromFloat32(0)
	(*b2ContactSim)(unsafe.Pointer(contactSim)).SimFlags = uint32(0)
	if (*b2Shape)(unsafe.Pointer(shapeA)).EnablePreSolveEvents != 0 || (*b2Shape)(unsafe.Pointer(shapeB)).EnablePreSolveEvents != 0 {
		*(*uint32_t)(unsafe.Add(nil, contactSim+164)) |= uint32(b2_simEnablePreSolveEvents)
	}
}

// C documentation
//
//	// A contact is destroyed when:
//	// - broad-phase proxies stop overlapping
//	// - a body is destroyed
//	// - a body is disabled
//	// - a body changes type from dynamic to kinematic or static
//	// - a shape is destroyed
//	// - contact filtering is modified
func b2DestroyContact(tls *_Stack, world uintptr, contact uintptr, wakeBodies uint8) {
	var bodyA, bodyB, edgeA, edgeB, movedContact, movedContactSim, nextContact, nextContact1, nextEdge, nextEdge1, prevContact, prevContact1, prevEdge, prevEdge1, set, shapeA, shapeB, v10, v12, v14, v16, v18, v2, v4, v6, v8 uintptr
	var bodyIdA, bodyIdB, contactId, edgeKeyA, edgeKeyB, movedIndex, movedIndex1, newCapacity, v11, v15, v19, v3, v7 int32
	var event ContactEndTouchEvent
	var flags uint32_t
	var pairKey uint64_t
	var shapeIdA, shapeIdB ShapeId
	var touching uint8
	var worldId uint16_t
	var v1 uint64
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, bodyIdA, bodyIdB, contactId, edgeA, edgeB, edgeKeyA, edgeKeyB, event, flags, movedContact, movedContactSim, movedIndex, movedIndex1, newCapacity, nextContact, nextContact1, nextEdge, nextEdge1, pairKey, prevContact, prevContact1, prevEdge, prevEdge1, set, shapeA, shapeB, shapeIdA, shapeIdB, touching, worldId, v1, v10, v11, v12, v14, v15, v16, v18, v19, v2, v3, v4, v6, v7, v8
	if (*b2Contact)(unsafe.Pointer(contact)).ShapeIdA < (*b2Contact)(unsafe.Pointer(contact)).ShapeIdB {
		v1 = uint64FromInt32((*b2Contact)(unsafe.Pointer(contact)).ShapeIdA)<<int32(32) | uint64FromInt32((*b2Contact)(unsafe.Pointer(contact)).ShapeIdB)
	} else {
		v1 = uint64FromInt32((*b2Contact)(unsafe.Pointer(contact)).ShapeIdB)<<int32(32) | uint64FromInt32((*b2Contact)(unsafe.Pointer(contact)).ShapeIdA)
	}
	// Remove pair from set
	pairKey = v1
	b2RemoveKey(tls, world+40+272, pairKey)
	edgeA = contact + 12 + uintptr(0)*12
	edgeB = contact + 12 + uintptr(1)*12
	bodyIdA = (*b2ContactEdge)(unsafe.Pointer(edgeA)).BodyId
	bodyIdB = (*b2ContactEdge)(unsafe.Pointer(edgeB)).BodyId
	v2 = world + 1024
	v3 = bodyIdA
	if !(0 <= v3 && v3 < (*b2BodyArray)(unsafe.Pointer(v2)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v4 = (*b2BodyArray)(unsafe.Pointer(v2)).Data + uintptr(v3)*128
	goto _5
_5:
	bodyA = v4
	v6 = world + 1024
	v7 = bodyIdB
	if !(0 <= v7 && v7 < (*b2BodyArray)(unsafe.Pointer(v6)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v8 = (*b2BodyArray)(unsafe.Pointer(v6)).Data + uintptr(v7)*128
	goto _9
_9:
	bodyB = v8
	flags = (*b2Contact)(unsafe.Pointer(contact)).Flags
	touching = boolUint8(flags&uint32(b2_contactTouchingFlag) != uint32(0))
	// End touch event
	if touching != 0 && flags&uint32(b2_contactEnableContactEvents) != uint32(0) {
		worldId = (*b2World)(unsafe.Pointer(world)).WorldId
		v10 = world + 1248
		v11 = (*b2Contact)(unsafe.Pointer(contact)).ShapeIdA
		if !(0 <= v11 && v11 < (*b2ShapeArray)(unsafe.Pointer(v10)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v12 = (*b2ShapeArray)(unsafe.Pointer(v10)).Data + uintptr(v11)*288
		goto _13
	_13:
		shapeA = v12
		v14 = world + 1248
		v15 = (*b2Contact)(unsafe.Pointer(contact)).ShapeIdB
		if !(0 <= v15 && v15 < (*b2ShapeArray)(unsafe.Pointer(v14)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v16 = (*b2ShapeArray)(unsafe.Pointer(v14)).Data + uintptr(v15)*288
		goto _17
	_17:
		shapeB = v16
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
		event = ContactEndTouchEvent{
			ShapeIdA: shapeIdA,
			ShapeIdB: shapeIdB,
		}
		v18 = world + 1408 + uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16
		if (*b2ContactEndTouchEventArray)(unsafe.Pointer(v18)).Count == (*b2ContactEndTouchEventArray)(unsafe.Pointer(v18)).Capacity {
			if (*b2ContactEndTouchEventArray)(unsafe.Pointer(v18)).Capacity < int32(2) {
				v19 = int32(2)
			} else {
				v19 = (*b2ContactEndTouchEventArray)(unsafe.Pointer(v18)).Capacity + (*b2ContactEndTouchEventArray)(unsafe.Pointer(v18)).Capacity>>int32(1)
			}
			newCapacity = v19
			b2ContactEndTouchEventArray_Reserve(tls, v18, newCapacity)
		}
		*(*ContactEndTouchEvent)(unsafe.Add(nil, (*b2ContactEndTouchEventArray)(unsafe.Pointer(v18)).Data+uintptr((*b2ContactEndTouchEventArray)(unsafe.Pointer(v18)).Count)*16)) = event
		*(*int32)(unsafe.Add(nil, v18+8)) += int32(1)
	}
	// Remove from body A
	if (*b2ContactEdge)(unsafe.Pointer(edgeA)).PrevKey != -int32(1) {
		v2 = world + 1144
		v3 = (*b2ContactEdge)(unsafe.Pointer(edgeA)).PrevKey >> int32(1)
		if !(0 <= v3 && v3 < (*b2ContactArray)(unsafe.Pointer(v2)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
			__builtin_trap(tls)
		}
		v4 = (*b2ContactArray)(unsafe.Pointer(v2)).Data + uintptr(v3)*68
		goto _23
	_23:
		prevContact = v4
		prevEdge = prevContact + 12 + uintptr((*b2ContactEdge)(unsafe.Pointer(edgeA)).PrevKey&int32FromInt32(1))*12
		(*b2ContactEdge)(unsafe.Pointer(prevEdge)).NextKey = (*b2ContactEdge)(unsafe.Pointer(edgeA)).NextKey
	}
	if (*b2ContactEdge)(unsafe.Pointer(edgeA)).NextKey != -int32(1) {
		v2 = world + 1144
		v3 = (*b2ContactEdge)(unsafe.Pointer(edgeA)).NextKey >> int32(1)
		if !(0 <= v3 && v3 < (*b2ContactArray)(unsafe.Pointer(v2)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
			__builtin_trap(tls)
		}
		v4 = (*b2ContactArray)(unsafe.Pointer(v2)).Data + uintptr(v3)*68
		goto _27
	_27:
		nextContact = v4
		nextEdge = nextContact + 12 + uintptr((*b2ContactEdge)(unsafe.Pointer(edgeA)).NextKey&int32FromInt32(1))*12
		(*b2ContactEdge)(unsafe.Pointer(nextEdge)).PrevKey = (*b2ContactEdge)(unsafe.Pointer(edgeA)).PrevKey
	}
	contactId = (*b2Contact)(unsafe.Pointer(contact)).ContactId
	edgeKeyA = contactId<<int32(1) | 0
	if (*b2Body)(unsafe.Pointer(bodyA)).HeadContactKey == edgeKeyA {
		(*b2Body)(unsafe.Pointer(bodyA)).HeadContactKey = (*b2ContactEdge)(unsafe.Pointer(edgeA)).NextKey
	}
	*(*int32)(unsafe.Add(nil, bodyA+52)) -= int32(1)
	// Remove from body B
	if (*b2ContactEdge)(unsafe.Pointer(edgeB)).PrevKey != -int32(1) {
		v2 = world + 1144
		v3 = (*b2ContactEdge)(unsafe.Pointer(edgeB)).PrevKey >> int32(1)
		if !(0 <= v3 && v3 < (*b2ContactArray)(unsafe.Pointer(v2)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
			__builtin_trap(tls)
		}
		v4 = (*b2ContactArray)(unsafe.Pointer(v2)).Data + uintptr(v3)*68
		goto _31
	_31:
		prevContact1 = v4
		prevEdge1 = prevContact1 + 12 + uintptr((*b2ContactEdge)(unsafe.Pointer(edgeB)).PrevKey&int32FromInt32(1))*12
		(*b2ContactEdge)(unsafe.Pointer(prevEdge1)).NextKey = (*b2ContactEdge)(unsafe.Pointer(edgeB)).NextKey
	}
	if (*b2ContactEdge)(unsafe.Pointer(edgeB)).NextKey != -int32(1) {
		v2 = world + 1144
		v3 = (*b2ContactEdge)(unsafe.Pointer(edgeB)).NextKey >> int32(1)
		if !(0 <= v3 && v3 < (*b2ContactArray)(unsafe.Pointer(v2)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
			__builtin_trap(tls)
		}
		v4 = (*b2ContactArray)(unsafe.Pointer(v2)).Data + uintptr(v3)*68
		goto _35
	_35:
		nextContact1 = v4
		nextEdge1 = nextContact1 + 12 + uintptr((*b2ContactEdge)(unsafe.Pointer(edgeB)).NextKey&int32FromInt32(1))*12
		(*b2ContactEdge)(unsafe.Pointer(nextEdge1)).PrevKey = (*b2ContactEdge)(unsafe.Pointer(edgeB)).PrevKey
	}
	edgeKeyB = contactId<<int32(1) | int32(1)
	if (*b2Body)(unsafe.Pointer(bodyB)).HeadContactKey == edgeKeyB {
		(*b2Body)(unsafe.Pointer(bodyB)).HeadContactKey = (*b2ContactEdge)(unsafe.Pointer(edgeB)).NextKey
	}
	*(*int32)(unsafe.Add(nil, bodyB+52)) -= int32(1)
	// Remove contact from the array that owns it
	if (*b2Contact)(unsafe.Pointer(contact)).IslandId != -int32(1) {
		b2UnlinkContact(tls, world, contact)
	}
	if (*b2Contact)(unsafe.Pointer(contact)).ColorIndex != -int32(1) {
		// contact is an active constraint
		if !((*b2Contact)(unsafe.Pointer(contact)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+3637, __ccgo_ts+3357, int32FromInt32(424)) != 0 {
			__builtin_trap(tls)
		}
		b2RemoveContactFromGraph(tls, world, bodyIdA, bodyIdB, (*b2Contact)(unsafe.Pointer(contact)).ColorIndex, (*b2Contact)(unsafe.Pointer(contact)).LocalIndex)
	} else {
		// contact is non-touching or is sleeping
		if !((*b2Contact)(unsafe.Pointer(contact)).SetIndex != int32(b2_awakeSet) || (*b2Contact)(unsafe.Pointer(contact)).Flags&uint32(b2_contactTouchingFlag) == uint32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+3670, __ccgo_ts+3357, int32FromInt32(430)) != 0 {
			__builtin_trap(tls)
		}
		v2 = world + 1064
		v3 = (*b2Contact)(unsafe.Pointer(contact)).SetIndex
		if !(0 <= v3 && v3 < (*b2SolverSetArray)(unsafe.Pointer(v2)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
			__builtin_trap(tls)
		}
		v4 = (*b2SolverSetArray)(unsafe.Pointer(v2)).Data + uintptr(v3)*88
		goto _39
	_39:
		set = v4
		v6 = set + 48
		v7 = (*b2Contact)(unsafe.Pointer(contact)).LocalIndex
		if !(0 <= v7 && v7 < (*b2ContactSimArray)(unsafe.Pointer(v6)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(147)) != 0 {
			__builtin_trap(tls)
		}
		movedIndex = -int32(1)
		if v7 != (*b2ContactSimArray)(unsafe.Pointer(v6)).Count-int32FromInt32(1) {
			movedIndex = (*b2ContactSimArray)(unsafe.Pointer(v6)).Count - int32(1)
			*(*b2ContactSim)(unsafe.Add(nil, (*b2ContactSimArray)(unsafe.Pointer(v6)).Data+uintptr(v7)*176)) = *(*b2ContactSim)(unsafe.Add(nil, (*b2ContactSimArray)(unsafe.Pointer(v6)).Data+uintptr(movedIndex)*176))
		}
		*(*int32)(unsafe.Add(nil, v6+8)) -= int32(1)
		v11 = movedIndex
		goto _43
	_43:
		movedIndex1 = v11
		if movedIndex1 != -int32(1) {
			movedContactSim = (*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Data + uintptr((*b2Contact)(unsafe.Pointer(contact)).LocalIndex)*176
			v8 = world + 1144
			v15 = (*b2ContactSim)(unsafe.Pointer(movedContactSim)).ContactId
			if !(0 <= v15 && v15 < (*b2ContactArray)(unsafe.Pointer(v8)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
				__builtin_trap(tls)
			}
			v10 = (*b2ContactArray)(unsafe.Pointer(v8)).Data + uintptr(v15)*68
			goto _47
		_47:
			movedContact = v10
			(*b2Contact)(unsafe.Pointer(movedContact)).LocalIndex = (*b2Contact)(unsafe.Pointer(contact)).LocalIndex
		}
	}
	(*b2Contact)(unsafe.Pointer(contact)).ContactId = -int32(1)
	(*b2Contact)(unsafe.Pointer(contact)).SetIndex = -int32(1)
	(*b2Contact)(unsafe.Pointer(contact)).ColorIndex = -int32(1)
	(*b2Contact)(unsafe.Pointer(contact)).LocalIndex = -int32(1)
	b2FreeId(tls, world+1120, contactId)
	if wakeBodies != 0 && touching != 0 {
		b2WakeBody(tls, world, bodyA)
		b2WakeBody(tls, world, bodyB)
	}
}

func b2GetContactSim(tls *_Stack, world uintptr, contact uintptr) (r uintptr) {
	var color, set, v1, v3, v5, v7 uintptr
	var v2, v6 int32
	_, _, _, _, _, _, _, _ = color, set, v1, v2, v3, v5, v6, v7
	if (*b2Contact)(unsafe.Pointer(contact)).SetIndex == int32(b2_awakeSet) && (*b2Contact)(unsafe.Pointer(contact)).ColorIndex != -int32(1) {
		// contact lives in constraint graph
		if !(0 <= (*b2Contact)(unsafe.Pointer(contact)).ColorIndex && (*b2Contact)(unsafe.Pointer(contact)).ColorIndex < int32(_B2_GRAPH_COLOR_COUNT)) && b2InternalAssertFcn(tls, __ccgo_ts+3755, __ccgo_ts+3357, int32FromInt32(461)) != 0 {
			__builtin_trap(tls)
		}
		color = world + 328 + uintptr((*b2Contact)(unsafe.Pointer(contact)).ColorIndex)*56
		v1 = color + 16
		v2 = (*b2Contact)(unsafe.Pointer(contact)).LocalIndex
		if !(0 <= v2 && v2 < (*b2ContactSimArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(147)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2ContactSimArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*176
		goto _4
	_4:
		return v3
	}
	v1 = world + 1064
	v2 = (*b2Contact)(unsafe.Pointer(contact)).SetIndex
	if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*88
	goto _8
_8:
	set = v3
	v5 = set + 48
	v6 = (*b2Contact)(unsafe.Pointer(contact)).LocalIndex
	if !(0 <= v6 && v6 < (*b2ContactSimArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(147)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2ContactSimArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*176
	goto _12
_12:
	return v7
}

// C documentation
//
//	// Update the contact manifold and touching status.
//	// Note: do not assume the shape AABBs are overlapping or are valid.
func b2UpdateContact(tls *_Stack, world uintptr, contactSim uintptr, shapeA uintptr, transformA Transform, centerOffsetA Vec2, shapeB uintptr, transformB Transform, centerOffsetB Vec2) (r uint8) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var fcn, mp1, mp2, v1, v4 uintptr
	var i, j, pointCount, unmatchedCount, v27 int32
	var id2 uint16_t
	var maxRadius, radiusA, radiusB, v11, v12, v13, v14, v16, v2, v5, v7, v8, v9 float32
	var shapeIdA, shapeIdB ShapeId
	var touching uint8
	var v18, v19, v20 Vec2
	var _ /* oldManifold at bp+0 */ Manifold
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = fcn, i, id2, j, maxRadius, mp1, mp2, pointCount, radiusA, radiusB, shapeIdA, shapeIdB, touching, unmatchedCount, v1, v11, v12, v13, v14, v16, v18, v19, v2, v20, v27, v4, v5, v7, v8, v9
	// Save old manifold
	*(*Manifold)(unsafe.Add(nil, bp)) = (*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold
	// Compute new manifold
	fcn = (*(*b2ContactRegister)(unsafe.Add(nil, uintptr(unsafe.Pointer(&s_registers))+uintptr((*b2Shape)(unsafe.Pointer(shapeA)).Type1)*80+uintptr((*b2Shape)(unsafe.Pointer(shapeB)).Type1)*16))).Fcn
	(*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold = (*(*func(*_Stack, uintptr, Transform, uintptr, Transform, uintptr) Manifold)(unsafe.Pointer(&struct{ uintptr }{fcn})))(tls, shapeA, transformA, shapeB, transformB, contactSim+168)
	// Keep these updated in case the values on the shapes are modified
	(*b2ContactSim)(unsafe.Pointer(contactSim)).Friction = (*(*func(*_Stack, float32, int32, float32, int32) float32)(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).FrictionCallback})))(tls, (*b2Shape)(unsafe.Pointer(shapeA)).Friction, (*b2Shape)(unsafe.Pointer(shapeA)).UserMaterialId, (*b2Shape)(unsafe.Pointer(shapeB)).Friction, (*b2Shape)(unsafe.Pointer(shapeB)).UserMaterialId)
	(*b2ContactSim)(unsafe.Pointer(contactSim)).Restitution = (*(*func(*_Stack, float32, int32, float32, int32) float32)(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).RestitutionCallback})))(tls, (*b2Shape)(unsafe.Pointer(shapeA)).Restitution, (*b2Shape)(unsafe.Pointer(shapeA)).UserMaterialId, (*b2Shape)(unsafe.Pointer(shapeB)).Restitution, (*b2Shape)(unsafe.Pointer(shapeB)).UserMaterialId)
	// todo branch improves perf?
	if (*b2Shape)(unsafe.Pointer(shapeA)).RollingResistance > float32FromFloat32(0) || (*b2Shape)(unsafe.Pointer(shapeB)).RollingResistance > float32FromFloat32(0) {
		v1 = shapeA
		switch (*b2Shape)(unsafe.Pointer(v1)).Type1 {
		case int32(b2_capsuleShape):
			v2 = (*b2Shape)(unsafe.Pointer(v1)).ｆ__ccgo19_132.Capsule.Radius
			goto _3
		case int32(b2_circleShape):
			v2 = (*(*Circle)(unsafe.Add(unsafe.Pointer(v1), 132))).Radius
			goto _3
		case int32(b2_polygonShape):
			v2 = (*(*Polygon)(unsafe.Add(unsafe.Pointer(v1), 132))).Radius
			goto _3
		default:
			v2 = float32FromFloat32(0)
			goto _3
		}
	_3:
		radiusA = v2
		v4 = shapeB
		switch (*b2Shape)(unsafe.Pointer(v4)).Type1 {
		case int32(b2_capsuleShape):
			v5 = (*b2Shape)(unsafe.Pointer(v4)).ｆ__ccgo19_132.Capsule.Radius
			goto _6
		case int32(b2_circleShape):
			v5 = (*(*Circle)(unsafe.Add(unsafe.Pointer(v4), 132))).Radius
			goto _6
		case int32(b2_polygonShape):
			v5 = (*(*Polygon)(unsafe.Add(unsafe.Pointer(v4), 132))).Radius
			goto _6
		default:
			v5 = float32FromFloat32(0)
			goto _6
		}
	_6:
		radiusB = v5
		v7 = radiusA
		v8 = radiusB
		if v7 > v8 {
			v11 = v7
		} else {
			v11 = v8
		}
		v9 = v11
		goto _10
	_10:
		maxRadius = v9
		v12 = (*b2Shape)(unsafe.Pointer(shapeA)).RollingResistance
		v13 = (*b2Shape)(unsafe.Pointer(shapeB)).RollingResistance
		if v12 > v13 {
			v16 = v12
		} else {
			v16 = v13
		}
		v14 = v16
		goto _15
	_15:
		(*b2ContactSim)(unsafe.Pointer(contactSim)).RollingResistance = float32(v14 * maxRadius)
	} else {
		(*b2ContactSim)(unsafe.Pointer(contactSim)).RollingResistance = float32FromFloat32(0)
	}
	(*b2ContactSim)(unsafe.Pointer(contactSim)).TangentSpeed = (*b2Shape)(unsafe.Pointer(shapeA)).TangentSpeed + (*b2Shape)(unsafe.Pointer(shapeB)).TangentSpeed
	pointCount = (*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.PointCount
	touching = boolUint8(pointCount > 0)
	if touching != 0 && (*b2World)(unsafe.Pointer(world)).PreSolveFcn != uintptrFromInt32(0) && (*b2ContactSim)(unsafe.Pointer(contactSim)).SimFlags&uint32(b2_simEnablePreSolveEvents) != uint32(0) {
		shapeIdA = ShapeId{
			Index1:     (*b2Shape)(unsafe.Pointer(shapeA)).Id + int32(1),
			World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
			Generation: (*b2Shape)(unsafe.Pointer(shapeA)).Generation,
		}
		shapeIdB = ShapeId{
			Index1:     (*b2Shape)(unsafe.Pointer(shapeB)).Id + int32(1),
			World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
			Generation: (*b2Shape)(unsafe.Pointer(shapeB)).Generation,
		}
		// this call assumes thread safety
		touching = (*(*func(*_Stack, ShapeId, ShapeId, uintptr, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).PreSolveFcn})))(tls, shapeIdA, shapeIdB, contactSim+36, (*b2World)(unsafe.Pointer(world)).PreSolveContext)
		if int32FromUint8(touching) == false1 {
			// disable contact
			pointCount = 0
			(*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.PointCount = 0
		}
	}
	// This flag is for testing
	if int32FromUint8((*b2World)(unsafe.Pointer(world)).EnableSpeculative) == false1 && pointCount == int32(2) {
		if (*(*ManifoldPoint)(unsafe.Add(nil, contactSim+36+12))).Separation > float32(float32FromFloat32(1.5)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
			*(*ManifoldPoint)(unsafe.Add(nil, contactSim+36+12)) = *(*ManifoldPoint)(unsafe.Add(nil, contactSim+36+12+1*48))
			(*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.PointCount = int32(1)
		} else {
			if (*(*ManifoldPoint)(unsafe.Add(nil, contactSim+36+12))).Separation > float32(float32FromFloat32(1.5)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
				(*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.PointCount = int32(1)
			}
		}
		pointCount = (*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.PointCount
	}
	if touching != 0 && ((*b2Shape)(unsafe.Pointer(shapeA)).EnableHitEvents != 0 || (*b2Shape)(unsafe.Pointer(shapeB)).EnableHitEvents != 0) {
		*(*uint32_t)(unsafe.Add(nil, contactSim+164)) |= uint32(b2_simEnableHitEvent)
	} else {
		*(*uint32_t)(unsafe.Add(nil, contactSim+164)) &= uint32FromInt32(^int32(b2_simEnableHitEvent))
	}
	if pointCount > 0 {
		(*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.RollingImpulse = (*(*Manifold)(unsafe.Add(nil, bp))).RollingImpulse
	}
	// Match old contact ids to new contact ids and copy the
	// stored impulses to warm start the solver.
	unmatchedCount = 0
	i = 0
	for {
		if !(i < pointCount) {
			break
		}
		mp2 = contactSim + 36 + 12 + uintptr(i)*48
		// shift anchors to be center of mass relative
		v18 = (*ManifoldPoint)(unsafe.Pointer(mp2)).AnchorA
		v19 = centerOffsetA
		v20 = Vec2{
			X: v18.X - v19.X,
			Y: v18.Y - v19.Y,
		}
		goto _21
	_21:
		(*ManifoldPoint)(unsafe.Pointer(mp2)).AnchorA = v20
		v18 = (*ManifoldPoint)(unsafe.Pointer(mp2)).AnchorB
		v19 = centerOffsetB
		v20 = Vec2{
			X: v18.X - v19.X,
			Y: v18.Y - v19.Y,
		}
		goto _25
	_25:
		(*ManifoldPoint)(unsafe.Pointer(mp2)).AnchorB = v20
		(*ManifoldPoint)(unsafe.Pointer(mp2)).NormalImpulse = float32FromFloat32(0)
		(*ManifoldPoint)(unsafe.Pointer(mp2)).TangentImpulse = float32FromFloat32(0)
		(*ManifoldPoint)(unsafe.Pointer(mp2)).TotalNormalImpulse = float32FromFloat32(0)
		(*ManifoldPoint)(unsafe.Pointer(mp2)).NormalVelocity = float32FromFloat32(0)
		(*ManifoldPoint)(unsafe.Pointer(mp2)).Persisted = boolUint8(false1 != 0)
		id2 = (*ManifoldPoint)(unsafe.Pointer(mp2)).Id
		j = 0
		for {
			if !(j < (*(*Manifold)(unsafe.Add(nil, bp))).PointCount) {
				break
			}
			mp1 = bp + 12 + uintptr(j)*48
			if int32FromUint16((*ManifoldPoint)(unsafe.Pointer(mp1)).Id) == int32FromUint16(id2) {
				(*ManifoldPoint)(unsafe.Pointer(mp2)).NormalImpulse = (*ManifoldPoint)(unsafe.Pointer(mp1)).NormalImpulse
				(*ManifoldPoint)(unsafe.Pointer(mp2)).TangentImpulse = (*ManifoldPoint)(unsafe.Pointer(mp1)).TangentImpulse
				(*ManifoldPoint)(unsafe.Pointer(mp2)).Persisted = boolUint8(true1 != 0)
				// clear old impulse
				(*ManifoldPoint)(unsafe.Pointer(mp1)).NormalImpulse = float32FromFloat32(0)
				(*ManifoldPoint)(unsafe.Pointer(mp1)).TangentImpulse = float32FromFloat32(0)
				break
			}
			goto _26
		_26:
			;
			j = j + 1
		}
		if (*ManifoldPoint)(unsafe.Pointer(mp2)).Persisted != 0 {
			v27 = 0
		} else {
			v27 = int32(1)
		}
		unmatchedCount = unmatchedCount + v27
		goto _17
	_17:
		;
		i = i + 1
	}
	_ = uint64FromInt64(4)
	if touching != 0 {
		*(*uint32_t)(unsafe.Add(nil, contactSim+164)) |= uint32(b2_simTouchingFlag)
	} else {
		*(*uint32_t)(unsafe.Add(nil, contactSim+164)) &= uint32FromInt32(^int32(b2_simTouchingFlag))
	}
	return touching
}

type b2ContactConstraintPoint struct {
	AnchorA            Vec2
	AnchorB            Vec2
	BaseSeparation     float32
	RelativeVelocity   float32
	NormalImpulse      float32
	TangentImpulse     float32
	TotalNormalImpulse float32
	NormalMass         float32
	TangentMass        float32
}

type b2ContactConstraint struct {
	IndexA            int32
	IndexB            int32
	Points            [2]b2ContactConstraintPoint
	Normal            Vec2
	InvMassA          float32
	InvMassB          float32
	InvIA             float32
	InvIB             float32
	Friction          float32
	Restitution       float32
	TangentSpeed      float32
	RollingResistance float32
	RollingMass       float32
	RollingImpulse    float32
	Softness          b2Softness
	PointCount        int32
}

var b2_identityBodyState4 = b2BodyState{
	DeltaRotation: Rot{
		C: float32FromFloat32(1),
	},
}

type b2ContactConstraintSIMD struct {
	IndexA              [4]int32
	IndexB              [4]int32
	InvMassA            b2FloatW
	InvMassB            b2FloatW
	InvIA               b2FloatW
	InvIB               b2FloatW
	Normal              b2Vec2W
	Friction            b2FloatW
	TangentSpeed        b2FloatW
	RollingResistance   b2FloatW
	RollingMass         b2FloatW
	RollingImpulse      b2FloatW
	BiasRate            b2FloatW
	MassScale           b2FloatW
	ImpulseScale        b2FloatW
	AnchorA1            b2Vec2W
	AnchorB1            b2Vec2W
	NormalMass1         b2FloatW
	TangentMass1        b2FloatW
	BaseSeparation1     b2FloatW
	NormalImpulse1      b2FloatW
	TotalNormalImpulse1 b2FloatW
	TangentImpulse1     b2FloatW
	AnchorA2            b2Vec2W
	AnchorB2            b2Vec2W
	BaseSeparation2     b2FloatW
	NormalImpulse2      b2FloatW
	TotalNormalImpulse2 b2FloatW
	TangentImpulse2     b2FloatW
	NormalMass2         b2FloatW
	TangentMass2        b2FloatW
	Restitution         b2FloatW
	RelativeVelocity1   b2FloatW
	RelativeVelocity2   b2FloatW
}

// contact separation for sub-stepping
// s = s0 + dot(cB + rB - cA - rA, normal)
// normal is held constant
// body positions c can translation and anchors r can rotate
// s(t) = s0 + dot(cB(t) + rB(t) - cA(t) - rA(t), normal)
// s(t) = s0 + dot(cB0 + dpB + rot(dqB, rB0) - cA0 - dpA - rot(dqA, rA0), normal)
// s(t) = s0 + dot(cB0 - cA0, normal) + dot(dpB - dpA + rot(dqB, rB0) - rot(dqA, rA0), normal)
// s_base = s0 + dot(cB0 - cA0, normal)

func b2PrepareOverflowContacts(tls *_Stack, context uintptr) {
	var awakeStates, color, constraint, constraints, contactSim, contacts, cp, graph, manifold, mp, stateA, stateB, world uintptr
	var contactCount, i, indexA, indexB, j, pointCount int32
	var contactSoftness, staticSoftness b2Softness
	var iA, iB, k, kNormal, kTangent, mA, mB, rnA, rnB, rtA, rtB, wA, wB, warmStartScale, v1, v14, v3 float32
	var normal, rA, rB, tangent, vA, vB, vrA, vrB, v10, v12, v13, v16, v17, v20, v21, v25, v26, v29, v30, v4, v5, v8, v9 Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = awakeStates, color, constraint, constraints, contactCount, contactSim, contactSoftness, contacts, cp, graph, i, iA, iB, indexA, indexB, j, k, kNormal, kTangent, mA, mB, manifold, mp, normal, pointCount, rA, rB, rnA, rnB, rtA, rtB, stateA, stateB, staticSoftness, tangent, vA, vB, vrA, vrB, wA, wB, warmStartScale, world, v1, v10, v12, v13, v14, v16, v17, v20, v21, v25, v26, v29, v3, v30, v4, v5, v8, v9
	world = (*b2StepContext)(unsafe.Pointer(context)).World
	graph = (*b2StepContext)(unsafe.Pointer(context)).Graph
	color = graph + uintptr(int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1))*56
	constraints = *(*uintptr)(unsafe.Add(unsafe.Pointer(color), 48))
	contactCount = (*b2GraphColor)(unsafe.Pointer(color)).ContactSims.Count
	contacts = (*b2GraphColor)(unsafe.Pointer(color)).ContactSims.Data
	awakeStates = (*b2StepContext)(unsafe.Pointer(context)).States
	// Stiffer for static contacts to avoid bodies getting pushed through the ground
	contactSoftness = (*b2StepContext)(unsafe.Pointer(context)).ContactSoftness
	staticSoftness = (*b2StepContext)(unsafe.Pointer(context)).StaticSoftness
	if (*b2World)(unsafe.Pointer(world)).EnableWarmStarting != 0 {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	warmStartScale = v1
	i = 0
	for {
		if !(i < contactCount) {
			break
		}
		contactSim = contacts + uintptr(i)*176
		manifold = contactSim + 36
		pointCount = (*Manifold)(unsafe.Pointer(manifold)).PointCount
		if !(0 < pointCount && pointCount <= int32(2)) && b2InternalAssertFcn(tls, __ccgo_ts+3826, __ccgo_ts+3860, int32FromInt32(53)) != 0 {
			__builtin_trap(tls)
		}
		indexA = (*b2ContactSim)(unsafe.Pointer(contactSim)).BodySimIndexA
		indexB = (*b2ContactSim)(unsafe.Pointer(contactSim)).BodySimIndexB
		constraint = constraints + uintptr(i)*160
		(*b2ContactConstraint)(unsafe.Pointer(constraint)).IndexA = indexA
		(*b2ContactConstraint)(unsafe.Pointer(constraint)).IndexB = indexB
		(*b2ContactConstraint)(unsafe.Pointer(constraint)).Normal = (*Manifold)(unsafe.Pointer(manifold)).Normal
		(*b2ContactConstraint)(unsafe.Pointer(constraint)).Friction = (*b2ContactSim)(unsafe.Pointer(contactSim)).Friction
		(*b2ContactConstraint)(unsafe.Pointer(constraint)).Restitution = (*b2ContactSim)(unsafe.Pointer(contactSim)).Restitution
		(*b2ContactConstraint)(unsafe.Pointer(constraint)).RollingResistance = (*b2ContactSim)(unsafe.Pointer(contactSim)).RollingResistance
		(*b2ContactConstraint)(unsafe.Pointer(constraint)).RollingImpulse = float32(warmStartScale * (*Manifold)(unsafe.Pointer(manifold)).RollingImpulse)
		(*b2ContactConstraint)(unsafe.Pointer(constraint)).TangentSpeed = (*b2ContactSim)(unsafe.Pointer(contactSim)).TangentSpeed
		(*b2ContactConstraint)(unsafe.Pointer(constraint)).PointCount = pointCount
		vA = b2Vec2_zero
		wA = float32FromFloat32(0)
		mA = (*b2ContactSim)(unsafe.Pointer(contactSim)).InvMassA
		iA = (*b2ContactSim)(unsafe.Pointer(contactSim)).InvIA
		if indexA != -int32(1) {
			stateA = awakeStates + uintptr(indexA)*32
			vA = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
			wA = (*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity
		}
		vB = b2Vec2_zero
		wB = float32FromFloat32(0)
		mB = (*b2ContactSim)(unsafe.Pointer(contactSim)).InvMassB
		iB = (*b2ContactSim)(unsafe.Pointer(contactSim)).InvIB
		if indexB != -int32(1) {
			stateB = awakeStates + uintptr(indexB)*32
			vB = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
			wB = (*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity
		}
		if indexA == -int32(1) || indexB == -int32(1) {
			(*b2ContactConstraint)(unsafe.Pointer(constraint)).Softness = staticSoftness
		} else {
			(*b2ContactConstraint)(unsafe.Pointer(constraint)).Softness = contactSoftness
		}
		// copy mass into constraint to avoid cache misses during sub-stepping
		(*b2ContactConstraint)(unsafe.Pointer(constraint)).InvMassA = mA
		(*b2ContactConstraint)(unsafe.Pointer(constraint)).InvIA = iA
		(*b2ContactConstraint)(unsafe.Pointer(constraint)).InvMassB = mB
		(*b2ContactConstraint)(unsafe.Pointer(constraint)).InvIB = iB
		k = iA + iB
		if k > float32FromFloat32(0) {
			v1 = float32FromFloat32(1) / k
		} else {
			v1 = float32FromFloat32(0)
		}
		(*b2ContactConstraint)(unsafe.Pointer(constraint)).RollingMass = v1
		normal = (*b2ContactConstraint)(unsafe.Pointer(constraint)).Normal
		v4 = (*b2ContactConstraint)(unsafe.Pointer(constraint)).Normal
		v5 = Vec2{
			X: v4.Y,
			Y: -v4.X,
		}
		goto _6
	_6:
		tangent = v5
		j = 0
		for {
			if !(j < pointCount) {
				break
			}
			mp = manifold + 12 + uintptr(j)*48
			cp = constraint + 8 + uintptr(j)*44
			(*b2ContactConstraintPoint)(unsafe.Pointer(cp)).NormalImpulse = float32(warmStartScale * (*ManifoldPoint)(unsafe.Pointer(mp)).NormalImpulse)
			(*b2ContactConstraintPoint)(unsafe.Pointer(cp)).TangentImpulse = float32(warmStartScale * (*ManifoldPoint)(unsafe.Pointer(mp)).TangentImpulse)
			(*b2ContactConstraintPoint)(unsafe.Pointer(cp)).TotalNormalImpulse = float32FromFloat32(0)
			rA = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
			rB = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorB
			(*b2ContactConstraintPoint)(unsafe.Pointer(cp)).AnchorA = rA
			(*b2ContactConstraintPoint)(unsafe.Pointer(cp)).AnchorB = rB
			v4 = rB
			v5 = rA
			v8 = Vec2{
				X: v4.X - v5.X,
				Y: v4.Y - v5.Y,
			}
			goto _11
		_11:
			v9 = v8
			v10 = normal
			v1 = float32(v9.X*v10.X) + float32(v9.Y*v10.Y)
			goto _15
		_15:
			(*b2ContactConstraintPoint)(unsafe.Pointer(cp)).BaseSeparation = (*ManifoldPoint)(unsafe.Pointer(mp)).Separation - v1
			v4 = rA
			v5 = normal
			v1 = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
			goto _19
		_19:
			rnA = v1
			v8 = rB
			v9 = normal
			v3 = float32(v8.X*v9.Y) - float32(v8.Y*v9.X)
			goto _23
		_23:
			rnB = v3
			kNormal = mA + mB + float32(float32(iA*rnA)*rnA) + float32(float32(iB*rnB)*rnB)
			if kNormal > float32FromFloat32(0) {
				v14 = float32FromFloat32(1) / kNormal
			} else {
				v14 = float32FromFloat32(0)
			}
			(*b2ContactConstraintPoint)(unsafe.Pointer(cp)).NormalMass = v14
			v4 = rA
			v5 = tangent
			v1 = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
			goto _28
		_28:
			rtA = v1
			v8 = rB
			v9 = tangent
			v3 = float32(v8.X*v9.Y) - float32(v8.Y*v9.X)
			goto _32
		_32:
			rtB = v3
			kTangent = mA + mB + float32(float32(iA*rtA)*rtA) + float32(float32(iB*rtB)*rtB)
			if kTangent > float32FromFloat32(0) {
				v14 = float32FromFloat32(1) / kTangent
			} else {
				v14 = float32FromFloat32(0)
			}
			(*b2ContactConstraintPoint)(unsafe.Pointer(cp)).TangentMass = v14
			v1 = wA
			v4 = rA
			v5 = Vec2{
				X: float32(-v1 * v4.Y),
				Y: float32(v1 * v4.X),
			}
			goto _37
		_37:
			v8 = vA
			v9 = v5
			v10 = Vec2{
				X: v8.X + v9.X,
				Y: v8.Y + v9.Y,
			}
			goto _41
		_41:
			// Save relative velocity for restitution
			vrA = v10
			v3 = wB
			v12 = rB
			v13 = Vec2{
				X: float32(-v3 * v12.Y),
				Y: float32(v3 * v12.X),
			}
			goto _45
		_45:
			v16 = vB
			v17 = v13
			v20 = Vec2{
				X: v16.X + v17.X,
				Y: v16.Y + v17.Y,
			}
			goto _49
		_49:
			vrB = v20
			v21 = vrB
			v25 = vrA
			v26 = Vec2{
				X: v21.X - v25.X,
				Y: v21.Y - v25.Y,
			}
			goto _53
		_53:
			v29 = normal
			v30 = v26
			v14 = float32(v29.X*v30.X) + float32(v29.Y*v30.Y)
			goto _57
		_57:
			(*b2ContactConstraintPoint)(unsafe.Pointer(cp)).RelativeVelocity = v14
			goto _7
		_7:
			;
			j = j + 1
		}
		goto _2
	_2:
		;
		i = i + 1
	}
}

func b2WarmStartOverflowContacts(tls *_Stack, context uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var P, normal, rA, rB, tangent, vA, vB, v13, v14, v17, v18, v20, v21, v22, v24, v25, v8, v9 Vec2
	var awakeSet, color, constraint, constraints, cp, graph, stateA, stateB, states, world, v1, v3, v6, v7 uintptr
	var contactCount, i, indexA, indexB, j, pointCount, v2 int32
	var iA, iB, mA, mB, wA, wB, v12, v16, v26 float32
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = P, awakeSet, color, constraint, constraints, contactCount, cp, graph, i, iA, iB, indexA, indexB, j, mA, mB, normal, pointCount, rA, rB, stateA, stateB, states, tangent, vA, vB, wA, wB, world, v1, v12, v13, v14, v16, v17, v18, v2, v20, v21, v22, v24, v25, v26, v3, v6, v7, v8, v9
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
	// This is a dummy state to represent a static body because static bodies don't have a solver body.
	*(*b2BodyState)(unsafe.Add(nil, bp)) = b2_identityBodyState4
	i = 0
	for {
		if !(i < contactCount) {
			break
		}
		constraint = constraints + uintptr(i)*160
		indexA = (*b2ContactConstraint)(unsafe.Pointer(constraint)).IndexA
		indexB = (*b2ContactConstraint)(unsafe.Pointer(constraint)).IndexB
		if indexA == -int32(1) {
			v6 = bp
		} else {
			v6 = states + uintptr(indexA)*32
		}
		stateA = v6
		if indexB == -int32(1) {
			v7 = bp
		} else {
			v7 = states + uintptr(indexB)*32
		}
		stateB = v7
		vA = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
		wA = (*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity
		vB = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
		wB = (*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity
		mA = (*b2ContactConstraint)(unsafe.Pointer(constraint)).InvMassA
		iA = (*b2ContactConstraint)(unsafe.Pointer(constraint)).InvIA
		mB = (*b2ContactConstraint)(unsafe.Pointer(constraint)).InvMassB
		iB = (*b2ContactConstraint)(unsafe.Pointer(constraint)).InvIB
		// Stiffer for static contacts to avoid bodies getting pushed through the ground
		normal = (*b2ContactConstraint)(unsafe.Pointer(constraint)).Normal
		v8 = (*b2ContactConstraint)(unsafe.Pointer(constraint)).Normal
		v9 = Vec2{
			X: v8.Y,
			Y: -v8.X,
		}
		goto _10
	_10:
		tangent = v9
		pointCount = (*b2ContactConstraint)(unsafe.Pointer(constraint)).PointCount
		j = 0
		for {
			if !(j < pointCount) {
				break
			}
			cp = constraint + 8 + uintptr(j)*44
			// fixed anchors
			rA = (*b2ContactConstraintPoint)(unsafe.Pointer(cp)).AnchorA
			rB = (*b2ContactConstraintPoint)(unsafe.Pointer(cp)).AnchorB
			v12 = (*b2ContactConstraintPoint)(unsafe.Pointer(cp)).NormalImpulse
			v13 = normal
			v14 = Vec2{
				X: float32(v12 * v13.X),
				Y: float32(v12 * v13.Y),
			}
			goto _15
		_15:
			v16 = (*b2ContactConstraintPoint)(unsafe.Pointer(cp)).TangentImpulse
			v17 = tangent
			v18 = Vec2{
				X: float32(v16 * v17.X),
				Y: float32(v16 * v17.Y),
			}
			goto _19
		_19:
			v20 = v14
			v21 = v18
			v22 = Vec2{
				X: v20.X + v21.X,
				Y: v20.Y + v21.Y,
			}
			goto _23
		_23:
			P = v22
			v24 = rA
			v25 = P
			v26 = float32(v24.X*v25.Y) - float32(v24.Y*v25.X)
			goto _27
		_27:
			wA = wA - float32(iA*v26)
			v8 = vA
			v12 = -mA
			v9 = P
			v13 = Vec2{
				X: v8.X + float32(v12*v9.X),
				Y: v8.Y + float32(v12*v9.Y),
			}
			goto _32
		_32:
			vA = v13
			v8 = rB
			v9 = P
			v12 = float32(v8.X*v9.Y) - float32(v8.Y*v9.X)
			goto _36
		_36:
			wB = wB + float32(iB*v12)
			v8 = vB
			v12 = mB
			v9 = P
			v13 = Vec2{
				X: v8.X + float32(v12*v9.X),
				Y: v8.Y + float32(v12*v9.Y),
			}
			goto _41
		_41:
			vB = v13
			goto _11
		_11:
			;
			j = j + 1
		}
		wA = wA - float32(iA*(*b2ContactConstraint)(unsafe.Pointer(constraint)).RollingImpulse)
		wB = wB + float32(iB*(*b2ContactConstraint)(unsafe.Pointer(constraint)).RollingImpulse)
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

func b2ApplyOverflowRestitution(tls *_Stack, context uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var P, normal, rA, rB, vA, vB, vrA, vrB, v10, v11, v13, v14, v15, v18, v19, v21, v22, v23, v25, v26, v27, v29, v30 Vec2
	var awakeSet, color, constraint, constraints, cp, graph, stateA, stateB, states, world, v1, v3 uintptr
	var contactCount, i, j, pointCount, v2 int32
	var iA, iB, impulse, mA, mB, newImpulse, restitution, threshold, vn, wA, wB, v17, v31, v33, v34, v35, v37, v9 float32
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = P, awakeSet, color, constraint, constraints, contactCount, cp, graph, i, iA, iB, impulse, j, mA, mB, newImpulse, normal, pointCount, rA, rB, restitution, stateA, stateB, states, threshold, vA, vB, vn, vrA, vrB, wA, wB, world, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v29, v3, v30, v31, v33, v34, v35, v37, v9
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
	threshold = (*b2World)(unsafe.Pointer((*b2StepContext)(unsafe.Pointer(context)).World)).RestitutionThreshold
	// dummy state to represent a static body
	*(*b2BodyState)(unsafe.Add(nil, bp)) = b2_identityBodyState4
	i = 0
	for {
		if !(i < contactCount) {
			break
		}
		constraint = constraints + uintptr(i)*160
		restitution = (*b2ContactConstraint)(unsafe.Pointer(constraint)).Restitution
		if restitution == float32FromFloat32(0) {
			goto _5
		}
		mA = (*b2ContactConstraint)(unsafe.Pointer(constraint)).InvMassA
		iA = (*b2ContactConstraint)(unsafe.Pointer(constraint)).InvIA
		mB = (*b2ContactConstraint)(unsafe.Pointer(constraint)).InvMassB
		iB = (*b2ContactConstraint)(unsafe.Pointer(constraint)).InvIB
		if (*b2ContactConstraint)(unsafe.Pointer(constraint)).IndexA == -int32(1) {
			v1 = bp
		} else {
			v1 = states + uintptr((*b2ContactConstraint)(unsafe.Pointer(constraint)).IndexA)*32
		}
		stateA = v1
		vA = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
		wA = (*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity
		if (*b2ContactConstraint)(unsafe.Pointer(constraint)).IndexB == -int32(1) {
			v3 = bp
		} else {
			v3 = states + uintptr((*b2ContactConstraint)(unsafe.Pointer(constraint)).IndexB)*32
		}
		stateB = v3
		vB = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
		wB = (*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity
		normal = (*b2ContactConstraint)(unsafe.Pointer(constraint)).Normal
		pointCount = (*b2ContactConstraint)(unsafe.Pointer(constraint)).PointCount
		// it is possible to get more accurate restitution by iterating
		// this only makes a difference if there are two contact points
		// for (int iter = 0; iter < 10; ++iter)
		j = 0
		for {
			if !(j < pointCount) {
				break
			}
			cp = constraint + 8 + uintptr(j)*44
			// if the normal impulse is zero then there was no collision
			// this skips speculative contact points that didn't generate an impulse
			// The max normal impulse is used in case there was a collision that moved away within the sub-step process
			if (*b2ContactConstraintPoint)(unsafe.Pointer(cp)).RelativeVelocity > -threshold || (*b2ContactConstraintPoint)(unsafe.Pointer(cp)).TotalNormalImpulse == float32FromFloat32(0) {
				goto _8
			}
			// fixed anchor points
			rA = (*b2ContactConstraintPoint)(unsafe.Pointer(cp)).AnchorA
			rB = (*b2ContactConstraintPoint)(unsafe.Pointer(cp)).AnchorB
			v9 = wB
			v10 = rB
			v11 = Vec2{
				X: float32(-v9 * v10.Y),
				Y: float32(v9 * v10.X),
			}
			goto _12
		_12:
			v13 = vB
			v14 = v11
			v15 = Vec2{
				X: v13.X + v14.X,
				Y: v13.Y + v14.Y,
			}
			goto _16
		_16:
			// relative normal velocity at contact
			vrB = v15
			v17 = wA
			v18 = rA
			v19 = Vec2{
				X: float32(-v17 * v18.Y),
				Y: float32(v17 * v18.X),
			}
			goto _20
		_20:
			v21 = vA
			v22 = v19
			v23 = Vec2{
				X: v21.X + v22.X,
				Y: v21.Y + v22.Y,
			}
			goto _24
		_24:
			vrA = v23
			v25 = vrB
			v26 = vrA
			v27 = Vec2{
				X: v25.X - v26.X,
				Y: v25.Y - v26.Y,
			}
			goto _28
		_28:
			v29 = v27
			v30 = normal
			v31 = float32(v29.X*v30.X) + float32(v29.Y*v30.Y)
			goto _32
		_32:
			vn = v31
			// compute normal impulse
			impulse = float32(-(*b2ContactConstraintPoint)(unsafe.Pointer(cp)).NormalMass * (vn + float32(restitution*(*b2ContactConstraintPoint)(unsafe.Pointer(cp)).RelativeVelocity)))
			v33 = (*b2ContactConstraintPoint)(unsafe.Pointer(cp)).NormalImpulse + impulse
			v34 = float32FromFloat32(0)
			if v33 > v34 {
				v37 = v33
			} else {
				v37 = v34
			}
			v35 = v37
			goto _36
		_36:
			// clamp the accumulated impulse
			// todo should this be stored?
			newImpulse = v35
			impulse = newImpulse - (*b2ContactConstraintPoint)(unsafe.Pointer(cp)).NormalImpulse
			(*b2ContactConstraintPoint)(unsafe.Pointer(cp)).NormalImpulse = newImpulse
			// Add the incremental impulse rather than the full impulse because this is not a sub-step
			*(*float32)(unsafe.Add(nil, cp+32)) += impulse
			v9 = impulse
			v10 = normal
			v11 = Vec2{
				X: float32(v9 * v10.X),
				Y: float32(v9 * v10.Y),
			}
			goto _41
		_41:
			// apply contact impulse
			P = v11
			v13 = vA
			v17 = mA
			v14 = P
			v15 = Vec2{
				X: v13.X - float32(v17*v14.X),
				Y: v13.Y - float32(v17*v14.Y),
			}
			goto _46
		_46:
			vA = v15
			v10 = rA
			v11 = P
			v9 = float32(v10.X*v11.Y) - float32(v10.Y*v11.X)
			goto _50
		_50:
			wA = wA - float32(iA*v9)
			v10 = vB
			v9 = mB
			v11 = P
			v13 = Vec2{
				X: v10.X + float32(v9*v11.X),
				Y: v10.Y + float32(v9*v11.Y),
			}
			goto _55
		_55:
			vB = v13
			v10 = rB
			v11 = P
			v9 = float32(v10.X*v11.Y) - float32(v10.Y*v11.X)
			goto _59
		_59:
			wB = wB + float32(iB*v9)
			goto _8
		_8:
			;
			j = j + 1
		}
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

func b2StoreOverflowImpulses(tls *_Stack, context uintptr) {
	var color, constraint, constraints, contact, contacts, graph, manifold uintptr
	var contactCount, i, j, pointCount int32
	_, _, _, _, _, _, _, _, _, _, _ = color, constraint, constraints, contact, contactCount, contacts, graph, i, j, manifold, pointCount
	graph = (*b2StepContext)(unsafe.Pointer(context)).Graph
	color = graph + uintptr(int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1))*56
	constraints = *(*uintptr)(unsafe.Add(unsafe.Pointer(color), 48))
	contacts = (*b2GraphColor)(unsafe.Pointer(color)).ContactSims.Data
	contactCount = (*b2GraphColor)(unsafe.Pointer(color)).ContactSims.Count
	i = 0
	for {
		if !(i < contactCount) {
			break
		}
		constraint = constraints + uintptr(i)*160
		contact = contacts + uintptr(i)*176
		manifold = contact + 36
		pointCount = (*Manifold)(unsafe.Pointer(manifold)).PointCount
		j = 0
		for {
			if !(j < pointCount) {
				break
			}
			(*(*ManifoldPoint)(unsafe.Add(nil, manifold+12+uintptr(j)*48))).NormalImpulse = (*(*b2ContactConstraintPoint)(unsafe.Add(nil, constraint+8+uintptr(j)*44))).NormalImpulse
			(*(*ManifoldPoint)(unsafe.Add(nil, manifold+12+uintptr(j)*48))).TangentImpulse = (*(*b2ContactConstraintPoint)(unsafe.Add(nil, constraint+8+uintptr(j)*44))).TangentImpulse
			(*(*ManifoldPoint)(unsafe.Add(nil, manifold+12+uintptr(j)*48))).TotalNormalImpulse = (*(*b2ContactConstraintPoint)(unsafe.Add(nil, constraint+8+uintptr(j)*44))).TotalNormalImpulse
			(*(*ManifoldPoint)(unsafe.Add(nil, manifold+12+uintptr(j)*48))).NormalVelocity = (*(*b2ContactConstraintPoint)(unsafe.Add(nil, constraint+8+uintptr(j)*44))).RelativeVelocity
			goto _2
		_2:
			;
			j = j + 1
		}
		(*Manifold)(unsafe.Pointer(manifold)).RollingImpulse = (*b2ContactConstraint)(unsafe.Pointer(constraint)).RollingImpulse
		goto _1
	_1:
		;
		i = i + 1
	}
}

// C documentation
//
//	// scalar math
type b2FloatW struct {
	X float32
	Y float32
	Z float32
	W float32
}

// C documentation
//
//	// Wide vec2
type b2Vec2W struct {
	X b2FloatW
	Y b2FloatW
}

// C documentation
//
//	// Wide rotation
type b2RotW struct {
	C b2FloatW
	S b2FloatW
}

func b2ZeroW(tls *_Stack) (r b2FloatW) {
	return b2FloatW{}
}

func b2SplatW(tls *_Stack, scalar float32) (r b2FloatW) {
	return b2FloatW{
		X: scalar,
		Y: scalar,
		Z: scalar,
		W: scalar,
	}
}

func b2AddW(tls *_Stack, a b2FloatW, b b2FloatW) (r b2FloatW) {
	return b2FloatW{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
		W: a.W + b.W,
	}
}

func b2SubW(tls *_Stack, a b2FloatW, b b2FloatW) (r b2FloatW) {
	return b2FloatW{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
		W: a.W - b.W,
	}
}

func b2MulW(tls *_Stack, a b2FloatW, b b2FloatW) (r b2FloatW) {
	return b2FloatW{
		X: float32(a.X * b.X),
		Y: float32(a.Y * b.Y),
		Z: float32(a.Z * b.Z),
		W: float32(a.W * b.W),
	}
}

func b2MulAddW(tls *_Stack, a b2FloatW, b b2FloatW, c b2FloatW) (r b2FloatW) {
	return b2FloatW{
		X: a.X + float32(b.X*c.X),
		Y: a.Y + float32(b.Y*c.Y),
		Z: a.Z + float32(b.Z*c.Z),
		W: a.W + float32(b.W*c.W),
	}
}

func b2MulSubW(tls *_Stack, a b2FloatW, b b2FloatW, c b2FloatW) (r b2FloatW) {
	return b2FloatW{
		X: a.X - float32(b.X*c.X),
		Y: a.Y - float32(b.Y*c.Y),
		Z: a.Z - float32(b.Z*c.Z),
		W: a.W - float32(b.W*c.W),
	}
}

func b2MinW(tls *_Stack, a b2FloatW, b b2FloatW) (r1 b2FloatW) {
	var r b2FloatW
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

func b2MaxW(tls *_Stack, a b2FloatW, b b2FloatW) (r1 b2FloatW) {
	var r b2FloatW
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
func b2SymClampW(tls *_Stack, a1 b2FloatW, b b2FloatW) (r1 b2FloatW) {
	var r b2FloatW
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

func b2OrW(tls *_Stack, a b2FloatW, b b2FloatW) (r1 b2FloatW) {
	var r b2FloatW
	var v1 float32
	_, _ = r, v1
	if a.X != float32FromFloat32(0) || b.X != float32FromFloat32(0) {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.X = v1
	if a.Y != float32FromFloat32(0) || b.Y != float32FromFloat32(0) {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Y = v1
	if a.Z != float32FromFloat32(0) || b.Z != float32FromFloat32(0) {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Z = v1
	if a.W != float32FromFloat32(0) || b.W != float32FromFloat32(0) {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.W = v1
	return r
}

func b2GreaterThanW(tls *_Stack, a b2FloatW, b b2FloatW) (r1 b2FloatW) {
	var r b2FloatW
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

func b2EqualsW(tls *_Stack, a b2FloatW, b b2FloatW) (r1 b2FloatW) {
	var r b2FloatW
	var v1 float32
	_, _ = r, v1
	if a.X == b.X {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.X = v1
	if a.Y == b.Y {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Y = v1
	if a.Z == b.Z {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.Z = v1
	if a.W == b.W {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	r.W = v1
	return r
}

func b2AllZeroW(tls *_Stack, a b2FloatW) (r uint8) {
	return boolUint8(a.X == float32FromFloat32(0) && a.Y == float32FromFloat32(0) && a.Z == float32FromFloat32(0) && a.W == float32FromFloat32(0))
}

// C documentation
//
//	// component-wise returns mask ? b : a
func b2BlendW(tls *_Stack, a b2FloatW, b b2FloatW, mask b2FloatW) (r1 b2FloatW) {
	var r b2FloatW
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

func b2DotW(tls *_Stack, a b2Vec2W, b b2Vec2W) (r b2FloatW) {
	return b2AddW(tls, b2MulW(tls, a.X, b.X), b2MulW(tls, a.Y, b.Y))
}

func b2CrossW(tls *_Stack, a b2Vec2W, b b2Vec2W) (r b2FloatW) {
	return b2SubW(tls, b2MulW(tls, a.X, b.Y), b2MulW(tls, a.Y, b.X))
}

func b2RotateVectorW(tls *_Stack, q b2RotW, v b2Vec2W) (r b2Vec2W) {
	return b2Vec2W{
		X: b2SubW(tls, b2MulW(tls, q.C, v.X), b2MulW(tls, q.S, v.Y)),
		Y: b2AddW(tls, b2MulW(tls, q.S, v.X), b2MulW(tls, q.C, v.Y)),
	}
}

func b2GetContactConstraintSIMDByteCount(tls *_Stack) (r int32) {
	return int32(624)
}

// C documentation
//
//	// wide version of b2BodyState
type b2BodyStateW struct {
	V     b2Vec2W
	W     b2FloatW
	Flags b2FloatW
	Dp    b2Vec2W
	Dq    b2RotW
}

// Custom gather/scatter for each SIMD type

// C documentation
//
//	// This is a load and transpose
func b2GatherBodies(tls *_Stack, states uintptr, indices uintptr) (r b2BodyStateW) {
	var identity, s1, s2, s3, s4, v1, v2, v3, v4 b2BodyState
	var simdBody b2BodyStateW
	_, _, _, _, _, _, _, _, _, _ = identity, s1, s2, s3, s4, simdBody, v1, v2, v3, v4
	identity = b2_identityBodyState4
	if *(*int32)(unsafe.Add(nil, indices)) == -int32(1) {
		v1 = identity
	} else {
		v1 = *(*b2BodyState)(unsafe.Add(nil, states+uintptr(*(*int32)(unsafe.Add(nil, indices)))*32))
	}
	s1 = v1
	if *(*int32)(unsafe.Add(nil, indices+1*4)) == -int32(1) {
		v2 = identity
	} else {
		v2 = *(*b2BodyState)(unsafe.Add(nil, states+uintptr(*(*int32)(unsafe.Add(nil, indices+1*4)))*32))
	}
	s2 = v2
	if *(*int32)(unsafe.Add(nil, indices+2*4)) == -int32(1) {
		v3 = identity
	} else {
		v3 = *(*b2BodyState)(unsafe.Add(nil, states+uintptr(*(*int32)(unsafe.Add(nil, indices+2*4)))*32))
	}
	s3 = v3
	if *(*int32)(unsafe.Add(nil, indices+3*4)) == -int32(1) {
		v4 = identity
	} else {
		v4 = *(*b2BodyState)(unsafe.Add(nil, states+uintptr(*(*int32)(unsafe.Add(nil, indices+3*4)))*32))
	}
	s4 = v4
	simdBody.V.X = b2FloatW{
		X: s1.LinearVelocity.X,
		Y: s2.LinearVelocity.X,
		Z: s3.LinearVelocity.X,
		W: s4.LinearVelocity.X,
	}
	simdBody.V.Y = b2FloatW{
		X: s1.LinearVelocity.Y,
		Y: s2.LinearVelocity.Y,
		Z: s3.LinearVelocity.Y,
		W: s4.LinearVelocity.Y,
	}
	simdBody.W = b2FloatW{
		X: s1.AngularVelocity,
		Y: s2.AngularVelocity,
		Z: s3.AngularVelocity,
		W: s4.AngularVelocity,
	}
	simdBody.Flags = b2FloatW{
		X: float32(s1.Flags),
		Y: float32(s2.Flags),
		Z: float32(s3.Flags),
		W: float32(s4.Flags),
	}
	simdBody.Dp.X = b2FloatW{
		X: s1.DeltaPosition.X,
		Y: s2.DeltaPosition.X,
		Z: s3.DeltaPosition.X,
		W: s4.DeltaPosition.X,
	}
	simdBody.Dp.Y = b2FloatW{
		X: s1.DeltaPosition.Y,
		Y: s2.DeltaPosition.Y,
		Z: s3.DeltaPosition.Y,
		W: s4.DeltaPosition.Y,
	}
	simdBody.Dq.C = b2FloatW{
		X: s1.DeltaRotation.C,
		Y: s2.DeltaRotation.C,
		Z: s3.DeltaRotation.C,
		W: s4.DeltaRotation.C,
	}
	simdBody.Dq.S = b2FloatW{
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
	var state, state1, state2, state3 uintptr
	_, _, _, _ = state, state1, state2, state3
	// todo somehow skip writing to kinematic bodies
	if *(*int32)(unsafe.Add(nil, indices)) != -int32(1) {
		state = states + uintptr(*(*int32)(unsafe.Add(nil, indices)))*32
		(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity.X = (*b2BodyStateW)(unsafe.Pointer(simdBody)).V.X.X
		(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity.Y = (*b2BodyStateW)(unsafe.Pointer(simdBody)).V.Y.X
		(*b2BodyState)(unsafe.Pointer(state)).AngularVelocity = (*b2BodyStateW)(unsafe.Pointer(simdBody)).W.X
	}
	if *(*int32)(unsafe.Add(nil, indices+1*4)) != -int32(1) {
		state1 = states + uintptr(*(*int32)(unsafe.Add(nil, indices+1*4)))*32
		(*b2BodyState)(unsafe.Pointer(state1)).LinearVelocity.X = (*b2BodyStateW)(unsafe.Pointer(simdBody)).V.X.Y
		(*b2BodyState)(unsafe.Pointer(state1)).LinearVelocity.Y = (*b2BodyStateW)(unsafe.Pointer(simdBody)).V.Y.Y
		(*b2BodyState)(unsafe.Pointer(state1)).AngularVelocity = (*b2BodyStateW)(unsafe.Pointer(simdBody)).W.Y
	}
	if *(*int32)(unsafe.Add(nil, indices+2*4)) != -int32(1) {
		state2 = states + uintptr(*(*int32)(unsafe.Add(nil, indices+2*4)))*32
		(*b2BodyState)(unsafe.Pointer(state2)).LinearVelocity.X = (*b2BodyStateW)(unsafe.Pointer(simdBody)).V.X.Z
		(*b2BodyState)(unsafe.Pointer(state2)).LinearVelocity.Y = (*b2BodyStateW)(unsafe.Pointer(simdBody)).V.Y.Z
		(*b2BodyState)(unsafe.Pointer(state2)).AngularVelocity = (*b2BodyStateW)(unsafe.Pointer(simdBody)).W.Z
	}
	if *(*int32)(unsafe.Add(nil, indices+3*4)) != -int32(1) {
		state3 = states + uintptr(*(*int32)(unsafe.Add(nil, indices+3*4)))*32
		(*b2BodyState)(unsafe.Pointer(state3)).LinearVelocity.X = (*b2BodyStateW)(unsafe.Pointer(simdBody)).V.X.W
		(*b2BodyState)(unsafe.Pointer(state3)).LinearVelocity.Y = (*b2BodyStateW)(unsafe.Pointer(simdBody)).V.Y.W
		(*b2BodyState)(unsafe.Pointer(state3)).AngularVelocity = (*b2BodyStateW)(unsafe.Pointer(simdBody)).W.W
	}
}

func b2PrepareContactsTask(tls *_Stack, startIndex int32, endIndex int32, context uintptr) {
	var awakeStates, constraint, constraints, contactSim, contacts, manifold, mp, mp1, stateA, stateB, world uintptr
	var contactSoftness, soft, staticSoftness, v5 b2Softness
	var i, indexA, indexB, j, pointCount int32
	var iA, iB, k, kNormal, kNormal1, kTangent, kTangent1, mA, mB, rnA, rnA1, rnB, rnB1, rtA, rtA1, rtB, rtB1, wA, wB, warmStartScale, v1, v15, v4 float32
	var normal, rA, rA1, rB, rB1, tangent, vA, vB, vrA, vrA1, vrB, vrB1, v10, v11, v13, v14, v17, v18, v21, v22, v26, v27, v30, v31, v6, v7, v9 Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = awakeStates, constraint, constraints, contactSim, contactSoftness, contacts, i, iA, iB, indexA, indexB, j, k, kNormal, kNormal1, kTangent, kTangent1, mA, mB, manifold, mp, mp1, normal, pointCount, rA, rA1, rB, rB1, rnA, rnA1, rnB, rnB1, rtA, rtA1, rtB, rtB1, soft, stateA, stateB, staticSoftness, tangent, vA, vB, vrA, vrA1, vrB, vrB1, wA, wB, warmStartScale, world, v1, v10, v11, v13, v14, v15, v17, v18, v21, v22, v26, v27, v30, v31, v4, v5, v6, v7, v9
	world = (*b2StepContext)(unsafe.Pointer(context)).World
	contacts = (*b2StepContext)(unsafe.Pointer(context)).Contacts
	constraints = (*b2StepContext)(unsafe.Pointer(context)).SimdContactConstraints
	awakeStates = (*b2StepContext)(unsafe.Pointer(context)).States
	// Stiffer for static contacts to avoid bodies getting pushed through the ground
	contactSoftness = (*b2StepContext)(unsafe.Pointer(context)).ContactSoftness
	staticSoftness = (*b2StepContext)(unsafe.Pointer(context)).StaticSoftness
	if (*b2World)(unsafe.Pointer(world)).EnableWarmStarting != 0 {
		v1 = float32FromFloat32(1)
	} else {
		v1 = float32FromFloat32(0)
	}
	warmStartScale = v1
	i = startIndex
	for {
		if !(i < endIndex) {
			break
		}
		constraint = constraints + uintptr(i)*624
		j = 0
		for {
			if !(j < int32(_B2_SIMD_WIDTH)) {
				break
			}
			contactSim = *(*uintptr)(unsafe.Add(nil, contacts+uintptr(int32(_B2_SIMD_WIDTH)*i+j)*8))
			if contactSim != uintptrFromInt32(0) {
				manifold = contactSim + 36
				indexA = (*b2ContactSim)(unsafe.Pointer(contactSim)).BodySimIndexA
				indexB = (*b2ContactSim)(unsafe.Pointer(contactSim)).BodySimIndexB
				*(*int32)(unsafe.Add(nil, constraint+uintptr(j)*4)) = indexA
				*(*int32)(unsafe.Add(nil, constraint+16+uintptr(j)*4)) = indexB
				vA = b2Vec2_zero
				wA = float32FromFloat32(0)
				mA = (*b2ContactSim)(unsafe.Pointer(contactSim)).InvMassA
				iA = (*b2ContactSim)(unsafe.Pointer(contactSim)).InvIA
				if indexA != -int32(1) {
					stateA = awakeStates + uintptr(indexA)*32
					vA = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
					wA = (*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity
				}
				vB = b2Vec2_zero
				wB = float32FromFloat32(0)
				mB = (*b2ContactSim)(unsafe.Pointer(contactSim)).InvMassB
				iB = (*b2ContactSim)(unsafe.Pointer(contactSim)).InvIB
				if indexB != -int32(1) {
					stateB = awakeStates + uintptr(indexB)*32
					vB = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
					wB = (*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity
				}
				*(*float32)(unsafe.Add(nil, constraint+32+uintptr(j)*4)) = mA
				*(*float32)(unsafe.Add(nil, constraint+48+uintptr(j)*4)) = mB
				*(*float32)(unsafe.Add(nil, constraint+64+uintptr(j)*4)) = iA
				*(*float32)(unsafe.Add(nil, constraint+80+uintptr(j)*4)) = iB
				k = iA + iB
				if k > float32FromFloat32(0) {
					v1 = float32FromFloat32(1) / k
				} else {
					v1 = float32FromFloat32(0)
				}
				*(*float32)(unsafe.Add(nil, constraint+176+uintptr(j)*4)) = v1
				if indexA == -int32(1) || indexB == -int32(1) {
					v5 = staticSoftness
				} else {
					v5 = contactSoftness
				}
				soft = v5
				normal = (*Manifold)(unsafe.Pointer(manifold)).Normal
				*(*float32)(unsafe.Add(nil, constraint+96+uintptr(j)*4)) = normal.X
				*(*float32)(unsafe.Add(nil, constraint+96+16+uintptr(j)*4)) = normal.Y
				*(*float32)(unsafe.Add(nil, constraint+128+uintptr(j)*4)) = (*b2ContactSim)(unsafe.Pointer(contactSim)).Friction
				*(*float32)(unsafe.Add(nil, constraint+144+uintptr(j)*4)) = (*b2ContactSim)(unsafe.Pointer(contactSim)).TangentSpeed
				*(*float32)(unsafe.Add(nil, constraint+576+uintptr(j)*4)) = (*b2ContactSim)(unsafe.Pointer(contactSim)).Restitution
				*(*float32)(unsafe.Add(nil, constraint+160+uintptr(j)*4)) = (*b2ContactSim)(unsafe.Pointer(contactSim)).RollingResistance
				*(*float32)(unsafe.Add(nil, constraint+192+uintptr(j)*4)) = float32(warmStartScale * (*Manifold)(unsafe.Pointer(manifold)).RollingImpulse)
				*(*float32)(unsafe.Add(nil, constraint+208+uintptr(j)*4)) = soft.BiasRate
				*(*float32)(unsafe.Add(nil, constraint+224+uintptr(j)*4)) = soft.MassScale
				*(*float32)(unsafe.Add(nil, constraint+240+uintptr(j)*4)) = soft.ImpulseScale
				v6 = normal
				v7 = Vec2{
					X: v6.Y,
					Y: -v6.X,
				}
				goto _8
			_8:
				tangent = v7
				mp = manifold + 12 + uintptr(0)*48
				rA = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
				rB = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorB
				*(*float32)(unsafe.Add(nil, constraint+256+uintptr(j)*4)) = rA.X
				*(*float32)(unsafe.Add(nil, constraint+256+16+uintptr(j)*4)) = rA.Y
				*(*float32)(unsafe.Add(nil, constraint+288+uintptr(j)*4)) = rB.X
				*(*float32)(unsafe.Add(nil, constraint+288+16+uintptr(j)*4)) = rB.Y
				v6 = rB
				v7 = rA
				v9 = Vec2{
					X: v6.X - v7.X,
					Y: v6.Y - v7.Y,
				}
				goto _12
			_12:
				v10 = v9
				v11 = normal
				v1 = float32(v10.X*v11.X) + float32(v10.Y*v11.Y)
				goto _16
			_16:
				*(*float32)(unsafe.Add(nil, constraint+352+uintptr(j)*4)) = (*ManifoldPoint)(unsafe.Pointer(mp)).Separation - v1
				*(*float32)(unsafe.Add(nil, constraint+368+uintptr(j)*4)) = float32(warmStartScale * (*ManifoldPoint)(unsafe.Pointer(mp)).NormalImpulse)
				*(*float32)(unsafe.Add(nil, constraint+400+uintptr(j)*4)) = float32(warmStartScale * (*ManifoldPoint)(unsafe.Pointer(mp)).TangentImpulse)
				*(*float32)(unsafe.Add(nil, constraint+384+uintptr(j)*4)) = float32FromFloat32(0)
				v6 = rA
				v7 = normal
				v1 = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
				goto _20
			_20:
				rnA = v1
				v9 = rB
				v10 = normal
				v4 = float32(v9.X*v10.Y) - float32(v9.Y*v10.X)
				goto _24
			_24:
				rnB = v4
				kNormal = mA + mB + float32(float32(iA*rnA)*rnA) + float32(float32(iB*rnB)*rnB)
				if kNormal > float32FromFloat32(0) {
					v15 = float32FromFloat32(1) / kNormal
				} else {
					v15 = float32FromFloat32(0)
				}
				*(*float32)(unsafe.Add(nil, constraint+320+uintptr(j)*4)) = v15
				v6 = rA
				v7 = tangent
				v1 = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
				goto _29
			_29:
				rtA = v1
				v9 = rB
				v10 = tangent
				v4 = float32(v9.X*v10.Y) - float32(v9.Y*v10.X)
				goto _33
			_33:
				rtB = v4
				kTangent = mA + mB + float32(float32(iA*rtA)*rtA) + float32(float32(iB*rtB)*rtB)
				if kTangent > float32FromFloat32(0) {
					v15 = float32FromFloat32(1) / kTangent
				} else {
					v15 = float32FromFloat32(0)
				}
				*(*float32)(unsafe.Add(nil, constraint+336+uintptr(j)*4)) = v15
				v1 = wA
				v6 = rA
				v7 = Vec2{
					X: float32(-v1 * v6.Y),
					Y: float32(v1 * v6.X),
				}
				goto _38
			_38:
				v9 = vA
				v10 = v7
				v11 = Vec2{
					X: v9.X + v10.X,
					Y: v9.Y + v10.Y,
				}
				goto _42
			_42:
				// relative velocity for restitution
				vrA = v11
				v4 = wB
				v13 = rB
				v14 = Vec2{
					X: float32(-v4 * v13.Y),
					Y: float32(v4 * v13.X),
				}
				goto _46
			_46:
				v17 = vB
				v18 = v14
				v21 = Vec2{
					X: v17.X + v18.X,
					Y: v17.Y + v18.Y,
				}
				goto _50
			_50:
				vrB = v21
				v22 = vrB
				v26 = vrA
				v27 = Vec2{
					X: v22.X - v26.X,
					Y: v22.Y - v26.Y,
				}
				goto _54
			_54:
				v30 = normal
				v31 = v27
				v15 = float32(v30.X*v31.X) + float32(v30.Y*v31.Y)
				goto _58
			_58:
				*(*float32)(unsafe.Add(nil, constraint+592+uintptr(j)*4)) = v15
				pointCount = (*Manifold)(unsafe.Pointer(manifold)).PointCount
				if !(0 < pointCount && pointCount <= int32(2)) && b2InternalAssertFcn(tls, __ccgo_ts+3826, __ccgo_ts+3860, int32FromInt32(1566)) != 0 {
					__builtin_trap(tls)
				}
				if pointCount == int32(2) {
					mp1 = manifold + 12 + uintptr(1)*48
					rA1 = (*ManifoldPoint)(unsafe.Pointer(mp1)).AnchorA
					rB1 = (*ManifoldPoint)(unsafe.Pointer(mp1)).AnchorB
					*(*float32)(unsafe.Add(nil, constraint+416+uintptr(j)*4)) = rA1.X
					*(*float32)(unsafe.Add(nil, constraint+416+16+uintptr(j)*4)) = rA1.Y
					*(*float32)(unsafe.Add(nil, constraint+448+uintptr(j)*4)) = rB1.X
					*(*float32)(unsafe.Add(nil, constraint+448+16+uintptr(j)*4)) = rB1.Y
					v6 = rB1
					v7 = rA1
					v9 = Vec2{
						X: v6.X - v7.X,
						Y: v6.Y - v7.Y,
					}
					goto _62
				_62:
					v10 = v9
					v11 = normal
					v1 = float32(v10.X*v11.X) + float32(v10.Y*v11.Y)
					goto _66
				_66:
					*(*float32)(unsafe.Add(nil, constraint+480+uintptr(j)*4)) = (*ManifoldPoint)(unsafe.Pointer(mp1)).Separation - v1
					*(*float32)(unsafe.Add(nil, constraint+496+uintptr(j)*4)) = float32(warmStartScale * (*ManifoldPoint)(unsafe.Pointer(mp1)).NormalImpulse)
					*(*float32)(unsafe.Add(nil, constraint+528+uintptr(j)*4)) = float32(warmStartScale * (*ManifoldPoint)(unsafe.Pointer(mp1)).TangentImpulse)
					*(*float32)(unsafe.Add(nil, constraint+512+uintptr(j)*4)) = float32FromFloat32(0)
					v6 = rA1
					v7 = normal
					v1 = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
					goto _70
				_70:
					rnA1 = v1
					v9 = rB1
					v10 = normal
					v4 = float32(v9.X*v10.Y) - float32(v9.Y*v10.X)
					goto _74
				_74:
					rnB1 = v4
					kNormal1 = mA + mB + float32(float32(iA*rnA1)*rnA1) + float32(float32(iB*rnB1)*rnB1)
					if kNormal1 > float32FromFloat32(0) {
						v15 = float32FromFloat32(1) / kNormal1
					} else {
						v15 = float32FromFloat32(0)
					}
					*(*float32)(unsafe.Add(nil, constraint+544+uintptr(j)*4)) = v15
					v6 = rA1
					v7 = tangent
					v1 = float32(v6.X*v7.Y) - float32(v6.Y*v7.X)
					goto _79
				_79:
					rtA1 = v1
					v9 = rB1
					v10 = tangent
					v4 = float32(v9.X*v10.Y) - float32(v9.Y*v10.X)
					goto _83
				_83:
					rtB1 = v4
					kTangent1 = mA + mB + float32(float32(iA*rtA1)*rtA1) + float32(float32(iB*rtB1)*rtB1)
					if kTangent1 > float32FromFloat32(0) {
						v15 = float32FromFloat32(1) / kTangent1
					} else {
						v15 = float32FromFloat32(0)
					}
					*(*float32)(unsafe.Add(nil, constraint+560+uintptr(j)*4)) = v15
					v1 = wA
					v6 = rA1
					v7 = Vec2{
						X: float32(-v1 * v6.Y),
						Y: float32(v1 * v6.X),
					}
					goto _88
				_88:
					v9 = vA
					v10 = v7
					v11 = Vec2{
						X: v9.X + v10.X,
						Y: v9.Y + v10.Y,
					}
					goto _92
				_92:
					// relative velocity for restitution
					vrA1 = v11
					v4 = wB
					v13 = rB1
					v14 = Vec2{
						X: float32(-v4 * v13.Y),
						Y: float32(v4 * v13.X),
					}
					goto _96
				_96:
					v17 = vB
					v18 = v14
					v21 = Vec2{
						X: v17.X + v18.X,
						Y: v17.Y + v18.Y,
					}
					goto _100
				_100:
					vrB1 = v21
					v22 = vrB1
					v26 = vrA1
					v27 = Vec2{
						X: v22.X - v26.X,
						Y: v22.Y - v26.Y,
					}
					goto _104
				_104:
					v30 = normal
					v31 = v27
					v15 = float32(v30.X*v31.X) + float32(v30.Y*v31.Y)
					goto _108
				_108:
					*(*float32)(unsafe.Add(nil, constraint+608+uintptr(j)*4)) = v15
				} else {
					// dummy data that has no effect
					*(*float32)(unsafe.Add(nil, constraint+480+uintptr(j)*4)) = float32FromFloat32(0)
					*(*float32)(unsafe.Add(nil, constraint+496+uintptr(j)*4)) = float32FromFloat32(0)
					*(*float32)(unsafe.Add(nil, constraint+528+uintptr(j)*4)) = float32FromFloat32(0)
					*(*float32)(unsafe.Add(nil, constraint+512+uintptr(j)*4)) = float32FromFloat32(0)
					*(*float32)(unsafe.Add(nil, constraint+416+uintptr(j)*4)) = float32FromFloat32(0)
					*(*float32)(unsafe.Add(nil, constraint+416+16+uintptr(j)*4)) = float32FromFloat32(0)
					*(*float32)(unsafe.Add(nil, constraint+448+uintptr(j)*4)) = float32FromFloat32(0)
					*(*float32)(unsafe.Add(nil, constraint+448+16+uintptr(j)*4)) = float32FromFloat32(0)
					*(*float32)(unsafe.Add(nil, constraint+544+uintptr(j)*4)) = float32FromFloat32(0)
					*(*float32)(unsafe.Add(nil, constraint+560+uintptr(j)*4)) = float32FromFloat32(0)
					*(*float32)(unsafe.Add(nil, constraint+608+uintptr(j)*4)) = float32FromFloat32(0)
				}
			} else {
				// SIMD remainder
				*(*int32)(unsafe.Add(nil, constraint+uintptr(j)*4)) = -int32(1)
				*(*int32)(unsafe.Add(nil, constraint+16+uintptr(j)*4)) = -int32(1)
				*(*float32)(unsafe.Add(nil, constraint+32+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+48+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+64+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+80+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+96+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+96+16+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+128+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+144+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+160+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+176+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+192+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+208+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+224+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+240+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+256+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+256+16+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+288+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+288+16+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+352+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+368+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+400+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+384+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+320+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+336+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+416+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+416+16+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+448+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+448+16+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+480+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+496+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+528+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+512+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+544+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+560+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+576+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+592+uintptr(j)*4)) = float32FromFloat32(0)
				*(*float32)(unsafe.Add(nil, constraint+608+uintptr(j)*4)) = float32FromFloat32(0)
			}
			goto _3
		_3:
			;
			j = j + 1
		}
		goto _2
	_2:
		;
		i = i + 1
	}
}

func b2WarmStartContactsTask(tls *_Stack, startIndex int32, endIndex int32, context uintptr, colorIndex int32) {
	bp := tls.Alloc(256)
	defer tls.Free(256)
	var P, P1, rA, rA1, rB, rB1 b2Vec2W
	var c, constraints, states uintptr
	var i int32
	var tangentX, tangentY b2FloatW
	var _ /* bA at bp+0 */ b2BodyStateW
	var _ /* bB at bp+128 */ b2BodyStateW
	_, _, _, _, _, _, _, _, _, _, _, _ = P, P1, c, constraints, i, rA, rA1, rB, rB1, states, tangentX, tangentY
	states = (*b2StepContext)(unsafe.Pointer(context)).States
	constraints = (*(*b2GraphColor)(unsafe.Add(nil, (*b2StepContext)(unsafe.Pointer(context)).Graph+uintptr(colorIndex)*56))).ｆ__ccgo3_48.SimdConstraints
	i = startIndex
	for {
		if !(i < endIndex) {
			break
		}
		c = constraints + uintptr(i)*624
		*(*b2BodyStateW)(unsafe.Add(nil, bp)) = b2GatherBodies(tls, states, c)
		*(*b2BodyStateW)(unsafe.Add(nil, bp+128)) = b2GatherBodies(tls, states, c+16)
		tangentX = (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Normal.Y
		tangentY = b2SubW(tls, b2ZeroW(tls), (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Normal.X)
		// fixed anchors
		rA = (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).AnchorA1
		rB = (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).AnchorB1
		P.X = b2AddW(tls, b2MulW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).NormalImpulse1, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Normal.X), b2MulW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).TangentImpulse1, tangentX))
		P.Y = b2AddW(tls, b2MulW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).NormalImpulse1, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Normal.Y), b2MulW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).TangentImpulse1, tangentY))
		(*(*b2BodyStateW)(unsafe.Add(nil, bp))).W = b2MulSubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).W, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvIA, b2CrossW(tls, rA, P))
		(*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.X = b2MulSubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.X, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassA, P.X)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.Y = b2MulSubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.Y, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassA, P.Y)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W = b2MulAddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvIB, b2CrossW(tls, rB, P))
		(*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.X = b2MulAddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.X, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassB, P.X)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.Y = b2MulAddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.Y, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassB, P.Y)
		// fixed anchors
		rA1 = (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).AnchorA2
		rB1 = (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).AnchorB2
		P1.X = b2AddW(tls, b2MulW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).NormalImpulse2, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Normal.X), b2MulW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).TangentImpulse2, tangentX))
		P1.Y = b2AddW(tls, b2MulW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).NormalImpulse2, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Normal.Y), b2MulW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).TangentImpulse2, tangentY))
		(*(*b2BodyStateW)(unsafe.Add(nil, bp))).W = b2MulSubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).W, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvIA, b2CrossW(tls, rA1, P1))
		(*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.X = b2MulSubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.X, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassA, P1.X)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.Y = b2MulSubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.Y, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassA, P1.Y)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W = b2MulAddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvIB, b2CrossW(tls, rB1, P1))
		(*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.X = b2MulAddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.X, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassB, P1.X)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.Y = b2MulAddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.Y, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassB, P1.Y)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp))).W = b2MulSubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).W, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvIA, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).RollingImpulse)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W = b2MulAddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvIB, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).RollingImpulse)
		b2ScatterBodies(tls, states, c, bp)
		b2ScatterBodies(tls, states, c+16, bp+128)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b2ApplyRestitutionTask(tls *_Stack, startIndex int32, endIndex int32, context uintptr, colorIndex int32) {
	bp := tls.Alloc(256)
	defer tls.Free(256)
	var Px, Px1, Py, Py1, dvx, dvx1, dvy, dvy1, impulse, impulse1, mask, mask1, mask11, mask12, mask2, mask21, mass, mass1, negImpulse, negImpulse1, newImpulse, newImpulse1, restitutionMask, threshold, vn, vn1, zero b2FloatW
	var c, constraints, states uintptr
	var i int32
	var rA, rA1, rB, rB1 b2Vec2W
	var _ /* bA at bp+0 */ b2BodyStateW
	var _ /* bB at bp+128 */ b2BodyStateW
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = Px, Px1, Py, Py1, c, constraints, dvx, dvx1, dvy, dvy1, i, impulse, impulse1, mask, mask1, mask11, mask12, mask2, mask21, mass, mass1, negImpulse, negImpulse1, newImpulse, newImpulse1, rA, rA1, rB, rB1, restitutionMask, states, threshold, vn, vn1, zero
	states = (*b2StepContext)(unsafe.Pointer(context)).States
	constraints = (*(*b2GraphColor)(unsafe.Add(nil, (*b2StepContext)(unsafe.Pointer(context)).Graph+uintptr(colorIndex)*56))).ｆ__ccgo3_48.SimdConstraints
	threshold = b2SplatW(tls, (*b2World)(unsafe.Pointer((*b2StepContext)(unsafe.Pointer(context)).World)).RestitutionThreshold)
	zero = b2ZeroW(tls)
	i = startIndex
	for {
		if !(i < endIndex) {
			break
		}
		c = constraints + uintptr(i)*624
		if b2AllZeroW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Restitution) != 0 {
			// No lanes have restitution. Common case.
			goto _1
		}
		// Create a mask based on restitution so that lanes with no restitution are not affected
		// by the calculations below.
		restitutionMask = b2EqualsW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Restitution, zero)
		*(*b2BodyStateW)(unsafe.Add(nil, bp)) = b2GatherBodies(tls, states, c)
		*(*b2BodyStateW)(unsafe.Add(nil, bp+128)) = b2GatherBodies(tls, states, c+16)
		// first point non-penetration constraint
		// Set effective mass to zero if restitution should not be applied
		mask11 = b2GreaterThanW(tls, b2AddW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).RelativeVelocity1, threshold), zero)
		mask2 = b2EqualsW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).TotalNormalImpulse1, zero)
		mask = b2OrW(tls, b2OrW(tls, mask11, mask2), restitutionMask)
		mass = b2BlendW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).NormalMass1, zero, mask)
		// fixed anchors for Jacobians
		rA = (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).AnchorA1
		rB = (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).AnchorB1
		// Relative velocity at contact
		dvx = b2SubW(tls, b2SubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.X, b2MulW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W, rB.Y)), b2SubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.X, b2MulW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).W, rA.Y)))
		dvy = b2SubW(tls, b2AddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.Y, b2MulW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W, rB.X)), b2AddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.Y, b2MulW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).W, rA.X)))
		vn = b2AddW(tls, b2MulW(tls, dvx, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Normal.X), b2MulW(tls, dvy, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Normal.Y))
		// Compute normal impulse
		negImpulse = b2MulW(tls, mass, b2AddW(tls, vn, b2MulW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Restitution, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).RelativeVelocity1)))
		// Clamp the accumulated impulse
		newImpulse = b2MaxW(tls, b2SubW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).NormalImpulse1, negImpulse), b2ZeroW(tls))
		impulse = b2SubW(tls, newImpulse, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).NormalImpulse1)
		(*b2ContactConstraintSIMD)(unsafe.Pointer(c)).NormalImpulse1 = newImpulse
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
		// Set effective mass to zero if restitution should not be applied
		mask12 = b2GreaterThanW(tls, b2AddW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).RelativeVelocity2, threshold), zero)
		mask21 = b2EqualsW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).TotalNormalImpulse2, zero)
		mask1 = b2OrW(tls, b2OrW(tls, mask12, mask21), restitutionMask)
		mass1 = b2BlendW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).NormalMass2, zero, mask1)
		// fixed anchors for Jacobians
		rA1 = (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).AnchorA2
		rB1 = (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).AnchorB2
		// Relative velocity at contact
		dvx1 = b2SubW(tls, b2SubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.X, b2MulW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W, rB1.Y)), b2SubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.X, b2MulW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).W, rA1.Y)))
		dvy1 = b2SubW(tls, b2AddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.Y, b2MulW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W, rB1.X)), b2AddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.Y, b2MulW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).W, rA1.X)))
		vn1 = b2AddW(tls, b2MulW(tls, dvx1, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Normal.X), b2MulW(tls, dvy1, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Normal.Y))
		// Compute normal impulse
		negImpulse1 = b2MulW(tls, mass1, b2AddW(tls, vn1, b2MulW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Restitution, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).RelativeVelocity2)))
		// Clamp the accumulated impulse
		newImpulse1 = b2MaxW(tls, b2SubW(tls, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).NormalImpulse2, negImpulse1), b2ZeroW(tls))
		impulse1 = b2SubW(tls, newImpulse1, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).NormalImpulse2)
		(*b2ContactConstraintSIMD)(unsafe.Pointer(c)).NormalImpulse2 = newImpulse1
		// Apply contact impulse
		Px1 = b2MulW(tls, impulse1, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Normal.X)
		Py1 = b2MulW(tls, impulse1, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).Normal.Y)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.X = b2MulSubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.X, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassA, Px1)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.Y = b2MulSubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).V.Y, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassA, Py1)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp))).W = b2MulSubW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp))).W, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvIA, b2SubW(tls, b2MulW(tls, rA1.X, Py1), b2MulW(tls, rA1.Y, Px1)))
		(*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.X = b2MulAddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.X, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassB, Px1)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.Y = b2MulAddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).V.Y, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvMassB, Py1)
		(*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W = b2MulAddW(tls, (*(*b2BodyStateW)(unsafe.Add(nil, bp+128))).W, (*b2ContactConstraintSIMD)(unsafe.Pointer(c)).InvIB, b2SubW(tls, b2MulW(tls, rB1.X, Py1), b2MulW(tls, rB1.Y, Px1)))
		b2ScatterBodies(tls, states, c, bp)
		b2ScatterBodies(tls, states, c+16, bp+128)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b2StoreImpulsesTask(tls *_Stack, startIndex int32, endIndex int32, context uintptr) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var baseIndex, constraintIndex, laneIndex int32
	var c, constraints, contacts, m, normalImpulse1, normalImpulse2, normalVelocity1, normalVelocity2, rollingImpulse, tangentImpulse1, tangentImpulse2, totalNormalImpulse1, totalNormalImpulse2, v3 uintptr
	var _ /* dummy at bp+0 */ Manifold
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = baseIndex, c, constraintIndex, constraints, contacts, laneIndex, m, normalImpulse1, normalImpulse2, normalVelocity1, normalVelocity2, rollingImpulse, tangentImpulse1, tangentImpulse2, totalNormalImpulse1, totalNormalImpulse2, v3
	contacts = (*b2StepContext)(unsafe.Pointer(context)).Contacts
	constraints = (*b2StepContext)(unsafe.Pointer(context)).SimdContactConstraints
	*(*Manifold)(unsafe.Add(nil, bp)) = Manifold{}
	constraintIndex = startIndex
	for {
		if !(constraintIndex < endIndex) {
			break
		}
		c = constraints + uintptr(constraintIndex)*624
		rollingImpulse = c + 192
		normalImpulse1 = c + 368
		normalImpulse2 = c + 496
		tangentImpulse1 = c + 400
		tangentImpulse2 = c + 528
		totalNormalImpulse1 = c + 384
		totalNormalImpulse2 = c + 512
		normalVelocity1 = c + 592
		normalVelocity2 = c + 608
		baseIndex = int32(_B2_SIMD_WIDTH) * constraintIndex
		laneIndex = 0
		for {
			if !(laneIndex < int32(_B2_SIMD_WIDTH)) {
				break
			}
			if *(*uintptr)(unsafe.Add(nil, contacts+uintptr(baseIndex+laneIndex)*8)) == uintptrFromInt32(0) {
				v3 = bp
			} else {
				v3 = *(*uintptr)(unsafe.Add(nil, contacts+uintptr(baseIndex+laneIndex)*8)) + 36
			}
			m = v3
			(*Manifold)(unsafe.Pointer(m)).RollingImpulse = *(*float32)(unsafe.Add(nil, rollingImpulse+uintptr(laneIndex)*4))
			(*(*ManifoldPoint)(unsafe.Add(nil, m+12))).NormalImpulse = *(*float32)(unsafe.Add(nil, normalImpulse1+uintptr(laneIndex)*4))
			(*(*ManifoldPoint)(unsafe.Add(nil, m+12))).TangentImpulse = *(*float32)(unsafe.Add(nil, tangentImpulse1+uintptr(laneIndex)*4))
			(*(*ManifoldPoint)(unsafe.Add(nil, m+12))).TotalNormalImpulse = *(*float32)(unsafe.Add(nil, totalNormalImpulse1+uintptr(laneIndex)*4))
			(*(*ManifoldPoint)(unsafe.Add(nil, m+12))).NormalVelocity = *(*float32)(unsafe.Add(nil, normalVelocity1+uintptr(laneIndex)*4))
			(*(*ManifoldPoint)(unsafe.Add(nil, m+12+1*48))).NormalImpulse = *(*float32)(unsafe.Add(nil, normalImpulse2+uintptr(laneIndex)*4))
			(*(*ManifoldPoint)(unsafe.Add(nil, m+12+1*48))).TangentImpulse = *(*float32)(unsafe.Add(nil, tangentImpulse2+uintptr(laneIndex)*4))
			(*(*ManifoldPoint)(unsafe.Add(nil, m+12+1*48))).TotalNormalImpulse = *(*float32)(unsafe.Add(nil, totalNormalImpulse2+uintptr(laneIndex)*4))
			(*(*ManifoldPoint)(unsafe.Add(nil, m+12+1*48))).NormalVelocity = *(*float32)(unsafe.Add(nil, normalVelocity2+uintptr(laneIndex)*4))
			goto _2
		_2:
			;
			laneIndex = laneIndex + 1
		}
		goto _1
	_1:
		;
		constraintIndex = constraintIndex + 1
	}
}

const _B2_ALIGNMENT = 32
const _BUFSIZ = 1024
const _EXIT_FAILURE = 1
const _EXIT_SUCCESS = 0
const _FILENAME_MAX = 4096
const _FOPEN_MAX = 1000
const _L_ctermid = 20
const _L_cuserid = 20
const _L_tmpnam = 20
const _P_tmpdir = "/tmp"
const _RAND_MAX = 0x7fffffff
const _TMP_MAX = 10000
const _WNOHANG = 1
const _WUNTRACED = 2
const _IOFBF = 0
const _IOLBF = 1
const _IONBF = 2
const alloca1 = "__builtin_alloca"

type div_t struct {
	Quot int32
	Rem  int32
}

type ldiv_t struct {
	Quot int64
	Rem  int64
}

type lldiv_t struct {
	Quot int64
	Rem  int64
}

type ssize_t = int64

type off_t = int64

type va_list = uintptr

type __isoc_va_list = uintptr

type fpos_t struct {
	ｆ__lldata [0]int64
	ｆ__align  [0]float64
	ｆ__opaque [16]uint8
}

type _G_fpos64_t = fpos_t

type cookie_io_functions_t struct {
	Read   uintptr
	Write  uintptr
	Seek   uintptr
	Close1 uintptr
}

type _IO_cookie_io_functions_t = cookie_io_functions_t

func b2SetLengthUnitsPerMeter(tls *_Stack, lengthUnits float32) {
	if !(b2IsValidFloat(tls, lengthUnits) != 0 && lengthUnits > float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+3892, __ccgo_ts+3944, int32FromInt32(37)) != 0 {
		__builtin_trap(tls)
	}
	b2_lengthUnitsPerMeter = lengthUnits
}

func b2GetLengthUnitsPerMeter(tls *_Stack) (r float32) {
	return b2_lengthUnitsPerMeter
}

func b2DefaultAssertFcn(tls *_Stack, condition uintptr, fileName uintptr, lineNumber int32) (r int32) {
	return b2DefaultAssertFcnGo(tls, condition, fileName, lineNumber)
}

type __ccgo_fp__Xb2SetAssertFcn_0 = func(*_Stack, uintptr, uintptr, int32) int32

func b2SetAssertFcn(tls *_Stack, __ccgo_fp_assertFcn uintptr) {
	if !(__ccgo_fp_assertFcn != uintptrFromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+4000, __ccgo_ts+3944, int32FromInt32(58)) != 0 {
		__builtin_trap(tls)
	}
	b2AssertHandler = __ccgo_fp_assertFcn
}

func b2InternalAssertFcn(tls *_Stack, condition uintptr, fileName uintptr, lineNumber int32) (r int32) {
	return (*(*func(*_Stack, uintptr, uintptr, int32) int32)(unsafe.Pointer(&struct{ uintptr }{b2AssertHandler})))(tls, condition, fileName, lineNumber)
}

func b2GetVersion(tls *_Stack) (r Version) {
	return Version{
		Major:    int32(3),
		Minor:    int32(1),
		Revision: int32(1),
	}
}

var b2_allocFcn = uintptr(0)

var b2_freeFcn = uintptr(0)

type __ccgo_fp__Xb2SetAllocator_0 = func(*_Stack, uint32, int32) uintptr

type __ccgo_fp__Xb2SetAllocator_1 = func(*_Stack, uintptr)

func b2SetAllocator(tls *_Stack, __ccgo_fp_allocFcn uintptr, __ccgo_fp_freeFcn uintptr) {
	b2_allocFcn = __ccgo_fp_allocFcn
	b2_freeFcn = __ccgo_fp_freeFcn
}

// Use 32 byte alignment for everything. Works with 256bit SIMD.

func b2Alloc(tls *_Stack, size int32) (r uintptr) {
	var ptr, ptr1 uintptr
	var size32 int32
	_, _, _ = ptr, ptr1, size32
	if size == 0 {
		return uintptrFromInt32(0)
	}
	// This could cause some sharing issues, however Box2D rarely calls b2Alloc.
	_ = __atomic_fetch_addInt32(tls, uintptr(unsafe.Pointer(&b2_byteCount)), size, int32FromInt32(__ATOMIC_SEQ_CST))
	goto _1
_1:
	;
	// Allocation must be a multiple of 32 or risk a seg fault
	// https://en.cppreference.com/w/c/memory/aligned_alloc
	size32 = size - int32(1) | int32(0x1F) + int32(1)
	if b2_allocFcn != uintptrFromInt32(0) {
		ptr = (*(*func(*_Stack, uint32, int32) uintptr)(unsafe.Pointer(&struct{ uintptr }{b2_allocFcn})))(tls, uint32FromInt32(size32), int32(_B2_ALIGNMENT))
		if !(ptr != uintptrFromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+4018, __ccgo_ts+3944, int32FromInt32(111)) != 0 {
			__builtin_trap(tls)
		}
		if !(uint64(ptr)&uint64FromInt32(0x1F) == uint64FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+4030, __ccgo_ts+3944, int32FromInt32(112)) != 0 {
			__builtin_trap(tls)
		}
		return ptr
	}
	ptr1 = __builtin_aligned_alloc(tls, uint64(_B2_ALIGNMENT), uint64FromInt32(size32))
	if !(ptr1 != uintptrFromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+4018, __ccgo_ts+3944, int32FromInt32(132)) != 0 {
		__builtin_trap(tls)
	}
	if !(uint64(ptr1)&uint64FromInt32(0x1F) == uint64FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+4030, __ccgo_ts+3944, int32FromInt32(133)) != 0 {
		__builtin_trap(tls)
	}
	return ptr1
}

func b2Free(tls *_Stack, mem uintptr, size int32) {
	if mem == uintptrFromInt32(0) {
		return
	}
	if b2_freeFcn != uintptrFromInt32(0) {
		(*(*func(*_Stack, uintptr))(unsafe.Pointer(&struct{ uintptr }{b2_freeFcn})))(tls, mem)
	} else {
		free(tls, mem)
	}
	_ = __atomic_fetch_addInt32(tls, uintptr(unsafe.Pointer(&b2_byteCount)), -size, int32FromInt32(__ATOMIC_SEQ_CST))
	goto _1
_1:
}

func b2GrowAlloc(tls *_Stack, oldMem uintptr, oldSize int32, newSize int32) (r uintptr) {
	var newMem uintptr
	_ = newMem
	if !(newSize > oldSize) && b2InternalAssertFcn(tls, __ccgo_ts+4061, __ccgo_ts+3944, int32FromInt32(165)) != 0 {
		__builtin_trap(tls)
	}
	newMem = b2Alloc(tls, newSize)
	if oldSize > 0 {
		memcpy(tls, newMem, oldMem, uint64FromInt32(oldSize))
		b2Free(tls, oldMem, oldSize)
	}
	return newMem
}

func b2GetByteCount(tls *_Stack) (r int32) {
	var v1 int32
	_ = v1
	v1 = atomicLoadNInt32(uintptr(unsafe.Pointer(&b2_byteCount)), int32FromInt32(__ATOMIC_SEQ_CST))
	goto _2
_2:
	return v1
}

/**@}*/

/**
 * @defgroup math_cpp C++ Math
 * @brief Math operator overloads for C++
 *
 * See math_functions.h for details.
 * @{
 */

/**@}*/

func b2GetSweepTransform(tls *_Stack, sweep uintptr, time float32) (r Transform) {
	var invMag, mag, v1, v5 float32
	var q2, qn, v13, v15 Rot
	var xf Transform
	var v10, v11, v2, v3, v6, v7, v9 Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = invMag, mag, q2, qn, xf, v1, v10, v11, v13, v15, v2, v3, v5, v6, v7, v9
	v1 = float32FromFloat32(1) - time
	v2 = (*Sweep)(unsafe.Pointer(sweep)).C1
	v3 = Vec2{
		X: float32(v1 * v2.X),
		Y: float32(v1 * v2.Y),
	}
	goto _4
_4:
	v5 = time
	v6 = (*Sweep)(unsafe.Pointer(sweep)).C2
	v7 = Vec2{
		X: float32(v5 * v6.X),
		Y: float32(v5 * v6.Y),
	}
	goto _8
_8:
	v9 = v3
	v10 = v7
	v11 = Vec2{
		X: v9.X + v10.X,
		Y: v9.Y + v10.Y,
	}
	goto _12
_12:
	xf.P = v11
	q2 = Rot{
		C: float32((float32FromFloat32(1)-time)*(*Sweep)(unsafe.Pointer(sweep)).Q1.C) + float32(time*(*Sweep)(unsafe.Pointer(sweep)).Q2.C),
		S: float32((float32FromFloat32(1)-time)*(*Sweep)(unsafe.Pointer(sweep)).Q1.S) + float32(time*(*Sweep)(unsafe.Pointer(sweep)).Q2.S),
	}
	v13 = q2
	mag = sqrtf(tls, float32(v13.S*v13.S)+float32(v13.C*v13.C))
	if float64(mag) > float64(0) {
		v1 = float32FromFloat32(1) / mag
	} else {
		v1 = float32FromFloat32(0)
	}
	invMag = v1
	qn = Rot{
		C: float32(v13.C * invMag),
		S: float32(v13.S * invMag),
	}
	v15 = qn
	goto _16
_16:
	xf.Q = v15
	// Shift to origin
	v13 = xf.Q
	v2 = (*Sweep)(unsafe.Pointer(sweep)).LocalCenter
	v3 = Vec2{
		X: float32(v13.C*v2.X) - float32(v13.S*v2.Y),
		Y: float32(v13.S*v2.X) + float32(v13.C*v2.Y),
	}
	goto _20
_20:
	v6 = xf.P
	v7 = v3
	v9 = Vec2{
		X: v6.X - v7.X,
		Y: v6.Y - v7.Y,
	}
	goto _24
_24:
	xf.P = v9
	return xf
}

// C documentation
//
//	/// Follows Ericson 5.1.9 Closest Points of Two Line Segments
func b2SegmentDistance(tls *_Stack, p1 Vec2, q1 Vec2, p2 Vec2, q2 Vec2) (r1 SegmentDistanceResult) {
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

func b2MakeProxy(tls *_Stack, points uintptr, count int32, radius float32) (r ShapeProxy) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var i, v1, v2, v3, v5 int32
	var _ /* proxy at bp+0 */ ShapeProxy
	_, _, _, _, _ = i, v1, v2, v3, v5
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
	i = 0
	for {
		if !(i < count) {
			break
		}
		*(*Vec2)(unsafe.Add(nil, bp+uintptr(i)*8)) = *(*Vec2)(unsafe.Add(nil, points+uintptr(i)*8))
		goto _6
	_6:
		;
		i = i + 1
	}
	(*(*ShapeProxy)(unsafe.Add(nil, bp))).Count = count
	(*(*ShapeProxy)(unsafe.Add(nil, bp))).Radius = radius
	return *(*ShapeProxy)(unsafe.Add(nil, bp))
}

func b2MakeOffsetProxy(tls *_Stack, points uintptr, count int32, radius float32, position Vec2, rotation Rot) (r ShapeProxy) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var i, v1, v2, v3, v5 int32
	var transform, v7 Transform
	var x, y float32
	var v8, v9 Vec2
	var _ /* proxy at bp+0 */ ShapeProxy
	_, _, _, _, _, _, _, _, _, _, _ = i, transform, x, y, v1, v2, v3, v5, v7, v8, v9
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
	transform = Transform{
		P: position,
		Q: rotation,
	}
	i = 0
	for {
		if !(i < count) {
			break
		}
		v7 = transform
		v8 = *(*Vec2)(unsafe.Add(nil, points+uintptr(i)*8))
		x = float32(v7.Q.C*v8.X) - float32(v7.Q.S*v8.Y) + v7.P.X
		y = float32(v7.Q.S*v8.X) + float32(v7.Q.C*v8.Y) + v7.P.Y
		v9 = Vec2{
			X: x,
			Y: y,
		}
		goto _10
	_10:
		*(*Vec2)(unsafe.Add(nil, bp+uintptr(i)*8)) = v9
		goto _6
	_6:
		;
		i = i + 1
	}
	(*(*ShapeProxy)(unsafe.Add(nil, bp))).Count = count
	(*(*ShapeProxy)(unsafe.Add(nil, bp))).Radius = radius
	return *(*ShapeProxy)(unsafe.Add(nil, bp))
}

func b2Weight2(tls *_Stack, a1 float32, w1 Vec2, a2 float32, w2 Vec2) (r Vec2) {
	return Vec2{
		X: float32(a1*w1.X) + float32(a2*w2.X),
		Y: float32(a1*w1.Y) + float32(a2*w2.Y),
	}
}

func b2Weight3(tls *_Stack, a1 float32, w1 Vec2, a2 float32, w2 Vec2, a3 float32, w3 Vec2) (r Vec2) {
	return Vec2{
		X: float32(a1*w1.X) + float32(a2*w2.X) + float32(a3*w3.X),
		Y: float32(a1*w1.Y) + float32(a2*w2.Y) + float32(a3*w3.Y),
	}
}

func b2FindSupport(tls *_Stack, proxy uintptr, direction Vec2) (r int32) {
	var bestIndex, count, i int32
	var bestValue, value, v3, v8 float32
	var points uintptr
	var v1, v2, v6, v7 Vec2
	_, _, _, _, _, _, _, _, _, _, _, _ = bestIndex, bestValue, count, i, points, value, v1, v2, v3, v6, v7, v8
	points = proxy
	count = (*ShapeProxy)(unsafe.Pointer(proxy)).Count
	bestIndex = 0
	v1 = *(*Vec2)(unsafe.Add(nil, points))
	v2 = direction
	v3 = float32(v1.X*v2.X) + float32(v1.Y*v2.Y)
	goto _4
_4:
	bestValue = v3
	i = int32(1)
	for {
		if !(i < count) {
			break
		}
		v6 = *(*Vec2)(unsafe.Add(nil, points+uintptr(i)*8))
		v7 = direction
		v8 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y)
		goto _9
	_9:
		value = v8
		if value > bestValue {
			bestIndex = i
			bestValue = value
		}
		goto _5
	_5:
		;
		i = i + 1
	}
	return bestIndex
}

func b2MakeSimplexFromCache(tls *_Stack, cache uintptr, proxyA uintptr, proxyB uintptr) (r Simplex) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var i int32
	var v, v1 uintptr
	var vertices [3]uintptr
	var v2, v3, v4 Vec2
	var _ /* s at bp+0 */ Simplex
	_, _, _, _, _, _, _ = i, v, v1, vertices, v2, v3, v4
	if !(int32FromUint16((*SimplexCache)(unsafe.Pointer(cache)).Count) <= int32FromInt32(3)) && b2InternalAssertFcn(tls, __ccgo_ts+4079, __ccgo_ts+4097, int32FromInt32(170)) != 0 {
		__builtin_trap(tls)
	}
	// Copy data from cache.
	(*(*Simplex)(unsafe.Add(nil, bp))).Count = int32FromUint16((*SimplexCache)(unsafe.Pointer(cache)).Count)
	vertices = [3]uintptr{
		0: bp,
		1: bp + 36,
		2: bp + 72,
	}
	i = 0
	for {
		if !(i < (*(*Simplex)(unsafe.Add(nil, bp))).Count) {
			break
		}
		v = vertices[i]
		(*SimplexVertex)(unsafe.Pointer(v)).IndexA = int32FromUint8(*(*uint8_t)(unsafe.Add(nil, cache+2+uintptr(i))))
		(*SimplexVertex)(unsafe.Pointer(v)).IndexB = int32FromUint8(*(*uint8_t)(unsafe.Add(nil, cache+5+uintptr(i))))
		(*SimplexVertex)(unsafe.Pointer(v)).WA = *(*Vec2)(unsafe.Add(nil, proxyA+uintptr((*SimplexVertex)(unsafe.Pointer(v)).IndexA)*8))
		(*SimplexVertex)(unsafe.Pointer(v)).WB = *(*Vec2)(unsafe.Add(nil, proxyB+uintptr((*SimplexVertex)(unsafe.Pointer(v)).IndexB)*8))
		v2 = (*SimplexVertex)(unsafe.Pointer(v)).WA
		v3 = (*SimplexVertex)(unsafe.Pointer(v)).WB
		v4 = Vec2{
			X: v2.X - v3.X,
			Y: v2.Y - v3.Y,
		}
		goto _5
	_5:
		(*SimplexVertex)(unsafe.Pointer(v)).W = v4
		// invalid
		(*SimplexVertex)(unsafe.Pointer(v)).A = -float32FromFloat32(1)
		goto _1
	_1:
		;
		i = i + 1
	}
	// If the cache is empty or invalid ...
	if (*(*Simplex)(unsafe.Add(nil, bp))).Count == 0 {
		v1 = vertices[0]
		(*SimplexVertex)(unsafe.Pointer(v1)).IndexA = 0
		(*SimplexVertex)(unsafe.Pointer(v1)).IndexB = 0
		(*SimplexVertex)(unsafe.Pointer(v1)).WA = *(*Vec2)(unsafe.Add(nil, proxyA))
		(*SimplexVertex)(unsafe.Pointer(v1)).WB = *(*Vec2)(unsafe.Add(nil, proxyB))
		v2 = (*SimplexVertex)(unsafe.Pointer(v1)).WA
		v3 = (*SimplexVertex)(unsafe.Pointer(v1)).WB
		v4 = Vec2{
			X: v2.X - v3.X,
			Y: v2.Y - v3.Y,
		}
		goto _9
	_9:
		(*SimplexVertex)(unsafe.Pointer(v1)).W = v4
		(*SimplexVertex)(unsafe.Pointer(v1)).A = float32FromFloat32(1)
		(*(*Simplex)(unsafe.Add(nil, bp))).Count = int32(1)
	}
	return *(*Simplex)(unsafe.Add(nil, bp))
}

func b2MakeSimplexCache(tls *_Stack, cache uintptr, simplex uintptr) {
	var i int32
	var vertices [3]uintptr
	_, _ = i, vertices
	(*SimplexCache)(unsafe.Pointer(cache)).Count = uint16FromInt32((*Simplex)(unsafe.Pointer(simplex)).Count)
	vertices = [3]uintptr{
		0: simplex,
		1: simplex + 36,
		2: simplex + 72,
	}
	i = 0
	for {
		if !(i < (*Simplex)(unsafe.Pointer(simplex)).Count) {
			break
		}
		*(*uint8_t)(unsafe.Add(nil, cache+2+uintptr(i))) = uint8FromInt32((*SimplexVertex)(unsafe.Pointer(vertices[i])).IndexA)
		*(*uint8_t)(unsafe.Add(nil, cache+5+uintptr(i))) = uint8FromInt32((*SimplexVertex)(unsafe.Pointer(vertices[i])).IndexB)
		goto _1
	_1:
		;
		i = i + 1
	}
}

// Warning: writing to these globals significantly slows multithreading performance

type b2SeparationType1 = int32

// Warning: writing to these globals significantly slows multithreading performance

type b2SeparationType = int32

const b2_pointsType = 0
const b2_faceAType = 1
const b2_faceBType = 2

type b2SeparationFunction struct {
	ProxyA     uintptr
	ProxyB     uintptr
	SweepA     Sweep
	SweepB     Sweep
	LocalPoint Vec2
	Axis       Vec2
	Type1      b2SeparationType1
}

func b2MakeSeparationFunction(tls *_Stack, cache uintptr, proxyA uintptr, sweepA uintptr, proxyB uintptr, sweepB uintptr, t1 float32) (r b2SeparationFunction) {
	var count int32
	var f b2SeparationFunction
	var invLength, length, s1, s2, x, y, v21 float32
	var localPointA, localPointA1, localPointA11, localPointA2, localPointB, localPointB1, localPointB11, localPointB2, n, normal, normal1, pointA, pointA1, pointA2, pointB, pointB1, pointB2, v10, v11, v13, v14, v16, v17, v2, v3, v6, v7, v9 Vec2
	var xfA, xfB, v1, v5 Transform
	var v27 Rot
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = count, f, invLength, length, localPointA, localPointA1, localPointA11, localPointA2, localPointB, localPointB1, localPointB11, localPointB2, n, normal, normal1, pointA, pointA1, pointA2, pointB, pointB1, pointB2, s1, s2, x, xfA, xfB, y, v1, v10, v11, v13, v14, v16, v17, v2, v21, v27, v3, v5, v6, v7, v9
	f.ProxyA = proxyA
	f.ProxyB = proxyB
	count = int32FromUint16((*SimplexCache)(unsafe.Pointer(cache)).Count)
	if !(0 < count && count < int32(3)) && b2InternalAssertFcn(tls, __ccgo_ts+4427, __ccgo_ts+4097, int32FromInt32(950)) != 0 {
		__builtin_trap(tls)
	}
	f.SweepA = *(*Sweep)(unsafe.Add(nil, sweepA))
	f.SweepB = *(*Sweep)(unsafe.Add(nil, sweepB))
	xfA = b2GetSweepTransform(tls, sweepA, t1)
	xfB = b2GetSweepTransform(tls, sweepB, t1)
	if count == int32(1) {
		f.Type1 = int32(b2_pointsType)
		localPointA = *(*Vec2)(unsafe.Add(nil, proxyA+uintptr(*(*uint8_t)(unsafe.Add(nil, cache+2)))*8))
		localPointB = *(*Vec2)(unsafe.Add(nil, proxyB+uintptr(*(*uint8_t)(unsafe.Add(nil, cache+5)))*8))
		v1 = xfA
		v2 = localPointA
		x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
		y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
		v3 = Vec2{
			X: x,
			Y: y,
		}
		goto _4
	_4:
		pointA = v3
		v5 = xfB
		v6 = localPointB
		x = float32(v5.Q.C*v6.X) - float32(v5.Q.S*v6.Y) + v5.P.X
		y = float32(v5.Q.S*v6.X) + float32(v5.Q.C*v6.Y) + v5.P.Y
		v7 = Vec2{
			X: x,
			Y: y,
		}
		goto _8
	_8:
		pointB = v7
		v9 = pointB
		v10 = pointA
		v11 = Vec2{
			X: v9.X - v10.X,
			Y: v9.Y - v10.Y,
		}
		goto _12
	_12:
		v13 = v11
		length = sqrtf(tls, float32(v13.X*v13.X)+float32(v13.Y*v13.Y))
		if length < float32FromFloat32(1.1920928955078125e-07) {
			v14 = Vec2{}
			goto _15
		}
		invLength = float32FromFloat32(1) / length
		n = Vec2{
			X: float32(invLength * v13.X),
			Y: float32(invLength * v13.Y),
		}
		v14 = n
		goto _15
	_15:
		f.Axis = v14
		f.LocalPoint = b2Vec2_zero
		return f
	}
	if int32FromUint8(*(*uint8_t)(unsafe.Add(nil, cache+2))) == int32FromUint8(*(*uint8_t)(unsafe.Add(nil, cache+2+1))) {
		// Two points on B and one on A.
		f.Type1 = int32(b2_faceBType)
		localPointB11 = *(*Vec2)(unsafe.Add(nil, proxyB+uintptr(*(*uint8_t)(unsafe.Add(nil, cache+5)))*8))
		localPointB2 = *(*Vec2)(unsafe.Add(nil, proxyB+uintptr(*(*uint8_t)(unsafe.Add(nil, cache+5+1)))*8))
		v2 = localPointB2
		v3 = localPointB11
		v6 = Vec2{
			X: v2.X - v3.X,
			Y: v2.Y - v3.Y,
		}
		goto _19
	_19:
		v7 = v6
		v21 = float32FromFloat32(1)
		v9 = Vec2{
			X: float32(v21 * v7.Y),
			Y: float32(-v21 * v7.X),
		}
		goto _23
	_23:
		f.Axis = v9
		v2 = f.Axis
		length = sqrtf(tls, float32(v2.X*v2.X)+float32(v2.Y*v2.Y))
		if length < float32FromFloat32(1.1920928955078125e-07) {
			v3 = Vec2{}
			goto _26
		}
		invLength = float32FromFloat32(1) / length
		n = Vec2{
			X: float32(invLength * v2.X),
			Y: float32(invLength * v2.Y),
		}
		v3 = n
		goto _26
	_26:
		f.Axis = v3
		v27 = xfB.Q
		v2 = f.Axis
		v3 = Vec2{
			X: float32(v27.C*v2.X) - float32(v27.S*v2.Y),
			Y: float32(v27.S*v2.X) + float32(v27.C*v2.Y),
		}
		goto _30
	_30:
		normal = v3
		f.LocalPoint = Vec2{
			X: float32(float32FromFloat32(0.5) * (localPointB11.X + localPointB2.X)),
			Y: float32(float32FromFloat32(0.5) * (localPointB11.Y + localPointB2.Y)),
		}
		v1 = xfB
		v2 = f.LocalPoint
		x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
		y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
		v3 = Vec2{
			X: x,
			Y: y,
		}
		goto _34
	_34:
		pointB1 = v3
		localPointA1 = *(*Vec2)(unsafe.Add(nil, proxyA+uintptr(*(*uint8_t)(unsafe.Add(nil, cache+2)))*8))
		v5 = xfA
		v6 = localPointA1
		x = float32(v5.Q.C*v6.X) - float32(v5.Q.S*v6.Y) + v5.P.X
		y = float32(v5.Q.S*v6.X) + float32(v5.Q.C*v6.Y) + v5.P.Y
		v7 = Vec2{
			X: x,
			Y: y,
		}
		goto _38
	_38:
		pointA1 = v7
		v9 = pointA1
		v10 = pointB1
		v11 = Vec2{
			X: v9.X - v10.X,
			Y: v9.Y - v10.Y,
		}
		goto _42
	_42:
		v13 = v11
		v14 = normal
		v21 = float32(v13.X*v14.X) + float32(v13.Y*v14.Y)
		goto _46
	_46:
		s1 = v21
		if s1 < float32FromFloat32(0) {
			v16 = f.Axis
			v17 = Vec2{
				X: -v16.X,
				Y: -v16.Y,
			}
			goto _49
		_49:
			f.Axis = v17
		}
		return f
	}
	// Two points on A and one or two points on B.
	f.Type1 = int32(b2_faceAType)
	localPointA11 = *(*Vec2)(unsafe.Add(nil, proxyA+uintptr(*(*uint8_t)(unsafe.Add(nil, cache+2)))*8))
	localPointA2 = *(*Vec2)(unsafe.Add(nil, proxyA+uintptr(*(*uint8_t)(unsafe.Add(nil, cache+2+1)))*8))
	v2 = localPointA2
	v3 = localPointA11
	v6 = Vec2{
		X: v2.X - v3.X,
		Y: v2.Y - v3.Y,
	}
	goto _53
_53:
	v7 = v6
	v21 = float32FromFloat32(1)
	v9 = Vec2{
		X: float32(v21 * v7.Y),
		Y: float32(-v21 * v7.X),
	}
	goto _57
_57:
	f.Axis = v9
	v2 = f.Axis
	length = sqrtf(tls, float32(v2.X*v2.X)+float32(v2.Y*v2.Y))
	if length < float32FromFloat32(1.1920928955078125e-07) {
		v3 = Vec2{}
		goto _60
	}
	invLength = float32FromFloat32(1) / length
	n = Vec2{
		X: float32(invLength * v2.X),
		Y: float32(invLength * v2.Y),
	}
	v3 = n
	goto _60
_60:
	f.Axis = v3
	v27 = xfA.Q
	v2 = f.Axis
	v3 = Vec2{
		X: float32(v27.C*v2.X) - float32(v27.S*v2.Y),
		Y: float32(v27.S*v2.X) + float32(v27.C*v2.Y),
	}
	goto _64
_64:
	normal1 = v3
	f.LocalPoint = Vec2{
		X: float32(float32FromFloat32(0.5) * (localPointA11.X + localPointA2.X)),
		Y: float32(float32FromFloat32(0.5) * (localPointA11.Y + localPointA2.Y)),
	}
	v1 = xfA
	v2 = f.LocalPoint
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _68
_68:
	pointA2 = v3
	localPointB1 = *(*Vec2)(unsafe.Add(nil, proxyB+uintptr(*(*uint8_t)(unsafe.Add(nil, cache+5)))*8))
	v5 = xfB
	v6 = localPointB1
	x = float32(v5.Q.C*v6.X) - float32(v5.Q.S*v6.Y) + v5.P.X
	y = float32(v5.Q.S*v6.X) + float32(v5.Q.C*v6.Y) + v5.P.Y
	v7 = Vec2{
		X: x,
		Y: y,
	}
	goto _72
_72:
	pointB2 = v7
	v9 = pointB2
	v10 = pointA2
	v11 = Vec2{
		X: v9.X - v10.X,
		Y: v9.Y - v10.Y,
	}
	goto _76
_76:
	v13 = v11
	v14 = normal1
	v21 = float32(v13.X*v14.X) + float32(v13.Y*v14.Y)
	goto _80
_80:
	s2 = v21
	if s2 < float32FromFloat32(0) {
		v16 = f.Axis
		v17 = Vec2{
			X: -v16.X,
			Y: -v16.Y,
		}
		goto _83
	_83:
		f.Axis = v17
	}
	return f
}

func b2FindMinSeparation(tls *_Stack, f uintptr, indexA uintptr, indexB uintptr, t1 float32) (r float32) {
	var axisA, axisA1, axisB, axisB1, localPointA, localPointA1, localPointB, localPointB1, normal, normal1, pointA, pointA1, pointA2, pointB, pointB1, pointB2, v10, v11, v14, v15, v18, v19, v22, v7, v8 Vec2
	var separation, separation1, separation2, x, y, v31 float32
	var xfA, xfB, v17, v21 Transform
	var v13, v6 Rot
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axisA, axisA1, axisB, axisB1, localPointA, localPointA1, localPointB, localPointB1, normal, normal1, pointA, pointA1, pointA2, pointB, pointB1, pointB2, separation, separation1, separation2, x, xfA, xfB, y, v10, v11, v13, v14, v15, v17, v18, v19, v21, v22, v31, v6, v7, v8
	xfA = b2GetSweepTransform(tls, f+16, t1)
	xfB = b2GetSweepTransform(tls, f+56, t1)
	switch (*b2SeparationFunction)(unsafe.Pointer(f)).Type1 {
	case int32(b2_pointsType):
		goto _1
	case int32(b2_faceAType):
		goto _2
	case int32(b2_faceBType):
		goto _3
	default:
		goto _4
	}
_1:
	;
	v6 = xfA.Q
	v7 = (*b2SeparationFunction)(unsafe.Pointer(f)).Axis
	v8 = Vec2{
		X: float32(v6.C*v7.X) + float32(v6.S*v7.Y),
		Y: float32(-v6.S*v7.X) + float32(v6.C*v7.Y),
	}
	goto _9
_9:
	axisA = v8
	v10 = (*b2SeparationFunction)(unsafe.Pointer(f)).Axis
	v11 = Vec2{
		X: -v10.X,
		Y: -v10.Y,
	}
	goto _12
_12:
	v13 = xfB.Q
	v14 = v11
	v15 = Vec2{
		X: float32(v13.C*v14.X) + float32(v13.S*v14.Y),
		Y: float32(-v13.S*v14.X) + float32(v13.C*v14.Y),
	}
	goto _16
_16:
	axisB = v15
	*(*int32)(unsafe.Add(nil, indexA)) = b2FindSupport(tls, (*b2SeparationFunction)(unsafe.Pointer(f)).ProxyA, axisA)
	*(*int32)(unsafe.Add(nil, indexB)) = b2FindSupport(tls, (*b2SeparationFunction)(unsafe.Pointer(f)).ProxyB, axisB)
	localPointA = *(*Vec2)(unsafe.Add(nil, (*b2SeparationFunction)(unsafe.Pointer(f)).ProxyA+uintptr(*(*int32)(unsafe.Add(nil, indexA)))*8))
	localPointB = *(*Vec2)(unsafe.Add(nil, (*b2SeparationFunction)(unsafe.Pointer(f)).ProxyB+uintptr(*(*int32)(unsafe.Add(nil, indexB)))*8))
	v17 = xfA
	v7 = localPointA
	x = float32(v17.Q.C*v7.X) - float32(v17.Q.S*v7.Y) + v17.P.X
	y = float32(v17.Q.S*v7.X) + float32(v17.Q.C*v7.Y) + v17.P.Y
	v8 = Vec2{
		X: x,
		Y: y,
	}
	goto _20
_20:
	pointA = v8
	v21 = xfB
	v10 = localPointB
	x = float32(v21.Q.C*v10.X) - float32(v21.Q.S*v10.Y) + v21.P.X
	y = float32(v21.Q.S*v10.X) + float32(v21.Q.C*v10.Y) + v21.P.Y
	v11 = Vec2{
		X: x,
		Y: y,
	}
	goto _24
_24:
	pointB = v11
	v14 = pointB
	v15 = pointA
	v18 = Vec2{
		X: v14.X - v15.X,
		Y: v14.Y - v15.Y,
	}
	goto _28
_28:
	v19 = v18
	v22 = (*b2SeparationFunction)(unsafe.Pointer(f)).Axis
	v31 = float32(v19.X*v22.X) + float32(v19.Y*v22.Y)
	goto _32
_32:
	separation = v31
	return separation
_2:
	;
	v6 = xfA.Q
	v7 = (*b2SeparationFunction)(unsafe.Pointer(f)).Axis
	v8 = Vec2{
		X: float32(v6.C*v7.X) - float32(v6.S*v7.Y),
		Y: float32(v6.S*v7.X) + float32(v6.C*v7.Y),
	}
	goto _36
_36:
	normal = v8
	v17 = xfA
	v10 = (*b2SeparationFunction)(unsafe.Pointer(f)).LocalPoint
	x = float32(v17.Q.C*v10.X) - float32(v17.Q.S*v10.Y) + v17.P.X
	y = float32(v17.Q.S*v10.X) + float32(v17.Q.C*v10.Y) + v17.P.Y
	v11 = Vec2{
		X: x,
		Y: y,
	}
	goto _40
_40:
	pointA1 = v11
	v14 = normal
	v15 = Vec2{
		X: -v14.X,
		Y: -v14.Y,
	}
	goto _43
_43:
	v13 = xfB.Q
	v18 = v15
	v19 = Vec2{
		X: float32(v13.C*v18.X) + float32(v13.S*v18.Y),
		Y: float32(-v13.S*v18.X) + float32(v13.C*v18.Y),
	}
	goto _47
_47:
	axisB1 = v19
	*(*int32)(unsafe.Add(nil, indexA)) = -int32(1)
	*(*int32)(unsafe.Add(nil, indexB)) = b2FindSupport(tls, (*b2SeparationFunction)(unsafe.Pointer(f)).ProxyB, axisB1)
	localPointB1 = *(*Vec2)(unsafe.Add(nil, (*b2SeparationFunction)(unsafe.Pointer(f)).ProxyB+uintptr(*(*int32)(unsafe.Add(nil, indexB)))*8))
	v17 = xfB
	v7 = localPointB1
	x = float32(v17.Q.C*v7.X) - float32(v17.Q.S*v7.Y) + v17.P.X
	y = float32(v17.Q.S*v7.X) + float32(v17.Q.C*v7.Y) + v17.P.Y
	v8 = Vec2{
		X: x,
		Y: y,
	}
	goto _51
_51:
	pointB1 = v8
	v10 = pointB1
	v11 = pointA1
	v14 = Vec2{
		X: v10.X - v11.X,
		Y: v10.Y - v11.Y,
	}
	goto _55
_55:
	v15 = v14
	v18 = normal
	v31 = float32(v15.X*v18.X) + float32(v15.Y*v18.Y)
	goto _59
_59:
	separation1 = v31
	return separation1
_3:
	;
	v6 = xfB.Q
	v7 = (*b2SeparationFunction)(unsafe.Pointer(f)).Axis
	v8 = Vec2{
		X: float32(v6.C*v7.X) - float32(v6.S*v7.Y),
		Y: float32(v6.S*v7.X) + float32(v6.C*v7.Y),
	}
	goto _63
_63:
	normal1 = v8
	v17 = xfB
	v10 = (*b2SeparationFunction)(unsafe.Pointer(f)).LocalPoint
	x = float32(v17.Q.C*v10.X) - float32(v17.Q.S*v10.Y) + v17.P.X
	y = float32(v17.Q.S*v10.X) + float32(v17.Q.C*v10.Y) + v17.P.Y
	v11 = Vec2{
		X: x,
		Y: y,
	}
	goto _67
_67:
	pointB2 = v11
	v14 = normal1
	v15 = Vec2{
		X: -v14.X,
		Y: -v14.Y,
	}
	goto _70
_70:
	v13 = xfA.Q
	v18 = v15
	v19 = Vec2{
		X: float32(v13.C*v18.X) + float32(v13.S*v18.Y),
		Y: float32(-v13.S*v18.X) + float32(v13.C*v18.Y),
	}
	goto _74
_74:
	axisA1 = v19
	*(*int32)(unsafe.Add(nil, indexB)) = -int32(1)
	*(*int32)(unsafe.Add(nil, indexA)) = b2FindSupport(tls, (*b2SeparationFunction)(unsafe.Pointer(f)).ProxyA, axisA1)
	localPointA1 = *(*Vec2)(unsafe.Add(nil, (*b2SeparationFunction)(unsafe.Pointer(f)).ProxyA+uintptr(*(*int32)(unsafe.Add(nil, indexA)))*8))
	v17 = xfA
	v7 = localPointA1
	x = float32(v17.Q.C*v7.X) - float32(v17.Q.S*v7.Y) + v17.P.X
	y = float32(v17.Q.S*v7.X) + float32(v17.Q.C*v7.Y) + v17.P.Y
	v8 = Vec2{
		X: x,
		Y: y,
	}
	goto _78
_78:
	pointA2 = v8
	v10 = pointA2
	v11 = pointB2
	v14 = Vec2{
		X: v10.X - v11.X,
		Y: v10.Y - v11.Y,
	}
	goto _82
_82:
	v15 = v14
	v18 = normal1
	v31 = float32(v15.X*v18.X) + float32(v15.Y*v18.Y)
	goto _86
_86:
	separation2 = v31
	return separation2
_4:
	;
_89:
	;
	if bool(!(int32FromInt32(false1) != 0)) && b2InternalAssertFcn(tls, __ccgo_ts+4123, __ccgo_ts+4097, int32FromInt32(1078)) != 0 {
		__builtin_trap(tls)
	}
	goto _88
_88:
	;
	if 0 != 0 {
		goto _89
	}
	goto _87
_87:
	;
	*(*int32)(unsafe.Add(nil, indexA)) = -int32(1)
	*(*int32)(unsafe.Add(nil, indexB)) = -int32(1)
	return float32FromFloat32(0)
}

// C documentation
//
//	//
func b2EvaluateSeparation(tls *_Stack, f uintptr, indexA int32, indexB int32, t1 float32) (r float32) {
	var localPointA, localPointA1, localPointB, localPointB1, normal, normal1, pointA, pointA1, pointA2, pointB, pointB1, pointB2, v11, v12, v14, v15, v16, v18, v19, v23, v24, v7, v8 Vec2
	var separation, separation1, separation2, x, y, v20 float32
	var xfA, xfB, v10, v6 Transform
	var v22 Rot
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = localPointA, localPointA1, localPointB, localPointB1, normal, normal1, pointA, pointA1, pointA2, pointB, pointB1, pointB2, separation, separation1, separation2, x, xfA, xfB, y, v10, v11, v12, v14, v15, v16, v18, v19, v20, v22, v23, v24, v6, v7, v8
	xfA = b2GetSweepTransform(tls, f+16, t1)
	xfB = b2GetSweepTransform(tls, f+56, t1)
	switch (*b2SeparationFunction)(unsafe.Pointer(f)).Type1 {
	case int32(b2_pointsType):
		goto _1
	case int32(b2_faceAType):
		goto _2
	case int32(b2_faceBType):
		goto _3
	default:
		goto _4
	}
_1:
	;
	localPointA = *(*Vec2)(unsafe.Add(nil, (*b2SeparationFunction)(unsafe.Pointer(f)).ProxyA+uintptr(indexA)*8))
	localPointB = *(*Vec2)(unsafe.Add(nil, (*b2SeparationFunction)(unsafe.Pointer(f)).ProxyB+uintptr(indexB)*8))
	v6 = xfA
	v7 = localPointA
	x = float32(v6.Q.C*v7.X) - float32(v6.Q.S*v7.Y) + v6.P.X
	y = float32(v6.Q.S*v7.X) + float32(v6.Q.C*v7.Y) + v6.P.Y
	v8 = Vec2{
		X: x,
		Y: y,
	}
	goto _9
_9:
	pointA = v8
	v10 = xfB
	v11 = localPointB
	x = float32(v10.Q.C*v11.X) - float32(v10.Q.S*v11.Y) + v10.P.X
	y = float32(v10.Q.S*v11.X) + float32(v10.Q.C*v11.Y) + v10.P.Y
	v12 = Vec2{
		X: x,
		Y: y,
	}
	goto _13
_13:
	pointB = v12
	v14 = pointB
	v15 = pointA
	v16 = Vec2{
		X: v14.X - v15.X,
		Y: v14.Y - v15.Y,
	}
	goto _17
_17:
	v18 = v16
	v19 = (*b2SeparationFunction)(unsafe.Pointer(f)).Axis
	v20 = float32(v18.X*v19.X) + float32(v18.Y*v19.Y)
	goto _21
_21:
	separation = v20
	return separation
_2:
	;
	v22 = xfA.Q
	v7 = (*b2SeparationFunction)(unsafe.Pointer(f)).Axis
	v8 = Vec2{
		X: float32(v22.C*v7.X) - float32(v22.S*v7.Y),
		Y: float32(v22.S*v7.X) + float32(v22.C*v7.Y),
	}
	goto _25
_25:
	normal = v8
	v6 = xfA
	v11 = (*b2SeparationFunction)(unsafe.Pointer(f)).LocalPoint
	x = float32(v6.Q.C*v11.X) - float32(v6.Q.S*v11.Y) + v6.P.X
	y = float32(v6.Q.S*v11.X) + float32(v6.Q.C*v11.Y) + v6.P.Y
	v12 = Vec2{
		X: x,
		Y: y,
	}
	goto _29
_29:
	pointA1 = v12
	localPointB1 = *(*Vec2)(unsafe.Add(nil, (*b2SeparationFunction)(unsafe.Pointer(f)).ProxyB+uintptr(indexB)*8))
	v10 = xfB
	v14 = localPointB1
	x = float32(v10.Q.C*v14.X) - float32(v10.Q.S*v14.Y) + v10.P.X
	y = float32(v10.Q.S*v14.X) + float32(v10.Q.C*v14.Y) + v10.P.Y
	v15 = Vec2{
		X: x,
		Y: y,
	}
	goto _33
_33:
	pointB1 = v15
	v16 = pointB1
	v18 = pointA1
	v19 = Vec2{
		X: v16.X - v18.X,
		Y: v16.Y - v18.Y,
	}
	goto _37
_37:
	v23 = v19
	v24 = normal
	v20 = float32(v23.X*v24.X) + float32(v23.Y*v24.Y)
	goto _41
_41:
	separation1 = v20
	return separation1
_3:
	;
	v22 = xfB.Q
	v7 = (*b2SeparationFunction)(unsafe.Pointer(f)).Axis
	v8 = Vec2{
		X: float32(v22.C*v7.X) - float32(v22.S*v7.Y),
		Y: float32(v22.S*v7.X) + float32(v22.C*v7.Y),
	}
	goto _45
_45:
	normal1 = v8
	v6 = xfB
	v11 = (*b2SeparationFunction)(unsafe.Pointer(f)).LocalPoint
	x = float32(v6.Q.C*v11.X) - float32(v6.Q.S*v11.Y) + v6.P.X
	y = float32(v6.Q.S*v11.X) + float32(v6.Q.C*v11.Y) + v6.P.Y
	v12 = Vec2{
		X: x,
		Y: y,
	}
	goto _49
_49:
	pointB2 = v12
	localPointA1 = *(*Vec2)(unsafe.Add(nil, (*b2SeparationFunction)(unsafe.Pointer(f)).ProxyA+uintptr(indexA)*8))
	v10 = xfA
	v14 = localPointA1
	x = float32(v10.Q.C*v14.X) - float32(v10.Q.S*v14.Y) + v10.P.X
	y = float32(v10.Q.S*v14.X) + float32(v10.Q.C*v14.Y) + v10.P.Y
	v15 = Vec2{
		X: x,
		Y: y,
	}
	goto _53
_53:
	pointA2 = v15
	v16 = pointA2
	v18 = pointB2
	v19 = Vec2{
		X: v16.X - v18.X,
		Y: v16.Y - v18.Y,
	}
	goto _57
_57:
	v23 = v19
	v24 = normal1
	v20 = float32(v23.X*v24.X) + float32(v23.Y*v24.Y)
	goto _61
_61:
	separation2 = v20
	return separation2
_4:
	;
_64:
	;
	if bool(!(int32FromInt32(false1) != 0)) && b2InternalAssertFcn(tls, __ccgo_ts+4123, __ccgo_ts+4097, int32FromInt32(1130)) != 0 {
		__builtin_trap(tls)
	}
	goto _63
_63:
	;
	if 0 != 0 {
		goto _64
	}
	goto _62
_62:
	;
	return float32FromFloat32(0)
}

// C documentation
//
//	// CCD via the local separating axis method. This seeks progression
//	// by computing the largest time at which separation is maintained.
func b2TimeOfImpact(tls *_Stack, input uintptr) (r TOIOutput) {
	bp := tls.Alloc(400)
	defer tls.Free(400)
	var a11, a2, qq, s, s1, s2, t, t1, t2, tMax, target, tolerance, totalRadius, v15, v16, v17, v19 float32
	var distanceIterations, k_maxIterations, pushBackIterations, rootIterationCount int32
	var distanceOutput DistanceOutput
	var done, v2, v5 uint8
	var output TOIOutput
	var proxyA, proxyB uintptr
	var xfA, xfB Transform
	var v1, v4 Rot
	var v7 bool
	var _ /* cache at bp+80 */ SimplexCache
	var _ /* distanceInput at bp+88 */ DistanceInput
	var _ /* fcn at bp+272 */ b2SeparationFunction
	var _ /* indexA at bp+392 */ int32
	var _ /* indexB at bp+396 */ int32
	var _ /* sweepA at bp+0 */ Sweep
	var _ /* sweepB at bp+40 */ Sweep
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a11, a2, distanceIterations, distanceOutput, done, k_maxIterations, output, proxyA, proxyB, pushBackIterations, qq, rootIterationCount, s, s1, s2, t, t1, t2, tMax, target, tolerance, totalRadius, xfA, xfB, v1, v15, v16, v17, v19, v2, v4, v5, v7
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
	if !(v7 && v5 != 0) && b2InternalAssertFcn(tls, __ccgo_ts+4450, __ccgo_ts+4097, int32FromInt32(1150)) != 0 {
		__builtin_trap(tls)
	}
	v1 = (*(*Sweep)(unsafe.Add(nil, bp+40))).Q1
	qq = float32(v1.S*v1.S) + float32(v1.C*v1.C)
	v2 = boolUint8(float32FromFloat32(1)-float32FromFloat32(0.0006) < qq && qq < float32FromFloat32(1)+float32FromFloat32(0.0006))
	goto _10
_10:
	;
	if v7 = v2 != 0; v7 {
		v4 = (*(*Sweep)(unsafe.Add(nil, bp+40))).Q2
		qq = float32(v4.S*v4.S) + float32(v4.C*v4.C)
		v5 = boolUint8(float32FromFloat32(1)-float32FromFloat32(0.0006) < qq && qq < float32FromFloat32(1)+float32FromFloat32(0.0006))
		goto _13
	_13:
	}
	if !(v7 && v5 != 0) && b2InternalAssertFcn(tls, __ccgo_ts+4515, __ccgo_ts+4097, int32FromInt32(1151)) != 0 {
		__builtin_trap(tls)
	}
	// todo_erin
	// c1 can be at the origin yet the points are far away
	// b2Vec2 origin = b2Add(sweepA.c1, input->proxyA.points[0]);
	proxyA = input
	proxyB = input + 72
	tMax = (*TOIInput)(unsafe.Pointer(input)).MaxFraction
	totalRadius = (*ShapeProxy)(unsafe.Pointer(proxyA)).Radius + (*ShapeProxy)(unsafe.Pointer(proxyB)).Radius
	v15 = float32(float32FromFloat32(0.005) * b2_lengthUnitsPerMeter)
	v16 = totalRadius - float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)
	if v15 > v16 {
		v19 = v15
	} else {
		v19 = v16
	}
	v17 = v19
	goto _18
_18:
	// todo_erin consider different target
	// float target = b2MaxFloat( B2_LINEAR_SLOP, totalRadius );
	target = v17
	tolerance = float32(float32FromFloat32(0.25) * float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter))
	if !(target > tolerance) && b2InternalAssertFcn(tls, __ccgo_ts+4263, __ccgo_ts+4097, int32FromInt32(1167)) != 0 {
		__builtin_trap(tls)
	}
	t1 = float32FromFloat32(0)
	k_maxIterations = int32(20)
	distanceIterations = 0
	// Prepare input for distance query.
	*(*SimplexCache)(unsafe.Add(nil, bp+80)) = SimplexCache{}
	(*(*DistanceInput)(unsafe.Add(nil, bp+88))).ProxyA = (*TOIInput)(unsafe.Pointer(input)).ProxyA
	(*(*DistanceInput)(unsafe.Add(nil, bp+88))).ProxyB = (*TOIInput)(unsafe.Pointer(input)).ProxyB
	(*(*DistanceInput)(unsafe.Add(nil, bp+88))).UseRadii = boolUint8(false1 != 0)
	// The outer loop progressively attempts to compute new separating axes.
	// This loop terminates when an axis is repeated (no progress is made).
	for {
		xfA = b2GetSweepTransform(tls, bp, t1)
		xfB = b2GetSweepTransform(tls, bp+40, t1)
		// Get the distance between shapes. We can also use the results
		// to get a separating axis.
		(*(*DistanceInput)(unsafe.Add(nil, bp+88))).TransformA = xfA
		(*(*DistanceInput)(unsafe.Add(nil, bp+88))).TransformB = xfB
		distanceOutput = b2ShapeDistance(tls, bp+88, bp+80, uintptrFromInt32(0), 0)
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
			// Victory!
			output.State = int32(b2_toiStateHit)
			output.Fraction = t1
			break
		}
		// Initialize the separating axis.
		*(*b2SeparationFunction)(unsafe.Add(nil, bp+272)) = b2MakeSeparationFunction(tls, bp+80, proxyA, bp, proxyB, bp+40, t1)
		// Compute the TOI on the separating axis. We do this by successively
		// resolving the deepest point. This loop is bounded by the number of vertices.
		done = boolUint8(false1 != 0)
		t2 = tMax
		pushBackIterations = 0
		for {
			s2 = b2FindMinSeparation(tls, bp+272, bp+392, bp+396, t2)
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
				t1 = t2
				break
			}
			// Compute the initial separation of the witness points.
			s1 = b2EvaluateSeparation(tls, bp+272, *(*int32)(unsafe.Add(nil, bp+392)), *(*int32)(unsafe.Add(nil, bp+396)), t1)
			// Check for initial overlap. This might happen if the root finder
			// runs out of iterations.
			if s1 < target-tolerance {
				output.State = int32(b2_toiStateFailed)
				output.Fraction = t1
				done = boolUint8(true1 != 0)
				break
			}
			// Check for touching
			if s1 <= target+tolerance {
				// Victory! t1 should hold the TOI (could be 0.0).
				output.State = int32(b2_toiStateHit)
				output.Fraction = t1
				done = boolUint8(true1 != 0)
				break
			}
			// Compute 1D root of: f(x) - target = 0
			rootIterationCount = 0
			a11 = t1
			a2 = t2
			for {
				if rootIterationCount&int32(1) != 0 {
					// Secant rule to improve convergence.
					t = a11 + float32((target-s1)*(a2-a11))/(s2-s1)
				} else {
					// Bisection to guarantee progress.
					t = float32(float32FromFloat32(0.5) * (a11 + a2))
				}
				rootIterationCount = rootIterationCount + int32(1)
				s = b2EvaluateSeparation(tls, bp+272, *(*int32)(unsafe.Add(nil, bp+392)), *(*int32)(unsafe.Add(nil, bp+396)), t)
				v15 = s - target
				if v15 < float32FromInt32(0) {
					v17 = -v15
				} else {
					v17 = v15
				}
				v16 = v17
				goto _25
			_25:
				if v16 < tolerance {
					// t2 holds a tentative value for t1
					t2 = t
					break
				}
				// Ensure we continue to bracket the root.
				if s > target {
					a11 = t
					s1 = s
				} else {
					a2 = t
					s2 = s
				}
				if rootIterationCount == int32(50) {
					break
				}
				goto _22
			_22:
			}
			pushBackIterations = pushBackIterations + int32(1)
			if pushBackIterations == int32(_B2_MAX_POLYGON_VERTICES) {
				break
			}
			goto _21
		_21:
		}
		if done != 0 {
			break
		}
		if distanceIterations == k_maxIterations {
			// Root finder got stuck. Semi-victory.
			output.State = int32(b2_toiStateFailed)
			output.Fraction = t1
			break
		}
		goto _20
	_20:
	}
	return output
}

var b2_identityBodyState5 = b2BodyState{
	DeltaRotation: Rot{
		C: float32FromFloat32(1),
	},
}

const _B2_TREE_HEURISTIC = 0
const _B2_TREE_STACK_SIZE = 1024
const _FLT_MAX2 = 3.4028234663852886e+38
const _UINT64_MAX3 = 18446744073709551615

// todo externalize this to visualize internal nodes and speed up FindPairs

// C documentation
//
//	// A node in the dynamic tree.
type b2TreeNode struct {
	Aabb         AABB
	CategoryBits uint64
	ｆ__ccgo2_24  struct {
		UserData [0]uint64_t
		Children struct {
			Child1 int32
			Child2 int32
		}
	}
	ｆ__ccgo3_32 struct {
		Next   [0]int32_t
		Parent int32
	}
	Height uint16
	Flags  uint16
}

var b2_defaultTreeNode = b2TreeNode{
	CategoryBits: uint64(_B2_DEFAULT_CATEGORY_BITS),
	ｆ__ccgo2_24: *(*struct {
		UserData [0]uint64_t
		Children struct {
			Child1 int32
			Child2 int32
		}
	})(unsafe.Pointer(&[1]uint64{0xffffffffffffffff})),
	ｆ__ccgo3_32: *(*struct {
		Next   [0]int32_t
		Parent int32
	})(unsafe.Pointer(&[1]uint32{0xffffffff})),
	Flags: uint16(b2_allocatedNode),
}

func b2IsLeaf(tls *_Stack, node uintptr) (r uint8) {
	return uint8FromInt32(boolInt32(int32FromUint16((*b2TreeNode)(unsafe.Pointer(node)).Flags)&int32(b2_leafNode) != 0))
}

func b2IsAllocated(tls *_Stack, node uintptr) (r uint8) {
	return uint8FromInt32(boolInt32(int32FromUint16((*b2TreeNode)(unsafe.Pointer(node)).Flags)&int32(b2_allocatedNode) != 0))
}

func b2MaxUInt16(tls *_Stack, a uint16, b uint16) (r uint16) {
	var v1 int32
	_ = v1
	if int32FromUint16(a) > int32FromUint16(b) {
		v1 = int32FromUint16(a)
	} else {
		v1 = int32FromUint16(b)
	}
	return uint16FromInt32(v1)
}

// C documentation
//
//	// Allocate a node from the pool. Grow the pool if necessary.
func b2AllocateNode(tls *_Stack, tree uintptr) (r int32) {
	var i, nodeIndex, oldCapacity int32
	var node, oldNodes uintptr
	_, _, _, _, _ = i, node, nodeIndex, oldCapacity, oldNodes
	// Expand the node pool as needed.
	if (*DynamicTree)(unsafe.Pointer(tree)).FreeList == -int32(1) {
		if !((*DynamicTree)(unsafe.Pointer(tree)).NodeCount == (*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity) && b2InternalAssertFcn(tls, __ccgo_ts+4708, __ccgo_ts+4746, int32FromInt32(127)) != 0 {
			__builtin_trap(tls)
		}
		// The free list is empty. Rebuild a bigger pool.
		oldNodes = (*DynamicTree)(unsafe.Pointer(tree)).Nodes
		oldCapacity = (*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity
		*(*int32)(unsafe.Add(nil, tree+16)) += oldCapacity >> int32(1)
		(*DynamicTree)(unsafe.Pointer(tree)).Nodes = b2Alloc(tls, int32FromUint64(uint64FromInt32((*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity)*uint64(40)))
		if !(oldNodes != uintptrFromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+4776, __ccgo_ts+4746, int32FromInt32(134)) != 0 {
			__builtin_trap(tls)
		}
		memcpy(tls, (*DynamicTree)(unsafe.Pointer(tree)).Nodes, oldNodes, uint64FromInt32((*DynamicTree)(unsafe.Pointer(tree)).NodeCount)*uint64(40))
		memset(tls, (*DynamicTree)(unsafe.Pointer(tree)).Nodes+uintptr((*DynamicTree)(unsafe.Pointer(tree)).NodeCount)*40, 0, uint64FromInt32((*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity-(*DynamicTree)(unsafe.Pointer(tree)).NodeCount)*uint64(40))
		b2Free(tls, oldNodes, int32FromUint64(uint64FromInt32(oldCapacity)*uint64(40)))
		// Build a linked list for the free list. The parent pointer becomes the "next" pointer.
		// todo avoid building freelist?
		i = (*DynamicTree)(unsafe.Pointer(tree)).NodeCount
		for {
			if !(i < (*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity-int32(1)) {
				break
			}
			*(*int32_t)(unsafe.Add(unsafe.Pointer(&*(*b2TreeNode)(unsafe.Add(nil, (*DynamicTree)(unsafe.Pointer(tree)).Nodes+uintptr(i)*40))), 32)) = i + int32(1)
			goto _1
		_1:
			;
			i = i + 1
		}
		*(*int32_t)(unsafe.Add(unsafe.Pointer(&*(*b2TreeNode)(unsafe.Add(nil, (*DynamicTree)(unsafe.Pointer(tree)).Nodes+uintptr((*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity-int32(1))*40))), 32)) = -int32FromInt32(1)
		(*DynamicTree)(unsafe.Pointer(tree)).FreeList = (*DynamicTree)(unsafe.Pointer(tree)).NodeCount
	}
	// Peel a node off the free list.
	nodeIndex = (*DynamicTree)(unsafe.Pointer(tree)).FreeList
	node = (*DynamicTree)(unsafe.Pointer(tree)).Nodes + uintptr(nodeIndex)*40
	(*DynamicTree)(unsafe.Pointer(tree)).FreeList = *(*int32_t)(unsafe.Add(unsafe.Pointer(node), 32))
	*(*b2TreeNode)(unsafe.Add(nil, node)) = b2_defaultTreeNode
	(*DynamicTree)(unsafe.Pointer(tree)).NodeCount = (*DynamicTree)(unsafe.Pointer(tree)).NodeCount + 1
	return nodeIndex
}

// C documentation
//
//	// Return a node to the pool.
func b2FreeNode(tls *_Stack, tree uintptr, nodeId int32) {
	if !(0 <= nodeId && nodeId < (*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity) && b2InternalAssertFcn(tls, __ccgo_ts+4793, __ccgo_ts+4746, int32FromInt32(162)) != 0 {
		__builtin_trap(tls)
	}
	if !(int32FromInt32(0) < (*DynamicTree)(unsafe.Pointer(tree)).NodeCount) && b2InternalAssertFcn(tls, __ccgo_ts+4836, __ccgo_ts+4746, int32FromInt32(163)) != 0 {
		__builtin_trap(tls)
	}
	*(*int32_t)(unsafe.Add(unsafe.Pointer(&*(*b2TreeNode)(unsafe.Add(nil, (*DynamicTree)(unsafe.Pointer(tree)).Nodes+uintptr(nodeId)*40))), 32)) = (*DynamicTree)(unsafe.Pointer(tree)).FreeList
	(*(*b2TreeNode)(unsafe.Add(nil, (*DynamicTree)(unsafe.Pointer(tree)).Nodes+uintptr(nodeId)*40))).Flags = uint16(0)
	(*DynamicTree)(unsafe.Pointer(tree)).FreeList = nodeId
	(*DynamicTree)(unsafe.Pointer(tree)).NodeCount = (*DynamicTree)(unsafe.Pointer(tree)).NodeCount - 1
}

// Greedy algorithm for sibling selection using the SAH
// We have three nodes A-(B,C) and want to add a leaf D, there are three choices.
// 1: make a new parent for A and D : E-(A-(B,C), D)
// 2: associate D with B
//   a: B is a leaf : A-(E-(B,D), C)
//   b: B is an internal node: A-(B{D},C)
// 3: associate D with C
//   a: C is a leaf : A-(B, E-(C,D))
//   b: C is an internal node: A-(B, C{D})
// All of these have a clear cost except when B or C is an internal node. Hence we need to be greedy.

// The cost for cases 1, 2a, and 3a can be computed using the sibling cost formula.
// cost of sibling H = area(union(H, D)) + increased area of ancestors

// C documentation
//
//	// Suppose B (or C) is an internal node, then the lowest cost would be one of two cases:
//	// case1: D becomes a sibling of B
//	// case2: D becomes a descendant of B along with a new internal node of area(D).
func b2FindBestSibling(tls *_Stack, tree uintptr, boxD AABB) (r int32) {
	var area1, area2, areaBase, areaD, bestCost, cost, cost1, cost2, directCost, directCost1, directCost2, inheritedCost, lowerCost1, lowerCost2, wx, wy, v12, v13, v14, v16, v17, v18, v19, v21, v22, v23, v24, v26, v27, v28, v29, v31, v35, v5, v8 float32
	var b3, centerD, d1, d2, v108, v110, v111, v112, v115, v117, v118, v119, v2 Vec2
	var bestSibling, child1, child2, index, rootIndex int32
	var box1, box2, c, rootBox, v1, v10, v11, v32, v34, v4, v7 AABB
	var leaf1, leaf2 uint8
	var nodes uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = area1, area2, areaBase, areaD, b3, bestCost, bestSibling, box1, box2, c, centerD, child1, child2, cost, cost1, cost2, d1, d2, directCost, directCost1, directCost2, index, inheritedCost, leaf1, leaf2, lowerCost1, lowerCost2, nodes, rootBox, rootIndex, wx, wy, v1, v10, v108, v11, v110, v111, v112, v115, v117, v118, v119, v12, v13, v14, v16, v17, v18, v19, v2, v21, v22, v23, v24, v26, v27, v28, v29, v31, v32, v34, v35, v4, v5, v7, v8
	v1 = boxD
	b3 = Vec2{
		X: float32(float32FromFloat32(0.5) * (v1.LowerBound.X + v1.UpperBound.X)),
		Y: float32(float32FromFloat32(0.5) * (v1.LowerBound.Y + v1.UpperBound.Y)),
	}
	v2 = b3
	goto _3
_3:
	centerD = v2
	v4 = boxD
	wx = v4.UpperBound.X - v4.LowerBound.X
	wy = v4.UpperBound.Y - v4.LowerBound.Y
	v5 = float32(float32FromFloat32(2) * (wx + wy))
	goto _6
_6:
	areaD = v5
	nodes = (*DynamicTree)(unsafe.Pointer(tree)).Nodes
	rootIndex = (*DynamicTree)(unsafe.Pointer(tree)).Root
	rootBox = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(rootIndex)*40))).Aabb
	v7 = rootBox
	wx = v7.UpperBound.X - v7.LowerBound.X
	wy = v7.UpperBound.Y - v7.LowerBound.Y
	v8 = float32(float32FromFloat32(2) * (wx + wy))
	goto _9
_9:
	// Area of current node
	areaBase = v8
	v10 = rootBox
	v11 = boxD
	v12 = v10.LowerBound.X
	v13 = v11.LowerBound.X
	if v12 < v13 {
		v16 = v12
	} else {
		v16 = v13
	}
	v14 = v16
	goto _15
_15:
	c.LowerBound.X = v14
	v17 = v10.LowerBound.Y
	v18 = v11.LowerBound.Y
	if v17 < v18 {
		v21 = v17
	} else {
		v21 = v18
	}
	v19 = v21
	goto _20
_20:
	c.LowerBound.Y = v19
	v22 = v10.UpperBound.X
	v23 = v11.UpperBound.X
	if v22 > v23 {
		v26 = v22
	} else {
		v26 = v23
	}
	v24 = v26
	goto _25
_25:
	c.UpperBound.X = v24
	v27 = v10.UpperBound.Y
	v28 = v11.UpperBound.Y
	if v27 > v28 {
		v31 = v27
	} else {
		v31 = v28
	}
	v29 = v31
	goto _30
_30:
	c.UpperBound.Y = v29
	v32 = c
	goto _33
_33:
	v34 = v32
	wx = v34.UpperBound.X - v34.LowerBound.X
	wy = v34.UpperBound.Y - v34.LowerBound.Y
	v35 = float32(float32FromFloat32(2) * (wx + wy))
	goto _36
_36:
	// Area of inflated node
	directCost = v35
	inheritedCost = float32FromFloat32(0)
	bestSibling = rootIndex
	bestCost = directCost
	// Descend the tree from root, following a single greedy path.
	index = rootIndex
	for int32FromUint16((*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(index)*40))).Height) > 0 {
		child1 = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(index)*40))).ｆ__ccgo2_24.Children.Child1
		child2 = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(index)*40))).ｆ__ccgo2_24.Children.Child2
		// Cost of creating a new parent for this node and the new leaf
		cost = directCost + inheritedCost
		// Sometimes there are multiple identical costs within tolerance.
		// This breaks the ties using the centroid distance.
		if cost < bestCost {
			bestSibling = index
			bestCost = cost
		}
		// Inheritance cost seen by children
		inheritedCost = inheritedCost + (directCost - areaBase)
		leaf1 = boolUint8(int32FromUint16((*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(child1)*40))).Height) == 0)
		leaf2 = boolUint8(int32FromUint16((*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(child2)*40))).Height) == 0)
		// Cost of descending into child 1
		lowerCost1 = float32FromFloat32(3.4028234663852886e+38)
		box1 = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(child1)*40))).Aabb
		v1 = box1
		v4 = boxD
		v5 = v1.LowerBound.X
		v8 = v4.LowerBound.X
		if v5 < v8 {
			v13 = v5
		} else {
			v13 = v8
		}
		v12 = v13
		goto _42
	_42:
		c.LowerBound.X = v12
		v14 = v1.LowerBound.Y
		v16 = v4.LowerBound.Y
		if v14 < v16 {
			v18 = v14
		} else {
			v18 = v16
		}
		v17 = v18
		goto _47
	_47:
		c.LowerBound.Y = v17
		v19 = v1.UpperBound.X
		v21 = v4.UpperBound.X
		if v19 > v21 {
			v23 = v19
		} else {
			v23 = v21
		}
		v22 = v23
		goto _52
	_52:
		c.UpperBound.X = v22
		v24 = v1.UpperBound.Y
		v26 = v4.UpperBound.Y
		if v24 > v26 {
			v28 = v24
		} else {
			v28 = v26
		}
		v27 = v28
		goto _57
	_57:
		c.UpperBound.Y = v27
		v7 = c
		goto _60
	_60:
		v10 = v7
		wx = v10.UpperBound.X - v10.LowerBound.X
		wy = v10.UpperBound.Y - v10.LowerBound.Y
		v29 = float32(float32FromFloat32(2) * (wx + wy))
		goto _63
	_63:
		directCost1 = v29
		area1 = float32FromFloat32(0)
		if leaf1 != 0 {
			// Child 1 is a leaf
			// Cost of creating new node and increasing area of node P
			cost1 = directCost1 + inheritedCost
			// Need this here due to while condition above
			if cost1 < bestCost {
				bestSibling = child1
				bestCost = cost1
			}
		} else {
			// Child 1 is an internal node
			v1 = box1
			wx = v1.UpperBound.X - v1.LowerBound.X
			wy = v1.UpperBound.Y - v1.LowerBound.Y
			v5 = float32(float32FromFloat32(2) * (wx + wy))
			goto _66
		_66:
			area1 = v5
			// Lower bound cost of inserting under child 1. The minimum accounts for two possibilities:
			// 1. Child1 could be the sibling with cost1 = inheritedCost + directCost1
			// 2. A descendent of child1 could be the sibling with the lower bound cost of
			//       cost1 = inheritedCost + (directCost1 - area1) + areaD
			// This minimum here leads to the minimum of these two costs.
			v5 = areaD - area1
			v8 = float32FromFloat32(0)
			if v5 < v8 {
				v13 = v5
			} else {
				v13 = v8
			}
			v12 = v13
			goto _70
		_70:
			lowerCost1 = inheritedCost + directCost1 + v12
		}
		// Cost of descending into child 2
		lowerCost2 = float32FromFloat32(3.4028234663852886e+38)
		box2 = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(child2)*40))).Aabb
		v1 = box2
		v4 = boxD
		v5 = v1.LowerBound.X
		v8 = v4.LowerBound.X
		if v5 < v8 {
			v13 = v5
		} else {
			v13 = v8
		}
		v12 = v13
		goto _77
	_77:
		c.LowerBound.X = v12
		v14 = v1.LowerBound.Y
		v16 = v4.LowerBound.Y
		if v14 < v16 {
			v18 = v14
		} else {
			v18 = v16
		}
		v17 = v18
		goto _82
	_82:
		c.LowerBound.Y = v17
		v19 = v1.UpperBound.X
		v21 = v4.UpperBound.X
		if v19 > v21 {
			v23 = v19
		} else {
			v23 = v21
		}
		v22 = v23
		goto _87
	_87:
		c.UpperBound.X = v22
		v24 = v1.UpperBound.Y
		v26 = v4.UpperBound.Y
		if v24 > v26 {
			v28 = v24
		} else {
			v28 = v26
		}
		v27 = v28
		goto _92
	_92:
		c.UpperBound.Y = v27
		v7 = c
		goto _95
	_95:
		v10 = v7
		wx = v10.UpperBound.X - v10.LowerBound.X
		wy = v10.UpperBound.Y - v10.LowerBound.Y
		v29 = float32(float32FromFloat32(2) * (wx + wy))
		goto _98
	_98:
		directCost2 = v29
		area2 = float32FromFloat32(0)
		if leaf2 != 0 {
			cost2 = directCost2 + inheritedCost
			if cost2 < bestCost {
				bestSibling = child2
				bestCost = cost2
			}
		} else {
			v1 = box2
			wx = v1.UpperBound.X - v1.LowerBound.X
			wy = v1.UpperBound.Y - v1.LowerBound.Y
			v5 = float32(float32FromFloat32(2) * (wx + wy))
			goto _101
		_101:
			area2 = v5
			v5 = areaD - area2
			v8 = float32FromFloat32(0)
			if v5 < v8 {
				v13 = v5
			} else {
				v13 = v8
			}
			v12 = v13
			goto _105
		_105:
			lowerCost2 = inheritedCost + directCost2 + v12
		}
		if leaf1 != 0 && leaf2 != 0 {
			break
		}
		// Can the cost possibly be decreased?
		if bestCost <= lowerCost1 && bestCost <= lowerCost2 {
			break
		}
		if lowerCost1 == lowerCost2 && int32FromUint8(leaf1) == false1 {
			if !(lowerCost1 < float32FromFloat32(3.4028234663852886e+38)) && b2InternalAssertFcn(tls, __ccgo_ts+4856, __ccgo_ts+4746, int32FromInt32(296)) != 0 {
				__builtin_trap(tls)
			}
			if !(lowerCost2 < float32FromFloat32(3.4028234663852886e+38)) && b2InternalAssertFcn(tls, __ccgo_ts+4877, __ccgo_ts+4746, int32FromInt32(297)) != 0 {
				__builtin_trap(tls)
			}
			v1 = box1
			b3 = Vec2{
				X: float32(float32FromFloat32(0.5) * (v1.LowerBound.X + v1.UpperBound.X)),
				Y: float32(float32FromFloat32(0.5) * (v1.LowerBound.Y + v1.UpperBound.Y)),
			}
			v2 = b3
			goto _109
		_109:
			v108 = v2
			v110 = centerD
			v111 = Vec2{
				X: v108.X - v110.X,
				Y: v108.Y - v110.Y,
			}
			goto _113
		_113:
			// No clear choice based on lower bound surface area. This can happen when both
			// children fully contain D. Fall back to node distance.
			d1 = v111
			v4 = box2
			b3 = Vec2{
				X: float32(float32FromFloat32(0.5) * (v4.LowerBound.X + v4.UpperBound.X)),
				Y: float32(float32FromFloat32(0.5) * (v4.LowerBound.Y + v4.UpperBound.Y)),
			}
			v112 = b3
			goto _116
		_116:
			v115 = v112
			v117 = centerD
			v118 = Vec2{
				X: v115.X - v117.X,
				Y: v115.Y - v117.Y,
			}
			goto _120
		_120:
			d2 = v118
			v119 = d1
			v5 = float32(v119.X*v119.X) + float32(v119.Y*v119.Y)
			goto _123
		_123:
			lowerCost1 = v5
			v2 = d2
			v5 = float32(v2.X*v2.X) + float32(v2.Y*v2.Y)
			goto _126
		_126:
			lowerCost2 = v5
		}
		// Descend
		if lowerCost1 < lowerCost2 && int32FromUint8(leaf1) == false1 {
			index = child1
			areaBase = area1
			directCost = directCost1
		} else {
			index = child2
			areaBase = area2
			directCost = directCost2
		}
		if !(int32FromUint16((*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(index)*40))).Height) > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+4898, __ccgo_ts+4746, int32FromInt32(321)) != 0 {
			__builtin_trap(tls)
		}
	}
	return bestSibling
}

type b2RotateType = int32

const b2_rotateNone = 0
const b2_rotateBF = 1
const b2_rotateBG = 2
const b2_rotateCD = 3
const b2_rotateCE = 4

// C documentation
//
//	// Perform a left or right rotation if node A is imbalanced.
//	// Returns the new root index.
func b2RotateNodes(tls *_Stack, tree uintptr, iA int32) {
	var A, B, C, D, D1, E, E1, F, F1, G, G1, nodes, v58 uintptr
	var aabbBF, aabbBF1, aabbBG, aabbBG1, aabbCD, aabbCD1, aabbCE, aabbCE1, c, v1, v26, v28, v31, v32, v4, v5, v53, v55 AABB
	var areaB, areaC, bestCost, costBF, costBF1, costBG, costBG1, costBase, costBase1, costBase2, costCD, costCD1, costCE, costCE1, wx, wy, v10, v11, v12, v13, v15, v16, v17, v18, v2, v20, v21, v22, v23, v25, v29, v33, v34, v35, v37, v38, v39, v40, v42, v43, v44, v45, v47, v48, v49, v50, v52, v56, v6, v7, v8 float32
	var bestRotation b2RotateType
	var iB, iC, iD, iD1, iE, iE1, iF, iF1, iG, iG1 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = A, B, C, D, D1, E, E1, F, F1, G, G1, aabbBF, aabbBF1, aabbBG, aabbBG1, aabbCD, aabbCD1, aabbCE, aabbCE1, areaB, areaC, bestCost, bestRotation, c, costBF, costBF1, costBG, costBG1, costBase, costBase1, costBase2, costCD, costCD1, costCE, costCE1, iB, iC, iD, iD1, iE, iE1, iF, iF1, iG, iG1, nodes, wx, wy, v1, v10, v11, v12, v13, v15, v16, v17, v18, v2, v20, v21, v22, v23, v25, v26, v28, v29, v31, v32, v33, v34, v35, v37, v38, v39, v4, v40, v42, v43, v44, v45, v47, v48, v49, v5, v50, v52, v53, v55, v56, v58, v6, v7, v8
	if !(iA != -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+4922, __ccgo_ts+4746, int32FromInt32(340)) != 0 {
		__builtin_trap(tls)
	}
	nodes = (*DynamicTree)(unsafe.Pointer(tree)).Nodes
	A = nodes + uintptr(iA)*40
	if int32FromUint16((*b2TreeNode)(unsafe.Pointer(A)).Height) < int32(2) {
		return
	}
	iB = (*b2TreeNode)(unsafe.Pointer(A)).ｆ__ccgo2_24.Children.Child1
	iC = (*b2TreeNode)(unsafe.Pointer(A)).ｆ__ccgo2_24.Children.Child2
	if !(0 <= iB && iB < (*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity) && b2InternalAssertFcn(tls, __ccgo_ts+4942, __ccgo_ts+4746, int32FromInt32(352)) != 0 {
		__builtin_trap(tls)
	}
	if !(0 <= iC && iC < (*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity) && b2InternalAssertFcn(tls, __ccgo_ts+4977, __ccgo_ts+4746, int32FromInt32(353)) != 0 {
		__builtin_trap(tls)
	}
	B = nodes + uintptr(iB)*40
	C = nodes + uintptr(iC)*40
	if int32FromUint16((*b2TreeNode)(unsafe.Pointer(B)).Height) == 0 {
		// B is a leaf and C is internal
		if !(int32FromUint16((*b2TreeNode)(unsafe.Pointer(C)).Height) > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+5012, __ccgo_ts+4746, int32FromInt32(361)) != 0 {
			__builtin_trap(tls)
		}
		iF = (*b2TreeNode)(unsafe.Pointer(C)).ｆ__ccgo2_24.Children.Child1
		iG = (*b2TreeNode)(unsafe.Pointer(C)).ｆ__ccgo2_24.Children.Child2
		F = nodes + uintptr(iF)*40
		G = nodes + uintptr(iG)*40
		if !(0 <= iF && iF < (*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity) && b2InternalAssertFcn(tls, __ccgo_ts+5026, __ccgo_ts+4746, int32FromInt32(367)) != 0 {
			__builtin_trap(tls)
		}
		if !(0 <= iG && iG < (*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity) && b2InternalAssertFcn(tls, __ccgo_ts+5061, __ccgo_ts+4746, int32FromInt32(368)) != 0 {
			__builtin_trap(tls)
		}
		v1 = (*b2TreeNode)(unsafe.Pointer(C)).Aabb
		wx = v1.UpperBound.X - v1.LowerBound.X
		wy = v1.UpperBound.Y - v1.LowerBound.Y
		v2 = float32(float32FromFloat32(2) * (wx + wy))
		goto _3
	_3:
		// Base cost
		costBase = v2
		v4 = (*b2TreeNode)(unsafe.Pointer(B)).Aabb
		v5 = (*b2TreeNode)(unsafe.Pointer(G)).Aabb
		v6 = v4.LowerBound.X
		v7 = v5.LowerBound.X
		if v6 < v7 {
			v10 = v6
		} else {
			v10 = v7
		}
		v8 = v10
		goto _9
	_9:
		c.LowerBound.X = v8
		v11 = v4.LowerBound.Y
		v12 = v5.LowerBound.Y
		if v11 < v12 {
			v15 = v11
		} else {
			v15 = v12
		}
		v13 = v15
		goto _14
	_14:
		c.LowerBound.Y = v13
		v16 = v4.UpperBound.X
		v17 = v5.UpperBound.X
		if v16 > v17 {
			v20 = v16
		} else {
			v20 = v17
		}
		v18 = v20
		goto _19
	_19:
		c.UpperBound.X = v18
		v21 = v4.UpperBound.Y
		v22 = v5.UpperBound.Y
		if v21 > v22 {
			v25 = v21
		} else {
			v25 = v22
		}
		v23 = v25
		goto _24
	_24:
		c.UpperBound.Y = v23
		v26 = c
		goto _27
	_27:
		// Cost of swapping B and F
		aabbBG = v26
		v28 = aabbBG
		wx = v28.UpperBound.X - v28.LowerBound.X
		wy = v28.UpperBound.Y - v28.LowerBound.Y
		v29 = float32(float32FromFloat32(2) * (wx + wy))
		goto _30
	_30:
		costBF = v29
		v31 = (*b2TreeNode)(unsafe.Pointer(B)).Aabb
		v32 = (*b2TreeNode)(unsafe.Pointer(F)).Aabb
		v33 = v31.LowerBound.X
		v34 = v32.LowerBound.X
		if v33 < v34 {
			v37 = v33
		} else {
			v37 = v34
		}
		v35 = v37
		goto _36
	_36:
		c.LowerBound.X = v35
		v38 = v31.LowerBound.Y
		v39 = v32.LowerBound.Y
		if v38 < v39 {
			v42 = v38
		} else {
			v42 = v39
		}
		v40 = v42
		goto _41
	_41:
		c.LowerBound.Y = v40
		v43 = v31.UpperBound.X
		v44 = v32.UpperBound.X
		if v43 > v44 {
			v47 = v43
		} else {
			v47 = v44
		}
		v45 = v47
		goto _46
	_46:
		c.UpperBound.X = v45
		v48 = v31.UpperBound.Y
		v49 = v32.UpperBound.Y
		if v48 > v49 {
			v52 = v48
		} else {
			v52 = v49
		}
		v50 = v52
		goto _51
	_51:
		c.UpperBound.Y = v50
		v53 = c
		goto _54
	_54:
		// Cost of swapping B and G
		aabbBF = v53
		v55 = aabbBF
		wx = v55.UpperBound.X - v55.LowerBound.X
		wy = v55.UpperBound.Y - v55.LowerBound.Y
		v56 = float32(float32FromFloat32(2) * (wx + wy))
		goto _57
	_57:
		costBG = v56
		if costBase < costBF && costBase < costBG {
			// Rotation does not improve cost
			return
		}
		if costBF < costBG {
			// Swap B and F
			(*b2TreeNode)(unsafe.Pointer(A)).ｆ__ccgo2_24.Children.Child1 = iF
			(*b2TreeNode)(unsafe.Pointer(C)).ｆ__ccgo2_24.Children.Child1 = iB
			(*b2TreeNode)(unsafe.Pointer(B)).ｆ__ccgo3_32.Parent = iC
			(*b2TreeNode)(unsafe.Pointer(F)).ｆ__ccgo3_32.Parent = iA
			(*b2TreeNode)(unsafe.Pointer(C)).Aabb = aabbBG
			(*b2TreeNode)(unsafe.Pointer(C)).Height = uint16FromInt32(int32(1) + int32FromUint16(b2MaxUInt16(tls, (*b2TreeNode)(unsafe.Pointer(B)).Height, (*b2TreeNode)(unsafe.Pointer(G)).Height)))
			(*b2TreeNode)(unsafe.Pointer(A)).Height = uint16FromInt32(int32(1) + int32FromUint16(b2MaxUInt16(tls, (*b2TreeNode)(unsafe.Pointer(C)).Height, (*b2TreeNode)(unsafe.Pointer(F)).Height)))
			(*b2TreeNode)(unsafe.Pointer(C)).CategoryBits = (*b2TreeNode)(unsafe.Pointer(B)).CategoryBits | (*b2TreeNode)(unsafe.Pointer(G)).CategoryBits
			(*b2TreeNode)(unsafe.Pointer(A)).CategoryBits = (*b2TreeNode)(unsafe.Pointer(C)).CategoryBits | (*b2TreeNode)(unsafe.Pointer(F)).CategoryBits
			v58 = C + 38
			*(*uint16_t)(unsafe.Pointer(v58)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v58))) | (int32FromUint16((*b2TreeNode)(unsafe.Pointer(B)).Flags)|int32FromUint16((*b2TreeNode)(unsafe.Pointer(G)).Flags))&int32(b2_enlargedNode))
			v58 = A + 38
			*(*uint16_t)(unsafe.Pointer(v58)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v58))) | (int32FromUint16((*b2TreeNode)(unsafe.Pointer(C)).Flags)|int32FromUint16((*b2TreeNode)(unsafe.Pointer(F)).Flags))&int32(b2_enlargedNode))
		} else {
			// Swap B and G
			(*b2TreeNode)(unsafe.Pointer(A)).ｆ__ccgo2_24.Children.Child1 = iG
			(*b2TreeNode)(unsafe.Pointer(C)).ｆ__ccgo2_24.Children.Child2 = iB
			(*b2TreeNode)(unsafe.Pointer(B)).ｆ__ccgo3_32.Parent = iC
			(*b2TreeNode)(unsafe.Pointer(G)).ｆ__ccgo3_32.Parent = iA
			(*b2TreeNode)(unsafe.Pointer(C)).Aabb = aabbBF
			(*b2TreeNode)(unsafe.Pointer(C)).Height = uint16FromInt32(int32(1) + int32FromUint16(b2MaxUInt16(tls, (*b2TreeNode)(unsafe.Pointer(B)).Height, (*b2TreeNode)(unsafe.Pointer(F)).Height)))
			(*b2TreeNode)(unsafe.Pointer(A)).Height = uint16FromInt32(int32(1) + int32FromUint16(b2MaxUInt16(tls, (*b2TreeNode)(unsafe.Pointer(C)).Height, (*b2TreeNode)(unsafe.Pointer(G)).Height)))
			(*b2TreeNode)(unsafe.Pointer(C)).CategoryBits = (*b2TreeNode)(unsafe.Pointer(B)).CategoryBits | (*b2TreeNode)(unsafe.Pointer(F)).CategoryBits
			(*b2TreeNode)(unsafe.Pointer(A)).CategoryBits = (*b2TreeNode)(unsafe.Pointer(C)).CategoryBits | (*b2TreeNode)(unsafe.Pointer(G)).CategoryBits
			v58 = C + 38
			*(*uint16_t)(unsafe.Pointer(v58)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v58))) | (int32FromUint16((*b2TreeNode)(unsafe.Pointer(B)).Flags)|int32FromUint16((*b2TreeNode)(unsafe.Pointer(F)).Flags))&int32(b2_enlargedNode))
			v58 = A + 38
			*(*uint16_t)(unsafe.Pointer(v58)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v58))) | (int32FromUint16((*b2TreeNode)(unsafe.Pointer(C)).Flags)|int32FromUint16((*b2TreeNode)(unsafe.Pointer(G)).Flags))&int32(b2_enlargedNode))
		}
	} else {
		if int32FromUint16((*b2TreeNode)(unsafe.Pointer(C)).Height) == 0 {
			// C is a leaf and B is internal
			if !(int32FromUint16((*b2TreeNode)(unsafe.Pointer(B)).Height) > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+5096, __ccgo_ts+4746, int32FromInt32(427)) != 0 {
				__builtin_trap(tls)
			}
			iD = (*b2TreeNode)(unsafe.Pointer(B)).ｆ__ccgo2_24.Children.Child1
			iE = (*b2TreeNode)(unsafe.Pointer(B)).ｆ__ccgo2_24.Children.Child2
			D = nodes + uintptr(iD)*40
			E = nodes + uintptr(iE)*40
			if !(0 <= iD && iD < (*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity) && b2InternalAssertFcn(tls, __ccgo_ts+5110, __ccgo_ts+4746, int32FromInt32(433)) != 0 {
				__builtin_trap(tls)
			}
			if !(0 <= iE && iE < (*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity) && b2InternalAssertFcn(tls, __ccgo_ts+5145, __ccgo_ts+4746, int32FromInt32(434)) != 0 {
				__builtin_trap(tls)
			}
			v1 = (*b2TreeNode)(unsafe.Pointer(B)).Aabb
			wx = v1.UpperBound.X - v1.LowerBound.X
			wy = v1.UpperBound.Y - v1.LowerBound.Y
			v2 = float32(float32FromFloat32(2) * (wx + wy))
			goto _64
		_64:
			// Base cost
			costBase1 = v2
			v4 = (*b2TreeNode)(unsafe.Pointer(C)).Aabb
			v5 = (*b2TreeNode)(unsafe.Pointer(E)).Aabb
			v6 = v4.LowerBound.X
			v7 = v5.LowerBound.X
			if v6 < v7 {
				v10 = v6
			} else {
				v10 = v7
			}
			v8 = v10
			goto _70
		_70:
			c.LowerBound.X = v8
			v11 = v4.LowerBound.Y
			v12 = v5.LowerBound.Y
			if v11 < v12 {
				v15 = v11
			} else {
				v15 = v12
			}
			v13 = v15
			goto _75
		_75:
			c.LowerBound.Y = v13
			v16 = v4.UpperBound.X
			v17 = v5.UpperBound.X
			if v16 > v17 {
				v20 = v16
			} else {
				v20 = v17
			}
			v18 = v20
			goto _80
		_80:
			c.UpperBound.X = v18
			v21 = v4.UpperBound.Y
			v22 = v5.UpperBound.Y
			if v21 > v22 {
				v25 = v21
			} else {
				v25 = v22
			}
			v23 = v25
			goto _85
		_85:
			c.UpperBound.Y = v23
			v26 = c
			goto _88
		_88:
			// Cost of swapping C and D
			aabbCE = v26
			v28 = aabbCE
			wx = v28.UpperBound.X - v28.LowerBound.X
			wy = v28.UpperBound.Y - v28.LowerBound.Y
			v29 = float32(float32FromFloat32(2) * (wx + wy))
			goto _91
		_91:
			costCD = v29
			v31 = (*b2TreeNode)(unsafe.Pointer(C)).Aabb
			v32 = (*b2TreeNode)(unsafe.Pointer(D)).Aabb
			v33 = v31.LowerBound.X
			v34 = v32.LowerBound.X
			if v33 < v34 {
				v37 = v33
			} else {
				v37 = v34
			}
			v35 = v37
			goto _97
		_97:
			c.LowerBound.X = v35
			v38 = v31.LowerBound.Y
			v39 = v32.LowerBound.Y
			if v38 < v39 {
				v42 = v38
			} else {
				v42 = v39
			}
			v40 = v42
			goto _102
		_102:
			c.LowerBound.Y = v40
			v43 = v31.UpperBound.X
			v44 = v32.UpperBound.X
			if v43 > v44 {
				v47 = v43
			} else {
				v47 = v44
			}
			v45 = v47
			goto _107
		_107:
			c.UpperBound.X = v45
			v48 = v31.UpperBound.Y
			v49 = v32.UpperBound.Y
			if v48 > v49 {
				v52 = v48
			} else {
				v52 = v49
			}
			v50 = v52
			goto _112
		_112:
			c.UpperBound.Y = v50
			v53 = c
			goto _115
		_115:
			// Cost of swapping C and E
			aabbCD = v53
			v55 = aabbCD
			wx = v55.UpperBound.X - v55.LowerBound.X
			wy = v55.UpperBound.Y - v55.LowerBound.Y
			v56 = float32(float32FromFloat32(2) * (wx + wy))
			goto _118
		_118:
			costCE = v56
			if costBase1 < costCD && costBase1 < costCE {
				// Rotation does not improve cost
				return
			}
			if costCD < costCE {
				// Swap C and D
				(*b2TreeNode)(unsafe.Pointer(A)).ｆ__ccgo2_24.Children.Child2 = iD
				(*b2TreeNode)(unsafe.Pointer(B)).ｆ__ccgo2_24.Children.Child1 = iC
				(*b2TreeNode)(unsafe.Pointer(C)).ｆ__ccgo3_32.Parent = iB
				(*b2TreeNode)(unsafe.Pointer(D)).ｆ__ccgo3_32.Parent = iA
				(*b2TreeNode)(unsafe.Pointer(B)).Aabb = aabbCE
				(*b2TreeNode)(unsafe.Pointer(B)).Height = uint16FromInt32(int32(1) + int32FromUint16(b2MaxUInt16(tls, (*b2TreeNode)(unsafe.Pointer(C)).Height, (*b2TreeNode)(unsafe.Pointer(E)).Height)))
				(*b2TreeNode)(unsafe.Pointer(A)).Height = uint16FromInt32(int32(1) + int32FromUint16(b2MaxUInt16(tls, (*b2TreeNode)(unsafe.Pointer(B)).Height, (*b2TreeNode)(unsafe.Pointer(D)).Height)))
				(*b2TreeNode)(unsafe.Pointer(B)).CategoryBits = (*b2TreeNode)(unsafe.Pointer(C)).CategoryBits | (*b2TreeNode)(unsafe.Pointer(E)).CategoryBits
				(*b2TreeNode)(unsafe.Pointer(A)).CategoryBits = (*b2TreeNode)(unsafe.Pointer(B)).CategoryBits | (*b2TreeNode)(unsafe.Pointer(D)).CategoryBits
				v58 = B + 38
				*(*uint16_t)(unsafe.Pointer(v58)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v58))) | (int32FromUint16((*b2TreeNode)(unsafe.Pointer(C)).Flags)|int32FromUint16((*b2TreeNode)(unsafe.Pointer(E)).Flags))&int32(b2_enlargedNode))
				v58 = A + 38
				*(*uint16_t)(unsafe.Pointer(v58)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v58))) | (int32FromUint16((*b2TreeNode)(unsafe.Pointer(B)).Flags)|int32FromUint16((*b2TreeNode)(unsafe.Pointer(D)).Flags))&int32(b2_enlargedNode))
			} else {
				// Swap C and E
				(*b2TreeNode)(unsafe.Pointer(A)).ｆ__ccgo2_24.Children.Child2 = iE
				(*b2TreeNode)(unsafe.Pointer(B)).ｆ__ccgo2_24.Children.Child2 = iC
				(*b2TreeNode)(unsafe.Pointer(C)).ｆ__ccgo3_32.Parent = iB
				(*b2TreeNode)(unsafe.Pointer(E)).ｆ__ccgo3_32.Parent = iA
				(*b2TreeNode)(unsafe.Pointer(B)).Aabb = aabbCD
				(*b2TreeNode)(unsafe.Pointer(B)).Height = uint16FromInt32(int32(1) + int32FromUint16(b2MaxUInt16(tls, (*b2TreeNode)(unsafe.Pointer(C)).Height, (*b2TreeNode)(unsafe.Pointer(D)).Height)))
				(*b2TreeNode)(unsafe.Pointer(A)).Height = uint16FromInt32(int32(1) + int32FromUint16(b2MaxUInt16(tls, (*b2TreeNode)(unsafe.Pointer(B)).Height, (*b2TreeNode)(unsafe.Pointer(E)).Height)))
				(*b2TreeNode)(unsafe.Pointer(B)).CategoryBits = (*b2TreeNode)(unsafe.Pointer(C)).CategoryBits | (*b2TreeNode)(unsafe.Pointer(D)).CategoryBits
				(*b2TreeNode)(unsafe.Pointer(A)).CategoryBits = (*b2TreeNode)(unsafe.Pointer(B)).CategoryBits | (*b2TreeNode)(unsafe.Pointer(E)).CategoryBits
				v58 = B + 38
				*(*uint16_t)(unsafe.Pointer(v58)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v58))) | (int32FromUint16((*b2TreeNode)(unsafe.Pointer(C)).Flags)|int32FromUint16((*b2TreeNode)(unsafe.Pointer(D)).Flags))&int32(b2_enlargedNode))
				v58 = A + 38
				*(*uint16_t)(unsafe.Pointer(v58)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v58))) | (int32FromUint16((*b2TreeNode)(unsafe.Pointer(B)).Flags)|int32FromUint16((*b2TreeNode)(unsafe.Pointer(E)).Flags))&int32(b2_enlargedNode))
			}
		} else {
			iD1 = (*b2TreeNode)(unsafe.Pointer(B)).ｆ__ccgo2_24.Children.Child1
			iE1 = (*b2TreeNode)(unsafe.Pointer(B)).ｆ__ccgo2_24.Children.Child2
			iF1 = (*b2TreeNode)(unsafe.Pointer(C)).ｆ__ccgo2_24.Children.Child1
			iG1 = (*b2TreeNode)(unsafe.Pointer(C)).ｆ__ccgo2_24.Children.Child2
			D1 = nodes + uintptr(iD1)*40
			E1 = nodes + uintptr(iE1)*40
			F1 = nodes + uintptr(iF1)*40
			G1 = nodes + uintptr(iG1)*40
			if !(0 <= iD1 && iD1 < (*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity) && b2InternalAssertFcn(tls, __ccgo_ts+5110, __ccgo_ts+4746, int32FromInt32(501)) != 0 {
				__builtin_trap(tls)
			}
			if !(0 <= iE1 && iE1 < (*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity) && b2InternalAssertFcn(tls, __ccgo_ts+5145, __ccgo_ts+4746, int32FromInt32(502)) != 0 {
				__builtin_trap(tls)
			}
			if !(0 <= iF1 && iF1 < (*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity) && b2InternalAssertFcn(tls, __ccgo_ts+5026, __ccgo_ts+4746, int32FromInt32(503)) != 0 {
				__builtin_trap(tls)
			}
			if !(0 <= iG1 && iG1 < (*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity) && b2InternalAssertFcn(tls, __ccgo_ts+5061, __ccgo_ts+4746, int32FromInt32(504)) != 0 {
				__builtin_trap(tls)
			}
			v1 = (*b2TreeNode)(unsafe.Pointer(B)).Aabb
			wx = v1.UpperBound.X - v1.LowerBound.X
			wy = v1.UpperBound.Y - v1.LowerBound.Y
			v2 = float32(float32FromFloat32(2) * (wx + wy))
			goto _125
		_125:
			// Base cost
			areaB = v2
			v4 = (*b2TreeNode)(unsafe.Pointer(C)).Aabb
			wx = v4.UpperBound.X - v4.LowerBound.X
			wy = v4.UpperBound.Y - v4.LowerBound.Y
			v6 = float32(float32FromFloat32(2) * (wx + wy))
			goto _128
		_128:
			areaC = v6
			costBase2 = areaB + areaC
			bestRotation = int32(b2_rotateNone)
			bestCost = costBase2
			v5 = (*b2TreeNode)(unsafe.Pointer(B)).Aabb
			v26 = (*b2TreeNode)(unsafe.Pointer(G1)).Aabb
			v7 = v5.LowerBound.X
			v8 = v26.LowerBound.X
			if v7 < v8 {
				v11 = v7
			} else {
				v11 = v8
			}
			v10 = v11
			goto _134
		_134:
			c.LowerBound.X = v10
			v12 = v5.LowerBound.Y
			v13 = v26.LowerBound.Y
			if v12 < v13 {
				v16 = v12
			} else {
				v16 = v13
			}
			v15 = v16
			goto _139
		_139:
			c.LowerBound.Y = v15
			v17 = v5.UpperBound.X
			v18 = v26.UpperBound.X
			if v17 > v18 {
				v21 = v17
			} else {
				v21 = v18
			}
			v20 = v21
			goto _144
		_144:
			c.UpperBound.X = v20
			v22 = v5.UpperBound.Y
			v23 = v26.UpperBound.Y
			if v22 > v23 {
				v29 = v22
			} else {
				v29 = v23
			}
			v25 = v29
			goto _149
		_149:
			c.UpperBound.Y = v25
			v28 = c
			goto _152
		_152:
			// Cost of swapping B and F
			aabbBG1 = v28
			v31 = aabbBG1
			wx = v31.UpperBound.X - v31.LowerBound.X
			wy = v31.UpperBound.Y - v31.LowerBound.Y
			v33 = float32(float32FromFloat32(2) * (wx + wy))
			goto _155
		_155:
			costBF1 = areaB + v33
			if costBF1 < bestCost {
				bestRotation = int32(b2_rotateBF)
				bestCost = costBF1
			}
			v1 = (*b2TreeNode)(unsafe.Pointer(B)).Aabb
			v4 = (*b2TreeNode)(unsafe.Pointer(F1)).Aabb
			v2 = v1.LowerBound.X
			v6 = v4.LowerBound.X
			if v2 < v6 {
				v8 = v2
			} else {
				v8 = v6
			}
			v7 = v8
			goto _161
		_161:
			c.LowerBound.X = v7
			v10 = v1.LowerBound.Y
			v11 = v4.LowerBound.Y
			if v10 < v11 {
				v13 = v10
			} else {
				v13 = v11
			}
			v12 = v13
			goto _166
		_166:
			c.LowerBound.Y = v12
			v15 = v1.UpperBound.X
			v16 = v4.UpperBound.X
			if v15 > v16 {
				v18 = v15
			} else {
				v18 = v16
			}
			v17 = v18
			goto _171
		_171:
			c.UpperBound.X = v17
			v20 = v1.UpperBound.Y
			v21 = v4.UpperBound.Y
			if v20 > v21 {
				v23 = v20
			} else {
				v23 = v21
			}
			v22 = v23
			goto _176
		_176:
			c.UpperBound.Y = v22
			v5 = c
			goto _179
		_179:
			// Cost of swapping B and G
			aabbBF1 = v5
			v26 = aabbBF1
			wx = v26.UpperBound.X - v26.LowerBound.X
			wy = v26.UpperBound.Y - v26.LowerBound.Y
			v25 = float32(float32FromFloat32(2) * (wx + wy))
			goto _182
		_182:
			costBG1 = areaB + v25
			if costBG1 < bestCost {
				bestRotation = int32(b2_rotateBG)
				bestCost = costBG1
			}
			v1 = (*b2TreeNode)(unsafe.Pointer(C)).Aabb
			v4 = (*b2TreeNode)(unsafe.Pointer(E1)).Aabb
			v2 = v1.LowerBound.X
			v6 = v4.LowerBound.X
			if v2 < v6 {
				v8 = v2
			} else {
				v8 = v6
			}
			v7 = v8
			goto _188
		_188:
			c.LowerBound.X = v7
			v10 = v1.LowerBound.Y
			v11 = v4.LowerBound.Y
			if v10 < v11 {
				v13 = v10
			} else {
				v13 = v11
			}
			v12 = v13
			goto _193
		_193:
			c.LowerBound.Y = v12
			v15 = v1.UpperBound.X
			v16 = v4.UpperBound.X
			if v15 > v16 {
				v18 = v15
			} else {
				v18 = v16
			}
			v17 = v18
			goto _198
		_198:
			c.UpperBound.X = v17
			v20 = v1.UpperBound.Y
			v21 = v4.UpperBound.Y
			if v20 > v21 {
				v23 = v20
			} else {
				v23 = v21
			}
			v22 = v23
			goto _203
		_203:
			c.UpperBound.Y = v22
			v5 = c
			goto _206
		_206:
			// Cost of swapping C and D
			aabbCE1 = v5
			v26 = aabbCE1
			wx = v26.UpperBound.X - v26.LowerBound.X
			wy = v26.UpperBound.Y - v26.LowerBound.Y
			v25 = float32(float32FromFloat32(2) * (wx + wy))
			goto _209
		_209:
			costCD1 = areaC + v25
			if costCD1 < bestCost {
				bestRotation = int32(b2_rotateCD)
				bestCost = costCD1
			}
			v1 = (*b2TreeNode)(unsafe.Pointer(C)).Aabb
			v4 = (*b2TreeNode)(unsafe.Pointer(D1)).Aabb
			v2 = v1.LowerBound.X
			v6 = v4.LowerBound.X
			if v2 < v6 {
				v8 = v2
			} else {
				v8 = v6
			}
			v7 = v8
			goto _215
		_215:
			c.LowerBound.X = v7
			v10 = v1.LowerBound.Y
			v11 = v4.LowerBound.Y
			if v10 < v11 {
				v13 = v10
			} else {
				v13 = v11
			}
			v12 = v13
			goto _220
		_220:
			c.LowerBound.Y = v12
			v15 = v1.UpperBound.X
			v16 = v4.UpperBound.X
			if v15 > v16 {
				v18 = v15
			} else {
				v18 = v16
			}
			v17 = v18
			goto _225
		_225:
			c.UpperBound.X = v17
			v20 = v1.UpperBound.Y
			v21 = v4.UpperBound.Y
			if v20 > v21 {
				v23 = v20
			} else {
				v23 = v21
			}
			v22 = v23
			goto _230
		_230:
			c.UpperBound.Y = v22
			v5 = c
			goto _233
		_233:
			// Cost of swapping C and E
			aabbCD1 = v5
			v26 = aabbCD1
			wx = v26.UpperBound.X - v26.LowerBound.X
			wy = v26.UpperBound.Y - v26.LowerBound.Y
			v25 = float32(float32FromFloat32(2) * (wx + wy))
			goto _236
		_236:
			costCE1 = areaC + v25
			if costCE1 < bestCost {
				bestRotation = int32(b2_rotateCE)
				// bestCost = costCE;
			}
			switch bestRotation {
			case int32(b2_rotateNone):
				goto _237
			case int32(b2_rotateBF):
				goto _238
			case int32(b2_rotateBG):
				goto _239
			case int32(b2_rotateCD):
				goto _240
			case int32(b2_rotateCE):
				goto _241
			default:
				goto _242
			}
		_237:
			;
			goto _243
		_238:
			;
			(*b2TreeNode)(unsafe.Pointer(A)).ｆ__ccgo2_24.Children.Child1 = iF1
			(*b2TreeNode)(unsafe.Pointer(C)).ｆ__ccgo2_24.Children.Child1 = iB
			(*b2TreeNode)(unsafe.Pointer(B)).ｆ__ccgo3_32.Parent = iC
			(*b2TreeNode)(unsafe.Pointer(F1)).ｆ__ccgo3_32.Parent = iA
			(*b2TreeNode)(unsafe.Pointer(C)).Aabb = aabbBG1
			(*b2TreeNode)(unsafe.Pointer(C)).Height = uint16FromInt32(int32(1) + int32FromUint16(b2MaxUInt16(tls, (*b2TreeNode)(unsafe.Pointer(B)).Height, (*b2TreeNode)(unsafe.Pointer(G1)).Height)))
			(*b2TreeNode)(unsafe.Pointer(A)).Height = uint16FromInt32(int32(1) + int32FromUint16(b2MaxUInt16(tls, (*b2TreeNode)(unsafe.Pointer(C)).Height, (*b2TreeNode)(unsafe.Pointer(F1)).Height)))
			(*b2TreeNode)(unsafe.Pointer(C)).CategoryBits = (*b2TreeNode)(unsafe.Pointer(B)).CategoryBits | (*b2TreeNode)(unsafe.Pointer(G1)).CategoryBits
			(*b2TreeNode)(unsafe.Pointer(A)).CategoryBits = (*b2TreeNode)(unsafe.Pointer(C)).CategoryBits | (*b2TreeNode)(unsafe.Pointer(F1)).CategoryBits
			v58 = C + 38
			*(*uint16_t)(unsafe.Pointer(v58)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v58))) | (int32FromUint16((*b2TreeNode)(unsafe.Pointer(B)).Flags)|int32FromUint16((*b2TreeNode)(unsafe.Pointer(G1)).Flags))&int32(b2_enlargedNode))
			v58 = A + 38
			*(*uint16_t)(unsafe.Pointer(v58)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v58))) | (int32FromUint16((*b2TreeNode)(unsafe.Pointer(C)).Flags)|int32FromUint16((*b2TreeNode)(unsafe.Pointer(F1)).Flags))&int32(b2_enlargedNode))
			goto _243
		_239:
			;
			(*b2TreeNode)(unsafe.Pointer(A)).ｆ__ccgo2_24.Children.Child1 = iG1
			(*b2TreeNode)(unsafe.Pointer(C)).ｆ__ccgo2_24.Children.Child2 = iB
			(*b2TreeNode)(unsafe.Pointer(B)).ｆ__ccgo3_32.Parent = iC
			(*b2TreeNode)(unsafe.Pointer(G1)).ｆ__ccgo3_32.Parent = iA
			(*b2TreeNode)(unsafe.Pointer(C)).Aabb = aabbBF1
			(*b2TreeNode)(unsafe.Pointer(C)).Height = uint16FromInt32(int32(1) + int32FromUint16(b2MaxUInt16(tls, (*b2TreeNode)(unsafe.Pointer(B)).Height, (*b2TreeNode)(unsafe.Pointer(F1)).Height)))
			(*b2TreeNode)(unsafe.Pointer(A)).Height = uint16FromInt32(int32(1) + int32FromUint16(b2MaxUInt16(tls, (*b2TreeNode)(unsafe.Pointer(C)).Height, (*b2TreeNode)(unsafe.Pointer(G1)).Height)))
			(*b2TreeNode)(unsafe.Pointer(C)).CategoryBits = (*b2TreeNode)(unsafe.Pointer(B)).CategoryBits | (*b2TreeNode)(unsafe.Pointer(F1)).CategoryBits
			(*b2TreeNode)(unsafe.Pointer(A)).CategoryBits = (*b2TreeNode)(unsafe.Pointer(C)).CategoryBits | (*b2TreeNode)(unsafe.Pointer(G1)).CategoryBits
			v58 = C + 38
			*(*uint16_t)(unsafe.Pointer(v58)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v58))) | (int32FromUint16((*b2TreeNode)(unsafe.Pointer(B)).Flags)|int32FromUint16((*b2TreeNode)(unsafe.Pointer(F1)).Flags))&int32(b2_enlargedNode))
			v58 = A + 38
			*(*uint16_t)(unsafe.Pointer(v58)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v58))) | (int32FromUint16((*b2TreeNode)(unsafe.Pointer(C)).Flags)|int32FromUint16((*b2TreeNode)(unsafe.Pointer(G1)).Flags))&int32(b2_enlargedNode))
			goto _243
		_240:
			;
			(*b2TreeNode)(unsafe.Pointer(A)).ｆ__ccgo2_24.Children.Child2 = iD1
			(*b2TreeNode)(unsafe.Pointer(B)).ｆ__ccgo2_24.Children.Child1 = iC
			(*b2TreeNode)(unsafe.Pointer(C)).ｆ__ccgo3_32.Parent = iB
			(*b2TreeNode)(unsafe.Pointer(D1)).ｆ__ccgo3_32.Parent = iA
			(*b2TreeNode)(unsafe.Pointer(B)).Aabb = aabbCE1
			(*b2TreeNode)(unsafe.Pointer(B)).Height = uint16FromInt32(int32(1) + int32FromUint16(b2MaxUInt16(tls, (*b2TreeNode)(unsafe.Pointer(C)).Height, (*b2TreeNode)(unsafe.Pointer(E1)).Height)))
			(*b2TreeNode)(unsafe.Pointer(A)).Height = uint16FromInt32(int32(1) + int32FromUint16(b2MaxUInt16(tls, (*b2TreeNode)(unsafe.Pointer(B)).Height, (*b2TreeNode)(unsafe.Pointer(D1)).Height)))
			(*b2TreeNode)(unsafe.Pointer(B)).CategoryBits = (*b2TreeNode)(unsafe.Pointer(C)).CategoryBits | (*b2TreeNode)(unsafe.Pointer(E1)).CategoryBits
			(*b2TreeNode)(unsafe.Pointer(A)).CategoryBits = (*b2TreeNode)(unsafe.Pointer(B)).CategoryBits | (*b2TreeNode)(unsafe.Pointer(D1)).CategoryBits
			v58 = B + 38
			*(*uint16_t)(unsafe.Pointer(v58)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v58))) | (int32FromUint16((*b2TreeNode)(unsafe.Pointer(C)).Flags)|int32FromUint16((*b2TreeNode)(unsafe.Pointer(E1)).Flags))&int32(b2_enlargedNode))
			v58 = A + 38
			*(*uint16_t)(unsafe.Pointer(v58)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v58))) | (int32FromUint16((*b2TreeNode)(unsafe.Pointer(B)).Flags)|int32FromUint16((*b2TreeNode)(unsafe.Pointer(D1)).Flags))&int32(b2_enlargedNode))
			goto _243
		_241:
			;
			(*b2TreeNode)(unsafe.Pointer(A)).ｆ__ccgo2_24.Children.Child2 = iE1
			(*b2TreeNode)(unsafe.Pointer(B)).ｆ__ccgo2_24.Children.Child2 = iC
			(*b2TreeNode)(unsafe.Pointer(C)).ｆ__ccgo3_32.Parent = iB
			(*b2TreeNode)(unsafe.Pointer(E1)).ｆ__ccgo3_32.Parent = iA
			(*b2TreeNode)(unsafe.Pointer(B)).Aabb = aabbCD1
			(*b2TreeNode)(unsafe.Pointer(B)).Height = uint16FromInt32(int32(1) + int32FromUint16(b2MaxUInt16(tls, (*b2TreeNode)(unsafe.Pointer(C)).Height, (*b2TreeNode)(unsafe.Pointer(D1)).Height)))
			(*b2TreeNode)(unsafe.Pointer(A)).Height = uint16FromInt32(int32(1) + int32FromUint16(b2MaxUInt16(tls, (*b2TreeNode)(unsafe.Pointer(B)).Height, (*b2TreeNode)(unsafe.Pointer(E1)).Height)))
			(*b2TreeNode)(unsafe.Pointer(B)).CategoryBits = (*b2TreeNode)(unsafe.Pointer(C)).CategoryBits | (*b2TreeNode)(unsafe.Pointer(D1)).CategoryBits
			(*b2TreeNode)(unsafe.Pointer(A)).CategoryBits = (*b2TreeNode)(unsafe.Pointer(B)).CategoryBits | (*b2TreeNode)(unsafe.Pointer(E1)).CategoryBits
			v58 = B + 38
			*(*uint16_t)(unsafe.Pointer(v58)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v58))) | (int32FromUint16((*b2TreeNode)(unsafe.Pointer(C)).Flags)|int32FromUint16((*b2TreeNode)(unsafe.Pointer(D1)).Flags))&int32(b2_enlargedNode))
			v58 = A + 38
			*(*uint16_t)(unsafe.Pointer(v58)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v58))) | (int32FromUint16((*b2TreeNode)(unsafe.Pointer(B)).Flags)|int32FromUint16((*b2TreeNode)(unsafe.Pointer(E1)).Flags))&int32(b2_enlargedNode))
			goto _243
		_242:
			;
		_254:
			;
			if bool(!(int32FromInt32(false1) != 0)) && b2InternalAssertFcn(tls, __ccgo_ts+4123, __ccgo_ts+4746, int32FromInt32(619)) != 0 {
				__builtin_trap(tls)
			}
			goto _253
		_253:
			;
			if 0 != 0 {
				goto _254
			}
			goto _252
		_252:
			;
			goto _243
		_243:
		}
	}
}

func b2InsertLeaf(tls *_Stack, tree uintptr, leaf int32, shouldRotate uint8) {
	var c, leafAABB, v1, v2, v23 AABB
	var child1, child2, index, newParent, oldParent, sibling int32
	var nodes, v49 uintptr
	var v10, v12, v13, v14, v15, v17, v18, v19, v20, v22, v3, v4, v5, v7, v8, v9 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, child1, child2, index, leafAABB, newParent, nodes, oldParent, sibling, v1, v10, v12, v13, v14, v15, v17, v18, v19, v2, v20, v22, v23, v3, v4, v49, v5, v7, v8, v9
	if (*DynamicTree)(unsafe.Pointer(tree)).Root == -int32(1) {
		(*DynamicTree)(unsafe.Pointer(tree)).Root = leaf
		(*(*b2TreeNode)(unsafe.Add(nil, (*DynamicTree)(unsafe.Pointer(tree)).Nodes+uintptr((*DynamicTree)(unsafe.Pointer(tree)).Root)*40))).ｆ__ccgo3_32.Parent = -int32FromInt32(1)
		return
	}
	// Stage 1: find the best sibling for this node
	leafAABB = (*(*b2TreeNode)(unsafe.Add(nil, (*DynamicTree)(unsafe.Pointer(tree)).Nodes+uintptr(leaf)*40))).Aabb
	sibling = b2FindBestSibling(tls, tree, leafAABB)
	// Stage 2: create a new parent for the leaf and sibling
	oldParent = (*(*b2TreeNode)(unsafe.Add(nil, (*DynamicTree)(unsafe.Pointer(tree)).Nodes+uintptr(sibling)*40))).ｆ__ccgo3_32.Parent
	newParent = b2AllocateNode(tls, tree)
	// warning: node pointer can change after allocation
	nodes = (*DynamicTree)(unsafe.Pointer(tree)).Nodes
	(*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(newParent)*40))).ｆ__ccgo3_32.Parent = oldParent
	*(*uint64_t)(unsafe.Add(unsafe.Pointer(&*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(newParent)*40))), 24)) = uint64FromUint64(0xffffffffffffffff)
	v1 = leafAABB
	v2 = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(sibling)*40))).Aabb
	v3 = v1.LowerBound.X
	v4 = v2.LowerBound.X
	if v3 < v4 {
		v7 = v3
	} else {
		v7 = v4
	}
	v5 = v7
	goto _6
_6:
	c.LowerBound.X = v5
	v8 = v1.LowerBound.Y
	v9 = v2.LowerBound.Y
	if v8 < v9 {
		v12 = v8
	} else {
		v12 = v9
	}
	v10 = v12
	goto _11
_11:
	c.LowerBound.Y = v10
	v13 = v1.UpperBound.X
	v14 = v2.UpperBound.X
	if v13 > v14 {
		v17 = v13
	} else {
		v17 = v14
	}
	v15 = v17
	goto _16
_16:
	c.UpperBound.X = v15
	v18 = v1.UpperBound.Y
	v19 = v2.UpperBound.Y
	if v18 > v19 {
		v22 = v18
	} else {
		v22 = v19
	}
	v20 = v22
	goto _21
_21:
	c.UpperBound.Y = v20
	v23 = c
	goto _24
_24:
	(*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(newParent)*40))).Aabb = v23
	(*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(newParent)*40))).CategoryBits = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(leaf)*40))).CategoryBits | (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(sibling)*40))).CategoryBits
	(*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(newParent)*40))).Height = uint16FromInt32(int32FromUint16((*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(sibling)*40))).Height) + int32(1))
	if oldParent != -int32(1) {
		// The sibling was not the root.
		if (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(oldParent)*40))).ｆ__ccgo2_24.Children.Child1 == sibling {
			(*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(oldParent)*40))).ｆ__ccgo2_24.Children.Child1 = newParent
		} else {
			(*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(oldParent)*40))).ｆ__ccgo2_24.Children.Child2 = newParent
		}
		(*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(newParent)*40))).ｆ__ccgo2_24.Children.Child1 = sibling
		(*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(newParent)*40))).ｆ__ccgo2_24.Children.Child2 = leaf
		(*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(sibling)*40))).ｆ__ccgo3_32.Parent = newParent
		(*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(leaf)*40))).ｆ__ccgo3_32.Parent = newParent
	} else {
		// The sibling was the root.
		(*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(newParent)*40))).ｆ__ccgo2_24.Children.Child1 = sibling
		(*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(newParent)*40))).ｆ__ccgo2_24.Children.Child2 = leaf
		(*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(sibling)*40))).ｆ__ccgo3_32.Parent = newParent
		(*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(leaf)*40))).ｆ__ccgo3_32.Parent = newParent
		(*DynamicTree)(unsafe.Pointer(tree)).Root = newParent
	}
	// Stage 3: walk back up the tree fixing heights and AABBs
	index = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(leaf)*40))).ｆ__ccgo3_32.Parent
	for index != -int32(1) {
		child1 = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(index)*40))).ｆ__ccgo2_24.Children.Child1
		child2 = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(index)*40))).ｆ__ccgo2_24.Children.Child2
		if !(child1 != -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+5180, __ccgo_ts+4746, int32FromInt32(684)) != 0 {
			__builtin_trap(tls)
		}
		if !(child2 != -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+5204, __ccgo_ts+4746, int32FromInt32(685)) != 0 {
			__builtin_trap(tls)
		}
		v1 = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(child1)*40))).Aabb
		v2 = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(child2)*40))).Aabb
		v3 = v1.LowerBound.X
		v4 = v2.LowerBound.X
		if v3 < v4 {
			v7 = v3
		} else {
			v7 = v4
		}
		v5 = v7
		goto _30
	_30:
		c.LowerBound.X = v5
		v8 = v1.LowerBound.Y
		v9 = v2.LowerBound.Y
		if v8 < v9 {
			v12 = v8
		} else {
			v12 = v9
		}
		v10 = v12
		goto _35
	_35:
		c.LowerBound.Y = v10
		v13 = v1.UpperBound.X
		v14 = v2.UpperBound.X
		if v13 > v14 {
			v17 = v13
		} else {
			v17 = v14
		}
		v15 = v17
		goto _40
	_40:
		c.UpperBound.X = v15
		v18 = v1.UpperBound.Y
		v19 = v2.UpperBound.Y
		if v18 > v19 {
			v22 = v18
		} else {
			v22 = v19
		}
		v20 = v22
		goto _45
	_45:
		c.UpperBound.Y = v20
		v23 = c
		goto _48
	_48:
		(*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(index)*40))).Aabb = v23
		(*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(index)*40))).CategoryBits = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(child1)*40))).CategoryBits | (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(child2)*40))).CategoryBits
		(*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(index)*40))).Height = uint16FromInt32(int32(1) + int32FromUint16(b2MaxUInt16(tls, (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(child1)*40))).Height, (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(child2)*40))).Height)))
		v49 = nodes + uintptr(index)*40 + 38
		*(*uint16_t)(unsafe.Pointer(v49)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v49))) | (int32FromUint16((*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(child1)*40))).Flags)|int32FromUint16((*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(child2)*40))).Flags))&int32(b2_enlargedNode))
		if shouldRotate != 0 {
			b2RotateNodes(tls, tree, index)
		}
		index = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(index)*40))).ｆ__ccgo3_32.Parent
	}
}

func b2RemoveLeaf(tls *_Stack, tree uintptr, leaf int32) {
	var c, v1, v2, v23 AABB
	var child1, child2, node, nodes uintptr
	var grandParent, index, parent, sibling int32
	var v10, v12, v13, v14, v15, v17, v18, v19, v20, v22, v3, v4, v5, v7, v8, v9 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, child1, child2, grandParent, index, node, nodes, parent, sibling, v1, v10, v12, v13, v14, v15, v17, v18, v19, v2, v20, v22, v23, v3, v4, v5, v7, v8, v9
	if leaf == (*DynamicTree)(unsafe.Pointer(tree)).Root {
		(*DynamicTree)(unsafe.Pointer(tree)).Root = -int32(1)
		return
	}
	nodes = (*DynamicTree)(unsafe.Pointer(tree)).Nodes
	parent = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(leaf)*40))).ｆ__ccgo3_32.Parent
	grandParent = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(parent)*40))).ｆ__ccgo3_32.Parent
	if (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(parent)*40))).ｆ__ccgo2_24.Children.Child1 == leaf {
		sibling = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(parent)*40))).ｆ__ccgo2_24.Children.Child2
	} else {
		sibling = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(parent)*40))).ｆ__ccgo2_24.Children.Child1
	}
	if grandParent != -int32(1) {
		// Destroy parent and connect sibling to grandParent.
		if (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(grandParent)*40))).ｆ__ccgo2_24.Children.Child1 == parent {
			(*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(grandParent)*40))).ｆ__ccgo2_24.Children.Child1 = sibling
		} else {
			(*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(grandParent)*40))).ｆ__ccgo2_24.Children.Child2 = sibling
		}
		(*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(sibling)*40))).ｆ__ccgo3_32.Parent = grandParent
		b2FreeNode(tls, tree, parent)
		// Adjust ancestor bounds.
		index = grandParent
		for index != -int32(1) {
			node = nodes + uintptr(index)*40
			child1 = nodes + uintptr((*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child1)*40
			child2 = nodes + uintptr((*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child2)*40
			// Fast union using SSE
			//__m128 aabb1 = _mm_load_ps(&child1->aabb.lowerBound.x);
			//__m128 aabb2 = _mm_load_ps(&child2->aabb.lowerBound.x);
			//__m128 lower = _mm_min_ps(aabb1, aabb2);
			//__m128 upper = _mm_max_ps(aabb1, aabb2);
			//__m128 aabb = _mm_shuffle_ps(lower, upper, _MM_SHUFFLE(3, 2, 1, 0));
			//_mm_store_ps(&node->aabb.lowerBound.x, aabb);
			v1 = (*b2TreeNode)(unsafe.Pointer(child1)).Aabb
			v2 = (*b2TreeNode)(unsafe.Pointer(child2)).Aabb
			v3 = v1.LowerBound.X
			v4 = v2.LowerBound.X
			if v3 < v4 {
				v7 = v3
			} else {
				v7 = v4
			}
			v5 = v7
			goto _6
		_6:
			c.LowerBound.X = v5
			v8 = v1.LowerBound.Y
			v9 = v2.LowerBound.Y
			if v8 < v9 {
				v12 = v8
			} else {
				v12 = v9
			}
			v10 = v12
			goto _11
		_11:
			c.LowerBound.Y = v10
			v13 = v1.UpperBound.X
			v14 = v2.UpperBound.X
			if v13 > v14 {
				v17 = v13
			} else {
				v17 = v14
			}
			v15 = v17
			goto _16
		_16:
			c.UpperBound.X = v15
			v18 = v1.UpperBound.Y
			v19 = v2.UpperBound.Y
			if v18 > v19 {
				v22 = v18
			} else {
				v22 = v19
			}
			v20 = v22
			goto _21
		_21:
			c.UpperBound.Y = v20
			v23 = c
			goto _24
		_24:
			(*b2TreeNode)(unsafe.Pointer(node)).Aabb = v23
			(*b2TreeNode)(unsafe.Pointer(node)).CategoryBits = (*b2TreeNode)(unsafe.Pointer(child1)).CategoryBits | (*b2TreeNode)(unsafe.Pointer(child2)).CategoryBits
			(*b2TreeNode)(unsafe.Pointer(node)).Height = uint16FromInt32(int32(1) + int32FromUint16(b2MaxUInt16(tls, (*b2TreeNode)(unsafe.Pointer(child1)).Height, (*b2TreeNode)(unsafe.Pointer(child2)).Height)))
			index = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo3_32.Parent
		}
	} else {
		(*DynamicTree)(unsafe.Pointer(tree)).Root = sibling
		(*(*b2TreeNode)(unsafe.Add(nil, (*DynamicTree)(unsafe.Pointer(tree)).Nodes+uintptr(sibling)*40))).ｆ__ccgo3_32.Parent = -int32FromInt32(1)
		b2FreeNode(tls, tree, parent)
	}
}

type __ccgo_fp__Xb2DynamicTree_Query_3 = func(*_Stack, int32, uint64, uintptr) uint8

type __ccgo_fp__Xb2DynamicTree_RayCast_3 = func(*_Stack, uintptr, int32, uint64, uintptr) float32

type __ccgo_fp__Xb2DynamicTree_ShapeCast_3 = func(*_Stack, uintptr, int32, uint64, uintptr) float32

// Median split == 0, Surface area heuristic == 1

// C documentation
//
//	// Median split heuristic
func b2PartitionMid(tls *_Stack, indices uintptr, centers uintptr, count int32) (r int32) {
	var c, c1, c2, d, lowerBound, temp1, temp3, upperBound, v14, v2, v3 Vec2
	var i, i1, i2, temp, temp2 int32
	var pivot, pivot1, v10, v11, v13, v4, v5, v6, v8, v9 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, c1, c2, d, i, i1, i2, lowerBound, pivot, pivot1, temp, temp1, temp2, temp3, upperBound, v10, v11, v13, v14, v2, v3, v4, v5, v6, v8, v9
	// Handle trivial case
	if count <= int32(2) {
		return count / int32(2)
	}
	lowerBound = *(*Vec2)(unsafe.Add(nil, centers))
	upperBound = *(*Vec2)(unsafe.Add(nil, centers))
	i = int32(1)
	for {
		if !(i < count) {
			break
		}
		v2 = lowerBound
		v3 = *(*Vec2)(unsafe.Add(nil, centers+uintptr(i)*8))
		v4 = v2.X
		v5 = v3.X
		if v4 < v5 {
			v8 = v4
		} else {
			v8 = v5
		}
		v6 = v8
		goto _7
	_7:
		c.X = v6
		v9 = v2.Y
		v10 = v3.Y
		if v9 < v10 {
			v13 = v9
		} else {
			v13 = v10
		}
		v11 = v13
		goto _12
	_12:
		c.Y = v11
		v14 = c
		goto _15
	_15:
		lowerBound = v14
		v2 = upperBound
		v3 = *(*Vec2)(unsafe.Add(nil, centers+uintptr(i)*8))
		v4 = v2.X
		v5 = v3.X
		if v4 > v5 {
			v8 = v4
		} else {
			v8 = v5
		}
		v6 = v8
		goto _21
	_21:
		c1.X = v6
		v9 = v2.Y
		v10 = v3.Y
		if v9 > v10 {
			v13 = v9
		} else {
			v13 = v10
		}
		v11 = v13
		goto _26
	_26:
		c1.Y = v11
		v14 = c1
		goto _29
	_29:
		upperBound = v14
		goto _1
	_1:
		;
		i = i + 1
	}
	v2 = upperBound
	v3 = lowerBound
	v14 = Vec2{
		X: v2.X - v3.X,
		Y: v2.Y - v3.Y,
	}
	goto _33
_33:
	d = v14
	c2 = Vec2{
		X: float32(float32FromFloat32(0.5) * (lowerBound.X + upperBound.X)),
		Y: float32(float32FromFloat32(0.5) * (lowerBound.Y + upperBound.Y)),
	}
	// Partition longest axis using the Hoare partition scheme
	// https://en.wikipedia.org/wiki/Quicksort
	// https://nicholasvadivelu.com/2021/01/11/array-partition/
	i1 = 0
	i2 = count
	if d.X > d.Y {
		pivot = c2.X
		for i1 < i2 {
			for i1 < i2 && (*(*Vec2)(unsafe.Add(nil, centers+uintptr(i1)*8))).X < pivot {
				i1 = i1 + int32(1)
			}
			for i1 < i2 && (*(*Vec2)(unsafe.Add(nil, centers+uintptr(i2-int32(1))*8))).X >= pivot {
				i2 = i2 - int32(1)
			}
			if i1 < i2 {
				// Swap indices
				temp = *(*int32)(unsafe.Add(nil, indices+uintptr(i1)*4))
				*(*int32)(unsafe.Add(nil, indices+uintptr(i1)*4)) = *(*int32)(unsafe.Add(nil, indices+uintptr(i2-int32(1))*4))
				*(*int32)(unsafe.Add(nil, indices+uintptr(i2-int32(1))*4)) = temp
				// Swap centers
				temp1 = *(*Vec2)(unsafe.Add(nil, centers+uintptr(i1)*8))
				*(*Vec2)(unsafe.Add(nil, centers+uintptr(i1)*8)) = *(*Vec2)(unsafe.Add(nil, centers+uintptr(i2-int32(1))*8))
				*(*Vec2)(unsafe.Add(nil, centers+uintptr(i2-int32(1))*8)) = temp1
				i1 = i1 + int32(1)
				i2 = i2 - int32(1)
			}
		}
	} else {
		pivot1 = c2.Y
		for i1 < i2 {
			for i1 < i2 && (*(*Vec2)(unsafe.Add(nil, centers+uintptr(i1)*8))).Y < pivot1 {
				i1 = i1 + int32(1)
			}
			for i1 < i2 && (*(*Vec2)(unsafe.Add(nil, centers+uintptr(i2-int32(1))*8))).Y >= pivot1 {
				i2 = i2 - int32(1)
			}
			if i1 < i2 {
				// Swap indices
				temp2 = *(*int32)(unsafe.Add(nil, indices+uintptr(i1)*4))
				*(*int32)(unsafe.Add(nil, indices+uintptr(i1)*4)) = *(*int32)(unsafe.Add(nil, indices+uintptr(i2-int32(1))*4))
				*(*int32)(unsafe.Add(nil, indices+uintptr(i2-int32(1))*4)) = temp2
				// Swap centers
				temp3 = *(*Vec2)(unsafe.Add(nil, centers+uintptr(i1)*8))
				*(*Vec2)(unsafe.Add(nil, centers+uintptr(i1)*8)) = *(*Vec2)(unsafe.Add(nil, centers+uintptr(i2-int32(1))*8))
				*(*Vec2)(unsafe.Add(nil, centers+uintptr(i2-int32(1))*8)) = temp3
				i1 = i1 + int32(1)
				i2 = i2 - int32(1)
			}
		}
	}
	if !(i1 == i2) && b2InternalAssertFcn(tls, __ccgo_ts+5924, __ccgo_ts+4746, int32FromInt32(1524)) != 0 {
		__builtin_trap(tls)
	}
	if i1 > 0 && i1 < count {
		return i1
	}
	return count / int32(2)
}

// Temporary data used to track the rebuild of a tree node
type b2RebuildItem struct {
	NodeIndex  int32
	ChildCount int32
	StartIndex int32
	SplitIndex int32
	EndIndex   int32
}

// C documentation
//
//	// Returns root node index
func b2BuildTree(tls *_Stack, tree uintptr, leafCount int32) (r int32) {
	bp := tls.Alloc(20480)
	defer tls.Free(20480)
	var c, v1, v2, v23 AABB
	var child1, child11, child2, child21, childNode, item, leafCenters, leafIndices, newItem, node, node1, nodes, parentItem, parentNode, rootNode uintptr
	var childIndex, count, endIndex, startIndex, top int32
	var v10, v12, v13, v14, v15, v17, v18, v19, v20, v22, v3, v4, v5, v7, v8, v9 float32
	var _ /* stack at bp+0 */ [1024]b2RebuildItem
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, child1, child11, child2, child21, childIndex, childNode, count, endIndex, item, leafCenters, leafIndices, newItem, node, node1, nodes, parentItem, parentNode, rootNode, startIndex, top, v1, v10, v12, v13, v14, v15, v17, v18, v19, v2, v20, v22, v23, v3, v4, v5, v7, v8, v9
	nodes = (*DynamicTree)(unsafe.Pointer(tree)).Nodes
	leafIndices = (*DynamicTree)(unsafe.Pointer(tree)).LeafIndices
	if leafCount == int32(1) {
		(*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(*(*int32)(unsafe.Add(nil, leafIndices)))*40))).ｆ__ccgo3_32.Parent = -int32FromInt32(1)
		return *(*int32)(unsafe.Add(nil, leafIndices))
	}
	leafCenters = (*DynamicTree)(unsafe.Pointer(tree)).LeafCenters
	top = 0
	(*(*[1024]b2RebuildItem)(unsafe.Add(nil, bp)))[0].NodeIndex = b2AllocateNode(tls, tree)
	(*(*[1024]b2RebuildItem)(unsafe.Add(nil, bp)))[0].ChildCount = -int32(1)
	(*(*[1024]b2RebuildItem)(unsafe.Add(nil, bp)))[0].StartIndex = 0
	(*(*[1024]b2RebuildItem)(unsafe.Add(nil, bp)))[0].EndIndex = leafCount
	(*(*[1024]b2RebuildItem)(unsafe.Add(nil, bp)))[0].SplitIndex = b2PartitionMid(tls, leafIndices, leafCenters, leafCount)
	for int32(true1) != 0 {
		item = bp + uintptr(top)*20
		*(*int32)(unsafe.Add(nil, item+4)) += int32(1)
		if (*b2RebuildItem)(unsafe.Pointer(item)).ChildCount == int32(2) {
			// This internal node has both children established
			if top == 0 {
				// all done
				break
			}
			parentItem = bp + uintptr(top-int32FromInt32(1))*20
			parentNode = nodes + uintptr((*b2RebuildItem)(unsafe.Pointer(parentItem)).NodeIndex)*40
			if (*b2RebuildItem)(unsafe.Pointer(parentItem)).ChildCount == 0 {
				if !((*b2TreeNode)(unsafe.Pointer(parentNode)).ｆ__ccgo2_24.Children.Child1 == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+5933, __ccgo_ts+4746, int32FromInt32(1769)) != 0 {
					__builtin_trap(tls)
				}
				(*b2TreeNode)(unsafe.Pointer(parentNode)).ｆ__ccgo2_24.Children.Child1 = (*b2RebuildItem)(unsafe.Pointer(item)).NodeIndex
			} else {
				if !((*b2RebuildItem)(unsafe.Pointer(parentItem)).ChildCount == int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+5978, __ccgo_ts+4746, int32FromInt32(1774)) != 0 {
					__builtin_trap(tls)
				}
				if !((*b2TreeNode)(unsafe.Pointer(parentNode)).ｆ__ccgo2_24.Children.Child2 == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+6006, __ccgo_ts+4746, int32FromInt32(1775)) != 0 {
					__builtin_trap(tls)
				}
				(*b2TreeNode)(unsafe.Pointer(parentNode)).ｆ__ccgo2_24.Children.Child2 = (*b2RebuildItem)(unsafe.Pointer(item)).NodeIndex
			}
			node = nodes + uintptr((*b2RebuildItem)(unsafe.Pointer(item)).NodeIndex)*40
			if !((*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo3_32.Parent == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+6051, __ccgo_ts+4746, int32FromInt32(1781)) != 0 {
				__builtin_trap(tls)
			}
			(*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo3_32.Parent = (*b2RebuildItem)(unsafe.Pointer(parentItem)).NodeIndex
			if !((*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child1 != -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+6081, __ccgo_ts+4746, int32FromInt32(1784)) != 0 {
				__builtin_trap(tls)
			}
			if !((*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child2 != -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+6120, __ccgo_ts+4746, int32FromInt32(1785)) != 0 {
				__builtin_trap(tls)
			}
			child1 = nodes + uintptr((*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child1)*40
			child2 = nodes + uintptr((*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child2)*40
			v1 = (*b2TreeNode)(unsafe.Pointer(child1)).Aabb
			v2 = (*b2TreeNode)(unsafe.Pointer(child2)).Aabb
			v3 = v1.LowerBound.X
			v4 = v2.LowerBound.X
			if v3 < v4 {
				v7 = v3
			} else {
				v7 = v4
			}
			v5 = v7
			goto _6
		_6:
			c.LowerBound.X = v5
			v8 = v1.LowerBound.Y
			v9 = v2.LowerBound.Y
			if v8 < v9 {
				v12 = v8
			} else {
				v12 = v9
			}
			v10 = v12
			goto _11
		_11:
			c.LowerBound.Y = v10
			v13 = v1.UpperBound.X
			v14 = v2.UpperBound.X
			if v13 > v14 {
				v17 = v13
			} else {
				v17 = v14
			}
			v15 = v17
			goto _16
		_16:
			c.UpperBound.X = v15
			v18 = v1.UpperBound.Y
			v19 = v2.UpperBound.Y
			if v18 > v19 {
				v22 = v18
			} else {
				v22 = v19
			}
			v20 = v22
			goto _21
		_21:
			c.UpperBound.Y = v20
			v23 = c
			goto _24
		_24:
			(*b2TreeNode)(unsafe.Pointer(node)).Aabb = v23
			(*b2TreeNode)(unsafe.Pointer(node)).Height = uint16FromInt32(int32(1) + int32FromUint16(b2MaxUInt16(tls, (*b2TreeNode)(unsafe.Pointer(child1)).Height, (*b2TreeNode)(unsafe.Pointer(child2)).Height)))
			(*b2TreeNode)(unsafe.Pointer(node)).CategoryBits = (*b2TreeNode)(unsafe.Pointer(child1)).CategoryBits | (*b2TreeNode)(unsafe.Pointer(child2)).CategoryBits
			// Pop stack
			top = top - int32(1)
		} else {
			if (*b2RebuildItem)(unsafe.Pointer(item)).ChildCount == 0 {
				startIndex = (*b2RebuildItem)(unsafe.Pointer(item)).StartIndex
				endIndex = (*b2RebuildItem)(unsafe.Pointer(item)).SplitIndex
			} else {
				if !((*b2RebuildItem)(unsafe.Pointer(item)).ChildCount == int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+6159, __ccgo_ts+4746, int32FromInt32(1806)) != 0 {
					__builtin_trap(tls)
				}
				startIndex = (*b2RebuildItem)(unsafe.Pointer(item)).SplitIndex
				endIndex = (*b2RebuildItem)(unsafe.Pointer(item)).EndIndex
			}
			count = endIndex - startIndex
			if count == int32(1) {
				childIndex = *(*int32)(unsafe.Add(nil, leafIndices+uintptr(startIndex)*4))
				node1 = nodes + uintptr((*b2RebuildItem)(unsafe.Pointer(item)).NodeIndex)*40
				if (*b2RebuildItem)(unsafe.Pointer(item)).ChildCount == 0 {
					if !((*b2TreeNode)(unsafe.Pointer(node1)).ｆ__ccgo2_24.Children.Child1 == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+6181, __ccgo_ts+4746, int32FromInt32(1820)) != 0 {
						__builtin_trap(tls)
					}
					(*b2TreeNode)(unsafe.Pointer(node1)).ｆ__ccgo2_24.Children.Child1 = childIndex
				} else {
					if !((*b2RebuildItem)(unsafe.Pointer(item)).ChildCount == int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+6159, __ccgo_ts+4746, int32FromInt32(1825)) != 0 {
						__builtin_trap(tls)
					}
					if !((*b2TreeNode)(unsafe.Pointer(node1)).ｆ__ccgo2_24.Children.Child2 == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+6220, __ccgo_ts+4746, int32FromInt32(1826)) != 0 {
						__builtin_trap(tls)
					}
					(*b2TreeNode)(unsafe.Pointer(node1)).ｆ__ccgo2_24.Children.Child2 = childIndex
				}
				childNode = nodes + uintptr(childIndex)*40
				if !((*b2TreeNode)(unsafe.Pointer(childNode)).ｆ__ccgo3_32.Parent == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+6259, __ccgo_ts+4746, int32FromInt32(1831)) != 0 {
					__builtin_trap(tls)
				}
				(*b2TreeNode)(unsafe.Pointer(childNode)).ｆ__ccgo3_32.Parent = (*b2RebuildItem)(unsafe.Pointer(item)).NodeIndex
			} else {
				if !(count > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+6294, __ccgo_ts+4746, int32FromInt32(1836)) != 0 {
					__builtin_trap(tls)
				}
				if !(top < int32FromInt32(_B2_TREE_STACK_SIZE)) && b2InternalAssertFcn(tls, __ccgo_ts+6304, __ccgo_ts+4746, int32FromInt32(1837)) != 0 {
					__builtin_trap(tls)
				}
				top = top + int32(1)
				newItem = bp + uintptr(top)*20
				(*b2RebuildItem)(unsafe.Pointer(newItem)).NodeIndex = b2AllocateNode(tls, tree)
				(*b2RebuildItem)(unsafe.Pointer(newItem)).ChildCount = -int32(1)
				(*b2RebuildItem)(unsafe.Pointer(newItem)).StartIndex = startIndex
				(*b2RebuildItem)(unsafe.Pointer(newItem)).EndIndex = endIndex
				(*b2RebuildItem)(unsafe.Pointer(newItem)).SplitIndex = b2PartitionMid(tls, leafIndices+uintptr(startIndex)*4, leafCenters+uintptr(startIndex)*8, count)
				*(*int32)(unsafe.Add(nil, newItem+12)) += startIndex
			}
		}
	}
	rootNode = nodes + uintptr((*(*[1024]b2RebuildItem)(unsafe.Add(nil, bp)))[0].NodeIndex)*40
	if !((*b2TreeNode)(unsafe.Pointer(rootNode)).ｆ__ccgo3_32.Parent == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+6329, __ccgo_ts+4746, int32FromInt32(1857)) != 0 {
		__builtin_trap(tls)
	}
	if !((*b2TreeNode)(unsafe.Pointer(rootNode)).ｆ__ccgo2_24.Children.Child1 != -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+6363, __ccgo_ts+4746, int32FromInt32(1858)) != 0 {
		__builtin_trap(tls)
	}
	if !((*b2TreeNode)(unsafe.Pointer(rootNode)).ｆ__ccgo2_24.Children.Child2 != -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+6406, __ccgo_ts+4746, int32FromInt32(1859)) != 0 {
		__builtin_trap(tls)
	}
	child11 = nodes + uintptr((*b2TreeNode)(unsafe.Pointer(rootNode)).ｆ__ccgo2_24.Children.Child1)*40
	child21 = nodes + uintptr((*b2TreeNode)(unsafe.Pointer(rootNode)).ｆ__ccgo2_24.Children.Child2)*40
	v1 = (*b2TreeNode)(unsafe.Pointer(child11)).Aabb
	v2 = (*b2TreeNode)(unsafe.Pointer(child21)).Aabb
	v3 = v1.LowerBound.X
	v4 = v2.LowerBound.X
	if v3 < v4 {
		v7 = v3
	} else {
		v7 = v4
	}
	v5 = v7
	goto _30
_30:
	c.LowerBound.X = v5
	v8 = v1.LowerBound.Y
	v9 = v2.LowerBound.Y
	if v8 < v9 {
		v12 = v8
	} else {
		v12 = v9
	}
	v10 = v12
	goto _35
_35:
	c.LowerBound.Y = v10
	v13 = v1.UpperBound.X
	v14 = v2.UpperBound.X
	if v13 > v14 {
		v17 = v13
	} else {
		v17 = v14
	}
	v15 = v17
	goto _40
_40:
	c.UpperBound.X = v15
	v18 = v1.UpperBound.Y
	v19 = v2.UpperBound.Y
	if v18 > v19 {
		v22 = v18
	} else {
		v22 = v19
	}
	v20 = v22
	goto _45
_45:
	c.UpperBound.Y = v20
	v23 = c
	goto _48
_48:
	(*b2TreeNode)(unsafe.Pointer(rootNode)).Aabb = v23
	(*b2TreeNode)(unsafe.Pointer(rootNode)).Height = uint16FromInt32(int32(1) + int32FromUint16(b2MaxUInt16(tls, (*b2TreeNode)(unsafe.Pointer(child11)).Height, (*b2TreeNode)(unsafe.Pointer(child21)).Height)))
	(*b2TreeNode)(unsafe.Pointer(rootNode)).CategoryBits = (*b2TreeNode)(unsafe.Pointer(child11)).CategoryBits | (*b2TreeNode)(unsafe.Pointer(child21)).CategoryBits
	return (*(*[1024]b2RebuildItem)(unsafe.Add(nil, bp)))[0].NodeIndex
}

const _FLT_MAX3 = 3.40282346638528859812e+38
const _UINT64_MAX4 = "0xffffffffffffffffu"

func b2IsValidRay(tls *_Stack, input uintptr) (r uint8) {
	var isValid uint8
	_ = isValid
	isValid = boolUint8(b2IsValidVec2(tls, (*RayCastInput)(unsafe.Pointer(input)).Origin) != 0 && b2IsValidVec2(tls, (*RayCastInput)(unsafe.Pointer(input)).Translation) != 0 && b2IsValidFloat(tls, (*RayCastInput)(unsafe.Pointer(input)).MaxFraction) != 0 && float32FromFloat32(0) <= (*RayCastInput)(unsafe.Pointer(input)).MaxFraction && (*RayCastInput)(unsafe.Pointer(input)).MaxFraction < float32(float32FromFloat32(100000)*b2_lengthUnitsPerMeter))
	return isValid
}

func b2MakePolygon(tls *_Stack, hull uintptr, radius float32) (r Polygon) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var edge, n, v4, v5, v6, v8, v9 Vec2
	var i, i1, i11, i2, v3 int32
	var invLength, length, v10 float32
	var _ /* shape at bp+0 */ Polygon
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = edge, i, i1, i11, i2, invLength, length, n, v10, v3, v4, v5, v6, v8, v9
	if !(b2ValidateHull(tls, hull) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+6550, __ccgo_ts+6524, int32FromInt32(58)) != 0 {
		__builtin_trap(tls)
	}
	if (*Hull)(unsafe.Pointer(hull)).Count < int32(3) {
		// Handle a bad hull when assertions are disabled
		return b2MakeSquare(tls, float32FromFloat32(0.5))
	}
	*(*Polygon)(unsafe.Add(nil, bp)) = Polygon{}
	(*(*Polygon)(unsafe.Add(nil, bp))).Count = (*Hull)(unsafe.Pointer(hull)).Count
	(*(*Polygon)(unsafe.Add(nil, bp))).Radius = radius
	// Copy vertices
	i = 0
	for {
		if !(i < (*(*Polygon)(unsafe.Add(nil, bp))).Count) {
			break
		}
		*(*Vec2)(unsafe.Add(nil, bp+uintptr(i)*8)) = *(*Vec2)(unsafe.Add(nil, hull+uintptr(i)*8))
		goto _1
	_1:
		;
		i = i + 1
	}
	// Compute normals. Ensure the edges have non-zero length.
	i1 = 0
	for {
		if !(i1 < (*(*Polygon)(unsafe.Add(nil, bp))).Count) {
			break
		}
		i11 = i1
		if i1+int32(1) < (*(*Polygon)(unsafe.Add(nil, bp))).Count {
			v3 = i1 + int32(1)
		} else {
			v3 = 0
		}
		i2 = v3
		v4 = *(*Vec2)(unsafe.Add(nil, bp+uintptr(i2)*8))
		v5 = *(*Vec2)(unsafe.Add(nil, bp+uintptr(i11)*8))
		v6 = Vec2{
			X: v4.X - v5.X,
			Y: v4.Y - v5.Y,
		}
		goto _7
	_7:
		edge = v6
		v8 = edge
		v9 = edge
		v10 = float32(v8.X*v9.X) + float32(v8.Y*v9.Y)
		goto _11
	_11:
		;
		if !(v10 > float32(float32FromFloat32(1.1920928955078125e-07)*float32FromFloat32(1.1920928955078125e-07))) && b2InternalAssertFcn(tls, __ccgo_ts+6573, __ccgo_ts+6524, int32FromInt32(82)) != 0 {
			__builtin_trap(tls)
		}
		v4 = edge
		v10 = float32FromFloat32(1)
		v5 = Vec2{
			X: float32(v10 * v4.Y),
			Y: float32(-v10 * v4.X),
		}
		goto _15
	_15:
		v6 = v5
		length = sqrtf(tls, float32(v6.X*v6.X)+float32(v6.Y*v6.Y))
		if length < float32FromFloat32(1.1920928955078125e-07) {
			v8 = Vec2{}
			goto _18
		}
		invLength = float32FromFloat32(1) / length
		n = Vec2{
			X: float32(invLength * v6.X),
			Y: float32(invLength * v6.Y),
		}
		v8 = n
		goto _18
	_18:
		*(*Vec2)(unsafe.Add(nil, bp+64+uintptr(i1)*8)) = v8
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	(*(*Polygon)(unsafe.Add(nil, bp))).Centroid = b2ComputePolygonCentroid(tls, bp, (*(*Polygon)(unsafe.Add(nil, bp))).Count)
	return *(*Polygon)(unsafe.Add(nil, bp))
}

func b2MakeOffsetPolygon(tls *_Stack, hull uintptr, position Vec2, rotation Rot) (r Polygon) {
	return b2MakeOffsetRoundedPolygon(tls, hull, position, rotation, float32FromFloat32(0))
}

func b2MakeOffsetRoundedPolygon(tls *_Stack, hull uintptr, position Vec2, rotation Rot, radius float32) (r Polygon) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var edge, n, v10, v3, v4, v8, v9 Vec2
	var i, i1, i11, i2, v7 int32
	var invLength, length, x, y, v14 float32
	var transform, v2 Transform
	var _ /* shape at bp+0 */ Polygon
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = edge, i, i1, i11, i2, invLength, length, n, transform, x, y, v10, v14, v2, v3, v4, v7, v8, v9
	if !(b2ValidateHull(tls, hull) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+6550, __ccgo_ts+6524, int32FromInt32(98)) != 0 {
		__builtin_trap(tls)
	}
	if (*Hull)(unsafe.Pointer(hull)).Count < int32(3) {
		// Handle a bad hull when assertions are disabled
		return b2MakeSquare(tls, float32FromFloat32(0.5))
	}
	transform = Transform{
		P: position,
		Q: rotation,
	}
	*(*Polygon)(unsafe.Add(nil, bp)) = Polygon{}
	(*(*Polygon)(unsafe.Add(nil, bp))).Count = (*Hull)(unsafe.Pointer(hull)).Count
	(*(*Polygon)(unsafe.Add(nil, bp))).Radius = radius
	// Copy vertices
	i = 0
	for {
		if !(i < (*(*Polygon)(unsafe.Add(nil, bp))).Count) {
			break
		}
		v2 = transform
		v3 = *(*Vec2)(unsafe.Add(nil, hull+uintptr(i)*8))
		x = float32(v2.Q.C*v3.X) - float32(v2.Q.S*v3.Y) + v2.P.X
		y = float32(v2.Q.S*v3.X) + float32(v2.Q.C*v3.Y) + v2.P.Y
		v4 = Vec2{
			X: x,
			Y: y,
		}
		goto _5
	_5:
		*(*Vec2)(unsafe.Add(nil, bp+uintptr(i)*8)) = v4
		goto _1
	_1:
		;
		i = i + 1
	}
	// Compute normals. Ensure the edges have non-zero length.
	i1 = 0
	for {
		if !(i1 < (*(*Polygon)(unsafe.Add(nil, bp))).Count) {
			break
		}
		i11 = i1
		if i1+int32(1) < (*(*Polygon)(unsafe.Add(nil, bp))).Count {
			v7 = i1 + int32(1)
		} else {
			v7 = 0
		}
		i2 = v7
		v3 = *(*Vec2)(unsafe.Add(nil, bp+uintptr(i2)*8))
		v4 = *(*Vec2)(unsafe.Add(nil, bp+uintptr(i11)*8))
		v8 = Vec2{
			X: v3.X - v4.X,
			Y: v3.Y - v4.Y,
		}
		goto _11
	_11:
		edge = v8
		v9 = edge
		v10 = edge
		v14 = float32(v9.X*v10.X) + float32(v9.Y*v10.Y)
		goto _15
	_15:
		;
		if !(v14 > float32(float32FromFloat32(1.1920928955078125e-07)*float32FromFloat32(1.1920928955078125e-07))) && b2InternalAssertFcn(tls, __ccgo_ts+6573, __ccgo_ts+6524, int32FromInt32(124)) != 0 {
			__builtin_trap(tls)
		}
		v3 = edge
		v14 = float32FromFloat32(1)
		v4 = Vec2{
			X: float32(v14 * v3.Y),
			Y: float32(-v14 * v3.X),
		}
		goto _19
	_19:
		v8 = v4
		length = sqrtf(tls, float32(v8.X*v8.X)+float32(v8.Y*v8.Y))
		if length < float32FromFloat32(1.1920928955078125e-07) {
			v9 = Vec2{}
			goto _22
		}
		invLength = float32FromFloat32(1) / length
		n = Vec2{
			X: float32(invLength * v8.X),
			Y: float32(invLength * v8.Y),
		}
		v9 = n
		goto _22
	_22:
		*(*Vec2)(unsafe.Add(nil, bp+64+uintptr(i1)*8)) = v9
		goto _6
	_6:
		;
		i1 = i1 + 1
	}
	(*(*Polygon)(unsafe.Add(nil, bp))).Centroid = b2ComputePolygonCentroid(tls, bp, (*(*Polygon)(unsafe.Add(nil, bp))).Count)
	return *(*Polygon)(unsafe.Add(nil, bp))
}

func b2MakeSquare(tls *_Stack, halfWidth float32) (r Polygon) {
	return b2MakeBox(tls, halfWidth, halfWidth)
}

func b2MakeBox(tls *_Stack, halfWidth float32, halfHeight float32) (r Polygon) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var _ /* shape at bp+0 */ Polygon
	if !(b2IsValidFloat(tls, halfWidth) != 0 && halfWidth > float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+6621, __ccgo_ts+6524, int32FromInt32(140)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2IsValidFloat(tls, halfHeight) != 0 && halfHeight > float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+6669, __ccgo_ts+6524, int32FromInt32(141)) != 0 {
		__builtin_trap(tls)
	}
	*(*Polygon)(unsafe.Add(nil, bp)) = Polygon{}
	(*(*Polygon)(unsafe.Add(nil, bp))).Count = int32(4)
	*(*Vec2)(unsafe.Add(nil, bp)) = Vec2{
		X: -halfWidth,
		Y: -halfHeight,
	}
	*(*Vec2)(unsafe.Add(nil, bp+1*8)) = Vec2{
		X: halfWidth,
		Y: -halfHeight,
	}
	*(*Vec2)(unsafe.Add(nil, bp+2*8)) = Vec2{
		X: halfWidth,
		Y: halfHeight,
	}
	*(*Vec2)(unsafe.Add(nil, bp+3*8)) = Vec2{
		X: -halfWidth,
		Y: halfHeight,
	}
	*(*Vec2)(unsafe.Add(nil, bp+64)) = Vec2{
		Y: -float32FromFloat32(1),
	}
	*(*Vec2)(unsafe.Add(nil, bp+64+1*8)) = Vec2{
		X: float32FromFloat32(1),
	}
	*(*Vec2)(unsafe.Add(nil, bp+64+2*8)) = Vec2{
		Y: float32FromFloat32(1),
	}
	*(*Vec2)(unsafe.Add(nil, bp+64+3*8)) = Vec2{
		X: -float32FromFloat32(1),
	}
	(*(*Polygon)(unsafe.Add(nil, bp))).Radius = float32FromFloat32(0)
	(*(*Polygon)(unsafe.Add(nil, bp))).Centroid = b2Vec2_zero
	return *(*Polygon)(unsafe.Add(nil, bp))
}

func b2MakeRoundedBox(tls *_Stack, halfWidth float32, halfHeight float32, radius float32) (r Polygon) {
	var shape Polygon
	_ = shape
	if !(b2IsValidFloat(tls, radius) != 0 && radius >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+6719, __ccgo_ts+6524, int32FromInt32(160)) != 0 {
		__builtin_trap(tls)
	}
	shape = b2MakeBox(tls, halfWidth, halfHeight)
	shape.Radius = radius
	return shape
}

func b2MakeOffsetBox(tls *_Stack, halfWidth float32, halfHeight float32, center Vec2, rotation Rot) (r Polygon) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var x, y float32
	var xf, v1 Transform
	var v2, v3 Vec2
	var v17 Rot
	var _ /* shape at bp+0 */ Polygon
	_, _, _, _, _, _, _ = x, xf, y, v1, v17, v2, v3
	xf = Transform{
		P: center,
		Q: rotation,
	}
	*(*Polygon)(unsafe.Add(nil, bp)) = Polygon{}
	(*(*Polygon)(unsafe.Add(nil, bp))).Count = int32(4)
	v1 = xf
	v2 = Vec2{
		X: -halfWidth,
		Y: -halfHeight,
	}
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	*(*Vec2)(unsafe.Add(nil, bp)) = v3
	v1 = xf
	v2 = Vec2{
		X: halfWidth,
		Y: -halfHeight,
	}
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _8
_8:
	*(*Vec2)(unsafe.Add(nil, bp+1*8)) = v3
	v1 = xf
	v2 = Vec2{
		X: halfWidth,
		Y: halfHeight,
	}
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _12
_12:
	*(*Vec2)(unsafe.Add(nil, bp+2*8)) = v3
	v1 = xf
	v2 = Vec2{
		X: -halfWidth,
		Y: halfHeight,
	}
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _16
_16:
	*(*Vec2)(unsafe.Add(nil, bp+3*8)) = v3
	v17 = xf.Q
	v2 = Vec2{
		Y: -float32FromFloat32(1),
	}
	v3 = Vec2{
		X: float32(v17.C*v2.X) - float32(v17.S*v2.Y),
		Y: float32(v17.S*v2.X) + float32(v17.C*v2.Y),
	}
	goto _20
_20:
	*(*Vec2)(unsafe.Add(nil, bp+64)) = v3
	v17 = xf.Q
	v2 = Vec2{
		X: float32FromFloat32(1),
	}
	v3 = Vec2{
		X: float32(v17.C*v2.X) - float32(v17.S*v2.Y),
		Y: float32(v17.S*v2.X) + float32(v17.C*v2.Y),
	}
	goto _24
_24:
	*(*Vec2)(unsafe.Add(nil, bp+64+1*8)) = v3
	v17 = xf.Q
	v2 = Vec2{
		Y: float32FromFloat32(1),
	}
	v3 = Vec2{
		X: float32(v17.C*v2.X) - float32(v17.S*v2.Y),
		Y: float32(v17.S*v2.X) + float32(v17.C*v2.Y),
	}
	goto _28
_28:
	*(*Vec2)(unsafe.Add(nil, bp+64+2*8)) = v3
	v17 = xf.Q
	v2 = Vec2{
		X: -float32FromFloat32(1),
	}
	v3 = Vec2{
		X: float32(v17.C*v2.X) - float32(v17.S*v2.Y),
		Y: float32(v17.S*v2.X) + float32(v17.C*v2.Y),
	}
	goto _32
_32:
	*(*Vec2)(unsafe.Add(nil, bp+64+3*8)) = v3
	(*(*Polygon)(unsafe.Add(nil, bp))).Radius = float32FromFloat32(0)
	(*(*Polygon)(unsafe.Add(nil, bp))).Centroid = xf.P
	return *(*Polygon)(unsafe.Add(nil, bp))
}

func b2MakeOffsetRoundedBox(tls *_Stack, halfWidth float32, halfHeight float32, center Vec2, rotation Rot, radius float32) (r Polygon) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var x, y float32
	var xf, v1 Transform
	var v2, v3 Vec2
	var v17 Rot
	var _ /* shape at bp+0 */ Polygon
	_, _, _, _, _, _, _ = x, xf, y, v1, v17, v2, v3
	if !(b2IsValidFloat(tls, radius) != 0 && radius >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+6719, __ccgo_ts+6524, int32FromInt32(187)) != 0 {
		__builtin_trap(tls)
	}
	xf = Transform{
		P: center,
		Q: rotation,
	}
	*(*Polygon)(unsafe.Add(nil, bp)) = Polygon{}
	(*(*Polygon)(unsafe.Add(nil, bp))).Count = int32(4)
	v1 = xf
	v2 = Vec2{
		X: -halfWidth,
		Y: -halfHeight,
	}
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	*(*Vec2)(unsafe.Add(nil, bp)) = v3
	v1 = xf
	v2 = Vec2{
		X: halfWidth,
		Y: -halfHeight,
	}
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _8
_8:
	*(*Vec2)(unsafe.Add(nil, bp+1*8)) = v3
	v1 = xf
	v2 = Vec2{
		X: halfWidth,
		Y: halfHeight,
	}
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _12
_12:
	*(*Vec2)(unsafe.Add(nil, bp+2*8)) = v3
	v1 = xf
	v2 = Vec2{
		X: -halfWidth,
		Y: halfHeight,
	}
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _16
_16:
	*(*Vec2)(unsafe.Add(nil, bp+3*8)) = v3
	v17 = xf.Q
	v2 = Vec2{
		Y: -float32FromFloat32(1),
	}
	v3 = Vec2{
		X: float32(v17.C*v2.X) - float32(v17.S*v2.Y),
		Y: float32(v17.S*v2.X) + float32(v17.C*v2.Y),
	}
	goto _20
_20:
	*(*Vec2)(unsafe.Add(nil, bp+64)) = v3
	v17 = xf.Q
	v2 = Vec2{
		X: float32FromFloat32(1),
	}
	v3 = Vec2{
		X: float32(v17.C*v2.X) - float32(v17.S*v2.Y),
		Y: float32(v17.S*v2.X) + float32(v17.C*v2.Y),
	}
	goto _24
_24:
	*(*Vec2)(unsafe.Add(nil, bp+64+1*8)) = v3
	v17 = xf.Q
	v2 = Vec2{
		Y: float32FromFloat32(1),
	}
	v3 = Vec2{
		X: float32(v17.C*v2.X) - float32(v17.S*v2.Y),
		Y: float32(v17.S*v2.X) + float32(v17.C*v2.Y),
	}
	goto _28
_28:
	*(*Vec2)(unsafe.Add(nil, bp+64+2*8)) = v3
	v17 = xf.Q
	v2 = Vec2{
		X: -float32FromFloat32(1),
	}
	v3 = Vec2{
		X: float32(v17.C*v2.X) - float32(v17.S*v2.Y),
		Y: float32(v17.S*v2.X) + float32(v17.C*v2.Y),
	}
	goto _32
_32:
	*(*Vec2)(unsafe.Add(nil, bp+64+3*8)) = v3
	(*(*Polygon)(unsafe.Add(nil, bp))).Radius = radius
	(*(*Polygon)(unsafe.Add(nil, bp))).Centroid = xf.P
	return *(*Polygon)(unsafe.Add(nil, bp))
}

func b2TransformPolygon(tls *_Stack, transform Transform, polygon uintptr) (r Polygon) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var i int32
	var x, y float32
	var v2 Transform
	var v3, v4 Vec2
	var v6 Rot
	var _ /* p at bp+0 */ Polygon
	_, _, _, _, _, _, _ = i, x, y, v2, v3, v4, v6
	*(*Polygon)(unsafe.Add(nil, bp)) = *(*Polygon)(unsafe.Add(nil, polygon))
	i = 0
	for {
		if !(i < (*(*Polygon)(unsafe.Add(nil, bp))).Count) {
			break
		}
		v2 = transform
		v3 = *(*Vec2)(unsafe.Add(nil, bp+uintptr(i)*8))
		x = float32(v2.Q.C*v3.X) - float32(v2.Q.S*v3.Y) + v2.P.X
		y = float32(v2.Q.S*v3.X) + float32(v2.Q.C*v3.Y) + v2.P.Y
		v4 = Vec2{
			X: x,
			Y: y,
		}
		goto _5
	_5:
		*(*Vec2)(unsafe.Add(nil, bp+uintptr(i)*8)) = v4
		v6 = transform.Q
		v3 = *(*Vec2)(unsafe.Add(nil, bp+64+uintptr(i)*8))
		v4 = Vec2{
			X: float32(v6.C*v3.X) - float32(v6.S*v3.Y),
			Y: float32(v6.S*v3.X) + float32(v6.C*v3.Y),
		}
		goto _9
	_9:
		*(*Vec2)(unsafe.Add(nil, bp+64+uintptr(i)*8)) = v4
		goto _1
	_1:
		;
		i = i + 1
	}
	v2 = transform
	v3 = (*(*Polygon)(unsafe.Add(nil, bp))).Centroid
	x = float32(v2.Q.C*v3.X) - float32(v2.Q.S*v3.Y) + v2.P.X
	y = float32(v2.Q.S*v3.X) + float32(v2.Q.C*v3.Y) + v2.P.Y
	v4 = Vec2{
		X: x,
		Y: y,
	}
	goto _13
_13:
	(*(*Polygon)(unsafe.Add(nil, bp))).Centroid = v4
	return *(*Polygon)(unsafe.Add(nil, bp))
}

func b2PointInCircle(tls *_Stack, point Vec2, shape uintptr) (r uint8) {
	var c, center, v1, v2 Vec2
	var v3 float32
	_, _, _, _, _ = c, center, v1, v2, v3
	center = (*Circle)(unsafe.Pointer(shape)).Center
	v1 = point
	v2 = center
	c = Vec2{
		X: v2.X - v1.X,
		Y: v2.Y - v1.Y,
	}
	v3 = float32(c.X*c.X) + float32(c.Y*c.Y)
	goto _4
_4:
	return boolUint8(v3 <= float32((*Circle)(unsafe.Pointer(shape)).Radius*(*Circle)(unsafe.Pointer(shape)).Radius))
}

func b2PointInCapsule(tls *_Stack, point Vec2, shape uintptr) (r uint8) {
	var c, c1, d, p1, p2, v1, v10, v2, v3, v5, v6, v9 Vec2
	var dd, rr, t, v11, v19, v21, v22, v23, v24, v7 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, c1, d, dd, p1, p2, rr, t, v1, v10, v11, v19, v2, v21, v22, v23, v24, v3, v5, v6, v7, v9
	rr = float32((*Capsule)(unsafe.Pointer(shape)).Radius * (*Capsule)(unsafe.Pointer(shape)).Radius)
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
	d = v3
	v5 = d
	v6 = d
	v7 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y)
	goto _8
_8:
	dd = v7
	if dd == float32FromFloat32(0) {
		// Capsule is really a circle
		v9 = point
		v10 = p1
		c = Vec2{
			X: v10.X - v9.X,
			Y: v10.Y - v9.Y,
		}
		v11 = float32(c.X*c.X) + float32(c.Y*c.Y)
		goto _12
	_12:
		return boolUint8(v11 <= rr)
	}
	v1 = point
	v2 = p1
	v3 = Vec2{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
	goto _16
_16:
	v5 = v3
	v6 = d
	v7 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y)
	goto _20
_20:
	// Get closest point on capsule segment
	// c = p1 + t * d
	// dot(point - c, d) = 0
	// dot(point - p1 - t * d, d) = 0
	// t = dot(point - p1, d) / dot(d, d)
	t = v7 / dd
	v11 = t
	v19 = float32FromFloat32(0)
	v21 = float32FromFloat32(1)
	if v11 < v19 {
		v23 = v19
	} else {
		if v11 > v21 {
			v24 = v21
		} else {
			v24 = v11
		}
		v23 = v24
	}
	v22 = v23
	goto _25
_25:
	t = v22
	v1 = p1
	v7 = t
	v2 = d
	v3 = Vec2{
		X: v1.X + float32(v7*v2.X),
		Y: v1.Y + float32(v7*v2.Y),
	}
	goto _32
_32:
	c1 = v3
	// Is query point within radius around closest point?
	v5 = point
	v6 = c1
	c = Vec2{
		X: v6.X - v5.X,
		Y: v6.Y - v5.Y,
	}
	v11 = float32(c.X*c.X) + float32(c.Y*c.Y)
	goto _36
_36:
	return boolUint8(v11 <= rr)
}

func b2PointInPolygon(tls *_Stack, _point Vec2, shape uintptr) (r uint8) {
	bp := tls.Alloc(208)
	defer tls.Free(208)
	*(*Vec2)(unsafe.Pointer(bp)) = _point
	var output DistanceOutput
	var _ /* cache at bp+192 */ SimplexCache
	var _ /* input at bp+8 */ DistanceInput
	_ = output
	*(*DistanceInput)(unsafe.Add(nil, bp+8)) = DistanceInput{}
	(*(*DistanceInput)(unsafe.Add(nil, bp+8))).ProxyA = b2MakeProxy(tls, shape, (*Polygon)(unsafe.Pointer(shape)).Count, float32FromFloat32(0))
	(*(*DistanceInput)(unsafe.Add(nil, bp+8))).ProxyB = b2MakeProxy(tls, bp, int32(1), float32FromFloat32(0))
	(*(*DistanceInput)(unsafe.Add(nil, bp+8))).TransformA = b2Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp+8))).TransformB = b2Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp+8))).UseRadii = boolUint8(false1 != 0)
	*(*SimplexCache)(unsafe.Add(nil, bp+192)) = SimplexCache{}
	output = b2ShapeDistance(tls, bp+8, bp+192, uintptrFromInt32(0), 0)
	return boolUint8(output.Distance <= (*Polygon)(unsafe.Pointer(shape)).Radius)
}

// C documentation
//
//	// Precision Improvements for Ray / Sphere Intersection - Ray Tracing Gems 2019
//	// http://www.codercorner.com/blog/?p=321
func b2RayCastCircle(tls *_Stack, input uintptr, shape uintptr) (r1 CastOutput) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var c, d, hitPoint, n, n1, p, s1, v1, v12, v2, v3, v6, v7, v9 Vec2
	var cc, fraction, h, invLength, invLength1, length, r, rr, t, v10, v14, v17 float32
	var output CastOutput
	var v5 uintptr
	var _ /* length at bp+0 */ float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, cc, d, fraction, h, hitPoint, invLength, invLength1, length, n, n1, output, p, r, rr, s1, t, v1, v10, v12, v14, v17, v2, v3, v5, v6, v7, v9
	if !(b2IsValidRay(tls, input) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+6779, __ccgo_ts+6524, int32FromInt32(508)) != 0 {
		__builtin_trap(tls)
	}
	p = (*Circle)(unsafe.Pointer(shape)).Center
	output = CastOutput{}
	v1 = (*RayCastInput)(unsafe.Pointer(input)).Origin
	v2 = p
	v3 = Vec2{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
	goto _4
_4:
	// Shift ray so circle center is the origin
	s1 = v3
	r = (*Circle)(unsafe.Pointer(shape)).Radius
	rr = float32(r * r)
	v5 = bp
	v6 = (*RayCastInput)(unsafe.Pointer(input)).Translation
	*(*float32)(unsafe.Add(nil, v5)) = sqrtf(tls, float32(v6.X*v6.X)+float32(v6.Y*v6.Y))
	if *(*float32)(unsafe.Add(nil, v5)) < float32FromFloat32(1.1920928955078125e-07) {
		v7 = Vec2{}
		goto _8
	}
	invLength1 = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v5))
	n1 = Vec2{
		X: float32(invLength1 * v6.X),
		Y: float32(invLength1 * v6.Y),
	}
	v7 = n1
	goto _8
_8:
	d = v7
	if *(*float32)(unsafe.Add(nil, bp)) == float32FromFloat32(0) {
		// zero length ray
		v9 = s1
		v10 = float32(v9.X*v9.X) + float32(v9.Y*v9.Y)
		goto _11
	_11:
		if v10 < rr {
			// initial overlap
			output.Point = (*RayCastInput)(unsafe.Pointer(input)).Origin
			output.Hit = boolUint8(true1 != 0)
		}
		return output
	}
	v1 = s1
	v2 = d
	v10 = float32(v1.X*v2.X) + float32(v1.Y*v2.Y)
	goto _15
_15:
	// Find closest point on ray to origin
	// solve: dot(s + t * d, d) = 0
	t = -v10
	v3 = s1
	v14 = t
	v6 = d
	v7 = Vec2{
		X: v3.X + float32(v14*v6.X),
		Y: v3.Y + float32(v14*v6.Y),
	}
	goto _20
_20:
	// c is the closest point on the line to the origin
	c = v7
	v9 = c
	v12 = c
	v17 = float32(v9.X*v12.X) + float32(v9.Y*v12.Y)
	goto _24
_24:
	cc = v17
	if cc > rr {
		// closest point is outside the circle
		return output
	}
	// Pythagoras
	h = sqrtf(tls, rr-cc)
	fraction = t - h
	if fraction < float32FromFloat32(0) || float32((*RayCastInput)(unsafe.Pointer(input)).MaxFraction**(*float32)(unsafe.Add(nil, bp))) < fraction {
		// intersection is point outside the range of the ray segment
		v1 = s1
		v10 = float32(v1.X*v1.X) + float32(v1.Y*v1.Y)
		goto _27
	_27:
		if v10 < rr {
			// initial overlap
			output.Point = (*RayCastInput)(unsafe.Pointer(input)).Origin
			output.Hit = boolUint8(true1 != 0)
		}
		return output
	}
	v1 = s1
	v10 = fraction
	v2 = d
	v3 = Vec2{
		X: v1.X + float32(v10*v2.X),
		Y: v1.Y + float32(v10*v2.Y),
	}
	goto _32
_32:
	// hit point relative to center
	hitPoint = v3
	output.Fraction = fraction / *(*float32)(unsafe.Add(nil, bp))
	v1 = hitPoint
	length = sqrtf(tls, float32(v1.X*v1.X)+float32(v1.Y*v1.Y))
	if length < float32FromFloat32(1.1920928955078125e-07) {
		v2 = Vec2{}
		goto _35
	}
	invLength = float32FromFloat32(1) / length
	n = Vec2{
		X: float32(invLength * v1.X),
		Y: float32(invLength * v1.Y),
	}
	v2 = n
	goto _35
_35:
	output.Normal = v2
	v1 = p
	v10 = (*Circle)(unsafe.Pointer(shape)).Radius
	v2 = output.Normal
	v3 = Vec2{
		X: v1.X + float32(v10*v2.X),
		Y: v1.Y + float32(v10*v2.Y),
	}
	goto _40
_40:
	output.Point = v3
	output.Hit = boolUint8(true1 != 0)
	return output
}

func b2RayCastCapsule(tls *_Stack, input uintptr, shape uintptr) (r CastOutput) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var a7, b11, b21, b6, d, e, n, n1, p1, q, qp, u, v11, v2, v1, v10, v111, v13, v14, v21, v3, v6, v7, v9 Vec2
	var den, invDen, invLength, qa, radius, s11, s21, s211, s22, v15, v18, v24 float32
	var output CastOutput
	var v5 uintptr
	var _ /* capsuleLength at bp+0 */ float32
	var _ /* circle at bp+16 */ Circle
	var _ /* circle at bp+28 */ Circle
	var _ /* circle at bp+4 */ Circle
	var _ /* circle at bp+44 */ Circle
	var _ /* circle at bp+56 */ Circle
	var _ /* rayLength at bp+40 */ float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a7, b11, b21, b6, d, den, e, invDen, invLength, n, n1, output, p1, q, qa, qp, radius, s11, s21, s211, s22, u, v11, v2, v1, v10, v111, v13, v14, v15, v18, v21, v24, v3, v5, v6, v7, v9
	if !(b2IsValidRay(tls, input) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+6779, __ccgo_ts+6524, int32FromInt32(584)) != 0 {
		__builtin_trap(tls)
	}
	output = CastOutput{}
	v11 = (*Capsule)(unsafe.Pointer(shape)).Center1
	v2 = (*Capsule)(unsafe.Pointer(shape)).Center2
	v1 = v2
	v21 = v11
	v3 = Vec2{
		X: v1.X - v21.X,
		Y: v1.Y - v21.Y,
	}
	goto _4
_4:
	e = v3
	v5 = bp
	v6 = e
	*(*float32)(unsafe.Add(nil, v5)) = sqrtf(tls, float32(v6.X*v6.X)+float32(v6.Y*v6.Y))
	if *(*float32)(unsafe.Add(nil, v5)) < float32FromFloat32(1.1920928955078125e-07) {
		v7 = Vec2{}
		goto _8
	}
	invLength = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v5))
	n = Vec2{
		X: float32(invLength * v6.X),
		Y: float32(invLength * v6.Y),
	}
	v7 = n
	goto _8
_8:
	a7 = v7
	if *(*float32)(unsafe.Add(nil, bp)) < float32FromFloat32(1.1920928955078125e-07) {
		// Capsule is really a circle
		*(*Circle)(unsafe.Add(nil, bp+4)) = Circle{
			Center: v11,
			Radius: (*Capsule)(unsafe.Pointer(shape)).Radius,
		}
		return b2RayCastCircle(tls, input, bp+4)
	}
	p1 = (*RayCastInput)(unsafe.Pointer(input)).Origin
	d = (*RayCastInput)(unsafe.Pointer(input)).Translation
	v1 = p1
	v21 = v11
	v3 = Vec2{
		X: v1.X - v21.X,
		Y: v1.Y - v21.Y,
	}
	goto _12
_12:
	// Ray from capsule start to ray start
	q = v3
	v6 = q
	v7 = a7
	v15 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y)
	goto _16
_16:
	qa = v15
	v9 = q
	v18 = -qa
	v10 = a7
	v111 = Vec2{
		X: v9.X + float32(v18*v10.X),
		Y: v9.Y + float32(v18*v10.Y),
	}
	goto _21
_21:
	// Vector to ray start that is perpendicular to capsule axis
	qp = v111
	radius = (*Capsule)(unsafe.Pointer(shape)).Radius
	// Does the ray start within the infinite length capsule?
	v13 = qp
	v14 = qp
	v24 = float32(v13.X*v14.X) + float32(v13.Y*v14.Y)
	goto _25
_25:
	if v24 < float32(radius*radius) {
		if qa < float32FromFloat32(0) {
			// start point behind capsule segment
			*(*Circle)(unsafe.Add(nil, bp+16)) = Circle{
				Center: v11,
				Radius: (*Capsule)(unsafe.Pointer(shape)).Radius,
			}
			return b2RayCastCircle(tls, input, bp+16)
		}
		if qa > *(*float32)(unsafe.Add(nil, bp)) {
			// start point ahead of capsule segment
			*(*Circle)(unsafe.Add(nil, bp+28)) = Circle{
				Center: v2,
				Radius: (*Capsule)(unsafe.Pointer(shape)).Radius,
			}
			return b2RayCastCircle(tls, input, bp+28)
		}
		// ray starts inside capsule -> no hit
		output.Point = (*RayCastInput)(unsafe.Pointer(input)).Origin
		output.Hit = boolUint8(true1 != 0)
		return output
	}
	// Perpendicular to capsule axis, pointing right
	n1 = Vec2{
		X: a7.Y,
		Y: -a7.X,
	}
	v5 = bp + 40
	v1 = d
	*(*float32)(unsafe.Add(nil, v5)) = sqrtf(tls, float32(v1.X*v1.X)+float32(v1.Y*v1.Y))
	if *(*float32)(unsafe.Add(nil, v5)) < float32FromFloat32(1.1920928955078125e-07) {
		v21 = Vec2{}
		goto _29
	}
	invLength = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v5))
	n = Vec2{
		X: float32(invLength * v1.X),
		Y: float32(invLength * v1.Y),
	}
	v21 = n
	goto _29
_29:
	u = v21
	// Intersect ray with infinite length capsule
	// v1 + radius * n + s1 * a = p1 + s2 * u
	// v1 - radius * n + s1 * a = p1 + s2 * u
	// s1 * a - s2 * u = b
	// b = q - radius * ap
	// or
	// b = q + radius * ap
	// Cramer's rule [a -u]
	den = float32(-a7.X*u.Y) + float32(u.X*a7.Y)
	if -float32FromFloat32(1.1920928955078125e-07) < den && den < float32FromFloat32(1.1920928955078125e-07) {
		// Ray is parallel to capsule and outside infinite length capsule
		return output
	}
	v1 = q
	v15 = radius
	v21 = n1
	v3 = Vec2{
		X: v1.X - float32(v15*v21.X),
		Y: v1.Y - float32(v15*v21.Y),
	}
	goto _34
_34:
	b11 = v3
	v6 = q
	v18 = radius
	v7 = n1
	v9 = Vec2{
		X: v6.X + float32(v18*v7.X),
		Y: v6.Y + float32(v18*v7.Y),
	}
	goto _39
_39:
	b21 = v9
	invDen = float32FromFloat32(1) / den
	// Cramer's rule [a b1]
	s211 = float32((float32(a7.X*b11.Y) - float32(b11.X*a7.Y)) * invDen)
	// Cramer's rule [a b2]
	s22 = float32((float32(a7.X*b21.Y) - float32(b21.X*a7.Y)) * invDen)
	if s211 < s22 {
		s21 = s211
		b6 = b11
	} else {
		s21 = s22
		b6 = b21
		v1 = n1
		v21 = Vec2{
			X: -v1.X,
			Y: -v1.Y,
		}
		goto _42
	_42:
		n1 = v21
	}
	if s21 < float32FromFloat32(0) || float32((*RayCastInput)(unsafe.Pointer(input)).MaxFraction**(*float32)(unsafe.Add(nil, bp+40))) < s21 {
		return output
	}
	// Cramer's rule [b -u]
	s11 = float32((float32(-b6.X*u.Y) + float32(u.X*b6.Y)) * invDen)
	if s11 < float32FromFloat32(0) {
		// ray passes behind capsule segment
		*(*Circle)(unsafe.Add(nil, bp+44)) = Circle{
			Center: v11,
			Radius: (*Capsule)(unsafe.Pointer(shape)).Radius,
		}
		return b2RayCastCircle(tls, input, bp+44)
	} else {
		if *(*float32)(unsafe.Add(nil, bp)) < s11 {
			// ray passes ahead of capsule segment
			*(*Circle)(unsafe.Add(nil, bp+56)) = Circle{
				Center: v2,
				Radius: (*Capsule)(unsafe.Pointer(shape)).Radius,
			}
			return b2RayCastCircle(tls, input, bp+56)
		} else {
			// ray hits capsule side
			output.Fraction = s21 / *(*float32)(unsafe.Add(nil, bp+40))
			v1 = v11
			v21 = v2
			v15 = s11 / *(*float32)(unsafe.Add(nil, bp))
			v3 = Vec2{
				X: float32((float32FromFloat32(1)-v15)*v1.X) + float32(v15*v21.X),
				Y: float32((float32FromFloat32(1)-v15)*v1.Y) + float32(v15*v21.Y),
			}
			goto _47
		_47:
			v18 = (*Capsule)(unsafe.Pointer(shape)).Radius
			v6 = n1
			v7 = Vec2{
				X: float32(v18 * v6.X),
				Y: float32(v18 * v6.Y),
			}
			goto _51
		_51:
			v9 = v3
			v10 = v7
			v111 = Vec2{
				X: v9.X + v10.X,
				Y: v9.Y + v10.Y,
			}
			goto _55
		_55:
			output.Point = v111
			output.Normal = n1
			output.Hit = boolUint8(true1 != 0)
			return output
		}
	}
}

// C documentation
//
//	// Ray vs line segment
func b2RayCastSegment(tls *_Stack, input uintptr, shape uintptr, oneSided uint8) (r CastOutput) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var d, e, eUnit, n, normal, p, p1, v11, v2, v1, v10, v13, v21, v3, v5, v6, v7, v9 Vec2
	var denominator, invLength, numerator, offset, s1, t, v111, v30 float32
	var output, output1 CastOutput
	var v17 uintptr
	var _ /* length at bp+0 */ float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = d, denominator, e, eUnit, invLength, n, normal, numerator, offset, output, output1, p, p1, s1, t, v11, v2, v1, v10, v111, v13, v17, v21, v3, v30, v5, v6, v7, v9
	if oneSided != 0 {
		v1 = (*RayCastInput)(unsafe.Pointer(input)).Origin
		v21 = (*Segment)(unsafe.Pointer(shape)).Point1
		v3 = Vec2{
			X: v1.X - v21.X,
			Y: v1.Y - v21.Y,
		}
		goto _4
	_4:
		v5 = (*Segment)(unsafe.Pointer(shape)).Point2
		v6 = (*Segment)(unsafe.Pointer(shape)).Point1
		v7 = Vec2{
			X: v5.X - v6.X,
			Y: v5.Y - v6.Y,
		}
		goto _8
	_8:
		v9 = v3
		v10 = v7
		v111 = float32(v9.X*v10.Y) - float32(v9.Y*v10.X)
		goto _12
	_12:
		// Skip left-side collision
		offset = v111
		if offset < float32FromFloat32(0) {
			output = CastOutput{}
			return output
		}
	}
	// Put the ray into the edge's frame of reference.
	p1 = (*RayCastInput)(unsafe.Pointer(input)).Origin
	d = (*RayCastInput)(unsafe.Pointer(input)).Translation
	v11 = (*Segment)(unsafe.Pointer(shape)).Point1
	v2 = (*Segment)(unsafe.Pointer(shape)).Point2
	v1 = v2
	v21 = v11
	v3 = Vec2{
		X: v1.X - v21.X,
		Y: v1.Y - v21.Y,
	}
	goto _16
_16:
	e = v3
	output1 = CastOutput{}
	v17 = bp
	v5 = e
	*(*float32)(unsafe.Add(nil, v17)) = sqrtf(tls, float32(v5.X*v5.X)+float32(v5.Y*v5.Y))
	if *(*float32)(unsafe.Add(nil, v17)) < float32FromFloat32(1.1920928955078125e-07) {
		v6 = Vec2{}
		goto _20
	}
	invLength = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v17))
	n = Vec2{
		X: float32(invLength * v5.X),
		Y: float32(invLength * v5.Y),
	}
	v6 = n
	goto _20
_20:
	eUnit = v6
	if *(*float32)(unsafe.Add(nil, bp)) == float32FromFloat32(0) {
		return output1
	}
	v1 = eUnit
	v21 = Vec2{
		X: v1.Y,
		Y: -v1.X,
	}
	goto _23
_23:
	// Normal points to the right, looking from v1 towards v2
	normal = v21
	v3 = v11
	v5 = p1
	v6 = Vec2{
		X: v3.X - v5.X,
		Y: v3.Y - v5.Y,
	}
	goto _27
_27:
	v7 = normal
	v9 = v6
	v111 = float32(v7.X*v9.X) + float32(v7.Y*v9.Y)
	goto _31
_31:
	// Intersect ray with infinite segment using normal
	// Similar to intersecting a ray with an infinite plane
	// p = p1 + t * d
	// dot(normal, p - v1) = 0
	// dot(normal, p1 - v1) + t * dot(normal, d) = 0
	numerator = v111
	v10 = normal
	v13 = d
	v30 = float32(v10.X*v13.X) + float32(v10.Y*v13.Y)
	goto _35
_35:
	denominator = v30
	if denominator == float32FromFloat32(0) {
		// parallel
		return output1
	}
	t = numerator / denominator
	if t < float32FromFloat32(0) || (*RayCastInput)(unsafe.Pointer(input)).MaxFraction < t {
		// out of ray range
		return output1
	}
	v1 = p1
	v111 = t
	v21 = d
	v3 = Vec2{
		X: v1.X + float32(v111*v21.X),
		Y: v1.Y + float32(v111*v21.Y),
	}
	goto _40
_40:
	// Intersection point on infinite segment
	p = v3
	v5 = p
	v6 = v11
	v7 = Vec2{
		X: v5.X - v6.X,
		Y: v5.Y - v6.Y,
	}
	goto _44
_44:
	v9 = v7
	v10 = eUnit
	v30 = float32(v9.X*v10.X) + float32(v9.Y*v10.Y)
	goto _48
_48:
	// Compute position of p along segment
	// p = v1 + s * e
	// s = dot(p - v1, e) / dot(e, e)
	s1 = v30
	if s1 < float32FromFloat32(0) || *(*float32)(unsafe.Add(nil, bp)) < s1 {
		// out of segment range
		return output1
	}
	if numerator > float32FromFloat32(0) {
		v1 = normal
		v21 = Vec2{
			X: -v1.X,
			Y: -v1.Y,
		}
		goto _51
	_51:
		normal = v21
	}
	output1.Fraction = t
	output1.Point = p
	output1.Normal = normal
	output1.Hit = boolUint8(true1 != 0)
	return output1
}

func b2RayCastPolygon(tls *_Stack, input uintptr, shape uintptr) (r CastOutput) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var base, d, p1, vertex, v1, v10, v11, v12, v14, v15, v18, v19, v2, v3, v6, v7, v8 Vec2
	var denominator, lower, numerator, upper, v16, v20 float32
	var i, index int32
	var output CastOutput
	var _ /* castInput at bp+0 */ ShapeCastPairInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = base, d, denominator, i, index, lower, numerator, output, p1, upper, vertex, v1, v10, v11, v12, v14, v15, v16, v18, v19, v2, v20, v3, v6, v7, v8
	if !(b2IsValidRay(tls, input) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+6779, __ccgo_ts+6524, int32FromInt32(801)) != 0 {
		__builtin_trap(tls)
	}
	if (*Polygon)(unsafe.Pointer(shape)).Radius == float32FromFloat32(0) {
		// Shift all math to first vertex since the polygon may be far
		// from the origin.
		base = *(*Vec2)(unsafe.Add(nil, shape))
		v1 = (*RayCastInput)(unsafe.Pointer(input)).Origin
		v2 = base
		v3 = Vec2{
			X: v1.X - v2.X,
			Y: v1.Y - v2.Y,
		}
		goto _4
	_4:
		p1 = v3
		d = (*RayCastInput)(unsafe.Pointer(input)).Translation
		lower = float32FromFloat32(0)
		upper = (*RayCastInput)(unsafe.Pointer(input)).MaxFraction
		index = -int32(1)
		output = CastOutput{}
		i = 0
		for {
			if !(i < (*Polygon)(unsafe.Pointer(shape)).Count) {
				break
			}
			v6 = *(*Vec2)(unsafe.Add(nil, shape+uintptr(i)*8))
			v7 = base
			v8 = Vec2{
				X: v6.X - v7.X,
				Y: v6.Y - v7.Y,
			}
			goto _9
		_9:
			// p = p1 + a * d
			// dot(normal, p - v) = 0
			// dot(normal, p1 - v) + a * dot(normal, d) = 0
			vertex = v8
			v10 = vertex
			v11 = p1
			v12 = Vec2{
				X: v10.X - v11.X,
				Y: v10.Y - v11.Y,
			}
			goto _13
		_13:
			v14 = *(*Vec2)(unsafe.Add(nil, shape+64+uintptr(i)*8))
			v15 = v12
			v16 = float32(v14.X*v15.X) + float32(v14.Y*v15.Y)
			goto _17
		_17:
			numerator = v16
			v18 = *(*Vec2)(unsafe.Add(nil, shape+64+uintptr(i)*8))
			v19 = d
			v20 = float32(v18.X*v19.X) + float32(v18.Y*v19.Y)
			goto _21
		_21:
			denominator = v20
			if denominator == float32FromFloat32(0) {
				if numerator < float32FromFloat32(0) {
					return output
				}
			} else {
				// Note: we want this predicate without division:
				// lower < numerator / denominator, where denominator < 0
				// Since denominator < 0, we have to flip the inequality:
				// lower < numerator / denominator <==> denominator * lower > numerator.
				if denominator < float32FromFloat32(0) && numerator < float32(lower*denominator) {
					// Increase lower.
					// The segment enters this half-space.
					lower = numerator / denominator
					index = i
				} else {
					if denominator > float32FromFloat32(0) && numerator < float32(upper*denominator) {
						// Decrease upper.
						// The segment exits this half-space.
						upper = numerator / denominator
					}
				}
			}
			if upper < lower {
				// Ray misses
				return output
			}
			goto _5
		_5:
			;
			i = i + 1
		}
		if !(float32FromFloat32(0) <= lower && lower <= (*RayCastInput)(unsafe.Pointer(input)).MaxFraction) && b2InternalAssertFcn(tls, __ccgo_ts+6801, __ccgo_ts+6524, int32FromInt32(862)) != 0 {
			__builtin_trap(tls)
		}
		if index >= 0 {
			output.Fraction = lower
			output.Normal = *(*Vec2)(unsafe.Add(nil, shape+64+uintptr(index)*8))
			v1 = (*RayCastInput)(unsafe.Pointer(input)).Origin
			v16 = lower
			v2 = d
			v3 = Vec2{
				X: v1.X + float32(v16*v2.X),
				Y: v1.Y + float32(v16*v2.Y),
			}
			goto _26
		_26:
			output.Point = v3
			output.Hit = boolUint8(true1 != 0)
		} else {
			// initial overlap
			output.Point = (*RayCastInput)(unsafe.Pointer(input)).Origin
			output.Hit = boolUint8(true1 != 0)
		}
		return output
	}
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).ProxyA = b2MakeProxy(tls, shape, (*Polygon)(unsafe.Pointer(shape)).Count, (*Polygon)(unsafe.Pointer(shape)).Radius)
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).ProxyB = b2MakeProxy(tls, input, int32(1), float32FromFloat32(0))
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).TransformA = b2Transform_identity
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).TransformB = b2Transform_identity
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).TranslationB = (*RayCastInput)(unsafe.Pointer(input)).Translation
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).MaxFraction = (*RayCastInput)(unsafe.Pointer(input)).MaxFraction
	(*(*ShapeCastPairInput)(unsafe.Add(nil, bp))).CanEncroach = boolUint8(false1 != 0)
	return b2ShapeCast(tls, bp)
}

const _FLT_MAX4 = 3.4028234663852886e+38

/**@}*/

/**
 * @defgroup math_cpp C++ Math
 * @brief Math operator overloads for C++
 *
 * See math_functions.h for details.
 * @{
 */

/**@}*/

// C documentation
//
//	// quickhull recursion
func b2RecurseHull(tls *_Stack, p1 Vec2, p2 Vec2, ps uintptr, count int32) (r Hull) {
	bp := tls.Alloc(272)
	defer tls.Free(272)
	var bestDistance, distance, invLength, length, v14 float32
	var bestIndex, i, i1, i2, rightCount, v16 int32
	var bestPoint, e, n, v1, v10, v12, v13, v2, v3, v5, v6, v8, v9 Vec2
	var v29 uintptr
	var _ /* hull at bp+0 */ Hull
	var _ /* hull1 at bp+132 */ Hull
	var _ /* hull2 at bp+200 */ Hull
	var _ /* rightPoints at bp+68 */ [8]Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bestDistance, bestIndex, bestPoint, distance, e, i, i1, i2, invLength, length, n, rightCount, v1, v10, v12, v13, v14, v16, v2, v29, v3, v5, v6, v8, v9
	(*(*Hull)(unsafe.Add(nil, bp))).Count = 0
	if count == 0 {
		return *(*Hull)(unsafe.Add(nil, bp))
	}
	v1 = p2
	v2 = p1
	v3 = Vec2{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
	goto _4
_4:
	v5 = v3
	length = sqrtf(tls, float32(v5.X*v5.X)+float32(v5.Y*v5.Y))
	if length < float32FromFloat32(1.1920928955078125e-07) {
		v6 = Vec2{}
		goto _7
	}
	invLength = float32FromFloat32(1) / length
	n = Vec2{
		X: float32(invLength * v5.X),
		Y: float32(invLength * v5.Y),
	}
	v6 = n
	goto _7
_7:
	// create an edge vector pointing from p1 to p2
	e = v6
	rightCount = 0
	bestIndex = 0
	v8 = *(*Vec2)(unsafe.Add(nil, ps+uintptr(bestIndex)*8))
	v9 = p1
	v10 = Vec2{
		X: v8.X - v9.X,
		Y: v8.Y - v9.Y,
	}
	goto _11
_11:
	v12 = v10
	v13 = e
	v14 = float32(v12.X*v13.Y) - float32(v12.Y*v13.X)
	goto _15
_15:
	bestDistance = v14
	if bestDistance > float32FromFloat32(0) {
		v16 = rightCount
		rightCount = rightCount + 1
		(*(*[8]Vec2)(unsafe.Add(nil, bp+68)))[v16] = *(*Vec2)(unsafe.Add(nil, ps+uintptr(bestIndex)*8))
	}
	i = int32(1)
	for {
		if !(i < count) {
			break
		}
		v1 = *(*Vec2)(unsafe.Add(nil, ps+uintptr(i)*8))
		v2 = p1
		v3 = Vec2{
			X: v1.X - v2.X,
			Y: v1.Y - v2.Y,
		}
		goto _21
	_21:
		v5 = v3
		v6 = e
		v14 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
		goto _25
	_25:
		distance = v14
		if distance > bestDistance {
			bestIndex = i
			bestDistance = distance
		}
		if distance > float32FromFloat32(0) {
			v16 = rightCount
			rightCount = rightCount + 1
			(*(*[8]Vec2)(unsafe.Add(nil, bp+68)))[v16] = *(*Vec2)(unsafe.Add(nil, ps+uintptr(i)*8))
		}
		goto _17
	_17:
		;
		i = i + 1
	}
	if bestDistance < float32(float32FromFloat32(2)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter)) {
		return *(*Hull)(unsafe.Add(nil, bp))
	}
	bestPoint = *(*Vec2)(unsafe.Add(nil, ps+uintptr(bestIndex)*8))
	// compute hull to the right of p1-bestPoint
	*(*Hull)(unsafe.Add(nil, bp+132)) = b2RecurseHull(tls, p1, bestPoint, bp+68, rightCount)
	// compute hull to the right of bestPoint-p2
	*(*Hull)(unsafe.Add(nil, bp+200)) = b2RecurseHull(tls, bestPoint, p2, bp+68, rightCount)
	// stitch together hulls
	i1 = 0
	for {
		if !(i1 < (*(*Hull)(unsafe.Add(nil, bp+132))).Count) {
			break
		}
		v29 = bp + 64
		v16 = *(*int32)(unsafe.Pointer(v29))
		*(*int32)(unsafe.Pointer(v29)) = *(*int32)(unsafe.Pointer(v29)) + 1
		*(*Vec2)(unsafe.Add(nil, bp+uintptr(v16)*8)) = *(*Vec2)(unsafe.Add(nil, bp+132+uintptr(i1)*8))
		goto _27
	_27:
		;
		i1 = i1 + 1
	}
	v29 = bp + 64
	v16 = *(*int32)(unsafe.Pointer(v29))
	*(*int32)(unsafe.Pointer(v29)) = *(*int32)(unsafe.Pointer(v29)) + 1
	*(*Vec2)(unsafe.Add(nil, bp+uintptr(v16)*8)) = bestPoint
	i2 = 0
	for {
		if !(i2 < (*(*Hull)(unsafe.Add(nil, bp+200))).Count) {
			break
		}
		v29 = bp + 64
		v16 = *(*int32)(unsafe.Pointer(v29))
		*(*int32)(unsafe.Pointer(v29)) = *(*int32)(unsafe.Pointer(v29)) + 1
		*(*Vec2)(unsafe.Add(nil, bp+uintptr(v16)*8)) = *(*Vec2)(unsafe.Add(nil, bp+200+uintptr(i2)*8))
		goto _32
	_32:
		;
		i2 = i2 + 1
	}
	if !((*(*Hull)(unsafe.Add(nil, bp))).Count < int32FromInt32(_B2_MAX_POLYGON_VERTICES)) && b2InternalAssertFcn(tls, __ccgo_ts+6846, __ccgo_ts+6883, int32FromInt32(78)) != 0 {
		__builtin_trap(tls)
	}
	return *(*Hull)(unsafe.Add(nil, bp))
}

func b2ValidateHull(tls *_Stack, hull uintptr) (r uint8) {
	var distance, distance1, invLength, length, linearSlop, v17 float32
	var e, e1, n, p, p1, p2, p3, v11, v12, v13, v15, v16, v3, v4, v5, v7, v8 Vec2
	var i, i1, i11, i12, i2, i21, i3, j, v2 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = distance, distance1, e, e1, i, i1, i11, i12, i2, i21, i3, invLength, j, length, linearSlop, n, p, p1, p2, p3, v11, v12, v13, v15, v16, v17, v2, v3, v4, v5, v7, v8
	if (*Hull)(unsafe.Pointer(hull)).Count < int32(3) || int32(_B2_MAX_POLYGON_VERTICES) < (*Hull)(unsafe.Pointer(hull)).Count {
		return boolUint8(false1 != 0)
	}
	// test that every point is behind every edge
	i = 0
	for {
		if !(i < (*Hull)(unsafe.Pointer(hull)).Count) {
			break
		}
		// create an edge vector
		i11 = i
		if i < (*Hull)(unsafe.Pointer(hull)).Count-int32(1) {
			v2 = i11 + int32(1)
		} else {
			v2 = 0
		}
		i2 = v2
		p = *(*Vec2)(unsafe.Add(nil, hull+uintptr(i11)*8))
		v3 = *(*Vec2)(unsafe.Add(nil, hull+uintptr(i2)*8))
		v4 = p
		v5 = Vec2{
			X: v3.X - v4.X,
			Y: v3.Y - v4.Y,
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
		e = v8
		j = 0
		for {
			if !(j < (*Hull)(unsafe.Pointer(hull)).Count) {
				break
			}
			// skip points that subtend the current edge
			if j == i11 || j == i2 {
				goto _10
			}
			v3 = *(*Vec2)(unsafe.Add(nil, hull+uintptr(j)*8))
			v4 = p
			v5 = Vec2{
				X: v3.X - v4.X,
				Y: v3.Y - v4.Y,
			}
			goto _14
		_14:
			v7 = v5
			v8 = e
			v17 = float32(v7.X*v8.Y) - float32(v7.Y*v8.X)
			goto _18
		_18:
			distance = v17
			if distance >= float32FromFloat32(0) {
				return boolUint8(false1 != 0)
			}
			goto _10
		_10:
			;
			j = j + 1
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	// test for collinear points
	linearSlop = float32(float32FromFloat32(0.005) * b2_lengthUnitsPerMeter)
	i1 = 0
	for {
		if !(i1 < (*Hull)(unsafe.Pointer(hull)).Count) {
			break
		}
		i12 = i1
		i21 = (i1 + int32(1)) % (*Hull)(unsafe.Pointer(hull)).Count
		i3 = (i1 + int32(2)) % (*Hull)(unsafe.Pointer(hull)).Count
		p1 = *(*Vec2)(unsafe.Add(nil, hull+uintptr(i12)*8))
		p2 = *(*Vec2)(unsafe.Add(nil, hull+uintptr(i21)*8))
		p3 = *(*Vec2)(unsafe.Add(nil, hull+uintptr(i3)*8))
		v3 = p3
		v4 = p1
		v5 = Vec2{
			X: v3.X - v4.X,
			Y: v3.Y - v4.Y,
		}
		goto _23
	_23:
		v7 = v5
		length = sqrtf(tls, float32(v7.X*v7.X)+float32(v7.Y*v7.Y))
		if length < float32FromFloat32(1.1920928955078125e-07) {
			v8 = Vec2{}
			goto _26
		}
		invLength = float32FromFloat32(1) / length
		n = Vec2{
			X: float32(invLength * v7.X),
			Y: float32(invLength * v7.Y),
		}
		v8 = n
		goto _26
	_26:
		e1 = v8
		v11 = p2
		v12 = p1
		v13 = Vec2{
			X: v11.X - v12.X,
			Y: v11.Y - v12.Y,
		}
		goto _30
	_30:
		v15 = v13
		v16 = e1
		v17 = float32(v15.X*v16.Y) - float32(v15.Y*v16.X)
		goto _34
	_34:
		distance1 = v17
		if distance1 <= linearSlop {
			// p1-p2-p3 are collinear
			return boolUint8(false1 != 0)
		}
		goto _19
	_19:
		;
		i1 = i1 + 1
	}
	return boolUint8(true1 != 0)
}

const _FLT_MAX5 = 3.40282346638528859812e+38

func b2CreateIdPool(tls *_Stack) (r b2IdPool) {
	var pool b2IdPool
	_ = pool
	pool = b2IdPool{}
	pool.FreeArray = b2IntArray_Create(tls, int32(32))
	return pool
}

func b2DestroyIdPool(tls *_Stack, pool uintptr) {
	b2IntArray_Destroy(tls, pool)
	*(*b2IdPool)(unsafe.Add(nil, pool)) = b2IdPool{}
}

func b2AllocId(tls *_Stack, pool uintptr) (r int32) {
	var count, id, id1, value, v2 int32
	var v1 uintptr
	_, _, _, _, _, _ = count, id, id1, value, v1, v2
	count = (*b2IdPool)(unsafe.Pointer(pool)).FreeArray.Count
	if count > 0 {
		v1 = pool
		if !((*b2IntArray)(unsafe.Pointer(v1)).Count > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+119, __ccgo_ts+2400, int32FromInt32(155)) != 0 {
			__builtin_trap(tls)
		}
		value = *(*int32)(unsafe.Add(nil, (*b2IntArray)(unsafe.Pointer(v1)).Data+uintptr((*b2IntArray)(unsafe.Pointer(v1)).Count-int32(1))*4))
		*(*int32)(unsafe.Add(nil, v1+8)) -= int32(1)
		v2 = value
		goto _3
	_3:
		id = v2
		return id
	}
	id1 = (*b2IdPool)(unsafe.Pointer(pool)).NextIndex
	*(*int32)(unsafe.Add(nil, pool+16)) += int32(1)
	return id1
}

func b2FreeId(tls *_Stack, pool uintptr, id int32) {
	var newCapacity, v2 int32
	var v1 uintptr
	_, _, _ = newCapacity, v1, v2
	if !((*b2IdPool)(unsafe.Pointer(pool)).NextIndex > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+6943, __ccgo_ts+6963, int32FromInt32(35)) != 0 {
		__builtin_trap(tls)
	}
	if !(0 <= id && id < (*b2IdPool)(unsafe.Pointer(pool)).NextIndex) && b2InternalAssertFcn(tls, __ccgo_ts+6988, __ccgo_ts+6963, int32FromInt32(36)) != 0 {
		__builtin_trap(tls)
	}
	v1 = pool
	if (*b2IntArray)(unsafe.Pointer(v1)).Count == (*b2IntArray)(unsafe.Pointer(v1)).Capacity {
		if (*b2IntArray)(unsafe.Pointer(v1)).Capacity < int32(2) {
			v2 = int32(2)
		} else {
			v2 = (*b2IntArray)(unsafe.Pointer(v1)).Capacity + (*b2IntArray)(unsafe.Pointer(v1)).Capacity>>int32(1)
		}
		newCapacity = v2
		b2IntArray_Reserve(tls, v1, newCapacity)
	}
	*(*int32)(unsafe.Add(nil, (*b2IntArray)(unsafe.Pointer(v1)).Data+uintptr((*b2IntArray)(unsafe.Pointer(v1)).Count)*4)) = id
	*(*int32)(unsafe.Add(nil, v1+8)) += int32(1)
}

func b2ValidateFreeId(tls *_Stack, pool uintptr, id int32) {
	_ = uint64FromInt64(4)
}

func b2ValidateUsedId(tls *_Stack, pool uintptr, id int32) {
	_ = uint64FromInt64(4)
}

const _B2_CONTACT_REMOVE_THRESHOLD = 1

var b2_identityBodyState6 = b2BodyState{
	DeltaRotation: Rot{
		C: float32FromFloat32(1),
	},
}

func b2IslandArray_Create(tls *_Stack, capacity int32) (r b2IslandArray) {
	var a b2IslandArray
	_ = a
	a = b2IslandArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(56)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2IslandArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2IslandArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2IslandArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2IslandArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2IslandArray)(unsafe.Pointer(a)).Capacity)*uint64(56)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(56)))
	(*b2IslandArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2IslandArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2IslandArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2IslandArray)(unsafe.Pointer(a)).Capacity)*uint64(56)))
	(*b2IslandArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2IslandArray)(unsafe.Pointer(a)).Count = 0
	(*b2IslandArray)(unsafe.Pointer(a)).Capacity = 0
}

func b2IslandSimArray_Create(tls *_Stack, capacity int32) (r b2IslandSimArray) {
	var a b2IslandSimArray
	_ = a
	a = b2IslandSimArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(4)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2IslandSimArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2IslandSimArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2IslandSimArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2IslandSimArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2IslandSimArray)(unsafe.Pointer(a)).Capacity)*uint64(4)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(4)))
	(*b2IslandSimArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2IslandSimArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2IslandSimArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2IslandSimArray)(unsafe.Pointer(a)).Capacity)*uint64(4)))
	(*b2IslandSimArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2IslandSimArray)(unsafe.Pointer(a)).Count = 0
	(*b2IslandSimArray)(unsafe.Pointer(a)).Capacity = 0
}

func b2CreateIsland(tls *_Stack, world uintptr, setIndex int32) (r uintptr) {
	var emptyIsland b2Island
	var island, islandSim, set, v1, v3, v5, v7 uintptr
	var islandId, newCapacity, newCapacity1, v2, v4 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _ = emptyIsland, island, islandId, islandSim, newCapacity, newCapacity1, set, v1, v2, v3, v4, v5, v7
	if !(setIndex == int32(b2_awakeSet) || setIndex >= int32(b2_firstSleepingSet)) && b2InternalAssertFcn(tls, __ccgo_ts+7020, __ccgo_ts+7079, int32FromInt32(20)) != 0 {
		__builtin_trap(tls)
	}
	islandId = b2AllocId(tls, world+1160)
	if islandId == (*b2World)(unsafe.Pointer(world)).Islands.Count {
		emptyIsland = b2Island{}
		v1 = world + 1184
		if (*b2IslandArray)(unsafe.Pointer(v1)).Count == (*b2IslandArray)(unsafe.Pointer(v1)).Capacity {
			if (*b2IslandArray)(unsafe.Pointer(v1)).Capacity < int32(2) {
				v2 = int32(2)
			} else {
				v2 = (*b2IslandArray)(unsafe.Pointer(v1)).Capacity + (*b2IslandArray)(unsafe.Pointer(v1)).Capacity>>int32(1)
			}
			newCapacity = v2
			b2IslandArray_Reserve(tls, v1, newCapacity)
		}
		*(*b2Island)(unsafe.Add(nil, (*b2IslandArray)(unsafe.Pointer(v1)).Data+uintptr((*b2IslandArray)(unsafe.Pointer(v1)).Count)*56)) = emptyIsland
		*(*int32)(unsafe.Add(nil, v1+8)) += int32(1)
	} else {
		if !((*(*b2Island)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Islands.Data+uintptr(islandId)*56))).SetIndex == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+7103, __ccgo_ts+7079, int32FromInt32(31)) != 0 {
			__builtin_trap(tls)
		}
	}
	v1 = world + 1064
	v2 = setIndex
	if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*88
	goto _6
_6:
	set = v3
	v5 = world + 1184
	v4 = islandId
	if !(0 <= v4 && v4 < (*b2IslandArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2IslandArray)(unsafe.Pointer(v5)).Data + uintptr(v4)*56
	goto _10
_10:
	island = v7
	(*b2Island)(unsafe.Pointer(island)).SetIndex = setIndex
	(*b2Island)(unsafe.Pointer(island)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(set)).IslandSims.Count
	(*b2Island)(unsafe.Pointer(island)).IslandId = islandId
	(*b2Island)(unsafe.Pointer(island)).HeadBody = -int32(1)
	(*b2Island)(unsafe.Pointer(island)).TailBody = -int32(1)
	(*b2Island)(unsafe.Pointer(island)).BodyCount = 0
	(*b2Island)(unsafe.Pointer(island)).HeadContact = -int32(1)
	(*b2Island)(unsafe.Pointer(island)).TailContact = -int32(1)
	(*b2Island)(unsafe.Pointer(island)).ContactCount = 0
	(*b2Island)(unsafe.Pointer(island)).HeadJoint = -int32(1)
	(*b2Island)(unsafe.Pointer(island)).TailJoint = -int32(1)
	(*b2Island)(unsafe.Pointer(island)).JointCount = 0
	(*b2Island)(unsafe.Pointer(island)).ParentIsland = -int32(1)
	(*b2Island)(unsafe.Pointer(island)).ConstraintRemoveCount = 0
	v1 = set + 64
	if (*b2IslandSimArray)(unsafe.Pointer(v1)).Count == (*b2IslandSimArray)(unsafe.Pointer(v1)).Capacity {
		if (*b2IslandSimArray)(unsafe.Pointer(v1)).Capacity < int32(2) {
			v2 = int32(2)
		} else {
			v2 = (*b2IslandSimArray)(unsafe.Pointer(v1)).Capacity + (*b2IslandSimArray)(unsafe.Pointer(v1)).Capacity>>int32(1)
		}
		newCapacity1 = v2
		b2IslandSimArray_Reserve(tls, v1, newCapacity1)
	}
	*(*int32)(unsafe.Add(nil, v1+8)) += int32(1)
	v3 = (*b2IslandSimArray)(unsafe.Pointer(v1)).Data + uintptr((*b2IslandSimArray)(unsafe.Pointer(v1)).Count-int32FromInt32(1))*4
	goto _14
_14:
	islandSim = v3
	(*b2IslandSim)(unsafe.Pointer(islandSim)).IslandId = islandId
	return island
}

func b2DestroyIsland(tls *_Stack, world uintptr, islandId int32) {
	var island, movedElement, movedIsland, set, v1, v13, v15, v3, v5, v7, v9 uintptr
	var movedId, movedIndex, movedIndex1, v10, v11, v14, v2, v6 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = island, movedElement, movedId, movedIndex, movedIndex1, movedIsland, set, v1, v10, v11, v13, v14, v15, v2, v3, v5, v6, v7, v9
	if (*b2World)(unsafe.Pointer(world)).SplitIslandId == islandId {
		(*b2World)(unsafe.Pointer(world)).SplitIslandId = -int32(1)
	}
	v1 = world + 1184
	v2 = islandId
	if !(0 <= v2 && v2 < (*b2IslandArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2IslandArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*56
	goto _4
_4:
	// assume island is empty
	island = v3
	v5 = world + 1064
	v6 = (*b2Island)(unsafe.Pointer(island)).SetIndex
	if !(0 <= v6 && v6 < (*b2SolverSetArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2SolverSetArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*88
	goto _8
_8:
	set = v7
	v9 = set + 64
	v10 = (*b2Island)(unsafe.Pointer(island)).LocalIndex
	if !(0 <= v10 && v10 < (*b2IslandSimArray)(unsafe.Pointer(v9)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(89)) != 0 {
		__builtin_trap(tls)
	}
	movedIndex = -int32(1)
	if v10 != (*b2IslandSimArray)(unsafe.Pointer(v9)).Count-int32FromInt32(1) {
		movedIndex = (*b2IslandSimArray)(unsafe.Pointer(v9)).Count - int32(1)
		*(*b2IslandSim)(unsafe.Add(nil, (*b2IslandSimArray)(unsafe.Pointer(v9)).Data+uintptr(v10)*4)) = *(*b2IslandSim)(unsafe.Add(nil, (*b2IslandSimArray)(unsafe.Pointer(v9)).Data+uintptr(movedIndex)*4))
	}
	*(*int32)(unsafe.Add(nil, v9+8)) -= int32(1)
	v11 = movedIndex
	goto _12
_12:
	movedIndex1 = v11
	if movedIndex1 != -int32(1) {
		// Fix index on moved element
		movedElement = (*b2SolverSet)(unsafe.Pointer(set)).IslandSims.Data + uintptr((*b2Island)(unsafe.Pointer(island)).LocalIndex)*4
		movedId = (*b2IslandSim)(unsafe.Pointer(movedElement)).IslandId
		v13 = world + 1184
		v14 = movedId
		if !(0 <= v14 && v14 < (*b2IslandArray)(unsafe.Pointer(v13)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
			__builtin_trap(tls)
		}
		v15 = (*b2IslandArray)(unsafe.Pointer(v13)).Data + uintptr(v14)*56
		goto _16
	_16:
		movedIsland = v15
		if !((*b2Island)(unsafe.Pointer(movedIsland)).LocalIndex == movedIndex1) && b2InternalAssertFcn(tls, __ccgo_ts+7159, __ccgo_ts+7079, int32FromInt32(75)) != 0 {
			__builtin_trap(tls)
		}
		(*b2Island)(unsafe.Pointer(movedIsland)).LocalIndex = (*b2Island)(unsafe.Pointer(island)).LocalIndex
	}
	// Free island and id (preserve island revision)
	(*b2Island)(unsafe.Pointer(island)).IslandId = -int32(1)
	(*b2Island)(unsafe.Pointer(island)).SetIndex = -int32(1)
	(*b2Island)(unsafe.Pointer(island)).LocalIndex = -int32(1)
	b2FreeId(tls, world+1160, islandId)
}

func b2AddContactToIsland(tls *_Stack, world uintptr, islandId int32, contact uintptr) {
	var headContact, island, v1, v3 uintptr
	var v2 int32
	_, _, _, _, _ = headContact, island, v1, v2, v3
	if !((*b2Contact)(unsafe.Pointer(contact)).IslandId == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+7197, __ccgo_ts+7079, int32FromInt32(88)) != 0 {
		__builtin_trap(tls)
	}
	if !((*b2Contact)(unsafe.Pointer(contact)).IslandPrev == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+7232, __ccgo_ts+7079, int32FromInt32(89)) != 0 {
		__builtin_trap(tls)
	}
	if !((*b2Contact)(unsafe.Pointer(contact)).IslandNext == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+7269, __ccgo_ts+7079, int32FromInt32(90)) != 0 {
		__builtin_trap(tls)
	}
	v1 = world + 1184
	v2 = islandId
	if !(0 <= v2 && v2 < (*b2IslandArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2IslandArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*56
	goto _4
_4:
	island = v3
	if (*b2Island)(unsafe.Pointer(island)).HeadContact != -int32(1) {
		(*b2Contact)(unsafe.Pointer(contact)).IslandNext = (*b2Island)(unsafe.Pointer(island)).HeadContact
		v1 = world + 1144
		v2 = (*b2Island)(unsafe.Pointer(island)).HeadContact
		if !(0 <= v2 && v2 < (*b2ContactArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2ContactArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*68
		goto _8
	_8:
		headContact = v3
		(*b2Contact)(unsafe.Pointer(headContact)).IslandPrev = (*b2Contact)(unsafe.Pointer(contact)).ContactId
	}
	(*b2Island)(unsafe.Pointer(island)).HeadContact = (*b2Contact)(unsafe.Pointer(contact)).ContactId
	if (*b2Island)(unsafe.Pointer(island)).TailContact == -int32(1) {
		(*b2Island)(unsafe.Pointer(island)).TailContact = (*b2Island)(unsafe.Pointer(island)).HeadContact
	}
	*(*int32)(unsafe.Add(nil, island+32)) += int32(1)
	(*b2Contact)(unsafe.Pointer(contact)).IslandId = islandId
	b2ValidateIsland(tls, world, islandId)
}

// C documentation
//
//	// Link a contact into an island.
//	// This performs union-find and path compression to join islands.
//	// https://en.wikipedia.org/wiki/Disjoint-set_data_structure
func b2LinkContact(tls *_Stack, world uintptr, contact uintptr) {
	var bodyA, bodyB, islandA, islandB, parent, parent1, v1, v3, v5, v7 uintptr
	var bodyIdA, bodyIdB, islandIdA, islandIdB, parentId, parentId1, v2, v6 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, bodyIdA, bodyIdB, islandA, islandB, islandIdA, islandIdB, parent, parent1, parentId, parentId1, v1, v2, v3, v5, v6, v7
	if !((*b2Contact)(unsafe.Pointer(contact)).Flags&uint32(b2_contactTouchingFlag) != uint32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+7306, __ccgo_ts+7079, int32FromInt32(118)) != 0 {
		__builtin_trap(tls)
	}
	bodyIdA = (*(*b2ContactEdge)(unsafe.Add(nil, contact+12))).BodyId
	bodyIdB = (*(*b2ContactEdge)(unsafe.Add(nil, contact+12+1*12))).BodyId
	v1 = world + 1024
	v2 = bodyIdA
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	bodyA = v3
	v5 = world + 1024
	v6 = bodyIdB
	if !(0 <= v6 && v6 < (*b2BodyArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2BodyArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*128
	goto _8
_8:
	bodyB = v7
	if !((*b2Body)(unsafe.Pointer(bodyA)).SetIndex != int32(b2_disabledSet) && (*b2Body)(unsafe.Pointer(bodyB)).SetIndex != int32(b2_disabledSet)) && b2InternalAssertFcn(tls, __ccgo_ts+3422, __ccgo_ts+7079, int32FromInt32(126)) != 0 {
		__builtin_trap(tls)
	}
	if !((*b2Body)(unsafe.Pointer(bodyA)).SetIndex != int32(b2_staticSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex != int32(b2_staticSet)) && b2InternalAssertFcn(tls, __ccgo_ts+3493, __ccgo_ts+7079, int32FromInt32(127)) != 0 {
		__builtin_trap(tls)
	}
	// Wake bodyB if bodyA is awake and bodyB is sleeping
	if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) && (*b2Body)(unsafe.Pointer(bodyB)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeSolverSet(tls, world, (*b2Body)(unsafe.Pointer(bodyB)).SetIndex)
	}
	// Wake bodyA if bodyB is awake and bodyA is sleeping
	if (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet) && (*b2Body)(unsafe.Pointer(bodyA)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeSolverSet(tls, world, (*b2Body)(unsafe.Pointer(bodyA)).SetIndex)
	}
	islandIdA = (*b2Body)(unsafe.Pointer(bodyA)).IslandId
	islandIdB = (*b2Body)(unsafe.Pointer(bodyB)).IslandId
	// Static bodies have null island indices.
	if !((*b2Body)(unsafe.Pointer(bodyA)).SetIndex != int32(b2_staticSet) || islandIdA == -int32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+7355, __ccgo_ts+7079, int32FromInt32(145)) != 0 {
		__builtin_trap(tls)
	}
	if !((*b2Body)(unsafe.Pointer(bodyB)).SetIndex != int32(b2_staticSet) || islandIdB == -int32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+7417, __ccgo_ts+7079, int32FromInt32(146)) != 0 {
		__builtin_trap(tls)
	}
	if !(islandIdA != -int32(1) || islandIdB != -int32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+7479, __ccgo_ts+7079, int32FromInt32(147)) != 0 {
		__builtin_trap(tls)
	}
	if islandIdA == islandIdB {
		// Contact in same island
		b2AddContactToIsland(tls, world, islandIdA, contact)
		return
	}
	// Union-find root of islandA
	islandA = uintptrFromInt32(0)
	if islandIdA != -int32(1) {
		v1 = world + 1184
		v2 = islandIdA
		if !(0 <= v2 && v2 < (*b2IslandArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2IslandArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*56
		goto _12
	_12:
		islandA = v3
		parentId = (*b2Island)(unsafe.Pointer(islandA)).ParentIsland
		for parentId != -int32(1) {
			v1 = world + 1184
			v2 = parentId
			if !(0 <= v2 && v2 < (*b2IslandArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
				__builtin_trap(tls)
			}
			v3 = (*b2IslandArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*56
			goto _16
		_16:
			parent = v3
			if (*b2Island)(unsafe.Pointer(parent)).ParentIsland != -int32(1) {
				// path compression
				(*b2Island)(unsafe.Pointer(islandA)).ParentIsland = (*b2Island)(unsafe.Pointer(parent)).ParentIsland
			}
			islandA = parent
			islandIdA = parentId
			parentId = (*b2Island)(unsafe.Pointer(islandA)).ParentIsland
		}
	}
	// Union-find root of islandB
	islandB = uintptrFromInt32(0)
	if islandIdB != -int32(1) {
		v1 = world + 1184
		v2 = islandIdB
		if !(0 <= v2 && v2 < (*b2IslandArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2IslandArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*56
		goto _20
	_20:
		islandB = v3
		parentId1 = (*b2Island)(unsafe.Pointer(islandB)).ParentIsland
		for (*b2Island)(unsafe.Pointer(islandB)).ParentIsland != -int32(1) {
			v1 = world + 1184
			v2 = parentId1
			if !(0 <= v2 && v2 < (*b2IslandArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
				__builtin_trap(tls)
			}
			v3 = (*b2IslandArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*56
			goto _24
		_24:
			parent1 = v3
			if (*b2Island)(unsafe.Pointer(parent1)).ParentIsland != -int32(1) {
				// path compression
				(*b2Island)(unsafe.Pointer(islandB)).ParentIsland = (*b2Island)(unsafe.Pointer(parent1)).ParentIsland
			}
			islandB = parent1
			islandIdB = parentId1
			parentId1 = (*b2Island)(unsafe.Pointer(islandB)).ParentIsland
		}
	}
	if !(islandA != uintptrFromInt32(0) || islandB != uintptrFromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+7536, __ccgo_ts+7079, int32FromInt32(198)) != 0 {
		__builtin_trap(tls)
	}
	// Union-Find link island roots
	if islandA != islandB && islandA != uintptrFromInt32(0) && islandB != uintptrFromInt32(0) {
		if !(islandA != islandB) && b2InternalAssertFcn(tls, __ccgo_ts+7571, __ccgo_ts+7079, int32FromInt32(203)) != 0 {
			__builtin_trap(tls)
		}
		if !((*b2Island)(unsafe.Pointer(islandB)).ParentIsland == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+7590, __ccgo_ts+7079, int32FromInt32(204)) != 0 {
			__builtin_trap(tls)
		}
		(*b2Island)(unsafe.Pointer(islandB)).ParentIsland = islandIdA
	}
	if islandA != uintptrFromInt32(0) {
		b2AddContactToIsland(tls, world, islandIdA, contact)
	} else {
		b2AddContactToIsland(tls, world, islandIdB, contact)
	}
	// todo why not merge the islands right here?
}

// C documentation
//
//	// This is called when a contact no longer has contact points or when a contact is destroyed.
func b2UnlinkContact(tls *_Stack, world uintptr, contact uintptr) {
	var island, nextContact, prevContact, v1, v3, v5, v7 uintptr
	var islandId, v2, v6 int32
	_, _, _, _, _, _, _, _, _, _ = island, islandId, nextContact, prevContact, v1, v2, v3, v5, v6, v7
	if !((*b2Contact)(unsafe.Pointer(contact)).IslandId != -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+7629, __ccgo_ts+7079, int32FromInt32(223)) != 0 {
		__builtin_trap(tls)
	}
	// remove from island
	islandId = (*b2Contact)(unsafe.Pointer(contact)).IslandId
	v1 = world + 1184
	v2 = islandId
	if !(0 <= v2 && v2 < (*b2IslandArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2IslandArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*56
	goto _4
_4:
	island = v3
	if (*b2Contact)(unsafe.Pointer(contact)).IslandPrev != -int32(1) {
		v5 = world + 1144
		v6 = (*b2Contact)(unsafe.Pointer(contact)).IslandPrev
		if !(0 <= v6 && v6 < (*b2ContactArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
			__builtin_trap(tls)
		}
		v7 = (*b2ContactArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*68
		goto _8
	_8:
		prevContact = v7
		if !((*b2Contact)(unsafe.Pointer(prevContact)).IslandNext == (*b2Contact)(unsafe.Pointer(contact)).ContactId) && b2InternalAssertFcn(tls, __ccgo_ts+7664, __ccgo_ts+7079, int32FromInt32(232)) != 0 {
			__builtin_trap(tls)
		}
		(*b2Contact)(unsafe.Pointer(prevContact)).IslandNext = (*b2Contact)(unsafe.Pointer(contact)).IslandNext
	}
	if (*b2Contact)(unsafe.Pointer(contact)).IslandNext != -int32(1) {
		v1 = world + 1144
		v2 = (*b2Contact)(unsafe.Pointer(contact)).IslandNext
		if !(0 <= v2 && v2 < (*b2ContactArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2ContactArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*68
		goto _12
	_12:
		nextContact = v3
		if !((*b2Contact)(unsafe.Pointer(nextContact)).IslandPrev == (*b2Contact)(unsafe.Pointer(contact)).ContactId) && b2InternalAssertFcn(tls, __ccgo_ts+7710, __ccgo_ts+7079, int32FromInt32(239)) != 0 {
			__builtin_trap(tls)
		}
		(*b2Contact)(unsafe.Pointer(nextContact)).IslandPrev = (*b2Contact)(unsafe.Pointer(contact)).IslandPrev
	}
	if (*b2Island)(unsafe.Pointer(island)).HeadContact == (*b2Contact)(unsafe.Pointer(contact)).ContactId {
		(*b2Island)(unsafe.Pointer(island)).HeadContact = (*b2Contact)(unsafe.Pointer(contact)).IslandNext
	}
	if (*b2Island)(unsafe.Pointer(island)).TailContact == (*b2Contact)(unsafe.Pointer(contact)).ContactId {
		(*b2Island)(unsafe.Pointer(island)).TailContact = (*b2Contact)(unsafe.Pointer(contact)).IslandPrev
	}
	if !((*b2Island)(unsafe.Pointer(island)).ContactCount > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+7756, __ccgo_ts+7079, int32FromInt32(253)) != 0 {
		__builtin_trap(tls)
	}
	*(*int32)(unsafe.Add(nil, island+32)) -= int32(1)
	*(*int32)(unsafe.Add(nil, island+52)) += int32(1)
	(*b2Contact)(unsafe.Pointer(contact)).IslandId = -int32(1)
	(*b2Contact)(unsafe.Pointer(contact)).IslandPrev = -int32(1)
	(*b2Contact)(unsafe.Pointer(contact)).IslandNext = -int32(1)
	b2ValidateIsland(tls, world, islandId)
}

// C documentation
//
//	// Merge an island into its root island.
//	// todo we can assume all islands are awake here
func b2MergeIsland(tls *_Stack, world uintptr, island uintptr) {
	var body, contact, headBody, headContact, headJoint, joint, rootIsland, tailBody, tailContact, tailJoint, v1, v3 uintptr
	var bodyId, contactId, jointId, rootId, v2 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodyId, contact, contactId, headBody, headContact, headJoint, joint, jointId, rootId, rootIsland, tailBody, tailContact, tailJoint, v1, v2, v3
	if !((*b2Island)(unsafe.Pointer(island)).ParentIsland != -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+7987, __ccgo_ts+7079, int32FromInt32(431)) != 0 {
		__builtin_trap(tls)
	}
	rootId = (*b2Island)(unsafe.Pointer(island)).ParentIsland
	v1 = world + 1184
	v2 = rootId
	if !(0 <= v2 && v2 < (*b2IslandArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2IslandArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*56
	goto _4
_4:
	rootIsland = v3
	if !((*b2Island)(unsafe.Pointer(rootIsland)).ParentIsland == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+8025, __ccgo_ts+7079, int32FromInt32(435)) != 0 {
		__builtin_trap(tls)
	}
	// remap island indices
	bodyId = (*b2Island)(unsafe.Pointer(island)).HeadBody
	for bodyId != -int32(1) {
		v1 = world + 1024
		v2 = bodyId
		if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
		goto _8
	_8:
		body = v3
		(*b2Body)(unsafe.Pointer(body)).IslandId = rootId
		bodyId = (*b2Body)(unsafe.Pointer(body)).IslandNext
	}
	contactId = (*b2Island)(unsafe.Pointer(island)).HeadContact
	for contactId != -int32(1) {
		v1 = world + 1144
		v2 = contactId
		if !(0 <= v2 && v2 < (*b2ContactArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2ContactArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*68
		goto _12
	_12:
		contact = v3
		(*b2Contact)(unsafe.Pointer(contact)).IslandId = rootId
		contactId = (*b2Contact)(unsafe.Pointer(contact)).IslandNext
	}
	jointId = (*b2Island)(unsafe.Pointer(island)).HeadJoint
	for jointId != -int32(1) {
		v1 = world + 1104
		v2 = jointId
		if !(0 <= v2 && v2 < (*b2JointArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2JointArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*72
		goto _16
	_16:
		joint = v3
		(*b2Joint)(unsafe.Pointer(joint)).IslandId = rootId
		jointId = (*b2Joint)(unsafe.Pointer(joint)).IslandNext
	}
	// connect body lists
	if !((*b2Island)(unsafe.Pointer(rootIsland)).TailBody != -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+8067, __ccgo_ts+7079, int32FromInt32(463)) != 0 {
		__builtin_trap(tls)
	}
	v1 = world + 1024
	v2 = (*b2Island)(unsafe.Pointer(rootIsland)).TailBody
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _20
_20:
	tailBody = v3
	if !((*b2Body)(unsafe.Pointer(tailBody)).IslandNext == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+8105, __ccgo_ts+7079, int32FromInt32(465)) != 0 {
		__builtin_trap(tls)
	}
	(*b2Body)(unsafe.Pointer(tailBody)).IslandNext = (*b2Island)(unsafe.Pointer(island)).HeadBody
	if !((*b2Island)(unsafe.Pointer(island)).HeadBody != -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+8143, __ccgo_ts+7079, int32FromInt32(468)) != 0 {
		__builtin_trap(tls)
	}
	v1 = world + 1024
	v2 = (*b2Island)(unsafe.Pointer(island)).HeadBody
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _24
_24:
	headBody = v3
	if !((*b2Body)(unsafe.Pointer(headBody)).IslandPrev == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+8177, __ccgo_ts+7079, int32FromInt32(470)) != 0 {
		__builtin_trap(tls)
	}
	(*b2Body)(unsafe.Pointer(headBody)).IslandPrev = (*b2Island)(unsafe.Pointer(rootIsland)).TailBody
	(*b2Island)(unsafe.Pointer(rootIsland)).TailBody = (*b2Island)(unsafe.Pointer(island)).TailBody
	*(*int32)(unsafe.Add(nil, rootIsland+20)) += (*b2Island)(unsafe.Pointer(island)).BodyCount
	// connect contact lists
	if (*b2Island)(unsafe.Pointer(rootIsland)).HeadContact == -int32(1) {
		// Root island has no contacts
		if !((*b2Island)(unsafe.Pointer(rootIsland)).TailContact == -int32(1) && (*b2Island)(unsafe.Pointer(rootIsland)).ContactCount == 0) && b2InternalAssertFcn(tls, __ccgo_ts+8215, __ccgo_ts+7079, int32FromInt32(480)) != 0 {
			__builtin_trap(tls)
		}
		(*b2Island)(unsafe.Pointer(rootIsland)).HeadContact = (*b2Island)(unsafe.Pointer(island)).HeadContact
		(*b2Island)(unsafe.Pointer(rootIsland)).TailContact = (*b2Island)(unsafe.Pointer(island)).TailContact
		(*b2Island)(unsafe.Pointer(rootIsland)).ContactCount = (*b2Island)(unsafe.Pointer(island)).ContactCount
	} else {
		if (*b2Island)(unsafe.Pointer(island)).HeadContact != -int32(1) {
			// Both islands have contacts
			if !((*b2Island)(unsafe.Pointer(island)).TailContact != -int32(1) && (*b2Island)(unsafe.Pointer(island)).ContactCount > 0) && b2InternalAssertFcn(tls, __ccgo_ts+8289, __ccgo_ts+7079, int32FromInt32(488)) != 0 {
				__builtin_trap(tls)
			}
			if !((*b2Island)(unsafe.Pointer(rootIsland)).TailContact != -int32(1) && (*b2Island)(unsafe.Pointer(rootIsland)).ContactCount > 0) && b2InternalAssertFcn(tls, __ccgo_ts+8354, __ccgo_ts+7079, int32FromInt32(489)) != 0 {
				__builtin_trap(tls)
			}
			v1 = world + 1144
			v2 = (*b2Island)(unsafe.Pointer(rootIsland)).TailContact
			if !(0 <= v2 && v2 < (*b2ContactArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
				__builtin_trap(tls)
			}
			v3 = (*b2ContactArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*68
			goto _28
		_28:
			tailContact = v3
			if !((*b2Contact)(unsafe.Pointer(tailContact)).IslandNext == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+8427, __ccgo_ts+7079, int32FromInt32(492)) != 0 {
				__builtin_trap(tls)
			}
			(*b2Contact)(unsafe.Pointer(tailContact)).IslandNext = (*b2Island)(unsafe.Pointer(island)).HeadContact
			v1 = world + 1144
			v2 = (*b2Island)(unsafe.Pointer(island)).HeadContact
			if !(0 <= v2 && v2 < (*b2ContactArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
				__builtin_trap(tls)
			}
			v3 = (*b2ContactArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*68
			goto _32
		_32:
			headContact = v3
			if !((*b2Contact)(unsafe.Pointer(headContact)).IslandPrev == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+8468, __ccgo_ts+7079, int32FromInt32(496)) != 0 {
				__builtin_trap(tls)
			}
			(*b2Contact)(unsafe.Pointer(headContact)).IslandPrev = (*b2Island)(unsafe.Pointer(rootIsland)).TailContact
			(*b2Island)(unsafe.Pointer(rootIsland)).TailContact = (*b2Island)(unsafe.Pointer(island)).TailContact
			*(*int32)(unsafe.Add(nil, rootIsland+32)) += (*b2Island)(unsafe.Pointer(island)).ContactCount
		}
	}
	if (*b2Island)(unsafe.Pointer(rootIsland)).HeadJoint == -int32(1) {
		// Root island has no joints
		if !((*b2Island)(unsafe.Pointer(rootIsland)).TailJoint == -int32(1) && (*b2Island)(unsafe.Pointer(rootIsland)).JointCount == 0) && b2InternalAssertFcn(tls, __ccgo_ts+8509, __ccgo_ts+7079, int32FromInt32(506)) != 0 {
			__builtin_trap(tls)
		}
		(*b2Island)(unsafe.Pointer(rootIsland)).HeadJoint = (*b2Island)(unsafe.Pointer(island)).HeadJoint
		(*b2Island)(unsafe.Pointer(rootIsland)).TailJoint = (*b2Island)(unsafe.Pointer(island)).TailJoint
		(*b2Island)(unsafe.Pointer(rootIsland)).JointCount = (*b2Island)(unsafe.Pointer(island)).JointCount
	} else {
		if (*b2Island)(unsafe.Pointer(island)).HeadJoint != -int32(1) {
			// Both islands have joints
			if !((*b2Island)(unsafe.Pointer(island)).TailJoint != -int32(1) && (*b2Island)(unsafe.Pointer(island)).JointCount > 0) && b2InternalAssertFcn(tls, __ccgo_ts+8579, __ccgo_ts+7079, int32FromInt32(514)) != 0 {
				__builtin_trap(tls)
			}
			if !((*b2Island)(unsafe.Pointer(rootIsland)).TailJoint != -int32(1) && (*b2Island)(unsafe.Pointer(rootIsland)).JointCount > 0) && b2InternalAssertFcn(tls, __ccgo_ts+8640, __ccgo_ts+7079, int32FromInt32(515)) != 0 {
				__builtin_trap(tls)
			}
			v1 = world + 1104
			v2 = (*b2Island)(unsafe.Pointer(rootIsland)).TailJoint
			if !(0 <= v2 && v2 < (*b2JointArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
				__builtin_trap(tls)
			}
			v3 = (*b2JointArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*72
			goto _36
		_36:
			tailJoint = v3
			if !((*b2Joint)(unsafe.Pointer(tailJoint)).IslandNext == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+8709, __ccgo_ts+7079, int32FromInt32(518)) != 0 {
				__builtin_trap(tls)
			}
			(*b2Joint)(unsafe.Pointer(tailJoint)).IslandNext = (*b2Island)(unsafe.Pointer(island)).HeadJoint
			v1 = world + 1104
			v2 = (*b2Island)(unsafe.Pointer(island)).HeadJoint
			if !(0 <= v2 && v2 < (*b2JointArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
				__builtin_trap(tls)
			}
			v3 = (*b2JointArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*72
			goto _40
		_40:
			headJoint = v3
			if !((*b2Joint)(unsafe.Pointer(headJoint)).IslandPrev == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+8748, __ccgo_ts+7079, int32FromInt32(522)) != 0 {
				__builtin_trap(tls)
			}
			(*b2Joint)(unsafe.Pointer(headJoint)).IslandPrev = (*b2Island)(unsafe.Pointer(rootIsland)).TailJoint
			(*b2Island)(unsafe.Pointer(rootIsland)).TailJoint = (*b2Island)(unsafe.Pointer(island)).TailJoint
			*(*int32)(unsafe.Add(nil, rootIsland+44)) += (*b2Island)(unsafe.Pointer(island)).JointCount
		}
	}
	// Track removed constraints
	*(*int32)(unsafe.Add(nil, rootIsland+52)) += (*b2Island)(unsafe.Pointer(island)).ConstraintRemoveCount
	b2ValidateIsland(tls, world, rootId)
}

// C documentation
//
//	// Iterate over all awake islands and merge any that need merging
//	// Islands that get merged into a root island will be removed from the awake island array
//	// and returned to the pool.
//	// todo this might be faster if b2IslandSim held the connectivity data
func b2MergeAwakeIslands(tls *_Stack, world uintptr) {
	var awakeIslandCount, i, i1, islandId, islandId1, rootId, v11, v2, v7 int32
	var awakeSet, island, island1, islandSims, parent, rootIsland, v1, v10, v12, v3, v6, v8 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = awakeIslandCount, awakeSet, i, i1, island, island1, islandId, islandId1, islandSims, parent, rootId, rootIsland, v1, v10, v11, v12, v2, v3, v6, v7, v8
	v1 = world + 1064
	v2 = int32(b2_awakeSet)
	if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*88
	goto _4
_4:
	awakeSet = v3
	islandSims = (*b2SolverSet)(unsafe.Pointer(awakeSet)).IslandSims.Data
	awakeIslandCount = (*b2SolverSet)(unsafe.Pointer(awakeSet)).IslandSims.Count
	// Step 1: Ensure every child island points to its root island. This avoids merging a child island with
	// a parent island that has already been merged with a grand-parent island.
	i = 0
	for {
		if !(i < awakeIslandCount) {
			break
		}
		islandId = (*(*b2IslandSim)(unsafe.Add(nil, islandSims+uintptr(i)*4))).IslandId
		v6 = world + 1184
		v7 = islandId
		if !(0 <= v7 && v7 < (*b2IslandArray)(unsafe.Pointer(v6)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
			__builtin_trap(tls)
		}
		v8 = (*b2IslandArray)(unsafe.Pointer(v6)).Data + uintptr(v7)*56
		goto _9
	_9:
		island = v8
		// find the root island
		rootId = islandId
		rootIsland = island
		for (*b2Island)(unsafe.Pointer(rootIsland)).ParentIsland != -int32(1) {
			v10 = world + 1184
			v11 = (*b2Island)(unsafe.Pointer(rootIsland)).ParentIsland
			if !(0 <= v11 && v11 < (*b2IslandArray)(unsafe.Pointer(v10)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
				__builtin_trap(tls)
			}
			v12 = (*b2IslandArray)(unsafe.Pointer(v10)).Data + uintptr(v11)*56
			goto _13
		_13:
			parent = v12
			if (*b2Island)(unsafe.Pointer(parent)).ParentIsland != -int32(1) {
				// path compression
				(*b2Island)(unsafe.Pointer(rootIsland)).ParentIsland = (*b2Island)(unsafe.Pointer(parent)).ParentIsland
			}
			rootId = (*b2Island)(unsafe.Pointer(rootIsland)).ParentIsland
			rootIsland = parent
		}
		if rootIsland != island {
			(*b2Island)(unsafe.Pointer(island)).ParentIsland = rootId
		}
		goto _5
	_5:
		;
		i = i + 1
	}
	// Step 2: merge every awake island into its parent (which must be a root island)
	// Reverse to support removal from awake array.
	i1 = awakeIslandCount - int32(1)
	for {
		if !(i1 >= 0) {
			break
		}
		islandId1 = (*(*b2IslandSim)(unsafe.Add(nil, islandSims+uintptr(i1)*4))).IslandId
		v1 = world + 1184
		v2 = islandId1
		if !(0 <= v2 && v2 < (*b2IslandArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2IslandArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*56
		goto _18
	_18:
		island1 = v3
		if (*b2Island)(unsafe.Pointer(island1)).ParentIsland == -int32(1) {
			goto _14
		}
		b2MergeIsland(tls, world, island1)
		// this call does a remove swap from the end of the island sim array
		b2DestroyIsland(tls, world, islandId1)
		goto _14
	_14:
		;
		i1 = i1 - 1
	}
	b2ValidateConnectivity(tls, world)
}

func b2SplitIsland(tls *_Stack, world uintptr, baseId int32) {
	var alloc, baseIsland, bodies, body, body1, bodyIds, contact, contact1, island, joint, joint1, otherBody, otherBody1, seed, stack, tailContact, tailJoint, v1, v3 uintptr
	var bodyCount, bodyId, contactId, contactKey, edgeIndex, edgeIndex1, i, index3, islandId, jointId, jointKey, nextBody, nextContactId, nextJoint, otherBodyId, otherBodyId1, otherEdgeIndex, otherEdgeIndex1, seedIndex, setIndex, stackCount, v2 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = alloc, baseIsland, bodies, body, body1, bodyCount, bodyId, bodyIds, contact, contact1, contactId, contactKey, edgeIndex, edgeIndex1, i, index3, island, islandId, joint, joint1, jointId, jointKey, nextBody, nextContactId, nextJoint, otherBody, otherBody1, otherBodyId, otherBodyId1, otherEdgeIndex, otherEdgeIndex1, seed, seedIndex, setIndex, stack, stackCount, tailContact, tailJoint, v1, v2, v3
	v1 = world + 1184
	v2 = baseId
	if !(0 <= v2 && v2 < (*b2IslandArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2IslandArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*56
	goto _4
_4:
	baseIsland = v3
	setIndex = (*b2Island)(unsafe.Pointer(baseIsland)).SetIndex
	if setIndex != int32(b2_awakeSet) {
		// can only split awake island
		return
	}
	if (*b2Island)(unsafe.Pointer(baseIsland)).ConstraintRemoveCount == 0 {
		// this island doesn't need to be split
		return
	}
	b2ValidateIsland(tls, world, baseId)
	bodyCount = (*b2Island)(unsafe.Pointer(baseIsland)).BodyCount
	bodies = (*b2World)(unsafe.Pointer(world)).Bodies.Data
	alloc = world
	// No lock is needed because I ensure the allocator is not used while this task is active.
	stack = b2AllocateArenaItem(tls, alloc, int32FromUint64(uint64FromInt32(bodyCount)*uint64(4)), __ccgo_ts+8787)
	bodyIds = b2AllocateArenaItem(tls, alloc, int32FromUint64(uint64FromInt32(bodyCount)*uint64(4)), __ccgo_ts+8800)
	// Build array containing all body indices from base island. These
	// serve as seed bodies for the depth first search (DFS).
	index3 = 0
	nextBody = (*b2Island)(unsafe.Pointer(baseIsland)).HeadBody
	for nextBody != -int32(1) {
		v2 = index3
		index3 = index3 + 1
		*(*int32)(unsafe.Add(nil, bodyIds+uintptr(v2)*4)) = nextBody
		body = bodies + uintptr(nextBody)*128
		// Clear visitation mark
		(*b2Body)(unsafe.Pointer(body)).IsMarked = boolUint8(false1 != 0)
		nextBody = (*b2Body)(unsafe.Pointer(body)).IslandNext
	}
	if !(index3 == bodyCount) && b2InternalAssertFcn(tls, __ccgo_ts+8809, __ccgo_ts+7079, int32FromInt32(644)) != 0 {
		__builtin_trap(tls)
	}
	// Clear contact island flags. Only need to consider contacts
	// already in the base island.
	nextContactId = (*b2Island)(unsafe.Pointer(baseIsland)).HeadContact
	for nextContactId != -int32(1) {
		v1 = world + 1144
		v2 = nextContactId
		if !(0 <= v2 && v2 < (*b2ContactArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2ContactArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*68
		goto _9
	_9:
		contact = v3
		(*b2Contact)(unsafe.Pointer(contact)).IsMarked = boolUint8(false1 != 0)
		nextContactId = (*b2Contact)(unsafe.Pointer(contact)).IslandNext
	}
	// Clear joint island flags.
	nextJoint = (*b2Island)(unsafe.Pointer(baseIsland)).HeadJoint
	for nextJoint != -int32(1) {
		v1 = world + 1104
		v2 = nextJoint
		if !(0 <= v2 && v2 < (*b2JointArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2JointArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*72
		goto _13
	_13:
		joint = v3
		(*b2Joint)(unsafe.Pointer(joint)).IsMarked = boolUint8(false1 != 0)
		nextJoint = (*b2Joint)(unsafe.Pointer(joint)).IslandNext
	}
	// Done with the base split island.
	b2DestroyIsland(tls, world, baseId)
	// Each island is found as a depth first search starting from a seed body
	i = 0
	for {
		if !(i < bodyCount) {
			break
		}
		seedIndex = *(*int32)(unsafe.Add(nil, bodyIds+uintptr(i)*4))
		seed = bodies + uintptr(seedIndex)*128
		if !((*b2Body)(unsafe.Pointer(seed)).SetIndex == setIndex) && b2InternalAssertFcn(tls, __ccgo_ts+8828, __ccgo_ts+7079, int32FromInt32(673)) != 0 {
			__builtin_trap(tls)
		}
		if int32FromUint8((*b2Body)(unsafe.Pointer(seed)).IsMarked) == int32(true1) {
			// The body has already been visited
			goto _14
		}
		stackCount = 0
		v2 = stackCount
		stackCount = stackCount + 1
		*(*int32)(unsafe.Add(nil, stack+uintptr(v2)*4)) = seedIndex
		(*b2Body)(unsafe.Pointer(seed)).IsMarked = boolUint8(true1 != 0)
		// Create new island
		// No lock needed because only a single island can split per time step. No islands are being used during the constraint
		// solve. However, islands are touched during body finalization.
		island = b2CreateIsland(tls, world, setIndex)
		islandId = (*b2Island)(unsafe.Pointer(island)).IslandId
		// Perform a depth first search (DFS) on the constraint graph.
		for stackCount > 0 {
			stackCount = stackCount - 1
			v2 = stackCount
			// Grab the next body off the stack and add it to the island.
			bodyId = *(*int32)(unsafe.Add(nil, stack+uintptr(v2)*4))
			body1 = bodies + uintptr(bodyId)*128
			if !((*b2Body)(unsafe.Pointer(body1)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+1662, __ccgo_ts+7079, int32FromInt32(698)) != 0 {
				__builtin_trap(tls)
			}
			if !(int32FromUint8((*b2Body)(unsafe.Pointer(body1)).IsMarked) == int32FromInt32(true1)) && b2InternalAssertFcn(tls, __ccgo_ts+8855, __ccgo_ts+7079, int32FromInt32(699)) != 0 {
				__builtin_trap(tls)
			}
			// Add body to island
			(*b2Body)(unsafe.Pointer(body1)).IslandId = islandId
			if (*b2Island)(unsafe.Pointer(island)).TailBody != -int32(1) {
				(*(*b2Body)(unsafe.Add(nil, bodies+uintptr((*b2Island)(unsafe.Pointer(island)).TailBody)*128))).IslandNext = bodyId
			}
			(*b2Body)(unsafe.Pointer(body1)).IslandPrev = (*b2Island)(unsafe.Pointer(island)).TailBody
			(*b2Body)(unsafe.Pointer(body1)).IslandNext = -int32(1)
			(*b2Island)(unsafe.Pointer(island)).TailBody = bodyId
			if (*b2Island)(unsafe.Pointer(island)).HeadBody == -int32(1) {
				(*b2Island)(unsafe.Pointer(island)).HeadBody = bodyId
			}
			*(*int32)(unsafe.Add(nil, island+20)) += int32(1)
			// Search all contacts connected to this body.
			contactKey = (*b2Body)(unsafe.Pointer(body1)).HeadContactKey
			for contactKey != -int32(1) {
				contactId = contactKey >> int32(1)
				edgeIndex = contactKey & int32(1)
				v1 = world + 1144
				v2 = contactId
				if !(0 <= v2 && v2 < (*b2ContactArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
					__builtin_trap(tls)
				}
				v3 = (*b2ContactArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*68
				goto _20
			_20:
				contact1 = v3
				if !((*b2Contact)(unsafe.Pointer(contact1)).ContactId == contactId) && b2InternalAssertFcn(tls, __ccgo_ts+8878, __ccgo_ts+7079, int32FromInt32(726)) != 0 {
					__builtin_trap(tls)
				}
				// Next key
				contactKey = (*(*b2ContactEdge)(unsafe.Add(nil, contact1+12+uintptr(edgeIndex)*12))).NextKey
				// Has this contact already been added to this island?
				if (*b2Contact)(unsafe.Pointer(contact1)).IsMarked != 0 {
					continue
				}
				// Is this contact enabled and touching?
				if (*b2Contact)(unsafe.Pointer(contact1)).Flags&uint32(b2_contactTouchingFlag) == uint32(0) {
					continue
				}
				(*b2Contact)(unsafe.Pointer(contact1)).IsMarked = boolUint8(true1 != 0)
				otherEdgeIndex = edgeIndex ^ int32(1)
				otherBodyId = (*(*b2ContactEdge)(unsafe.Add(nil, contact1+12+uintptr(otherEdgeIndex)*12))).BodyId
				otherBody = bodies + uintptr(otherBodyId)*128
				// Maybe add other body to stack
				if int32FromUint8((*b2Body)(unsafe.Pointer(otherBody)).IsMarked) == false1 && (*b2Body)(unsafe.Pointer(otherBody)).SetIndex != int32(b2_staticSet) {
					if !(stackCount < bodyCount) && b2InternalAssertFcn(tls, __ccgo_ts+8910, __ccgo_ts+7079, int32FromInt32(752)) != 0 {
						__builtin_trap(tls)
					}
					v2 = stackCount
					stackCount = stackCount + 1
					*(*int32)(unsafe.Add(nil, stack+uintptr(v2)*4)) = otherBodyId
					(*b2Body)(unsafe.Pointer(otherBody)).IsMarked = boolUint8(true1 != 0)
				}
				// Add contact to island
				(*b2Contact)(unsafe.Pointer(contact1)).IslandId = islandId
				if (*b2Island)(unsafe.Pointer(island)).TailContact != -int32(1) {
					v1 = world + 1144
					v2 = (*b2Island)(unsafe.Pointer(island)).TailContact
					if !(0 <= v2 && v2 < (*b2ContactArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
						__builtin_trap(tls)
					}
					v3 = (*b2ContactArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*68
					goto _25
				_25:
					tailContact = v3
					(*b2Contact)(unsafe.Pointer(tailContact)).IslandNext = contactId
				}
				(*b2Contact)(unsafe.Pointer(contact1)).IslandPrev = (*b2Island)(unsafe.Pointer(island)).TailContact
				(*b2Contact)(unsafe.Pointer(contact1)).IslandNext = -int32(1)
				(*b2Island)(unsafe.Pointer(island)).TailContact = contactId
				if (*b2Island)(unsafe.Pointer(island)).HeadContact == -int32(1) {
					(*b2Island)(unsafe.Pointer(island)).HeadContact = contactId
				}
				*(*int32)(unsafe.Add(nil, island+32)) += int32(1)
			}
			// Search all joints connect to this body.
			jointKey = (*b2Body)(unsafe.Pointer(body1)).HeadJointKey
			for jointKey != -int32(1) {
				jointId = jointKey >> int32(1)
				edgeIndex1 = jointKey & int32(1)
				v1 = world + 1104
				v2 = jointId
				if !(0 <= v2 && v2 < (*b2JointArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
					__builtin_trap(tls)
				}
				v3 = (*b2JointArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*72
				goto _29
			_29:
				joint1 = v3
				if !((*b2Joint)(unsafe.Pointer(joint1)).JointId == jointId) && b2InternalAssertFcn(tls, __ccgo_ts+8933, __ccgo_ts+7079, int32FromInt32(784)) != 0 {
					__builtin_trap(tls)
				}
				// Next key
				jointKey = (*(*b2JointEdge)(unsafe.Add(nil, joint1+20+uintptr(edgeIndex1)*12))).NextKey
				// Has this joint already been added to this island?
				if (*b2Joint)(unsafe.Pointer(joint1)).IsMarked != 0 {
					continue
				}
				(*b2Joint)(unsafe.Pointer(joint1)).IsMarked = boolUint8(true1 != 0)
				otherEdgeIndex1 = edgeIndex1 ^ int32(1)
				otherBodyId1 = (*(*b2JointEdge)(unsafe.Add(nil, joint1+20+uintptr(otherEdgeIndex1)*12))).BodyId
				otherBody1 = bodies + uintptr(otherBodyId1)*128
				// Don't simulate joints connected to disabled bodies.
				if (*b2Body)(unsafe.Pointer(otherBody1)).SetIndex == int32(b2_disabledSet) {
					continue
				}
				// Maybe add other body to stack
				if int32FromUint8((*b2Body)(unsafe.Pointer(otherBody1)).IsMarked) == false1 && (*b2Body)(unsafe.Pointer(otherBody1)).SetIndex == int32(b2_awakeSet) {
					if !(stackCount < bodyCount) && b2InternalAssertFcn(tls, __ccgo_ts+8910, __ccgo_ts+7079, int32FromInt32(810)) != 0 {
						__builtin_trap(tls)
					}
					v2 = stackCount
					stackCount = stackCount + 1
					*(*int32)(unsafe.Add(nil, stack+uintptr(v2)*4)) = otherBodyId1
					(*b2Body)(unsafe.Pointer(otherBody1)).IsMarked = boolUint8(true1 != 0)
				}
				// Add joint to island
				(*b2Joint)(unsafe.Pointer(joint1)).IslandId = islandId
				if (*b2Island)(unsafe.Pointer(island)).TailJoint != -int32(1) {
					v1 = world + 1104
					v2 = (*b2Island)(unsafe.Pointer(island)).TailJoint
					if !(0 <= v2 && v2 < (*b2JointArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
						__builtin_trap(tls)
					}
					v3 = (*b2JointArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*72
					goto _34
				_34:
					tailJoint = v3
					(*b2Joint)(unsafe.Pointer(tailJoint)).IslandNext = jointId
				}
				(*b2Joint)(unsafe.Pointer(joint1)).IslandPrev = (*b2Island)(unsafe.Pointer(island)).TailJoint
				(*b2Joint)(unsafe.Pointer(joint1)).IslandNext = -int32(1)
				(*b2Island)(unsafe.Pointer(island)).TailJoint = jointId
				if (*b2Island)(unsafe.Pointer(island)).HeadJoint == -int32(1) {
					(*b2Island)(unsafe.Pointer(island)).HeadJoint = jointId
				}
				*(*int32)(unsafe.Add(nil, island+44)) += int32(1)
			}
		}
		b2ValidateIsland(tls, world, islandId)
		goto _14
	_14:
		;
		i = i + 1
	}
	b2FreeArenaItem(tls, alloc, bodyIds)
	b2FreeArenaItem(tls, alloc, stack)
}

// C documentation
//
//	// Split an island because some contacts and/or joints have been removed.
//	// This is called during the constraint solve while islands are not being touched. This uses DFS and touches a lot of memory,
//	// so it can be quite slow.
//	// Note: contacts/joints connected to static bodies must belong to an island but don't affect island connectivity
//	// Note: static bodies are never in an island
//	// Note: this task interacts with some allocators without locks under the assumption that no other tasks
//	// are interacting with these data structures.
func b2SplitIslandTask(tls *_Stack, startIndex int32, endIndex int32, threadIndex uint32, context uintptr) {
	var ticks uint64_t
	var world uintptr
	_, _ = ticks, world
	_ = uint64FromInt64(4)
	ticks = b2GetTicks(tls)
	world = context
	if !((*b2World)(unsafe.Pointer(world)).SplitIslandId != -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+8959, __ccgo_ts+7079, int32FromInt32(858)) != 0 {
		__builtin_trap(tls)
	}
	b2SplitIsland(tls, world, (*b2World)(unsafe.Pointer(world)).SplitIslandId)
	(*b2World)(unsafe.Pointer(world)).Profile.SplitIslands += b2GetMilliseconds(tls, ticks)
}

func b2ValidateIsland(tls *_Stack, world uintptr, islandId int32) {
	_ = uint64FromInt64(4)
	_ = uint64FromInt64(4)
}

const _B2_JOINT_CONSTRAINT_DAMPING_RATIO1 = 2
const _B2_JOINT_CONSTRAINT_HERTZ1 = 60
const _UINT64_MAX5 = 18446744073709551615

var b2_identityBodyState7 = b2BodyState{
	DeltaRotation: Rot{
		C: float32FromFloat32(1),
	},
}

func b2DefaultExplosionDef(tls *_Stack) (r ExplosionDef) {
	var def ExplosionDef
	_ = def
	def = ExplosionDef{}
	def.MaskBits = uint64FromUint64(0xffffffffffffffff)
	return def
}

type b2JointPair struct {
	Joint    uintptr
	JointSim uintptr
}

func b2DestroyContactsBetweenBodies(tls *_Stack, world uintptr, bodyA uintptr, bodyB uintptr) {
	var contact, v1, v3 uintptr
	var contactId, contactKey, edgeIndex, otherBodyId, otherEdgeIndex, v2 int32
	var wakeBodies uint8
	_, _, _, _, _, _, _, _, _, _ = contact, contactId, contactKey, edgeIndex, otherBodyId, otherEdgeIndex, wakeBodies, v1, v2, v3
	// use the smaller of the two contact lists
	if (*b2Body)(unsafe.Pointer(bodyA)).ContactCount < (*b2Body)(unsafe.Pointer(bodyB)).ContactCount {
		contactKey = (*b2Body)(unsafe.Pointer(bodyA)).HeadContactKey
		otherBodyId = (*b2Body)(unsafe.Pointer(bodyB)).Id
	} else {
		contactKey = (*b2Body)(unsafe.Pointer(bodyB)).HeadContactKey
		otherBodyId = (*b2Body)(unsafe.Pointer(bodyA)).Id
	}
	// no need to wake bodies when a joint removes collision between them
	wakeBodies = boolUint8(false1 != 0)
	// destroy the contacts
	for contactKey != -int32(1) {
		contactId = contactKey >> int32(1)
		edgeIndex = contactKey & int32(1)
		v1 = world + 1144
		v2 = contactId
		if !(0 <= v2 && v2 < (*b2ContactArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2ContactArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*68
		goto _4
	_4:
		contact = v3
		contactKey = (*(*b2ContactEdge)(unsafe.Add(nil, contact+12+uintptr(edgeIndex)*12))).NextKey
		otherEdgeIndex = edgeIndex ^ int32(1)
		if (*(*b2ContactEdge)(unsafe.Add(nil, contact+12+uintptr(otherEdgeIndex)*12))).BodyId == otherBodyId {
			// Careful, this removes the contact from the current doubly linked list
			b2DestroyContact(tls, world, contact, wakeBodies)
		}
	}
	b2ValidateSolverSets(tls, world)
}

const _B2_JOINT_CONSTRAINT_DAMPING_RATIO2 = "2.0f"
const _B2_JOINT_CONSTRAINT_HERTZ2 = "60.0f"
const _FLT_MAX6 = 3.4028234663852886e+38
const _UINT64_MAX6 = "0xffffffffffffffffu"

func b2MakeCapsule(tls *_Stack, p1 Vec2, p2 Vec2, radius float32) (r Polygon) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var axis, d, n, normal, v1, v2, v4, v6 Vec2
	var invLength, length, v3 float32
	var _ /* shape at bp+0 */ Polygon
	_, _, _, _, _, _, _, _, _, _, _ = axis, d, invLength, length, n, normal, v1, v2, v3, v4, v6
	*(*Polygon)(unsafe.Add(nil, bp)) = Polygon{}
	*(*Vec2)(unsafe.Add(nil, bp)) = p1
	*(*Vec2)(unsafe.Add(nil, bp+1*8)) = p2
	v1 = p1
	v2 = p2
	v3 = float32FromFloat32(0.5)
	v4 = Vec2{
		X: float32((float32FromFloat32(1)-v3)*v1.X) + float32(v3*v2.X),
		Y: float32((float32FromFloat32(1)-v3)*v1.Y) + float32(v3*v2.Y),
	}
	goto _5
_5:
	(*(*Polygon)(unsafe.Add(nil, bp))).Centroid = v4
	v1 = p2
	v2 = p1
	v4 = Vec2{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
	goto _9
_9:
	d = v4
	v6 = d
	v3 = float32(v6.X*v6.X) + float32(v6.Y*v6.Y)
	goto _12
_12:
	;
	if !(v3 > float32FromFloat32(1.1920928955078125e-07)) && b2InternalAssertFcn(tls, __ccgo_ts+9901, __ccgo_ts+9936, int32FromInt32(24)) != 0 {
		__builtin_trap(tls)
	}
	v1 = d
	length = sqrtf(tls, float32(v1.X*v1.X)+float32(v1.Y*v1.Y))
	if length < float32FromFloat32(1.1920928955078125e-07) {
		v2 = Vec2{}
		goto _15
	}
	invLength = float32FromFloat32(1) / length
	n = Vec2{
		X: float32(invLength * v1.X),
		Y: float32(invLength * v1.Y),
	}
	v2 = n
	goto _15
_15:
	axis = v2
	v4 = axis
	v6 = Vec2{
		X: v4.Y,
		Y: -v4.X,
	}
	goto _18
_18:
	normal = v6
	*(*Vec2)(unsafe.Add(nil, bp+64)) = normal
	v1 = normal
	v2 = Vec2{
		X: -v1.X,
		Y: -v1.Y,
	}
	goto _21
_21:
	*(*Vec2)(unsafe.Add(nil, bp+64+1*8)) = v2
	(*(*Polygon)(unsafe.Add(nil, bp))).Count = int32(2)
	(*(*Polygon)(unsafe.Add(nil, bp))).Radius = radius
	return *(*Polygon)(unsafe.Add(nil, bp))
}

// C documentation
//
//	// Polygon clipper used to compute contact points when there are potentially two contact points.
func b2ClipPolygons(tls *_Stack, polyA uintptr, polyB uintptr, edgeA int32, edgeB int32, flip uint8) (r Manifold) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var cp, cp1, poly1, poly2 uintptr
	var i11, i12, i21, i22, v1 int32
	var lower1, lower2, r1, r2, radius, separationLower, separationUpper, upper1, upper2, v15, v23, v31, v5 float32
	var normal, tangent, v11, v12, v21, v22, vLower, vUpper, v10, v111, v13, v14, v17, v18, v19, v211, v221, v25, v26, v27, v29, v30, v6, v7, v9 Vec2
	var _ /* manifold at bp+0 */ Manifold
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = cp, cp1, i11, i12, i21, i22, lower1, lower2, normal, poly1, poly2, r1, r2, radius, separationLower, separationUpper, tangent, upper1, upper2, v11, v12, v21, v22, vLower, vUpper, v1, v10, v111, v13, v14, v15, v17, v18, v19, v211, v221, v23, v25, v26, v27, v29, v30, v31, v5, v6, v7, v9
	*(*Manifold)(unsafe.Add(nil, bp)) = Manifold{}
	if flip != 0 {
		poly1 = polyB
		poly2 = polyA
		i11 = edgeB
		if edgeB+int32(1) < (*Polygon)(unsafe.Pointer(polyB)).Count {
			v1 = edgeB + int32(1)
		} else {
			v1 = 0
		}
		i12 = v1
		i21 = edgeA
		if edgeA+int32(1) < (*Polygon)(unsafe.Pointer(polyA)).Count {
			v1 = edgeA + int32(1)
		} else {
			v1 = 0
		}
		i22 = v1
	} else {
		poly1 = polyA
		poly2 = polyB
		i11 = edgeA
		if edgeA+int32(1) < (*Polygon)(unsafe.Pointer(polyA)).Count {
			v1 = edgeA + int32(1)
		} else {
			v1 = 0
		}
		i12 = v1
		i21 = edgeB
		if edgeB+int32(1) < (*Polygon)(unsafe.Pointer(polyB)).Count {
			v1 = edgeB + int32(1)
		} else {
			v1 = 0
		}
		i22 = v1
	}
	normal = *(*Vec2)(unsafe.Add(nil, poly1+64+uintptr(i11)*8))
	// Reference edge vertices
	v11 = *(*Vec2)(unsafe.Add(nil, poly1+uintptr(i11)*8))
	v12 = *(*Vec2)(unsafe.Add(nil, poly1+uintptr(i12)*8))
	// Incident edge vertices
	v21 = *(*Vec2)(unsafe.Add(nil, poly2+uintptr(i21)*8))
	v22 = *(*Vec2)(unsafe.Add(nil, poly2+uintptr(i22)*8))
	v5 = float32FromFloat32(1)
	v6 = normal
	v7 = Vec2{
		X: float32(-v5 * v6.Y),
		Y: float32(v5 * v6.X),
	}
	goto _8
_8:
	tangent = v7
	lower1 = float32FromFloat32(0)
	v9 = v12
	v10 = v11
	v111 = Vec2{
		X: v9.X - v10.X,
		Y: v9.Y - v10.Y,
	}
	goto _12
_12:
	v13 = v111
	v14 = tangent
	v15 = float32(v13.X*v14.X) + float32(v13.Y*v14.Y)
	goto _16
_16:
	upper1 = v15
	v17 = v21
	v18 = v11
	v19 = Vec2{
		X: v17.X - v18.X,
		Y: v17.Y - v18.Y,
	}
	goto _20
_20:
	v211 = v19
	v221 = tangent
	v23 = float32(v211.X*v221.X) + float32(v211.Y*v221.Y)
	goto _24
_24:
	// Incident edge points opposite of tangent due to CCW winding
	upper2 = v23
	v25 = v22
	v26 = v11
	v27 = Vec2{
		X: v25.X - v26.X,
		Y: v25.Y - v26.Y,
	}
	goto _28
_28:
	v29 = v27
	v30 = tangent
	v31 = float32(v29.X*v30.X) + float32(v29.Y*v30.Y)
	goto _32
_32:
	lower2 = v31
	// Are the segments disjoint?
	if upper2 < lower1 || upper1 < lower2 {
		return *(*Manifold)(unsafe.Add(nil, bp))
	}
	if lower2 < lower1 && upper2-lower2 > float32FromFloat32(1.1920928955078125e-07) {
		v6 = v22
		v7 = v21
		v5 = (lower1 - lower2) / (upper2 - lower2)
		v9 = Vec2{
			X: float32((float32FromFloat32(1)-v5)*v6.X) + float32(v5*v7.X),
			Y: float32((float32FromFloat32(1)-v5)*v6.Y) + float32(v5*v7.Y),
		}
		goto _37
	_37:
		vLower = v9
	} else {
		vLower = v22
	}
	if upper2 > upper1 && upper2-lower2 > float32FromFloat32(1.1920928955078125e-07) {
		v6 = v22
		v7 = v21
		v5 = (upper1 - lower2) / (upper2 - lower2)
		v9 = Vec2{
			X: float32((float32FromFloat32(1)-v5)*v6.X) + float32(v5*v7.X),
			Y: float32((float32FromFloat32(1)-v5)*v6.Y) + float32(v5*v7.Y),
		}
		goto _42
	_42:
		vUpper = v9
	} else {
		vUpper = v21
	}
	v6 = vLower
	v7 = v11
	v9 = Vec2{
		X: v6.X - v7.X,
		Y: v6.Y - v7.Y,
	}
	goto _46
_46:
	v10 = v9
	v111 = normal
	v5 = float32(v10.X*v111.X) + float32(v10.Y*v111.Y)
	goto _50
_50:
	// todo vLower can be very close to vUpper, reduce to one point?
	separationLower = v5
	v13 = vUpper
	v14 = v11
	v17 = Vec2{
		X: v13.X - v14.X,
		Y: v13.Y - v14.Y,
	}
	goto _54
_54:
	v18 = v17
	v19 = normal
	v15 = float32(v18.X*v19.X) + float32(v18.Y*v19.Y)
	goto _58
_58:
	separationUpper = v15
	r1 = (*Polygon)(unsafe.Pointer(poly1)).Radius
	r2 = (*Polygon)(unsafe.Pointer(poly2)).Radius
	// Put contact points at midpoint, accounting for radii
	v211 = vLower
	v23 = float32(float32FromFloat32(0.5) * (r1 - r2 - separationLower))
	v221 = normal
	v25 = Vec2{
		X: v211.X + float32(v23*v221.X),
		Y: v211.Y + float32(v23*v221.Y),
	}
	goto _63
_63:
	vLower = v25
	v6 = vUpper
	v5 = float32(float32FromFloat32(0.5) * (r1 - r2 - separationUpper))
	v7 = normal
	v9 = Vec2{
		X: v6.X + float32(v5*v7.X),
		Y: v6.Y + float32(v5*v7.Y),
	}
	goto _68
_68:
	vUpper = v9
	radius = r1 + r2
	if int32FromUint8(flip) == false1 {
		(*(*Manifold)(unsafe.Add(nil, bp))).Normal = normal
		cp = bp + 12 + uintptr(0)*48
		(*ManifoldPoint)(unsafe.Pointer(cp)).AnchorA = vLower
		(*ManifoldPoint)(unsafe.Pointer(cp)).Separation = separationLower - radius
		(*ManifoldPoint)(unsafe.Pointer(cp)).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(i11))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(i22)))
		(*(*Manifold)(unsafe.Add(nil, bp))).PointCount += int32(1)
		cp = cp + uintptr(1)*48
		(*ManifoldPoint)(unsafe.Pointer(cp)).AnchorA = vUpper
		(*ManifoldPoint)(unsafe.Pointer(cp)).Separation = separationUpper - radius
		(*ManifoldPoint)(unsafe.Pointer(cp)).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(i12))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(i21)))
		(*(*Manifold)(unsafe.Add(nil, bp))).PointCount += int32(1)
	} else {
		v6 = normal
		v7 = Vec2{
			X: -v6.X,
			Y: -v6.Y,
		}
		goto _71
	_71:
		(*(*Manifold)(unsafe.Add(nil, bp))).Normal = v7
		cp1 = bp + 12 + uintptr(0)*48
		(*ManifoldPoint)(unsafe.Pointer(cp1)).AnchorA = vUpper
		(*ManifoldPoint)(unsafe.Pointer(cp1)).Separation = separationUpper - radius
		(*ManifoldPoint)(unsafe.Pointer(cp1)).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(i21))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(i12)))
		(*(*Manifold)(unsafe.Add(nil, bp))).PointCount += int32(1)
		cp1 = cp1 + uintptr(1)*48
		(*ManifoldPoint)(unsafe.Pointer(cp1)).AnchorA = vLower
		(*ManifoldPoint)(unsafe.Pointer(cp1)).Separation = separationLower - radius
		(*ManifoldPoint)(unsafe.Pointer(cp1)).Id = uint16FromInt32(int32FromUint8(uint8FromInt32(i22))<<int32FromInt32(8) | int32FromUint8(uint8FromInt32(i11)))
		(*(*Manifold)(unsafe.Add(nil, bp))).PointCount += int32(1)
	}
	return *(*Manifold)(unsafe.Add(nil, bp))
}

// C documentation
//
//	// Find the max separation between poly1 and poly2 using edge normals from poly1.
func b2FindMaxSeparation(tls *_Stack, edgeIndex uintptr, poly1 uintptr, poly2 uintptr) (r float32) {
	var bestIndex, count1, count2, i, j int32
	var maxSeparation, si, sij, v9 float32
	var n, v1, v3, v4, v5, v7, v8 Vec2
	var n1s, v1s, v2s uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bestIndex, count1, count2, i, j, maxSeparation, n, n1s, si, sij, v1, v1s, v2s, v3, v4, v5, v7, v8, v9
	count1 = (*Polygon)(unsafe.Pointer(poly1)).Count
	count2 = (*Polygon)(unsafe.Pointer(poly2)).Count
	n1s = poly1 + 64
	v1s = poly1
	v2s = poly2
	bestIndex = 0
	maxSeparation = -float32FromFloat32(3.4028234663852886e+38)
	i = 0
	for {
		if !(i < count1) {
			break
		}
		// Get poly1 normal in frame2.
		n = *(*Vec2)(unsafe.Add(nil, n1s+uintptr(i)*8))
		v1 = *(*Vec2)(unsafe.Add(nil, v1s+uintptr(i)*8))
		// Find the deepest point for normal i.
		si = float32FromFloat32(3.4028234663852886e+38)
		j = 0
		for {
			if !(j < count2) {
				break
			}
			v3 = *(*Vec2)(unsafe.Add(nil, v2s+uintptr(j)*8))
			v4 = v1
			v5 = Vec2{
				X: v3.X - v4.X,
				Y: v3.Y - v4.Y,
			}
			goto _6
		_6:
			v7 = n
			v8 = v5
			v9 = float32(v7.X*v8.X) + float32(v7.Y*v8.Y)
			goto _10
		_10:
			sij = v9
			if sij < si {
				si = sij
			}
			goto _2
		_2:
			;
			j = j + 1
		}
		if si > maxSeparation {
			maxSeparation = si
			bestIndex = i
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	*(*int32)(unsafe.Add(nil, edgeIndex)) = bestIndex
	return maxSeparation
}

func b2ClipSegments(tls *_Stack, a11 Vec2, a21 Vec2, b11 Vec2, b21 Vec2, normal Vec2, ra float32, rb float32, id1 uint16, id2 uint16) (r Manifold) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var cp, cp1 uintptr
	var lower1, lower2, radius, separationLower, separationUpper, upper1, upper2, v10, v18, v26 float32
	var tangent, vLower, vUpper, v1, v12, v13, v14, v16, v17, v2, v20, v21, v22, v24, v25, v4, v5, v6, v8, v9 Vec2
	var _ /* manifold at bp+0 */ Manifold
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = cp, cp1, lower1, lower2, radius, separationLower, separationUpper, tangent, upper1, upper2, vLower, vUpper, v1, v10, v12, v13, v14, v16, v17, v18, v2, v20, v21, v22, v24, v25, v26, v4, v5, v6, v8, v9
	*(*Manifold)(unsafe.Add(nil, bp)) = Manifold{}
	v1 = normal
	v2 = Vec2{
		X: -v1.Y,
		Y: v1.X,
	}
	goto _3
_3:
	tangent = v2
	// Barycentric coordinates of each point relative to a1 along tangent
	lower1 = float32FromFloat32(0)
	v4 = a21
	v5 = a11
	v6 = Vec2{
		X: v4.X - v5.X,
		Y: v4.Y - v5.Y,
	}
	goto _7
_7:
	v8 = v6
	v9 = tangent
	v10 = float32(v8.X*v9.X) + float32(v8.Y*v9.Y)
	goto _11
_11:
	upper1 = v10
	v12 = b11
	v13 = a11
	v14 = Vec2{
		X: v12.X - v13.X,
		Y: v12.Y - v13.Y,
	}
	goto _15
_15:
	v16 = v14
	v17 = tangent
	v18 = float32(v16.X*v17.X) + float32(v16.Y*v17.Y)
	goto _19
_19:
	// Incident edge points opposite of tangent due to CCW winding
	upper2 = v18
	v20 = b21
	v21 = a11
	v22 = Vec2{
		X: v20.X - v21.X,
		Y: v20.Y - v21.Y,
	}
	goto _23
_23:
	v24 = v22
	v25 = tangent
	v26 = float32(v24.X*v25.X) + float32(v24.Y*v25.Y)
	goto _27
_27:
	lower2 = v26
	// Do segments overlap?
	if upper2 < lower1 || upper1 < lower2 {
		return *(*Manifold)(unsafe.Add(nil, bp))
	}
	if lower2 < lower1 && upper2-lower2 > float32FromFloat32(1.1920928955078125e-07) {
		v1 = b21
		v2 = b11
		v10 = (lower1 - lower2) / (upper2 - lower2)
		v4 = Vec2{
			X: float32((float32FromFloat32(1)-v10)*v1.X) + float32(v10*v2.X),
			Y: float32((float32FromFloat32(1)-v10)*v1.Y) + float32(v10*v2.Y),
		}
		goto _32
	_32:
		vLower = v4
	} else {
		vLower = b21
	}
	if upper2 > upper1 && upper2-lower2 > float32FromFloat32(1.1920928955078125e-07) {
		v1 = b21
		v2 = b11
		v10 = (upper1 - lower2) / (upper2 - lower2)
		v4 = Vec2{
			X: float32((float32FromFloat32(1)-v10)*v1.X) + float32(v10*v2.X),
			Y: float32((float32FromFloat32(1)-v10)*v1.Y) + float32(v10*v2.Y),
		}
		goto _37
	_37:
		vUpper = v4
	} else {
		vUpper = b11
	}
	v1 = vLower
	v2 = a11
	v4 = Vec2{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
	goto _41
_41:
	v5 = v4
	v6 = normal
	v10 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y)
	goto _45
_45:
	// todo vLower can be very close to vUpper, reduce to one point?
	separationLower = v10
	v8 = vUpper
	v9 = a11
	v12 = Vec2{
		X: v8.X - v9.X,
		Y: v8.Y - v9.Y,
	}
	goto _49
_49:
	v13 = v12
	v14 = normal
	v18 = float32(v13.X*v14.X) + float32(v13.Y*v14.Y)
	goto _53
_53:
	separationUpper = v18
	// Put contact points at midpoint, accounting for radii
	v16 = vLower
	v26 = float32(float32FromFloat32(0.5) * (ra - rb - separationLower))
	v17 = normal
	v20 = Vec2{
		X: v16.X + float32(v26*v17.X),
		Y: v16.Y + float32(v26*v17.Y),
	}
	goto _58
_58:
	vLower = v20
	v1 = vUpper
	v10 = float32(float32FromFloat32(0.5) * (ra - rb - separationUpper))
	v2 = normal
	v4 = Vec2{
		X: v1.X + float32(v10*v2.X),
		Y: v1.Y + float32(v10*v2.Y),
	}
	goto _63
_63:
	vUpper = v4
	radius = ra + rb
	(*(*Manifold)(unsafe.Add(nil, bp))).Normal = normal
	cp = bp + 12 + uintptr(0)*48
	(*ManifoldPoint)(unsafe.Pointer(cp)).AnchorA = vLower
	(*ManifoldPoint)(unsafe.Pointer(cp)).Separation = separationLower - radius
	(*ManifoldPoint)(unsafe.Pointer(cp)).Id = id1
	cp1 = bp + 12 + uintptr(1)*48
	(*ManifoldPoint)(unsafe.Pointer(cp1)).AnchorA = vUpper
	(*ManifoldPoint)(unsafe.Pointer(cp1)).Separation = separationUpper - radius
	(*ManifoldPoint)(unsafe.Pointer(cp1)).Id = id2
	(*(*Manifold)(unsafe.Add(nil, bp))).PointCount = int32(2)
	return *(*Manifold)(unsafe.Add(nil, bp))
}

type b2NormalType = int32

const
// This means the normal points in a direction that is non-smooth relative to a convex vertex and should be skipped
b2_normalSkip = 0
const

// This means the normal points in a direction that is smooth relative to a convex vertex and should be used for collision
b2_normalAdmit = 1
const

// This means the normal is in a region of a concave vertex and should be snapped to the segment normal
b2_normalSnap = 2

type b2ChainSegmentParams struct {
	Edge1   Vec2
	Normal0 Vec2
	Normal2 Vec2
	Convex1 uint8
	Convex2 uint8
}

// C documentation
//
//	// Evaluate Gauss map
//	// See https://box2d.org/posts/2020/06/ghost-collisions/
func b2ClassifyNormal(tls *_Stack, params b2ChainSegmentParams, normal Vec2) (r b2NormalType) {
	var sinTol, v3, v7 float32
	var v1, v2, v5, v6 Vec2
	_, _, _, _, _, _, _ = sinTol, v1, v2, v3, v5, v6, v7
	sinTol = float32FromFloat32(0.01)
	v1 = normal
	v2 = params.Edge1
	v3 = float32(v1.X*v2.X) + float32(v1.Y*v2.Y)
	goto _4
_4:
	if v3 <= float32FromFloat32(0) {
		// Normal points towards the segment tail
		if params.Convex1 != 0 {
			v5 = normal
			v6 = params.Normal0
			v7 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
			goto _8
		_8:
			if v7 > sinTol {
				return int32(b2_normalSkip)
			}
			return int32(b2_normalAdmit)
		} else {
			return int32(b2_normalSnap)
		}
	} else {
		// Normal points towards segment head
		if params.Convex2 != 0 {
			v1 = params.Normal2
			v2 = normal
			v3 = float32(v1.X*v2.Y) - float32(v1.Y*v2.X)
			goto _12
		_12:
			if v3 > sinTol {
				return int32(b2_normalSkip)
			}
			return int32(b2_normalAdmit)
		} else {
			return int32(b2_normalSnap)
		}
	}
}

const _FLT_MAX7 = 3.40282346638528859812e+38

func b2IsValidFloat(tls *_Stack, a float32) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var v1 uint32
	var _ /* __u at bp+0 */ struct {
		ｆ__i [0]uint32
		ｆ__f float32
	}
	_ = v1
	*(*float32)(unsafe.Pointer(bp)) = a
	v1 = *(*uint32)(unsafe.Pointer(bp))
	goto _2
_2:
	if boolInt32(v1&uint32(0x7fffffff) > uint32(0x7f800000)) != 0 {
		return boolUint8(false1 != 0)
	}
	*(*float32)(unsafe.Pointer(bp)) = a
	v1 = *(*uint32)(unsafe.Pointer(bp))
	goto _4
_4:
	if boolInt32(v1&uint32(0x7fffffff) == uint32(0x7f800000)) != 0 {
		return boolUint8(false1 != 0)
	}
	return boolUint8(true1 != 0)
}

func b2IsValidVec2(tls *_Stack, v Vec2) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var v1, v3 uint32
	var v5 bool
	var _ /* __u at bp+0 */ struct {
		ｆ__i [0]uint32
		ｆ__f float32
	}
	_, _, _ = v1, v3, v5
	*(*float32)(unsafe.Pointer(bp)) = v.X
	v1 = *(*uint32)(unsafe.Pointer(bp))
	goto _2
_2:
	;
	if v5 = boolInt32(v1&uint32(0x7fffffff) > uint32(0x7f800000)) != 0; !v5 {
		*(*float32)(unsafe.Pointer(bp)) = v.Y
		v3 = *(*uint32)(unsafe.Pointer(bp))
		goto _4
	_4:
	}
	if v5 || boolInt32(v3&uint32(0x7fffffff) > uint32(0x7f800000)) != 0 {
		return boolUint8(false1 != 0)
	}
	*(*float32)(unsafe.Pointer(bp)) = v.X
	v1 = *(*uint32)(unsafe.Pointer(bp))
	goto _7
_7:
	;
	if v5 = boolInt32(v1&uint32(0x7fffffff) == uint32(0x7f800000)) != 0; !v5 {
		*(*float32)(unsafe.Pointer(bp)) = v.Y
		v3 = *(*uint32)(unsafe.Pointer(bp))
		goto _9
	_9:
	}
	if v5 || boolInt32(v3&uint32(0x7fffffff) == uint32(0x7f800000)) != 0 {
		return boolUint8(false1 != 0)
	}
	return boolUint8(true1 != 0)
}

func b2IsValidRotation(tls *_Stack, q1 Rot) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var qq float32
	var v1, v3 uint32
	var v5 bool
	var v11 Rot
	var v12 uint8
	var _ /* __u at bp+0 */ struct {
		ｆ__i [0]uint32
		ｆ__f float32
	}
	_, _, _, _, _, _ = qq, v1, v11, v12, v3, v5
	*(*float32)(unsafe.Pointer(bp)) = q1.S
	v1 = *(*uint32)(unsafe.Pointer(bp))
	goto _2
_2:
	;
	if v5 = boolInt32(v1&uint32(0x7fffffff) > uint32(0x7f800000)) != 0; !v5 {
		*(*float32)(unsafe.Pointer(bp)) = q1.C
		v3 = *(*uint32)(unsafe.Pointer(bp))
		goto _4
	_4:
	}
	if v5 || boolInt32(v3&uint32(0x7fffffff) > uint32(0x7f800000)) != 0 {
		return boolUint8(false1 != 0)
	}
	*(*float32)(unsafe.Pointer(bp)) = q1.S
	v1 = *(*uint32)(unsafe.Pointer(bp))
	goto _7
_7:
	;
	if v5 = boolInt32(v1&uint32(0x7fffffff) == uint32(0x7f800000)) != 0; !v5 {
		*(*float32)(unsafe.Pointer(bp)) = q1.C
		v3 = *(*uint32)(unsafe.Pointer(bp))
		goto _9
	_9:
	}
	if v5 || boolInt32(v3&uint32(0x7fffffff) == uint32(0x7f800000)) != 0 {
		return boolUint8(false1 != 0)
	}
	v11 = q1
	qq = float32(v11.S*v11.S) + float32(v11.C*v11.C)
	v12 = boolUint8(float32FromFloat32(1)-float32FromFloat32(0.0006) < qq && qq < float32FromFloat32(1)+float32FromFloat32(0.0006))
	goto _13
_13:
	return v12
}

func b2IsValidPlane(tls *_Stack, a3 Plane) (r uint8) {
	var aa, v11, v4, v8, v9 float32
	var v1, v2, v3 Vec2
	var v12 bool
	var v6 uint8
	_, _, _, _, _, _, _, _, _, _ = aa, v1, v11, v12, v2, v3, v4, v6, v8, v9
	if v12 = b2IsValidVec2(tls, a3.Normal) != 0; v12 {
		v1 = a3.Normal
		v2 = v1
		v3 = v1
		v4 = float32(v2.X*v3.X) + float32(v2.Y*v3.Y)
		goto _5
	_5:
		aa = v4
		v8 = float32FromFloat32(1) - aa
		if v8 < float32FromInt32(0) {
			v11 = -v8
		} else {
			v11 = v8
		}
		v9 = v11
		goto _10
	_10:
		v6 = boolUint8(v9 < float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07)))
		goto _7
	_7:
	}
	return boolUint8(v12 && v6 != 0 && b2IsValidFloat(tls, a3.Offset) != 0)
}

// C documentation
//
//	// https://stackoverflow.com/questions/46210708/atan2-approximation-with-11bits-in-mantissa-on-x86with-sse2-and-armwith-vfpv4
func b2Atan2(tls *_Stack, y float32, x float32) (r1 float32) {
	var a3, ax, ay, c, mn, mx, q, r, s, t, v1, v10, v11, v13, v14, v15, v16, v18, v2, v4, v5, v6, v8, v9 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a3, ax, ay, c, mn, mx, q, r, s, t, v1, v10, v11, v13, v14, v15, v16, v18, v2, v4, v5, v6, v8, v9
	// Added check for (0,0) to match atan2f and avoid NaN
	if x == float32FromFloat32(0) && y == float32FromFloat32(0) {
		return float32FromFloat32(0)
	}
	v1 = x
	if v1 < float32FromInt32(0) {
		v4 = -v1
	} else {
		v4 = v1
	}
	v2 = v4
	goto _3
_3:
	ax = v2
	v5 = y
	if v5 < float32FromInt32(0) {
		v8 = -v5
	} else {
		v8 = v5
	}
	v6 = v8
	goto _7
_7:
	ay = v6
	v9 = ay
	v10 = ax
	if v9 > v10 {
		v13 = v9
	} else {
		v13 = v10
	}
	v11 = v13
	goto _12
_12:
	mx = v11
	v14 = ay
	v15 = ax
	if v14 < v15 {
		v18 = v14
	} else {
		v18 = v15
	}
	v16 = v18
	goto _17
_17:
	mn = v16
	a3 = mn / mx
	// Minimax polynomial approximation to atan(a) on [0,1]
	s = float32(a3 * a3)
	c = float32(s * a3)
	q = float32(s * s)
	r = float32(float32FromFloat32(0.024840285)*q) + float32FromFloat32(0.18681418)
	t = float32(-float32FromFloat32(0.094097948)*q) - float32FromFloat32(0.33213072)
	r = float32(r*s) + t
	r = float32(r*c) + a3
	// Map to full circle
	if ay > ax {
		r = float32FromFloat32(1.57079637) - r
	}
	if x < float32FromInt32(0) {
		r = float32FromFloat32(3.14159274) - r
	}
	if y < float32FromInt32(0) {
		r = -r
	}
	return r
}

var b2_identityBodyState8 = b2BodyState{
	DeltaRotation: Rot{
		C: float32FromFloat32(1),
	},
}

var b2_identityBodyState9 = b2BodyState{
	DeltaRotation: Rot{
		C: float32FromFloat32(1),
	},
}

func b2ClipVector(tls *_Stack, vector Vec2, planes uintptr, count int32) (r Vec2) {
	var plane uintptr
	var planeIndex int32
	var v, v11, v13, v14, v2, v3 Vec2
	var v10, v12, v4, v6, v7, v8 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = plane, planeIndex, v, v10, v11, v12, v13, v14, v2, v3, v4, v6, v7, v8
	v = vector
	planeIndex = 0
	for {
		if !(planeIndex < count) {
			break
		}
		plane = planes + uintptr(planeIndex)*24
		if (*CollisionPlane)(unsafe.Pointer(plane)).Push == float32FromFloat32(0) || int32FromUint8((*CollisionPlane)(unsafe.Pointer(plane)).ClipVelocity) == false1 {
			goto _1
		}
		v2 = v
		v3 = (*CollisionPlane)(unsafe.Pointer(plane)).Plane.Normal
		v4 = float32(v2.X*v3.X) + float32(v2.Y*v3.Y)
		goto _5
	_5:
		v6 = float32FromFloat32(0)
		v7 = v4
		if v6 < v7 {
			v10 = v6
		} else {
			v10 = v7
		}
		v8 = v10
		goto _9
	_9:
		v11 = v
		v12 = v8
		v13 = (*CollisionPlane)(unsafe.Pointer(plane)).Plane.Normal
		v14 = Vec2{
			X: v11.X - float32(v12*v13.X),
			Y: v11.Y - float32(v12*v13.Y),
		}
		goto _15
	_15:
		v = v14
		goto _1
	_1:
		;
		planeIndex = planeIndex + 1
	}
	return v
}

var b2_identityBodyState10 = b2BodyState{
	DeltaRotation: Rot{
		C: float32FromFloat32(1),
	},
}

var b2_identityBodyState11 = b2BodyState{
	DeltaRotation: Rot{
		C: float32FromFloat32(1),
	},
}

var b2_identityBodyState12 = b2BodyState{
	DeltaRotation: Rot{
		C: float32FromFloat32(1),
	},
}

type b2SensorQueryContext struct {
	World       uintptr
	TaskContext uintptr
	Sensor      uintptr
	SensorShape uintptr
	Transform   Transform
}

var b2_identityBodyState13 = b2BodyState{
	DeltaRotation: Rot{
		C: float32FromFloat32(1),
	},
}

func b2CreateChain(tls *_Stack, bodyId BodyId, def uintptr) (r ChainId) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var body, chainShape, material, points, shape, shape1, shape2, shape3, world, v1, v3 uintptr
	var chainId, i, i1, i2, materialCount, materialIndex, materialIndex1, materialIndex2, materialIndex3, n, newCapacity, prevIndex, v2 int32
	var id ChainId
	var transform Transform
	var _ /* chainSegment at bp+116 */ ChainSegment
	var _ /* chainSegment at bp+80 */ ChainSegment
	var _ /* shapeDef at bp+0 */ ShapeDef
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, chainId, chainShape, i, i1, i2, id, material, materialCount, materialIndex, materialIndex1, materialIndex2, materialIndex3, n, newCapacity, points, prevIndex, shape, shape1, shape2, shape3, transform, world, v1, v2, v3
	if !((*ChainDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B2_SECRET_COOKIE)) && b2InternalAssertFcn(tls, __ccgo_ts+730, __ccgo_ts+11037, int32FromInt32(341)) != 0 {
		__builtin_trap(tls)
	}
	if !((*ChainDef)(unsafe.Pointer(def)).Count >= int32FromInt32(4)) && b2InternalAssertFcn(tls, __ccgo_ts+11577, __ccgo_ts+11037, int32FromInt32(342)) != 0 {
		__builtin_trap(tls)
	}
	if !((*ChainDef)(unsafe.Pointer(def)).MaterialCount == int32(1) || (*ChainDef)(unsafe.Pointer(def)).MaterialCount == (*ChainDef)(unsafe.Pointer(def)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+11593, __ccgo_ts+11037, int32FromInt32(343)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return ChainId{}
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	transform = b2GetBodyTransformQuick(tls, world, body)
	chainId = b2AllocId(tls, world+1224)
	if chainId == (*b2World)(unsafe.Pointer(world)).ChainShapes.Count {
		v1 = world + 1264
		if (*b2ChainShapeArray)(unsafe.Pointer(v1)).Count == (*b2ChainShapeArray)(unsafe.Pointer(v1)).Capacity {
			if (*b2ChainShapeArray)(unsafe.Pointer(v1)).Capacity < int32(2) {
				v2 = int32(2)
			} else {
				v2 = (*b2ChainShapeArray)(unsafe.Pointer(v1)).Capacity + (*b2ChainShapeArray)(unsafe.Pointer(v1)).Capacity>>int32(1)
			}
			newCapacity = v2
			b2ChainShapeArray_Reserve(tls, v1, newCapacity)
		}
		*(*b2ChainShape)(unsafe.Add(nil, (*b2ChainShapeArray)(unsafe.Pointer(v1)).Data+uintptr((*b2ChainShapeArray)(unsafe.Pointer(v1)).Count)*48)) = b2ChainShape{}
		*(*int32)(unsafe.Add(nil, v1+8)) += int32(1)
	} else {
		if !((*(*b2ChainShape)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).ChainShapes.Data+uintptr(chainId)*48))).Id == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+11653, __ccgo_ts+11037, int32FromInt32(362)) != 0 {
			__builtin_trap(tls)
		}
	}
	v1 = world + 1264
	v2 = chainId
	if !(0 <= v2 && v2 < (*b2ChainShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(137)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2ChainShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*48
	goto _6
_6:
	chainShape = v3
	(*b2ChainShape)(unsafe.Pointer(chainShape)).Id = chainId
	(*b2ChainShape)(unsafe.Pointer(chainShape)).BodyId = (*b2Body)(unsafe.Pointer(body)).Id
	(*b2ChainShape)(unsafe.Pointer(chainShape)).NextChainId = (*b2Body)(unsafe.Pointer(body)).HeadChainId
	v1 = chainShape + 40
	*(*uint16_t)(unsafe.Pointer(v1)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v1))) + int32FromInt32(1))
	materialCount = (*ChainDef)(unsafe.Pointer(def)).MaterialCount
	(*b2ChainShape)(unsafe.Pointer(chainShape)).MaterialCount = materialCount
	(*b2ChainShape)(unsafe.Pointer(chainShape)).Materials = b2Alloc(tls, int32FromUint64(uint64FromInt32(materialCount)*uint64(24)))
	i = 0
	for {
		if !(i < materialCount) {
			break
		}
		material = (*ChainDef)(unsafe.Pointer(def)).Materials + uintptr(i)*24
		if !(b2IsValidFloat(tls, (*SurfaceMaterial)(unsafe.Pointer(material)).Friction) != 0 && (*SurfaceMaterial)(unsafe.Pointer(material)).Friction >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+11706, __ccgo_ts+11037, int32FromInt32(379)) != 0 {
			__builtin_trap(tls)
		}
		if !(b2IsValidFloat(tls, (*SurfaceMaterial)(unsafe.Pointer(material)).Restitution) != 0 && (*SurfaceMaterial)(unsafe.Pointer(material)).Restitution >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+11773, __ccgo_ts+11037, int32FromInt32(380)) != 0 {
			__builtin_trap(tls)
		}
		if !(b2IsValidFloat(tls, (*SurfaceMaterial)(unsafe.Pointer(material)).RollingResistance) != 0 && (*SurfaceMaterial)(unsafe.Pointer(material)).RollingResistance >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+11846, __ccgo_ts+11037, int32FromInt32(381)) != 0 {
			__builtin_trap(tls)
		}
		if !(b2IsValidFloat(tls, (*SurfaceMaterial)(unsafe.Pointer(material)).TangentSpeed) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+11931, __ccgo_ts+11037, int32FromInt32(382)) != 0 {
			__builtin_trap(tls)
		}
		*(*SurfaceMaterial)(unsafe.Add(nil, (*b2ChainShape)(unsafe.Pointer(chainShape)).Materials+uintptr(i)*24)) = *(*SurfaceMaterial)(unsafe.Add(nil, material))
		goto _8
	_8:
		;
		i = i + 1
	}
	(*b2Body)(unsafe.Pointer(body)).HeadChainId = chainId
	*(*ShapeDef)(unsafe.Add(nil, bp)) = b2DefaultShapeDef(tls)
	(*(*ShapeDef)(unsafe.Add(nil, bp))).UserData = (*ChainDef)(unsafe.Pointer(def)).UserData
	(*(*ShapeDef)(unsafe.Add(nil, bp))).Filter = (*ChainDef)(unsafe.Pointer(def)).Filter
	(*(*ShapeDef)(unsafe.Add(nil, bp))).EnableSensorEvents = (*ChainDef)(unsafe.Pointer(def)).EnableSensorEvents
	(*(*ShapeDef)(unsafe.Add(nil, bp))).EnableContactEvents = boolUint8(false1 != 0)
	(*(*ShapeDef)(unsafe.Add(nil, bp))).EnableHitEvents = boolUint8(false1 != 0)
	points = (*ChainDef)(unsafe.Pointer(def)).Points
	n = (*ChainDef)(unsafe.Pointer(def)).Count
	if (*ChainDef)(unsafe.Pointer(def)).IsLoop != 0 {
		(*b2ChainShape)(unsafe.Pointer(chainShape)).Count = n
		(*b2ChainShape)(unsafe.Pointer(chainShape)).ShapeIndices = b2Alloc(tls, int32FromUint64(uint64FromInt32((*b2ChainShape)(unsafe.Pointer(chainShape)).Count)*uint64(4)))
		prevIndex = n - int32(1)
		i1 = 0
		for {
			if !(i1 < n-int32(2)) {
				break
			}
			(*(*ChainSegment)(unsafe.Add(nil, bp+80))).Ghost1 = *(*Vec2)(unsafe.Add(nil, points+uintptr(prevIndex)*8))
			(*(*ChainSegment)(unsafe.Add(nil, bp+80))).Segment.Point1 = *(*Vec2)(unsafe.Add(nil, points+uintptr(i1)*8))
			(*(*ChainSegment)(unsafe.Add(nil, bp+80))).Segment.Point2 = *(*Vec2)(unsafe.Add(nil, points+uintptr(i1+int32(1))*8))
			(*(*ChainSegment)(unsafe.Add(nil, bp+80))).Ghost2 = *(*Vec2)(unsafe.Add(nil, points+uintptr(i1+int32(2))*8))
			(*(*ChainSegment)(unsafe.Add(nil, bp+80))).ChainId = chainId
			prevIndex = i1
			if materialCount == int32(1) {
				v2 = 0
			} else {
				v2 = i1
			}
			materialIndex = v2
			(*(*ShapeDef)(unsafe.Add(nil, bp))).Material = *(*SurfaceMaterial)(unsafe.Add(nil, (*ChainDef)(unsafe.Pointer(def)).Materials+uintptr(materialIndex)*24))
			shape = b2CreateShapeInternal(tls, world, body, transform, bp, bp+80, int32(b2_chainSegmentShape))
			*(*int32)(unsafe.Add(nil, (*b2ChainShape)(unsafe.Pointer(chainShape)).ShapeIndices+uintptr(i1)*4)) = (*b2Shape)(unsafe.Pointer(shape)).Id
			goto _9
		_9:
			;
			i1 = i1 + 1
		}
		(*(*ChainSegment)(unsafe.Add(nil, bp+80))).Ghost1 = *(*Vec2)(unsafe.Add(nil, points+uintptr(n-int32(3))*8))
		(*(*ChainSegment)(unsafe.Add(nil, bp+80))).Segment.Point1 = *(*Vec2)(unsafe.Add(nil, points+uintptr(n-int32(2))*8))
		(*(*ChainSegment)(unsafe.Add(nil, bp+80))).Segment.Point2 = *(*Vec2)(unsafe.Add(nil, points+uintptr(n-int32(1))*8))
		(*(*ChainSegment)(unsafe.Add(nil, bp+80))).Ghost2 = *(*Vec2)(unsafe.Add(nil, points))
		(*(*ChainSegment)(unsafe.Add(nil, bp+80))).ChainId = chainId
		if materialCount == int32(1) {
			v2 = 0
		} else {
			v2 = n - int32(2)
		}
		materialIndex1 = v2
		(*(*ShapeDef)(unsafe.Add(nil, bp))).Material = *(*SurfaceMaterial)(unsafe.Add(nil, (*ChainDef)(unsafe.Pointer(def)).Materials+uintptr(materialIndex1)*24))
		shape1 = b2CreateShapeInternal(tls, world, body, transform, bp, bp+80, int32(b2_chainSegmentShape))
		*(*int32)(unsafe.Add(nil, (*b2ChainShape)(unsafe.Pointer(chainShape)).ShapeIndices+uintptr(n-int32(2))*4)) = (*b2Shape)(unsafe.Pointer(shape1)).Id
		(*(*ChainSegment)(unsafe.Add(nil, bp+80))).Ghost1 = *(*Vec2)(unsafe.Add(nil, points+uintptr(n-int32(2))*8))
		(*(*ChainSegment)(unsafe.Add(nil, bp+80))).Segment.Point1 = *(*Vec2)(unsafe.Add(nil, points+uintptr(n-int32(1))*8))
		(*(*ChainSegment)(unsafe.Add(nil, bp+80))).Segment.Point2 = *(*Vec2)(unsafe.Add(nil, points))
		(*(*ChainSegment)(unsafe.Add(nil, bp+80))).Ghost2 = *(*Vec2)(unsafe.Add(nil, points+1*8))
		(*(*ChainSegment)(unsafe.Add(nil, bp+80))).ChainId = chainId
		if materialCount == int32(1) {
			v2 = 0
		} else {
			v2 = n - int32(1)
		}
		materialIndex2 = v2
		(*(*ShapeDef)(unsafe.Add(nil, bp))).Material = *(*SurfaceMaterial)(unsafe.Add(nil, (*ChainDef)(unsafe.Pointer(def)).Materials+uintptr(materialIndex2)*24))
		shape2 = b2CreateShapeInternal(tls, world, body, transform, bp, bp+80, int32(b2_chainSegmentShape))
		*(*int32)(unsafe.Add(nil, (*b2ChainShape)(unsafe.Pointer(chainShape)).ShapeIndices+uintptr(n-int32(1))*4)) = (*b2Shape)(unsafe.Pointer(shape2)).Id
	} else {
		(*b2ChainShape)(unsafe.Pointer(chainShape)).Count = n - int32(3)
		(*b2ChainShape)(unsafe.Pointer(chainShape)).ShapeIndices = b2Alloc(tls, int32FromUint64(uint64FromInt32((*b2ChainShape)(unsafe.Pointer(chainShape)).Count)*uint64(4)))
		i2 = 0
		for {
			if !(i2 < n-int32(3)) {
				break
			}
			(*(*ChainSegment)(unsafe.Add(nil, bp+116))).Ghost1 = *(*Vec2)(unsafe.Add(nil, points+uintptr(i2)*8))
			(*(*ChainSegment)(unsafe.Add(nil, bp+116))).Segment.Point1 = *(*Vec2)(unsafe.Add(nil, points+uintptr(i2+int32(1))*8))
			(*(*ChainSegment)(unsafe.Add(nil, bp+116))).Segment.Point2 = *(*Vec2)(unsafe.Add(nil, points+uintptr(i2+int32(2))*8))
			(*(*ChainSegment)(unsafe.Add(nil, bp+116))).Ghost2 = *(*Vec2)(unsafe.Add(nil, points+uintptr(i2+int32(3))*8))
			(*(*ChainSegment)(unsafe.Add(nil, bp+116))).ChainId = chainId
			if materialCount == int32(1) {
				v2 = 0
			} else {
				v2 = i2 + int32(1)
			}
			// Material is associated with leading point of solid segment
			materialIndex3 = v2
			(*(*ShapeDef)(unsafe.Add(nil, bp))).Material = *(*SurfaceMaterial)(unsafe.Add(nil, (*ChainDef)(unsafe.Pointer(def)).Materials+uintptr(materialIndex3)*24))
			shape3 = b2CreateShapeInternal(tls, world, body, transform, bp, bp+116, int32(b2_chainSegmentShape))
			*(*int32)(unsafe.Add(nil, (*b2ChainShape)(unsafe.Pointer(chainShape)).ShapeIndices+uintptr(i2)*4)) = (*b2Shape)(unsafe.Pointer(shape3)).Id
			goto _13
		_13:
			;
			i2 = i2 + 1
		}
	}
	id = ChainId{
		Index1:     chainId + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2ChainShape)(unsafe.Pointer(chainShape)).Generation,
	}
	return id
}

func b2FreeChainData(tls *_Stack, chain uintptr) {
	b2Free(tls, (*b2ChainShape)(unsafe.Pointer(chain)).ShapeIndices, int32FromUint64(uint64FromInt32((*b2ChainShape)(unsafe.Pointer(chain)).Count)*uint64(4)))
	(*b2ChainShape)(unsafe.Pointer(chain)).ShapeIndices = uintptrFromInt32(0)
	b2Free(tls, (*b2ChainShape)(unsafe.Pointer(chain)).Materials, int32FromUint64(uint64FromInt32((*b2ChainShape)(unsafe.Pointer(chain)).MaterialCount)*uint64(24)))
	(*b2ChainShape)(unsafe.Pointer(chain)).Materials = uintptrFromInt32(0)
}

func b2DestroyChain(tls *_Stack, chainId ChainId) {
	var body, chain, chainIdPtr, shape, world, v1, v3 uintptr
	var count, i, shapeId, v2 int32
	var found, wakeBodies uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _ = body, chain, chainIdPtr, count, found, i, shape, shapeId, wakeBodies, world, v1, v2, v3
	world = b2GetWorldLocked(tls, int32FromUint16(chainId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	chain = b2GetChainShape(tls, world, chainId)
	v1 = world + 1024
	v2 = (*b2ChainShape)(unsafe.Pointer(chain)).BodyId
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	body = v3
	// Remove the chain from the body's singly linked list.
	chainIdPtr = body + 64
	found = boolUint8(false1 != 0)
	for *(*int32)(unsafe.Add(nil, chainIdPtr)) != -int32(1) {
		if *(*int32)(unsafe.Add(nil, chainIdPtr)) == (*b2ChainShape)(unsafe.Pointer(chain)).Id {
			*(*int32)(unsafe.Add(nil, chainIdPtr)) = (*b2ChainShape)(unsafe.Pointer(chain)).NextChainId
			found = boolUint8(true1 != 0)
			break
		}
		chainIdPtr = (*b2World)(unsafe.Pointer(world)).ChainShapes.Data + uintptr(*(*int32)(unsafe.Add(nil, chainIdPtr)))*48 + 8
	}
	if !(int32FromUint8(found) == int32FromInt32(true1)) && b2InternalAssertFcn(tls, __ccgo_ts+11972, __ccgo_ts+11037, int32FromInt32(515)) != 0 {
		__builtin_trap(tls)
	}
	if int32FromUint8(found) == false1 {
		return
	}
	count = (*b2ChainShape)(unsafe.Pointer(chain)).Count
	i = 0
	for {
		if !(i < count) {
			break
		}
		shapeId = *(*int32)(unsafe.Add(nil, (*b2ChainShape)(unsafe.Pointer(chain)).ShapeIndices+uintptr(i)*4))
		v1 = world + 1248
		v2 = shapeId
		if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
		goto _9
	_9:
		shape = v3
		wakeBodies = boolUint8(true1 != 0)
		b2DestroyShapeInternal(tls, world, shape, body, wakeBodies)
		goto _5
	_5:
		;
		i = i + 1
	}
	b2FreeChainData(tls, chain)
	// Return chain to free list.
	b2FreeId(tls, world+1224, (*b2ChainShape)(unsafe.Pointer(chain)).Id)
	(*b2ChainShape)(unsafe.Pointer(chain)).Id = -int32(1)
	b2ValidateSolverSets(tls, world)
}

func b2Chain_GetWorld(tls *_Stack, chainId ChainId) (r WorldId) {
	var world uintptr
	_ = world
	world = b2GetWorld(tls, int32FromUint16(chainId.World0))
	return WorldId{
		Index1:     uint16FromInt32(int32FromUint16(chainId.World0) + int32(1)),
		Generation: (*b2World)(unsafe.Pointer(world)).Generation,
	}
}

func b2Chain_GetSegmentCount(tls *_Stack, chainId ChainId) (r int32) {
	var chain, world uintptr
	_, _ = chain, world
	world = b2GetWorldLocked(tls, int32FromUint16(chainId.World0))
	if world == uintptrFromInt32(0) {
		return 0
	}
	chain = b2GetChainShape(tls, world, chainId)
	return (*b2ChainShape)(unsafe.Pointer(chain)).Count
}

func b2Chain_GetSegments(tls *_Stack, chainId ChainId, segmentArray uintptr, capacity int32) (r int32) {
	var chain, shape, world, v7, v9 uintptr
	var count, i, shapeId, v1, v2, v3, v5, v8 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _ = chain, count, i, shape, shapeId, world, v1, v2, v3, v5, v7, v8, v9
	world = b2GetWorldLocked(tls, int32FromUint16(chainId.World0))
	if world == uintptrFromInt32(0) {
		return 0
	}
	chain = b2GetChainShape(tls, world, chainId)
	v1 = (*b2ChainShape)(unsafe.Pointer(chain)).Count
	v2 = capacity
	if v1 < v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _4
_4:
	count = v3
	i = 0
	for {
		if !(i < count) {
			break
		}
		shapeId = *(*int32)(unsafe.Add(nil, (*b2ChainShape)(unsafe.Pointer(chain)).ShapeIndices+uintptr(i)*4))
		v7 = world + 1248
		v8 = shapeId
		if !(0 <= v8 && v8 < (*b2ShapeArray)(unsafe.Pointer(v7)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v9 = (*b2ShapeArray)(unsafe.Pointer(v7)).Data + uintptr(v8)*288
		goto _10
	_10:
		shape = v9
		*(*ShapeId)(unsafe.Add(nil, segmentArray+uintptr(i)*8)) = ShapeId{
			Index1:     shapeId + int32(1),
			World0:     chainId.World0,
			Generation: (*b2Shape)(unsafe.Pointer(shape)).Generation,
		}
		goto _6
	_6:
		;
		i = i + 1
	}
	return count
}

func b2ResetProxy(tls *_Stack, world uintptr, shape uintptr, wakeBodies uint8, destroyProxy uint8) {
	var body, contact, v1, v3, v5, v7 uintptr
	var contactId, contactKey, edgeIndex, shapeId, v2, v6 int32
	var forcePairCreation uint8
	var proxyType, proxyType1 b2BodyType1
	var transform Transform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, contact, contactId, contactKey, edgeIndex, forcePairCreation, proxyType, proxyType1, shapeId, transform, v1, v2, v3, v5, v6, v7
	v1 = world + 1024
	v2 = (*b2Shape)(unsafe.Pointer(shape)).BodyId
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	body = v3
	shapeId = (*b2Shape)(unsafe.Pointer(shape)).Id
	// destroy all contacts associated with this shape
	contactKey = (*b2Body)(unsafe.Pointer(body)).HeadContactKey
	for contactKey != -int32(1) {
		contactId = contactKey >> int32(1)
		edgeIndex = contactKey & int32(1)
		v5 = world + 1144
		v6 = contactId
		if !(0 <= v6 && v6 < (*b2ContactArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
			__builtin_trap(tls)
		}
		v7 = (*b2ContactArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*68
		goto _8
	_8:
		contact = v7
		contactKey = (*(*b2ContactEdge)(unsafe.Add(nil, contact+12+uintptr(edgeIndex)*12))).NextKey
		if (*b2Contact)(unsafe.Pointer(contact)).ShapeIdA == shapeId || (*b2Contact)(unsafe.Pointer(contact)).ShapeIdB == shapeId {
			b2DestroyContact(tls, world, contact, wakeBodies)
		}
	}
	transform = b2GetBodyTransformQuick(tls, world, body)
	if (*b2Shape)(unsafe.Pointer(shape)).ProxyKey != -int32(1) {
		proxyType = (*b2Shape)(unsafe.Pointer(shape)).ProxyKey & int32FromInt32(3)
		b2UpdateShapeAABBs(tls, shape, transform, proxyType)
		if destroyProxy != 0 {
			b2BroadPhase_DestroyProxy(tls, world+40, (*b2Shape)(unsafe.Pointer(shape)).ProxyKey)
			forcePairCreation = boolUint8(true1 != 0)
			(*b2Shape)(unsafe.Pointer(shape)).ProxyKey = b2BroadPhase_CreateProxy(tls, world+40, proxyType, (*b2Shape)(unsafe.Pointer(shape)).FatAABB, (*b2Shape)(unsafe.Pointer(shape)).Filter.CategoryBits, shapeId, forcePairCreation)
		} else {
			b2BroadPhase_MoveProxy(tls, world+40, (*b2Shape)(unsafe.Pointer(shape)).ProxyKey, (*b2Shape)(unsafe.Pointer(shape)).FatAABB)
		}
	} else {
		proxyType1 = (*b2Body)(unsafe.Pointer(body)).Type1
		b2UpdateShapeAABBs(tls, shape, transform, proxyType1)
	}
	b2ValidateSolverSets(tls, world)
}

func b2Chain_SetFriction(tls *_Stack, chainId ChainId, friction float32) {
	var chainShape, shape, world, v3, v5 uintptr
	var count, i, i1, materialCount, shapeId, v4 int32
	_, _, _, _, _, _, _, _, _, _, _ = chainShape, count, i, i1, materialCount, shape, shapeId, world, v3, v4, v5
	if !(b2IsValidFloat(tls, friction) != 0 && friction >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+12116, __ccgo_ts+11037, int32FromInt32(1478)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorldLocked(tls, int32FromUint16(chainId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	chainShape = b2GetChainShape(tls, world, chainId)
	materialCount = (*b2ChainShape)(unsafe.Pointer(chainShape)).MaterialCount
	i = 0
	for {
		if !(i < materialCount) {
			break
		}
		(*(*SurfaceMaterial)(unsafe.Add(nil, (*b2ChainShape)(unsafe.Pointer(chainShape)).Materials+uintptr(i)*24))).Friction = friction
		goto _1
	_1:
		;
		i = i + 1
	}
	count = (*b2ChainShape)(unsafe.Pointer(chainShape)).Count
	i1 = 0
	for {
		if !(i1 < count) {
			break
		}
		shapeId = *(*int32)(unsafe.Add(nil, (*b2ChainShape)(unsafe.Pointer(chainShape)).ShapeIndices+uintptr(i1)*4))
		v3 = world + 1248
		v4 = shapeId
		if !(0 <= v4 && v4 < (*b2ShapeArray)(unsafe.Pointer(v3)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v5 = (*b2ShapeArray)(unsafe.Pointer(v3)).Data + uintptr(v4)*288
		goto _6
	_6:
		shape = v5
		(*b2Shape)(unsafe.Pointer(shape)).Friction = friction
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
}

func b2Chain_GetFriction(tls *_Stack, chainId ChainId) (r float32) {
	var chainShape, world uintptr
	_, _ = chainShape, world
	world = b2GetWorld(tls, int32FromUint16(chainId.World0))
	chainShape = b2GetChainShape(tls, world, chainId)
	return (*(*SurfaceMaterial)(unsafe.Add(nil, (*b2ChainShape)(unsafe.Pointer(chainShape)).Materials))).Friction
}

func b2Chain_SetRestitution(tls *_Stack, chainId ChainId, restitution float32) {
	var chainShape, shape, world, v3, v5 uintptr
	var count, i, i1, materialCount, shapeId, v4 int32
	_, _, _, _, _, _, _, _, _, _, _ = chainShape, count, i, i1, materialCount, shape, shapeId, world, v3, v4, v5
	if !(b2IsValidFloat(tls, restitution) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+12375, __ccgo_ts+11037, int32FromInt32(1513)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorldLocked(tls, int32FromUint16(chainId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	chainShape = b2GetChainShape(tls, world, chainId)
	materialCount = (*b2ChainShape)(unsafe.Pointer(chainShape)).MaterialCount
	i = 0
	for {
		if !(i < materialCount) {
			break
		}
		(*(*SurfaceMaterial)(unsafe.Add(nil, (*b2ChainShape)(unsafe.Pointer(chainShape)).Materials+uintptr(i)*24))).Restitution = restitution
		goto _1
	_1:
		;
		i = i + 1
	}
	count = (*b2ChainShape)(unsafe.Pointer(chainShape)).Count
	i1 = 0
	for {
		if !(i1 < count) {
			break
		}
		shapeId = *(*int32)(unsafe.Add(nil, (*b2ChainShape)(unsafe.Pointer(chainShape)).ShapeIndices+uintptr(i1)*4))
		v3 = world + 1248
		v4 = shapeId
		if !(0 <= v4 && v4 < (*b2ShapeArray)(unsafe.Pointer(v3)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v5 = (*b2ShapeArray)(unsafe.Pointer(v3)).Data + uintptr(v4)*288
		goto _6
	_6:
		shape = v5
		(*b2Shape)(unsafe.Pointer(shape)).Restitution = restitution
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
}

func b2Chain_GetRestitution(tls *_Stack, chainId ChainId) (r float32) {
	var chainShape, world uintptr
	_, _ = chainShape, world
	world = b2GetWorld(tls, int32FromUint16(chainId.World0))
	chainShape = b2GetChainShape(tls, world, chainId)
	return (*(*SurfaceMaterial)(unsafe.Add(nil, (*b2ChainShape)(unsafe.Pointer(chainShape)).Materials))).Restitution
}

func b2Chain_SetMaterial(tls *_Stack, chainId ChainId, material int32) {
	var chainShape, shape, world, v3, v5 uintptr
	var count, i, i1, materialCount, shapeId, v4 int32
	_, _, _, _, _, _, _, _, _, _, _ = chainShape, count, i, i1, materialCount, shape, shapeId, world, v3, v4, v5
	world = b2GetWorldLocked(tls, int32FromUint16(chainId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	chainShape = b2GetChainShape(tls, world, chainId)
	materialCount = (*b2ChainShape)(unsafe.Pointer(chainShape)).MaterialCount
	i = 0
	for {
		if !(i < materialCount) {
			break
		}
		(*(*SurfaceMaterial)(unsafe.Add(nil, (*b2ChainShape)(unsafe.Pointer(chainShape)).Materials+uintptr(i)*24))).UserMaterialId = material
		goto _1
	_1:
		;
		i = i + 1
	}
	count = (*b2ChainShape)(unsafe.Pointer(chainShape)).Count
	i1 = 0
	for {
		if !(i1 < count) {
			break
		}
		shapeId = *(*int32)(unsafe.Add(nil, (*b2ChainShape)(unsafe.Pointer(chainShape)).ShapeIndices+uintptr(i1)*4))
		v3 = world + 1248
		v4 = shapeId
		if !(0 <= v4 && v4 < (*b2ShapeArray)(unsafe.Pointer(v3)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v5 = (*b2ShapeArray)(unsafe.Pointer(v3)).Data + uintptr(v4)*288
		goto _6
	_6:
		shape = v5
		(*b2Shape)(unsafe.Pointer(shape)).UserMaterialId = material
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
}

func b2Chain_GetMaterial(tls *_Stack, chainId ChainId) (r int32) {
	var chainShape, world uintptr
	_, _ = chainShape, world
	world = b2GetWorld(tls, int32FromUint16(chainId.World0))
	chainShape = b2GetChainShape(tls, world, chainId)
	return (*(*SurfaceMaterial)(unsafe.Add(nil, (*b2ChainShape)(unsafe.Pointer(chainShape)).Materials))).UserMaterialId
}

const _ARG_MAX = 131072
const _B2_SIMD_SHIFT = 2
const _B2_TIME_TO_SLEEP1 = 0.5
const _BC_BASE_MAX = 99
const _BC_DIM_MAX = 2048
const _BC_SCALE_MAX = 99
const _BC_STRING_MAX = 1000
const _CHARCLASS_NAME_MAX = 14
const _CHAR_BIT = 8
const _CHAR_MAX = 255
const _CHAR_MIN = 0
const _COLL_WEIGHTS_MAX = 2
const _DELAYTIMER_MAX = 0x7fffffff
const _EXPR_NEST_MAX = 32
const _FILESIZEBITS = 64
const _HOST_NAME_MAX = 255
const _INT_MAX = 0x7fffffff
const _IOV_MAX = 1024
const _ITERATIONS = 1
const _LINE_MAX = 4096
const _LLONG_MAX = 0x7fffffffffffffff
const _LOGIN_NAME_MAX = 256
const _LONG_BIT = 64
const _LONG_MAX = "__LONG_MAX"
const _MB_LEN_MAX = 4
const _MQ_PRIO_MAX = 32768
const _NAME_MAX = 255
const _NGROUPS_MAX = 32
const _NL_ARGMAX = 9
const _NL_LANGMAX = 32
const _NL_MSGMAX = 32767
const _NL_NMAX = 16
const _NL_SETMAX = 255
const _NL_TEXTMAX = 2048
const _NZERO = 20
const _PATH_MAX = 4096
const _PIPE_BUF = 4096
const _PTHREAD_DESTRUCTOR_ITERATIONS = 4
const _PTHREAD_KEYS_MAX = 128
const _PTHREAD_STACK_MIN = 2048
const _RELAX_ITERATIONS = 1
const _RE_DUP_MAX = 255
const _SCHAR_MAX = 127
const _SEM_NSEMS_MAX = 256
const _SEM_VALUE_MAX = 0x7fffffff
const _SHRT_MAX = 0x7fff
const _SSIZE_MAX = "LONG_MAX"
const _SYMLOOP_MAX = 40
const _TTY_NAME_MAX = 32
const _TZNAME_MAX = 6
const _UCHAR_MAX = 255
const _UINT64_MAX7 = 18446744073709551615
const _UINT_MAX = 4294967295
const _USHRT_MAX = 0xffff
const _WORD_BIT = 32
const _POSIX2_BC_BASE_MAX = 99
const _POSIX2_BC_DIM_MAX = 2048
const _POSIX2_BC_SCALE_MAX = 99
const _POSIX2_BC_STRING_MAX = 1000
const _POSIX2_CHARCLASS_NAME_MAX = 14
const _POSIX2_COLL_WEIGHTS_MAX = 2
const _POSIX2_EXPR_NEST_MAX = 32
const _POSIX2_LINE_MAX = 2048
const _POSIX2_RE_DUP_MAX = 255
const _POSIX_AIO_LISTIO_MAX = 2
const _POSIX_AIO_MAX = 1
const _POSIX_ARG_MAX = 4096
const _POSIX_CHILD_MAX = 25
const _POSIX_CLOCKRES_MIN = 20000000
const _POSIX_DELAYTIMER_MAX = 32
const _POSIX_HOST_NAME_MAX = 255
const _POSIX_LINK_MAX = 8
const _POSIX_LOGIN_NAME_MAX = 9
const _POSIX_MAX_CANON = 255
const _POSIX_MAX_INPUT = 255
const _POSIX_MQ_OPEN_MAX = 8
const _POSIX_MQ_PRIO_MAX = 32
const _POSIX_NAME_MAX = 14
const _POSIX_NGROUPS_MAX = 8
const _POSIX_OPEN_MAX = 20
const _POSIX_PATH_MAX = 256
const _POSIX_PIPE_BUF = 512
const _POSIX_RE_DUP_MAX = 255
const _POSIX_RTSIG_MAX = 8
const _POSIX_SEM_NSEMS_MAX = 256
const _POSIX_SEM_VALUE_MAX = 32767
const _POSIX_SIGQUEUE_MAX = 32
const _POSIX_SSIZE_MAX = 32767
const _POSIX_SS_REPL_MAX = 4
const _POSIX_STREAM_MAX = 8
const _POSIX_SYMLINK_MAX = 255
const _POSIX_SYMLOOP_MAX = 8
const _POSIX_THREAD_DESTRUCTOR_ITERATIONS = 4
const _POSIX_THREAD_KEYS_MAX = 128
const _POSIX_THREAD_THREADS_MAX = 64
const _POSIX_TIMER_MAX = 32
const _POSIX_TRACE_EVENT_NAME_MAX = 30
const _POSIX_TRACE_NAME_MAX = 8
const _POSIX_TRACE_SYS_MAX = 8
const _POSIX_TRACE_USER_EVENT_MAX = 32
const _POSIX_TTY_NAME_MAX = 9
const _POSIX_TZNAME_MAX = 6
const _XOPEN_IOV_MAX = 16
const _XOPEN_NAME_MAX = 255
const _XOPEN_PATH_MAX = 1024

var b2_identityBodyState14 = b2BodyState{
	DeltaRotation: Rot{
		C: float32FromFloat32(1),
	},
}

func b2Pause(tls *_Stack) {
	__builtin_Gosched(tls)
}

type b2WorkerContext struct {
	Context     uintptr
	WorkerIndex int32
	UserTask    uintptr
}

// C documentation
//
//	// Integrate velocities and apply damping
func b2IntegrateVelocitiesTask(tls *_Stack, startIndex int32, endIndex int32, context uintptr) {
	var angularDamping, angularVelocityDelta, gravityScale, h, linearDamping, maxAngularSpeed, maxAngularSpeedSquared, maxLinearSpeed, maxLinearSpeedSquared, ratio, ratio1, w, v16, v21, v3, v7 float32
	var gravity, linearVelocityDelta, v2, v11, v12, v13, v15, v17, v18, v4, v5, v8, v9 Vec2
	var i int32
	var sim, sims, state, states uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = angularDamping, angularVelocityDelta, gravity, gravityScale, h, i, linearDamping, linearVelocityDelta, maxAngularSpeed, maxAngularSpeedSquared, maxLinearSpeed, maxLinearSpeedSquared, ratio, ratio1, sim, sims, state, states, v2, w, v11, v12, v13, v15, v16, v17, v18, v21, v3, v4, v5, v7, v8, v9
	states = (*b2StepContext)(unsafe.Pointer(context)).States
	sims = (*b2StepContext)(unsafe.Pointer(context)).Sims
	gravity = (*b2World)(unsafe.Pointer((*b2StepContext)(unsafe.Pointer(context)).World)).Gravity
	h = (*b2StepContext)(unsafe.Pointer(context)).H
	maxLinearSpeed = (*b2StepContext)(unsafe.Pointer(context)).MaxLinearVelocity
	maxAngularSpeed = float32(float32(float32FromFloat32(0.25)*float32FromFloat32(3.14159265359)) * (*b2StepContext)(unsafe.Pointer(context)).Inv_dt)
	maxLinearSpeedSquared = float32(maxLinearSpeed * maxLinearSpeed)
	maxAngularSpeedSquared = float32(maxAngularSpeed * maxAngularSpeed)
	i = startIndex
	for {
		if !(i < endIndex) {
			break
		}
		sim = sims + uintptr(i)*100
		state = states + uintptr(i)*32
		v2 = (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
		w = (*b2BodyState)(unsafe.Pointer(state)).AngularVelocity
		// Apply forces, torque, gravity, and damping
		// Apply damping.
		// Differential equation: dv/dt + c * v = 0
		// Solution: v(t) = v0 * exp(-c * t)
		// Time step: v(t + dt) = v0 * exp(-c * (t + dt)) = v0 * exp(-c * t) * exp(-c * dt) = v(t) * exp(-c * dt)
		// v2 = exp(-c * dt) * v1
		// Pade approximation:
		// v2 = v1 * 1 / (1 + c * dt)
		linearDamping = float32FromFloat32(1) / (float32FromFloat32(1) + float32(h*(*b2BodySim)(unsafe.Pointer(sim)).LinearDamping))
		angularDamping = float32FromFloat32(1) / (float32FromFloat32(1) + float32(h*(*b2BodySim)(unsafe.Pointer(sim)).AngularDamping))
		if (*b2BodySim)(unsafe.Pointer(sim)).InvMass > float32FromFloat32(0) {
			v21 = (*b2BodySim)(unsafe.Pointer(sim)).GravityScale
		} else {
			v21 = float32FromFloat32(0)
		}
		// Gravity scale will be zero for kinematic bodies
		gravityScale = v21
		v3 = float32(h * (*b2BodySim)(unsafe.Pointer(sim)).InvMass)
		v4 = (*b2BodySim)(unsafe.Pointer(sim)).Force
		v5 = Vec2{
			X: float32(v3 * v4.X),
			Y: float32(v3 * v4.Y),
		}
		goto _6
	_6:
		v7 = float32(h * gravityScale)
		v8 = gravity
		v9 = Vec2{
			X: float32(v7 * v8.X),
			Y: float32(v7 * v8.Y),
		}
		goto _10
	_10:
		v11 = v5
		v12 = v9
		v13 = Vec2{
			X: v11.X + v12.X,
			Y: v11.Y + v12.Y,
		}
		goto _14
	_14:
		// lvd = h * im * f + h * g
		linearVelocityDelta = v13
		angularVelocityDelta = float32(float32(h*(*b2BodySim)(unsafe.Pointer(sim)).InvInertia) * (*b2BodySim)(unsafe.Pointer(sim)).Torque)
		v15 = linearVelocityDelta
		v16 = linearDamping
		v17 = v2
		v18 = Vec2{
			X: v15.X + float32(v16*v17.X),
			Y: v15.Y + float32(v16*v17.Y),
		}
		goto _19
	_19:
		v2 = v18
		w = angularVelocityDelta + float32(angularDamping*w)
		// Clamp to max linear speed
		v4 = v2
		v5 = v2
		v21 = float32(v4.X*v5.X) + float32(v4.Y*v5.Y)
		goto _23
	_23:
		if v21 > maxLinearSpeedSquared {
			v8 = v2
			v3 = sqrtf(tls, float32(v8.X*v8.X)+float32(v8.Y*v8.Y))
			goto _26
		_26:
			ratio = maxLinearSpeed / v3
			v7 = ratio
			v9 = v2
			v11 = Vec2{
				X: float32(v7 * v9.X),
				Y: float32(v7 * v9.Y),
			}
			goto _30
		_30:
			v2 = v11
			(*b2BodySim)(unsafe.Pointer(sim)).IsSpeedCapped = boolUint8(true1 != 0)
		}
		// Clamp to max angular speed
		if float32(w*w) > maxAngularSpeedSquared && int32FromUint8((*b2BodySim)(unsafe.Pointer(sim)).AllowFastRotation) == false1 {
			v21 = w
			if v21 < float32FromInt32(0) {
				v7 = -v21
			} else {
				v7 = v21
			}
			v3 = v7
			goto _33
		_33:
			ratio1 = maxAngularSpeed / v3
			w = float32(w * ratio1)
			(*b2BodySim)(unsafe.Pointer(sim)).IsSpeedCapped = boolUint8(true1 != 0)
		}
		(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity = v2
		(*b2BodyState)(unsafe.Pointer(state)).AngularVelocity = w
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b2IntegratePositionsTask(tls *_Stack, startIndex int32, endIndex int32, context uintptr) {
	var h, invMag, mag, v3, v4 float32
	var i int32
	var q2, qn, v2, v5 Rot
	var state, states uintptr
	var v10, v7, v9 Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = h, i, invMag, mag, q2, qn, state, states, v10, v2, v3, v4, v5, v7, v9
	states = (*b2StepContext)(unsafe.Pointer(context)).States
	h = (*b2StepContext)(unsafe.Pointer(context)).H
	if !(startIndex <= endIndex) && b2InternalAssertFcn(tls, __ccgo_ts+12545, __ccgo_ts+12460, int32FromInt32(196)) != 0 {
		__builtin_trap(tls)
	}
	i = startIndex
	for {
		if !(i < endIndex) {
			break
		}
		state = states + uintptr(i)*32
		v2 = (*b2BodyState)(unsafe.Pointer(state)).DeltaRotation
		v3 = float32(h * (*b2BodyState)(unsafe.Pointer(state)).AngularVelocity)
		q2 = Rot{
			C: v2.C - float32(v3*v2.S),
			S: v2.S + float32(v3*v2.C),
		}
		mag = sqrtf(tls, float32(q2.S*q2.S)+float32(q2.C*q2.C))
		if float64(mag) > float64(0) {
			v4 = float32FromFloat32(1) / mag
		} else {
			v4 = float32FromFloat32(0)
		}
		invMag = v4
		qn = Rot{
			C: float32(q2.C * invMag),
			S: float32(q2.S * invMag),
		}
		v5 = qn
		goto _6
	_6:
		(*b2BodyState)(unsafe.Pointer(state)).DeltaRotation = v5
		v7 = (*b2BodyState)(unsafe.Pointer(state)).DeltaPosition
		v3 = h
		v9 = (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
		v10 = Vec2{
			X: v7.X + float32(v3*v9.X),
			Y: v7.Y + float32(v3*v9.Y),
		}
		goto _11
	_11:
		(*b2BodyState)(unsafe.Pointer(state)).DeltaPosition = v10
		goto _1
	_1:
		;
		i = i + 1
	}
}

type b2ContinuousContext struct {
	World       uintptr
	FastBodySim uintptr
	FastShape   uintptr
	Centroid1   Vec2
	Centroid2   Vec2
	Sweep       Sweep
	Fraction    float32
}

// C documentation
//
//	// This is called from b2DynamicTree_Query for continuous collision
func b2ContinuousQueryCallback(tls *_Stack, proxyId int32, userData uint64, context uintptr) (r uint8) {
	bp := tls.Alloc(352)
	defer tls.Free(352)
	var allowedFraction, hitFraction, invLength, offset1, offset2, radius, x, y, v39, v47 float32
	var body, bodySim1, continuousContext, customFilterFcn, fastBody, fastBodySim, fastShape, shape, world, v1, v3 uintptr
	var c1, c2, e, n, p1, p2, v18, v19, v22, v23, v25, v26, v27, v30, v31, v33 Vec2
	var canCollide, didHit, v7 uint8
	var extent b2ShapeExtent
	var idA, idB, shapeIdA, shapeIdB ShapeId
	var output TOIOutput
	var s, v50 Sweep
	var shapeId, v2 int32
	var transform, transformA, transformB, v17, v21 Transform
	var v5, v6 Filter
	var _ /* centroid at bp+232 */ Vec2
	var _ /* input at bp+4 */ TOIInput
	var _ /* length at bp+0 */ float32
	var _ /* manifold at bp+240 */ Manifold
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = allowedFraction, body, bodySim1, c1, c2, canCollide, continuousContext, customFilterFcn, didHit, e, extent, fastBody, fastBodySim, fastShape, hitFraction, idA, idB, invLength, n, offset1, offset2, output, p1, p2, radius, s, shape, shapeId, shapeIdA, shapeIdB, transform, transformA, transformB, world, x, y, v1, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v3, v30, v31, v33, v39, v47, v5, v50, v6, v7
	_ = uint64FromInt64(4)
	shapeId = int32FromUint64(userData)
	continuousContext = context
	fastShape = (*b2ContinuousContext)(unsafe.Pointer(continuousContext)).FastShape
	fastBodySim = (*b2ContinuousContext)(unsafe.Pointer(continuousContext)).FastBodySim
	// Skip same shape
	if shapeId == (*b2Shape)(unsafe.Pointer(fastShape)).Id {
		return boolUint8(true1 != 0)
	}
	world = (*b2ContinuousContext)(unsafe.Pointer(continuousContext)).World
	v1 = world + 1248
	v2 = shapeId
	if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
	goto _4
_4:
	shape = v3
	// Skip same body
	if (*b2Shape)(unsafe.Pointer(shape)).BodyId == (*b2Shape)(unsafe.Pointer(fastShape)).BodyId {
		return boolUint8(true1 != 0)
	}
	// Skip sensors
	if (*b2Shape)(unsafe.Pointer(shape)).SensorIndex != -int32(1) {
		return boolUint8(true1 != 0)
	}
	v5 = (*b2Shape)(unsafe.Pointer(fastShape)).Filter
	v6 = (*b2Shape)(unsafe.Pointer(shape)).Filter
	if v5.GroupIndex == v6.GroupIndex && v5.GroupIndex != 0 {
		v7 = boolUint8(v5.GroupIndex > 0)
		goto _8
	}
	v7 = boolUint8(v5.MaskBits&v6.CategoryBits != uint64(0) && v5.CategoryBits&v6.MaskBits != uint64(0))
	goto _8
_8:
	// Skip filtered shapes
	canCollide = v7
	if int32FromUint8(canCollide) == false1 {
		return boolUint8(true1 != 0)
	}
	v1 = world + 1024
	v2 = (*b2Shape)(unsafe.Pointer(shape)).BodyId
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _12
_12:
	body = v3
	bodySim1 = b2GetBodySim(tls, world, body)
	if !((*b2Body)(unsafe.Pointer(body)).Type1 == int32(b2_staticBody) || (*b2BodySim)(unsafe.Pointer(fastBodySim)).IsBullet != 0) && b2InternalAssertFcn(tls, __ccgo_ts+12568, __ccgo_ts+12460, int32FromInt32(261)) != 0 {
		__builtin_trap(tls)
	}
	// Skip bullets
	if (*b2BodySim)(unsafe.Pointer(bodySim1)).IsBullet != 0 {
		return boolUint8(true1 != 0)
	}
	v1 = world + 1024
	v2 = (*b2BodySim)(unsafe.Pointer(fastBodySim)).BodyId
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _16
_16:
	// Skip filtered bodies
	fastBody = v3
	canCollide = b2ShouldBodiesCollide(tls, world, fastBody, body)
	if int32FromUint8(canCollide) == false1 {
		return boolUint8(true1 != 0)
	}
	// Custom user filtering
	customFilterFcn = (*b2World)(unsafe.Pointer(world)).CustomFilterFcn
	if customFilterFcn != uintptrFromInt32(0) {
		idA = ShapeId{
			Index1:     (*b2Shape)(unsafe.Pointer(shape)).Id + int32(1),
			World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
			Generation: (*b2Shape)(unsafe.Pointer(shape)).Generation,
		}
		idB = ShapeId{
			Index1:     (*b2Shape)(unsafe.Pointer(fastShape)).Id + int32(1),
			World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
			Generation: (*b2Shape)(unsafe.Pointer(fastShape)).Generation,
		}
		canCollide = (*(*func(*_Stack, ShapeId, ShapeId, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{customFilterFcn})))(tls, idA, idB, (*b2World)(unsafe.Pointer(world)).CustomFilterContext)
		if int32FromUint8(canCollide) == false1 {
			return boolUint8(true1 != 0)
		}
	}
	// Prevent pausing on chain segment junctions
	if (*b2Shape)(unsafe.Pointer(shape)).Type1 == int32(b2_chainSegmentShape) {
		transform = (*b2BodySim)(unsafe.Pointer(bodySim1)).Transform
		v17 = transform
		v18 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 132))).Segment.Point1
		x = float32(v17.Q.C*v18.X) - float32(v17.Q.S*v18.Y) + v17.P.X
		y = float32(v17.Q.S*v18.X) + float32(v17.Q.C*v18.Y) + v17.P.Y
		v19 = Vec2{
			X: x,
			Y: y,
		}
		goto _20
	_20:
		p1 = v19
		v21 = transform
		v22 = (*(*ChainSegment)(unsafe.Add(unsafe.Pointer(shape), 132))).Segment.Point2
		x = float32(v21.Q.C*v22.X) - float32(v21.Q.S*v22.Y) + v21.P.X
		y = float32(v21.Q.S*v22.X) + float32(v21.Q.C*v22.Y) + v21.P.Y
		v23 = Vec2{
			X: x,
			Y: y,
		}
		goto _24
	_24:
		p2 = v23
		v25 = p2
		v26 = p1
		v27 = Vec2{
			X: v25.X - v26.X,
			Y: v25.Y - v26.Y,
		}
		goto _28
	_28:
		e = v27
		v1 = bp
		v30 = e
		*(*float32)(unsafe.Add(nil, v1)) = sqrtf(tls, float32(v30.X*v30.X)+float32(v30.Y*v30.Y))
		if *(*float32)(unsafe.Add(nil, v1)) < float32FromFloat32(1.1920928955078125e-07) {
			v31 = Vec2{}
			goto _32
		}
		invLength = float32FromFloat32(1) / *(*float32)(unsafe.Add(nil, v1))
		n = Vec2{
			X: float32(invLength * v30.X),
			Y: float32(invLength * v30.Y),
		}
		v31 = n
		goto _32
	_32:
		e = v31
		if *(*float32)(unsafe.Add(nil, bp)) > float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter) {
			c1 = (*b2ContinuousContext)(unsafe.Pointer(continuousContext)).Centroid1
			v18 = c1
			v19 = p1
			v22 = Vec2{
				X: v18.X - v19.X,
				Y: v18.Y - v19.Y,
			}
			goto _36
		_36:
			v23 = v22
			v25 = e
			v39 = float32(v23.X*v25.Y) - float32(v23.Y*v25.X)
			goto _40
		_40:
			offset1 = v39
			c2 = (*b2ContinuousContext)(unsafe.Pointer(continuousContext)).Centroid2
			v26 = c2
			v27 = p1
			v30 = Vec2{
				X: v26.X - v27.X,
				Y: v26.Y - v27.Y,
			}
			goto _44
		_44:
			v31 = v30
			v33 = e
			v47 = float32(v31.X*v33.Y) - float32(v31.Y*v33.X)
			goto _48
		_48:
			offset2 = v47
			// todo this should use the min extent of the fast shape, not the body
			allowedFraction = float32FromFloat32(0.25)
			if offset1 < float32FromFloat32(0) || offset1-offset2 < float32(allowedFraction*(*b2BodySim)(unsafe.Pointer(fastBodySim)).MinExtent) {
				// Minimal clipping
				return boolUint8(true1 != 0)
			}
		}
	}
	(*(*TOIInput)(unsafe.Add(nil, bp+4))).ProxyA = b2MakeShapeDistanceProxy(tls, shape)
	(*(*TOIInput)(unsafe.Add(nil, bp+4))).ProxyB = b2MakeShapeDistanceProxy(tls, fastShape)
	v1 = bodySim1
	s.C1 = (*b2BodySim)(unsafe.Pointer(v1)).Center0
	s.C2 = (*b2BodySim)(unsafe.Pointer(v1)).Center
	s.Q1 = (*b2BodySim)(unsafe.Pointer(v1)).Rotation0
	s.Q2 = (*b2BodySim)(unsafe.Pointer(v1)).Transform.Q
	s.LocalCenter = (*b2BodySim)(unsafe.Pointer(v1)).LocalCenter
	v50 = s
	goto _51
_51:
	(*(*TOIInput)(unsafe.Add(nil, bp+4))).SweepA = v50
	(*(*TOIInput)(unsafe.Add(nil, bp+4))).SweepB = (*b2ContinuousContext)(unsafe.Pointer(continuousContext)).Sweep
	(*(*TOIInput)(unsafe.Add(nil, bp+4))).MaxFraction = (*b2ContinuousContext)(unsafe.Pointer(continuousContext)).Fraction
	hitFraction = (*b2ContinuousContext)(unsafe.Pointer(continuousContext)).Fraction
	didHit = boolUint8(false1 != 0)
	output = b2TimeOfImpact(tls, bp+4)
	if float32FromFloat32(0) < output.Fraction && output.Fraction < (*b2ContinuousContext)(unsafe.Pointer(continuousContext)).Fraction {
		hitFraction = output.Fraction
		didHit = boolUint8(true1 != 0)
	} else {
		if float32FromFloat32(0) == output.Fraction {
			// fallback to TOI of a small circle around the fast shape centroid
			*(*Vec2)(unsafe.Add(nil, bp+232)) = b2GetShapeCentroid(tls, fastShape)
			extent = b2ComputeShapeExtent(tls, fastShape, *(*Vec2)(unsafe.Add(nil, bp+232)))
			radius = float32(float32FromFloat32(0.25) * extent.MinExtent)
			(*(*TOIInput)(unsafe.Add(nil, bp+4))).ProxyB = b2MakeProxy(tls, bp+232, int32(1), radius)
			output = b2TimeOfImpact(tls, bp+4)
			if float32FromFloat32(0) < output.Fraction && output.Fraction < (*b2ContinuousContext)(unsafe.Pointer(continuousContext)).Fraction {
				hitFraction = output.Fraction
				didHit = boolUint8(true1 != 0)
			}
		}
	}
	if didHit != 0 && ((*b2Shape)(unsafe.Pointer(shape)).EnablePreSolveEvents != 0 || (*b2Shape)(unsafe.Pointer(fastShape)).EnablePreSolveEvents != 0) && (*b2World)(unsafe.Pointer(world)).PreSolveFcn != uintptrFromInt32(0) {
		// Pre-solve is expensive because I need to compute a temporary manifold
		transformA = b2GetSweepTransform(tls, bp+4+144, hitFraction)
		transformB = b2GetSweepTransform(tls, bp+4+184, hitFraction)
		*(*Manifold)(unsafe.Add(nil, bp+240)) = b2ComputeManifold(tls, shape, transformA, fastShape, transformB)
		shapeIdA = ShapeId{
			Index1:     (*b2Shape)(unsafe.Pointer(shape)).Id + int32(1),
			World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
			Generation: (*b2Shape)(unsafe.Pointer(shape)).Generation,
		}
		shapeIdB = ShapeId{
			Index1:     (*b2Shape)(unsafe.Pointer(fastShape)).Id + int32(1),
			World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
			Generation: (*b2Shape)(unsafe.Pointer(fastShape)).Generation,
		}
		// The user may modify the temporary manifold here but it doesn't matter. They will be able to
		// modify the real manifold in the discrete solver.
		didHit = (*(*func(*_Stack, ShapeId, ShapeId, uintptr, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).PreSolveFcn})))(tls, shapeIdA, shapeIdB, bp+240, (*b2World)(unsafe.Pointer(world)).PreSolveContext)
	}
	if didHit != 0 {
		(*b2ContinuousContext)(unsafe.Pointer(continuousContext)).Fraction = hitFraction
	}
	return boolUint8(true1 != 0)
}

func b2FinalizeBodiesTask(tls *_Stack, startIndex int32, endIndex int32, threadIndex uint32, context uintptr) {
	var aabb, fatAABB, v55, v56 AABB
	var aabbMargin, invMag, invTimeStep, mag, maxDeltaPosition, maxVelocity, positionSleepFactor, sleepVelocity, speculativeDistance, timeStep, w, v11, v15, v17, v18, v20, v22, v24, v25, v27, v28, v29, v30 float32
	var awakeIslandBitSet, bodies, body, enlargedSimBitSet, island, moveEvents, shape, sim, sims, state, states, stepContext, taskContext, world, v43, v45, v47 uintptr
	var blockIndex, v48 uint32_t
	var bulletIndex, islandIndex, shapeId, simIndex, v41 int32
	var enableContinuous, enableSleep, isFast, s, v57 uint8
	var qn, qr, v10, v12, v6, v7, v8 Rot
	var transform Transform
	var v2, v14, v21, v211, v3, v4 Vec2
	var worldId uint16_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, aabbMargin, awakeIslandBitSet, blockIndex, bodies, body, bulletIndex, enableContinuous, enableSleep, enlargedSimBitSet, fatAABB, invMag, invTimeStep, isFast, island, islandIndex, mag, maxDeltaPosition, maxVelocity, moveEvents, positionSleepFactor, qn, qr, s, shape, shapeId, sim, simIndex, sims, sleepVelocity, speculativeDistance, state, states, stepContext, taskContext, timeStep, transform, v2, w, world, worldId, v10, v11, v12, v14, v15, v17, v18, v21, v20, v211, v22, v24, v25, v27, v28, v29, v3, v30, v4, v41, v43, v45, v47, v48, v55, v56, v57, v6, v7, v8
	stepContext = context
	world = (*b2StepContext)(unsafe.Pointer(stepContext)).World
	enableSleep = (*b2World)(unsafe.Pointer(world)).EnableSleep
	states = (*b2StepContext)(unsafe.Pointer(stepContext)).States
	sims = (*b2StepContext)(unsafe.Pointer(stepContext)).Sims
	bodies = (*b2World)(unsafe.Pointer(world)).Bodies.Data
	timeStep = (*b2StepContext)(unsafe.Pointer(stepContext)).Dt
	invTimeStep = (*b2StepContext)(unsafe.Pointer(stepContext)).Inv_dt
	worldId = (*b2World)(unsafe.Pointer(world)).WorldId
	// The body move event array should already have the correct size
	if !(endIndex <= (*b2World)(unsafe.Pointer(world)).BodyMoveEvents.Count) && b2InternalAssertFcn(tls, __ccgo_ts+12664, __ccgo_ts+12460, int32FromInt32(566)) != 0 {
		__builtin_trap(tls)
	}
	moveEvents = (*b2World)(unsafe.Pointer(world)).BodyMoveEvents.Data
	enlargedSimBitSet = (*b2World)(unsafe.Pointer(world)).TaskContexts.Data + uintptr(threadIndex)*56 + 16
	awakeIslandBitSet = (*b2World)(unsafe.Pointer(world)).TaskContexts.Data + uintptr(threadIndex)*56 + 32
	taskContext = (*b2World)(unsafe.Pointer(world)).TaskContexts.Data + uintptr(threadIndex)*56
	enableContinuous = (*b2World)(unsafe.Pointer(world)).EnableContinuous
	speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter))
	aabbMargin = float32(float32FromFloat32(0.05) * b2_lengthUnitsPerMeter)
	if !(startIndex <= endIndex) && b2InternalAssertFcn(tls, __ccgo_ts+12545, __ccgo_ts+12460, int32FromInt32(578)) != 0 {
		__builtin_trap(tls)
	}
	simIndex = startIndex
	for {
		if !(simIndex < endIndex) {
			break
		}
		state = states + uintptr(simIndex)*32
		sim = sims + uintptr(simIndex)*100
		v2 = (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
		w = (*b2BodyState)(unsafe.Pointer(state)).AngularVelocity
		if !(b2IsValidVec2(tls, v2) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+12704, __ccgo_ts+12460, int32FromInt32(588)) != 0 {
			__builtin_trap(tls)
		}
		if !(b2IsValidFloat(tls, w) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+12723, __ccgo_ts+12460, int32FromInt32(589)) != 0 {
			__builtin_trap(tls)
		}
		v21 = (*b2BodySim)(unsafe.Pointer(sim)).Center
		v3 = (*b2BodyState)(unsafe.Pointer(state)).DeltaPosition
		v4 = Vec2{
			X: v21.X + v3.X,
			Y: v21.Y + v3.Y,
		}
		goto _5
	_5:
		(*b2BodySim)(unsafe.Pointer(sim)).Center = v4
		v6 = (*b2BodyState)(unsafe.Pointer(state)).DeltaRotation
		v7 = (*b2BodySim)(unsafe.Pointer(sim)).Transform.Q
		qr.S = float32(v6.S*v7.C) + float32(v6.C*v7.S)
		qr.C = float32(v6.C*v7.C) - float32(v6.S*v7.S)
		v8 = qr
		goto _9
	_9:
		v10 = v8
		mag = sqrtf(tls, float32(v10.S*v10.S)+float32(v10.C*v10.C))
		if float64(mag) > float64(0) {
			v11 = float32FromFloat32(1) / mag
		} else {
			v11 = float32FromFloat32(0)
		}
		invMag = v11
		qn = Rot{
			C: float32(v10.C * invMag),
			S: float32(v10.S * invMag),
		}
		v12 = qn
		goto _13
	_13:
		(*b2BodySim)(unsafe.Pointer(sim)).Transform.Q = v12
		v21 = v2
		v11 = sqrtf(tls, float32(v21.X*v21.X)+float32(v21.Y*v21.Y))
		goto _16
	_16:
		v15 = w
		if v15 < float32FromInt32(0) {
			v18 = -v15
		} else {
			v18 = v15
		}
		v17 = v18
		goto _19
	_19:
		// Use the velocity of the farthest point on the body to account for rotation.
		maxVelocity = v11 + float32(v17*(*b2BodySim)(unsafe.Pointer(sim)).MaxExtent)
		v3 = (*b2BodyState)(unsafe.Pointer(state)).DeltaPosition
		v20 = sqrtf(tls, float32(v3.X*v3.X)+float32(v3.Y*v3.Y))
		goto _23
	_23:
		v22 = (*b2BodyState)(unsafe.Pointer(state)).DeltaRotation.S
		if v22 < float32FromInt32(0) {
			v25 = -v22
		} else {
			v25 = v22
		}
		v24 = v25
		goto _26
	_26:
		// Sleep needs to observe position correction as well as true velocity.
		maxDeltaPosition = v20 + float32(v24*(*b2BodySim)(unsafe.Pointer(sim)).MaxExtent)
		// Position correction is not as important for sleep as true velocity.
		positionSleepFactor = float32FromFloat32(0.5)
		v27 = maxVelocity
		v28 = float32(float32(positionSleepFactor*invTimeStep) * maxDeltaPosition)
		if v27 > v28 {
			v30 = v27
		} else {
			v30 = v28
		}
		v29 = v30
		goto _31
	_31:
		sleepVelocity = v29
		// reset state deltas
		(*b2BodyState)(unsafe.Pointer(state)).DeltaPosition = b2Vec2_zero
		(*b2BodyState)(unsafe.Pointer(state)).DeltaRotation = b2Rot_identity
		v6 = (*b2BodySim)(unsafe.Pointer(sim)).Transform.Q
		v21 = (*b2BodySim)(unsafe.Pointer(sim)).LocalCenter
		v3 = Vec2{
			X: float32(v6.C*v21.X) - float32(v6.S*v21.Y),
			Y: float32(v6.S*v21.X) + float32(v6.C*v21.Y),
		}
		goto _36
	_36:
		v4 = (*b2BodySim)(unsafe.Pointer(sim)).Center
		v14 = v3
		v211 = Vec2{
			X: v4.X - v14.X,
			Y: v4.Y - v14.Y,
		}
		goto _40
	_40:
		(*b2BodySim)(unsafe.Pointer(sim)).Transform.P = v211
		// cache miss here, however I need the shape list below
		body = bodies + uintptr((*b2BodySim)(unsafe.Pointer(sim)).BodyId)*128
		(*b2Body)(unsafe.Pointer(body)).BodyMoveIndex = simIndex
		(*(*BodyMoveEvent)(unsafe.Add(nil, moveEvents+uintptr(simIndex)*40))).Transform = (*b2BodySim)(unsafe.Pointer(sim)).Transform
		(*(*BodyMoveEvent)(unsafe.Add(nil, moveEvents+uintptr(simIndex)*40))).BodyId = BodyId{
			Index1:     (*b2BodySim)(unsafe.Pointer(sim)).BodyId + int32(1),
			World0:     worldId,
			Generation: (*b2Body)(unsafe.Pointer(body)).Generation,
		}
		(*(*BodyMoveEvent)(unsafe.Add(nil, moveEvents+uintptr(simIndex)*40))).UserData = (*b2Body)(unsafe.Pointer(body)).UserData
		(*(*BodyMoveEvent)(unsafe.Add(nil, moveEvents+uintptr(simIndex)*40))).FellAsleep = boolUint8(false1 != 0)
		// reset applied force and torque
		(*b2BodySim)(unsafe.Pointer(sim)).Force = b2Vec2_zero
		(*b2BodySim)(unsafe.Pointer(sim)).Torque = float32FromFloat32(0)
		(*b2Body)(unsafe.Pointer(body)).IsSpeedCapped = (*b2BodySim)(unsafe.Pointer(sim)).IsSpeedCapped
		(*b2BodySim)(unsafe.Pointer(sim)).IsSpeedCapped = boolUint8(false1 != 0)
		(*b2BodySim)(unsafe.Pointer(sim)).IsFast = boolUint8(false1 != 0)
		if int32FromUint8(enableSleep) == false1 || int32FromUint8((*b2Body)(unsafe.Pointer(body)).EnableSleep) == false1 || sleepVelocity > (*b2Body)(unsafe.Pointer(body)).SleepThreshold {
			// Body is not sleepy
			(*b2Body)(unsafe.Pointer(body)).SleepTime = float32FromFloat32(0)
			if (*b2Body)(unsafe.Pointer(body)).Type1 == int32(b2_dynamicBody) && enableContinuous != 0 && float32(maxVelocity*timeStep) > float32(float32FromFloat32(0.5)*(*b2BodySim)(unsafe.Pointer(sim)).MinExtent) {
				// This flag is only retained for debug draw
				(*b2BodySim)(unsafe.Pointer(sim)).IsFast = boolUint8(true1 != 0)
				// Store in fast array for the continuous collision stage
				// This is deterministic because the order of TOI sweeps doesn't matter
				if (*b2BodySim)(unsafe.Pointer(sim)).IsBullet != 0 {
					v41 = __atomic_fetch_addInt32(tls, stepContext+112, int32(1), int32FromInt32(__ATOMIC_SEQ_CST))
					goto _42
				_42:
					bulletIndex = v41
					*(*int32)(unsafe.Add(nil, (*b2StepContext)(unsafe.Pointer(stepContext)).BulletBodies+uintptr(bulletIndex)*4)) = simIndex
				} else {
					b2SolveContinuous(tls, world, simIndex)
				}
			} else {
				// Body is safe to advance
				(*b2BodySim)(unsafe.Pointer(sim)).Center0 = (*b2BodySim)(unsafe.Pointer(sim)).Center
				(*b2BodySim)(unsafe.Pointer(sim)).Rotation0 = (*b2BodySim)(unsafe.Pointer(sim)).Transform.Q
			}
		} else {
			// Body is safe to advance and is falling asleep
			(*b2BodySim)(unsafe.Pointer(sim)).Center0 = (*b2BodySim)(unsafe.Pointer(sim)).Center
			(*b2BodySim)(unsafe.Pointer(sim)).Rotation0 = (*b2BodySim)(unsafe.Pointer(sim)).Transform.Q
			*(*float32)(unsafe.Add(nil, body+100)) += timeStep
		}
		v43 = world + 1184
		v41 = (*b2Body)(unsafe.Pointer(body)).IslandId
		if !(0 <= v41 && v41 < (*b2IslandArray)(unsafe.Pointer(v43)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
			__builtin_trap(tls)
		}
		v45 = (*b2IslandArray)(unsafe.Pointer(v43)).Data + uintptr(v41)*56
		goto _46
	_46:
		// Any single body in an island can keep it awake
		island = v45
		if (*b2Body)(unsafe.Pointer(body)).SleepTime < float32FromFloat32(0.5) {
			// keep island awake
			islandIndex = (*b2Island)(unsafe.Pointer(island)).LocalIndex
			v47 = awakeIslandBitSet
			v48 = uint32FromInt32(islandIndex)
			blockIndex = v48 / uint32(64)
			if !(blockIndex < (*b2BitSet)(unsafe.Pointer(v47)).BlockCount) && b2InternalAssertFcn(tls, __ccgo_ts+10859, __ccgo_ts+10891, int32FromInt32(28)) != 0 {
				__builtin_trap(tls)
			}
			*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v47)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v48 % uint32FromInt32(64))
		} else {
			if (*b2Island)(unsafe.Pointer(island)).ConstraintRemoveCount > 0 {
				// body wants to sleep but its island needs splitting first
				if (*b2Body)(unsafe.Pointer(body)).SleepTime > (*b2TaskContext)(unsafe.Pointer(taskContext)).SplitSleepTime {
					// pick the sleepiest candidate
					(*b2TaskContext)(unsafe.Pointer(taskContext)).SplitIslandId = (*b2Body)(unsafe.Pointer(body)).IslandId
					(*b2TaskContext)(unsafe.Pointer(taskContext)).SplitSleepTime = (*b2Body)(unsafe.Pointer(body)).SleepTime
				}
			}
		}
		// Update shapes AABBs
		transform = (*b2BodySim)(unsafe.Pointer(sim)).Transform
		isFast = (*b2BodySim)(unsafe.Pointer(sim)).IsFast
		shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
		for shapeId != -int32(1) {
			v43 = world + 1248
			v41 = shapeId
			if !(0 <= v41 && v41 < (*b2ShapeArray)(unsafe.Pointer(v43)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
				__builtin_trap(tls)
			}
			v45 = (*b2ShapeArray)(unsafe.Pointer(v43)).Data + uintptr(v41)*288
			goto _52
		_52:
			shape = v45
			if isFast != 0 {
				// For fast non-bullet bodies the AABB has already been updated in b2SolveContinuous
				// For fast bullet bodies the AABB will be updated at a later stage
				// Add to enlarged shapes regardless of AABB changes.
				// Bit-set to keep the move array sorted
				v47 = enlargedSimBitSet
				v48 = uint32FromInt32(simIndex)
				blockIndex = v48 / uint32(64)
				if !(blockIndex < (*b2BitSet)(unsafe.Pointer(v47)).BlockCount) && b2InternalAssertFcn(tls, __ccgo_ts+10859, __ccgo_ts+10891, int32FromInt32(28)) != 0 {
					__builtin_trap(tls)
				}
				*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v47)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v48 % uint32FromInt32(64))
			} else {
				aabb = b2ComputeShapeAABB(tls, shape, transform)
				aabb.LowerBound.X -= speculativeDistance
				aabb.LowerBound.Y -= speculativeDistance
				aabb.UpperBound.X += speculativeDistance
				aabb.UpperBound.Y += speculativeDistance
				(*b2Shape)(unsafe.Pointer(shape)).Aabb = aabb
				if !(int32FromUint8((*b2Shape)(unsafe.Pointer(shape)).EnlargedAABB) == int32FromInt32(false1)) && b2InternalAssertFcn(tls, __ccgo_ts+12743, __ccgo_ts+12460, int32FromInt32(710)) != 0 {
					__builtin_trap(tls)
				}
				v55 = (*b2Shape)(unsafe.Pointer(shape)).FatAABB
				v56 = aabb
				s = boolUint8(true1 != 0)
				s = boolUint8(s != 0 && v55.LowerBound.X <= v56.LowerBound.X)
				s = boolUint8(s != 0 && v55.LowerBound.Y <= v56.LowerBound.Y)
				s = boolUint8(s != 0 && v56.UpperBound.X <= v55.UpperBound.X)
				s = boolUint8(s != 0 && v56.UpperBound.Y <= v55.UpperBound.Y)
				v57 = s
				goto _58
			_58:
				if int32FromUint8(v57) == false1 {
					fatAABB.LowerBound.X = aabb.LowerBound.X - aabbMargin
					fatAABB.LowerBound.Y = aabb.LowerBound.Y - aabbMargin
					fatAABB.UpperBound.X = aabb.UpperBound.X + aabbMargin
					fatAABB.UpperBound.Y = aabb.UpperBound.Y + aabbMargin
					(*b2Shape)(unsafe.Pointer(shape)).FatAABB = fatAABB
					(*b2Shape)(unsafe.Pointer(shape)).EnlargedAABB = boolUint8(true1 != 0)
					// Bit-set to keep the move array sorted
					v43 = enlargedSimBitSet
					v48 = uint32FromInt32(simIndex)
					blockIndex = v48 / uint32(64)
					if !(blockIndex < (*b2BitSet)(unsafe.Pointer(v43)).BlockCount) && b2InternalAssertFcn(tls, __ccgo_ts+10859, __ccgo_ts+10891, int32FromInt32(28)) != 0 {
						__builtin_trap(tls)
					}
					*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v43)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v48 % uint32FromInt32(64))
				}
			}
			shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
		}
		goto _1
	_1:
		;
		simIndex = simIndex + 1
	}
}

/*
 typedef enum b2SolverStageType
{
	b2_stagePrepareJoints,
	b2_stagePrepareContacts,
	b2_stageIntegrateVelocities,
	b2_stageWarmStart,
	b2_stageSolve,
	b2_stageIntegratePositions,
	b2_stageRelax,
	b2_stageRestitution,
	b2_stageStoreImpulses
} b2SolverStageType;

typedef enum b2SolverBlockType
{
	b2_bodyBlock,
	b2_jointBlock,
	b2_contactBlock,
	b2_graphJointBlock,
	b2_graphContactBlock
} b2SolverBlockType;
*/

func b2ExecuteBlock(tls *_Stack, stage uintptr, context uintptr, block uintptr) {
	var blockType b2SolverBlockType1
	var endIndex, startIndex int32
	var stageType b2SolverStageType1
	_, _, _, _ = blockType, endIndex, stageType, startIndex
	stageType = (*b2SolverStage)(unsafe.Pointer(stage)).Type1
	blockType = int32((*b2SolverBlock)(unsafe.Pointer(block)).BlockType)
	startIndex = (*b2SolverBlock)(unsafe.Pointer(block)).StartIndex
	endIndex = startIndex + int32((*b2SolverBlock)(unsafe.Pointer(block)).Count)
	switch stageType {
	case int32(b2_stagePrepareJoints):
		b2PrepareJointsTask(tls, startIndex, endIndex, context)
	case int32(b2_stagePrepareContacts):
		b2PrepareContactsTask(tls, startIndex, endIndex, context)
	case int32(b2_stageIntegrateVelocities):
		b2IntegrateVelocitiesTask(tls, startIndex, endIndex, context)
	case int32(b2_stageWarmStart):
		if blockType == int32(b2_graphContactBlock) {
			b2WarmStartContactsTask(tls, startIndex, endIndex, context, (*b2SolverStage)(unsafe.Pointer(stage)).ColorIndex)
		} else {
			if blockType == int32(b2_graphJointBlock) {
				b2WarmStartJointsTask(tls, startIndex, endIndex, context, (*b2SolverStage)(unsafe.Pointer(stage)).ColorIndex)
			}
		}
	case int32(b2_stageSolve):
		if blockType == int32(b2_graphContactBlock) {
			b2SolveContactsTask(tls, startIndex, endIndex, context, (*b2SolverStage)(unsafe.Pointer(stage)).ColorIndex, boolUint8(true1 != 0))
		} else {
			if blockType == int32(b2_graphJointBlock) {
				b2SolveJointsTask(tls, startIndex, endIndex, context, (*b2SolverStage)(unsafe.Pointer(stage)).ColorIndex, boolUint8(true1 != 0))
			}
		}
	case int32(b2_stageIntegratePositions):
		b2IntegratePositionsTask(tls, startIndex, endIndex, context)
	case int32(b2_stageRelax):
		if blockType == int32(b2_graphContactBlock) {
			b2SolveContactsTask(tls, startIndex, endIndex, context, (*b2SolverStage)(unsafe.Pointer(stage)).ColorIndex, boolUint8(false1 != 0))
		} else {
			if blockType == int32(b2_graphJointBlock) {
				b2SolveJointsTask(tls, startIndex, endIndex, context, (*b2SolverStage)(unsafe.Pointer(stage)).ColorIndex, boolUint8(false1 != 0))
			}
		}
	case int32(b2_stageRestitution):
		if blockType == int32(b2_graphContactBlock) {
			b2ApplyRestitutionTask(tls, startIndex, endIndex, context, (*b2SolverStage)(unsafe.Pointer(stage)).ColorIndex)
		}
	case int32(b2_stageStoreImpulses):
		b2StoreImpulsesTask(tls, startIndex, endIndex, context)
		break
	}
}

func GetWorkerStartIndex(tls *_Stack, workerIndex int32, blockCount int32, workerCount int32) (r int32) {
	var blocksPerWorker, remainder, v1, v2, v3, v4 int32
	_, _, _, _, _, _ = blocksPerWorker, remainder, v1, v2, v3, v4
	if blockCount <= workerCount {
		if workerIndex < blockCount {
			v1 = workerIndex
		} else {
			v1 = -int32(1)
		}
		return v1
	}
	blocksPerWorker = blockCount / workerCount
	remainder = blockCount - blocksPerWorker*workerCount
	v1 = remainder
	v2 = workerIndex
	if v1 < v2 {
		v4 = v1
	} else {
		v4 = v2
	}
	v3 = v4
	goto _5
_5:
	return blocksPerWorker*workerIndex + v3
}

func b2ExecuteStage(tls *_Stack, stage uintptr, context uintptr, previousSyncIndex int32, syncIndex int32, workerIndex int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var blockCount, blockIndex, completedCount, expectedSyncIndex, startIndex, v1 int32
	var blocks uintptr
	var v2 uint8
	_, _, _, _, _, _, _, _ = blockCount, blockIndex, blocks, completedCount, expectedSyncIndex, startIndex, v1, v2
	completedCount = 0
	blocks = (*b2SolverStage)(unsafe.Pointer(stage)).Blocks
	blockCount = (*b2SolverStage)(unsafe.Pointer(stage)).BlockCount
	expectedSyncIndex = previousSyncIndex
	startIndex = GetWorkerStartIndex(tls, workerIndex, blockCount, (*b2StepContext)(unsafe.Pointer(context)).WorkerCount)
	if startIndex == -int32(1) {
		return
	}
	if !(0 <= startIndex && startIndex < blockCount) && b2InternalAssertFcn(tls, __ccgo_ts+12772, __ccgo_ts+12460, int32FromInt32(856)) != 0 {
		__builtin_trap(tls)
	}
	blockIndex = startIndex
	for {
		v1 = expectedSyncIndex
		*(*int32)(unsafe.Pointer(bp)) = v1
		v2 = __builtin___atomic_compare_exchange_n(tls, blocks+uintptr(blockIndex)*12+8, bp, syncIndex, boolUint8(false1 != 0), int32(__ATOMIC_SEQ_CST), int32(__ATOMIC_SEQ_CST))
		goto _3
	_3:
		if !(int32FromUint8(v2) == int32(true1)) {
			break
		}
		if !((*b2SolverStage)(unsafe.Pointer(stage)).Type1 != int32(b2_stagePrepareContacts) || syncIndex < int32(2)) && b2InternalAssertFcn(tls, __ccgo_ts+12815, __ccgo_ts+12460, int32FromInt32(862)) != 0 {
			__builtin_trap(tls)
		}
		if !(completedCount < blockCount) && b2InternalAssertFcn(tls, __ccgo_ts+12871, __ccgo_ts+12460, int32FromInt32(864)) != 0 {
			__builtin_trap(tls)
		}
		b2ExecuteBlock(tls, stage, context, blocks+uintptr(blockIndex)*12)
		completedCount = completedCount + int32(1)
		blockIndex = blockIndex + int32(1)
		if blockIndex >= blockCount {
			// Keep looking for work
			blockIndex = 0
		}
		expectedSyncIndex = previousSyncIndex
	}
	// Search backwards for blocks
	blockIndex = startIndex - int32(1)
	for int32(true1) != 0 {
		if blockIndex < 0 {
			blockIndex = blockCount - int32(1)
		}
		expectedSyncIndex = previousSyncIndex
		v1 = expectedSyncIndex
		*(*int32)(unsafe.Pointer(bp)) = v1
		v2 = __builtin___atomic_compare_exchange_n(tls, blocks+uintptr(blockIndex)*12+8, bp, syncIndex, boolUint8(false1 != 0), int32(__ATOMIC_SEQ_CST), int32(__ATOMIC_SEQ_CST))
		goto _6
	_6:
		if int32FromUint8(v2) == false1 {
			break
		}
		b2ExecuteBlock(tls, stage, context, blocks+uintptr(blockIndex)*12)
		completedCount = completedCount + int32(1)
		blockIndex = blockIndex - int32(1)
	}
	_ = __atomic_fetch_addInt32(tls, stage+24, completedCount, int32FromInt32(__ATOMIC_SEQ_CST))
	goto _7
_7:
}

func b2ExecuteMainStage(tls *_Stack, stage uintptr, context uintptr, syncBits uint32) {
	var blockCount, previousSyncIndex, syncIndex, v1 int32
	_, _, _, _ = blockCount, previousSyncIndex, syncIndex, v1
	blockCount = (*b2SolverStage)(unsafe.Pointer(stage)).BlockCount
	if blockCount == 0 {
		return
	}
	if blockCount == int32(1) {
		b2ExecuteBlock(tls, stage, context, (*b2SolverStage)(unsafe.Pointer(stage)).Blocks)
	} else {
		atomicStoreNUint32(context+232, syncBits, int32FromInt32(__ATOMIC_SEQ_CST))
		syncIndex = int32FromUint32(syncBits >> int32FromInt32(16) & uint32(0xFFFF))
		if !(syncIndex > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+12899, __ccgo_ts+12460, int32FromInt32(920)) != 0 {
			__builtin_trap(tls)
		}
		previousSyncIndex = syncIndex - int32(1)
		b2ExecuteStage(tls, stage, context, previousSyncIndex, syncIndex, 0)
		// todo consider using the cycle counter as well
		for {
			v1 = atomicLoadNInt32(stage+24, int32FromInt32(__ATOMIC_SEQ_CST))
			goto _2
		_2:
			if !(v1 != blockCount) {
				break
			}
			b2Pause(tls)
		}
		atomicStoreNInt32(stage+24, 0, int32FromInt32(__ATOMIC_SEQ_CST))
	}
}

const _B2_TIME_TO_SLEEP2 = "0.5f"
const _UINT64_MAX8 = "0xffffffffffffffffu"

var b2_identityBodyState15 = b2BodyState{
	DeltaRotation: Rot{
		C: float32FromFloat32(1),
	},
}

func b2TrySleepIsland(tls *_Stack, world uintptr, islandId int32) {
	var awakeBodyIndex, bodyId, colorIndex, colorIndex1, contactId, contactId1, contactKey, edgeIndex, islandIndex, jointId, localIndex, localIndex1, localIndex2, movedId, movedId1, movedIndex, movedIndex1, movedIndex2, movedIndex3, movedIndex4, movedIndex5, movedIndex6, movedIslandId, movedIslandIndex, movedLocalIndex, movedLocalIndex1, newCapacity, newCapacity1, newCapacity2, newCapacity3, newCapacity4, otherBodyId, otherEdgeIndex, sleepBodyIndex, sleepContactIndex, sleepJointIndex, sleepSetId, v2, v6, v8 int32
	var awakeContactSim, awakeJointSim, awakeSet, awakeSim, body, color, color1, contact, contact1, contactSim, disabledContactSim, disabledSet, island, joint, moveEvent, movedBody, movedContact, movedContact1, movedContactSim, movedContactSim1, movedIsland, movedIslandSim, movedJoint, movedJointSim, movedSim, otherBody, sleepBodySim, sleepContactSim, sleepIsland, sleepJointSim, sleepSet, v1, v3, v5, v7 uintptr
	var blockIndex, v75 uint32_t
	var set b2SolverSet
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = awakeBodyIndex, awakeContactSim, awakeJointSim, awakeSet, awakeSim, blockIndex, body, bodyId, color, color1, colorIndex, colorIndex1, contact, contact1, contactId, contactId1, contactKey, contactSim, disabledContactSim, disabledSet, edgeIndex, island, islandIndex, joint, jointId, localIndex, localIndex1, localIndex2, moveEvent, movedBody, movedContact, movedContact1, movedContactSim, movedContactSim1, movedId, movedId1, movedIndex, movedIndex1, movedIndex2, movedIndex3, movedIndex4, movedIndex5, movedIndex6, movedIsland, movedIslandId, movedIslandIndex, movedIslandSim, movedJoint, movedJointSim, movedLocalIndex, movedLocalIndex1, movedSim, newCapacity, newCapacity1, newCapacity2, newCapacity3, newCapacity4, otherBody, otherBodyId, otherEdgeIndex, set, sleepBodyIndex, sleepBodySim, sleepContactIndex, sleepContactSim, sleepIsland, sleepJointIndex, sleepJointSim, sleepSet, sleepSetId, v1, v2, v3, v5, v6, v7, v75, v8
	v1 = world + 1184
	v2 = islandId
	if !(0 <= v2 && v2 < (*b2IslandArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2IslandArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*56
	goto _4
_4:
	island = v3
	if !((*b2Island)(unsafe.Pointer(island)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+14347, __ccgo_ts+14063, int32FromInt32(159)) != 0 {
		__builtin_trap(tls)
	}
	// cannot put an island to sleep while it has a pending split
	if (*b2Island)(unsafe.Pointer(island)).ConstraintRemoveCount > 0 {
		return
	}
	// island is sleeping
	// - create new sleeping solver set
	// - move island to sleeping solver set
	// - identify non-touching contacts that should move to sleeping solver set or disabled set
	// - remove old island
	// - fix island
	sleepSetId = b2AllocId(tls, world+1040)
	if sleepSetId == (*b2World)(unsafe.Pointer(world)).SolverSets.Count {
		set = b2SolverSet{}
		set.SetIndex = -int32(1)
		v1 = world + 1064
		if (*b2SolverSetArray)(unsafe.Pointer(v1)).Count == (*b2SolverSetArray)(unsafe.Pointer(v1)).Capacity {
			if (*b2SolverSetArray)(unsafe.Pointer(v1)).Capacity < int32(2) {
				v2 = int32(2)
			} else {
				v2 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Capacity + (*b2SolverSetArray)(unsafe.Pointer(v1)).Capacity>>int32(1)
			}
			newCapacity = v2
			b2SolverSetArray_Reserve(tls, v1, newCapacity)
		}
		*(*b2SolverSet)(unsafe.Add(nil, (*b2SolverSetArray)(unsafe.Pointer(v1)).Data+uintptr((*b2SolverSetArray)(unsafe.Pointer(v1)).Count)*88)) = set
		*(*int32)(unsafe.Add(nil, v1+8)) += int32(1)
	}
	v1 = world + 1064
	v2 = sleepSetId
	if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*88
	goto _10
_10:
	sleepSet = v3
	*(*b2SolverSet)(unsafe.Add(nil, sleepSet)) = b2SolverSet{}
	v1 = world + 1064
	v2 = int32(b2_awakeSet)
	if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*88
	goto _14
_14:
	// grab awake set after creating the sleep set because the solver set array may have been resized
	awakeSet = v3
	if !(0 <= (*b2Island)(unsafe.Pointer(island)).LocalIndex && (*b2Island)(unsafe.Pointer(island)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(awakeSet)).IslandSims.Count) && b2InternalAssertFcn(tls, __ccgo_ts+14379, __ccgo_ts+14063, int32FromInt32(186)) != 0 {
		__builtin_trap(tls)
	}
	(*b2SolverSet)(unsafe.Pointer(sleepSet)).SetIndex = sleepSetId
	(*b2SolverSet)(unsafe.Pointer(sleepSet)).BodySims = b2BodySimArray_Create(tls, (*b2Island)(unsafe.Pointer(island)).BodyCount)
	(*b2SolverSet)(unsafe.Pointer(sleepSet)).ContactSims = b2ContactSimArray_Create(tls, (*b2Island)(unsafe.Pointer(island)).ContactCount)
	(*b2SolverSet)(unsafe.Pointer(sleepSet)).JointSims = b2JointSimArray_Create(tls, (*b2Island)(unsafe.Pointer(island)).JointCount)
	// move awake bodies to sleeping set
	// this shuffles around bodies in the awake set
	v1 = world + 1064
	v2 = int32(b2_disabledSet)
	if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*88
	goto _18
_18:
	disabledSet = v3
	bodyId = (*b2Island)(unsafe.Pointer(island)).HeadBody
	for bodyId != -int32(1) {
		v5 = world + 1024
		v6 = bodyId
		if !(0 <= v6 && v6 < (*b2BodyArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
			__builtin_trap(tls)
		}
		v7 = (*b2BodyArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*128
		goto _22
	_22:
		body = v7
		if !((*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+1662, __ccgo_ts+14063, int32FromInt32(201)) != 0 {
			__builtin_trap(tls)
		}
		if !((*b2Body)(unsafe.Pointer(body)).IslandId == islandId) && b2InternalAssertFcn(tls, __ccgo_ts+14454, __ccgo_ts+14063, int32FromInt32(202)) != 0 {
			__builtin_trap(tls)
		}
		// Update the body move event to indicate this body fell asleep
		// It could happen the body is forced asleep before it ever moves.
		if (*b2Body)(unsafe.Pointer(body)).BodyMoveIndex != -int32(1) {
			v1 = world + 1328
			v2 = (*b2Body)(unsafe.Pointer(body)).BodyMoveIndex
			if !(0 <= v2 && v2 < (*b2BodyMoveEventArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+12641, int32FromInt32(185)) != 0 {
				__builtin_trap(tls)
			}
			v3 = (*b2BodyMoveEventArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*40
			goto _26
		_26:
			moveEvent = v3
			if !((*BodyMoveEvent)(unsafe.Pointer(moveEvent)).BodyId.Index1-int32FromInt32(1) == bodyId) && b2InternalAssertFcn(tls, __ccgo_ts+14481, __ccgo_ts+14063, int32FromInt32(209)) != 0 {
				__builtin_trap(tls)
			}
			if !(int32FromUint16((*BodyMoveEvent)(unsafe.Pointer(moveEvent)).BodyId.Generation) == int32FromUint16((*b2Body)(unsafe.Pointer(body)).Generation)) && b2InternalAssertFcn(tls, __ccgo_ts+14520, __ccgo_ts+14063, int32FromInt32(210)) != 0 {
				__builtin_trap(tls)
			}
			(*BodyMoveEvent)(unsafe.Pointer(moveEvent)).FellAsleep = boolUint8(true1 != 0)
			(*b2Body)(unsafe.Pointer(body)).BodyMoveIndex = -int32(1)
		}
		awakeBodyIndex = (*b2Body)(unsafe.Pointer(body)).LocalIndex
		v1 = awakeSet
		v2 = awakeBodyIndex
		if !(0 <= v2 && v2 < (*b2BodySimArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2BodySimArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*100
		goto _30
	_30:
		awakeSim = v3
		// move body sim to sleep set
		sleepBodyIndex = (*b2SolverSet)(unsafe.Pointer(sleepSet)).BodySims.Count
		v5 = sleepSet
		if (*b2BodySimArray)(unsafe.Pointer(v5)).Count == (*b2BodySimArray)(unsafe.Pointer(v5)).Capacity {
			if (*b2BodySimArray)(unsafe.Pointer(v5)).Capacity < int32(2) {
				v6 = int32(2)
			} else {
				v6 = (*b2BodySimArray)(unsafe.Pointer(v5)).Capacity + (*b2BodySimArray)(unsafe.Pointer(v5)).Capacity>>int32(1)
			}
			newCapacity1 = v6
			b2BodySimArray_Reserve(tls, v5, newCapacity1)
		}
		*(*int32)(unsafe.Add(nil, v5+8)) += int32(1)
		v7 = (*b2BodySimArray)(unsafe.Pointer(v5)).Data + uintptr((*b2BodySimArray)(unsafe.Pointer(v5)).Count-int32FromInt32(1))*100
		goto _34
	_34:
		sleepBodySim = v7
		memcpy(tls, sleepBodySim, awakeSim, uint64(100))
		v1 = awakeSet
		v2 = awakeBodyIndex
		if !(0 <= v2 && v2 < (*b2BodySimArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
			__builtin_trap(tls)
		}
		movedIndex = -int32(1)
		if v2 != (*b2BodySimArray)(unsafe.Pointer(v1)).Count-int32FromInt32(1) {
			movedIndex = (*b2BodySimArray)(unsafe.Pointer(v1)).Count - int32(1)
			*(*b2BodySim)(unsafe.Add(nil, (*b2BodySimArray)(unsafe.Pointer(v1)).Data+uintptr(v2)*100)) = *(*b2BodySim)(unsafe.Add(nil, (*b2BodySimArray)(unsafe.Pointer(v1)).Data+uintptr(movedIndex)*100))
		}
		*(*int32)(unsafe.Add(nil, v1+8)) -= int32(1)
		v6 = movedIndex
		goto _38
	_38:
		movedIndex5 = v6
		if movedIndex5 != -int32(1) {
			// fix local index on moved element
			movedSim = (*b2SolverSet)(unsafe.Pointer(awakeSet)).BodySims.Data + uintptr(awakeBodyIndex)*100
			movedId = (*b2BodySim)(unsafe.Pointer(movedSim)).BodyId
			v3 = world + 1024
			v8 = movedId
			if !(0 <= v8 && v8 < (*b2BodyArray)(unsafe.Pointer(v3)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
				__builtin_trap(tls)
			}
			v5 = (*b2BodyArray)(unsafe.Pointer(v3)).Data + uintptr(v8)*128
			goto _42
		_42:
			movedBody = v5
			if !((*b2Body)(unsafe.Pointer(movedBody)).LocalIndex == movedIndex5) && b2InternalAssertFcn(tls, __ccgo_ts+1407, __ccgo_ts+14063, int32FromInt32(230)) != 0 {
				__builtin_trap(tls)
			}
			(*b2Body)(unsafe.Pointer(movedBody)).LocalIndex = awakeBodyIndex
		}
		// destroy state, no need to clone
		v1 = awakeSet + 16
		v2 = awakeBodyIndex
		if !(0 <= v2 && v2 < (*b2BodyStateArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(194)) != 0 {
			__builtin_trap(tls)
		}
		movedIndex1 = -int32(1)
		if v2 != (*b2BodyStateArray)(unsafe.Pointer(v1)).Count-int32FromInt32(1) {
			movedIndex1 = (*b2BodyStateArray)(unsafe.Pointer(v1)).Count - int32(1)
			*(*b2BodyState)(unsafe.Add(nil, (*b2BodyStateArray)(unsafe.Pointer(v1)).Data+uintptr(v2)*32)) = *(*b2BodyState)(unsafe.Add(nil, (*b2BodyStateArray)(unsafe.Pointer(v1)).Data+uintptr(movedIndex1)*32))
		}
		*(*int32)(unsafe.Add(nil, v1+8)) -= int32(1)
		_ = movedIndex1
		goto _45
	_45:
		;
		(*b2Body)(unsafe.Pointer(body)).SetIndex = sleepSetId
		(*b2Body)(unsafe.Pointer(body)).LocalIndex = sleepBodyIndex
		// Move non-touching contacts to the disabled set.
		// Non-touching contacts may exist between sleeping islands and there is no clear ownership.
		contactKey = (*b2Body)(unsafe.Pointer(body)).HeadContactKey
		for contactKey != -int32(1) {
			contactId = contactKey >> int32(1)
			edgeIndex = contactKey & int32(1)
			v1 = world + 1144
			v2 = contactId
			if !(0 <= v2 && v2 < (*b2ContactArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
				__builtin_trap(tls)
			}
			v3 = (*b2ContactArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*68
			goto _49
		_49:
			contact = v3
			if !((*b2Contact)(unsafe.Pointer(contact)).SetIndex == int32(b2_awakeSet) || (*b2Contact)(unsafe.Pointer(contact)).SetIndex == int32(b2_disabledSet)) && b2InternalAssertFcn(tls, __ccgo_ts+14569, __ccgo_ts+14063, int32FromInt32(250)) != 0 {
				__builtin_trap(tls)
			}
			contactKey = (*(*b2ContactEdge)(unsafe.Add(nil, contact+12+uintptr(edgeIndex)*12))).NextKey
			if (*b2Contact)(unsafe.Pointer(contact)).SetIndex == int32(b2_disabledSet) {
				// already moved to disabled set by another body in the island
				continue
			}
			if (*b2Contact)(unsafe.Pointer(contact)).ColorIndex != -int32(1) {
				// contact is touching and will be moved separately
				if !((*b2Contact)(unsafe.Pointer(contact)).Flags&uint32(b2_contactTouchingFlag) != uint32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+7306, __ccgo_ts+14063, int32FromInt32(262)) != 0 {
					__builtin_trap(tls)
				}
				continue
			}
			// the other body may still be awake, it still may go to sleep and then it will be responsible
			// for moving this contact to the disabled set.
			otherEdgeIndex = edgeIndex ^ int32(1)
			otherBodyId = (*(*b2ContactEdge)(unsafe.Add(nil, contact+12+uintptr(otherEdgeIndex)*12))).BodyId
			v1 = world + 1024
			v2 = otherBodyId
			if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
				__builtin_trap(tls)
			}
			v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
			goto _53
		_53:
			otherBody = v3
			if (*b2Body)(unsafe.Pointer(otherBody)).SetIndex == int32(b2_awakeSet) {
				continue
			}
			localIndex = (*b2Contact)(unsafe.Pointer(contact)).LocalIndex
			v1 = awakeSet + 48
			v2 = localIndex
			if !(0 <= v2 && v2 < (*b2ContactSimArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(147)) != 0 {
				__builtin_trap(tls)
			}
			v3 = (*b2ContactSimArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*176
			goto _57
		_57:
			contactSim = v3
			if !((*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.PointCount == int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+14641, __ccgo_ts+14063, int32FromInt32(279)) != 0 {
				__builtin_trap(tls)
			}
			if !((*b2Contact)(unsafe.Pointer(contact)).Flags&uint32(b2_contactTouchingFlag) == uint32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+14678, __ccgo_ts+14063, int32FromInt32(280)) != 0 {
				__builtin_trap(tls)
			}
			// move the non-touching contact to the disabled set
			(*b2Contact)(unsafe.Pointer(contact)).SetIndex = int32(b2_disabledSet)
			(*b2Contact)(unsafe.Pointer(contact)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(disabledSet)).ContactSims.Count
			v1 = disabledSet + 48
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
			goto _61
		_61:
			disabledContactSim = v3
			memcpy(tls, disabledContactSim, contactSim, uint64(176))
			v1 = awakeSet + 48
			v2 = localIndex
			if !(0 <= v2 && v2 < (*b2ContactSimArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(147)) != 0 {
				__builtin_trap(tls)
			}
			movedIndex2 = -int32(1)
			if v2 != (*b2ContactSimArray)(unsafe.Pointer(v1)).Count-int32FromInt32(1) {
				movedIndex2 = (*b2ContactSimArray)(unsafe.Pointer(v1)).Count - int32(1)
				*(*b2ContactSim)(unsafe.Add(nil, (*b2ContactSimArray)(unsafe.Pointer(v1)).Data+uintptr(v2)*176)) = *(*b2ContactSim)(unsafe.Add(nil, (*b2ContactSimArray)(unsafe.Pointer(v1)).Data+uintptr(movedIndex2)*176))
			}
			*(*int32)(unsafe.Add(nil, v1+8)) -= int32(1)
			v6 = movedIndex2
			goto _65
		_65:
			movedLocalIndex = v6
			if movedLocalIndex != -int32(1) {
				// fix moved element
				movedContactSim = (*b2SolverSet)(unsafe.Pointer(awakeSet)).ContactSims.Data + uintptr(localIndex)*176
				v3 = world + 1144
				v8 = (*b2ContactSim)(unsafe.Pointer(movedContactSim)).ContactId
				if !(0 <= v8 && v8 < (*b2ContactArray)(unsafe.Pointer(v3)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
					__builtin_trap(tls)
				}
				v5 = (*b2ContactArray)(unsafe.Pointer(v3)).Data + uintptr(v8)*68
				goto _69
			_69:
				movedContact = v5
				if !((*b2Contact)(unsafe.Pointer(movedContact)).LocalIndex == movedLocalIndex) && b2InternalAssertFcn(tls, __ccgo_ts+14273, __ccgo_ts+14063, int32FromInt32(294)) != 0 {
					__builtin_trap(tls)
				}
				(*b2Contact)(unsafe.Pointer(movedContact)).LocalIndex = localIndex
			}
		}
		bodyId = (*b2Body)(unsafe.Pointer(body)).IslandNext
	}
	// move touching contacts
	// this shuffles contacts in the awake set
	contactId1 = (*b2Island)(unsafe.Pointer(island)).HeadContact
	for contactId1 != -int32(1) {
		v1 = world + 1144
		v2 = contactId1
		if !(0 <= v2 && v2 < (*b2ContactArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2ContactArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*68
		goto _73
	_73:
		contact1 = v3
		if !((*b2Contact)(unsafe.Pointer(contact1)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+3637, __ccgo_ts+14063, int32FromInt32(310)) != 0 {
			__builtin_trap(tls)
		}
		if !((*b2Contact)(unsafe.Pointer(contact1)).IslandId == islandId) && b2InternalAssertFcn(tls, __ccgo_ts+14727, __ccgo_ts+14063, int32FromInt32(311)) != 0 {
			__builtin_trap(tls)
		}
		colorIndex = (*b2Contact)(unsafe.Pointer(contact1)).ColorIndex
		if !(0 <= colorIndex && colorIndex < int32(_B2_GRAPH_COLOR_COUNT)) && b2InternalAssertFcn(tls, __ccgo_ts+3038, __ccgo_ts+14063, int32FromInt32(313)) != 0 {
			__builtin_trap(tls)
		}
		color = world + 328 + uintptr(colorIndex)*56
		// Remove bodies from graph coloring associated with this constraint
		if colorIndex != int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1) {
			// might clear a bit for a static body, but this has no effect
			v1 = color
			v75 = uint32FromInt32((*(*b2ContactEdge)(unsafe.Add(nil, contact1+12))).BodyId)
			blockIndex = v75 / uint32(64)
			if blockIndex >= (*b2BitSet)(unsafe.Pointer(v1)).BlockCount {
				goto _76
			}
			*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v1)).Bits+uintptr(blockIndex)*8)) &= ^(uint64FromInt32(1) << (v75 % uint32FromInt32(64)))
		_76:
			;
			v1 = color
			v75 = uint32FromInt32((*(*b2ContactEdge)(unsafe.Add(nil, contact1+12+1*12))).BodyId)
			blockIndex = v75 / uint32(64)
			if blockIndex >= (*b2BitSet)(unsafe.Pointer(v1)).BlockCount {
				goto _79
			}
			*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v1)).Bits+uintptr(blockIndex)*8)) &= ^(uint64FromInt32(1) << (v75 % uint32FromInt32(64)))
		_79:
		}
		localIndex1 = (*b2Contact)(unsafe.Pointer(contact1)).LocalIndex
		v1 = color + 16
		v2 = localIndex1
		if !(0 <= v2 && v2 < (*b2ContactSimArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(147)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2ContactSimArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*176
		goto _83
	_83:
		awakeContactSim = v3
		sleepContactIndex = (*b2SolverSet)(unsafe.Pointer(sleepSet)).ContactSims.Count
		v5 = sleepSet + 48
		if (*b2ContactSimArray)(unsafe.Pointer(v5)).Count == (*b2ContactSimArray)(unsafe.Pointer(v5)).Capacity {
			if (*b2ContactSimArray)(unsafe.Pointer(v5)).Capacity < int32(2) {
				v6 = int32(2)
			} else {
				v6 = (*b2ContactSimArray)(unsafe.Pointer(v5)).Capacity + (*b2ContactSimArray)(unsafe.Pointer(v5)).Capacity>>int32(1)
			}
			newCapacity2 = v6
			b2ContactSimArray_Reserve(tls, v5, newCapacity2)
		}
		*(*int32)(unsafe.Add(nil, v5+8)) += int32(1)
		v7 = (*b2ContactSimArray)(unsafe.Pointer(v5)).Data + uintptr((*b2ContactSimArray)(unsafe.Pointer(v5)).Count-int32FromInt32(1))*176
		goto _87
	_87:
		sleepContactSim = v7
		memcpy(tls, sleepContactSim, awakeContactSim, uint64(176))
		v1 = color + 16
		v2 = localIndex1
		if !(0 <= v2 && v2 < (*b2ContactSimArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(147)) != 0 {
			__builtin_trap(tls)
		}
		movedIndex2 = -int32(1)
		if v2 != (*b2ContactSimArray)(unsafe.Pointer(v1)).Count-int32FromInt32(1) {
			movedIndex2 = (*b2ContactSimArray)(unsafe.Pointer(v1)).Count - int32(1)
			*(*b2ContactSim)(unsafe.Add(nil, (*b2ContactSimArray)(unsafe.Pointer(v1)).Data+uintptr(v2)*176)) = *(*b2ContactSim)(unsafe.Add(nil, (*b2ContactSimArray)(unsafe.Pointer(v1)).Data+uintptr(movedIndex2)*176))
		}
		*(*int32)(unsafe.Add(nil, v1+8)) -= int32(1)
		v6 = movedIndex2
		goto _91
	_91:
		movedLocalIndex1 = v6
		if movedLocalIndex1 != -int32(1) {
			// fix moved element
			movedContactSim1 = (*b2GraphColor)(unsafe.Pointer(color)).ContactSims.Data + uintptr(localIndex1)*176
			v3 = world + 1144
			v8 = (*b2ContactSim)(unsafe.Pointer(movedContactSim1)).ContactId
			if !(0 <= v8 && v8 < (*b2ContactArray)(unsafe.Pointer(v3)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
				__builtin_trap(tls)
			}
			v5 = (*b2ContactArray)(unsafe.Pointer(v3)).Data + uintptr(v8)*68
			goto _95
		_95:
			movedContact1 = v5
			if !((*b2Contact)(unsafe.Pointer(movedContact1)).LocalIndex == movedLocalIndex1) && b2InternalAssertFcn(tls, __ccgo_ts+14273, __ccgo_ts+14063, int32FromInt32(338)) != 0 {
				__builtin_trap(tls)
			}
			(*b2Contact)(unsafe.Pointer(movedContact1)).LocalIndex = localIndex1
		}
		(*b2Contact)(unsafe.Pointer(contact1)).SetIndex = sleepSetId
		(*b2Contact)(unsafe.Pointer(contact1)).ColorIndex = -int32(1)
		(*b2Contact)(unsafe.Pointer(contact1)).LocalIndex = sleepContactIndex
		contactId1 = (*b2Contact)(unsafe.Pointer(contact1)).IslandNext
	}
	// move joints
	// this shuffles joints in the awake set
	jointId = (*b2Island)(unsafe.Pointer(island)).HeadJoint
	for jointId != -int32(1) {
		v1 = world + 1104
		v2 = jointId
		if !(0 <= v2 && v2 < (*b2JointArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2JointArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*72
		goto _99
	_99:
		joint = v3
		if !((*b2Joint)(unsafe.Pointer(joint)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+1730, __ccgo_ts+14063, int32FromInt32(357)) != 0 {
			__builtin_trap(tls)
		}
		if !((*b2Joint)(unsafe.Pointer(joint)).IslandId == islandId) && b2InternalAssertFcn(tls, __ccgo_ts+14757, __ccgo_ts+14063, int32FromInt32(358)) != 0 {
			__builtin_trap(tls)
		}
		colorIndex1 = (*b2Joint)(unsafe.Pointer(joint)).ColorIndex
		localIndex2 = (*b2Joint)(unsafe.Pointer(joint)).LocalIndex
		if !(0 <= colorIndex1 && colorIndex1 < int32(_B2_GRAPH_COLOR_COUNT)) && b2InternalAssertFcn(tls, __ccgo_ts+3038, __ccgo_ts+14063, int32FromInt32(362)) != 0 {
			__builtin_trap(tls)
		}
		color1 = world + 328 + uintptr(colorIndex1)*56
		v1 = color1 + 32
		v2 = localIndex2
		if !(0 <= v2 && v2 < (*b2JointSimArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(342)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2JointSimArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*196
		goto _103
	_103:
		awakeJointSim = v3
		if colorIndex1 != int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1) {
			// might clear a bit for a static body, but this has no effect
			v5 = color1
			v75 = uint32FromInt32((*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId)
			blockIndex = v75 / uint32(64)
			if blockIndex >= (*b2BitSet)(unsafe.Pointer(v5)).BlockCount {
				goto _106
			}
			*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v5)).Bits+uintptr(blockIndex)*8)) &= ^(uint64FromInt32(1) << (v75 % uint32FromInt32(64)))
		_106:
			;
			v1 = color1
			v75 = uint32FromInt32((*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId)
			blockIndex = v75 / uint32(64)
			if blockIndex >= (*b2BitSet)(unsafe.Pointer(v1)).BlockCount {
				goto _109
			}
			*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v1)).Bits+uintptr(blockIndex)*8)) &= ^(uint64FromInt32(1) << (v75 % uint32FromInt32(64)))
		_109:
		}
		sleepJointIndex = (*b2SolverSet)(unsafe.Pointer(sleepSet)).JointSims.Count
		v1 = sleepSet + 32
		if (*b2JointSimArray)(unsafe.Pointer(v1)).Count == (*b2JointSimArray)(unsafe.Pointer(v1)).Capacity {
			if (*b2JointSimArray)(unsafe.Pointer(v1)).Capacity < int32(2) {
				v2 = int32(2)
			} else {
				v2 = (*b2JointSimArray)(unsafe.Pointer(v1)).Capacity + (*b2JointSimArray)(unsafe.Pointer(v1)).Capacity>>int32(1)
			}
			newCapacity4 = v2
			b2JointSimArray_Reserve(tls, v1, newCapacity4)
		}
		*(*int32)(unsafe.Add(nil, v1+8)) += int32(1)
		v3 = (*b2JointSimArray)(unsafe.Pointer(v1)).Data + uintptr((*b2JointSimArray)(unsafe.Pointer(v1)).Count-int32FromInt32(1))*196
		goto _113
	_113:
		sleepJointSim = v3
		memcpy(tls, sleepJointSim, awakeJointSim, uint64(196))
		v1 = color1 + 32
		v2 = localIndex2
		if !(0 <= v2 && v2 < (*b2JointSimArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(342)) != 0 {
			__builtin_trap(tls)
		}
		movedIndex4 = -int32(1)
		if v2 != (*b2JointSimArray)(unsafe.Pointer(v1)).Count-int32FromInt32(1) {
			movedIndex4 = (*b2JointSimArray)(unsafe.Pointer(v1)).Count - int32(1)
			*(*b2JointSim)(unsafe.Add(nil, (*b2JointSimArray)(unsafe.Pointer(v1)).Data+uintptr(v2)*196)) = *(*b2JointSim)(unsafe.Add(nil, (*b2JointSimArray)(unsafe.Pointer(v1)).Data+uintptr(movedIndex4)*196))
		}
		*(*int32)(unsafe.Add(nil, v1+8)) -= int32(1)
		v6 = movedIndex4
		goto _117
	_117:
		movedIndex6 = v6
		if movedIndex6 != -int32(1) {
			// fix moved element
			movedJointSim = (*b2GraphColor)(unsafe.Pointer(color1)).JointSims.Data + uintptr(localIndex2)*196
			movedId1 = (*b2JointSim)(unsafe.Pointer(movedJointSim)).JointId
			v3 = world + 1104
			v8 = movedId1
			if !(0 <= v8 && v8 < (*b2JointArray)(unsafe.Pointer(v3)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
				__builtin_trap(tls)
			}
			v5 = (*b2JointArray)(unsafe.Pointer(v3)).Data + uintptr(v8)*72
			goto _121
		_121:
			movedJoint = v5
			if !((*b2Joint)(unsafe.Pointer(movedJoint)).LocalIndex == movedIndex6) && b2InternalAssertFcn(tls, __ccgo_ts+3280, __ccgo_ts+14063, int32FromInt32(386)) != 0 {
				__builtin_trap(tls)
			}
			(*b2Joint)(unsafe.Pointer(movedJoint)).LocalIndex = localIndex2
		}
		(*b2Joint)(unsafe.Pointer(joint)).SetIndex = sleepSetId
		(*b2Joint)(unsafe.Pointer(joint)).ColorIndex = -int32(1)
		(*b2Joint)(unsafe.Pointer(joint)).LocalIndex = sleepJointIndex
		jointId = (*b2Joint)(unsafe.Pointer(joint)).IslandNext
	}
	// move island struct
	if !((*b2Island)(unsafe.Pointer(island)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+14347, __ccgo_ts+14063, int32FromInt32(400)) != 0 {
		__builtin_trap(tls)
	}
	islandIndex = (*b2Island)(unsafe.Pointer(island)).LocalIndex
	v1 = sleepSet + 64
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
	goto _125
_125:
	sleepIsland = v3
	(*b2IslandSim)(unsafe.Pointer(sleepIsland)).IslandId = islandId
	v1 = awakeSet + 64
	v2 = islandIndex
	if !(0 <= v2 && v2 < (*b2IslandSimArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(89)) != 0 {
		__builtin_trap(tls)
	}
	movedIndex3 = -int32(1)
	if v2 != (*b2IslandSimArray)(unsafe.Pointer(v1)).Count-int32FromInt32(1) {
		movedIndex3 = (*b2IslandSimArray)(unsafe.Pointer(v1)).Count - int32(1)
		*(*b2IslandSim)(unsafe.Add(nil, (*b2IslandSimArray)(unsafe.Pointer(v1)).Data+uintptr(v2)*4)) = *(*b2IslandSim)(unsafe.Add(nil, (*b2IslandSimArray)(unsafe.Pointer(v1)).Data+uintptr(movedIndex3)*4))
	}
	*(*int32)(unsafe.Add(nil, v1+8)) -= int32(1)
	v6 = movedIndex3
	goto _129
_129:
	movedIslandIndex = v6
	if movedIslandIndex != -int32(1) {
		// fix index on moved element
		movedIslandSim = (*b2SolverSet)(unsafe.Pointer(awakeSet)).IslandSims.Data + uintptr(islandIndex)*4
		movedIslandId = (*b2IslandSim)(unsafe.Pointer(movedIslandSim)).IslandId
		v3 = world + 1184
		v8 = movedIslandId
		if !(0 <= v8 && v8 < (*b2IslandArray)(unsafe.Pointer(v3)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
			__builtin_trap(tls)
		}
		v5 = (*b2IslandArray)(unsafe.Pointer(v3)).Data + uintptr(v8)*56
		goto _133
	_133:
		movedIsland = v5
		if !((*b2Island)(unsafe.Pointer(movedIsland)).LocalIndex == movedIslandIndex) && b2InternalAssertFcn(tls, __ccgo_ts+14785, __ccgo_ts+14063, int32FromInt32(413)) != 0 {
			__builtin_trap(tls)
		}
		(*b2Island)(unsafe.Pointer(movedIsland)).LocalIndex = islandIndex
	}
	(*b2Island)(unsafe.Pointer(island)).SetIndex = sleepSetId
	(*b2Island)(unsafe.Pointer(island)).LocalIndex = 0
	b2ValidateSolverSets(tls, world)
}

// todo compare with https://github.com/skeeto/scratch/blob/master/set32/set32.h

func b2CreateSet(tls *_Stack, capacity int32) (r b2HashSet) {
	var set b2HashSet
	var v1, v2 int32
	var v4 uint32_t
	_, _, _, _ = set, v1, v2, v4
	set = b2HashSet{}
	// Capacity must be a power of 2
	if capacity > int32(16) {
		v1 = capacity
		if v1 <= int32FromInt32(1) {
			v2 = int32(1)
			goto _3
		}
		v4 = uint32FromInt32(__builtin_clz(tls, uint32FromInt32(v1)-uint32(1)))
		goto _5
	_5:
		v2 = int32(1) << (int32(32) - int32FromUint32(v4))
		goto _3
	_3:
		set.Capacity = uint32FromInt32(v2)
	} else {
		set.Capacity = uint32(16)
	}
	set.Count = uint32(0)
	set.Items = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(16)))
	memset(tls, set.Items, 0, uint64FromInt32(capacity)*uint64(16))
	return set
}

func b2DestroySet(tls *_Stack, set uintptr) {
	b2Free(tls, (*b2HashSet)(unsafe.Pointer(set)).Items, int32FromUint64(uint64((*b2HashSet)(unsafe.Pointer(set)).Capacity)*uint64(16)))
	(*b2HashSet)(unsafe.Pointer(set)).Items = uintptrFromInt32(0)
	(*b2HashSet)(unsafe.Pointer(set)).Count = uint32(0)
	(*b2HashSet)(unsafe.Pointer(set)).Capacity = uint32(0)
}

func b2ClearSet(tls *_Stack, set uintptr) {
	(*b2HashSet)(unsafe.Pointer(set)).Count = uint32(0)
	memset(tls, (*b2HashSet)(unsafe.Pointer(set)).Items, 0, uint64((*b2HashSet)(unsafe.Pointer(set)).Capacity)*uint64(16))
}

// C documentation
//
//	// I need a good hash because the keys are built from pairs of increasing integers.
//	// A simple hash like hash = (integer1 XOR integer2) has many collisions.
//	// https://lemire.me/blog/2018/08/15/fast-strongly-universal-64-bit-hashing-everywhere/
//	// https://preshing.com/20130107/this-hash-set-is-faster-than-a-judy-array/
//	// todo try: https://www.jandrewrogers.com/2019/02/12/fast-perfect-hashing/
//	// todo try: https://probablydance.com/2018/06/16/fibonacci-hashing-the-optimization-that-the-world-forgot-or-a-better-alternative-to-integer-modulo/
func b2KeyHash(tls *_Stack, key uint64) (r uint32) {
	var h uint64_t
	_ = h
	// Murmur hash
	h = key
	h = h ^ h>>int32(33)
	h = uint64(h * uint64FromUint64(0xff51afd7ed558ccd))
	h = h ^ h>>int32(33)
	h = uint64(h * uint64FromUint64(0xc4ceb9fe1a85ec53))
	h = h ^ h>>int32(33)
	return uint32(h)
}

func b2FindSlot(tls *_Stack, set uintptr, key uint64, hash uint32) (r int32) {
	var capacity uint32_t
	var index int32
	var items uintptr
	_, _, _ = capacity, index, items
	capacity = (*b2HashSet)(unsafe.Pointer(set)).Capacity
	index = int32FromUint32(hash & (capacity - uint32(1)))
	items = (*b2HashSet)(unsafe.Pointer(set)).Items
	for (*(*b2SetItem)(unsafe.Add(nil, items+uintptr(index)*16))).Hash != uint32(0) && (*(*b2SetItem)(unsafe.Add(nil, items+uintptr(index)*16))).Key != key {
		index = int32FromUint32(uint32FromInt32(index+int32FromInt32(1)) & (capacity - uint32(1)))
	}
	return index
}

func b2AddKeyHaveCapacity(tls *_Stack, set uintptr, key uint64, hash uint32) {
	var index int32
	var items uintptr
	_, _ = index, items
	index = b2FindSlot(tls, set, key, hash)
	items = (*b2HashSet)(unsafe.Pointer(set)).Items
	if !((*(*b2SetItem)(unsafe.Add(nil, items+uintptr(index)*16))).Hash == uint32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+15019, __ccgo_ts+15042, int32FromInt32(98)) != 0 {
		__builtin_trap(tls)
	}
	(*(*b2SetItem)(unsafe.Add(nil, items+uintptr(index)*16))).Key = key
	(*(*b2SetItem)(unsafe.Add(nil, items+uintptr(index)*16))).Hash = hash
	*(*uint32_t)(unsafe.Add(nil, set+12)) += uint32(1)
}

func b2GrowTable(tls *_Stack, set uintptr) {
	var i, oldCapacity, oldCount uint32_t
	var item, oldItems uintptr
	_, _, _, _, _ = i, item, oldCapacity, oldCount, oldItems
	oldCount = (*b2HashSet)(unsafe.Pointer(set)).Count
	_ = uint64FromInt64(4)
	oldCapacity = (*b2HashSet)(unsafe.Pointer(set)).Capacity
	oldItems = (*b2HashSet)(unsafe.Pointer(set)).Items
	(*b2HashSet)(unsafe.Pointer(set)).Count = uint32(0)
	// Capacity must be a power of 2
	(*b2HashSet)(unsafe.Pointer(set)).Capacity = uint32(2) * oldCapacity
	(*b2HashSet)(unsafe.Pointer(set)).Items = b2Alloc(tls, int32FromUint64(uint64((*b2HashSet)(unsafe.Pointer(set)).Capacity)*uint64(16)))
	memset(tls, (*b2HashSet)(unsafe.Pointer(set)).Items, 0, uint64((*b2HashSet)(unsafe.Pointer(set)).Capacity)*uint64(16))
	// Transfer items into new array
	i = uint32(0)
	for {
		if !(i < oldCapacity) {
			break
		}
		item = oldItems + uintptr(i)*16
		if (*b2SetItem)(unsafe.Pointer(item)).Hash == uint32(0) {
			// this item was empty
			goto _1
		}
		b2AddKeyHaveCapacity(tls, set, (*b2SetItem)(unsafe.Pointer(item)).Key, (*b2SetItem)(unsafe.Pointer(item)).Hash)
		goto _1
	_1:
		;
		i = i + 1
	}
	if !((*b2HashSet)(unsafe.Pointer(set)).Count == oldCount) && b2InternalAssertFcn(tls, __ccgo_ts+15065, __ccgo_ts+15042, int32FromInt32(132)) != 0 {
		__builtin_trap(tls)
	}
	b2Free(tls, oldItems, int32FromUint64(uint64(oldCapacity)*uint64(16)))
}

func b2ContainsKey(tls *_Stack, set uintptr, key uint64) (r uint8) {
	var hash uint32_t
	var index int32
	_, _ = hash, index
	// key of zero is a sentinel
	if !(key != uint64FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+15088, __ccgo_ts+15042, int32FromInt32(140)) != 0 {
		__builtin_trap(tls)
	}
	hash = b2KeyHash(tls, key)
	index = b2FindSlot(tls, set, key, hash)
	return boolUint8((*(*b2SetItem)(unsafe.Add(nil, (*b2HashSet)(unsafe.Pointer(set)).Items+uintptr(index)*16))).Key == key)
}

func b2GetHashSetBytes(tls *_Stack, set uintptr) (r int32) {
	return int32FromUint32((*b2HashSet)(unsafe.Pointer(set)).Capacity * uint32FromInt32(int32FromInt64(16)))
}

func b2AddKey(tls *_Stack, set uintptr, key uint64) (r uint8) {
	var hash uint32_t
	var index int32
	_, _ = hash, index
	// key of zero is a sentinel
	if !(key != uint64FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+15088, __ccgo_ts+15042, int32FromInt32(154)) != 0 {
		__builtin_trap(tls)
	}
	hash = b2KeyHash(tls, key)
	if !(hash != uint32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+15097, __ccgo_ts+15042, int32FromInt32(157)) != 0 {
		__builtin_trap(tls)
	}
	index = b2FindSlot(tls, set, key, hash)
	if (*(*b2SetItem)(unsafe.Add(nil, (*b2HashSet)(unsafe.Pointer(set)).Items+uintptr(index)*16))).Hash != uint32(0) {
		// Already in set
		if !((*(*b2SetItem)(unsafe.Add(nil, (*b2HashSet)(unsafe.Pointer(set)).Items+uintptr(index)*16))).Hash == hash && (*(*b2SetItem)(unsafe.Add(nil, (*b2HashSet)(unsafe.Pointer(set)).Items+uintptr(index)*16))).Key == key) && b2InternalAssertFcn(tls, __ccgo_ts+15107, __ccgo_ts+15042, int32FromInt32(163)) != 0 {
			__builtin_trap(tls)
		}
		return boolUint8(true1 != 0)
	}
	if uint32(2)*(*b2HashSet)(unsafe.Pointer(set)).Count >= (*b2HashSet)(unsafe.Pointer(set)).Capacity {
		b2GrowTable(tls, set)
	}
	b2AddKeyHaveCapacity(tls, set, key, hash)
	return boolUint8(false1 != 0)
}

// C documentation
//
//	// See https://en.wikipedia.org/wiki/Open_addressing
func b2RemoveKey(tls *_Stack, set uintptr, key uint64) (r uint8) {
	var capacity, hash uint32_t
	var i, j, k int32
	var items uintptr
	_, _, _, _, _, _ = capacity, hash, i, items, j, k
	hash = b2KeyHash(tls, key)
	i = b2FindSlot(tls, set, key, hash)
	items = (*b2HashSet)(unsafe.Pointer(set)).Items
	if (*(*b2SetItem)(unsafe.Add(nil, items+uintptr(i)*16))).Hash == uint32(0) {
		// Not in set
		return boolUint8(false1 != 0)
	}
	// Mark item i as unoccupied
	(*(*b2SetItem)(unsafe.Add(nil, items+uintptr(i)*16))).Key = uint64(0)
	(*(*b2SetItem)(unsafe.Add(nil, items+uintptr(i)*16))).Hash = uint32(0)
	if !((*b2HashSet)(unsafe.Pointer(set)).Count > uint32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+15170, __ccgo_ts+15042, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	*(*uint32_t)(unsafe.Add(nil, set+12)) -= uint32(1)
	// Attempt to fill item i
	j = i
	capacity = (*b2HashSet)(unsafe.Pointer(set)).Capacity
	for {
		j = int32FromUint32(uint32FromInt32(j+int32FromInt32(1)) & (capacity - uint32(1)))
		if (*(*b2SetItem)(unsafe.Add(nil, items+uintptr(j)*16))).Hash == uint32(0) {
			break
		}
		// k is the first item for the hash of j
		k = int32FromUint32((*(*b2SetItem)(unsafe.Add(nil, items+uintptr(j)*16))).Hash & (capacity - uint32(1)))
		// determine if k lies cyclically in (i,j]
		// i <= j: | i..k..j |
		// i > j: |.k..j  i....| or |....j     i..k.|
		if i <= j {
			if i < k && k <= j {
				goto _1
			}
		} else {
			if i < k || k <= j {
				goto _1
			}
		}
		// Move j into i
		*(*b2SetItem)(unsafe.Add(nil, items+uintptr(i)*16)) = *(*b2SetItem)(unsafe.Add(nil, items+uintptr(j)*16))
		// Mark item j as unoccupied
		(*(*b2SetItem)(unsafe.Add(nil, items+uintptr(j)*16))).Key = uint64(0)
		(*(*b2SetItem)(unsafe.Add(nil, items+uintptr(j)*16))).Hash = uint32(0)
		i = j
		goto _1
	_1:
	}
	return boolUint8(true1 != 0)
}

const _CLOCKS_PER_SEC = 1000000
const _CLOCK_BOOTTIME = 7
const _CLOCK_BOOTTIME_ALARM = 9
const _CLOCK_MONOTONIC = 1
const _CLOCK_MONOTONIC_COARSE = 6
const _CLOCK_MONOTONIC_RAW = 4
const _CLOCK_PROCESS_CPUTIME_ID = 2
const _CLOCK_REALTIME = 0
const _CLOCK_REALTIME_ALARM = 8
const _CLOCK_REALTIME_COARSE = 5
const _CLOCK_SGI_CYCLE = 10
const _CLOCK_TAI = 11
const _CLOCK_THREAD_CPUTIME_ID = 3
const _CLONE_CHILD_CLEARTID = 0x00200000
const _CLONE_CHILD_SETTID = 0x01000000
const _CLONE_DETACHED = 0x00400000
const _CLONE_FILES = 0x00000400
const _CLONE_FS = 0x00000200
const _CLONE_IO = 0x80000000
const _CLONE_NEWCGROUP = 0x02000000
const _CLONE_NEWIPC = 0x08000000
const _CLONE_NEWNET = 0x40000000
const _CLONE_NEWNS = 0x00020000
const _CLONE_NEWPID = 0x20000000
const _CLONE_NEWTIME = 0x00000080
const _CLONE_NEWUSER = 0x10000000
const _CLONE_NEWUTS = 0x04000000
const _CLONE_PARENT = 0x00008000
const _CLONE_PARENT_SETTID = 0x00100000
const _CLONE_PIDFD = 0x00001000
const _CLONE_PTRACE = 0x00002000
const _CLONE_SETTLS = 0x00080000
const _CLONE_SIGHAND = 0x00000800
const _CLONE_SYSVSEM = 0x00040000
const _CLONE_THREAD = 0x00010000
const _CLONE_UNTRACED = 0x00800000
const _CLONE_VFORK = 0x00004000
const _CLONE_VM = 0x00000100
const _CPU_SETSIZE = 1024
const _CSIGNAL = 0x000000ff
const _SCHED_BATCH = 3
const _SCHED_DEADLINE = 6
const _SCHED_FIFO = 1
const _SCHED_IDLE = 5
const _SCHED_OTHER = 0
const _SCHED_RESET_ON_FORK = 0x40000000
const _SCHED_RR = 2
const _TIMER_ABSTIME = 1
const _TIME_UTC = 1
const __tm_gmtoff = "tm_gmtoff"
const __tm_zone = "tm_zone"

type time_t = int64

type timespec struct {
	Tv_sec  time_t
	Tv_nsec int64
}

type pid_t = int32

type sched_param struct {
	Sched_priority int32
	ｆ__reserved1   int32
	ｆ__reserved2   [2]struct {
		ｆ__reserved1 time_t
		ｆ__reserved2 int64
	}
	ｆ__reserved3 int32
}

type cpu_set_t struct {
	ｆ__bits [16]uint64
}

type timer_t = uintptr

type clockid_t = int32

type clock_t = int64

type tm struct {
	Tm_sec    int32
	Tm_min    int32
	Tm_hour   int32
	Tm_mday   int32
	Tm_mon    int32
	Tm_year   int32
	Tm_wday   int32
	Tm_yday   int32
	Tm_isdst  int32
	Tm_gmtoff int64
	Tm_zone   uintptr
}

type itimerspec struct {
	It_interval timespec
	It_value    timespec
}

func b2GetTicks(tls *_Stack) (r uint64) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* ts at bp+0 */ timespec
	clock_gettime(tls, int32(_CLOCK_MONOTONIC), bp)
	return uint64FromInt64((*(*timespec)(unsafe.Add(nil, bp))).Tv_sec*int64(1000000000) + (*(*timespec)(unsafe.Add(nil, bp))).Tv_nsec)
}

func b2GetMilliseconds(tls *_Stack, ticks uint64) (r float32) {
	var ticksNow uint64_t
	_ = ticksNow
	ticksNow = b2GetTicks(tls)
	return float32(float64(ticksNow-ticks) / float64FromFloat64(1e+06))
}

func b2GetMillisecondsAndReset(tls *_Stack, ticks uintptr) (r float32) {
	var ms float32
	var ticksNow uint64_t
	_, _ = ms, ticksNow
	ticksNow = b2GetTicks(tls)
	ms = float32(float64(ticksNow-*(*uint64_t)(unsafe.Add(nil, ticks))) / float64FromFloat64(1e+06))
	*(*uint64_t)(unsafe.Add(nil, ticks)) = ticksNow
	return ms
}

func b2Yield(tls *_Stack) {
	sched_yield(tls)
}

// C documentation
//
//	// djb2 hash
//	// https://en.wikipedia.org/wiki/List_of_hash_functions
func b2Hash(tls *_Stack, hash uint32, data uintptr, count int32) (r uint32) {
	var i int32
	var result uint32_t
	_, _ = i, result
	result = hash
	i = 0
	for {
		if !(i < count) {
			break
		}
		result = result<<int32FromInt32(5) + result + uint32(*(*uint8_t)(unsafe.Add(nil, data+uintptr(i))))
		goto _1
	_1:
		;
		i = i + 1
	}
	return result
}

const _UINT64_MAX9 = 18446744073709551615

func b2DefaultFilter(tls *_Stack) (r Filter) {
	var filter Filter
	_ = filter
	filter = Filter{
		CategoryBits: uint64(_B2_DEFAULT_CATEGORY_BITS),
		MaskBits:     uint64FromUint64(0xffffffffffffffff),
	}
	return filter
}

func b2DefaultQueryFilter(tls *_Stack) (r QueryFilter) {
	var filter QueryFilter
	_ = filter
	filter = QueryFilter{
		CategoryBits: uint64(_B2_DEFAULT_CATEGORY_BITS),
		MaskBits:     uint64FromUint64(0xffffffffffffffff),
	}
	return filter
}

func b2DefaultSurfaceMaterial(tls *_Stack) (r SurfaceMaterial) {
	var material SurfaceMaterial
	_ = material
	material = SurfaceMaterial{
		Friction: float32FromFloat32(0.6),
	}
	return material
}

func b2DefaultChainDef(tls *_Stack) (r ChainDef) {
	var def ChainDef
	_ = def
	def = ChainDef{}
	def.Materials = uintptr(unsafe.Pointer(&defaultMaterial))
	def.MaterialCount = int32(1)
	def.Filter = b2DefaultFilter(tls)
	def.InternalValue = int32(_B2_SECRET_COOKIE)
	return def
}

var defaultMaterial = SurfaceMaterial{
	Friction: float32FromFloat32(0.6),
}

func b2EmptyDrawPolygon(tls *_Stack, vertices uintptr, vertexCount int32, color HexColor, context uintptr) {
	_ = uint64FromInt64(4)
}

func b2EmptyDrawSolidPolygon(tls *_Stack, transform Transform, vertices uintptr, vertexCount int32, radius float32, color HexColor, context uintptr) {
	_ = uint64FromInt64(4)
}

func b2EmptyDrawCircle(tls *_Stack, center Vec2, radius float32, color HexColor, context uintptr) {
	_ = uint64FromInt64(4)
}

func b2EmptyDrawSolidCircle(tls *_Stack, transform Transform, radius float32, color HexColor, context uintptr) {
	_ = uint64FromInt64(4)
}

func b2EmptyDrawSolidCapsule(tls *_Stack, p1 Vec2, p2 Vec2, radius float32, color HexColor, context uintptr) {
	_ = uint64FromInt64(4)
}

func b2EmptyDrawSegment(tls *_Stack, p1 Vec2, p2 Vec2, color HexColor, context uintptr) {
	_ = uint64FromInt64(4)
}

func b2EmptyDrawTransform(tls *_Stack, transform Transform, context uintptr) {
	_ = uint64FromInt64(4)
}

func b2EmptyDrawPoint(tls *_Stack, p Vec2, size float32, color HexColor, context uintptr) {
	_ = uint64FromInt64(4)
}

func b2EmptyDrawString(tls *_Stack, p Vec2, s uintptr, color HexColor, context uintptr) {
	_ = uint64FromInt64(4)
}

func b2DefaultDebugDraw(tls *_Stack) (r b2DebugDraw) {
	var draw b2DebugDraw
	_ = draw
	draw = b2DebugDraw{}
	// These allow the user to skip some implementations and not hit null exceptions.
	draw.DrawPolygonFcn = __ccgo_fp(b2EmptyDrawPolygon)
	draw.DrawSolidPolygonFcn = __ccgo_fp(b2EmptyDrawSolidPolygon)
	draw.DrawCircleFcn = __ccgo_fp(b2EmptyDrawCircle)
	draw.DrawSolidCircleFcn = __ccgo_fp(b2EmptyDrawSolidCircle)
	draw.DrawSolidCapsuleFcn = __ccgo_fp(b2EmptyDrawSolidCapsule)
	draw.DrawSegmentFcn = __ccgo_fp(b2EmptyDrawSegment)
	draw.DrawTransformFcn = __ccgo_fp(b2EmptyDrawTransform)
	draw.DrawPointFcn = __ccgo_fp(b2EmptyDrawPoint)
	draw.DrawStringFcn = __ccgo_fp(b2EmptyDrawString)
	return draw
}

const _UINT64_MAX10 = "0xffffffffffffffffu"

var b2_identityBodyState16 = b2BodyState{
	DeltaRotation: Rot{
		C: float32FromFloat32(1),
	},
}

var b2_identityBodyState17 = b2BodyState{
	DeltaRotation: Rot{
		C: float32FromFloat32(1),
	},
}

const _FLT_MAX8 = 3.4028234663852886e+38
const _UINT16_MAX1 = 65535
const _UINT64_MAX11 = 18446744073709551615

var b2_identityBodyState18 = b2BodyState{
	DeltaRotation: Rot{
		C: float32FromFloat32(1),
	},
}

func b2ContactBeginTouchEventArray_Create(tls *_Stack, capacity int32) (r b2ContactBeginTouchEventArray) {
	var a b2ContactBeginTouchEventArray
	_ = a
	a = b2ContactBeginTouchEventArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(128)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2ContactBeginTouchEventArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2ContactBeginTouchEventArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2ContactBeginTouchEventArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2ContactBeginTouchEventArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2ContactBeginTouchEventArray)(unsafe.Pointer(a)).Capacity)*uint64(128)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(128)))
	(*b2ContactBeginTouchEventArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2ContactBeginTouchEventArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2ContactBeginTouchEventArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2ContactBeginTouchEventArray)(unsafe.Pointer(a)).Capacity)*uint64(128)))
	(*b2ContactBeginTouchEventArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2ContactBeginTouchEventArray)(unsafe.Pointer(a)).Count = 0
	(*b2ContactBeginTouchEventArray)(unsafe.Pointer(a)).Capacity = 0
}

func b2ContactEndTouchEventArray_Create(tls *_Stack, capacity int32) (r b2ContactEndTouchEventArray) {
	var a b2ContactEndTouchEventArray
	_ = a
	a = b2ContactEndTouchEventArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(16)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2ContactEndTouchEventArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2ContactEndTouchEventArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2ContactEndTouchEventArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2ContactEndTouchEventArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2ContactEndTouchEventArray)(unsafe.Pointer(a)).Capacity)*uint64(16)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(16)))
	(*b2ContactEndTouchEventArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2ContactEndTouchEventArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2ContactEndTouchEventArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2ContactEndTouchEventArray)(unsafe.Pointer(a)).Capacity)*uint64(16)))
	(*b2ContactEndTouchEventArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2ContactEndTouchEventArray)(unsafe.Pointer(a)).Count = 0
	(*b2ContactEndTouchEventArray)(unsafe.Pointer(a)).Capacity = 0
}

func b2ContactHitEventArray_Create(tls *_Stack, capacity int32) (r b2ContactHitEventArray) {
	var a b2ContactHitEventArray
	_ = a
	a = b2ContactHitEventArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(36)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2ContactHitEventArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2ContactHitEventArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2ContactHitEventArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2ContactHitEventArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2ContactHitEventArray)(unsafe.Pointer(a)).Capacity)*uint64(36)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(36)))
	(*b2ContactHitEventArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2ContactHitEventArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2ContactHitEventArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2ContactHitEventArray)(unsafe.Pointer(a)).Capacity)*uint64(36)))
	(*b2ContactHitEventArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2ContactHitEventArray)(unsafe.Pointer(a)).Count = 0
	(*b2ContactHitEventArray)(unsafe.Pointer(a)).Capacity = 0
}

func b2TaskContextArray_Create(tls *_Stack, capacity int32) (r b2TaskContextArray) {
	var a b2TaskContextArray
	_ = a
	a = b2TaskContextArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(56)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2TaskContextArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2TaskContextArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2TaskContextArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2TaskContextArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2TaskContextArray)(unsafe.Pointer(a)).Capacity)*uint64(56)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(56)))
	(*b2TaskContextArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2TaskContextArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2TaskContextArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2TaskContextArray)(unsafe.Pointer(a)).Capacity)*uint64(56)))
	(*b2TaskContextArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2TaskContextArray)(unsafe.Pointer(a)).Count = 0
	(*b2TaskContextArray)(unsafe.Pointer(a)).Capacity = 0
}

func b2DefaultAddTaskFcn(tls *_Stack, __ccgo_fp_task uintptr, count int32, minRange int32, taskContext uintptr, userContext uintptr) (r uintptr) {
	_ = uint64FromInt64(4)
	(*(*func(*_Stack, int32, int32, uint32, uintptr))(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_task})))(tls, 0, count, uint32(0), taskContext)
	return uintptrFromInt32(0)
}

func b2DefaultFinishTaskFcn(tls *_Stack, userTask uintptr, userContext uintptr) {
	_ = uint64FromInt64(4)
}

func b2DefaultFrictionCallback(tls *_Stack, frictionA float32, materialA int32, frictionB float32, materialB int32) (r float32) {
	_ = uint64FromInt64(4)
	return sqrtf(tls, float32(frictionA*frictionB))
}

func b2DefaultRestitutionCallback(tls *_Stack, restitutionA float32, materialA int32, restitutionB float32, materialB int32) (r float32) {
	var v1, v2, v3, v5 float32
	_, _, _, _ = v1, v2, v3, v5
	_ = uint64FromInt64(4)
	v1 = restitutionA
	v2 = restitutionB
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _4
_4:
	return v3
}

func b2UpdateTreesTask(tls *_Stack, startIndex int32, endIndex int32, threadIndex uint32, context uintptr) {
	var world uintptr
	_ = world
	_ = uint64FromInt64(4)
	_ = uint64FromInt64(4)
	_ = uint64FromInt64(4)
	world = context
	b2BroadPhase_RebuildTrees(tls, world+40)
}

func b2AddNonTouchingContact(tls *_Stack, world uintptr, contact uintptr, contactSim uintptr) {
	var newCapacity, v2 int32
	var newContactSim, set, v1, v3 uintptr
	_, _, _, _, _, _ = newCapacity, newContactSim, set, v1, v2, v3
	if !((*b2Contact)(unsafe.Pointer(contact)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+3637, __ccgo_ts+15342, int32FromInt32(463)) != 0 {
		__builtin_trap(tls)
	}
	v1 = world + 1064
	v2 = int32(b2_awakeSet)
	if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*88
	goto _4
_4:
	set = v3
	(*b2Contact)(unsafe.Pointer(contact)).ColorIndex = -int32(1)
	(*b2Contact)(unsafe.Pointer(contact)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Count
	v1 = set + 48
	if (*b2ContactSimArray)(unsafe.Pointer(v1)).Count == (*b2ContactSimArray)(unsafe.Pointer(v1)).Capacity {
		if (*b2ContactSimArray)(unsafe.Pointer(v1)).Capacity < int32(2) {
			v2 = int32(2)
		} else {
			v2 = (*b2ContactSimArray)(unsafe.Pointer(v1)).Capacity + (*b2ContactSimArray)(unsafe.Pointer(v1)).Capacity>>int32(1)
		}
		newCapacity = v2
		b2ContactSimArray_Reserve(tls, v1, newCapacity)
	}
	*(*int32)(unsafe.Add(nil, v1+8)) += int32(1)
	v3 = (*b2ContactSimArray)(unsafe.Pointer(v1)).Data + uintptr((*b2ContactSimArray)(unsafe.Pointer(v1)).Count-int32FromInt32(1))*176
	goto _8
_8:
	newContactSim = v3
	memcpy(tls, newContactSim, contactSim, uint64(176))
}

func b2RemoveNonTouchingContact(tls *_Stack, world uintptr, setIndex int32, localIndex int32) {
	var movedContact, movedContactSim, set, v1, v11, v3, v5, v9 uintptr
	var movedIndex, movedIndex1, v10, v2, v6, v7 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = movedContact, movedContactSim, movedIndex, movedIndex1, set, v1, v10, v11, v2, v3, v5, v6, v7, v9
	v1 = world + 1064
	v2 = setIndex
	if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*88
	goto _4
_4:
	set = v3
	v5 = set + 48
	v6 = localIndex
	if !(0 <= v6 && v6 < (*b2ContactSimArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(147)) != 0 {
		__builtin_trap(tls)
	}
	movedIndex = -int32(1)
	if v6 != (*b2ContactSimArray)(unsafe.Pointer(v5)).Count-int32FromInt32(1) {
		movedIndex = (*b2ContactSimArray)(unsafe.Pointer(v5)).Count - int32(1)
		*(*b2ContactSim)(unsafe.Add(nil, (*b2ContactSimArray)(unsafe.Pointer(v5)).Data+uintptr(v6)*176)) = *(*b2ContactSim)(unsafe.Add(nil, (*b2ContactSimArray)(unsafe.Pointer(v5)).Data+uintptr(movedIndex)*176))
	}
	*(*int32)(unsafe.Add(nil, v5+8)) -= int32(1)
	v7 = movedIndex
	goto _8
_8:
	movedIndex1 = v7
	if movedIndex1 != -int32(1) {
		movedContactSim = (*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Data + uintptr(localIndex)*176
		v9 = world + 1144
		v10 = (*b2ContactSim)(unsafe.Pointer(movedContactSim)).ContactId
		if !(0 <= v10 && v10 < (*b2ContactArray)(unsafe.Pointer(v9)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
			__builtin_trap(tls)
		}
		v11 = (*b2ContactArray)(unsafe.Pointer(v9)).Data + uintptr(v10)*68
		goto _12
	_12:
		movedContact = v11
		if !((*b2Contact)(unsafe.Pointer(movedContact)).SetIndex == setIndex) && b2InternalAssertFcn(tls, __ccgo_ts+15733, __ccgo_ts+15342, int32FromInt32(480)) != 0 {
			__builtin_trap(tls)
		}
		if !((*b2Contact)(unsafe.Pointer(movedContact)).LocalIndex == movedIndex1) && b2InternalAssertFcn(tls, __ccgo_ts+3168, __ccgo_ts+15342, int32FromInt32(481)) != 0 {
			__builtin_trap(tls)
		}
		if !((*b2Contact)(unsafe.Pointer(movedContact)).ColorIndex == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+15768, __ccgo_ts+15342, int32FromInt32(482)) != 0 {
			__builtin_trap(tls)
		}
		(*b2Contact)(unsafe.Pointer(movedContact)).LocalIndex = localIndex
	}
}

type DrawContext struct {
	World uintptr
	Draw  uintptr
}

func DrawQueryCallback(tls *_Stack, proxyId int32, userData uint64, context uintptr) (r uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var aabb AABB
	var blockIndex, v6 uint32_t
	var body, bodySim, draw, drawContext, shape, world, v1, v3 uintptr
	var color b2HexColor1
	var shapeId, v2 int32
	var _ /* vs at bp+0 */ [4]Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, blockIndex, body, bodySim, color, draw, drawContext, shape, shapeId, world, v1, v2, v3, v6
	_ = uint64FromInt64(4)
	shapeId = int32FromUint64(userData)
	drawContext = context
	world = (*DrawContext)(unsafe.Pointer(drawContext)).World
	draw = (*DrawContext)(unsafe.Pointer(drawContext)).Draw
	v1 = world + 1248
	v2 = shapeId
	if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
	goto _4
_4:
	shape = v3
	if !((*b2Shape)(unsafe.Pointer(shape)).Id == shapeId) && b2InternalAssertFcn(tls, __ccgo_ts+16034, __ccgo_ts+15342, int32FromInt32(884)) != 0 {
		__builtin_trap(tls)
	}
	v1 = world + 1464
	v6 = uint32FromInt32((*b2Shape)(unsafe.Pointer(shape)).BodyId)
	blockIndex = v6 / uint32(64)
	if !(blockIndex < (*b2BitSet)(unsafe.Pointer(v1)).BlockCount) && b2InternalAssertFcn(tls, __ccgo_ts+10859, __ccgo_ts+10891, int32FromInt32(28)) != 0 {
		__builtin_trap(tls)
	}
	*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v1)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v6 % uint32FromInt32(64))
	if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawShapes != 0 {
		v1 = world + 1024
		v2 = (*b2Shape)(unsafe.Pointer(shape)).BodyId
		if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
		goto _10
	_10:
		body = v3
		bodySim = b2GetBodySim(tls, world, body)
		if (*b2Shape)(unsafe.Pointer(shape)).CustomColor != uint32(0) {
			color = int32FromUint32((*b2Shape)(unsafe.Pointer(shape)).CustomColor)
		} else {
			if (*b2Body)(unsafe.Pointer(body)).Type1 == int32(b2_dynamicBody) && (*b2Body)(unsafe.Pointer(body)).Mass == float32FromFloat32(0) {
				// Bad body
				color = int32(b2_colorRed)
			} else {
				if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_disabledSet) {
					color = int32(b2_colorSlateGray)
				} else {
					if (*b2Shape)(unsafe.Pointer(shape)).SensorIndex != -int32(1) {
						color = int32(b2_colorWheat)
					} else {
						if (*b2BodySim)(unsafe.Pointer(bodySim)).IsBullet != 0 && (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
							color = int32(b2_colorTurquoise)
						} else {
							if (*b2Body)(unsafe.Pointer(body)).IsSpeedCapped != 0 {
								color = int32(b2_colorYellow)
							} else {
								if (*b2BodySim)(unsafe.Pointer(bodySim)).IsFast != 0 {
									color = int32(b2_colorSalmon)
								} else {
									if (*b2Body)(unsafe.Pointer(body)).Type1 == int32(b2_staticBody) {
										color = int32(b2_colorPaleGreen)
									} else {
										if (*b2Body)(unsafe.Pointer(body)).Type1 == int32(b2_kinematicBody) {
											color = int32(b2_colorRoyalBlue)
										} else {
											if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
												color = int32(b2_colorPink)
											} else {
												color = int32(b2_colorGray)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
		b2DrawShape(tls, draw, shape, (*b2BodySim)(unsafe.Pointer(bodySim)).Transform, color)
	}
	if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawBounds != 0 {
		aabb = (*b2Shape)(unsafe.Pointer(shape)).FatAABB
		*(*[4]Vec2)(unsafe.Add(nil, bp)) = [4]Vec2{
			0: {
				X: aabb.LowerBound.X,
				Y: aabb.LowerBound.Y,
			},
			1: {
				X: aabb.UpperBound.X,
				Y: aabb.LowerBound.Y,
			},
			2: {
				X: aabb.UpperBound.X,
				Y: aabb.UpperBound.Y,
			},
			3: {
				X: aabb.LowerBound.X,
				Y: aabb.UpperBound.Y,
			},
		}
		(*(*func(*_Stack, uintptr, int32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPolygonFcn})))(tls, bp, int32(4), int32(b2_colorGold), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	}
	return boolUint8(true1 != 0)
}

// C documentation
//
//	// todo this has varying order for moving shapes, causing flicker when overlapping shapes are moving
//	// solution: display order by shape id modulus 3, keep 3 buckets in GLSolid* and flush in 3 passes.
func b2DrawWithBounds(tls *_Stack, world uintptr, draw uintptr) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var addColor, frictionColor, impulseColor, normalColor, persistColor, speculativeColor b2HexColor1
	var bits, body, bodySim, bodySim1, contact, contactSim, gc, joint, point, v11, v13, v23 uintptr
	var blockIndex, blockIndex1, bodyId, ctz, k, wordCount, v9 uint32_t
	var bodyCapacity, contactCapacity, contactId, contactKey, edgeIndex, edgeIndex1, i, j, jointCapacity, jointId, jointKey, pointCount, v1 int32
	var graphColors [12]b2HexColor1
	var k_axisScale, k_impulseScale, linearSlop, pointSize, x, y, v46 float32
	var normal, offset, offset1, p1, p11, p12, p13, p2, p21, p22, p23, tangent, v16, v17, v20, v21, v47 Vec2
	var transform, transform1, v15 Transform
	var word uint64_t
	var v29 uint8
	var _ /* buffer at bp+16 */ [32]uint8
	var _ /* buffer at bp+48 */ [32]uint8
	var _ /* drawContext at bp+0 */ DrawContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = addColor, bits, blockIndex, blockIndex1, body, bodyCapacity, bodyId, bodySim, bodySim1, contact, contactCapacity, contactId, contactKey, contactSim, ctz, edgeIndex, edgeIndex1, frictionColor, gc, graphColors, i, impulseColor, j, joint, jointCapacity, jointId, jointKey, k, k_axisScale, k_impulseScale, linearSlop, normal, normalColor, offset, offset1, p1, p11, p12, p13, p2, p21, p22, p23, persistColor, point, pointCount, pointSize, speculativeColor, tangent, transform, transform1, word, wordCount, x, y, v1, v11, v13, v15, v16, v17, v20, v21, v23, v29, v46, v47, v9
	if !(b2IsValidAABB(tls, (*b2DebugDraw)(unsafe.Pointer(draw)).DrawingBounds) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+16055, __ccgo_ts+15342, int32FromInt32(963)) != 0 {
		__builtin_trap(tls)
	}
	k_impulseScale = float32FromFloat32(1)
	k_axisScale = float32FromFloat32(0.3)
	speculativeColor = int32(b2_colorGainsboro)
	addColor = int32(b2_colorGreen)
	persistColor = int32(b2_colorBlue)
	normalColor = int32(b2_colorDimGray)
	impulseColor = int32(b2_colorMagenta)
	frictionColor = int32(b2_colorYellow)
	graphColors = [12]b2HexColor1{
		0:  int32(b2_colorRed),
		1:  int32(b2_colorOrange),
		2:  int32(b2_colorYellow),
		3:  int32(b2_colorGreen),
		4:  int32(b2_colorCyan),
		5:  int32(b2_colorBlue),
		6:  int32(b2_colorViolet),
		7:  int32(b2_colorPink),
		8:  int32(b2_colorChocolate),
		9:  int32(b2_colorGoldenRod),
		10: int32(b2_colorCoral),
	}
	v1 = (*b2IdPool)(unsafe.Pointer(world + 1000)).NextIndex
	goto _2
_2:
	bodyCapacity = v1
	b2SetBitCountAndClear(tls, world+1464, uint32FromInt32(bodyCapacity))
	v1 = (*b2IdPool)(unsafe.Pointer(world + 1080)).NextIndex
	goto _4
_4:
	jointCapacity = v1
	b2SetBitCountAndClear(tls, world+1480, uint32FromInt32(jointCapacity))
	v1 = (*b2IdPool)(unsafe.Pointer(world + 1120)).NextIndex
	goto _6
_6:
	contactCapacity = v1
	b2SetBitCountAndClear(tls, world+1496, uint32FromInt32(contactCapacity))
	*(*DrawContext)(unsafe.Add(nil, bp)) = DrawContext{
		World: world,
		Draw:  draw,
	}
	i = 0
	for {
		if !(i < int32(b2_bodyTypeCount)) {
			break
		}
		b2DynamicTree_Query(tls, world+40+uintptr(i)*72, (*b2DebugDraw)(unsafe.Pointer(draw)).DrawingBounds, uint64FromUint64(0xffffffffffffffff), __ccgo_fp(DrawQueryCallback), bp)
		goto _7
	_7:
		;
		i = i + 1
	}
	wordCount = (*b2World)(unsafe.Pointer(world)).DebugBodySet.BlockCount
	bits = (*b2World)(unsafe.Pointer(world)).DebugBodySet.Bits
	k = uint32(0)
	for {
		if !(k < wordCount) {
			break
		}
		word = *(*uint64_t)(unsafe.Add(nil, bits+uintptr(k)*8))
		for word != uint64(0) {
			v9 = uint32FromInt32(__builtin_ctzll(tls, word))
			goto _10
		_10:
			ctz = v9
			bodyId = uint32(64)*k + ctz
			v11 = world + 1024
			v1 = int32FromUint32(bodyId)
			if !(0 <= v1 && v1 < (*b2BodyArray)(unsafe.Pointer(v11)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
				__builtin_trap(tls)
			}
			v13 = (*b2BodyArray)(unsafe.Pointer(v11)).Data + uintptr(v1)*128
			goto _14
		_14:
			body = v13
			if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawBodyNames != 0 && int32FromUint8(*(*uint8)(unsafe.Add(nil, body))) != 0 {
				offset = Vec2{
					X: float32FromFloat32(0.1),
					Y: float32FromFloat32(0.1),
				}
				bodySim = b2GetBodySim(tls, world, body)
				transform = Transform{
					P: (*b2BodySim)(unsafe.Pointer(bodySim)).Center,
					Q: (*b2BodySim)(unsafe.Pointer(bodySim)).Transform.Q,
				}
				v15 = transform
				v16 = offset
				x = float32(v15.Q.C*v16.X) - float32(v15.Q.S*v16.Y) + v15.P.X
				y = float32(v15.Q.S*v16.X) + float32(v15.Q.C*v16.Y) + v15.P.Y
				v17 = Vec2{
					X: x,
					Y: y,
				}
				goto _18
			_18:
				p1 = v17
				(*(*func(*_Stack, Vec2, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, p1, body, int32(b2_colorBlueViolet), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
			}
			if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawMass != 0 && (*b2Body)(unsafe.Pointer(body)).Type1 == int32(b2_dynamicBody) {
				offset1 = Vec2{
					X: float32FromFloat32(0.1),
					Y: float32FromFloat32(0.1),
				}
				bodySim1 = b2GetBodySim(tls, world, body)
				transform1 = Transform{
					P: (*b2BodySim)(unsafe.Pointer(bodySim1)).Center,
					Q: (*b2BodySim)(unsafe.Pointer(bodySim1)).Transform.Q,
				}
				(*(*func(*_Stack, Transform, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawTransformFcn})))(tls, transform1, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
				v15 = transform1
				v16 = offset1
				x = float32(v15.Q.C*v16.X) - float32(v15.Q.S*v16.Y) + v15.P.X
				y = float32(v15.Q.S*v16.X) + float32(v15.Q.C*v16.Y) + v15.P.Y
				v17 = Vec2{
					X: x,
					Y: y,
				}
				goto _22
			_22:
				p2 = v17
				__builtin_snprintf(tls, bp+16, uint64(32), __ccgo_ts+16092, vaList(bp+88, float64((*b2Body)(unsafe.Pointer(body)).Mass)))
				(*(*func(*_Stack, Vec2, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, p2, bp+16, int32(b2_colorWhite), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
			}
			if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawJoints != 0 {
				jointKey = (*b2Body)(unsafe.Pointer(body)).HeadJointKey
				for jointKey != -int32(1) {
					jointId = jointKey >> int32(1)
					edgeIndex = jointKey & int32(1)
					v11 = world + 1104
					v1 = jointId
					if !(0 <= v1 && v1 < (*b2JointArray)(unsafe.Pointer(v11)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
						__builtin_trap(tls)
					}
					v13 = (*b2JointArray)(unsafe.Pointer(v11)).Data + uintptr(v1)*72
					goto _26
				_26:
					joint = v13
					// avoid double draw
					v23 = world + 1480
					v9 = uint32FromInt32(jointId)
					blockIndex1 = v9 / uint32(64)
					if blockIndex1 >= (*b2BitSet)(unsafe.Pointer(v23)).BlockCount {
						v29 = boolUint8(false1 != 0)
						goto _30
					}
					v29 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v23)).Bits+uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v9%uint32FromInt32(64))) != uint64(0))
					goto _30
				_30:
					if int32FromUint8(v29) == false1 {
						b2DrawJoint(tls, draw, world, joint)
						v11 = world + 1480
						v9 = uint32FromInt32(jointId)
						blockIndex = v9 / uint32(64)
						if !(blockIndex < (*b2BitSet)(unsafe.Pointer(v11)).BlockCount) && b2InternalAssertFcn(tls, __ccgo_ts+10859, __ccgo_ts+10891, int32FromInt32(28)) != 0 {
							__builtin_trap(tls)
						}
						*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v11)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v9 % uint32FromInt32(64))
					} else {
						// todo testing
						edgeIndex = edgeIndex + 0
					}
					jointKey = (*(*b2JointEdge)(unsafe.Add(nil, joint+20+uintptr(edgeIndex)*12))).NextKey
				}
			}
			linearSlop = float32(float32FromFloat32(0.005) * b2_lengthUnitsPerMeter)
			if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawContacts != 0 && (*b2Body)(unsafe.Pointer(body)).Type1 == int32(b2_dynamicBody) && (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
				contactKey = (*b2Body)(unsafe.Pointer(body)).HeadContactKey
				for contactKey != -int32(1) {
					contactId = contactKey >> int32(1)
					edgeIndex1 = contactKey & int32(1)
					v11 = world + 1144
					v1 = contactId
					if !(0 <= v1 && v1 < (*b2ContactArray)(unsafe.Pointer(v11)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
						__builtin_trap(tls)
					}
					v13 = (*b2ContactArray)(unsafe.Pointer(v11)).Data + uintptr(v1)*68
					goto _36
				_36:
					contact = v13
					contactKey = (*(*b2ContactEdge)(unsafe.Add(nil, contact+12+uintptr(edgeIndex1)*12))).NextKey
					if (*b2Contact)(unsafe.Pointer(contact)).SetIndex != int32(b2_awakeSet) || (*b2Contact)(unsafe.Pointer(contact)).ColorIndex == -int32(1) {
						continue
					}
					// avoid double draw
					v11 = world + 1496
					v9 = uint32FromInt32(contactId)
					blockIndex1 = v9 / uint32(64)
					if blockIndex1 >= (*b2BitSet)(unsafe.Pointer(v11)).BlockCount {
						v29 = boolUint8(false1 != 0)
						goto _40
					}
					v29 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v11)).Bits+uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v9%uint32FromInt32(64))) != uint64(0))
					goto _40
				_40:
					if int32FromUint8(v29) == false1 {
						if !(0 <= (*b2Contact)(unsafe.Pointer(contact)).ColorIndex && (*b2Contact)(unsafe.Pointer(contact)).ColorIndex < int32(_B2_GRAPH_COLOR_COUNT)) && b2InternalAssertFcn(tls, __ccgo_ts+3755, __ccgo_ts+15342, int32FromInt32(1076)) != 0 {
							__builtin_trap(tls)
						}
						gc = world + 328 + uintptr((*b2Contact)(unsafe.Pointer(contact)).ColorIndex)*56
						v11 = gc + 16
						v1 = (*b2Contact)(unsafe.Pointer(contact)).LocalIndex
						if !(0 <= v1 && v1 < (*b2ContactSimArray)(unsafe.Pointer(v11)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(147)) != 0 {
							__builtin_trap(tls)
						}
						v13 = (*b2ContactSimArray)(unsafe.Pointer(v11)).Data + uintptr(v1)*176
						goto _44
					_44:
						contactSim = v13
						pointCount = (*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.PointCount
						normal = (*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.Normal
						j = 0
						for {
							if !(j < pointCount) {
								break
							}
							point = contactSim + 36 + 12 + uintptr(j)*48
							if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawGraphColors != 0 {
								if (*b2Contact)(unsafe.Pointer(contact)).ColorIndex == int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1) {
									v46 = float32FromFloat32(7.5)
								} else {
									v46 = float32FromFloat32(5)
								}
								// graph color
								pointSize = v46
								(*(*func(*_Stack, Vec2, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, (*ManifoldPoint)(unsafe.Pointer(point)).Point, pointSize, graphColors[(*b2Contact)(unsafe.Pointer(contact)).ColorIndex], (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
								// m_context->draw.DrawString(point->position, "%d", point->color);
							} else {
								if (*ManifoldPoint)(unsafe.Pointer(point)).Separation > linearSlop {
									// Speculative
									(*(*func(*_Stack, Vec2, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, (*ManifoldPoint)(unsafe.Pointer(point)).Point, float32FromFloat32(5), speculativeColor, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
								} else {
									if int32FromUint8((*ManifoldPoint)(unsafe.Pointer(point)).Persisted) == false1 {
										// Add
										(*(*func(*_Stack, Vec2, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, (*ManifoldPoint)(unsafe.Pointer(point)).Point, float32FromFloat32(10), addColor, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
									} else {
										if int32FromUint8((*ManifoldPoint)(unsafe.Pointer(point)).Persisted) == int32(true1) {
											// Persist
											(*(*func(*_Stack, Vec2, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, (*ManifoldPoint)(unsafe.Pointer(point)).Point, float32FromFloat32(5), persistColor, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
										}
									}
								}
							}
							if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawContactNormals != 0 {
								p11 = (*ManifoldPoint)(unsafe.Pointer(point)).Point
								v16 = p11
								v46 = k_axisScale
								v17 = normal
								v20 = Vec2{
									X: v16.X + float32(v46*v17.X),
									Y: v16.Y + float32(v46*v17.Y),
								}
								goto _51
							_51:
								p21 = v20
								(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, p11, p21, normalColor, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
							} else {
								if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawContactImpulses != 0 {
									p12 = (*ManifoldPoint)(unsafe.Pointer(point)).Point
									v16 = p12
									v46 = float32(k_impulseScale * (*ManifoldPoint)(unsafe.Pointer(point)).NormalImpulse)
									v17 = normal
									v20 = Vec2{
										X: v16.X + float32(v46*v17.X),
										Y: v16.Y + float32(v46*v17.Y),
									}
									goto _56
								_56:
									p22 = v20
									(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, p12, p22, impulseColor, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
									__builtin_snprintf(tls, bp+48, uint64FromInt32(int32FromUint64(uint64FromInt64(32)/uint64FromInt64(1))), __ccgo_ts+16099, vaList(bp+88, float64(float32FromFloat32(1000)*(*ManifoldPoint)(unsafe.Pointer(point)).NormalImpulse)))
									(*(*func(*_Stack, Vec2, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, p12, bp+48, int32(b2_colorWhite), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
								}
							}
							if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawContactFeatures != 0 {
								__builtin_snprintf(tls, bp+48, uint64FromInt32(int32FromUint64(uint64FromInt64(32)/uint64FromInt64(1))), __ccgo_ts+16104, vaList(bp+88, int32FromUint16((*ManifoldPoint)(unsafe.Pointer(point)).Id)))
								(*(*func(*_Stack, Vec2, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, (*ManifoldPoint)(unsafe.Pointer(point)).Point, bp+48, int32(b2_colorOrange), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
							}
							if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawFrictionImpulses != 0 {
								v16 = normal
								v17 = Vec2{
									X: v16.Y,
									Y: -v16.X,
								}
								goto _59
							_59:
								tangent = v17
								p13 = (*ManifoldPoint)(unsafe.Pointer(point)).Point
								v20 = p13
								v46 = float32(k_impulseScale * (*ManifoldPoint)(unsafe.Pointer(point)).TangentImpulse)
								v21 = tangent
								v47 = Vec2{
									X: v20.X + float32(v46*v21.X),
									Y: v20.Y + float32(v46*v21.Y),
								}
								goto _64
							_64:
								p23 = v47
								(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, p13, p23, frictionColor, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
								__builtin_snprintf(tls, bp+48, uint64FromInt32(int32FromUint64(uint64FromInt64(32)/uint64FromInt64(1))), __ccgo_ts+16099, vaList(bp+88, float64(float32FromFloat32(1000)*(*ManifoldPoint)(unsafe.Pointer(point)).TangentImpulse)))
								(*(*func(*_Stack, Vec2, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, p13, bp+48, int32(b2_colorWhite), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
							}
							goto _45
						_45:
							;
							j = j + 1
						}
						v11 = world + 1496
						v9 = uint32FromInt32(contactId)
						blockIndex = v9 / uint32(64)
						if !(blockIndex < (*b2BitSet)(unsafe.Pointer(v11)).BlockCount) && b2InternalAssertFcn(tls, __ccgo_ts+10859, __ccgo_ts+10891, int32FromInt32(28)) != 0 {
							__builtin_trap(tls)
						}
						*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v11)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v9 % uint32FromInt32(64))
					} else {
						// todo testing
						edgeIndex1 = edgeIndex1 + 0
					}
					contactKey = (*(*b2ContactEdge)(unsafe.Add(nil, contact+12+uintptr(edgeIndex1)*12))).NextKey
				}
			}
			// Clear the smallest set bit
			word = word & (word - uint64(1))
		}
		goto _8
	_8:
		;
		k = k + 1
	}
}

func b2Chain_IsValid(tls *_Stack, id ChainId) (r uint8) {
	var chain, world uintptr
	var chainId int32
	_, _, _ = chain, chainId, world
	if int32(_B2_MAX_WORLDS) <= int32FromUint16(id.World0) {
		return boolUint8(false1 != 0)
	}
	world = uintptr(unsafe.Pointer(&b2_worlds)) + uintptr(id.World0)*1792
	if int32FromUint16((*b2World)(unsafe.Pointer(world)).WorldId) != int32FromUint16(id.World0) {
		// world is free
		return boolUint8(false1 != 0)
	}
	chainId = id.Index1 - int32(1)
	if chainId < 0 || (*b2World)(unsafe.Pointer(world)).ChainShapes.Count <= chainId {
		return boolUint8(false1 != 0)
	}
	chain = (*b2World)(unsafe.Pointer(world)).ChainShapes.Data + uintptr(chainId)*48
	if (*b2ChainShape)(unsafe.Pointer(chain)).Id == -int32(1) {
		// chain is free
		return boolUint8(false1 != 0)
	}
	if !((*b2ChainShape)(unsafe.Pointer(chain)).Id == chainId) && b2InternalAssertFcn(tls, __ccgo_ts+16146, __ccgo_ts+15342, int32FromInt32(1673)) != 0 {
		__builtin_trap(tls)
	}
	return boolUint8(int32FromUint16(id.Generation) == int32FromUint16((*b2ChainShape)(unsafe.Pointer(chain)).Generation))
}

type __ccgo_fp__Xb2World_SetFrictionCallback_1 = func(*_Stack, float32, int32, float32, int32) float32

type __ccgo_fp__Xb2World_SetRestitutionCallback_1 = func(*_Stack, float32, int32, float32, int32) float32

type WorldQueryContext struct {
	World       uintptr
	Fcn         uintptr
	Filter      QueryFilter
	UserContext uintptr
}

func TreeQueryCallback(tls *_Stack, proxyId int32, userData uint64, context uintptr) (r uint8) {
	var id ShapeId
	var result, v7 uint8
	var shape, world, worldContext, v1, v3 uintptr
	var shapeId, v2 int32
	var v5 Filter
	var v6 QueryFilter
	_, _, _, _, _, _, _, _, _, _, _, _ = id, result, shape, shapeId, world, worldContext, v1, v2, v3, v5, v6, v7
	_ = uint64FromInt64(4)
	shapeId = int32FromUint64(userData)
	worldContext = context
	world = (*WorldQueryContext)(unsafe.Pointer(worldContext)).World
	v1 = world + 1248
	v2 = shapeId
	if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
	goto _4
_4:
	shape = v3
	v5 = (*b2Shape)(unsafe.Pointer(shape)).Filter
	v6 = (*WorldQueryContext)(unsafe.Pointer(worldContext)).Filter
	v7 = boolUint8(v5.CategoryBits&v6.MaskBits != uint64(0) && v5.MaskBits&v6.CategoryBits != uint64(0))
	goto _8
_8:
	if int32FromUint8(v7) == false1 {
		return boolUint8(true1 != 0)
	}
	id = ShapeId{
		Index1:     shapeId + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2Shape)(unsafe.Pointer(shape)).Generation,
	}
	result = (*(*func(*_Stack, ShapeId, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*WorldQueryContext)(unsafe.Pointer(worldContext)).Fcn})))(tls, id, (*WorldQueryContext)(unsafe.Pointer(worldContext)).UserContext)
	return result
}

type __ccgo_fp__Xb2World_OverlapAABB_3 = func(*_Stack, struct {
	Index1     int32
	World0     uint16
	Generation uint16
}, uintptr) uint8

type WorldOverlapContext struct {
	World       uintptr
	Fcn         uintptr
	Filter      QueryFilter
	Proxy       uintptr
	UserContext uintptr
}

func TreeOverlapCallback(tls *_Stack, proxyId int32, userData uint64, context uintptr) (r uint8) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var body, shape, world, worldContext, v1, v3 uintptr
	var id ShapeId
	var output DistanceOutput
	var result, v7 uint8
	var shapeId, v2 int32
	var tolerance float32
	var transform Transform
	var v5 Filter
	var v6 QueryFilter
	var _ /* cache at bp+184 */ SimplexCache
	var _ /* input at bp+0 */ DistanceInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, id, output, result, shape, shapeId, tolerance, transform, world, worldContext, v1, v2, v3, v5, v6, v7
	_ = uint64FromInt64(4)
	shapeId = int32FromUint64(userData)
	worldContext = context
	world = (*WorldOverlapContext)(unsafe.Pointer(worldContext)).World
	v1 = world + 1248
	v2 = shapeId
	if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
	goto _4
_4:
	shape = v3
	v5 = (*b2Shape)(unsafe.Pointer(shape)).Filter
	v6 = (*WorldOverlapContext)(unsafe.Pointer(worldContext)).Filter
	v7 = boolUint8(v5.CategoryBits&v6.MaskBits != uint64(0) && v5.MaskBits&v6.CategoryBits != uint64(0))
	goto _8
_8:
	if int32FromUint8(v7) == false1 {
		return boolUint8(true1 != 0)
	}
	v1 = world + 1024
	v2 = (*b2Shape)(unsafe.Pointer(shape)).BodyId
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _12
_12:
	body = v3
	transform = b2GetBodyTransformQuick(tls, world, body)
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyA = *(*ShapeProxy)(unsafe.Add(nil, (*WorldOverlapContext)(unsafe.Pointer(worldContext)).Proxy))
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyB = b2MakeShapeDistanceProxy(tls, shape)
	(*(*DistanceInput)(unsafe.Add(nil, bp))).TransformA = b2Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp))).TransformB = transform
	(*(*DistanceInput)(unsafe.Add(nil, bp))).UseRadii = boolUint8(true1 != 0)
	*(*SimplexCache)(unsafe.Add(nil, bp+184)) = SimplexCache{}
	output = b2ShapeDistance(tls, bp, bp+184, uintptrFromInt32(0), 0)
	tolerance = float32(float32FromFloat32(0.1) * float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter))
	if output.Distance > tolerance {
		return boolUint8(true1 != 0)
	}
	id = ShapeId{
		Index1:     (*b2Shape)(unsafe.Pointer(shape)).Id + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2Shape)(unsafe.Pointer(shape)).Generation,
	}
	result = (*(*func(*_Stack, ShapeId, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*WorldOverlapContext)(unsafe.Pointer(worldContext)).Fcn})))(tls, id, (*WorldOverlapContext)(unsafe.Pointer(worldContext)).UserContext)
	return result
}

type __ccgo_fp__Xb2World_OverlapShape_3 = func(*_Stack, struct {
	Index1     int32
	World0     uint16
	Generation uint16
}, uintptr) uint8

type WorldRayCastContext struct {
	World       uintptr
	Fcn         uintptr
	Filter      QueryFilter
	Fraction    float32
	UserContext uintptr
}

func RayCastCallback(tls *_Stack, input uintptr, proxyId int32, userData uint64, context uintptr) (r float32) {
	var body, shape, world, worldContext, v1, v3 uintptr
	var fraction float32
	var id ShapeId
	var output CastOutput
	var shapeId, v2 int32
	var transform Transform
	var v5 Filter
	var v6 QueryFilter
	var v7 uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, fraction, id, output, shape, shapeId, transform, world, worldContext, v1, v2, v3, v5, v6, v7
	_ = uint64FromInt64(4)
	shapeId = int32FromUint64(userData)
	worldContext = context
	world = (*WorldRayCastContext)(unsafe.Pointer(worldContext)).World
	v1 = world + 1248
	v2 = shapeId
	if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
	goto _4
_4:
	shape = v3
	v5 = (*b2Shape)(unsafe.Pointer(shape)).Filter
	v6 = (*WorldRayCastContext)(unsafe.Pointer(worldContext)).Filter
	v7 = boolUint8(v5.CategoryBits&v6.MaskBits != uint64(0) && v5.MaskBits&v6.CategoryBits != uint64(0))
	goto _8
_8:
	if int32FromUint8(v7) == false1 {
		return (*RayCastInput)(unsafe.Pointer(input)).MaxFraction
	}
	v1 = world + 1024
	v2 = (*b2Shape)(unsafe.Pointer(shape)).BodyId
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _12
_12:
	body = v3
	transform = b2GetBodyTransformQuick(tls, world, body)
	output = b2RayCastShape(tls, input, shape, transform)
	if output.Hit != 0 {
		id = ShapeId{
			Index1:     shapeId + int32(1),
			World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
			Generation: (*b2Shape)(unsafe.Pointer(shape)).Generation,
		}
		fraction = (*(*func(*_Stack, ShapeId, Vec2, Vec2, float32, uintptr) float32)(unsafe.Pointer(&struct{ uintptr }{(*WorldRayCastContext)(unsafe.Pointer(worldContext)).Fcn})))(tls, id, output.Point, output.Normal, output.Fraction, (*WorldRayCastContext)(unsafe.Pointer(worldContext)).UserContext)
		// The user may return -1 to skip this shape
		if float32FromFloat32(0) <= fraction && fraction <= float32FromFloat32(1) {
			(*WorldRayCastContext)(unsafe.Pointer(worldContext)).Fraction = fraction
		}
		return fraction
	}
	return (*RayCastInput)(unsafe.Pointer(input)).MaxFraction
}

type __ccgo_fp__Xb2World_CastRay_4 = func(*_Stack, struct {
	Index1     int32
	World0     uint16
	Generation uint16
}, struct {
	X float32
	Y float32
}, struct {
	X float32
	Y float32
}, float32, uintptr) float32

// C documentation
//
//	// This callback finds the closest hit. This is the most common callback used in games.
func b2RayCastClosestFcn(tls *_Stack, shapeId ShapeId, point Vec2, normal Vec2, fraction float32, context uintptr) (r float32) {
	var rayResult uintptr
	_ = rayResult
	// Ignore initial overlap
	if fraction == float32FromFloat32(0) {
		return -float32FromFloat32(1)
	}
	rayResult = context
	(*RayResult)(unsafe.Pointer(rayResult)).ShapeId = shapeId
	(*RayResult)(unsafe.Pointer(rayResult)).Point = point
	(*RayResult)(unsafe.Pointer(rayResult)).Normal = normal
	(*RayResult)(unsafe.Pointer(rayResult)).Fraction = fraction
	(*RayResult)(unsafe.Pointer(rayResult)).Hit = boolUint8(true1 != 0)
	return fraction
}

func ShapeCastCallback(tls *_Stack, input uintptr, proxyId int32, userData uint64, context uintptr) (r float32) {
	var body, shape, world, worldContext, v1, v3 uintptr
	var fraction float32
	var id ShapeId
	var output CastOutput
	var shapeId, v2 int32
	var transform Transform
	var v5 Filter
	var v6 QueryFilter
	var v7 uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, fraction, id, output, shape, shapeId, transform, world, worldContext, v1, v2, v3, v5, v6, v7
	_ = uint64FromInt64(4)
	shapeId = int32FromUint64(userData)
	worldContext = context
	world = (*WorldRayCastContext)(unsafe.Pointer(worldContext)).World
	v1 = world + 1248
	v2 = shapeId
	if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
	goto _4
_4:
	shape = v3
	v5 = (*b2Shape)(unsafe.Pointer(shape)).Filter
	v6 = (*WorldRayCastContext)(unsafe.Pointer(worldContext)).Filter
	v7 = boolUint8(v5.CategoryBits&v6.MaskBits != uint64(0) && v5.MaskBits&v6.CategoryBits != uint64(0))
	goto _8
_8:
	if int32FromUint8(v7) == false1 {
		return (*ShapeCastInput)(unsafe.Pointer(input)).MaxFraction
	}
	v1 = world + 1024
	v2 = (*b2Shape)(unsafe.Pointer(shape)).BodyId
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _12
_12:
	body = v3
	transform = b2GetBodyTransformQuick(tls, world, body)
	output = b2ShapeCastShape(tls, input, shape, transform)
	if output.Hit != 0 {
		id = ShapeId{
			Index1:     shapeId + int32(1),
			World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
			Generation: (*b2Shape)(unsafe.Pointer(shape)).Generation,
		}
		fraction = (*(*func(*_Stack, ShapeId, Vec2, Vec2, float32, uintptr) float32)(unsafe.Pointer(&struct{ uintptr }{(*WorldRayCastContext)(unsafe.Pointer(worldContext)).Fcn})))(tls, id, output.Point, output.Normal, output.Fraction, (*WorldRayCastContext)(unsafe.Pointer(worldContext)).UserContext)
		// The user may return -1 to skip this shape
		if float32FromFloat32(0) <= fraction && fraction <= float32FromFloat32(1) {
			(*WorldRayCastContext)(unsafe.Pointer(worldContext)).Fraction = fraction
		}
		return fraction
	}
	return (*ShapeCastInput)(unsafe.Pointer(input)).MaxFraction
}

type __ccgo_fp__Xb2World_CastShape_4 = func(*_Stack, struct {
	Index1     int32
	World0     uint16
	Generation uint16
}, struct {
	X float32
	Y float32
}, struct {
	X float32
	Y float32
}, float32, uintptr) float32

type b2CharacterCallbackContext struct {
	World       uintptr
	Filter      QueryFilter
	Proxy       ShapeProxy
	Transform   Transform
	UserContext uintptr
}

type b2MoverContext = b2CharacterCallbackContext

type WorldMoverCastContext struct {
	World    uintptr
	Filter   QueryFilter
	Fraction float32
}

func MoverCastCallback(tls *_Stack, input uintptr, proxyId int32, userData uint64, context uintptr) (r float32) {
	var body, shape, world, worldContext, v1, v3 uintptr
	var output CastOutput
	var shapeId, v2 int32
	var transform Transform
	var v5 Filter
	var v6 QueryFilter
	var v7 uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _ = body, output, shape, shapeId, transform, world, worldContext, v1, v2, v3, v5, v6, v7
	_ = uint64FromInt64(4)
	shapeId = int32FromUint64(userData)
	worldContext = context
	world = (*WorldMoverCastContext)(unsafe.Pointer(worldContext)).World
	v1 = world + 1248
	v2 = shapeId
	if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
	goto _4
_4:
	shape = v3
	v5 = (*b2Shape)(unsafe.Pointer(shape)).Filter
	v6 = (*WorldMoverCastContext)(unsafe.Pointer(worldContext)).Filter
	v7 = boolUint8(v5.CategoryBits&v6.MaskBits != uint64(0) && v5.MaskBits&v6.CategoryBits != uint64(0))
	goto _8
_8:
	if int32FromUint8(v7) == false1 {
		return (*WorldMoverCastContext)(unsafe.Pointer(worldContext)).Fraction
	}
	v1 = world + 1024
	v2 = (*b2Shape)(unsafe.Pointer(shape)).BodyId
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _12
_12:
	body = v3
	transform = b2GetBodyTransformQuick(tls, world, body)
	output = b2ShapeCastShape(tls, input, shape, transform)
	if output.Fraction == float32FromFloat32(0) {
		// Ignore overlapping shapes
		return (*WorldMoverCastContext)(unsafe.Pointer(worldContext)).Fraction
	}
	(*WorldMoverCastContext)(unsafe.Pointer(worldContext)).Fraction = output.Fraction
	return output.Fraction
}

type WorldMoverContext struct {
	World       uintptr
	Fcn         uintptr
	Filter      QueryFilter
	Mover       Capsule
	UserContext uintptr
}

func TreeCollideCallback(tls *_Stack, proxyId int32, userData uint64, context uintptr) (r uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var aa, v16, v20, v21, v23 float32
	var body, shape, world, worldContext, v1, v3 uintptr
	var id ShapeId
	var shapeId, v2 int32
	var transform Transform
	var v13, v14, v15 Vec2
	var v7 uint8
	var v24 bool
	var v5 Filter
	var v6 QueryFilter
	var _ /* result at bp+0 */ PlaneResult
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aa, body, id, shape, shapeId, transform, world, worldContext, v1, v13, v14, v15, v16, v2, v20, v21, v23, v24, v3, v5, v6, v7
	_ = uint64FromInt64(4)
	shapeId = int32FromUint64(userData)
	worldContext = context
	world = (*WorldMoverContext)(unsafe.Pointer(worldContext)).World
	v1 = world + 1248
	v2 = shapeId
	if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
	goto _4
_4:
	shape = v3
	v5 = (*b2Shape)(unsafe.Pointer(shape)).Filter
	v6 = (*WorldMoverContext)(unsafe.Pointer(worldContext)).Filter
	v7 = boolUint8(v5.CategoryBits&v6.MaskBits != uint64(0) && v5.MaskBits&v6.CategoryBits != uint64(0))
	goto _8
_8:
	if int32FromUint8(v7) == false1 {
		return boolUint8(true1 != 0)
	}
	v1 = world + 1024
	v2 = (*b2Shape)(unsafe.Pointer(shape)).BodyId
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _12
_12:
	body = v3
	transform = b2GetBodyTransformQuick(tls, world, body)
	*(*PlaneResult)(unsafe.Add(nil, bp)) = b2CollideMover(tls, shape, transform, worldContext+32)
	// todo handle deep overlap
	if v24 = (*(*PlaneResult)(unsafe.Add(nil, bp))).Hit != 0; v24 {
		v13 = (*(*PlaneResult)(unsafe.Add(nil, bp))).Plane.Normal
		v14 = v13
		v15 = v13
		v16 = float32(v14.X*v15.X) + float32(v14.Y*v15.Y)
		goto _17
	_17:
		aa = v16
		v20 = float32FromFloat32(1) - aa
		if v20 < float32FromInt32(0) {
			v23 = -v20
		} else {
			v23 = v20
		}
		v21 = v23
		goto _22
	_22:
		v7 = boolUint8(v21 < float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07)))
		goto _19
	_19:
	}
	if v24 && v7 != 0 {
		id = ShapeId{
			Index1:     (*b2Shape)(unsafe.Pointer(shape)).Id + int32(1),
			World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
			Generation: (*b2Shape)(unsafe.Pointer(shape)).Generation,
		}
		return (*(*func(*_Stack, ShapeId, uintptr, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*WorldMoverContext)(unsafe.Pointer(worldContext)).Fcn})))(tls, id, bp, (*WorldMoverContext)(unsafe.Pointer(worldContext)).UserContext)
	}
	return boolUint8(true1 != 0)
}

type __ccgo_fp__Xb2World_CollideMover_3 = func(*_Stack, struct {
	Index1     int32
	World0     uint16
	Generation uint16
}, uintptr, uintptr) uint8

type __ccgo_fp__Xb2World_SetCustomFilterCallback_1 = func(*_Stack, struct {
	Index1     int32
	World0     uint16
	Generation uint16
}, struct {
	Index1     int32
	World0     uint16
	Generation uint16
}, uintptr) uint8

type __ccgo_fp__Xb2World_SetPreSolveCallback_1 = func(*_Stack, struct {
	Index1     int32
	World0     uint16
	Generation uint16
}, struct {
	Index1     int32
	World0     uint16
	Generation uint16
}, uintptr, uintptr) uint8

type ExplosionContext struct {
	World            uintptr
	Position         Vec2
	Radius           float32
	Falloff          float32
	ImpulsePerLength float32
}

func ExplosionCallback(tls *_Stack, proxyId int32, userData uint64, context uintptr) (r uint8) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var body, bodySim, explosionContext, set, shape, state, world, v1, v3, v41, v43, v5, v7 uintptr
	var closestPoint, direction, impulse, localCentroid, localLine, n, v10, v11, v13, v14, v15, v17 Vec2
	var falloff, invLength, length, magnitude, perimeter, radius, scale, x, y, v18, v30, v31, v32, v33, v35 float32
	var localIndex, shapeId, v2, v42, v6 int32
	var output DistanceOutput
	var transform, v9 Transform
	var v26 Rot
	var _ /* cache at bp+184 */ SimplexCache
	var _ /* input at bp+0 */ DistanceInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodySim, closestPoint, direction, explosionContext, falloff, impulse, invLength, length, localCentroid, localIndex, localLine, magnitude, n, output, perimeter, radius, scale, set, shape, shapeId, state, transform, world, x, y, v1, v10, v11, v13, v14, v15, v17, v18, v2, v26, v3, v30, v31, v32, v33, v35, v41, v42, v43, v5, v6, v7, v9
	_ = uint64FromInt64(4)
	shapeId = int32FromUint64(userData)
	explosionContext = context
	world = (*ExplosionContext)(unsafe.Pointer(explosionContext)).World
	v1 = world + 1248
	v2 = shapeId
	if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
	goto _4
_4:
	shape = v3
	v5 = world + 1024
	v6 = (*b2Shape)(unsafe.Pointer(shape)).BodyId
	if !(0 <= v6 && v6 < (*b2BodyArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2BodyArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*128
	goto _8
_8:
	body = v7
	if !((*b2Body)(unsafe.Pointer(body)).Type1 == int32(b2_dynamicBody)) && b2InternalAssertFcn(tls, __ccgo_ts+16892, __ccgo_ts+15342, int32FromInt32(2652)) != 0 {
		__builtin_trap(tls)
	}
	transform = b2GetBodyTransformQuick(tls, world, body)
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyA = b2MakeShapeDistanceProxy(tls, shape)
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyB = b2MakeProxy(tls, explosionContext+8, int32(1), float32FromFloat32(0))
	(*(*DistanceInput)(unsafe.Add(nil, bp))).TransformA = transform
	(*(*DistanceInput)(unsafe.Add(nil, bp))).TransformB = b2Transform_identity
	(*(*DistanceInput)(unsafe.Add(nil, bp))).UseRadii = boolUint8(true1 != 0)
	*(*SimplexCache)(unsafe.Add(nil, bp+184)) = SimplexCache{}
	output = b2ShapeDistance(tls, bp, bp+184, uintptrFromInt32(0), 0)
	radius = (*ExplosionContext)(unsafe.Pointer(explosionContext)).Radius
	falloff = (*ExplosionContext)(unsafe.Pointer(explosionContext)).Falloff
	if output.Distance > radius+falloff {
		return boolUint8(true1 != 0)
	}
	b2WakeBody(tls, world, body)
	if (*b2Body)(unsafe.Pointer(body)).SetIndex != int32(b2_awakeSet) {
		return boolUint8(true1 != 0)
	}
	closestPoint = output.PointA
	if output.Distance == float32FromFloat32(0) {
		localCentroid = b2GetShapeCentroid(tls, shape)
		v9 = transform
		v10 = localCentroid
		x = float32(v9.Q.C*v10.X) - float32(v9.Q.S*v10.Y) + v9.P.X
		y = float32(v9.Q.S*v10.X) + float32(v9.Q.C*v10.Y) + v9.P.Y
		v11 = Vec2{
			X: x,
			Y: y,
		}
		goto _12
	_12:
		closestPoint = v11
	}
	v10 = closestPoint
	v11 = (*ExplosionContext)(unsafe.Pointer(explosionContext)).Position
	v13 = Vec2{
		X: v10.X - v11.X,
		Y: v10.Y - v11.Y,
	}
	goto _16
_16:
	direction = v13
	v14 = direction
	v18 = float32(v14.X*v14.X) + float32(v14.Y*v14.Y)
	goto _19
_19:
	if v18 > float32(float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07))*float32FromFloat32(1.1920928955078125e-07)) {
		v15 = direction
		length = sqrtf(tls, float32(v15.X*v15.X)+float32(v15.Y*v15.Y))
		if length < float32FromFloat32(1.1920928955078125e-07) {
			v17 = Vec2{}
			goto _22
		}
		invLength = float32FromFloat32(1) / length
		n = Vec2{
			X: float32(invLength * v15.X),
			Y: float32(invLength * v15.Y),
		}
		v17 = n
		goto _22
	_22:
		direction = v17
	} else {
		direction = Vec2{
			X: float32FromFloat32(1),
		}
	}
	v10 = direction
	v11 = Vec2{
		X: -v10.Y,
		Y: v10.X,
	}
	goto _25
_25:
	v26 = transform.Q
	v13 = v11
	v14 = Vec2{
		X: float32(v26.C*v13.X) + float32(v26.S*v13.Y),
		Y: float32(-v26.S*v13.X) + float32(v26.C*v13.Y),
	}
	goto _29
_29:
	localLine = v14
	perimeter = b2GetShapeProjectedPerimeter(tls, shape, localLine)
	scale = float32FromFloat32(1)
	if output.Distance > radius && falloff > float32FromFloat32(0) {
		v18 = (radius + falloff - output.Distance) / falloff
		v30 = float32FromFloat32(0)
		v31 = float32FromFloat32(1)
		if v18 < v30 {
			v33 = v30
		} else {
			if v18 > v31 {
				v35 = v31
			} else {
				v35 = v18
			}
			v33 = v35
		}
		v32 = v33
		goto _34
	_34:
		scale = v32
	}
	magnitude = float32(float32((*ExplosionContext)(unsafe.Pointer(explosionContext)).ImpulsePerLength*perimeter) * scale)
	v18 = magnitude
	v10 = direction
	v11 = Vec2{
		X: float32(v18 * v10.X),
		Y: float32(v18 * v10.Y),
	}
	goto _40
_40:
	impulse = v11
	localIndex = (*b2Body)(unsafe.Pointer(body)).LocalIndex
	v1 = world + 1064
	v2 = int32(b2_awakeSet)
	if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*88
	goto _44
_44:
	set = v3
	v5 = set + 16
	v6 = localIndex
	if !(0 <= v6 && v6 < (*b2BodyStateArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(194)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2BodyStateArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*32
	goto _48
_48:
	state = v7
	v41 = set
	v42 = localIndex
	if !(0 <= v42 && v42 < (*b2BodySimArray)(unsafe.Pointer(v41)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
		__builtin_trap(tls)
	}
	v43 = (*b2BodySimArray)(unsafe.Pointer(v41)).Data + uintptr(v42)*100
	goto _52
_52:
	bodySim = v43
	v13 = (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
	v30 = (*b2BodySim)(unsafe.Pointer(bodySim)).InvMass
	v14 = impulse
	v15 = Vec2{
		X: v13.X + float32(v30*v14.X),
		Y: v13.Y + float32(v30*v14.Y),
	}
	goto _57
_57:
	(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity = v15
	v10 = closestPoint
	v11 = (*b2BodySim)(unsafe.Pointer(bodySim)).Center
	v13 = Vec2{
		X: v10.X - v11.X,
		Y: v10.Y - v11.Y,
	}
	goto _61
_61:
	v14 = v13
	v15 = impulse
	v18 = float32(v14.X*v15.Y) - float32(v14.Y*v15.X)
	goto _65
_65:
	*(*float32)(unsafe.Add(nil, state+8)) += float32((*b2BodySim)(unsafe.Pointer(bodySim)).InvInertia * v18)
	return boolUint8(true1 != 0)
}

func b2ValidateConnectivity(tls *_Stack, world uintptr) {
	_ = uint64FromInt64(4)
}

func b2ValidateContacts(tls *_Stack, world uintptr) {
	_ = uint64FromInt64(4)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var b2AssertHandler = uintptr(0)

var b2Mat22_zero = Mat22{}

var b2Rot_identity = Rot{
	C: float32FromFloat32(1),
}

var b2Transform_identity = Transform{
	Q: Rot{
		C: float32FromFloat32(1),
	},
}

var b2Vec2_zero = Vec2{}

var b2_byteCount b2AtomicInt

var b2_emptySimplexCache = SimplexCache{}

// C documentation
//
//	// This allows the user to change the length units at runtime
var b2_lengthUnitsPerMeter = float32FromFloat32(1)

var b2_nullBodyId = BodyId{}

var b2_nullChainId = ChainId{}

var b2_nullJointId = JointId{}

var b2_nullShapeId = ShapeId{}

var b2_nullWorldId = WorldId{}

var b2_worlds [128]b2World

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "capacity >= 0\x00../box2d-c/src/arena_allocator.c\x00( (uintptr_t)entry.data & 0x1F ) == 0\x00entryCount > 0\x00mem == entry->data\x00a->count > 0\x00../box2d-c/src/arena_allocator.h\x00alloc->allocation == 0\x00blockCount > bitSet->blockCount\x00../box2d-c/src/bitset.c\x00bitSet->bits != NULL\x00setA->blockCount == setB->blockCount\x00b2Body_IsValid( bodyId )\x00../box2d-c/src/body.c\x000 <= index && index < a->count\x00../box2d-c/src/body.h\x00../box2d-c/src/solver_set.h\x00body->islandId == B2_NULL_INDEX\x00body->islandPrev == B2_NULL_INDEX\x00body->islandNext == B2_NULL_INDEX\x00setIndex != b2_disabledSet\x00../box2d-c/src/island.h\x00island->bodyCount > 0\x00island->tailBody == body->id\x00island->bodyCount == 0\x00island->contactCount == 0\x00island->jointCount == 0\x00../box2d-c/src/contact.h\x00def->internalValue == B2_SECRET_COOKIE\x00b2IsValidVec2( def->position )\x00b2IsValidRotation( def->rotation )\x00b2IsValidVec2( def->linearVelocity )\x00b2IsValidFloat( def->angularVelocity )\x00b2IsValidFloat( def->linearDamping ) && def->linearDamping >= 0.0f\x00b2IsValidFloat( def->angularDamping ) && def->angularDamping >= 0.0f\x00b2IsValidFloat( def->sleepThreshold ) && def->sleepThreshold >= 0.0f\x00b2IsValidFloat( def->gravityScale )\x00world->locked == false\x00world->solverSets.data[setId].setIndex == B2_NULL_INDEX\x000 <= setId && setId < world->solverSets.count\x00( (uintptr_t)bodyState & 0x1F ) == 0\x00world->bodies.data[bodyId].id == B2_NULL_INDEX\x00../box2d-c/src/joint.h\x00../box2d-c/src/shape.h\x00movedBody->localIndex == movedIndex\x00result == movedIndex\x00index <= capacity\x00body->inertia > 0.0f\x00b2IsValidVec2( position )\x00b2IsValidRotation( rotation )\x00body->generation == bodyId.generation\x00body->setIndex == b2_staticSet\x00joint->setIndex == b2_disabledSet\x00body->setIndex == b2_awakeSet\x00otherBody->setIndex == b2_disabledSet\x00joint->setIndex == b2_awakeSet\x00otherBody->setIndex == b2_awakeSet\x000 <= joint->colorIndex && joint->colorIndex < B2_GRAPH_COLOR_COUNT\x00originalType == b2_dynamicBody || originalType == b2_kinematicBody\x00type == b2_dynamicBody || type == b2_kinematicBody\x00b2IsValidFloat( massData.mass ) && massData.mass >= 0.0f\x00b2IsValidFloat( massData.rotationalInertia ) && massData.rotationalInertia >= 0.0f\x00b2IsValidVec2( massData.center )\x00b2IsValidFloat( linearDamping ) && linearDamping >= 0.0f\x00b2IsValidFloat( angularDamping ) && angularDamping >= 0.0f\x00b2IsValidFloat( gravityScale )\x00joint->setIndex == set->setIndex || set->setIndex == b2_staticSet\x00joint->islandId == B2_NULL_INDEX\x00../box2d-c/src/array.h\x000 <= proxyType && proxyType < b2_bodyTypeCount\x00../box2d-c/src/broad_phase.c\x00bp->moveArray.count == (int)bp->moveSet.count\x000 <= proxyType && proxyType <= b2_bodyTypeCount\x00proxyKey != B2_NULL_INDEX\x00typeIndex != b2_staticBody\x00treeType == b2_dynamicBody\x00moveCount == (int)bp->moveSet.count\x00move results\x00move pairs\x00i != B2_OVERFLOW_INDEX || color->bodySet.bits == NULL\x00../box2d-c/src/constraint_graph.c\x00contactSim->manifold.pointCount > 0\x00contactSim->simFlags & b2_simTouchingFlag\x00contact->flags & b2_contactTouchingFlag\x00staticA == false || staticB == false\x00bodyA->setIndex == b2_awakeSet\x00bodyB->setIndex == b2_awakeSet\x000 <= colorIndex && colorIndex < B2_GRAPH_COLOR_COUNT\x00movedContact->setIndex == b2_awakeSet\x00movedContact->colorIndex == colorIndex\x00movedContact->localIndex == movedIndex\x00movedJoint->setIndex == b2_awakeSet\x00movedJoint->colorIndex == colorIndex\x00movedJoint->localIndex == movedIndex\x000 <= type1 && type1 < b2_shapeTypeCount\x00../box2d-c/src/contact.c\x000 <= type2 && type2 < b2_shapeTypeCount\x00bodyA->setIndex != b2_disabledSet && bodyB->setIndex != b2_disabledSet\x00bodyA->setIndex != b2_staticSet || bodyB->setIndex != b2_staticSet\x00shapeA->sensorIndex == B2_NULL_INDEX && shapeB->sensorIndex == B2_NULL_INDEX\x00contact->setIndex == b2_awakeSet\x00contact->setIndex != b2_awakeSet || ( contact->flags & b2_contactTouchingFlag ) == 0\x000 <= contact->colorIndex && contact->colorIndex < B2_GRAPH_COLOR_COUNT\x000 < pointCount && pointCount <= 2\x00../box2d-c/src/contact_solver.c\x00b2IsValidFloat( lengthUnits ) && lengthUnits > 0.0f\x00../box2d-c/src/core.c\x00BOX2D ASSERTION: %s, %s, line %d\n\x00assertFcn != NULL\x00ptr != NULL\x00( (uintptr_t)ptr & 0x1F ) == 0\x00newSize > oldSize\x00cache->count <= 3\x00../box2d-c/src/distance.c\x00false\x00input->proxyA.count > 0 && input->proxyB.count > 0\x00input->proxyA.radius >= 0.0f\x00input->proxyB.radius >= 0.0f\x00b2IsNormalized( normal )\x00target > tolerance\x00distanceOutput.distance > 0.0f && b2IsNormalized( distanceOutput.normal )\x00distanceOutput.distance > 0.0f\x00b2IsNormalized( distanceOutput.normal )\x000 < count && count < 3\x00b2IsNormalizedRot( sweepA.q1 ) && b2IsNormalizedRot( sweepA.q2 )\x00b2IsNormalizedRot( sweepB.q1 ) && b2IsNormalizedRot( sweepB.q2 )\x00../box2d-c/src/distance_joint.c\x00base->type == b2_distanceJoint\x00bodyA->setIndex == b2_awakeSet || bodyB->setIndex == b2_awakeSet\x00tree->nodeCount == tree->nodeCapacity\x00../box2d-c/src/dynamic_tree.c\x00oldNodes != NULL\x000 <= nodeId && nodeId < tree->nodeCapacity\x000 < tree->nodeCount\x00lowerCost1 < FLT_MAX\x00lowerCost2 < FLT_MAX\x00nodes[index].height > 0\x00iA != B2_NULL_INDEX\x000 <= iB && iB < tree->nodeCapacity\x000 <= iC && iC < tree->nodeCapacity\x00C->height > 0\x000 <= iF && iF < tree->nodeCapacity\x000 <= iG && iG < tree->nodeCapacity\x00B->height > 0\x000 <= iD && iD < tree->nodeCapacity\x000 <= iE && iE < tree->nodeCapacity\x00child1 != B2_NULL_INDEX\x00child2 != B2_NULL_INDEX\x00-B2_HUGE < aabb.lowerBound.x && aabb.lowerBound.x < B2_HUGE\x00-B2_HUGE < aabb.lowerBound.y && aabb.lowerBound.y < B2_HUGE\x00-B2_HUGE < aabb.upperBound.x && aabb.upperBound.x < B2_HUGE\x00-B2_HUGE < aabb.upperBound.y && aabb.upperBound.y < B2_HUGE\x000 <= proxyId && proxyId < tree->nodeCapacity\x00b2IsLeaf( tree->nodes + proxyId )\x00tree->proxyCount > 0\x00b2IsValidAABB( aabb )\x00aabb.upperBound.x - aabb.lowerBound.x < B2_HUGE\x00aabb.upperBound.y - aabb.lowerBound.y < B2_HUGE\x00b2AABB_Contains( nodes[proxyId].aabb, aabb ) == false\x00nodes[proxyId].children.child1 == B2_NULL_INDEX\x00nodes[proxyId].children.child2 == B2_NULL_INDEX\x00(nodes[proxyId].flags & b2_leafNode) == b2_leafNode\x00stackCount < B2_TREE_STACK_SIZE - 1\x00i1 == i2\x00parentNode->children.child1 == B2_NULL_INDEX\x00parentItem->childCount == 1\x00parentNode->children.child2 == B2_NULL_INDEX\x00node->parent == B2_NULL_INDEX\x00node->children.child1 != B2_NULL_INDEX\x00node->children.child2 != B2_NULL_INDEX\x00item->childCount == 1\x00node->children.child1 == B2_NULL_INDEX\x00node->children.child2 == B2_NULL_INDEX\x00childNode->parent == B2_NULL_INDEX\x00count > 0\x00top < B2_TREE_STACK_SIZE\x00rootNode->parent == B2_NULL_INDEX\x00rootNode->children.child1 != B2_NULL_INDEX\x00rootNode->children.child2 != B2_NULL_INDEX\x00stackCount < B2_TREE_STACK_SIZE\x00leafCount <= proxyCount\x00area > FLT_EPSILON\x00../box2d-c/src/geometry.c\x00b2ValidateHull( hull )\x00b2Dot( edge, edge ) > FLT_EPSILON * FLT_EPSILON\x00b2IsValidFloat( halfWidth ) && halfWidth > 0.0f\x00b2IsValidFloat( halfHeight ) && halfHeight > 0.0f\x00b2IsValidFloat( radius ) && radius >= 0.0f\x00shape->count > 0\x00b2IsValidRay( input )\x000.0f <= lower && lower <= input->maxFraction\x00hull.count < B2_MAX_POLYGON_VERTICES\x00../box2d-c/src/hull.c\x00hull.count <= B2_MAX_POLYGON_VERTICES\x00pool->nextIndex > 0\x00../box2d-c/src/id_pool.c\x000 <= id && id < pool->nextIndex\x00setIndex == b2_awakeSet || setIndex >= b2_firstSleepingSet\x00../box2d-c/src/island.c\x00world->islands.data[islandId].setIndex == B2_NULL_INDEX\x00movedIsland->localIndex == movedIndex\x00contact->islandId == B2_NULL_INDEX\x00contact->islandPrev == B2_NULL_INDEX\x00contact->islandNext == B2_NULL_INDEX\x00( contact->flags & b2_contactTouchingFlag ) != 0\x00bodyA->setIndex != b2_staticSet || islandIdA == B2_NULL_INDEX\x00bodyB->setIndex != b2_staticSet || islandIdB == B2_NULL_INDEX\x00islandIdA != B2_NULL_INDEX || islandIdB != B2_NULL_INDEX\x00islandA != NULL || islandB != NULL\x00islandA != islandB\x00islandB->parentIsland == B2_NULL_INDEX\x00contact->islandId != B2_NULL_INDEX\x00prevContact->islandNext == contact->contactId\x00nextContact->islandPrev == contact->contactId\x00island->contactCount > 0\x00joint->islandPrev == B2_NULL_INDEX\x00joint->islandNext == B2_NULL_INDEX\x00joint->islandId != B2_NULL_INDEX\x00prevJoint->islandNext == joint->jointId\x00nextJoint->islandPrev == joint->jointId\x00island->jointCount > 0\x00island->parentIsland != B2_NULL_INDEX\x00rootIsland->parentIsland == B2_NULL_INDEX\x00rootIsland->tailBody != B2_NULL_INDEX\x00tailBody->islandNext == B2_NULL_INDEX\x00island->headBody != B2_NULL_INDEX\x00headBody->islandPrev == B2_NULL_INDEX\x00rootIsland->tailContact == B2_NULL_INDEX && rootIsland->contactCount == 0\x00island->tailContact != B2_NULL_INDEX && island->contactCount > 0\x00rootIsland->tailContact != B2_NULL_INDEX && rootIsland->contactCount > 0\x00tailContact->islandNext == B2_NULL_INDEX\x00headContact->islandPrev == B2_NULL_INDEX\x00rootIsland->tailJoint == B2_NULL_INDEX && rootIsland->jointCount == 0\x00island->tailJoint != B2_NULL_INDEX && island->jointCount > 0\x00rootIsland->tailJoint != B2_NULL_INDEX && rootIsland->jointCount > 0\x00tailJoint->islandNext == B2_NULL_INDEX\x00headJoint->islandPrev == B2_NULL_INDEX\x00island stack\x00body ids\x00index == bodyCount\x00seed->setIndex == setIndex\x00body->isMarked == true\x00contact->contactId == contactId\x00stackCount < bodyCount\x00joint->jointId == jointId\x00world->splitIslandId != B2_NULL_INDEX\x00joint->jointId == id && joint->generation == jointId.generation\x00../box2d-c/src/joint.c\x00joint->type == type\x00jointSim->type == type\x00bodyA->setIndex >= b2_firstSleepingSet || bodyB->setIndex >= b2_firstSleepingSet\x00bodyA->setIndex == bodyB->setIndex\x00joint->setIndex == setIndex\x00jointSim->jointId == jointId\x00jointSim->bodyIdA == bodyIdA\x00jointSim->bodyIdB == bodyIdB\x00b2Body_IsValid( def->bodyIdA )\x00b2Body_IsValid( def->bodyIdB )\x00b2IsValidFloat( def->length ) && def->length > 0.0f\x00def->lowerAngle <= def->upperAngle\x00def->lowerAngle >= -0.99f * B2_PI\x00def->upperAngle <= 0.99f * B2_PI\x00def->lowerTranslation <= def->upperTranslation\x00joint->setIndex > b2_disabledSet\x00joint->setIndex <= b2_disabledSet\x00b2IsValidVec2( localAnchor )\x00b2IsValidFloat( angleInRadians )\x00b2IsValidVec2( localAxis )\x00b2IsNormalized( localAxis )\x00b2IsValidFloat( hertz ) && hertz >= 0.0f\x00b2IsValidFloat( dampingRatio ) && dampingRatio >= 0.0f\x00b2LengthSquared( d ) > FLT_EPSILON\x00../box2d-c/src/manifold.c\x00dd1 > epsSqr && dd2 > epsSqr\x00result.distanceSquared > 0.0f\x00cache->count == 2\x00ib1 != ib2\x00manifold.pointCount == 0 || manifold.pointCount == 2\x00incidentNormal != -1 || incidentIndex != -1\x00b2AbsFloat( 1.0f - b2Length( v1 ) ) < 100.0f * FLT_EPSILON\x00../box2d-c/src/math_functions.c\x00b2AbsFloat( 1.0f - b2Length( v2 ) ) < 100.0f * FLT_EPSILON\x00base->type == b2_motorJoint\x00../box2d-c/src/motor_joint.c\x00b2IsValidVec2( target )\x00../box2d-c/src/mouse_joint.c\x00b2IsValidFloat( maxForce ) && maxForce >= 0.0f\x00base->type == b2_mouseJoint\x00lower <= upper\x00../box2d-c/src/prismatic_joint.c\x00joint->type == b2_prismaticJoint\x00jointSim->type == b2_prismaticJoint\x00base->type == b2_prismaticJoint\x00../box2d-c/src/revolute_joint.c\x00lower >= -0.99f * B2_PI\x00upper <= 0.99f * B2_PI\x00base->type == b2_revoluteJoint\x00 %.1f deg\x00(int)threadIndex < world->workerCount\x00../box2d-c/src/sensor.c\x00startIndex < endIndex\x00../box2d-c/src/sensor.h\x00blockIndex < bitSet->blockCount\x00../box2d-c/src/bitset.h\x00sensorShape->sensorIndex == sensorIndex\x00world->workerCount > 0\x00shape->id == id && shape->generation == shapeId.generation\x00../box2d-c/src/shape.c\x00chain->id == id && chain->generation == chainId.generation\x00world->shapes.data[shapeId].id == B2_NULL_INDEX\x00b2IsValidFloat( def->density ) && def->density >= 0.0f\x00b2IsValidFloat( def->material.friction ) && def->material.friction >= 0.0f\x00b2IsValidFloat( def->material.restitution ) && def->material.restitution >= 0.0f\x00b2IsValidFloat( def->material.rollingResistance ) && def->material.rollingResistance >= 0.0f\x00b2IsValidFloat( def->material.tangentSpeed )\x00b2IsValidFloat( polygon->radius ) && polygon->radius >= 0.0f\x00def->count >= 4\x00def->materialCount == 1 || def->materialCount == def->count\x00world->chainShapes.data[chainId].id == B2_NULL_INDEX\x00b2IsValidFloat( material->friction ) && material->friction >= 0.0f\x00b2IsValidFloat( material->restitution ) && material->restitution >= 0.0f\x00b2IsValidFloat( material->rollingResistance ) && material->rollingResistance >= 0.0f\x00b2IsValidFloat( material->tangentSpeed )\x00found == true\x00shape->proxyKey == B2_NULL_INDEX\x00B2_PROXY_TYPE( shape->proxyKey ) < b2_bodyTypeCount\x00b2IsValidFloat( density ) && density >= 0.0f\x00b2IsValidFloat( friction ) && friction >= 0.0f\x00b2IsValidFloat( restitution ) && restitution >= 0.0f\x00shape->type == b2_circleShape\x00shape->type == b2_segmentShape\x00shape->type == b2_chainSegmentShape\x00shape->type == b2_capsuleShape\x00shape->type == b2_polygonShape\x00b2IsValidFloat( restitution )\x000 <= startIndex && startIndex < color->jointSims.count\x00../box2d-c/src/solver.c\x00startIndex <= endIndex && endIndex <= color->jointSims.count\x00startIndex <= endIndex\x00body->type == b2_staticBody || fastBodySim->isBullet\x00fastBodySim->isFast\x00../box2d-c/src/world.h\x00endIndex <= world->bodyMoveEvents.count\x00b2IsValidVec2( v )\x00b2IsValidFloat( w )\x00shape->enlargedAABB == false\x000 <= startIndex && startIndex < blockCount\x00stage->type != b2_stagePrepareContacts || syncIndex < 2\x00completedCount < blockCount\x00syncIndex > 0\x00stages[stageIndex].type == b2_stagePrepareJoints\x00stages[stageIndex].type == b2_stagePrepareContacts\x00stages[iterStageIndex].type == b2_stageIntegrateVelocities\x00stages[iterStageIndex].type == b2_stageWarmStart\x00stages[iterStageIndex].type == b2_stageSolve\x00stages[iterStageIndex].type == b2_stageIntegratePositions\x00stages[iterStageIndex].type == b2_stageRelax\x00stages[iterStageIndex].type == b2_stageRestitution\x00stages[stageIndex].type == b2_stageStoreImpulses\x00stageIndex + 1 == context->stageCount\x00stageIndex < context->stageCount\x00bullet bodies\x00contact pointers\x00joint pointers\x00contact constraint\x00overflow contact constraint\x00contactBase == simdContactCount\x00jointBase == awakeJointCount\x00stages\x00body blocks\x00contact blocks\x00joint blocks\x00graph blocks\x00(ptrdiff_t)(baseGraphBlock - graphBlocks) == graphBlockCount\x00(int)( stage - stages ) == stageCount\x00workerCount <= B2_MAX_WORKERS\x00world->contactHitEvents.count == 0\x000 <= bodyId && bodyId < world->bodies.count\x00B2_PROXY_TYPE( proxyKey ) == b2_dynamicBody\x00b2ContainsKey( &broadPhase->moveSet, proxyKey + 1 )\x00world->splitIslandId == B2_NULL_INDEX\x00taskContext->splitSleepTime > 0.0f\x00setIndex >= b2_firstSleepingSet\x00../box2d-c/src/solver_set.c\x00body->setIndex == setIndex\x00contact->setIndex == b2_awakeSet || contact->setIndex == setIndex\x00( contact->flags & b2_contactTouchingFlag ) == 0 && contactSim->manifold.pointCount == 0\x00movedContact->localIndex == movedLocalIndex\x00contact->setIndex == setIndex\x00island->setIndex == b2_awakeSet\x000 <= island->localIndex && island->localIndex < awakeSet->islandSims.count\x00body->islandId == islandId\x00moveEvent->bodyId.index1 - 1 == bodyId\x00moveEvent->bodyId.generation == body->generation\x00contact->setIndex == b2_awakeSet || contact->setIndex == b2_disabledSet\x00contactSim->manifold.pointCount == 0\x00( contact->flags & b2_contactTouchingFlag ) == 0\x00contact->islandId == islandId\x00joint->islandId == islandId\x00movedIsland->localIndex == movedIslandIndex\x00setId1 >= b2_firstSleepingSet\x00setId2 >= b2_firstSleepingSet\x00body->setIndex == setId2\x00contact->setIndex == setId2\x00joint->setIndex == setId2\x00targetSet != sourceSet\x00colorIndex == B2_NULL_INDEX\x00items[index].hash == 0\x00../box2d-c/src/table.c\x00set->count == oldCount\x00key != 0\x00hash != 0\x00set->items[index].hash == hash && set->items[index].key == key\x00set->count > 0\x00../box2d-c/src/weld_joint.c\x00base->type == b2_weldJoint\x00../box2d-c/src/wheel_joint.c\x00base->type == b2_wheelJoint\x001 <= id.index1 && id.index1 <= B2_MAX_WORLDS\x00../box2d-c/src/world.c\x00id.index1 == world->worldId + 1\x00id.generation == world->generation\x000 <= index && index < B2_MAX_WORLDS\x00world->worldId == index\x00world->solverSets.data[b2_staticSet].setIndex == b2_staticSet\x00world->solverSets.data[b2_disabledSet].setIndex == b2_disabledSet\x00world->solverSets.data[b2_awakeSet].setIndex == b2_awakeSet\x00chain->shapeIndices == NULL\x00chain->materials == NULL\x00movedContact->setIndex == setIndex\x00movedContact->colorIndex == B2_NULL_INDEX\x00contacts\x00contactIndex == contactCount\x00contact->colorIndex == B2_NULL_INDEX\x00contact->localIndex == localIndex\x00b2IsValidFloat( timeStep )\x000 < subStepCount\x00b2GetArenaAllocation( &world->arena ) == 0\x00world->activeTaskCount == 0\x00shape->id == shapeId\x00b2IsValidAABB( draw->drawingBounds )\x00  %.2f\x00%.1f\x00%d\x00%.2f\x00body->localIndex != B2_NULL_INDEX\x00chain->id == chainId\x00b2IsValidFloat( maximumLinearSpeed ) && maximumLinearSpeed > 0.0f\x00box2d_memory.txt\x00w\x00id pools\n\x00body ids: %d\n\x00solver set ids: %d\n\x00joint ids: %d\n\x00contact ids: %d\n\x00island ids: %d\n\x00shape ids: %d\n\x00chain ids: %d\n\x00\n\x00world arrays\n\x00bodies: %d\n\x00solver sets: %d\n\x00joints: %d\n\x00contacts: %d\n\x00islands: %d\n\x00shapes: %d\n\x00chains: %d\n\x00broad-phase\n\x00static tree: %d\n\x00kinematic tree: %d\n\x00dynamic tree: %d\n\x00moveSet: %d (%d, %d)\n\x00moveArray: %d\n\x00pairSet: %d (%d, %d)\n\x00solver sets\n\x00body sim: %d\n\x00body state: %d\n\x00joint sim: %d\n\x00contact sim: %d\n\x00island sim: %d\n\x00constraint graph\n\x00body bit sets: %d\n\x00stack allocator: %d\n\n\x00../box2d-c/include/box2d/math_functions.h\x00b2IsValidVec2( origin )\x00b2IsValidVec2( translation )\x00mover->radius > 2.0f * B2_LINEAR_SLOP\x00body->type == b2_dynamicBody\x00b2IsValidFloat( falloff ) && falloff >= 0.0f\x00b2IsValidFloat( impulsePerLength )\x00"
