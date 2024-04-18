package server

import (
	"github.com/arthit666/shop_api/pkg/productManaging/controller"
	_productManagingRepo "github.com/arthit666/shop_api/pkg/productManaging/repository"
	"github.com/arthit666/shop_api/pkg/productManaging/service"
	_productShopRepo "github.com/arthit666/shop_api/pkg/productShop/repository"
)

func (s *echoServer) initProductManagingRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/product_managing")
	shopRepo := _productShopRepo.NewProductShopRepositoryImpl(s.db, s.app.Logger)
	managingRepo := _productManagingRepo.NewProductManagingRepositoryImpl(s.db, s.app.Logger)
	ser := service.NewProductManagingService(managingRepo, shopRepo)
	con := controller.NewProductManagingController(ser)

	router.POST("", con.Creating, m.AdminAuthorizing)
	router.PATCH("/:productId", con.Editing, m.AdminAuthorizing)
	router.DELETE("/:productId", con.Archiving, m.AdminAuthorizing)
}
