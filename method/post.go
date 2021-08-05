package method

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func FnPost1(c echo.Context) error {
	return c.String(http.StatusOK, "TEST POST1\n")
}

func FnPost2(c echo.Context) error {
	return c.String(http.StatusOK, "TEST POST2\n")
}

var MapPostAPI = map[string]func(c echo.Context) error{
	"/test-post1": FnPost1,
	"/test-post2": FnPost2,
}
