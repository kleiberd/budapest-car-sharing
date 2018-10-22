package services

import (
	"budapest-car-sharing-backend/collector/domain"
	"budapest-car-sharing-backend/collector/infrastucture"
	"budapest-car-sharing-backend/collector/providers"
	"budapest-car-sharing-backend/collector/providers/blinkee"
	"budapest-car-sharing-backend/collector/providers/greengo"
	"budapest-car-sharing-backend/collector/providers/mollimo"
	"fmt"
	"github.com/micro/go-config"
	"time"
)

type Collector struct {
	providers  []providers.Provider
	repository *domain.VehicleRepository
}

func NewCollector(database *infrastucture.Database) *Collector {
	return &Collector{
		providers: []providers.Provider{
			mollimo.NewProvider(),
			greengo.NewProvider(),
			blinkee.NewProvider(),
		},
		repository: domain.NewVehicleRepository(database.Connection),
	}
}

func (c *Collector) Collect() {
	for {
		fmt.Printf("Collected data at: %s", time.Now().Format(time.RFC3339))

		c.doCollect()
		time.Sleep(config.Get("collector", "sleep").Duration(1) * time.Minute)
	}
}

func (c *Collector) doCollect() {
	var vehicles []domain.Vehicle

	for _, provider := range c.providers {
		collected, err := provider.GetVehicles()
		if nil != err {
			fmt.Println(err)
		}
		vehicles = append(vehicles, collected...)
	}

	c.repository.DeleteAll()
	c.repository.StoreAll(vehicles)
}
