package gateway

import (
	"github.com/labstack/echo/v4"
)

func InterceptRoutes(e *echo.Echo) error {
	gw, err := New()
	if err != nil {
		return err
	}
	e.Any("/*", gw.Handler())
	return nil
}
