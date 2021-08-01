package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	fGet := func(c echo.Context) error {
		return c.String(http.StatusOK, "TEST GET")
	}
	fPost := func(c echo.Context) error {
		return c.String(http.StatusOK, "TEST POST")
	}

	bf := func(c echo.Context) error {
		fmt.Printf("before REST @ %s %s %s\n", c.RealIP(), c.Request().RequestURI, c.Request().Method)
		return nil
	}
	af := func(c echo.Context) error {
		fmt.Printf("after REST @ %s %s %s\n", c.RealIP(), c.Request().RequestURI, c.Request().Method)
		return nil
	}

	mGetWrap := map[string]func(c echo.Context) error{
		"/test1": echoWrap(fGet, bf, af),
		"/test2": echoWrap(fGet, bf, af),
	}
	mPostWrap := map[string]func(c echo.Context) error{
		"/test1": echoWrap(fPost, bf, af),
		"/test2": echoWrap(fPost, bf, af),
	}

	for k, v := range mGetWrap {
		e.GET(k, v)
	}
	for k, v := range mPostWrap {
		e.POST(k, v)
	}

	e.Logger.Fatal(e.Start(":1323"))
}

func echoWrap(f, before, after func(c echo.Context) error) func(c echo.Context) error {
	return func(c echo.Context) (err error) {
		if before != nil {
			if err = before(c); err != nil {
				return
			}
		}
		if f == nil {
			panic("The 1st parameter for echo function must be given")
		}
		if err = f(c); err != nil {
			return
		}
		if after != nil {
			err = after(c)
		}
		return
	}
}
