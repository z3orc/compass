package middleware

import (
	"time"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func RegisterMiddleware(e *echo.Echo, ratelimit int, useHttps bool) {
	e.Use(middleware.Secure())
	e.Use(loggerMiddleware())
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

	if useHttps {
		log.Info("Enabling http to https redirection. This only works if running Compass with https!")
		e.Use(middleware.HTTPSRedirect())
	}
}
