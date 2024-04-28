package api

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/humbruno/echo-templ/views/templates"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Static("/public", "./api/public")
	e.GET("/", handleHome)
	e.Logger.Fatal(e.Start(":8000"))
}

func render(ctx echo.Context, component templ.Component) error {
	return component.Render(ctx.Request().Context(), ctx.Response())
}

func handleHome(c echo.Context) error {
	return render(c, templates.Home("bruno"))
}
