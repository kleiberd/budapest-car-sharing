package greengo

import (
	p "budapest-car-sharing-backend/collector/providers"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	providerName = "greengo"
	endpointURL  = "https://www.greengo.hu/divcontent.php?rnd={{RAND}}&funct=callAPI&APIname=getVehicleList&params[P_ICON_SIZE]=48&_={{TIMESTAMP}}"
	referer      = "https://www.greengo.hu/"
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

	c := p.NewClient(generateUrl())

	req, err := c.GetRequest(getHeader())
	if err != nil {
		return nil, err
	}

	resp, err := c.SendRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&vehicles)

	return vehicles, err
}

func generateUrl() string {
	source := rand.NewSource(time.Now().UnixNano())
	timestamp := fmt.Sprintf("%d", time.Now().UnixNano())

	url := strings.Replace(endpointURL, "{{RAND}}", strconv.FormatFloat(rand.New(source).Float64(), 'f', 16, 64), -1)
	url = strings.Replace(url, "{{TIMESTAMP}}", timestamp[0:len(timestamp)-6], -1)

	return url
}

func getHeader() *http.Header {
	header := http.Header{}
	header.Set("X-Requested-With", "XMLHttpRequest")
	header.Set("Referer", referer)

	return &header
}
