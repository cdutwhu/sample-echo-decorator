package main

import (
	"fmt"
	"net/http"

	m "github.com/cdutwhu/sample-echo-decorator/method"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	bf := func(c echo.Context) error {
		s := fmt.Sprintf("before Method @ %s %s %s\n", c.RealIP(), c.Request().RequestURI, c.Request().Method)
		return c.String(http.StatusOK, s) //errors.New("error in bf")
	}
	af := func(c echo.Context) error {
		s := fmt.Sprintf("after Method @ %s %s %s\n", c.RealIP(), c.Request().RequestURI, c.Request().Method)
		return c.String(http.StatusOK, s)
	}

	wrap(e.GET, m.MapGetAPI, bf, af)
	wrap(e.POST, m.MapPostAPI, bf, af)

	e.Logger.Fatal(e.Start(":80"))
}
