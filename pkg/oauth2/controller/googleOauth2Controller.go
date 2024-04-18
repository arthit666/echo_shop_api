package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"sync"
	"time"

	_adminModel "github.com/arthit666/shop_api/pkg/admin/model"
	"github.com/arthit666/shop_api/pkg/custom"
	_customerModel "github.com/arthit666/shop_api/pkg/customer/model"
	"github.com/arthit666/shop_api/pkg/oauth2/exception"
	_oauth2Model "github.com/arthit666/shop_api/pkg/oauth2/model"
	_oauth2Service "github.com/arthit666/shop_api/pkg/oauth2/service"
	"github.com/avast/retry-go"

	"github.com/arthit666/shop_api/config"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type googleOAuth2Controller struct {
	oauth2Service _oauth2Service.OAuth2Service
	oauth2Conf    *config.OAuth2
	logger        echo.Logger
}

var (
	customGoogleOAuth2 *oauth2.Config
	adminGoogleOAuth2  *oauth2.Config
	once               sync.Once

	accessTokenCookieName  = "act"
	refreshTokenCookieName = "rft"
	stateCookieName        = "state"

	letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func NewGoogleOAuth2Controller(
	oauth2Service _oauth2Service.OAuth2Service,
	oauth2Conf *config.OAuth2,
	logger echo.Logger,
) OAuth2Controller {
	once.Do(func() {
		setGooleOAuth2Config(oauth2Conf)
	})

	return &googleOAuth2Controller{
		oauth2Service: oauth2Service,
		oauth2Conf:    oauth2Conf,
		logger:        logger,
	}
}

func setGooleOAuth2Config(oauth2Conf *config.OAuth2) {
	customGoogleOAuth2 = &oauth2.Config{
		ClientID:     oauth2Conf.ClientId,
		ClientSecret: oauth2Conf.ClientSecret,
		RedirectURL:  oauth2Conf.CustomerRedirectUrl,
		Scopes:       oauth2Conf.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:       oauth2Conf.Endpoints.AuthUrl,
			TokenURL:      oauth2Conf.Endpoints.TokenUrl,
			DeviceAuthURL: oauth2Conf.Endpoints.DeviceAuthUrl,
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}

	adminGoogleOAuth2 = &oauth2.Config{
		ClientID:     oauth2Conf.ClientId,
		ClientSecret: oauth2Conf.ClientSecret,
		RedirectURL:  oauth2Conf.AdminRedirectUrl,
		Scopes:       oauth2Conf.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:       oauth2Conf.Endpoints.AuthUrl,
			TokenURL:      oauth2Conf.Endpoints.TokenUrl,
			DeviceAuthURL: oauth2Conf.Endpoints.DeviceAuthUrl,
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}
}

func (c *googleOAuth2Controller) CustomerLogin(e echo.Context) error {
	state := c.randomState()
	c.setCookie(e, stateCookieName, state)
	return e.Redirect(http.StatusFound, customGoogleOAuth2.AuthCodeURL(state))
}

func (c *googleOAuth2Controller) AdminLogin(e echo.Context) error {
	state := c.randomState()
	c.setCookie(e, stateCookieName, state)
	return e.Redirect(http.StatusFound, adminGoogleOAuth2.AuthCodeURL(state))
}

func (c *googleOAuth2Controller) CustomerLoginCallback(e echo.Context) error {
	ctx := context.Background()

	if err := retry.Do(func() error {
		return c.callbackValidating(e)
	}, retry.Attempts(3), retry.Delay(3*time.Second)); err != nil {
		c.logger.Errorf("Error validating callback: %s", err.Error())
		return custom.Error(e, http.StatusBadRequest, err)
	}

	token, err := customGoogleOAuth2.Exchange(ctx, e.QueryParam("code"))
	if err != nil {
		c.logger.Errorf("Error exchanging code for token: %s", err.Error())
		return custom.Error(e, http.StatusBadRequest, &exception.Unauthorized{})
	}

	client := customGoogleOAuth2.Client(ctx, token)

	customerInfo, err := c.getUserInfo(client)
	if err != nil {
		c.logger.Errorf("Error reading user info: %s", err.Error())
		return custom.Error(e, http.StatusBadRequest, &exception.Unauthorized{})

	}

	customerCreatingReq := &_customerModel.CustomerCreatingReq{
		ID:     customerInfo.ID,
		Email:  customerInfo.Email,
		Name:   customerInfo.Name,
		Avatar: customerInfo.Picture,
	}

	if err := c.oauth2Service.CustomerAccountCreating(customerCreatingReq); err != nil {
		return custom.Error(e, http.StatusInternalServerError, &exception.OAuth2Processing{})
	}

	c.setSameSiteCookie(e, accessTokenCookieName, token.AccessToken)
	c.setSameSiteCookie(e, refreshTokenCookieName, token.RefreshToken)

	return e.JSON(http.StatusOK, &_oauth2Model.LoginResponse{Message: "Login successful"})
}

func (c *googleOAuth2Controller) AdminLoginCallback(e echo.Context) error {
	ctx := context.Background()

	if err := retry.Do(func() error {
		return c.callbackValidating(e)
	}, retry.Attempts(3), retry.Delay(3*time.Second)); err != nil {
		c.logger.Errorf("Error validating callback: %s", err.Error())
		return custom.Error(e, http.StatusBadRequest, err)
	}

	token, err := adminGoogleOAuth2.Exchange(ctx, e.QueryParam("code"))
	if err != nil {
		c.logger.Errorf("Error exchanging code for token: %s", err.Error())
		return custom.Error(e, http.StatusBadRequest, &exception.Unauthorized{})
	}

	client := adminGoogleOAuth2.Client(ctx, token)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		c.logger.Errorf("Error reading user info: %s", err.Error())
		return custom.Error(e, http.StatusBadRequest, &exception.Unauthorized{})

	}

	createAdminReq := &_adminModel.AdminCreatingReq{
		ID:     userInfo.ID,
		Email:  userInfo.Email,
		Name:   userInfo.Name,
		Avatar: userInfo.Picture,
	}

	if err := c.oauth2Service.AdminAccountCreating(createAdminReq); err != nil {
		return custom.Error(e, http.StatusInternalServerError, &exception.Unauthorized{})
	}

	c.setSameSiteCookie(e, accessTokenCookieName, token.AccessToken)
	c.setSameSiteCookie(e, refreshTokenCookieName, token.RefreshToken)

	return e.JSON(http.StatusOK, &_oauth2Model.LoginResponse{Message: "Login successful"})
}

func (c *googleOAuth2Controller) Logout(e echo.Context) error {
	accessToken, err := e.Cookie(accessTokenCookieName)
	fmt.Println(accessToken)
	if err != nil {
		c.logger.Errorf("Error reading access token: %s", err.Error())
		return custom.Error(e, http.StatusBadRequest, &exception.Logout{})
	}

	if err := c.revokeToken(accessToken.Value); err != nil {
		c.logger.Errorf("Error revoking token: %s", err.Error())
		return custom.Error(e, http.StatusInternalServerError, &exception.Logout{})
	}

	c.removeSameSiteCookie(e, accessTokenCookieName)
	c.removeSameSiteCookie(e, refreshTokenCookieName)

	return e.JSON(http.StatusOK, &_oauth2Model.LogoutResponse{Message: "Logout successful"})
}

func (c *googleOAuth2Controller) revokeToken(accessToken string) error {
	revokeURL := fmt.Sprintf("%s?token=%s", c.oauth2Conf.RevokeUrl, accessToken)

	resp, err := http.Post(revokeURL, "application/x-www-form-urlencoded", nil)
	if err != nil {
		fmt.Println("Error revoking token:", err)
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *googleOAuth2Controller) callbackValidating(e echo.Context) error {
	state := e.QueryParam("state")

	stateFromCookie, err := e.Cookie(stateCookieName)
	if err != nil {
		c.logger.Errorf("Error reading state: %s", err.Error())
		return &exception.InvalidState{}
	}

	if state == "" || state != stateFromCookie.Value {
		c.logger.Errorf("Invalid state: %s != %s", state)
		return &exception.InvalidState{}
	}

	c.removeCookie(e, stateCookieName)

	return nil
}

func (c *googleOAuth2Controller) getUserInfo(client *http.Client) (*_oauth2Model.CustomerInfo, error) {
	resp, err := client.Get(c.oauth2Conf.CustomerInfoUrl)
	if err != nil {
		c.logger.Errorf("Error getting user info: %s", err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	customerInfoInBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Errorf("Error reading user info: %s", err.Error())
		return nil, err
	}

	customerInfo := &_oauth2Model.CustomerInfo{}
	if err := json.Unmarshal(customerInfoInBytes, &customerInfo); err != nil {
		c.logger.Errorf("Error unmarshalling user info: %s", err.Error())
		return nil, err
	}

	return customerInfo, nil
}

func (c *googleOAuth2Controller) setSameSiteCookie(e echo.Context, name, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
	}

	e.SetCookie(cookie)
}

func (c *googleOAuth2Controller) removeSameSiteCookie(e echo.Context, name string) {
	cookie := &http.Cookie{
		Name:     name,
		Path:     "/",
		MaxAge:   -1,
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
	}

	e.SetCookie(cookie)
}

func (c *googleOAuth2Controller) setCookie(e echo.Context, name, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
	}

	e.SetCookie(cookie)
}

func (c *googleOAuth2Controller) removeCookie(pctx echo.Context, name string) {
	cookie := &http.Cookie{
		Name:     name,
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}

	pctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) randomState() string {
	b := make([]byte, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
