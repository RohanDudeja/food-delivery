package repository

import (
	"context"

	"food-delivery/internal/application/restaurants/model"

	"github.com/jinzhu/gorm"
)

type (
	Repo interface {
		GetAllRestaurants(ctx context.Context) ([]*model.Restaurant, error)
	}
	restaurantRepo struct {
		DB *gorm.DB
	}
)

func (r *restaurantRepo) GetAllRestaurants(ctx context.Context) ([]*model.Restaurant, error) {
	return nil, nil
}

func NewRepo(db *gorm.DB) Repo {
	return &restaurantRepo{
		DB: db,
	}
}
