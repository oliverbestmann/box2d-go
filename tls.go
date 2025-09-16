package b2

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

var theStack *_Stack = newStack(1024 * 256)

type _Stack struct {
	Stack  []byte
	allocs []int
	pos    int

	heap    map[uintptr]unsafe.Pointer
	objects map[uintptr]any
}

//go:noinline
func newStack(stackSize int) *_Stack {
	stack := &_Stack{
		Stack:   make([]byte, stackSize),
		heap:    make(map[uintptr]unsafe.Pointer, 4096),
		objects: make(map[uintptr]any),
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

func (s *_Stack) RegisterObject(obj any) uintptr {
	escapes(obj)

	rObj := reflect.ValueOf(obj)
	if rObj.Kind() != reflect.Pointer {
		panic(errors.New("not a pointer"))
	}

	addr := rObj.Elem().UnsafeAddr()

	for s.objects[addr] != nil {
		addr += 8
	}

	s.objects[addr] = obj

	return addr
}

func (s *_Stack) GetObject(addr uintptr) any {
	obj, ok := s.objects[addr]
	if !ok {
		err := fmt.Errorf("no object at addr %#x", addr)
		panic(err)
	}

	return obj
}

func (s *_Stack) ForgetObject(addr uintptr) {
	if _, ok := s.objects[addr]; !ok {
		err := fmt.Errorf("no object at addr %#x", addr)
		panic(err)
	}

	delete(s.objects, addr)
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
	n := unsafe.Sizeof(value)
	stack := tls.Alloc(int(n))
	memcpy(tls, stack, uintptr(unsafe.Pointer(&value)), size_t(n))

	runtime.KeepAlive(value)

	return alloc{Addr: stack, tls: tls, size: int(n)}
}

func derefToValue[T any](ptr uintptr) T {
	src := toSlice[T](ptr, 1)
	return src[0]
}
