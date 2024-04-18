package controller

import (
	"context"
	"net/http"

	"github.com/arthit666/shop_api/pkg/custom"
	_oauth2 "github.com/arthit666/shop_api/pkg/oauth2/exception"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

func (c *googleOAuth2Controller) CustomerAuthorizing(e echo.Context, next echo.HandlerFunc) error {
	ctx := context.Background()

	tokenSource, err := c.getToken(e)
	if err != nil {
		c.logger.Errorf("Error reading token: %s", err.Error())
		return custom.Error(
			e, http.StatusUnauthorized,
			&_oauth2.Unauthorized{},
		)

	}

	// Validate the token
	if !tokenSource.Valid() {
		c.logger.Errorf("Token is not valid")

		// Refresh the token
		tokenSource, err = c.customerTokenRefreshing(e, tokenSource)
		if err != nil {
			c.logger.Errorf("Error refreshing token: %s", err.Error())
			return custom.Error(e, http.StatusUnauthorized, err)
		}
	}

	// Get user info
	client := customGoogleOAuth2.Client(ctx, tokenSource)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		c.logger.Errorf("Error reading user info: %s", err.Error())
		return custom.Error(e, http.StatusUnauthorized, err)

	}

	if !c.oauth2Service.IsCustomerExists(userInfo.ID) {
		return custom.Error(e, http.StatusUnauthorized, &_oauth2.NoPermission{})
	}

	e.Set("playerID", userInfo.ID)

	return next(e)
}

func (c *googleOAuth2Controller) AdminAuthorizing(e echo.Context, next echo.HandlerFunc) error {
	ctx := context.Background()

	tokenSource, err := c.getToken(e)
	if err != nil {
		c.logger.Errorf("Error reading token: %s", err.Error())
		return custom.Error(e, http.StatusUnauthorized, err)

	}

	// Validate the token
	if !tokenSource.Valid() {
		c.logger.Errorf("Token is not valid")

		// Refresh the token
		tokenSource, err = c.adminTokenRefreshing(e, tokenSource)
		if err != nil {
			c.logger.Errorf("Error refreshing token: %s", err.Error())
			return custom.Error(e, http.StatusUnauthorized, err)
		}
	}

	// Get user info
	client := adminGoogleOAuth2.Client(ctx, tokenSource)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		c.logger.Errorf("Error reading user info: %s", err.Error())
		return custom.Error(e, http.StatusUnauthorized, &_oauth2.Unauthorized{})

	}

	if !c.oauth2Service.IsAdminExists(userInfo.ID) {
		return custom.Error(e, http.StatusUnauthorized, &_oauth2.NoPermission{})
	}

	e.Set("adminID", userInfo.ID)

	return next(e)
}

func (c *googleOAuth2Controller) customerTokenRefreshing(e echo.Context, token *oauth2.Token) (*oauth2.Token, error) {
	ctx := context.Background()

	updatedToken, err := customGoogleOAuth2.TokenSource(ctx, token).Token()
	if err != nil {
		c.logger.Errorf("Error refreshing token: %s", err.Error())
		return nil, &_oauth2.Unauthorized{}
	}

	// Update cookies
	c.setSameSiteCookie(e, accessTokenCookieName, updatedToken.AccessToken)
	c.setSameSiteCookie(e, refreshTokenCookieName, updatedToken.RefreshToken)

	return updatedToken, nil
}

func (c *googleOAuth2Controller) adminTokenRefreshing(e echo.Context, token *oauth2.Token) (*oauth2.Token, error) {
	ctx := context.Background()

	updatedToken, err := adminGoogleOAuth2.TokenSource(ctx, token).Token()
	if err != nil {
		c.logger.Errorf("Error refreshing token: %s", err.Error())
		return nil, &_oauth2.Unauthorized{}
	}

	// Update cookies
	c.setSameSiteCookie(e, accessTokenCookieName, updatedToken.AccessToken)
	c.setSameSiteCookie(e, refreshTokenCookieName, updatedToken.RefreshToken)

	return updatedToken, nil
}

func (c *googleOAuth2Controller) getToken(e echo.Context) (*oauth2.Token, error) {
	accessToken, err := e.Cookie(accessTokenCookieName)
	if err != nil {
		c.logger.Errorf("Error reading access token: %s", err.Error())
		return nil, &_oauth2.Unauthorized{}
	}

	refreshToken, err := e.Cookie(accessTokenCookieName)
	if err != nil {
		c.logger.Errorf("Error reading refresh token: %s", err.Error())
		return nil, &_oauth2.Unauthorized{}
	}

	return &oauth2.Token{
		AccessToken:  accessToken.Value,
		RefreshToken: refreshToken.Value,
	}, nil
}
