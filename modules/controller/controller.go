package controller

import (
	"fmt"

	"github.com/RetailPulse/modules/job"
	"github.com/RetailPulse/modules/parser"
	"github.com/RetailPulse/types"
)

type Controller struct {
	collection map[string]*job.Job
}

func New() *Controller {
	var c Controller
	c.collection = make(map[string]*job.Job)
	return &c
}

func (c *Controller) Add(p *parser.ParsedData) string {
	j := job.New(p)
	c.collection[j.Id] = j
	go j.Execute()
	return j.Id
}

func (c *Controller) StateById(jobId string) (job.State, error) {
	if job, ok := c.collection[jobId]; ok {
		return job.State, nil
	}
	return "", fmt.Errorf("job not found")
}

func (c *Controller) ErrorById(jobId string) ([]types.Error, error) {
	if job, ok := c.collection[jobId]; ok {
		return job.Errors, nil
	}
	return []types.Error{}, fmt.Errorf("job not found")
}
