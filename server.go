package main

import (
	controllers "fcm-golang/controllers/fcm"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/fcm",controllers.GetFcm)
	e.POST("/fcm", controllers.RegisterFcm)
	//e.POST("/test",controllers.Test)
	//e.POST("/fcm", func(c echo.Context) (err error) {
	//	m := echo.Map{}
	//	if err := c.Bind(&m); err != nil {
	//		return err
	//	}
	//	return c.JSON(200, m)
	//})
	e.Logger.Fatal(e.Start(":1323"))
}