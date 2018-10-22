package greengo

import (
	"budapest-car-sharing-backend/collector/domain"
	p "budapest-car-sharing-backend/collector/providers"
	"strconv"
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
	latitude, _ := strconv.ParseFloat(vehicle.GpsLat, 64)
	longitude, _ := strconv.ParseFloat(vehicle.GpsLong, 64)
	estimatedRange, _ := strconv.Atoi(vehicle.EstimatedKm)

	transformedVehicle := domain.Vehicle{
		ExternalID: vehicle.VehicleID,
		Provider:   providerName,
		Latitude:   latitude,
		Longitude:  longitude,
		Type:       p.Car,
		FuelType:   p.Electric,
		Brand:      vehicle.MakeDesc,
		Model:      p.EUp,
		Plate:      vehicle.PlateNumber,
		Range:      estimatedRange,
	}

	return transformedVehicle
}
