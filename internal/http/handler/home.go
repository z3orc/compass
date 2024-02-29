package handler

import (
	"fmt"
	"net/http"

	"github.com/common-nighthawk/go-figure"
	"github.com/z3orc/compass/internal/env"
)

func Home(w http.ResponseWriter, r *http.Request) {
	banner := figure.NewColorFigure("COMPASS", "slant", "cyan", true)
	figure.Write(w, banner)
	w.Write([]byte("\nWelcome! Compase is here to help you fetch whichever version of the Minecraft Server Jar you would like \nboth from a browser and from the command-line! The tool has been made to work with both vanilla, paper\nand purpur. The only thing that remain is giving it a try! Usage: /{flavour}/{version} \n"))
	w.Write([]byte(fmt.Sprintf("\n\nThis instance is current running v%s_%s.", env.Version, env.Build)))
}
