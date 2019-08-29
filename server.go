package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/rocketblack/smartkidney-api/config"
	route "github.com/rocketblack/smartkidney-api/routes"
)

func main() {
	// Use labstack/echo for rich routing.
	// See https://echo.labstack.com/
	e := echo.New()
	s := config.Spec()

	// Middleware
	e.Logger.SetLevel(log.ERROR)
	e.Use(
		middleware.CORS(),
		middleware.Recover(),
		middleware.Logger(),
	)

	// Respond to API health checks.
	// Indicate the server is healthy.
	e.GET("/_ah/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "smart kidney api : ok!")
	})

	// Initialize routes
	route.Init(e)

	// Start server
	e.Logger.Fatal(e.Start(s.APIPort))
}
