package greengo

import (
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

func (t *Transformer) Transform() ([]p.Vehicle, error) {
	var vehicles []p.Vehicle

	for _, value := range t.responseData {
		vehicles = append(vehicles, t.transformItem(value))
	}

	return vehicles, nil
}

func (t *Transformer) transformItem(vehicle Vehicle) p.Vehicle {
	latitude, _ := strconv.ParseFloat(vehicle.GpsLat, 64)
	longitude, _ := strconv.ParseFloat(vehicle.GpsLong, 64)
	estimatedRange, _ := strconv.ParseInt(vehicle.EstimatedKm, 10, 2)

	transformedVehicle := p.Vehicle{
		ExternalID: vehicle.VehicleID,
		Provider:   providerName,
		Latitude:   latitude,
		Longitude:  longitude,
		Type:       p.CAR,
		FuelType:   p.ELECTRIC,
		Brand:      vehicle.MakeDesc,
		Model:      p.EUP,
		Plate:      vehicle.PlateNumber,
		Range:      int(estimatedRange),
	}

	return transformedVehicle
}
