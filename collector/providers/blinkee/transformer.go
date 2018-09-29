package blinkee

import (
	p "budapest-car-sharing-backend/collector/providers"
	"fmt"
)

func Transform(vehicle Vehicle) (p.Vehicle, error) {
	transformedVehicle := p.Vehicle{
		ExternalID: fmt.Sprintf("%d", vehicle.ID),
		Provider:   providerName,
		Position: p.Position{
			Latitude:  vehicle.Position.Lat,
			Longitude: vehicle.Position.Lng,
		},
	}

	return transformedVehicle, nil
}
