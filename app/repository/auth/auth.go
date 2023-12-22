package auth

import (
	"github.com/absormu/go-jti/app/entity"
	md "github.com/absormu/go-jti/app/middleware"
	db "github.com/absormu/go-jti/pkg/mariadb"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func GetAuth(c echo.Context, params map[string]string) (user entity.User, e error) {
	logger := md.GetLogger(c)
	logger.WithFields(logrus.Fields{
		"params": params,
	}).Info("repository: GetAuth")

	db := db.MariaDBInit()

	defer db.Close()

	query := "SELECT id, name, email, password FROM user"
	var condition string
	// Combine where clause
	clause := false
	for key, value := range params {

		if clause == false {
			condition += " WHERE "
		} else {
			condition += " AND "
		}
		condition += "user" + "." + key + " = '" + value + "'"
		clause = true
	}

	query += condition

	logger.WithFields(logrus.Fields{"query": query}).Info("repository: GetAuth-Query")

	result, err := db.Query(query)
	if err != nil {
		return
	}

	defer result.Close()

	for result.Next() {
		if e = result.Scan(&user.ID, &user.Name, &user.Email, &user.Password); e != nil {
			return
		}
	}

	return
}
