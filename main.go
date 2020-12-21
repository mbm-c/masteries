package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/maneeshbabu/masteries/pkg/worker"
	"github.com/maneeshbabu/masteries/pkg/worker/jobs"
)

func main() {
	fmt.Println("Starting")
	dummyWorker := worker.Start("dummy-worker", 3)

	for i := 0; i < 1000; i++ {
		dummyWorker <- &worker.Worker{Job: Job(strconv.Itoa(i))}
	}

	time.Sleep(500 * time.Minute)
}

func Job(id string) jobs.IJob {
	return &LocalJob{Id: id}
}

type LocalJob struct {
	Id string
	jobs.Job
}

func (j *LocalJob) ID() string {
	return j.Id
}

func (j *LocalJob) Run() error {
	rnd := rand.Intn(60)
	fmt.Println("Executing ", j.Id, " for ", rnd, " seconds")
	time.Sleep(time.Duration(rnd) * time.Second)
	fmt.Println("Completed", j.Id)
	return nil
}

func (j *LocalJob) Status() string {
	fmt.Sprintln("jobs#Status status of the job", j.State)
	return j.State
}

func (j *LocalJob) SetState(state string) {
	j.State = state
}
