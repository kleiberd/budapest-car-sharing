package blinkee

import (
	"budapest-car-sharing-backend/collector/domain"
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

func (t *Transformer) Transform() ([]domain.Vehicle, error) {
	var vehicles []domain.Vehicle

	for _, value := range t.responseData {
		vehicles = append(vehicles, t.transformItem(value))
	}

	return vehicles, nil
}

func (t *Transformer) transformItem(vehicle Vehicle) domain.Vehicle {
	transformedVehicle := domain.Vehicle{
		ExternalID: fmt.Sprintf("%d", vehicle.ID),
		Provider:   providerName,
		Latitude:   vehicle.Position.Lat,
		Longitude:  vehicle.Position.Lng,
		Type:       vehicle.Type,
		FuelType:   p.Electric,
		Brand:      "Blinkee",
		Model:      "Scooter",
	}

	return transformedVehicle
}
