package blinkee

import (
	p "budapest-car-sharing-backend/collector/providers"
	"fmt"
)

func Transform(vehicle Vehicle) (p.Vehicle, error) {
	transformedVehicle := p.Vehicle{
		ExternalID: fmt.Sprintf("%d", vehicle.ID),
		Provider:   providerName,
		Latitude:   vehicle.Position.Lat,
		Longitude:  vehicle.Position.Lng,
		Type:       vehicle.Type,
		FuelType:   p.ELECTRIC,
		Brand:      "Blinkee",
		Model:      "Scooter",
	}

	return transformedVehicle, nil
}
