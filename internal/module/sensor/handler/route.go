package handler

import "github.com/labstack/echo/v4"

func NewSensorRoute(h Handler, route *echo.Group) {
	sensor := route.Group("/streaming")
	sensor.GET("/", h.Health)
	sensor.POST("/sensor-generate", h.GenerateStream)
}
