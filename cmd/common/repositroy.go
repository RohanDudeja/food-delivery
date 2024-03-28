package common

import (
	address "food-delivery/internal/application/address/repository"
	restaurants "food-delivery/internal/application/restaurants/repository"
	users "food-delivery/internal/application/users/repository"
)

type Repository struct {
	Restaurant restaurants.Repo
	Users      users.Repo
	Address    address.Repo
}

func bindRepo(app *App) error {
	app.Repository.Restaurant = restaurants.NewRepo(app.Storage.DB)
	app.Repository.Users = users.NewRepo(app.Storage.DB)
	app.Repository.Address = address.NewRepo(app.Storage.DB)
	return nil
}
