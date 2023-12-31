package mariadb

import (
	"database/sql"
	"fmt"

	cm "github.com/absormu/go-jti/pkg/configuration"
	_ "github.com/go-sql-driver/mysql"
)

func MariaDBInit() *sql.DB {

	db, err := sql.Open("mysql", cm.Config.MariaDBUser+":"+cm.Config.MariaDBPassword+"@tcp("+cm.Config.MariaDBAddr+":"+cm.Config.MariaDBPort+")/"+cm.Config.MariaDBDatabase+"?loc=Asia%2FJakarta")
	if err != nil {
		panic(err)
	}

	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return db
}
