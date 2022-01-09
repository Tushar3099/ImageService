package store

import (
	"fmt"
	"testing"

	"github.com/RetailPulse/types"
)

type Test struct {
	id   string
	urls []string
}

var test = []Test{
	{
		id: "store1",
		urls: []string{"https://www.gstatic.com/webp/gallery/4.jpg",
			"https://www.gstatic.com/webp/gallery/1.jpg",
			"https://www.gstatic.com/webp/gallery/2.jpg",
			"https://www.gstatic.com/webp/gallery/3.jpg",
		},
	},
	{
		id: "store2",
		urls: []string{
			"https://www.gstatic.com/webp/gallery/5.jpg",
			"https://www.gstatic.com/webp/gallery/6.jpg",
			"https://www.gstatic.com/webp/gallery/7.jpg",
			"https://www.gstatic.com/webp/gallery/8.jpg",
		},
	},
	{
		id: "store3",
		urls: []string{
			"https://www.gstatic.com/webp/gallery/5.jpg",
			"https://www.gstatic.com/webp/gallery/2.jpg",
			"https://www.gstatic.com/webp/gallery/1.jpg",
			"https://www.gstatic.com/webp/gallery/7.jpg",
		},
	},
}

func TestStore(t *testing.T) {
	var Errors []types.Error
	resCh := make(chan types.Error)
	go func() {
		for _, s := range test {
			st := New(s.id, s.urls)
			go st.Execute(resCh)
		}
	}()

	for i := 0; i < len(test); i++ {
		res, ok := <-resCh
		if !ok {
			panic("res Chanel is Closed")
		}

		if res.StoreId != "" {
			// fmt.Printf("Found Err : %v\n", res.StoreId)
			Errors = append(Errors, res)
		}
	}
	// time.Sleep(time.Second * 10)
	fmt.Printf("%+v\n", Errors)
	close(resCh)
}
