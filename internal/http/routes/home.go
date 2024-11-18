package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
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
