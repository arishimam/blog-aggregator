package main

import (
	"fmt"
	"github.com/arishimam/blog-aggregator/internal/config"
	"log"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("arish")
	if err != nil {
		log.Fatalf("Couldn't set current user: %v", err)

	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}
