package db

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func CreateCon() *sql.DB {
	db, err := gorm.Open("mysql", "custinfo:custinfodev@(10.16.5.162:3306)/custinfo")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	//return db
}
