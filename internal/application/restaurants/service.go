package restaurant

import (
	"context"
	"sync"

	addressrepository "food-delivery/internal/application/address/repository"
	"food-delivery/internal/application/restaurants/dtos"
	"food-delivery/internal/application/restaurants/repository"
)

type service struct {
	Repo        repository.Repo
	Mutex       *sync.Mutex
	AddressRepo addressrepository.Repo
}

func (s *service) AddRestaurant(ctx context.Context, req *dtos.AddRestaurantRequest) (*dtos.Restaurant, error) {
	return nil, nil
}

func (s *service) GetAllRestaurants(ctx context.Context) ([]*dtos.Restaurant, error) {
	return nil, nil
}

func (s *service) GetRestaurantByID(ctx context.Context, id string) (*dtos.Restaurant, error) {
	return nil, nil
}

func (s *service) GetRestaurantByName(ctx context.Context, name string) (*dtos.Restaurant, error) {
	return nil, nil
}

func NewService(
	repo repository.Repo,
	addressRepo addressrepository.Repo,
	mutex *sync.Mutex,
) IRestaurant {
	return &service{
		AddressRepo: addressRepo,
		Repo:        repo,
		Mutex:       mutex,
	}
}
