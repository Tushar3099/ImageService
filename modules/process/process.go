package process

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"time"
)

type Process struct {
	Url    string
	Result int
}

func New(url string) *Process {
	var p Process
	p.Url = url
	return &p
}

func (p *Process) Execute() error {

	// Execute the process

	response, err := http.Get(p.Url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	img, _, err := image.Decode(response.Body)
	if err != nil {
		return err
	}

	b := img.Bounds()
	p.Result = 2 * (b.Max.X + b.Max.Y)
	fmt.Printf("Img Width : %v, Img Height : %v\n", b.Max.X, b.Max.Y)
	// fmt.Printf("Img Width : %v, Img Height : %v", img.Height, img.Width)
	time.Sleep(time.Millisecond * 5)

	return nil
}
