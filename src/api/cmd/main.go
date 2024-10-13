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
	group.POST("", PostHandler)
	e.Logger.Fatal(e.Start(":8080"))
}

func GetHandler(e echo.Context) error {
	param := e.Param("caption")
	e.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	for _, v := range items {
		if v.caption == param {
			return e.JSON(http.StatusOK, map[string]interface{}{
				"caption": v.caption,
				"weight":  v.weight,
				"number":  v.number,
			})
		}
	}
	return e.JSONPretty(http.StatusInternalServerError, map[string]interface{}{
		"status": http.StatusInternalServerError,
	}, "  ")
}

func PostHandler(e echo.Context) error {
	item, err := Parsing(e)
	e.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	if err != nil {
		return e.JSONPretty(http.StatusInternalServerError, map[string]interface{}{
			"status": http.StatusBadRequest,
		}, "  ")
	}
	items = append(items, item)
	return e.JSONPretty(http.StatusInternalServerError, map[string]interface{}{
		"status": http.StatusOK,
	}, "  ")
}

func Parsing(e echo.Context) (Items, error) {
	decoder := json.NewDecoder(e.Request().Body)
	_, err := decoder.Token()
	if err != nil {
		return Items{}, err
	}
	item := Items{}
	for decoder.More() {
		key, err := decoder.Token()
		if err != nil {
			return Items{}, err
		}

		switch key {
		case "caption":
			err := decoder.Decode(&item.caption)
			if err != nil {
				return Items{}, err
			}
		case "weight":
			err := decoder.Decode(&item.weight)
			if err != nil {
				return Items{}, err
			}
		case "number":
			err := decoder.Decode(&item.number)
			if err != nil {
				return Items{}, err
			}
		}

	}

	return item, nil
}
