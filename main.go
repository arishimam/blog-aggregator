package main

import "github.com/arishimam/blog-aggregator/internal/config"

func main() {
	cfg := config.Read()
	cfg.SetUser("arish")
	cfg = config.Read()
}
