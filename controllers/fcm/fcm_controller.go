package controllers

import(
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
	//db:= db.CreateCon()
	//sqlStatement := "INSERT INTO gcms (mdn, reg_id,device_model)VALUES (?, ?, ?)"
	//res, err := db.Queryx(sqlStatement, u.Mdn, u.Reg_id, u.Device_model)

	result := models.RegisterFcm(c)
	if result != nil {
		fmt.Println(result)
	} else {
		fmt.Println(result)
		return c.JSON(http.StatusCreated, result)
	}
	return c.String(http.StatusOK, "ok")
}
//
//func UpdateFcm() error{
//
//}