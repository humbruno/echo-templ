package api

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/humbruno/echo-templ/views/templates"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// //go:embed public
// var publicAssets embed.FS
//
// // GetPublicAssetsFileSystem returns a http.FileSystem for the public assets so that
// // we can embed them into the binary so it is self-contained.
// func GetPublicAssetsFileSystem() http.FileSystem {
// 	fsys, err := fs.Sub(publicAssets, "public")
// 	if err != nil {
// 		panic(err)
// 	}
// 	return http.FS(fsys)
// }

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
