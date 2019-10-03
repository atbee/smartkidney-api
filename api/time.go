package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// GetTime to response local time.
func (db *MongoDB) GetTime(c echo.Context) (err error) {
	t := time.Now()

	return c.JSON(http.StatusOK, t)
}