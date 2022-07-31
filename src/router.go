package main

import (
	"github.com/labstack/echo/v4"
)

func NewRouter(init *Initialize)(e *echo.Echo){
	e = echo.New()
	e.GET("/",check)
	e.POST("/callback", init.lineHandler.EventHandler)
	// batch.weather.Getみたいにしたい
	e.POST("/weather", init.lineBatch.GetWeather)
	return e
}

func check(c echo.Context) error{
	return c.String(200, "Hello world")
}
