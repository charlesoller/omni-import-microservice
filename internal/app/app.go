package app

import (
	"log"
	"net/http"
	"os"

	"github.com/charlesoller/omni-import-microservice/internal/database"
	"github.com/charlesoller/omni-import-microservice/internal/db"
	"github.com/charlesoller/omni-import-microservice/internal/services"
	"github.com/labstack/echo/v4"
)

func Setup() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	tmdbAuth := os.Getenv("TMDB_ACCESS_TOKEN")
	if tmdbAuth == "" {
		log.Fatalln("Unable to load TMDB Access Token")
	}

	// Database
	data := database.SetupDatabase()
	queries := db.New(data)
	store := database.NewStore(data, queries)

	// Services
	tmdbService := services.NewTmdbService(tmdbAuth)
	movieImportService := services.NewMovieImportService(tmdbService, store)
	movieImportService.StartImport(0)

	return e
}