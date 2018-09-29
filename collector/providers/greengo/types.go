package greengo

type Vehicle struct {
	Address          string `json:"address"`
	BatteryLevel     string `json:"battery_level"`
	BatteryState     string `json:"battery_state"`
	BatteryStateDesc string `json:"battery_state_desc"`
	Color            string `json:"color"`
	ColorDesc        string `json:"color_desc"`
	Distance         string `json:"distance"`
	Doors            string `json:"doors"`
	EstimatedKm      string `json:"estimated_km"`
	GpsLat           string `json:"gps_lat"`
	GpsLong          string `json:"gps_long"`
	Icon             string `json:"icon"`
	Make             string `json:"make"`
	MakeDesc         string `json:"make_desc"`
	MaxRange         string `json:"max_range"`
	Model            string `json:"model"`
	PlateNumber      string `json:"plate_number"`
	Seats            string `json:"seats"`
	State            string `json:"state"`
	StateDesc        string `json:"state_desc"`
	VehicleID        string `json:"vehicle_id"`
}
