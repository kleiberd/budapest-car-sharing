package greengo

import (
	"budapest-car-sharing-backend/collector/providers"
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

type Provider struct {
	client *providers.Client
}

func NewProvider() *Provider {
	return &Provider{client: providers.NewClient(generateURL())}
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

	err = json.NewDecoder(resp.Body).Decode(&response)

	for _, value := range response {
		transformedVehicle, _ := Transform(value)
		vehicles = append(vehicles, transformedVehicle)
	}

	return vehicles, err
}

func generateURL() string {
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
