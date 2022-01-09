package process

import (
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
	client := &http.Client{}
	req, _ := http.NewRequest("GET", p.Url, nil)
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("Accept-Language", "en-US")
	req.Header.Set("User-Agent", "Mozilla/5.0")
	response, err := client.Do(req)
	// response, err := client.Get(p.Url)
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
	// fmt.Printf("Img Width : %v, Img Height : %v\n", b.Max.X, b.Max.Y)
	// fmt.Printf("Img Width : %v, Img Height : %v", img.Height, img.Width)
	time.Sleep(time.Millisecond * 4)

	return nil
}
