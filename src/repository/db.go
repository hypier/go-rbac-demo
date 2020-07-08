package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-rbac-demo/custerror"
)

var (
	userName  string = "root"
	password  string = "root"
	ipAddress string = "localhost"
	port      int    = 3306
	dbName    string = "db_admin"
	charset   string = "utf8"
)

func connectMysql() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddress, port, dbName, charset)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, custerror.NewError(err)
	}

	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(3)
	_ = conn.Ping()

	return conn, nil

}
