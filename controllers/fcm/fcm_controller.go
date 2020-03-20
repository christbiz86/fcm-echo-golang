package controllers

import(
	"github.com/labstack/echo"
	"net/http"
	"fcm-golang/models/fcm"
)

func GetFcm(c echo.Context) error{
	result := models.GetAllFcm
	return c.Render(http.StatusOK, "list data", result)
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