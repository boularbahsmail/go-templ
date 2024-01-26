package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(context echo.Context, component templ.Component) error {
	return component.Render(context.Request().Context(), context.Response())
}
