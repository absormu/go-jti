package handler

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/absormu/go-jti/app/entity"
	md "github.com/absormu/go-jti/app/middleware"
	usecaseauth "github.com/absormu/go-jti/app/usecase/auth"
	cm "github.com/absormu/go-jti/pkg/configuration"
	lg "github.com/absormu/go-jti/pkg/response"
	resp "github.com/absormu/go-jti/pkg/response"
	sdk "github.com/absormu/go-jti/pkg/utils"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	log "github.com/sirupsen/logrus"
)

func LoginHandler(c echo.Context) (e error) {
	logger := md.GetLogger(c)
	logger.Info("handler: LoginHandler")

	req := entity.Auth{}
	if e = c.Bind(&req); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error bind request")
		e = resp.CustomError(c, http.StatusBadRequest, sdk.ERR_PARAM_ILLEGAL,
			lg.Language{Bahasa: e.Error(), English: e.Error()}, nil, nil)
		return
	}

	e = usecaseauth.Login(c, req)

	return
}

func AuthHandler(c echo.Context) (e error) {
	logger := md.GetLogger(c)
	logger.Info("handler: AuthHandler")

	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(c.Response().Writer, false)

	return
}

func ProviderAuthHandler(c echo.Context) (e error) {
	logger := md.GetLogger(c)
	logger.Info("handler: ProviderAuthHandler")

	store := sessions.NewCookieStore([]byte("Secret-session-key"))
	store.MaxAge(86400 * 30) // 30 days
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = false

	gothic.Store = store

	goth.UseProviders(
		google.New(cm.Config.AuthClientKey, cm.Config.AuthClientSecret, cm.Config.AuthCallbackURL, "email", "profile"),
	)

	gothic.BeginAuthHandler(c.Response().Writer, c.Request())

	return
}

func ProviderCallbackAuthHandler(c echo.Context) (e error) {
	logger := md.GetLogger(c)
	logger.Info("handler: ProviderCallbackAuthHandler")

	store := sessions.NewCookieStore([]byte("Secret-session-key"))
	store.MaxAge(86400 * 30) // 30 days
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = false

	gothic.Store = store

	goth.UseProviders(
		google.New(cm.Config.AuthClientKey, cm.Config.AuthClientSecret, cm.Config.AuthCallbackURL, "email", "profile"),
	)

	user, err := gothic.CompleteUserAuth(c.Response().Writer, c.Request())
	if err != nil {
		fmt.Fprintln(c.Response().Writer, err)
		return
	}

	logger.WithFields(log.Fields{"id": user.UserID, "email": user.Email}).Info("handler: ProviderCallbackAuthHandler-Success")

	t, _ := template.ParseFiles("templates/success.html")
	t.Execute(c.Response().Writer, user)

	return
}
