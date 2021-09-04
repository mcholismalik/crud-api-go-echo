package middleware

import (
	res "codeid-boiler/pkg/util/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	var errCustom *res.Error

	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	switch report.Code {
	case http.StatusNotFound:
		errCustom = res.ErrorBuilder(&res.ErrorConstant.RouteNotFound, err)
	default:
		errCustom = res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	res.ErrorResponse(errCustom).Send(c)
}
