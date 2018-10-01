package providers

type Provider interface {
	GetVehicles([]Vehicle, error)
}
