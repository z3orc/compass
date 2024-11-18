package main

import (
	"flag"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/z3orc/compass/internal/data"
	"github.com/z3orc/compass/internal/http"
	"github.com/z3orc/compass/internal/http/middleware"
	"github.com/z3orc/compass/internal/http/routes"
	"github.com/z3orc/compass/internal/repo"
)

var port int
var debug bool
var https bool
var ratelimit int

func main() {
	log.Info("Starting COMPASS🧭")

	flag.IntVar(&port, "port", 8000, "port for webserver to listen on")
	flag.IntVar(&ratelimit, "ratelimit", 31, "max amount of requests per minute. 0 to disable")
	flag.BoolVar(&debug, "debug", false, "enable debug logging")
	flag.BoolVar(&https, "https", false, "enable https redirection")
	flag.Parse()

	if debug {
		log.Warn("Enabling debug logging")
		log.SetLevel(log.DebugLevel)
	}

	e := http.NewEcho()
	repo := repo.NewVersionRepository(data.NewPistonDataSource())

	routes.RegisterRoutes(e, repo)
	middleware.RegisterMiddleware(e, ratelimit, https)

	log.Infof("Running webserver on port %d🚀", port)
	log.Fatal(e.Start(":" + strconv.Itoa(port)))
}
