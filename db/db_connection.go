package db

import (
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func CreateCon() *sqlx.DB {
	db, err := sqlx.Open("mysql", "custinfo:custinfodev@(10.16.5.162:3306)/custinfo")
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}
