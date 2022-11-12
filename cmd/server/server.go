package server

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/iugstav/colatech-api/internal/filesystem"
	"github.com/iugstav/colatech-api/internal/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitServer() error {
	// load environment variables from local .env
	envError := godotenv.Load(".env")
	if envError != nil {
		return envError
	}

	// assets watcher
	watcher := filesystem.WatchDirectories()
	defer watcher.Close()

	// connection to local database
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_CONNECTION_INFO"))
	if err != nil {
		return err
	}
	defer db.Close()

	// server
	r := gin.New()
	r.Use(middleware.DefaultStructuredLogger())
	r.Use(gin.Recovery())

	rg := r.Group("/")

	SetupRoutes(rg, db)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	return nil
}
