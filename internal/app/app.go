package app

import (
	_ "net/http/pprof"
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
	embeddingService := services.NewEmbeddingService()
	movieImportService := services.NewMovieImportService(tmdbService, embeddingService, store)

	go func() {
			log.Println("Starting pprof server on :6060")
			log.Println(http.ListenAndServe("localhost:6060", nil)) // Default pprof endpoints are served here
	}()

	// movieImportService.StartImport(20)	// Start Index
	movieImportService.StartMultithreadedImport(8, 1000, 540)	// Num Workers, Num Movies, Start Index

	return e
}
