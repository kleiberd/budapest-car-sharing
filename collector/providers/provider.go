package providers

type Provider interface {
	GetResponseData() ([]Vehicle, error)
	Transform(interface{}) (Vehicle, error)
}

func Provide(p *Provider) ([]Vehicle, error) {
	data, err := (*p).GetResponseData()
	if nil != err {
		return nil, err
	}

	var vehicles []Vehicle
	for _, value := range data {
		transformedVehicle, _ := (*p).Transform(value)
		vehicles = append(vehicles, transformedVehicle)
	}

	return vehicles, nil
}
