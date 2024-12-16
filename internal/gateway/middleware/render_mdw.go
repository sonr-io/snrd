package middleware

import (
	"bytes"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/views"
	config "github.com/onsonr/sonr/pkg/config/hway"
)

type RenderContext struct {
	echo.Context
	env config.Hway
}

func UseRender(env config.Hway) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &RenderContext{Context: c, env: env}
			return next(ctx)
		}
	}
}

func Render(c echo.Context, cmp templ.Component) error {
	// Create a buffer to store the rendered HTML
	buf := &bytes.Buffer{}
	// Render the component to the buffer
	err := cmp.Render(c.Request().Context(), buf)
	if err != nil {
		return err
	}

	// Set the content type
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	// Write the buffered content to the response
	_, err = c.Response().Write(buf.Bytes())
	if err != nil {
		return err
	}
	c.Response().WriteHeader(200)
	return nil
}

func RenderError(c echo.Context, err error) error {
	return Render(c, views.ErrorView(err.Error()))
}

func RenderInitial(c echo.Context) error {
	return Render(c, views.InitialView())
}

func RenderLoading(c echo.Context) error {
	return Render(c, views.LoadingView())
}
