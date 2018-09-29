package main

import (
	"budapest-car-sharing-backend/collector/providers/blinkee"
	"budapest-car-sharing-backend/collector/providers/greengo"
	"budapest-car-sharing-backend/collector/providers/mollimo"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	data, err := greengo.Provide()

	spew.Dump(data, err)

	data, err = mollimo.Provide()

	spew.Dump(data, err)

	data, err = blinkee.Provide()

	spew.Dump(data, err)
}
