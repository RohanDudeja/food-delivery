package common

import (
	restaurant "food-delivery/internal/application/restaurants"
	"food-delivery/internal/application/users"
)

type Services struct {
	Restaurant restaurant.IRestaurant
	Users      users.IUsers
}

func bindServices(app *App) error {
	app.Services.Restaurant = restaurant.NewService(app.Repository.Restaurant, app.Repository.Address, NewMutex())
	app.Services.Users = users.NewService(app.Repository.Users, app.Repository.Address, NewMutex())
	return nil
}
