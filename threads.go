package b2

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

// Threading primitives of box2d (src/core.h), used by its built-in task scheduler and the
// recorder. A thread is a goroutine with its own stack. box2d only passes the handles around.

// handles keeps the Go objects behind the handles that box2d holds, for threads and for
// callback contexts like the DebugDraw in World.Draw.
var handles sync.Map

var lastHandle atomic.Uintptr

func newHandle(obj any) uintptr {
	h := lastHandle.Add(1)
	handles.Store(h, obj)
	return h
}

func handleOf[T any](h uintptr) T {
	obj, ok := handles.Load(h)
	if !ok {
		panic("unknown handle")
	}

	return obj.(T)
}

func __builtin_b2CreateMutex(_ *_Stack) uintptr {
	return newHandle(&sync.Mutex{})
}

func __builtin_b2DestroyMutex(_ *_Stack, m uintptr) {
	handles.Delete(m)
}

func __builtin_b2LockMutex(_ *_Stack, m uintptr) {
	handleOf[*sync.Mutex](m).Lock()
}

func __builtin_b2UnlockMutex(_ *_Stack, m uintptr) {
	handleOf[*sync.Mutex](m).Unlock()
}

type semaphore struct {
	mu    sync.Mutex
	cond  sync.Cond
	count int32
}

func __builtin_b2CreateSemaphore(_ *_Stack, initCount int32) uintptr {
	s := &semaphore{count: initCount}
	s.cond.L = &s.mu
	return newHandle(s)
}

func __builtin_b2DestroySemaphore(_ *_Stack, s uintptr) {
	handles.Delete(s)
}

func __builtin_b2WaitSemaphore(_ *_Stack, h uintptr) {
	s := handleOf[*semaphore](h)

	s.mu.Lock()
	defer s.mu.Unlock()

	for s.count == 0 {
		s.cond.Wait()
	}

	s.count--
}

func __builtin_b2SignalSemaphore(_ *_Stack, h uintptr) {
	s := handleOf[*semaphore](h)

	s.mu.Lock()
	s.count++
	s.mu.Unlock()

	s.cond.Signal()
}

func __builtin_b2CreateThread(_ *_Stack, function, context, _ uintptr) uintptr {
	type functionPtr struct{ uintptr }
	fn := *(*threadFunction)(unsafe.Pointer(&functionPtr{function}))

	done := make(chan struct{})

	go func() {
		defer close(done)

		stack := newStack(1024 * 64)
		fn(stack, context)
		stack.CheckEmpty()
	}()

	return newHandle(done)
}

func __builtin_b2JoinThread(_ *_Stack, t uintptr) {
	<-handleOf[chan struct{}](t)
	handles.Delete(t)
}
