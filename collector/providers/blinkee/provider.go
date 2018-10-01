package blinkee

import (
	"budapest-car-sharing-backend/collector/providers"
	"encoding/json"
	"net/http"
)

const (
	providerName = "blinkee"
	endpointURL  = "https://blinkee.city/api/regions/11/vehicles"
	referer      = "https://blinkee.city/hu/"
)

type Provider struct {
	client *providers.Client
}

func NewProvider() *Provider {
	return &Provider{client: providers.NewClient(endpointURL)}
}

func (p *Provider) GetVehicles() ([]providers.Vehicle, error) {
	var response Response
	var vehicles []providers.Vehicle

	req, err := p.client.GetRequest(getHeader())
	if err != nil {
		return nil, err
	}

	resp, err := p.client.SendRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)

	for _, value := range response.Data.Items {
		transformedVehicle, _ := Transform(value)
		vehicles = append(vehicles, transformedVehicle)
	}

	return vehicles, err
}

func getHeader() *http.Header {
	header := http.Header{}
	header.Set("Referer", referer)

	return &header
}
