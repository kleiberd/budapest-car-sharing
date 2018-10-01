package main

import (
	"budapest-car-sharing-backend/collector/providers/blinkee"
	"budapest-car-sharing-backend/collector/providers/greengo"
	"budapest-car-sharing-backend/collector/providers/mollimo"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	spew.Dump(mollimo.NewProvider().GetVehicles())
	spew.Dump(greengo.NewProvider().GetVehicles())
	spew.Dump(blinkee.NewProvider().GetVehicles())
}
