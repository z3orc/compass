package routes

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/z3orc/compass/internal/model"
	"github.com/z3orc/compass/internal/repo"
	"github.com/z3orc/compass/internal/util"
)

func VersionHandler(c echo.Context, repo repo.IVersionRepository, flavour model.Flavour) error {
	id := c.Param("id")

	version, err := repo.GetVersion(flavour, id)
	if err != nil {
		return c.JSON(int(err.StatusCode()), util.ErrorToJson(err))
	}

	versionJson, errJson := version.ToJson()
	if errJson != nil {
		log.Error("Error converting version to json", "error", errJson)
		return c.JSON(http.StatusInternalServerError, util.ErrorToJson(errJson))
	}

	return c.JSONBlob(200, versionJson)
}
