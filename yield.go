//go:build !race

package b2

import "runtime"

// sched_yield is b2Yield, which box2d calls in spin loops while it waits for another worker.
// Like the syscall it must not sleep. runtime.Gosched costs far more than the syscall when
// there are idle Ps, so it keeps spinning and only yields every yieldEvery calls. The yield
// matters when there are more workers than GOMAXPROCS. wasm has a single thread, where
// spinning can't help and Gosched is cheap, so it always yields there.
func sched_yield(tls *_Stack) int32 {
	tls.yields++
	if runtime.GOARCH == "wasm" || tls.yields%yieldEvery == 0 {
		runtime.Gosched()
	} else {
		pause()
	}

	return 0
}

// yieldEvery is about 0.2 ms of spinning on arm64 and longer on amd64, where PAUSE is slower.
const yieldEvery = 4096
