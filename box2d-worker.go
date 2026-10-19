package b2

import (
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"
)

// EnableConcurrency lets the world run its solver on multiple goroutines.
// It does nothing on wasm or with a single CPU, where workers would only add overhead.
func EnableConcurrency(def *WorldDef) {
	if runtime.GOARCH == "wasm" || runtime.NumCPU() == 1 {
		return
	}

	startWorker()

	def.WorkerCount = 8
	def.EnqueueTask = __ccgo_fp(enqueueTask)
	def.FinishTask = __ccgo_fp(finishTask)
}

type task struct {
	remainingSteps atomic.Int64
	doneSignal     chan struct{}

	taskCallback uintptr
	taskContext  uintptr
	itemCount    int32
	minRange     int32

	parentStack *_Stack
}

func (t *task) Await() {
	<-t.doneSignal
}

type taskStep struct {
	task  *task
	start int32
	end   int32
}

var taskStepQueue = make(chan *taskStep, 1024)

var startWorker = sync.OnceFunc(func() {

	for workerId := range 8 {
		go worker(uint32(workerId))
	}
})

func worker(workerId uint32) {
	stack := newStack(1024 * 64)

	for {
		step := <-taskStepQueue
		work(stack, workerId, step)
	}
}

func work(stack *_Stack, workerId uint32, step *taskStep) {
	defer stack.CheckEmpty()

	task := step.task

	// int startIndex, int endIndex, uint32_t workerIndex, void* taskContext
	type TaskCallback func(*_Stack, int32, int32, uint32, uintptr)
	type FunctionPtr struct{ uintptr }

	(*(*TaskCallback)(unsafe.Pointer(&FunctionPtr{task.taskCallback})))(stack, step.start, step.end, workerId, task.taskContext)

	if task.remainingSteps.Add(-1) == 0 {
		close(task.doneSignal)
	}
}

func enqueueTask(tls *_Stack, taskCallback uintptr, itemCount int32, minRange int32, taskContext uintptr, userContext uintptr) uintptr {
	task := &task{
		taskCallback: taskCallback,
		taskContext:  taskContext,
		itemCount:    itemCount,
		minRange:     minRange,
	}

	task.doneSignal = make(chan struct{})

	var steps []*taskStep

	var startAt = int32(0)
	for startAt < itemCount {
		step := &taskStep{task: task, start: startAt}
		step.end = min(startAt+minRange, itemCount)

		startAt += minRange

		steps = append(steps, step)
	}

	task.remainingSteps.Store(int64(len(steps)))
	if len(steps) == 0 {
		// nothing to do, finishTask must not block
		close(task.doneSignal)
	}

	for _, step := range steps {
		taskStepQueue <- step
	}

	return tls.RegisterObject(task)
}

func finishTask(tls *_Stack, userTask uintptr, userContext uintptr) {
	task := tls.GetObject(userTask).(*task)

	// remove from heap
	tls.ForgetObject(userTask)

	// wait for userTask to finish before return
	task.Await()
}
