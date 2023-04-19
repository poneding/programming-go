package task

type Task interface {
	Do()
}

type taskQueue chan Task

type worker struct {
	tq taskQueue
}

func newWorker() worker {
	return worker{tq: make(taskQueue)}
}

func (w worker) work(wq chan taskQueue) {
	go func() {
		for {
			wq <- w.tq
			select {
			case job := <-w.tq:
				job.Do()
			}
		}
	}()
}

type taskPool struct {
	routines int
	tq       taskQueue
	wq       chan taskQueue
}

func NewTaskPool(routines int) *taskPool {
	return &taskPool{
		routines: routines,
		tq:       make(taskQueue),
		wq:       make(chan taskQueue, routines),
	}
}
func (pool *taskPool) Open() {
	for i := 0; i < pool.routines; i++ {
		worker := newWorker()
		worker.work(pool.wq)
	}

	go func() {
		for {
			select {
			case task := <-pool.tq:
				// Acquire a worker.
				worker := <-pool.wq
				// Assign task to a worker's task queue.
				worker <- task
			}
		}
	}()
}

func (pool *taskPool) Accept(t Task) {
	pool.tq <- t
}

func (pool *taskPool) Close() {
	pool = nil
}
