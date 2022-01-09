package job

import (
	"crypto/rand"
	"fmt"
	"log"

	"github.com/RetailPulse/modules/parser"
	"github.com/RetailPulse/modules/store"
	"github.com/RetailPulse/types"
)

type State string

const (
	FailedState  State = "failed"
	SuccessState State = "success"
	OngoingState State = "ongoing"
)

type Job struct {
	Id     string
	Count  int
	Stores []*store.Store
	State  State
	ResCh  chan types.Error
	Errors []types.Error
}

func New(data *parser.ParsedData) *Job {
	var j Job
	j.Id = generateuuId()
	j.ResCh = make(chan types.Error)
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

	for i := 0; i < len(j.Stores); i++ {
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

func generateuuId() string {
	b := make([]byte, 4)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x", b[0:2], b[2:])
	return uuid
}
