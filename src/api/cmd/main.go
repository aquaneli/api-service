package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

type Items struct {
	caption string
	weight  float32
	number  int
}

var items []Items

func main() {
	e := echo.New()

	group := e.Group("/item")
	group.GET("/:caption", GetHandler)
	group.POST("/cap", PostHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func GetHandler(e echo.Context) error {
	param := e.Param("caption")
	for _, v := range items {
		if v.caption == param {
			return e.JSON(http.StatusOK, v)
		}
	}
	return e.String(http.StatusInternalServerError, "{\n    \"status_err\": 500\n}")
}

func GetData() {

}

func PostHandler(e echo.Context) error {
	item, err := Parsing(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "{\n\"status_err\": 400\n}")
	}
	items = append(items, *item)
	return e.JSON(http.StatusInternalServerError, "{\n\"status_ok\": 200\n}")
}

func Parsing(e echo.Context) (*Items, error) {
	decoder := json.NewDecoder(e.Request().Body)
	decoder.Token()
	item := Items{}
	for decoder.More() {
		key, err := decoder.Token()
		if err != nil {
			return nil, err
		}
		value, err := decoder.Token()
		if err != nil {
			return nil, err
		}

		switch key {
		case "caption":
			val, ok := value.(string)
			if !ok {
				return nil, err
			}
			item.caption = val
		case "weight":
			val, ok := value.(float32)
			if !ok {
				return nil, err
			}
			item.weight = val
		case "number":
			val, ok := value.(int)
			if !ok {
				return nil, err
			}
			item.number = val
		}

	}
	return &item, nil
}
