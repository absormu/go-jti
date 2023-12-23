package handler

import (
	"text/template"

	md "github.com/absormu/go-jti/app/middleware"
	"github.com/labstack/echo/v4"
)

func WebInputHandler(c echo.Context) (e error) {
	logger := md.GetLogger(c)
	logger.Info("handler: WebInputHandler")

	t, _ := template.ParseFiles("templates/input.html")
	t.Execute(c.Response().Writer, false)

	return
}
