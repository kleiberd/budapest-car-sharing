package infrastucture

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/gobuffalo/packr"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-config"
	"github.com/rubenv/sql-migrate"
	"time"
)

const (
	Dialect        = "mysql"
	MigrationPath  = "./migrations"
	MigrationTable = "migrations"
)

type Database struct {
	Config     mysql.Config
	Connection *gorm.DB
}

func NewDatabase() *Database {
	databaseConfig := mysql.Config{
		User:   config.Get("database", "user").String("root"),
		Passwd: config.Get("database", "password").String("root"),
		Addr:   config.Get("database", "host").String("127.0.0.1"),
		DBName: config.Get("database", "db").String("budapest-car-sharing"),
		Net:    "tcp",
		Loc:    time.Local,
		Params: map[string]string{
			"parseTime":            "true",
			"allowNativePasswords": "true",
		},
	}

	databaseMigration(databaseConfig)
	db := databaseConnection(databaseConfig)

	return &Database{
		Config:     databaseConfig,
		Connection: db,
	}
}

func databaseConnection(config mysql.Config) *gorm.DB {
	db, err := gorm.Open(Dialect, config.FormatDSN())
	if err != nil {
		panic("failed to connect databaseInit")
	}

	return db
}

func databaseMigration(config mysql.Config) {
	migrate.SetTable(MigrationTable)
	migrations := &migrate.PackrMigrationSource{
		Box: packr.NewBox(MigrationPath),
	}

	db, err := sql.Open(Dialect, config.FormatDSN())
	if err != nil {
		fmt.Println(err)
	}

	n, err := migrate.Exec(db, Dialect, migrations, migrate.Up)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
