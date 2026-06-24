package main

import (
	"log"

	"github.com/bv344/goalapp/internal/config"
	"github.com/bv344/goalapp/internal/routes"
)

func main() {
	cfg := config.Load()
	r := routes.Setup(cfg)
	log.Fatal(r.Run(":" + cfg.Port))
}
