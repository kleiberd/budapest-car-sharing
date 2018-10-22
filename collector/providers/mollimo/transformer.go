package mollimo

import (
	"budapest-car-sharing-backend/collector/domain"
	p "budapest-car-sharing-backend/collector/providers"
)

var modelMap = map[int]map[string]string{
	14: {
		p.Brand:    "Volkswagen",
		p.Model:    p.EUp,
		p.FuelType: p.Electric,
	},
	15: {
		p.Brand:    "Volkswagen",
		p.Model:    "Up",
		p.FuelType: p.Petrol,
	},
	18: {
		p.Brand:    "Mercedes",
		p.Model:    "A200",
		p.FuelType: p.Petrol,
	},
}

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
		ExternalID: vehicle.Description.ID,
		Provider:   providerName,
		Latitude:   vehicle.Location.Position.Lat,
		Longitude:  vehicle.Location.Position.Lon,
		Type:       p.Car,
		FuelType:   modelMap[vehicle.Description.ModelID][p.FuelType],
		Brand:      modelMap[vehicle.Description.ModelID][p.Brand],
		Model:      modelMap[vehicle.Description.ModelID][p.Model],
		Plate:      vehicle.Description.Name[0:6],
		Range:      vehicle.Status.EnergyLevel,
	}

	return transformedVehicle
}
