package server

import (
	_adminRepository "github.com/arthit666/shop_api/pkg/admin/repository"
	_customerRepository "github.com/arthit666/shop_api/pkg/customer/repository"
	_oauth2Controller "github.com/arthit666/shop_api/pkg/oauth2/controller"
	_oauth2Service "github.com/arthit666/shop_api/pkg/oauth2/service"
)

func (s *echoServer) initOAuth2Router() {
	router := s.app.Group("/v1/oauth2/google")

	playerRepository := _customerRepository.NewCustomerRepositoryImpl(s.db, s.app.Logger)
	adminRepository := _adminRepository.NewAdminRepositoryImpl(s.db, s.app.Logger)

	oauth2Service := _oauth2Service.NewGoogleOAuth2Service(
		playerRepository,
		adminRepository,
	)

	oauth2Controller := _oauth2Controller.NewGoogleOAuth2Controller(
		oauth2Service,
		s.conf.OAuth2,
		s.app.Logger,
	)

	router.GET("/customer/login", oauth2Controller.CustomerLogin)
	router.GET("/admin/login", oauth2Controller.AdminLogin)
	router.GET("/customer/login/callback", oauth2Controller.CustomerLoginCallback)
	router.GET("/admin/login/callback", oauth2Controller.AdminLoginCallback)
	router.DELETE("/logout", oauth2Controller.Logout)
}
