package main

import (
	"service/internal/handler"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("item/:caption", handler.GetHandler)
	e.POST("item", handler.PostHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
