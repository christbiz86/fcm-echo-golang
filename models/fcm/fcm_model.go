package models

import (
	"log"
	"fcm-golang/db"
)

type Gcms struct {
	Mdn  string `db:"MDN"`
	Reg_id string
	Device_model  string
	First_login string
	Last_login string
}

var gcmsList []Gcms

func GetAllFcm() []Gcms{
	gcms := Gcms{}
	db := db.CreateCon()
	rows, err := db.Queryx("select * from gcms")
	if err != nil {
		log.Printf("%v\n", err)
	}
	for rows.Next() {
		err := rows.StructScan(&gcms)
		if err != nil {
			log.Printf("%v\n", err)
		}
		gcmsList = append(gcmsList, gcms)
	}
	return gcmsList
}