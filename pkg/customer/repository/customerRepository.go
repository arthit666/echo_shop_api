package repository

import entities "github.com/arthit666/shop_api/entities"

type CustomerRepository interface {
	Creating(playerEntity *entities.Customer) (*entities.Customer, error)
	FindByID(playerID string) (*entities.Customer, error)
}
