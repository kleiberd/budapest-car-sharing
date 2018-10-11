package mollimo

import (
	"budapest-car-sharing-backend/collector/providers"
	"encoding/json"
	"strings"
)

const (
	providerName = "mollimo"
	endpointURL  = "https://www.mollimo.hu/data/cars.js"
	referer      = "https://www.mollimo.hu/"
)

type provider struct {
	client *providers.Client
}

func NewProvider() *provider {
	return &provider{client: providers.NewClient(endpointURL, referer, nil)}
}

func (p *provider) GetVehicles() ([]providers.Vehicle, error) {
	var responseData []Vehicle

	responseBytes, _ := p.client.SendRequest()
	responseString := string(responseBytes)
	responseString = strings.Replace(responseString, "window.cars = ", "", -1)

	err := json.Unmarshal([]byte(responseString), &responseData)
	if nil != err {
		return nil, err
	}

	return NewTransformer(responseData).Transform()
}
