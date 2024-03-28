package users

import (
	"context"
	"log"
	"sync"

	addressrepository "food-delivery/internal/application/address/repository"
	"food-delivery/internal/application/users/dtos"
	"food-delivery/internal/application/users/repository"
)

type service struct {
	Repo        repository.Repo
	AddressRepo addressrepository.Repo
	Mutex       *sync.Mutex
}

func (s *service) GetUserById(ctx context.Context, id string) (*dtos.Users, error) {
	user, err := s.Repo.GetUserById(ctx, id)
	if err != nil {
		log.Println("GetUserById repo error")
		return nil, err
	}
	address, err := s.AddressRepo.GetAddressById(ctx, user.AddressID)
	if err != nil {
		log.Println("GetAddressById repo error")
	}
	if address != nil {
		return user.ToDto(address.ToDto()), nil
	}
	return user.ToDto(nil), nil
}

func (s *service) AddUser(ctx context.Context, request *dtos.AddUserRequest) (*dtos.Users, error) {
	return nil, nil
}

func (s *service) DeleteUser(ctx context.Context, id string) error { return nil }

func NewService(
	repo repository.Repo,
	addressRepo addressrepository.Repo,
	mutex *sync.Mutex,
) IUsers {
	return &service{
		Repo:        repo,
		Mutex:       mutex,
		AddressRepo: addressRepo,
	}
}
