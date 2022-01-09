package parser

import (
	"encoding/json"
	"fmt"
)

type Visit struct {
	StoreId string   `json:"store_id"`
	URLs    []string `json:"image_url"`
}

type ParsedData struct {
	Count  int     `json:"count"`
	Visits []Visit `json:"visits"`
}

func Parse(byts []byte) (*ParsedData, error) {
	if byts == nil {
		return nil, fmt.Errorf("bytes is nil")
	}
	var data ParsedData
	err := json.Unmarshal(byts, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
