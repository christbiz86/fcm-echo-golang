package main

import (
	"fcm-golang/controllers/fcm"
	"fcm-golang/db"
	models "fcm-golang/models/fcm"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/fcm",controllers.GetFcm)
	e.POST("/fcm", func(c echo.Context) error {
		u := new(models.Gcms)
		if err := c.Bind(u); err != nil {
			return err
		}
		db.CreateCon()
		sqlStatement := "INSERT INTO gcms (mdn, reg_id,device_model)VALUES ($1, $2, $3)"
		res, err := db.Queryx(sqlStatement, u.Mdn, u.Reg_id, u.Device_model)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, u)
		}
		return c.String(http.StatusOK, "ok")
	})
	//e.POST("/fcm", controllers.RegisterFcm)
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