package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/RetailPulse/modules/controller"
	"github.com/RetailPulse/modules/parser"
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
		return
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
	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		writeError(w, err)
		return
	}

	p, err := parser.Parse(b)
	if err != nil {
		fmt.Printf("Parsing error : %v", err)
		return
	}
	if p.Count != len(p.Visits) {
		writeError(w, fmt.Errorf("count not equal to visit length"))
		return
	}
	id := s.Controller.Add(p)
	var response = struct {
		JobId string `json:"job_id"`
	}{
		JobId: id,
	}
	w.WriteHeader(http.StatusCreated)
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
	res := struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write(marshall(res))
}

func marshall(i interface{}) []byte {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	return b
}
