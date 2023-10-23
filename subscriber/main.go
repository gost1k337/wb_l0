package main

import (
	"github.com/gost1k337/wb_demo/subscriber/config"
	"github.com/gost1k337/wb_demo/subscriber/internal/app"
	"log"
)

const configPath = "config/config.yml"

func main() {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
