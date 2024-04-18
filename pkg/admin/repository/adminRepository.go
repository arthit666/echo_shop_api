package repository

import (
	entities "github.com/arthit666/shop_api/entities"
)

type AdminRepository interface {
	Creating(adminEntity *entities.Admin) (string, error)
	FindByID(adminID string) (*entities.Admin, error)
}
