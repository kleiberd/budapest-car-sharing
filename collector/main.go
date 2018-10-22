package main

import (
	"budapest-car-sharing-backend/collector/infrastucture"
	"budapest-car-sharing-backend/collector/services"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/env"
)

func main() {
	config.Load(env.NewSource())

	database := infrastucture.NewDatabase()
	defer database.Connection.Close()

	services.NewCollector(database).Collect()
}
