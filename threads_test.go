package b2

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Ported from box2d test/test_thread.c

func TestSemaphoreInitialCount(t *testing.T) {
	s := __builtin_b2CreateSemaphore(nil, 3)
	defer __builtin_b2DestroySemaphore(nil, s)

	// these do not block, the count starts at 3
	for range 3 {
		__builtin_b2WaitSemaphore(nil, s)
	}

	__builtin_b2SignalSemaphore(nil, s)
	__builtin_b2WaitSemaphore(nil, s)
}

func TestThreadSignalAndJoin(t *testing.T) {
	sem := __builtin_b2CreateSemaphore(nil, 0)
	defer __builtin_b2DestroySemaphore(nil, sem)

	var value int
	worker := threadFunction(func(_ *_Stack, _ uintptr) {
		value = 99
		__builtin_b2SignalSemaphore(nil, sem)
	})

	// __ccgo_fp hides the closure from the compiler and GC: move it to the heap and keep it alive
	escapes(worker)
	defer runtime.KeepAlive(worker)

	thread := __builtin_b2CreateThread(nil, __ccgo_fp(worker), 0, 0)
	__builtin_b2WaitSemaphore(nil, sem)
	assert.Equal(t, 99, value)

	__builtin_b2JoinThread(nil, thread)
}

func TestThreadMutex(t *testing.T) {
	mutex := __builtin_b2CreateMutex(nil)
	defer __builtin_b2DestroyMutex(nil, mutex)

	var sum int
	worker := threadFunction(func(_ *_Stack, _ uintptr) {
		for range 1000 {
			__builtin_b2LockMutex(nil, mutex)
			sum++
			__builtin_b2UnlockMutex(nil, mutex)
		}
	})

	// __ccgo_fp hides the closure from the compiler and GC: move it to the heap and keep it alive
	escapes(worker)
	defer runtime.KeepAlive(worker)

	var threads []uintptr
	for range 4 {
		threads = append(threads, __builtin_b2CreateThread(nil, __ccgo_fp(worker), 0, 0))
	}

	for _, thread := range threads {
		__builtin_b2JoinThread(nil, thread)
	}

	assert.Equal(t, 4000, sum)
}

// Ported from TestSetWorkerCount in box2d test/test_world.c
func TestWorld_SetWorkerCount(t *testing.T) {
	def := DefaultWorldDef()
	def.WorkerCount = 1
	world, _ := newTestWorld(t, def)

	for idx := range 50 {
		createBox(world, Vec2{X: float32(idx%10) - 5, Y: float32(idx / 10)}, DefaultShapeDef())
	}

	for _, workers := range []int32{4, 4, 0, -5, _B2_MAX_WORKERS + 10, 2} {
		world.SetWorkerCount(workers)
		assert.EqualValues(t, min(max(workers, 1), _B2_MAX_WORKERS), world.GetWorkerCount())

		for range 30 {
			world.Step(1.0/60.0, 4)
		}
	}
}
