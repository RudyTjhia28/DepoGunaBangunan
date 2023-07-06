package main

import (
	"depogunabangunan/apps/repository"
	"depogunabangunan/apps/services"
	"depogunabangunan/config"
	"depogunabangunan/endpoints"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load the configuration
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Establish a connection to the PostgreSQL database
	db, err := services.ConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	//create repository instance
	userRepository := repository.NewUserRepository(db)
	productRepository := repository.NewProductRepository(db)

	// Initialize the Gin router
	router := gin.Default()

	// Pass the database connection to the handlers
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Initialize the endpoints
	endpoints.InitializeEndpoints(router.Group(""), db, userRepository, productRepository)

	// Start the server
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
