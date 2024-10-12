package main

import (
	"encoding/json"
	"net/http"
	"reflect"

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
	items = append(items, item)
	return e.JSON(http.StatusInternalServerError, "{\n\"status_ok\": 200\n}")
}

func Parsing(e echo.Context) (Items, error) {
	decoder := json.NewDecoder(e.Request().Body)
	decoder.Token()
	item := Items{}
	for decoder.More() {
		key, err := decoder.Token()
		if err != nil {
			return Items{}, err
		}
		value, err := decoder.Token()
		if err != nil {
			return Items{}, err
		}

		switch key {
		case "caption":
			val := reflect.ValueOf(value)
			item.caption = val.String()
		case "weight":
			val := reflect.ValueOf(value)
			ok := val.CanFloat()
			if !ok {
				return Items{}, err
			}
			item.weight = float32(val.Float())
		case "number":
			val := reflect.ValueOf(value)
			ok := val.CanInt()
			if !ok {
				return Items{}, err
			}
			item.number = int(val.Int())
		}

	}
	return item, nil
}
