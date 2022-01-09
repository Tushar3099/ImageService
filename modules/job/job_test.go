package job

import (
	"fmt"
	"testing"
	"time"

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

func TestJob(t *testing.T) {
	collection := make(map[string]*Job)
	for _, tst := range test {
		j := New(&tst)
		collection[j.Id] = j
		go j.Execute()
	}

	for i := 0; i < 30; i++ {
		for _, job := range collection {
			fmt.Printf("id : %v , State : %v ,  Error : %+v\n", job.Id, job.State, job.Errors)
		}

		fmt.Printf("\n\n")
		time.Sleep(time.Second * 1)
	}
}
