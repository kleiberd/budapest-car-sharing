package mollimo

import (
	p "budapest-car-sharing-backend/collector/providers"
)

func Transform(vehicle Vehicle) (p.Vehicle, error) {
	transformedVehicle := p.Vehicle{
		ExternalID: vehicle.Description.ID,
		Provider:   providerName,
		Position: p.Position{
			Latitude:  vehicle.Location.Position.Lat,
			Longitude: vehicle.Location.Position.Lon,
		},
	}

	return transformedVehicle, nil
}
