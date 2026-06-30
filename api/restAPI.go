package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

// Starts listening for requests on the given port
func HandleRequests(port int) {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// GETs
	e.GET("/", healthcheck)

	if err := e.Start(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatal(err)
	}
}

// Routes

// GETs

func healthcheck(c *echo.Context) error {
	return c.String(http.StatusOK, "Health check successful.")
}
