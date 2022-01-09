package service

import (
	"encoding/json"
	"net/http"

	"github.com/RetailPulse/modules/controller"
	"github.com/RetailPulse/types"
)

type Service struct {
	Controller *controller.Controller
}

func New(c *controller.Controller) *Service {
	var s Service
	s.Controller = c
	return &s
}

func (s *Service) GetJob(w http.ResponseWriter, r *http.Request) {
	jobid := r.URL.Query().Get("jobid")
	st, err := s.Controller.StateById(jobid)
	if err != nil {
		writeError(w, err)
	}
	errors, _ := s.Controller.ErrorById(jobid)
	s.Controller.StateById(jobid)

	var response struct {
		Status string        `json:"status"`
		JobId  string        `json:"job_id"`
		Error  []types.Error `json:"error"`
	}
	response.Status = string(st)
	response.JobId = jobid
	response.Error = errors

	w.Write(marshall(response))
}

func (s *Service) PostJob(w http.ResponseWriter, r *http.Request) {
	var response = struct {
		Message string `json:"message"`
	}{
		Message: "Posting Job",
	}

	w.Write(marshall(response))
}

func (s *Service) GetIndex(w http.ResponseWriter, r *http.Request) {

	var response = struct {
		Message string `json:"message"`
	}{
		Message: "Welcome to Home Page",
	}

	w.Write(marshall(response))
}

func writeError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write(marshall(err.Error()))
}

func marshall(i interface{}) []byte {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	return b
}
