package handler

import (
	"text/template"

	"github.com/labstack/echo/v4"

	md "github.com/absormu/go-jti/app/middleware"
)

func WebInputHandler(c echo.Context) (e error) {
	logger := md.GetLogger(c)
	logger.Info("handler: WebInputHandler")

	t, _ := template.ParseFiles("templates/input.html")
	t.Execute(c.Response().Writer, false)

	return
}

func WebOutputHandler(c echo.Context) (e error) {
	logger := md.GetLogger(c)
	logger.Info("handler: WebOutputHandler")

	t, _ := template.ParseFiles("templates/output.html")
	t.Execute(c.Response().Writer, false)

	return
}
