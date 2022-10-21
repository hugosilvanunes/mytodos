package main

import (
	"log"

	"github.com/hugosilvanunes/mytodos/config"
	"github.com/hugosilvanunes/mytodos/server"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Run(*cfg); err != nil {
		log.Fatal(err)
	}
}
