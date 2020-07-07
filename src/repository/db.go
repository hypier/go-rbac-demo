package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	userName  string = "root"
	password  string = "root"
	ipAddress string = "localhost"
	port      int    = 3306
	dbName    string = "db_admin"
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

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
