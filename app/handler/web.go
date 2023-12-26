package handler

import (
	"text/template"

	md "github.com/absormu/go-jti/app/middleware"
	usecaseauth "github.com/absormu/go-jti/app/usecase/auth"
	"github.com/labstack/echo/v4"
)

func WebInputHandler(c echo.Context) (e error) {
	logger := md.GetLogger(c)
	logger.Info("handler: WebInputHandler")

	t, _ := template.ParseFiles("public/input.html")
	t.Execute(c.Response().Writer, false)

	return
}

func WebOutputHandler(c echo.Context) (e error) {
	logger := md.GetLogger(c)
	logger.Info("handler: WebOutputHandler")

	t, _ := template.ParseFiles("public/output.html")
	t.Execute(c.Response().Writer, false)

	return
}

func WebTokenHandler(c echo.Context) (e error) {
	logger := md.GetLogger(c)
	logger.Info("handler: WebTokenHandler")

	e = usecaseauth.GetTokenUser(c)

	return
}
