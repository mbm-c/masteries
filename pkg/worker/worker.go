package worker

import (
	"fmt"
	"time"

	"github.com/maneeshbabu/masteries/pkg/worker/jobs"
)

type Worker struct {
	Type string
	Job  interface{}
}

func (w *Worker) Work(id int, works <-chan *Worker) {
	for j := range works {
		fmt.Println("worker#Work worker", w.Type, id, "is executing job")
		if job, ok := j.Job.(jobs.IJob); ok {
			startedAt := time.Now()
			fmt.Println("worker#Work job", job.ID(), "started executing at", startedAt)
			err := job.Run()
			if err != nil {
				fmt.Println("worker#Work", job.ID(), "failed with err", err)
				job.SetState(string(jobs.FAILED))
			}
			fmt.Println("worker#Work", job.ID(), "completed")
			job.SetState(string(jobs.COMPLETED))
		} else {
			fmt.Println("Invalid job", j)
		}
	}
}

// Start a worker
func Start(workerType string, count int) chan *Worker {
	var workerChannel = make(chan *Worker)
	worker := Worker{Type: workerType}
	for w := 1; w <= count; w++ {
		go worker.Work(w, workerChannel)
	}
	return workerChannel
}
