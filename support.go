package b2

import (
	"fmt"
	"math"
	"math/bits"
	"regexp"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

func sliceOf(p uintptr, n size_t) []byte {
	arr := (*[math.MaxInt32]byte)(unsafe.Pointer(p))
	return (*arr)[:n]
}

func __atomic_fetch_addInt32(tls *_Stack, addr uintptr, inc int32, mode int32) int32 {
	value := (*int32)(unsafe.Pointer(addr))

	for {
		old := atomic.LoadInt32(value)
		if atomic.CompareAndSwapInt32(value, old, old+inc) {
			return old
		}
	}
}

func __atomic_fetch_addInt64(_ *_Stack, addr uintptr, inc int64, _ int32) int64 {
	return atomic.AddInt64((*int64)(unsafe.Pointer(addr)), inc) - inc
}

func __atomic_fetch_orUint32(_ *_Stack, addr uintptr, mask uint32, _ int32) uint32 {
	value := (*uint32)(unsafe.Pointer(addr))

	// atomic.OrUint32 needs go 1.23
	for {
		old := atomic.LoadUint32(value)
		if atomic.CompareAndSwapUint32(value, old, old|mask) {
			return old
		}
	}
}

func __builtin___atomic_compare_exchange_n(tls *_Stack, addr, expected uintptr, desired int32, weak uint8, success_memorder, failure_memorder int32) uint8 {
	ptrValue := (*int32)(unsafe.Pointer(addr))
	ptrExpected := (*int32)(unsafe.Pointer(expected))

	exp := *ptrExpected
	for {
		if atomic.CompareAndSwapInt32(ptrValue, exp, desired) {
			return true1
		}

		// like C, report the current value on failure. Retry if it changed back to exp in between.
		if current := atomic.LoadInt32(ptrValue); current != exp {
			*ptrExpected = current
			return false1
		}
	}
}

func remainderf(_ *_Stack, a, b float32) float32 {
	return float32(math.Remainder(float64(a), float64(b)))
}

func __builtin_trap(_ *_Stack) {
	panic("trap")
}

func __builtin_clz(t *_Stack, n uint32) int32 {
	return int32(bits.LeadingZeros32(n))
}

func __builtin_ctzll(_ *_Stack, value uint64) int32 {
	return int32(bits.TrailingZeros64(value))
}

func __builtin_aligned_alloc(tls *_Stack, align, size uint64) uintptr {
	buf := unsafe.Pointer(unsafe.SliceData(make([]byte, size+align)))
	addr := uintptr(buf)

	if addr%uintptr(align) != 0 {
		addr += uintptr(align) - addr%uintptr(align)
	}

	// address is now aligned
	if addr%uintptr(align) != 0 {
		panic("not aligned")
	}

	heap.Store(addr, buf)
	return addr
}

// heap keeps allocations alive. It is shared by all stacks, as box2d may allocate
// on a worker goroutine and free on another one.
var heap sync.Map

func memset(_ *_Stack, dest uintptr, c int32, n size_t) (r uintptr) {
	slice := sliceOf(dest, n)

	if c == 0 {
		clear(slice)
	} else {
		for idx := range slice {
			slice[idx] = byte(c)
		}
	}

	return dest
}

func memcpy(_ *_Stack, a uintptr, b uintptr, n size_t) uintptr {
	dst := sliceOf(a, n)
	src := sliceOf(b, n)
	copy(dst, src)
	return a
}

func free(tls *_Stack, ptr uintptr) {
	if _, ok := heap.LoadAndDelete(ptr); !ok {
		err := fmt.Errorf("free(%#x): unknown address", ptr)
		panic(err)
	}
}

// strncpy copies src up to n bytes and pads dest with zeros, like C.
func strncpy(_ *_Stack, dest, src uintptr, n size_t) uintptr {
	target := sliceOf(dest, n)
	clear(target[copy(target, toString(src)):])
	return dest
}

func sqrtf(_ *_Stack, f float32) float32 {
	return float32(math.Sqrt(float64(f)))
}

type _qsort struct {
	tls      *_Stack
	memory   []byte
	compare  func(_ *_Stack, a, b uintptr) int32
	itemSize int
	len      int
}

func (q *_qsort) sliceOf(idx int) []byte {
	start := idx * q.itemSize
	return q.memory[start:][:q.itemSize]
}

func (q *_qsort) Len() int {
	return q.len
}

func (q *_qsort) Less(i, j int) bool {
	lhs := q.sliceOf(i)
	rhs := q.sliceOf(j)

	rc := q.compare(
		q.tls,
		uintptr(unsafe.Pointer(unsafe.SliceData(lhs))),
		uintptr(unsafe.Pointer(unsafe.SliceData(rhs))),
	)

	return rc < 0
}

func (q *_qsort) Swap(i, j int) {
	lhs := q.sliceOf(i)
	rhs := q.sliceOf(j)

	for idx := range len(lhs) {
		// swap the bytes
		lhs[idx], rhs[idx] = rhs[idx], lhs[idx]
	}
}

func qsort(tls *_Stack, base uintptr, nel size_t, width size_t, compare func(_ *_Stack, a, b uintptr) int32) {
	sort.Sort(&_qsort{
		tls:      tls,
		memory:   sliceOf(base, nel*width),
		compare:  compare,
		itemSize: int(width),
		len:      int(nel),
	})
}

func clock_gettime(tls *_Stack, clk clockid_t, ts uintptr) (r int32) {
	now := time.Now()

	(*timespec)(unsafe.Pointer(ts)).Tv_sec = now.Unix()
	(*timespec)(unsafe.Pointer(ts)).Tv_nsec = now.UnixNano() - now.Unix()*1_000_000_000

	return 0
}

func sched_yield(tls *_Stack) int32 {
	time.Sleep(20 * time.Microsecond)
	return 0
}

func __builtin_Gosched(tls *_Stack) {
	pause()
}

func printf(tls *_Stack, format uintptr, va uintptr) (r int32) {
	panic("not required, only used in default assert function")
}

var reFormatStr = regexp.MustCompile("%[^a-zA-Z]*([dfgs])")

// cFormat formats like C printf, for the patterns that box2d uses.
// Like the varargs of ccgo, every argument takes 8 bytes at args.
func cFormat(format, args uintptr) string {
	formatStr := toString(format)

	var values []any
	for _, match := range reFormatStr.FindAllStringSubmatch(formatStr, -1) {
		switch match[1] {
		case "d":
			values = append(values, *(*int64)(unsafe.Pointer(args)))
		case "f", "g":
			values = append(values, *(*float64)(unsafe.Pointer(args)))
		case "s":
			values = append(values, toString(*(*uintptr)(unsafe.Pointer(args))))
		}

		args += 8
	}

	return fmt.Sprintf(formatStr, values...)
}

func __builtin_snprintf(_ *_Stack, str uintptr, size size_t, format, args uintptr) int32 {
	return __builtin_vsnprintf(nil, str, size, format, args)
}

func __builtin_vsnprintf(_ *_Stack, str uintptr, size size_t, format, args uintptr) int32 {
	result := cFormat(format, args)

	if size > 0 {
		target := sliceOf(str, size)
		clear(target)
		copy(target[:size-1], result)
	}

	return int32(len(result))
}

func __builtin_inff(_ *_Stack) float32 {
	return float32(math.Inf(1))
}

func __builtin_ctz(_ *_Stack, value uint32) int32 {
	return int32(bits.TrailingZeros32(value))
}

func __builtin_popcountll(_ *_Stack, value uint64) int32 {
	return int32(bits.OnesCount64(value))
}

func fprintf(tls *_Stack, f uintptr, fmt uintptr, va uintptr) (r int32) {
	panic("fprintf: not implemented")
}

func fopen(tls *_Stack, filename uintptr, mode uintptr) (r uintptr) {
	panic("fopen: not implemented")
}

func fwrite(tls *_Stack, ptr uintptr, size, n size_t, f uintptr) size_t {
	panic("fwrite: not implemented")
}

func fread(tls *_Stack, ptr uintptr, size, n size_t, f uintptr) size_t {
	panic("fread: not implemented")
}

func fseek(tls *_Stack, f uintptr, offset int64, whence int32) int32 {
	panic("fseek: not implemented")
}

func ftell(tls *_Stack, f uintptr) int64 {
	panic("ftell: not implemented")
}

func fclose(tls *_Stack, f uintptr) (r1 int32) {
	panic("fclose: not implemented")
}

// Dummy annotation marking that the value x escapes,
// for use in cases where the reflect code is so clever that
// the compiler cannot follow.
func escapes(x any) {
	if dummy.b {
		dummy.x = x
	}
}

var dummy struct {
	b bool
	x any
}

func toString(base uintptr) string {
	var str strings.Builder

	for n := 0; ; n += 1 {
		ch := *(*byte)(unsafe.Pointer(base + uintptr(n)))
		if ch == 0 {
			return str.String()
		}

		str.WriteByte(ch)
	}
}
