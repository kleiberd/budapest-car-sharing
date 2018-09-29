package providers

type Vehicle struct {
	ExternalID string   `json:"externalId"`
	Provider   string   `json:"provider"`
	Position   Position `json:"position"`
}

type Position struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
