package controllers

import(
	"fcm-golang/db"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"fcm-golang/models/fcm"
)

func GetFcm(c echo.Context) error{
	result := models.GetAllFcm()
	var jsonResult map[string]interface{}
	jsonResult = make(map[string]interface{})
	jsonResult["status"] = "1"
	jsonResult["message"] = "Success"
	jsonResult["result"] = result
	return c.JSON(http.StatusOK, jsonResult)
}

func RegisterFcm(c echo.Context) error{
	//u := new(models.Gcms)
	//if err := c.Bind(u); err != nil {
	//	return err
	//}
	db:= db.CreateCon()
	//sqlStatement := "INSERT INTO gcms (mdn, reg_id,device_model)VALUES (?, ?, ?)"
	//res, err := db.Queryx(sqlStatement, u.Mdn, u.Reg_id, u.Device_model)


	//result := models.RegisterFcm(c)
	//return c.JSON(http.StatusOK, result)
	u := new(models.Gcms)
	if err := c.Bind(u); err != nil {
		return err
	}
	sqlStatement := "INSERT INTO gcms (mdn, reg_id,device_model)VALUES ($1, $2, $3)"
	res, err := db.Queryx(sqlStatement, u.Mdn, u.Reg_id, u.Device_model)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusCreated, u)
	}
	return c.String(http.StatusOK, "ok")

}
//
//func UpdateFcm() error{
//
//}