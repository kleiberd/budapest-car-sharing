package providers

import "budapest-car-sharing-backend/collector/domain"

type Provider interface {
	GetVehicles() ([]domain.Vehicle, error)
}
