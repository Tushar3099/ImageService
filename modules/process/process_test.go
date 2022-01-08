package process

import (
	"fmt"
	"testing"
)

type Test struct {
	url string
}

var test = []Test{
	{
		url: "https://www.gstatic.com/webp/gallery/4.jpg",
	},
	{
		url: "https://www.gstatic.com/webp/gallery/1.jpg",
	},
	{
		url: "https://www.gstatic.com/webp/gallery/2.jpg",
	},
	{
		url: "https://www.gstatic.com/webp/gallery/3.jpg",
	},
	{
		url: "https://www.gstatic.com/webp/gallery/5.jpg",
	},
}

func TestProcess(t *testing.T) {

	for _, tst := range test {
		p := New(tst.url)
		err := p.Execute()
		if err != nil {
			fmt.Printf("Error found : %v", err.Error())
			t.Fail()
			return
		}
		// fmt.Printf("Img Width : %v, Img Height : %v", img.Max.X, img.Max.Y)
	}
}
