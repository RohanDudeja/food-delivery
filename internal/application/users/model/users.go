package model

import (
	address "food-delivery/internal/application/address/dtos"
	"food-delivery/internal/application/users/dtos"

	"github.com/jinzhu/gorm"
)

const users = "users"

type Users struct {
	ID        string `json:"id" gorm:"id"`
	Name      string `json:"name" gorm:"name"`
	AddressID string `json:"address_id" gorm:"address_id"`
	gorm.Model
}

func (u *Users) TableName() string {
	return users
}

func (u *Users) ToDto(address *address.Address) *dtos.Users {
	return &dtos.Users{
		ID:      u.ID,
		Name:    u.Name,
		Address: address,
	}
}
