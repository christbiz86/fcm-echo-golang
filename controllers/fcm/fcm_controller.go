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
	var jsonResult map[string]interface{}
	jsonResult = make(map[string]interface{})
	jsonResult["status"] = "1"
	if result != nil {
		jsonResult["message"] = "Data insert successfully!"
		jsonResult["result"] = result
		return c.JSON(http.StatusCreated, jsonResult)
	} else {
		jsonResult["message"] = "Failed insert new data!"
		return c.JSON(http.StatusBadRequest, jsonResult)
	}
}

func GetFcmById(c echo.Context) error {
	result := models.GetFcmById(c)
	var jsonResult map[string]interface{}
	jsonResult = make(map[string]interface{})
	jsonResult["status"] = "1"
	if result != nil {
		jsonResult["message"] = "Data retrieve successfully!"
		jsonResult["result"] = result
		return c.JSON(http.StatusOK, jsonResult)
	} else {
		jsonResult["message"] = "Failed retrieved row!"
		return c.JSON(http.StatusBadRequest, jsonResult)
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

func DeleteFcm(c echo.Context) error{
	result := models.DeleteGcm(c)
	var jsonResult map[string]interface{}
	jsonResult = make(map[string]interface{})
	jsonResult["status"] = "1"
	if result != nil {
		jsonResult["message"] = "Data deleted successfully!"
		jsonResult["result"] = result
		return c.JSON(http.StatusOK, jsonResult)
	} else {
		jsonResult["message"] = "Failed deleted row!"
		return c.JSON(http.StatusBadRequest, jsonResult)
	}
}