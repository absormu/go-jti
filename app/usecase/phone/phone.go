package phone

import (
	"net/http"
	"strconv"
	"time"

	"github.com/absormu/go-jti/app/entity"
	md "github.com/absormu/go-jti/app/middleware"
	repoPhone "github.com/absormu/go-jti/app/repository/phone"
	cm "github.com/absormu/go-jti/pkg/configuration"
	lg "github.com/absormu/go-jti/pkg/response"
	resp "github.com/absormu/go-jti/pkg/response"
	sdk "github.com/absormu/go-jti/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
)

func GetProviders(c echo.Context) (e error) {
	logger := md.GetLogger(c)
	logger.Info("usecase: GetProviders")

	var results []entity.Provider
	if results, e = repoPhone.GetProviders(c); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error failure query GetProviders")
		e = resp.CustomError(c, http.StatusInternalServerError, sdk.ERR_DATABASE,
			lg.Language{Bahasa: nil, English: "Failure query get GetProviders"}, nil, nil)
		return
	}

	if len(results) == 0 {
		logger.Error("Catch error data not found")
		e = resp.CustomError(c, http.StatusNotFound, sdk.ERR_DATA_NOT_FOUND,
			lg.Language{Bahasa: "Data tidak tersedia", English: "Data not found"}, nil, nil)
		return
	}

	var responseData []entity.Provider
	var data entity.Provider
	for _, result := range results {

		data.ID = result.ID
		data.Code = result.Code
		data.Name = result.Name

		responseData = append(responseData, data)
	}

	e = resp.CustomError(c, http.StatusOK, sdk.ERR_SUCCESS,
		lg.Language{Bahasa: "Sukses", English: "Success"}, nil, responseData)
	return
}

func CreateAutoNumberPhone(c echo.Context, extractToken entity.ExtractToken) (e error) {
	logger := md.GetLogger(c)
	logger.WithField("request", cm.Config.LimitGenerate).Info("usecase: CreateAutoNumberPhone")

	for index := 0; index < cm.Config.LimitGenerate; index++ {

		var result entity.Provider
		if result, e = repoPhone.GetProviderRandom(c); e != nil {
			logger.WithField("error", e.Error()).Error("Catch error failure query GetProviderRandom")
			e = resp.CustomError(c, http.StatusInternalServerError, sdk.ERR_DATABASE,
				lg.Language{Bahasa: nil, English: "Failure query get GetProviderRandom"}, nil, nil)
			return
		}

		var empty entity.Provider
		if result == empty {
			logger.Error("Catch error data not found")
			e = resp.CustomError(c, http.StatusNotFound, sdk.ERR_DATA_NOT_FOUND,
				lg.Language{Bahasa: "Data tidak tersedia", English: "Data not found"}, nil, nil)
			return
		}

		resultNumber, resultType := autoNumberNineDigit(c, xid.New())
		autoNumber := result.Code + resultNumber
		autoProviderID := result.ID

		var encryptedNumber string
		encryptedNumber, e = sdk.GetAESEncrypted(autoNumber)
		if e != nil {
			logger.WithField("error", e.Error()).Error("Error during encryption")
			e = resp.CustomError(c, http.StatusInternalServerError, sdk.ERR_INVALID_FORMAT,
				lg.Language{Bahasa: nil, English: "Error during encryption"}, nil, nil)
			return
		}

		params := make(map[string]interface{})
		params["number"] = encryptedNumber
		params["type"] = resultType
		params["provider_id"] = autoProviderID
		params["created_by"] = extractToken.Name
		params["created_at"] = time.Now()

		if _, e = repoPhone.CreateNumberPhone(c, params); e != nil {
			logger.WithField("error", e.Error()).Error("Catch error failure query CreateNumberPhone")
			e = resp.CustomError(c, http.StatusInternalServerError, sdk.ERR_DATABASE,
				lg.Language{Bahasa: nil, English: "Failure query create CreateNumberPhone"}, nil, nil)
			return
		}

		logger.WithFields(log.Fields{
			"number": autoNumber,
			"data":   index + 1,
		}).Info("result: CreateNumberPhone")
	}

	e = resp.CustomError(c, http.StatusOK, sdk.ERR_SUCCESS,
		lg.Language{Bahasa: "Sukses", English: "Success"}, nil, nil)
	return
}

func CreateNumberPhone(c echo.Context, req entity.PhoneData, extractToken entity.ExtractToken) (e error) {
	logger := md.GetLogger(c)
	logger.WithField("request", req).Info("usecase: CreateNumberPhone")

	if req.Number == "" || req.Provider.ID == 0 {
		logger.Error("Missing mandatory parameter")
		e = resp.CustomError(c, http.StatusBadRequest, sdk.ERR_PARAM_MISSING,
			lg.Language{Bahasa: "Missing mandatory parameter", English: "Missing mandatory parameter"}, nil, nil)
		return
	}

	if !isPhoneNumber(req.Number) {
		logger.Error("invalid phone number format ", req.Number)
		e = resp.CustomError(c, http.StatusBadRequest, sdk.ERR_PARAM_MISSING,
			lg.Language{Bahasa: "Format nomor telepon tidak valid", English: "Invalid phone number format"}, nil, nil)
		return
	}

	encryptedNumber, e := sdk.GetAESEncrypted(req.Number)
	if e != nil {
		logger.WithField("error", e.Error()).Error("Error during encryption")
		e = resp.CustomError(c, http.StatusInternalServerError, sdk.ERR_INVALID_FORMAT,
			lg.Language{Bahasa: nil, English: "Error during encryption"}, nil, nil)
		return
	}

	var phoneData entity.PhoneData
	if phoneData, e = repoPhone.GetNumberPhoneByNumber(c, encryptedNumber); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error failure query GetProviderByID")
		e = resp.CustomError(c, http.StatusInternalServerError, sdk.ERR_DATABASE,
			lg.Language{Bahasa: nil, English: "Failure query get GetProviderByID"}, nil, nil)
		return
	}

	var emptyPhone entity.PhoneData
	if phoneData != emptyPhone {
		logger.Error("Catch error number exist ", req.Number)
		e = resp.CustomError(c, http.StatusConflict, sdk.ERR_DATA_ALREADY_EXIST,
			lg.Language{Bahasa: "Nomor Handphone Sudah Ada", English: "Number Phone Already Exist"}, nil, nil)
		return
	}

	var result entity.Provider
	if result, e = repoPhone.GetProviderByID(c, int(req.Provider.ID)); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error failure query GetProviderByID")
		e = resp.CustomError(c, http.StatusInternalServerError, sdk.ERR_DATABASE,
			lg.Language{Bahasa: nil, English: "Failure query get GetProviderByID"}, nil, nil)
		return
	}

	var empty entity.Provider
	if result == empty {
		logger.Error("Catch error provider id not found")
		e = resp.CustomError(c, http.StatusNotFound, sdk.ERR_DATA_NOT_FOUND,
			lg.Language{Bahasa: "Provider tidak tersedia", English: "Provider not found"}, nil, nil)
		return
	}

	var intNumber int
	if intNumber, e = strconv.Atoi(req.Number); e != nil {
		logger.WithField("error", e.Error()).Info("Bad Request")
		e = resp.CustomError(c, http.StatusBadRequest, sdk.ERR_PARAM_MISSING,
			lg.Language{Bahasa: "Bad Request", English: "Bad Request"}, nil, nil)
		return
	}

	params := make(map[string]interface{})
	params["number"] = encryptedNumber
	params["type"] = getType(intNumber)
	params["provider_id"] = req.Provider.ID
	params["created_by"] = extractToken.Name
	params["created_at"] = time.Now()

	if _, e = repoPhone.CreateNumberPhone(c, params); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error failure query CreateNumberPhone")
		e = resp.CustomError(c, http.StatusInternalServerError, sdk.ERR_DATABASE,
			lg.Language{Bahasa: nil, English: "Failure query create CreateNumberPhone"}, nil, nil)
		return
	}

	e = resp.CustomError(c, http.StatusOK, sdk.ERR_SUCCESS,
		lg.Language{Bahasa: "Sukses", English: "Success"}, nil, nil)
	return
}

func GetNumberPhones(c echo.Context) (e error) {
	logger := md.GetLogger(c)
	logger.Info("usecase: GetNumberPhones")

	var results []entity.PhoneData
	if results, e = repoPhone.GetNumberPhones(c); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error failure query GetNumberPhones")
		e = resp.CustomError(c, http.StatusInternalServerError, sdk.ERR_DATABASE,
			lg.Language{Bahasa: nil, English: "Failure query get GetNumberPhones"}, nil, nil)
		return
	}

	if len(results) == 0 {
		logger.Error("Catch error data not found")
		e = resp.CustomError(c, http.StatusNotFound, sdk.ERR_DATA_NOT_FOUND,
			lg.Language{Bahasa: "Data tidak tersedia", English: "Data not found"}, nil, nil)
		return
	}

	var responseData []entity.PhoneData
	var data entity.PhoneData
	for _, result := range results {

		decryptedNumber, _ := sdk.GetAESDecrypted(result.Number)

		data.ID = result.ID
		data.Number = string(decryptedNumber)
		data.Type = result.Type

		data.Provider.ID = result.Provider.ID
		data.Provider.Code = result.Provider.Code
		data.Provider.Name = result.Provider.Name

		data.CreatedAt = result.CreatedAt
		data.CreatedBy = result.CreatedBy
		data.ModifiedAt = result.ModifiedAt
		data.ModifiedBy = result.ModifiedBy
		responseData = append(responseData, data)

	}

	e = resp.CustomError(c, http.StatusOK, sdk.ERR_SUCCESS,
		lg.Language{Bahasa: "Sukses", English: "Success"}, nil, responseData)
	return
}

func GetNumberPhoneByID(c echo.Context) (e error) {
	logger := md.GetLogger(c)
	logger.WithField("id", c.Param("id")).Info("usecase: GetNumberPhoneByID")

	idStr := c.Param("id")
	var id int
	if id, e = strconv.Atoi(idStr); e != nil {
		logger.WithField("error", e.Error()).Info("Bad Request")
		e = resp.CustomError(c, http.StatusBadRequest, sdk.ERR_PARAM_MISSING,
			lg.Language{Bahasa: "Bad Request", English: "Bad Request"}, nil, nil)
		return
	}

	if id < 1 {
		logger.Error("Catch error id : ", id)
		e = resp.CustomError(c, http.StatusBadRequest, sdk.ERR_PARAM_MISSING,
			lg.Language{Bahasa: "Bad Request", English: "Bad Request"}, nil, nil)
		return
	}

	var result entity.PhoneData
	if result, e = repoPhone.GetNumberPhoneByID(c, id); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error failure query GetNumberPhoneByID")
		e = resp.CustomError(c, http.StatusInternalServerError, sdk.ERR_DATABASE,
			lg.Language{Bahasa: nil, English: "Failure query get GetNumberPhoneByID"}, nil, nil)
		return
	}

	var empty entity.PhoneData
	if result == empty {
		logger.Error("Catch error data not found ", id)
		e = resp.CustomError(c, http.StatusNotFound, sdk.ERR_DATA_NOT_FOUND,
			lg.Language{Bahasa: "Data tidak tersedia", English: "Data not found"}, nil, nil)
		return
	}

	decryptedNumber, e := sdk.GetAESDecrypted(result.Number)
	if e != nil {
		logger.WithField("error", e.Error()).Error("Error during decryption")
		e = resp.CustomError(c, http.StatusInternalServerError, sdk.ERR_INVALID_FORMAT,
			lg.Language{Bahasa: nil, English: "Error during decryption"}, nil, nil)
		return
	}

	var data entity.PhoneData
	data.ID = result.ID
	data.Number = string(decryptedNumber)
	data.Type = result.Type

	data.Provider.ID = result.Provider.ID
	data.Provider.Code = result.Provider.Code
	data.Provider.Name = result.Provider.Name

	data.CreatedAt = result.CreatedAt
	data.CreatedBy = result.CreatedBy
	data.ModifiedAt = result.ModifiedAt
	data.ModifiedBy = result.ModifiedBy

	e = resp.CustomError(c, http.StatusOK, sdk.ERR_SUCCESS,
		lg.Language{Bahasa: "Sukses", English: "Success"}, nil, data)
	return
}

func UpdateNumberPhone(c echo.Context, req entity.PhoneData, extractToken entity.ExtractToken) (e error) {
	logger := md.GetLogger(c)

	params := make(map[string]interface{})

	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	if id > 0 {
		params["id"] = id
	} else {
		e = resp.CustomError(c, http.StatusBadRequest, sdk.ERR_PARAM_MISSING,
			lg.Language{Bahasa: nil, English: "Bad Request"}, nil, nil)
		return
	}

	logger.WithFields(log.Fields{"id": id, "request": req}).Info("usecase: UpdateNumberPhone")

	if id == 0 {
		logger.Error("Missing mandatory parameter")
		e = resp.CustomError(c, http.StatusBadRequest, sdk.ERR_PARAM_MISSING,
			lg.Language{Bahasa: "Missing mandatory parameter", English: "Missing mandatory parameter"}, nil, nil)
		return
	}

	if req.Provider.ID != 0 {
		var result entity.Provider
		if result, e = repoPhone.GetProviderByID(c, int(req.Provider.ID)); e != nil {
			logger.WithField("error", e.Error()).Error("Catch error failure query GetProviderByID")
			e = resp.CustomError(c, http.StatusInternalServerError, sdk.ERR_DATABASE,
				lg.Language{Bahasa: nil, English: "Failure query get GetProviderByID"}, nil, nil)
			return
		}

		var empty entity.Provider
		if result == empty {
			logger.Error("Catch error provider id not found")
			e = resp.CustomError(c, http.StatusNotFound, sdk.ERR_DATA_NOT_FOUND,
				lg.Language{Bahasa: "Provider tidak tersedia", English: "Provider not found"}, nil, nil)
			return
		}
	}

	if req.Number != "" {
		var encryptedNumber string
		encryptedNumber, e = sdk.GetAESEncrypted(req.Number)
		if e != nil {
			logger.WithField("error", e.Error()).Error("Error during encryption")
			e = resp.CustomError(c, http.StatusInternalServerError, sdk.ERR_INVALID_FORMAT,
				lg.Language{Bahasa: nil, English: "Error during encryption"}, nil, nil)
			return
		}

		var phoneData entity.PhoneData
		if phoneData, e = repoPhone.GetNumberPhoneByNumber(c, encryptedNumber); e != nil {
			logger.WithField("error", e.Error()).Error("Catch error failure query GetProviderByID")
			e = resp.CustomError(c, http.StatusInternalServerError, sdk.ERR_DATABASE,
				lg.Language{Bahasa: nil, English: "Failure query get GetProviderByID"}, nil, nil)
			return
		}

		var emptyPhone entity.PhoneData
		if phoneData != emptyPhone {
			logger.Error("Catch error number exist ", req.Number)
			e = resp.CustomError(c, http.StatusConflict, sdk.ERR_DATA_ALREADY_EXIST,
				lg.Language{Bahasa: "Nomor Handphone Sudah Ada", English: "Number Phone Already Exist"}, nil, nil)
			return
		}

		params["number"] = encryptedNumber
	}

	var intNumber int
	if intNumber, e = strconv.Atoi(req.Number); e != nil {
		logger.WithField("error", e.Error()).Info("Bad Request")
		e = resp.CustomError(c, http.StatusBadRequest, sdk.ERR_PARAM_MISSING,
			lg.Language{Bahasa: "Bad Request", English: "Bad Request"}, nil, nil)
		return
	}

	if intNumber > 0 {
		params["type"] = getType(intNumber)
	}

	if req.Provider.ID > 0 {
		params["provider_id"] = req.Provider.ID
	}

	params["modified_by"] = extractToken.Name

	time := time.Now()
	params["modified_at"] = time

	var rowsAffected int64
	rowsAffected, e = repoPhone.UpdateNumberPhone(c, params)
	if e != nil {
		logger.WithField("error", e.Error()).Error("Catch error failure query UpdateSeller")
		e = resp.CustomError(c, http.StatusInternalServerError, sdk.ERR_DATABASE,
			lg.Language{Bahasa: nil, English: "Failure query update seller"}, nil, nil)
		return
	}

	if rowsAffected == 0 {
		logger.Error("Catch error id not found")
		e = resp.CustomError(c, http.StatusNotFound, sdk.ERR_DATA_NOT_FOUND,
			lg.Language{Bahasa: "Data tidak tersedia", English: "Data not found"}, nil, nil)
		return
	}

	e = resp.CustomError(c, http.StatusOK, sdk.ERR_SUCCESS,
		lg.Language{Bahasa: "Sukses", English: "Success"}, nil, nil)
	return
}

func DeleteNumberPhoneByID(c echo.Context, extractToken entity.ExtractToken) (e error) {
	logger := md.GetLogger(c)

	// Get parameter id
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	if id < 1 {
		logger.Error("Catch error id")
		e = resp.CustomError(c, http.StatusBadRequest, sdk.ERR_PARAM_MISSING,
			lg.Language{Bahasa: nil, English: "Bad Request"}, nil, nil)
		return
	}

	logger.WithFields(log.Fields{"id": id, "name": extractToken.Name}).Info("usecase: DeleteNumberPhoneByID")

	var rowsAffected int64
	rowsAffected, e = repoPhone.DeleteNumberPhoneByID(c, id, extractToken.Name)
	if e != nil {
		logger.WithField("error", e.Error()).Error("Catch error failure query DeleteNumberPhoneByID")
		e = resp.CustomError(c, http.StatusInternalServerError, sdk.ERR_DATABASE,
			lg.Language{Bahasa: nil, English: "Failure query delete seller"}, nil, nil)
		return
	}

	if rowsAffected == 0 {
		logger.Error("Catch error id not found")
		e = resp.CustomError(c, http.StatusNotFound, sdk.ERR_DATA_NOT_FOUND,
			lg.Language{Bahasa: "Data tidak tersedia", English: "Data not found"}, nil, nil)
		return
	}

	e = resp.CustomError(c, http.StatusOK, sdk.ERR_SUCCESS,
		lg.Language{Bahasa: "Sukses", English: "Success"}, nil, nil)
	return
}
