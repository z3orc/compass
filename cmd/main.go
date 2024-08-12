package main

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/z3orc/compass/internal/data"
	"github.com/z3orc/compass/internal/model"
	"github.com/z3orc/compass/internal/repo"
	"github.com/z3orc/compass/internal/util"
)

func main() {
	log.SetLevel(log.DebugLevel)

	e := echo.New()
	e.GET("/piston/:id", pistonHandler)
	log.Fatal(e.Start(":8000"))
}

func pistonHandler(c echo.Context) error {
	id := c.Param("id")

	src := data.NewPistonDataSource()
	repo := repo.NewVersionRepository(src)

	version, err := repo.GetVersion(model.FlavourPiston, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.ErrorToJson(err))
	}

	return c.JSON(200, version)
}
