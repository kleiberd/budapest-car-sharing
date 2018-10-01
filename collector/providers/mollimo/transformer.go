package mollimo

import (
	p "budapest-car-sharing-backend/collector/providers"
)

func Transform(vehicle Vehicle) (p.Vehicle, error) {
	modelMap := map[int]map[string]string{
		14: {
			p.BRAND:    "Volkswagen",
			p.MODEL:    p.EUP,
			p.FUELTYPE: p.ELECTRIC,
		},
		15: {
			p.BRAND:    "Volkswagen",
			p.MODEL:    "Up",
			p.FUELTYPE: p.PETROL,
		},
		18: {
			p.BRAND:    "Mercedes",
			p.MODEL:    "A200",
			p.FUELTYPE: p.PETROL,
		},
	}

	transformedVehicle := p.Vehicle{
		ExternalID: vehicle.Description.ID,
		Provider:   providerName,
		Latitude:   vehicle.Location.Position.Lat,
		Longitude:  vehicle.Location.Position.Lon,
		Type:       p.CAR,
		FuelType:   modelMap[vehicle.Description.ModelID][p.FUELTYPE],
		Brand:      modelMap[vehicle.Description.ModelID][p.BRAND],
		Model:      modelMap[vehicle.Description.ModelID][p.MODEL],
		Plate:      vehicle.Description.Name,
		Range:      vehicle.Status.EnergyLevel,
	}

	return transformedVehicle, nil
}
