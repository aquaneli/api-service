package main

import (
	"service/internal/handler"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("api/:caption", handler.GetHandler)
	e.POST("api", handler.PostHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
