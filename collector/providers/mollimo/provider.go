package mollimo

import (
	"budapest-car-sharing-backend/collector/providers"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	providerName = "mollimo"
	endpointURL  = "https://www.mollimo.hu/data/cars.js"
	referer      = "https://www.mollimo.hu/"
)

type Provider struct {
	client *providers.Client
}

func NewProvider() *Provider {
	return &Provider{client: providers.NewClient(endpointURL)}
}

func (p *Provider) GetVehicles() ([]providers.Vehicle, error) {
	var response []Vehicle
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

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	bodyString = strings.Replace(bodyString, "window.cars = ", "", -1)

	err = json.Unmarshal([]byte(bodyString), &response)

	for _, value := range response {
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
