package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"service/internal/model"

	"github.com/labstack/echo"
)

func PostHandler(e echo.Context) error {
	item, err := Parsing(e)
	e.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	if err != nil {
		return e.JSONPretty(http.StatusInternalServerError, map[string]interface{}{
			"status": http.StatusBadRequest,
		}, "  ")
	}
	model.Data = append(model.Data, item)
	return e.JSONPretty(http.StatusInternalServerError, map[string]interface{}{
		"status": http.StatusOK,
	}, "  ")
}

func Parsing(e echo.Context) (model.Items, error) {
	decoder := json.NewDecoder(e.Request().Body)
	_, err := decoder.Token()
	if err != nil {
		return model.Items{}, err
	}
	item := model.Items{}
	for decoder.More() {
		key, err := decoder.Token()
		if err != nil {
			return model.Items{}, err
		}

		switch key {
		case "caption":
			err := decoder.Decode(&item.Caption)
			if err != nil {
				return model.Items{}, err
			}
		case "weight":
			err := decoder.Decode(&item.Weight)
			if err != nil {
				return model.Items{}, err
			}
		case "number":
			err := decoder.Decode(&item.Number)
			if err != nil {
				return model.Items{}, err
			}
		default:
			return model.Items{}, errors.New("incorrect data")
		}
	}

	return item, nil
}
