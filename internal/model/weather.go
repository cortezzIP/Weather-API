package model

import "encoding/json"

type Weather struct {
	Location `json:"location"`
	Current  `json:"current"`
}

type Location struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Current struct {
	Condition   `json:"condition"`
	TempCelsius float32 `json:"temp_c"`
}

type Condition struct {
	Text string `json:"text"`
}

func (w Weather) MarshalBinary() ([]byte, error) {
	return json.Marshal(w)
}

func (w Weather) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &w)
}