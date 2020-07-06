package repository

import (
	"database/sql"
	"fmt"
)

var (
	userName  string = "chenkai"
	password  string = "chenkai"
	ipAddress string = "192.168.0.115"
	port      int    = 3306
	dbName    string = "test"
	charset   string = "utf8"
)

func connectMysql() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddress, port, dbName, charset)
	Db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}
	return Db
}
