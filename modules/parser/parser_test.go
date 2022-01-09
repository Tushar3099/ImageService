package parser

import (
	"fmt"
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	b, err := os.ReadFile("test.json")

	if err != nil {
		fmt.Printf("Read file error : %v", err)
		return
	}
	if b == nil {
		fmt.Println("file is empty")
		return
	}
	p, err := Parse(b)
	// var p ParsedData
	// err = json.Unmarshal(b, &p)

	if err != nil {
		fmt.Printf("Parsing error : %v", err)
		return
	}
	fmt.Printf("%+v", p)
}
