package repository

import (
	"github.com/arthit666/shop_api/databases"
	"github.com/arthit666/shop_api/entities"
	"github.com/arthit666/shop_api/pkg/admin/exception"
	"github.com/labstack/echo/v4"
)

type adminRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewAdminRepositoryImpl(db databases.Database, logger echo.Logger) AdminRepository {
	return &adminRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *adminRepositoryImpl) Creating(adminEntity *entities.Admin) (string, error) {
	tx := r.db.Connect().Create(adminEntity)

	if tx.Error != nil {
		r.logger.Errorf("Error inserting admin: %s", tx.Error.Error())
		return "", &exception.AdminCreating{}
	}

	return adminEntity.ID, nil
}

func (r *adminRepositoryImpl) FindByID(adminID string) (*entities.Admin, error) {
	admin := &entities.Admin{}

	if err := r.db.Connect().Where("id = ?", adminID).First(admin).Error; err != nil {
		r.logger.Errorf("Error finding admin: %s", err.Error())
		return nil, &exception.AdminNotFound{AdminID: adminID}
	}

	return admin, nil
}
