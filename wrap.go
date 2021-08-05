package main

import "github.com/labstack/echo/v4"

func echoWrap(f, before, after func(c echo.Context) error) func(c echo.Context) error {
	return func(c echo.Context) (err error) {
		if before != nil {
			if err = before(c); err != nil {
				return
			}
		}
		if f == nil {
			panic("The 1st parameter for echo method must be given")
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

func wrap(f func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route,
	mUrlApi map[string]func(c echo.Context) error,
	bf, af func(c echo.Context) error) {

	for k, v := range mUrlApi {
		f(k, echoWrap(v, bf, af))
	}
}
