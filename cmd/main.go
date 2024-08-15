package main

import (
	"flag"
	"strconv"
	"time"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/z3orc/compass/internal/data"
	"github.com/z3orc/compass/internal/handler"
	"github.com/z3orc/compass/internal/model"
	"github.com/z3orc/compass/internal/repo"
	"github.com/z3orc/compass/internal/util"
	"golang.org/x/time/rate"
)

var port int
var debug bool
var https bool
var ratelimit int

func main() {
	log.Info("Starting COMPASS🧭")

	flag.IntVar(&port, "port", 8000, "port for webserver to listen on")
	flag.IntVar(&ratelimit, "rate", 31, "max amount of requests per minute. 0 to disable")
	flag.BoolVar(&debug, "debug", false, "enable debug logging")
	flag.BoolVar(&https, "https", false, "enable https redirection")
	flag.Parse()

	if debug {
		log.Warn("Enabling debug logging")
		log.SetLevel(log.DebugLevel)
	}

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.Secure())
	e.Use(util.LoggerMiddleware())
	e.Use(middleware.Recover())

	if ratelimit != 0 {
		e.Use(middleware.RateLimiter(
			middleware.NewRateLimiterMemoryStoreWithConfig(
				middleware.RateLimiterMemoryStoreConfig{
					Rate:      rate.Every(time.Minute),
					Burst:     ratelimit,
					ExpiresIn: 3 * time.Minute},
			)))
	} else {
		log.Warn("Disabling ratelimiting")
	}

	if https {
		log.Info("Enabling http to https redirection. This only works if running Compass with https!")
		e.Use(middleware.HTTPSRedirect())
	}

	repo := repo.NewVersionRepository(data.NewPistonDataSource())

	//Users can use either piston or vanilla to get vanilla jars
	e.Pre(middleware.Rewrite(map[string]string{
		"/vanilla/*": "/piston/$1",
	}))

	e.GET("/", handler.HomeHandler)

	e.GET("/piston/:id", func(c echo.Context) error {
		return handler.VersionHandler(c, repo, model.FlavourPiston)
	})

	// e.GET("/paper/:id", func(c echo.Context) error {
	// 	return versionHandler(c, repo, model.FlavourPaper)
	// })

	// e.GET("/purpur/:id", func(c echo.Context) error {
	// 	return versionHandler(c, repo, model.FlavourPurpur)
	// })

	log.Infof("Running webserver on port %d🚀", port)
	log.Fatal(e.Start(":" + strconv.Itoa(port)))
}
