package users

import (
	"context"

	"food-delivery/internal/application/users/dtos"
)

type IUsers interface {
	GetUserById(ctx context.Context, id string) (*dtos.Users, error)
	AddUser(ctx context.Context, request *dtos.AddUserRequest) (*dtos.Users, error)
	DeleteUser(ctx context.Context, id string) error
}
