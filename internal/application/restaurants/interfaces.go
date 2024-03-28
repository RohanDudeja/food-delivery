package restaurant

import (
	"context"

	"food-delivery/internal/application/restaurants/dtos"
)

type IRestaurant interface {
	AddRestaurant(ctx context.Context, req *dtos.AddRestaurantRequest) (*dtos.Restaurant, error)
	GetAllRestaurants(ctx context.Context) ([]*dtos.Restaurant, error)
	GetRestaurantByID(ctx context.Context, id string) (*dtos.Restaurant, error)
	GetRestaurantByName(ctx context.Context, name string) (*dtos.Restaurant, error)
}
