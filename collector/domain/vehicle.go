package domain

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Vehicle struct {
	ExternalID string    `json:"external_id" gorm:"primary_key"`
	Provider   string    `json:"provider" gorm:"primary_key"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	Type       string    `json:"type"`
	FuelType   string    `json:"fuel_type"`
	Brand      string    `json:"brand"`
	Model      string    `json:"model"`
	Plate      string    `json:"plate"`
	Range      int       `json:"range"`
	CreatedAt  time.Time `json:"created_at"`
}

type VehicleRepository struct {
	db *gorm.DB
}

func NewVehicleRepository(db *gorm.DB) *VehicleRepository {
	return &VehicleRepository{
		db: db,
	}
}

func (r *VehicleRepository) StoreAll(vehicles []Vehicle) error {
	tx := r.db.Begin()
	defer func() {
		if rec := recover(); rec != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	for _, vehicle := range vehicles {
		if err := tx.Create(&vehicle).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (r *VehicleRepository) DeleteAll() {
	r.db.Delete(Vehicle{})
}
