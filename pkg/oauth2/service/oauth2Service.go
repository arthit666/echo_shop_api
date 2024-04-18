package service

import (
	_adminModel "github.com/arthit666/shop_api/pkg/admin/model"
	_customerModel "github.com/arthit666/shop_api/pkg/customer/model"
)

type OAuth2Service interface {
	CustomerAccountCreating(playerCreatingReq *_customerModel.CustomerCreatingReq) error
	AdminAccountCreating(adminCreatingInfoReq *_adminModel.AdminCreatingReq) error
	IsCustomerExists(playerID string) bool
	IsAdminExists(adminID string) bool
}
