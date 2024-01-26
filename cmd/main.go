package main

import (
	"context"

	"github.com/boularbahsmail/go-templ/handler"
	"github.com/labstack/echo/v4"
)

// Echo func(context) error
// Fiber func(context) error

func main() {
	// Init a new echo app
	app := echo.New()

	// Routes
	userHandler := handler.UserHandler{}
	// Injecting a middleware
	app.Use(withUser)
	app.GET("/users", userHandler.HandlerUserShow)

	app.Start(":3000")
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "User", "boularbahismail01@gmail.com")
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
