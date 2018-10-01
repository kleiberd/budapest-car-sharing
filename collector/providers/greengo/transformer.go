package greengo

import (
	p "budapest-car-sharing-backend/collector/providers"
	"strconv"
)

func Transform(vehicle Vehicle) (p.Vehicle, error) {
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

	return transformedVehicle, nil
}
