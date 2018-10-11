package blinkee

import (
	p "budapest-car-sharing-backend/collector/providers"
	"fmt"
)

type Transformer struct {
	responseData []Vehicle
}

func NewTransformer(responseData []Vehicle) *Transformer {
	return &Transformer{
		responseData: responseData,
	}
}

func (t *Transformer) Transform() ([]p.Vehicle, error) {
	var vehicles []p.Vehicle

	for _, value := range t.responseData {
		vehicles = append(vehicles, t.transformItem(value))
	}

	return vehicles, nil
}

func (t *Transformer) transformItem(vehicle Vehicle) p.Vehicle {
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

	return transformedVehicle
}
