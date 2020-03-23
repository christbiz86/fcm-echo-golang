package models

import (
	"github.com/labstack/echo"
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
var gcms Gcms

func GetAllFcm() []Gcms{
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
	reqBodyMap := echo.Map{}
	if err := c.Bind(&reqBodyMap); err != nil {
		return nil
	}
	device_model := reqBodyMap["device_model"]
	reg_id := reqBodyMap["reg_id"]
	session_id := reqBodyMap["session_id"]
	var u = &Gcms{
		Mdn:          session_id,
		Reg_id:       reg_id,
		Device_model: device_model,
	}
	return u
}

//func InsertFcm() string{
	//db := db.CreateCon()
	//insertGcm := `INSERT INTO gcms (MDN, reg_id,device_model) VALUES (?, ?, ?)`
	//db.MustExec(insertGcm, "6288123214124","reg id sample","device model sample")
//}

