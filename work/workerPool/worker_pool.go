package workerPool

import (
	"fmt"
)

type Task struct{ I int }

func (t *Task) Do() int {
	return t.I
}

type Result struct {
	Resp int
}

type WorkerPool struct {
	numberOfWorkers int
	respChan        chan Result
	inChans         []chan Task
}

func NewWorkerPool(numberOfWorkers int) *WorkerPool {
	respChan := make(chan Result, numberOfWorkers)
	return &WorkerPool{
		respChan:        respChan,
		numberOfWorkers: numberOfWorkers,
	}
}

func (w *WorkerPool) AddTasks(tasks []Task) <-chan Result {
	chunks := chunkSlice(tasks, len(tasks)/w.numberOfWorkers)
	for k, v := range chunks {
		go worker(k, v, w.respChan)
	}
	return w.respChan
}

func chunkSlice(slice []Task, chunkSize int) [][]Task {
	var chunks [][]Task
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func worker(i int, tasks []Task, respChan chan<- Result) {
	for _, task := range tasks {
		fmt.Printf("Worker number:%v\n", i)
		resp := task.Do()
		respChan <- Result{Resp: resp * 2}
	}
}
