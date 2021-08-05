package method

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func FnGet1(c echo.Context) error {
	return c.String(http.StatusOK, "TEST GET1\n")
}

func FnGet2(c echo.Context) error {
	return c.String(http.StatusOK, "TEST GET2\n")
}

var MapGetAPI = map[string]func(c echo.Context) error{
	"/test-get1": FnGet1,
	"/test-get2": FnGet2,
}
