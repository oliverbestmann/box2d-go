package b2

import "runtime"

// EnableConcurrency lets the world run its solver on multiple goroutines, using the task
// scheduler built into box2d. It does nothing on wasm or with a single CPU, where workers
// would only add overhead.
func EnableConcurrency(def *WorldDef) {
	if runtime.GOARCH == "wasm" || runtime.NumCPU() == 1 {
		return
	}

	def.WorkerCount = int32(min(runtime.NumCPU(), 8))
}
