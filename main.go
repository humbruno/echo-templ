package main

import (
	"github.com/a-h/templ"
	"github.com/humbruno/echo-templ/templates"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", handleHome)
	e.Logger.Fatal(e.Start(":8000"))
}

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func handleHome(c echo.Context) error {
	return render(c, templates.Home("bruno"))
}
