package phone

import (
	"context"
	"strconv"
	"time"

	"github.com/absormu/go-jti/app/entity"
	md "github.com/absormu/go-jti/app/middleware"
	db "github.com/absormu/go-jti/pkg/mariadb"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func GetProviders(c echo.Context) (results []entity.Provider, e error) {
	logger := md.GetLogger(c)
	logger.Info("repository: GetProviders")

	db := db.MariaDBInit()
	defer db.Close()

	query := "SELECT p.id, p.code, p.name " +
		"FROM provider AS p "

	logger.WithFields(logrus.Fields{"query": query}).Info("repository: GetProviders-query")

	result, e := db.Query(query)
	if e != nil {
		return
	}

	defer result.Close()
	for result.Next() {
		var data entity.Provider
		if e = result.Scan(&data.ID, &data.Code, &data.Name); e != nil {
			return
		}
		results = append(results, data)
	}

	return
}

func CreateNumberPhone(c echo.Context, params map[string]interface{}) (numberPhoneID int64, e error) {
	logger := md.GetLogger(c)
	logger.WithField("request", params).Info("repository: CreateNumberPhone")

	db := db.MariaDBInit()
	defer db.Close()

	ctx := context.Background()
	tx, e := db.BeginTx(ctx, nil)
	if e != nil {
		return
	}

	defer tx.Rollback()

	query := "INSERT INTO phone_number ("
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

	logger.WithFields(logrus.Fields{"query": query}).Info("repository: CreateNumberPhone-query")

	result, e := tx.ExecContext(ctx, query, valuestr...)
	if e != nil {
		return
	}
	numberPhoneID, e = result.LastInsertId()
	if e != nil {
		return
	}

	// Commit the transaction.
	if e = tx.Commit(); e != nil {
		return
	}

	return
}

func GetNumberPhones(c echo.Context) (results []entity.PhoneData, e error) {
	logger := md.GetLogger(c)
	logger.Info("repository: GetNumberPhones")

	db := db.MariaDBInit()
	defer db.Close()

	query := "SELECT pn.id, pn.number, pn.type, " +
		"p.id, p.code, p.name, " +
		"pn.created_at, pn.created_by, IFNULL(pn.modified_at, ''), IFNULL(pn.modified_by, ''), pn.is_deleted " +
		"FROM phone_number AS pn " +
		"LEFT JOIN provider AS p ON pn.provider_id  = p.id " +
		"WHERE pn.is_deleted = 0 ORDER BY pn.id DESC LIMIT 100"

	logger.WithFields(logrus.Fields{"query": query}).Info("repository: GetNumberPhones-query")

	result, e := db.Query(query)
	if e != nil {
		return
	}

	defer result.Close()
	for result.Next() {
		var data entity.PhoneData
		if e = result.Scan(&data.ID, &data.Number, &data.Type,
			&data.Provider.ID, &data.Provider.Code, &data.Provider.Name,
			&data.CreatedAt, &data.CreatedBy, &data.ModifiedAt, &data.ModifiedBy, &data.IsDeleted); e != nil {
			return
		}
		results = append(results, data)
	}

	return
}

func GetNumberPhoneByID(c echo.Context, id int) (data entity.PhoneData, e error) {
	logger := md.GetLogger(c)
	logger.WithFields(logrus.Fields{
		"param": id,
	}).Info("repository: GetNumberPhoneByID")

	db := db.MariaDBInit()
	defer db.Close()

	query := "SELECT pn.id, pn.number, pn.type, " +
		"p.id, p.name, " +
		"pn.created_at, pn.created_by, IFNULL(pn.modified_at, ''), IFNULL(pn.modified_by, ''), pn.is_deleted " +
		"FROM phone_number AS pn " +
		"LEFT JOIN provider AS p ON pn.provider_id  = p.id " +
		"WHERE pn.is_deleted = 0 AND pn.id = '" + strconv.Itoa(id) + "' "

	logger.WithFields(logrus.Fields{"query": query}).Info("repository: GetNumberPhones-query")

	result, e := db.Query(query)
	if e != nil {
		return
	}

	defer result.Close()
	for result.Next() {
		if e = result.Scan(&data.ID, &data.Number, &data.Type,
			&data.Provider.ID, &data.Provider.Name,
			&data.CreatedAt, &data.CreatedBy, &data.ModifiedAt, &data.ModifiedBy, &data.IsDeleted); e != nil {
			return
		}
	}

	return
}

func GetNumberPhoneByNumber(c echo.Context, number string) (data entity.PhoneData, e error) {
	logger := md.GetLogger(c)
	logger.WithFields(logrus.Fields{
		"param": number,
	}).Info("repository: GetNumberPhoneByNumber")

	db := db.MariaDBInit()
	defer db.Close()

	query := "SELECT pn.id, pn.number, pn.type, " +
		"p.id, p.name, " +
		"pn.created_at, pn.created_by, IFNULL(pn.modified_at, ''), IFNULL(pn.modified_by, ''), pn.is_deleted " +
		"FROM phone_number AS pn " +
		"LEFT JOIN provider AS p ON pn.provider_id  = p.id " +
		"WHERE pn.is_deleted = 0 AND pn.number = '" + number + "' "

	logger.WithFields(logrus.Fields{"query": query}).Info("repository: GetNumberPhoneByNumber-query")

	result, e := db.Query(query)
	if e != nil {
		return
	}

	defer result.Close()
	for result.Next() {
		if e = result.Scan(&data.ID, &data.Number, &data.Type,
			&data.Provider.ID, &data.Provider.Name,
			&data.CreatedAt, &data.CreatedBy, &data.ModifiedAt, &data.ModifiedBy, &data.IsDeleted); e != nil {
			return
		}
	}

	return
}

func GetProviderByID(c echo.Context, id int) (data entity.Provider, e error) {
	logger := md.GetLogger(c)
	logger.WithFields(logrus.Fields{
		"param": id,
	}).Info("repository: GetProviderByID")

	db := db.MariaDBInit()
	defer db.Close()

	query := "SELECT p.id, p.code, p.name " +
		"FROM provider AS p " +
		"WHERE p.id = '" + strconv.Itoa(id) + "' "

	logger.WithFields(logrus.Fields{"query": query}).Info("repository: GetProviderByID-query")

	result, e := db.Query(query)
	if e != nil {
		return
	}

	defer result.Close()
	for result.Next() {
		if e = result.Scan(&data.ID, &data.Code, &data.Name); e != nil {
			return
		}
	}

	return
}

func GetProviderRandom(c echo.Context) (data entity.Provider, e error) {
	logger := md.GetLogger(c)
	logger.Info("repository: GetProviderRandom")

	db := db.MariaDBInit()
	defer db.Close()

	query := "SELECT p.id, p.code, p.name " +
		"FROM provider AS p ORDER BY RAND() LIMIT 1"

	logger.WithFields(logrus.Fields{"query": query}).Info("repository: GetProviderRandom-query")

	result, e := db.Query(query)
	if e != nil {
		return
	}

	defer result.Close()
	for result.Next() {
		if e = result.Scan(&data.ID, &data.Code, &data.Name); e != nil {
			return
		}
	}

	return
}

func UpdateNumberPhone(c echo.Context, params map[string]interface{}) (rowsAffected int64, e error) {
	logger := md.GetLogger(c)
	logger.WithFields(logrus.Fields{
		"params": params,
	}).Info("repository: UpdateNumberPhone")

	db := db.MariaDBInit()

	defer db.Close()

	ctx := context.Background()
	tx, e := db.BeginTx(ctx, nil)
	if e != nil {
		return
	}

	defer tx.Rollback()

	query := "UPDATE phone_number SET "
	i := 0
	var values []interface{}
	for key, value := range params {
		if key != "id" {
			query += key + " = ?"
			if (len(params) - 2) > i {
				query += ", "
			}
			i++
			values = append(values, value)
		}
	}

	query += " WHERE id = ? AND is_deleted = ?"
	values = append(values, params["id"], 0)

	logger.WithFields(logrus.Fields{"query": query}).Info("repository: UpdateNumberPhone-query")

	results, e := tx.ExecContext(ctx, query, values...)
	if e != nil {
		return
	}

	rowsAffected, e = results.RowsAffected()
	if e != nil {
		return
	}

	// Commit the transaction.
	if e = tx.Commit(); e != nil {
		return
	}

	return
}

func DeleteNumberPhoneByID(c echo.Context, id int, modifiedBy interface{}) (rowsAffected int64, e error) {
	logger := md.GetLogger(c)
	logger.WithFields(log.Fields{"di": id, "name": modifiedBy}).Info("repository: DeleteNumberPhoneByID")

	time := time.Now()
	db := db.MariaDBInit()
	defer db.Close()

	ctx := context.Background()
	tx, e := db.BeginTx(ctx, nil)
	if e != nil {
		return
	}

	defer tx.Rollback()
	results, e := tx.ExecContext(ctx, "UPDATE phone_number SET is_deleted = ?, modified_by = ?, modified_at = ? WHERE id = ? AND is_deleted = ?", 1, modifiedBy, time, id, 0)
	if e != nil {
		return
	}

	rowsAffected, e = results.RowsAffected()
	if e != nil {
		return
	}

	if e = tx.Commit(); e != nil {
		return
	}

	return
}
