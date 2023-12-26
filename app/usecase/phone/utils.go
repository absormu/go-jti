package phone

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

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

	uniqNumber := randomString(8)
	uniqNumberInt, _ := strconv.Atoi(uniqNumber)
	resultType = getType(uniqNumberInt)

	resultNumber = string(uniqNumber)

	logger.WithFields(log.Fields{
		"number": resultNumber,
		"lenght": len(uniqNumber),
	}).Info("result: autoNumberNineDigit")

	return
}

func randomString(length int) string {
	var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	var letters = []rune("0123456789")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}

	return string(b)
}
