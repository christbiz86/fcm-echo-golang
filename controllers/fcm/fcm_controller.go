package controllers

import(
	"github.com/labstack/echo"
	"net/http"
	"fcm-golang/models/fcm"
)

func GetFcm(c echo.Context) error{
	result := models.GetAllFcm()
	var jsonResult map[string]interface{}
	jsonResult = make(map[string]interface{})
	jsonResult["status"] = "Sukses"
	jsonResult["result"] = result
	return c.JSON(http.StatusOK, jsonResult)
}

func CheckFcm(c echo.Context) error{
	return c.String(http.StatusOK, "Hello, post FCM!")
}

//func RegisterFcm() error{
//
//}
//
//func UpdateFcm() error{
//
//}