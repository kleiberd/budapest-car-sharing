package mollimo

import (
	p "budapest-car-sharing-backend/collector/providers"
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
	var vehicles []Vehicle

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

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	bodyString = strings.Replace(bodyString, "window.cars = ", "", -1)

	err = json.Unmarshal([]byte(bodyString), &vehicles)

	return vehicles, err
}

func getHeader() *http.Header {
	header := http.Header{}
	header.Set("Referer", referer)

	return &header
}
