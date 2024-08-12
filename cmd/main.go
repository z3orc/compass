package main

import (
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/z3orc/compass/internal/data"
	"github.com/z3orc/compass/internal/model"
	"github.com/z3orc/compass/internal/repo"
	"github.com/z3orc/compass/internal/util"
	"github.com/z3orc/compass/internal/web"
)

func main() {
	log.SetLevel(log.DebugLevel)

	e := echo.New()
	e.Use(web.LoggerMiddleware())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	// e.Use(middleware.HTTPSRedirect())

	repo := repo.NewVersionRepository(data.NewPistonDataSource())

	e.GET("/piston/:id", func(c echo.Context) error {
		return pistonHandler(c, repo)
	})

	log.Fatal(e.Start(":8000"))
}

func pistonHandler(c echo.Context, r repo.IVersionRepository) error {
	id := c.Param("id")

	version, err := r.GetVersion(model.FlavourPiston, id)
	if err != nil {
		return c.JSON(int(err.StatusCode()), util.ErrorToJson(err))
	}

	return c.JSON(200, version)
}
