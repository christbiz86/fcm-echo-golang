package models

import (
	"fcm-golang/db"
	"fmt"
	"log"
	"github.com/labstack/echo"
	"time"
)

type Gcms struct {
	Mdn          string `db:"MDN"`
	Reg_id       string `json:"reg_id" form:"reg_id" query:"reg_id"`
	Device_model string `json:"device_model" form:"device_model" query:"device_model"`
	First_login  string `json:"first_login" form:"first_login" query:"first_login"`
	Last_login   string `json:"last_login" form:"last_login" query:"last_login"`
}

var currentTime = time.Now()

func GetAllFcm() []Gcms {
	var gcmsList []Gcms
	var gcms Gcms
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

func RegisterFcm(c echo.Context) *Gcms {
	u := new(Gcms)
	if err := c.Bind(u); err != nil {
		log.Println(err)
		return nil
	}
	db := db.CreateCon()

	sqlStatement1 := "SELECT mdn FROM gcms where mdn = ?"
	_, check := db.Queryx(sqlStatement1, u.Mdn)
	if check != nil {
		sqlStatement := "UPDATE gcms set reg_id=?, device_model=?, last_login=? where id=?"
		_, err := db.Queryx(sqlStatement, u.Reg_id, u.Device_model, currentTime.Format("2006-01-02 15:04:05"),u.Mdn)
		if err != nil {
			fmt.Println("no error")
		} else {
			fmt.Println("err")
		}
		return u
	} else {
		sqlStatement := "INSERT INTO gcms (mdn, reg_id,device_model,first_login)VALUES (?, ?, ?, ?)"
		_, err := db.Queryx(sqlStatement, u.Mdn, u.Reg_id, u.Device_model, currentTime.Format("2006-01-02 15:04:05"))
		if err != nil {
			fmt.Println("no error")
		} else {
			fmt.Println("err")
		}
		return u
	}

}