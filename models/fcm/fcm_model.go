package models

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"log"
	"fcm-golang/db"
)

type Gcms struct {
	Mdn  string `db:"MDN"`
	Reg_id string `json:"reg_id" form:"reg_id" query:"reg_id"`
	Device_model  string `json:"device_model" form:"device_model" query:"device_model"`
	First_login string `json:"first_login" form:"first_login" query:"first_login"`
	Last_login string `json:"last_login" form:"last_login" query:"last_login"`
}

func GetAllFcm() []Gcms{
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

func RegisterFcm(c echo.Context) *sqlx.Rows {
	u := new(Gcms)
	if err := c.Bind(u); err != nil {
		return nil
	}
	db:= db.CreateCon()
	sqlStatement := "INSERT INTO gcms (mdn, reg_id,device_model)VALUES (?, ?, ?)"
	res, _ := db.Queryx(sqlStatement, u.Mdn, u.Reg_id, u.Device_model)

	return res
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(res)
	//	return c.JSON(http.StatusCreated, u)
	//}
	//return c.String(http.StatusOK, "ok")
}

//func RegisterFcm(c echo.Context) *Gcms {
//	db := db.CreateCon()
//	reqBodyMap := echo.Map{}
//	if err := c.Bind(&reqBodyMap); err != nil {
//		return nil
//	}
//	device_model := reqBodyMap["device_model"]
//	reg_id := reqBodyMap["reg_id"]
//	session_id := reqBodyMap["session_id"]
//
//	defer db.Close()
//
//	_, err := db.Exec("insert into gcms values (?, ?, ?)", session_id, reg_id, device_model)
//	if err != nil {
//		return nil
//	}
//	return &Gcms{}
//}

//func InsertFcm() string{
	//db := db.CreateCon()
	//insertGcm := `INSERT INTO gcms (MDN, reg_id,device_model) VALUES (?, ?, ?)`
	//db.MustExec(insertGcm, "6288123214124","reg id sample","device model sample")
//}

