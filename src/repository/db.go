package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-rbac-demo/custerror"
)

var (
	userName  string = "root"
	password  string = "root1"
	ipAddress string = "localhost"
	port      int    = 3306
	dbName    string = "db_admin"
	charset   string = "utf8"
)

func connectMysql() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddress, port, dbName, charset)
	conn, err := sql.Open("mysql", dsn)

	return conn, custerror.NewError(err)

}
