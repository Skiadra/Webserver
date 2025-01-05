package main

import (
	"log"
	"time"
	"webserver/config"
	"webserver/database/migrations"
	"webserver/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// Load the database URI from the config package
	config := config.LoadConfig()
	dsn := config.DBUri // Ensure this is configured in config.go

	// Open a database connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	migrations.RunMigrations(db)

	// Initialize Gin
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://example.com"}, // Allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},                 // Allowed HTTP methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},      // Allowed headers
		ExposeHeaders:    []string{"Content-Length"},                               // Headers exposed to the browser
		AllowCredentials: true,                                                     // Allow cookies or Authorization headers
		MaxAge:           12 * time.Hour,                                           // Preflight request cache duration
	}))

	// Initialize routes
	routes.InitRoutes(router, db)

	// Run the server
	router.Run(config.ServerPort)
}
