package mollimo

type Description struct {
	CityID   int    `json:"cityId"`
	IconsURL string `json:"iconsUrl"`
	ID       string `json:"id"`
	Model    string `json:"model"`
	ModelID  int    `json:"modelId"`
	Name     string `json:"name"`
	Plate    string `json:"string"`
}

type Location struct {
	Address  Address  `json:"address"`
	Position Position `json:"position"`
}

type Address struct {
	StreetAddress string `json:"streetAddress"`
}

type Position struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Status struct {
	EnergyLevel int `json:"energyLevel"`
}

type Vehicle struct {
	Description Description `json:"description"`
	Location    Location    `json:"location"`
	Status      Status      `json:"status"`
}
