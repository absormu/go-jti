package handler

import (
	"net/http"

	usecaseauth "github.com/absormu/go-jti/app/usecase/auth"
	"github.com/labstack/echo/v4"

	"github.com/absormu/go-jti/app/entity"
	md "github.com/absormu/go-jti/app/middleware"
	lg "github.com/absormu/go-jti/pkg/response"
	resp "github.com/absormu/go-jti/pkg/response"
	sdk "github.com/absormu/go-jti/pkg/utils"
)

func PingHandler(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "PONG")
}

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
