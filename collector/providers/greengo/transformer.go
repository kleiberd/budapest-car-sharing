package greengo

import (
	p "budapest-car-sharing-backend/collector/providers"
	"strconv"
)

func Transform(vehicle Vehicle) (p.Vehicle, error) {
	latitude, _ := strconv.ParseFloat(vehicle.GpsLat, 64)
	longitude, _ := strconv.ParseFloat(vehicle.GpsLong, 64)

	transformedVehicle := p.Vehicle{
		ExternalID: vehicle.VehicleID,
		Provider:   providerName,
		Position: p.Position{
			Latitude:  latitude,
			Longitude: longitude,
		},
	}

	return transformedVehicle, nil
}
