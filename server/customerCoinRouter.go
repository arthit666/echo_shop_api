package server

import (
	_customerCoinController "github.com/arthit666/shop_api/pkg/customerCoin/controller"
	_customerCoinRepository "github.com/arthit666/shop_api/pkg/customerCoin/repository"
	_customerCoinService "github.com/arthit666/shop_api/pkg/customerCoin/service"
)

func (s *echoServer) initCustomerCoinRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/customer_coin")

	customerCoinRepository := _customerCoinRepository.NewCustomerCoinRepositoryImpl(s.db, s.app.Logger)

	customerCoinService := _customerCoinService.NewCustomerCoinServiceImpl(
		customerCoinRepository,
	)
	customerCoinController := _customerCoinController.NewCustomerCoinControllerImpl(customerCoinService)

	router.POST("", customerCoinController.CoinAdding, m.AdminAuthorizing)
	router.GET("", customerCoinController.CustomerCoinShowing, m.CustomerAuthorizing)
}
