package auth

import (
	"context"

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

	query := "SELECT id, name, email, password, IFNULL(token, '') FROM user"
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
		if e = result.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.AccessToken); e != nil {
			return
		}
	}

	return
}

func CreateUser(c echo.Context, params map[string]interface{}) (e error) {
	logger := md.GetLogger(c)
	logger.WithField("request", params).Info("repository: CreateUser")

	db := db.MariaDBInit()
	defer db.Close()

	ctx := context.Background()
	tx, e := db.BeginTx(ctx, nil)
	if e != nil {
		return
	}

	defer tx.Rollback()

	query := "INSERT INTO user ("
	var fields = ""
	var values = ""
	i := 0
	var valuestr []interface{}
	for key, value := range params {
		fields += "`" + key + "`"
		values += "" + "?" + ""
		valuestr = append(valuestr, value)
		if (len(params) - 1) > i {
			fields += ", "
			values += ", "
		}
		i++
	}

	query += fields + ") VALUES(" + values + ")"

	logger.WithFields(logrus.Fields{"query": query}).Info("repository: CreateUser-query")

	_, e = tx.ExecContext(ctx, query, valuestr...)
	if e != nil {
		return
	}

	// Commit the transaction.
	if e = tx.Commit(); e != nil {
		return
	}

	return
}

func UpdateUserToken(c echo.Context, params map[string]interface{}) (e error) {
	logger := md.GetLogger(c)
	logger.WithFields(logrus.Fields{
		"params": params,
	}).Info("repository: UpdateUserToken")

	db := db.MariaDBInit()

	defer db.Close()

	ctx := context.Background()
	tx, e := db.BeginTx(ctx, nil)
	if e != nil {
		return
	}

	defer tx.Rollback()

	query := "UPDATE user SET "
	i := 0
	var values []interface{}
	for key, value := range params {
		if key != "email" {
			query += key + " = ?"
			if (len(params) - 2) > i {
				query += ", "
			}
			i++
			values = append(values, value)
		}
	}

	query += " WHERE email = ? AND is_deleted = ?"
	values = append(values, params["email"], 0)

	logger.WithFields(logrus.Fields{"query": query}).Info("repository: UpdateUserToken-query")

	_, e = tx.ExecContext(ctx, query, values...)
	if e != nil {
		return
	}

	// Commit the transaction.
	if e = tx.Commit(); e != nil {
		return
	}

	return
}
