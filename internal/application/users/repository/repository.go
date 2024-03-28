package repository

import (
	"context"
	"errors"

	"food-delivery/internal/application/users/model"

	"github.com/jinzhu/gorm"
)

type (
	Repo interface {
		GetUserById(ctx context.Context, id string) (*model.Users, error)
	}
	usersRepo struct {
		DB *gorm.DB
	}
)

func (r *usersRepo) GetUserById(ctx context.Context, id string) (*model.Users, error) {
	var user *model.Users
	r.DB.Table("users").Where("id = ?", id).Find(&user).Limit(1)
	if user == nil {
		return nil, errors.New("invalid id")
	}
	return user, nil
}

func NewRepo(db *gorm.DB) Repo {
	return &usersRepo{
		DB: db,
	}
}
