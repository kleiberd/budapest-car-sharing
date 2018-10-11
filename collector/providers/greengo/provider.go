package greengo

import (
	"budapest-car-sharing-backend/collector/providers"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	providerName = "greengo"
	endpointURL  = "https://www.greengo.hu/divcontent.php?rnd={{RAND}}&funct=callAPI&APIname=getVehicleList&params[P_ICON_SIZE]=48&_={{TIMESTAMP}}"
	referer      = "https://www.greengo.hu/"
)

var header = map[string]string{
	"X-Requested-With": "XMLHttpRequest",
}

type provider struct {
	client *providers.Client
}

func NewProvider() *provider {
	return &provider{client: providers.NewClient(generateURL(), referer, header)}
}

func (p *provider) GetVehicles() ([]providers.Vehicle, error) {
	var responseData []Vehicle

	response, err := p.client.SendRequest()

	err = json.Unmarshal([]byte(response), &responseData)

	if nil != err {
		return nil, err
	}

	return NewTransformer(responseData).Transform()
}

func generateURL() string {
	source := rand.NewSource(time.Now().UnixNano())
	timestamp := fmt.Sprintf("%d", time.Now().UnixNano())

	url := strings.Replace(endpointURL, "{{RAND}}", strconv.FormatFloat(rand.New(source).Float64(), 'f', 16, 64), -1)
	url = strings.Replace(url, "{{TIMESTAMP}}", timestamp[0:len(timestamp)-6], -1)

	return url
}
