package blinkee

type Response struct {
	Data Data `json:"data"`
}

type Data struct {
	Items []Vehicle `json:"items"`
}

type Vehicle struct {
	ID       int      `json:"id"`
	Position Position `json:"position"`
	Type     string   `json:"type"`
}

type Position struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
