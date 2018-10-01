package providers

type Vehicle struct {
	ExternalID string  `json:"external_id"`
	Provider   string  `json:"provider"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Type       string  `json:"type"`
	FuelType   string  `json:"fuel_type"`
	Brand      string  `json:"brand"`
	Model      string  `json:"model"`
	Plate      string  `json:"plate"`
	Range      int     `json:"range"`
}
