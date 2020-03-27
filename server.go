package main

import (
	controllers "fcm-golang/controllers/fcm"
	"fcm-golang/middlewares"
	//"fcm-golang/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"
)

type jwtCustomClaims struct {
	SessionMdn  string `json:"sessionMdn"`
	jwt.StandardClaims
}

func main() {
	e := echo.New()
	//config := middleware.JWTConfig{
	//	Claims:	&jwtCustomClaims{},
	//	SigningKey: []byte(services.GoDotEnvVariable("JWT_KEY")),
	//}
	//e.Use(middleware.JWTWithConfig(config))
	e.Use(echo.WrapMiddleware(middlewares.SessionMiddleware))
	e.GET("/fcm",controllers.GetFcm)
	e.POST("/fcm", controllers.RegisterFcm)
	//e.DELETE("/fcm/:mdn", controllers.DeleteFcm)
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