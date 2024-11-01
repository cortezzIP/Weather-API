package model

type Weather struct {
	Location `json:"location"`
	Current  `json:"current"`
}

type Location struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Current struct {
	TempCelsius float32 `json:"temp_c"`
}
