package repository

import (
	"github.com/arthit666/shop_api/databases"
	entities "github.com/arthit666/shop_api/entities"
	"github.com/arthit666/shop_api/pkg/customer/exception"

	"github.com/labstack/echo/v4"
)

type CustomerRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewCustomerRepositoryImpl(db databases.Database, logger echo.Logger) CustomerRepository {
	return &CustomerRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *CustomerRepositoryImpl) Creating(CustomerEntity *entities.Customer) (*entities.Customer, error) {
	insertedCustomer := &entities.Customer{}

	if err := r.db.Connect().Create(CustomerEntity).Scan(insertedCustomer).Error; err != nil {
		r.logger.Error("Creating Customer failed", err.Error())
		return nil, &exception.CustomerCreating{}
	}

	return insertedCustomer, nil
}

func (r *CustomerRepositoryImpl) FindByID(CustomerID string) (*entities.Customer, error) {
	Customer := &entities.Customer{}

	if err := r.db.Connect().Where("id = ?", CustomerID).First(Customer).Error; err != nil {
		r.logger.Errorf("Finding Customer failed: %s", err.Error())
		return nil, &exception.CustomerNotFound{CustomID: CustomerID}
	}

	return Customer, nil
}
