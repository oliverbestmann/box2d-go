//go:build race

package b2

import "time"

// sched_yield is b2Yield, see yield.go. With the race detector, the atomic loads in the spin
// loops of box2d are so slow that spinning starves the worker it waits for, sleeping is faster.
func sched_yield(_ *_Stack) int32 {
	time.Sleep(20 * time.Microsecond)
	return 0
}
