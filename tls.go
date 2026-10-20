package b2

import (
	"unsafe"
)

var theStack *_Stack = newStack(1024 * 256)

type _Stack struct {
	Stack  []byte
	allocs []int
	pos    int

	// yields counts the calls to sched_yield, each goroutine of box2d has its own stack
	yields uint32
}

//go:noinline
func newStack(stackSize int) *_Stack {
	stack := &_Stack{
		Stack: make([]byte, stackSize),
	}

	escapes(stack)

	return stack
}

func (s *_Stack) CheckEmpty() {
	if s.pos != 0 || len(s.allocs) > 0 {
		panic("stack is not empty")
	}
}

func (s *_Stack) Alloc(n int) uintptr {
	if n%8 > 0 {
		// round up to a multiple of 8
		n += 8 - n%8
	}

	if s.pos+n > len(s.Stack) {
		panic("out of stack space")
	}

	buf := uintptr(unsafe.Pointer(unsafe.SliceData(s.Stack[s.pos:])))
	s.pos += n
	s.allocs = append(s.allocs, n)

	return buf
}

func (s *_Stack) Free(n int) uintptr {
	if n%8 > 0 {
		// round up to a multiple of 8
		n += 8 - n%8
	}

	lastAlloc := s.allocs[len(s.allocs)-1]
	if lastAlloc != n {
		panic("alloc/free is not paired")
	}

	s.allocs = s.allocs[:len(s.allocs)-1]
	s.pos -= n

	return 0
}

func (s *_Stack) toCString(src string) alloc {
	n := len(src) + 1
	mem := s.Alloc(n)

	buf := sliceOf(mem, size_t(n))
	copy(buf, src)
	buf[n-1] = 0

	return alloc{Addr: mem, tls: s, size: n}
}

type alloc struct {
	Addr uintptr
	tls  *_Stack
	size int
}

func (a alloc) Free() {
	a.tls.Free(a.size)
}

func copyToStack[T any](tls *_Stack, value T) alloc {
	n := int(unsafe.Sizeof(value))
	stack := tls.Alloc(n)
	*(*T)(unsafe.Pointer(stack)) = value

	return alloc{Addr: stack, tls: tls, size: n}
}

func sliceToStack[T any](tls *_Stack, values []T) alloc {
	n := len(values) * int(unsafe.Sizeof(*new(T)))
	stack := tls.Alloc(n)
	copy(unsafe.Slice((*T)(unsafe.Pointer(stack)), len(values)), values)

	return alloc{Addr: stack, tls: tls, size: n}
}
