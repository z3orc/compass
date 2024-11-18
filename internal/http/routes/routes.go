package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/z3orc/compass/internal/model"
	"github.com/z3orc/compass/internal/repo"
)

func RegisterRoutes(e *echo.Echo, repo repo.IVersionRepository) {

	//Users can use either piston or vanilla to get vanilla jars
	e.Pre(middleware.Rewrite(map[string]string{
		"/vanilla/*": "/piston/$1",
	}))

	e.GET("/", HomeHandler)

	e.GET("/piston/:id", func(c echo.Context) error {
		return VersionHandler(c, repo, model.FlavourPiston)
	})

	// e.GET("/paper/:id", func(c echo.Context) error {
	// 	return versionHandler(c, repo, model.FlavourPaper)
	// })

	// e.GET("/purpur/:id", func(c echo.Context) error {
	// 	return versionHandler(c, repo, model.FlavourPurpur)
	// })
}
