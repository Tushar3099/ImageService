package store

import (
	"github.com/RetailPulse/modules/process"
	"github.com/RetailPulse/types"
)

type Store struct {
	StoreId   string
	Processes []*process.Process
}

func New(id string, urls []string) *Store {
	var s Store
	s.StoreId = id
	for _, url := range urls {
		s.Processes = append(s.Processes, process.New(url))
	}

	return &s
}

func (s *Store) Execute(jobCh chan<- types.Error) {

	for _, p := range s.Processes {
		err := p.Execute()
		if err != nil {
			jobCh <- types.Error{StoreId: s.StoreId, Message: err.Error()}
			return
		}
	}
	jobCh <- types.Error{}
}
