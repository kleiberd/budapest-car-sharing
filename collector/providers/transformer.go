package providers

import "budapest-car-sharing-backend/collector/domain"

type Transformer interface {
	Transform() ([]domain.Vehicle, error)
}
