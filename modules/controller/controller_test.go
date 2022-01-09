package controller

import (
	"fmt"
	"testing"
	"time"

	"github.com/RetailPulse/modules/job"
	"github.com/RetailPulse/modules/parser"
)

var test = []parser.ParsedData{
	{
		Visits: []parser.Visit{
			{
				StoreId: "store1",
				URLs: []string{"https://www.gstatic.com/webp/gallery/4.jpg",
					"https://www.gstatic.com/webp/gallery/1.jpg",
					"https://www.gstatic.com/webp/gallery/2.jpg",
					"https://www.gstatic.com/webp/gallery/3.jpg",
				},
			},
			{
				StoreId: "store2",
				URLs: []string{
					"https://www.gstatic.com/webp/gallery/5.jpg",
					"https://www.gstatic.com/webp/gallery/6.jpg",
					"https://www.gstatic.com/webp/gallery/7.jpg",
					"https://www.gstatic.com/webp/gallery/8.jpg",
				},
			},
			{
				StoreId: "store3",
				URLs: []string{
					"https://www.gstatic.com/webp/gallery/5.jpg",
					"https://www.gstatic.com/webp/gallery/2.jpg",
					"https://www.gstatic.com/webp/gallery/1.jpg",
					"https://www.gstatic.com/webp/gallery/7.jpg",
				},
			},
		},
	},
	{
		Visits: []parser.Visit{
			{
				StoreId: "store1",
				URLs: []string{"https://www.gstatic.com/webp/gallery/4.jpg",
					"https://www.gstatic.com/webp/gallery/1.jpg",
					"https://www.gstatic.com/webp/gallery/2.jpg",
					"https://www.gstatic.com/webp/gallery/3.jpg",
				},
			},
			{
				StoreId: "store2",
				URLs: []string{"https://www.gstatic.com/webp/gallery/4.jpg",
					"https://www.gstatic.com/webp/gallery/1.jpg",
					"https://www.gstatic.com/webp/gallery/2.jpg",
					"https://www.gstatic.com/webp/gallery/3.jpg",
				},
			},
			{
				StoreId: "store3",
				URLs: []string{"https://www.gstatic.com/webp/gallery/4.jpg",
					"https://www.gstatic.com/webp/gallery/1.jpg",
					"https://www.gstatic.com/webp/gallery/2.jpg",
					"https://www.gstatic.com/webp/gallery/3.jpg",
				},
			},
		},
	},
}

func TestController(t *testing.T) {
	c := New()
	var jobs []string
	for _, tst := range test {
		jobs = append(jobs, c.Add(&tst))
	}

	for {
		var hasOngoing bool
		for _, j := range jobs {
			st, err := c.StateById(j)
			if err != nil {
				t.Fail()
			}
			errors, _ := c.ErrorById(j)
			if st == job.OngoingState {
				hasOngoing = true
			}
			fmt.Printf("id : %v , State : %v ,  Error : %+v\n", j, st, errors)
		}
		if !hasOngoing {
			break
		}
		fmt.Printf("\n\n")
		time.Sleep(time.Second * 1)
	}
}
