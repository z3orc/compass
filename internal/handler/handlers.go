package handler

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/z3orc/compass/internal/model"
	"github.com/z3orc/compass/internal/repo"
	"github.com/z3orc/compass/internal/util"
)

func HomeHandler(c echo.Context) error {
	content := []byte(`
	Welcome to COMPASS! Compass is here to help you fetch whichever version 
	of the Minecraft Server Jar you would like, both from a browser and 
	from the command-line! It current only supports server jars from mojang,
	but support for other sources will be added in the future! The only thing 
	that remain now is giving it a try!
	
	Usage: .../{flavour}/{version}, e.g. /piston/1.21 for vanilla 1.21`)

	return c.Blob(http.StatusOK, "text/plain", content)
}

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

	return c.Blob(200, "application/json", versionJson)
}
