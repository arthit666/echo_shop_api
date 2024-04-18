package service

import (
	entities "github.com/arthit666/shop_api/entities"
	_adminModel "github.com/arthit666/shop_api/pkg/admin/model"
	_adminRepository "github.com/arthit666/shop_api/pkg/admin/repository"
	_customerModel "github.com/arthit666/shop_api/pkg/customer/model"
	_customerRepository "github.com/arthit666/shop_api/pkg/customer/repository"
)

type googleOAuth2Service struct {
	customerRepository _customerRepository.CustomerRepository
	adminRepository    _adminRepository.AdminRepository
}

func NewGoogleOAuth2Service(
	customerRepository _customerRepository.CustomerRepository,
	adminRepository _adminRepository.AdminRepository,
) OAuth2Service {
	return &googleOAuth2Service{
		customerRepository: customerRepository,
		adminRepository:    adminRepository,
	}
}

func (s *googleOAuth2Service) CustomerAccountCreating(customerCreatingReq *_customerModel.CustomerCreatingReq) error {
	if !s.IsCustomerExists(customerCreatingReq.ID) {
		customerEntity := &entities.Customer{
			ID:     customerCreatingReq.ID,
			Email:  customerCreatingReq.Email,
			Name:   customerCreatingReq.Name,
			Avatar: customerCreatingReq.Avatar,
		}

		if _, err := s.customerRepository.Creating(customerEntity); err != nil {
			return err
		}
	}

	return nil
}

func (s *googleOAuth2Service) AdminAccountCreating(adminCreatingInfoReq *_adminModel.AdminCreatingReq) error {
	if !s.IsAdminExists(adminCreatingInfoReq.ID) {
		adminEntity := &entities.Admin{
			ID:     adminCreatingInfoReq.ID,
			Email:  adminCreatingInfoReq.Email,
			Name:   adminCreatingInfoReq.Name,
			Avatar: adminCreatingInfoReq.Avatar,
		}

		if _, err := s.adminRepository.Creating(adminEntity); err != nil {
			return err
		}
	}

	return nil

}

func (s *googleOAuth2Service) IsCustomerExists(palyerId string) bool {
	player, err := s.customerRepository.FindByID(palyerId)
	if err != nil {
		return false
	}

	return player != nil
}

func (s *googleOAuth2Service) IsAdminExists(adminId string) bool {
	admin, err := s.adminRepository.FindByID(adminId)
	if err != nil {
		return false
	}

	return admin != nil
}
