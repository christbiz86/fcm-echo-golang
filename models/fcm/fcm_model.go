package models

import (
	sql "database/sql"
	"fcm-golang/db"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"time"
)

type Gcms struct {
	Mdn         string `db:"MDN"`
	Reg_id       string `json:"reg_id" form:"reg_id" query:"reg_id"`
	Device_model string `json:"device_model" form:"device_model" query:"device_model"`
	First_login  string `json:"first_login" form:"first_login" query:"first_login"`
	Last_login   string `json:"last_login" form:"last_login" query:"last_login"`
}

var currentTime = time.Now()

func GetMdnCookie(r *http.Request) string {
	c, _ := r.Cookie("CookieData")
	return c.Value
}

func GetAllFcm() []Gcms {
	var gcmsList []Gcms
	var gcms Gcms
	dbConn := db.CreateCon()
	rows, err := dbConn.Queryx("select * from gcms")
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
	mdn := GetMdnCookie(c.Request())

	u := new(Gcms)
	if err := c.Bind(u); err != nil {
		return nil
	}
	dbConn := db.CreateCon()
	sqlStatement1 := "SELECT mdn FROM gcms where mdn = ?"
	//row := db.QueryRow(sqlStatement1, u.Mdn)
	row := dbConn.QueryRow(sqlStatement1, mdn)
	switch err := row.Scan(&u.Mdn); err {
		case sql.ErrNoRows:
			InsetNewGcm(u)
		case nil:
			UpdateGcm(u)
	}
	return u
}

func InsetNewGcm(u *Gcms) {
	dbConn := db.CreateCon()
	sqlStatement := "INSERT INTO gcms (mdn, reg_id,device_model,first_login)VALUES (?, ?, ?, ?)"
	dbConn.Queryx(sqlStatement, u.Mdn, u.Reg_id, u.Device_model, currentTime.Format("2006-01-02 15:04:05"))
}

func UpdateGcm(u *Gcms){
	dbConn := db.CreateCon()
	sqlStatement := "UPDATE gcms set reg_id=?, device_model=?, last_login=? where MDN=?"
	dbConn.Queryx(sqlStatement, u.Reg_id, u.Device_model, currentTime.Format("2006-01-02 15:04:05"),u.Mdn)
}