package api

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var makingRequest sync.Mutex

// Starts listening for requests on the given port
func HandleRequests(port int) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// GETs
	e.GET("/", healthcheck)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

// Routes

// GETs

func healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "Health check successful.")
}
