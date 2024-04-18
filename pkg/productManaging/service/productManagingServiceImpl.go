package service

import (
	"github.com/arthit666/shop_api/entities"
	_productManagingModel "github.com/arthit666/shop_api/pkg/productManaging/model"
	_productManagingRepo "github.com/arthit666/shop_api/pkg/productManaging/repository"
	_productShopModel "github.com/arthit666/shop_api/pkg/productShop/model"
	_productShopRepo "github.com/arthit666/shop_api/pkg/productShop/repository"
)

type productManagingServiceImpl struct {
	productManagingRepo _productManagingRepo.ProductManagingRepository
	productShopRepo     _productShopRepo.ProductShopRepository
}

func NewProductManagingService(managingRepo _productManagingRepo.ProductManagingRepository, shopRepo _productShopRepo.ProductShopRepository) ProductManagingService {
	return &productManagingServiceImpl{productManagingRepo: managingRepo, productShopRepo: shopRepo}
}

func (s *productManagingServiceImpl) Creating(productCreateReq *_productManagingModel.ProductCreatingReq) (*_productShopModel.Product, error) {
	product := &entities.Product{
		AdminID:     &productCreateReq.AdminID,
		Name:        productCreateReq.Name,
		Description: productCreateReq.Description,
		Picture:     productCreateReq.Picture,
		Price:       productCreateReq.Price,
	}

	itemEntity, err := s.productManagingRepo.Creating(product)
	if err != nil {
		return nil, err
	}

	return itemEntity.ToProductModel(), nil
}

func (s *productManagingServiceImpl) Editing(productID uint64, itemEditingReq *_productManagingModel.ProductEditingReq) (*_productShopModel.Product, error) {
	updatedProductID, err := s.productManagingRepo.Editing(productID, itemEditingReq)
	if err != nil {
		return nil, err
	}

	productEntity, err := s.productShopRepo.FindByID(updatedProductID)
	if err != nil {
		return nil, err
	}

	return productEntity.ToProductModel(), nil
}

func (s *productManagingServiceImpl) Archiving(itemID uint64) error {
	return s.productManagingRepo.Archiving(itemID)
}
