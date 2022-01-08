package job

import (
	"github.com/RetailPulse/modules/parser"
	"github.com/RetailPulse/modules/store"
	"github.com/RetailPulse/types"
)

type State int64

const (
	FailedState State = iota
	SuccessState
	OngoingState
)

type Job struct {
	Count  int
	Stores []*store.Store
	State  State
	ResCh  chan types.Error
	Errors []types.Error
}

func New(data *parser.ParsedData) *Job {
	var j Job
	for _, d := range data.Visits {
		j.Stores = append(j.Stores, store.New(d.StoreId, d.URLs))
	}
	j.Count = len(data.Visits)
	j.State = OngoingState
	return &j
}

func (j *Job) Execute() {
	go func() {
		for _, s := range j.Stores {
			go s.Execute(j.ResCh)
		}
	}()

	for i := 0; i < j.Count; i++ {
		res, ok := <-j.ResCh
		if !ok {
			panic("res Chanel is Closed")
		}
		if res.StoreId != "" {
			j.State = FailedState
			j.Errors = append(j.Errors, res)
		}
	}
	close(j.ResCh)
	if j.State != FailedState {
		j.State = SuccessState
	}
}
