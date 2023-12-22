package middleware

import (
	cm "github.com/absormu/go-jti/pkg/configuration"
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
)

func AddLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("logger", log.WithFields(log.Fields{
			"app":    cm.Config.OriginHost,
			"msg_id": xid.New().String(),
		}))
		return next(c)
	}
}

func GetLogger(ctx echo.Context) *log.Entry {
	logger := ctx.Get("logger")
	if logger != nil {
		return logger.(*log.Entry)
	}
	return log.WithFields(log.Fields{
		"app": cm.Config.OriginHost,
	})

}
