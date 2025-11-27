package main

import ("fmt"
	"log"
	"github.com/mnadalc/traffic-tracker/internal/config"
)

func main() {
  fmt.Println("Hello, World!")
	cfg, err := config.LoadConfig()
	
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	
	fmt.Println(cfg)
}