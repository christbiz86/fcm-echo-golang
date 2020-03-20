package models

import (
	"time"
	"github.com/jinzhu/gorm"
	//"fcm-golang/db"
)

type Gcms struct {
	gorm.Model
	Mdn  string
	Reg_id string
	Device_model  string
	First_login time.Time
	Last_login time.Time
}

var gcms Gcms

func GetAllFcm(){
	db, err := gorm.Open("mysql", "custinfo:custinfodev@(10.16.5.162:3306)/custinfo")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.Find(&gcms)
}