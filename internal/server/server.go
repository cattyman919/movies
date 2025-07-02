package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"movies/internal/database"
	"movies/internal/tmdb"
)

type Server struct {
	port int

	db   database.Service
	tmdb *tmdb.Client
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	tmdbApiKey := os.Getenv("TMDB_API_KEY")
	NewServer := &Server{
		port: port,

		db:   database.New(),
		tmdb: tmdb.NewClient(tmdbApiKey),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
