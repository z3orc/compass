package main

import (
	"net/http"

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

	log.Info("Starting COMPASS🧭")

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.Secure())
	e.Use(web.LoggerMiddleware())
	e.Use(middleware.Recover())
	// e.Use(middleware.HTTPSRedirect())

	repo := repo.NewVersionRepository(data.NewPistonDataSource())

	//Users can use either piston or vanilla to get vanilla jars
	e.Pre(middleware.Rewrite(map[string]string{
		"/vanilla/*": "/piston/$1",
	}))

	e.GET("/piston/:id", func(c echo.Context) error {
		return versionHandler(c, repo, model.FlavourPiston)
	})

	e.GET("/paper/:id", func(c echo.Context) error {
		return versionHandler(c, repo, model.FlavourPaper)
	})

	e.GET("/purpur/:id", func(c echo.Context) error {
		return versionHandler(c, repo, model.FlavourPurpur)
	})

	log.Info("Running webserver on port 8000🚀")
	log.Fatal(e.Start(":8000"))
}

func versionHandler(c echo.Context, repo repo.IVersionRepository, flavour model.Flavour) error {
	id := c.Param("id")

	version, err := repo.GetVersion(flavour, id)
	if err != nil {
		return c.JSON(int(err.StatusCode()), util.ErrorToJson(err))
	}

	versionJson, errJson := version.ToJson()
	if errJson != nil {
		return c.JSON(http.StatusInternalServerError, util.ErrorToJson(errJson))
	}

	return c.Blob(200, "application/json", versionJson)
}
