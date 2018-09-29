package blinkee

import (
	p "budapest-car-sharing-backend/collector/providers"
	"encoding/json"
	"net/http"
)

const (
	providerName = "blinkee"
	endpointURL  = "https://blinkee.city/api/regions/11/vehicles"
	referer      = "https://blinkee.city/hu/"
)

func Provide() ([]p.Vehicle, error) {
	data, err := getResponseData()
	if nil != err {
		return nil, err
	}

	var vehicles []p.Vehicle
	for _, value := range data {
		transformedVehicle, _ := Transform(value)
		vehicles = append(vehicles, transformedVehicle)
	}

	return vehicles, nil
}

func getResponseData() ([]Vehicle, error) {
	var response Response

	c := p.NewClient(endpointURL)

	req, err := c.GetRequest(getHeader())
	if err != nil {
		return nil, err
	}

	resp, err := c.SendRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)

	return response.Data.Items, err
}

func getHeader() *http.Header {
	header := http.Header{}
	header.Set("Referer", referer)

	return &header
}
