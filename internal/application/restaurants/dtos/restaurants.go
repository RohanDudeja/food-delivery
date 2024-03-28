package dtos

import (
	address "food-delivery/internal/application/address/dtos"
)

type Restaurant struct {
	ID      string          `json:"id"`
	Name    string          `json:"name"`
	Address address.Address `json:"address"`
}
