package blinkee

import (
	"budapest-car-sharing-backend/collector/domain"
	"budapest-car-sharing-backend/collector/providers"
	"encoding/json"
)

const (
	providerName = "blinkee"
	endpointURL  = "https://blinkee.city/api/regions/11/vehicles"
	referer      = "https://blinkee.city/hu/"
)

type provider struct {
	client *providers.Client
}

func NewProvider() *provider {
	return &provider{client: providers.NewClient(endpointURL, referer, nil)}
}

func (p *provider) GetVehicles() ([]domain.Vehicle, error) {
	var responseData Response

	response, err := p.client.SendRequest()

	err = json.Unmarshal([]byte(response), &responseData)

	if nil != err {
		return nil, err
	}

	return NewTransformer(responseData.Data.Items).Transform()
}
