package model

import (
	"food-delivery/internal/application/address/dtos"

	"github.com/jinzhu/gorm"
)

type (
	Address struct {
		ID      string `json:"id" gorm:"id"`
		Line1   string `json:"line1" gorm:"line1"`
		Line2   string `json:"line2" gorm:"line2"`
		City    string `json:"city" gorm:"city"`
		State   string `json:"state" gorm:"state"`
		Country string `json:"country" gorm:"country"`
		gorm.Model
	}
)

func (a *Address) TableName() string {
	return "addresses"
}

func (a *Address) ToDto() *dtos.Address {
	return &dtos.Address{
		Line1:   a.Line1,
		Line2:   a.Line2,
		City:    a.City,
		State:   a.State,
		Country: a.Country,
	}
}
