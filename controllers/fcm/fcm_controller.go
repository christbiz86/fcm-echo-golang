package controllers

import (
	models "fcm-golang/models/fcm"
	"github.com/labstack/echo"
	"net/http"
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

//func DeleteFcm(c echo.Context) error{
//	mdn := c.Param("mdn")
//	sqlStatement := "DELETE FROM employees WHERE MDN = ?"
//	res, err := db.Query(sqlStatement, mdn)
//	if err != nil {
//		fmt.Println(err)
//		//return c.JSON(http.StatusCreated, u);
//	} else {
//		fmt.Println(res)
//		return c.JSON(http.StatusOK, "Deleted")
//	}
//	return c.String(http.StatusOK, id+"Deleted")
//}

//
//func UpdateFcm() error{
//
//}