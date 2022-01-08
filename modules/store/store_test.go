package store

import (
	"fmt"
	"testing"

	"github.com/RetailPulse/modules/process"
	"github.com/RetailPulse/types"
)

var test = []Store{
	{
		StoreId: "store1",
		Processes: []*process.Process{
			{
				Url: "https://www.gstatic.com/webp/gallery/4.jpg",
			},
			{
				Url: "https://www.gstatic.com/webp/gallery/1.jpg",
			},
			{
				Url: "https://www.gstatic.com/webp/gallery/2.jpg",
			},
			{
				Url: "https://www.gstatic.com/webp/gallery/3.jpg",
			},
		},
	},
	{
		StoreId: "store2",
		Processes: []*process.Process{
			{
				Url: "https://www.gstatic.com/webp/gallery/5.jpg",
			},
			{
				Url: "https://www.gstatic.com/webp/gallery/6.jpg",
			},
			{
				Url: "https://www.gstatic.com/webp/gallery/7.jpg",
			},
			{
				Url: "https://www.gstatic.com/webp/gallery/8.jpg",
			},
		},
	},
	// {
	// 	StoreId: "store3",
	// 	Processes: []*process.Process{
	// 		{
	// 			Url: "https://www.gstatic.com/webp/gallery/9.jpg",
	// 		},
	// 		{
	// 			Url: "https://www.gstatic.com/webp/gallery/10.jpg",
	// 		},
	// 		{
	// 			Url: "https://www.gstatic.com/webp/gallery/11.jpg",
	// 		},
	// 		{
	// 			Url: "https://www.gstatic.com/webp/gallery/12.jpg",
	// 		},
	// 	},
	// },
}

func TestStore(t *testing.T) {
	var Errors []types.Error
	resCh := make(chan types.Error)
	go func() {
		for _, s := range test {
			go s.Execute(resCh)
		}
	}()

	for i := 0; i < len(test); i++ {
		res, ok := <-resCh
		if !ok {
			panic("res Chanel is Closed")
		}

		if res.StoreId != "" {
			fmt.Printf("Found Err : %v\n", res.StoreId)
			Errors = append(Errors, res)
		}
	}
	fmt.Printf("%+v", Errors)
	close(resCh)
}
