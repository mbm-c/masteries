package jobs

import (
	"fmt"
)

type State string

const (
	RUNNING   State = "RUNNING"
	FAILED    State = "FAILED"
	COMPLETED State = "COMPLETED"
)

// Job ...
type Job struct {
	Id    string
	State string
}

type IJob interface {
	ID() string
	Run() error
	Status() string
	SetState(string)
}

func (j *Job) ID() string {
	return j.Id
}

// Run ...
func (j *Job) Run() error {
	fmt.Sprintln("jobs#Run executing job", j.Id)
	return nil
}

// Status ...
func (j *Job) Status() string {
	fmt.Sprintln("jobs#Status status of the job", j.State)
	return j.State
}

func (j *Job) SetState(state string) {
	j.State = state
}
