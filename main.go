package main

import (
	"depogunabangunan/apps/services"
	"depogunabangunan/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := services.ConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	router := gin.Default()

	// Initialize the endpoints

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
