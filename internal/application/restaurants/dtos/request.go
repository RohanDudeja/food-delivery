package dtos

import (
	address "food-delivery/internal/application/address/dtos"
)

type (
	AddRestaurantRequest struct {
		Name    string          `json:"name"`
		Address address.Address `json:"address"`
	}
)
