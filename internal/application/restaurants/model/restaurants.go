package model

import "github.com/jinzhu/gorm"

const restaurants = "restaurants"

type Restaurant struct {
	ID        string `json:"id" gorm:"id"`
	Name      string `json:"name" gorm:"name"`
	AddressID string `json:"address_id" gorm:"address_id"`
	gorm.Model
}

func (r *Restaurant) TableName() string {
	return restaurants
}
