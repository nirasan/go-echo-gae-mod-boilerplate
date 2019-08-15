package app

import (
	"net/http"

	"github.com/labstack/echo"
)

func IndexHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", "hello world")
}
