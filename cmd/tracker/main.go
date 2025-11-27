package main

import (
	"fmt"
	"log"

	"github.com/mnadalc/traffic-tracker/internal/config"
	"github.com/mnadalc/traffic-tracker/internal/database"
)

func main() {
	// Test
  fmt.Println("Hello, World!")
	cfg, _ := config.LoadConfig()
	fmt.Println(cfg)

	// Database
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	err = db.Init()
	if err != nil {
		log.Fatalf("Failed to init tables: %v", err)
	}
}