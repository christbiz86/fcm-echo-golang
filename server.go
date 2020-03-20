package main

import (
	"github.com/labstack/echo"
	//"fcm-golang/db"
	"fcm-golang/controllers/fcm"
)

func main() {
	e := echo.New()
	//db.CreateCon()
	e.GET("/fcm",controllers.GetFcm)
	e.POST("/fcm",controllers.CheckFcm)
	e.Logger.Fatal(e.Start(":1323"))
}