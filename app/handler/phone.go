package handler

import (
	"net/http"

	"github.com/absormu/go-jti/app/entity"
	usecasePhone "github.com/absormu/go-jti/app/usecase/phone"
	pkgjwt "github.com/absormu/go-jti/pkg/jwt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	md "github.com/absormu/go-jti/app/middleware"
	lg "github.com/absormu/go-jti/pkg/response"
	resp "github.com/absormu/go-jti/pkg/response"
	sdk "github.com/absormu/go-jti/pkg/utils"
)

func GetProvidersHandler(c echo.Context) (e error) {
	logger := md.GetLogger(c)
	logger.Info("handler: GetProvidersHandler")

	extractToken, e := pkgjwt.ExtractToken(c)
	if e != nil {
		logger.Error("Catch error extractToken")
		e = resp.CustomError(c, http.StatusUnauthorized, sdk.ERR_UNAUTHORIZED,
			lg.Language{Bahasa: nil, English: "Authorization missing"}, nil, nil)
		return
	}

	logger.WithFields(log.Fields{
		"extractToken": extractToken,
	}).Info("ExtractToken")

	e = usecasePhone.GetProviders(c)

	return
}

func CreateAutoNumberPhoneHandler(c echo.Context) (e error) {
	logger := md.GetLogger(c)
	logger.Info("handler: CreateAutoNumberPhoneHandler")

	extractToken, e := pkgjwt.ExtractToken(c)
	if e != nil {
		logger.Error("Catch error extractToken")
		e = resp.CustomError(c, http.StatusUnauthorized, sdk.ERR_UNAUTHORIZED,
			lg.Language{Bahasa: nil, English: "Authorization missing"}, nil, nil)
		return
	}

	logger.WithFields(log.Fields{
		"extractToken": extractToken,
	}).Info("ExtractToken")

	e = usecasePhone.CreateAutoNumberPhone(c, extractToken)

	return
}

func CreateNumberPhoneHandler(c echo.Context) (e error) {
	logger := md.GetLogger(c)
	logger.Info("handler: CreateNumberPhoneHandler")

	extractToken, e := pkgjwt.ExtractToken(c)
	if e != nil {
		logger.Error("Catch error extractToken")
		e = resp.CustomError(c, http.StatusUnauthorized, sdk.ERR_UNAUTHORIZED,
			lg.Language{Bahasa: nil, English: "Authorization missing"}, nil, nil)
		return
	}

	logger.WithFields(log.Fields{
		"extractToken": extractToken,
	}).Info("ExtractToken")

	req := entity.PhoneData{}
	if e = c.Bind(&req); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error bind request")
		e = resp.CustomError(c, http.StatusBadRequest, sdk.ERR_PARAM_ILLEGAL,
			lg.Language{Bahasa: nil, English: e.Error()}, nil, nil)
		return
	}

	e = usecasePhone.CreateNumberPhone(c, req, extractToken)

	return
}

func GetNumberPhonesHandler(c echo.Context) (e error) {
	logger := md.GetLogger(c)
	logger.Info("handler: GetNumberPhonesHandler")

	extractToken, e := pkgjwt.ExtractToken(c)
	if e != nil {
		logger.Error("Catch error extractToken")
		e = resp.CustomError(c, http.StatusUnauthorized, sdk.ERR_UNAUTHORIZED,
			lg.Language{Bahasa: nil, English: "Authorization missing"}, nil, nil)
		return
	}

	logger.WithFields(log.Fields{
		"extractToken": extractToken,
	}).Info("ExtractToken")

	e = usecasePhone.GetNumberPhones(c)

	return
}

func GetNumberPhoneByIDHandler(c echo.Context) (e error) {
	logger := md.GetLogger(c)
	logger.Info("handler: GetNumberPhoneByIDHandler")

	extractToken, e := pkgjwt.ExtractToken(c)
	if e != nil {
		logger.Error("Catch error extractToken")
		e = resp.CustomError(c, http.StatusUnauthorized, sdk.ERR_UNAUTHORIZED,
			lg.Language{Bahasa: "Authorization missing", English: "Authorization missing"}, nil, nil)
		return
	}

	logger.WithFields(log.Fields{
		"extractToken": extractToken,
	}).Info("ExtractToken")

	e = usecasePhone.GetNumberPhoneByID(c)

	return
}

func UpdateNumberPhoneHandler(c echo.Context) (e error) {
	logger := md.GetLogger(c)
	logger.Info("handler: UpdateNumberPhoneHandler")

	extractToken, e := pkgjwt.ExtractToken(c)
	if e != nil {
		logger.Error("Catch error extractToken")
		e = resp.CustomError(c, http.StatusUnauthorized, sdk.ERR_UNAUTHORIZED,
			lg.Language{Bahasa: nil, English: "Authorization missing"}, nil, nil)
		return
	}

	logger.WithFields(log.Fields{
		"extractToken": extractToken,
	}).Info("ExtractToken")

	req := entity.PhoneData{}
	if e = c.Bind(&req); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error bind request")
		e = resp.CustomError(c, http.StatusBadRequest, sdk.ERR_PARAM_ILLEGAL,
			lg.Language{Bahasa: nil, English: e.Error()}, nil, nil)
		return
	}

	e = usecasePhone.UpdateNumberPhone(c, req, extractToken)

	return
}

func DeleteNumberPhoneByIDHandler(c echo.Context) (e error) {
	logger := md.GetLogger(c)
	logger.Info("handler: DeleteNumberPhoneByIDHandler")

	extractToken, e := pkgjwt.ExtractToken(c)
	if e != nil {
		logger.Error("Catch error extractToken")
		e = resp.CustomError(c, http.StatusUnauthorized, sdk.ERR_UNAUTHORIZED,
			lg.Language{Bahasa: nil, English: "Authorization missing"}, nil, nil)
		return
	}

	logger.WithFields(log.Fields{
		"extractToken": extractToken,
	}).Info("ExtractToken")

	e = usecasePhone.DeleteNumberPhoneByID(c, extractToken)

	return
}
