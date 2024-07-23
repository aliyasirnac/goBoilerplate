package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/aliyasirnac/goBackendBoilerplate/internal/config"
)

func Run(cfg *config.Config) {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong!")
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.App.Port)))
}
