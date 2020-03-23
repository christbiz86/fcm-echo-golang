package controllers

import(
	"github.com/labstack/echo"
	"net/http"
	"fcm-golang/models/fcm"
)

type test_struct struct {
	Mdn  string `db:"MDN"`
	Reg_id string `json:"reg_id" form:"reg_id" query:"reg_id"`
	Device_model  string `json:"device_model" form:"device_model" query:"device_model"`
	First_login string `json:"first_login" form:"first_login" query:"first_login"`
	Last_login string `json:"last_login" form:"last_login" query:"last_login"`
}

func GetFcm(c echo.Context) error{
	result := models.GetAllFcm()
	var jsonResult map[string]interface{}
	jsonResult = make(map[string]interface{})
	jsonResult["status"] = "1"
	jsonResult["message"] = "Success"
	jsonResult["result"] = result
	return c.JSON(http.StatusOK, jsonResult)
}

//func RegisterFcm(c echo.Context) error{
//	result := models.RegisterFcm(c)
//	return c.JSON(http.StatusCreated, result)
//}
//
//func UpdateFcm() error{
//
//}