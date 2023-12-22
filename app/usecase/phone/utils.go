package phone

import (
	"strconv"
	"strings"

	md "github.com/absormu/go-jti/app/middleware"
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
)

func getType(number int) int {

	if number%2 == 0 {
		return 2
	}

	return 1
}

func isPhoneNumber(phone string) bool {
	if strings.HasPrefix(phone, "+62") && len(phone) >= 8 && len(phone) <= 15 {
		return true
	} else if strings.HasPrefix(phone, "62") && len(phone) >= 8 && len(phone) <= 14 {
		return true
	} else if strings.HasPrefix(phone, "08") && len(phone) >= 8 && len(phone) <= 13 {
		return true
	} else {
		return false
	}
}

func delChar(s []rune, index int) []rune {
	return append(s[0:index], s[index+1:]...)
}

func autoNumberNineDigit(c echo.Context, guid xid.ID) (resultNumber string, resultType int) {
	logger := md.GetLogger(c)
	logger.WithFields(log.Fields{
		"guid": guid,
	}).Info("generate: autoNumberNineDigit")

	uniqNumber := guid.Time().UTC().Unix()
	resultType = getType(int(uniqNumber))

	var s, res []rune
	s = []rune(strconv.Itoa(int(uniqNumber)))
	res = delChar(s, 0)

	s = []rune(res)
	res = delChar(s, 0)

	resultNumber = string(res)

	logger.WithFields(log.Fields{
		"number": resultNumber,
		"lenght": len(string(res)),
	}).Info("result: autoNumberNineDigit")

	return
}
