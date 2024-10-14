package handler

import (
	"net/http"
	"service/internal/model"

	"github.com/labstack/echo"
)

func GetHandler(e echo.Context) error {
	param := e.Param("caption")
	e.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	for _, v := range model.Data {
		if v.Caption == param {
			return e.JSON(http.StatusOK, map[string]interface{}{
				"caption": v.Caption,
				"weight":  v.Weight,
				"number":  v.Number,
			})
		}
	}
	return e.JSONPretty(http.StatusInternalServerError, map[string]interface{}{
		"status": http.StatusInternalServerError,
	}, "  ")
}
