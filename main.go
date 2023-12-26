package main

import (
	"os"

	handler "github.com/absormu/go-jti/app/handler"
	md "github.com/absormu/go-jti/app/middleware"
	cm "github.com/absormu/go-jti/pkg/configuration"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

func initHandlers(e *echo.Echo) {

	root := e.Group(cm.Config.RootURL)
	root.GET("/ping", handler.PingHandler)
	root.POST("/api/v1/login", handler.LoginHandler)

	r := e.Group(cm.Config.RootURL)
	config := middleware.JWTConfig{
		Claims:     &md.JwtCustomClaims{},
		SigningKey: []byte(cm.Config.ClientSecret),
	}

	r.Use(middleware.JWTWithConfig(config))

	v1 := r.Group("/api/v1")
	v1.GET("/providers", handler.GetProvidersHandler)
	v1.POST("/auto-number-phones", handler.CreateAutoNumberPhoneHandler)
	v1.POST("/number-phone", handler.CreateNumberPhoneHandler)

	np := v1.Group("/number-phone")
	v1.GET("/number-phones", handler.GetNumberPhonesHandler)
	np.GET("/:id", handler.GetNumberPhoneByIDHandler)
	np.PUT("/:id", handler.UpdateNumberPhoneHandler)
	np.DELETE("/:id", handler.DeleteNumberPhoneByIDHandler)

	log.Info("Staring server ...")
}

func initLogger() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999",
	})
}

func main() {
	e := echo.New()
	initLogger()

	cm.LoadConfig()

	e.Use(md.AddLogger)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	initHandlers(e)

	err := e.Start(cm.Config.ListenPort)

	if err != nil {
		log.WithField("error", err).Error("Unable to start the server")
		os.Exit(1)
	}

}
