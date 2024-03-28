package repository

import (
	"context"

	"food-delivery/internal/application/address/model"

	"github.com/jinzhu/gorm"
)

type (
	Repo interface {
		GetAddressById(ctx context.Context, id string) (*model.Address, error)
	}
	restaurantRepo struct {
		DB *gorm.DB
	}
)

func (r *restaurantRepo) GetAddressById(ctx context.Context, id string) (*model.Address, error) {
	return nil, nil
}

func NewRepo(db *gorm.DB) Repo {
	return &restaurantRepo{
		DB: db,
	}
}
