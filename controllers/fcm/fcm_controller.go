package controllers

import (
	models "fcm-golang/models/fcm"
	"net/http"
	"github.com/labstack/echo"
)

func GetFcm(c echo.Context) error {
	result := models.GetAllFcm()
	var jsonResult map[string]interface{}
	jsonResult = make(map[string]interface{})
	jsonResult["status"] = "1"
	jsonResult["message"] = "Success"
	jsonResult["result"] = result
	return c.JSON(http.StatusOK, jsonResult)
}

func RegisterFcm(c echo.Context) error {
	result := models.RegisterFcm(c)
	if result != nil {
		var jsonResult map[string]interface{}
		jsonResult = make(map[string]interface{})
		jsonResult["status"] = "1"
		jsonResult["message"] = "Data insert successfully!"
		jsonResult["result"] = result
		return c.JSON(http.StatusCreated, jsonResult)
	} else {
		return c.String(http.StatusOK, "Failed insert new data")
	}
}

//
//func UpdateFcm() error{
//
//}