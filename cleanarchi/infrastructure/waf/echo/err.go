package echo

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/po3rin/cleanarchi/infrastructure/waf/echo/response"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "Internal Server Error"
	if e, ok := err.(*echo.HTTPError); ok {
		code = e.Code
		message = e.Message.(string)
	}
	body := response.Error{
		StatusCode: code,
		Message:    message,
	}
	c.JSON(code, body)
}
